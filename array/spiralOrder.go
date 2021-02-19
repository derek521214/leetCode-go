package main

import "fmt"

/**
54. 螺旋矩阵
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:
输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 */
var result []int
func spiralOrder(matrix [][]int) []int {
	result = make([]int, 0)
	addInts(matrix)
	return result
}

func addInts(matrix [][]int)  {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	columns := len(matrix[0])
	if columns == 0 {
		return
	}
	if rows == 1 {
		result = append(result, matrix[0]...)
		return
	}
	if columns == 1 {
		for n := 0; n < rows; n ++ {
			result = append(result, matrix[n][0])
		}
		return
	}
	//第一行
	result = append(result, matrix[0]...)
	//最后一列 不包括第一行和最后一行
	for i := 1; i < rows-1; i++ {
		result = append(result, matrix[i][columns-1])
	}
	//最后一行
	for j :=columns-1; j >= 0; j-- {
		result = append(result, matrix[rows-1][j])
	}
	//第一列
	for k := rows-2; k > 0; k-- {
		result = append(result, matrix[k][0])
	}
	for m:= 1; m < rows-1; m++ {
		matrix[m] = matrix[m][1:columns-1]
	}
	//把第一行和最后一行去掉
	matrix = matrix[1:rows-1]
	addInts(matrix)
}

/**
范例
 */
func spiralOrder1(matrix [][]int) []int {
	top := 0
	bottom := len(matrix)-1
	if bottom < 0 {
		return []int{}
	}
	left :=0
	right := len(matrix[0])-1
	ret := []int{}
	for left <= right && top <= bottom{
		printCircle(matrix,top,bottom,left,right,&ret)
		left++
		top++
		right--
		bottom--
	}
	return ret
}

func printCircle(matrix [][]int,top int,bottom int,left int,right int,ret *[]int){
	for i:=left;i<=right;i++{
		*ret =append(*ret,matrix[top][i])
	}

	for i:=top+1;i <= bottom;i++{
		*ret = append(*ret,matrix[i][right])
	}

	if left <right && top < bottom{
		for i:=right-1;i>=left;i--{
			*ret = append(*ret,matrix[bottom][i])
		}
		for i:= bottom-1;i>top;i--{
			*ret = append(*ret,matrix[i][left])
		}
	}
}

func main()  {
	test := [][] int {
	{ 1, 2, 3 },
	{ 4, 5, 6 },
	{ 7, 8, 9 },
	}
	/*test := [][]int{
		{1},
		{2},
		{3},
	}*/
	res :=spiralOrder(test)
	fmt.Print(res)
}