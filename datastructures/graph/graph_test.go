package graph

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()

	if len(g.Nodes()) != 0 {
		t.Fatalf("unexpected node")
	}
}

func TestNewNode(t *testing.T) {
	x := &struct{}{}
	g := NewNode(x)

	if g.Value() != x {
		t.Fatalf("unexpected node")
	}
}

func TestGraphAppend(t *testing.T) {
	g := NewGraph()
	a, b := NewNode(1), NewNode(2)

	g.Append(a, b)

	if len(g.Nodes()) != 2 {
		t.Fatalf("unexpected node")
	}

	occs := map[Node]int{}
	for _, v := range g.Nodes() {
		occs[v]++
	}

	if occs[a] != 1 && occs[b] != 1 {
		t.Fatalf("unexpected count")
	}
}

func TestNodeAppend(t *testing.T) {
	a, b, c := NewNode(1), NewNode(2), NewNode(3)

	a.Append(b, c)

	if len(a.Adjacent()) != 2 {
		t.Fatalf("unexpected len")
	}

	occs := map[Node]int{}
	for _, v := range a.Adjacent() {
		occs[v]++
	}

	if occs[b] != 1 && occs[c] != 1 {
		t.Fatalf("unexpected count")
	}
}
