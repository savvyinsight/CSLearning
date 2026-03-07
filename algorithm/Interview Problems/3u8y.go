package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/*
Thinking process:
1. Read the compressed line and split it into parts.
2. Convert the compressed parts into integers and store them in a slice.
3. Read the target line and split it into parts to get the target row and column.
4. Calculate the index in the uncompressed data using the formula: index = targetRow * C + targetCol, where C is the number of columns.
5. Iterate through the compressed data to find the value at the target index by keeping track of the current index and comparing it with the target index.
6. Print the value at the target position.

Approach:
- Use bufio.Scanner to read input lines.
- Use strings.Fields to split the input lines into parts.
- Use strconv.Atoi to convert string parts to integers.
- Implement a loop to find the value at the target index based on the compressed data.

Ways to solve:
1. Directly iterate through the compressed data and keep track of the current index until we find the target index. TC: O(n), SC: O(1).
2. fully reconstruct the 2D matrix from the compressed data and then directly access the target index. TC: O(n), SC: O(m*n) where m and n are the dimensions of the uncompressed matrix.
*/

func approach1() {
	// This approach directly iterates through the compressed data and keeps track of the current index until it finds the target index.
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	compressedLine := scanner.Text()

	compressedParts := strings.Fields(compressedLine)

	data := make([]int, len(compressedParts))
	for i, s := range compressedParts {
		data[i], _ = strconv.Atoi(s)
	}

	scanner.Scan()
	targetLine := scanner.Text()
	targetParts := strings.Fields(targetLine)
	targetRow, _ := strconv.Atoi(targetParts[0])
	targetCol, _ := strconv.Atoi(targetParts[1])

	C := data[1]
	index := targetRow*C + targetCol
	currindex := 0
	result := 0
	for i := 2; i < len(data); i += 2 {
		value := data[i]
		count := data[i+1]
		if currindex+count > index {
			result = value
			break
		}
		currindex += count
	}
	println("Value at target position:", result)
}

func approach2() {
	// fully reconstruct the 2D matrix
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	compressedLine := scanner.Text()
	compressedParts := strings.Fields(compressedLine)
	data := make([]int, len(compressedParts))
	for i, s := range compressedParts {
		data[i], _ = strconv.Atoi(s)
	}

	scanner.Scan()
	targetLine := scanner.Text()
	targetParts := strings.Fields(targetLine)
	targetRow, _ := strconv.Atoi(targetParts[0])
	targetCol, _ := strconv.Atoi(targetParts[1])

	C := data[1]
	matrix := make([][]int, 0)
	row := make([]int, 0)
	for i := 2; i < len(data); i += 2 {
		value := data[i]
		count := data[i+1]
		for j := 0; j < count; j++ {
			row = append(row, value)
			if len(row) == C {
				matrix = append(matrix, row)
				row = make([]int, 0)
			}
		}
	}

	result := matrix[targetRow][targetCol]
	println("Value at target position:", result)
}
func main() {
	/*
		test cases:
		Input:
		10 10 255 34 0 1 255 8 0 3 255 6 0 5 255 4 0 7 255 2 0 9 255 21
		3 4
		Output: Value at target position: 0
	*/
	// approach1()
	approach2()
}

/*
The closest related LeetCode problems are:
RLE Iterator – Iterate through a run-length encoded sequence.
Flatten 2D Vector – Access elements in a flattened 2D structure.
Design Compressed String Iterator – Iterate over compressed data.
*/
