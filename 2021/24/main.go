package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const inputs = 14

func printZ(z int) string {
	var zs string
	for z > 0 {
		zs = string([]rune{'a' + rune(z%26)}) + zs
		z /= 26
	}
	return zs
}

func printIndex(index int) string {
	return fmt.Sprintf("#%02d", index)
}

type variables [4]int

func (v variables) String() string {
	z := v[3]
	return fmt.Sprintf("w:%2d x:%2d y:%2d z:%s", v[0], v[1], v[2], printZ(z))
}

func (v variables) z() int {
	return v[3]
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

type instruction interface {
	String() string

	execute(vars *variables, input int)
	binary() *binary
}

type inp struct {
	a value
}

func (i *inp) String() string {
	return fmt.Sprintf("(inp %v)", i.a)
}

func (i *inp) execute(vars *variables, input int) {
	i.a.set(vars, input)
}

func (*inp) binary() *binary {
	return nil
}

type binary struct {
	operator string
	a, b     value
}

func (b *binary) String() string {
	return fmt.Sprintf("(%s %v %v)", b.operator, b.a, b.b)
}

func (b *binary) binary() *binary {
	return b
}

func (b *binary) execute(vars *variables, _ int) {
	var n int
	left := b.a.get(vars)
	right := b.b.get(vars)
	switch b.operator {
	case "add":
		n = left + right
	case "mul":
		n = left * right
	case "div":
		n = left / right
	case "mod":
		n = left % right
	case "eql":
		if left == right {
			n = 1
		}
	}
	b.a.set(vars, n)
}

func readInput(input string) []instruction {
	var instructions []instruction
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

func printDiff(instructions []instruction) {
	n := inputs
	m := len(instructions) / n
	fmt.Println(len(instructions), m*n)
	for i := 0; i < m; i++ {
		last := instructions[i].String()
		fmt.Printf("%v: %v\n", printIndex(i), last)
		for j := 1; j < n; j++ {
			index := j*m + i
			curr := instructions[index].String()
			if curr != last {
				fmt.Printf("  %v: %v\n", printIndex(j), curr)
			}
			last = curr
		}
	}
}

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() int {
	index := len(*s) - 1
	v := (*s)[index]
	*s = (*s)[:index]
	return v
}

const (
	offsetX = 5
	offsetY = 15
)

func generateYmap(instructions []instruction) map[int]int {
	m := len(instructions) / inputs
	ymap := make(map[int]int, inputs/2)
	var s stack
	for i := 0; i < inputs; i++ {
		base := m * i
		argX := instructions[base+offsetX].binary().b.get(nil)
		if argX > 0 {
			s.push(i)
		} else {
			j := s.pop()
			ymap[j] = argX
		}
	}
	return ymap
}

func chooseMin(min, max int) int {
	return min
}

func chooseMax(min, max int) int {
	return max
}

func resolve(
	instructions []instruction,
	ymap map[int]int,
	choose func(min, max int) int,
) int {
	var vars variables
	var ws [inputs]int
	m := len(instructions) / inputs
	for i := 0; i < inputs; i++ {
		base := m * i
		min := 1
		max := 9

		argX := instructions[base+offsetX].binary().b.get(&vars)
		argY := instructions[base+offsetY].binary().b.get(&vars)
		if argX < 0 {
			w := vars.z()%26 + argX
			min = w
			max = w
		} else {
			sum := argY + ymap[i]
			if sum < 0 {
				min = 1 - sum
			} else if sum > 0 {
				max = 9 - sum
			}
		}

		w := choose(min, max)
		ws[i] = w
		for _, ins := range instructions[base : base+m] {
			ins.execute(&vars, w)
		}
		fmt.Println(printIndex(i), ws, printZ(vars[3]))
	}
	var n int
	for _, w := range ws {
		n *= 10
		n += w
	}
	return n
}

func main() {
	instructions := readInput(input)
	printDiff(instructions)
	ymap := generateYmap(instructions)
	fmt.Println(ymap)
	fmt.Println(resolve(instructions, ymap, chooseMax))
	fmt.Println(resolve(instructions, ymap, chooseMin))
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
