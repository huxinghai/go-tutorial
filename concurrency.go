package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func sum(a []int, c chan int) {
	sum := 0
	for _, value := range a {
		sum += value
	}
	c <- sum
}

func main() {
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out)
	fmt.Println(<-out)

	ci := make(chan int)

	arr := []int{1, 2, 3, 4, 5, 6}

	go sum(arr[:len(arr)/2], ci)
	go sum(arr[len(arr)/2:], ci)

	x, y := <-ci, <-ci
	fmt.Println("x+y=", x+y)
	close(ci)
}
