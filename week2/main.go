package main

import "fmt"

func main()  {
	Q1()
	Q2()
}

func Q1() {
	fmt.Println("--------------")
	fmt.Printf("Q1 \n")

	scoreList := [5]float64{59,65,78,99,87}

	var sum, avg float64
	for _, score := range scoreList {
		fmt.Printf("分數: %+v \n", score)
		sum += score
	}
	avg = sum/float64(len(scoreList))

	fmt.Printf("總分： %+v \n", sum)
	fmt.Printf("平均： %+v \n", avg)
}

func Q2() {
	x := []int {
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	var min int
	for i:=0; i<len(x)-1; i++ {
		if x[i] > x[i+1] {
			min = x[i+1]
		}
	}

	fmt.Println("--------------")
	fmt.Println("Q2")
	fmt.Printf("x[]= %+v \n", x)
	fmt.Printf("min = %+v \n", min)
}