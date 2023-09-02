# Go Tree Package

![Go Version](https://img.shields.io/badge/Go-v1.21%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)

The Go Tree Package provides a flexible and generic tree data structure along with useful functions for working with trees in Go. It is designed to be easy to use and adaptable to various data types.

## Features

- **Generic Tree Structure**: The package supports a generic `Node[T]` structure, allowing you to create trees with nodes containing data of any type `T`.

- **Serialization**: Serialize your tree to JSON and deserialize it back, making it easy to save and load tree structures.

- **Tree Traversal**: Implement depth-first and breadth-first traversal algorithms to navigate your tree.

- **Searching**: Find nodes by name or other criteria.

- **Lowest Common Ancestor (LCA)**: Determine the lowest common ancestor of two nodes in a tree.

- **List Nodes at Certain Depth**: Get a list of nodes at a specified depth in the tree.

- **Building Trees**: Build trees from slices of data using a comparison function to determine parent-child relationships.

- **Node Identification**: Nodes can be identified by `Id`, making it efficient for fast searching without comparison functions.

- **Full Node Details**: Retrieve detailed information about a node, including its parent, depth, and siblings.

- **Leaves**: Find all leaf nodes in the tree.

- **Tree Size**: Get the size (number of nodes) of the tree.

- **Slice Conversion**: Convert the tree into a slice of data.

## Installation

To use the Go Tree Package in your project, you can install it using `go get`:

```bash
go get github.com/yourusername/gotrees
```

## Examples

### Searching for Nodes

```go
// Find a node by its Id.
foundNode := tree.FindId("Node A")

// Find a node using a custom comparison function and Breadth First Search (BFS).
targetNode := tree.FindBFS("TargetNode", func(node *gotrees.Node[T], target interface{}) bool {
    // Implement your custom comparison logic here.
    return node.Name == target
})

// Find a node using a custom comparison function and Depth First Search (DFS).
targetNode := tree.FindDFS("TargetNode", func(node *gotrees.Node[T], target interface{}) bool {
    // Implement your custom comparison logic here.
    return node.Name == target
})
```
### Serialization and Deserialization

```go
// Serialize the tree to JSON.
jsonString, err := tree.SerializeJSON()
if err != nil {
    // Handle error
}

// Deserialize a JSON string into a tree.
newTree, err := gotrees.DeserializeJSONToTree(jsonString)
if err != nil {
    // Handle error
}

```

### Lowest Common Ancestor (LCA)

```go
// Find the Lowest Common Ancestor (LCA) of two nodes.
lcaNode := tree.LCA(nodeA, nodeB)
```

### List Nodes at Certain Depth

```go
// List all nodes at a specific depth in the tree.
nodesAtDepth := tree.Level(2)
```

### Building Trees

```go
// Build a tree from a slice of data using a custom comparison function.
// This function determines the parent-child relationships.
values := []T{ /* Your data slice here */ }
treeNodes := gotrees.Build(values, func(parent, child T) bool {
    // Implement your custom logic for parent-child relationships here.
    return /* Your logic here */
})
```