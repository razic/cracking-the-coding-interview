package stack

// Stack represents a stack data structure
type Stack struct {
	top *Node
}

// Node represents a node in the stack
type Node struct {
	Value interface{}
	next  *Node
}

// Pop will pop a node off the top of the stack
func (s *Stack) Pop() interface{} {
	if &s.top == nil {
		return nil
	}

	node := s.top
	s.top = node.next

	return node.Value
}

// Push will push a node onto the stack
func (s *Stack) Push(val interface{}) {
	s.top = &Node{
		Value: val,
		next:  s.top,
	}
}

// Peek returns the node on the top of the stack
func (s *Stack) Peek() interface{} {
	return s.top.Value
}

// IsEmpty returns true or false if the stack is empty
func (s *Stack) IsEmpty() bool {
	return &s.top == nil
}
