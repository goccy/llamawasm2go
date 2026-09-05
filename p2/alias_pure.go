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

//go:linkname Fn358 github.com/goccy/llamawasm2go/p1.Fn358
func Fn358(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn362 github.com/goccy/llamawasm2go/p1.Fn362
func Fn362(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn363 github.com/goccy/llamawasm2go/p1.Fn363
func Fn363(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn364 github.com/goccy/llamawasm2go/p1.Fn364
func Fn364(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn368 github.com/goccy/llamawasm2go/p1.Fn368
func Fn368(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn376 github.com/goccy/llamawasm2go/p1.Fn376
func Fn376(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn396 github.com/goccy/llamawasm2go/p1.Fn396
func Fn396(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn397 github.com/goccy/llamawasm2go/p0.Fn397
func Fn397(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn402 github.com/goccy/llamawasm2go/p1.Fn402
func Fn402(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn405 github.com/goccy/llamawasm2go/p1.Fn405
func Fn405(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn406 github.com/goccy/llamawasm2go/p1.Fn406
func Fn406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn449 github.com/goccy/llamawasm2go/p1.Fn449
func Fn449(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn539 github.com/goccy/llamawasm2go/p1.Fn539
func Fn539(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn541 github.com/goccy/llamawasm2go/p1.Fn541
func Fn541(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn547 github.com/goccy/llamawasm2go/p1.Fn547
func Fn547(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn557 github.com/goccy/llamawasm2go/p1.Fn557
func Fn557(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn561 github.com/goccy/llamawasm2go/p1.Fn561
func Fn561(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn612 github.com/goccy/llamawasm2go/p0.Fn612
func Fn612(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn653 github.com/goccy/llamawasm2go/p1.Fn653
func Fn653(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn657 github.com/goccy/llamawasm2go/p0.Fn657
func Fn657(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn667 github.com/goccy/llamawasm2go/p1.Fn667
func Fn667(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn691 github.com/goccy/llamawasm2go/p1.Fn691
func Fn691(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn709 github.com/goccy/llamawasm2go/p1.Fn709
func Fn709(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn740 github.com/goccy/llamawasm2go/p1.Fn740
func Fn740(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn741 github.com/goccy/llamawasm2go/p1.Fn741
func Fn741(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn755 github.com/goccy/llamawasm2go/p1.Fn755
func Fn755(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn757 github.com/goccy/llamawasm2go/p0.Fn757
func Fn757(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn761 github.com/goccy/llamawasm2go/p1.Fn761
func Fn761(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn762 github.com/goccy/llamawasm2go/p1.Fn762
func Fn762(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn779 github.com/goccy/llamawasm2go/p0.Fn779
func Fn779(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn781 github.com/goccy/llamawasm2go/p0.Fn781
func Fn781(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn789 github.com/goccy/llamawasm2go/p0.Fn789
func Fn789(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn791 github.com/goccy/llamawasm2go/p1.Fn791
func Fn791(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn799 github.com/goccy/llamawasm2go/p1.Fn799
func Fn799(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn804 github.com/goccy/llamawasm2go/p1.Fn804
func Fn804(m *base.Module)

//go:linkname Fn805 github.com/goccy/llamawasm2go/p0.Fn805
func Fn805(m *base.Module, l0 int64)

//go:linkname Fn947 github.com/goccy/llamawasm2go/p0.Fn947
func Fn947(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1012 github.com/goccy/llamawasm2go/p1.Fn1012
func Fn1012(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1016 github.com/goccy/llamawasm2go/p1.Fn1016
func Fn1016(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1029 github.com/goccy/llamawasm2go/p1.Fn1029
func Fn1029(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1034 github.com/goccy/llamawasm2go/p1.Fn1034
func Fn1034(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1049 github.com/goccy/llamawasm2go/p1.Fn1049
func Fn1049(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1054 github.com/goccy/llamawasm2go/p1.Fn1054
func Fn1054(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1063 github.com/goccy/llamawasm2go/p1.Fn1063
func Fn1063(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1066 github.com/goccy/llamawasm2go/p1.Fn1066
func Fn1066(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1068 github.com/goccy/llamawasm2go/p1.Fn1068
func Fn1068(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1077 github.com/goccy/llamawasm2go/p1.Fn1077
func Fn1077(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1095 github.com/goccy/llamawasm2go/p0.Fn1095
func Fn1095(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1099 github.com/goccy/llamawasm2go/p0.Fn1099
func Fn1099(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1103 github.com/goccy/llamawasm2go/p1.Fn1103
func Fn1103(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1107 github.com/goccy/llamawasm2go/p1.Fn1107
func Fn1107(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1116 github.com/goccy/llamawasm2go/p0.Fn1116
func Fn1116(m *base.Module, l0 int64) int64

//go:linkname Fn1285 github.com/goccy/llamawasm2go/p1.Fn1285
func Fn1285(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1287 github.com/goccy/llamawasm2go/p0.Fn1287
func Fn1287(m *base.Module, l0 int64)

//go:linkname Fn1294 github.com/goccy/llamawasm2go/p0.Fn1294
func Fn1294(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1362 github.com/goccy/llamawasm2go/p1.Fn1362
func Fn1362(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1374 github.com/goccy/llamawasm2go/p1.Fn1374
func Fn1374(m *base.Module, l0 int64) int64

//go:linkname Fn1415 github.com/goccy/llamawasm2go/p1.Fn1415
func Fn1415(m *base.Module, l0 int64)

//go:linkname Fn1425 github.com/goccy/llamawasm2go/p1.Fn1425
func Fn1425(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p1.Fn1426
func Fn1426(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1449 github.com/goccy/llamawasm2go/p1.Fn1449
func Fn1449(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1452 github.com/goccy/llamawasm2go/p1.Fn1452
func Fn1452(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1461 github.com/goccy/llamawasm2go/p1.Fn1461
func Fn1461(m *base.Module, l0 int64)

//go:linkname Fn1466 github.com/goccy/llamawasm2go/p1.Fn1466
func Fn1466(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1470 github.com/goccy/llamawasm2go/p0.Fn1470
func Fn1470(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1518 github.com/goccy/llamawasm2go/p1.Fn1518
func Fn1518(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1529 github.com/goccy/llamawasm2go/p1.Fn1529
func Fn1529(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1546 github.com/goccy/llamawasm2go/p1.Fn1546
func Fn1546(m *base.Module, l0 int64)

//go:linkname Fn1554 github.com/goccy/llamawasm2go/p1.Fn1554
func Fn1554(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1560 github.com/goccy/llamawasm2go/p1.Fn1560
func Fn1560(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1561 github.com/goccy/llamawasm2go/p1.Fn1561
func Fn1561(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p1.Fn1563
func Fn1563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1574 github.com/goccy/llamawasm2go/p1.Fn1574
func Fn1574(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1672 github.com/goccy/llamawasm2go/p0.Fn1672
func Fn1672(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1703 github.com/goccy/llamawasm2go/p0.Fn1703
func Fn1703(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1711 github.com/goccy/llamawasm2go/p1.Fn1711
func Fn1711(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1751 github.com/goccy/llamawasm2go/p1.Fn1751
func Fn1751(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1752 github.com/goccy/llamawasm2go/p0.Fn1752
func Fn1752(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1778 github.com/goccy/llamawasm2go/p1.Fn1778
func Fn1778(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1849 github.com/goccy/llamawasm2go/p1.Fn1849
func Fn1849(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1862 github.com/goccy/llamawasm2go/p1.Fn1862
func Fn1862(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1947 github.com/goccy/llamawasm2go/p1.Fn1947
func Fn1947(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2004 github.com/goccy/llamawasm2go/p1.Fn2004
func Fn2004(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2007 github.com/goccy/llamawasm2go/p1.Fn2007
func Fn2007(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2015 github.com/goccy/llamawasm2go/p1.Fn2015
func Fn2015(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2031 github.com/goccy/llamawasm2go/p1.Fn2031
func Fn2031(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2051 github.com/goccy/llamawasm2go/p1.Fn2051
func Fn2051(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2063 github.com/goccy/llamawasm2go/p1.Fn2063
func Fn2063(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2120 github.com/goccy/llamawasm2go/p1.Fn2120
func Fn2120(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2121 github.com/goccy/llamawasm2go/p1.Fn2121
func Fn2121(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2225 github.com/goccy/llamawasm2go/p1.Fn2225
func Fn2225(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2233 github.com/goccy/llamawasm2go/p0.Fn2233
func Fn2233(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p0.Fn2249
func Fn2249(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2253 github.com/goccy/llamawasm2go/p1.Fn2253
func Fn2253(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2255 github.com/goccy/llamawasm2go/p1.Fn2255
func Fn2255(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2267 github.com/goccy/llamawasm2go/p1.Fn2267
func Fn2267(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2315 github.com/goccy/llamawasm2go/p1.Fn2315
func Fn2315(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2318 github.com/goccy/llamawasm2go/p1.Fn2318
func Fn2318(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2343 github.com/goccy/llamawasm2go/p1.Fn2343
func Fn2343(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2369 github.com/goccy/llamawasm2go/p0.Fn2369
func Fn2369(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2374 github.com/goccy/llamawasm2go/p1.Fn2374
func Fn2374(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2375 github.com/goccy/llamawasm2go/p0.Fn2375
func Fn2375(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2381 github.com/goccy/llamawasm2go/p1.Fn2381
func Fn2381(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2413 github.com/goccy/llamawasm2go/p1.Fn2413
func Fn2413(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2421 github.com/goccy/llamawasm2go/p0.Fn2421
func Fn2421(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2433 github.com/goccy/llamawasm2go/p1.Fn2433
func Fn2433(m *base.Module)

//go:linkname Fn2469 github.com/goccy/llamawasm2go/p0.Fn2469
func Fn2469(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2479 github.com/goccy/llamawasm2go/p1.Fn2479
func Fn2479(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2502 github.com/goccy/llamawasm2go/p1.Fn2502
func Fn2502(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2513 github.com/goccy/llamawasm2go/p1.Fn2513
func Fn2513(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2514 github.com/goccy/llamawasm2go/p1.Fn2514
func Fn2514(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2583 github.com/goccy/llamawasm2go/p1.Fn2583
func Fn2583(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2589 github.com/goccy/llamawasm2go/p1.Fn2589
func Fn2589(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2600 github.com/goccy/llamawasm2go/p1.Fn2600
func Fn2600(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2611 github.com/goccy/llamawasm2go/p1.Fn2611
func Fn2611(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2651 github.com/goccy/llamawasm2go/p1.Fn2651
func Fn2651(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2658 github.com/goccy/llamawasm2go/p1.Fn2658
func Fn2658(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2659 github.com/goccy/llamawasm2go/p1.Fn2659
func Fn2659(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2713 github.com/goccy/llamawasm2go/p0.Fn2713
func Fn2713(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2745 github.com/goccy/llamawasm2go/p1.Fn2745
func Fn2745(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2747 github.com/goccy/llamawasm2go/p1.Fn2747
func Fn2747(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2774 github.com/goccy/llamawasm2go/p0.Fn2774
func Fn2774(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2877 github.com/goccy/llamawasm2go/p1.Fn2877
func Fn2877(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2884 github.com/goccy/llamawasm2go/p1.Fn2884
func Fn2884(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2954 github.com/goccy/llamawasm2go/p1.Fn2954
func Fn2954(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2963 github.com/goccy/llamawasm2go/p1.Fn2963
func Fn2963(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2978 github.com/goccy/llamawasm2go/p1.Fn2978
func Fn2978(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn3003 github.com/goccy/llamawasm2go/p1.Fn3003
func Fn3003(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3030 github.com/goccy/llamawasm2go/p1.Fn3030
func Fn3030(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3042 github.com/goccy/llamawasm2go/p1.Fn3042
func Fn3042(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3057 github.com/goccy/llamawasm2go/p1.Fn3057
func Fn3057(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3063 github.com/goccy/llamawasm2go/p0.Fn3063
func Fn3063(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3067 github.com/goccy/llamawasm2go/p0.Fn3067
func Fn3067(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3072 github.com/goccy/llamawasm2go/p0.Fn3072
func Fn3072(m *base.Module, l0 int64) int64

//go:linkname Fn3073 github.com/goccy/llamawasm2go/p1.Fn3073
func Fn3073(m *base.Module, l0 int64)

//go:linkname Fn3075 github.com/goccy/llamawasm2go/p1.Fn3075
func Fn3075(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3085 github.com/goccy/llamawasm2go/p1.Fn3085
func Fn3085(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3087 github.com/goccy/llamawasm2go/p1.Fn3087
func Fn3087(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3088 github.com/goccy/llamawasm2go/p1.Fn3088
func Fn3088(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3089 github.com/goccy/llamawasm2go/p1.Fn3089
func Fn3089(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3121 github.com/goccy/llamawasm2go/p1.Fn3121
func Fn3121(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)
