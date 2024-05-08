package handlers

import (
	. "main/backend/models"
)

func GetNotifications(Id int64) ([]Notification, error) {
	var notifications []Notification
	err := Db.Where("Uid = ?", Id).Find(&notifications)
	if err != nil {
		return []Notification{}, err
	}
	return notifications, nil
}
