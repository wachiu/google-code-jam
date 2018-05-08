package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	cases, _ := strconv.Atoi(text[:len(text)-1])

	type Input struct {
		cakes []bool
		k     int
	}

	input := make([]Input, cases)
	for i := 0; i < cases; i++ {
		line, _ := reader.ReadString('\n')
		split := strings.Split(line[:len(line)-1], " ")
		k, _ := strconv.Atoi(split[1])
		var cakes []bool
		for j := 0; j < len(split[0]); j++ {
			if string(split[0][j]) == "+" {
				cakes = append(cakes, true)
			} else if string(split[0][j]) == "-" {
				cakes = append(cakes, false)
			}
		}
		input[i] = Input{cakes: cakes, k: k}
	}

	for i := range input {
		count := 0
		impossible := false

		j := 0
		for j < len(input[i].cakes) {
			if input[i].cakes[j] {
				j++
			} else {
				end := j + input[i].k
				if end > len(input[i].cakes) {
					impossible = true
					break
				}
				for k := j; k < end; k++ {
					input[i].cakes[k] = !input[i].cakes[k]
				}
				j++
				count++
			}
		}

		if !impossible {
			fmt.Printf("Case #%d: %d\n", i+1, count)
		} else {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i+1)
		}
	}
}
