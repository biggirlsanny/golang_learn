package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

type StuSlice []Student

func (s StuSlice) Len() int {
	// sort.Sort()
	// fmt.Println("say()")
	return len(s)
}
func (s StuSlice) Less(i, j int) bool {
	// sort.Sort()
	// fmt.Println("say()")
	return true
}
func (s StuSlice) Swap(i, j int) {
	// sort.Sort()
	// fmt.Println("say()")
	//return true
}

func main() {
	var stus StuSlice
	for i := 0; i <= 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("英雄|%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: rand.Float64(),
		}
		stus = append(stus, stu)
	}

	for _, v := range stus {
		fmt.Println(v)
	}
	fmt.Println("-----------排序后------------")

	sort.Sort(stus)
	for _, v := range stus {
		fmt.Println(v)
	}

}
