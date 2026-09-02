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

//go:linkname Fn354 github.com/goccy/llamawasm2go/p0.Fn354
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

//go:linkname Fn505 github.com/goccy/llamawasm2go/p1.Fn505
func Fn505(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 float32, l6 float32, l7 float32) int64

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

//go:linkname Fn986 github.com/goccy/llamawasm2go/p0.Fn986
func Fn986(m *base.Module, l0 int64) int64

//go:linkname Fn1138 github.com/goccy/llamawasm2go/p0.Fn1138
func Fn1138(m *base.Module, l0 int64)

//go:linkname Fn1145 github.com/goccy/llamawasm2go/p0.Fn1145
func Fn1145(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1224 github.com/goccy/llamawasm2go/p1.Fn1224
func Fn1224(m *base.Module, l0 int64) int64

//go:linkname Fn1229 github.com/goccy/llamawasm2go/p0.Fn1229
func Fn1229(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1272 github.com/goccy/llamawasm2go/p1.Fn1272
func Fn1272(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1273 github.com/goccy/llamawasm2go/p1.Fn1273
func Fn1273(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1295 github.com/goccy/llamawasm2go/p1.Fn1295
func Fn1295(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1297 github.com/goccy/llamawasm2go/p1.Fn1297
func Fn1297(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1311 github.com/goccy/llamawasm2go/p1.Fn1311
func Fn1311(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1315 github.com/goccy/llamawasm2go/p0.Fn1315
func Fn1315(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1371 github.com/goccy/llamawasm2go/p1.Fn1371
func Fn1371(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1388 github.com/goccy/llamawasm2go/p1.Fn1388
func Fn1388(m *base.Module, l0 int64)

//go:linkname Fn1396 github.com/goccy/llamawasm2go/p1.Fn1396
func Fn1396(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p1.Fn1401
func Fn1401(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p1.Fn1402
func Fn1402(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1404 github.com/goccy/llamawasm2go/p1.Fn1404
func Fn1404(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1415 github.com/goccy/llamawasm2go/p1.Fn1415
func Fn1415(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1504 github.com/goccy/llamawasm2go/p0.Fn1504
func Fn1504(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1534 github.com/goccy/llamawasm2go/p0.Fn1534
func Fn1534(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p1.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p0.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1609 github.com/goccy/llamawasm2go/p1.Fn1609
func Fn1609(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1636 github.com/goccy/llamawasm2go/p1.Fn1636
func Fn1636(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1680 github.com/goccy/llamawasm2go/p1.Fn1680
func Fn1680(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1693 github.com/goccy/llamawasm2go/p1.Fn1693
func Fn1693(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1776 github.com/goccy/llamawasm2go/p1.Fn1776
func Fn1776(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1827 github.com/goccy/llamawasm2go/p1.Fn1827
func Fn1827(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1835 github.com/goccy/llamawasm2go/p1.Fn1835
func Fn1835(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn1851 github.com/goccy/llamawasm2go/p1.Fn1851
func Fn1851(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1871 github.com/goccy/llamawasm2go/p1.Fn1871
func Fn1871(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1883 github.com/goccy/llamawasm2go/p1.Fn1883
func Fn1883(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1890 github.com/goccy/llamawasm2go/p0.Fn1890
func Fn1890(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1940 github.com/goccy/llamawasm2go/p1.Fn1940
func Fn1940(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1941 github.com/goccy/llamawasm2go/p1.Fn1941
func Fn1941(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2038 github.com/goccy/llamawasm2go/p1.Fn2038
func Fn2038(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2046 github.com/goccy/llamawasm2go/p0.Fn2046
func Fn2046(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2062 github.com/goccy/llamawasm2go/p1.Fn2062
func Fn2062(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2066 github.com/goccy/llamawasm2go/p1.Fn2066
func Fn2066(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2068 github.com/goccy/llamawasm2go/p1.Fn2068
func Fn2068(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2080 github.com/goccy/llamawasm2go/p1.Fn2080
func Fn2080(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2127 github.com/goccy/llamawasm2go/p1.Fn2127
func Fn2127(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2130 github.com/goccy/llamawasm2go/p1.Fn2130
func Fn2130(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2155 github.com/goccy/llamawasm2go/p1.Fn2155
func Fn2155(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2181 github.com/goccy/llamawasm2go/p0.Fn2181
func Fn2181(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2186 github.com/goccy/llamawasm2go/p1.Fn2186
func Fn2186(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2187 github.com/goccy/llamawasm2go/p0.Fn2187
func Fn2187(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2193 github.com/goccy/llamawasm2go/p1.Fn2193
func Fn2193(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2222 github.com/goccy/llamawasm2go/p1.Fn2222
func Fn2222(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2230 github.com/goccy/llamawasm2go/p1.Fn2230
func Fn2230(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2241 github.com/goccy/llamawasm2go/p1.Fn2241
func Fn2241(m *base.Module)

//go:linkname Fn2277 github.com/goccy/llamawasm2go/p1.Fn2277
func Fn2277(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2287 github.com/goccy/llamawasm2go/p1.Fn2287
func Fn2287(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2310 github.com/goccy/llamawasm2go/p1.Fn2310
func Fn2310(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2321 github.com/goccy/llamawasm2go/p1.Fn2321
func Fn2321(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2322 github.com/goccy/llamawasm2go/p1.Fn2322
func Fn2322(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2391 github.com/goccy/llamawasm2go/p1.Fn2391
func Fn2391(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2397 github.com/goccy/llamawasm2go/p1.Fn2397
func Fn2397(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2408 github.com/goccy/llamawasm2go/p1.Fn2408
func Fn2408(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2419 github.com/goccy/llamawasm2go/p1.Fn2419
func Fn2419(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2459 github.com/goccy/llamawasm2go/p1.Fn2459
func Fn2459(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2466 github.com/goccy/llamawasm2go/p1.Fn2466
func Fn2466(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2467 github.com/goccy/llamawasm2go/p1.Fn2467
func Fn2467(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2521 github.com/goccy/llamawasm2go/p0.Fn2521
func Fn2521(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2553 github.com/goccy/llamawasm2go/p1.Fn2553
func Fn2553(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2555 github.com/goccy/llamawasm2go/p1.Fn2555
func Fn2555(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2582 github.com/goccy/llamawasm2go/p0.Fn2582
func Fn2582(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2685 github.com/goccy/llamawasm2go/p1.Fn2685
func Fn2685(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2692 github.com/goccy/llamawasm2go/p1.Fn2692
func Fn2692(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2773 github.com/goccy/llamawasm2go/p1.Fn2773
func Fn2773(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2816 github.com/goccy/llamawasm2go/p1.Fn2816
func Fn2816(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2828 github.com/goccy/llamawasm2go/p1.Fn2828
func Fn2828(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn2849 github.com/goccy/llamawasm2go/p0.Fn2849
func Fn2849(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2853 github.com/goccy/llamawasm2go/p0.Fn2853
func Fn2853(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2857 github.com/goccy/llamawasm2go/p0.Fn2857
func Fn2857(m *base.Module, l0 int64) int64

//go:linkname Fn2858 github.com/goccy/llamawasm2go/p1.Fn2858
func Fn2858(m *base.Module, l0 int64)

//go:linkname Fn2918 github.com/goccy/llamawasm2go/p1.Fn2918
func Fn2918(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn2931 github.com/goccy/llamawasm2go/p1.Fn2931
func Fn2931(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2932 github.com/goccy/llamawasm2go/p1.Fn2932
func Fn2932(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2933 github.com/goccy/llamawasm2go/p1.Fn2933
func Fn2933(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2934 github.com/goccy/llamawasm2go/p1.Fn2934
func Fn2934(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32
