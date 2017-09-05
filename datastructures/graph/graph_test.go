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
	a, b := &node{}, &node{}

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
