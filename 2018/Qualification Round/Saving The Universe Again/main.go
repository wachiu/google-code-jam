package main

import (
	"fmt"
	"strings"
)

func main() {
	var cases int
	fmt.Scanf("%d", &cases)

	type Input struct {
		Defence int64
		Pattern string
	}

	input := make([]Input, cases)
	for i := 0; i < cases; i++ {
		fmt.Scanf("%d %s\n", &input[i].Defence, &input[i].Pattern)
	}

	for i := 0; i < len(input); i++ {
		damage := count(input[i].Pattern)
		p := input[i].Pattern
		y := 0
		for damage > input[i].Defence {
			p = hack(p)
			if damage <= count(p) {
				fmt.Printf("Case #%d: IMPOSSIBLE\n", i+1)
				break
			}
			damage = count(p)
			y++
		}
		if damage <= input[i].Defence {
			fmt.Printf("Case #%d: %d\n", i+1, y)
		}
	}

}

func count(pattern string) int64 {
	var ret int64
	str := int64(1)
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == []byte("C")[0] {
			str *= 2
		} else if pattern[i] == []byte("S")[0] {
			ret += str
		}
	}

	return ret
}

func hack(pattern string) string {
	p := []byte(pattern)
	CS := strings.LastIndex(pattern, "CS")
	if CS != -1 {
		temp := p[CS]
		p[CS] = p[CS+1]
		p[CS+1] = temp
	}
	return string(p)
}
