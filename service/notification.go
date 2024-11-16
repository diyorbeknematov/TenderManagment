package service

import (
	"fmt"
	"tender/model"
)

func (s *Service) CreateNotification(notif model.Notification) (*model.NotificationResp, error) {
	resp, err := s.Storage.NotificationRepository().CreateNotification(notif)

	if err != nil {
		s.Log.Error(fmt.Sprintf("Notoficationni yaratishda xatolik bor: %v", err))
		return nil, err
	}

	return resp, nil
}

func (s *Service) UpdateNotification(notif model.UpdateNotification) (*model.UpdateNotificationResp, error) {
	resp, err := s.Storage.NotificationRepository().UpdateNotification(notif.ID)

	if err != nil {
		s.Log.Error(fmt.Sprintf("Notificationni update qilishda xatolik bor: %v", err))
		return nil, err
	}

	return resp, nil
}

func (s *Service) GetAllNotifications(filter model.NotifFilter) (*model.AllNotifications, error) {
	resp, err := s.Storage.NotificationRepository().GetAllNotifications(filter)

	if err != nil {
		s.Log.Error(fmt.Sprintf("Barcha notificationlarni olishda xatolik bor: %v", err))
		return nil, err
	}

	return resp, nil
}
