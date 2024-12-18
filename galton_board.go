package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

const (
	defaultRows  = 9
	defaultBalls = 256
)

func getInput(prompt string, defaultValue int) int {
	var input string
	fmt.Printf("%s [%d]: ", prompt, defaultValue)
	_, err := fmt.Scanln(&input)
	if err != nil || input == "" {
		return defaultValue
	}
	num, err := strconv.Atoi(input)
	if err != nil || num <= 0 {
		return defaultValue
	}
	return num
}

func simulate(rows, balls int) []int {
	dist := make([]int, rows+1)
	for i := 0; i < balls; i++ {
		pos := 0
		for j := 0; j < rows; j++ {
			pos += rand.Intn(2)
		}
		dist[pos]++
	}
	return dist
}

func printDistribution(dist []int) {
	fmt.Println("\nDistribution:")
	for i, count := range dist {
		fmt.Printf("Bin %d: %s\n", i, strings.Repeat("*", count))
	}
}

func calculateStats(dist []int) (float64, float64, float64) {
	total := float64(len(dist))
	mean := 0.0
	for i, count := range dist {
		mean += float64(i) * float64(count) / total
	}

	variance := 0.0
	for i, count := range dist {
		variance += float64(count) * math.Pow(float64(i)-mean, 2) / total
	}

	stdDev := math.Sqrt(variance)
	return mean, variance, stdDev
}

func main() {
	rows := getInput("Enter number of rows", defaultRows)
	balls := getInput("Enter number of balls", defaultBalls)

	distribution := simulate(rows, balls)
	printDistribution(distribution)

	mean, variance, stdDev := calculateStats(distribution)
	fmt.Printf("\nMean: %.2f\nVariance: %.2f\nStandard Deviation: %.2f\n", mean, variance, stdDev)
}
