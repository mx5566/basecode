package main

func MartixRotate(matrix [][]int) {
	var length = len(matrix)

	for i := 0; i < length; i++ {
		for j := 0; j < length-i; j++ {
			matrix[i][j], matrix[length-j-1][length-i-1] = matrix[length-j-1][length-i-1], matrix[i][j]
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[j][i], matrix[length-j-1][i] = matrix[length-j-1][i], matrix[j][i]
		}
	}

}
