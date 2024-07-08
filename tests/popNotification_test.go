package tests

import (
	"example/services"
	"example/structures"
	"testing"
)

func TestPopNotification(t *testing.T) {
	Queue := structures.NotificationQueue{
		Notifications: []structures.Notification{
			{Time: "23:23", Text: "First"},
			{Time: "23:24", Text: "Second"},
			{Time: "23:25", Text: "Third"},
		},
	}

	initialLength := len(Queue.Notifications)

	Notification, err := services.PopNotification(&Queue)

	if Notification == (structures.Notification{}) || len(Queue.Notifications) != initialLength-1 || err != nil {
		t.Fatal("Test TestPopNotification failed")
	}
}
