package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

func main() {
	fmt.Println("Ready? Steady? GO!")

	argc := len(os.Args)
	if argc < 3 {
		fmt.Printf("Usage: %s <n1> <n2> <nn>...\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	nums := make([]float64, 0, argc-1)

	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", arg)
			os.Exit(1)
		}
		nums = append(nums, n)
	}

	sum := 0.0
	for _, n := range nums {
		sum += n
	}

	sub := nums[0]
	for _, n := range nums {
		sub -= n
	}

	mul := 1.0
	for _, n := range nums {
		mul *= n
	}

	div := nums[0]
	for _, n := range nums[1:] {
		if n == 0 {
			fmt.Println("Cannot divide by zero!")
			os.Exit(1)
		}
		div /= n
	}

	fmt.Printf("%-5s = %5.2f\n", "sum", sum)
	fmt.Printf("%-5s = %5.2f\n", "sub", sub)
	fmt.Printf("%-5s = %10.2f\n", "mul", mul)
	fmt.Printf("%-5s = %2.2f\n", "div", div)
	os.Exit(0)
}
