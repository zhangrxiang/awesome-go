package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(_ int, _ int) int  { return 0 }

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	fmt.Println(hypot(3, 4))
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"
}
