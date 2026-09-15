//go:build !arm64 && (!amd64 || !amd64.v2)

package p1

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn21 github.com/goccy/llamawasm2go/p2.Fn21
func Fn21(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn48 github.com/goccy/llamawasm2go/p2.Fn48
func Fn48(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn51 github.com/goccy/llamawasm2go/p2.Fn51
func Fn51(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn55 github.com/goccy/llamawasm2go/p2.Fn55
func Fn55(m *base.Module, l0 int64) int64

//go:linkname Fn56 github.com/goccy/llamawasm2go/p2.Fn56
func Fn56(m *base.Module, l0 int64)

//go:linkname Fn58 github.com/goccy/llamawasm2go/p2.Fn58
func Fn58(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn59 github.com/goccy/llamawasm2go/p2.Fn59
func Fn59(m *base.Module, l0 int64) int64

//go:linkname Fn60 github.com/goccy/llamawasm2go/p2.Fn60
func Fn60(m *base.Module)

//go:linkname Fn66 github.com/goccy/llamawasm2go/p0.Fn66
func Fn66(m *base.Module, l0 int64) int64

//go:linkname Fn67 github.com/goccy/llamawasm2go/p2.Fn67
func Fn67(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn68 github.com/goccy/llamawasm2go/p0.Fn68
func Fn68(m *base.Module, l0 int64) int64

//go:linkname Fn84 github.com/goccy/llamawasm2go/p2.Fn84
func Fn84(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn129 github.com/goccy/llamawasm2go/p2.Fn129
func Fn129(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn135 github.com/goccy/llamawasm2go/p2.Fn135
func Fn135(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn168 github.com/goccy/llamawasm2go/p2.Fn168
func Fn168(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn242 github.com/goccy/llamawasm2go/p2.Fn242
func Fn242(m *base.Module)

//go:linkname Fn243 github.com/goccy/llamawasm2go/p2.Fn243
func Fn243(m *base.Module, l0 int64)

//go:linkname Fn244 github.com/goccy/llamawasm2go/p2.Fn244
func Fn244(m *base.Module, l0 int64) int64

//go:linkname Fn251 github.com/goccy/llamawasm2go/p2.Fn251
func Fn251(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn252 github.com/goccy/llamawasm2go/p2.Fn252
func Fn252(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn254 github.com/goccy/llamawasm2go/p2.Fn254
func Fn254(m *base.Module)

//go:linkname Fn255 github.com/goccy/llamawasm2go/p2.Fn255
func Fn255(m *base.Module)

//go:linkname Fn256 github.com/goccy/llamawasm2go/p2.Fn256
func Fn256(m *base.Module, l0 int64)

//go:linkname Fn260 github.com/goccy/llamawasm2go/p2.Fn260
func Fn260(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn261 github.com/goccy/llamawasm2go/p2.Fn261
func Fn261(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn262 github.com/goccy/llamawasm2go/p2.Fn262
func Fn262(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn264 github.com/goccy/llamawasm2go/p2.Fn264
func Fn264(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn265 github.com/goccy/llamawasm2go/p2.Fn265
func Fn265(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn267 github.com/goccy/llamawasm2go/p2.Fn267
func Fn267(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn268 github.com/goccy/llamawasm2go/p2.Fn268
func Fn268(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn269 github.com/goccy/llamawasm2go/p2.Fn269
func Fn269(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn270 github.com/goccy/llamawasm2go/p2.Fn270
func Fn270(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn271 github.com/goccy/llamawasm2go/p2.Fn271
func Fn271(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn273 github.com/goccy/llamawasm2go/p2.Fn273
func Fn273(m *base.Module)

//go:linkname Fn274 github.com/goccy/llamawasm2go/p2.Fn274
func Fn274(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn275 github.com/goccy/llamawasm2go/p2.Fn275
func Fn275(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64)

//go:linkname Fn277 github.com/goccy/llamawasm2go/p2.Fn277
func Fn277(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn278 github.com/goccy/llamawasm2go/p2.Fn278
func Fn278(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn279 github.com/goccy/llamawasm2go/p2.Fn279
func Fn279(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn282 github.com/goccy/llamawasm2go/p2.Fn282
func Fn282(m *base.Module, l0 int64) int64

//go:linkname Fn283 github.com/goccy/llamawasm2go/p2.Fn283
func Fn283(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn286 github.com/goccy/llamawasm2go/p2.Fn286
func Fn286(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn287 github.com/goccy/llamawasm2go/p2.Fn287
func Fn287(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn293 github.com/goccy/llamawasm2go/p2.Fn293
func Fn293(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn294 github.com/goccy/llamawasm2go/p2.Fn294
func Fn294(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn342 github.com/goccy/llamawasm2go/p2.Fn342
func Fn342(m *base.Module, l0 int64) int32

//go:linkname Fn343 github.com/goccy/llamawasm2go/p2.Fn343
func Fn343(m *base.Module, l0 int64)

//go:linkname Fn344 github.com/goccy/llamawasm2go/p2.Fn344
func Fn344(m *base.Module, l0 int64)

//go:linkname Fn361 github.com/goccy/llamawasm2go/p2.Fn361
func Fn361(m *base.Module) int64

//go:linkname Fn364 github.com/goccy/llamawasm2go/p2.Fn364
func Fn364(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn365 github.com/goccy/llamawasm2go/p2.Fn365
func Fn365(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn368 github.com/goccy/llamawasm2go/p2.Fn368
func Fn368(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn369 github.com/goccy/llamawasm2go/p2.Fn369
func Fn369(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn373 github.com/goccy/llamawasm2go/p2.Fn373
func Fn373(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn374 github.com/goccy/llamawasm2go/p2.Fn374
func Fn374(m *base.Module, l0 int64)

//go:linkname Fn375 github.com/goccy/llamawasm2go/p2.Fn375
func Fn375(m *base.Module, l0 int64)

//go:linkname Fn377 github.com/goccy/llamawasm2go/p2.Fn377
func Fn377(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn378 github.com/goccy/llamawasm2go/p2.Fn378
func Fn378(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn384 github.com/goccy/llamawasm2go/p2.Fn384
func Fn384(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn385 github.com/goccy/llamawasm2go/p2.Fn385
func Fn385(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn387 github.com/goccy/llamawasm2go/p2.Fn387
func Fn387(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn391 github.com/goccy/llamawasm2go/p0.Fn391
func Fn391(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int64)

//go:linkname Fn394 github.com/goccy/llamawasm2go/p2.Fn394
func Fn394(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn395 github.com/goccy/llamawasm2go/p2.Fn395
func Fn395(m *base.Module, l0 int64)

//go:linkname Fn398 github.com/goccy/llamawasm2go/p2.Fn398
func Fn398(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn399 github.com/goccy/llamawasm2go/p2.Fn399
func Fn399(m *base.Module, l0 int64)

//go:linkname Fn400 github.com/goccy/llamawasm2go/p2.Fn400
func Fn400(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn401 github.com/goccy/llamawasm2go/p2.Fn401
func Fn401(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn402 github.com/goccy/llamawasm2go/p2.Fn402
func Fn402(m *base.Module)

//go:linkname Fn413 github.com/goccy/llamawasm2go/p0.Fn413
func Fn413(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int32)

//go:linkname Fn417 github.com/goccy/llamawasm2go/p0.Fn417
func Fn417(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 int64)

//go:linkname Fn431 github.com/goccy/llamawasm2go/p2.Fn431
func Fn431(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn435 github.com/goccy/llamawasm2go/p2.Fn435
func Fn435(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn436 github.com/goccy/llamawasm2go/p2.Fn436
func Fn436(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn438 github.com/goccy/llamawasm2go/p2.Fn438
func Fn438(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn439 github.com/goccy/llamawasm2go/p2.Fn439
func Fn439(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn460 github.com/goccy/llamawasm2go/p2.Fn460
func Fn460(m *base.Module, l0 int64) int32

//go:linkname Fn461 github.com/goccy/llamawasm2go/p2.Fn461
func Fn461(m *base.Module, l0 int64) int32

//go:linkname Fn472 github.com/goccy/llamawasm2go/p2.Fn472
func Fn472(m *base.Module, l0 int64) int64

//go:linkname Fn477 github.com/goccy/llamawasm2go/p2.Fn477
func Fn477(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn479 github.com/goccy/llamawasm2go/p2.Fn479
func Fn479(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn480 github.com/goccy/llamawasm2go/p2.Fn480
func Fn480(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn481 github.com/goccy/llamawasm2go/p2.Fn481
func Fn481(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn482 github.com/goccy/llamawasm2go/p2.Fn482
func Fn482(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn483 github.com/goccy/llamawasm2go/p2.Fn483
func Fn483(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn484 github.com/goccy/llamawasm2go/p2.Fn484
func Fn484(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn487 github.com/goccy/llamawasm2go/p2.Fn487
func Fn487(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn488 github.com/goccy/llamawasm2go/p2.Fn488
func Fn488(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn491 github.com/goccy/llamawasm2go/p2.Fn491
func Fn491(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn492 github.com/goccy/llamawasm2go/p2.Fn492
func Fn492(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn493 github.com/goccy/llamawasm2go/p2.Fn493
func Fn493(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn494 github.com/goccy/llamawasm2go/p2.Fn494
func Fn494(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn495 github.com/goccy/llamawasm2go/p2.Fn495
func Fn495(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn496 github.com/goccy/llamawasm2go/p2.Fn496
func Fn496(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn498 github.com/goccy/llamawasm2go/p2.Fn498
func Fn498(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn500 github.com/goccy/llamawasm2go/p2.Fn500
func Fn500(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn502 github.com/goccy/llamawasm2go/p2.Fn502
func Fn502(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn503 github.com/goccy/llamawasm2go/p2.Fn503
func Fn503(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn504 github.com/goccy/llamawasm2go/p2.Fn504
func Fn504(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn505 github.com/goccy/llamawasm2go/p2.Fn505
func Fn505(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn506 github.com/goccy/llamawasm2go/p2.Fn506
func Fn506(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn508 github.com/goccy/llamawasm2go/p2.Fn508
func Fn508(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn509 github.com/goccy/llamawasm2go/p2.Fn509
func Fn509(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn510 github.com/goccy/llamawasm2go/p2.Fn510
func Fn510(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn511 github.com/goccy/llamawasm2go/p2.Fn511
func Fn511(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn512 github.com/goccy/llamawasm2go/p2.Fn512
func Fn512(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn513 github.com/goccy/llamawasm2go/p2.Fn513
func Fn513(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn515 github.com/goccy/llamawasm2go/p2.Fn515
func Fn515(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn516 github.com/goccy/llamawasm2go/p2.Fn516
func Fn516(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn517 github.com/goccy/llamawasm2go/p2.Fn517
func Fn517(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn518 github.com/goccy/llamawasm2go/p2.Fn518
func Fn518(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn519 github.com/goccy/llamawasm2go/p2.Fn519
func Fn519(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn520 github.com/goccy/llamawasm2go/p2.Fn520
func Fn520(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn521 github.com/goccy/llamawasm2go/p2.Fn521
func Fn521(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn522 github.com/goccy/llamawasm2go/p2.Fn522
func Fn522(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn523 github.com/goccy/llamawasm2go/p2.Fn523
func Fn523(m *base.Module, l0 int64)

//go:linkname Fn524 github.com/goccy/llamawasm2go/p2.Fn524
func Fn524(m *base.Module, l0 int64)

//go:linkname Fn525 github.com/goccy/llamawasm2go/p2.Fn525
func Fn525(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn526 github.com/goccy/llamawasm2go/p2.Fn526
func Fn526(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn527 github.com/goccy/llamawasm2go/p2.Fn527
func Fn527(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn528 github.com/goccy/llamawasm2go/p2.Fn528
func Fn528(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn529 github.com/goccy/llamawasm2go/p2.Fn529
func Fn529(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn530 github.com/goccy/llamawasm2go/p2.Fn530
func Fn530(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn531 github.com/goccy/llamawasm2go/p2.Fn531
func Fn531(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn532 github.com/goccy/llamawasm2go/p2.Fn532
func Fn532(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn533 github.com/goccy/llamawasm2go/p2.Fn533
func Fn533(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn534 github.com/goccy/llamawasm2go/p2.Fn534
func Fn534(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn536 github.com/goccy/llamawasm2go/p2.Fn536
func Fn536(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn537 github.com/goccy/llamawasm2go/p2.Fn537
func Fn537(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn538 github.com/goccy/llamawasm2go/p2.Fn538
func Fn538(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn539 github.com/goccy/llamawasm2go/p2.Fn539
func Fn539(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn540 github.com/goccy/llamawasm2go/p2.Fn540
func Fn540(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn541 github.com/goccy/llamawasm2go/p2.Fn541
func Fn541(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn542 github.com/goccy/llamawasm2go/p2.Fn542
func Fn542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64) int64

//go:linkname Fn543 github.com/goccy/llamawasm2go/p2.Fn543
func Fn543(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn544 github.com/goccy/llamawasm2go/p2.Fn544
func Fn544(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn545 github.com/goccy/llamawasm2go/p2.Fn545
func Fn545(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn546 github.com/goccy/llamawasm2go/p2.Fn546
func Fn546(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn547 github.com/goccy/llamawasm2go/p2.Fn547
func Fn547(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn549 github.com/goccy/llamawasm2go/p2.Fn549
func Fn549(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn552 github.com/goccy/llamawasm2go/p2.Fn552
func Fn552(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32) int64

//go:linkname Fn553 github.com/goccy/llamawasm2go/p2.Fn553
func Fn553(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32, l7 float32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32) int64

//go:linkname Fn555 github.com/goccy/llamawasm2go/p2.Fn555
func Fn555(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn556 github.com/goccy/llamawasm2go/p2.Fn556
func Fn556(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn557 github.com/goccy/llamawasm2go/p2.Fn557
func Fn557(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn559 github.com/goccy/llamawasm2go/p2.Fn559
func Fn559(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn561 github.com/goccy/llamawasm2go/p2.Fn561
func Fn561(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn562 github.com/goccy/llamawasm2go/p2.Fn562
func Fn562(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn563 github.com/goccy/llamawasm2go/p2.Fn563
func Fn563(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn565 github.com/goccy/llamawasm2go/p2.Fn565
func Fn565(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 float32, l6 float32, l7 float32) int64

//go:linkname Fn566 github.com/goccy/llamawasm2go/p2.Fn566
func Fn566(m *base.Module, l0 int64)

//go:linkname Fn567 github.com/goccy/llamawasm2go/p2.Fn567
func Fn567(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn571 github.com/goccy/llamawasm2go/p2.Fn571
func Fn571(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn572 github.com/goccy/llamawasm2go/p2.Fn572
func Fn572(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn575 github.com/goccy/llamawasm2go/p2.Fn575
func Fn575(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn577 github.com/goccy/llamawasm2go/p2.Fn577
func Fn577(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn579 github.com/goccy/llamawasm2go/p2.Fn579
func Fn579(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn585 github.com/goccy/llamawasm2go/p2.Fn585
func Fn585(m *base.Module, l0 int64)

//go:linkname Fn588 github.com/goccy/llamawasm2go/p2.Fn588
func Fn588(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn589 github.com/goccy/llamawasm2go/p2.Fn589
func Fn589(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn591 github.com/goccy/llamawasm2go/p2.Fn591
func Fn591(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn594 github.com/goccy/llamawasm2go/p2.Fn594
func Fn594(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn595 github.com/goccy/llamawasm2go/p2.Fn595
func Fn595(m *base.Module, l0 int64) int64

//go:linkname Fn596 github.com/goccy/llamawasm2go/p2.Fn596
func Fn596(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn597 github.com/goccy/llamawasm2go/p2.Fn597
func Fn597(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn598 github.com/goccy/llamawasm2go/p2.Fn598
func Fn598(m *base.Module, l0 int64) int64

//go:linkname Fn599 github.com/goccy/llamawasm2go/p2.Fn599
func Fn599(m *base.Module, l0 int64) int64

//go:linkname Fn600 github.com/goccy/llamawasm2go/p2.Fn600
func Fn600(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn602 github.com/goccy/llamawasm2go/p2.Fn602
func Fn602(m *base.Module, l0 int64) int64

//go:linkname Fn603 github.com/goccy/llamawasm2go/p2.Fn603
func Fn603(m *base.Module, l0 int64) int64

//go:linkname Fn605 github.com/goccy/llamawasm2go/p2.Fn605
func Fn605(m *base.Module, l0 int64)

//go:linkname Fn606 github.com/goccy/llamawasm2go/p2.Fn606
func Fn606(m *base.Module, l0 int64) int64

//go:linkname Fn607 github.com/goccy/llamawasm2go/p2.Fn607
func Fn607(m *base.Module, l0 int64) int64

//go:linkname Fn608 github.com/goccy/llamawasm2go/p2.Fn608
func Fn608(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn609 github.com/goccy/llamawasm2go/p2.Fn609
func Fn609(m *base.Module, l0 int64) int32

//go:linkname Fn610 github.com/goccy/llamawasm2go/p2.Fn610
func Fn610(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn612 github.com/goccy/llamawasm2go/p2.Fn612
func Fn612(m *base.Module, l0 int64)

//go:linkname Fn613 github.com/goccy/llamawasm2go/p2.Fn613
func Fn613(m *base.Module, l0 int64)

//go:linkname Fn614 github.com/goccy/llamawasm2go/p2.Fn614
func Fn614(m *base.Module, l0 int64) int64

//go:linkname Fn616 github.com/goccy/llamawasm2go/p2.Fn616
func Fn616(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn617 github.com/goccy/llamawasm2go/p2.Fn617
func Fn617(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn618 github.com/goccy/llamawasm2go/p2.Fn618
func Fn618(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn619 github.com/goccy/llamawasm2go/p2.Fn619
func Fn619(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn620 github.com/goccy/llamawasm2go/p2.Fn620
func Fn620(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn621 github.com/goccy/llamawasm2go/p2.Fn621
func Fn621(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn624 github.com/goccy/llamawasm2go/p2.Fn624
func Fn624(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn625 github.com/goccy/llamawasm2go/p2.Fn625
func Fn625(m *base.Module, l0 int64) int64

//go:linkname Fn626 github.com/goccy/llamawasm2go/p2.Fn626
func Fn626(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn629 github.com/goccy/llamawasm2go/p2.Fn629
func Fn629(m *base.Module, l0 int64) int64

//go:linkname Fn630 github.com/goccy/llamawasm2go/p2.Fn630
func Fn630(m *base.Module, l0 int64) int64

//go:linkname Fn634 github.com/goccy/llamawasm2go/p2.Fn634
func Fn634(m *base.Module, l0 int64) int64

//go:linkname Fn635 github.com/goccy/llamawasm2go/p2.Fn635
func Fn635(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn636 github.com/goccy/llamawasm2go/p2.Fn636
func Fn636(m *base.Module, l0 int64) int64

//go:linkname Fn640 github.com/goccy/llamawasm2go/p2.Fn640
func Fn640(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn641 github.com/goccy/llamawasm2go/p0.Fn641
func Fn641(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn646 github.com/goccy/llamawasm2go/p2.Fn646
func Fn646(m *base.Module, l0 int64)

//go:linkname Fn647 github.com/goccy/llamawasm2go/p2.Fn647
func Fn647(m *base.Module, l0 int64)

//go:linkname Fn648 github.com/goccy/llamawasm2go/p2.Fn648
func Fn648(m *base.Module, l0 int64)

//go:linkname Fn649 github.com/goccy/llamawasm2go/p2.Fn649
func Fn649(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn650 github.com/goccy/llamawasm2go/p2.Fn650
func Fn650(m *base.Module, l0 int64) int32

//go:linkname Fn651 github.com/goccy/llamawasm2go/p2.Fn651
func Fn651(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn652 github.com/goccy/llamawasm2go/p2.Fn652
func Fn652(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn653 github.com/goccy/llamawasm2go/p2.Fn653
func Fn653(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn654 github.com/goccy/llamawasm2go/p2.Fn654
func Fn654(m *base.Module, l0 int64) int32

//go:linkname Fn655 github.com/goccy/llamawasm2go/p2.Fn655
func Fn655(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn673 github.com/goccy/llamawasm2go/p2.Fn673
func Fn673(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn679 github.com/goccy/llamawasm2go/p2.Fn679
func Fn679(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn680 github.com/goccy/llamawasm2go/p2.Fn680
func Fn680(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn681 github.com/goccy/llamawasm2go/p2.Fn681
func Fn681(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn683 github.com/goccy/llamawasm2go/p2.Fn683
func Fn683(m *base.Module, l0 int64)

//go:linkname Fn685 github.com/goccy/llamawasm2go/p2.Fn685
func Fn685(m *base.Module, l0 int64)

//go:linkname Fn686 github.com/goccy/llamawasm2go/p0.Fn686
func Fn686(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn688 github.com/goccy/llamawasm2go/p2.Fn688
func Fn688(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn702 github.com/goccy/llamawasm2go/p2.Fn702
func Fn702(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn704 github.com/goccy/llamawasm2go/p2.Fn704
func Fn704(m *base.Module, l0 int64)

//go:linkname Fn705 github.com/goccy/llamawasm2go/p2.Fn705
func Fn705(m *base.Module, l0 int64) int64

//go:linkname Fn728 github.com/goccy/llamawasm2go/p2.Fn728
func Fn728(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn739 github.com/goccy/llamawasm2go/p2.Fn739
func Fn739(m *base.Module, l0 int64)

//go:linkname Fn758 github.com/goccy/llamawasm2go/p2.Fn758
func Fn758(m *base.Module, l0 int64)

//go:linkname Fn759 github.com/goccy/llamawasm2go/p2.Fn759
func Fn759(m *base.Module, l0 int64)

//go:linkname Fn829 github.com/goccy/llamawasm2go/p2.Fn829
func Fn829(m *base.Module, l0 int64) int64

//go:linkname Fn890 github.com/goccy/llamawasm2go/p0.Fn890
func Fn890(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn938 github.com/goccy/llamawasm2go/p2.Fn938
func Fn938(m *base.Module) int64

//go:linkname Fn975 github.com/goccy/llamawasm2go/p2.Fn975
func Fn975(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32)

//go:linkname Fn985 github.com/goccy/llamawasm2go/p2.Fn985
func Fn985(m *base.Module, l0 int64)

//go:linkname Fn1010 github.com/goccy/llamawasm2go/p2.Fn1010
func Fn1010(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn1015 github.com/goccy/llamawasm2go/p2.Fn1015
func Fn1015(m *base.Module, l0 int64) int64

//go:linkname Fn1025 github.com/goccy/llamawasm2go/p2.Fn1025
func Fn1025(m *base.Module, l0 int64)

//go:linkname Fn1030 github.com/goccy/llamawasm2go/p2.Fn1030
func Fn1030(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1031 github.com/goccy/llamawasm2go/p2.Fn1031
func Fn1031(m *base.Module)

//go:linkname Fn1040 github.com/goccy/llamawasm2go/p2.Fn1040
func Fn1040(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1048 github.com/goccy/llamawasm2go/p2.Fn1048
func Fn1048(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1049 github.com/goccy/llamawasm2go/p2.Fn1049
func Fn1049(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1050 github.com/goccy/llamawasm2go/p2.Fn1050
func Fn1050(m *base.Module, l0 int64, l1 int64, l2 int64) float32

//go:linkname Fn1052 github.com/goccy/llamawasm2go/p2.Fn1052
func Fn1052(m *base.Module, l0 int64, l1 int64, l2 int64) float64

//go:linkname Fn1054 github.com/goccy/llamawasm2go/p2.Fn1054
func Fn1054(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1057 github.com/goccy/llamawasm2go/p2.Fn1057
func Fn1057(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1060 github.com/goccy/llamawasm2go/p2.Fn1060
func Fn1060(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1066 github.com/goccy/llamawasm2go/p2.Fn1066
func Fn1066(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1067 github.com/goccy/llamawasm2go/p2.Fn1067
func Fn1067(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1075 github.com/goccy/llamawasm2go/p2.Fn1075
func Fn1075(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1080 github.com/goccy/llamawasm2go/p2.Fn1080
func Fn1080(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1081 github.com/goccy/llamawasm2go/p2.Fn1081
func Fn1081(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1089 github.com/goccy/llamawasm2go/p2.Fn1089
func Fn1089(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1105 github.com/goccy/llamawasm2go/p2.Fn1105
func Fn1105(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1113 github.com/goccy/llamawasm2go/p2.Fn1113
func Fn1113(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1129 github.com/goccy/llamawasm2go/p0.Fn1129
func Fn1129(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1132 github.com/goccy/llamawasm2go/p2.Fn1132
func Fn1132(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1136 github.com/goccy/llamawasm2go/p2.Fn1136
func Fn1136(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1146 github.com/goccy/llamawasm2go/p0.Fn1146
func Fn1146(m *base.Module, l0 int64) int64

//go:linkname Fn1147 github.com/goccy/llamawasm2go/p2.Fn1147
func Fn1147(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1148 github.com/goccy/llamawasm2go/p2.Fn1148
func Fn1148(m *base.Module, l0 int64)

//go:linkname Fn1150 github.com/goccy/llamawasm2go/p2.Fn1150
func Fn1150(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1240 github.com/goccy/llamawasm2go/p2.Fn1240
func Fn1240(m *base.Module, l0 int64)

//go:linkname Fn1262 github.com/goccy/llamawasm2go/p2.Fn1262
func Fn1262(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1272 github.com/goccy/llamawasm2go/p2.Fn1272
func Fn1272(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1278 github.com/goccy/llamawasm2go/p2.Fn1278
func Fn1278(m *base.Module)

//go:linkname Fn1282 github.com/goccy/llamawasm2go/p2.Fn1282
func Fn1282(m *base.Module, l0 int64) int64

//go:linkname Fn1310 github.com/goccy/llamawasm2go/p2.Fn1310
func Fn1310(m *base.Module, l0 int64)

//go:linkname Fn1314 github.com/goccy/llamawasm2go/p2.Fn1314
func Fn1314(m *base.Module, l0 int32) int64

//go:linkname Fn1325 github.com/goccy/llamawasm2go/p2.Fn1325
func Fn1325(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1326 github.com/goccy/llamawasm2go/p2.Fn1326
func Fn1326(m *base.Module, l0 int64)

//go:linkname Fn1328 github.com/goccy/llamawasm2go/p2.Fn1328
func Fn1328(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1350 github.com/goccy/llamawasm2go/p2.Fn1350
func Fn1350(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1351 github.com/goccy/llamawasm2go/p2.Fn1351
func Fn1351(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1352 github.com/goccy/llamawasm2go/p2.Fn1352
func Fn1352(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1353 github.com/goccy/llamawasm2go/p2.Fn1353
func Fn1353(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1354 github.com/goccy/llamawasm2go/p2.Fn1354
func Fn1354(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1355 github.com/goccy/llamawasm2go/p2.Fn1355
func Fn1355(m *base.Module, l0 int64) int64

//go:linkname Fn1358 github.com/goccy/llamawasm2go/p2.Fn1358
func Fn1358(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1362 github.com/goccy/llamawasm2go/p2.Fn1362
func Fn1362(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1363 github.com/goccy/llamawasm2go/p2.Fn1363
func Fn1363(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1364 github.com/goccy/llamawasm2go/p2.Fn1364
func Fn1364(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1365 github.com/goccy/llamawasm2go/p2.Fn1365
func Fn1365(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1366 github.com/goccy/llamawasm2go/p2.Fn1366
func Fn1366(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1367 github.com/goccy/llamawasm2go/p2.Fn1367
func Fn1367(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1368 github.com/goccy/llamawasm2go/p2.Fn1368
func Fn1368(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1369 github.com/goccy/llamawasm2go/p2.Fn1369
func Fn1369(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1370 github.com/goccy/llamawasm2go/p2.Fn1370
func Fn1370(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1371 github.com/goccy/llamawasm2go/p2.Fn1371
func Fn1371(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1372 github.com/goccy/llamawasm2go/p2.Fn1372
func Fn1372(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1373 github.com/goccy/llamawasm2go/p2.Fn1373
func Fn1373(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1374 github.com/goccy/llamawasm2go/p2.Fn1374
func Fn1374(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1375 github.com/goccy/llamawasm2go/p2.Fn1375
func Fn1375(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1376 github.com/goccy/llamawasm2go/p2.Fn1376
func Fn1376(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1377 github.com/goccy/llamawasm2go/p2.Fn1377
func Fn1377(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1378 github.com/goccy/llamawasm2go/p2.Fn1378
func Fn1378(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1379 github.com/goccy/llamawasm2go/p2.Fn1379
func Fn1379(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1380 github.com/goccy/llamawasm2go/p2.Fn1380
func Fn1380(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1381 github.com/goccy/llamawasm2go/p2.Fn1381
func Fn1381(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1382 github.com/goccy/llamawasm2go/p2.Fn1382
func Fn1382(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1383 github.com/goccy/llamawasm2go/p2.Fn1383
func Fn1383(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1384 github.com/goccy/llamawasm2go/p2.Fn1384
func Fn1384(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1385 github.com/goccy/llamawasm2go/p2.Fn1385
func Fn1385(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1386 github.com/goccy/llamawasm2go/p2.Fn1386
func Fn1386(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1387 github.com/goccy/llamawasm2go/p2.Fn1387
func Fn1387(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1388 github.com/goccy/llamawasm2go/p2.Fn1388
func Fn1388(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1390 github.com/goccy/llamawasm2go/p2.Fn1390
func Fn1390(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p2.Fn1392
func Fn1392(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1393 github.com/goccy/llamawasm2go/p2.Fn1393
func Fn1393(m *base.Module)

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p2.Fn1394
func Fn1394(m *base.Module)

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p0.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module)

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p0.Fn1407
func Fn1407(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p2.Fn1408
func Fn1408(m *base.Module, l0 int64) int64

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p2.Fn1413
func Fn1413(m *base.Module, l0 int64)

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1423 github.com/goccy/llamawasm2go/p2.Fn1423
func Fn1423(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1425 github.com/goccy/llamawasm2go/p2.Fn1425
func Fn1425(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1427 github.com/goccy/llamawasm2go/p2.Fn1427
func Fn1427(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1429 github.com/goccy/llamawasm2go/p2.Fn1429
func Fn1429(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1431 github.com/goccy/llamawasm2go/p2.Fn1431
func Fn1431(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1433 github.com/goccy/llamawasm2go/p2.Fn1433
func Fn1433(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1439 github.com/goccy/llamawasm2go/p2.Fn1439
func Fn1439(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1440 github.com/goccy/llamawasm2go/p2.Fn1440
func Fn1440(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1441 github.com/goccy/llamawasm2go/p2.Fn1441
func Fn1441(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p0.Fn1442
func Fn1442(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64)

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int64) int64

//go:linkname Fn1448 github.com/goccy/llamawasm2go/p2.Fn1448
func Fn1448(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1449 github.com/goccy/llamawasm2go/p2.Fn1449
func Fn1449(m *base.Module, l0 int64)

//go:linkname Fn1450 github.com/goccy/llamawasm2go/p2.Fn1450
func Fn1450(m *base.Module, l0 int64)

//go:linkname Fn1451 github.com/goccy/llamawasm2go/p2.Fn1451
func Fn1451(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1452 github.com/goccy/llamawasm2go/p2.Fn1452
func Fn1452(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1455 github.com/goccy/llamawasm2go/p2.Fn1455
func Fn1455(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1463 github.com/goccy/llamawasm2go/p2.Fn1463
func Fn1463(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1465 github.com/goccy/llamawasm2go/p0.Fn1465
func Fn1465(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1467 github.com/goccy/llamawasm2go/p2.Fn1467
func Fn1467(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1470 github.com/goccy/llamawasm2go/p0.Fn1470
func Fn1470(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1471 github.com/goccy/llamawasm2go/p2.Fn1471
func Fn1471(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1475 github.com/goccy/llamawasm2go/p2.Fn1475
func Fn1475(m *base.Module, l0 int64)

//go:linkname Fn1478 github.com/goccy/llamawasm2go/p2.Fn1478
func Fn1478(m *base.Module, l0 int64)

//go:linkname Fn1481 github.com/goccy/llamawasm2go/p2.Fn1481
func Fn1481(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1482 github.com/goccy/llamawasm2go/p2.Fn1482
func Fn1482(m *base.Module, l0 int64) int64

//go:linkname Fn1483 github.com/goccy/llamawasm2go/p2.Fn1483
func Fn1483(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1484 github.com/goccy/llamawasm2go/p2.Fn1484
func Fn1484(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1485 github.com/goccy/llamawasm2go/p2.Fn1485
func Fn1485(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1486 github.com/goccy/llamawasm2go/p2.Fn1486
func Fn1486(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1487 github.com/goccy/llamawasm2go/p2.Fn1487
func Fn1487(m *base.Module, l0 int64)

//go:linkname Fn1490 github.com/goccy/llamawasm2go/p2.Fn1490
func Fn1490(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1492 github.com/goccy/llamawasm2go/p2.Fn1492
func Fn1492(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1495 github.com/goccy/llamawasm2go/p2.Fn1495
func Fn1495(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1502 github.com/goccy/llamawasm2go/p2.Fn1502
func Fn1502(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1503 github.com/goccy/llamawasm2go/p2.Fn1503
func Fn1503(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1505 github.com/goccy/llamawasm2go/p2.Fn1505
func Fn1505(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1509 github.com/goccy/llamawasm2go/p2.Fn1509
func Fn1509(m *base.Module, l0 int64)

//go:linkname Fn1510 github.com/goccy/llamawasm2go/p2.Fn1510
func Fn1510(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1513 github.com/goccy/llamawasm2go/p2.Fn1513
func Fn1513(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1515 github.com/goccy/llamawasm2go/p2.Fn1515
func Fn1515(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1516 github.com/goccy/llamawasm2go/p2.Fn1516
func Fn1516(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1517 github.com/goccy/llamawasm2go/p2.Fn1517
func Fn1517(m *base.Module, l0 int64)

//go:linkname Fn1518 github.com/goccy/llamawasm2go/p2.Fn1518
func Fn1518(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1520 github.com/goccy/llamawasm2go/p2.Fn1520
func Fn1520(m *base.Module, l0 int64) int32

//go:linkname Fn1521 github.com/goccy/llamawasm2go/p2.Fn1521
func Fn1521(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p2.Fn1563
func Fn1563(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1579 github.com/goccy/llamawasm2go/p2.Fn1579
func Fn1579(m *base.Module, l0 int64)

//go:linkname Fn1580 github.com/goccy/llamawasm2go/p2.Fn1580
func Fn1580(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1581 github.com/goccy/llamawasm2go/p2.Fn1581
func Fn1581(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1582 github.com/goccy/llamawasm2go/p2.Fn1582
func Fn1582(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p2.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1584 github.com/goccy/llamawasm2go/p2.Fn1584
func Fn1584(m *base.Module, l0 int64)

//go:linkname Fn1585 github.com/goccy/llamawasm2go/p2.Fn1585
func Fn1585(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1587 github.com/goccy/llamawasm2go/p2.Fn1587
func Fn1587(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1588 github.com/goccy/llamawasm2go/p2.Fn1588
func Fn1588(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1589 github.com/goccy/llamawasm2go/p2.Fn1589
func Fn1589(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1590 github.com/goccy/llamawasm2go/p2.Fn1590
func Fn1590(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1591 github.com/goccy/llamawasm2go/p2.Fn1591
func Fn1591(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1594 github.com/goccy/llamawasm2go/p2.Fn1594
func Fn1594(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1596 github.com/goccy/llamawasm2go/p2.Fn1596
func Fn1596(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1597 github.com/goccy/llamawasm2go/p2.Fn1597
func Fn1597(m *base.Module, l0 int64) int64

//go:linkname Fn1598 github.com/goccy/llamawasm2go/p2.Fn1598
func Fn1598(m *base.Module, l0 int64) int64

//go:linkname Fn1599 github.com/goccy/llamawasm2go/p2.Fn1599
func Fn1599(m *base.Module, l0 int64) int64

//go:linkname Fn1600 github.com/goccy/llamawasm2go/p2.Fn1600
func Fn1600(m *base.Module, l0 int64) int64

//go:linkname Fn1601 github.com/goccy/llamawasm2go/p2.Fn1601
func Fn1601(m *base.Module, l0 int64) int64

//go:linkname Fn1602 github.com/goccy/llamawasm2go/p2.Fn1602
func Fn1602(m *base.Module, l0 int64) int64

//go:linkname Fn1604 github.com/goccy/llamawasm2go/p2.Fn1604
func Fn1604(m *base.Module, l0 int64) int64

//go:linkname Fn1605 github.com/goccy/llamawasm2go/p2.Fn1605
func Fn1605(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1607 github.com/goccy/llamawasm2go/p2.Fn1607
func Fn1607(m *base.Module, l0 int64) int64

//go:linkname Fn1608 github.com/goccy/llamawasm2go/p2.Fn1608
func Fn1608(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1609 github.com/goccy/llamawasm2go/p2.Fn1609
func Fn1609(m *base.Module, l0 int64) int64

//go:linkname Fn1611 github.com/goccy/llamawasm2go/p2.Fn1611
func Fn1611(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1612 github.com/goccy/llamawasm2go/p2.Fn1612
func Fn1612(m *base.Module, l0 int64) int64

//go:linkname Fn1614 github.com/goccy/llamawasm2go/p2.Fn1614
func Fn1614(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1615 github.com/goccy/llamawasm2go/p2.Fn1615
func Fn1615(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1616 github.com/goccy/llamawasm2go/p2.Fn1616
func Fn1616(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1617 github.com/goccy/llamawasm2go/p2.Fn1617
func Fn1617(m *base.Module, l0 int64) int64

//go:linkname Fn1618 github.com/goccy/llamawasm2go/p2.Fn1618
func Fn1618(m *base.Module, l0 int64) int64

//go:linkname Fn1619 github.com/goccy/llamawasm2go/p2.Fn1619
func Fn1619(m *base.Module, l0 int64) int64

//go:linkname Fn1625 github.com/goccy/llamawasm2go/p2.Fn1625
func Fn1625(m *base.Module, l0 int64) int64

//go:linkname Fn1627 github.com/goccy/llamawasm2go/p2.Fn1627
func Fn1627(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1628 github.com/goccy/llamawasm2go/p2.Fn1628
func Fn1628(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1630 github.com/goccy/llamawasm2go/p2.Fn1630
func Fn1630(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1631 github.com/goccy/llamawasm2go/p2.Fn1631
func Fn1631(m *base.Module, l0 int64) int64

//go:linkname Fn1632 github.com/goccy/llamawasm2go/p2.Fn1632
func Fn1632(m *base.Module, l0 int64) int64

//go:linkname Fn1633 github.com/goccy/llamawasm2go/p2.Fn1633
func Fn1633(m *base.Module, l0 int64) int64

//go:linkname Fn1634 github.com/goccy/llamawasm2go/p2.Fn1634
func Fn1634(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1635 github.com/goccy/llamawasm2go/p2.Fn1635
func Fn1635(m *base.Module, l0 int64)

//go:linkname Fn1664 github.com/goccy/llamawasm2go/p2.Fn1664
func Fn1664(m *base.Module, l0 int64) int64

//go:linkname Fn1671 github.com/goccy/llamawasm2go/p2.Fn1671
func Fn1671(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1672 github.com/goccy/llamawasm2go/p2.Fn1672
func Fn1672(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1673 github.com/goccy/llamawasm2go/p2.Fn1673
func Fn1673(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1675 github.com/goccy/llamawasm2go/p2.Fn1675
func Fn1675(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1676 github.com/goccy/llamawasm2go/p2.Fn1676
func Fn1676(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1680 github.com/goccy/llamawasm2go/p2.Fn1680
func Fn1680(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1681 github.com/goccy/llamawasm2go/p2.Fn1681
func Fn1681(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1682 github.com/goccy/llamawasm2go/p2.Fn1682
func Fn1682(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1683 github.com/goccy/llamawasm2go/p2.Fn1683
func Fn1683(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1684 github.com/goccy/llamawasm2go/p2.Fn1684
func Fn1684(m *base.Module, l0 int64) int32

//go:linkname Fn1685 github.com/goccy/llamawasm2go/p2.Fn1685
func Fn1685(m *base.Module, l0 int64) int32

//go:linkname Fn1686 github.com/goccy/llamawasm2go/p2.Fn1686
func Fn1686(m *base.Module, l0 int64) int32

//go:linkname Fn1687 github.com/goccy/llamawasm2go/p2.Fn1687
func Fn1687(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1690 github.com/goccy/llamawasm2go/p2.Fn1690
func Fn1690(m *base.Module, l0 int64) int32

//go:linkname Fn1691 github.com/goccy/llamawasm2go/p2.Fn1691
func Fn1691(m *base.Module, l0 int64) int32

//go:linkname Fn1697 github.com/goccy/llamawasm2go/p2.Fn1697
func Fn1697(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1698 github.com/goccy/llamawasm2go/p2.Fn1698
func Fn1698(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1699 github.com/goccy/llamawasm2go/p2.Fn1699
func Fn1699(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1700 github.com/goccy/llamawasm2go/p2.Fn1700
func Fn1700(m *base.Module)

//go:linkname Fn1701 github.com/goccy/llamawasm2go/p2.Fn1701
func Fn1701(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1703 github.com/goccy/llamawasm2go/p2.Fn1703
func Fn1703(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1705 github.com/goccy/llamawasm2go/p2.Fn1705
func Fn1705(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1706 github.com/goccy/llamawasm2go/p2.Fn1706
func Fn1706(m *base.Module, l0 int64)

//go:linkname Fn1710 github.com/goccy/llamawasm2go/p2.Fn1710
func Fn1710(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1711 github.com/goccy/llamawasm2go/p2.Fn1711
func Fn1711(m *base.Module, l0 int64)

//go:linkname Fn1714 github.com/goccy/llamawasm2go/p2.Fn1714
func Fn1714(m *base.Module, l0 int64)

//go:linkname Fn1724 github.com/goccy/llamawasm2go/p2.Fn1724
func Fn1724(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1725 github.com/goccy/llamawasm2go/p2.Fn1725
func Fn1725(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1734 github.com/goccy/llamawasm2go/p2.Fn1734
func Fn1734(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1735 github.com/goccy/llamawasm2go/p0.Fn1735
func Fn1735(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1737 github.com/goccy/llamawasm2go/p2.Fn1737
func Fn1737(m *base.Module, l0 int64)

//go:linkname Fn1739 github.com/goccy/llamawasm2go/p2.Fn1739
func Fn1739(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1746 github.com/goccy/llamawasm2go/p2.Fn1746
func Fn1746(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1747 github.com/goccy/llamawasm2go/p2.Fn1747
func Fn1747(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1754 github.com/goccy/llamawasm2go/p2.Fn1754
func Fn1754(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1761 github.com/goccy/llamawasm2go/p2.Fn1761
func Fn1761(m *base.Module, l0 int64)

//go:linkname Fn1764 github.com/goccy/llamawasm2go/p2.Fn1764
func Fn1764(m *base.Module, l0 int64) int32

//go:linkname Fn1774 github.com/goccy/llamawasm2go/p2.Fn1774
func Fn1774(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1776 github.com/goccy/llamawasm2go/p2.Fn1776
func Fn1776(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1777 github.com/goccy/llamawasm2go/p2.Fn1777
func Fn1777(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1796 github.com/goccy/llamawasm2go/p2.Fn1796
func Fn1796(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1809 github.com/goccy/llamawasm2go/p2.Fn1809
func Fn1809(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1816 github.com/goccy/llamawasm2go/p2.Fn1816
func Fn1816(m *base.Module, l0 int64)

//go:linkname Fn1850 github.com/goccy/llamawasm2go/p2.Fn1850
func Fn1850(m *base.Module, l0 int64)

//go:linkname Fn1853 github.com/goccy/llamawasm2go/p2.Fn1853
func Fn1853(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1867 github.com/goccy/llamawasm2go/p2.Fn1867
func Fn1867(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1868 github.com/goccy/llamawasm2go/p2.Fn1868
func Fn1868(m *base.Module, l0 int64) int64

//go:linkname Fn1869 github.com/goccy/llamawasm2go/p2.Fn1869
func Fn1869(m *base.Module, l0 int64)

//go:linkname Fn1873 github.com/goccy/llamawasm2go/p0.Fn1873
func Fn1873(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1874 github.com/goccy/llamawasm2go/p2.Fn1874
func Fn1874(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1875 github.com/goccy/llamawasm2go/p2.Fn1875
func Fn1875(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1876 github.com/goccy/llamawasm2go/p2.Fn1876
func Fn1876(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1877 github.com/goccy/llamawasm2go/p2.Fn1877
func Fn1877(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1895 github.com/goccy/llamawasm2go/p2.Fn1895
func Fn1895(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1906 github.com/goccy/llamawasm2go/p2.Fn1906
func Fn1906(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1908 github.com/goccy/llamawasm2go/p2.Fn1908
func Fn1908(m *base.Module, l0 int64) int64

//go:linkname Fn1909 github.com/goccy/llamawasm2go/p2.Fn1909
func Fn1909(m *base.Module, l0 int64)

//go:linkname Fn1912 github.com/goccy/llamawasm2go/p0.Fn1912
func Fn1912(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1914 github.com/goccy/llamawasm2go/p2.Fn1914
func Fn1914(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1925 github.com/goccy/llamawasm2go/p2.Fn1925
func Fn1925(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1928 github.com/goccy/llamawasm2go/p2.Fn1928
func Fn1928(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1929 github.com/goccy/llamawasm2go/p2.Fn1929
func Fn1929(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1939 github.com/goccy/llamawasm2go/p2.Fn1939
func Fn1939(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1966 github.com/goccy/llamawasm2go/p2.Fn1966
func Fn1966(m *base.Module, l0 int64)

//go:linkname Fn1967 github.com/goccy/llamawasm2go/p2.Fn1967
func Fn1967(m *base.Module, l0 int64)

//go:linkname Fn1978 github.com/goccy/llamawasm2go/p2.Fn1978
func Fn1978(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2010 github.com/goccy/llamawasm2go/p2.Fn2010
func Fn2010(m *base.Module) int64

//go:linkname Fn2015 github.com/goccy/llamawasm2go/p2.Fn2015
func Fn2015(m *base.Module, l0 int64) int64

//go:linkname Fn2016 github.com/goccy/llamawasm2go/p2.Fn2016
func Fn2016(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2023 github.com/goccy/llamawasm2go/p2.Fn2023
func Fn2023(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2030 github.com/goccy/llamawasm2go/p2.Fn2030
func Fn2030(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2037 github.com/goccy/llamawasm2go/p2.Fn2037
func Fn2037(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2044 github.com/goccy/llamawasm2go/p2.Fn2044
func Fn2044(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2046 github.com/goccy/llamawasm2go/p2.Fn2046
func Fn2046(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2049 github.com/goccy/llamawasm2go/p2.Fn2049
func Fn2049(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2050 github.com/goccy/llamawasm2go/p2.Fn2050
func Fn2050(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2051 github.com/goccy/llamawasm2go/p2.Fn2051
func Fn2051(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn2054 github.com/goccy/llamawasm2go/p2.Fn2054
func Fn2054(m *base.Module, l0 int64)

//go:linkname Fn2064 github.com/goccy/llamawasm2go/p2.Fn2064
func Fn2064(m *base.Module, l0 int64)

//go:linkname Fn2066 github.com/goccy/llamawasm2go/p2.Fn2066
func Fn2066(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2067 github.com/goccy/llamawasm2go/p2.Fn2067
func Fn2067(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2080 github.com/goccy/llamawasm2go/p2.Fn2080
func Fn2080(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2081 github.com/goccy/llamawasm2go/p2.Fn2081
func Fn2081(m *base.Module, l0 int64) int64

//go:linkname Fn2082 github.com/goccy/llamawasm2go/p2.Fn2082
func Fn2082(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2084 github.com/goccy/llamawasm2go/p2.Fn2084
func Fn2084(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2098 github.com/goccy/llamawasm2go/p2.Fn2098
func Fn2098(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2099 github.com/goccy/llamawasm2go/p2.Fn2099
func Fn2099(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2101 github.com/goccy/llamawasm2go/p2.Fn2101
func Fn2101(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2102 github.com/goccy/llamawasm2go/p2.Fn2102
func Fn2102(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2103 github.com/goccy/llamawasm2go/p2.Fn2103
func Fn2103(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2104 github.com/goccy/llamawasm2go/p2.Fn2104
func Fn2104(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2105 github.com/goccy/llamawasm2go/p2.Fn2105
func Fn2105(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2106 github.com/goccy/llamawasm2go/p2.Fn2106
func Fn2106(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2125 github.com/goccy/llamawasm2go/p2.Fn2125
func Fn2125(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2126 github.com/goccy/llamawasm2go/p2.Fn2126
func Fn2126(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2127 github.com/goccy/llamawasm2go/p2.Fn2127
func Fn2127(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2129 github.com/goccy/llamawasm2go/p2.Fn2129
func Fn2129(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2130 github.com/goccy/llamawasm2go/p2.Fn2130
func Fn2130(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2131 github.com/goccy/llamawasm2go/p2.Fn2131
func Fn2131(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2132 github.com/goccy/llamawasm2go/p2.Fn2132
func Fn2132(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2133 github.com/goccy/llamawasm2go/p2.Fn2133
func Fn2133(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2135 github.com/goccy/llamawasm2go/p2.Fn2135
func Fn2135(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2137 github.com/goccy/llamawasm2go/p2.Fn2137
func Fn2137(m *base.Module, l0 int64)

//go:linkname Fn2154 github.com/goccy/llamawasm2go/p2.Fn2154
func Fn2154(m *base.Module, l0 int64)

//go:linkname Fn2155 github.com/goccy/llamawasm2go/p2.Fn2155
func Fn2155(m *base.Module, l0 int64)

//go:linkname Fn2156 github.com/goccy/llamawasm2go/p2.Fn2156
func Fn2156(m *base.Module, l0 int64)

//go:linkname Fn2158 github.com/goccy/llamawasm2go/p2.Fn2158
func Fn2158(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2159 github.com/goccy/llamawasm2go/p2.Fn2159
func Fn2159(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2196 github.com/goccy/llamawasm2go/p2.Fn2196
func Fn2196(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2205 github.com/goccy/llamawasm2go/p2.Fn2205
func Fn2205(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2207 github.com/goccy/llamawasm2go/p2.Fn2207
func Fn2207(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2209 github.com/goccy/llamawasm2go/p2.Fn2209
func Fn2209(m *base.Module, l0 int64) int64

//go:linkname Fn2212 github.com/goccy/llamawasm2go/p2.Fn2212
func Fn2212(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2216 github.com/goccy/llamawasm2go/p2.Fn2216
func Fn2216(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2221 github.com/goccy/llamawasm2go/p2.Fn2221
func Fn2221(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2235 github.com/goccy/llamawasm2go/p2.Fn2235
func Fn2235(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2237 github.com/goccy/llamawasm2go/p2.Fn2237
func Fn2237(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2238 github.com/goccy/llamawasm2go/p2.Fn2238
func Fn2238(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2252 github.com/goccy/llamawasm2go/p2.Fn2252
func Fn2252(m *base.Module, l0 int64) int64

//go:linkname Fn2254 github.com/goccy/llamawasm2go/p2.Fn2254
func Fn2254(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2257 github.com/goccy/llamawasm2go/p2.Fn2257
func Fn2257(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2262 github.com/goccy/llamawasm2go/p2.Fn2262
func Fn2262(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2271 github.com/goccy/llamawasm2go/p2.Fn2271
func Fn2271(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2275 github.com/goccy/llamawasm2go/p2.Fn2275
func Fn2275(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2276 github.com/goccy/llamawasm2go/p2.Fn2276
func Fn2276(m *base.Module, l0 int64)

//go:linkname Fn2277 github.com/goccy/llamawasm2go/p2.Fn2277
func Fn2277(m *base.Module, l0 int64)

//go:linkname Fn2278 github.com/goccy/llamawasm2go/p2.Fn2278
func Fn2278(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2280 github.com/goccy/llamawasm2go/p2.Fn2280
func Fn2280(m *base.Module, l0 int64)

//go:linkname Fn2281 github.com/goccy/llamawasm2go/p2.Fn2281
func Fn2281(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2282 github.com/goccy/llamawasm2go/p2.Fn2282
func Fn2282(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2286 github.com/goccy/llamawasm2go/p2.Fn2286
func Fn2286(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2289 github.com/goccy/llamawasm2go/p2.Fn2289
func Fn2289(m *base.Module, l0 int64) int64

//go:linkname Fn2290 github.com/goccy/llamawasm2go/p2.Fn2290
func Fn2290(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2291 github.com/goccy/llamawasm2go/p2.Fn2291
func Fn2291(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2298 github.com/goccy/llamawasm2go/p2.Fn2298
func Fn2298(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2299 github.com/goccy/llamawasm2go/p2.Fn2299
func Fn2299(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2300 github.com/goccy/llamawasm2go/p2.Fn2300
func Fn2300(m *base.Module, l0 int64)

//go:linkname Fn2301 github.com/goccy/llamawasm2go/p2.Fn2301
func Fn2301(m *base.Module) int64

//go:linkname Fn2303 github.com/goccy/llamawasm2go/p2.Fn2303
func Fn2303(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2304 github.com/goccy/llamawasm2go/p2.Fn2304
func Fn2304(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2306 github.com/goccy/llamawasm2go/p2.Fn2306
func Fn2306(m *base.Module) int64

//go:linkname Fn2308 github.com/goccy/llamawasm2go/p2.Fn2308
func Fn2308(m *base.Module, l0 int32) int64

//go:linkname Fn2309 github.com/goccy/llamawasm2go/p2.Fn2309
func Fn2309(m *base.Module, l0 int32) int32

//go:linkname Fn2310 github.com/goccy/llamawasm2go/p2.Fn2310
func Fn2310(m *base.Module, l0 int32) int64

//go:linkname Fn2311 github.com/goccy/llamawasm2go/p2.Fn2311
func Fn2311(m *base.Module, l0 float32) int64

//go:linkname Fn2312 github.com/goccy/llamawasm2go/p2.Fn2312
func Fn2312(m *base.Module, l0 float32) int64

//go:linkname Fn2313 github.com/goccy/llamawasm2go/p2.Fn2313
func Fn2313(m *base.Module, l0 float32) int64

//go:linkname Fn2315 github.com/goccy/llamawasm2go/p2.Fn2315
func Fn2315(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2316 github.com/goccy/llamawasm2go/p2.Fn2316
func Fn2316(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2357 github.com/goccy/llamawasm2go/p2.Fn2357
func Fn2357(m *base.Module, l0 int64)

//go:linkname Fn2359 github.com/goccy/llamawasm2go/p2.Fn2359
func Fn2359(m *base.Module, l0 int64)

//go:linkname Fn2367 github.com/goccy/llamawasm2go/p2.Fn2367
func Fn2367(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2401 github.com/goccy/llamawasm2go/p2.Fn2401
func Fn2401(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2412 github.com/goccy/llamawasm2go/p2.Fn2412
func Fn2412(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2415 github.com/goccy/llamawasm2go/p2.Fn2415
func Fn2415(m *base.Module, l0 int64)

//go:linkname Fn2416 github.com/goccy/llamawasm2go/p2.Fn2416
func Fn2416(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2417 github.com/goccy/llamawasm2go/p2.Fn2417
func Fn2417(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2419 github.com/goccy/llamawasm2go/p2.Fn2419
func Fn2419(m *base.Module, l0 int64)

//go:linkname Fn2423 github.com/goccy/llamawasm2go/p2.Fn2423
func Fn2423(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2429 github.com/goccy/llamawasm2go/p2.Fn2429
func Fn2429(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2444 github.com/goccy/llamawasm2go/p2.Fn2444
func Fn2444(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2451 github.com/goccy/llamawasm2go/p0.Fn2451
func Fn2451(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2454 github.com/goccy/llamawasm2go/p2.Fn2454
func Fn2454(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2457 github.com/goccy/llamawasm2go/p2.Fn2457
func Fn2457(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2460 github.com/goccy/llamawasm2go/p2.Fn2460
func Fn2460(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2461 github.com/goccy/llamawasm2go/p2.Fn2461
func Fn2461(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2467 github.com/goccy/llamawasm2go/p2.Fn2467
func Fn2467(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2469 github.com/goccy/llamawasm2go/p2.Fn2469
func Fn2469(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2485 github.com/goccy/llamawasm2go/p2.Fn2485
func Fn2485(m *base.Module, l0 int64)

//go:linkname Fn2486 github.com/goccy/llamawasm2go/p2.Fn2486
func Fn2486(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2491 github.com/goccy/llamawasm2go/p2.Fn2491
func Fn2491(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2496 github.com/goccy/llamawasm2go/p2.Fn2496
func Fn2496(m *base.Module, l0 int64) int64

//go:linkname Fn2497 github.com/goccy/llamawasm2go/p2.Fn2497
func Fn2497(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2498 github.com/goccy/llamawasm2go/p0.Fn2498
func Fn2498(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2509 github.com/goccy/llamawasm2go/p2.Fn2509
func Fn2509(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2510 github.com/goccy/llamawasm2go/p2.Fn2510
func Fn2510(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2511 github.com/goccy/llamawasm2go/p2.Fn2511
func Fn2511(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2512 github.com/goccy/llamawasm2go/p2.Fn2512
func Fn2512(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2532 github.com/goccy/llamawasm2go/p2.Fn2532
func Fn2532(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2533 github.com/goccy/llamawasm2go/p2.Fn2533
func Fn2533(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2536 github.com/goccy/llamawasm2go/p2.Fn2536
func Fn2536(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2622 github.com/goccy/llamawasm2go/p2.Fn2622
func Fn2622(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2741 github.com/goccy/llamawasm2go/p2.Fn2741
func Fn2741(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2742 github.com/goccy/llamawasm2go/p0.Fn2742
func Fn2742(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2743 github.com/goccy/llamawasm2go/p2.Fn2743
func Fn2743(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2744 github.com/goccy/llamawasm2go/p2.Fn2744
func Fn2744(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2748 github.com/goccy/llamawasm2go/p2.Fn2748
func Fn2748(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2773 github.com/goccy/llamawasm2go/p2.Fn2773
func Fn2773(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2776 github.com/goccy/llamawasm2go/p0.Fn2776
func Fn2776(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2905 github.com/goccy/llamawasm2go/p2.Fn2905
func Fn2905(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2940 github.com/goccy/llamawasm2go/p2.Fn2940
func Fn2940(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2954 github.com/goccy/llamawasm2go/p2.Fn2954
func Fn2954(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2966 github.com/goccy/llamawasm2go/p2.Fn2966
func Fn2966(m *base.Module, l0 int32)

//go:linkname Fn2968 github.com/goccy/llamawasm2go/p2.Fn2968
func Fn2968(m *base.Module, l0 int64) int64

//go:linkname Fn2969 github.com/goccy/llamawasm2go/p2.Fn2969
func Fn2969(m *base.Module, l0 int64)

//go:linkname Fn2972 github.com/goccy/llamawasm2go/p2.Fn2972
func Fn2972(m *base.Module, l0 int64)

//go:linkname Fn2973 github.com/goccy/llamawasm2go/p2.Fn2973
func Fn2973(m *base.Module, l0 int64)

//go:linkname Fn2975 github.com/goccy/llamawasm2go/p2.Fn2975
func Fn2975(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2976 github.com/goccy/llamawasm2go/p2.Fn2976
func Fn2976(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2982 github.com/goccy/llamawasm2go/p2.Fn2982
func Fn2982(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2988 github.com/goccy/llamawasm2go/p2.Fn2988
func Fn2988(m *base.Module) int32

//go:linkname Fn2999 github.com/goccy/llamawasm2go/p2.Fn2999
func Fn2999(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3003 github.com/goccy/llamawasm2go/p2.Fn3003
func Fn3003(m *base.Module) int64

//go:linkname Fn3005 github.com/goccy/llamawasm2go/p2.Fn3005
func Fn3005(m *base.Module, l0 float64) float32

//go:linkname Fn3006 github.com/goccy/llamawasm2go/p2.Fn3006
func Fn3006(m *base.Module, l0 float64) float32

//go:linkname Fn3010 github.com/goccy/llamawasm2go/p2.Fn3010
func Fn3010(m *base.Module, l0 float64) float64

//go:linkname Fn3013 github.com/goccy/llamawasm2go/p2.Fn3013
func Fn3013(m *base.Module, l0 int32) float32

//go:linkname Fn3014 github.com/goccy/llamawasm2go/p2.Fn3014
func Fn3014(m *base.Module, l0 int32) float32

//go:linkname Fn3017 github.com/goccy/llamawasm2go/p2.Fn3017
func Fn3017(m *base.Module, l0 float32) float32

//go:linkname Fn3020 github.com/goccy/llamawasm2go/p2.Fn3020
func Fn3020(m *base.Module, l0 float64) float64

//go:linkname Fn3021 github.com/goccy/llamawasm2go/p2.Fn3021
func Fn3021(m *base.Module, l0 float64) float64

//go:linkname Fn3022 github.com/goccy/llamawasm2go/p2.Fn3022
func Fn3022(m *base.Module, l0 float32) float32

//go:linkname Fn3024 github.com/goccy/llamawasm2go/p2.Fn3024
func Fn3024(m *base.Module, l0 float32) float32

//go:linkname Fn3026 github.com/goccy/llamawasm2go/p2.Fn3026
func Fn3026(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn3027 github.com/goccy/llamawasm2go/p2.Fn3027
func Fn3027(m *base.Module, l0 float32) float32

//go:linkname Fn3044 github.com/goccy/llamawasm2go/p2.Fn3044
func Fn3044(m *base.Module, l0 int64) int32

//go:linkname Fn3045 github.com/goccy/llamawasm2go/p2.Fn3045
func Fn3045(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3047 github.com/goccy/llamawasm2go/p2.Fn3047
func Fn3047(m *base.Module, l0 int64)

//go:linkname Fn3048 github.com/goccy/llamawasm2go/p2.Fn3048
func Fn3048(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3049 github.com/goccy/llamawasm2go/p2.Fn3049
func Fn3049(m *base.Module, l0 int64) int32

//go:linkname Fn3056 github.com/goccy/llamawasm2go/p2.Fn3056
func Fn3056(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3058 github.com/goccy/llamawasm2go/p2.Fn3058
func Fn3058(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3064 github.com/goccy/llamawasm2go/p2.Fn3064
func Fn3064(m *base.Module, l0 int64) int32

//go:linkname Fn3067 github.com/goccy/llamawasm2go/p2.Fn3067
func Fn3067(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3070 github.com/goccy/llamawasm2go/p2.Fn3070
func Fn3070(m *base.Module, l0 int64) int32

//go:linkname Fn3072 github.com/goccy/llamawasm2go/p2.Fn3072
func Fn3072(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3074 github.com/goccy/llamawasm2go/p2.Fn3074
func Fn3074(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3075 github.com/goccy/llamawasm2go/p2.Fn3075
func Fn3075(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3078 github.com/goccy/llamawasm2go/p2.Fn3078
func Fn3078(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3081 github.com/goccy/llamawasm2go/p2.Fn3081
func Fn3081(m *base.Module, l0 int64) int64

//go:linkname Fn3085 github.com/goccy/llamawasm2go/p2.Fn3085
func Fn3085(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3095 github.com/goccy/llamawasm2go/p2.Fn3095
func Fn3095(m *base.Module)

//go:linkname Fn3096 github.com/goccy/llamawasm2go/p0.Fn3096
func Fn3096(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3099 github.com/goccy/llamawasm2go/p2.Fn3099
func Fn3099(m *base.Module)

//go:linkname Fn3101 github.com/goccy/llamawasm2go/p0.Fn3101
func Fn3101(m *base.Module, l0 int64) int64

//go:linkname Fn3103 github.com/goccy/llamawasm2go/p2.Fn3103
func Fn3103(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3107 github.com/goccy/llamawasm2go/p2.Fn3107
func Fn3107(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3176 github.com/goccy/llamawasm2go/p2.Fn3176
func Fn3176(m *base.Module, l0 int32)
