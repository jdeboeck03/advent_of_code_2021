package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// loads input of file into a slice of integer variables
func loadInput() ([]string, error) {
	input := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		input = append(input, line)
	}
	return input, nil
}

func remove_index(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func contains(s []int, value int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func remove_white_spaces(input []string) []string {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == "" {
			input = remove_index(input, i)
		}
	}
	return input
}

func process_bingo_boards(input []string) [][][]string {
	// number of bingo boards
	// fmt.Println(len(input))
	input = remove_white_spaces(input)
	board_width := 5
	board_height := 5
	number_of_boards := len(input) / board_height
	var bingo_boards = make([][][]string, number_of_boards)
	// init board
	for i := range bingo_boards {
		// height of bingo_board
		bingo_boards[i] = make([][]string, board_height)
		for j := range bingo_boards[i] {
			// width of bingo_board
			bingo_boards[i][j] = make([]string, board_width)
		}
	}
	// fill up board
	board_index := 0
	board_height_index := 0
	for i := 0; i < len(input); i++ {
		input_split := strings.Split(input[i], " ")
		input_split = remove_white_spaces(input_split)
		for board_width_index, value := range input_split {
			bingo_boards[board_index][board_height_index][board_width_index] = value
		}
		board_height_index++
		if board_height_index >= board_height {
			board_height_index = 0
			board_index++
		}
	}
	fmt.Println(bingo_boards)
	return bingo_boards
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	bingo_draws := input[0]
	input = remove_index(input, 0)
	fmt.Println("Solution 1:")
	fmt.Println(bingo_draws)
	bingo_boards := process_bingo_boards(input)
	fmt.Println(bingo_boards)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	fmt.Println("Solution 2:")
	fmt.Printf("Time: %s\n", time.Since(start2))
}
