package main

// 36. Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/

/*
Thinking process:
1.Brute Force : check each row, column and 3*3 box one by one
	- use 3 hash set to store the number in each row, column and box
	- if the number already exists in the hash set, return false
	- TC:O(n^2) SC:O(n)
2.Hash Set : Instead of checking each row, column and box seperately, we can validate the entire board in one pass.
	- use 3 hash set to store the number in each row, column and box
	- for each number, check if it already exists in the corresponding row, column and box hash set
	- if it already exists, return false
	- if not, add the number to the corresponding hash set
	- TC:O(n^2) SC:O(n)

*/

// brute force
func isValidSudoku_1(board [][]byte) bool {
	// row
	for i := 0; i < 9; i++ {
		set := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if set[board[i][j]] {
					return false
				}
				set[board[i][j]] = true
			}
		}
	}

	// column
	for j := 0; j < 9; j++ {
		set := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if set[board[i][j]] {
					return false
				}
				set[board[i][j]] = true
			}
		}
	}

	// box
	for square := 0; square < 9; square++ {
		seen := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				row := (square/3)*3 + i
				col := (square%3)*3 + j
				if board[row][col] != '.' {
					if seen[board[row][col]] {
						return false
					}
					seen[board[row][col]] = true
				}
			}
		}
	}
	return true
}

// hash set
func isValidSudoku_2(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	square := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		square[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			squareIndex := (i/3)*3 + i/3

			if board[i][j] != '.' {
				if rows[i][ch] || cols[j][ch] || square[squareIndex][ch] {
					return false
				}
				rows[i][ch] = true
				cols[j][ch] = true
				square[squareIndex][ch] = true
			}

		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '3', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	println(isValidSudoku_2(board))
}
