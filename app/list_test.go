package app

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	l := NewList(3 * time.Second)

	t.Run("Check List len", func(t *testing.T) {
		assert.Equal(t, 0, l.len, "len is 0")

		assert.Equal(t, l.Tail, l.Head.Next, "Head equal to Tail")

	})

	t.Run("List add node", func(t *testing.T) {
		node1 := l.Append("e1", 5)

		assert.Greater(t, node1.TTL.Unix(), time.Now().Unix(), "Time.Now().Unix() greater than node.TTL.Unix()")
		assert.Equal(t, 5, l.Head.Next.Value.(int), "First node value is 5")

		node2 := l.Append("e2", 10)

		assert.Equal(t, 10, node1.Next.Value.(int), "Second node value is 10")
		assert.Equal(t, node2, node1.Next, "node1 next is node2")

		node3 := l.Append("e3", 15)
		assert.Equal(t, "e3", node3.Key, "Node3 key is equal to e3")
	})

	t.Run("List pop node", func(t *testing.T) {
		node := l.Pop()
		assert.Equal(t, 2, l.len, "list len is 2")
		assert.Equal(t, node.Next, l.Head.Next, "Head.Next equal poped node.Next")
	})

	t.Run("List iteration test", func(t *testing.T) {
		current := l.Head.Next
		i := 10
		for current != l.Tail {
			assert.Equal(t, i, current.Value.(int), "i equal to current.Value.(int)")
			current = current.Next
			i += 5
		}
	})
}

func InsertXList(x int, b *testing.B) {
	list := NewList(2 * time.Second)
	b.ResetTimer()
	for i := 0; i < x; i++ {
		list.Append(strconv.Itoa(i), i)
	}
}

func BenchmarkList_Append1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXList(1000, b)
	}
}

func BenchmarkList_Append10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXList(10000, b)
	}
}

func BenchmarkList_Append100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXList(100000, b)
	}
}

func BenchmarkList_Append1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertXList(1000000, b)
	}
}
