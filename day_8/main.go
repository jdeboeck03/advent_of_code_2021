package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type number struct {
	number int
	output string
}

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

func process_input(input []string) ([]string, []string) {
	signal_patterns := []string{}
	output_values := []string{}

	//split input
	for _, input_line := range input {
		input_line_split := strings.Split(input_line, " | ")
		signal_patterns = append(signal_patterns, input_line_split[0])
		output_values = append(output_values, input_line_split[1])
	}
	return signal_patterns, output_values
}

func process_input_2(input []string) ([][][]string, [][][]string) {
	signal_patterns := make([][][]string, len(input))
	for i := range signal_patterns {
		// height of bingo_board
		signal_patterns[i] = make([][]string, 10)
	}
	output_values := make([][][]string, len(input))
	for i := range output_values {
		// height of bingo_board
		output_values[i] = make([][]string, 4)
	}

	//split input
	for i, input_line := range input {
		input_line_split := strings.Split(input_line, " | ")
		signal_values := strings.Split(input_line_split[0], " ")
		output_value := strings.Split(input_line_split[1], " ")
		for j, signal_value := range signal_values {
			signal_value_split := strings.Split(signal_value, "")
			sort.Strings(signal_value_split)
			signal_patterns[i][j] = signal_value_split
		}
		for j, output := range output_value {
			output_split := strings.Split(output, "")
			sort.Strings(output_split)
			output_values[i][j] = output_split
		}
	}
	return signal_patterns, output_values
}

func calculate_total_output_values(output_values []string) int {
	total_output_count := 0
	for _, output_line := range output_values {
		outputs := strings.Split(output_line, " ")
		for _, output := range outputs {
			if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
				//fmt.Println(output)
				total_output_count++
			}
		}
		//fmt.Println(outputs)
	}
	return total_output_count
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	fmt.Println("Solution 1:")
	signal_patterns, output_values := process_input(input)
	fmt.Println(signal_patterns)
	//fmt.Println(output_values)
	total_output_values := calculate_total_output_values(output_values)
	fmt.Printf("Total count: %d\n", total_output_values)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	fmt.Println("Solution 2:")
	signal_patterns_2, output_values_2 := process_input_2(input)
	fmt.Println(signal_patterns_2)
	fmt.Println(output_values_2)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
