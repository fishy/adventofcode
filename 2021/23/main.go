package main

import (
	"bufio"
	"fmt"
	"strings"
)

var amphipods = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

var amphiString = map[int]string{
	0:    ".",
	1:    "A",
	10:   "B",
	100:  "C",
	1000: "D",
}

var targetAmphi = []int{1, 10, 100, 1000}

var targetRoom = map[int]int{
	1:    0,
	10:   1,
	100:  2,
	1000: 3,
}

var hallStoppable = [11]bool{true, true, false, true, false, true, false, true, false, true, true}

var roomPosition = []int{2, 4, 6, 8}

type state struct {
	roomSize int
	rooms    [4][4]int
	hall     [11]int
}

func (s state) isTarget() bool {
	for i, room := range s.rooms {
		for j := 0; j < s.roomSize; j++ {
			amphi := room[j]
			if amphi != targetAmphi[i] {
				return false
			}
		}
	}
	return true
}

func (s state) hallClear(from, to int, fromRoom bool) (ret bool) {
	if from == to {
		return true
	}

	var min, max int
	if from < to {
		min = from
		if !fromRoom {
			min++
		}
		max = to + 1
	} else {
		min = to
		max = from + 1
		if !fromRoom {
			max--
		}
	}
	for i := min; i < max; i++ {
		if s.hall[i] != 0 {
			return false
		}
	}
	return true
}

func (s state) String() string {
	var sb strings.Builder
	sb.WriteString("#############\n#")
	for _, amphi := range s.hall {
		sb.WriteString(amphiString[amphi])
	}
	sb.WriteString("#\n###")
	for _, room := range s.rooms {
		sb.WriteString(amphiString[room[0]])
		sb.WriteString("#")
	}
	sb.WriteString("##\n  #")
	for _, room := range s.rooms {
		sb.WriteString(amphiString[room[1]])
		sb.WriteString("#")
	}
	for i := 2; i < s.roomSize; i++ {
		sb.WriteString("\n  #")
		for _, room := range s.rooms {
			sb.WriteString(amphiString[room[i]])
			sb.WriteString("#")
		}
	}
	sb.WriteString("\n  #########")
	return sb.String()
}

func readInput(input string) state {
	var s state
	scanner := bufio.NewScanner(strings.NewReader(input))
	var i int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == '#'
		})
		if len(fields) != 4 {
			continue
		}
		for j, f := range fields {
			s.rooms[j][i] = amphipods[f]
		}
		i++
	}
	s.roomSize = i
	return s
}

var cache = make(map[state]int)

func hallSteps(from, to int) int {
	if from > to {
		return from - to
	}
	return to - from
}

func resolve(s state, visited map[state]bool) (min int) {
	if cost, ok := cache[s]; ok {
		return cost
	}
	if visited[s] {
		return -1
	}
	defer func() {
		cache[s] = min
	}()

	visited[s] = true
	if s.isTarget() {
		return 0
	}

	min = -1
	replaceMin := func(cost int) {
		if min < 0 || cost < min {
			min = cost
		}
	}

	for from, amphi := range s.hall {
		if amphi == 0 {
			continue
		}
		roomIndex := targetRoom[amphi]
		room := s.rooms[roomIndex]
		if room[0] != 0 {
			continue
		}
		to := roomPosition[roomIndex]
		if !s.hallClear(from, to, false) {
			continue
		}
		var extra int
		legalRoom := true
		for i := 1; i < s.roomSize; i++ {
			if room[i] == 0 {
				extra++
				continue
			}
			if room[i] != amphi {
				legalRoom = false
				break
			}
		}
		if !legalRoom {
			continue
		}
		next := s
		next.rooms[roomIndex][extra] = amphi
		next.hall[from] = 0
		nextCost := resolve(next, visited)
		if nextCost < 0 {
			continue
		}
		cost := amphi*(hallSteps(from, to)+1+extra) + nextCost
		replaceMin(cost)
	}

	for i, room := range s.rooms {
		allClear := true
		for j := 0; j < s.roomSize; j++ {
			if room[j] != 0 && room[j] != targetAmphi[i] {
				allClear = false
				break
			}
		}
		if allClear {
			continue
		}
		from := roomPosition[i]
		for to := range s.hall {
			if s.hall[to] != 0 {
				continue
			}
			if !hallStoppable[to] {
				continue
			}
			if !s.hallClear(from, to, true) {
				continue
			}
			var extra int
			next := s
			var amphi int
			for j := 0; j < s.roomSize; j++ {
				amphi = room[j]
				next.rooms[i][j] = 0
				extra++
				if amphi != 0 {
					break
				}
			}
			next.hall[to] = amphi
			nextCost := resolve(next, visited)
			if nextCost < 0 {
				continue
			}
			cost := amphi*(hallSteps(from, to)+extra) + nextCost
			replaceMin(cost)
		}
	}
	return min
}

func main() {
	s1 := readInput(input1)
	fmt.Println(s1)
	fmt.Println(resolve(s1, make(map[state]bool)))

	s2 := readInput(input2)
	fmt.Println(s2)
	fmt.Println(resolve(s2, make(map[state]bool)))
}

const input1 = `
#############
#...........#
###B#A#A#D###
  #B#C#D#C#
  #########
`

const input2 = `
#############
#...........#
###B#A#A#D###
  #D#C#B#A#
  #D#B#A#C#
  #B#C#D#C#
  #########
`
