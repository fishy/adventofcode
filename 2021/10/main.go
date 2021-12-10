package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strings"
)

var pairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var scores = map[rune]int{
	// part1 scores
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
	// part2 scores
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

type stack []rune

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s stack) empty() bool {
	return len(s) == 0
}

func (s *stack) pop() (r rune, ok bool) {
	if s.empty() {
		return 0, false
	}
	r = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r, true
}

func main() {
	var part1 int
	var part2 []int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var s stack
		var corrupt bool
		for _, r := range line {
			switch r {
			case '(', '[', '{', '<':
				s.push(r)
				continue
			}
			opening, ok := s.pop()
			if !ok {
				// should not happen
				log.Printf("line %q does not have enough opening chunks\n", line)
				corrupt = true
				break
			}
			closing := pairs[opening]
			if closing != r {
				// corrupt
				part1 += scores[r]
				corrupt = true
				break
			}
		}
		if corrupt {
			continue
		}
		var n int
		for {
			r, ok := s.pop()
			if !ok {
				break
			}
			n = n*5 + scores[r]
		}
		part2 = append(part2, n)
	}
	fmt.Println(part1)
	sort.Ints(part2)
	fmt.Println(part2[len(part2)/2])
}

const input = `
{{[<[{({<[{(<[<>{}]({}[])>)({<()<>>(<>[])}[[()<>]{<>[]}])}[{{({}<>)[(){}]}{{<>[]}(<><>)}}<{[<>{}]{[]<>}}>]]>[
{{<((<<<[{(([[<>[]]]<(()<>){[]{}}>])}]([{[<<[]()>[()[]]>{(()[]){<><>}}][{(<>())([]{})}{{<>{}}[{}[
{{[<{{<[[<[<<([]())>>[<<(){}>{<><>}><{<><>}<[]<>>>]]]{<([<<>[]>[[]<>]]{{()[]}})<{{[][]}(()<>)}((
{((<<[{{{{<{[(<>)]<[<>{}](<>{})>}>{(<[[]<>]([][])>[<<>()>([]<>)])<<[{}<>](<>[])><<[]<>>{{}<>}>>}}{(
{{{[{([[<<{{(<[]<>><[]>){[<>()][[]{}]}}{({[][]}(<>[]))}}{(((()())(<>{}))({<>()}(<><>)))([{()[]}[()[]]][{
{[<[{((((<([<<<>[]>>(<[]{}>]]([(<>())<()<>>]))>((<{<<><>><{}<>>}[{{}{}}{{}[]}]>(<{[]()}>[{<>()}(<
[(({({{<{[[([[()()]]{{()<>}[{}{}]}><(([]{}){[]{}}){(()<>)[()[]]}>]<<({[]()}{()()})({[]()}{<><
{<<<<({{([[{[[()<>]{[]()}]([(){}]({}()))}{{<{}<>>}{[{}()]<(){})}}]{[{[()()]<[][]>}]}]([({<{}>(<>())})
({(((([{{[[[{(<><>)({}<>)}]{{{[][]}{{}()}}{{<><>}(()<>)}}]][{[{{[]<>}([]()))]}[<({[]<>}<(){}>){<<><>>({}{}
<[[{[<{<{[{[({{}()}[{}<>])[([]())[<><>]]]<{{{}<>}[[]{}]}>}[{[{<>{}}([]{})]}])}>[{<[(({<>{}}([]()))(<{}<>><{}(
({[({{([(<[{(([]()))<{<>()}[()<>]>}](({{()()}[<>()]}<<<>[]>[<><>]>))>{(({<[][]>}<{[][]}[()[]]>)<[<
{{<(<[<{{<[{[(()[])([]())}<[<>()]<<>>>}{[[()[]]{()[]}][{{}}{{}()}]}]{<<({}<>)[{}{}]>(<{}[]>{()[]})>{(
<{<<<({[({[{{<<>()>{[]{}}}}{<<()<>>[[][]]>[<{}<>>[{}{}]]}}}{({{{()[]}[{}]}<((){}){()<>}>}<<[()<>]>[{[][]
<<[[<{<{(<[[(({}[])[()[]]){<[]()>}]<(<<>[]>(<>[]))[[[]()]<[]()>]}]{[[{<>[]}[{}]][<[]{}>({}<>
((({[{{{<[([({<>[]}[()()]}[[{}[]]([]<>)]]<(((){})(<><>))([()[]](<>()))>)]<(<([[]()]({}<>))[({}[])]>)<[<<()[
<<[{({[[([{{<<<>{}><[]()>><([][])[()[]]>}<([()[]]([]{})){(()){[][]}}}}]<{<<[<><>]<(){}>>(([]<>)<{}{}>)><
(<[([<<((<<[({[]{}}([]{}))({[]()}<()>)]{{<()<>><(){}>}{([]{})}}>[(((()()){()}){({}<>)[()()]}){{(())<<>[]>}}
[<{{[(<[(<[[({<>{}}<<>[]>){[<><>]{<><>}}]<<[<><>]({}())>[{{}[]}[[]{}]]>]{{{{<>()}}{<{}[]){[]{}}}}<
(<(<(({(<<(<[{{}{}}[[]<>]]>{[<()[]>(()<>)]({[]<>}}})>>)})[[<[[{[[<[][]>{()[]}]]}[[{<{}>{[][]}}]<{
(<{{{<[<<{{[[[{}<>]<[]>]](<[[]{}]((){})>((()[])))}<([{()<>}{<>{}})[[[]{}]{<>()}])>}>{<<[({{}{}}([]()))<([]
(<(({[{<{<<<(([][])(()()))(<{}()>)><({[]<>}[()[]])<<(){}>[{}{}]>>>[<[[[]<>]<<>()>][[(){}](<><>)]
[[[[[<{{{<<<{{<>()}(()[])}<{<>{}}<<>[]>>>([[{}<>][[]()]]({()()}(<>[])))>>}}[(([[((()()))[([]())[{}{}]]]{(<{
{{({([<([{[<{[<>[]]<<><>>}[(<><>)([][])]>({[()<>]<[]()>}{<{}()>})]<([<{}{}>([]())))>}{[<(([]<>)(
<<{<((<(<{{[([<>[]][{}<>]){({}())[()()]}]([({}())<()[]>]({{}}>)}[<<(<>{})>(<()[]>([][]))>(([<>[]]<()<>>)<[{
{[<[<<(<{{((<{{}()}<[][]>>[[(){}]{()()}])<{[<>[]][{}{}]}{[[]{}]}>){<<{()()}{<>()}>>({[<>[]]<(){}
{<[((<{<(<([(<()<>><[]{}>){{<>}[{}[]]}]{(<()()>((){}))({{}<>}[()[]])})[{({<>[]})<{<>{}}<[][]>
<(<{{<(([<<{<([]{})[<><>]>{<<><>><(){}>}}<<<[]{}>>{[<><>](<>{})}>>({<{()[]}({}())><[{}[]]<(){}>>})>[{<[{{
<<(([<[<({{{([[]<>]<{}{}>)}[{(<>[])<[][]>}{{{}{}}(<>())}]}})<({[[<<><>><<>()>][{<>[]}({}())]]<<[{}()
<<<{{[<({<<{{{[]{}}{{}[]}}{{()[]}<{}{}>}}>{({{()<>}[<>[]]}({[]<>}(()()))){<<[]<>>{[][]}><<
[([[{[(<[{<<[((){})<()>]><{(()())([]())}{<()<>]<[][]>}>>[{{((){}){<>[]}}<[()()](<><>)>}[[([
(<{(<{<<{{<<[<[]<>><[][]>]{{{}{}}}>([[()[]]{[]<>}]<[()[]]<<><>>>)>[({[[]{}]({}[])}{{(){}}{<><
[{<<<<({{[({({[]{}}[<>{}])[<<>{}>({}<>)]}{{({})}({[]{}}({}{}))}){<<{{}[]}<{}<>>>[{<>[]}<[][]>]
{{<<({<{<[[([[()<>][[]{}>]((()[])))]([{((){}){()[]}}<<<>{}>{{}<>}>])]>}>}){([[[[(<([<>()])<<()()><[]()>>>
[<<[(<<<<<{([{<>}]<({}{})<[]<>>>)<[[(){}]<<><>>]{{{}[]}{{}[]}}>}<{<{[]<>}([][])><[{}()]{{}[]}>}(<{<><>}{<>()
(([(([<{(<{{<[()[]][{}[]]><<()>([]())>}[(<()<>>({}()))<{[]()}{[]()}>]}[({{[]{}}}<{<><>}(()[])>)<[{<>[]}{<>[]
([<{<{<[{({<{[[]()](<><>)}(<<>{}>{{}{}})><<<[]<>>{{}()}>{<(){}>[[]()]}>}[{[{{}}([]{})]]])(({{{
[<{[{[{<[{{[(([]<>)<()<>>)(<()()><()()>)][<([]<>)<<>{}>>[(<>{}){(){}}]]}{(<{()<>}{[]{}}>(((){
<(<{[[{<{<{(<(<>[])[{}<>]>(<{}<>>[{}[]]))[{{<>[]}([]{})}[(<>()){<>()}]]}<[((<>){()[]})[({}<>)(()[]
{{(<[{<{<(<[<(<>[])([][])>[([][])]][<({}<>)<<><>>><[()<>]{<>{}}>]>)[(([[[]]]<<()[]>>)<[<()[]>[()[]]]([()()]
{<({{<[([([{[[<><>]{{}{}}]}<{[{}<>]({}{})}<<()>[<>[]]>>])<(<{([]{}){[]()}}[<{}()>{{}}]>(<(<
{<[{{{{{<[{(<({}{})[()]>[([]{}){<>()}])[{[<>[]](<>{})}<<<>[]>[()<>]>]}]>([({(<()[]>[{}[]]){{[][
{<{{<{[[{[<<{(<>()){(){}}}>{({[][]}[[]()])([{}{}](()))}><{{<<>[]>{{}{}>}<([]<>)<<>()>>}({[[]{}][[]<>]}
<(<[(((([{<<{{()[]}[()[]]}({[]<>}<(){}>)>({{[]{}}<<>>}{[<>[]]<()<>>})><<[<[]<>>{<><>}]({[]{}
<([[[[<(((({[({}()){<>{}}][{<>}<()[]>]}(([()[]]<{}<>>)[<<>()>(<>[])]))<[<[[]()]{[]<>}>([<>{}]<[]{}>)](<(()[
([<{[{([[{[{(<(){}>(<>())){([]())[()()]}}([(<>{})<[]<>>]{<{}()>[[]()]})]<(([<>{}](<>[]))[<{}
[[<<{[{[<[{(<{{}<>}<()<>>>([[]{}]{<><>}))<<[{}()]{<>[]}>[{{}[]}{{}<>}]>}]>[{{{<<()<>>{{}[]}
((<[{<({([{{(([]{})(<>()))}[[<[]<>>][[[]<>][()()]]]}[{[({})({}[])]<[[]()][[]()}>}[{[()[]]<<>
{[(<((<{(<[[{([]()){{}()}}][{<[][]>[<>()]}{[[]][[]{}]}]][{{<()()>{<><>}}<{{}[]}{<>()}>}]>(
{{[<(({<[{{[{(<>{})([]<>)}{{()<>}[[][]]}](((<>{})[{}[]])([{}<>]{{}[]}))}[<<[()()][<><>]><[
{(({[[([[({{[({}{})[{}{}]][<()()>({}{})]}{(<<>{}>{()[]}){[()]({}[])}>}){((<<{}{}>{{}{}}>[<()<
[<{(<[([(((<{<<>[]>[<>()]}({[]()}({}()))>(({<>{}}{[]})<{()<>}(<><>)>)))([<<<{}()>{<>{}}>[<[
[{{[(<{[<([[<{<>{}}<(){}>>]{{(())}((<>[]))}]({({{}<>}[(){}])>))({<{(<><>)[<>{}]}<<[][]><<>{}>>>}((
([[(({({(<([({{}}[<>()])<({}{})>]((([]()>[[][]])))>)})(<<<[<{[<>](()())}[(()[])(()())]>]{([[[][]]{[]{}}
{([({{<([[<<[[{}<>]<()()>]><{{[]}(<><>)}>><<{<{}()>[<><>]}>[[{()}{{}{}}]]>]])[({([<<()<>>([]<
<{{{(({{[(((<([][])[{}[]]>((<>())[{}{}]))[<[[]()](<>{})>[(()()){{}[]}]]))][<[[[<<>[])(<>[])][
<<<[[[<[[[<[[({}[])<[]<>>]([{}<>][{}[]])]><(({{}()}))({[()]<[]<>>}<[<>{}]({}{})>)]]<[[<{[]
({{[<[(<{<{{([<>[]])(([][]){<><>})}}{(({[]{}})<{{}}[<>{}]>){[<[]<>>{()<>}]}}>}{<(<<[<>[]](<><>)>[<()>{()
<<[(<[{[[(({{(<>{})<{}{}>}}))<<(((()[]){(){}})([[]<>]<<>{}>})<[<{}{}>]([()]<<>[]>)>>>]]<{(((
({({(<<{<(<([{<>[]}<<><>>])>[(<[<>()](<>())>[<[]{}><()()>]){([[]{}])[[{}{}]{[]}]}])<[((<(){}>){([]
{{[<<[{[<[(((({}<>){{}{}})<<<>[]>[{}[]]>)[{<{}[]>{<>()}}[<{}[]>[<>[]]]])<(<[[]()>><(<>{})<()<
<(<(<<[<({[[(({}{})[[][]])<[{}{}][<>[]]>]]{<[{<>[]}<[]<>>]{{<>{}}{[]{}}}>{<[<><>](<>())>[[<>
{(((({{<<<[[{[[][]>{<>{}}}{[{}{}]{()<>}}]{{[[]]}[[[][]][()()]]}]({{[[]<>][[]<>]}<{{}{}}({}
({([{<{(<{({[([]())[()[]]]{([]<>){<>[])}}[([[]{}](()<>))<<()<>><{}[]>>])<((<()[]>{<>[]}){{{}()
[{<{[{[<[{(<<({}{})<()[])>(<()<>>(()))>{((<>())[{}()])<{<><>}<[]<>>>}){{(<{}{}>{<>[]})[{{}}[[]{}]]}<[
{[{<[[[[[(<{{[()[]]{{}<>}}{([][])[(){}]}}>{{<{<>()>([]<>)>[({}{})]}{{<[]>}[[[]<>]([]<>)]}})<
{[<{{({((<<(<({}<>)<<>[]>>({{}}[()()]))[[{<><>}[[]{}]]]>>){{<{{([]{})[<><>]}}>(([<{}{}><(){}>]{({}<>)
{{{<[((({(<{{({}<>)[<>[]]}[([]<>){()[]}]}{(<[]<>>[[]{}])}>){<<{{[]()}[<>[]]}>[<<[]{}>[(){}]>]>
[{{[<{[{({{<<(())<()<>}>(<{}{}><<>[]>)>[[<{}>([]{})]<<<>[]><{}[]>>]}((<{{}<>}{()()}>([()()]{{}<>})))
({<[<[<[<<{<[<{}[]>]>{<{{}<>}{()<>}>([()<>][<><>])}}{{[<<>()>]{{[][]}{<>()}}}{<(<>[])<<><>>><(<>{})<<><
{{{([[<[[(((<(<>[])<<>[]>>{<{}<>>{()}}><[{(){}}({}<>)][[<>[]][{}{}]]>)([({[][]}[[]<>]){<(){}>}]
{(<{{{<[<[[[(<<>[]>([][]))]]]>]>}}([((<<[{(<{}>(<>{})){[()()]{[][]}}}{{([]<>><{}<>>}<{{}()}(
([(<[[((([[((((){})({}()))({[]()}{<>()}))[([<><>])[[<>[]]<{}[]>]]]]<<({<(){}>{{}<>}}({(){}}[{}<>]))([(
{[[[[{{[[{{<(([][]))(([]<>){[]{}})><{<{}{}>{[]<>}}<{{}<>}(()())>>}([{<[][]>{[][]}}<[<>]{<>[]}>]<{<()()>[()
<({{[((<<<([<({}())[()<>]>(<()<>>[<><>])][{(<><>)[<>{}]}([<>{}](<>()))])(<[{(){}}[[]{}]>>)
<((({([{<<<[((()[])[<>[]]){[<>{}](()<>)}]{[<{}[]>]<([]<>){<>{}}>}>[<<[{}[]]<(){}>>>]>{<(((<>[]){<>()}){{<
[<(<<({<<[<{{<[]()>{{}()}}[{{}{}}{[][]}]}>[{(({}[])({}{}))<((){}){{}[]}>}{(([]{})([]{}))}]][({<<()>(<>())>{{<
({<({(([<{[(({(){}}{<><>}){({}[])<()()>})]<<{[[]()]<[][]>}>[{<<>[]>[[]{}]}[([]<>]]]>}[<{({{}()})((()<>)({}())
{({[[{({{<<(<[[]{}]{<>[]}>([[]<>]))((<{}()>{()()})<<{}()><(){}>>}>{{[{[]{}}<[]()>](({}[])[<>{}])}<{[<>[]]
[([<({{<<([<[[{}{}]((){})]<(<>[])>>][{[[<>[]]{<>()}]{<<>{}>({}())}}({[<>()]{<>[])}{<<>()>({}())})]
[[([[(<<({([{<[]()>({})}[<()()>{()[]}]][<[{}<>>{[]{}}>])<[(<[]()>)({<>[]}<[]{}>)]>})>{[<<{{{{}[]}[{}<
[<[{<([(<[([[([]<>){()()}][[[]<>]{{}<>}]]{[[<><>]][<<>{}><{}{}>]}){{(([])[()<>])(<<>{}><{}[]>)}}]({<((<>
{(<(<{<<({{[[<()()><{}[]>]{<()<>>[{}{}]}][(<()()>({}{}))({<><>}<{}[]>)]}{<<([]{}){[]()}>[((){})[()
{({[[<[<<{[(<([][])><[<><>][{}<>]>)<<{<>()}(<>[])>{({}{})[()]}>]{[{<()[]>[(){}]}[<[]()>[(){}]]]{<
{(<[[[<(<{{{[{()()}[<>]](({}[])[[][]])}({([][])[(){}]}{{[]{}}{()<>}}}}}<{[(({}<>)<()()>){<(){}>[{}<>]}](<<<
{{{{{[[[<[(<<([][])<<>()>>[{{}{}}{[]<>}]><{<{}{}>({}())}>)(({{<>{}}}))][({<<[]{}>>{<()()>}
(([[({{[<<[{{({}<>)(<><>)>}([<[]<>>{{}()}]({[]()}[(){}]))]><[(((()<>)<<>[]>))[(<()[]>[()<>]){[()()][{}()]}]
[({[[<{<{((<<([][])<()()>>{([]{})([]{})}>(<[[]<>](())><{{}[]>({}[])>))({<[{}]<()[]>><({}<>)([]<>)>}<[<()()>
{([<<({<{[{[{([]<>){()()}}][[(<>[])<()()>]({(){}})]}[{[{[][]}<{}[]>]{<()()>>}{<{[]()}<<>()>>[({}<>)[(
[[((<{(([<((<(<>)[()]>(<{}{}>)){<{<><>}<()[]>>([<>()]({}[]))})<{<(()()){()[]}>}[({{}{}}([]()))<{
<{[(<{({<(<[{[{}[]]}(({}()){{}<>})]{((()[])[[]<>])<[()<>]>}>((<{<><>}([]<>)>))){[[({(){}}<<><>>)<{{
({([{[<{(<({<[()<>]<{}()>>[[{}{}][()()]]}[(<[]<>><{}{}>)<{{}<>}[<>()]>])>{([{<[]{}><{}<>>}<({}{}){
<([[<<[<{({<([()<>](()<>))[{<>())<{}{}>]><[{[]<>}](([]<>)([]<>))>})[<[{((){})[()<>]}{[{}<>]<<><>>}]>]}><
<{<[((<{[{([{(()<>)([][])}][(({}[])(()[])){{()[]}<<><>>}])<{[{[]()}<[]{}>]}<{({}{})<<><>>}<{()()}(<>{})>>>}
{{<[<({(<[[<{(())}([<><>]([]<>))>(({<>()}<()()>){{<>()}{[]{}}})]<{([[]<>]{[]()})<<[]()>{()<>)>}<{{[
[<([{{[[<(({{(<>())<<>()>}<[[]<>][[]{}]>}<(<{}()>(<>[]))[{[]()}[{}<>]}>)[[{[[]<>](<>{})}[[[]{}][[]{}]
{[{([<({<{[{<(<><>)([][])>([[]<>](()()))}<{<[]<>>}[[{}()]{<>[]}]>](<<[<>()]<<>()>>([[]<>])>{{[{}{}][<>[]]}{
{{<<<<((({(([{{}()}(<>{})]{{<>()}[()]})<[(()()){<>[]}]{<<>{}>}>)([({[]{}}<{}<>>)]<[([][]){(){})]([[]
({[(((<[([[[(({}<>)<()<>>)<{<>()}>]]{{(<{}[]>[<>[]])<({}<>)>}[{<<>>}{[[]{}][{}()]}]}])<<(<{{<>{}}[{}[]
`
