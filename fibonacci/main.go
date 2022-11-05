package main

import (
	"fmt"
)

//递归法
func fbn1(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn1(n-1) + fbn1(n-2)
	}
}

func main() {

	res := fbn1(3)
	fmt.Println("res=", res)

	fmt.Println("res=", fbn1(6))
}
