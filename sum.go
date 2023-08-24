package main

import (
	
	_ "fmt"
)

type Number interface {
	int | float64
}



// func sums() {
// 	numsInt := []int{1, 2, 3, 4, 5}
// 	numsFloat64 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

// 	fmt.Println(sum(numsInt))
// 	fmt.Println(sum(numsFloat64))
// }

// func sum[V Number](nums []V) V {
// 	var sum V
// 	for _, num := range nums{
// 		sum += num
// 	}
// 	return sum
// }

// func sumInt(nums []int) int{
// 	var sum int
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }

// func sumFloat64(nums []float64) float64{
// 	var sum float64
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }