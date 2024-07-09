package structures

func (nq *NotificationQueue) AddNotification(notification Notification) {
	newNode := &Node{Notification: notification}

	if nq.Head == nil {
		nq.Head = newNode
		nq.Tail = newNode
	} else {
		current := nq.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		nq.Tail = newNode
	}

}
