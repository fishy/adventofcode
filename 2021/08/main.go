package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func overlaps(a, b string) int {
	var n int
	for _, r := range a {
		if strings.ContainsRune(b, r) {
			n++
		}
	}
	return n
}

func fields(s string) []string {
	fields := strings.Fields(s)
	for i, f := range fields {
		b := []byte(f)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		fields[i] = string(b)
	}
	return fields
}

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var part1 int
	var part2 int
	for scanner.Scan() {
		split := strings.Split(strings.TrimSpace(scanner.Text()), " | ")

		f := fields(split[0])
		// sort by length
		sort.Slice(f, func(i, j int) bool {
			return len(f[i]) < len(f[j])
		})
		m := make(map[string]int)
		reverse := make([]string, 10)
		store := func(index, n int) {
			m[f[index]] = n
			reverse[n] = f[index]
		}
		store(0, 1)               // len: 2
		store(1, 7)               // len: 3
		store(2, 4)               // len: 4
		store(9, 8)               // len: 7
		for i := 6; i <= 8; i++ { // len: 6, 0/6/9
			switch {
			case overlaps(f[i], reverse[1]) == 1:
				// 0 and 9 have 2 overlaps with 1, 6 has 1 overlap with 1
				store(i, 6)
			case overlaps(f[i], reverse[4]) == 4:
				// 0 and 6 have 3 iverlaps with 4, 9 has 4 overlaps with 4
				store(i, 9)
			default:
				store(i, 0)
			}
		}
		for i := 3; i <= 5; i++ { // len: 5, 2/3/5
			switch {
			case overlaps(f[i], reverse[6]) == 5:
				store(i, 5)
			case overlaps(f[i], reverse[1]) == 2:
				store(i, 3)
			default:
				store(i, 2)
			}
		}

		f = fields(split[1])
		for _, f := range f {
			switch len(f) {
			case 2, 3, 4, 7:
				part1++
			}
		}
		part2 += m[f[0]] * 1000
		part2 += m[f[1]] * 100
		part2 += m[f[2]] * 10
		part2 += m[f[3]]
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

const input = `fcdeba edcbag decab adcefg acdfb gdcfb acf fabe fa eacfgbd | aefb cfa acf cdabf
adbec fabeg fgda gafedb fadeb cdebgf cfaebdg fd bdf cgfbae | ebdfga fbd bdagcef dfb
cagefd fegabc gbde fcagebd bcedf gefcd bec cefbdg dfabc eb | gdeb dgbe defcb ebc
cegbf acfeb cafde ab abfecd ecagfd edbcga adbf ecfadgb abc | degafc gefbc ab adcef
bc gdeca bec fcab cegfdb edbgcfa gbefda fedba dceba badecf | fdbeacg dgbaef bec cbe
caebdg gfc gaceb adcbf bgfac gf eafg dfbegc abfcge ebdcgaf | bcfag bdgface fegcba gdbeca
fedga fgebdc dfb edcagb edcab afbc fdecba dbaef fb aefbgcd | bf bcade bf cafb
cb cfgea badcge gdfabce bec fedbg egdafb cgdfbe cdfb cbgef | dbfc gecbfd fbdc ceb
cbfdag fcbgd afbgced cgfab ebfgad agcbef dfg fecdb dcga gd | gacfedb cbfgd dgf gdfbc
abgfd dgcbe dfbgce acbged eab ae daebg edac gdbfcae cgbeaf | bcedg eba eab edac
gea ag cfaeg agcd befagd fcdae adcgfe ecbgafd cdbfea gefbc | bdeafc defacg gae dgcaef
bfeca adfbce febd beagc cdeafg fgacedb efc dcbagf abcdf fe | agedfc cefba ef fgcbeda
cbeafd gfabd dabecg cfeg dfc gcadf cgaed efacgdb cf gfcdae | cefg fgdca decagfb ebfadc
fg edabf gbedf gedbfa ebafcg egbdacf fgad cedgb cdebfa egf | bgdec egdfab gf acfgbe
cfe bacge ecfabd dcbfega defa cdfba defbcg fe feabc dbfcag | cabge ef fadbc eafd
ag gadb edacfb gcaefd eafbd dbeafg efcgdba fag egcbf gfbea | agf gbfcdea cbaefd ag
adegb gbdfe bcdafge gdfbce fcgeab geafbd ga dgfa age ecadb | ag gae gfda eag
beagdfc faedc caegbf fgeba caefg cg ceg dbgfae aebgcd cfgb | ceg bgefca egc gce
dcaf cd cdeabf cefab fgbde efbdc bcd eadcgb gebcadf caegbf | eacfgb cd feacb eafdcbg
begadcf dbf ecdbg aedfc fbae cdfbga dcbaef bf deacfg cfbde | feba eafb bf cdfbe
fbgced gcedba gbc fgadc gafcdbe cebf egfbd dbafge cdfgb cb | gcb begfd dfbaecg dbgfc
dbcega beg gb cfdeba gaceb gfeca gdab fcedbg beacd cfagedb | cfbgade bg bg fdgcbe
abdcgf agebc dacbg fbadce fcaged gd fdcebga cfdab gbdf dag | gd cefgdba dgbf fbdg
gb ecbadf gbad edfgcb bgf dbcaf afbcg cgfea cdfgbae dcfbag | gb bg bgfceda dafcbe
gaefd cg adbfce bdfecg fcbdgea cbafe agefc cgf abcegf gabc | gcf acedfb febgca cafbdeg
deacgfb bfeacg ebgcdf dc cfadgb abegd gbcfe bcgde gdc dfec | gcd begad befcg dgc
cdge dfebga gcbef abdcfe ec gebcdf cbe cedabfg fgdeb bfgca | ebgdfca ce fedbgca fdaecb
bfde aeb aebfcgd gabdfc dgafb bfage be dfgeab fcega gdaceb | edbgac eab eba gaefc
dcgae afebcg eabcgdf bg agb febda dcbg eagcdb abged edfcga | egcad cgdb deagb gb
fdceag dbgca beg cegaf begcfa baef eb abgce adgebfc fbgdec | be eb afeb eb
beadgc gdcf ebacf acdefg bfgdeca afcge dafgeb gf fga daegc | cdgf fg aecfb cbfea
gfca gfdcabe ecg gc abecdf fbcea fagceb ebfdg efcbg dabegc | dfgbe gdebf cefbag ecafb
feadg gbfdea acdefb gaebf cfged ad beagfcd dae ecagbf bagd | ad cgabfe cfdaeb dbga
bcdge adcfeb cadbge bgdafe dbg gcedf acbg ebacd gb dcgeabf | abcgdef ebgcda bdgec bdcaeg
cfga gacde fdgea cgefbd fgaedc cbadfge gc edbac gebdfa dgc | gcd cgaf eagfdb eagdf
bg fdcega bgf dbgfcea gdcfe gbcaef bfged dbcg cdbgfe dfbae | gcbd ecfdg fbgde bfg
gcdabf bgafe dbe eafbdc ecagfdb bedaf gadebc cdfe fdcab ed | dbfea gabdcef de gbdeca
bcdaf gacfd deacgf bd gbdc adb aecfb gfcbad fdagbe edgcfab | eagbdf bgcd afbdge dab
gface cbea cfegad bfa ab fdgbc acbgf ebdfga dbcafge cebfga | eagbdcf dacegbf ab ceab
eabdf bcgafe gfedbac agcdbe gedcb abecd gcad ac bac cedbfg | deacb cab bdcfge dcfgaeb
cedgf cbgdef gcafde fdae ebgfac edagc cae ea bdefagc bcdag | defcg gebdcf aec bgdac
cdbae edf afbegd fgdebc aegfcb aebdgfc febcg dcfg df ecdfb | debagf df fcgd eadfgb
adc fedcbga ad adgb ecagdb dcebg cgead dfgbec eafgc afcdbe | cda bgda ebdafc egbcfd
cdfeb eafcdg acfde bdac edb gcfeb eagdbf db afdecb ebacgfd | abcd abcd db gbefc
cdefa bf dcbag cdabf bcefgda ebcadg debfga afb bfcg cgabdf | bf gcabd bdagfec gfcb
ecb eabgdc cgaeb bcagfd gfedcb beacdgf aedc ec beafg dgacb | gcdfba fgcbead ce ce
fcebg gcfdeb bdgaec bfdg facged cbg egdcf facbe bg badfgce | fdgb bfcge fgbd egbcfd
ebfcg bd adgfce acgbdef fegdb fbd gdba gbeafd fbacde agfde | cfbge db gdacfbe abfgde
bgefc gcde afgcbe dgcfeb eadcgbf dg fbdeg gafcdb bdg fbade | dcgabf bcfdga egcd gd
fbgadc gcb fcbagde bgaef bc cbdf fcgade cagbf bgaced gfacd | bcg dcebga cdgfa bcdf
cgebaf egdac aefbdc cbafg becag bfge eb adbfegc ebc bdagcf | be bec egdca cabdgf
fgd afgdce dg dage decgfb fcdag bcafgde gfcae egcafb fabcd | fcbdaeg gfcdbe agde fgd
cbfga abg dgbfea dbafegc fceagb ba cfgea bcea gbdfc fgdeca | bcea bafdceg egafc cbdgf
ebgcda acfegb gbdc befdcag fceadg afdbe cegad ebdga ebg bg | dfebgac dfeab gb gb
bacdf cgdfae abc ecdaf agefcb ebdc dfbag cb cgfbdae daefcb | bc gcdafeb bca gcfeda
fegabd egbfa gfa debag decfagb becfa cgadbf defg fg bgdeca | gfde agdcfb bgdcaf bafeg
gf cbeafgd gedbaf cbfdeg egcfb cfbae gef cfgd egdcb gbdeca | gf fg gef dagcbe
bac agcbd ba edcga gdfbce abfd egfbca fbgcd bacgdf becadgf | cfbeagd bdfa adbcg badf
gd cdgeb aecfgb egcabd dcbagf ceagb ebcfd adge dgb gfbdcea | ebfdc deag ebfcga fabcgd
dcbae agec dacebfg ebdag beacgd afcbd ced decfbg ec dfgeba | gdebcaf dbaegf ce aceg
dagefb fbecgda adgec afcgde gacf bcdef fa efa cfead geadbc | dfebc bafcdge efa aef
gbdae cgad bfcde ac dceabg bdgefa bfcaedg cab debca acgefb | ebacdgf abcgedf cab cefdgab
dbegf adegbc gefdba acbgfd ebfcg gfbad deb edfa ed agedbfc | bfdgcae debgaf bdcgeaf gebcf
fcgeb dagebc acfeg aecdfbg gedbaf bcgefd dfbeg bc bfcd cgb | cb efbgd cbg bc
dba gcdbfe egabdc aedbfc fbced bgacf cfbdeag ad faed cbafd | cbfad ad ad cbfga
edbcag ea ade cfbed fcdag gdaefc eafg faecd fdceagb gfbdca | bgadcef ae eafg ea
eafbd gfbeda bc dacb fcbega bedcf debfacg edgfc cebafd ceb | dgbecaf gcedf bdca cgabfe
bacegf deca gefbad adebcg gbcae gbacedf dgcbe ed cgdbf edb | bde ebd efadbg ceda
dca acdfe bfegac fbeadcg deab dcgfe da acdbgf fbeca edcabf | deacbf dca cdbfage dac
febg ecfdab abe deagf efgdba gcabd fdegca eb egfacbd dgaeb | cfedga abcefd efagdc bea
fed feadc ef aecbdf gadbce cgdaf gdcaefb dbgefa bfce debca | dafec cdaeb fbec ef
fecbdg cefdb gcbe cabdfe cdbgf gb cebafgd agfdc dbg dafgeb | ecbg cgfad cgdaf agdfc
feacg cgfeda fbcagd fcbge cfa baegdc acegd gfadebc eafd fa | begfc fa afc dbagcf
afgcbe cfg cf cfde cgfbd gbdfeca gbedac debgc egcfbd gfdab | cf cgfebd cfg fbegcad
adecfgb bgdcea cbegaf efcbad ebg ge bdeac bcdeg fbcdg dega | dbcgf agde bfcega bge
cg cbag dcg gdcfae dbfgcea dagefb cgabfd fbgda dgbcf fecbd | gdafeb dgc gc dgfabce
cdfea dfbcg ga gac eafg adefcg acfdbeg gadfc cfedba edgbac | eafg cag bfgcd cdbgf
dbaefcg cfdea eadgf dcba cfbaed befcd ace ca abgecf cfdebg | gbdface cegfdb edagf ecdbf
dcf fcagd gcbf edbgcfa cf baegfd ecgad fedabc gfadb dcagfb | bgfc acgde cf fgbacd
bg fedgc bfcega dbaegc gfacebd decbg gdba bcafed bedca cbg | dcefg bg bg bg
geadcf ecfbd bagedc afcdb gcfde cefgdba bce begf be fdcebg | fgeb fegb cgedf ebc
gbdcaf acgfed gc dafcb cag fbgea cbfga cbfdea bgdc dcfegab | bcgd aedgfc gfeba cg
efdga afe ea gacbdf gbafd ageb dfaceb aedbfg dgfec bagfcde | fegdc dagfe ea egba
bfg gf edabg fecbag abfgdec gacf acbfe bdecfg facbde fbega | fgaebc befcda fg abfge
fceadg beacf cgbdae bgfd cedbf ecfdbag fgdceb edf bdecg df | cfabedg fed dfe gdbf
fdeac dgfaeb dgfb gfadceb ceabdg gcebaf bf baf daefb aegdb | bdfg gdeab dbgacfe acfed
ge cbfeda eag fagde fedab ceagdbf fbge ebcagd dcfga dagefb | cdfga ebfcda cadgbfe dfcag
egcbadf ce afdgc cdefa abedf ecfdab efc bdegfa begafc ecbd | bcafde ecf ce ecf
ebc gdacbfe gcefdb abefd dgec ec fbgdc fcadgb fdceb gfaceb | cged fadbceg ce bec
ace egbdac acdef cafbgd efgdc cfabd ea feab bdaefc ebdfgac | adebcg cfbad ea dcfaebg
fc bagcfed fdabg acefgb bcf bcgead egbca dfgceb fabgc acef | cfb beadgc begca fc
cdefa agfecd fbcaed bgdeac gaefc ag eag cgbfe gafd cgbadef | gadf age ag gecbf
dfgcea dgac egabdfc cfdbae edcgf fegbd cge cegbaf cg adfce | cbdaef gec cfedg cagd
gcdfb dcbfa bcfgae gdbfe gfabed bafcdge cg cbg decg bfcgde | cg dabegf gc gecd
bf dcefb fbe edagfb gefcd ebcad cbaf dcfeba ecagbd edbgcfa | bfeadc baefdg efdgc cfab
debga ebgafd gcdafb cd cgd decabg ceagd egbfdca bcde aefgc | cd acbdgf ebcd cbde
dfbaegc fgcae bgcead gfaedb cdefgb adfge fd degab def fbda | fcdgbe gbcaed dcegab febcdg
gde gabcde afcbeg dg ebdaf degbfca cfgd bdegcf fgceb fdgbe | fgedbc dbagec gde eadbfgc
aebdcgf bd ebdag agceb fegad gdeafb gdb fbed cgbfda cgefda | db fbaecdg eadbg befd
cebdf bgecda efbdacg af aecfgd adf aegdbf abdef ebadg gfab | bdegaf gafb fbegad abdecgf
adcb daf afdgce fgdbca gfbdc da bfadg dfebgc gbafe fdacbge | febga dfa dcafeg gafcdb
bgacf dabg ebadcfg bfgdca ceadbf fecgd db dcb egfcab gfcdb | cabegf cdb cdfge cbd
bg bedfc bedafc cebdfg fedgcba beg gdbcae adefg dgfbe gcbf | edbcaf fgbc egb cgbf
abegd ecfbga cadfbg gbdecf adbgcfe ac gabec faec ebfcg abc | ceaf dgbecf bdgcfa efca
degba begdaf gbdcea ecgd gceadfb bgadfc ecdba bfeca cd dbc | cadgbef gbdea dgec cdge
ceabgd fcbge cdeba cdf fgecdab fd afcgdb dbcfea dbefc afed | dcebf df fd df
ebd fgbdc cfbe bgdcef fegad be befdg acbgdef cegdab agcdfb | bed gfedb cefb fgbcd
deb gbcd ecfbgd cfedga ecbdf db edcgf afbdge cefba eabcgfd | bde bd cfeba gfadce
afbec bcefd afgec cfaegd fgbace ba deagbf abe bcag dafebcg | cefba fgaedc fbgaed bfced
fabed cadfbeg ecf afedc gfdac agec facgdb ec cdefga dfgcbe | agbdcef efc dbfeagc agdfc
gc baedg gaebdc edbfga dabfc dbagc gdec gbc dcfegba cfabeg | agdeb eabdg dbagc gdfabec
dfegab fgd gadc agfdcb gd cfdbg bcagf edbcf fcgeadb aefcgb | cfabg dgac dg gd
cagfbde dcbfge bgefd ebdfag af caedb daf eabdf bfgacd egfa | ebcad adegbf dcaeb edgfbac
fecd gbadf baecdg ebfacg fbedcag ef gdafec egdaf gecda afe | fe fe efdc defcabg
fgbdac egbca gecfda fgadc gfb befgda bcafg dcbaegf bdfc bf | bf fb egcab caefgd
fb efcdag gebcd faedg dagfeb cfgadb cfgbdae bfae fdb degbf | gbdef afbe gecbfad adgef
gdcf abegdcf dg dga cdefba bgcae afegdb deafc eafcdg dcega | cegab eacbg dga gedac
caegfb gfcad ebcdfg cbdfega gdcbf bf cbf edgcb bdfe gedacb | dgcaf gcdfb bfed edbcafg
bdcfe fgdbaec ecgafd befgad eafbd adf gabd fbgae ad gfcaeb | bgda ad ad fcedag
agecfd bcegfa fbecgad fgcad gefdc fcdegb fbgad cga aedc ac | cdfag cegdf gdceaf cdbgfe
bgd befgd fdbceag db dcgfe cbde dfgcba gafecd dcbgfe gfabe | cdbe egdcbf egfcd dbcegf
ed gdefc dcegfb fecbg fcdag edg bacedg edbf becafg adecfgb | ed ed gbfcae bgfdce
fbdagc gabdf ecdabfg caegbd ebgaf cfabde cgfd agd gd dbafc | fbdcae dfcaegb gad gabfe
fbeacg fbcaged fdaeg fde becfgd agfbe de eabd abedgf acfdg | ed dgebafc agefbd fgdac
deagf efcad dga dfgeca dcaebf gecd dg faebdgc agefb bdgcaf | dg gbeaf fcbdgea gd
fb gaedcfb gebda gabdec dagbfe fbdae aecdf fbd gcebdf gbfa | afecbdg bdecga abged bdecgfa
gb bgedcfa adgefc dbgef deafg eabdgf ecdfb acgbde afgb ebg | gb afgb cefgbda gb
gb edacfgb cdfegb bgc cgaefd bdga fbeac gacfd fgdacb bacgf | gdacf gdba abdg bfcga
ce befcag ebdca adgcb aec dgebaf bdefca badfe edcf bfegacd | cbfaed begcfa abegdcf cfde
ceabdg bceafdg gabfed caef ac aegfb bafcg bfaceg cdbfg cga | ac aedfgb gfecba ca
gdebf aecf bcgfda fcd gcebda fbcde ecbad gebfdac cfaebd cf | cdf dcbage cf befdg
fdga df fdbage bdeaf adegb agedcb fbd fbcea dbfgec cfbegad | fd bfdgcea aecgbd agedb
cbegda agcbfd gbeafdc dacge fgbec bd edfacg bead gdb dgbce | bdg acgde daeb gbafcd
degcbf fg acgedb cdgbfea adgf bfg eagbd fgbae becfa geadfb | dgaf febag gdeba gdbcea
fcgaeb aebdg gcadeb cegfdb afdeg ebcfgda dgb cdba bd bcage | agedf eagcb dcgaeb cgfbae
gbecd abedc fegbcd efgacd fegdab bg bgcdafe geb fdceg gbcf | beg bdgeaf gdfec efgcda
acbfe cef egfabd afbeg bdfeacg cage gfdcbe ec beagfc facbd | cef fce afebdcg acefb
da afd gfebdc cadg deafc ecfdg gdafbec fagcde edbafg eabfc | aefbdg ad fcdge dgbfce
ecgdaf df dbeac fbegc cdf bfedc bacgde dfaecb dbfa bfgdace | defcag bcfed fedcb begcf
ceafbd ecfgbad afgdbe gbdac fcgbed ebg eg gfae egadb baedf | ge eg gafedb bge
bf abcdfe afb acdbf fdaebg cdgabe cbfe cdbea dgaefcb gfcad | cagedbf acefdbg febadgc edbac
aegcdf abgecf abefdgc daefg agbfd eacgd cgeadb decf fe gfe | fe gcafedb fe bdcgea
be cgfdea bfce fbegd bagefdc bge gbfda fedgc ecdabg bdefgc | begdf dgaefbc bge edbcfg
facebgd fcge caf agbecf dbafg aegcb dcabeg bagfc fc fcadeb | cf caf fc ecafbg
fbegca bfegad cafe caegb beagf gfcdab ca cdegb bfdegca cag | fgaeb agc ecfa bfaeg
ab fceadgb bae gceda afbced gfdeb edbafg fdbecg aebgd bfga | ba ebfdca gfedb dagbe
abdfcg egcf efd efgacdb fe ebacd dbcfeg gbfaed bdcfg cdbef | cadfgb gfce def geabfdc
beafdc bgdec efgbac gab gdbcaf ebcga gfea ga cagebfd fcabe | gba cbega gaef eabcfdg
de fdcegb fcebd afegbc efd bdcfa decgbaf gfecb gefdac dbge | gedb edf gbfec ed
dbfacge eabdg beg adbcg eg abdfe fbadeg adcbef bfecdg gaef | bedfga bdfae dafcbe beg
acfbg bfgdea eabcf fagbced gcfd cgabfd gf cadbge fbg acgbd | gfdc bcgfeda acbgf dcagb
cdeba gfdbce cad acbf gedba ecabdgf cfadeb egadfc ca cdbef | defbgc cdbef caedb acebd
edfabg ab cfegdb adgcebf eacb dcafb ebdcf efbdac adfgc abd | efbdag abec cfbad ab
abgdcf fdg fg edfabc dabfegc gfac bdcaf egdcb gcbfd fdaebg | dfbgeca fg bfdgac cgdefab
gcdafb ce cbde fecgd acgebdf gcbfd aefdg gbfdce gcfeab cef | bced dacgfb dbgecfa decb
gaecfdb cge dacef cg bgead gcead fbecag cbeafd dfgc fdaegc | dfcg gc agced dfcg
eb cegfab gdbe ecagfd dafcb decagbf fbe abefd gfead eafdgb | gbafed fgeda abfde be
gebadf ebfdg bgefdc bdea bafcedg feagc adbfcg befag ba bga | gacebfd geadfb acfgbd bade
becda dgaeb dg dcefab gbd cgfdba cabged gfabe egdc gafedbc | gabef dgb fcdabg dceg
aebf bfc bfecg bf febcgad egbafc fgedc eacgbd dafbcg ebacg | aefb gfbaec bf fb
cdgaeb bdacefg afecgd edaf cagfe ae gbcfa cae bgcfde fgecd | cfdage abdgfce defa cagfb
bedag gdcbea gfebdca adg dgfbac dg faegcb eafdb egdc acegb | dg efgadcb gdce bcegad
dacfgbe fg fbace fbgae gef cefbad debga gefbcd ebcgfa gfca | cafg decbfg efg gf
deca cbaefd fdgebca fca bfdga gbdcfe ac fcbeag dcfeb badcf | adcfb dfcbgae gecbaf eabcdf
gdabefc ca gbefda aec cafg ebafc gcdeba gefba ecbdf feabcg | cfga fgaebd dcbgae bfage
ecd fagedc dcaegb efgd febdcga fcade efcgab fgcea afdbc de | gbceda de gefbca efdg
afgcb ecbdgfa bdgecf ad dgea cadfg afd afecbd cedagf fecgd | beafdc afd geda adgfec
gbdcfe cdeab fd fecda adcgebf adebgc badf egcaf dfc fbdeac | cgafe fbdeca beagfcd caebfgd
abgfc gfedabc cb dcefga bdca fcb egbcfd dbgafc dafcg gabfe | acefgd bcf dbgfac acgdf
facge ecadgb fecdba dfcg fc fac eagfcd bgecfda afgeb eacdg | cdfg caefg fc bdeafcg
efdgb cebadg ceafbg fadcbg bcafdge fag af efgba cabeg faec | abgce ceaf gaf gbdeac
cafegb gcadfb begfacd bfad db dbgfc dbc cgfde bgcdea cbfga | gbfadc edabcg agfbec dbfa
facb acbefdg baedfg fbegca gfeadc bae dgcbe geacf ab gcaeb | ba eba baegdfc ba
edfgba gbfed dfc afdcebg cd cbdef fbaec gedcbf gcdb dgefca | gfbade cbdg dfc cbgd
bfc fgcbda dgfecb dbegac gcaf afcbd dcbfeag dagbc adbef fc | afcg afgc fcb gbedfc
dfagb edagbfc dceb agebcf bc gaedfc fbc cfbaed fcdba adecf | bc cb ecdfba bdec
gc cgdfbe agcb bfgdcea egacfb gefdba dcfea fcg fgeac aebgf | fbgdec cbafge gbac befga
fgdac agfbed ec bcegfd eadgbc gdcae gdaeb abec ceg bcdgfae | fgcad dagce adfcg ec
gfcad aecbgdf feacgd ceda fgdae fca ac bagedf acfegb cfbdg | acf afc fac ac
agcdeb ga aebdgf bedfgc age feadg cadfe egdbfca efgdb gfba | edbgac fdeca acedbg cgbdfe
gfdc acebf dce adgcbe egcdaf aedgfbc gfdae afgbed aedcf cd | gadcef dc fdecag aegfdb
agcdeb gdfc gebdc bcaegf cf dabef cfdegb ecbfd aecfdbg fbc | fbc debfc gecabfd gedcb
aebfc aecgf cb daegcf dbaecg cbdgafe adfeb bce bfgc bgcaef | ceb fbgc cbgf bdcefag
egfdab cefbgad cbfdea gdea egbdf bed de gbefc gbfad gbfacd | eacdgbf agbfdc dgcaefb de
adbfce be afcebgd fadbc bed dfaeb badcfg acdebg eafdg cbfe | edb cbfe fdgeacb deb
bad feab cfdga fcgdbe dcfbe gedacb cadbf becfagd dbeafc ab | efcbgda aebf bdcgafe abef
dbfeg dfbecag dace cfdbag gcabd cegdb egafcb ec ceb debcga | bce ce adbecgf abgcd
cabdgf dea abcgd abfdeg acdbge befdgac edacg fegdc abce ea | edacg gedcf dceag bagdfc
aceg ac bacgfd ceabf fcaegb edbgfc dbfea bgecdaf ebgfc acf | agce ac ecbfa agfbecd
bgcfda egc dbefcag befdg ec aedc cadfeg cgdaf gefcd fcageb | cebdagf dace ecfdag bcdaegf
bfdgc gbe acgef gbcade be egfcb afbe bacfeg dfbgeac gcfaed | abfe gfcbd fbgeac dcgbf
efcbd adegbcf agefdb cdabf bedgca abgdf dgabfc ac fcga bac | ac abc debgac fbdgac
ebgacd fd bdface cabdf dfa gfcab gbeafdc ecfd ecbda egfdba | fd df dcef fd
cbadge geb cfdge bg bfeca bgecaf edcabf gcbfe afgb efgdabc | fabg gabf gfbec afebc
fedgabc cbedg fcaegd ecg dcfgb ce bdfgae gaebd cbea ecdbag | efcbagd bfeagcd ce ce
gebdcf cdbfga fead aedfcgb ad afegbd gdfbe bedag ecagb agd | ad bfcagd ceafdgb fcgadbe
fdegb egdbfc cegfdab dacfge dge gadbf edcfb edcabf egcb ge | edg bagdcfe egcb ged
ad cfbeag adb efbad feabc acfd cgdeba fdacgbe dgefb feadbc | fabce cfbega defbca bad
facbge cfabd bgcafde agdb fab becdf cfaegd dafgc agfcbd ba | cbfda ab cbfdaeg dfagbc
eg dgaebc dfcgab cdgfe edfcb adcgf cedagf geaf bacdefg gde | dfcbe eagf agdecf ge`
