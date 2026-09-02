//go:build !arm64 && (!amd64 || !amd64.v2)

package p2

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn48 github.com/goccy/llamawasm2go/p1.Fn48
func Fn48(m *base.Module, l0 int64)

//go:linkname Fn63 github.com/goccy/llamawasm2go/p1.Fn63
func Fn63(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn352 github.com/goccy/llamawasm2go/p1.Fn352
func Fn352(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn354 github.com/goccy/llamawasm2go/p1.Fn354
func Fn354(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn358 github.com/goccy/llamawasm2go/p1.Fn358
func Fn358(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn362 github.com/goccy/llamawasm2go/p1.Fn362
func Fn362(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn367 github.com/goccy/llamawasm2go/p1.Fn367
func Fn367(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn370 github.com/goccy/llamawasm2go/p1.Fn370
func Fn370(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn390 github.com/goccy/llamawasm2go/p1.Fn390
func Fn390(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn391 github.com/goccy/llamawasm2go/p0.Fn391
func Fn391(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn396 github.com/goccy/llamawasm2go/p1.Fn396
func Fn396(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn397 github.com/goccy/llamawasm2go/p1.Fn397
func Fn397(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn398 github.com/goccy/llamawasm2go/p1.Fn398
func Fn398(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32)

//go:linkname Fn399 github.com/goccy/llamawasm2go/p1.Fn399
func Fn399(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn429 github.com/goccy/llamawasm2go/p1.Fn429
func Fn429(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn508 github.com/goccy/llamawasm2go/p1.Fn508
func Fn508(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn510 github.com/goccy/llamawasm2go/p1.Fn510
func Fn510(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn515 github.com/goccy/llamawasm2go/p1.Fn515
func Fn515(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn523 github.com/goccy/llamawasm2go/p1.Fn523
func Fn523(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn526 github.com/goccy/llamawasm2go/p1.Fn526
func Fn526(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn576 github.com/goccy/llamawasm2go/p0.Fn576
func Fn576(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn616 github.com/goccy/llamawasm2go/p1.Fn616
func Fn616(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn620 github.com/goccy/llamawasm2go/p0.Fn620
func Fn620(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn630 github.com/goccy/llamawasm2go/p1.Fn630
func Fn630(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn654 github.com/goccy/llamawasm2go/p1.Fn654
func Fn654(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn670 github.com/goccy/llamawasm2go/p1.Fn670
func Fn670(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn700 github.com/goccy/llamawasm2go/p1.Fn700
func Fn700(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn701 github.com/goccy/llamawasm2go/p1.Fn701
func Fn701(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn715 github.com/goccy/llamawasm2go/p1.Fn715
func Fn715(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn717 github.com/goccy/llamawasm2go/p1.Fn717
func Fn717(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn721 github.com/goccy/llamawasm2go/p1.Fn721
func Fn721(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn722 github.com/goccy/llamawasm2go/p1.Fn722
func Fn722(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn739 github.com/goccy/llamawasm2go/p1.Fn739
func Fn739(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn741 github.com/goccy/llamawasm2go/p0.Fn741
func Fn741(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn749 github.com/goccy/llamawasm2go/p0.Fn749
func Fn749(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn751 github.com/goccy/llamawasm2go/p1.Fn751
func Fn751(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn756 github.com/goccy/llamawasm2go/p1.Fn756
func Fn756(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn761 github.com/goccy/llamawasm2go/p1.Fn761
func Fn761(m *base.Module)

//go:linkname Fn762 github.com/goccy/llamawasm2go/p0.Fn762
func Fn762(m *base.Module, l0 int64)

//go:linkname Fn817 github.com/goccy/llamawasm2go/p1.Fn817
func Fn817(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int32

//go:linkname Fn818 github.com/goccy/llamawasm2go/p1.Fn818
func Fn818(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int32

//go:linkname Fn819 github.com/goccy/llamawasm2go/p1.Fn819
func Fn819(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int32

//go:linkname Fn820 github.com/goccy/llamawasm2go/p1.Fn820
func Fn820(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int32

//go:linkname Fn992 github.com/goccy/llamawasm2go/p0.Fn992
func Fn992(m *base.Module, l0 int64) int64

//go:linkname Fn1145 github.com/goccy/llamawasm2go/p0.Fn1145
func Fn1145(m *base.Module, l0 int64)

//go:linkname Fn1152 github.com/goccy/llamawasm2go/p0.Fn1152
func Fn1152(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1231 github.com/goccy/llamawasm2go/p1.Fn1231
func Fn1231(m *base.Module, l0 int64) int64

//go:linkname Fn1236 github.com/goccy/llamawasm2go/p0.Fn1236
func Fn1236(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1279 github.com/goccy/llamawasm2go/p1.Fn1279
func Fn1279(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1280 github.com/goccy/llamawasm2go/p1.Fn1280
func Fn1280(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1302 github.com/goccy/llamawasm2go/p1.Fn1302
func Fn1302(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1304 github.com/goccy/llamawasm2go/p1.Fn1304
func Fn1304(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1318 github.com/goccy/llamawasm2go/p1.Fn1318
func Fn1318(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1322 github.com/goccy/llamawasm2go/p0.Fn1322
func Fn1322(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1378 github.com/goccy/llamawasm2go/p1.Fn1378
func Fn1378(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p1.Fn1395
func Fn1395(m *base.Module, l0 int64)

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p1.Fn1403
func Fn1403(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1409 github.com/goccy/llamawasm2go/p1.Fn1409
func Fn1409(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1411 github.com/goccy/llamawasm2go/p1.Fn1411
func Fn1411(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1422 github.com/goccy/llamawasm2go/p1.Fn1422
func Fn1422(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1511 github.com/goccy/llamawasm2go/p0.Fn1511
func Fn1511(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p0.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1549 github.com/goccy/llamawasm2go/p1.Fn1549
func Fn1549(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1590 github.com/goccy/llamawasm2go/p0.Fn1590
func Fn1590(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1616 github.com/goccy/llamawasm2go/p1.Fn1616
func Fn1616(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1645 github.com/goccy/llamawasm2go/p1.Fn1645
func Fn1645(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1689 github.com/goccy/llamawasm2go/p1.Fn1689
func Fn1689(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1702 github.com/goccy/llamawasm2go/p1.Fn1702
func Fn1702(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1789 github.com/goccy/llamawasm2go/p1.Fn1789
func Fn1789(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1841 github.com/goccy/llamawasm2go/p1.Fn1841
func Fn1841(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1849 github.com/goccy/llamawasm2go/p1.Fn1849
func Fn1849(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn1865 github.com/goccy/llamawasm2go/p1.Fn1865
func Fn1865(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1889 github.com/goccy/llamawasm2go/p1.Fn1889
func Fn1889(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1904 github.com/goccy/llamawasm2go/p1.Fn1904
func Fn1904(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1969 github.com/goccy/llamawasm2go/p1.Fn1969
func Fn1969(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1970 github.com/goccy/llamawasm2go/p1.Fn1970
func Fn1970(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2067 github.com/goccy/llamawasm2go/p1.Fn2067
func Fn2067(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2075 github.com/goccy/llamawasm2go/p0.Fn2075
func Fn2075(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2091 github.com/goccy/llamawasm2go/p1.Fn2091
func Fn2091(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2095 github.com/goccy/llamawasm2go/p1.Fn2095
func Fn2095(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2097 github.com/goccy/llamawasm2go/p1.Fn2097
func Fn2097(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2109 github.com/goccy/llamawasm2go/p1.Fn2109
func Fn2109(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2156 github.com/goccy/llamawasm2go/p1.Fn2156
func Fn2156(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2159 github.com/goccy/llamawasm2go/p1.Fn2159
func Fn2159(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2184 github.com/goccy/llamawasm2go/p1.Fn2184
func Fn2184(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2210 github.com/goccy/llamawasm2go/p0.Fn2210
func Fn2210(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2215 github.com/goccy/llamawasm2go/p1.Fn2215
func Fn2215(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2216 github.com/goccy/llamawasm2go/p0.Fn2216
func Fn2216(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2222 github.com/goccy/llamawasm2go/p1.Fn2222
func Fn2222(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2259 github.com/goccy/llamawasm2go/p1.Fn2259
func Fn2259(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2270 github.com/goccy/llamawasm2go/p1.Fn2270
func Fn2270(m *base.Module)

//go:linkname Fn2306 github.com/goccy/llamawasm2go/p1.Fn2306
func Fn2306(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2316 github.com/goccy/llamawasm2go/p1.Fn2316
func Fn2316(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2339 github.com/goccy/llamawasm2go/p1.Fn2339
func Fn2339(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2350 github.com/goccy/llamawasm2go/p1.Fn2350
func Fn2350(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2351 github.com/goccy/llamawasm2go/p1.Fn2351
func Fn2351(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2420 github.com/goccy/llamawasm2go/p1.Fn2420
func Fn2420(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2426 github.com/goccy/llamawasm2go/p1.Fn2426
func Fn2426(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2437 github.com/goccy/llamawasm2go/p1.Fn2437
func Fn2437(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2448 github.com/goccy/llamawasm2go/p1.Fn2448
func Fn2448(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2488 github.com/goccy/llamawasm2go/p1.Fn2488
func Fn2488(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2495 github.com/goccy/llamawasm2go/p1.Fn2495
func Fn2495(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2496 github.com/goccy/llamawasm2go/p1.Fn2496
func Fn2496(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2550 github.com/goccy/llamawasm2go/p0.Fn2550
func Fn2550(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2582 github.com/goccy/llamawasm2go/p1.Fn2582
func Fn2582(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2584 github.com/goccy/llamawasm2go/p1.Fn2584
func Fn2584(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2611 github.com/goccy/llamawasm2go/p0.Fn2611
func Fn2611(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2714 github.com/goccy/llamawasm2go/p1.Fn2714
func Fn2714(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2721 github.com/goccy/llamawasm2go/p1.Fn2721
func Fn2721(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2802 github.com/goccy/llamawasm2go/p1.Fn2802
func Fn2802(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2845 github.com/goccy/llamawasm2go/p1.Fn2845
func Fn2845(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2857 github.com/goccy/llamawasm2go/p1.Fn2857
func Fn2857(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn2878 github.com/goccy/llamawasm2go/p0.Fn2878
func Fn2878(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2882 github.com/goccy/llamawasm2go/p0.Fn2882
func Fn2882(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2886 github.com/goccy/llamawasm2go/p0.Fn2886
func Fn2886(m *base.Module, l0 int64) int64

//go:linkname Fn2887 github.com/goccy/llamawasm2go/p1.Fn2887
func Fn2887(m *base.Module, l0 int64)
