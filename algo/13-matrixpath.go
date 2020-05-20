package main

import "fmt"

//题目13:矩阵中的路径
/*
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[
["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]
]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
 */

func exist(board [][]byte, word string) bool {
	row := len(board) //特别注意对数组的判断,比如传进来的是[],先判断row,有row后再判断col
	if row < 1 {
		return false
	}

	col := len(board[0])
	if col < 1 {
		return false
	}

	for i, v1 := range board {
		for j, _ := range v1 { //每个元素开始搜寻时都是新的
			if pathed(board, word, i, j, row, col, 0) {
				return true
			}
		}
	}

	return false
}

//已board中i,j位置处的元素为起点搜索是否具有匹配的路径
func pathed(board [][]byte, word string, i, j int, row, col int, k int) bool {
	if board[i][j] != word[k] {
		return false
	}

	if k == len(word) - 1 {
		return true
	}

	tmp := board[i][j]
	board[j][j] = byte(' ')   //避免被折回

	//向上搜寻
	if i-1 >= 0 && pathed(board, word, i-1, j, row, col, k+1) {
		return true
	}

	//向下搜寻
	if i+1 < row && pathed(board, word, i+1, j, row, col, k+1) {
		return true
	}

	//向左搜寻
	if j-1 >= 0 && pathed(board, word, i, j-1, row, col, k+1) {
		return true
	}

	//向右搜寻
	if j+1 < col && pathed(board, word, i, j+1, row, col, k+1) {
		return true
	}

	board[i][j] = tmp

	return false
}

func main() {
	board := [][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'},
	}

	str := "bfce"
	find := exist(board, str)
	fmt.Printf("str:%s exists in borad:%t\n", str, find)
}
