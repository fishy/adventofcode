package main

import (
	"bufio"
	"fmt"
	"strings"
)

type coordinate struct {
	x, y int
}

var directions = []coordinate{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func readInput() [][]int {
	var m [][]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var l []int
		for _, r := range line {
			l = append(l, int(r-'0'))
		}
		m = append(m, l)
	}
	return m
}

func flash(m [][]int, flashed map[coordinate]bool, c coordinate) {
	if flashed[c] {
		return
	}
	if m[c.x][c.y] <= 9 {
		return
	}
	flashed[c] = true
	maxX := len(m)
	maxY := len(m[0])
	for _, d := range directions {
		x := c.x + d.x
		y := c.y + d.y
		if x < 0 || y < 0 || x >= maxX || y >= maxY {
			continue
		}
		m[x][y]++
		cc := coordinate{x, y}
		if m[x][y] > 9 {
			flash(m, flashed, cc)
		}
	}
}

func iterate(m [][]int) int {
	flashed := make(map[coordinate]bool)
	for i := range m {
		for j := range m[i] {
			m[i][j]++
		}
	}
	for i := range m {
		for j := range m[i] {
			if m[i][j] > 9 {
				flash(m, flashed, coordinate{i, j})
			}
		}
	}
	for c := range flashed {
		m[c.x][c.y] = 0
	}
	return len(flashed)
}

func main() {
	m := readInput()
	var part1 int
	for i := 0; ; i++ {
		n := iterate(m)
		part1 += n
		fmt.Printf("step %d: %d, %d\n", i+1, n, part1)
		if n == 100 && i >= 100 {
			break
		}
	}
}

const input = `6636827465
6774248431
4227386366
7447452613
6223122545
2814388766
6615551144
4836235836
5334783256
4128344843`
