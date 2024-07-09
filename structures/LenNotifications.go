package structures

func (nq *NotificationQueue) Len() int {
	count := 0

	current := nq.Head

	for current != nil {
		count += 1
		current = current.Next
	}

	return count
}
