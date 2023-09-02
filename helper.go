// Copyright 2023 Hany Mamdouh. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package gotrees

// helper used in recursive search for finding a leaves using DFS algorithm.
// leaves are evaluated starting from input `root` as the root node.
func findLeavesDFS[T any](node *Node[T], leaves []*Node[T]) []*Node[T] {
	if len(node.Children) == 0 {
		leaves = append(leaves, node)
	} else {
		for _, child := range node.Children {
			leaves = findLeavesDFS(child, leaves)
		}
	}

	return leaves
}

// helper used in recursive search for finding a node full details using DFS algorithm
func findNodeFullDFS[T any](node *Node[T], parent *Node[T], depth int, target interface{}, f func(*Node[T], interface{}) bool) Details[T] {
	if f(node, target) {
		siblings := make([]*Node[T], 0)
		for _, n := range parent.Children {
			if n != node {
				siblings = append(siblings, n)
			}
		}
		return Details[T]{
			Node:     node,
			Parent:   parent,
			Depth:    depth,
			Siblings: siblings,
		}
	}

	if len(node.Children) > 0 {
		for _, child := range node.Children {
			if found := findNodeFullDFS(child, node, depth+1, target, f); found.Node != nil {
				return found
			}
		}
	}
	return Details[T]{}
}

// Returns all nodes that satisfy target depth.
// The depth is computed starting from `root` argument which is considered level 0
func listNodesAtDepth[T any](root *Node[T], targetDepth int, currentDepth int, result []*Node[T]) []*Node[T] {
	if root == nil {
		return result
	}

	// If the current depth matches the target depth, add this node to the result.
	if currentDepth == targetDepth {
		result = append(result, root)
	} else if currentDepth < targetDepth {
		// If the current depth is less than the target depth, recursively search children.
		for _, child := range root.Children {
			result = listNodesAtDepth(child, targetDepth, currentDepth+1, result)
		}
	}

	return result
}

// builds tree starting from root node as an argument
func buildForRoot[T any](root T, values []T, compareFunc CompareFunc[T]) *Node[T] {
	// Create the root node.
	rootNode := &Node[T]{Data: root}

	// Add the root node to the parent map.
	// parentMap[root] = rootNode

	// Build the tree structure.
	for _, value := range values {
		if compareFunc(root, value) {
			// If the value is a child of the root, create the child node and add it to the parent node.
			childNode := buildForRoot(value, values, compareFunc)
			rootNode.Children = append(rootNode.Children, childNode)
		}
	}

	return rootNode
}

// FindRoots finds the root nodes from a slice of values using the provided comparison function.
// A slice of type T and a correct comaprison function must be provided.
// If comparison function is not correctly implemented, no roots will be returned
func findRoots[T any](values []T, compareFunc CompareFunc[T]) []T {
	if len(values) == 0 {
		return nil
	}

	// Create a map to store values with matching parents.
	parentMap := make(map[interface{}]struct{})

	for i, value := range values {
		for j, other := range values {
			if i != j && compareFunc(value, other) {
				// If a matching parent is found, mark it.
				parentMap[other] = struct{}{}
			}
		}
	}

	// Create a slice to store root nodes (values with no matching parent).
	var roots []T

	for _, value := range values {
		if _, isParent := parentMap[value]; !isParent {
			roots = append(roots, value)
		}
	}

	return roots
}

// Generate slice of T starting from object node as the root
// This is anti-build process
func toSlice[T any](node *Node[T], s *[]T) []T {
	// Append root to the slice
	*s = append(*s, node.Data)

	// Traverse children and add to slice
	if len(node.Children) != 0 {
		for _, c := range node.Children {
			toSlice[T](c, s)
			// toAdd := toSlice[T](c, s)
			// s = append(s, toAdd...)
		}
	}

	return *s
}
