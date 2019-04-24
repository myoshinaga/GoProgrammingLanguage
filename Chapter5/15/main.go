package main

import (
	"errors"
	"fmt"
)

func main() {
	max, _ := max(1, 2, 3, 4)
	fmt.Printf("max is [%d]\n", max)
	min, _ := min(1, 2, 3, 4)
	fmt.Printf("min is [%d]\n", min)
	fmt.Printf("max is [%d]\n", max2(1, 2, 3, 4))
	fmt.Printf("min is [%d]\n", min2(1, 2, 3, 4))
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("max: no parameter\n")
	}

	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}

	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("min: no parameter\n")
	}

	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}

	return min, nil
}

func max2(num int, vals ...int) int {
	max := num
	for _, val := range vals {
		if val > max {
			max = val
		}
	}

	return max
}

func min2(num int, vals ...int) int {
	min := num
	for _, val := range vals {
		if val < min {
			min = val
		}
	}

	return min
}
