//go:build !arm64 && (!amd64 || !amd64.v2)

package p2

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn23 github.com/goccy/llamawasm2go/p1.Fn23
func Fn23(m *base.Module, l0 int64)

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

//go:linkname Fn989 github.com/goccy/llamawasm2go/p0.Fn989
func Fn989(m *base.Module, l0 int64) int64

//go:linkname Fn1141 github.com/goccy/llamawasm2go/p0.Fn1141
func Fn1141(m *base.Module, l0 int64)

//go:linkname Fn1148 github.com/goccy/llamawasm2go/p0.Fn1148
func Fn1148(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1227 github.com/goccy/llamawasm2go/p1.Fn1227
func Fn1227(m *base.Module, l0 int64) int64

//go:linkname Fn1232 github.com/goccy/llamawasm2go/p0.Fn1232
func Fn1232(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1275 github.com/goccy/llamawasm2go/p1.Fn1275
func Fn1275(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1276 github.com/goccy/llamawasm2go/p1.Fn1276
func Fn1276(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1298 github.com/goccy/llamawasm2go/p1.Fn1298
func Fn1298(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1300 github.com/goccy/llamawasm2go/p1.Fn1300
func Fn1300(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1314 github.com/goccy/llamawasm2go/p1.Fn1314
func Fn1314(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1318 github.com/goccy/llamawasm2go/p0.Fn1318
func Fn1318(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1374 github.com/goccy/llamawasm2go/p1.Fn1374
func Fn1374(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1391 github.com/goccy/llamawasm2go/p1.Fn1391
func Fn1391(m *base.Module, l0 int64)

//go:linkname Fn1399 github.com/goccy/llamawasm2go/p1.Fn1399
func Fn1399(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1404 github.com/goccy/llamawasm2go/p1.Fn1404
func Fn1404(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p1.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p1.Fn1407
func Fn1407(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1418 github.com/goccy/llamawasm2go/p1.Fn1418
func Fn1418(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1507 github.com/goccy/llamawasm2go/p0.Fn1507
func Fn1507(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1537 github.com/goccy/llamawasm2go/p0.Fn1537
func Fn1537(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p1.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1586 github.com/goccy/llamawasm2go/p0.Fn1586
func Fn1586(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1612 github.com/goccy/llamawasm2go/p1.Fn1612
func Fn1612(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1639 github.com/goccy/llamawasm2go/p1.Fn1639
func Fn1639(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1683 github.com/goccy/llamawasm2go/p1.Fn1683
func Fn1683(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1696 github.com/goccy/llamawasm2go/p1.Fn1696
func Fn1696(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1779 github.com/goccy/llamawasm2go/p1.Fn1779
func Fn1779(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1830 github.com/goccy/llamawasm2go/p1.Fn1830
func Fn1830(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1838 github.com/goccy/llamawasm2go/p1.Fn1838
func Fn1838(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn1854 github.com/goccy/llamawasm2go/p1.Fn1854
func Fn1854(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1874 github.com/goccy/llamawasm2go/p1.Fn1874
func Fn1874(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1886 github.com/goccy/llamawasm2go/p1.Fn1886
func Fn1886(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1943 github.com/goccy/llamawasm2go/p1.Fn1943
func Fn1943(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn1944 github.com/goccy/llamawasm2go/p1.Fn1944
func Fn1944(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2041 github.com/goccy/llamawasm2go/p1.Fn2041
func Fn2041(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2049 github.com/goccy/llamawasm2go/p0.Fn2049
func Fn2049(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2065 github.com/goccy/llamawasm2go/p1.Fn2065
func Fn2065(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2069 github.com/goccy/llamawasm2go/p1.Fn2069
func Fn2069(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2071 github.com/goccy/llamawasm2go/p1.Fn2071
func Fn2071(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2083 github.com/goccy/llamawasm2go/p1.Fn2083
func Fn2083(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2130 github.com/goccy/llamawasm2go/p1.Fn2130
func Fn2130(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2133 github.com/goccy/llamawasm2go/p1.Fn2133
func Fn2133(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2158 github.com/goccy/llamawasm2go/p1.Fn2158
func Fn2158(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2184 github.com/goccy/llamawasm2go/p0.Fn2184
func Fn2184(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2189 github.com/goccy/llamawasm2go/p1.Fn2189
func Fn2189(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2190 github.com/goccy/llamawasm2go/p0.Fn2190
func Fn2190(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2196 github.com/goccy/llamawasm2go/p1.Fn2196
func Fn2196(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2225 github.com/goccy/llamawasm2go/p1.Fn2225
func Fn2225(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2233 github.com/goccy/llamawasm2go/p1.Fn2233
func Fn2233(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2244 github.com/goccy/llamawasm2go/p1.Fn2244
func Fn2244(m *base.Module)

//go:linkname Fn2280 github.com/goccy/llamawasm2go/p0.Fn2280
func Fn2280(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2290 github.com/goccy/llamawasm2go/p1.Fn2290
func Fn2290(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2313 github.com/goccy/llamawasm2go/p1.Fn2313
func Fn2313(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2324 github.com/goccy/llamawasm2go/p1.Fn2324
func Fn2324(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2325 github.com/goccy/llamawasm2go/p1.Fn2325
func Fn2325(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2394 github.com/goccy/llamawasm2go/p1.Fn2394
func Fn2394(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2400 github.com/goccy/llamawasm2go/p1.Fn2400
func Fn2400(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2411 github.com/goccy/llamawasm2go/p1.Fn2411
func Fn2411(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2422 github.com/goccy/llamawasm2go/p1.Fn2422
func Fn2422(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2462 github.com/goccy/llamawasm2go/p1.Fn2462
func Fn2462(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2469 github.com/goccy/llamawasm2go/p1.Fn2469
func Fn2469(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2470 github.com/goccy/llamawasm2go/p1.Fn2470
func Fn2470(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2524 github.com/goccy/llamawasm2go/p0.Fn2524
func Fn2524(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2556 github.com/goccy/llamawasm2go/p1.Fn2556
func Fn2556(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2558 github.com/goccy/llamawasm2go/p1.Fn2558
func Fn2558(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2585 github.com/goccy/llamawasm2go/p0.Fn2585
func Fn2585(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2688 github.com/goccy/llamawasm2go/p1.Fn2688
func Fn2688(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2695 github.com/goccy/llamawasm2go/p1.Fn2695
func Fn2695(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2776 github.com/goccy/llamawasm2go/p1.Fn2776
func Fn2776(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn2819 github.com/goccy/llamawasm2go/p1.Fn2819
func Fn2819(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2831 github.com/goccy/llamawasm2go/p1.Fn2831
func Fn2831(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn2852 github.com/goccy/llamawasm2go/p0.Fn2852
func Fn2852(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2856 github.com/goccy/llamawasm2go/p0.Fn2856
func Fn2856(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2860 github.com/goccy/llamawasm2go/p0.Fn2860
func Fn2860(m *base.Module, l0 int64) int64

//go:linkname Fn2861 github.com/goccy/llamawasm2go/p1.Fn2861
func Fn2861(m *base.Module, l0 int64)

//go:linkname Fn2921 github.com/goccy/llamawasm2go/p1.Fn2921
func Fn2921(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn2934 github.com/goccy/llamawasm2go/p1.Fn2934
func Fn2934(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2935 github.com/goccy/llamawasm2go/p1.Fn2935
func Fn2935(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2936 github.com/goccy/llamawasm2go/p1.Fn2936
func Fn2936(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn2937 github.com/goccy/llamawasm2go/p1.Fn2937
func Fn2937(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32
