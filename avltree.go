package main

import (
	"fmt"
)

type AVLNode struct {
	Val    int
	Height int
	Left   *AVLNode
	Right  *AVLNode
}

func (node *AVLNode) GetHeight() int {
	if node == nil {
		return -1
	}
	return node.Height
}

func (node *AVLNode) UpdateHeight() {
	node.Height = max(node.Left.GetHeight(), node.Right.GetHeight()) + 1
}

func (node *AVLNode) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.Left.GetHeight() - node.Right.GetHeight()
}

func (node *AVLNode) RightRotate() *AVLNode {
	left := node.Left
	t3 := left.Right

	left.Right = node
	node.Left = t3

	node.UpdateHeight()
	left.UpdateHeight()

	return left
}

func RemoveDuplicates(nums []int) int {
	last := nums[0]
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != last {
			nums[j] = nums[i]
			j++
		}
		last = nums[i]
	}
	return j
}

func PreOrderTrav(tree *AVLNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	PreOrderTrav(tree.Left)
	PreOrderTrav(tree.Right)
}
