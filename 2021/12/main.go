package main

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	start = "start"
	end   = "end"
)

type stringSet = map[string]bool

type maps = map[string]stringSet

func isSmall(name string) bool {
	return strings.ToLower(name) == name
}

func readMaps(input string) maps {
	scanner := bufio.NewScanner(strings.NewReader(input))
	m := make(maps)
	for scanner.Scan() {
		pair := strings.Split(strings.TrimSpace(scanner.Text()), "-")
		if len(pair) != 2 {
			continue
		}
		for i := 0; i < 2; i++ {
			j := 1 - i
			if _, ok := m[pair[i]]; !ok {
				m[pair[i]] = make(stringSet)
			}
			m[pair[i]][pair[j]] = true
		}
	}
	return m
}

func recursion(m maps, s string, visited stringSet, twiceSmall string) int {
	if s == end {
		return 1
	}
	if visited[s] {
		if isSmall(s) {
			if twiceSmall == "" {
				twiceSmall = s
			} else {
				// Already visited, dead end.
				return 0
			}
		}
	} else {
		defer func() {
			delete(visited, s)
		}()
	}
	visited[s] = true
	var n int
	for next := range m[s] {
		if next == start {
			continue
		}
		n += recursion(m, next, visited, twiceSmall)
	}
	return n
}

func main() {
	m := readMaps(input)
	fmt.Println(recursion(m, "start", make(stringSet), end))
	fmt.Println(recursion(m, "start", make(stringSet), ""))
}

const input = `
LA-sn
LA-mo
LA-zs
end-RD
sn-mo
end-zs
vx-start
mh-mo
mh-start
zs-JI
JQ-mo
zs-mo
start-JQ
rk-zs
mh-sn
mh-JQ
RD-mo
zs-JQ
vx-sn
RD-sn
vx-mh
JQ-vx
LA-end
JQ-sn
`
