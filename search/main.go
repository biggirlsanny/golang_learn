package main

import (
	"fmt"
	"time"
)

//冒泡排序，配合二分查找
func maoPaoSort(arr *[6]int) {

	temp := 0
	for j := 0; j < len(*arr)-1; j++ {
		for i := 0; i < len(*arr)-1-j; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
		}
	}
}

//二分查找
func BinaryFind(arr *[6]int, left int, right int, find int) {
	middle := (left + right) / 2
	fmt.Println("中间：", left, right, middle)

	if left > right {
		fmt.Println("没找到")
		return
	}
	if (*arr)[middle] > find {
		BinaryFind(arr, left, middle-1, find)
	} else if (*arr)[middle] < find {
		BinaryFind(arr, middle+1, right, find)
	} else {
		fmt.Printf("找到了,下标为%v\n", middle)
	}
}

//顺序查找
func search(arr *[5]string) {
	var name = "肖战"
	fmt.Println("请输入查找的人名")
	//fmt.Scanln(&name)
	for i := 0; i < len(arr); i++ {
		if arr[i] == name {
			fmt.Printf("找到%v,下标%v \n", name, i)
			break
		} else if i == (len(arr) - 1) {
			fmt.Printf("没有找到%v \n", name)

		}
	}
}

func main() {

	//顺序查找
	start := time.Now().UnixMicro()
	arr := [5]string{"肖战", "王一博", "赵丽颖", "金晨", "胡歌"}
	search(&arr)
	end := time.Now().UnixMicro()
	fmt.Println("顺序查找执行时间为：", end-start)

	//二分查找
	start1 := time.Now().UnixMicro()
	arr1 := [6]int{53, 24, 56, 77, 90, 9}
	maoPaoSort(&arr1)
	BinaryFind(&arr1, 0, len(arr1)-1, 77)
	end1 := time.Now().UnixMicro()
	fmt.Println("二分查找执行时间为：", end1-start1)
}
