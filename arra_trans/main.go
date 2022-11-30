package main

import "fmt"

type myArray [3][3]int

//数组翻转
func (arr *myArray) trans() {
	var arr_trans [3][3]int
	for k, v := range *arr {
		for k1, v1 := range v {
			arr_trans[k1][k] = v1
		}

	}
	fmt.Println(arr_trans)
}

func main() {
	s := myArray{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	s.trans()
}
