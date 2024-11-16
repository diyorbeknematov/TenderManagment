package model

type Notification struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Message    string `json:"message"`
	RelationID string `json:"relation_id"`
	Type       string `json:"type"`
	IsRead     bool   `json:"is_read"`
}

type NotificationResp struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type NotifFilter struct {
	UserID   string `json:"user_id"`
	IsRead   string `json:"is_read"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type AllNotifications struct {
	Notifications []Notification `json:"notifications"`
	TotalCount    int            `json:"total_count"`
	Limit         int            `json:"limit"`
	Page          int            `json:"page"`
}

type UpdateNotification struct {
	ID string `json:"id"`
}

type UpdateNotificationResp struct {
	ID        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
	Message   string `json:"message"`
}
