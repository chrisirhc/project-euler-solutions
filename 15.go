package main

import (
	"fmt"
)

func main() {
	const GRIDWIDTH, GRIDHEIGHT = 20, 20
	const HEIGHT, WIDTH = GRIDWIDTH + 1, GRIDHEIGHT + 1
	ways := make([][WIDTH]int, HEIGHT)

	for i := 0; i < WIDTH; i++ {
		ways[0][i] = 1
	}

	for y := 1; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			ways[y][x] += ways[y-1][x]
			if x > 0 {
				ways[y][x] += ways[y][x-1]
			}
		}
	}

	fmt.Println(ways[GRIDHEIGHT][GRIDWIDTH])
}

