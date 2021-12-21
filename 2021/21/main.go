package main

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	n    = 3
	min1 = 1000
	min2 = 21
)

var inputRE = regexp.MustCompile(`starting position: ([0-9]+)`)

type deterministicDice int

func (d *deterministicDice) roll() int {
	*d++
	ret := *d % 100
	if ret == 0 {
		ret = 100
	}
	return int(ret)
}

// sparse array
var quantumDice [10]int64

const (
	qdMin = 3
	qdMax = 9
)

func init() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				quantumDice[i+j+k]++
			}
		}
	}
	fmt.Println(quantumDice)
}

type player struct {
	position int
	score    int
}

func (p *player) move(m int) int {
	p.position += m
	for p.position > 10 {
		p.position -= 10
	}
	p.score += p.position
	return p.score
}

func readPlayers(input string) []*player {
	var players []*player
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matches := inputRE.FindStringSubmatch(scanner.Text())
		if len(matches) == 2 {
			p, _ := strconv.ParseInt(matches[1], 10, 64)
			players = append(players, &player{
				position: int(p),
			})
		}
	}
	return players
}

func recursion(player1, player2 *player, p1win, p2win *int64, base int64, i int) {
	i++
	// fmt.Println(i, player1, player2)
	for i := qdMin; i <= qdMax; i++ {
		cp := *player1
		univ := base * quantumDice[i]
		score := cp.move(i)
		if score >= min2 {
			*p1win += univ
		} else {
			recursion(player2, &cp, p2win, p1win, univ, i)
		}
	}
}

func main() {
	var dice deterministicDice
	players := readPlayers(input)
	for i := 0; ; i++ {
		player := players[i%len(players)]
		var m int
		for j := 0; j < n; j++ {
			m += dice.roll()
		}
		if player.move(m) >= min1 {
			i++
			op := players[i%len(players)]
			fmt.Println(op.score * i * n)
			break
		}
	}

	players = readPlayers(input)
	var p1win, p2win int64
	recursion(players[0], players[1], &p1win, &p2win, 1, 0)
	winners := []int{int(p1win), int(p2win)}
	sort.Ints(winners)
	fmt.Println(winners)
}

const input = `
Player 1 starting position: 10
Player 2 starting position: 9
`
