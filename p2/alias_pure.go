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

//go:linkname Fn990 github.com/goccy/llamawasm2go/p0.Fn990
func Fn990(m *base.Module, l0 int64) int64

//go:linkname Fn1142 github.com/goccy/llamawasm2go/p0.Fn1142
func Fn1142(m *base.Module, l0 int64)

//go:linkname Fn1149 github.com/goccy/llamawasm2go/p0.Fn1149
func Fn1149(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1228 github.com/goccy/llamawasm2go/p1.Fn1228
func Fn1228(m *base.Module, l0 int64) int64

//go:linkname Fn1233 github.com/goccy/llamawasm2go/p0.Fn1233
func Fn1233(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1276 github.com/goccy/llamawasm2go/p1.Fn1276
func Fn1276(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1277 github.com/goccy/llamawasm2go/p1.Fn1277
func Fn1277(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1299 github.com/goccy/llamawasm2go/p1.Fn1299
func Fn1299(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1301 github.com/goccy/llamawasm2go/p1.Fn1301
func Fn1301(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1315 github.com/goccy/llamawasm2go/p1.Fn1315
func Fn1315(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1319 github.com/goccy/llamawasm2go/p0.Fn1319
func Fn1319(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1375 github.com/goccy/llamawasm2go/p1.Fn1375
func Fn1375(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p1.Fn1392
func Fn1392(m *base.Module, l0 int64)

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p1.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p1.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p1.Fn1406
func Fn1406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p1.Fn1408
func Fn1408(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1419 github.com/goccy/llamawasm2go/p1.Fn1419
func Fn1419(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1508 github.com/goccy/llamawasm2go/p0.Fn1508
func Fn1508(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1538 github.com/goccy/llamawasm2go/p0.Fn1538
func Fn1538(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1546 github.com/goccy/llamawasm2go/p1.Fn1546
func Fn1546(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1587 github.com/goccy/llamawasm2go/p0.Fn1587
func Fn1587(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1613 github.com/goccy/llamawasm2go/p1.Fn1613
func Fn1613(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1640 github.com/goccy/llamawasm2go/p1.Fn1640
func Fn1640(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1684 github.com/goccy/llamawasm2go/p1.Fn1684
func Fn1684(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1697 github.com/goccy/llamawasm2go/p1.Fn1697
func Fn1697(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1780 github.com/goccy/llamawasm2go/p1.Fn1780
func Fn1780(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1831 github.com/goccy/llamawasm2go/p1.Fn1831
func Fn1831(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1839 github.com/goccy/llamawasm2go/p1.Fn1839
func Fn1839(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn1855 github.com/goccy/llamawasm2go/p1.Fn1855
func Fn1855(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1875 github.com/goccy/llamawasm2go/p1.Fn1875
func Fn1875(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1887 github.com/goccy/llamawasm2go/p1.Fn1887
func Fn1887(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1944 github.com/goccy/llamawasm2go/p1.Fn1944
func Fn1944(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1945 github.com/goccy/llamawasm2go/p1.Fn1945
func Fn1945(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2042 github.com/goccy/llamawasm2go/p1.Fn2042
func Fn2042(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2050 github.com/goccy/llamawasm2go/p0.Fn2050
func Fn2050(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2066 github.com/goccy/llamawasm2go/p1.Fn2066
func Fn2066(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2070 github.com/goccy/llamawasm2go/p1.Fn2070
func Fn2070(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2072 github.com/goccy/llamawasm2go/p1.Fn2072
func Fn2072(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2084 github.com/goccy/llamawasm2go/p1.Fn2084
func Fn2084(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2131 github.com/goccy/llamawasm2go/p1.Fn2131
func Fn2131(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2134 github.com/goccy/llamawasm2go/p1.Fn2134
func Fn2134(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2159 github.com/goccy/llamawasm2go/p1.Fn2159
func Fn2159(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2185 github.com/goccy/llamawasm2go/p0.Fn2185
func Fn2185(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2190 github.com/goccy/llamawasm2go/p1.Fn2190
func Fn2190(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2191 github.com/goccy/llamawasm2go/p0.Fn2191
func Fn2191(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2197 github.com/goccy/llamawasm2go/p1.Fn2197
func Fn2197(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2226 github.com/goccy/llamawasm2go/p1.Fn2226
func Fn2226(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2234 github.com/goccy/llamawasm2go/p1.Fn2234
func Fn2234(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2245 github.com/goccy/llamawasm2go/p1.Fn2245
func Fn2245(m *base.Module)

//go:linkname Fn2281 github.com/goccy/llamawasm2go/p0.Fn2281
func Fn2281(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2291 github.com/goccy/llamawasm2go/p1.Fn2291
func Fn2291(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2314 github.com/goccy/llamawasm2go/p1.Fn2314
func Fn2314(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2325 github.com/goccy/llamawasm2go/p1.Fn2325
func Fn2325(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2326 github.com/goccy/llamawasm2go/p1.Fn2326
func Fn2326(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2395 github.com/goccy/llamawasm2go/p1.Fn2395
func Fn2395(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2401 github.com/goccy/llamawasm2go/p1.Fn2401
func Fn2401(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2412 github.com/goccy/llamawasm2go/p1.Fn2412
func Fn2412(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2423 github.com/goccy/llamawasm2go/p1.Fn2423
func Fn2423(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2463 github.com/goccy/llamawasm2go/p1.Fn2463
func Fn2463(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2470 github.com/goccy/llamawasm2go/p1.Fn2470
func Fn2470(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2471 github.com/goccy/llamawasm2go/p1.Fn2471
func Fn2471(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2525 github.com/goccy/llamawasm2go/p0.Fn2525
func Fn2525(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2557 github.com/goccy/llamawasm2go/p1.Fn2557
func Fn2557(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2559 github.com/goccy/llamawasm2go/p1.Fn2559
func Fn2559(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2586 github.com/goccy/llamawasm2go/p0.Fn2586
func Fn2586(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2689 github.com/goccy/llamawasm2go/p1.Fn2689
func Fn2689(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2696 github.com/goccy/llamawasm2go/p1.Fn2696
func Fn2696(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2777 github.com/goccy/llamawasm2go/p1.Fn2777
func Fn2777(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2820 github.com/goccy/llamawasm2go/p1.Fn2820
func Fn2820(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2832 github.com/goccy/llamawasm2go/p1.Fn2832
func Fn2832(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn2853 github.com/goccy/llamawasm2go/p0.Fn2853
func Fn2853(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2857 github.com/goccy/llamawasm2go/p0.Fn2857
func Fn2857(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2861 github.com/goccy/llamawasm2go/p0.Fn2861
func Fn2861(m *base.Module, l0 int64) int64

//go:linkname Fn2862 github.com/goccy/llamawasm2go/p1.Fn2862
func Fn2862(m *base.Module, l0 int64)

//go:linkname Fn2922 github.com/goccy/llamawasm2go/p1.Fn2922
func Fn2922(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn2935 github.com/goccy/llamawasm2go/p1.Fn2935
func Fn2935(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2936 github.com/goccy/llamawasm2go/p1.Fn2936
func Fn2936(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2937 github.com/goccy/llamawasm2go/p1.Fn2937
func Fn2937(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2938 github.com/goccy/llamawasm2go/p1.Fn2938
func Fn2938(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32
