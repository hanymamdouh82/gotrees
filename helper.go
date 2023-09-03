// Copyright 2023 Hany Mamdouh. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package gotrees

import (
	"encoding/json"
)

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

// Tree size use for recuresive operation to get tree size
func size[T any](node *Node[T], currentSize *int) int {
	*currentSize++
	if len(node.Children) != 0 {
		for _, n := range node.Children {
			size(n, currentSize)
		}
	}
	return *currentSize
}

// FindLowestCommonAncestor finds the lowest common ancestor of two nodes in a tree.
func findLowestCommonAncestor[T any](root, p, q *Node[T]) *Node[T] {
	if root == nil {
		return nil
	}

	// If the current node matches either p or q, it is the LCA.
	if root == p || root == q {
		return root
	}

	// Recursively search for p and q in the children nodes.
	var lca *Node[T]
	for _, child := range root.Children {
		childLCA := findLowestCommonAncestor(child, p, q)
		if childLCA != nil {
			if lca != nil {
				// If a previous LCA was found, this node is the new LCA.
				return root
			}
			lca = childLCA
		}
	}

	// Return the LCA found in the children (if any).
	return lca
}

// Recursive function to serialize a node and its children.
func serializeNode[T any](node *Node[T]) (map[string]interface{}, error) {
	if node == nil {
		return nil, nil
	}

	// Future refactor: replace json marshal/unmarsham with another efficient implementation
	nodeData := make(map[string]interface{}, 0)
	j, err := json.Marshal(node.Data)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(j, &nodeData)

	// Serialize children recursively.
	if len(node.Children) > 0 {
		childNodes := make([]map[string]interface{}, len(node.Children))
		for i, child := range node.Children {
			childNodes[i], err = serializeNode(child)
			if err != nil {
				return nil, err
			}
		}
		nodeData["Children"] = childNodes
	}

	// Add the node data to the parent node map.
	return nodeData, nil
}

// Recursive function to deserialize a node and its children from a map.
func deserializeJSON[T any](nodeData map[string]interface{}) *Node[T] {
	if nodeData == nil {
		return nil
	}

	// Extract Id if provided
	var id string
	id, ok := nodeData["Id"].(string)
	if !ok {
		id = ""
	}

	// Convert map to T struct
	var temp T
	j, _ := json.Marshal(nodeData)
	json.Unmarshal(j, &temp)

	// Create a new node.
	node := &Node[T]{
		Id:   id,
		Data: temp,
	}

	// Deserialize children recursively.
	childrenData, hasChildren := nodeData["Children"]
	if hasChildren {
		children := childrenData.([]interface{})
		for _, childData := range children {
			childNodeData := childData.(map[string]interface{})
			childNode := deserializeJSON[T](childNodeData)
			node.Children = append(node.Children, childNode)
		}
	}

	return node
}

// Recursive helper function to get all nodes from root to a specific node
func rootToNode[T any](root *Node[T], target *Node[T]) []*Node[T] {
	if root == nil {
		return nil
	}

	if root == target {
		return []*Node[T]{root}
	}

	for _, child := range root.Children {
		path := rootToNode(child, target)
		if len(path) > 0 {
			return append([]*Node[T]{root}, path...)
		}
	}

	return nil // Target node not found in the subtree.
}
