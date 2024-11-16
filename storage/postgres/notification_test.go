package postgres_test

import (
	"tender/model"
	"tender/storage/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNotification(t *testing.T) {
	db := postgres.ConnectDB(postgres.Logger)
	NotifRepo := postgres.NewNotificationRepository(db, postgres.Logger)

	resp, err := NotifRepo.CreateNotification(model.Notification{
		UserID:     "f76b4c4b-da00-4957-bb88-b197e0ce739a",
		Message:    "message",
		RelationID: "97c71420-b902-434d-9ba0-e2d718451ead",
		Type:       "",
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Notification created successfully", resp.Message)
}

func TestUpdateNotification(t *testing.T) {
	db := postgres.ConnectDB(postgres.Logger)
	NotifRepo := postgres.NewNotificationRepository(db, postgres.Logger)

	resp, err := NotifRepo.UpdateNotification("ff221779-f22f-496d-aba2-61d27acd14ce")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Notification updated successfully", resp.Message)
}

func TestGetAllNotifications(t *testing.T) {
	db := postgres.ConnectDB(postgres.Logger)

	NotifRepo := postgres.NewNotificationRepository(db, postgres.Logger)

	resp, err := NotifRepo.GetAllNotifications(model.NotifFilter{
		UserID: "f76b4c4b-da00-4957-bb88-b197e0ce739a",
		IsRead: "true",
		Limit:  10,
		Offset: 0,
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, resp)
	assert.NotEqual(t, len(resp.Notifications), 0)
}
