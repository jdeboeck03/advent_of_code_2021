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

func process_input(input []string) [][]int {
	heightmap := make([][]int, len(input))

	//split input
	for i, input_line := range input {
		heights := strings.Split(input_line, "")
		for _, height := range heights {
			height_int, _ := strconv.Atoi(height)
			heightmap[i] = append(heightmap[i], height_int)
		}
	}
	return heightmap
}

func calculate_risk_level(heightmap [][]int) (int, map[int][]int) {
	risk_level_sum := 0
	basin_indexes := make(map[int][]int)
	basin_counter := 0
	for i := 0; i < len(heightmap); i++ {
		//fmt.Println(i)
		for j := 0; j < len(heightmap[i]); j++ {
			//fmt.Println(j)
			//Now get the current height value and check the nearby height values
			curr_height := heightmap[i][j]
			if i == 0 {
				//Check two adjacent locations
				if j == 0 {
					//fmt.Println(curr_height)
					// Check low point
					if curr_height < heightmap[i+1][j] && curr_height < heightmap[i][j+1] {
						risk_level_sum += curr_height + 1
						//fmt.Println(curr_height)
						basin_indexes[basin_counter] = []int{i, j}
						basin_counter++
					}

				} else if j == len(heightmap[i])-1 {
					//fmt.Println(curr_height)
					// Check low point
					if curr_height < heightmap[i+1][j] && curr_height < heightmap[i][j-1] {
						risk_level_sum += curr_height + 1
						//fmt.Println(curr_height)
						basin_indexes[basin_counter] = []int{i, j}
						basin_counter++
					}
				} else {
					//Three adjacent locations
					//Check low point
					if curr_height < heightmap[i+1][j] && curr_height < heightmap[i][j-1] && curr_height < heightmap[i][j+1] {
						risk_level_sum += curr_height + 1
						//fmt.Println(curr_height)
						basin_indexes[basin_counter] = []int{i, j}
						basin_counter++
					}
				}
			} else if i == len(heightmap)-1 {
				//Check two adjacent locations
				if j == 0 {
					//fmt.Println(curr_height)
					// Check low point
					if curr_height < heightmap[i-1][j] && curr_height < heightmap[i][j+1] {
						risk_level_sum += curr_height + 1
						//fmt.Println(curr_height)
						basin_indexes[basin_counter] = []int{i, j}
						basin_counter++
					}

				} else if j == len(heightmap[i])-1 {
					//fmt.Println(curr_height)
					// Check low point
					if curr_height < heightmap[i-1][j] && curr_height < heightmap[i][j-1] {
						risk_level_sum += curr_height + 1
						//fmt.Println(curr_height)
						basin_indexes[basin_counter] = []int{i, j}
						basin_counter++
					}
				} else {
					//Three adjacent locations
					//Check low point
					if curr_height < heightmap[i-1][j] && curr_height < heightmap[i][j-1] && curr_height < heightmap[i][j+1] {
						risk_level_sum += curr_height + 1
						//fmt.Println(curr_height)
						basin_indexes[basin_counter] = []int{i, j}
						basin_counter++
					}
				}
			} else if j == 0 {
				// Check 3 adjacent locations
				if curr_height < heightmap[i-1][j] && curr_height < heightmap[i+1][j] && curr_height < heightmap[i][j+1] {
					risk_level_sum += curr_height + 1
					//fmt.Println(curr_height)
					basin_indexes[basin_counter] = []int{i, j}
					basin_counter++
				}
			} else if j == len(heightmap[i])-1 {
				// Check 3 adjacent locations
				if curr_height < heightmap[i-1][j] && curr_height < heightmap[i+1][j] && curr_height < heightmap[i][j-1] {
					risk_level_sum += curr_height + 1
					//fmt.Println(curr_height)
					basin_indexes[basin_counter] = []int{i, j}
					basin_counter++
				}
			} else {
				//fmt.Println(curr_height)
				//Check 4 adjacent locations
				// Check 3 adjacent locations
				if curr_height < heightmap[i-1][j] && curr_height < heightmap[i+1][j] && curr_height < heightmap[i][j-1] && curr_height < heightmap[i][j+1] {
					risk_level_sum += curr_height + 1
					//fmt.Println(curr_height)
					basin_indexes[basin_counter] = []int{i, j}
					basin_counter++
				}
			}
		}
	}
	return risk_level_sum, basin_indexes
}

func calculate_basin_size(heightmap [][]int, basin_indexes map[int][]int) int {
	total_basin_size := 0
	basin_sizes := []int{}
	highest_basin_sizes := []int{0, 0, 0}
	//Have to loop through all values in basin_indexes map and find the basin size
	for _, basin_index := range basin_indexes {
		fmt.Println(basin_sizes)
		fmt.Println(highest_basin_sizes)
		searching := true
		var height = basin_index[0]
		var width = basin_index[1]
		//Searching top
		fmt.Println(heightmap[height][width])
		basin_size := heightmap[height][width]
		//fmt.Println(height)
		for searching {
			// Check all neighbours
			if height != 0 {
				left_neighbour := heightmap[height+1][width]
			}
		}

		for searching_left {
			//First go up and down on current point, then go left and repeat
			for searching_top {
				if height != 0 {
					height--
					//fmt.Println(height)
					if heightmap[height][width] != 9 {
						fmt.Println(heightmap[height][width])
						basin_size += heightmap[height][width]
					} else {
						searching_top = false
						height = basin_index[0]
					}
				} else {
					searching_top = false
					height = basin_index[0]
				}
			}
			if width != 0 {
				width--
				//fmt.Println(height)
				if heightmap[height][width] != 9 {
					//fmt.Println(heightmap[height][width])
					basin_size += heightmap[height][width]
				} else {
					searching_left = false
					width = basin_index[1]
				}
			} else {
				searching_left = false
				width = basin_index[1]
			}
		}
	}
	return total_basin_size
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	fmt.Println("Solution 1:")
	heightmap := process_input(input)
	fmt.Println(heightmap)
	risk_level, basin_indexes := calculate_risk_level(heightmap)
	fmt.Printf("Risk Level: %d\n", risk_level)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	fmt.Println("Solution 2:")
	fmt.Println(basin_indexes)
	basin_size := calculate_basin_size(heightmap, basin_indexes)
	fmt.Printf("Total basin size: %d\n", basin_size)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
