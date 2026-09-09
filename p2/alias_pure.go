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

//go:linkname Fn364 github.com/goccy/llamawasm2go/p1.Fn364
func Fn364(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn368 github.com/goccy/llamawasm2go/p1.Fn368
func Fn368(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn369 github.com/goccy/llamawasm2go/p1.Fn369
func Fn369(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn370 github.com/goccy/llamawasm2go/p1.Fn370
func Fn370(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn374 github.com/goccy/llamawasm2go/p1.Fn374
func Fn374(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn378 github.com/goccy/llamawasm2go/p1.Fn378
func Fn378(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32)

//go:linkname Fn381 github.com/goccy/llamawasm2go/p1.Fn381
func Fn381(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn383 github.com/goccy/llamawasm2go/p1.Fn383
func Fn383(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn398 github.com/goccy/llamawasm2go/p0.Fn398
func Fn398(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn402 github.com/goccy/llamawasm2go/p0.Fn402
func Fn402(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn405 github.com/goccy/llamawasm2go/p1.Fn405
func Fn405(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn406 github.com/goccy/llamawasm2go/p1.Fn406
func Fn406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn415 github.com/goccy/llamawasm2go/p1.Fn415
func Fn415(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn416 github.com/goccy/llamawasm2go/p0.Fn416
func Fn416(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn420 github.com/goccy/llamawasm2go/p1.Fn420
func Fn420(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn422 github.com/goccy/llamawasm2go/p1.Fn422
func Fn422(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn423 github.com/goccy/llamawasm2go/p1.Fn423
func Fn423(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32)

//go:linkname Fn424 github.com/goccy/llamawasm2go/p1.Fn424
func Fn424(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn425 github.com/goccy/llamawasm2go/p1.Fn425
func Fn425(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn473 github.com/goccy/llamawasm2go/p1.Fn473
func Fn473(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn563 github.com/goccy/llamawasm2go/p1.Fn563
func Fn563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn565 github.com/goccy/llamawasm2go/p1.Fn565
func Fn565(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn571 github.com/goccy/llamawasm2go/p1.Fn571
func Fn571(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn581 github.com/goccy/llamawasm2go/p1.Fn581
func Fn581(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn585 github.com/goccy/llamawasm2go/p1.Fn585
func Fn585(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn636 github.com/goccy/llamawasm2go/p0.Fn636
func Fn636(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn677 github.com/goccy/llamawasm2go/p1.Fn677
func Fn677(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn681 github.com/goccy/llamawasm2go/p0.Fn681
func Fn681(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn691 github.com/goccy/llamawasm2go/p1.Fn691
func Fn691(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn715 github.com/goccy/llamawasm2go/p1.Fn715
func Fn715(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn733 github.com/goccy/llamawasm2go/p1.Fn733
func Fn733(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn764 github.com/goccy/llamawasm2go/p1.Fn764
func Fn764(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn765 github.com/goccy/llamawasm2go/p1.Fn765
func Fn765(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn779 github.com/goccy/llamawasm2go/p1.Fn779
func Fn779(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn781 github.com/goccy/llamawasm2go/p0.Fn781
func Fn781(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn785 github.com/goccy/llamawasm2go/p1.Fn785
func Fn785(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn786 github.com/goccy/llamawasm2go/p1.Fn786
func Fn786(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn803 github.com/goccy/llamawasm2go/p0.Fn803
func Fn803(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn805 github.com/goccy/llamawasm2go/p0.Fn805
func Fn805(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn813 github.com/goccy/llamawasm2go/p0.Fn813
func Fn813(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn815 github.com/goccy/llamawasm2go/p1.Fn815
func Fn815(m *base.Module, l0 int32, l1 int64, l2 int64) int32

//go:linkname Fn823 github.com/goccy/llamawasm2go/p1.Fn823
func Fn823(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn828 github.com/goccy/llamawasm2go/p1.Fn828
func Fn828(m *base.Module)

//go:linkname Fn829 github.com/goccy/llamawasm2go/p0.Fn829
func Fn829(m *base.Module, l0 int64)

//go:linkname Fn972 github.com/goccy/llamawasm2go/p0.Fn972
func Fn972(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1037 github.com/goccy/llamawasm2go/p1.Fn1037
func Fn1037(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1041 github.com/goccy/llamawasm2go/p1.Fn1041
func Fn1041(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1054 github.com/goccy/llamawasm2go/p1.Fn1054
func Fn1054(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1059 github.com/goccy/llamawasm2go/p1.Fn1059
func Fn1059(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1074 github.com/goccy/llamawasm2go/p1.Fn1074
func Fn1074(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1079 github.com/goccy/llamawasm2go/p1.Fn1079
func Fn1079(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1088 github.com/goccy/llamawasm2go/p1.Fn1088
func Fn1088(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1091 github.com/goccy/llamawasm2go/p1.Fn1091
func Fn1091(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1093 github.com/goccy/llamawasm2go/p1.Fn1093
func Fn1093(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1102 github.com/goccy/llamawasm2go/p1.Fn1102
func Fn1102(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1120 github.com/goccy/llamawasm2go/p0.Fn1120
func Fn1120(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1124 github.com/goccy/llamawasm2go/p0.Fn1124
func Fn1124(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1128 github.com/goccy/llamawasm2go/p1.Fn1128
func Fn1128(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1132 github.com/goccy/llamawasm2go/p1.Fn1132
func Fn1132(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int32, l8 int64, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int32)

//go:linkname Fn1141 github.com/goccy/llamawasm2go/p0.Fn1141
func Fn1141(m *base.Module, l0 int64) int64

//go:linkname Fn1310 github.com/goccy/llamawasm2go/p1.Fn1310
func Fn1310(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1312 github.com/goccy/llamawasm2go/p0.Fn1312
func Fn1312(m *base.Module, l0 int64)

//go:linkname Fn1319 github.com/goccy/llamawasm2go/p0.Fn1319
func Fn1319(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p1.Fn1397
func Fn1397(m *base.Module, l0 int64) int64

//go:linkname Fn1438 github.com/goccy/llamawasm2go/p1.Fn1438
func Fn1438(m *base.Module, l0 int64)

//go:linkname Fn1448 github.com/goccy/llamawasm2go/p1.Fn1448
func Fn1448(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1449 github.com/goccy/llamawasm2go/p1.Fn1449
func Fn1449(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32)

//go:linkname Fn1472 github.com/goccy/llamawasm2go/p1.Fn1472
func Fn1472(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1475 github.com/goccy/llamawasm2go/p1.Fn1475
func Fn1475(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1484 github.com/goccy/llamawasm2go/p1.Fn1484
func Fn1484(m *base.Module, l0 int64)

//go:linkname Fn1489 github.com/goccy/llamawasm2go/p1.Fn1489
func Fn1489(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1493 github.com/goccy/llamawasm2go/p0.Fn1493
func Fn1493(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p1.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p1.Fn1553
func Fn1553(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64)

//go:linkname Fn1570 github.com/goccy/llamawasm2go/p1.Fn1570
func Fn1570(m *base.Module, l0 int64)

//go:linkname Fn1578 github.com/goccy/llamawasm2go/p1.Fn1578
func Fn1578(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1584 github.com/goccy/llamawasm2go/p1.Fn1584
func Fn1584(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1585 github.com/goccy/llamawasm2go/p1.Fn1585
func Fn1585(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1587 github.com/goccy/llamawasm2go/p1.Fn1587
func Fn1587(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64) int64

//go:linkname Fn1598 github.com/goccy/llamawasm2go/p1.Fn1598
func Fn1598(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1696 github.com/goccy/llamawasm2go/p0.Fn1696
func Fn1696(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1727 github.com/goccy/llamawasm2go/p0.Fn1727
func Fn1727(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1735 github.com/goccy/llamawasm2go/p1.Fn1735
func Fn1735(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1775 github.com/goccy/llamawasm2go/p1.Fn1775
func Fn1775(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1776 github.com/goccy/llamawasm2go/p0.Fn1776
func Fn1776(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1802 github.com/goccy/llamawasm2go/p1.Fn1802
func Fn1802(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1873 github.com/goccy/llamawasm2go/p1.Fn1873
func Fn1873(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn1886 github.com/goccy/llamawasm2go/p1.Fn1886
func Fn1886(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1971 github.com/goccy/llamawasm2go/p1.Fn1971
func Fn1971(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2028 github.com/goccy/llamawasm2go/p1.Fn2028
func Fn2028(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2031 github.com/goccy/llamawasm2go/p1.Fn2031
func Fn2031(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2039 github.com/goccy/llamawasm2go/p1.Fn2039
func Fn2039(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2055 github.com/goccy/llamawasm2go/p1.Fn2055
func Fn2055(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2075 github.com/goccy/llamawasm2go/p1.Fn2075
func Fn2075(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2087 github.com/goccy/llamawasm2go/p1.Fn2087
func Fn2087(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2119 github.com/goccy/llamawasm2go/p1.Fn2119
func Fn2119(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2144 github.com/goccy/llamawasm2go/p1.Fn2144
func Fn2144(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2145 github.com/goccy/llamawasm2go/p1.Fn2145
func Fn2145(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2247 github.com/goccy/llamawasm2go/p1.Fn2247
func Fn2247(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn2255 github.com/goccy/llamawasm2go/p0.Fn2255
func Fn2255(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2271 github.com/goccy/llamawasm2go/p1.Fn2271
func Fn2271(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2275 github.com/goccy/llamawasm2go/p1.Fn2275
func Fn2275(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2277 github.com/goccy/llamawasm2go/p1.Fn2277
func Fn2277(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2289 github.com/goccy/llamawasm2go/p1.Fn2289
func Fn2289(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2337 github.com/goccy/llamawasm2go/p1.Fn2337
func Fn2337(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2340 github.com/goccy/llamawasm2go/p1.Fn2340
func Fn2340(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2365 github.com/goccy/llamawasm2go/p1.Fn2365
func Fn2365(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2391 github.com/goccy/llamawasm2go/p0.Fn2391
func Fn2391(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2396 github.com/goccy/llamawasm2go/p1.Fn2396
func Fn2396(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2397 github.com/goccy/llamawasm2go/p0.Fn2397
func Fn2397(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn2403 github.com/goccy/llamawasm2go/p1.Fn2403
func Fn2403(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2435 github.com/goccy/llamawasm2go/p1.Fn2435
func Fn2435(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2443 github.com/goccy/llamawasm2go/p0.Fn2443
func Fn2443(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2455 github.com/goccy/llamawasm2go/p1.Fn2455
func Fn2455(m *base.Module)

//go:linkname Fn2490 github.com/goccy/llamawasm2go/p0.Fn2490
func Fn2490(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2500 github.com/goccy/llamawasm2go/p1.Fn2500
func Fn2500(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2523 github.com/goccy/llamawasm2go/p1.Fn2523
func Fn2523(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2534 github.com/goccy/llamawasm2go/p1.Fn2534
func Fn2534(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2535 github.com/goccy/llamawasm2go/p1.Fn2535
func Fn2535(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int32

//go:linkname Fn2604 github.com/goccy/llamawasm2go/p1.Fn2604
func Fn2604(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2610 github.com/goccy/llamawasm2go/p1.Fn2610
func Fn2610(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2621 github.com/goccy/llamawasm2go/p1.Fn2621
func Fn2621(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2632 github.com/goccy/llamawasm2go/p1.Fn2632
func Fn2632(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2672 github.com/goccy/llamawasm2go/p1.Fn2672
func Fn2672(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2679 github.com/goccy/llamawasm2go/p1.Fn2679
func Fn2679(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2680 github.com/goccy/llamawasm2go/p1.Fn2680
func Fn2680(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2734 github.com/goccy/llamawasm2go/p0.Fn2734
func Fn2734(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2766 github.com/goccy/llamawasm2go/p1.Fn2766
func Fn2766(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2768 github.com/goccy/llamawasm2go/p1.Fn2768
func Fn2768(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2795 github.com/goccy/llamawasm2go/p0.Fn2795
func Fn2795(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2898 github.com/goccy/llamawasm2go/p1.Fn2898
func Fn2898(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2905 github.com/goccy/llamawasm2go/p1.Fn2905
func Fn2905(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2984 github.com/goccy/llamawasm2go/p1.Fn2984
func Fn2984(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2999 github.com/goccy/llamawasm2go/p1.Fn2999
func Fn2999(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn3051 github.com/goccy/llamawasm2go/p1.Fn3051
func Fn3051(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3063 github.com/goccy/llamawasm2go/p1.Fn3063
func Fn3063(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3078 github.com/goccy/llamawasm2go/p1.Fn3078
func Fn3078(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3084 github.com/goccy/llamawasm2go/p0.Fn3084
func Fn3084(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3088 github.com/goccy/llamawasm2go/p0.Fn3088
func Fn3088(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3093 github.com/goccy/llamawasm2go/p0.Fn3093
func Fn3093(m *base.Module, l0 int64) int64

//go:linkname Fn3094 github.com/goccy/llamawasm2go/p1.Fn3094
func Fn3094(m *base.Module, l0 int64)

//go:linkname Fn3096 github.com/goccy/llamawasm2go/p1.Fn3096
func Fn3096(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3106 github.com/goccy/llamawasm2go/p1.Fn3106
func Fn3106(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3108 github.com/goccy/llamawasm2go/p1.Fn3108
func Fn3108(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3109 github.com/goccy/llamawasm2go/p1.Fn3109
func Fn3109(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3110 github.com/goccy/llamawasm2go/p1.Fn3110
func Fn3110(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3142 github.com/goccy/llamawasm2go/p1.Fn3142
func Fn3142(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)
