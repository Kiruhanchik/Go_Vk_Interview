package structures

type NotificationQueue struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Notification Notification
	Next         *Node
}
