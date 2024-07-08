package main

import (
	"example/services"
	"example/structures"
	"fmt"
	"log"
)

func main() {
	Queue := structures.NotificationQueue{Notifications: []structures.Notification{}}

	Notification := structures.Notification{
		Text: "Hi",
		Time: "12:30",
	}

	Notification2 := structures.Notification{
		Text: "Goodbye",
		Time: "12:34",
	}

	Notification3 := structures.Notification{
		Text: "Ok",
		Time: "13:00",
	}

	services.AddNotification(Notification, &Queue)

	services.AddNotification(Notification2, &Queue)

	services.AddNotification(Notification3, &Queue)

	fmt.Println(Queue.Notifications)

	Notification, err := services.GetRandomNotification(&Queue)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(Notification)
	}

	FirstNotification, err := services.PopNotification(&Queue)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(FirstNotification)
	}

	fmt.Println(Queue.Notifications)

}
