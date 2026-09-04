//go:build !arm64 && (!amd64 || !amd64.v2)

package p2

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn49 github.com/goccy/llamawasm2go/p1.Fn49
func Fn49(m *base.Module, l0 int64)

//go:linkname Fn65 github.com/goccy/llamawasm2go/p1.Fn65
func Fn65(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn163 github.com/goccy/llamawasm2go/p1.Fn163
func Fn163(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn357 github.com/goccy/llamawasm2go/p1.Fn357
func Fn357(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn359 github.com/goccy/llamawasm2go/p1.Fn359
func Fn359(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn362 github.com/goccy/llamawasm2go/p1.Fn362
func Fn362(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn363 github.com/goccy/llamawasm2go/p1.Fn363
func Fn363(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn367 github.com/goccy/llamawasm2go/p1.Fn367
func Fn367(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn375 github.com/goccy/llamawasm2go/p1.Fn375
func Fn375(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn395 github.com/goccy/llamawasm2go/p1.Fn395
func Fn395(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn396 github.com/goccy/llamawasm2go/p0.Fn396
func Fn396(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn401 github.com/goccy/llamawasm2go/p1.Fn401
func Fn401(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn404 github.com/goccy/llamawasm2go/p1.Fn404
func Fn404(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn405 github.com/goccy/llamawasm2go/p1.Fn405
func Fn405(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn448 github.com/goccy/llamawasm2go/p1.Fn448
func Fn448(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn538 github.com/goccy/llamawasm2go/p1.Fn538
func Fn538(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn540 github.com/goccy/llamawasm2go/p1.Fn540
func Fn540(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn546 github.com/goccy/llamawasm2go/p1.Fn546
func Fn546(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn556 github.com/goccy/llamawasm2go/p1.Fn556
func Fn556(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn560 github.com/goccy/llamawasm2go/p1.Fn560
func Fn560(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn611 github.com/goccy/llamawasm2go/p0.Fn611
func Fn611(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn652 github.com/goccy/llamawasm2go/p1.Fn652
func Fn652(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn656 github.com/goccy/llamawasm2go/p0.Fn656
func Fn656(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn666 github.com/goccy/llamawasm2go/p1.Fn666
func Fn666(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn690 github.com/goccy/llamawasm2go/p1.Fn690
func Fn690(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn708 github.com/goccy/llamawasm2go/p1.Fn708
func Fn708(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn739 github.com/goccy/llamawasm2go/p1.Fn739
func Fn739(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn740 github.com/goccy/llamawasm2go/p1.Fn740
func Fn740(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn754 github.com/goccy/llamawasm2go/p1.Fn754
func Fn754(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn756 github.com/goccy/llamawasm2go/p1.Fn756
func Fn756(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn760 github.com/goccy/llamawasm2go/p1.Fn760
func Fn760(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn761 github.com/goccy/llamawasm2go/p1.Fn761
func Fn761(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn778 github.com/goccy/llamawasm2go/p0.Fn778
func Fn778(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn780 github.com/goccy/llamawasm2go/p0.Fn780
func Fn780(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn788 github.com/goccy/llamawasm2go/p0.Fn788
func Fn788(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn790 github.com/goccy/llamawasm2go/p1.Fn790
func Fn790(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn798 github.com/goccy/llamawasm2go/p1.Fn798
func Fn798(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn803 github.com/goccy/llamawasm2go/p1.Fn803
func Fn803(m *base.Module)

//go:linkname Fn804 github.com/goccy/llamawasm2go/p0.Fn804
func Fn804(m *base.Module, l0 int64)

//go:linkname Fn940 github.com/goccy/llamawasm2go/p0.Fn940
func Fn940(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1001 github.com/goccy/llamawasm2go/p1.Fn1001
func Fn1001(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1005 github.com/goccy/llamawasm2go/p1.Fn1005
func Fn1005(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1018 github.com/goccy/llamawasm2go/p1.Fn1018
func Fn1018(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1023 github.com/goccy/llamawasm2go/p1.Fn1023
func Fn1023(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1032 github.com/goccy/llamawasm2go/p1.Fn1032
func Fn1032(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn1038 github.com/goccy/llamawasm2go/p1.Fn1038
func Fn1038(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1043 github.com/goccy/llamawasm2go/p1.Fn1043
func Fn1043(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1052 github.com/goccy/llamawasm2go/p1.Fn1052
func Fn1052(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1055 github.com/goccy/llamawasm2go/p1.Fn1055
func Fn1055(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1057 github.com/goccy/llamawasm2go/p1.Fn1057
func Fn1057(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1066 github.com/goccy/llamawasm2go/p1.Fn1066
func Fn1066(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1084 github.com/goccy/llamawasm2go/p1.Fn1084
func Fn1084(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1088 github.com/goccy/llamawasm2go/p0.Fn1088
func Fn1088(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1091 github.com/goccy/llamawasm2go/p1.Fn1091
func Fn1091(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1092 github.com/goccy/llamawasm2go/p1.Fn1092
func Fn1092(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1096 github.com/goccy/llamawasm2go/p1.Fn1096
func Fn1096(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1105 github.com/goccy/llamawasm2go/p0.Fn1105
func Fn1105(m *base.Module, l0 int64) int64

//go:linkname Fn1274 github.com/goccy/llamawasm2go/p1.Fn1274
func Fn1274(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1276 github.com/goccy/llamawasm2go/p0.Fn1276
func Fn1276(m *base.Module, l0 int64)

//go:linkname Fn1283 github.com/goccy/llamawasm2go/p0.Fn1283
func Fn1283(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1351 github.com/goccy/llamawasm2go/p1.Fn1351
func Fn1351(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1363 github.com/goccy/llamawasm2go/p1.Fn1363
func Fn1363(m *base.Module, l0 int64) int64

//go:linkname Fn1404 github.com/goccy/llamawasm2go/p1.Fn1404
func Fn1404(m *base.Module, l0 int64)

//go:linkname Fn1414 github.com/goccy/llamawasm2go/p1.Fn1414
func Fn1414(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1415 github.com/goccy/llamawasm2go/p1.Fn1415
func Fn1415(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1438 github.com/goccy/llamawasm2go/p1.Fn1438
func Fn1438(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1441 github.com/goccy/llamawasm2go/p1.Fn1441
func Fn1441(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1450 github.com/goccy/llamawasm2go/p1.Fn1450
func Fn1450(m *base.Module, l0 int64)

//go:linkname Fn1455 github.com/goccy/llamawasm2go/p1.Fn1455
func Fn1455(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1459 github.com/goccy/llamawasm2go/p0.Fn1459
func Fn1459(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1507 github.com/goccy/llamawasm2go/p1.Fn1507
func Fn1507(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1518 github.com/goccy/llamawasm2go/p1.Fn1518
func Fn1518(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1535 github.com/goccy/llamawasm2go/p1.Fn1535
func Fn1535(m *base.Module, l0 int64)

//go:linkname Fn1543 github.com/goccy/llamawasm2go/p1.Fn1543
func Fn1543(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1549 github.com/goccy/llamawasm2go/p1.Fn1549
func Fn1549(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1550 github.com/goccy/llamawasm2go/p1.Fn1550
func Fn1550(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1552 github.com/goccy/llamawasm2go/p1.Fn1552
func Fn1552(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p1.Fn1563
func Fn1563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1661 github.com/goccy/llamawasm2go/p0.Fn1661
func Fn1661(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1692 github.com/goccy/llamawasm2go/p0.Fn1692
func Fn1692(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1700 github.com/goccy/llamawasm2go/p1.Fn1700
func Fn1700(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1740 github.com/goccy/llamawasm2go/p1.Fn1740
func Fn1740(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1741 github.com/goccy/llamawasm2go/p0.Fn1741
func Fn1741(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1767 github.com/goccy/llamawasm2go/p1.Fn1767
func Fn1767(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1838 github.com/goccy/llamawasm2go/p1.Fn1838
func Fn1838(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1851 github.com/goccy/llamawasm2go/p1.Fn1851
func Fn1851(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1936 github.com/goccy/llamawasm2go/p1.Fn1936
func Fn1936(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1993 github.com/goccy/llamawasm2go/p1.Fn1993
func Fn1993(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1996 github.com/goccy/llamawasm2go/p1.Fn1996
func Fn1996(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2004 github.com/goccy/llamawasm2go/p1.Fn2004
func Fn2004(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2020 github.com/goccy/llamawasm2go/p1.Fn2020
func Fn2020(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2040 github.com/goccy/llamawasm2go/p1.Fn2040
func Fn2040(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2052 github.com/goccy/llamawasm2go/p1.Fn2052
func Fn2052(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2109 github.com/goccy/llamawasm2go/p1.Fn2109
func Fn2109(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2110 github.com/goccy/llamawasm2go/p1.Fn2110
func Fn2110(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2214 github.com/goccy/llamawasm2go/p1.Fn2214
func Fn2214(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2229 github.com/goccy/llamawasm2go/p1.Fn2229
func Fn2229(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2238 github.com/goccy/llamawasm2go/p1.Fn2238
func Fn2238(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2242 github.com/goccy/llamawasm2go/p1.Fn2242
func Fn2242(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2244 github.com/goccy/llamawasm2go/p1.Fn2244
func Fn2244(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2256 github.com/goccy/llamawasm2go/p1.Fn2256
func Fn2256(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2304 github.com/goccy/llamawasm2go/p1.Fn2304
func Fn2304(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2307 github.com/goccy/llamawasm2go/p1.Fn2307
func Fn2307(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2332 github.com/goccy/llamawasm2go/p1.Fn2332
func Fn2332(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2358 github.com/goccy/llamawasm2go/p0.Fn2358
func Fn2358(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2363 github.com/goccy/llamawasm2go/p1.Fn2363
func Fn2363(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2364 github.com/goccy/llamawasm2go/p0.Fn2364
func Fn2364(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2370 github.com/goccy/llamawasm2go/p1.Fn2370
func Fn2370(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2402 github.com/goccy/llamawasm2go/p1.Fn2402
func Fn2402(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2410 github.com/goccy/llamawasm2go/p0.Fn2410
func Fn2410(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2422 github.com/goccy/llamawasm2go/p1.Fn2422
func Fn2422(m *base.Module)

//go:linkname Fn2458 github.com/goccy/llamawasm2go/p0.Fn2458
func Fn2458(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2468 github.com/goccy/llamawasm2go/p1.Fn2468
func Fn2468(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2491 github.com/goccy/llamawasm2go/p1.Fn2491
func Fn2491(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2502 github.com/goccy/llamawasm2go/p1.Fn2502
func Fn2502(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2503 github.com/goccy/llamawasm2go/p1.Fn2503
func Fn2503(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2572 github.com/goccy/llamawasm2go/p1.Fn2572
func Fn2572(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2578 github.com/goccy/llamawasm2go/p1.Fn2578
func Fn2578(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2589 github.com/goccy/llamawasm2go/p1.Fn2589
func Fn2589(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2600 github.com/goccy/llamawasm2go/p1.Fn2600
func Fn2600(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2640 github.com/goccy/llamawasm2go/p1.Fn2640
func Fn2640(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2647 github.com/goccy/llamawasm2go/p1.Fn2647
func Fn2647(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2648 github.com/goccy/llamawasm2go/p1.Fn2648
func Fn2648(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2702 github.com/goccy/llamawasm2go/p0.Fn2702
func Fn2702(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2734 github.com/goccy/llamawasm2go/p1.Fn2734
func Fn2734(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2736 github.com/goccy/llamawasm2go/p1.Fn2736
func Fn2736(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2763 github.com/goccy/llamawasm2go/p0.Fn2763
func Fn2763(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2866 github.com/goccy/llamawasm2go/p1.Fn2866
func Fn2866(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2873 github.com/goccy/llamawasm2go/p1.Fn2873
func Fn2873(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2943 github.com/goccy/llamawasm2go/p1.Fn2943
func Fn2943(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2952 github.com/goccy/llamawasm2go/p1.Fn2952
func Fn2952(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2967 github.com/goccy/llamawasm2go/p1.Fn2967
func Fn2967(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2992 github.com/goccy/llamawasm2go/p1.Fn2992
func Fn2992(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3019 github.com/goccy/llamawasm2go/p1.Fn3019
func Fn3019(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3031 github.com/goccy/llamawasm2go/p1.Fn3031
func Fn3031(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3046 github.com/goccy/llamawasm2go/p1.Fn3046
func Fn3046(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3052 github.com/goccy/llamawasm2go/p0.Fn3052
func Fn3052(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3056 github.com/goccy/llamawasm2go/p0.Fn3056
func Fn3056(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3061 github.com/goccy/llamawasm2go/p0.Fn3061
func Fn3061(m *base.Module, l0 int64) int64

//go:linkname Fn3062 github.com/goccy/llamawasm2go/p1.Fn3062
func Fn3062(m *base.Module, l0 int64)

//go:linkname Fn3064 github.com/goccy/llamawasm2go/p1.Fn3064
func Fn3064(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3074 github.com/goccy/llamawasm2go/p1.Fn3074
func Fn3074(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3076 github.com/goccy/llamawasm2go/p1.Fn3076
func Fn3076(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3077 github.com/goccy/llamawasm2go/p1.Fn3077
func Fn3077(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3078 github.com/goccy/llamawasm2go/p1.Fn3078
func Fn3078(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3110 github.com/goccy/llamawasm2go/p1.Fn3110
func Fn3110(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)
