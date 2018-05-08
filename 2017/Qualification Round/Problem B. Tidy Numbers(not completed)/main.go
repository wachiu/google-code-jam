package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	cases, _ := strconv.Atoi(text[:len(text)-1])

	input := make([]uint64, cases)

	for i := range input {
		line, _ := reader.ReadString('\n')
		n, _ := strconv.ParseUint(line[:len(line)-1], 10, 64)
		input[i] = n
	}

}
