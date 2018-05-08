package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	type Input struct {
		no    uint64
		value []uint64
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	cases, _ := strconv.Atoi(strings.Trim(text, "\n"))
	for i := 0; i < cases; i++ {
		var input Input
		no, _ := reader.ReadString('\n')
		input.no, _ = strconv.ParseUint(strings.Trim(no, "\n"), 10, 64)
		input.value = make([]uint64, input.no)
		values, _ := reader.ReadString('\n')
		vs := strings.Split(strings.Trim(values, "\n"), " ")
		for j := range input.value {
			input.value[j], _ = strconv.ParseUint(vs[j], 10, 64)
		}

		// bubblesort(input.value)
		troublesort(input.value)
		var wrong bool
		for j := 0; j < len(input.value)-1; j++ {
			if input.value[j] > input.value[j+1] {
				wrong = true
				fmt.Printf("Case #%d: %d\n", i+1, j)
				break
			}
		}
		if !wrong {
			fmt.Printf("Case #%d: OK\n", i+1)
		}
	}
}

//bubblesort
func bubblesort(L []uint64) {
	var done bool
	for !done {
		done = true
		for i := 0; i < len(L)-2; i++ {
			if L[i] > L[i+2] {
				done = false
				reverse(L, i, i+2)
			}
		}
	}
}

func reverse(L []uint64, start, end int) {
	temp := L[start]
	L[start] = L[end]
	L[end] = temp
}

//troublesort
func troublesort(L []uint64) {
	even, odd := splitOddEvenSlice(L)
	sort.Slice(even, func(i, j int) bool { return even[i] < even[j] })
	sort.Slice(odd, func(i, j int) bool { return odd[i] < odd[j] })
	for i := range L {
		if i%2 == 0 {
			L[i] = odd[i/2]
		} else {
			L[i] = even[i/2]
		}
	}
}

func splitOddEvenSlice(L []uint64) (even, odd []uint64) {
	var o, e int
	if len(L)%2 == 0 {
		o = len(L) / 2
		e = len(L) / 2
	} else {
		o = len(L)/2 + 1
		e = len(L)
	}
	even = make([]uint64, e)
	odd = make([]uint64, o)
	for i, j := 0, 0; i < len(L); i += 2 {
		odd[j] = L[i]
		j++
	}
	for i, j := 1, 0; i < len(L); i += 2 {
		odd[j] = L[i]
		j++
	}
	return
}
