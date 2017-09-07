package tree

type BinaryTree interface {
	Root() BinaryNode
	SetRoot(BinaryNode)
}

type BinaryNode interface {
	Value() interface{}
	SetValue(interface{})
	Left() BinaryNode
	SetLeft(BinaryNode)
	Right() BinaryNode
	SetRight(BinaryNode)
}

type binaryTree struct {
	root BinaryNode
}

type binaryNode struct {
	value interface{}
	left  BinaryNode
	right BinaryNode
}

func NewBinaryNode(v interface{}) BinaryNode {
	return &binaryNode{value: v}
}

func NewBinaryTree(root BinaryNode) BinaryTree {
	return &binaryTree{root: root}
}

func (n *binaryNode) Left() BinaryNode {
	return n.left
}

func (n *binaryNode) SetLeft(v BinaryNode) {
	n.left = v
}

func (n *binaryNode) Right() BinaryNode {
	return n.right
}

func (n *binaryNode) SetRight(v BinaryNode) {
	n.right = v
}

func (n *binaryNode) Value() interface{} {
	return n.value
}

func (n *binaryNode) SetValue(v interface{}) {
	n.value = v
}

func (t *binaryTree) Root() BinaryNode {
	return t.root
}

func (t *binaryTree) SetRoot(root BinaryNode) {
	t.root = root
}
