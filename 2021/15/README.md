# 2021 day 15

https://adventofcode.com/2021/day/15

## Backtracking

For my input,
there are total of 4 backtracking required for part 2.
Although one of them are in the range of the map of part 1,
that was not on the path to `{99,99}`
(no backtracking required for part 1).

Here are the 4 partial maps around the backtrackings
(the coordinates are the top-left corners):

```
{35,22}:
6211915
1245886
5181898
2843922
8519791
2619983
6352291

{114,65}:
4936629
7251566
5422111
2141884
1218288
1699812
2332879

{179,145}:
2263336
5343514
2931212
4133933
3329492
2454843
7527436

{352,340}:
8266754
7642563
7762812
5315668
7673336
4663371
5663968
```

The (partial) routes around them:

```
{35,24}->{35,25}->{36,25}->{37,25}->{38,25}->{38,24}->{39,24}->{40,24}->{41,24}

{117,65}->{117,66}->{117,67}->{117,68}->{116,68}->{116,69}->{116,70}->{116,71}->{117,71}

{182,145}->{182,146}->{182,147}->{182,148}->{181,148}->{181,149}->{181,150}->{181,151}->{182,151}

{355,340}->{355,341}->{355,342}->{355,343}->{354,343}->{354,344}->{354,345}->{354,346}
```

Visualize the (partial) routes and the big numbers they are walking around:

```
1>1
  v
  5
  v
8 1
  v
4<3


      2>1>1>1
      ^
2>1>4>1 8 8


      1>2>1>1
      ^     v
4>1>3>3 9   3


      2>8>1>2
      ^
5>3>1>5 6 6 8
```
