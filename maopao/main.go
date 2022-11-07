package main

import "fmt"

//冒泡
func maoPaoSort(arr *[5]int) {

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

func main() {
	arr := [5]int{15, 90, 40, 32, 76}
	maoPaoSort(&arr)
	fmt.Println(arr)
}
