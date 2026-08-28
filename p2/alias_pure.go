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

//go:linkname Fn1501 github.com/goccy/llamawasm2go/p1.Fn1501
func Fn1501(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1512 github.com/goccy/llamawasm2go/p1.Fn1512
func Fn1512(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1529 github.com/goccy/llamawasm2go/p1.Fn1529
func Fn1529(m *base.Module, l0 int64)

//go:linkname Fn1537 github.com/goccy/llamawasm2go/p1.Fn1537
func Fn1537(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1543 github.com/goccy/llamawasm2go/p1.Fn1543
func Fn1543(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1544 github.com/goccy/llamawasm2go/p1.Fn1544
func Fn1544(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1546 github.com/goccy/llamawasm2go/p1.Fn1546
func Fn1546(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1557 github.com/goccy/llamawasm2go/p1.Fn1557
func Fn1557(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1655 github.com/goccy/llamawasm2go/p0.Fn1655
func Fn1655(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1686 github.com/goccy/llamawasm2go/p0.Fn1686
func Fn1686(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1694 github.com/goccy/llamawasm2go/p1.Fn1694
func Fn1694(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1734 github.com/goccy/llamawasm2go/p1.Fn1734
func Fn1734(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1735 github.com/goccy/llamawasm2go/p0.Fn1735
func Fn1735(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1761 github.com/goccy/llamawasm2go/p1.Fn1761
func Fn1761(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1832 github.com/goccy/llamawasm2go/p1.Fn1832
func Fn1832(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1845 github.com/goccy/llamawasm2go/p1.Fn1845
func Fn1845(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1930 github.com/goccy/llamawasm2go/p1.Fn1930
func Fn1930(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1987 github.com/goccy/llamawasm2go/p1.Fn1987
func Fn1987(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1990 github.com/goccy/llamawasm2go/p1.Fn1990
func Fn1990(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1998 github.com/goccy/llamawasm2go/p1.Fn1998
func Fn1998(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2014 github.com/goccy/llamawasm2go/p1.Fn2014
func Fn2014(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2034 github.com/goccy/llamawasm2go/p1.Fn2034
func Fn2034(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2046 github.com/goccy/llamawasm2go/p1.Fn2046
func Fn2046(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2103 github.com/goccy/llamawasm2go/p1.Fn2103
func Fn2103(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2104 github.com/goccy/llamawasm2go/p1.Fn2104
func Fn2104(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2208 github.com/goccy/llamawasm2go/p1.Fn2208
func Fn2208(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2223 github.com/goccy/llamawasm2go/p1.Fn2223
func Fn2223(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2232 github.com/goccy/llamawasm2go/p1.Fn2232
func Fn2232(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2236 github.com/goccy/llamawasm2go/p1.Fn2236
func Fn2236(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2238 github.com/goccy/llamawasm2go/p1.Fn2238
func Fn2238(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2250 github.com/goccy/llamawasm2go/p1.Fn2250
func Fn2250(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2298 github.com/goccy/llamawasm2go/p1.Fn2298
func Fn2298(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2301 github.com/goccy/llamawasm2go/p1.Fn2301
func Fn2301(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2326 github.com/goccy/llamawasm2go/p1.Fn2326
func Fn2326(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2352 github.com/goccy/llamawasm2go/p0.Fn2352
func Fn2352(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2357 github.com/goccy/llamawasm2go/p1.Fn2357
func Fn2357(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2358 github.com/goccy/llamawasm2go/p0.Fn2358
func Fn2358(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2364 github.com/goccy/llamawasm2go/p1.Fn2364
func Fn2364(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2396 github.com/goccy/llamawasm2go/p1.Fn2396
func Fn2396(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2404 github.com/goccy/llamawasm2go/p0.Fn2404
func Fn2404(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2416 github.com/goccy/llamawasm2go/p1.Fn2416
func Fn2416(m *base.Module)

//go:linkname Fn2452 github.com/goccy/llamawasm2go/p0.Fn2452
func Fn2452(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2462 github.com/goccy/llamawasm2go/p1.Fn2462
func Fn2462(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2485 github.com/goccy/llamawasm2go/p1.Fn2485
func Fn2485(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2496 github.com/goccy/llamawasm2go/p1.Fn2496
func Fn2496(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2497 github.com/goccy/llamawasm2go/p1.Fn2497
func Fn2497(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2566 github.com/goccy/llamawasm2go/p1.Fn2566
func Fn2566(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2572 github.com/goccy/llamawasm2go/p1.Fn2572
func Fn2572(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2583 github.com/goccy/llamawasm2go/p1.Fn2583
func Fn2583(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2594 github.com/goccy/llamawasm2go/p1.Fn2594
func Fn2594(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2634 github.com/goccy/llamawasm2go/p1.Fn2634
func Fn2634(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2641 github.com/goccy/llamawasm2go/p1.Fn2641
func Fn2641(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2642 github.com/goccy/llamawasm2go/p1.Fn2642
func Fn2642(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2696 github.com/goccy/llamawasm2go/p0.Fn2696
func Fn2696(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2728 github.com/goccy/llamawasm2go/p1.Fn2728
func Fn2728(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2730 github.com/goccy/llamawasm2go/p1.Fn2730
func Fn2730(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2757 github.com/goccy/llamawasm2go/p0.Fn2757
func Fn2757(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2860 github.com/goccy/llamawasm2go/p1.Fn2860
func Fn2860(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2867 github.com/goccy/llamawasm2go/p1.Fn2867
func Fn2867(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2937 github.com/goccy/llamawasm2go/p1.Fn2937
func Fn2937(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2946 github.com/goccy/llamawasm2go/p1.Fn2946
func Fn2946(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2961 github.com/goccy/llamawasm2go/p1.Fn2961
func Fn2961(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2986 github.com/goccy/llamawasm2go/p1.Fn2986
func Fn2986(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3013 github.com/goccy/llamawasm2go/p1.Fn3013
func Fn3013(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3025 github.com/goccy/llamawasm2go/p1.Fn3025
func Fn3025(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3040 github.com/goccy/llamawasm2go/p1.Fn3040
func Fn3040(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3046 github.com/goccy/llamawasm2go/p0.Fn3046
func Fn3046(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3050 github.com/goccy/llamawasm2go/p0.Fn3050
func Fn3050(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3055 github.com/goccy/llamawasm2go/p0.Fn3055
func Fn3055(m *base.Module, l0 int64) int64

//go:linkname Fn3056 github.com/goccy/llamawasm2go/p1.Fn3056
func Fn3056(m *base.Module, l0 int64)

//go:linkname Fn3058 github.com/goccy/llamawasm2go/p1.Fn3058
func Fn3058(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3068 github.com/goccy/llamawasm2go/p1.Fn3068
func Fn3068(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3070 github.com/goccy/llamawasm2go/p1.Fn3070
func Fn3070(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3071 github.com/goccy/llamawasm2go/p1.Fn3071
func Fn3071(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3072 github.com/goccy/llamawasm2go/p1.Fn3072
func Fn3072(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3104 github.com/goccy/llamawasm2go/p1.Fn3104
func Fn3104(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)
