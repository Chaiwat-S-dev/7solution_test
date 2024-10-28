package main

import (
	"bufio"
	"fmt"
	"os"
)

func decode(text string) []int {
	n := len(text)
	seq := make([]int, n+1)

	seq[0] = 0

	for i := 0; i < n; i++ {
		switch text[i] {
		case 'L':
			seq[i+1] = seq[i] - 1
		case 'R':
			seq[i+1] = seq[i] + 1
		case '=':
			seq[i+1] = seq[i]
		}
	}

	minVal := 0
	for _, val := range seq {
		if val < minVal {
			minVal = val
		}
	}

	for i := range seq {
		seq[i] -= minVal
	}
	return seq
}

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the encoded sequence: ")
	scanner.Scan()
	text := scanner.Text()

	sequence := decode(text)
	minSum := sum(sequence)

	fmt.Println("Decoded sequence:", sequence)
	fmt.Println("Minimum sum:", minSum)
}
