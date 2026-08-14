//go:build !arm64 && (!amd64 || !amd64.v2)

package p2

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn46 github.com/goccy/llamawasm2go/p1.Fn46
func Fn46(m *base.Module, l0 int64)

//go:linkname Fn62 github.com/goccy/llamawasm2go/p1.Fn62
func Fn62(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn160 github.com/goccy/llamawasm2go/p1.Fn160
func Fn160(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn349 github.com/goccy/llamawasm2go/p1.Fn349
func Fn349(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn351 github.com/goccy/llamawasm2go/p1.Fn351
func Fn351(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn354 github.com/goccy/llamawasm2go/p1.Fn354
func Fn354(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn359 github.com/goccy/llamawasm2go/p1.Fn359
func Fn359(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn367 github.com/goccy/llamawasm2go/p1.Fn367
func Fn367(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn391 github.com/goccy/llamawasm2go/p1.Fn391
func Fn391(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn398 github.com/goccy/llamawasm2go/p1.Fn398
func Fn398(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn421 github.com/goccy/llamawasm2go/p1.Fn421
func Fn421(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn425 github.com/goccy/llamawasm2go/p1.Fn425
func Fn425(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn443 github.com/goccy/llamawasm2go/p1.Fn443
func Fn443(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn447 github.com/goccy/llamawasm2go/p0.Fn447
func Fn447(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn457 github.com/goccy/llamawasm2go/p1.Fn457
func Fn457(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn481 github.com/goccy/llamawasm2go/p1.Fn481
func Fn481(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn540 github.com/goccy/llamawasm2go/p0.Fn540
func Fn540(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn576 github.com/goccy/llamawasm2go/p1.Fn576
func Fn576(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn577 github.com/goccy/llamawasm2go/p1.Fn577
func Fn577(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn591 github.com/goccy/llamawasm2go/p1.Fn591
func Fn591(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn593 github.com/goccy/llamawasm2go/p1.Fn593
func Fn593(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn597 github.com/goccy/llamawasm2go/p1.Fn597
func Fn597(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn598 github.com/goccy/llamawasm2go/p1.Fn598
func Fn598(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn615 github.com/goccy/llamawasm2go/p0.Fn615
func Fn615(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn617 github.com/goccy/llamawasm2go/p0.Fn617
func Fn617(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn625 github.com/goccy/llamawasm2go/p0.Fn625
func Fn625(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn627 github.com/goccy/llamawasm2go/p1.Fn627
func Fn627(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn669 github.com/goccy/llamawasm2go/p1.Fn669
func Fn669(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn759 github.com/goccy/llamawasm2go/p1.Fn759
func Fn759(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn761 github.com/goccy/llamawasm2go/p1.Fn761
func Fn761(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn767 github.com/goccy/llamawasm2go/p1.Fn767
func Fn767(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn802 github.com/goccy/llamawasm2go/p1.Fn802
func Fn802(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn806 github.com/goccy/llamawasm2go/p1.Fn806
func Fn806(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn816 github.com/goccy/llamawasm2go/p1.Fn816
func Fn816(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn821 github.com/goccy/llamawasm2go/p1.Fn821
func Fn821(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn830 github.com/goccy/llamawasm2go/p1.Fn830
func Fn830(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn836 github.com/goccy/llamawasm2go/p1.Fn836
func Fn836(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn841 github.com/goccy/llamawasm2go/p1.Fn841
func Fn841(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn850 github.com/goccy/llamawasm2go/p1.Fn850
func Fn850(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn853 github.com/goccy/llamawasm2go/p1.Fn853
func Fn853(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn855 github.com/goccy/llamawasm2go/p1.Fn855
func Fn855(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn864 github.com/goccy/llamawasm2go/p1.Fn864
func Fn864(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn883 github.com/goccy/llamawasm2go/p1.Fn883
func Fn883(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn887 github.com/goccy/llamawasm2go/p0.Fn887
func Fn887(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn891 github.com/goccy/llamawasm2go/p1.Fn891
func Fn891(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn895 github.com/goccy/llamawasm2go/p1.Fn895
func Fn895(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn903 github.com/goccy/llamawasm2go/p0.Fn903
func Fn903(m *base.Module, l0 int64) int64

//go:linkname Fn1071 github.com/goccy/llamawasm2go/p0.Fn1071
func Fn1071(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1140 github.com/goccy/llamawasm2go/p1.Fn1140
func Fn1140(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1145 github.com/goccy/llamawasm2go/p1.Fn1145
func Fn1145(m *base.Module)

//go:linkname Fn1146 github.com/goccy/llamawasm2go/p0.Fn1146
func Fn1146(m *base.Module, l0 int64)

//go:linkname Fn1251 github.com/goccy/llamawasm2go/p1.Fn1251
func Fn1251(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1253 github.com/goccy/llamawasm2go/p0.Fn1253
func Fn1253(m *base.Module, l0 int64)

//go:linkname Fn1260 github.com/goccy/llamawasm2go/p0.Fn1260
func Fn1260(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1328 github.com/goccy/llamawasm2go/p1.Fn1328
func Fn1328(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1341 github.com/goccy/llamawasm2go/p1.Fn1341
func Fn1341(m *base.Module, l0 int64) int64

//go:linkname Fn1382 github.com/goccy/llamawasm2go/p1.Fn1382
func Fn1382(m *base.Module, l0 int64)

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p1.Fn1392
func Fn1392(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1393 github.com/goccy/llamawasm2go/p1.Fn1393
func Fn1393(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1419 github.com/goccy/llamawasm2go/p1.Fn1419
func Fn1419(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1422 github.com/goccy/llamawasm2go/p1.Fn1422
func Fn1422(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1431 github.com/goccy/llamawasm2go/p1.Fn1431
func Fn1431(m *base.Module, l0 int64)

//go:linkname Fn1440 github.com/goccy/llamawasm2go/p0.Fn1440
func Fn1440(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1486 github.com/goccy/llamawasm2go/p1.Fn1486
func Fn1486(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1487 github.com/goccy/llamawasm2go/p1.Fn1487
func Fn1487(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1497 github.com/goccy/llamawasm2go/p1.Fn1497
func Fn1497(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1514 github.com/goccy/llamawasm2go/p1.Fn1514
func Fn1514(m *base.Module, l0 int64)

//go:linkname Fn1522 github.com/goccy/llamawasm2go/p1.Fn1522
func Fn1522(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1528 github.com/goccy/llamawasm2go/p1.Fn1528
func Fn1528(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1529 github.com/goccy/llamawasm2go/p1.Fn1529
func Fn1529(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1531 github.com/goccy/llamawasm2go/p1.Fn1531
func Fn1531(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p1.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1640 github.com/goccy/llamawasm2go/p0.Fn1640
func Fn1640(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1671 github.com/goccy/llamawasm2go/p0.Fn1671
func Fn1671(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1679 github.com/goccy/llamawasm2go/p1.Fn1679
func Fn1679(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1719 github.com/goccy/llamawasm2go/p1.Fn1719
func Fn1719(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1720 github.com/goccy/llamawasm2go/p0.Fn1720
func Fn1720(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1746 github.com/goccy/llamawasm2go/p1.Fn1746
func Fn1746(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1817 github.com/goccy/llamawasm2go/p1.Fn1817
func Fn1817(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1830 github.com/goccy/llamawasm2go/p1.Fn1830
func Fn1830(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1915 github.com/goccy/llamawasm2go/p1.Fn1915
func Fn1915(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1961 github.com/goccy/llamawasm2go/p1.Fn1961
func Fn1961(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1972 github.com/goccy/llamawasm2go/p1.Fn1972
func Fn1972(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1974 github.com/goccy/llamawasm2go/p1.Fn1974
func Fn1974(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1975 github.com/goccy/llamawasm2go/p1.Fn1975
func Fn1975(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1983 github.com/goccy/llamawasm2go/p1.Fn1983
func Fn1983(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn1998 github.com/goccy/llamawasm2go/p1.Fn1998
func Fn1998(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2018 github.com/goccy/llamawasm2go/p1.Fn2018
func Fn2018(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2030 github.com/goccy/llamawasm2go/p1.Fn2030
func Fn2030(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2087 github.com/goccy/llamawasm2go/p1.Fn2087
func Fn2087(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2088 github.com/goccy/llamawasm2go/p1.Fn2088
func Fn2088(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2192 github.com/goccy/llamawasm2go/p1.Fn2192
func Fn2192(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2207 github.com/goccy/llamawasm2go/p1.Fn2207
func Fn2207(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2216 github.com/goccy/llamawasm2go/p1.Fn2216
func Fn2216(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2220 github.com/goccy/llamawasm2go/p1.Fn2220
func Fn2220(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2222 github.com/goccy/llamawasm2go/p1.Fn2222
func Fn2222(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2234 github.com/goccy/llamawasm2go/p1.Fn2234
func Fn2234(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2282 github.com/goccy/llamawasm2go/p1.Fn2282
func Fn2282(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2285 github.com/goccy/llamawasm2go/p1.Fn2285
func Fn2285(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2310 github.com/goccy/llamawasm2go/p1.Fn2310
func Fn2310(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2336 github.com/goccy/llamawasm2go/p0.Fn2336
func Fn2336(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2341 github.com/goccy/llamawasm2go/p1.Fn2341
func Fn2341(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2342 github.com/goccy/llamawasm2go/p0.Fn2342
func Fn2342(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2348 github.com/goccy/llamawasm2go/p1.Fn2348
func Fn2348(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2380 github.com/goccy/llamawasm2go/p1.Fn2380
func Fn2380(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2388 github.com/goccy/llamawasm2go/p0.Fn2388
func Fn2388(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2400 github.com/goccy/llamawasm2go/p1.Fn2400
func Fn2400(m *base.Module)

//go:linkname Fn2434 github.com/goccy/llamawasm2go/p0.Fn2434
func Fn2434(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2444 github.com/goccy/llamawasm2go/p1.Fn2444
func Fn2444(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2467 github.com/goccy/llamawasm2go/p1.Fn2467
func Fn2467(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2478 github.com/goccy/llamawasm2go/p1.Fn2478
func Fn2478(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2479 github.com/goccy/llamawasm2go/p1.Fn2479
func Fn2479(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2548 github.com/goccy/llamawasm2go/p1.Fn2548
func Fn2548(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2554 github.com/goccy/llamawasm2go/p1.Fn2554
func Fn2554(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2565 github.com/goccy/llamawasm2go/p1.Fn2565
func Fn2565(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2576 github.com/goccy/llamawasm2go/p1.Fn2576
func Fn2576(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2616 github.com/goccy/llamawasm2go/p1.Fn2616
func Fn2616(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2623 github.com/goccy/llamawasm2go/p1.Fn2623
func Fn2623(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2624 github.com/goccy/llamawasm2go/p1.Fn2624
func Fn2624(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2678 github.com/goccy/llamawasm2go/p0.Fn2678
func Fn2678(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2710 github.com/goccy/llamawasm2go/p1.Fn2710
func Fn2710(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2712 github.com/goccy/llamawasm2go/p1.Fn2712
func Fn2712(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2739 github.com/goccy/llamawasm2go/p0.Fn2739
func Fn2739(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2842 github.com/goccy/llamawasm2go/p1.Fn2842
func Fn2842(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2849 github.com/goccy/llamawasm2go/p1.Fn2849
func Fn2849(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2915 github.com/goccy/llamawasm2go/p1.Fn2915
func Fn2915(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2921 github.com/goccy/llamawasm2go/p1.Fn2921
func Fn2921(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2939 github.com/goccy/llamawasm2go/p1.Fn2939
func Fn2939(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2964 github.com/goccy/llamawasm2go/p1.Fn2964
func Fn2964(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2990 github.com/goccy/llamawasm2go/p1.Fn2990
func Fn2990(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3002 github.com/goccy/llamawasm2go/p1.Fn3002
func Fn3002(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3020 github.com/goccy/llamawasm2go/p0.Fn3020
func Fn3020(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3024 github.com/goccy/llamawasm2go/p0.Fn3024
func Fn3024(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3028 github.com/goccy/llamawasm2go/p0.Fn3028
func Fn3028(m *base.Module, l0 int64) int64

//go:linkname Fn3029 github.com/goccy/llamawasm2go/p1.Fn3029
func Fn3029(m *base.Module, l0 int64)

//go:linkname Fn3031 github.com/goccy/llamawasm2go/p1.Fn3031
func Fn3031(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3041 github.com/goccy/llamawasm2go/p1.Fn3041
func Fn3041(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3043 github.com/goccy/llamawasm2go/p1.Fn3043
func Fn3043(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3044 github.com/goccy/llamawasm2go/p0.Fn3044
func Fn3044(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3045 github.com/goccy/llamawasm2go/p1.Fn3045
func Fn3045(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3076 github.com/goccy/llamawasm2go/p1.Fn3076
func Fn3076(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)
