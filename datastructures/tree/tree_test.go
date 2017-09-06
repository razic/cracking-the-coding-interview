package tree

import "testing"

func TestNewNode(t *testing.T) {
	if NewNode(1) == nil {
		t.Fatal("expected a node")
	}
}

func TestNewTree(t *testing.T) {
	if NewTree(nil) == nil {
		t.Fatal("expected a tree")
	}
}

func TestTreeRoot(t *testing.T) {
	if NewTree(1).Root() == nil {
		t.Fatal("unexpected value")
	}
}

func TestNodeChildren(t *testing.T) {
	a := NewNode(3)
	b := NewNode(5)
	c := NewNode(1)

	if len(a.Children()) != 0 {
		t.Fatal("unexpected number of child nodes")
	}

	a.Append(b, c)

	if len(a.Children()) != 2 {
		t.Fatal("unexpected number of child nodes")
	}

	occs := map[Node]int{}
	for _, v := range a.Children() {
		occs[v]++
	}

	if occs[b] != 1 {
		t.Fatal("unexpected child node")
	}

	if occs[b] != 1 {
		t.Fatal("unexpected child node")
	}
}

func TestNodeValue(t *testing.T) {
	a := NewNode(3)

	if a.Value() != 3 {
		t.Fatal("unexpected value")
	}

	a.SetValue(4)

	if a.Value() != 4 {
		t.Fatal("unexpected value")
	}
}
