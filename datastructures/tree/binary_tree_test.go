package tree

import "testing"

func TestNewBinaryNode(t *testing.T) {
	if NewBinaryNode(nil) == nil {
		t.Fatal("expected a node")
	}
}

func TestNewBinaryTree(t *testing.T) {
	if NewBinaryTree(nil) == nil {
		t.Fatal("expected a tree")
	}
}

func TestBinaryTreeRoot(t *testing.T) {
	a := NewBinaryNode(3)
	b := NewBinaryNode(5)
	tr := NewBinaryTree(a)

	if NewBinaryTree(nil).Root() != nil {
		t.Fatal("unexpected value")
	}

	if tr.Root() != a {
		t.Fatal("unexpected value")
	}

	tr.SetRoot(b)

	if tr.Root() != b {
		t.Fatal("unexpected value")
	}
}

func TestBinaryNodeLeftRight(t *testing.T) {
	a := NewBinaryNode(3)
	b := NewBinaryNode(5)
	c := NewBinaryNode(1)

	if a.Left() != nil && a.Right() != nil {
		t.Fatal("unexpected value")
	}

	a.SetLeft(b)
	a.SetRight(c)

	if a.Left() != b && a.Right() != c {
		t.Fatal("unexpected value")
	}
}

func TestBinaryNodeValue(t *testing.T) {
	a := NewBinaryNode(3)

	if a.Value() != 3 {
		t.Fatal("unexpected value")
	}

	a.SetValue(4)

	if a.Value() != 4 {
		t.Fatal("unexpected value")
	}
}
