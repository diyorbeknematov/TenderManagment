package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"tender/model"
)

type RegistrationRepository interface {
	CreateUser(user model.UserRegisterReq) (*model.UserRegisterResp, error)
	GetUserByUsername(email  string) (*model.GetUser, error)
	IsUserExists(email, username string) (bool, error)
}

type registrationRepositoryImpl struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewRegistrationRepository(db *sql.DB, logger *slog.Logger) RegistrationRepository {
	return registrationRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (repo registrationRepositoryImpl) CreateUser(user model.UserRegisterReq) (*model.UserRegisterResp, error) {
	var userResp model.UserRegisterResp

	err := repo.db.QueryRow(`
		INSERT INTO users (
			username,
			email,
			role,
			password
		) 
		VALUES (
			$1, $2, $3, $4
		)
		RETURNING
			id,
			created_at
	`, user.Username, user.Email, user.Role, user.Password).
		Scan(
			&userResp.ID,
			&userResp.CreatedAt,
		)

	if err != nil {
		repo.logger.Error(fmt.Sprintf("Register qilishda xatolik bor: %v", err))
		return nil, err
	}

	userResp.Message = "User registered successfully"

	return &userResp, err
}

func (repo registrationRepositoryImpl) GetUserByUsername(email string) (*model.GetUser, error) {
	var user model.GetUser

	err := repo.db.QueryRow(`
		SELECT 
			id,
			username,
			role,
			password
		FROM users
		WHERE
			deleted_at IS NULL AND username = $1;
	`, email).Scan(
		&user.ID,
		&user.Username,
		&user.Role,
		&user.Password,
	)

	if err != nil {
		repo.logger.Error(fmt.Sprintf("Userni email bo'yicha olishda xatolik: %v", err))
		return nil, err
	}

	return &user, err
}

func (repo registrationRepositoryImpl) IsUserExists(email, username string) (bool, error) {
	var exists bool

	err := repo.db.QueryRow(`
		SELECT EXISTS(
            SELECT 1 FROM users 
            WHERE email = $1 OR username = $2
        )
	`, email, username).Scan(&exists)

	if err != nil {
		repo.logger.Error("Username or Email already exists")
		return false, err
	}

	return exists, nil
}
