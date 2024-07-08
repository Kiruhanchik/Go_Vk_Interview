package services

import (
	"errors"
	"example/structures"
)

func PopNotification(queue *structures.NotificationQueue) (structures.Notification, error) {
	if len(queue.Notifications) == 0 {
		return structures.Notification{}, errors.New("Queue is empty")
	}

	first := queue.Notifications[0]

	queue.Notifications = queue.Notifications[1:]

	return first, nil
}
