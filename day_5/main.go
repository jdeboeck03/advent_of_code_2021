package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	return bingo_boards
}

func drawing(bingo_boards [][][]string, bingo_draws string) int {
	solution := 0
	winning_board := 0
	winning_number := 0
	// fmt.Println(bingo_boards)
	// fmt.Println(bingo_draws)
	bingo_draws_split := strings.Split(bingo_draws, ",")
	// fmt.Println(bingo_draws_split)
	// Drawing
	fmt.Println(bingo_boards)
	quit := false
	number_of_boards_left := len(bingo_boards)
	fmt.Println(number_of_boards_left)
	for _, bingo_draw := range bingo_draws_split {
		// Check each board for a possible match on bingo board
		// fmt.Println(bingo_draw)
		for i, board := range bingo_boards {
			for j, board_line := range board {
				for k, board_piece := range board_line {
					//fmt.Println(board_piece)
					if board_piece == bingo_draw {
						// fmt.Println("Het is van dat, we hebben een hit")
						// fmt.Println(board_piece)
						// fmt.Println(bingo_draw)
						bingo_boards[i][j][k] = "x"

						// Check if bingo board is complete
						// Loop through the corresponding row

						row_checking := true
						bingo := false
						for l, check_piece := range board_line {
							//fmt.Println(check_piece)
							if check_piece == "x" {
								if l == len(board_line)-1 {
									// BINGO FIND BINGO BOARD
									// fmt.Println("BINGO")
									row_checking = false
									bingo = true
									number_of_boards_left -= 1
									fmt.Println(number_of_boards_left)
									if number_of_boards_left == 0 {
										winning_board = i
										winning_number, _ = strconv.Atoi(bingo_draw)
										quit = true
									}
								}
							} else {
								row_checking = false
							}
							if !row_checking {
								break
							}
						}

						if quit {
							break
						}

						// Loop through the corresponding column
						// Double loop to check every column
						// Loop through the corresponding column
						// Double loop to check every column
						column_checking := true
						for m, check_board_line := range board {
							//fmt.Println(check_board_line[k])
							if check_board_line[k] == "x" {
								if m == len(check_board_line)-1 {
									// BINGO FIND BINGO BOARD
									// fmt.Println("BINGO")
									//fmt.Println(bingo_boards)
									column_checking = false
									if !bingo {
										number_of_boards_left -= 1
										fmt.Println(number_of_boards_left)
									}
									if number_of_boards_left == 0 {
										winning_board = i
										winning_number, _ = strconv.Atoi(bingo_draw)
										quit = true
									}
								}
							} else {
								column_checking = false
							}
							if !column_checking {
								break
							}
						}
					}
					if quit {
						break
					}
				}
				if quit {
					break
				}
			}
			if quit {
				break
			}
		}
		if quit {
			break
		}
	}
	fmt.Println(bingo_boards)
	fmt.Println(winning_number)
	fmt.Println(winning_board)
	sum := 0
	// Count score!
	for _, winning_board_line := range bingo_boards[winning_board] {
		for _, winning_board_piece := range winning_board_line {
			if winning_board_piece != "x" {
				winning_board_piece_int, err := strconv.Atoi(winning_board_piece)
				if err != nil {
					log.Fatalf("Program failed: %s", err)
				}
				sum += winning_board_piece_int
			}
		}
	}
	fmt.Println(sum)
	solution = sum * winning_number
	return solution
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
	// fmt.Println(bingo_draws)
	bingo_boards := process_bingo_boards(input)
	// fmt.Println(bingo_boards)
	solution := drawing(bingo_boards, bingo_draws)
	fmt.Printf("Solution: %d\n", solution)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	fmt.Println("Solution 2:")
	fmt.Printf("Time: %s\n", time.Since(start2))
}
