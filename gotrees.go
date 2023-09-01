package gotrees

import (
	"fmt"
	"strings"
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
	fmt.Printf("%v%s%v\n", level, strings.Repeat("  ", level), node.Data)
	for _, child := range node.Children {
		PrintTree(child, level+1)
	}
}
