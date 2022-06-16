package main

import "fmt"

type State struct {
	state bool
}

var grid = [4][3]State{}

func main() {
	fmt.Print(1 / 5)
	// board := [4][3]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}

	// for i, v1 := range board {
	// 	for j, v2 := range v1 {
	// 		if v2 == 0 {
	// 			grid[i][j].state = false
	// 		} else {
	// 			grid[i][j].state = true
	// 		}
	// 	}
	// }
	// for i, v1 := range board {
	// 	for j, v2 := range v1 {
	// 		aliveCountAround := aliveCountAround(i, j)
	// 		if aliveCountAround >= 4 {
	// 			v2 = 0
	// 		} else {
	// 			v2 = 1
	// 		}
	// 	}
	// }
}

//计算某个生命的邻居生存个数
func aliveCountAround(x, y int) int {
	around := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	count := 0
	for i := 0; i < 8; i++ {
		if grid[x+around[i][0]][y+around[i][1]].state {
			count++
		}
	}
	return count
}
