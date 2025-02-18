package main

import (
	"bufio"
	"fmt"
	"linearstats/internal"
	"os"
	"strconv"
)

func main() {
	var yValue int
	var xValues, yValues []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		yValue, _ = strconv.Atoi(line)
		xValues = append(xValues, len(xValues))
		yValues = append(yValues, yValue)
	}

	intercept, slope := internal.CalculateLinearRegression(xValues, yValues)
	pcc := internal.CalculatePCC(xValues, yValues)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pcc)
}
