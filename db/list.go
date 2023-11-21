package db

type listNode struct {
	data []byte
	prev *listNode
	next *listNode
}

type List struct {
	len        int
	head, tail *listNode
}

func (l *List) LPush(data []byte) {
	n := &listNode{
		data: data,
	}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.len += 1
		return
	}

	if l.head == l.tail {
		head := l.head
		l.head = n
		l.tail = head
		l.head.next = l.tail
		l.tail.prev = l.head
		l.len += 1
		return
	}

	l.head.prev = n
	n.next = l.head
	l.head = n
	l.len += 1
}

func (l *List) RPush(data []byte) {
	n := &listNode{
		data: data,
	}

	if l.tail == nil {
		l.tail = n
		l.head = n
		l.len += 1
		return
	}

	if l.head == l.tail {
		tail := l.tail
		l.tail = n
		l.head = tail
		l.head.next = l.tail
		l.tail.prev = l.head
		l.len += 1
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.len += 1
}

func (l *List) RPop() []byte {
	tail := l.tail
	l.tail = tail.prev
	tail.prev = nil
	l.tail.next = nil
	l.len -= 1

	return tail.data
}

func (l *List) Len() int {
	return l.len
}

func (l *List) LPop() []byte {
	head := l.head
	l.head = head.next
	l.head.prev = nil
	head.next = nil
	l.len -= 1

	return head.data
}

func (l *List) Index(index int) []byte {
	if index >= l.len {
		return nil
	}
	n := l.head
	for i := 0; i <= index; i++ {
		n = n.next
	}

	if n != nil {
		return n.data
	}

	return nil
}

func (l *List) Range(start, end int) [][]byte {
	n := end - start
	b := make([][]byte, 0, n)
	
	h := l.head
	for i := 0; i <= start; i++ {
		h = h.next
	}

	for i := start; i < end; i++ {
		b = append(b, h.data)
		h = h.next
	}
	
	return b
}