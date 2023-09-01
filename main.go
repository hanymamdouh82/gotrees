package main

import "fmt"

func main() {

	fmt.Println("Running: FindBFS()")
	resultRFBFS := boss.FindBFS("Amr", func(n *Node[Person], s interface{}) bool {
		return n.Data.Name == s
	})
	fmt.Printf("target\t: %s\t Result Age\t: %v\n", "Amr", resultRFBFS.Data.Age)

	fmt.Println("Running: FindDFS()")
	resultRFDFS := boss.FindDFS("Amr", func(n *Node[Person], s interface{}) bool {
		return n.Data.Name == s
	})
	fmt.Printf("target\t: %s\t Result Age\t: %v\n", "Amr", resultRFDFS.Data.Age)

	fmt.Println("Running: FindId()")
	resultRFBFS = boss.FindId("44")
	fmt.Printf("target\t: %s\t Result Age\t: %v\n", "44", resultRFBFS.Data.Age)

	fmt.Println("Running: Leaves()")
	leaves := boss.Leaves()
	for i, l := range leaves {
		fmt.Printf("Leaf idx: %v\t Leaf Name: %s\n", i, l.Data.Name)
	}

	fmt.Println("Running: Depth()")
	depth := boss.Depth()
	fmt.Printf("Node Name: %s\t Depth: %v\n", boss.Data.Name, depth)

	fmt.Println("Running: FindFullDFS()")
	det := boss.FindFullDFS("Amr", func(n *Node[Person], t interface{}) bool {
		return n.Data.Name == t
	})
	if det.Node != nil {
		fmt.Printf("Node Name\t: %s\n", det.Node.Data.Name)
		fmt.Printf("Parent Name\t: %s\n", det.Parent.Data.Name)
		fmt.Printf("Node Depth\t: %v\n", det.Depth)
		for _, s := range det.Siblings {
			fmt.Printf("Sibling Name\t: %s\n", s.Data.Name)
		}
	}

	fmt.Println("Running: Level()")
	levelNodes := boss.Level(2)
	for _, n := range levelNodes {
		fmt.Printf("Node Name\t:%s\n", n.Data.Name)
	}

}
