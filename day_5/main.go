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

func create_empty_matrix(height int, width int) [][]int {

	var vents = make([][]int, height)
	for i := range vents {
		// height of bingo_board
		vents[i] = make([]int, width)
	}

	//fmt.Println(len(vents))
	//fmt.Println(len(vents[0]))
	return vents
}

func process_lines(input []string) [][]int {
	var processed_lines = make([][]int, len(input))
	for j := range processed_lines {
		// width of bingo_board
		processed_lines[j] = make([]int, 4)
	}
	for i, line := range input {
		line_split := strings.Split(line, ",")
		// Put first and last coordinate in processed lines
		processed_lines[i][0], _ = strconv.Atoi(line_split[0])
		processed_lines[i][3], _ = strconv.Atoi(line_split[2])
		//split the last two coordinates even further
		line_split_2 := strings.Split(line_split[1], " -> ")
		processed_lines[i][1], _ = strconv.Atoi(line_split_2[0])
		processed_lines[i][2], _ = strconv.Atoi(line_split_2[1])
	}
	return processed_lines
}

func find_vert_horiz_lines(processed_lines [][]int) [][]int {
	vert_horiz_lines := [][]int{}
	for _, line := range processed_lines {
		if (line[0] == line[2]) || (line[1] == line[3]) {
			vert_horiz_lines = append(vert_horiz_lines, line)
		}
	}

	return vert_horiz_lines
}

func find_diagonal_lines(processed_lines [][]int) [][]int {
	diagonal_lines := [][]int{}
	for _, line := range processed_lines {
		if (line[0] == line[2]) || (line[1] == line[3]) {
		} else {
			diagonal_lines = append(diagonal_lines, line)
		}
	}
	return diagonal_lines
}

func fill_vents(vents [][]int, vert_horiz_lines [][]int) [][]int {

	for _, line := range vert_horiz_lines {
		//Check for vertical vents
		if line[0] == line[2] {
			width := line[0]
			if line[1] < line[3] {
				for i := line[1]; i <= line[3]; i++ {
					vents[i][width]++
				}
			} else if line[3] < line[1] {
				for i := line[3]; i <= line[1]; i++ {
					vents[i][width]++
				}
			}
		} else if line[1] == line[3] {
			height := line[1]
			if line[0] < line[2] {
				for i := line[0]; i <= line[2]; i++ {
					vents[height][i]++
				}
			} else if line[2] < line[0] {
				for i := line[2]; i <= line[0]; i++ {
					vents[height][i]++
				}
			}
		}
	}
	return vents
}

func fill_diagonal_vents(vents [][]int, vert_horiz_lines [][]int) [][]int {
	for _, line := range vert_horiz_lines {
		if line[0] > line[2] && line[1] < line[3] {
			//Start from first coordinate
			height := line[1]
			for i := line[0]; i >= line[2]; i-- {
				vents[height][i]++
				height++
			}
		} else if line[0] > line[2] && line[1] > line[3] {
			//Start from first coordinate
			height := line[1]
			for i := line[0]; i >= line[2]; i-- {
				vents[height][i]++
				height--
			}
		} else if line[0] < line[2] && line[1] < line[3] {
			//Start from first coordinate
			height := line[1]
			for i := line[0]; i <= line[2]; i++ {
				vents[height][i]++
				height++
			}
		} else if line[0] < line[2] && line[1] > line[3] {
			//Start from first coordinate
			height := line[1]
			//fmt.Println(line)
			//fmt.Println(width)
			//fmt.Println(height)
			for i := line[0]; i <= line[2]; i++ {
				//fmt.Println(i)
				//fmt.Println(height)
				vents[height][i]++
				height--
			}
			//fmt.Println("Processed")
			//fmt.Println(vents)
		}
	}
	return vents
}
func find_points(filled_vents [][]int) int {
	number_of_points := 0
	for _, line := range filled_vents {
		for _, point := range line {
			if point >= 2 {
				number_of_points++
			}
		}
	}
	return number_of_points
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	fmt.Println("Solution 1:")
	processed_lines := process_lines(input)
	vert_horiz_lines := find_vert_horiz_lines(processed_lines)
	fmt.Println(vert_horiz_lines)
	height := 991
	width := 991
	vents := create_empty_matrix(height, width)
	filled_vents := fill_vents(vents, vert_horiz_lines)
	fmt.Println(filled_vents)
	number_of_points := find_points(filled_vents)
	fmt.Printf("number of points: %d\n", number_of_points)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	fmt.Println("Solution 2:")
	diagonal_lines := find_diagonal_lines(processed_lines)
	filled_diagonal_vents := fill_diagonal_vents(filled_vents, diagonal_lines)
	number_of_diagonal_points := find_points(filled_diagonal_vents)
	fmt.Println(filled_diagonal_vents)
	fmt.Printf("number of points: %d\n", number_of_diagonal_points)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
