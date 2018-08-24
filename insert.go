package main

import "fmt"

/*
	插入排序就是每一步都将一个待排数据按其大小插入到已经排序的数据中的适当位置，直到全部插入完毕。
	array 待排序数组
	length 数组长度
 */


func insert(array []int,length int)  {
	var  k int
	for n :=1;n <length;n++ {
		if array[n] < array[n-1] {
			temp := array[n]
			for k =n-1; array[k]>temp; k--{
				array[k+1] = array[k]
			}
			array[k+1] = temp
		}

	}
}

func insert1(array []int,length int)  {
	var  j,k int
	for n :=1;n <length;n++ {

		//在已经排序的数组中 找到一个比array[n]小的
		for j =n-1;j>=0;j-- {
			if array[j]< array[n]{
				break
			}
		}
		if j != n -1 {
			temp := array[n]
			for k =n-1; k>j; k--{
				array[k+1] = array[k]
			}
			array[k+1] = temp
		}


	}
}

func main()  {
	a := []int{1,2,3,6,5,4,88,56,4,33,45,23,17}
	insert(a,len(a))
	fmt.Println(a,"finish")
}