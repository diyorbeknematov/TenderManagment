package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"tender/model"
)

type NotificationRepository interface {
	CreateNotification(notif model.Notification) (*model.NotificationResp, error)
	UpdateNotification(id string) (*model.UpdateNotificationResp, error)
	GetUnreadNotifications(userID string) ([]model.Notification, error)
}

type notificationRepositoryImpl struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewNotificationRepository(db *sql.DB, logger *slog.Logger) NotificationRepository {
	return &notificationRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (repo notificationRepositoryImpl) CreateNotification(notif model.Notification) (*model.NotificationResp, error) {
	var notifResp model.NotificationResp

	err := repo.db.QueryRow(`
		INSERT INTO notifications (
			user_id,
			message,
			relation_id,
			type
		)
		VALUES (
			$1, $2, $3, $4
		)
		RETURNING 
			id,
			created_at
	`, notif.UserID, notif.Message, notif.RelationID, notif.Type).Scan(
		&notifResp.ID,
		&notifResp.CreatedAt,
	)

	if err != nil {
		repo.logger.Error(fmt.Sprintf("NOtification created qilinganda xatolik bor: %v", err))
		return nil, err
	}

	notifResp.Message = "Notification created successfully"

	return &notifResp, nil
}

func (repo notificationRepositoryImpl) UpdateNotification(id string) (*model.UpdateNotificationResp, error) {
	var notif model.UpdateNotificationResp

	err := repo.db.QueryRow(`
		UPDATE notifications
		SET
			is_read = true
		WHERE id = $1
		RETURNING 
			id,
			updated_at
	`, id).Scan(&notif.ID, &notif.UpdatedAt)

	if err != nil {
		repo.logger.Error(fmt.Sprintf("Notificationni yangilashda xatolik bor: %v", err))
		return nil, err
	}

	notif.Message = "Notification updated successfully"

	return &notif, err
}

var Query = `
	SELECT 
		id,
		user_id,
		message,
		relation_id,
		type,
		is_read
	FROM notifications
	WHERE user_id = $1
`
var CountQuery = `
	SELECT
		COUNT(*)
	FROM notifications
	WHERE user_id = $1
`

func (repo notificationRepositoryImpl) GetUnreadNotifications(userID string) ([]model.Notification, error) {
	var notifications []model.Notification

	query := `
        SELECT 
            id, 
            user_id, 
            message, 
            relation_id, 
            type, 
            is_read
        FROM 
            notifications 
        WHERE 
            user_id = $1 
            AND is_read = FALSE;
    `

	rows, err := repo.db.Query(query, userID)
	if err != nil {
		repo.logger.Error(fmt.Sprintf("Error fetching unread notifications: %v", err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var notification model.Notification
		err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.Message,
			&notification.RelationID,
			&notification.Type,
			&notification.IsRead,
		)
		if err != nil {
			repo.logger.Error(fmt.Sprintf("Error scanning row: %v", err))
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}
