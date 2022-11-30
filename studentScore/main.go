package main

import (
	"fmt"
)

func main() {
	// var score [3][5]float64 = [3][5]float64{
	// 	{10, 20, 30, 40, 50}, {60, 70, 80, 90, 100}, {10, 20, 30, 40, 50},
	// }

	var score [3][5]float64

	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			fmt.Printf("输输入第%d班，第%d个学生成绩", i+1, j+1)
			fmt.Scanln(&score[i][j])
		}
	}

	for k, v := range score {
		sum := 0.0
		for _, v1 := range v {
			sum += v1
		}
		fmt.Printf("第%d班,总分%v,平均分%v\n", k+1, sum, sum/float64(len(v)))
	}

	fmt.Println(score)
}
