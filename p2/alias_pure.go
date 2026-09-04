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

//go:linkname Fn739 github.com/goccy/llamawasm2go/p0.Fn739
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

//go:linkname Fn994 github.com/goccy/llamawasm2go/p0.Fn994
func Fn994(m *base.Module, l0 int64) int64

//go:linkname Fn1146 github.com/goccy/llamawasm2go/p0.Fn1146
func Fn1146(m *base.Module, l0 int64)

//go:linkname Fn1153 github.com/goccy/llamawasm2go/p0.Fn1153
func Fn1153(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1232 github.com/goccy/llamawasm2go/p1.Fn1232
func Fn1232(m *base.Module, l0 int64) int64

//go:linkname Fn1237 github.com/goccy/llamawasm2go/p0.Fn1237
func Fn1237(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1280 github.com/goccy/llamawasm2go/p1.Fn1280
func Fn1280(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1281 github.com/goccy/llamawasm2go/p1.Fn1281
func Fn1281(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1303 github.com/goccy/llamawasm2go/p1.Fn1303
func Fn1303(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1305 github.com/goccy/llamawasm2go/p1.Fn1305
func Fn1305(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1319 github.com/goccy/llamawasm2go/p1.Fn1319
func Fn1319(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p0.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1379 github.com/goccy/llamawasm2go/p1.Fn1379
func Fn1379(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1396 github.com/goccy/llamawasm2go/p1.Fn1396
func Fn1396(m *base.Module, l0 int64)

//go:linkname Fn1404 github.com/goccy/llamawasm2go/p1.Fn1404
func Fn1404(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1410 github.com/goccy/llamawasm2go/p1.Fn1410
func Fn1410(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1412 github.com/goccy/llamawasm2go/p1.Fn1412
func Fn1412(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1423 github.com/goccy/llamawasm2go/p1.Fn1423
func Fn1423(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1512 github.com/goccy/llamawasm2go/p0.Fn1512
func Fn1512(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p0.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1550 github.com/goccy/llamawasm2go/p1.Fn1550
func Fn1550(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1591 github.com/goccy/llamawasm2go/p0.Fn1591
func Fn1591(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1617 github.com/goccy/llamawasm2go/p1.Fn1617
func Fn1617(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1644 github.com/goccy/llamawasm2go/p1.Fn1644
func Fn1644(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1688 github.com/goccy/llamawasm2go/p1.Fn1688
func Fn1688(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1701 github.com/goccy/llamawasm2go/p1.Fn1701
func Fn1701(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1784 github.com/goccy/llamawasm2go/p1.Fn1784
func Fn1784(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1835 github.com/goccy/llamawasm2go/p1.Fn1835
func Fn1835(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1843 github.com/goccy/llamawasm2go/p1.Fn1843
func Fn1843(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn1859 github.com/goccy/llamawasm2go/p1.Fn1859
func Fn1859(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1879 github.com/goccy/llamawasm2go/p1.Fn1879
func Fn1879(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1891 github.com/goccy/llamawasm2go/p1.Fn1891
func Fn1891(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1948 github.com/goccy/llamawasm2go/p1.Fn1948
func Fn1948(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1949 github.com/goccy/llamawasm2go/p1.Fn1949
func Fn1949(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2046 github.com/goccy/llamawasm2go/p1.Fn2046
func Fn2046(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2054 github.com/goccy/llamawasm2go/p0.Fn2054
func Fn2054(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2070 github.com/goccy/llamawasm2go/p1.Fn2070
func Fn2070(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2074 github.com/goccy/llamawasm2go/p1.Fn2074
func Fn2074(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2076 github.com/goccy/llamawasm2go/p1.Fn2076
func Fn2076(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2088 github.com/goccy/llamawasm2go/p1.Fn2088
func Fn2088(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2135 github.com/goccy/llamawasm2go/p1.Fn2135
func Fn2135(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2138 github.com/goccy/llamawasm2go/p1.Fn2138
func Fn2138(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2163 github.com/goccy/llamawasm2go/p1.Fn2163
func Fn2163(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2189 github.com/goccy/llamawasm2go/p0.Fn2189
func Fn2189(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2194 github.com/goccy/llamawasm2go/p1.Fn2194
func Fn2194(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2195 github.com/goccy/llamawasm2go/p0.Fn2195
func Fn2195(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2201 github.com/goccy/llamawasm2go/p1.Fn2201
func Fn2201(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2238 github.com/goccy/llamawasm2go/p1.Fn2238
func Fn2238(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p1.Fn2249
func Fn2249(m *base.Module)

//go:linkname Fn2285 github.com/goccy/llamawasm2go/p0.Fn2285
func Fn2285(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2295 github.com/goccy/llamawasm2go/p1.Fn2295
func Fn2295(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2318 github.com/goccy/llamawasm2go/p1.Fn2318
func Fn2318(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2329 github.com/goccy/llamawasm2go/p1.Fn2329
func Fn2329(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2330 github.com/goccy/llamawasm2go/p1.Fn2330
func Fn2330(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2399 github.com/goccy/llamawasm2go/p1.Fn2399
func Fn2399(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2405 github.com/goccy/llamawasm2go/p1.Fn2405
func Fn2405(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2416 github.com/goccy/llamawasm2go/p1.Fn2416
func Fn2416(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2427 github.com/goccy/llamawasm2go/p1.Fn2427
func Fn2427(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2467 github.com/goccy/llamawasm2go/p1.Fn2467
func Fn2467(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2474 github.com/goccy/llamawasm2go/p1.Fn2474
func Fn2474(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2475 github.com/goccy/llamawasm2go/p1.Fn2475
func Fn2475(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2529 github.com/goccy/llamawasm2go/p0.Fn2529
func Fn2529(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2561 github.com/goccy/llamawasm2go/p1.Fn2561
func Fn2561(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2563 github.com/goccy/llamawasm2go/p1.Fn2563
func Fn2563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2590 github.com/goccy/llamawasm2go/p0.Fn2590
func Fn2590(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2693 github.com/goccy/llamawasm2go/p1.Fn2693
func Fn2693(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2700 github.com/goccy/llamawasm2go/p1.Fn2700
func Fn2700(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2781 github.com/goccy/llamawasm2go/p1.Fn2781
func Fn2781(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2824 github.com/goccy/llamawasm2go/p1.Fn2824
func Fn2824(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2836 github.com/goccy/llamawasm2go/p1.Fn2836
func Fn2836(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn2857 github.com/goccy/llamawasm2go/p0.Fn2857
func Fn2857(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2861 github.com/goccy/llamawasm2go/p0.Fn2861
func Fn2861(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2865 github.com/goccy/llamawasm2go/p0.Fn2865
func Fn2865(m *base.Module, l0 int64) int64

//go:linkname Fn2866 github.com/goccy/llamawasm2go/p1.Fn2866
func Fn2866(m *base.Module, l0 int64)

//go:linkname Fn2926 github.com/goccy/llamawasm2go/p1.Fn2926
func Fn2926(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn2939 github.com/goccy/llamawasm2go/p1.Fn2939
func Fn2939(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2940 github.com/goccy/llamawasm2go/p1.Fn2940
func Fn2940(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2941 github.com/goccy/llamawasm2go/p1.Fn2941
func Fn2941(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2942 github.com/goccy/llamawasm2go/p1.Fn2942
func Fn2942(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32
