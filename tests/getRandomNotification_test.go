package tests

import (
	"practice/services"
	"practice/structures"
	"testing"
)

func TestGetRandomNotification(t *testing.T) {
	Queue := structures.NotificationQueue{
		Notifications: []structures.Notification{
			{Time: "23:23", Text: "First"},
			{Time: "23:24", Text: "Second"},
			{Time: "23:25", Text: "Third"},
		},
	}

	Notification, err := services.GetRandomNotification(&Queue)

	if Notification == (structures.Notification{}) || err != nil {
		t.Fatal("Test TestGetRandomNotification failed")
	}
}
