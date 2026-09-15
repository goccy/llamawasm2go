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

//go:linkname Fn366 github.com/goccy/llamawasm2go/p1.Fn366
func Fn366(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn370 github.com/goccy/llamawasm2go/p1.Fn370
func Fn370(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn371 github.com/goccy/llamawasm2go/p1.Fn371
func Fn371(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn372 github.com/goccy/llamawasm2go/p1.Fn372
func Fn372(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn376 github.com/goccy/llamawasm2go/p1.Fn376
func Fn376(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn383 github.com/goccy/llamawasm2go/p1.Fn383
func Fn383(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32)

//go:linkname Fn386 github.com/goccy/llamawasm2go/p1.Fn386
func Fn386(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn390 github.com/goccy/llamawasm2go/p1.Fn390
func Fn390(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn403 github.com/goccy/llamawasm2go/p0.Fn403
func Fn403(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn405 github.com/goccy/llamawasm2go/p1.Fn405
func Fn405(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn407 github.com/goccy/llamawasm2go/p0.Fn407
func Fn407(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn410 github.com/goccy/llamawasm2go/p1.Fn410
func Fn410(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn411 github.com/goccy/llamawasm2go/p1.Fn411
func Fn411(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn420 github.com/goccy/llamawasm2go/p1.Fn420
func Fn420(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn421 github.com/goccy/llamawasm2go/p0.Fn421
func Fn421(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn425 github.com/goccy/llamawasm2go/p1.Fn425
func Fn425(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn427 github.com/goccy/llamawasm2go/p1.Fn427
func Fn427(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn428 github.com/goccy/llamawasm2go/p1.Fn428
func Fn428(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32)

//go:linkname Fn429 github.com/goccy/llamawasm2go/p1.Fn429
func Fn429(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn430 github.com/goccy/llamawasm2go/p1.Fn430
func Fn430(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn478 github.com/goccy/llamawasm2go/p1.Fn478
func Fn478(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn568 github.com/goccy/llamawasm2go/p1.Fn568
func Fn568(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn570 github.com/goccy/llamawasm2go/p1.Fn570
func Fn570(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn576 github.com/goccy/llamawasm2go/p1.Fn576
func Fn576(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn586 github.com/goccy/llamawasm2go/p1.Fn586
func Fn586(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn590 github.com/goccy/llamawasm2go/p1.Fn590
func Fn590(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn641 github.com/goccy/llamawasm2go/p0.Fn641
func Fn641(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn682 github.com/goccy/llamawasm2go/p1.Fn682
func Fn682(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn686 github.com/goccy/llamawasm2go/p0.Fn686
func Fn686(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn696 github.com/goccy/llamawasm2go/p1.Fn696
func Fn696(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn720 github.com/goccy/llamawasm2go/p1.Fn720
func Fn720(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn738 github.com/goccy/llamawasm2go/p1.Fn738
func Fn738(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn769 github.com/goccy/llamawasm2go/p1.Fn769
func Fn769(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn770 github.com/goccy/llamawasm2go/p1.Fn770
func Fn770(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn784 github.com/goccy/llamawasm2go/p1.Fn784
func Fn784(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn786 github.com/goccy/llamawasm2go/p0.Fn786
func Fn786(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn790 github.com/goccy/llamawasm2go/p1.Fn790
func Fn790(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn791 github.com/goccy/llamawasm2go/p1.Fn791
func Fn791(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn808 github.com/goccy/llamawasm2go/p0.Fn808
func Fn808(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn810 github.com/goccy/llamawasm2go/p0.Fn810
func Fn810(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn818 github.com/goccy/llamawasm2go/p0.Fn818
func Fn818(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn820 github.com/goccy/llamawasm2go/p1.Fn820
func Fn820(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn828 github.com/goccy/llamawasm2go/p1.Fn828
func Fn828(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn833 github.com/goccy/llamawasm2go/p1.Fn833
func Fn833(m *base.Module)

//go:linkname Fn834 github.com/goccy/llamawasm2go/p0.Fn834
func Fn834(m *base.Module, l0 int64)

//go:linkname Fn977 github.com/goccy/llamawasm2go/p0.Fn977
func Fn977(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1042 github.com/goccy/llamawasm2go/p1.Fn1042
func Fn1042(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1046 github.com/goccy/llamawasm2go/p1.Fn1046
func Fn1046(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1059 github.com/goccy/llamawasm2go/p1.Fn1059
func Fn1059(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1064 github.com/goccy/llamawasm2go/p1.Fn1064
func Fn1064(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1079 github.com/goccy/llamawasm2go/p1.Fn1079
func Fn1079(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1084 github.com/goccy/llamawasm2go/p1.Fn1084
func Fn1084(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1093 github.com/goccy/llamawasm2go/p1.Fn1093
func Fn1093(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1096 github.com/goccy/llamawasm2go/p1.Fn1096
func Fn1096(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1098 github.com/goccy/llamawasm2go/p1.Fn1098
func Fn1098(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1107 github.com/goccy/llamawasm2go/p1.Fn1107
func Fn1107(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1125 github.com/goccy/llamawasm2go/p0.Fn1125
func Fn1125(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1129 github.com/goccy/llamawasm2go/p0.Fn1129
func Fn1129(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1133 github.com/goccy/llamawasm2go/p1.Fn1133
func Fn1133(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1137 github.com/goccy/llamawasm2go/p1.Fn1137
func Fn1137(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1146 github.com/goccy/llamawasm2go/p0.Fn1146
func Fn1146(m *base.Module, l0 int64) int64

//go:linkname Fn1315 github.com/goccy/llamawasm2go/p1.Fn1315
func Fn1315(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1317 github.com/goccy/llamawasm2go/p0.Fn1317
func Fn1317(m *base.Module, l0 int64)

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p0.Fn1324
func Fn1324(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p1.Fn1402
func Fn1402(m *base.Module, l0 int64) int64

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p1.Fn1443
func Fn1443(m *base.Module, l0 int64)

//go:linkname Fn1453 github.com/goccy/llamawasm2go/p1.Fn1453
func Fn1453(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1454 github.com/goccy/llamawasm2go/p1.Fn1454
func Fn1454(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1477 github.com/goccy/llamawasm2go/p1.Fn1477
func Fn1477(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1480 github.com/goccy/llamawasm2go/p1.Fn1480
func Fn1480(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1489 github.com/goccy/llamawasm2go/p1.Fn1489
func Fn1489(m *base.Module, l0 int64)

//go:linkname Fn1494 github.com/goccy/llamawasm2go/p1.Fn1494
func Fn1494(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1498 github.com/goccy/llamawasm2go/p0.Fn1498
func Fn1498(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1511 github.com/goccy/llamawasm2go/p1.Fn1511
func Fn1511(m *base.Module, l0 int64)

//go:linkname Fn1550 github.com/goccy/llamawasm2go/p1.Fn1550
func Fn1550(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1561 github.com/goccy/llamawasm2go/p1.Fn1561
func Fn1561(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1578 github.com/goccy/llamawasm2go/p1.Fn1578
func Fn1578(m *base.Module, l0 int64)

//go:linkname Fn1586 github.com/goccy/llamawasm2go/p1.Fn1586
func Fn1586(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1592 github.com/goccy/llamawasm2go/p1.Fn1592
func Fn1592(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1593 github.com/goccy/llamawasm2go/p1.Fn1593
func Fn1593(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1595 github.com/goccy/llamawasm2go/p1.Fn1595
func Fn1595(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1606 github.com/goccy/llamawasm2go/p1.Fn1606
func Fn1606(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1704 github.com/goccy/llamawasm2go/p0.Fn1704
func Fn1704(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1735 github.com/goccy/llamawasm2go/p0.Fn1735
func Fn1735(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1743 github.com/goccy/llamawasm2go/p1.Fn1743
func Fn1743(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1783 github.com/goccy/llamawasm2go/p1.Fn1783
func Fn1783(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1784 github.com/goccy/llamawasm2go/p0.Fn1784
func Fn1784(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1810 github.com/goccy/llamawasm2go/p1.Fn1810
func Fn1810(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1881 github.com/goccy/llamawasm2go/p1.Fn1881
func Fn1881(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1894 github.com/goccy/llamawasm2go/p1.Fn1894
func Fn1894(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1979 github.com/goccy/llamawasm2go/p1.Fn1979
func Fn1979(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2036 github.com/goccy/llamawasm2go/p1.Fn2036
func Fn2036(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2039 github.com/goccy/llamawasm2go/p1.Fn2039
func Fn2039(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2047 github.com/goccy/llamawasm2go/p1.Fn2047
func Fn2047(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2063 github.com/goccy/llamawasm2go/p1.Fn2063
func Fn2063(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2083 github.com/goccy/llamawasm2go/p1.Fn2083
func Fn2083(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2095 github.com/goccy/llamawasm2go/p1.Fn2095
func Fn2095(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2152 github.com/goccy/llamawasm2go/p1.Fn2152
func Fn2152(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2153 github.com/goccy/llamawasm2go/p1.Fn2153
func Fn2153(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2255 github.com/goccy/llamawasm2go/p1.Fn2255
func Fn2255(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2263 github.com/goccy/llamawasm2go/p0.Fn2263
func Fn2263(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2279 github.com/goccy/llamawasm2go/p1.Fn2279
func Fn2279(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2283 github.com/goccy/llamawasm2go/p1.Fn2283
func Fn2283(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2285 github.com/goccy/llamawasm2go/p1.Fn2285
func Fn2285(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2297 github.com/goccy/llamawasm2go/p1.Fn2297
func Fn2297(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2345 github.com/goccy/llamawasm2go/p1.Fn2345
func Fn2345(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2348 github.com/goccy/llamawasm2go/p1.Fn2348
func Fn2348(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2373 github.com/goccy/llamawasm2go/p1.Fn2373
func Fn2373(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2399 github.com/goccy/llamawasm2go/p0.Fn2399
func Fn2399(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2404 github.com/goccy/llamawasm2go/p1.Fn2404
func Fn2404(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2405 github.com/goccy/llamawasm2go/p0.Fn2405
func Fn2405(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2411 github.com/goccy/llamawasm2go/p1.Fn2411
func Fn2411(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2443 github.com/goccy/llamawasm2go/p1.Fn2443
func Fn2443(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2451 github.com/goccy/llamawasm2go/p0.Fn2451
func Fn2451(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2463 github.com/goccy/llamawasm2go/p1.Fn2463
func Fn2463(m *base.Module)

//go:linkname Fn2498 github.com/goccy/llamawasm2go/p0.Fn2498
func Fn2498(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2508 github.com/goccy/llamawasm2go/p1.Fn2508
func Fn2508(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2531 github.com/goccy/llamawasm2go/p1.Fn2531
func Fn2531(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2542 github.com/goccy/llamawasm2go/p1.Fn2542
func Fn2542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2543 github.com/goccy/llamawasm2go/p1.Fn2543
func Fn2543(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2612 github.com/goccy/llamawasm2go/p1.Fn2612
func Fn2612(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2618 github.com/goccy/llamawasm2go/p1.Fn2618
func Fn2618(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2629 github.com/goccy/llamawasm2go/p1.Fn2629
func Fn2629(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2640 github.com/goccy/llamawasm2go/p1.Fn2640
func Fn2640(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2680 github.com/goccy/llamawasm2go/p1.Fn2680
func Fn2680(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2687 github.com/goccy/llamawasm2go/p1.Fn2687
func Fn2687(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2688 github.com/goccy/llamawasm2go/p1.Fn2688
func Fn2688(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2742 github.com/goccy/llamawasm2go/p0.Fn2742
func Fn2742(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2774 github.com/goccy/llamawasm2go/p1.Fn2774
func Fn2774(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2776 github.com/goccy/llamawasm2go/p0.Fn2776
func Fn2776(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2803 github.com/goccy/llamawasm2go/p0.Fn2803
func Fn2803(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2906 github.com/goccy/llamawasm2go/p1.Fn2906
func Fn2906(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2913 github.com/goccy/llamawasm2go/p1.Fn2913
func Fn2913(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2992 github.com/goccy/llamawasm2go/p1.Fn2992
func Fn2992(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3007 github.com/goccy/llamawasm2go/p1.Fn3007
func Fn3007(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn3059 github.com/goccy/llamawasm2go/p1.Fn3059
func Fn3059(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3071 github.com/goccy/llamawasm2go/p1.Fn3071
func Fn3071(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3086 github.com/goccy/llamawasm2go/p1.Fn3086
func Fn3086(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3092 github.com/goccy/llamawasm2go/p0.Fn3092
func Fn3092(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3096 github.com/goccy/llamawasm2go/p0.Fn3096
func Fn3096(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3101 github.com/goccy/llamawasm2go/p0.Fn3101
func Fn3101(m *base.Module, l0 int64) int64

//go:linkname Fn3102 github.com/goccy/llamawasm2go/p1.Fn3102
func Fn3102(m *base.Module, l0 int64)

//go:linkname Fn3104 github.com/goccy/llamawasm2go/p1.Fn3104
func Fn3104(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3114 github.com/goccy/llamawasm2go/p1.Fn3114
func Fn3114(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3116 github.com/goccy/llamawasm2go/p1.Fn3116
func Fn3116(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3117 github.com/goccy/llamawasm2go/p1.Fn3117
func Fn3117(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3118 github.com/goccy/llamawasm2go/p1.Fn3118
func Fn3118(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3150 github.com/goccy/llamawasm2go/p1.Fn3150
func Fn3150(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)
