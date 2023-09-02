// Copyright 2023 Hany Mamdouh. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// All tests are excuted against dataset in provided file.
// Test is comapred for memory addresses.
package gotrees

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test_FindBFS(t *testing.T) {
	expect := "4"
	got := boss.FindBFS("Amr", func(n *Node[Person], s interface{}) bool {
		return n.Data.Name == s
	})

	if got.Id != expect {
		t.Errorf("Expected %s", expect)
	}
}

// Tests find node using DFS
func Test_FindDFS(t *testing.T) {
	expect := &developer2
	got := boss.FindDFS("Amr", func(n *Node[Person], s interface{}) bool {
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
	expect := []*Node[Person]{&developer1, &developer2, &developer3, &developer4}
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
	det := boss.FindFullDFS("Amr", func(n *Node[Person], t interface{}) bool {
		return n.Data.Name == t
	})

	if det.Node != expect {
		t.Errorf("Expected %v", expect)
	}
}

// Test scan specific level
func Test_Level(t *testing.T) {
	expect := []*Node[Person]{&developer1, &developer2, &developer3, &developer4}
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

	roots := Build[Person](rawData, compareFunc)
	if len(roots) != expectLen {
		t.Errorf("Expected length %v", expectLen)
	}
	if roots[0].Data.Name != expectName {
		t.Errorf("Expected Root with name %s", expectName)
	}
}

// Test convert Node to slice of type T
func Test_Slice(t *testing.T) {
	expect := 7

	if got := len(boss.Slice()); got != expect {
		t.Errorf("Expected %v", expect)
	}
}

// Testing printout of tree
// This test print out tree to stdout, it is not intended to be used as a varification in build line
func Test_PrintTree(t *testing.T) {
	PrintTree[Person](&boss, 0)
}

// Testing tree size
func Test_Size(t *testing.T) {
	expect := 7

	if got := boss.Size(); got != expect {
		t.Errorf("Expected %v\n", expect)
	}
}

// Testing LCA
func Test_LCA(t *testing.T) {
	expect := &teamleader1
	if gotlca := boss.LCA(&developer1, &developer2); gotlca != expect {
		t.Errorf("Expected %v", expect)
	}
}

// Testing serialization to JSON
func Test_SerializeJSON(t *testing.T) {

	if _, err := boss.SerializeJSON(); err != nil {
		t.Errorf("Failed to serialize")
	}
}

// Testing deserialize
func Test_DeserializeJSON(t *testing.T) {
	j, _ := boss.SerializeJSON()
	expect := 4
	got, err := DeserializeJSONToTree[Person](j)
	leaves := len(got.Leaves())
	if err != nil || leaves != expect {
		t.Errorf("Failed to serialize")
	}
}
