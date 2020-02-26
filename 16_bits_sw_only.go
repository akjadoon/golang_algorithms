// Author: Asfandyar Jadoon
// DMOJ Problem Link: https://dmoj.ca/problem/16bitswonly

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func parse_line(input string) {
	strings := strings.Split(input, " ")
	nums := make([]int64, len(strings))

	for i, s := range strings {
		nums[i], _ = strconv.ParseInt(s, 10, 64)
	}
	res := nums[0] * nums[1]
	if res == nums[2] {
		fmt.Println("POSSIBLE DOUBLE SIGMA")
	} else {
		fmt.Println("16 BIT S/W ONLY")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	num_lines, err := strconv.Atoi(input[:len(input) - 1])
	if err != nil {
		fmt.Println(err)
	}
	inputs := make([]string, num_lines)

	for i:=0; i<int(num_lines); i++ {
		inputs[i], _ = reader.ReadString('\n')
	}

	for i:=0; i<int(num_lines); i++ {
		parse_line(inputs[i][:len(inputs[i]) - 1])
	}
}