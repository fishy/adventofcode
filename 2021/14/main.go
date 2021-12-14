package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

const n = 40

type countsMap map[byte]int

func (m countsMap) merge(m2 countsMap) {
	for b, c := range m2 {
		m[b] += c
	}
}

func (m countsMap) toCounts() counts {
	cc := make(counts, 0, len(m))
	for b, c := range m {
		cc = append(cc, count{
			b: b,
			c: c,
		})
	}
	sort.Sort(cc)
	return cc
}

type count struct {
	b byte
	c int
}

func (c count) String() string {
	return fmt.Sprintf("%s:%d", []byte{c.b}, c.c)
}

type counts []count

func (c counts) Len() int           { return len(c) }
func (c counts) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c counts) Less(i, j int) bool { return c[i].c < c[j].c }

func (c counts) result() int {
	return c[len(c)-1].c - c[0].c
}

func makeKey(a, b byte) string {
	return string([]byte{a, b})
}

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	scanner.Scan()
	start := strings.TrimSpace(scanner.Text())
	scanner.Scan()
	rules := make(map[string]byte)
	for scanner.Scan() {
		split := strings.Split(strings.TrimSpace(scanner.Text()), " -> ")
		if len(split) != 2 {
			continue
		}
		from := split[0]
		to := split[1][0]
		rules[from] = to
	}

	cache := make(map[string][]countsMap)
	for a := byte('A'); a <= byte('Z'); a++ {
		for b := byte('A'); b <= byte('Z'); b++ {
			key := makeKey(a, b)
			cache[key] = make([]countsMap, n+1)
			v := make(countsMap)
			v[a]++
			v[b]++
			cache[key][0] = v
		}
	}
	for s := 1; s <= n; s++ {
		for a := byte('A'); a <= byte('Z'); a++ {
			for b := byte('A'); b <= byte('Z'); b++ {
				key := makeKey(a, b)
				if to, ok := rules[key]; ok {
					v := make(countsMap)
					v.merge(cache[makeKey(a, to)][s-1])
					v.merge(cache[makeKey(to, b)][s-1])
					v[to]-- // overlaps and double counted
					cache[key][s] = v
				} else {
					v := make(countsMap)
					v[a]++
					v[b]++
					cache[key][s] = v
				}
			}
		}
	}

	for nn := 1; nn <= n; nn++ {
		m := make(countsMap)
		m[start[0]]++
		for i := 1; i < len(start); i++ {
			key := makeKey(start[i-1], start[i])
			v := cache[key][nn]
			m.merge(v)
			m[start[i-1]]-- // overlaps and double counted
		}
		fmt.Println(nn, m.toCounts().result(), m.toCounts())
	}
}

const input = `
CVKKFSSNNHNPSPPKBHPB

OF -> S
VO -> F
BP -> S
FC -> S
PN -> K
HC -> P
PP -> N
FK -> V
KN -> C
BO -> O
KS -> B
FF -> S
KC -> B
FV -> C
VF -> N
HS -> H
OS -> F
VC -> S
VP -> P
BC -> O
HF -> F
HO -> F
PC -> B
CC -> K
NB -> N
KK -> N
KP -> V
BH -> H
BF -> O
OB -> F
VK -> P
FB -> O
NP -> B
CB -> C
PS -> S
KO -> V
SP -> C
BK -> O
NN -> O
OC -> F
VB -> B
ON -> K
NK -> B
CK -> H
NH -> N
CV -> C
PF -> P
PV -> V
CP -> N
FP -> N
SB -> B
SN -> N
KF -> F
HP -> S
BN -> V
NF -> B
PO -> O
CH -> O
VV -> S
OV -> V
SF -> P
BV -> S
FH -> V
CN -> H
VH -> V
HB -> B
FN -> P
OH -> S
SK -> H
OP -> H
VN -> V
HN -> P
BS -> S
CF -> B
PB -> H
SS -> K
NV -> P
FS -> N
CS -> O
OK -> B
CO -> O
VS -> F
OO -> B
NO -> H
SO -> F
HH -> K
FO -> H
SH -> O
HV -> B
SV -> N
PH -> F
BB -> P
KV -> B
KB -> H
KH -> N
NC -> P
SC -> S
PK -> B
NS -> V
HK -> B
`
