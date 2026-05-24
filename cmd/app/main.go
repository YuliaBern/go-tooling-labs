package main

import (
	"fmt"
	"github.com/YuliaBern/lab1-tooling/internal"
)

func main() {
	sum := internal.Add(5, 3)

	fmt.Println("Sum:", sum)

	_, err := fmt.Printf("Hello, World from lab1!\n")
	if err != nil {
		fmt.Println("Помилка друку:", err)
	}
}
