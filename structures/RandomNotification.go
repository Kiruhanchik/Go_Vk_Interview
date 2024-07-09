package structures

import (
	"errors"
	"math/rand"
)

func (nq *NotificationQueue) RandomNotification() (*Notification, error) {
	if nq.Head == nil {
		return nil, errors.New("Your queue is empty")
	}

	current := nq.Head

	randNum := rand.Intn(nq.Len())

	i := 0

	for {
		if i == randNum {
			break
		}

		i += 1
		current = current.Next
	}

	return &current.Notification, nil

}
