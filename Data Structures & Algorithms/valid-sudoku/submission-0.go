func isValidSudoku(board [][]byte) bool {
	rows := [9]map[byte]bool{}
	cols := [9]map[byte]bool{}
	boxes := [9]map[byte]bool{}

	for i := 0; i < 9; i++ {
		rows[i] = map[byte]bool{}
		cols[i] = map[byte]bool{}
		boxes[i] = map[byte]bool{}
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cell := board[row][col]

			if cell == '.' {
				continue
			}

			boxIndex := (row/3)*3 + col/3

			if rows[row][cell] || cols[col][cell] || boxes[boxIndex][cell] {
				return false
			}

			rows[row][cell] = true
			cols[col][cell] = true
			boxes[boxIndex][cell] = true
		}
	}

	return true
}
