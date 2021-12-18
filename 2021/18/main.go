package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type element interface {
	String() string

	magnitude() int
	isNumber() bool
}

type node struct {
	element

	parent *node
	isLeft bool
	level  int
}

func (n *node) incLevel() {
	n.level++
	if p, ok := n.element.(*pair); ok {
		p.left.incLevel()
		p.right.incLevel()
	}
}

func (n *node) lastNumber() *number {
	if p, ok := n.element.(*pair); ok {
		return p.right.lastNumber()
	}
	return n.element.(*number)
}

func (n *node) leftNumber() *number {
	parent := n.parent
	if parent == nil {
		return nil
	}
	if n.isLeft {
		return parent.leftNumber()
	}
	p := parent.element.(*pair)
	return p.left.lastNumber()
}

func (n *node) firstNumber() *number {
	if p, ok := n.element.(*pair); ok {
		return p.left.firstNumber()
	}
	return n.element.(*number)
}

func (n *node) rightNumber() *number {
	parent := n.parent
	if parent == nil {
		return nil
	}
	if n.isLeft {
		return parent.element.(*pair).right.firstNumber()
	}
	return parent.rightNumber()
}

func (n *node) explode() bool {
	p, ok := n.element.(*pair)
	if !ok {
		return false
	}

	if n.level > 4 && p.left.isNumber() && p.right.isNumber() {
		if left := n.leftNumber(); left != nil {
			*left += number(p.left.magnitude())
		}
		if right := n.rightNumber(); right != nil {
			*right += number(p.right.magnitude())
		}
		parent := n.parent.element.(*pair)
		if n.isLeft {
			parent.left.element = numberPtr(0)
		} else {
			parent.right.element = numberPtr(0)
		}
		return true
	}

	for _, e := range []*node{p.left, p.right} {
		if e.explode() {
			return true
		}
	}
	return false
}

func (n *node) split() bool {
	if m := n.magnitude(); m >= 10 && n.isNumber() {
		l := m / 2
		r := m - l
		left := &node{
			element: numberPtr(number(l)),
			parent:  n,
			isLeft:  true,
			level:   n.level + 1,
		}
		right := &node{
			element: numberPtr(number(r)),
			parent:  n,
			isLeft:  false,
			level:   n.level + 1,
		}
		n.element = &pair{
			left:  left,
			right: right,
		}
		return true
	}

	p, ok := n.element.(*pair)
	if !ok {
		return false
	}
	for _, e := range []*node{p.left, p.right} {
		if e.split() {
			return true
		}
	}
	return false
}

func (n *node) reduce() *node {
	for n.explode() || n.split() {
	}
	return n
}

func (n *node) clone(parent *node) *node {
	c := &node{
		isLeft: n.isLeft,
		level:  n.level,
		parent: parent,
	}
	if p, ok := n.element.(*pair); ok {
		c.element = &pair{
			left:  p.left.clone(c),
			right: p.right.clone(c),
		}
	} else {
		c.element = numberPtr(number(n.magnitude()))
	}
	return c
}

func add(a, b *node) *node {
	a = a.clone(nil)
	b = b.clone(nil)
	p := &node{
		element: &pair{
			left:  a,
			right: b,
		},
	}
	a.parent = p
	a.isLeft = true
	b.parent = p
	b.isLeft = false
	p.incLevel()
	return p.clone(nil)
}

type number int

func (n *number) magnitude() int {
	return int(*n)
}

func (*number) isNumber() bool {
	return true
}

func (n *number) String() string {
	return fmt.Sprintf("%d", *n)
}

func numberPtr(n number) *number {
	return &n
}

type pair struct {
	left  *node
	right *node
}

func (p *pair) magnitude() int {
	return p.left.magnitude()*3 + p.right.magnitude()*2
}

func (*pair) isNumber() bool {
	return false
}

func (p *pair) String() string {
	return fmt.Sprintf("[%v,%v]", p.left, p.right)
}

func byteToNumber(b byte) number {
	return number(b - '0')
}

func readElement(buf *bytes.Buffer) *node {
	b, _ := buf.ReadByte()
	if b != '[' {
		// number
		var n number
		for ; b >= '0' && b <= '9'; b, _ = buf.ReadByte() {
			n *= 10
			n += byteToNumber(b)
		}
		buf.UnreadByte()
		return &node{
			element: numberPtr(n),
			level:   1,
		}
	}
	left := readElement(buf)
	buf.ReadByte() // ,
	right := readElement(buf)
	buf.ReadByte() // ]
	return add(left, right)
}

func readString(s string) *node {
	return readElement(bytes.NewBuffer([]byte(s)))
}

func readInput(input string) []*node {
	var lines []*node
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lines = append(lines, readString(line))
	}
	return lines
}

func main() {
	lines := readInput(input)
	v := lines[0]
	for i := 1; i < len(lines); i++ {
		v = add(v, lines[i]).reduce()
	}
	fmt.Println(v.magnitude())

	lines = readInput(input)
	var max int
	for i := range lines {
		for j := range lines {
			if i == j {
				continue
			}
			m := add(lines[i], lines[j]).reduce().magnitude()
			if m > max {
				max = m
			}
		}
	}
	fmt.Println(max)
}

const input = `
[[[[9,5],[9,4]],[[6,5],[7,0]]],4]
[[[5,2],[[7,2],1]],[[[7,5],[0,8]],[[6,9],[7,3]]]]
[[[9,7],[0,1]],9]
[1,[[7,3],[[3,7],[3,2]]]]
[[9,[[0,8],7]],[[3,1],[[6,6],[9,0]]]]
[4,[[4,4],[[7,7],1]]]
[[[[6,2],[5,1]],[[3,3],9]],[7,[[5,7],[5,0]]]]
[[[[4,8],[4,9]],[1,[9,3]]],[1,[1,[6,1]]]]
[[[[4,7],[3,4]],[8,3]],[[3,7],[0,[1,8]]]]
[[[6,[4,8]],[4,5]],[4,[1,3]]]
[[[0,7],0],[[6,[1,8]],[9,[7,9]]]]
[[[[4,8],[3,9]],[4,5]],[1,1]]
[[[4,2],[0,[6,7]]],[[[1,8],2],[8,8]]]
[[[[1,1],7],5],[[6,[5,6]],[6,[7,5]]]]
[[[[3,2],5],[[5,3],1]],[[[0,4],[9,6]],9]]
[[6,[7,6]],9]
[[[[4,0],[0,1]],7],1]
[[[[1,3],4],6],[[1,[4,2]],[1,4]]]
[[[[6,9],[4,1]],[[6,3],[0,8]]],[[4,0],[[3,2],[2,9]]]]
[[[3,6],[[2,0],[3,2]]],[2,5]]
[[[[4,3],5],5],[[4,[4,0]],6]]
[[[[4,0],3],[[3,5],8]],[[8,[4,4]],[[9,9],[4,1]]]]
[[[2,7],6],1]
[[[[5,3],[8,4]],[0,0]],4]
[[[0,[8,1]],0],3]
[[[6,5],[8,2]],[[[6,9],[6,1]],[9,9]]]
[0,[[4,9],6]]
[[9,[[9,9],4]],[[[4,7],1],2]]
[[8,0],[[[0,7],6],[[6,4],2]]]
[[1,[[2,4],8]],1]
[[[[1,3],4],[[1,3],0]],[[[1,2],3],2]]
[[[[2,1],2],[5,[2,8]]],[2,[[6,0],2]]]
[[[8,[1,0]],[[6,7],[9,6]]],[[2,[9,7]],5]]
[[[3,[2,0]],[[3,2],[0,0]]],[[[4,6],[9,4]],[[7,8],[5,1]]]]
[[3,[[9,9],[7,2]]],[[1,3],[2,[3,2]]]]
[4,[4,[[9,5],6]]]
[[[[5,7],7],[[3,4],0]],[[9,[8,2]],[2,3]]]
[[[[2,1],[5,7]],4],[[[6,3],8],[[1,6],[5,1]]]]
[[[4,4],[[0,9],[7,8]]],[[2,[2,5]],5]]
[1,[5,[[3,7],[8,2]]]]
[[[[9,5],[8,6]],[5,5]],[[[9,2],8],[[9,3],[3,8]]]]
[0,[[9,5],[[3,7],7]]]
[[[8,[0,4]],[[2,9],6]],[[6,[8,0]],4]]
[[0,[3,5]],[[5,[0,1]],[[3,6],7]]]
[[2,[7,1]],[[[5,0],[7,7]],[[2,3],9]]]
[[5,[9,[3,9]]],[[8,[3,7]],[[7,6],[3,0]]]]
[[[4,[2,5]],5],[3,1]]
[[[[4,3],1],[[5,7],6]],[0,[3,1]]]
[[8,9],[[[0,7],5],[6,[5,7]]]]
[[6,8],[[5,8],[[8,2],[6,0]]]]
[[1,[5,6]],5]
[[[6,1],[9,[1,2]]],1]
[[5,[7,[4,8]]],[[4,[2,9]],5]]
[[[2,2],[[7,1],3]],[[[9,7],[4,6]],[1,[0,1]]]]
[[3,[6,[4,5]]],2]
[[[0,2],[[8,1],[0,6]]],[[7,[9,6]],0]]
[[[[1,0],[5,1]],[[0,6],5]],[[[1,8],8],[[0,2],5]]]
[[6,[[3,6],6]],[[[9,7],[6,4]],[[9,5],1]]]
[[[0,[5,6]],[9,0]],[[2,9],9]]
[1,[[4,[9,3]],0]]
[[1,0],[[1,9],[4,8]]]
[[[9,3],[7,0]],[[[5,1],[3,8]],9]]
[[[3,9],[[5,9],2]],[[7,2],1]]
[[1,[[3,0],[7,6]]],[7,[8,1]]]
[0,[6,[[7,1],[1,1]]]]
[[4,[[5,0],[2,1]]],[[[8,8],[8,1]],7]]
[[[[9,3],[4,3]],4],[7,5]]
[[9,[[7,4],[8,3]]],[[[1,9],7],[[1,6],[3,1]]]]
[[6,9],[5,[0,[5,1]]]]
[[[8,7],3],[[4,8],[0,7]]]
[[[[3,1],2],[[1,6],[4,3]]],[0,6]]
[[5,[[5,4],3]],[[8,8],9]]
[[5,[3,[4,5]]],[[2,[6,0]],[6,1]]]
[[[[9,5],3],6],[[8,[1,9]],[[5,2],5]]]
[[[7,5],[[3,6],4]],[6,[[5,1],[0,1]]]]
[[1,[[4,8],[1,3]]],7]
[[4,[[4,0],5]],[[[6,2],7],[[4,8],[4,9]]]]
[[[[2,3],[0,9]],[7,2]],[4,5]]
[[[[7,7],[8,0]],[7,7]],[[[6,6],[3,2]],[4,[4,3]]]]
[[[[8,7],6],[[5,5],0]],[[6,[7,3]],[[4,1],[1,7]]]]
[[[2,[2,2]],[[5,2],1]],[[9,[9,2]],6]]
[[[[1,7],6],[[8,8],5]],[6,[1,[1,7]]]]
[[[[8,6],[3,2]],[[5,2],[2,0]]],[[[8,7],2],[[5,5],2]]]
[[[8,[9,0]],[[9,5],[7,5]]],[[5,1],[[1,1],[4,6]]]]
[5,[9,[[0,2],7]]]
[8,[[0,[4,9]],[[7,4],9]]]
[[[[2,9],5],[[0,6],[6,6]]],[[0,6],[[4,2],[9,9]]]]
[7,[[[4,3],3],[[5,4],[6,0]]]]
[[0,[8,[1,1]]],5]
[[[1,8],[[4,6],[9,7]]],[[[6,6],[2,6]],[4,3]]]
[[0,[[7,5],[9,9]]],[[9,7],[6,2]]]
[[[9,[3,0]],[[1,4],0]],[[1,1],1]]
[[[0,7],[[3,0],8]],[[6,[8,0]],[[4,5],[4,0]]]]
[[[[2,9],[4,2]],[5,[9,3]]],[4,[2,[3,4]]]]
[[[1,[7,3]],[[5,7],0]],[6,[[6,5],2]]]
[4,5]
[[7,9],[6,[[6,5],[1,0]]]]
[[4,[[7,5],8]],[[4,0],[[6,6],[0,4]]]]
[[[9,[7,7]],[[4,2],7]],4]
[[0,[0,3]],5]
`
