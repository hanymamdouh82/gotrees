package gotrees

import (
	"testing"

	gotrees "github.com/hanymamdouh82/gotrees"
	"golang.org/x/exp/slices"
)

func Test_FindBFS(t *testing.T) {
	expect := "4"
	got := boss.FindBFS("Amr", func(n *gotrees.Node[Person], s interface{}) bool {
		return n.Data.Name == s
	})

	if got.Id != expect {
		t.Errorf("Expected %s", expect)
	}
}

// Tests find node using DFS
// Test is excuted against dataset in provided file
// Test is comapred for memory addresses
func Test_FindDFS(t *testing.T) {
	expect := &developer2
	got := boss.FindDFS("Amr", func(n *gotrees.Node[Person], s interface{}) bool {
		return n.Data.Name == s
	})

	if got != expect {
		t.Errorf("Expected %v", expect)
	}
}

// Tests find node using Id
// Test is excuted against dataset in provided file
// Test is comapred for memory addresses
func Test_FindId(t *testing.T) {
	expect := &developer2
	got := boss.FindId("4")

	if got != expect {
		t.Errorf("Expected %v", expect)
	}
}

// Test leaves starting from root
func Test_Leaves(t *testing.T) {
	expect := []*gotrees.Node[Person]{&developer1, &developer2, &developer3, &developer4}
	got := boss.Leaves()

	for _, v := range got {
		if idx := slices.Index(expect, v); idx == -1 {
			t.Errorf("Wrong memory addresses")
		}
	}
}

// Testing maximum depth
func Test_Depth(t *testing.T) {
	expect := 3
	if got := boss.Depth(); got != expect {
		t.Errorf("Expected %v", expect)
	}
}

// Test find node and return full details
// Test uses compare function
func Test_FindFullDFS(t *testing.T) {
	expect := &developer2
	det := boss.FindFullDFS("Amr", func(n *gotrees.Node[Person], t interface{}) bool {
		return n.Data.Name == t
	})

	if det.Node != expect {
		t.Errorf("Expected %v", expect)
	}
}

// Test scan specific level
func Test_Level(t *testing.T) {
	expect := []*gotrees.Node[Person]{&developer1, &developer2, &developer3, &developer4}
	got := boss.Level(3)

	for _, v := range got {
		if idx := slices.Index(expect, v); idx == -1 {
			t.Errorf("Wrong memory addresses")
		}
	}
}

// Test build tree from array
func Test_Build(t *testing.T) {
	expectLen := 1
	expectName := "Hany"

	compareFunc := func(p, c Person) bool {
		return p.Name == c.Boss
	}

	roots := gotrees.Build[Person](rawData, compareFunc)
	if len(roots) != expectLen {
		t.Errorf("Expected length %v", expectLen)
	}
	if roots[0].Data.Name != expectName {
		t.Errorf("Expected Root with name %s", expectName)
	}
}
