package main

import (
	"fmt"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	var arr = [][]byte{{
		'8', '3', '.', '.', '7', '.', '.', '.', '.'}, {
		'8', '.', '.', '1', '9', '5', '.', '.', '.'}, {
		'.', '9', '1', '.', '.', '.', '.', '6', '.'}, {
		'1', '.', '.', '.', '6', '.', '.', '.', '3'}, {
		'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {
		'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {
		'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {
		'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {
		'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(IsValidSudoku(arr))
}
