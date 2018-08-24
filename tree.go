package main

import (
	"fmt"
)

type tree struct{
	left *tree
	data  int
	right *tree
}

func createTree(arr []int) []tree  {
	array := make([]tree,0)
	for i, val := range arr {
		array = append(array,tree{})
		array[i].data = val
	}

	for i := 0; i <len(arr)/2; i++ {
		array[i].left = &array[i*2+1]
		if i*2+2 < len(array) {
			array[i].right = &array[i*2+2]
		}
	}

	return array
}


//前序遍历
func preOrder(root tree,preList *[]int) {
	*preList = append(*preList,root.data)

	if root.left != nil {
		preOrder(*root.left,preList)
	}

	if root.right != nil {
		preOrder(*root.right,preList)
	}

}

//中序遍历
func inOrder(root tree,inList *[]int)  {
	if root.left != nil {
		inOrder(*root.left,inList)
	}
	*inList = append(*inList,root.data)
	if root.right != nil {
		inOrder(*root.right,inList)
	}
}


//后序遍历
func afterOrder(root tree,afterList *[]int)  {
	if root.left != nil {
		afterOrder(*root.left,afterList)
	}

	if root.right != nil {
		afterOrder(*root.right,afterList)
	}

	*afterList = append(*afterList,root.data)
}

func main()  {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	array := createTree(arr)
	preList := make([]int,0)
	preOrder(array[0],&preList)
	fmt.Println(preList,"前序遍历")
	inList := make([]int,0)
	inOrder(array[0],&inList)
	fmt.Println(inList,"中序遍历")
	afterList := make([]int,0)
	afterOrder(array[0],&afterList)
	fmt.Println(afterList,"后序遍历")
}