package main

import "fmt"

/*冒泡
	array 待排序数组
*/


func bubble(array []int)  {
	var count,count1 int
	for i := len(array)-1;i>0;i-- {
		count++
		for j:=0;j<i;j++ {

			if array[j]< array[j+1] {
				count1++
				array[j],array[j+1] = array[j+1],array[j]
			}
		}
	}
	fmt.Println(array)
	fmt.Println(count,count1)
}


func bubble1(array []int)  {
	var count,count1 int
	for i := 0;i< len(array)-1;i++ {
		count++
		for j:=len(array)-1;j >i;j-- {

			if array[j-1]< array[j] {
				count1++
				array[j],array[j-1] = array[j-1],array[j]
			}
		}
	}
	fmt.Println(array)
	fmt.Println(count,count1)
}
//
func bubble2(array []int)  {
	var count,count1 int
	for i :=0;i <len(array)-1;i++ {
		count++
		for j:=i+1;j <len(array);j++ {
			if array[i] <array[j] {
				count1++
				array[j],array[i] = array[i],array[j]
			}
		}
	}
	fmt.Println(array)
	fmt.Println(count,count1)
}

// 冒泡算法优化版
func BubbleSort(attr []int) []int {
	changeInx := 0
	sortBorder :=len(attr)-1
	for i :=0;i <len(attr);i++ {
		isSorted := true
		for j :=0;j <sortBorder;j++ {
			if attr[j] >attr[j+1] {
				attr[j],attr[j+1] = attr[j+1],attr[j]
				isSorted = false
				changeInx = j
			}
		}
		sortBorder = changeInx
		if isSorted {
			break
		}
	}
	return attr
}

func main()  {
	a := []int{1,2,3,6,5,4,88,56,4,33,45,23,17}
	bubble(a)
	//fmt.Println(a)
	//bubble1(a)
}
