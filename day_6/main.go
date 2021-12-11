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

func process_input(input []string) []int {
	var processed_lines = []int{}

	//split input
	input_split := strings.Split(input[0], ",")
	for _, number := range input_split {
		// width of bingo_board
		number_int, _ := strconv.Atoi(number)
		processed_lines = append(processed_lines, number_int)
	}
	return processed_lines
}

func process_input_2(input []string) []int {
	var processed_lines = make([]int, 9)

	//split input
	input_split := strings.Split(input[0], ",")
	for _, number := range input_split {
		// width of bingo_board
		number_int, _ := strconv.Atoi(number)
		processed_lines[number_int]++
	}
	return processed_lines
}

func create_fishes(input []int) int {
	days := 80
	fishes_cnt := 0
	var fishes = input
	for i := 0; i < days; i++ {
		for j, fish := range fishes {
			if fish == 0 {
				fish = 6
				fishes = append(fishes, 8)
			} else {
				fish--
			}
			fishes[j] = fish
		}
	}
	fishes_cnt = len(fishes)
	return fishes_cnt
}

func create_fishes_2(input []int) int {
	days := 256
	fishes_cnt := 0
	var fishes = input
	next_fish := fishes[8]
	for i := 0; i < days; i++ {
		for j := 8; j >= 0; j-- {
			if j == 0 {
				fishes[8] = next_fish
				fishes[6] += next_fish
			} else {
				fish_backup := fishes[j-1]
				fishes[j-1] = next_fish
				next_fish = fish_backup
			}
		}
		//fmt.Println(fishes)
	}
	for _, fish := range fishes {
		fishes_cnt += fish
	}
	return fishes_cnt
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	var input_2 = input
	start1 := time.Now()
	fmt.Println("Solution 1:")
	processed_input := process_input(input)
	fishes_cnt := create_fishes(processed_input)
	fmt.Println(fishes_cnt)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	fmt.Println("Solution 2:")
	processed_input_2 := process_input_2(input_2)
	fishes_cnt_2 := create_fishes_2(processed_input_2)
	fmt.Println(fishes_cnt_2)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
