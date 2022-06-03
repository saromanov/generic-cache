package cache

type Node[V any] struct {
	Value      V
	Prev, Next *Node[V]
}

type List[V any] struct {
	Top, Back *Node[V]
}

func NewList[V any]()*List[V]{
	return &List[V]{}
}

func (l *List[V]) PushBack(v V) {
	l.PushBackNode(&Node[V]{
		Value: v,
	})
}

func (l *List[V]) PushFront(v V) {
	l.PushTopNode(&Node[V]{
		Value: v,
	})
}

func (l *List[V]) PushBackNode(n *Node[V]) {
	n.Next = nil
	n.Prev = l.Back
	if l.Back != nil {
		l.Back.Next = n
	} else {
		l.Top = n
	}
	l.Back = n
}

func (l *List[V]) PushTopNode(n *Node[V]) {
	n.Next = l.Top
	n.Prev = nil
	if l.Top != nil {
		l.Top.Prev = n
	} else {
		l.Back = n
	}
	l.Top = n
}

func (l *List[V]) Remove(n *Node[V]) {
	if n.Next != nil {
		n.Next.Prev = n.Prev
	} else {
		l.Back = n.Prev
	}
	if n.Prev != nil {
		n.Prev.Next = n.Next
	} else {
		l.Top = n.Next
	}
}

func (n *Node[V]) Each(fn func(val V)) {
	node := n
	for node != nil {
		fn(node.Value)
		node = node.Next
	}
}