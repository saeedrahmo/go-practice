package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi)

	fmt.Println(add(5, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// what is naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
