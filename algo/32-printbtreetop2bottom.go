package main

import (
	"container/list"
	"fmt"
)

//题目32:从上至下打印二叉树
/*
  从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]

 */

//二叉树的按层次遍历(广度优先)
//一般是借助队列
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0) //结果序列
	queue := list.New()   //辅助队列
	queue.PushBack(root)  //根结点先入队列
	for {
		if queue.Front() == nil { //队列为空,则退出
			break
		}
		topElement, _ := queue.Front().Value.(*TreeNode)
		res = append(res, topElement.Val)
		queue.Remove(queue.Front()) //队头出队
		if topElement.Left != nil {
			queue.PushBack(topElement.Left)
		}

		if topElement.Right != nil {
			queue.PushBack(topElement.Right)
		}
	}

	return res
}

//题目32-II,从上到下层次打印二叉树,同一层的结点按从左到右的顺序打印
//每一层打印到一行
/*
 给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
 */
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)               //结果序列
	queue := list.New()                   //辅助队列
	queue.PushBack(root)                  //根结点先入队列

	nextLevelNodes := 0                   //下一行的结点数
	toBePrinted := 1                      //当前需要打印的结点数,第一行是1
	rowRes := make([]int, 0, toBePrinted) //当前行的元素

	for {
		if queue.Front() == nil { //队列为空,则退出
			break
		}
		topElement, _ := queue.Front().Value.(*TreeNode)
		rowRes = append(rowRes, topElement.Val)

		queue.Remove(queue.Front()) //队头出队
		if topElement.Left != nil {
			queue.PushBack(topElement.Left)
			nextLevelNodes++
		}

		if topElement.Right != nil {
			queue.PushBack(topElement.Right)
			nextLevelNodes++
		}
		toBePrinted--

		if toBePrinted == 0 { //当前层打印的结点数完成
			toBePrinted = nextLevelNodes
			nextLevelNodes = 0
			res = append(res, rowRes)
			rowRes = make([]int, 0, toBePrinted)
		}
	}

	return res
}

//题目32-III
/*
  请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]
 */

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)               //结果序列
	queue := list.New()                   //辅助队列
	queue.PushBack(root)                  //根结点先入队列

	nextLevelNodes := 0                   //下一行的结点数
	toBePrinted := 1                      //当前需要打印的结点数,第一行是1
	rowRes := make([]int, 0, toBePrinted) //当前行的元素
	direction := 1                        //当前行的打印方向,0:从左到右,1:从右到左

	for {
		if queue.Front() == nil { //队列为空,则退出
			break
		}
		topElement, _ := queue.Front().Value.(*TreeNode)
		rowRes = append(rowRes, topElement.Val)

		queue.Remove(queue.Front()) //队头出队

		if direction == 0 {
			if topElement.Left != nil {
				queue.PushBack(topElement.Left)
				nextLevelNodes++
			}

			if topElement.Right != nil {
				queue.PushBack(topElement.Right)
				nextLevelNodes++
			}
		} else if direction == 1 {
			if topElement.Right != nil {
				queue.PushBack(topElement.Right)
				nextLevelNodes++
			}

			if topElement.Left != nil {
				queue.PushBack(topElement.Left)
				nextLevelNodes++
			}
		}

		toBePrinted--

		if toBePrinted == 0 { //当前层打印的结点数完成
			toBePrinted = nextLevelNodes
			nextLevelNodes = 0
			res = append(res, rowRes)
			rowRes = make([]int, 0, toBePrinted)
			if direction == 0 {
				direction = 1
			} else if direction == 1 {
				direction = 0
			}
		}
	}

	return res
}

func main() {
	root := &TreeNode{3, nil, nil}
	node2 := &TreeNode{9, nil, nil}
	node3 := &TreeNode{20, nil, nil}
	node4 := &TreeNode{15, nil, nil}
	node5 := &TreeNode{7, nil, nil}

	root.Left = node2
	root.Right = node3
	node3.Left = node4
	node3.Right = node5

	res := levelOrder3(root)
	fmt.Printf("层次遍历结果为:%v\n", res)

}
