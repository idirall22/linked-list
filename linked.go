package linked

import "fmt"

type node struct {
	Value int
	Next  *node
}

type LinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func (l *LinkedList) Insert(value int) {
	n := &node{Value: value, Next: nil}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	}
	l.Tail.Next = n
	l.Tail = n
	l.Length++
}

func (l *LinkedList) transverse(index int) *node {
	n := l.Head
	for i := 0; i <= index; i++ {
		if index == i {
			return n
		}
		n = n.Next
	}
	return nil
}

func (l *LinkedList) Remove(index int) {
	if index > l.Length-1 {
		return
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Length--
	} else if index == l.Length-1 {
		prevNode := l.transverse(index - 1)
		if prevNode == nil {
			return
		}
		prevNode.Next = nil
		l.Tail = prevNode
		l.Length--
	} else {
		prevNode := l.transverse(index - 1)
		prevNode.Next = prevNode.Next.Next
		l.Length--
	}
}

func (l *LinkedList) Lookup(index int) int {
	n := l.transverse(index)
	return n.Value
}

func (l *LinkedList) Update(index, value int) {
	n := l.transverse(index)
	n.Value = value
}

func (l *LinkedList) Print() {
	n := l.Head
	for {
		if n == nil {
			break
		}
		if n != l.Tail {
			fmt.Print(n.Value, " --> ")
		} else {
			fmt.Print(n.Value)
		}
		n = n.Next
	}
	fmt.Println()
	fmt.Println("Length: ", l.Length)
}
