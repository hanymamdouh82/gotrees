# Gotrees Features

the following summarizes features to be added to package.

## Traversal Algorithms:

- [x] Depth-First Search (DFS): You've already implemented this, but you can extend it to perform pre-order, in-order, or post-order traversals.
- [x] Breadth-First Search (BFS): Similar to DFS, but it explores all nodes at the current level before moving to the next level.

# Tree Operations:

- [x] Insertion: Add a new node to the tree.
- [x] Deletion: Remove a node from the tree.
- [x] Search: Find a specific node in the tree.
- [ ] Update: Modify the data of a node in the tree.

## Properties:

- [x] Size: Determine the number of nodes in the tree.
- [x] Height/Depth: Calculate the height or depth of the tree.
- [ ] Is Balanced: Check if the tree is balanced (e.g., AVL tree or Red-Black tree).

## Traversal Variants:

- [ ] Level Order Traversal: Visit nodes level by level, from left to right.
- [ ] In-order Successor/Predecessor: Find the next or previous node in in-order traversal.
- [ ] Kth Smallest/Largest Node: Find the kth smallest or largest node in the tree.

## Tree Transformation:

- [ ] Mirror/Invert Tree: Swap left and right subtrees at each node.
- [x] Convert to/from Array/List: Transform the tree into an array or list, and vice versa.
- [ ] Flatten to Linked List: Convert the tree into a singly linked list in pre-order or in-order sequence.

## Validation and Properties:

- [ ] Is Binary Search Tree (BST): Check if the tree satisfies the properties of a BST.
- [ ] Is Full/Complete Binary Tree: Verify if the tree is a full or complete binary tree.

## Path and Sum Operations:

- [x] Root-to-Leaf Paths: Find all paths from the root to the leaf nodes.
- [ ] Path Sum: Check if there exists a path with a given sum in the tree.
- [x] Lowest Common Ancestor (LCA): Find the lowest common ancestor of two nodes.

## Balancing (for Balanced Trees):

- [ ] Rotation: Perform rotations (e.g., left rotation, right rotation) to balance the tree.
- [ ] Rebalancing: Ensure the tree remains balanced after insertions and deletions.

## Serialization and Deserialization:

- [x] Serialize: Convert the tree into a string or an array for storage or transmission.
- [x] Deserialize: Recreate the tree from a serialized form.

## Deletion Strategies:

- [ ] Lazy Deletion: Mark nodes as deleted without actually removing them.
- [ ] Real Deletion: Physically remove nodes marked for deletion.

## Tree Trimming:

- [ ] Trim Tree: Remove nodes that do not meet certain criteria (e.g., values less than a threshold).

## Iterative Traversal:

- [x] Implement DFS and BFS using iterative (non-recursive) methods.