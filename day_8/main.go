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

type numbers struct {
	numbers []int
	outputs [][]string
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

func find_index(input []string, value string) int {
	for p, v := range input {
		if v == value {
			return p
		}
	}
	return -1
}

func find_index_2(input [][]string, value []string) int {
	for p, v := range input {
		if equal(v, value) {
			return p
		}
	}
	return -1
}

func contains(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
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
		signal_patterns[i] = make([][]string, 10)
	}
	output_values := make([][][]string, len(input))
	for i := range output_values {
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

func processed_signal_patterns(signal_patterns_2 [][][]string) []numbers {
	processed_signal_patterns := make([]numbers, len(signal_patterns_2[0]))
	for i := range processed_signal_patterns {
		//fmt.Println(i)
		//Create slice for numbers
		processed_signal_patterns[i].numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		processed_signal_patterns[i].outputs = make([][]string, 10)
	}
	signal_wires := make(map[string]string)
	for i, signal_pattern := range signal_patterns_2 {
		//Find all numbers!
		for _, signal := range signal_pattern {
			//Find 1, 4, 7, 8
			if len(signal) == 2 {
				processed_signal_patterns[i].outputs[1] = signal
			} else if len(signal) == 3 {
				processed_signal_patterns[i].outputs[7] = signal
			} else if len(signal) == 4 {
				processed_signal_patterns[i].outputs[4] = signal
			} else if len(signal) == 7 {
				processed_signal_patterns[i].outputs[8] = signal
			}
		}
		for _, signal := range signal_pattern {
			//Find 6 (missing number of 1)
			if len(signal) == 6 {
				//fmt.Println(signal)
				for _, connection := range processed_signal_patterns[i].outputs[1] {
					//Find which element of 1 is missing
					present := contains(signal, connection)
					if !present {
						//fmt.Println(signal)
						//fmt.Println(connection)
						signal_wires["c"] = connection
						connection_index := find_index(processed_signal_patterns[i].outputs[1], connection)
						if connection_index == 0 {
							signal_wires["f"] = processed_signal_patterns[i].outputs[1][1]
						} else {
							signal_wires["f"] = processed_signal_patterns[i].outputs[1][0]
						}
						processed_signal_patterns[i].outputs[6] = signal
					}
				}
			}
		}
		for _, signal := range signal_pattern {
			//Find 5 (also misses c)
			if len(signal) == 5 {
				present := contains(signal, signal_wires["c"])
				//Check if c is present in numbers (if no we found 5)
				if !present {
					processed_signal_patterns[i].outputs[5] = signal
					//find e (only missing letter with 5 & 6)
					for _, connection := range processed_signal_patterns[i].outputs[6] {
						e_present := contains(signal, connection)
						//fmt.Println(e_present)
						if !e_present {
							signal_wires["e"] = connection
							fmt.Println(signal_wires)
						}
					}
				}
			}
		}
		for _, signal := range signal_pattern {
			if len(signal) == 6 {
				present := contains(signal, signal_wires["e"])
				if !present {
					processed_signal_patterns[i].outputs[9] = signal
				}
			}
		}
		for _, signal := range signal_pattern {
			if len(signal) == 6 {
				if !equal(signal, processed_signal_patterns[i].outputs[9]) && !equal(signal, processed_signal_patterns[i].outputs[6]) {
					processed_signal_patterns[i].outputs[0] = signal
				}
			}
		}
		for _, signal := range signal_pattern {
			if len(signal) == 5 {
				present := contains(signal, signal_wires["f"])
				if !present {
					processed_signal_patterns[i].outputs[2] = signal
				}
			}
		}
		for _, signal := range signal_pattern {
			if len(signal) == 5 {
				if !equal(signal, processed_signal_patterns[i].outputs[2]) && !equal(signal, processed_signal_patterns[i].outputs[5]) {
					processed_signal_patterns[i].outputs[3] = signal
				}
			}
		}
	}
	return processed_signal_patterns
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
	processed_signal_patterns := processed_signal_patterns(signal_patterns_2)
	fmt.Println(processed_signal_patterns)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
