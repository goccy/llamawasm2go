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

//go:linkname Fn244 github.com/goccy/llamawasm2go/p0.Fn244
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

//go:linkname Fn292 github.com/goccy/llamawasm2go/p2.Fn292
func Fn292(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn293 github.com/goccy/llamawasm2go/p2.Fn293
func Fn293(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn332 github.com/goccy/llamawasm2go/p2.Fn332
func Fn332(m *base.Module, l0 int64) int32

//go:linkname Fn333 github.com/goccy/llamawasm2go/p2.Fn333
func Fn333(m *base.Module, l0 int64)

//go:linkname Fn334 github.com/goccy/llamawasm2go/p2.Fn334
func Fn334(m *base.Module, l0 int64)

//go:linkname Fn353 github.com/goccy/llamawasm2go/p2.Fn353
func Fn353(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn354 github.com/goccy/llamawasm2go/p2.Fn354
func Fn354(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn358 github.com/goccy/llamawasm2go/p2.Fn358
func Fn358(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn359 github.com/goccy/llamawasm2go/p2.Fn359
func Fn359(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn362 github.com/goccy/llamawasm2go/p2.Fn362
func Fn362(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn363 github.com/goccy/llamawasm2go/p2.Fn363
func Fn363(m *base.Module, l0 int64)

//go:linkname Fn364 github.com/goccy/llamawasm2go/p2.Fn364
func Fn364(m *base.Module, l0 int64)

//go:linkname Fn366 github.com/goccy/llamawasm2go/p2.Fn366
func Fn366(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn368 github.com/goccy/llamawasm2go/p2.Fn368
func Fn368(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn369 github.com/goccy/llamawasm2go/p2.Fn369
func Fn369(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn371 github.com/goccy/llamawasm2go/p2.Fn371
func Fn371(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn374 github.com/goccy/llamawasm2go/p0.Fn374
func Fn374(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int64)

//go:linkname Fn377 github.com/goccy/llamawasm2go/p2.Fn377
func Fn377(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn378 github.com/goccy/llamawasm2go/p2.Fn378
func Fn378(m *base.Module, l0 int64)

//go:linkname Fn381 github.com/goccy/llamawasm2go/p2.Fn381
func Fn381(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn382 github.com/goccy/llamawasm2go/p2.Fn382
func Fn382(m *base.Module, l0 int64)

//go:linkname Fn383 github.com/goccy/llamawasm2go/p2.Fn383
func Fn383(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn384 github.com/goccy/llamawasm2go/p2.Fn384
func Fn384(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn385 github.com/goccy/llamawasm2go/p2.Fn385
func Fn385(m *base.Module)

//go:linkname Fn386 github.com/goccy/llamawasm2go/p0.Fn386
func Fn386(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int32)

//go:linkname Fn390 github.com/goccy/llamawasm2go/p0.Fn390
func Fn390(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 int64)

//go:linkname Fn394 github.com/goccy/llamawasm2go/p2.Fn394
func Fn394(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn400 github.com/goccy/llamawasm2go/p2.Fn400
func Fn400(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn402 github.com/goccy/llamawasm2go/p2.Fn402
func Fn402(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn403 github.com/goccy/llamawasm2go/p2.Fn403
func Fn403(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn424 github.com/goccy/llamawasm2go/p2.Fn424
func Fn424(m *base.Module, l0 int64) int32

//go:linkname Fn425 github.com/goccy/llamawasm2go/p2.Fn425
func Fn425(m *base.Module, l0 int64) int32

//go:linkname Fn436 github.com/goccy/llamawasm2go/p2.Fn436
func Fn436(m *base.Module, l0 int64) int64

//go:linkname Fn441 github.com/goccy/llamawasm2go/p2.Fn441
func Fn441(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn443 github.com/goccy/llamawasm2go/p2.Fn443
func Fn443(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn444 github.com/goccy/llamawasm2go/p2.Fn444
func Fn444(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn445 github.com/goccy/llamawasm2go/p2.Fn445
func Fn445(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn446 github.com/goccy/llamawasm2go/p2.Fn446
func Fn446(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn447 github.com/goccy/llamawasm2go/p2.Fn447
func Fn447(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn448 github.com/goccy/llamawasm2go/p2.Fn448
func Fn448(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn451 github.com/goccy/llamawasm2go/p2.Fn451
func Fn451(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn452 github.com/goccy/llamawasm2go/p2.Fn452
func Fn452(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn455 github.com/goccy/llamawasm2go/p2.Fn455
func Fn455(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn456 github.com/goccy/llamawasm2go/p2.Fn456
func Fn456(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn457 github.com/goccy/llamawasm2go/p2.Fn457
func Fn457(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn458 github.com/goccy/llamawasm2go/p2.Fn458
func Fn458(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn459 github.com/goccy/llamawasm2go/p2.Fn459
func Fn459(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn460 github.com/goccy/llamawasm2go/p2.Fn460
func Fn460(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn462 github.com/goccy/llamawasm2go/p2.Fn462
func Fn462(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn464 github.com/goccy/llamawasm2go/p2.Fn464
func Fn464(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn466 github.com/goccy/llamawasm2go/p2.Fn466
func Fn466(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn467 github.com/goccy/llamawasm2go/p2.Fn467
func Fn467(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn468 github.com/goccy/llamawasm2go/p2.Fn468
func Fn468(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn469 github.com/goccy/llamawasm2go/p2.Fn469
func Fn469(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn470 github.com/goccy/llamawasm2go/p2.Fn470
func Fn470(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn472 github.com/goccy/llamawasm2go/p2.Fn472
func Fn472(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn473 github.com/goccy/llamawasm2go/p2.Fn473
func Fn473(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn474 github.com/goccy/llamawasm2go/p2.Fn474
func Fn474(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn475 github.com/goccy/llamawasm2go/p2.Fn475
func Fn475(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn476 github.com/goccy/llamawasm2go/p2.Fn476
func Fn476(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn477 github.com/goccy/llamawasm2go/p2.Fn477
func Fn477(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn479 github.com/goccy/llamawasm2go/p2.Fn479
func Fn479(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn480 github.com/goccy/llamawasm2go/p2.Fn480
func Fn480(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn481 github.com/goccy/llamawasm2go/p2.Fn481
func Fn481(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn482 github.com/goccy/llamawasm2go/p2.Fn482
func Fn482(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn483 github.com/goccy/llamawasm2go/p2.Fn483
func Fn483(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn484 github.com/goccy/llamawasm2go/p2.Fn484
func Fn484(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn485 github.com/goccy/llamawasm2go/p2.Fn485
func Fn485(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn486 github.com/goccy/llamawasm2go/p2.Fn486
func Fn486(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn487 github.com/goccy/llamawasm2go/p2.Fn487
func Fn487(m *base.Module, l0 int64)

//go:linkname Fn488 github.com/goccy/llamawasm2go/p2.Fn488
func Fn488(m *base.Module, l0 int64)

//go:linkname Fn489 github.com/goccy/llamawasm2go/p2.Fn489
func Fn489(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn490 github.com/goccy/llamawasm2go/p2.Fn490
func Fn490(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn491 github.com/goccy/llamawasm2go/p2.Fn491
func Fn491(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn492 github.com/goccy/llamawasm2go/p2.Fn492
func Fn492(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn493 github.com/goccy/llamawasm2go/p2.Fn493
func Fn493(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn494 github.com/goccy/llamawasm2go/p2.Fn494
func Fn494(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn495 github.com/goccy/llamawasm2go/p2.Fn495
func Fn495(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn496 github.com/goccy/llamawasm2go/p2.Fn496
func Fn496(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn497 github.com/goccy/llamawasm2go/p2.Fn497
func Fn497(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn498 github.com/goccy/llamawasm2go/p2.Fn498
func Fn498(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn500 github.com/goccy/llamawasm2go/p2.Fn500
func Fn500(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn501 github.com/goccy/llamawasm2go/p2.Fn501
func Fn501(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn502 github.com/goccy/llamawasm2go/p2.Fn502
func Fn502(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn503 github.com/goccy/llamawasm2go/p2.Fn503
func Fn503(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn504 github.com/goccy/llamawasm2go/p2.Fn504
func Fn504(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn505 github.com/goccy/llamawasm2go/p2.Fn505
func Fn505(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn506 github.com/goccy/llamawasm2go/p2.Fn506
func Fn506(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64) int64

//go:linkname Fn507 github.com/goccy/llamawasm2go/p2.Fn507
func Fn507(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn508 github.com/goccy/llamawasm2go/p2.Fn508
func Fn508(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn509 github.com/goccy/llamawasm2go/p2.Fn509
func Fn509(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn510 github.com/goccy/llamawasm2go/p2.Fn510
func Fn510(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn511 github.com/goccy/llamawasm2go/p2.Fn511
func Fn511(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn513 github.com/goccy/llamawasm2go/p2.Fn513
func Fn513(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn516 github.com/goccy/llamawasm2go/p2.Fn516
func Fn516(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32) int64

//go:linkname Fn517 github.com/goccy/llamawasm2go/p2.Fn517
func Fn517(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32, l7 float32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32) int64

//go:linkname Fn519 github.com/goccy/llamawasm2go/p2.Fn519
func Fn519(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn520 github.com/goccy/llamawasm2go/p2.Fn520
func Fn520(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn521 github.com/goccy/llamawasm2go/p2.Fn521
func Fn521(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn523 github.com/goccy/llamawasm2go/p2.Fn523
func Fn523(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn525 github.com/goccy/llamawasm2go/p2.Fn525
func Fn525(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn526 github.com/goccy/llamawasm2go/p2.Fn526
func Fn526(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn527 github.com/goccy/llamawasm2go/p2.Fn527
func Fn527(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn529 github.com/goccy/llamawasm2go/p2.Fn529
func Fn529(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 float32, l6 float32, l7 float32) int64

//go:linkname Fn530 github.com/goccy/llamawasm2go/p2.Fn530
func Fn530(m *base.Module, l0 int64)

//go:linkname Fn531 github.com/goccy/llamawasm2go/p2.Fn531
func Fn531(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn535 github.com/goccy/llamawasm2go/p2.Fn535
func Fn535(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn536 github.com/goccy/llamawasm2go/p2.Fn536
func Fn536(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn539 github.com/goccy/llamawasm2go/p2.Fn539
func Fn539(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn541 github.com/goccy/llamawasm2go/p2.Fn541
func Fn541(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn543 github.com/goccy/llamawasm2go/p2.Fn543
func Fn543(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn549 github.com/goccy/llamawasm2go/p2.Fn549
func Fn549(m *base.Module, l0 int64)

//go:linkname Fn552 github.com/goccy/llamawasm2go/p2.Fn552
func Fn552(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn553 github.com/goccy/llamawasm2go/p2.Fn553
func Fn553(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn555 github.com/goccy/llamawasm2go/p2.Fn555
func Fn555(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn558 github.com/goccy/llamawasm2go/p2.Fn558
func Fn558(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn559 github.com/goccy/llamawasm2go/p2.Fn559
func Fn559(m *base.Module, l0 int64) int64

//go:linkname Fn560 github.com/goccy/llamawasm2go/p2.Fn560
func Fn560(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn561 github.com/goccy/llamawasm2go/p2.Fn561
func Fn561(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn562 github.com/goccy/llamawasm2go/p2.Fn562
func Fn562(m *base.Module, l0 int64) int64

//go:linkname Fn563 github.com/goccy/llamawasm2go/p2.Fn563
func Fn563(m *base.Module, l0 int64) int64

//go:linkname Fn564 github.com/goccy/llamawasm2go/p2.Fn564
func Fn564(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn566 github.com/goccy/llamawasm2go/p2.Fn566
func Fn566(m *base.Module, l0 int64) int64

//go:linkname Fn567 github.com/goccy/llamawasm2go/p2.Fn567
func Fn567(m *base.Module, l0 int64) int64

//go:linkname Fn569 github.com/goccy/llamawasm2go/p2.Fn569
func Fn569(m *base.Module, l0 int64)

//go:linkname Fn570 github.com/goccy/llamawasm2go/p2.Fn570
func Fn570(m *base.Module, l0 int64) int64

//go:linkname Fn571 github.com/goccy/llamawasm2go/p2.Fn571
func Fn571(m *base.Module, l0 int64) int64

//go:linkname Fn572 github.com/goccy/llamawasm2go/p2.Fn572
func Fn572(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn573 github.com/goccy/llamawasm2go/p2.Fn573
func Fn573(m *base.Module, l0 int64) int32

//go:linkname Fn574 github.com/goccy/llamawasm2go/p2.Fn574
func Fn574(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn576 github.com/goccy/llamawasm2go/p2.Fn576
func Fn576(m *base.Module, l0 int64)

//go:linkname Fn577 github.com/goccy/llamawasm2go/p2.Fn577
func Fn577(m *base.Module, l0 int64)

//go:linkname Fn578 github.com/goccy/llamawasm2go/p2.Fn578
func Fn578(m *base.Module, l0 int64) int64

//go:linkname Fn580 github.com/goccy/llamawasm2go/p2.Fn580
func Fn580(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn581 github.com/goccy/llamawasm2go/p2.Fn581
func Fn581(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn582 github.com/goccy/llamawasm2go/p2.Fn582
func Fn582(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn583 github.com/goccy/llamawasm2go/p2.Fn583
func Fn583(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn584 github.com/goccy/llamawasm2go/p2.Fn584
func Fn584(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn585 github.com/goccy/llamawasm2go/p2.Fn585
func Fn585(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn588 github.com/goccy/llamawasm2go/p2.Fn588
func Fn588(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn589 github.com/goccy/llamawasm2go/p2.Fn589
func Fn589(m *base.Module, l0 int64) int64

//go:linkname Fn590 github.com/goccy/llamawasm2go/p2.Fn590
func Fn590(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn593 github.com/goccy/llamawasm2go/p2.Fn593
func Fn593(m *base.Module, l0 int64) int64

//go:linkname Fn594 github.com/goccy/llamawasm2go/p2.Fn594
func Fn594(m *base.Module, l0 int64) int64

//go:linkname Fn598 github.com/goccy/llamawasm2go/p2.Fn598
func Fn598(m *base.Module, l0 int64) int64

//go:linkname Fn599 github.com/goccy/llamawasm2go/p2.Fn599
func Fn599(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn600 github.com/goccy/llamawasm2go/p2.Fn600
func Fn600(m *base.Module, l0 int64) int64

//go:linkname Fn604 github.com/goccy/llamawasm2go/p2.Fn604
func Fn604(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn605 github.com/goccy/llamawasm2go/p0.Fn605
func Fn605(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn610 github.com/goccy/llamawasm2go/p2.Fn610
func Fn610(m *base.Module, l0 int64)

//go:linkname Fn611 github.com/goccy/llamawasm2go/p2.Fn611
func Fn611(m *base.Module, l0 int64)

//go:linkname Fn612 github.com/goccy/llamawasm2go/p2.Fn612
func Fn612(m *base.Module, l0 int64)

//go:linkname Fn613 github.com/goccy/llamawasm2go/p2.Fn613
func Fn613(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn614 github.com/goccy/llamawasm2go/p2.Fn614
func Fn614(m *base.Module, l0 int64) int32

//go:linkname Fn615 github.com/goccy/llamawasm2go/p2.Fn615
func Fn615(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn616 github.com/goccy/llamawasm2go/p2.Fn616
func Fn616(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn617 github.com/goccy/llamawasm2go/p2.Fn617
func Fn617(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn618 github.com/goccy/llamawasm2go/p2.Fn618
func Fn618(m *base.Module, l0 int64) int32

//go:linkname Fn619 github.com/goccy/llamawasm2go/p2.Fn619
func Fn619(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn637 github.com/goccy/llamawasm2go/p2.Fn637
func Fn637(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn643 github.com/goccy/llamawasm2go/p2.Fn643
func Fn643(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn645 github.com/goccy/llamawasm2go/p2.Fn645
func Fn645(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn647 github.com/goccy/llamawasm2go/p2.Fn647
func Fn647(m *base.Module, l0 int64)

//go:linkname Fn649 github.com/goccy/llamawasm2go/p2.Fn649
func Fn649(m *base.Module, l0 int64)

//go:linkname Fn650 github.com/goccy/llamawasm2go/p0.Fn650
func Fn650(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn652 github.com/goccy/llamawasm2go/p2.Fn652
func Fn652(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn666 github.com/goccy/llamawasm2go/p2.Fn666
func Fn666(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn668 github.com/goccy/llamawasm2go/p2.Fn668
func Fn668(m *base.Module, l0 int64)

//go:linkname Fn669 github.com/goccy/llamawasm2go/p2.Fn669
func Fn669(m *base.Module, l0 int64) int64

//go:linkname Fn683 github.com/goccy/llamawasm2go/p2.Fn683
func Fn683(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn692 github.com/goccy/llamawasm2go/p2.Fn692
func Fn692(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn703 github.com/goccy/llamawasm2go/p2.Fn703
func Fn703(m *base.Module, l0 int64)

//go:linkname Fn722 github.com/goccy/llamawasm2go/p2.Fn722
func Fn722(m *base.Module, l0 int64)

//go:linkname Fn723 github.com/goccy/llamawasm2go/p2.Fn723
func Fn723(m *base.Module, l0 int64)

//go:linkname Fn791 github.com/goccy/llamawasm2go/p2.Fn791
func Fn791(m *base.Module, l0 int64)

//go:linkname Fn793 github.com/goccy/llamawasm2go/p2.Fn793
func Fn793(m *base.Module, l0 int64) int64

//go:linkname Fn895 github.com/goccy/llamawasm2go/p2.Fn895
func Fn895(m *base.Module) int64

//go:linkname Fn966 github.com/goccy/llamawasm2go/p2.Fn966
func Fn966(m *base.Module, l0 int64) int64

//go:linkname Fn976 github.com/goccy/llamawasm2go/p2.Fn976
func Fn976(m *base.Module, l0 int64)

//go:linkname Fn981 github.com/goccy/llamawasm2go/p2.Fn981
func Fn981(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn982 github.com/goccy/llamawasm2go/p2.Fn982
func Fn982(m *base.Module)

//go:linkname Fn991 github.com/goccy/llamawasm2go/p2.Fn991
func Fn991(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn999 github.com/goccy/llamawasm2go/p2.Fn999
func Fn999(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1000 github.com/goccy/llamawasm2go/p2.Fn1000
func Fn1000(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1001 github.com/goccy/llamawasm2go/p2.Fn1001
func Fn1001(m *base.Module, l0 int64, l1 int64, l2 int64) float32

//go:linkname Fn1003 github.com/goccy/llamawasm2go/p2.Fn1003
func Fn1003(m *base.Module, l0 int64, l1 int64, l2 int64) float64

//go:linkname Fn1005 github.com/goccy/llamawasm2go/p2.Fn1005
func Fn1005(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1008 github.com/goccy/llamawasm2go/p2.Fn1008
func Fn1008(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1011 github.com/goccy/llamawasm2go/p2.Fn1011
func Fn1011(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1017 github.com/goccy/llamawasm2go/p2.Fn1017
func Fn1017(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1018 github.com/goccy/llamawasm2go/p2.Fn1018
func Fn1018(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1026 github.com/goccy/llamawasm2go/p2.Fn1026
func Fn1026(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1031 github.com/goccy/llamawasm2go/p2.Fn1031
func Fn1031(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1032 github.com/goccy/llamawasm2go/p2.Fn1032
func Fn1032(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1040 github.com/goccy/llamawasm2go/p2.Fn1040
func Fn1040(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1056 github.com/goccy/llamawasm2go/p2.Fn1056
func Fn1056(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1064 github.com/goccy/llamawasm2go/p2.Fn1064
func Fn1064(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1080 github.com/goccy/llamawasm2go/p0.Fn1080
func Fn1080(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1083 github.com/goccy/llamawasm2go/p2.Fn1083
func Fn1083(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1087 github.com/goccy/llamawasm2go/p2.Fn1087
func Fn1087(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1097 github.com/goccy/llamawasm2go/p0.Fn1097
func Fn1097(m *base.Module, l0 int64) int64

//go:linkname Fn1098 github.com/goccy/llamawasm2go/p2.Fn1098
func Fn1098(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1099 github.com/goccy/llamawasm2go/p2.Fn1099
func Fn1099(m *base.Module, l0 int64)

//go:linkname Fn1101 github.com/goccy/llamawasm2go/p2.Fn1101
func Fn1101(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1191 github.com/goccy/llamawasm2go/p2.Fn1191
func Fn1191(m *base.Module, l0 int64)

//go:linkname Fn1213 github.com/goccy/llamawasm2go/p2.Fn1213
func Fn1213(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1223 github.com/goccy/llamawasm2go/p2.Fn1223
func Fn1223(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1229 github.com/goccy/llamawasm2go/p2.Fn1229
func Fn1229(m *base.Module)

//go:linkname Fn1233 github.com/goccy/llamawasm2go/p2.Fn1233
func Fn1233(m *base.Module, l0 int64) int64

//go:linkname Fn1261 github.com/goccy/llamawasm2go/p2.Fn1261
func Fn1261(m *base.Module, l0 int64)

//go:linkname Fn1265 github.com/goccy/llamawasm2go/p2.Fn1265
func Fn1265(m *base.Module, l0 int32) int64

//go:linkname Fn1276 github.com/goccy/llamawasm2go/p2.Fn1276
func Fn1276(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1277 github.com/goccy/llamawasm2go/p2.Fn1277
func Fn1277(m *base.Module, l0 int64)

//go:linkname Fn1279 github.com/goccy/llamawasm2go/p2.Fn1279
func Fn1279(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1300 github.com/goccy/llamawasm2go/p2.Fn1300
func Fn1300(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1301 github.com/goccy/llamawasm2go/p2.Fn1301
func Fn1301(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1302 github.com/goccy/llamawasm2go/p2.Fn1302
func Fn1302(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1303 github.com/goccy/llamawasm2go/p2.Fn1303
func Fn1303(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1304 github.com/goccy/llamawasm2go/p2.Fn1304
func Fn1304(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1305 github.com/goccy/llamawasm2go/p2.Fn1305
func Fn1305(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1306 github.com/goccy/llamawasm2go/p2.Fn1306
func Fn1306(m *base.Module, l0 int64) int64

//go:linkname Fn1309 github.com/goccy/llamawasm2go/p2.Fn1309
func Fn1309(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1313 github.com/goccy/llamawasm2go/p2.Fn1313
func Fn1313(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1314 github.com/goccy/llamawasm2go/p2.Fn1314
func Fn1314(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1315 github.com/goccy/llamawasm2go/p2.Fn1315
func Fn1315(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1316 github.com/goccy/llamawasm2go/p2.Fn1316
func Fn1316(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1317 github.com/goccy/llamawasm2go/p2.Fn1317
func Fn1317(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1318 github.com/goccy/llamawasm2go/p2.Fn1318
func Fn1318(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1319 github.com/goccy/llamawasm2go/p2.Fn1319
func Fn1319(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1320 github.com/goccy/llamawasm2go/p2.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1321 github.com/goccy/llamawasm2go/p2.Fn1321
func Fn1321(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1322 github.com/goccy/llamawasm2go/p2.Fn1322
func Fn1322(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p2.Fn1324
func Fn1324(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1325 github.com/goccy/llamawasm2go/p2.Fn1325
func Fn1325(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1326 github.com/goccy/llamawasm2go/p2.Fn1326
func Fn1326(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1327 github.com/goccy/llamawasm2go/p2.Fn1327
func Fn1327(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1328 github.com/goccy/llamawasm2go/p2.Fn1328
func Fn1328(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1329 github.com/goccy/llamawasm2go/p2.Fn1329
func Fn1329(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1330 github.com/goccy/llamawasm2go/p2.Fn1330
func Fn1330(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1331 github.com/goccy/llamawasm2go/p2.Fn1331
func Fn1331(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1332 github.com/goccy/llamawasm2go/p2.Fn1332
func Fn1332(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1333 github.com/goccy/llamawasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1335 github.com/goccy/llamawasm2go/p2.Fn1335
func Fn1335(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1336 github.com/goccy/llamawasm2go/p2.Fn1336
func Fn1336(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1337 github.com/goccy/llamawasm2go/p2.Fn1337
func Fn1337(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1338 github.com/goccy/llamawasm2go/p2.Fn1338
func Fn1338(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1339 github.com/goccy/llamawasm2go/p2.Fn1339
func Fn1339(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1341 github.com/goccy/llamawasm2go/p2.Fn1341
func Fn1341(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1345 github.com/goccy/llamawasm2go/p2.Fn1345
func Fn1345(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1346 github.com/goccy/llamawasm2go/p2.Fn1346
func Fn1346(m *base.Module)

//go:linkname Fn1347 github.com/goccy/llamawasm2go/p2.Fn1347
func Fn1347(m *base.Module)

//go:linkname Fn1348 github.com/goccy/llamawasm2go/p0.Fn1348
func Fn1348(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1357 github.com/goccy/llamawasm2go/p2.Fn1357
func Fn1357(m *base.Module)

//go:linkname Fn1359 github.com/goccy/llamawasm2go/p2.Fn1359
func Fn1359(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1361 github.com/goccy/llamawasm2go/p0.Fn1361
func Fn1361(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1362 github.com/goccy/llamawasm2go/p2.Fn1362
func Fn1362(m *base.Module, l0 int64) int64

//go:linkname Fn1367 github.com/goccy/llamawasm2go/p2.Fn1367
func Fn1367(m *base.Module, l0 int64)

//go:linkname Fn1374 github.com/goccy/llamawasm2go/p2.Fn1374
func Fn1374(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1377 github.com/goccy/llamawasm2go/p2.Fn1377
func Fn1377(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1379 github.com/goccy/llamawasm2go/p2.Fn1379
func Fn1379(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1381 github.com/goccy/llamawasm2go/p2.Fn1381
func Fn1381(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1383 github.com/goccy/llamawasm2go/p2.Fn1383
func Fn1383(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1385 github.com/goccy/llamawasm2go/p2.Fn1385
func Fn1385(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1387 github.com/goccy/llamawasm2go/p2.Fn1387
func Fn1387(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1393 github.com/goccy/llamawasm2go/p2.Fn1393
func Fn1393(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p2.Fn1394
func Fn1394(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p2.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1396 github.com/goccy/llamawasm2go/p0.Fn1396
func Fn1396(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64)

//go:linkname Fn1399 github.com/goccy/llamawasm2go/p2.Fn1399
func Fn1399(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p2.Fn1401
func Fn1401(m *base.Module, l0 int64) int64

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p2.Fn1402
func Fn1402(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module, l0 int64)

//go:linkname Fn1404 github.com/goccy/llamawasm2go/p2.Fn1404
func Fn1404(m *base.Module, l0 int64)

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p2.Fn1406
func Fn1406(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1411 github.com/goccy/llamawasm2go/p2.Fn1411
func Fn1411(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1419 github.com/goccy/llamawasm2go/p2.Fn1419
func Fn1419(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1421 github.com/goccy/llamawasm2go/p0.Fn1421
func Fn1421(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1423 github.com/goccy/llamawasm2go/p2.Fn1423
func Fn1423(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p0.Fn1426
func Fn1426(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1427 github.com/goccy/llamawasm2go/p2.Fn1427
func Fn1427(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1431 github.com/goccy/llamawasm2go/p2.Fn1431
func Fn1431(m *base.Module, l0 int64)

//go:linkname Fn1434 github.com/goccy/llamawasm2go/p2.Fn1434
func Fn1434(m *base.Module, l0 int64)

//go:linkname Fn1437 github.com/goccy/llamawasm2go/p2.Fn1437
func Fn1437(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1438 github.com/goccy/llamawasm2go/p2.Fn1438
func Fn1438(m *base.Module, l0 int64) int64

//go:linkname Fn1439 github.com/goccy/llamawasm2go/p2.Fn1439
func Fn1439(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1440 github.com/goccy/llamawasm2go/p2.Fn1440
func Fn1440(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1441 github.com/goccy/llamawasm2go/p2.Fn1441
func Fn1441(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64)

//go:linkname Fn1446 github.com/goccy/llamawasm2go/p2.Fn1446
func Fn1446(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1448 github.com/goccy/llamawasm2go/p2.Fn1448
func Fn1448(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1451 github.com/goccy/llamawasm2go/p2.Fn1451
func Fn1451(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1458 github.com/goccy/llamawasm2go/p2.Fn1458
func Fn1458(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1459 github.com/goccy/llamawasm2go/p2.Fn1459
func Fn1459(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1461 github.com/goccy/llamawasm2go/p2.Fn1461
func Fn1461(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1465 github.com/goccy/llamawasm2go/p2.Fn1465
func Fn1465(m *base.Module, l0 int64)

//go:linkname Fn1466 github.com/goccy/llamawasm2go/p2.Fn1466
func Fn1466(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1468 github.com/goccy/llamawasm2go/p2.Fn1468
func Fn1468(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1469 github.com/goccy/llamawasm2go/p2.Fn1469
func Fn1469(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1470 github.com/goccy/llamawasm2go/p2.Fn1470
func Fn1470(m *base.Module, l0 int64)

//go:linkname Fn1472 github.com/goccy/llamawasm2go/p2.Fn1472
func Fn1472(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1514 github.com/goccy/llamawasm2go/p2.Fn1514
func Fn1514(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1530 github.com/goccy/llamawasm2go/p2.Fn1530
func Fn1530(m *base.Module, l0 int64)

//go:linkname Fn1531 github.com/goccy/llamawasm2go/p2.Fn1531
func Fn1531(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1532 github.com/goccy/llamawasm2go/p2.Fn1532
func Fn1532(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1533 github.com/goccy/llamawasm2go/p2.Fn1533
func Fn1533(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1534 github.com/goccy/llamawasm2go/p2.Fn1534
func Fn1534(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1535 github.com/goccy/llamawasm2go/p2.Fn1535
func Fn1535(m *base.Module, l0 int64)

//go:linkname Fn1536 github.com/goccy/llamawasm2go/p2.Fn1536
func Fn1536(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1538 github.com/goccy/llamawasm2go/p2.Fn1538
func Fn1538(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1539 github.com/goccy/llamawasm2go/p2.Fn1539
func Fn1539(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1540 github.com/goccy/llamawasm2go/p2.Fn1540
func Fn1540(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p2.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p2.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1547 github.com/goccy/llamawasm2go/p2.Fn1547
func Fn1547(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1548 github.com/goccy/llamawasm2go/p2.Fn1548
func Fn1548(m *base.Module, l0 int64) int64

//go:linkname Fn1549 github.com/goccy/llamawasm2go/p2.Fn1549
func Fn1549(m *base.Module, l0 int64) int64

//go:linkname Fn1550 github.com/goccy/llamawasm2go/p2.Fn1550
func Fn1550(m *base.Module, l0 int64) int64

//go:linkname Fn1551 github.com/goccy/llamawasm2go/p2.Fn1551
func Fn1551(m *base.Module, l0 int64) int64

//go:linkname Fn1552 github.com/goccy/llamawasm2go/p2.Fn1552
func Fn1552(m *base.Module, l0 int64) int64

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64) int64

//go:linkname Fn1555 github.com/goccy/llamawasm2go/p2.Fn1555
func Fn1555(m *base.Module, l0 int64) int64

//go:linkname Fn1556 github.com/goccy/llamawasm2go/p2.Fn1556
func Fn1556(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1558 github.com/goccy/llamawasm2go/p2.Fn1558
func Fn1558(m *base.Module, l0 int64) int64

//go:linkname Fn1559 github.com/goccy/llamawasm2go/p2.Fn1559
func Fn1559(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1560 github.com/goccy/llamawasm2go/p2.Fn1560
func Fn1560(m *base.Module, l0 int64) int64

//go:linkname Fn1562 github.com/goccy/llamawasm2go/p2.Fn1562
func Fn1562(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p2.Fn1563
func Fn1563(m *base.Module, l0 int64) int64

//go:linkname Fn1565 github.com/goccy/llamawasm2go/p2.Fn1565
func Fn1565(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1566 github.com/goccy/llamawasm2go/p2.Fn1566
func Fn1566(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1567 github.com/goccy/llamawasm2go/p2.Fn1567
func Fn1567(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1568 github.com/goccy/llamawasm2go/p2.Fn1568
func Fn1568(m *base.Module, l0 int64) int64

//go:linkname Fn1570 github.com/goccy/llamawasm2go/p2.Fn1570
func Fn1570(m *base.Module, l0 int64) int64

//go:linkname Fn1571 github.com/goccy/llamawasm2go/p2.Fn1571
func Fn1571(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1576 github.com/goccy/llamawasm2go/p2.Fn1576
func Fn1576(m *base.Module, l0 int64) int64

//go:linkname Fn1577 github.com/goccy/llamawasm2go/p2.Fn1577
func Fn1577(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1578 github.com/goccy/llamawasm2go/p2.Fn1578
func Fn1578(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1579 github.com/goccy/llamawasm2go/p2.Fn1579
func Fn1579(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1581 github.com/goccy/llamawasm2go/p2.Fn1581
func Fn1581(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1582 github.com/goccy/llamawasm2go/p2.Fn1582
func Fn1582(m *base.Module, l0 int64) int64

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p2.Fn1583
func Fn1583(m *base.Module, l0 int64) int64

//go:linkname Fn1585 github.com/goccy/llamawasm2go/p2.Fn1585
func Fn1585(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1586 github.com/goccy/llamawasm2go/p2.Fn1586
func Fn1586(m *base.Module, l0 int64)

//go:linkname Fn1615 github.com/goccy/llamawasm2go/p2.Fn1615
func Fn1615(m *base.Module, l0 int64) int64

//go:linkname Fn1622 github.com/goccy/llamawasm2go/p2.Fn1622
func Fn1622(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1623 github.com/goccy/llamawasm2go/p2.Fn1623
func Fn1623(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1624 github.com/goccy/llamawasm2go/p2.Fn1624
func Fn1624(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1626 github.com/goccy/llamawasm2go/p2.Fn1626
func Fn1626(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1627 github.com/goccy/llamawasm2go/p2.Fn1627
func Fn1627(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1631 github.com/goccy/llamawasm2go/p2.Fn1631
func Fn1631(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1632 github.com/goccy/llamawasm2go/p2.Fn1632
func Fn1632(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1633 github.com/goccy/llamawasm2go/p2.Fn1633
func Fn1633(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1634 github.com/goccy/llamawasm2go/p2.Fn1634
func Fn1634(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1635 github.com/goccy/llamawasm2go/p2.Fn1635
func Fn1635(m *base.Module, l0 int64) int32

//go:linkname Fn1636 github.com/goccy/llamawasm2go/p2.Fn1636
func Fn1636(m *base.Module, l0 int64) int32

//go:linkname Fn1637 github.com/goccy/llamawasm2go/p2.Fn1637
func Fn1637(m *base.Module, l0 int64) int32

//go:linkname Fn1638 github.com/goccy/llamawasm2go/p2.Fn1638
func Fn1638(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1641 github.com/goccy/llamawasm2go/p2.Fn1641
func Fn1641(m *base.Module, l0 int64) int32

//go:linkname Fn1642 github.com/goccy/llamawasm2go/p2.Fn1642
func Fn1642(m *base.Module, l0 int64) int32

//go:linkname Fn1648 github.com/goccy/llamawasm2go/p2.Fn1648
func Fn1648(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1649 github.com/goccy/llamawasm2go/p2.Fn1649
func Fn1649(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1650 github.com/goccy/llamawasm2go/p2.Fn1650
func Fn1650(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1651 github.com/goccy/llamawasm2go/p2.Fn1651
func Fn1651(m *base.Module)

//go:linkname Fn1652 github.com/goccy/llamawasm2go/p2.Fn1652
func Fn1652(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1654 github.com/goccy/llamawasm2go/p2.Fn1654
func Fn1654(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1656 github.com/goccy/llamawasm2go/p2.Fn1656
func Fn1656(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1657 github.com/goccy/llamawasm2go/p2.Fn1657
func Fn1657(m *base.Module, l0 int64)

//go:linkname Fn1661 github.com/goccy/llamawasm2go/p2.Fn1661
func Fn1661(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1662 github.com/goccy/llamawasm2go/p2.Fn1662
func Fn1662(m *base.Module, l0 int64)

//go:linkname Fn1665 github.com/goccy/llamawasm2go/p2.Fn1665
func Fn1665(m *base.Module, l0 int64)

//go:linkname Fn1675 github.com/goccy/llamawasm2go/p2.Fn1675
func Fn1675(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1676 github.com/goccy/llamawasm2go/p2.Fn1676
func Fn1676(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1685 github.com/goccy/llamawasm2go/p2.Fn1685
func Fn1685(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1686 github.com/goccy/llamawasm2go/p0.Fn1686
func Fn1686(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1688 github.com/goccy/llamawasm2go/p2.Fn1688
func Fn1688(m *base.Module, l0 int64)

//go:linkname Fn1690 github.com/goccy/llamawasm2go/p2.Fn1690
func Fn1690(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1697 github.com/goccy/llamawasm2go/p2.Fn1697
func Fn1697(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1698 github.com/goccy/llamawasm2go/p2.Fn1698
func Fn1698(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1699 github.com/goccy/llamawasm2go/p2.Fn1699
func Fn1699(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1705 github.com/goccy/llamawasm2go/p2.Fn1705
func Fn1705(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1712 github.com/goccy/llamawasm2go/p2.Fn1712
func Fn1712(m *base.Module, l0 int64)

//go:linkname Fn1715 github.com/goccy/llamawasm2go/p2.Fn1715
func Fn1715(m *base.Module, l0 int64) int32

//go:linkname Fn1725 github.com/goccy/llamawasm2go/p2.Fn1725
func Fn1725(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1727 github.com/goccy/llamawasm2go/p2.Fn1727
func Fn1727(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1728 github.com/goccy/llamawasm2go/p2.Fn1728
func Fn1728(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1729 github.com/goccy/llamawasm2go/p2.Fn1729
func Fn1729(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1730 github.com/goccy/llamawasm2go/p2.Fn1730
func Fn1730(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1731 github.com/goccy/llamawasm2go/p2.Fn1731
func Fn1731(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1747 github.com/goccy/llamawasm2go/p2.Fn1747
func Fn1747(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1760 github.com/goccy/llamawasm2go/p2.Fn1760
func Fn1760(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1767 github.com/goccy/llamawasm2go/p2.Fn1767
func Fn1767(m *base.Module, l0 int64)

//go:linkname Fn1801 github.com/goccy/llamawasm2go/p2.Fn1801
func Fn1801(m *base.Module, l0 int64)

//go:linkname Fn1804 github.com/goccy/llamawasm2go/p2.Fn1804
func Fn1804(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1818 github.com/goccy/llamawasm2go/p2.Fn1818
func Fn1818(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1819 github.com/goccy/llamawasm2go/p2.Fn1819
func Fn1819(m *base.Module, l0 int64) int64

//go:linkname Fn1820 github.com/goccy/llamawasm2go/p2.Fn1820
func Fn1820(m *base.Module, l0 int64)

//go:linkname Fn1824 github.com/goccy/llamawasm2go/p0.Fn1824
func Fn1824(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1825 github.com/goccy/llamawasm2go/p2.Fn1825
func Fn1825(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1826 github.com/goccy/llamawasm2go/p2.Fn1826
func Fn1826(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1827 github.com/goccy/llamawasm2go/p2.Fn1827
func Fn1827(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1828 github.com/goccy/llamawasm2go/p2.Fn1828
func Fn1828(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1846 github.com/goccy/llamawasm2go/p2.Fn1846
func Fn1846(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1857 github.com/goccy/llamawasm2go/p2.Fn1857
func Fn1857(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1859 github.com/goccy/llamawasm2go/p2.Fn1859
func Fn1859(m *base.Module, l0 int64) int64

//go:linkname Fn1860 github.com/goccy/llamawasm2go/p2.Fn1860
func Fn1860(m *base.Module, l0 int64)

//go:linkname Fn1863 github.com/goccy/llamawasm2go/p0.Fn1863
func Fn1863(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1865 github.com/goccy/llamawasm2go/p2.Fn1865
func Fn1865(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1876 github.com/goccy/llamawasm2go/p2.Fn1876
func Fn1876(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1879 github.com/goccy/llamawasm2go/p2.Fn1879
func Fn1879(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1880 github.com/goccy/llamawasm2go/p2.Fn1880
func Fn1880(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1890 github.com/goccy/llamawasm2go/p2.Fn1890
func Fn1890(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1917 github.com/goccy/llamawasm2go/p2.Fn1917
func Fn1917(m *base.Module, l0 int64)

//go:linkname Fn1918 github.com/goccy/llamawasm2go/p2.Fn1918
func Fn1918(m *base.Module, l0 int64)

//go:linkname Fn1929 github.com/goccy/llamawasm2go/p2.Fn1929
func Fn1929(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1961 github.com/goccy/llamawasm2go/p2.Fn1961
func Fn1961(m *base.Module) int64

//go:linkname Fn1966 github.com/goccy/llamawasm2go/p2.Fn1966
func Fn1966(m *base.Module, l0 int64) int64

//go:linkname Fn1967 github.com/goccy/llamawasm2go/p2.Fn1967
func Fn1967(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1974 github.com/goccy/llamawasm2go/p2.Fn1974
func Fn1974(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1978 github.com/goccy/llamawasm2go/p2.Fn1978
func Fn1978(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1980 github.com/goccy/llamawasm2go/p2.Fn1980
func Fn1980(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1981 github.com/goccy/llamawasm2go/p2.Fn1981
func Fn1981(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1982 github.com/goccy/llamawasm2go/p2.Fn1982
func Fn1982(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1983 github.com/goccy/llamawasm2go/p2.Fn1983
func Fn1983(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1984 github.com/goccy/llamawasm2go/p2.Fn1984
func Fn1984(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1988 github.com/goccy/llamawasm2go/p2.Fn1988
func Fn1988(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1989 github.com/goccy/llamawasm2go/p2.Fn1989
func Fn1989(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1995 github.com/goccy/llamawasm2go/p2.Fn1995
func Fn1995(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1997 github.com/goccy/llamawasm2go/p2.Fn1997
func Fn1997(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2000 github.com/goccy/llamawasm2go/p2.Fn2000
func Fn2000(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2001 github.com/goccy/llamawasm2go/p2.Fn2001
func Fn2001(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2002 github.com/goccy/llamawasm2go/p2.Fn2002
func Fn2002(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn2005 github.com/goccy/llamawasm2go/p2.Fn2005
func Fn2005(m *base.Module, l0 int64)

//go:linkname Fn2015 github.com/goccy/llamawasm2go/p2.Fn2015
func Fn2015(m *base.Module, l0 int64)

//go:linkname Fn2017 github.com/goccy/llamawasm2go/p2.Fn2017
func Fn2017(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2018 github.com/goccy/llamawasm2go/p2.Fn2018
func Fn2018(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2031 github.com/goccy/llamawasm2go/p2.Fn2031
func Fn2031(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2032 github.com/goccy/llamawasm2go/p2.Fn2032
func Fn2032(m *base.Module, l0 int64) int64

//go:linkname Fn2033 github.com/goccy/llamawasm2go/p2.Fn2033
func Fn2033(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2035 github.com/goccy/llamawasm2go/p2.Fn2035
func Fn2035(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2049 github.com/goccy/llamawasm2go/p2.Fn2049
func Fn2049(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2050 github.com/goccy/llamawasm2go/p2.Fn2050
func Fn2050(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2052 github.com/goccy/llamawasm2go/p2.Fn2052
func Fn2052(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2053 github.com/goccy/llamawasm2go/p2.Fn2053
func Fn2053(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2054 github.com/goccy/llamawasm2go/p2.Fn2054
func Fn2054(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2055 github.com/goccy/llamawasm2go/p2.Fn2055
func Fn2055(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2056 github.com/goccy/llamawasm2go/p2.Fn2056
func Fn2056(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2057 github.com/goccy/llamawasm2go/p2.Fn2057
func Fn2057(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2076 github.com/goccy/llamawasm2go/p2.Fn2076
func Fn2076(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2077 github.com/goccy/llamawasm2go/p2.Fn2077
func Fn2077(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2078 github.com/goccy/llamawasm2go/p2.Fn2078
func Fn2078(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2080 github.com/goccy/llamawasm2go/p2.Fn2080
func Fn2080(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2081 github.com/goccy/llamawasm2go/p2.Fn2081
func Fn2081(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2082 github.com/goccy/llamawasm2go/p2.Fn2082
func Fn2082(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2083 github.com/goccy/llamawasm2go/p2.Fn2083
func Fn2083(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2084 github.com/goccy/llamawasm2go/p2.Fn2084
func Fn2084(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2086 github.com/goccy/llamawasm2go/p2.Fn2086
func Fn2086(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2088 github.com/goccy/llamawasm2go/p2.Fn2088
func Fn2088(m *base.Module, l0 int64)

//go:linkname Fn2105 github.com/goccy/llamawasm2go/p2.Fn2105
func Fn2105(m *base.Module, l0 int64)

//go:linkname Fn2106 github.com/goccy/llamawasm2go/p2.Fn2106
func Fn2106(m *base.Module, l0 int64)

//go:linkname Fn2107 github.com/goccy/llamawasm2go/p2.Fn2107
func Fn2107(m *base.Module, l0 int64)

//go:linkname Fn2109 github.com/goccy/llamawasm2go/p2.Fn2109
func Fn2109(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2110 github.com/goccy/llamawasm2go/p2.Fn2110
func Fn2110(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2149 github.com/goccy/llamawasm2go/p2.Fn2149
func Fn2149(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2158 github.com/goccy/llamawasm2go/p2.Fn2158
func Fn2158(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2160 github.com/goccy/llamawasm2go/p2.Fn2160
func Fn2160(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2162 github.com/goccy/llamawasm2go/p2.Fn2162
func Fn2162(m *base.Module, l0 int64) int64

//go:linkname Fn2165 github.com/goccy/llamawasm2go/p2.Fn2165
func Fn2165(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2169 github.com/goccy/llamawasm2go/p2.Fn2169
func Fn2169(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2174 github.com/goccy/llamawasm2go/p2.Fn2174
func Fn2174(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2188 github.com/goccy/llamawasm2go/p2.Fn2188
func Fn2188(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2190 github.com/goccy/llamawasm2go/p2.Fn2190
func Fn2190(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2191 github.com/goccy/llamawasm2go/p2.Fn2191
func Fn2191(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2205 github.com/goccy/llamawasm2go/p2.Fn2205
func Fn2205(m *base.Module, l0 int64) int64

//go:linkname Fn2207 github.com/goccy/llamawasm2go/p2.Fn2207
func Fn2207(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2210 github.com/goccy/llamawasm2go/p2.Fn2210
func Fn2210(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2214 github.com/goccy/llamawasm2go/p2.Fn2214
func Fn2214(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2215 github.com/goccy/llamawasm2go/p2.Fn2215
func Fn2215(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2216 github.com/goccy/llamawasm2go/p0.Fn2216
func Fn2216(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2218 github.com/goccy/llamawasm2go/p2.Fn2218
func Fn2218(m *base.Module, l0 int64) int64

//go:linkname Fn2224 github.com/goccy/llamawasm2go/p2.Fn2224
func Fn2224(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2228 github.com/goccy/llamawasm2go/p2.Fn2228
func Fn2228(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2229 github.com/goccy/llamawasm2go/p2.Fn2229
func Fn2229(m *base.Module, l0 int64)

//go:linkname Fn2230 github.com/goccy/llamawasm2go/p2.Fn2230
func Fn2230(m *base.Module, l0 int64)

//go:linkname Fn2231 github.com/goccy/llamawasm2go/p2.Fn2231
func Fn2231(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2233 github.com/goccy/llamawasm2go/p2.Fn2233
func Fn2233(m *base.Module, l0 int64)

//go:linkname Fn2234 github.com/goccy/llamawasm2go/p2.Fn2234
func Fn2234(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2235 github.com/goccy/llamawasm2go/p2.Fn2235
func Fn2235(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2239 github.com/goccy/llamawasm2go/p2.Fn2239
func Fn2239(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2242 github.com/goccy/llamawasm2go/p2.Fn2242
func Fn2242(m *base.Module, l0 int64) int64

//go:linkname Fn2243 github.com/goccy/llamawasm2go/p2.Fn2243
func Fn2243(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2244 github.com/goccy/llamawasm2go/p2.Fn2244
func Fn2244(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2251 github.com/goccy/llamawasm2go/p2.Fn2251
func Fn2251(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2252 github.com/goccy/llamawasm2go/p2.Fn2252
func Fn2252(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2254 github.com/goccy/llamawasm2go/p2.Fn2254
func Fn2254(m *base.Module) int64

//go:linkname Fn2256 github.com/goccy/llamawasm2go/p2.Fn2256
func Fn2256(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2257 github.com/goccy/llamawasm2go/p2.Fn2257
func Fn2257(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2259 github.com/goccy/llamawasm2go/p2.Fn2259
func Fn2259(m *base.Module) int64

//go:linkname Fn2261 github.com/goccy/llamawasm2go/p2.Fn2261
func Fn2261(m *base.Module, l0 int32) int64

//go:linkname Fn2262 github.com/goccy/llamawasm2go/p2.Fn2262
func Fn2262(m *base.Module, l0 int32) int32

//go:linkname Fn2263 github.com/goccy/llamawasm2go/p2.Fn2263
func Fn2263(m *base.Module, l0 int32) int64

//go:linkname Fn2264 github.com/goccy/llamawasm2go/p2.Fn2264
func Fn2264(m *base.Module, l0 float32) int64

//go:linkname Fn2265 github.com/goccy/llamawasm2go/p2.Fn2265
func Fn2265(m *base.Module, l0 float32) int64

//go:linkname Fn2266 github.com/goccy/llamawasm2go/p2.Fn2266
func Fn2266(m *base.Module, l0 float32) int64

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2269 github.com/goccy/llamawasm2go/p2.Fn2269
func Fn2269(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2310 github.com/goccy/llamawasm2go/p2.Fn2310
func Fn2310(m *base.Module, l0 int64)

//go:linkname Fn2312 github.com/goccy/llamawasm2go/p2.Fn2312
func Fn2312(m *base.Module, l0 int64)

//go:linkname Fn2354 github.com/goccy/llamawasm2go/p2.Fn2354
func Fn2354(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2365 github.com/goccy/llamawasm2go/p2.Fn2365
func Fn2365(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2368 github.com/goccy/llamawasm2go/p2.Fn2368
func Fn2368(m *base.Module, l0 int64)

//go:linkname Fn2369 github.com/goccy/llamawasm2go/p2.Fn2369
func Fn2369(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2370 github.com/goccy/llamawasm2go/p2.Fn2370
func Fn2370(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2372 github.com/goccy/llamawasm2go/p2.Fn2372
func Fn2372(m *base.Module, l0 int64)

//go:linkname Fn2376 github.com/goccy/llamawasm2go/p2.Fn2376
func Fn2376(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2382 github.com/goccy/llamawasm2go/p2.Fn2382
func Fn2382(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2397 github.com/goccy/llamawasm2go/p2.Fn2397
func Fn2397(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2404 github.com/goccy/llamawasm2go/p0.Fn2404
func Fn2404(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2407 github.com/goccy/llamawasm2go/p2.Fn2407
func Fn2407(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2410 github.com/goccy/llamawasm2go/p2.Fn2410
func Fn2410(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2413 github.com/goccy/llamawasm2go/p2.Fn2413
func Fn2413(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2414 github.com/goccy/llamawasm2go/p2.Fn2414
func Fn2414(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2415 github.com/goccy/llamawasm2go/p2.Fn2415
func Fn2415(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2420 github.com/goccy/llamawasm2go/p2.Fn2420
func Fn2420(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2422 github.com/goccy/llamawasm2go/p2.Fn2422
func Fn2422(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2439 github.com/goccy/llamawasm2go/p2.Fn2439
func Fn2439(m *base.Module, l0 int64)

//go:linkname Fn2440 github.com/goccy/llamawasm2go/p2.Fn2440
func Fn2440(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2445 github.com/goccy/llamawasm2go/p2.Fn2445
func Fn2445(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2450 github.com/goccy/llamawasm2go/p2.Fn2450
func Fn2450(m *base.Module, l0 int64) int64

//go:linkname Fn2451 github.com/goccy/llamawasm2go/p2.Fn2451
func Fn2451(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2452 github.com/goccy/llamawasm2go/p0.Fn2452
func Fn2452(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2463 github.com/goccy/llamawasm2go/p2.Fn2463
func Fn2463(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2464 github.com/goccy/llamawasm2go/p2.Fn2464
func Fn2464(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2465 github.com/goccy/llamawasm2go/p2.Fn2465
func Fn2465(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2466 github.com/goccy/llamawasm2go/p2.Fn2466
func Fn2466(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2486 github.com/goccy/llamawasm2go/p2.Fn2486
func Fn2486(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2487 github.com/goccy/llamawasm2go/p2.Fn2487
func Fn2487(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2490 github.com/goccy/llamawasm2go/p2.Fn2490
func Fn2490(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2576 github.com/goccy/llamawasm2go/p2.Fn2576
func Fn2576(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2695 github.com/goccy/llamawasm2go/p2.Fn2695
func Fn2695(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2696 github.com/goccy/llamawasm2go/p0.Fn2696
func Fn2696(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2697 github.com/goccy/llamawasm2go/p2.Fn2697
func Fn2697(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2698 github.com/goccy/llamawasm2go/p2.Fn2698
func Fn2698(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2702 github.com/goccy/llamawasm2go/p2.Fn2702
func Fn2702(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2727 github.com/goccy/llamawasm2go/p2.Fn2727
func Fn2727(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2859 github.com/goccy/llamawasm2go/p2.Fn2859
func Fn2859(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2866 github.com/goccy/llamawasm2go/p2.Fn2866
func Fn2866(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2894 github.com/goccy/llamawasm2go/p2.Fn2894
func Fn2894(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2908 github.com/goccy/llamawasm2go/p2.Fn2908
func Fn2908(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2920 github.com/goccy/llamawasm2go/p2.Fn2920
func Fn2920(m *base.Module, l0 int32)

//go:linkname Fn2922 github.com/goccy/llamawasm2go/p2.Fn2922
func Fn2922(m *base.Module, l0 int64) int64

//go:linkname Fn2923 github.com/goccy/llamawasm2go/p2.Fn2923
func Fn2923(m *base.Module, l0 int64)

//go:linkname Fn2926 github.com/goccy/llamawasm2go/p2.Fn2926
func Fn2926(m *base.Module, l0 int64)

//go:linkname Fn2927 github.com/goccy/llamawasm2go/p2.Fn2927
func Fn2927(m *base.Module, l0 int64)

//go:linkname Fn2929 github.com/goccy/llamawasm2go/p2.Fn2929
func Fn2929(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2930 github.com/goccy/llamawasm2go/p2.Fn2930
func Fn2930(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2936 github.com/goccy/llamawasm2go/p2.Fn2936
func Fn2936(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2938 github.com/goccy/llamawasm2go/p2.Fn2938
func Fn2938(m *base.Module, l0 int64) int32

//go:linkname Fn2942 github.com/goccy/llamawasm2go/p2.Fn2942
func Fn2942(m *base.Module) int32

//go:linkname Fn2953 github.com/goccy/llamawasm2go/p2.Fn2953
func Fn2953(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2957 github.com/goccy/llamawasm2go/p2.Fn2957
func Fn2957(m *base.Module) int64

//go:linkname Fn2959 github.com/goccy/llamawasm2go/p2.Fn2959
func Fn2959(m *base.Module, l0 float64) float32

//go:linkname Fn2960 github.com/goccy/llamawasm2go/p2.Fn2960
func Fn2960(m *base.Module, l0 float64) float32

//go:linkname Fn2964 github.com/goccy/llamawasm2go/p2.Fn2964
func Fn2964(m *base.Module, l0 float64) float64

//go:linkname Fn2967 github.com/goccy/llamawasm2go/p2.Fn2967
func Fn2967(m *base.Module, l0 int32) float32

//go:linkname Fn2968 github.com/goccy/llamawasm2go/p2.Fn2968
func Fn2968(m *base.Module, l0 int32) float32

//go:linkname Fn2971 github.com/goccy/llamawasm2go/p2.Fn2971
func Fn2971(m *base.Module, l0 float32) float32

//go:linkname Fn2974 github.com/goccy/llamawasm2go/p2.Fn2974
func Fn2974(m *base.Module, l0 float64) float64

//go:linkname Fn2975 github.com/goccy/llamawasm2go/p2.Fn2975
func Fn2975(m *base.Module, l0 float64) float64

//go:linkname Fn2976 github.com/goccy/llamawasm2go/p2.Fn2976
func Fn2976(m *base.Module, l0 float32) float32

//go:linkname Fn2978 github.com/goccy/llamawasm2go/p2.Fn2978
func Fn2978(m *base.Module, l0 float32) float32

//go:linkname Fn2980 github.com/goccy/llamawasm2go/p2.Fn2980
func Fn2980(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn2981 github.com/goccy/llamawasm2go/p2.Fn2981
func Fn2981(m *base.Module, l0 float32) float32

//go:linkname Fn2998 github.com/goccy/llamawasm2go/p2.Fn2998
func Fn2998(m *base.Module, l0 int64) int32

//go:linkname Fn2999 github.com/goccy/llamawasm2go/p2.Fn2999
func Fn2999(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3001 github.com/goccy/llamawasm2go/p2.Fn3001
func Fn3001(m *base.Module, l0 int64)

//go:linkname Fn3002 github.com/goccy/llamawasm2go/p2.Fn3002
func Fn3002(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3003 github.com/goccy/llamawasm2go/p2.Fn3003
func Fn3003(m *base.Module, l0 int64) int32

//go:linkname Fn3010 github.com/goccy/llamawasm2go/p2.Fn3010
func Fn3010(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3012 github.com/goccy/llamawasm2go/p2.Fn3012
func Fn3012(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3018 github.com/goccy/llamawasm2go/p2.Fn3018
func Fn3018(m *base.Module, l0 int64) int32

//go:linkname Fn3021 github.com/goccy/llamawasm2go/p2.Fn3021
func Fn3021(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3024 github.com/goccy/llamawasm2go/p2.Fn3024
func Fn3024(m *base.Module, l0 int64) int32

//go:linkname Fn3026 github.com/goccy/llamawasm2go/p2.Fn3026
func Fn3026(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3028 github.com/goccy/llamawasm2go/p2.Fn3028
func Fn3028(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3029 github.com/goccy/llamawasm2go/p2.Fn3029
func Fn3029(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3032 github.com/goccy/llamawasm2go/p2.Fn3032
func Fn3032(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3035 github.com/goccy/llamawasm2go/p2.Fn3035
func Fn3035(m *base.Module, l0 int64) int64

//go:linkname Fn3039 github.com/goccy/llamawasm2go/p2.Fn3039
func Fn3039(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3049 github.com/goccy/llamawasm2go/p2.Fn3049
func Fn3049(m *base.Module)

//go:linkname Fn3050 github.com/goccy/llamawasm2go/p0.Fn3050
func Fn3050(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3053 github.com/goccy/llamawasm2go/p2.Fn3053
func Fn3053(m *base.Module)

//go:linkname Fn3055 github.com/goccy/llamawasm2go/p0.Fn3055
func Fn3055(m *base.Module, l0 int64) int64

//go:linkname Fn3057 github.com/goccy/llamawasm2go/p2.Fn3057
func Fn3057(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3061 github.com/goccy/llamawasm2go/p2.Fn3061
func Fn3061(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3130 github.com/goccy/llamawasm2go/p2.Fn3130
func Fn3130(m *base.Module, l0 int32)

//go:linkname Fn957rows github.com/goccy/llamawasm2go/p2.Fn957rows
func Fn957rows(m *base.Module)
