// Copyright 2023 Hany Mamdouh. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package gotrees

import "encoding/json"

// Comparison function for building a tree.
// First argument is the parent, second argument is the child
// you can encapsulate your logic for parent/child relationship inside the function
type CompareFunc[T any] func(T, T) bool

// Comparison function for searching node.
// First argument is the node, second argument is the search parameter
// You can encapsulate your logic for search inside it
type FindFunc[T any] func(n *Node[T], C interface{}) bool

// The node structure. Each node is a container for any type of structs or primative types.
// Node can be identified by `Id`, which helps in fast searching and doesn't require comparison function.
// Id value is the responsibility of the consumer, you can use any identification method to identify nodes.
// If Id is repeated within a tree, first occurance will be picked during FindId()
type Node[T any] struct {
	Id       string
	Data     T
	Children []*Node[T]
}

// Describes the full details of a Node.
// Is used only as a result of FindFullDFS() receiver function.
type Details[T any] struct {
	Node     *Node[T]
	Parent   *Node[T]
	Depth    int
	Siblings []*Node[T]
}

// Adds node to the current node and returns its memory reference.
func (n *Node[T]) AddNode() *Node[T] {
	node := Node[T]{}
	n.Children = append(n.Children, &node)
	return &node
}

// find node by its Id and return it
func (n *Node[T]) FindId(id string) *Node[T] {
	if n.Id == id {
		return n
	}

	if len(n.Children) > 0 {
		for _, child := range n.Children {
			if foundNode := child.FindId(id); foundNode != nil {
				return foundNode
			}
		}
	}
	return nil
}

// Find node by compare function, using Breadth First Search (BFS) algorithm.
func (n *Node[T]) FindBFS(target interface{}, f FindFunc[T]) *Node[T] {

	queue := make([]*Node[T], 0)
	queue = append(queue, n)

	for len(queue) > 0 {
		nextUp := queue[0] // take first element in the queue for insepction
		queue = queue[1:]  // pop firest element from the slice

		// check if this is the Node we are searching for
		if f(nextUp, target) {
			return nextUp
		}

		// otherwise, add its children to the queue
		if len(nextUp.Children) > 0 {
			queue = append(queue, nextUp.Children...)
		}
	}
	return nil
}

// Find node by comparison function, using Depth First Search (DFS) algorithm.
func (n *Node[T]) FindDFS(target interface{}, f FindFunc[T]) *Node[T] {
	if f(n, target) {
		return n
	}

	if len(n.Children) > 0 {
		for _, child := range n.Children {
			// if foundNode := findByNameDFS(child, name); foundNode != nil {
			if foundNode := child.FindDFS(target, f); foundNode != nil {
				return foundNode
			}
		}
	}
	return nil
}

// Find node by comparison function, using Depth First Search (DFS) algorithm, and return full node details.
func (n *Node[T]) FindFullDFS(target interface{}, f FindFunc[T]) Details[T] {
	det := findNodeFullDFS(n, nil, 0, target, f)
	return det
}

// Find all leaves starting from object node
// Object node is conisdered root node
func (n *Node[T]) Leaves() []*Node[T] {
	leaves := findLeavesDFS[T](n, []*Node[T]{})
	return leaves
}

// Find depth starting from object node.
// Node is considered as Root node
func (n *Node[T]) Depth() int {
	if n == nil {
		return 0
	}

	// Initialize the depth to 1, as the root node is at depth 1.
	depth := 1
	maxChildDepth := 0

	// Calculate the maximum depth of children nodes.
	for _, child := range n.Children {
		childDepth := child.Depth()
		if childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}

	// The depth of the current node is the maximum child depth plus one.
	depth += maxChildDepth

	return depth
}

// Returns tree size.
// Tree size is the count of all nodes inside a tree. Consider current node object as root node.
func (n *Node[T]) Size() int {
	s := new(int)
	size := size(n, s)
	return size
}

// List all nodes at certain depth, starting from object node which is considered as root node
func (n *Node[T]) Level(d int) []*Node[T] {
	result := listNodesAtDepth(n, d, 0, []*Node[T]{})
	return result
}

// Returns slice of T objects from current Node.
// Current Node object is considered as root node.
func (n *Node[T]) Slice() []T {
	s := toSlice[T](n, &[]T{})
	return s
}

// Returns Lowest Common Ancestor for current Node Object
func (n *Node[T]) LCA(p, q *Node[T]) *Node[T] {
	node := findLowestCommonAncestor[T](n, p, q)
	return node
}

// Serialize a tree into JSON format.
// Use as compatibile format to transfer over wire or store into NoSQL
func (n *Node[T]) SerializeJSON() (string, error) {
	// Use a map to represent each node as a JSON object.
	nodeMap, err := serializeNode[T](n)
	if err != nil {
		return "", err
	}

	// Convert the map to a JSON string.
	jsonData, err := json.Marshal(nodeMap)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
