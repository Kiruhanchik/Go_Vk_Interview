package structures

import "errors"

func (nq *NotificationQueue) PopNotification() (*Notification, error) {
	if nq.Head == nil {
		return nil, errors.New("Your queue is empty")
	}

	dequened := nq.Head

	nq.Head = nq.Head.Next

	dequened.Next = nil

	if nq.Head == nil {
		nq.Tail = nil
	}

	return &dequened.Notification, nil
}
