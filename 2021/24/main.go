package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const inputs = 14

type variables [4]int

func printz(z int) string {
	var zs string
	for z > 0 {
		zs = string([]byte{'a' + byte(z%26)}) + zs
		z /= 26
	}
	return zs
}

func printlevel(level int) string {
	return fmt.Sprintf("#%02d", level)
}

func (v variables) String() string {
	z := v[3]
	return fmt.Sprintf("w:%2d x:%2d y:%2d z:%s", v[0], v[1], v[2], printz(z))
}

func readString(s string) value {
	switch s {
	default:
		v, _ := strconv.ParseInt(s, 10, 64)
		return number(v)
	case "w", "x", "y", "z":
		return variable(s[0] - 'w')
	}
}

type value interface {
	get(vars *variables) int
	set(vars *variables, n int)
}

type number int

func (n number) get(*variables) int {
	return int(n)
}

// no-op
func (number) set(*variables, int) {}

func (n number) String() string {
	return fmt.Sprintf("%d", n)
}

type variable int

func (v variable) get(vars *variables) int {
	return int(vars[v])
}

func (v variable) set(vars *variables, n int) {
	vars[v] = n
}

func (v variable) String() string {
	return string([]rune{rune(v) + 'w'})
}

type inp struct {
	a value
}

func (i inp) String() string {
	return fmt.Sprintf("(inp %v)", i.a)
}

type binary struct {
	operator string
	a, b     value
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %v %v)", b.operator, b.a, b.b)
}

func readInput(input string) []interface{} {
	var instructions []interface{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		switch fields[0] {
		default:
			instructions = append(instructions, &binary{
				operator: fields[0],
				a:        readString(fields[1]),
				b:        readString(fields[2]),
			})
		case "inp":
			instructions = append(instructions, &inp{a: readString(fields[1])})
		}
	}
	return instructions
}

func printDiff(instructions []interface{}) {
	n := inputs
	m := len(instructions) / n
	fmt.Println(len(instructions), m*n)
	for i := 0; i < m; i++ {
		last := fmt.Sprintf("%v", instructions[i])
		fmt.Printf("%v: %s\n", printlevel(i), last)
		for j := 1; j < n; j++ {
			index := j*m + i
			curr := fmt.Sprintf("%v", instructions[index])
			if curr != last {
				fmt.Printf("  %v: %s\n", printlevel(j), curr)
			}
			last = curr
		}
	}
}

func generateMods(instructions []interface{}) (mods [inputs][]int, ymap map[int]int) {
	m := len(instructions) / inputs
	for i := 0; i < inputs; i++ {
		base := m * i
		mods[i] = []int{
			instructions[base+4].(*binary).b.get(nil),
			instructions[base+5].(*binary).b.get(nil),
			instructions[base+15].(*binary).b.get(nil),
		}
	}
	ymap = make(map[int]int)
	var stack []int
	for i, mod := range mods {
		xmod := mod[1]
		if xmod > 0 {
			stack = append(stack, i)
		} else {
			n := len(stack) - 1
			j := stack[n]
			stack = stack[:n]
			ymap[j] = mods[i][1]
		}
	}
	return mods, ymap
}

func min(min, max int) int {
	return min
}

func max(min, max int) int {
	return max
}

func resolve(
	mods [inputs][]int,
	ymap map[int]int,
	choice func(min, max int) int,
	z int,
	level int,
	ws [inputs]int,
) int {
	if level >= inputs {
		if z != 0 {
			return 0
		}
		var n int
		for _, w := range ws {
			n *= 10
			n += w
		}
		return n
	}
	zmod := mods[level][0]
	xmod := mods[level][1]
	ymod := mods[level][2]
	min := 1
	max := 9
	if xmod < 0 {
		w := z%26 + xmod
		if w < 1 || w > 9 {
			return 0
		}
		min = w
		max = w
	} else {
		sum := ymod + ymap[level]
		if sum < 0 {
			min = 1 - sum
		} else if sum > 0 {
			max = 9 - sum
		}
	}

	w := choice(min, max)
	ws[level] = w
	x := z
	x %= 26
	z /= zmod
	x += xmod
	var y int
	if x == w {
		x = 0
		y = 1
	} else {
		x = 1
		y = 26
	}
	z *= y
	y = w
	y += ymod
	y *= x
	z += y
	fmt.Println(printlevel(level), ws, printz(z))
	return resolve(mods, ymap, choice, z, level+1, ws)
}

func main() {
	instructions := readInput(input)
	printDiff(instructions)
	mods, ymap := generateMods(instructions)
	fmt.Println(resolve(mods, ymap, max, 0, 0, [inputs]int{}))
	fmt.Println(resolve(mods, ymap, min, 0, 0, [inputs]int{}))
}

const input = `
inp w
mul x 0
add x z
mod x 26
div z 1
add x 12
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 6
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 12
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 10
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 5
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 10
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 10
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -16
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 7
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 14
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 0
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 12
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 4
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -4
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 12
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 15
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 14
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -7
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 13
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -8
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 10
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -4
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 11
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -15
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 9
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -8
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 9
mul y x
add z y
`
