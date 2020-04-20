package main

import (
	"fmt"
	"time"
)

func main() {

	start1 := time.Now() // 获取当前时间
	//gas  = [2,3,4]
	//cost = [3,4,3]
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	fmt.Println(buildTree(preorder, inorder))
	elapsed1 := time.Since(start1)
	fmt.Println("该函数执行完成耗时test：", elapsed1)

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 *  前序遍历：根左右
 *  中序遍历：左根右
 *  后序遍历：左右根
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	return nil
}
