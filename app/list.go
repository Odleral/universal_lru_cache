package app

import "time"

type List struct {
	cap         int
	len         int
	TTLDuration time.Duration
	Head        *Node
	Tail        *Node
}

type Node struct {
	Value any
	Key   string
	TTL   time.Time
	Next  *Node
	Prev  *Node
}

func NewList(duration time.Duration) *List {
	list := List{
		len:         0,
		Head:        &Node{},
		Tail:        &Node{},
		TTLDuration: duration,
	}

	list.Head.Next = list.Tail
	list.Tail.Prev = list.Head

	return &list
}

func (l *List) Append(key string, v any) *Node {
	node := Node{Key: key, Value: v, TTL: time.Now().Add(l.TTLDuration)}

	p := l.Tail.Prev
	p.Next = &node
	l.Tail.Prev = &node
	node.Prev = p
	node.Next = l.Tail

	l.len++

	return &node
}

func (l *List) Pop() *Node {
	return l.Remove(l.Head.Next)
}

func (l *List) Remove(node *Node) *Node {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	l.len -= 1

	return node
}
