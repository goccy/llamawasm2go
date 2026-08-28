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

//go:linkname Fn939 github.com/goccy/llamawasm2go/p0.Fn939
func Fn939(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn999 github.com/goccy/llamawasm2go/p1.Fn999
func Fn999(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1003 github.com/goccy/llamawasm2go/p1.Fn1003
func Fn1003(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1016 github.com/goccy/llamawasm2go/p1.Fn1016
func Fn1016(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1021 github.com/goccy/llamawasm2go/p1.Fn1021
func Fn1021(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1030 github.com/goccy/llamawasm2go/p1.Fn1030
func Fn1030(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn1036 github.com/goccy/llamawasm2go/p1.Fn1036
func Fn1036(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1041 github.com/goccy/llamawasm2go/p1.Fn1041
func Fn1041(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1050 github.com/goccy/llamawasm2go/p1.Fn1050
func Fn1050(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1053 github.com/goccy/llamawasm2go/p1.Fn1053
func Fn1053(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1055 github.com/goccy/llamawasm2go/p1.Fn1055
func Fn1055(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1064 github.com/goccy/llamawasm2go/p1.Fn1064
func Fn1064(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1082 github.com/goccy/llamawasm2go/p1.Fn1082
func Fn1082(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1086 github.com/goccy/llamawasm2go/p0.Fn1086
func Fn1086(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1090 github.com/goccy/llamawasm2go/p1.Fn1090
func Fn1090(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1094 github.com/goccy/llamawasm2go/p1.Fn1094
func Fn1094(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1103 github.com/goccy/llamawasm2go/p0.Fn1103
func Fn1103(m *base.Module, l0 int64) int64

//go:linkname Fn1272 github.com/goccy/llamawasm2go/p1.Fn1272
func Fn1272(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1274 github.com/goccy/llamawasm2go/p0.Fn1274
func Fn1274(m *base.Module, l0 int64)

//go:linkname Fn1281 github.com/goccy/llamawasm2go/p0.Fn1281
func Fn1281(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1349 github.com/goccy/llamawasm2go/p1.Fn1349
func Fn1349(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1361 github.com/goccy/llamawasm2go/p1.Fn1361
func Fn1361(m *base.Module, l0 int64) int64

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p1.Fn1402
func Fn1402(m *base.Module, l0 int64)

//go:linkname Fn1412 github.com/goccy/llamawasm2go/p1.Fn1412
func Fn1412(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p1.Fn1413
func Fn1413(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1436 github.com/goccy/llamawasm2go/p1.Fn1436
func Fn1436(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1439 github.com/goccy/llamawasm2go/p1.Fn1439
func Fn1439(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1448 github.com/goccy/llamawasm2go/p1.Fn1448
func Fn1448(m *base.Module, l0 int64)

//go:linkname Fn1453 github.com/goccy/llamawasm2go/p1.Fn1453
func Fn1453(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1457 github.com/goccy/llamawasm2go/p0.Fn1457
func Fn1457(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1505 github.com/goccy/llamawasm2go/p1.Fn1505
func Fn1505(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1516 github.com/goccy/llamawasm2go/p1.Fn1516
func Fn1516(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1533 github.com/goccy/llamawasm2go/p1.Fn1533
func Fn1533(m *base.Module, l0 int64)

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p1.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1547 github.com/goccy/llamawasm2go/p1.Fn1547
func Fn1547(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1548 github.com/goccy/llamawasm2go/p1.Fn1548
func Fn1548(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1550 github.com/goccy/llamawasm2go/p1.Fn1550
func Fn1550(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1561 github.com/goccy/llamawasm2go/p1.Fn1561
func Fn1561(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1659 github.com/goccy/llamawasm2go/p0.Fn1659
func Fn1659(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1690 github.com/goccy/llamawasm2go/p0.Fn1690
func Fn1690(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1698 github.com/goccy/llamawasm2go/p1.Fn1698
func Fn1698(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1738 github.com/goccy/llamawasm2go/p1.Fn1738
func Fn1738(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1739 github.com/goccy/llamawasm2go/p0.Fn1739
func Fn1739(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1765 github.com/goccy/llamawasm2go/p1.Fn1765
func Fn1765(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1836 github.com/goccy/llamawasm2go/p1.Fn1836
func Fn1836(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1849 github.com/goccy/llamawasm2go/p1.Fn1849
func Fn1849(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1934 github.com/goccy/llamawasm2go/p1.Fn1934
func Fn1934(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1991 github.com/goccy/llamawasm2go/p1.Fn1991
func Fn1991(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1994 github.com/goccy/llamawasm2go/p1.Fn1994
func Fn1994(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2002 github.com/goccy/llamawasm2go/p1.Fn2002
func Fn2002(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2018 github.com/goccy/llamawasm2go/p1.Fn2018
func Fn2018(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2038 github.com/goccy/llamawasm2go/p1.Fn2038
func Fn2038(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2050 github.com/goccy/llamawasm2go/p1.Fn2050
func Fn2050(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2107 github.com/goccy/llamawasm2go/p1.Fn2107
func Fn2107(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2108 github.com/goccy/llamawasm2go/p1.Fn2108
func Fn2108(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2212 github.com/goccy/llamawasm2go/p1.Fn2212
func Fn2212(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2227 github.com/goccy/llamawasm2go/p1.Fn2227
func Fn2227(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2236 github.com/goccy/llamawasm2go/p1.Fn2236
func Fn2236(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2240 github.com/goccy/llamawasm2go/p1.Fn2240
func Fn2240(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2242 github.com/goccy/llamawasm2go/p1.Fn2242
func Fn2242(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2254 github.com/goccy/llamawasm2go/p1.Fn2254
func Fn2254(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2302 github.com/goccy/llamawasm2go/p1.Fn2302
func Fn2302(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2305 github.com/goccy/llamawasm2go/p1.Fn2305
func Fn2305(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2330 github.com/goccy/llamawasm2go/p1.Fn2330
func Fn2330(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2356 github.com/goccy/llamawasm2go/p0.Fn2356
func Fn2356(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2361 github.com/goccy/llamawasm2go/p1.Fn2361
func Fn2361(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2362 github.com/goccy/llamawasm2go/p0.Fn2362
func Fn2362(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2368 github.com/goccy/llamawasm2go/p1.Fn2368
func Fn2368(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2400 github.com/goccy/llamawasm2go/p1.Fn2400
func Fn2400(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2408 github.com/goccy/llamawasm2go/p0.Fn2408
func Fn2408(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2420 github.com/goccy/llamawasm2go/p1.Fn2420
func Fn2420(m *base.Module)

//go:linkname Fn2456 github.com/goccy/llamawasm2go/p0.Fn2456
func Fn2456(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2466 github.com/goccy/llamawasm2go/p1.Fn2466
func Fn2466(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2489 github.com/goccy/llamawasm2go/p1.Fn2489
func Fn2489(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2500 github.com/goccy/llamawasm2go/p1.Fn2500
func Fn2500(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2501 github.com/goccy/llamawasm2go/p1.Fn2501
func Fn2501(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2570 github.com/goccy/llamawasm2go/p1.Fn2570
func Fn2570(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2576 github.com/goccy/llamawasm2go/p1.Fn2576
func Fn2576(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2587 github.com/goccy/llamawasm2go/p1.Fn2587
func Fn2587(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2598 github.com/goccy/llamawasm2go/p1.Fn2598
func Fn2598(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2638 github.com/goccy/llamawasm2go/p1.Fn2638
func Fn2638(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2645 github.com/goccy/llamawasm2go/p1.Fn2645
func Fn2645(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2646 github.com/goccy/llamawasm2go/p1.Fn2646
func Fn2646(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2700 github.com/goccy/llamawasm2go/p1.Fn2700
func Fn2700(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2732 github.com/goccy/llamawasm2go/p1.Fn2732
func Fn2732(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2734 github.com/goccy/llamawasm2go/p1.Fn2734
func Fn2734(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2761 github.com/goccy/llamawasm2go/p0.Fn2761
func Fn2761(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2864 github.com/goccy/llamawasm2go/p1.Fn2864
func Fn2864(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2871 github.com/goccy/llamawasm2go/p1.Fn2871
func Fn2871(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2941 github.com/goccy/llamawasm2go/p1.Fn2941
func Fn2941(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2950 github.com/goccy/llamawasm2go/p1.Fn2950
func Fn2950(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2965 github.com/goccy/llamawasm2go/p1.Fn2965
func Fn2965(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2990 github.com/goccy/llamawasm2go/p1.Fn2990
func Fn2990(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3017 github.com/goccy/llamawasm2go/p1.Fn3017
func Fn3017(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3029 github.com/goccy/llamawasm2go/p1.Fn3029
func Fn3029(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3044 github.com/goccy/llamawasm2go/p1.Fn3044
func Fn3044(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3050 github.com/goccy/llamawasm2go/p0.Fn3050
func Fn3050(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3054 github.com/goccy/llamawasm2go/p0.Fn3054
func Fn3054(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3059 github.com/goccy/llamawasm2go/p0.Fn3059
func Fn3059(m *base.Module, l0 int64) int64

//go:linkname Fn3060 github.com/goccy/llamawasm2go/p1.Fn3060
func Fn3060(m *base.Module, l0 int64)

//go:linkname Fn3062 github.com/goccy/llamawasm2go/p1.Fn3062
func Fn3062(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3072 github.com/goccy/llamawasm2go/p1.Fn3072
func Fn3072(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3074 github.com/goccy/llamawasm2go/p1.Fn3074
func Fn3074(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3075 github.com/goccy/llamawasm2go/p1.Fn3075
func Fn3075(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3076 github.com/goccy/llamawasm2go/p1.Fn3076
func Fn3076(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3108 github.com/goccy/llamawasm2go/p1.Fn3108
func Fn3108(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)
