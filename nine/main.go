package main

import "fmt"

type MethodUtils struct {
}

func (method *MethodUtils) jiujiu(n int) {
	if n == 0 {
		//fmt.Printf("请输入整数\n")
		fmt.Scanln(&n)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v ", i, j, i*j)

		}
		fmt.Println()
	}
}

func main() {
	var mu MethodUtils
	//mu.Print()
	// var n int
	// fmt.Printf("请输入整数\n")
	// fmt.Scanln(&n)
	mu.jiujiu(0)
}
