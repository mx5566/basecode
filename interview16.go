package main

func IsValidSudoku(board [][]byte) bool {
	row := [9]byte{}
	col := [9]byte{}
	artxi := [9]byte{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			temp := board[i][j] - '0'

			cur := byte(1 << (temp - 1))
			artx := i/3*3 + j/3

			if (row[i]&cur)|(col[j]&cur)|(artxi[artx]&cur) != 0 {
				return false
			}

			row[i] |= cur
			col[j] |= cur
			artxi[artx] |= cur
		}
	}

	return true
	/*
		var temp = func(nums []byte) bool {
			var maps map[byte]int

			maps = map[byte]int{}

			for _, value := range nums {
				if value == '.' {
					continue
				}
				maps[value]++

				if maps[value] >= 2 {
					return true
				}
			}

			return false
		}

		// 判断行
		for _, oneArray := range board {
			var retBool = temp(oneArray)
			if retBool == true {
				return !retBool
			}
		}

		// 判断列
		var length = len(board)
		for i := 0; i < length; i++ {
			var arr []byte
			for j := 0; j < length; j++ {
				arr = append(arr, board[j][i])
			}

			var retBool = temp(arr)
			if retBool == true {
				return !retBool
			}
		}

		// 宫判断
		for i := 0; i < length; i += 3 {
			for j := 0; j < length; j += 3 {
				var arr []byte
				arr = append(arr, board[i][j:j+3]...)
				arr = append(arr, board[i+1][j:j+3]...)
				arr = append(arr, board[i+2][j:j+3]...)

				var retBool = temp(arr)
				if retBool == true {
					return !retBool
				}
			}
		}

		return true
	*/
}
