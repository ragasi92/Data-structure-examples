package main

import "log"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (node *Node) Insert(value int) {
	if value < node.Key {
		if node.Left == nil {
			node.Left = &Node{Key: value}
		} else {
			node.Left.Insert(value)
		}
	} else if value > node.Key {
		if node.Right == nil {
			node.Right = &Node{Key: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

func (node *Node) Search(value int) bool {

	if node == nil {
		return false
	}

	if value < node.Key {
		return node.Left.Search(value)
	} else if value > node.Key {
		return node.Right.Search(value)
	}

	return true
}

func main() {
	binaryTree := &Node{Key: 500}

	buildBinaryTree := []int{600, 30, 1, 2, 5, 78, 90}

	for _, value := range buildBinaryTree {
		binaryTree.Insert(value)
		log.Println(binaryTree)
	}

	seachValues := []int{3, 5, 1, 900, 5000}
	for _, value := range seachValues {
		log.Printf("Is %d exist in the tree: %v", value, binaryTree.Search(value))
	}
}
