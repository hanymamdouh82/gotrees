package main

// func customCompare[C any](node *Node[Person], name string) bool {
// 	// Change this comparison logic according to your criteria.
// 	return node.Data.Name == name
// }

// func findByNameBFS[T any, C any](root *Node[T], target C, f func(*Node[T], C) bool) *Node[T] {
// 	queue := make([]*Node[T], 0)
// 	queue = append(queue, root)

// 	for len(queue) > 0 {
// 		nextUp := queue[0] // take first element in the queue for insepction
// 		queue = queue[1:]  // pop firest element from the slice

// 		// check if this is the Node we are searching for
// 		// if it is, return
// 		// if nextUp.Data.Name == name {
// 		// 	return nextUp
// 		// }
// 		if f(nextUp, target) {
// 			return nextUp
// 		}

// 		// otherwise, add its children to the queue
// 		if len(nextUp.Children) > 0 {
// 			queue = append(queue, nextUp.Children...)
// 		}
// 	}
// 	return nil
// }

// func findByNameDFS(node *Node, name string) *Node {
// 	if node.Name == name {
// 		return node
// 	}

// 	if len(node.Children) > 0 {
// 		for _, child := range node.Children {
// 			if foundNode := findByNameDFS(child, name); foundNode != nil {
// 				return foundNode
// 			}
// 		}
// 	}
// 	return nil
// }

// func findLeavesDFS(node *Node, leaves []*Node) []*Node {
// 	if len(node.Children) == 0 {
// 		leaves = append(leaves, node)
// 	} else {
// 		for _, child := range node.Children {
// 			leaves = findLeaves(child, leaves)
// 		}
// 	}

// 	return leaves
// }

// func findNodeParentByNameDFS(node *Node, parent *Node, depth int, name string) (result SearchResult) {
// 	if node.Name == name {
// 		siblings := make([]*Node, 0)
// 		for _, n := range parent.Children {
// 			if n != node {
// 				siblings = append(siblings, n)
// 			}
// 		}
// 		return SearchResult{
// 			Node:     node,
// 			Parent:   parent,
// 			Depth:    depth,
// 			Siblings: siblings,
// 		}
// 	}

// 	if len(node.Children) > 0 {
// 		for _, child := range node.Children {
// 			if found := findNodeParentByNameDFS(child, node, depth+1, name); found.Node != nil {
// 				return found
// 			}
// 		}
// 	}
// 	return SearchResult{}
// }
