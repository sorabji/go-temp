package main

import (
	"fmt"
)

func main(){
	var low, high, step float32
	var results [][]float32

	fmt.Print("low value: ")
	fmt.Scanf("%f", &low)
	fmt.Print("high value: ")
	fmt.Scanf("%f", &high)
	fmt.Print("step: ")
	fmt.Scanf("%f", &step)

	for i := low; i < high; i = i + step {
		tmp := i * 1.8 + 32
		results = append(results, []float32{i, tmp})
	}

	for _, val := range results {
		fmt.Printf("%f\t%f\n", val[0], val[1])
	}
}
