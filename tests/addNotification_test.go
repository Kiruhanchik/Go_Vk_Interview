package tests

import (
	"example/services"
	"example/structures"
	"testing"
)

func TestAddNotification(t *testing.T) {
	Queue := structures.NotificationQueue{Notifications: []structures.Notification{}}

	Notification := structures.Notification{Text: "Check", Time: "12:32"}

	services.AddNotification(Notification, &Queue)

	if len(Queue.Notifications) == 0 {
		t.Fatal("Test TestAddNotification failed")
	}
}
