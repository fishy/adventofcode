package main

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const nDimensions = 3

const minOverlap = 12

const nScanners = 33

var reScannerLine = regexp.MustCompile(`scanner ([0-9]+)`)

type point [nDimensions]int

func (p point) String() string {
	s := make([]string, nDimensions)
	for i := range p {
		s[i] = fmt.Sprintf("%d", p[i])
	}
	return strings.Join(s, ",")
}

func (p point) transform(transforms [nDimensions]dimensionTransform) point {
	var r point
	for i, t := range transforms {
		r[i] = p[t.dimension]
		if t.flip {
			r[i] = -r[i]
		}
		r[i] += t.diff
	}
	return r
}

func (p point) reverse(transforms [nDimensions]dimensionTransform) point {
	var r point
	for i, t := range transforms {
		j := t.dimension
		r[i] = p[j]
		if t.flip {
			r[i] = -r[i]
		}
		r[i] -= t.diff
	}
	return r
}

type dimensionTransform struct {
	dimension int
	diff      int
	flip      bool
}

type scanner struct {
	beacons []point

	distances [][]point

	p *point
}

func (s *scanner) addBeacon(p point) {
	s.beacons = append(s.beacons, p)
}

func (s *scanner) buildDistances() {
	s.distances = make([][]point, len(s.beacons))
	for i := range s.beacons {
		s.distances[i] = make([]point, len(s.beacons))
		for j := range s.beacons {
			for k := 0; k < nDimensions; k++ {
				s.distances[i][j][k] = s.beacons[i][k] - s.beacons[j][k]
			}
		}
	}
}

func overlapBeacons(a, b []point) map[int]int {
	m := make(map[int]int)
	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				m[i] = j
			}
		}
	}
	return m
}

func flipDistances(d []point, transforms [nDimensions]dimensionTransform) []point {
	r := make([]point, len(d))
	for i := range d {
		r[i] = d[i].transform(transforms)
	}
	return r
}

func (s scanner) overlap(o scanner) []dimensionTransform {
	for i := range s.beacons {
		a := s.distances[i]
		for j := range o.beacons {
			for d1 := 0; d1 < nDimensions; d1++ {
				for d2 := 0; d2 < nDimensions; d2++ {
					if d2 == d1 {
						continue
					}
					d3 := 0 + 1 + 2 - d1 - d2
					for _, f1 := range []bool{false, true} {
						for _, f2 := range []bool{false, true} {
							for _, f3 := range []bool{false, true} {
								transforms := [nDimensions]dimensionTransform{
									{
										dimension: d1,
										flip:      f1,
									},
									{
										dimension: d2,
										flip:      f2,
									},
									{
										dimension: d3,
										flip:      f3,
									},
								}
								b := flipDistances(o.distances[j], transforms)
								m := overlapBeacons(a, b)
								if len(m) >= minOverlap {
									var p1, p2 point
									for i, j := range m {
										p1 = s.beacons[i]
										p2 = o.beacons[j]
										break
									}
									p2 = p2.reverse(transforms)
									for k := 0; k < nDimensions; k++ {
										transforms[k].diff = p2[k] - p1[k]
									}
									return transforms[:]
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func (s *scanner) transform(transforms []dimensionTransform) {
	var t [nDimensions]dimensionTransform
	for i := range transforms {
		t[i] = transforms[i]
	}
	for i := range s.beacons {
		s.beacons[i] = s.beacons[i].reverse(t)
	}
	s.buildDistances()
}

func main() {
	scn := bufio.NewScanner(strings.NewReader(input))
	scanners := make([]scanner, nScanners)
	var totalReadings int
	for scn.Scan() {
		line := strings.TrimSpace(scn.Text())
		matches := reScannerLine.FindStringSubmatch(line)
		if len(matches) < 2 {
			continue
		}
		i64, _ := strconv.ParseInt(matches[1], 10, 64)
		i := int(i64)
		for scn.Scan() {
			line := strings.TrimSpace(scn.Text())
			split := strings.Split(line, ",")
			if len(split) != nDimensions {
				break
			}
			x, _ := strconv.ParseInt(split[0], 10, 64)
			y, _ := strconv.ParseInt(split[1], 10, 64)
			z, _ := strconv.ParseInt(split[2], 10, 64)
			scanners[i].addBeacon(point{int(x), int(y), int(z)})
			totalReadings++
		}
		scanners[i].buildDistances()
	}
	convertedScanners := make(map[int]bool, nScanners)
	convertedScanners[0] = true
	scanners[0].p = &point{0, 0, 0}
	beacons := make([]point, 0, totalReadings)
	beacons = append(beacons, scanners[0].beacons...)
	for len(convertedScanners) < nScanners {
		for i := range convertedScanners {
			for j := range scanners {
				if scanners[j].p != nil {
					continue
				}
				t := scanners[i].overlap(scanners[j])
				if len(t) != nDimensions {
					continue
				}
				scanners[j].transform(t)
				beacons = append(beacons, scanners[j].beacons...)
				convertedScanners[j] = true
				var p point
				for k := 0; k < nDimensions; k++ {
					p[k] = -t[k].diff
				}
				scanners[j].p = &p
				fmt.Printf("Found %2d against %2d: %17v, %+v\n", j, i, p, t)
			}
		}
	}
	sort.Slice(beacons, func(i, j int) bool {
		for k := 0; k < nDimensions; k++ {
			if beacons[i][k] == beacons[j][k] {
				continue
			}
			return beacons[i][k] < beacons[j][k]
		}
		return false
	})
	last := beacons[0]
	var uniq []point
	uniq = append(uniq, last)
	for i := 1; i < len(beacons); i++ {
		if beacons[i] == last {
			continue
		}
		last = beacons[i]
		uniq = append(uniq, last)
	}
	fmt.Println(len(uniq))

	var maxManhattanDistance int
	for i := range scanners {
		for j := i + 1; j < len(scanners); j++ {
			var md int
			for k := 0; k < nDimensions; k++ {
				md += int(math.Abs(float64(scanners[i].p[k] - scanners[j].p[k])))
			}
			if md > maxManhattanDistance {
				maxManhattanDistance = md
				fmt.Println(md, *scanners[i].p, *scanners[j].p)
			}
		}
	}
}

const input = `
--- scanner 0 ---
507,-531,548
401,525,521
764,817,-797
-492,-808,548
-640,576,-732
7,-18,-66
-786,-540,-464
-815,-426,-432
556,-417,485
265,-614,-484
-738,-570,-474
-788,754,757
574,-611,471
-143,70,34
270,633,578
697,844,-775
259,-683,-478
-643,602,-605
-486,-694,513
250,-771,-455
450,684,573
-459,-707,400
-720,775,668
-736,605,-807
822,863,-756
-654,684,780

--- scanner 1 ---
592,569,-925
-372,-699,-541
902,-823,-595
-595,-894,631
809,-461,714
896,-880,-730
678,-453,815
736,416,410
-805,-929,633
661,404,-924
843,-510,760
-786,578,-503
-806,428,-534
890,425,388
928,-805,-744
-374,-790,-658
-780,367,-438
-354,478,846
-395,422,767
-469,-669,-599
101,-45,-104
-351,560,805
561,463,-908
-731,-894,726
799,414,549

--- scanner 2 ---
406,-343,763
682,570,-319
-508,503,634
-521,364,675
813,305,667
-540,724,-553
874,367,718
-446,-676,604
767,-732,-466
506,-452,796
-679,682,-436
767,489,-312
817,657,-349
-357,-766,709
-544,671,-503
678,329,737
-418,-591,-716
-445,-836,690
758,-764,-472
497,-472,755
28,-48,92
795,-906,-520
-465,422,698
-326,-681,-711
-367,-555,-749

--- scanner 3 ---
-475,-787,-901
743,-663,-469
-723,240,744
-616,-396,497
776,-726,807
652,-780,-467
-510,-413,466
781,547,-879
-625,-818,-881
468,735,416
-368,693,-437
480,681,377
29,10,-92
833,-591,801
-451,-397,566
-414,-883,-875
637,683,-888
-612,359,752
622,-722,-564
-503,260,704
758,-577,833
-475,624,-517
823,643,-884
-533,660,-364
392,595,492

--- scanner 4 ---
764,-534,380
819,496,480
743,-539,-587
-637,643,-343
-793,-406,381
-701,544,-382
-625,605,-342
775,-699,374
693,-443,-734
-116,15,-74
-822,420,328
776,-452,-622
734,-580,316
312,409,-411
-821,-452,382
352,271,-303
625,530,495
-801,632,329
-688,561,328
55,-182,-47
-742,-948,-602
20,-47,106
-789,-495,516
-577,-953,-487
387,370,-408
-573,-982,-634
697,565,509

--- scanner 5 ---
-724,330,944
327,-646,-389
-578,-849,664
-546,260,898
-137,-161,177
773,588,464
-804,402,-391
-647,-762,-384
501,-915,628
-825,470,-480
-716,-791,-262
758,698,-600
3,-19,144
698,640,385
721,671,519
657,731,-748
364,-641,-466
595,-911,771
-734,536,-368
-519,-803,554
-555,-800,451
-565,369,843
644,704,-635
606,-873,545
-682,-776,-456
218,-736,-446

--- scanner 6 ---
342,-659,-336
-813,731,-628
354,343,-418
299,443,612
-92,43,-66
375,-624,-423
731,-445,486
806,-448,525
-627,385,750
343,337,-583
-414,434,773
766,-340,387
-699,711,-778
493,458,646
-939,-758,755
-517,431,809
-834,-341,-303
-879,-763,809
421,472,681
384,279,-464
-864,-343,-286
-803,722,-863
-846,-734,680
-784,-342,-299
425,-771,-327
-9,-33,91

--- scanner 7 ---
678,469,-586
535,-777,-610
-474,-647,-677
569,-792,-804
-523,374,514
505,546,687
508,453,889
504,-944,703
-365,-737,-620
-628,-649,566
-24,-2,-38
-905,697,-466
500,-763,-682
-76,-143,108
-804,730,-318
-709,-615,522
-637,403,546
-511,-642,-600
433,-943,816
587,535,-698
-670,-633,656
-667,451,556
657,401,-675
479,549,814
645,-944,837
-791,703,-431

--- scanner 8 ---
37,41,-38
-579,445,-499
-885,799,441
-628,559,-562
-313,-835,742
-876,711,464
-371,-789,783
471,-661,679
525,538,-775
656,424,525
428,-547,588
425,618,-841
828,-722,-365
606,344,388
868,-733,-353
-870,642,470
-438,-816,761
875,-563,-389
-734,-497,-535
544,469,456
549,-561,609
-865,-605,-592
-561,639,-485
-685,-656,-587
416,514,-681

--- scanner 9 ---
-608,733,767
-763,-907,-585
421,-600,-318
-57,-121,48
805,449,631
585,408,-861
-796,309,-747
-694,573,766
430,-723,-471
512,-788,699
614,-825,842
-749,-873,-584
-733,-809,-676
-783,301,-757
-623,-467,691
550,-704,837
844,513,730
-664,-606,704
700,573,-866
-600,-525,747
-628,325,-803
553,577,-895
-589,663,855
845,525,692
507,-757,-372

--- scanner 10 ---
370,735,428
230,-341,-879
-402,812,-518
309,712,-606
-568,-614,-620
-553,-808,-654
318,667,306
335,-376,-780
-401,833,-730
-730,518,624
263,-595,845
-582,-686,-585
338,-301,-741
-698,-681,389
-712,652,585
-703,-604,426
269,636,-545
277,-698,667
453,675,354
337,-611,752
-401,699,-589
-696,562,543
-144,92,13
-608,-624,280
270,648,-507
18,-45,-102

--- scanner 11 ---
811,-924,-797
535,569,-785
-730,-780,751
373,-818,221
519,527,-695
358,-739,346
-574,-690,-908
-384,-755,-922
637,-845,-797
769,272,402
34,-24,-179
-565,530,-795
647,-907,-697
-488,674,-837
-116,19,-25
792,256,326
-400,567,601
-915,-718,760
501,488,-829
-520,550,577
409,-862,282
-380,541,580
-559,-863,-932
-517,560,-805
709,256,249
-946,-799,749

--- scanner 12 ---
-582,613,328
-653,500,329
537,-572,-757
567,-642,717
13,132,47
927,559,-602
-626,676,-342
932,724,-487
-510,-636,496
-633,-637,-675
385,508,381
474,512,516
961,728,-634
686,-743,686
-441,479,328
438,-606,-680
384,-449,-766
-664,664,-343
630,-687,763
-776,-674,-573
-771,-553,-637
-430,-641,517
-565,600,-382
-424,-676,488
153,51,-85
431,480,382

--- scanner 13 ---
-298,-733,-669
429,-705,-512
-599,681,766
773,661,811
839,-349,740
411,-595,-559
819,-327,569
-341,-465,668
-589,725,-237
36,41,93
-468,-491,671
776,670,951
753,504,-407
640,516,-469
-641,537,843
-317,-752,-794
-595,794,-231
422,-734,-707
-618,882,-222
164,123,-12
797,644,931
801,-412,643
-266,-494,764
-276,-706,-736
754,507,-541
-609,651,814

--- scanner 14 ---
558,-399,-314
-762,-706,948
674,739,432
-383,821,-498
-804,647,949
274,769,-636
-607,-420,-704
646,-767,551
-399,762,-398
627,-549,-306
-900,-814,908
638,797,585
-869,858,943
731,-707,638
254,549,-676
32,169,181
-382,656,-429
-478,-391,-737
707,904,505
-36,8,9
441,-509,-252
-801,-709,873
824,-781,654
-773,797,969
373,660,-684
-532,-396,-797

--- scanner 15 ---
413,-407,801
-737,-743,619
629,461,254
-838,-712,500
668,771,-589
-667,731,649
47,-88,-75
-791,447,-659
-801,320,-571
-91,55,-105
633,608,225
-342,-612,-755
-806,508,-673
701,619,-551
640,683,-505
-562,-541,-752
-544,740,802
417,-542,712
-782,-789,556
-701,654,748
-507,-630,-864
386,-509,-789
385,-408,755
387,-362,-775
321,-466,-731
464,576,259

--- scanner 16 ---
466,731,704
467,501,-625
-820,697,-446
-666,-666,692
-726,714,-309
-498,-630,-617
-526,585,339
449,416,-699
495,-572,-457
-594,771,363
-38,-34,70
429,438,-528
670,-850,832
-619,-538,-586
471,651,725
37,85,-37
746,-866,855
-851,643,-346
431,737,733
-604,657,318
471,-555,-337
-678,-683,-628
-691,-565,550
492,-581,-317
823,-866,772
-779,-593,688

--- scanner 17 ---
-607,890,-709
419,798,602
-713,-494,-404
643,-623,-695
-679,-440,-531
708,-273,375
498,459,-551
-254,560,609
402,767,496
-538,882,-559
99,137,-35
552,399,-556
627,492,-580
810,-247,254
-234,520,389
-421,944,-669
31,7,-165
-623,-461,-567
-589,-646,403
-491,-607,386
651,-650,-598
-353,512,500
403,698,661
788,-251,227
646,-543,-746
-651,-620,349

--- scanner 18 ---
428,512,-741
-872,-620,284
-686,790,677
-166,143,-113
-17,85,34
400,596,-829
681,782,304
747,738,408
435,-218,-666
-881,-661,-619
-685,738,643
-902,-666,348
682,796,451
367,558,-677
651,-705,484
-831,-757,-766
-952,711,-507
604,-252,-634
-914,-748,-654
439,-247,-519
-913,-501,342
-706,659,804
655,-794,349
683,-748,350
-915,753,-354
-970,787,-552

--- scanner 19 ---
-600,-739,557
-759,-424,-712
-790,551,313
689,799,-511
-594,236,-900
-6,-158,-188
-540,410,-872
-546,378,-863
390,-866,-858
530,466,341
722,690,-611
-630,-836,615
616,306,340
681,670,-602
-721,-834,498
-823,587,262
647,-422,653
436,-428,718
635,376,252
555,-939,-825
497,-786,-789
-615,-415,-848
99,21,-73
-785,461,221
-801,-419,-808
497,-402,622

--- scanner 20 ---
-701,607,670
-14,69,-61
653,936,355
-622,-294,-760
670,776,-800
-659,-459,636
441,935,376
-933,815,-408
-933,898,-485
512,936,489
434,-582,-764
371,-485,-802
584,768,-760
-917,875,-561
-639,-414,-736
349,-656,-782
487,-557,378
-620,641,626
513,-554,505
516,-461,303
664,747,-765
-629,-411,646
-515,-382,678
-648,-274,-739
-662,625,536

--- scanner 21 ---
325,713,596
683,-524,445
383,690,540
328,670,-799
290,651,-603
33,56,155
-511,633,-422
562,-460,450
464,754,583
-636,-706,-256
-568,-654,-378
646,-422,588
-382,625,-331
-491,592,-431
719,-557,-501
-133,106,44
-591,-715,656
647,-605,-395
-805,679,432
-628,-571,-304
721,-637,-587
-667,-755,595
266,721,-738
-655,-754,557
-715,784,521
-875,698,520

--- scanner 22 ---
840,-602,325
827,588,-776
378,291,762
-513,317,667
419,416,730
-36,-21,-55
712,-535,-730
-500,-613,593
716,624,-760
854,645,-843
24,-152,108
668,-616,-798
891,-654,415
773,-587,492
-613,-669,557
602,-580,-845
-609,475,629
-267,-619,-608
-418,-748,570
-519,266,-848
-444,435,629
-412,-593,-541
-411,-695,-623
427,367,835
-401,243,-725
-373,271,-871

--- scanner 23 ---
424,-915,857
-819,-486,664
489,-921,646
-280,750,-812
910,526,820
-417,-489,-531
-293,670,-813
393,-880,-815
911,457,655
-60,-133,158
-781,-530,802
815,765,-392
492,-845,-710
-321,-539,-384
520,-941,796
-860,-489,746
118,-29,12
761,825,-397
-330,-456,-496
894,426,761
-297,720,-814
686,732,-421
-402,755,799
531,-920,-707
-563,664,810
-424,601,808

--- scanner 24 ---
-642,-412,665
623,843,-461
708,-808,-518
35,-20,-85
107,111,82
-529,-450,-367
-572,594,455
646,-843,681
-578,675,462
-610,-486,473
853,578,563
716,-867,816
-427,-487,-510
-562,-414,579
-470,644,-522
523,908,-377
-535,801,-492
665,-805,877
613,821,-362
936,572,566
-617,582,447
-465,-571,-372
749,-669,-505
685,-750,-462
987,690,563
-523,704,-466

--- scanner 25 ---
-456,718,-439
-465,650,819
-532,-510,-468
783,207,-481
-614,-729,471
-520,-607,-442
852,356,675
626,-468,-437
-575,-636,486
107,-169,72
939,-739,590
780,419,-504
990,-621,587
-341,770,-376
-614,617,866
167,2,-48
-425,678,-421
616,-490,-338
-569,-717,479
-467,-667,-498
856,-733,587
-417,537,844
987,237,635
729,358,-401
893,211,742
526,-403,-353

--- scanner 26 ---
-730,-771,858
606,-881,887
-621,808,-322
-785,-663,799
635,-763,867
647,-789,806
617,-465,-243
-736,614,800
455,453,722
-739,-542,840
-420,-668,-435
-854,547,877
-803,599,967
537,588,699
-413,-587,-350
654,374,-620
636,-427,-439
-573,666,-359
-450,-755,-300
695,354,-611
650,532,735
566,-403,-304
18,-155,-1
662,245,-711
-567,634,-350

--- scanner 27 ---
87,-20,-104
-774,-567,-781
802,443,-335
964,-328,687
582,483,-355
-389,771,-830
937,-324,709
-358,-305,832
-24,42,39
-302,-415,737
589,721,318
-730,-623,-637
-242,-440,786
695,690,301
-313,695,-690
805,-684,-592
530,621,302
-679,354,482
-789,-711,-688
-401,810,-776
-697,380,631
-754,386,496
816,-350,607
788,-723,-690
842,-697,-752
691,435,-454

--- scanner 28 ---
273,-544,-591
784,931,470
-436,-654,-453
701,755,-837
-773,-698,671
-479,683,-721
-637,794,632
291,-497,-481
706,905,603
284,-692,-452
-606,686,-695
798,-301,369
667,-387,288
-414,-721,-480
637,-356,288
-404,779,-692
681,584,-859
60,171,42
749,840,476
-799,-645,810
-868,-767,779
663,658,-774
-455,-779,-434
-656,743,536
-116,89,-15
-703,645,660

--- scanner 29 ---
-484,437,833
-507,-728,602
-424,354,759
102,-127,108
850,398,-551
-442,445,-524
-641,-746,563
785,551,-514
470,-838,-787
166,2,-37
865,371,550
863,451,636
-492,302,881
409,-874,-743
863,422,-525
-509,-600,-832
791,403,742
-421,531,-520
437,-538,331
534,-889,-664
-609,-610,-770
514,-455,331
590,-478,347
-609,-436,-849
-446,504,-746
-597,-712,644

--- scanner 30 ---
-672,-857,273
-559,457,788
400,-453,422
-115,-128,-157
-521,-644,-529
501,410,441
-526,292,715
-512,-638,-561
-773,624,-841
-80,19,-22
-782,583,-821
573,579,-663
638,-807,-581
295,-461,373
805,-874,-596
-484,-692,-387
715,-925,-648
-576,-899,268
-595,399,635
368,-412,297
463,479,323
501,575,481
560,384,-742
-536,-713,264
-873,591,-950
537,466,-607

--- scanner 31 ---
-524,280,430
713,709,-629
-554,734,-896
-506,-521,611
-705,-600,-731
698,-638,729
-623,-735,-708
631,-600,-850
732,637,638
781,-578,655
833,-646,777
71,-69,-83
-543,318,234
-465,-501,602
737,696,-655
-594,256,295
-510,-679,-744
604,-575,-624
684,709,532
-498,-392,702
690,-596,-628
584,631,-618
-513,642,-815
-525,616,-876
764,711,446

--- scanner 32 ---
721,678,399
-782,-722,762
-749,408,-490
722,-815,709
-889,-791,842
-157,-126,119
-843,785,404
790,482,-538
796,-908,693
782,-730,796
829,816,439
-598,-535,-694
-590,-444,-823
-557,-544,-714
-663,407,-330
689,-491,-637
724,717,474
-644,319,-400
11,1,-39
741,-369,-569
-782,689,378
-724,-758,915
789,444,-550
750,543,-614
606,-417,-577
-632,720,398
`