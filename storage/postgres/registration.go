package postgres

import "database/sql"

type RegistrationRepository interface {

}

type registrationRepository struct {
	db *sql.DB
}


func NewRegistrationRepository(db *sql.DB) RegistrationRepository {
	return registrationRepository{
		db: db,
	}
}