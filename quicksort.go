package main

import "fmt"

/*
	快排
	array  待排序数组
	left -- 数组的左边界(例如，从起始位置开始排序，则left=0)
	right -- 数组的右边界(例如，排序截至到数组末尾，则right=len(array)-1)
 */
func quickSort(array []int,left,right int)  {
	if left > right{
		return
	}
	l,r := left,right
	val := array[l]
	for  l < r {
		//从右向左找到第一个比val小的元素
		for l < r && array[r]>=val{
			r--
		}

		if l < r {
			array[l] = array[r]
		}

		//从左向右找到第一个比val大的元素
		for l<r && array[l]<= val {
			l++
		}

		if l<r {
			array[r] = array[l]
		}

	}

	array[l] = val
	quickSort(array,left,l-1)
	quickSort(array,l+1,right)

}

func main()  {
	//a := []int{30,40,60,10,20,50}
	a := []int{1,2,3,6,5,4,88,56,4,33,45,23,17}

	quickSort(a,0,len(a)-1)
	fmt.Println(a,"finish")
}