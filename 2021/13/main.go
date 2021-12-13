package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const foldPrefix = "fold along "

type coordinate struct {
	x, y int64
}

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var points []coordinate
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		split := strings.Split(line, ",")
		if len(split) != 2 {
			continue
		}
		x, _ := strconv.ParseInt(split[0], 10, 64)
		y, _ := strconv.ParseInt(split[1], 10, 64)
		points = append(points, coordinate{x, y})
	}
	fmt.Println(len(points))
	scanner = bufio.NewScanner(strings.NewReader(foldsInput))
	var m map[coordinate]bool
	var maxX, maxY int64
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, foldPrefix) {
			continue
		}
		line = line[len(foldPrefix):]
		n, _ := strconv.ParseInt(line[2:], 10, 64)
		var foldX bool
		if line[0] == 'x' {
			foldX = true
			maxX = n
		} else {
			maxY = n
		}
		m = make(map[coordinate]bool)
		for i, c := range points {
			if foldX {
				if c.x == n {
					c.x = -1
				}
				if c.x > n {
					c.x = n*2 - c.x
				}
			} else {
				if c.y == n {
					c.y = -1
				}
				if c.y > n {
					c.y = n*2 - c.y
				}
			}
			if c.x >= 0 && c.y >= 0 {
				m[c] = true
			} else {
				log.Printf("%q: %v", line, points[i])
			}
			points[i] = c
		}
		fmt.Println(len(m))
	}

	for y := int64(0); y < maxY; y++ {
		for x := int64(0); x < maxX; x++ {
			if m[coordinate{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

const input = `
726,774
246,695
579,249
691,724
246,820
738,887
1088,75
264,887
704,775
907,625
676,117
507,658
1009,24
547,735
157,126
599,113
445,226
363,691
918,794
927,113
999,400
443,305
654,729
408,767
1066,863
1148,473
321,35
1093,803
1044,718
202,889
262,164
378,541
619,662
1034,849
432,595
1145,656
1295,668
1125,705
1161,529
759,619
1170,147
688,742
328,729
718,439
935,701
246,647
594,110
845,495
160,189
1225,315
580,486
469,481
440,401
584,774
897,719
1007,516
547,159
112,117
982,645
62,439
192,441
631,211
654,16
113,698
378,865
373,19
441,777
390,7
1263,312
1121,610
509,582
893,352
44,131
1092,19
592,719
918,100
326,820
62,719
520,889
718,103
571,579
1165,267
1208,334
525,415
268,588
769,464
596,716
734,436
1283,747
35,457
982,94
1235,849
932,752
1265,243
262,752
99,704
547,732
1096,829
791,329
1222,320
199,568
671,494
1079,93
569,315
129,415
868,878
788,803
1175,889
965,676
904,660
552,560
619,484
507,236
566,768
1215,30
45,410
885,441
478,336
945,397
596,306
1145,686
189,495
1153,126
79,444
719,403
903,91
27,105
441,329
768,565
671,400
507,205
224,390
867,753
425,441
1123,67
833,14
793,278
1237,724
238,808
1099,32
411,243
1,574
726,120
862,441
212,94
80,199
1047,312
164,389
440,773
185,189
412,114
537,792
403,697
1208,768
408,527
1135,831
746,826
1064,110
15,226
102,574
1001,442
769,16
1236,408
440,121
1309,131
771,473
1064,527
189,284
427,704
276,865
986,593
1009,870
745,276
965,291
1190,5
1084,768
313,750
976,544
1222,658
223,619
718,791
976,574
705,822
785,863
388,390
1235,716
579,645
90,378
497,849
1308,145
658,301
648,94
1150,189
231,129
408,639
390,119
132,556
639,220
885,453
413,651
639,68
177,544
552,782
239,130
867,645
187,67
1163,494
1227,19
249,658
132,270
740,82
73,206
1166,525
262,506
1092,875
923,509
967,400
813,849
734,364
604,541
566,758
90,852
190,831
1047,65
313,144
542,157
720,304
129,863
1280,149
448,441
947,484
1048,506
408,863
657,803
480,834
201,369
489,208
1275,36
1064,784
781,845
1071,152
1251,61
455,500
564,68
758,798
922,390
443,421
505,792
60,784
1059,645
741,315
89,164
142,637
348,705
763,735
165,686
711,781
1220,852
455,120
1205,859
1208,413
1098,240
33,505
821,208
35,289
1251,833
1131,226
745,730
75,625
905,75
738,7
1223,316
923,395
7,884
1153,299
552,96
1047,134
642,266
537,698
1211,190
959,235
1235,625
711,614
510,798
73,306
47,582
657,875
468,75
638,346
144,525
612,665
917,57
1235,402
211,249
831,63
1197,250
493,686
887,801
85,315
263,65
924,306
140,595
997,141
657,173
579,850
489,686
448,5
572,7
1248,103
445,674
711,113
1232,884
1121,278
845,732
786,630
114,712
691,662
140,147
1088,299
408,191
897,691
92,662
904,754
1309,621
102,768
7,10
33,429
903,803
2,749
1083,704
157,819
325,91
830,386
763,284
175,63
902,199
887,129
1285,297
1287,683
590,304
714,716
6,686
1136,834
452,889
653,315
135,145
683,329
251,473
1110,455
959,516
1091,838
407,220
654,800
549,239
765,355
113,250
771,93
194,270
864,803
517,417
345,291
253,724
365,145
522,624
692,495
830,834
1021,301
825,595
145,890
1125,481
48,119
82,7
965,666
540,630
542,121
731,290
85,763
957,297
1277,617
1089,239
619,612
62,175
427,190
401,724
1133,798
475,882
1062,140
246,784
1001,515
244,31
759,171
1246,282
249,236
919,816
907,25
755,403
557,725
15,332
840,861
1031,760
965,218
813,45
440,829
885,457
45,243
351,435
191,565
984,820
715,275
689,641
289,845
575,892
605,822
1136,508
137,278
870,513
59,621
79,539
89,276
1,621
753,725
1119,301
1,763
1159,539
1059,25
267,843
1072,236
1203,607
425,453
62,551
788,624
576,884
326,430
345,452
539,473
1195,337
981,432
23,683
455,394
1098,654
1118,441
1277,429
567,625
773,698
23,459
441,464
870,488
567,269
517,399
935,302
436,875
309,442
900,191
1277,277
1310,165
383,781
149,792
492,578
88,658
736,500
103,275
731,44
720,794
1278,600
981,462
1020,623
329,462
333,483
977,880
1181,863
509,890
846,831
246,527
570,754
841,481
78,884
1231,355
618,47
413,719
850,786
174,386
364,175
1193,196
162,421
221,239
373,875
718,7
1231,539
1218,662
333,868
885,9
30,149
1299,511
227,190
1197,698
1150,705
657,238
783,717
1159,383
212,688
241,719
618,495
1232,458
985,432
27,75
918,506
328,800
145,582
387,395
90,516
1019,773
334,574
1285,149
716,558
413,236
353,477
607,193
845,844
874,875
1099,648
264,63
99,491
263,312
303,67
463,565
433,239
902,527
132,355
1300,550
606,119
291,849
25,86
301,865
1046,831
805,102
408,598
1101,693
1215,864
947,691
64,730
1237,170
970,628
1136,60
947,730
334,544
1149,686
1009,198
691,282
691,457
378,142
599,614
648,320
507,400
1178,803
478,558
267,51
279,701
965,452
691,232
401,170
219,838
1292,413
296,373
246,127
446,91
894,86
115,480
1287,155
433,655
263,134
87,630
965,403
107,623
189,610
330,749
1121,271
965,603
135,749
1059,473
328,165
27,147
443,134
1210,320
1211,470
415,824
835,882
405,819
957,870
493,721
1246,164
935,591
895,824
264,455
99,470
408,296
803,400
1084,126
1135,63
835,46
830,60
653,721
1104,453
525,863
102,334
1235,45
870,121
375,591
60,336
348,880
895,600
517,477
427,526
100,551
37,301
477,880
0,94
383,390
266,718
212,240
634,329
291,493
976,320
751,30
1119,593
443,753
1099,645
79,450
189,278
358,745
870,355
1064,820
353,597
770,630
157,75
22,371
214,493
465,726
1205,655
740,469
125,49
1019,849
735,556
1148,421
832,336
803,338
848,441
946,175
301,149
115,305
528,215
6,208
6,320
691,410
238,658
16,346
735,892
689,725
661,320
1136,386
1153,810
363,730
326,464
321,819
885,885
1048,388
425,9
127,809
656,16
599,280
74,408
387,509
73,724
293,877
557,687
656,878
1304,320
1034,865
704,119
126,663
141,656
914,859
1230,647
340,266
33,501
1262,352
505,344
1283,105
1198,289
855,500
574,871
540,598
465,844
1210,343
825,96
290,623
174,508
132,539
867,93
1101,201
392,142
472,126
552,320
631,739
867,473
763,758
467,67
251,269
1273,301
619,282
1148,130
801,4
566,136
1098,94
1274,47
333,411
401,82
48,352
1230,522
907,269
870,829
631,683
406,660
440,488
800,798
244,765
657,768
662,94
33,465
785,415
83,450
1159,355
653,686
392,730
417,352
691,829
902,598
947,282
242,469
135,301
174,834
816,189
914,894
671,562
460,786
1079,765
125,525
213,430
480,60
1086,390
1111,809
375,302
1227,390
1072,684
343,494
443,473
904,234
485,45
1285,86
977,299
758,782
242,425
460,718
37,593
730,486
559,877
505,102
401,812
231,254
403,269
493,238
1181,479
189,29
545,355
264,439
547,60
900,695
689,393
1079,254
408,199
846,383
1223,630
1121,284
345,666
441,117
237,800
191,525
1066,255
552,768
639,494
798,705
1004,189
1169,861
686,745
1121,29
554,126
1277,465
977,432
711,399
12,628
`

const foldsInput = `
fold along x=655
fold along y=447
fold along x=327
fold along y=223
fold along x=163
fold along y=111
fold along x=81
fold along y=55
fold along x=40
fold along y=27
fold along y=13
fold along y=6
`
