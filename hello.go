package main

import "fmt"

func main() {

	fmt.Printf("hello world!!")

	// for i := 0; i < 5; i++ {
	// 	defer fmt.Printf("%d ", i)
	// }

	array := [2]int{1, 2}
	fmt.Printf("%v \n\r", array)

	array1 := [4]int{1: 6, 3: 8}
	fmt.Printf("%v \n\r", array1)
	//平均值

	slice := []int{1, 2, 3, 45}

	sliceNew := slice[1:3]
	fmt.Printf("%v \n\r", sliceNew)
	fmt.Printf("%v \n\r", slice)

	source := []string{"apple", "orange", "plum", "banana", "grape"}

	slice2 := source[2:3:3]

	fmt.Printf("%v \n\r", slice2)

	fmt.Printf("avg %f \n", average([]float64{1, 2, 3, 4, 5}))
	var s stack
	s.push(20)
	s.push(40)

	fmt.Printf("%v\n", s)

	fmt.Printf("%v", fibonacci(10))
}

func fibonacci(value int) []int {
	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x
}

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return
	}

	s.data[s.i] = k
	s.i++
}

func average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))

	}
	return
}
