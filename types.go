package main

type Node[T any] struct {
	Id       string
	Data     T
	Children []*Node[T]
}

type Details[T any] struct {
	Node     *Node[T]
	Parent   *Node[T]
	Depth    int
	Siblings []*Node[T]
}

type Person struct {
	Name string
	Age  int
}

// ----------------------------------------------------------------------------------------------------------------------------------------

// find node by its Id
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

// ----------------------------------------------------------------------------------------------------------------------------------------

// Find node by compare function, using Breadth First Search (BFS) algorithm.
func (n *Node[T]) FindBFS(target interface{}, f func(*Node[T], interface{}) bool) *Node[T] {

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

// ----------------------------------------------------------------------------------------------------------------------------------------

// Find node by comparison function, using Depth First Search (DFS) algorithm.
func (n *Node[T]) FindDFS(target interface{}, f func(*Node[T], interface{}) bool) *Node[T] {
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

// ----------------------------------------------------------------------------------------------------------------------------------------

// Find node by comparison function, using Depth First Search (DFS) algorithm, and return full node details.
func (n *Node[T]) FindFullDFS(target interface{}, f func(*Node[T], interface{}) bool) Details[T] {
	det := findNodeFullDFS(n, nil, 0, target, f)
	return det
}

// ----------------------------------------------------------------------------------------------------------------------------------------

// Find all leaves starting from node
func (n *Node[T]) Leaves() []*Node[T] {
	leaves := findLeavesDFS[T](n, []*Node[T]{})
	return leaves
}

// ----------------------------------------------------------------------------------------------------------------------------------------

// Find depth starting from current node. Node is considered as Root node
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

// ----------------------------------------------------------------------------------------------------------------------------------------

// List all nodes at certain depth, starting from provided node which is considered as root node
func (n *Node[T]) Level(d int) []*Node[T] {
	result := listNodesAtDepth(n, d, 0, []*Node[T]{})
	return result
}
