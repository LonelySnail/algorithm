package main

import "fmt"

/*
初始时在序列中找到最小（大）元素，
放到序列的起始位置作为已排序序列；
然后，再从剩余未排序元素中继续寻找最小（大）元素，放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。

 */

func selectSort(array []int,length int)  {
	var min,count int
	for i :=0; i<length;i++ {
		min =i
		for j:=i+1;j<length;j++ {
			if array[j]<array[min] {
				min = j
			}
		}
		if min != i {
			count++
			array[i],array[min] = array[min],array[i]
		}

	}
	fmt.Println(count)
}


func main()  {
	a := []int{1,2,3,6,5,4,88,56,4,33,45,23,17}
	selectSort(a,len(a))
	fmt.Println(a,"finish")
}