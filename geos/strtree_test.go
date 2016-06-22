package geos

import (
	"testing"
)

func TestNewSTRTree(t *testing.T) {
	tree := NewSTRTree(10)

	if tree == nil {
		t.Error("STRTree should not be nil")
	}
}

type itemTest string

func (i itemTest) Parse() []byte {
	return []byte(i)
}

func TestStrTree_Insert(t *testing.T) {
	tree := NewSTRTree(10)

	item := itemTest("test1")
	g, _ := FromWKT("POINT(-46 -23)")
	tree.Insert(g, item)

	items := 0

	cb := func(bitem []byte) {
		items++

		if items != 1 {
			t.Error("Should have only one item")
		}

		if itemTest(bitem) != item {
			t.Errorf("Expect item: %s, got %s", item, bitem)
		}
	}

	tree.Iterate(cb)
}

// TODO Improve tests