package main

import (
	"fmt"
	"os"
)

func main() {
	var cases int
	fmt.Scanf("%d", &cases)

	for i := 0; i < cases; i++ {
		orchard := generateOrchard()
		var count int
		var a int
		fmt.Scanf("%d", &a)
		for loop(a, orchard) && count < 1000 {
			count++
		}
	}
}

func loop(a int, orchard [][]bool) bool {
	//#1 examine orchard
	x, y := think(a, orchard)
	//#2 send target
	fmt.Printf("%d %d\n", x+1, y+1)
	//#3 receive actual digged
	var xdash, ydash int
	fmt.Scanf("%d %d", &xdash, &ydash)
	//#4 react to response
	//end
	if xdash == 0 && ydash == 0 {
		return false
	}
	//error
	if xdash == -1 && ydash == -1 {
		os.Exit(1)
	}
	//update orchard
	orchard[xdash-1][ydash-1] = true

	return true
}

func generateOrchard() [][]bool {
	ret := make([][]bool, 1000)
	for i := range ret {
		ret[i] = make([]bool, 1000)
	}
	return ret
}

func think(a int, orchard [][]bool) (int, int) {
	//max = 999, min =2
	var x, y, max int
	if a == 20 {
		for i := 1; i < 5; i++ {
			for j := 1; j < 6; j++ {
				if cur := checkEmptySpace(i, j, orchard); cur > max {
					max = cur
					x = i
					y = j
				}
			}
		}
	} else if a == 200 {
		for i := 1; i < 11; i++ {
			for j := 1; j < 21; j++ {
				if cur := checkEmptySpace(i, j, orchard); cur > max {
					max = cur
					x = i
					y = j
				}
			}
		}
	}
	return x, y
}

func checkEmptySpace(x, y int, orchard [][]bool) int {
	var count int
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !orchard[i][j] {
				count++
			}
		}
	}
	return count
}
