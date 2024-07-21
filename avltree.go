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

func (node *AVLNode) getHeight() int {
	if node == nil {
		return -1
	}
	return node.Height
}

func (node *AVLNode) updateHeight() {
	node.Height = max(node.Left.getHeight(), node.Right.getHeight()) + 1
}

func (node *AVLNode) getBalance() int {
	if node == nil {
		return 0
	}
	return node.Left.getHeight() - node.Right.getHeight()
}

func (node *AVLNode) rightRotate() *AVLNode {
	left := node.Left
	t3 := left.Right

	// Выполнение поворота
	left.Right = node
	node.Left = t3

	// Обновление высот после поворота
	node.updateHeight()
	left.updateHeight()

	// Возвращение нового корня после поворота
	return left
}

func removeDuplicates(nums []int) int {
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

func preOrderTrav(tree *AVLNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Val)
	preOrderTrav(tree.Left)
	preOrderTrav(tree.Right)
}

func main() {
	one := []int{1, 1, 2, 2, 3, 3}
	fmt.Print(one)
	removeDuplicates(one)
	fmt.Print(one)
}
