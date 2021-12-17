package main

import (
	"flag"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var maxY = flag.Int(
	"y",
	100,
	"the max y to try",
)

var re = regexp.MustCompile(`target area: x=(-?[0-9]*)\.\.(-?[0-9]*), y=(-?[0-9]*)\.\.(-?[0-9]*)`)

type point struct {
	x, y int
}

type area struct {
	low  point
	high point
}

func (a area) checkX(p point) int {
	if p.x > a.high.x {
		return 1
	}
	if p.x < a.low.x {
		return -1
	}
	return 0
}

func (a area) checkY(p point) int {
	if p.y < a.low.y {
		return 1
	}
	if p.y > a.high.y {
		return -1
	}
	return 0
}

// return 0 means hit, 1 means overshoot, -1 means not yet.
func (a area) check(p point) int {
	x := a.checkX(p)
	y := a.checkY(p)
	if x > 0 || y > 0 {
		return 1
	}
	if x < 0 || y < 0 {
		return -1
	}
	return 0
}

func readInput(s string) area {
	matches := re.FindStringSubmatch(s)
	x1, _ := strconv.ParseInt(matches[1], 10, 64)
	x2, _ := strconv.ParseInt(matches[2], 10, 64)
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	y1, _ := strconv.ParseInt(matches[3], 10, 64)
	y2, _ := strconv.ParseInt(matches[4], 10, 64)
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return area{
		low: point{
			x: int(x1),
			y: int(y1),
		},
		high: point{
			x: int(x2),
			y: int(y2),
		},
	}
}

func iterate(v point, target area) (maxY int, hit bool) {
	var p point
	for {
		if p.y > maxY {
			maxY = p.y
		}
		c := target.check(p)
		if c == 0 {
			return maxY, true
		}
		if c == 1 {
			return maxY, false
		}
		p.x += v.x
		p.y += v.y
		if v.x > 0 {
			v.x--
		} else {
			if target.checkX(p) < 0 {
				return maxY, false
			}
		}
		v.y--
	}
}

func main() {
	flag.Parse()
	minY := -*maxY
	target := readInput(input)
	fmt.Println("target:", target)
	minX := int(math.Floor(math.Sqrt(float64(target.low.x) * 2)))
	for (1+minX)*minX/2 < target.low.x {
		minX++
	}
	// part 1
	for y := *maxY; ; y-- {
		v := point{minX, y}
		if maxY, hit := iterate(v, target); hit {
			fmt.Println(maxY)
			break
		}
	}
	// part 2
	var count int
	for x := minX; x <= target.high.x; x++ {
		for y := *maxY; y >= minY; y-- {
			v := point{x, y}
			if _, hit := iterate(v, target); hit {
				count++
				// fmt.Println(v)
				continue
			}
		}
	}
	fmt.Println(count)
}

const input = `target area: x=277..318, y=-92..-53`
