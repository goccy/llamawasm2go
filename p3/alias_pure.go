//go:build !arm64 && (!amd64 || !amd64.v2)

package p3

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn21 github.com/goccy/llamawasm2go/p2.Fn21
func Fn21(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn22 github.com/goccy/llamawasm2go/p2.Fn22
func Fn22(m *base.Module, l0 int64) int64

//go:linkname Fn23 github.com/goccy/llamawasm2go/p2.Fn23
func Fn23(m *base.Module, l0 int64)

//go:linkname Fn56 github.com/goccy/llamawasm2go/p2.Fn56
func Fn56(m *base.Module, l0 int64) int64

//go:linkname Fn57 github.com/goccy/llamawasm2go/p2.Fn57
func Fn57(m *base.Module)

//go:linkname Fn59 github.com/goccy/llamawasm2go/p2.Fn59
func Fn59(m *base.Module)

//go:linkname Fn127 github.com/goccy/llamawasm2go/p2.Fn127
func Fn127(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn133 github.com/goccy/llamawasm2go/p2.Fn133
func Fn133(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn240 github.com/goccy/llamawasm2go/p2.Fn240
func Fn240(m *base.Module)

//go:linkname Fn241 github.com/goccy/llamawasm2go/p2.Fn241
func Fn241(m *base.Module, l0 int64)

//go:linkname Fn258 github.com/goccy/llamawasm2go/p2.Fn258
func Fn258(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn259 github.com/goccy/llamawasm2go/p2.Fn259
func Fn259(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn261 github.com/goccy/llamawasm2go/p2.Fn261
func Fn261(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn274 github.com/goccy/llamawasm2go/p2.Fn274
func Fn274(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64)

//go:linkname Fn277 github.com/goccy/llamawasm2go/p2.Fn277
func Fn277(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn293 github.com/goccy/llamawasm2go/p2.Fn293
func Fn293(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn341 github.com/goccy/llamawasm2go/p2.Fn341
func Fn341(m *base.Module, l0 int64) int32

//go:linkname Fn347 github.com/goccy/llamawasm2go/p2.Fn347
func Fn347(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn355 github.com/goccy/llamawasm2go/p2.Fn355
func Fn355(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn357 github.com/goccy/llamawasm2go/p2.Fn357
func Fn357(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn367 github.com/goccy/llamawasm2go/p2.Fn367
func Fn367(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn414 github.com/goccy/llamawasm2go/p2.Fn414
func Fn414(m *base.Module, l0 int64)

//go:linkname Fn450 github.com/goccy/llamawasm2go/p2.Fn450
func Fn450(m *base.Module, l0 int64) int64

//go:linkname Fn474 github.com/goccy/llamawasm2go/p2.Fn474
func Fn474(m *base.Module, l0 int64) int64

//go:linkname Fn481 github.com/goccy/llamawasm2go/p2.Fn481
func Fn481(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn493 github.com/goccy/llamawasm2go/p2.Fn493
func Fn493(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn520 github.com/goccy/llamawasm2go/p2.Fn520
func Fn520(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn531 github.com/goccy/llamawasm2go/p2.Fn531
func Fn531(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn537 github.com/goccy/llamawasm2go/p2.Fn537
func Fn537(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn544 github.com/goccy/llamawasm2go/p2.Fn544
func Fn544(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn545 github.com/goccy/llamawasm2go/p2.Fn545
func Fn545(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn547 github.com/goccy/llamawasm2go/p2.Fn547
func Fn547(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn554 github.com/goccy/llamawasm2go/p2.Fn554
func Fn554(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn555 github.com/goccy/llamawasm2go/p2.Fn555
func Fn555(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn558 github.com/goccy/llamawasm2go/p2.Fn558
func Fn558(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn561 github.com/goccy/llamawasm2go/p2.Fn561
func Fn561(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32, l14 int32) int64

//go:linkname Fn579 github.com/goccy/llamawasm2go/p1.Fn579
func Fn579(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64) int64

//go:linkname Fn585 github.com/goccy/llamawasm2go/p2.Fn585
func Fn585(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 int32) int64

//go:linkname Fn591 github.com/goccy/llamawasm2go/p2.Fn591
func Fn591(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn603 github.com/goccy/llamawasm2go/p1.Fn603
func Fn603(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn612 github.com/goccy/llamawasm2go/p2.Fn612
func Fn612(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn618 github.com/goccy/llamawasm2go/p2.Fn618
func Fn618(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn642 github.com/goccy/llamawasm2go/p2.Fn642
func Fn642(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn644 github.com/goccy/llamawasm2go/p2.Fn644
func Fn644(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn718 github.com/goccy/llamawasm2go/p1.Fn718
func Fn718(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn746 github.com/goccy/llamawasm2go/p1.Fn746
func Fn746(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn749 github.com/goccy/llamawasm2go/p2.Fn749
func Fn749(m *base.Module, l0 int64) int64

//go:linkname Fn778 github.com/goccy/llamawasm2go/p2.Fn778
func Fn778(m *base.Module, l0 int32)

//go:linkname Fn802 github.com/goccy/llamawasm2go/p1.Fn802
func Fn802(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn803 github.com/goccy/llamawasm2go/p1.Fn803
func Fn803(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn817 github.com/goccy/llamawasm2go/p1.Fn817
func Fn817(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn819 github.com/goccy/llamawasm2go/p1.Fn819
func Fn819(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn823 github.com/goccy/llamawasm2go/p1.Fn823
func Fn823(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn824 github.com/goccy/llamawasm2go/p2.Fn824
func Fn824(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn841 github.com/goccy/llamawasm2go/p1.Fn841
func Fn841(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn843 github.com/goccy/llamawasm2go/p1.Fn843
func Fn843(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn851 github.com/goccy/llamawasm2go/p1.Fn851
func Fn851(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn862 github.com/goccy/llamawasm2go/p1.Fn862
func Fn862(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64)

//go:linkname Fn864 github.com/goccy/llamawasm2go/p2.Fn864
func Fn864(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn866 github.com/goccy/llamawasm2go/p2.Fn866
func Fn866(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn867 github.com/goccy/llamawasm2go/p2.Fn867
func Fn867(m *base.Module)

//go:linkname Fn934 github.com/goccy/llamawasm2go/p2.Fn934
func Fn934(m *base.Module, l0 int64) int64

//go:linkname Fn978 github.com/goccy/llamawasm2go/p2.Fn978
func Fn978(m *base.Module)

//go:linkname Fn1019 github.com/goccy/llamawasm2go/p0.Fn1019
func Fn1019(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1063 github.com/goccy/llamawasm2go/p2.Fn1063
func Fn1063(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1072 github.com/goccy/llamawasm2go/p2.Fn1072
func Fn1072(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1084 github.com/goccy/llamawasm2go/p1.Fn1084
func Fn1084(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1088 github.com/goccy/llamawasm2go/p1.Fn1088
func Fn1088(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1101 github.com/goccy/llamawasm2go/p1.Fn1101
func Fn1101(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1106 github.com/goccy/llamawasm2go/p1.Fn1106
func Fn1106(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1119 github.com/goccy/llamawasm2go/p2.Fn1119
func Fn1119(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64) int64

//go:linkname Fn1121 github.com/goccy/llamawasm2go/p2.Fn1121
func Fn1121(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1126 github.com/goccy/llamawasm2go/p2.Fn1126
func Fn1126(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1133 github.com/goccy/llamawasm2go/p2.Fn1133
func Fn1133(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64) int64

//go:linkname Fn1135 github.com/goccy/llamawasm2go/p2.Fn1135
func Fn1135(m *base.Module, l0 int64, l1 int64, l2 int32, l3 float64) int64

//go:linkname Fn1138 github.com/goccy/llamawasm2go/p2.Fn1138
func Fn1138(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1140 github.com/goccy/llamawasm2go/p2.Fn1140
func Fn1140(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1149 github.com/goccy/llamawasm2go/p2.Fn1149
func Fn1149(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1188 github.com/goccy/llamawasm2go/p0.Fn1188
func Fn1188(m *base.Module, l0 int64) int64

//go:linkname Fn1305 github.com/goccy/llamawasm2go/p2.Fn1305
func Fn1305(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1307 github.com/goccy/llamawasm2go/p2.Fn1307
func Fn1307(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1341 github.com/goccy/llamawasm2go/p2.Fn1341
func Fn1341(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1345 github.com/goccy/llamawasm2go/p2.Fn1345
func Fn1345(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1349 github.com/goccy/llamawasm2go/p2.Fn1349
func Fn1349(m *base.Module)

//go:linkname Fn1355 github.com/goccy/llamawasm2go/p2.Fn1355
func Fn1355(m *base.Module, l0 int64) int64

//go:linkname Fn1358 github.com/goccy/llamawasm2go/p2.Fn1358
func Fn1358(m *base.Module)

//go:linkname Fn1365 github.com/goccy/llamawasm2go/p0.Fn1365
func Fn1365(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1435 github.com/goccy/llamawasm2go/p2.Fn1435
func Fn1435(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1460 github.com/goccy/llamawasm2go/p2.Fn1460
func Fn1460(m *base.Module, l0 int64)

//go:linkname Fn1476 github.com/goccy/llamawasm2go/p2.Fn1476
func Fn1476(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1510 github.com/goccy/llamawasm2go/p1.Fn1510
func Fn1510(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1531 github.com/goccy/llamawasm2go/p2.Fn1531
func Fn1531(m *base.Module, l0 int64) int64

//go:linkname Fn1533 github.com/goccy/llamawasm2go/p2.Fn1533
func Fn1533(m *base.Module, l0 int64)

//go:linkname Fn1552 github.com/goccy/llamawasm2go/p2.Fn1552
func Fn1552(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1559 github.com/goccy/llamawasm2go/p0.Fn1559
func Fn1559(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1574 github.com/goccy/llamawasm2go/p2.Fn1574
func Fn1574(m *base.Module, l0 int64)

//go:linkname Fn1653 github.com/goccy/llamawasm2go/p2.Fn1653
func Fn1653(m *base.Module, l0 int64)

//go:linkname Fn1655 github.com/goccy/llamawasm2go/p2.Fn1655
func Fn1655(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1659 github.com/goccy/llamawasm2go/p2.Fn1659
func Fn1659(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1663 github.com/goccy/llamawasm2go/p2.Fn1663
func Fn1663(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int32, l11 int32)

//go:linkname Fn1666 github.com/goccy/llamawasm2go/p1.Fn1666
func Fn1666(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int32, l14 int32, l15 float32, l16 int32, l17 int32, l18 int64, l19 int64, l20 int64, l21 int64, l22 int64, l23 int64, l24 int64, l25 float32) int64

//go:linkname Fn1680 github.com/goccy/llamawasm2go/p2.Fn1680
func Fn1680(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1687 github.com/goccy/llamawasm2go/p2.Fn1687
func Fn1687(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1702 github.com/goccy/llamawasm2go/p2.Fn1702
func Fn1702(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int64) int64

//go:linkname Fn1745 github.com/goccy/llamawasm2go/p2.Fn1745
func Fn1745(m *base.Module, l0 int64) int64

//go:linkname Fn1770 github.com/goccy/llamawasm2go/p2.Fn1770
func Fn1770(m *base.Module, l0 int64) int32

//go:linkname Fn1784 github.com/goccy/llamawasm2go/p2.Fn1784
func Fn1784(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1791 github.com/goccy/llamawasm2go/p2.Fn1791
func Fn1791(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1792 github.com/goccy/llamawasm2go/p2.Fn1792
func Fn1792(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1807 github.com/goccy/llamawasm2go/p2.Fn1807
func Fn1807(m *base.Module, l0 int64)

//go:linkname Fn1813 github.com/goccy/llamawasm2go/p2.Fn1813
func Fn1813(m *base.Module, l0 int64) int64

//go:linkname Fn1841 github.com/goccy/llamawasm2go/p2.Fn1841
func Fn1841(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1848 github.com/goccy/llamawasm2go/p2.Fn1848
func Fn1848(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1850 github.com/goccy/llamawasm2go/p2.Fn1850
func Fn1850(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64) int64

//go:linkname Fn1851 github.com/goccy/llamawasm2go/p2.Fn1851
func Fn1851(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1853 github.com/goccy/llamawasm2go/p2.Fn1853
func Fn1853(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1856 github.com/goccy/llamawasm2go/p2.Fn1856
func Fn1856(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1861 github.com/goccy/llamawasm2go/p1.Fn1861
func Fn1861(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64)

//go:linkname Fn1867 github.com/goccy/llamawasm2go/p2.Fn1867
func Fn1867(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1872 github.com/goccy/llamawasm2go/p2.Fn1872
func Fn1872(m *base.Module, l0 int64) int64

//go:linkname Fn1886 github.com/goccy/llamawasm2go/p2.Fn1886
func Fn1886(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1887 github.com/goccy/llamawasm2go/p0.Fn1887
func Fn1887(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1889 github.com/goccy/llamawasm2go/p2.Fn1889
func Fn1889(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1890 github.com/goccy/llamawasm2go/p2.Fn1890
func Fn1890(m *base.Module, l0 int64) int64

//go:linkname Fn1898 github.com/goccy/llamawasm2go/p2.Fn1898
func Fn1898(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1941 github.com/goccy/llamawasm2go/p2.Fn1941
func Fn1941(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1943 github.com/goccy/llamawasm2go/p2.Fn1943
func Fn1943(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1977 github.com/goccy/llamawasm2go/p2.Fn1977
func Fn1977(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1979 github.com/goccy/llamawasm2go/p2.Fn1979
func Fn1979(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2004 github.com/goccy/llamawasm2go/p2.Fn2004
func Fn2004(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2005 github.com/goccy/llamawasm2go/p2.Fn2005
func Fn2005(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2013 github.com/goccy/llamawasm2go/p2.Fn2013
func Fn2013(m *base.Module, l0 int64) int64

//go:linkname Fn2015 github.com/goccy/llamawasm2go/p2.Fn2015
func Fn2015(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2023 github.com/goccy/llamawasm2go/p2.Fn2023
func Fn2023(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2065 github.com/goccy/llamawasm2go/p2.Fn2065
func Fn2065(m *base.Module, l0 int64) int64

//go:linkname Fn2074 github.com/goccy/llamawasm2go/p2.Fn2074
func Fn2074(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn2075 github.com/goccy/llamawasm2go/p2.Fn2075
func Fn2075(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2078 github.com/goccy/llamawasm2go/p2.Fn2078
func Fn2078(m *base.Module, l0 int64) int64

//go:linkname Fn2080 github.com/goccy/llamawasm2go/p2.Fn2080
func Fn2080(m *base.Module, l0 int64) int64

//go:linkname Fn2092 github.com/goccy/llamawasm2go/p0.Fn2092
func Fn2092(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2138 github.com/goccy/llamawasm2go/p1.Fn2138
func Fn2138(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2157 github.com/goccy/llamawasm2go/p2.Fn2157
func Fn2157(m *base.Module, l0 int64) int64

//go:linkname Fn2186 github.com/goccy/llamawasm2go/p2.Fn2186
func Fn2186(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2189 github.com/goccy/llamawasm2go/p2.Fn2189
func Fn2189(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2192 github.com/goccy/llamawasm2go/p2.Fn2192
func Fn2192(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2194 github.com/goccy/llamawasm2go/p2.Fn2194
func Fn2194(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2198 github.com/goccy/llamawasm2go/p2.Fn2198
func Fn2198(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2200 github.com/goccy/llamawasm2go/p2.Fn2200
func Fn2200(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2201 github.com/goccy/llamawasm2go/p2.Fn2201
func Fn2201(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2206 github.com/goccy/llamawasm2go/p2.Fn2206
func Fn2206(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2209 github.com/goccy/llamawasm2go/p1.Fn2209
func Fn2209(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2230 github.com/goccy/llamawasm2go/p1.Fn2230
func Fn2230(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2338 github.com/goccy/llamawasm2go/p2.Fn2338
func Fn2338(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2349 github.com/goccy/llamawasm2go/p2.Fn2349
func Fn2349(m *base.Module, l0 int64) int64

//go:linkname Fn2352 github.com/goccy/llamawasm2go/p2.Fn2352
func Fn2352(m *base.Module, l0 int64) int64

//go:linkname Fn2362 github.com/goccy/llamawasm2go/p0.Fn2362
func Fn2362(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn2391 github.com/goccy/llamawasm2go/p1.Fn2391
func Fn2391(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2392 github.com/goccy/llamawasm2go/p1.Fn2392
func Fn2392(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2399 github.com/goccy/llamawasm2go/p1.Fn2399
func Fn2399(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2405 github.com/goccy/llamawasm2go/p2.Fn2405
func Fn2405(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2446 github.com/goccy/llamawasm2go/p2.Fn2446
func Fn2446(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2447 github.com/goccy/llamawasm2go/p2.Fn2447
func Fn2447(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2532 github.com/goccy/llamawasm2go/p2.Fn2532
func Fn2532(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2543 github.com/goccy/llamawasm2go/p2.Fn2543
func Fn2543(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2601 github.com/goccy/llamawasm2go/p1.Fn2601
func Fn2601(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2657 github.com/goccy/llamawasm2go/p1.Fn2657
func Fn2657(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2658 github.com/goccy/llamawasm2go/p2.Fn2658
func Fn2658(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2660 github.com/goccy/llamawasm2go/p2.Fn2660
func Fn2660(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2666 github.com/goccy/llamawasm2go/p2.Fn2666
func Fn2666(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2673 github.com/goccy/llamawasm2go/p1.Fn2673
func Fn2673(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2710 github.com/goccy/llamawasm2go/p1.Fn2710
func Fn2710(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2824 github.com/goccy/llamawasm2go/p1.Fn2824
func Fn2824(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2831 github.com/goccy/llamawasm2go/p1.Fn2831
func Fn2831(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2836 github.com/goccy/llamawasm2go/p1.Fn2836
func Fn2836(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2853 github.com/goccy/llamawasm2go/p2.Fn2853
func Fn2853(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2854 github.com/goccy/llamawasm2go/p2.Fn2854
func Fn2854(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3050 github.com/goccy/llamawasm2go/p2.Fn3050
func Fn3050(m *base.Module, l0 int64) int64

//go:linkname Fn3077 github.com/goccy/llamawasm2go/p2.Fn3077
func Fn3077(m *base.Module, l0 int64)

//go:linkname Fn3078 github.com/goccy/llamawasm2go/p2.Fn3078
func Fn3078(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3100 github.com/goccy/llamawasm2go/p2.Fn3100
func Fn3100(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3101 github.com/goccy/llamawasm2go/p2.Fn3101
func Fn3101(m *base.Module, l0 int32) int64

//go:linkname Fn3110 github.com/goccy/llamawasm2go/p2.Fn3110
func Fn3110(m *base.Module, l0 int64, l1 int32, l2 int32, l3 float32, l4 float32, l5 float32) int64

//go:linkname Fn3111 github.com/goccy/llamawasm2go/p2.Fn3111
func Fn3111(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn3130 github.com/goccy/llamawasm2go/p2.Fn3130
func Fn3130(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3149 github.com/goccy/llamawasm2go/p2.Fn3149
func Fn3149(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3161 github.com/goccy/llamawasm2go/p2.Fn3161
func Fn3161(m *base.Module, l0 int64)

//go:linkname Fn3176 github.com/goccy/llamawasm2go/p2.Fn3176
func Fn3176(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3193 github.com/goccy/llamawasm2go/p2.Fn3193
func Fn3193(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3211 github.com/goccy/llamawasm2go/p0.Fn3211
func Fn3211(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3224 github.com/goccy/llamawasm2go/p1.Fn3224
func Fn3224(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn3236 github.com/goccy/llamawasm2go/p2.Fn3236
func Fn3236(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3282 github.com/goccy/llamawasm2go/p2.Fn3282
func Fn3282(m *base.Module)

//go:linkname Fn3383 github.com/goccy/llamawasm2go/p2.Fn3383
func Fn3383(m *base.Module)

//go:linkname Fn3384 github.com/goccy/llamawasm2go/p2.Fn3384
func Fn3384(m *base.Module, l0 int64) int64

//go:linkname Fn3388 github.com/goccy/llamawasm2go/p2.Fn3388
func Fn3388(m *base.Module, l0 int64)

//go:linkname Fn3390 github.com/goccy/llamawasm2go/p2.Fn3390
func Fn3390(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3399 github.com/goccy/llamawasm2go/p2.Fn3399
func Fn3399(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3400 github.com/goccy/llamawasm2go/p2.Fn3400
func Fn3400(m *base.Module, l0 int64) int32

//go:linkname Fn3408 github.com/goccy/llamawasm2go/p1.Fn3408
func Fn3408(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3416 github.com/goccy/llamawasm2go/p2.Fn3416
func Fn3416(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3420 github.com/goccy/llamawasm2go/p2.Fn3420
func Fn3420(m *base.Module) int64

//go:linkname Fn3447 github.com/goccy/llamawasm2go/p2.Fn3447
func Fn3447(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn3465 github.com/goccy/llamawasm2go/p2.Fn3465
func Fn3465(m *base.Module, l0 int64) int32

//go:linkname Fn3471 github.com/goccy/llamawasm2go/p2.Fn3471
func Fn3471(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3475 github.com/goccy/llamawasm2go/p2.Fn3475
func Fn3475(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3492 github.com/goccy/llamawasm2go/p2.Fn3492
func Fn3492(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3498 github.com/goccy/llamawasm2go/p2.Fn3498
func Fn3498(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3512 github.com/goccy/llamawasm2go/p2.Fn3512
func Fn3512(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3517 github.com/goccy/llamawasm2go/p1.Fn3517
func Fn3517(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3522 github.com/goccy/llamawasm2go/p0.Fn3522
func Fn3522(m *base.Module, l0 int64) int64

//go:linkname Fn3523 github.com/goccy/llamawasm2go/p2.Fn3523
func Fn3523(m *base.Module, l0 int64)

//go:linkname Fn3525 github.com/goccy/llamawasm2go/p2.Fn3525
func Fn3525(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3527 github.com/goccy/llamawasm2go/p2.Fn3527
func Fn3527(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3531 github.com/goccy/llamawasm2go/p2.Fn3531
func Fn3531(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3537 github.com/goccy/llamawasm2go/p2.Fn3537
func Fn3537(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3539 github.com/goccy/llamawasm2go/p2.Fn3539
func Fn3539(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3540 github.com/goccy/llamawasm2go/p1.Fn3540
func Fn3540(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3541 github.com/goccy/llamawasm2go/p1.Fn3541
func Fn3541(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3542 github.com/goccy/llamawasm2go/p1.Fn3542
func Fn3542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3543 github.com/goccy/llamawasm2go/p2.Fn3543
func Fn3543(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn3552 github.com/goccy/llamawasm2go/p2.Fn3552
func Fn3552(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn3553 github.com/goccy/llamawasm2go/p2.Fn3553
func Fn3553(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn3554 github.com/goccy/llamawasm2go/p2.Fn3554
func Fn3554(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn3555 github.com/goccy/llamawasm2go/p2.Fn3555
func Fn3555(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn3556 github.com/goccy/llamawasm2go/p2.Fn3556
func Fn3556(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn3557 github.com/goccy/llamawasm2go/p2.Fn3557
func Fn3557(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn3558 github.com/goccy/llamawasm2go/p2.Fn3558
func Fn3558(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn3559 github.com/goccy/llamawasm2go/p2.Fn3559
func Fn3559(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn3562 github.com/goccy/llamawasm2go/p2.Fn3562
func Fn3562(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3563 github.com/goccy/llamawasm2go/p2.Fn3563
func Fn3563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int64

//go:linkname Fn3566 github.com/goccy/llamawasm2go/p2.Fn3566
func Fn3566(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64) int64

//go:linkname Fn3568 github.com/goccy/llamawasm2go/p1.Fn3568
func Fn3568(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int32, l7 int64, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int64, l17 int32, l18 int64, l19 int32, l20 int64, l21 int64)

//go:linkname Fn3569 github.com/goccy/llamawasm2go/p2.Fn3569
func Fn3569(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int32) int64

//go:linkname Fn3573 github.com/goccy/llamawasm2go/p2.Fn3573
func Fn3573(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn3580 github.com/goccy/llamawasm2go/p2.Fn3580
func Fn3580(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3594 github.com/goccy/llamawasm2go/p2.Fn3594
func Fn3594(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn3595 github.com/goccy/llamawasm2go/p2.Fn3595
func Fn3595(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn3602 github.com/goccy/llamawasm2go/p2.Fn3602
func Fn3602(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64)
