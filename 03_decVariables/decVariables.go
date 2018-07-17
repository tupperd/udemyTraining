package main

import "fmt"

const p = "hello"

func main() {
	fmt.Println(p)
	const fail = "10"
	fmt.Println(fail)
}
