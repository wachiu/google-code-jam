package main

import "fmt"

func main() {
	var cases int
	fmt.Scanf("%d", &cases)

	for i := 0; i < cases; i++ {
		var R, C, H, V int
		fmt.Scanf("%d %d %d %d", &R, &C, &H, &V)

		waffle := make([][]bool, R)
		for j := 0; j < R; j++ {
			waffle[j] = make([]bool, C)
		}

		for j := 0; j < R; j++ {
			var row string
			fmt.Scanf("%s", &row)
			for k := 0; k < C; k++ {
				if row[k] == '@' {
					waffle[j][k] = true
				}
			}
		}

		horizontalChoco := make([]int, R)
		verticalChoco := make([]int, C)
		var total int

		for j := 0; j < R; j++ {
			for k := 0; k < C; k++ {
				if waffle[j][k] {
					horizontalChoco[j]++
					verticalChoco[k]++
					total++
				}
			}
		}

		//horizontal accumunlate
		horizontalAccum := make([][]int, R)
		for j := 0; j < R; j++ {
			horizontalAccum[j] = make([]int, C)
		}
		for j := 0; j < R; j++ {
			for k := 0; k < C; k++ {
				if waffle[j][k] {
					horizontalAccum[j][k] = 1
				}
				if k > 0 {
					horizontalAccum[j][k] += horizontalAccum[j][k-1]
				}
			}
		}

		maxSizeChoco := total / ((H + 1) * (V + 1))

		//horizontal cut
		horizontalCut := make([]int, H)
		var curr int
		var failure bool
		var currCut int
		for j := 0; j < R; j++ {
			if currCut < H {
				if curr < maxSizeChoco*(V+1) {
					curr += horizontalChoco[j]
				} else if curr > maxSizeChoco*(V+1) {
					failure = true
					break
				} else if curr == maxSizeChoco*(V+1) {
					horizontalCut[currCut] = j
					currCut++
					curr = 0
				}
			} else {
				break
			}
		}
		if failure {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i+1)
			continue
		}

		fmt.Printf("%+v\n", horizontalCut)

		//vertical cut
		verticallCut := make([]int, H)
		curr = 0
		currCut = 0
		for j := 0; j < C; j++ {
			if currCut < H {
				if curr < maxSizeChoco*(H+1) {
					curr += verticalChoco[j]
				} else if curr > maxSizeChoco*(H+1) {
					failure = true
					break
				} else if curr == maxSizeChoco*(H+1) {
					if !checkHorizontal(horizontalAccum, horizontalCut, j) {
						failure = true
						break
					}
					verticallCut[currCut] = j
					currCut++
					curr = 0
				}
			} else {
				break
			}
		}
		if failure {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", i+1)
			continue
		}

		fmt.Printf("Case #%d: POSSIBLE\n", i+1)
	}
}

func checkHorizontal(horizontalAccum [][]int, horizontalCut []int, verticalCut []int) bool {
	fmt.Printf("%+v\n", horizontalAccum)
	fmt.Println(verticalCut)
	for i := range horizontalCut {
		for z := range verticalCut {
			var prevHorizontalCut int
			if i != 0 {
				prevHorizontalCut = horizontalCut[i-1]
			}
			var left int
			for j := prevHorizontalCut; j < horizontalCut[i]; j++ {
				left += horizontalAccum[j][verticalCut[z]-1]
			}
			var right int
			for j := prevHorizontalCut; j < horizontalCut[i]; j++ {
				right += horizontalAccum[j][len(horizontalAccum[j])-1]
				if verticalCut[z] != 0 {
					right -= horizontalAccum[j][verticalCut-1]
					fmt.Printf("horizontalAccum[%d][%d]: %d\n", j, verticalCut-1, horizontalAccum[j][verticalCut-1])
				}
			}
			fmt.Printf("left: %d, right: %d\n", left, right)
			if left != right {
				return false
			}
		}
	}
	return true
}
