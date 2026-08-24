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

//go:linkname Fn244 github.com/goccy/llamawasm2go/p0.Fn244
func Fn244(m *base.Module, l0 int64) int64

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

//go:linkname Fn933 github.com/goccy/llamawasm2go/p0.Fn933
func Fn933(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn993 github.com/goccy/llamawasm2go/p1.Fn993
func Fn993(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn997 github.com/goccy/llamawasm2go/p1.Fn997
func Fn997(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1010 github.com/goccy/llamawasm2go/p1.Fn1010
func Fn1010(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1015 github.com/goccy/llamawasm2go/p1.Fn1015
func Fn1015(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1024 github.com/goccy/llamawasm2go/p1.Fn1024
func Fn1024(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn1030 github.com/goccy/llamawasm2go/p1.Fn1030
func Fn1030(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1035 github.com/goccy/llamawasm2go/p1.Fn1035
func Fn1035(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1044 github.com/goccy/llamawasm2go/p1.Fn1044
func Fn1044(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1047 github.com/goccy/llamawasm2go/p1.Fn1047
func Fn1047(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1049 github.com/goccy/llamawasm2go/p1.Fn1049
func Fn1049(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1058 github.com/goccy/llamawasm2go/p1.Fn1058
func Fn1058(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1076 github.com/goccy/llamawasm2go/p1.Fn1076
func Fn1076(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1080 github.com/goccy/llamawasm2go/p0.Fn1080
func Fn1080(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1084 github.com/goccy/llamawasm2go/p1.Fn1084
func Fn1084(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1088 github.com/goccy/llamawasm2go/p1.Fn1088
func Fn1088(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1097 github.com/goccy/llamawasm2go/p0.Fn1097
func Fn1097(m *base.Module, l0 int64) int64

//go:linkname Fn1266 github.com/goccy/llamawasm2go/p1.Fn1266
func Fn1266(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1268 github.com/goccy/llamawasm2go/p0.Fn1268
func Fn1268(m *base.Module, l0 int64)

//go:linkname Fn1275 github.com/goccy/llamawasm2go/p0.Fn1275
func Fn1275(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1343 github.com/goccy/llamawasm2go/p1.Fn1343
func Fn1343(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1356 github.com/goccy/llamawasm2go/p1.Fn1356
func Fn1356(m *base.Module, l0 int64) int64

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p1.Fn1397
func Fn1397(m *base.Module, l0 int64)

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p1.Fn1407
func Fn1407(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p1.Fn1408
func Fn1408(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1433 github.com/goccy/llamawasm2go/p1.Fn1433
func Fn1433(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1436 github.com/goccy/llamawasm2go/p1.Fn1436
func Fn1436(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p1.Fn1445
func Fn1445(m *base.Module, l0 int64)

//go:linkname Fn1454 github.com/goccy/llamawasm2go/p0.Fn1454
func Fn1454(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1500 github.com/goccy/llamawasm2go/p1.Fn1500
func Fn1500(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1511 github.com/goccy/llamawasm2go/p1.Fn1511
func Fn1511(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1528 github.com/goccy/llamawasm2go/p1.Fn1528
func Fn1528(m *base.Module, l0 int64)

//go:linkname Fn1536 github.com/goccy/llamawasm2go/p1.Fn1536
func Fn1536(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p1.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1543 github.com/goccy/llamawasm2go/p1.Fn1543
func Fn1543(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p1.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1556 github.com/goccy/llamawasm2go/p1.Fn1556
func Fn1556(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1654 github.com/goccy/llamawasm2go/p0.Fn1654
func Fn1654(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1685 github.com/goccy/llamawasm2go/p0.Fn1685
func Fn1685(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1693 github.com/goccy/llamawasm2go/p1.Fn1693
func Fn1693(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1733 github.com/goccy/llamawasm2go/p1.Fn1733
func Fn1733(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1734 github.com/goccy/llamawasm2go/p0.Fn1734
func Fn1734(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1760 github.com/goccy/llamawasm2go/p1.Fn1760
func Fn1760(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1831 github.com/goccy/llamawasm2go/p1.Fn1831
func Fn1831(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1844 github.com/goccy/llamawasm2go/p1.Fn1844
func Fn1844(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1929 github.com/goccy/llamawasm2go/p1.Fn1929
func Fn1929(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1986 github.com/goccy/llamawasm2go/p1.Fn1986
func Fn1986(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1989 github.com/goccy/llamawasm2go/p1.Fn1989
func Fn1989(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1997 github.com/goccy/llamawasm2go/p1.Fn1997
func Fn1997(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2013 github.com/goccy/llamawasm2go/p1.Fn2013
func Fn2013(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2033 github.com/goccy/llamawasm2go/p1.Fn2033
func Fn2033(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2045 github.com/goccy/llamawasm2go/p1.Fn2045
func Fn2045(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2102 github.com/goccy/llamawasm2go/p1.Fn2102
func Fn2102(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2103 github.com/goccy/llamawasm2go/p1.Fn2103
func Fn2103(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2207 github.com/goccy/llamawasm2go/p1.Fn2207
func Fn2207(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2222 github.com/goccy/llamawasm2go/p1.Fn2222
func Fn2222(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2231 github.com/goccy/llamawasm2go/p1.Fn2231
func Fn2231(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2235 github.com/goccy/llamawasm2go/p1.Fn2235
func Fn2235(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2237 github.com/goccy/llamawasm2go/p1.Fn2237
func Fn2237(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p1.Fn2249
func Fn2249(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2297 github.com/goccy/llamawasm2go/p1.Fn2297
func Fn2297(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2300 github.com/goccy/llamawasm2go/p1.Fn2300
func Fn2300(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2325 github.com/goccy/llamawasm2go/p1.Fn2325
func Fn2325(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2351 github.com/goccy/llamawasm2go/p0.Fn2351
func Fn2351(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2356 github.com/goccy/llamawasm2go/p1.Fn2356
func Fn2356(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2357 github.com/goccy/llamawasm2go/p0.Fn2357
func Fn2357(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2363 github.com/goccy/llamawasm2go/p1.Fn2363
func Fn2363(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2395 github.com/goccy/llamawasm2go/p1.Fn2395
func Fn2395(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2403 github.com/goccy/llamawasm2go/p0.Fn2403
func Fn2403(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2415 github.com/goccy/llamawasm2go/p1.Fn2415
func Fn2415(m *base.Module)

//go:linkname Fn2451 github.com/goccy/llamawasm2go/p0.Fn2451
func Fn2451(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2461 github.com/goccy/llamawasm2go/p1.Fn2461
func Fn2461(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2484 github.com/goccy/llamawasm2go/p1.Fn2484
func Fn2484(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2495 github.com/goccy/llamawasm2go/p1.Fn2495
func Fn2495(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2496 github.com/goccy/llamawasm2go/p1.Fn2496
func Fn2496(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2565 github.com/goccy/llamawasm2go/p1.Fn2565
func Fn2565(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2571 github.com/goccy/llamawasm2go/p1.Fn2571
func Fn2571(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2582 github.com/goccy/llamawasm2go/p1.Fn2582
func Fn2582(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2593 github.com/goccy/llamawasm2go/p1.Fn2593
func Fn2593(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2633 github.com/goccy/llamawasm2go/p1.Fn2633
func Fn2633(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2640 github.com/goccy/llamawasm2go/p1.Fn2640
func Fn2640(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2641 github.com/goccy/llamawasm2go/p1.Fn2641
func Fn2641(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2695 github.com/goccy/llamawasm2go/p0.Fn2695
func Fn2695(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2727 github.com/goccy/llamawasm2go/p1.Fn2727
func Fn2727(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2729 github.com/goccy/llamawasm2go/p1.Fn2729
func Fn2729(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2756 github.com/goccy/llamawasm2go/p0.Fn2756
func Fn2756(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2859 github.com/goccy/llamawasm2go/p1.Fn2859
func Fn2859(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2866 github.com/goccy/llamawasm2go/p1.Fn2866
func Fn2866(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2936 github.com/goccy/llamawasm2go/p1.Fn2936
func Fn2936(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2945 github.com/goccy/llamawasm2go/p1.Fn2945
func Fn2945(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2960 github.com/goccy/llamawasm2go/p1.Fn2960
func Fn2960(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2985 github.com/goccy/llamawasm2go/p1.Fn2985
func Fn2985(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3012 github.com/goccy/llamawasm2go/p1.Fn3012
func Fn3012(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3024 github.com/goccy/llamawasm2go/p1.Fn3024
func Fn3024(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3039 github.com/goccy/llamawasm2go/p1.Fn3039
func Fn3039(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3045 github.com/goccy/llamawasm2go/p0.Fn3045
func Fn3045(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3049 github.com/goccy/llamawasm2go/p0.Fn3049
func Fn3049(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3054 github.com/goccy/llamawasm2go/p0.Fn3054
func Fn3054(m *base.Module, l0 int64) int64

//go:linkname Fn3055 github.com/goccy/llamawasm2go/p1.Fn3055
func Fn3055(m *base.Module, l0 int64)

//go:linkname Fn3057 github.com/goccy/llamawasm2go/p1.Fn3057
func Fn3057(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3067 github.com/goccy/llamawasm2go/p1.Fn3067
func Fn3067(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3069 github.com/goccy/llamawasm2go/p1.Fn3069
func Fn3069(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3070 github.com/goccy/llamawasm2go/p1.Fn3070
func Fn3070(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3071 github.com/goccy/llamawasm2go/p1.Fn3071
func Fn3071(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3103 github.com/goccy/llamawasm2go/p1.Fn3103
func Fn3103(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)
