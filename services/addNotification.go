package services

import "example/structures"

func AddNotification(Notification structures.Notification, queue *structures.NotificationQueue) {
	queue.Notifications = append(queue.Notifications, Notification)
}
