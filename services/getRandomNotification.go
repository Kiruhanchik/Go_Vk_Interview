package services

import (
	"errors"
	"example/structures"
	"math/rand"
)

func GetRandomNotification(queue *structures.NotificationQueue) (structures.Notification, error) {

	if len(queue.Notifications) == 0 {
		return structures.Notification{}, errors.New("Queue is empty")
	}

	ind := rand.Intn(len(queue.Notifications))

	return queue.Notifications[ind], nil
}
