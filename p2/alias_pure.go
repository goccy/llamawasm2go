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

//go:linkname Fn355 github.com/goccy/llamawasm2go/p1.Fn355
func Fn355(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn357 github.com/goccy/llamawasm2go/p1.Fn357
func Fn357(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn360 github.com/goccy/llamawasm2go/p1.Fn360
func Fn360(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn365 github.com/goccy/llamawasm2go/p1.Fn365
func Fn365(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn373 github.com/goccy/llamawasm2go/p1.Fn373
func Fn373(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn395 github.com/goccy/llamawasm2go/p1.Fn395
func Fn395(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn398 github.com/goccy/llamawasm2go/p1.Fn398
func Fn398(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn442 github.com/goccy/llamawasm2go/p1.Fn442
func Fn442(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn532 github.com/goccy/llamawasm2go/p1.Fn532
func Fn532(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn534 github.com/goccy/llamawasm2go/p1.Fn534
func Fn534(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn540 github.com/goccy/llamawasm2go/p1.Fn540
func Fn540(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn550 github.com/goccy/llamawasm2go/p1.Fn550
func Fn550(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn554 github.com/goccy/llamawasm2go/p1.Fn554
func Fn554(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn605 github.com/goccy/llamawasm2go/p0.Fn605
func Fn605(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn646 github.com/goccy/llamawasm2go/p1.Fn646
func Fn646(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn650 github.com/goccy/llamawasm2go/p0.Fn650
func Fn650(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn660 github.com/goccy/llamawasm2go/p1.Fn660
func Fn660(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn684 github.com/goccy/llamawasm2go/p1.Fn684
func Fn684(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn702 github.com/goccy/llamawasm2go/p1.Fn702
func Fn702(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn733 github.com/goccy/llamawasm2go/p1.Fn733
func Fn733(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn734 github.com/goccy/llamawasm2go/p1.Fn734
func Fn734(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn748 github.com/goccy/llamawasm2go/p1.Fn748
func Fn748(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn750 github.com/goccy/llamawasm2go/p1.Fn750
func Fn750(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn754 github.com/goccy/llamawasm2go/p1.Fn754
func Fn754(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn755 github.com/goccy/llamawasm2go/p1.Fn755
func Fn755(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn772 github.com/goccy/llamawasm2go/p0.Fn772
func Fn772(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn774 github.com/goccy/llamawasm2go/p0.Fn774
func Fn774(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn782 github.com/goccy/llamawasm2go/p0.Fn782
func Fn782(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn784 github.com/goccy/llamawasm2go/p1.Fn784
func Fn784(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn792 github.com/goccy/llamawasm2go/p1.Fn792
func Fn792(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn797 github.com/goccy/llamawasm2go/p1.Fn797
func Fn797(m *base.Module)

//go:linkname Fn798 github.com/goccy/llamawasm2go/p0.Fn798
func Fn798(m *base.Module, l0 int64)

//go:linkname Fn859 github.com/goccy/llamawasm2go/p1.Fn859
func Fn859(m *base.Module, l0 int64) int64

//go:linkname Fn934 github.com/goccy/llamawasm2go/p0.Fn934
func Fn934(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn992 github.com/goccy/llamawasm2go/p1.Fn992
func Fn992(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn996 github.com/goccy/llamawasm2go/p1.Fn996
func Fn996(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1009 github.com/goccy/llamawasm2go/p1.Fn1009
func Fn1009(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1014 github.com/goccy/llamawasm2go/p1.Fn1014
func Fn1014(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1023 github.com/goccy/llamawasm2go/p1.Fn1023
func Fn1023(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn1029 github.com/goccy/llamawasm2go/p1.Fn1029
func Fn1029(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1034 github.com/goccy/llamawasm2go/p1.Fn1034
func Fn1034(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1043 github.com/goccy/llamawasm2go/p1.Fn1043
func Fn1043(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1046 github.com/goccy/llamawasm2go/p1.Fn1046
func Fn1046(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1048 github.com/goccy/llamawasm2go/p1.Fn1048
func Fn1048(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1057 github.com/goccy/llamawasm2go/p1.Fn1057
func Fn1057(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1075 github.com/goccy/llamawasm2go/p1.Fn1075
func Fn1075(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1079 github.com/goccy/llamawasm2go/p0.Fn1079
func Fn1079(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1083 github.com/goccy/llamawasm2go/p1.Fn1083
func Fn1083(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1087 github.com/goccy/llamawasm2go/p1.Fn1087
func Fn1087(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1096 github.com/goccy/llamawasm2go/p0.Fn1096
func Fn1096(m *base.Module, l0 int64) int64

//go:linkname Fn1265 github.com/goccy/llamawasm2go/p1.Fn1265
func Fn1265(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1267 github.com/goccy/llamawasm2go/p0.Fn1267
func Fn1267(m *base.Module, l0 int64)

//go:linkname Fn1274 github.com/goccy/llamawasm2go/p0.Fn1274
func Fn1274(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1342 github.com/goccy/llamawasm2go/p1.Fn1342
func Fn1342(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1355 github.com/goccy/llamawasm2go/p1.Fn1355
func Fn1355(m *base.Module, l0 int64) int64

//go:linkname Fn1396 github.com/goccy/llamawasm2go/p1.Fn1396
func Fn1396(m *base.Module, l0 int64)

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p1.Fn1406
func Fn1406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p1.Fn1407
func Fn1407(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1432 github.com/goccy/llamawasm2go/p1.Fn1432
func Fn1432(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1435 github.com/goccy/llamawasm2go/p1.Fn1435
func Fn1435(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p1.Fn1444
func Fn1444(m *base.Module, l0 int64)

//go:linkname Fn1453 github.com/goccy/llamawasm2go/p0.Fn1453
func Fn1453(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1499 github.com/goccy/llamawasm2go/p1.Fn1499
func Fn1499(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1510 github.com/goccy/llamawasm2go/p1.Fn1510
func Fn1510(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1527 github.com/goccy/llamawasm2go/p1.Fn1527
func Fn1527(m *base.Module, l0 int64)

//go:linkname Fn1535 github.com/goccy/llamawasm2go/p1.Fn1535
func Fn1535(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p1.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p1.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1544 github.com/goccy/llamawasm2go/p1.Fn1544
func Fn1544(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1555 github.com/goccy/llamawasm2go/p1.Fn1555
func Fn1555(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1653 github.com/goccy/llamawasm2go/p0.Fn1653
func Fn1653(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1684 github.com/goccy/llamawasm2go/p0.Fn1684
func Fn1684(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1692 github.com/goccy/llamawasm2go/p1.Fn1692
func Fn1692(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1732 github.com/goccy/llamawasm2go/p1.Fn1732
func Fn1732(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1733 github.com/goccy/llamawasm2go/p0.Fn1733
func Fn1733(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1759 github.com/goccy/llamawasm2go/p1.Fn1759
func Fn1759(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1830 github.com/goccy/llamawasm2go/p1.Fn1830
func Fn1830(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1843 github.com/goccy/llamawasm2go/p1.Fn1843
func Fn1843(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1928 github.com/goccy/llamawasm2go/p1.Fn1928
func Fn1928(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1985 github.com/goccy/llamawasm2go/p1.Fn1985
func Fn1985(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1988 github.com/goccy/llamawasm2go/p1.Fn1988
func Fn1988(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1996 github.com/goccy/llamawasm2go/p1.Fn1996
func Fn1996(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2012 github.com/goccy/llamawasm2go/p1.Fn2012
func Fn2012(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2032 github.com/goccy/llamawasm2go/p1.Fn2032
func Fn2032(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2044 github.com/goccy/llamawasm2go/p1.Fn2044
func Fn2044(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2101 github.com/goccy/llamawasm2go/p1.Fn2101
func Fn2101(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2102 github.com/goccy/llamawasm2go/p1.Fn2102
func Fn2102(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2206 github.com/goccy/llamawasm2go/p1.Fn2206
func Fn2206(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2221 github.com/goccy/llamawasm2go/p1.Fn2221
func Fn2221(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2230 github.com/goccy/llamawasm2go/p1.Fn2230
func Fn2230(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2234 github.com/goccy/llamawasm2go/p1.Fn2234
func Fn2234(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2236 github.com/goccy/llamawasm2go/p1.Fn2236
func Fn2236(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2248 github.com/goccy/llamawasm2go/p1.Fn2248
func Fn2248(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2296 github.com/goccy/llamawasm2go/p1.Fn2296
func Fn2296(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2299 github.com/goccy/llamawasm2go/p1.Fn2299
func Fn2299(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2324 github.com/goccy/llamawasm2go/p1.Fn2324
func Fn2324(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2350 github.com/goccy/llamawasm2go/p0.Fn2350
func Fn2350(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2355 github.com/goccy/llamawasm2go/p1.Fn2355
func Fn2355(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2356 github.com/goccy/llamawasm2go/p0.Fn2356
func Fn2356(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2362 github.com/goccy/llamawasm2go/p1.Fn2362
func Fn2362(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2394 github.com/goccy/llamawasm2go/p1.Fn2394
func Fn2394(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2402 github.com/goccy/llamawasm2go/p0.Fn2402
func Fn2402(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2414 github.com/goccy/llamawasm2go/p1.Fn2414
func Fn2414(m *base.Module)

//go:linkname Fn2450 github.com/goccy/llamawasm2go/p0.Fn2450
func Fn2450(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2460 github.com/goccy/llamawasm2go/p1.Fn2460
func Fn2460(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2483 github.com/goccy/llamawasm2go/p1.Fn2483
func Fn2483(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2494 github.com/goccy/llamawasm2go/p1.Fn2494
func Fn2494(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2495 github.com/goccy/llamawasm2go/p1.Fn2495
func Fn2495(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2564 github.com/goccy/llamawasm2go/p1.Fn2564
func Fn2564(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2570 github.com/goccy/llamawasm2go/p1.Fn2570
func Fn2570(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2581 github.com/goccy/llamawasm2go/p1.Fn2581
func Fn2581(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2592 github.com/goccy/llamawasm2go/p1.Fn2592
func Fn2592(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2632 github.com/goccy/llamawasm2go/p1.Fn2632
func Fn2632(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2639 github.com/goccy/llamawasm2go/p1.Fn2639
func Fn2639(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2640 github.com/goccy/llamawasm2go/p1.Fn2640
func Fn2640(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2694 github.com/goccy/llamawasm2go/p1.Fn2694
func Fn2694(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2726 github.com/goccy/llamawasm2go/p1.Fn2726
func Fn2726(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2728 github.com/goccy/llamawasm2go/p1.Fn2728
func Fn2728(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2755 github.com/goccy/llamawasm2go/p0.Fn2755
func Fn2755(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2858 github.com/goccy/llamawasm2go/p1.Fn2858
func Fn2858(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2865 github.com/goccy/llamawasm2go/p1.Fn2865
func Fn2865(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2935 github.com/goccy/llamawasm2go/p1.Fn2935
func Fn2935(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2944 github.com/goccy/llamawasm2go/p1.Fn2944
func Fn2944(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2959 github.com/goccy/llamawasm2go/p1.Fn2959
func Fn2959(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2984 github.com/goccy/llamawasm2go/p1.Fn2984
func Fn2984(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3011 github.com/goccy/llamawasm2go/p1.Fn3011
func Fn3011(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3023 github.com/goccy/llamawasm2go/p1.Fn3023
func Fn3023(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3038 github.com/goccy/llamawasm2go/p1.Fn3038
func Fn3038(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3044 github.com/goccy/llamawasm2go/p0.Fn3044
func Fn3044(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3048 github.com/goccy/llamawasm2go/p0.Fn3048
func Fn3048(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3053 github.com/goccy/llamawasm2go/p0.Fn3053
func Fn3053(m *base.Module, l0 int64) int64

//go:linkname Fn3054 github.com/goccy/llamawasm2go/p1.Fn3054
func Fn3054(m *base.Module, l0 int64)

//go:linkname Fn3056 github.com/goccy/llamawasm2go/p1.Fn3056
func Fn3056(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3066 github.com/goccy/llamawasm2go/p1.Fn3066
func Fn3066(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3068 github.com/goccy/llamawasm2go/p1.Fn3068
func Fn3068(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3069 github.com/goccy/llamawasm2go/p1.Fn3069
func Fn3069(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3070 github.com/goccy/llamawasm2go/p1.Fn3070
func Fn3070(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3102 github.com/goccy/llamawasm2go/p1.Fn3102
func Fn3102(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)
