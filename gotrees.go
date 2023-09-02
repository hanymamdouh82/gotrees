// Copyright 2023 Hany Mamdouh. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package gotrees

import (
	"encoding/json"
	"fmt"
)

// Build a tree out from slice of objects using comparison function to determine parent/child relationship.
// Implement your own logic in compareFunc to specify parent/child relationship
func Build[T any](values []T, compareFunc CompareFunc[T]) []*Node[T] {
	nodes := make([]*Node[T], 0)
	roots := findRoots[T](values, compareFunc)

	for _, root := range roots {
		node := buildForRoot[T](root, values, compareFunc)
		nodes = append(nodes, node)
	}

	return nodes
}

// Utility function to print the tree structure.
// It prints the depth and node values in hirarichal view to stdout.
// Use for debugging only
func PrintTree[T any](node *Node[T], level int) {
	if node == nil {
		return
	}
	sep := ""
	for i := 0; i < level+1; i++ {
		sep = sep + "  "
	}
	fmt.Printf("%v%s%v\n", level, sep, node.Data)
	for _, child := range node.Children {
		PrintTree(child, level+1)
	}
}

// DeserializeJSONToTree deserializes a JSON representation into a tree structure.
// JSON must contain `Children` key, that will be used to identify children and relationship
func DeserializeJSONToTree[T any](jsonData string) (*Node[T], error) {
	// Unmarshal the JSON data into a map.
	var nodeMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &nodeMap)
	if err != nil {
		return nil, err
	}

	// Reconstruct the tree from the map.
	return deserializeJSON[T](nodeMap), nil
}
