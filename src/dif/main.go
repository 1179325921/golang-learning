package main

import "fmt"

type Node struct {
	Name       string
	ProviderID string
	InternalIP string
}

func main() {
	node1 := Node{"1", "1", "1"}
	node2 := Node{"2", "2", "2"}
	node3 := Node{"3", "3", "3"}

	allNodes := []Node{node1, node2, node3}
	masterNodes := []Node{node1}
	var n []Node
	var flag bool = false
	// 两个list取差集
	for _, value1 := range allNodes {
		flag = false
		for _, value2 := range masterNodes {
			if value1 == value2 {
				flag = true
				break
			}
		}
		if !flag {
			n = append(n, value1)
		}

	}
	fmt.Println(n)
}
