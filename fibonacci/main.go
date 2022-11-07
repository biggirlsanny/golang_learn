package main

import (
	"fmt"
)

//递归法：返回对应数
func fbn1(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn1(n-1) + fbn1(n-2)
	}
}

//切片法：返回数组
func fbn2(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	if n == 1 {
		fbnSlice[0] = 1
	}
	if n >= 2 {
		fbnSlice[0] = 1
		fbnSlice[1] = 1
		for i := 2; i < n; i++ {
			fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
		}
	}

	return fbnSlice
}

func main() {

	res := fbn2(3)
	fmt.Println("res=", res)

	fmt.Println("res=", fbn1(6))
}
