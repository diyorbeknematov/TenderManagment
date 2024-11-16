package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"tender/model"
)

type NotificationRepository interface {
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
		FROM notifications 
		VALUES (
			$1, $2, $3, $4
		)
		RETURNING 
			id,
			created_at
	`).Scan(
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

func (repo notificationRepositoryImpl) UpdateNotification(id string) error {
	_, err := repo.db.Exec(`
		UPDATE notifications
		SET
			is_read = true
		WHERE id = $1
	`, id)

	if err != nil {
		repo.logger.Error(fmt.Sprintf("Notificationni yangilashda xatolik bor: %v", err))
		return err
	}

	return err
}

var Query = `
	SELECT 
		id,
		user_id,
		message,
		relation_id,
		type,
		isRead
	FROM notifications
	WHERE user_id = $1
`
var CountQuery = `
	SELECT
		COUNT(*)
	FROM notifications
	WHERE user_id = $1
`

func (repo notificationRepositoryImpl) GetAllNotifications(notFilter model.NotifFilter) (*model.AllNotifications, error) {
	var (
		args          []interface{}
		filter        string
		notifications []model.Notification
		totalCount    int
	)

	args = append(args, notFilter.UserID)

	if notFilter.IsRead != "" {
		filter += fmt.Sprintf(" AND is_read = $%d", len(args)+1)
		args = append(args, notFilter.IsRead)
	}

	if notFilter.DateFrom != "" {
		filter += fmt.Sprintf(" AND created_at > $%d", len(args)+1)
		args = append(args, notFilter.DateFrom)
	}

	if notFilter.DateTo != "" {
		filter += fmt.Sprintf(" AND created_at < $%d", len(args)+1)
		args = append(args, notFilter.DateTo)
	}

	err := repo.db.QueryRow(CountQuery+filter, args...).Scan(&totalCount)
	if err != nil {
		repo.logger.Error(fmt.Sprintf("Notiticationslar sonini get qilishda xatolik: %v", err))
		return nil, err
	}

	filter += fmt.Sprintf(" OFFSET %d LIMIT %d", notFilter.Offset, notFilter.Limit)

	rows, err := repo.db.Query(Query+filter, args...)
	if err != nil {
		repo.logger.Error(fmt.Sprintf("Notificationlarni olishda xatolik bor: %v", err))
		return nil, err
	}

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
			repo.logger.Error(fmt.Sprintf("rowsdan o'qishda xatoik bor: %v", err))
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return &model.AllNotifications{
		Notifications: notifications,
		TotalCount:    totalCount,
		Limit:         notFilter.Limit,
		Page:          notFilter.Offset,
	}, nil
}
