package linked_list

// Node represents a node in a singly linked list
type Node struct {
	Value int
	Next  *Node
}

// SinglyLinkedList represents a singly linked list
type SinglyLinkedList struct {
	Head *Node
	Size int
}

// New creates a new empty singly linked list
func New() *SinglyLinkedList {
	return &SinglyLinkedList{
		Head: nil,
		Size: 0,
	}
}

// Append adds a new node with the given value to the end of the list
func (l *SinglyLinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: nil}

	if l.IsEmpty() {
		l.Head = newNode
		l.Size++
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	l.Size++
}

// Prepend adds a new node with the given value to the beginning of the list
func (l *SinglyLinkedList) Prepend(value int) {
	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
	l.Size++
}

// InsertAt inserts a value at the specified index
func (l *SinglyLinkedList) InsertAt(index, value int) bool {
	if index < 0 || index > l.Size {
		return false
	}

	if index == 0 {
		l.Prepend(value)
		return true
	}

	newNode := &Node{Value: value}
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	l.Size++
	return true
}

// Delete removes the first node with the given value
func (l *SinglyLinkedList) Delete(value int) bool {
	if l.IsEmpty() {
		return false
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Size--
		return true
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			l.Size--
			return true
		}
		current = current.Next
	}

	return false
}

// DeleteAt removes the node at the specified index
func (l *SinglyLinkedList) DeleteAt(index int) bool {
	if index < 0 || index >= l.Size || l.IsEmpty() {
		return false
	}

	if index == 0 {
		l.Head = l.Head.Next
		l.Size--
		return true
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	l.Size--
	return true
}

// Find searches for a node with the given value
func (l *SinglyLinkedList) Find(value int) *Node {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// Get returns the value at the specified index
func (l *SinglyLinkedList) Get(index int) (int, bool) {
	if index < 0 || index >= l.Size || l.IsEmpty() {
		return 0, false
	}

	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, true
}

// Reverse reverses the linked list in place
func (l *SinglyLinkedList) Reverse() {
	var prev *Node
	current := l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}

// ToSlice returns a slice representation of the list
func (l *SinglyLinkedList) ToSlice() []int {
	result := make([]int, 0, l.Size)
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// IsEmpty checks if the list is empty
func (l *SinglyLinkedList) IsEmpty() bool {
	return l.Head == nil
}

// Clear removes all nodes from the list
func (l *SinglyLinkedList) Clear() {
	l.Head = nil
	l.Size = 0
}
