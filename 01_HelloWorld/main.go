package main

import "fmt"
import "math"

func main() {
	fmt.Println("Hello World!")
	for i := 0.0; i < 10; i++ {
		fmt.Println(math.Pow(i, 2))
	}
}
