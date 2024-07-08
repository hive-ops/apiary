package utils

type Node struct {
	Key   string
	Value []byte
	Prev  *Node
	Next  *Node
}

type DoublyLinkedList struct {
	Head    *Node
	Tail    *Node
	Size    int
	MaxSize int
}

func NewDoublyLinkedList(maxSize int) *DoublyLinkedList {
	return &DoublyLinkedList{
		MaxSize: maxSize,
	}
}

func (dll *DoublyLinkedList) PushFront(key string, value []byte) *Node {
	newNode := &Node{
		Key:   key,
		Value: value,
	}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
	dll.Size++
	return newNode
}

func (dll *DoublyLinkedList) MoveToFront(node *Node) {
	if node == dll.Head || dll.Head == nil || dll.Tail == nil {
		return
	}
	if node == dll.Tail {
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node
}

func (dll *DoublyLinkedList) RemoveLast() *Node {
	tail := dll.Tail
	if dll.Tail != nil {
		if dll.Head == dll.Tail {
			dll.Head = nil
			dll.Tail = nil
		} else {
			dll.Tail = dll.Tail.Prev
			dll.Tail.Next = nil
		}
		dll.Size--
	}
	return tail
}

func (dll *DoublyLinkedList) Remove(node *Node) {
	if node == dll.Head {
		dll.Head = node.Next
	}
	if node == dll.Tail {
		dll.Tail = node.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	dll.Size--
}

func (dll *DoublyLinkedList) Clear() {
	dll.Head = nil
	dll.Tail = nil
	dll.Size = 0
}

type HashMap struct {
	nodes map[string]*Node
}

func NewHashMap() *HashMap {
	return &HashMap{
		nodes: make(map[string]*Node),
	}
}

func (hm *HashMap) Get(key string) (*Node, bool) {
	node, ok := hm.nodes[key]
	return node, ok
}

func (hm *HashMap) Set(key string, node *Node) {
	hm.nodes[key] = node
}

func (hm *HashMap) Delete(key string) {
	delete(hm.nodes, key)
}

func (hm *HashMap) Clear() {
	hm.nodes = make(map[string]*Node)
}
