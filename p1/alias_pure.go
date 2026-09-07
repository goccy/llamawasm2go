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

//go:linkname Fn340 github.com/goccy/llamawasm2go/p2.Fn340
func Fn340(m *base.Module, l0 int64) int32

//go:linkname Fn341 github.com/goccy/llamawasm2go/p2.Fn341
func Fn341(m *base.Module, l0 int64)

//go:linkname Fn342 github.com/goccy/llamawasm2go/p2.Fn342
func Fn342(m *base.Module, l0 int64)

//go:linkname Fn359 github.com/goccy/llamawasm2go/p2.Fn359
func Fn359(m *base.Module) int64

//go:linkname Fn362 github.com/goccy/llamawasm2go/p2.Fn362
func Fn362(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn363 github.com/goccy/llamawasm2go/p2.Fn363
func Fn363(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn366 github.com/goccy/llamawasm2go/p2.Fn366
func Fn366(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn367 github.com/goccy/llamawasm2go/p2.Fn367
func Fn367(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn371 github.com/goccy/llamawasm2go/p2.Fn371
func Fn371(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn372 github.com/goccy/llamawasm2go/p2.Fn372
func Fn372(m *base.Module, l0 int64)

//go:linkname Fn373 github.com/goccy/llamawasm2go/p2.Fn373
func Fn373(m *base.Module, l0 int64)

//go:linkname Fn375 github.com/goccy/llamawasm2go/p2.Fn375
func Fn375(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn376 github.com/goccy/llamawasm2go/p2.Fn376
func Fn376(m *base.Module, l0 int64)

//go:linkname Fn379 github.com/goccy/llamawasm2go/p2.Fn379
func Fn379(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn380 github.com/goccy/llamawasm2go/p2.Fn380
func Fn380(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn382 github.com/goccy/llamawasm2go/p2.Fn382
func Fn382(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn386 github.com/goccy/llamawasm2go/p0.Fn386
func Fn386(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int64)

//go:linkname Fn389 github.com/goccy/llamawasm2go/p2.Fn389
func Fn389(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn390 github.com/goccy/llamawasm2go/p2.Fn390
func Fn390(m *base.Module, l0 int64)

//go:linkname Fn393 github.com/goccy/llamawasm2go/p2.Fn393
func Fn393(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn394 github.com/goccy/llamawasm2go/p2.Fn394
func Fn394(m *base.Module, l0 int64)

//go:linkname Fn395 github.com/goccy/llamawasm2go/p2.Fn395
func Fn395(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn396 github.com/goccy/llamawasm2go/p2.Fn396
func Fn396(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn397 github.com/goccy/llamawasm2go/p2.Fn397
func Fn397(m *base.Module)

//go:linkname Fn408 github.com/goccy/llamawasm2go/p0.Fn408
func Fn408(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int32)

//go:linkname Fn412 github.com/goccy/llamawasm2go/p0.Fn412
func Fn412(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 int64)

//go:linkname Fn426 github.com/goccy/llamawasm2go/p2.Fn426
func Fn426(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn430 github.com/goccy/llamawasm2go/p2.Fn430
func Fn430(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn431 github.com/goccy/llamawasm2go/p2.Fn431
func Fn431(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn433 github.com/goccy/llamawasm2go/p2.Fn433
func Fn433(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn434 github.com/goccy/llamawasm2go/p2.Fn434
func Fn434(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn455 github.com/goccy/llamawasm2go/p2.Fn455
func Fn455(m *base.Module, l0 int64) int32

//go:linkname Fn456 github.com/goccy/llamawasm2go/p2.Fn456
func Fn456(m *base.Module, l0 int64) int32

//go:linkname Fn467 github.com/goccy/llamawasm2go/p2.Fn467
func Fn467(m *base.Module, l0 int64) int64

//go:linkname Fn472 github.com/goccy/llamawasm2go/p2.Fn472
func Fn472(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn474 github.com/goccy/llamawasm2go/p2.Fn474
func Fn474(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn475 github.com/goccy/llamawasm2go/p2.Fn475
func Fn475(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn476 github.com/goccy/llamawasm2go/p2.Fn476
func Fn476(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn477 github.com/goccy/llamawasm2go/p2.Fn477
func Fn477(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn478 github.com/goccy/llamawasm2go/p2.Fn478
func Fn478(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn479 github.com/goccy/llamawasm2go/p2.Fn479
func Fn479(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn482 github.com/goccy/llamawasm2go/p2.Fn482
func Fn482(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn483 github.com/goccy/llamawasm2go/p2.Fn483
func Fn483(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn486 github.com/goccy/llamawasm2go/p2.Fn486
func Fn486(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn487 github.com/goccy/llamawasm2go/p2.Fn487
func Fn487(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn488 github.com/goccy/llamawasm2go/p2.Fn488
func Fn488(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn489 github.com/goccy/llamawasm2go/p2.Fn489
func Fn489(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn490 github.com/goccy/llamawasm2go/p2.Fn490
func Fn490(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn491 github.com/goccy/llamawasm2go/p2.Fn491
func Fn491(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn493 github.com/goccy/llamawasm2go/p2.Fn493
func Fn493(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn495 github.com/goccy/llamawasm2go/p2.Fn495
func Fn495(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn497 github.com/goccy/llamawasm2go/p2.Fn497
func Fn497(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn498 github.com/goccy/llamawasm2go/p2.Fn498
func Fn498(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn499 github.com/goccy/llamawasm2go/p2.Fn499
func Fn499(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn500 github.com/goccy/llamawasm2go/p2.Fn500
func Fn500(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn501 github.com/goccy/llamawasm2go/p2.Fn501
func Fn501(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn503 github.com/goccy/llamawasm2go/p2.Fn503
func Fn503(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn504 github.com/goccy/llamawasm2go/p2.Fn504
func Fn504(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn505 github.com/goccy/llamawasm2go/p2.Fn505
func Fn505(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn506 github.com/goccy/llamawasm2go/p2.Fn506
func Fn506(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn507 github.com/goccy/llamawasm2go/p2.Fn507
func Fn507(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn508 github.com/goccy/llamawasm2go/p2.Fn508
func Fn508(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn510 github.com/goccy/llamawasm2go/p2.Fn510
func Fn510(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn511 github.com/goccy/llamawasm2go/p2.Fn511
func Fn511(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn512 github.com/goccy/llamawasm2go/p2.Fn512
func Fn512(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn513 github.com/goccy/llamawasm2go/p2.Fn513
func Fn513(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn514 github.com/goccy/llamawasm2go/p2.Fn514
func Fn514(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn515 github.com/goccy/llamawasm2go/p2.Fn515
func Fn515(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn516 github.com/goccy/llamawasm2go/p2.Fn516
func Fn516(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn517 github.com/goccy/llamawasm2go/p2.Fn517
func Fn517(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn518 github.com/goccy/llamawasm2go/p2.Fn518
func Fn518(m *base.Module, l0 int64)

//go:linkname Fn519 github.com/goccy/llamawasm2go/p2.Fn519
func Fn519(m *base.Module, l0 int64)

//go:linkname Fn520 github.com/goccy/llamawasm2go/p2.Fn520
func Fn520(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn521 github.com/goccy/llamawasm2go/p2.Fn521
func Fn521(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn522 github.com/goccy/llamawasm2go/p2.Fn522
func Fn522(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn523 github.com/goccy/llamawasm2go/p2.Fn523
func Fn523(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn524 github.com/goccy/llamawasm2go/p2.Fn524
func Fn524(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn525 github.com/goccy/llamawasm2go/p2.Fn525
func Fn525(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn526 github.com/goccy/llamawasm2go/p2.Fn526
func Fn526(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn527 github.com/goccy/llamawasm2go/p2.Fn527
func Fn527(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn528 github.com/goccy/llamawasm2go/p2.Fn528
func Fn528(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn529 github.com/goccy/llamawasm2go/p2.Fn529
func Fn529(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn531 github.com/goccy/llamawasm2go/p2.Fn531
func Fn531(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn532 github.com/goccy/llamawasm2go/p2.Fn532
func Fn532(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn533 github.com/goccy/llamawasm2go/p2.Fn533
func Fn533(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn534 github.com/goccy/llamawasm2go/p2.Fn534
func Fn534(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn535 github.com/goccy/llamawasm2go/p2.Fn535
func Fn535(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn536 github.com/goccy/llamawasm2go/p2.Fn536
func Fn536(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn537 github.com/goccy/llamawasm2go/p2.Fn537
func Fn537(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64) int64

//go:linkname Fn538 github.com/goccy/llamawasm2go/p2.Fn538
func Fn538(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn539 github.com/goccy/llamawasm2go/p2.Fn539
func Fn539(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn540 github.com/goccy/llamawasm2go/p2.Fn540
func Fn540(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn541 github.com/goccy/llamawasm2go/p2.Fn541
func Fn541(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn542 github.com/goccy/llamawasm2go/p2.Fn542
func Fn542(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn544 github.com/goccy/llamawasm2go/p2.Fn544
func Fn544(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn547 github.com/goccy/llamawasm2go/p2.Fn547
func Fn547(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32) int64

//go:linkname Fn548 github.com/goccy/llamawasm2go/p2.Fn548
func Fn548(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32, l7 float32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32) int64

//go:linkname Fn550 github.com/goccy/llamawasm2go/p2.Fn550
func Fn550(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn551 github.com/goccy/llamawasm2go/p2.Fn551
func Fn551(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn552 github.com/goccy/llamawasm2go/p2.Fn552
func Fn552(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn554 github.com/goccy/llamawasm2go/p2.Fn554
func Fn554(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn556 github.com/goccy/llamawasm2go/p2.Fn556
func Fn556(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn557 github.com/goccy/llamawasm2go/p2.Fn557
func Fn557(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn558 github.com/goccy/llamawasm2go/p2.Fn558
func Fn558(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn560 github.com/goccy/llamawasm2go/p2.Fn560
func Fn560(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 float32, l6 float32, l7 float32) int64

//go:linkname Fn561 github.com/goccy/llamawasm2go/p2.Fn561
func Fn561(m *base.Module, l0 int64)

//go:linkname Fn562 github.com/goccy/llamawasm2go/p2.Fn562
func Fn562(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn566 github.com/goccy/llamawasm2go/p2.Fn566
func Fn566(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn567 github.com/goccy/llamawasm2go/p2.Fn567
func Fn567(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn570 github.com/goccy/llamawasm2go/p2.Fn570
func Fn570(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn572 github.com/goccy/llamawasm2go/p2.Fn572
func Fn572(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn574 github.com/goccy/llamawasm2go/p2.Fn574
func Fn574(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn580 github.com/goccy/llamawasm2go/p2.Fn580
func Fn580(m *base.Module, l0 int64)

//go:linkname Fn583 github.com/goccy/llamawasm2go/p2.Fn583
func Fn583(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn584 github.com/goccy/llamawasm2go/p2.Fn584
func Fn584(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn586 github.com/goccy/llamawasm2go/p2.Fn586
func Fn586(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn589 github.com/goccy/llamawasm2go/p2.Fn589
func Fn589(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn590 github.com/goccy/llamawasm2go/p2.Fn590
func Fn590(m *base.Module, l0 int64) int64

//go:linkname Fn591 github.com/goccy/llamawasm2go/p2.Fn591
func Fn591(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn592 github.com/goccy/llamawasm2go/p2.Fn592
func Fn592(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn593 github.com/goccy/llamawasm2go/p2.Fn593
func Fn593(m *base.Module, l0 int64) int64

//go:linkname Fn594 github.com/goccy/llamawasm2go/p2.Fn594
func Fn594(m *base.Module, l0 int64) int64

//go:linkname Fn595 github.com/goccy/llamawasm2go/p2.Fn595
func Fn595(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn597 github.com/goccy/llamawasm2go/p2.Fn597
func Fn597(m *base.Module, l0 int64) int64

//go:linkname Fn598 github.com/goccy/llamawasm2go/p2.Fn598
func Fn598(m *base.Module, l0 int64) int64

//go:linkname Fn600 github.com/goccy/llamawasm2go/p2.Fn600
func Fn600(m *base.Module, l0 int64)

//go:linkname Fn601 github.com/goccy/llamawasm2go/p2.Fn601
func Fn601(m *base.Module, l0 int64) int64

//go:linkname Fn602 github.com/goccy/llamawasm2go/p2.Fn602
func Fn602(m *base.Module, l0 int64) int64

//go:linkname Fn603 github.com/goccy/llamawasm2go/p2.Fn603
func Fn603(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn604 github.com/goccy/llamawasm2go/p2.Fn604
func Fn604(m *base.Module, l0 int64) int32

//go:linkname Fn605 github.com/goccy/llamawasm2go/p2.Fn605
func Fn605(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn607 github.com/goccy/llamawasm2go/p2.Fn607
func Fn607(m *base.Module, l0 int64)

//go:linkname Fn608 github.com/goccy/llamawasm2go/p2.Fn608
func Fn608(m *base.Module, l0 int64)

//go:linkname Fn609 github.com/goccy/llamawasm2go/p2.Fn609
func Fn609(m *base.Module, l0 int64) int64

//go:linkname Fn611 github.com/goccy/llamawasm2go/p2.Fn611
func Fn611(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn612 github.com/goccy/llamawasm2go/p2.Fn612
func Fn612(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn613 github.com/goccy/llamawasm2go/p2.Fn613
func Fn613(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn614 github.com/goccy/llamawasm2go/p2.Fn614
func Fn614(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn615 github.com/goccy/llamawasm2go/p2.Fn615
func Fn615(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn616 github.com/goccy/llamawasm2go/p2.Fn616
func Fn616(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn619 github.com/goccy/llamawasm2go/p2.Fn619
func Fn619(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn620 github.com/goccy/llamawasm2go/p2.Fn620
func Fn620(m *base.Module, l0 int64) int64

//go:linkname Fn621 github.com/goccy/llamawasm2go/p2.Fn621
func Fn621(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn624 github.com/goccy/llamawasm2go/p2.Fn624
func Fn624(m *base.Module, l0 int64) int64

//go:linkname Fn625 github.com/goccy/llamawasm2go/p2.Fn625
func Fn625(m *base.Module, l0 int64) int64

//go:linkname Fn629 github.com/goccy/llamawasm2go/p2.Fn629
func Fn629(m *base.Module, l0 int64) int64

//go:linkname Fn630 github.com/goccy/llamawasm2go/p2.Fn630
func Fn630(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn631 github.com/goccy/llamawasm2go/p2.Fn631
func Fn631(m *base.Module, l0 int64) int64

//go:linkname Fn635 github.com/goccy/llamawasm2go/p2.Fn635
func Fn635(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn636 github.com/goccy/llamawasm2go/p0.Fn636
func Fn636(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn641 github.com/goccy/llamawasm2go/p2.Fn641
func Fn641(m *base.Module, l0 int64)

//go:linkname Fn642 github.com/goccy/llamawasm2go/p2.Fn642
func Fn642(m *base.Module, l0 int64)

//go:linkname Fn643 github.com/goccy/llamawasm2go/p2.Fn643
func Fn643(m *base.Module, l0 int64)

//go:linkname Fn644 github.com/goccy/llamawasm2go/p2.Fn644
func Fn644(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn645 github.com/goccy/llamawasm2go/p2.Fn645
func Fn645(m *base.Module, l0 int64) int32

//go:linkname Fn646 github.com/goccy/llamawasm2go/p2.Fn646
func Fn646(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn647 github.com/goccy/llamawasm2go/p2.Fn647
func Fn647(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn648 github.com/goccy/llamawasm2go/p2.Fn648
func Fn648(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn649 github.com/goccy/llamawasm2go/p2.Fn649
func Fn649(m *base.Module, l0 int64) int32

//go:linkname Fn650 github.com/goccy/llamawasm2go/p2.Fn650
func Fn650(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn668 github.com/goccy/llamawasm2go/p2.Fn668
func Fn668(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn674 github.com/goccy/llamawasm2go/p2.Fn674
func Fn674(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn675 github.com/goccy/llamawasm2go/p2.Fn675
func Fn675(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn676 github.com/goccy/llamawasm2go/p2.Fn676
func Fn676(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn678 github.com/goccy/llamawasm2go/p2.Fn678
func Fn678(m *base.Module, l0 int64)

//go:linkname Fn680 github.com/goccy/llamawasm2go/p2.Fn680
func Fn680(m *base.Module, l0 int64)

//go:linkname Fn681 github.com/goccy/llamawasm2go/p0.Fn681
func Fn681(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn683 github.com/goccy/llamawasm2go/p2.Fn683
func Fn683(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn697 github.com/goccy/llamawasm2go/p2.Fn697
func Fn697(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn699 github.com/goccy/llamawasm2go/p2.Fn699
func Fn699(m *base.Module, l0 int64)

//go:linkname Fn700 github.com/goccy/llamawasm2go/p2.Fn700
func Fn700(m *base.Module, l0 int64) int64

//go:linkname Fn723 github.com/goccy/llamawasm2go/p2.Fn723
func Fn723(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn734 github.com/goccy/llamawasm2go/p2.Fn734
func Fn734(m *base.Module, l0 int64)

//go:linkname Fn753 github.com/goccy/llamawasm2go/p2.Fn753
func Fn753(m *base.Module, l0 int64)

//go:linkname Fn754 github.com/goccy/llamawasm2go/p2.Fn754
func Fn754(m *base.Module, l0 int64)

//go:linkname Fn822 github.com/goccy/llamawasm2go/p2.Fn822
func Fn822(m *base.Module, l0 int64)

//go:linkname Fn824 github.com/goccy/llamawasm2go/p2.Fn824
func Fn824(m *base.Module, l0 int64) int64

//go:linkname Fn885 github.com/goccy/llamawasm2go/p0.Fn885
func Fn885(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn933 github.com/goccy/llamawasm2go/p2.Fn933
func Fn933(m *base.Module) int64

//go:linkname Fn970 github.com/goccy/llamawasm2go/p2.Fn970
func Fn970(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32)

//go:linkname Fn980 github.com/goccy/llamawasm2go/p2.Fn980
func Fn980(m *base.Module, l0 int64)

//go:linkname Fn1005 github.com/goccy/llamawasm2go/p2.Fn1005
func Fn1005(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn1010 github.com/goccy/llamawasm2go/p2.Fn1010
func Fn1010(m *base.Module, l0 int64) int64

//go:linkname Fn1020 github.com/goccy/llamawasm2go/p2.Fn1020
func Fn1020(m *base.Module, l0 int64)

//go:linkname Fn1025 github.com/goccy/llamawasm2go/p2.Fn1025
func Fn1025(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1026 github.com/goccy/llamawasm2go/p2.Fn1026
func Fn1026(m *base.Module)

//go:linkname Fn1035 github.com/goccy/llamawasm2go/p2.Fn1035
func Fn1035(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1043 github.com/goccy/llamawasm2go/p2.Fn1043
func Fn1043(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1044 github.com/goccy/llamawasm2go/p2.Fn1044
func Fn1044(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1045 github.com/goccy/llamawasm2go/p2.Fn1045
func Fn1045(m *base.Module, l0 int64, l1 int64, l2 int64) float32

//go:linkname Fn1047 github.com/goccy/llamawasm2go/p2.Fn1047
func Fn1047(m *base.Module, l0 int64, l1 int64, l2 int64) float64

//go:linkname Fn1049 github.com/goccy/llamawasm2go/p2.Fn1049
func Fn1049(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1052 github.com/goccy/llamawasm2go/p2.Fn1052
func Fn1052(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1055 github.com/goccy/llamawasm2go/p2.Fn1055
func Fn1055(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1061 github.com/goccy/llamawasm2go/p2.Fn1061
func Fn1061(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1062 github.com/goccy/llamawasm2go/p2.Fn1062
func Fn1062(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1070 github.com/goccy/llamawasm2go/p2.Fn1070
func Fn1070(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1075 github.com/goccy/llamawasm2go/p2.Fn1075
func Fn1075(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1076 github.com/goccy/llamawasm2go/p2.Fn1076
func Fn1076(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1084 github.com/goccy/llamawasm2go/p2.Fn1084
func Fn1084(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1100 github.com/goccy/llamawasm2go/p2.Fn1100
func Fn1100(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1108 github.com/goccy/llamawasm2go/p2.Fn1108
func Fn1108(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1124 github.com/goccy/llamawasm2go/p0.Fn1124
func Fn1124(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1127 github.com/goccy/llamawasm2go/p2.Fn1127
func Fn1127(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1131 github.com/goccy/llamawasm2go/p2.Fn1131
func Fn1131(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1141 github.com/goccy/llamawasm2go/p0.Fn1141
func Fn1141(m *base.Module, l0 int64) int64

//go:linkname Fn1142 github.com/goccy/llamawasm2go/p2.Fn1142
func Fn1142(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1143 github.com/goccy/llamawasm2go/p2.Fn1143
func Fn1143(m *base.Module, l0 int64)

//go:linkname Fn1145 github.com/goccy/llamawasm2go/p2.Fn1145
func Fn1145(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1235 github.com/goccy/llamawasm2go/p2.Fn1235
func Fn1235(m *base.Module, l0 int64)

//go:linkname Fn1257 github.com/goccy/llamawasm2go/p2.Fn1257
func Fn1257(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1267 github.com/goccy/llamawasm2go/p2.Fn1267
func Fn1267(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1273 github.com/goccy/llamawasm2go/p2.Fn1273
func Fn1273(m *base.Module)

//go:linkname Fn1277 github.com/goccy/llamawasm2go/p2.Fn1277
func Fn1277(m *base.Module, l0 int64) int64

//go:linkname Fn1305 github.com/goccy/llamawasm2go/p2.Fn1305
func Fn1305(m *base.Module, l0 int64)

//go:linkname Fn1309 github.com/goccy/llamawasm2go/p2.Fn1309
func Fn1309(m *base.Module, l0 int32) int64

//go:linkname Fn1320 github.com/goccy/llamawasm2go/p2.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1321 github.com/goccy/llamawasm2go/p2.Fn1321
func Fn1321(m *base.Module, l0 int64)

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1345 github.com/goccy/llamawasm2go/p2.Fn1345
func Fn1345(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1346 github.com/goccy/llamawasm2go/p2.Fn1346
func Fn1346(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1347 github.com/goccy/llamawasm2go/p2.Fn1347
func Fn1347(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1348 github.com/goccy/llamawasm2go/p2.Fn1348
func Fn1348(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1349 github.com/goccy/llamawasm2go/p2.Fn1349
func Fn1349(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1350 github.com/goccy/llamawasm2go/p2.Fn1350
func Fn1350(m *base.Module, l0 int64) int64

//go:linkname Fn1353 github.com/goccy/llamawasm2go/p2.Fn1353
func Fn1353(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1357 github.com/goccy/llamawasm2go/p2.Fn1357
func Fn1357(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1358 github.com/goccy/llamawasm2go/p2.Fn1358
func Fn1358(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1359 github.com/goccy/llamawasm2go/p2.Fn1359
func Fn1359(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1360 github.com/goccy/llamawasm2go/p2.Fn1360
func Fn1360(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1361 github.com/goccy/llamawasm2go/p2.Fn1361
func Fn1361(m *base.Module, l0 int64, l1 int64, l2 int64) int32

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
func Fn1371(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1372 github.com/goccy/llamawasm2go/p2.Fn1372
func Fn1372(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1373 github.com/goccy/llamawasm2go/p2.Fn1373
func Fn1373(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1374 github.com/goccy/llamawasm2go/p2.Fn1374
func Fn1374(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1375 github.com/goccy/llamawasm2go/p2.Fn1375
func Fn1375(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1376 github.com/goccy/llamawasm2go/p2.Fn1376
func Fn1376(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1377 github.com/goccy/llamawasm2go/p2.Fn1377
func Fn1377(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1378 github.com/goccy/llamawasm2go/p2.Fn1378
func Fn1378(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1379 github.com/goccy/llamawasm2go/p2.Fn1379
func Fn1379(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1380 github.com/goccy/llamawasm2go/p2.Fn1380
func Fn1380(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1381 github.com/goccy/llamawasm2go/p2.Fn1381
func Fn1381(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1382 github.com/goccy/llamawasm2go/p2.Fn1382
func Fn1382(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1383 github.com/goccy/llamawasm2go/p2.Fn1383
func Fn1383(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1385 github.com/goccy/llamawasm2go/p2.Fn1385
func Fn1385(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1387 github.com/goccy/llamawasm2go/p2.Fn1387
func Fn1387(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1388 github.com/goccy/llamawasm2go/p2.Fn1388
func Fn1388(m *base.Module)

//go:linkname Fn1389 github.com/goccy/llamawasm2go/p2.Fn1389
func Fn1389(m *base.Module)

//go:linkname Fn1390 github.com/goccy/llamawasm2go/p0.Fn1390
func Fn1390(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module)

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p0.Fn1402
func Fn1402(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module, l0 int64) int64

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p2.Fn1408
func Fn1408(m *base.Module, l0 int64)

//go:linkname Fn1415 github.com/goccy/llamawasm2go/p2.Fn1415
func Fn1415(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1418 github.com/goccy/llamawasm2go/p2.Fn1418
func Fn1418(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1422 github.com/goccy/llamawasm2go/p2.Fn1422
func Fn1422(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p2.Fn1426
func Fn1426(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1428 github.com/goccy/llamawasm2go/p2.Fn1428
func Fn1428(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1434 github.com/goccy/llamawasm2go/p2.Fn1434
func Fn1434(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1435 github.com/goccy/llamawasm2go/p2.Fn1435
func Fn1435(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1436 github.com/goccy/llamawasm2go/p2.Fn1436
func Fn1436(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1437 github.com/goccy/llamawasm2go/p0.Fn1437
func Fn1437(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1439 github.com/goccy/llamawasm2go/p2.Fn1439
func Fn1439(m *base.Module, l0 int64)

//go:linkname Fn1440 github.com/goccy/llamawasm2go/p2.Fn1440
func Fn1440(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int64) int64

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64)

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int64)

//go:linkname Fn1446 github.com/goccy/llamawasm2go/p2.Fn1446
func Fn1446(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1450 github.com/goccy/llamawasm2go/p2.Fn1450
func Fn1450(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1458 github.com/goccy/llamawasm2go/p2.Fn1458
func Fn1458(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1460 github.com/goccy/llamawasm2go/p0.Fn1460
func Fn1460(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1462 github.com/goccy/llamawasm2go/p2.Fn1462
func Fn1462(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1465 github.com/goccy/llamawasm2go/p0.Fn1465
func Fn1465(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1466 github.com/goccy/llamawasm2go/p2.Fn1466
func Fn1466(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1470 github.com/goccy/llamawasm2go/p2.Fn1470
func Fn1470(m *base.Module, l0 int64)

//go:linkname Fn1473 github.com/goccy/llamawasm2go/p2.Fn1473
func Fn1473(m *base.Module, l0 int64)

//go:linkname Fn1476 github.com/goccy/llamawasm2go/p2.Fn1476
func Fn1476(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1477 github.com/goccy/llamawasm2go/p2.Fn1477
func Fn1477(m *base.Module, l0 int64) int64

//go:linkname Fn1478 github.com/goccy/llamawasm2go/p2.Fn1478
func Fn1478(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1479 github.com/goccy/llamawasm2go/p2.Fn1479
func Fn1479(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1480 github.com/goccy/llamawasm2go/p2.Fn1480
func Fn1480(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1481 github.com/goccy/llamawasm2go/p2.Fn1481
func Fn1481(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1482 github.com/goccy/llamawasm2go/p2.Fn1482
func Fn1482(m *base.Module, l0 int64)

//go:linkname Fn1485 github.com/goccy/llamawasm2go/p2.Fn1485
func Fn1485(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1487 github.com/goccy/llamawasm2go/p2.Fn1487
func Fn1487(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1490 github.com/goccy/llamawasm2go/p2.Fn1490
func Fn1490(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1497 github.com/goccy/llamawasm2go/p2.Fn1497
func Fn1497(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1498 github.com/goccy/llamawasm2go/p2.Fn1498
func Fn1498(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1500 github.com/goccy/llamawasm2go/p2.Fn1500
func Fn1500(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1504 github.com/goccy/llamawasm2go/p2.Fn1504
func Fn1504(m *base.Module, l0 int64)

//go:linkname Fn1505 github.com/goccy/llamawasm2go/p2.Fn1505
func Fn1505(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1507 github.com/goccy/llamawasm2go/p2.Fn1507
func Fn1507(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1508 github.com/goccy/llamawasm2go/p2.Fn1508
func Fn1508(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1509 github.com/goccy/llamawasm2go/p2.Fn1509
func Fn1509(m *base.Module, l0 int64)

//go:linkname Fn1510 github.com/goccy/llamawasm2go/p2.Fn1510
func Fn1510(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1512 github.com/goccy/llamawasm2go/p2.Fn1512
func Fn1512(m *base.Module, l0 int64) int32

//go:linkname Fn1513 github.com/goccy/llamawasm2go/p2.Fn1513
func Fn1513(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1555 github.com/goccy/llamawasm2go/p2.Fn1555
func Fn1555(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1571 github.com/goccy/llamawasm2go/p2.Fn1571
func Fn1571(m *base.Module, l0 int64)

//go:linkname Fn1572 github.com/goccy/llamawasm2go/p2.Fn1572
func Fn1572(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1573 github.com/goccy/llamawasm2go/p2.Fn1573
func Fn1573(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1574 github.com/goccy/llamawasm2go/p2.Fn1574
func Fn1574(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1575 github.com/goccy/llamawasm2go/p2.Fn1575
func Fn1575(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1576 github.com/goccy/llamawasm2go/p2.Fn1576
func Fn1576(m *base.Module, l0 int64)

//go:linkname Fn1577 github.com/goccy/llamawasm2go/p2.Fn1577
func Fn1577(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1579 github.com/goccy/llamawasm2go/p2.Fn1579
func Fn1579(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1580 github.com/goccy/llamawasm2go/p2.Fn1580
func Fn1580(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1581 github.com/goccy/llamawasm2go/p2.Fn1581
func Fn1581(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1582 github.com/goccy/llamawasm2go/p2.Fn1582
func Fn1582(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p2.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1586 github.com/goccy/llamawasm2go/p2.Fn1586
func Fn1586(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1588 github.com/goccy/llamawasm2go/p2.Fn1588
func Fn1588(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1589 github.com/goccy/llamawasm2go/p2.Fn1589
func Fn1589(m *base.Module, l0 int64) int64

//go:linkname Fn1590 github.com/goccy/llamawasm2go/p2.Fn1590
func Fn1590(m *base.Module, l0 int64) int64

//go:linkname Fn1591 github.com/goccy/llamawasm2go/p2.Fn1591
func Fn1591(m *base.Module, l0 int64) int64

//go:linkname Fn1592 github.com/goccy/llamawasm2go/p2.Fn1592
func Fn1592(m *base.Module, l0 int64) int64

//go:linkname Fn1593 github.com/goccy/llamawasm2go/p2.Fn1593
func Fn1593(m *base.Module, l0 int64) int64

//go:linkname Fn1594 github.com/goccy/llamawasm2go/p2.Fn1594
func Fn1594(m *base.Module, l0 int64) int64

//go:linkname Fn1596 github.com/goccy/llamawasm2go/p2.Fn1596
func Fn1596(m *base.Module, l0 int64) int64

//go:linkname Fn1597 github.com/goccy/llamawasm2go/p2.Fn1597
func Fn1597(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1599 github.com/goccy/llamawasm2go/p2.Fn1599
func Fn1599(m *base.Module, l0 int64) int64

//go:linkname Fn1600 github.com/goccy/llamawasm2go/p2.Fn1600
func Fn1600(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1601 github.com/goccy/llamawasm2go/p2.Fn1601
func Fn1601(m *base.Module, l0 int64) int64

//go:linkname Fn1603 github.com/goccy/llamawasm2go/p2.Fn1603
func Fn1603(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1604 github.com/goccy/llamawasm2go/p2.Fn1604
func Fn1604(m *base.Module, l0 int64) int64

//go:linkname Fn1606 github.com/goccy/llamawasm2go/p2.Fn1606
func Fn1606(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1607 github.com/goccy/llamawasm2go/p2.Fn1607
func Fn1607(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1608 github.com/goccy/llamawasm2go/p2.Fn1608
func Fn1608(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1609 github.com/goccy/llamawasm2go/p2.Fn1609
func Fn1609(m *base.Module, l0 int64) int64

//go:linkname Fn1610 github.com/goccy/llamawasm2go/p2.Fn1610
func Fn1610(m *base.Module, l0 int64) int64

//go:linkname Fn1611 github.com/goccy/llamawasm2go/p2.Fn1611
func Fn1611(m *base.Module, l0 int64) int64

//go:linkname Fn1617 github.com/goccy/llamawasm2go/p2.Fn1617
func Fn1617(m *base.Module, l0 int64) int64

//go:linkname Fn1619 github.com/goccy/llamawasm2go/p2.Fn1619
func Fn1619(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1620 github.com/goccy/llamawasm2go/p2.Fn1620
func Fn1620(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1622 github.com/goccy/llamawasm2go/p2.Fn1622
func Fn1622(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1623 github.com/goccy/llamawasm2go/p2.Fn1623
func Fn1623(m *base.Module, l0 int64) int64

//go:linkname Fn1624 github.com/goccy/llamawasm2go/p2.Fn1624
func Fn1624(m *base.Module, l0 int64) int64

//go:linkname Fn1625 github.com/goccy/llamawasm2go/p2.Fn1625
func Fn1625(m *base.Module, l0 int64) int64

//go:linkname Fn1626 github.com/goccy/llamawasm2go/p2.Fn1626
func Fn1626(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1627 github.com/goccy/llamawasm2go/p2.Fn1627
func Fn1627(m *base.Module, l0 int64)

//go:linkname Fn1656 github.com/goccy/llamawasm2go/p2.Fn1656
func Fn1656(m *base.Module, l0 int64) int64

//go:linkname Fn1663 github.com/goccy/llamawasm2go/p2.Fn1663
func Fn1663(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1664 github.com/goccy/llamawasm2go/p2.Fn1664
func Fn1664(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1665 github.com/goccy/llamawasm2go/p2.Fn1665
func Fn1665(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1667 github.com/goccy/llamawasm2go/p2.Fn1667
func Fn1667(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1668 github.com/goccy/llamawasm2go/p2.Fn1668
func Fn1668(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1672 github.com/goccy/llamawasm2go/p2.Fn1672
func Fn1672(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1673 github.com/goccy/llamawasm2go/p2.Fn1673
func Fn1673(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1674 github.com/goccy/llamawasm2go/p2.Fn1674
func Fn1674(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1675 github.com/goccy/llamawasm2go/p2.Fn1675
func Fn1675(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1676 github.com/goccy/llamawasm2go/p2.Fn1676
func Fn1676(m *base.Module, l0 int64) int32

//go:linkname Fn1677 github.com/goccy/llamawasm2go/p2.Fn1677
func Fn1677(m *base.Module, l0 int64) int32

//go:linkname Fn1678 github.com/goccy/llamawasm2go/p2.Fn1678
func Fn1678(m *base.Module, l0 int64) int32

//go:linkname Fn1679 github.com/goccy/llamawasm2go/p2.Fn1679
func Fn1679(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1682 github.com/goccy/llamawasm2go/p2.Fn1682
func Fn1682(m *base.Module, l0 int64) int32

//go:linkname Fn1683 github.com/goccy/llamawasm2go/p2.Fn1683
func Fn1683(m *base.Module, l0 int64) int32

//go:linkname Fn1689 github.com/goccy/llamawasm2go/p2.Fn1689
func Fn1689(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1690 github.com/goccy/llamawasm2go/p2.Fn1690
func Fn1690(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1691 github.com/goccy/llamawasm2go/p2.Fn1691
func Fn1691(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1692 github.com/goccy/llamawasm2go/p2.Fn1692
func Fn1692(m *base.Module)

//go:linkname Fn1693 github.com/goccy/llamawasm2go/p2.Fn1693
func Fn1693(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1695 github.com/goccy/llamawasm2go/p2.Fn1695
func Fn1695(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1697 github.com/goccy/llamawasm2go/p2.Fn1697
func Fn1697(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1698 github.com/goccy/llamawasm2go/p2.Fn1698
func Fn1698(m *base.Module, l0 int64)

//go:linkname Fn1702 github.com/goccy/llamawasm2go/p2.Fn1702
func Fn1702(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1703 github.com/goccy/llamawasm2go/p2.Fn1703
func Fn1703(m *base.Module, l0 int64)

//go:linkname Fn1706 github.com/goccy/llamawasm2go/p2.Fn1706
func Fn1706(m *base.Module, l0 int64)

//go:linkname Fn1716 github.com/goccy/llamawasm2go/p2.Fn1716
func Fn1716(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1717 github.com/goccy/llamawasm2go/p2.Fn1717
func Fn1717(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1726 github.com/goccy/llamawasm2go/p2.Fn1726
func Fn1726(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1727 github.com/goccy/llamawasm2go/p0.Fn1727
func Fn1727(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1729 github.com/goccy/llamawasm2go/p2.Fn1729
func Fn1729(m *base.Module, l0 int64)

//go:linkname Fn1731 github.com/goccy/llamawasm2go/p2.Fn1731
func Fn1731(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1738 github.com/goccy/llamawasm2go/p2.Fn1738
func Fn1738(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1739 github.com/goccy/llamawasm2go/p2.Fn1739
func Fn1739(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1746 github.com/goccy/llamawasm2go/p2.Fn1746
func Fn1746(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1753 github.com/goccy/llamawasm2go/p2.Fn1753
func Fn1753(m *base.Module, l0 int64)

//go:linkname Fn1756 github.com/goccy/llamawasm2go/p2.Fn1756
func Fn1756(m *base.Module, l0 int64) int32

//go:linkname Fn1766 github.com/goccy/llamawasm2go/p2.Fn1766
func Fn1766(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1768 github.com/goccy/llamawasm2go/p2.Fn1768
func Fn1768(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1769 github.com/goccy/llamawasm2go/p2.Fn1769
func Fn1769(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1788 github.com/goccy/llamawasm2go/p2.Fn1788
func Fn1788(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1801 github.com/goccy/llamawasm2go/p2.Fn1801
func Fn1801(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1808 github.com/goccy/llamawasm2go/p2.Fn1808
func Fn1808(m *base.Module, l0 int64)

//go:linkname Fn1842 github.com/goccy/llamawasm2go/p2.Fn1842
func Fn1842(m *base.Module, l0 int64)

//go:linkname Fn1845 github.com/goccy/llamawasm2go/p2.Fn1845
func Fn1845(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1859 github.com/goccy/llamawasm2go/p2.Fn1859
func Fn1859(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1860 github.com/goccy/llamawasm2go/p2.Fn1860
func Fn1860(m *base.Module, l0 int64) int64

//go:linkname Fn1861 github.com/goccy/llamawasm2go/p2.Fn1861
func Fn1861(m *base.Module, l0 int64)

//go:linkname Fn1865 github.com/goccy/llamawasm2go/p0.Fn1865
func Fn1865(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1866 github.com/goccy/llamawasm2go/p2.Fn1866
func Fn1866(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1867 github.com/goccy/llamawasm2go/p2.Fn1867
func Fn1867(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1868 github.com/goccy/llamawasm2go/p2.Fn1868
func Fn1868(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1869 github.com/goccy/llamawasm2go/p2.Fn1869
func Fn1869(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1887 github.com/goccy/llamawasm2go/p2.Fn1887
func Fn1887(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1898 github.com/goccy/llamawasm2go/p2.Fn1898
func Fn1898(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1900 github.com/goccy/llamawasm2go/p2.Fn1900
func Fn1900(m *base.Module, l0 int64) int64

//go:linkname Fn1901 github.com/goccy/llamawasm2go/p2.Fn1901
func Fn1901(m *base.Module, l0 int64)

//go:linkname Fn1904 github.com/goccy/llamawasm2go/p0.Fn1904
func Fn1904(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1906 github.com/goccy/llamawasm2go/p2.Fn1906
func Fn1906(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1917 github.com/goccy/llamawasm2go/p2.Fn1917
func Fn1917(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1920 github.com/goccy/llamawasm2go/p2.Fn1920
func Fn1920(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1921 github.com/goccy/llamawasm2go/p2.Fn1921
func Fn1921(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1931 github.com/goccy/llamawasm2go/p2.Fn1931
func Fn1931(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1958 github.com/goccy/llamawasm2go/p2.Fn1958
func Fn1958(m *base.Module, l0 int64)

//go:linkname Fn1959 github.com/goccy/llamawasm2go/p2.Fn1959
func Fn1959(m *base.Module, l0 int64)

//go:linkname Fn1970 github.com/goccy/llamawasm2go/p2.Fn1970
func Fn1970(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2002 github.com/goccy/llamawasm2go/p2.Fn2002
func Fn2002(m *base.Module) int64

//go:linkname Fn2007 github.com/goccy/llamawasm2go/p2.Fn2007
func Fn2007(m *base.Module, l0 int64) int64

//go:linkname Fn2008 github.com/goccy/llamawasm2go/p2.Fn2008
func Fn2008(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2015 github.com/goccy/llamawasm2go/p2.Fn2015
func Fn2015(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2022 github.com/goccy/llamawasm2go/p2.Fn2022
func Fn2022(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2029 github.com/goccy/llamawasm2go/p2.Fn2029
func Fn2029(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2036 github.com/goccy/llamawasm2go/p2.Fn2036
func Fn2036(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2038 github.com/goccy/llamawasm2go/p2.Fn2038
func Fn2038(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2041 github.com/goccy/llamawasm2go/p2.Fn2041
func Fn2041(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2042 github.com/goccy/llamawasm2go/p2.Fn2042
func Fn2042(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2043 github.com/goccy/llamawasm2go/p2.Fn2043
func Fn2043(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn2046 github.com/goccy/llamawasm2go/p2.Fn2046
func Fn2046(m *base.Module, l0 int64)

//go:linkname Fn2056 github.com/goccy/llamawasm2go/p2.Fn2056
func Fn2056(m *base.Module, l0 int64)

//go:linkname Fn2058 github.com/goccy/llamawasm2go/p2.Fn2058
func Fn2058(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2059 github.com/goccy/llamawasm2go/p2.Fn2059
func Fn2059(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2072 github.com/goccy/llamawasm2go/p2.Fn2072
func Fn2072(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2073 github.com/goccy/llamawasm2go/p2.Fn2073
func Fn2073(m *base.Module, l0 int64) int64

//go:linkname Fn2074 github.com/goccy/llamawasm2go/p2.Fn2074
func Fn2074(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2076 github.com/goccy/llamawasm2go/p2.Fn2076
func Fn2076(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2090 github.com/goccy/llamawasm2go/p2.Fn2090
func Fn2090(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2091 github.com/goccy/llamawasm2go/p2.Fn2091
func Fn2091(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2093 github.com/goccy/llamawasm2go/p2.Fn2093
func Fn2093(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2095 github.com/goccy/llamawasm2go/p2.Fn2095
func Fn2095(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2096 github.com/goccy/llamawasm2go/p2.Fn2096
func Fn2096(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2097 github.com/goccy/llamawasm2go/p2.Fn2097
func Fn2097(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2098 github.com/goccy/llamawasm2go/p2.Fn2098
func Fn2098(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2117 github.com/goccy/llamawasm2go/p2.Fn2117
func Fn2117(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2118 github.com/goccy/llamawasm2go/p2.Fn2118
func Fn2118(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2121 github.com/goccy/llamawasm2go/p2.Fn2121
func Fn2121(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2122 github.com/goccy/llamawasm2go/p2.Fn2122
func Fn2122(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2123 github.com/goccy/llamawasm2go/p2.Fn2123
func Fn2123(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2124 github.com/goccy/llamawasm2go/p2.Fn2124
func Fn2124(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2125 github.com/goccy/llamawasm2go/p2.Fn2125
func Fn2125(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2127 github.com/goccy/llamawasm2go/p2.Fn2127
func Fn2127(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2129 github.com/goccy/llamawasm2go/p2.Fn2129
func Fn2129(m *base.Module, l0 int64)

//go:linkname Fn2146 github.com/goccy/llamawasm2go/p2.Fn2146
func Fn2146(m *base.Module, l0 int64)

//go:linkname Fn2147 github.com/goccy/llamawasm2go/p2.Fn2147
func Fn2147(m *base.Module, l0 int64)

//go:linkname Fn2148 github.com/goccy/llamawasm2go/p2.Fn2148
func Fn2148(m *base.Module, l0 int64)

//go:linkname Fn2150 github.com/goccy/llamawasm2go/p2.Fn2150
func Fn2150(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2151 github.com/goccy/llamawasm2go/p2.Fn2151
func Fn2151(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2188 github.com/goccy/llamawasm2go/p2.Fn2188
func Fn2188(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2197 github.com/goccy/llamawasm2go/p2.Fn2197
func Fn2197(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2199 github.com/goccy/llamawasm2go/p2.Fn2199
func Fn2199(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2201 github.com/goccy/llamawasm2go/p2.Fn2201
func Fn2201(m *base.Module, l0 int64) int64

//go:linkname Fn2204 github.com/goccy/llamawasm2go/p2.Fn2204
func Fn2204(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2208 github.com/goccy/llamawasm2go/p2.Fn2208
func Fn2208(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2213 github.com/goccy/llamawasm2go/p2.Fn2213
func Fn2213(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2227 github.com/goccy/llamawasm2go/p2.Fn2227
func Fn2227(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2229 github.com/goccy/llamawasm2go/p2.Fn2229
func Fn2229(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2230 github.com/goccy/llamawasm2go/p2.Fn2230
func Fn2230(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2244 github.com/goccy/llamawasm2go/p2.Fn2244
func Fn2244(m *base.Module, l0 int64) int64

//go:linkname Fn2246 github.com/goccy/llamawasm2go/p2.Fn2246
func Fn2246(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p2.Fn2249
func Fn2249(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2254 github.com/goccy/llamawasm2go/p2.Fn2254
func Fn2254(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2263 github.com/goccy/llamawasm2go/p2.Fn2263
func Fn2263(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2267 github.com/goccy/llamawasm2go/p2.Fn2267
func Fn2267(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 int64)

//go:linkname Fn2269 github.com/goccy/llamawasm2go/p2.Fn2269
func Fn2269(m *base.Module, l0 int64)

//go:linkname Fn2270 github.com/goccy/llamawasm2go/p2.Fn2270
func Fn2270(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2272 github.com/goccy/llamawasm2go/p2.Fn2272
func Fn2272(m *base.Module, l0 int64)

//go:linkname Fn2273 github.com/goccy/llamawasm2go/p2.Fn2273
func Fn2273(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2274 github.com/goccy/llamawasm2go/p2.Fn2274
func Fn2274(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2278 github.com/goccy/llamawasm2go/p2.Fn2278
func Fn2278(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2281 github.com/goccy/llamawasm2go/p2.Fn2281
func Fn2281(m *base.Module, l0 int64) int64

//go:linkname Fn2282 github.com/goccy/llamawasm2go/p2.Fn2282
func Fn2282(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2283 github.com/goccy/llamawasm2go/p2.Fn2283
func Fn2283(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2290 github.com/goccy/llamawasm2go/p2.Fn2290
func Fn2290(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2291 github.com/goccy/llamawasm2go/p2.Fn2291
func Fn2291(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2292 github.com/goccy/llamawasm2go/p2.Fn2292
func Fn2292(m *base.Module, l0 int64)

//go:linkname Fn2293 github.com/goccy/llamawasm2go/p2.Fn2293
func Fn2293(m *base.Module) int64

//go:linkname Fn2295 github.com/goccy/llamawasm2go/p2.Fn2295
func Fn2295(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2296 github.com/goccy/llamawasm2go/p2.Fn2296
func Fn2296(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2298 github.com/goccy/llamawasm2go/p2.Fn2298
func Fn2298(m *base.Module) int64

//go:linkname Fn2300 github.com/goccy/llamawasm2go/p2.Fn2300
func Fn2300(m *base.Module, l0 int32) int64

//go:linkname Fn2301 github.com/goccy/llamawasm2go/p2.Fn2301
func Fn2301(m *base.Module, l0 int32) int32

//go:linkname Fn2302 github.com/goccy/llamawasm2go/p2.Fn2302
func Fn2302(m *base.Module, l0 int32) int64

//go:linkname Fn2303 github.com/goccy/llamawasm2go/p2.Fn2303
func Fn2303(m *base.Module, l0 float32) int64

//go:linkname Fn2304 github.com/goccy/llamawasm2go/p2.Fn2304
func Fn2304(m *base.Module, l0 float32) int64

//go:linkname Fn2305 github.com/goccy/llamawasm2go/p2.Fn2305
func Fn2305(m *base.Module, l0 float32) int64

//go:linkname Fn2307 github.com/goccy/llamawasm2go/p2.Fn2307
func Fn2307(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2308 github.com/goccy/llamawasm2go/p2.Fn2308
func Fn2308(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2349 github.com/goccy/llamawasm2go/p2.Fn2349
func Fn2349(m *base.Module, l0 int64)

//go:linkname Fn2351 github.com/goccy/llamawasm2go/p2.Fn2351
func Fn2351(m *base.Module, l0 int64)

//go:linkname Fn2359 github.com/goccy/llamawasm2go/p2.Fn2359
func Fn2359(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2393 github.com/goccy/llamawasm2go/p2.Fn2393
func Fn2393(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2404 github.com/goccy/llamawasm2go/p2.Fn2404
func Fn2404(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2407 github.com/goccy/llamawasm2go/p2.Fn2407
func Fn2407(m *base.Module, l0 int64)

//go:linkname Fn2408 github.com/goccy/llamawasm2go/p2.Fn2408
func Fn2408(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2409 github.com/goccy/llamawasm2go/p2.Fn2409
func Fn2409(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2411 github.com/goccy/llamawasm2go/p2.Fn2411
func Fn2411(m *base.Module, l0 int64)

//go:linkname Fn2415 github.com/goccy/llamawasm2go/p2.Fn2415
func Fn2415(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2421 github.com/goccy/llamawasm2go/p2.Fn2421
func Fn2421(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2436 github.com/goccy/llamawasm2go/p2.Fn2436
func Fn2436(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2443 github.com/goccy/llamawasm2go/p0.Fn2443
func Fn2443(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2446 github.com/goccy/llamawasm2go/p2.Fn2446
func Fn2446(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2449 github.com/goccy/llamawasm2go/p2.Fn2449
func Fn2449(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2452 github.com/goccy/llamawasm2go/p2.Fn2452
func Fn2452(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2453 github.com/goccy/llamawasm2go/p2.Fn2453
func Fn2453(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2459 github.com/goccy/llamawasm2go/p2.Fn2459
func Fn2459(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2461 github.com/goccy/llamawasm2go/p2.Fn2461
func Fn2461(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2477 github.com/goccy/llamawasm2go/p2.Fn2477
func Fn2477(m *base.Module, l0 int64)

//go:linkname Fn2478 github.com/goccy/llamawasm2go/p2.Fn2478
func Fn2478(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2483 github.com/goccy/llamawasm2go/p2.Fn2483
func Fn2483(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2488 github.com/goccy/llamawasm2go/p2.Fn2488
func Fn2488(m *base.Module, l0 int64) int64

//go:linkname Fn2489 github.com/goccy/llamawasm2go/p2.Fn2489
func Fn2489(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2490 github.com/goccy/llamawasm2go/p0.Fn2490
func Fn2490(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2501 github.com/goccy/llamawasm2go/p2.Fn2501
func Fn2501(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2502 github.com/goccy/llamawasm2go/p2.Fn2502
func Fn2502(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2503 github.com/goccy/llamawasm2go/p2.Fn2503
func Fn2503(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2504 github.com/goccy/llamawasm2go/p2.Fn2504
func Fn2504(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2524 github.com/goccy/llamawasm2go/p2.Fn2524
func Fn2524(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2525 github.com/goccy/llamawasm2go/p2.Fn2525
func Fn2525(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2528 github.com/goccy/llamawasm2go/p2.Fn2528
func Fn2528(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2614 github.com/goccy/llamawasm2go/p2.Fn2614
func Fn2614(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2733 github.com/goccy/llamawasm2go/p2.Fn2733
func Fn2733(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2734 github.com/goccy/llamawasm2go/p0.Fn2734
func Fn2734(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2735 github.com/goccy/llamawasm2go/p2.Fn2735
func Fn2735(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2736 github.com/goccy/llamawasm2go/p2.Fn2736
func Fn2736(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2740 github.com/goccy/llamawasm2go/p2.Fn2740
func Fn2740(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2765 github.com/goccy/llamawasm2go/p2.Fn2765
func Fn2765(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2897 github.com/goccy/llamawasm2go/p2.Fn2897
func Fn2897(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2932 github.com/goccy/llamawasm2go/p2.Fn2932
func Fn2932(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2946 github.com/goccy/llamawasm2go/p2.Fn2946
func Fn2946(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2958 github.com/goccy/llamawasm2go/p2.Fn2958
func Fn2958(m *base.Module, l0 int32)

//go:linkname Fn2960 github.com/goccy/llamawasm2go/p2.Fn2960
func Fn2960(m *base.Module, l0 int64) int64

//go:linkname Fn2961 github.com/goccy/llamawasm2go/p2.Fn2961
func Fn2961(m *base.Module, l0 int64)

//go:linkname Fn2964 github.com/goccy/llamawasm2go/p2.Fn2964
func Fn2964(m *base.Module, l0 int64)

//go:linkname Fn2965 github.com/goccy/llamawasm2go/p2.Fn2965
func Fn2965(m *base.Module, l0 int64)

//go:linkname Fn2967 github.com/goccy/llamawasm2go/p2.Fn2967
func Fn2967(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2968 github.com/goccy/llamawasm2go/p2.Fn2968
func Fn2968(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2974 github.com/goccy/llamawasm2go/p2.Fn2974
func Fn2974(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2980 github.com/goccy/llamawasm2go/p2.Fn2980
func Fn2980(m *base.Module) int32

//go:linkname Fn2991 github.com/goccy/llamawasm2go/p2.Fn2991
func Fn2991(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2995 github.com/goccy/llamawasm2go/p2.Fn2995
func Fn2995(m *base.Module) int64

//go:linkname Fn2997 github.com/goccy/llamawasm2go/p2.Fn2997
func Fn2997(m *base.Module, l0 float64) float32

//go:linkname Fn2998 github.com/goccy/llamawasm2go/p2.Fn2998
func Fn2998(m *base.Module, l0 float64) float32

//go:linkname Fn3002 github.com/goccy/llamawasm2go/p2.Fn3002
func Fn3002(m *base.Module, l0 float64) float64

//go:linkname Fn3005 github.com/goccy/llamawasm2go/p2.Fn3005
func Fn3005(m *base.Module, l0 int32) float32

//go:linkname Fn3006 github.com/goccy/llamawasm2go/p2.Fn3006
func Fn3006(m *base.Module, l0 int32) float32

//go:linkname Fn3009 github.com/goccy/llamawasm2go/p2.Fn3009
func Fn3009(m *base.Module, l0 float32) float32

//go:linkname Fn3012 github.com/goccy/llamawasm2go/p2.Fn3012
func Fn3012(m *base.Module, l0 float64) float64

//go:linkname Fn3013 github.com/goccy/llamawasm2go/p2.Fn3013
func Fn3013(m *base.Module, l0 float64) float64

//go:linkname Fn3014 github.com/goccy/llamawasm2go/p2.Fn3014
func Fn3014(m *base.Module, l0 float32) float32

//go:linkname Fn3016 github.com/goccy/llamawasm2go/p2.Fn3016
func Fn3016(m *base.Module, l0 float32) float32

//go:linkname Fn3018 github.com/goccy/llamawasm2go/p2.Fn3018
func Fn3018(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn3019 github.com/goccy/llamawasm2go/p2.Fn3019
func Fn3019(m *base.Module, l0 float32) float32

//go:linkname Fn3036 github.com/goccy/llamawasm2go/p2.Fn3036
func Fn3036(m *base.Module, l0 int64) int32

//go:linkname Fn3037 github.com/goccy/llamawasm2go/p2.Fn3037
func Fn3037(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3039 github.com/goccy/llamawasm2go/p2.Fn3039
func Fn3039(m *base.Module, l0 int64)

//go:linkname Fn3040 github.com/goccy/llamawasm2go/p2.Fn3040
func Fn3040(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3041 github.com/goccy/llamawasm2go/p2.Fn3041
func Fn3041(m *base.Module, l0 int64) int32

//go:linkname Fn3048 github.com/goccy/llamawasm2go/p2.Fn3048
func Fn3048(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3050 github.com/goccy/llamawasm2go/p2.Fn3050
func Fn3050(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3056 github.com/goccy/llamawasm2go/p2.Fn3056
func Fn3056(m *base.Module, l0 int64) int32

//go:linkname Fn3059 github.com/goccy/llamawasm2go/p2.Fn3059
func Fn3059(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3062 github.com/goccy/llamawasm2go/p2.Fn3062
func Fn3062(m *base.Module, l0 int64) int32

//go:linkname Fn3064 github.com/goccy/llamawasm2go/p2.Fn3064
func Fn3064(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3066 github.com/goccy/llamawasm2go/p2.Fn3066
func Fn3066(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3067 github.com/goccy/llamawasm2go/p2.Fn3067
func Fn3067(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3070 github.com/goccy/llamawasm2go/p2.Fn3070
func Fn3070(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3073 github.com/goccy/llamawasm2go/p2.Fn3073
func Fn3073(m *base.Module, l0 int64) int64

//go:linkname Fn3077 github.com/goccy/llamawasm2go/p2.Fn3077
func Fn3077(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3087 github.com/goccy/llamawasm2go/p2.Fn3087
func Fn3087(m *base.Module)

//go:linkname Fn3088 github.com/goccy/llamawasm2go/p0.Fn3088
func Fn3088(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3091 github.com/goccy/llamawasm2go/p2.Fn3091
func Fn3091(m *base.Module)

//go:linkname Fn3093 github.com/goccy/llamawasm2go/p0.Fn3093
func Fn3093(m *base.Module, l0 int64) int64

//go:linkname Fn3095 github.com/goccy/llamawasm2go/p2.Fn3095
func Fn3095(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3099 github.com/goccy/llamawasm2go/p2.Fn3099
func Fn3099(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3168 github.com/goccy/llamawasm2go/p2.Fn3168
func Fn3168(m *base.Module, l0 int32)
