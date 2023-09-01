package main

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

// ----------------------------------------------------------------------------------------------------------------------------------------

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

// ----------------------------------------------------------------------------------------------------------------------------------------

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
