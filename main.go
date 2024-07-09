package main

import (
	"example/structures"
	"fmt"
	"log"
)

func main() {
	Queue := &structures.NotificationQueue{}

	Notification := structures.Notification{
		Text: "First",
		Time: "12:30",
	}

	Notification2 := structures.Notification{
		Text: "Second",
		Time: "12:33",
	}

	Notification3 := structures.Notification{
		Text: "Third",
		Time: "12:36",
	}

	Queue.AddNotification(Notification)

	Queue.AddNotification(Notification2)

	Queue.AddNotification(Notification3)

	n, err := Queue.RandomNotification()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Random: ", *n)

	n, err = Queue.PopNotification()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("First notification: ", *n)

	current := Queue.Head

	for current != nil {
		fmt.Println(current.Notification)
		current = current.Next
	}

}
