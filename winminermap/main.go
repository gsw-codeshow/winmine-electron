package main

import (
	"flag"
	"fmt"
	"math/rand"
)

type pointxy struct {
	x, y int
}

var winminexy = [8]pointxy{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1}}

var winminemap = [12][12]int{}

func randomSquare(n int) {
	if n > 12*12 {
		n = 12 * 12
	}
	for i := 0; i < n; {
		x := rand.Intn(12)
		y := rand.Intn(12)
		if 0 == winminemap[x][y] {
			winminemap[x][y] = 9
			i++
		}
	}
	return
}

func makeSquare() {
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			if 9 == winminemap[i][j] {
				for _, xy := range winminexy {
					posx := xy.x + i
					if posx >= 12 || posx < 0 {
						continue
					}
					posy := xy.y + j
					if posy >= 12 || posy < 0 {
						continue
					}
					if 9 == winminemap[posx][posy] {
						continue
					}
					winminemap[posx][posy]++
				}
			}
		}
	}
}

func main() {
	n := flag.Int("n", 0, "n")
	flag.Parse()
	randomSquare(*n)
	makeSquare()
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Print(winminemap[i][j])
		}
	}
}
