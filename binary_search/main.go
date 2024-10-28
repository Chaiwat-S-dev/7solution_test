package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func loadData(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var triangle [][]int
	err = json.Unmarshal(data, &triangle)
	if err != nil {
		return nil, err
	}

	return triangle, nil
}

func maxPathSumTopDown(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	dp[0][0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = max(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	maxSum := dp[len(triangle)-1][0]
	for _, value := range dp[len(triangle)-1] {
		if value > maxSum {
			maxSum = value
		}
	}
	return maxSum
}

func main() {
	triangle, err := loadData("hard.json")
	if err != nil {
		log.Fatalf("Failed to load file: %v", err)
	}

	result := maxPathSumTopDown(triangle)
	fmt.Println("Result max:", result)
}
