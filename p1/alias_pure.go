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

//go:linkname Fn465 github.com/goccy/llamawasm2go/p2.Fn465
func Fn465(m *base.Module, l0 int64, l1 int64) int64

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

//go:linkname Fn524 github.com/goccy/llamawasm2go/p2.Fn524
func Fn524(m *base.Module, l0 int64, l1 int64, l2 int32) int64

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

//go:linkname Fn896 github.com/goccy/llamawasm2go/p2.Fn896
func Fn896(m *base.Module) int64

//go:linkname Fn965 github.com/goccy/llamawasm2go/p2.Fn965
func Fn965(m *base.Module, l0 int64) int64

//go:linkname Fn975 github.com/goccy/llamawasm2go/p2.Fn975
func Fn975(m *base.Module, l0 int64)

//go:linkname Fn980 github.com/goccy/llamawasm2go/p2.Fn980
func Fn980(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn981 github.com/goccy/llamawasm2go/p2.Fn981
func Fn981(m *base.Module)

//go:linkname Fn990 github.com/goccy/llamawasm2go/p2.Fn990
func Fn990(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn998 github.com/goccy/llamawasm2go/p2.Fn998
func Fn998(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn999 github.com/goccy/llamawasm2go/p2.Fn999
func Fn999(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1000 github.com/goccy/llamawasm2go/p2.Fn1000
func Fn1000(m *base.Module, l0 int64, l1 int64, l2 int64) float32

//go:linkname Fn1002 github.com/goccy/llamawasm2go/p2.Fn1002
func Fn1002(m *base.Module, l0 int64, l1 int64, l2 int64) float64

//go:linkname Fn1004 github.com/goccy/llamawasm2go/p2.Fn1004
func Fn1004(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1007 github.com/goccy/llamawasm2go/p2.Fn1007
func Fn1007(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1010 github.com/goccy/llamawasm2go/p2.Fn1010
func Fn1010(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1016 github.com/goccy/llamawasm2go/p2.Fn1016
func Fn1016(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1017 github.com/goccy/llamawasm2go/p2.Fn1017
func Fn1017(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1025 github.com/goccy/llamawasm2go/p2.Fn1025
func Fn1025(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1030 github.com/goccy/llamawasm2go/p2.Fn1030
func Fn1030(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1031 github.com/goccy/llamawasm2go/p2.Fn1031
func Fn1031(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1039 github.com/goccy/llamawasm2go/p2.Fn1039
func Fn1039(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1055 github.com/goccy/llamawasm2go/p2.Fn1055
func Fn1055(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1063 github.com/goccy/llamawasm2go/p2.Fn1063
func Fn1063(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1079 github.com/goccy/llamawasm2go/p0.Fn1079
func Fn1079(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1082 github.com/goccy/llamawasm2go/p2.Fn1082
func Fn1082(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1086 github.com/goccy/llamawasm2go/p2.Fn1086
func Fn1086(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1096 github.com/goccy/llamawasm2go/p0.Fn1096
func Fn1096(m *base.Module, l0 int64) int64

//go:linkname Fn1097 github.com/goccy/llamawasm2go/p2.Fn1097
func Fn1097(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1098 github.com/goccy/llamawasm2go/p2.Fn1098
func Fn1098(m *base.Module, l0 int64)

//go:linkname Fn1100 github.com/goccy/llamawasm2go/p2.Fn1100
func Fn1100(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1190 github.com/goccy/llamawasm2go/p2.Fn1190
func Fn1190(m *base.Module, l0 int64)

//go:linkname Fn1212 github.com/goccy/llamawasm2go/p2.Fn1212
func Fn1212(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1222 github.com/goccy/llamawasm2go/p2.Fn1222
func Fn1222(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1228 github.com/goccy/llamawasm2go/p2.Fn1228
func Fn1228(m *base.Module)

//go:linkname Fn1232 github.com/goccy/llamawasm2go/p2.Fn1232
func Fn1232(m *base.Module, l0 int64) int64

//go:linkname Fn1260 github.com/goccy/llamawasm2go/p2.Fn1260
func Fn1260(m *base.Module, l0 int64)

//go:linkname Fn1264 github.com/goccy/llamawasm2go/p2.Fn1264
func Fn1264(m *base.Module, l0 int32) int64

//go:linkname Fn1275 github.com/goccy/llamawasm2go/p2.Fn1275
func Fn1275(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1276 github.com/goccy/llamawasm2go/p2.Fn1276
func Fn1276(m *base.Module, l0 int64)

//go:linkname Fn1278 github.com/goccy/llamawasm2go/p2.Fn1278
func Fn1278(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1299 github.com/goccy/llamawasm2go/p2.Fn1299
func Fn1299(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1300 github.com/goccy/llamawasm2go/p2.Fn1300
func Fn1300(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1301 github.com/goccy/llamawasm2go/p2.Fn1301
func Fn1301(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1302 github.com/goccy/llamawasm2go/p2.Fn1302
func Fn1302(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1303 github.com/goccy/llamawasm2go/p2.Fn1303
func Fn1303(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1304 github.com/goccy/llamawasm2go/p2.Fn1304
func Fn1304(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1305 github.com/goccy/llamawasm2go/p2.Fn1305
func Fn1305(m *base.Module, l0 int64) int64

//go:linkname Fn1308 github.com/goccy/llamawasm2go/p2.Fn1308
func Fn1308(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1312 github.com/goccy/llamawasm2go/p2.Fn1312
func Fn1312(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1313 github.com/goccy/llamawasm2go/p2.Fn1313
func Fn1313(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1314 github.com/goccy/llamawasm2go/p2.Fn1314
func Fn1314(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1315 github.com/goccy/llamawasm2go/p2.Fn1315
func Fn1315(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1316 github.com/goccy/llamawasm2go/p2.Fn1316
func Fn1316(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1317 github.com/goccy/llamawasm2go/p2.Fn1317
func Fn1317(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1318 github.com/goccy/llamawasm2go/p2.Fn1318
func Fn1318(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1319 github.com/goccy/llamawasm2go/p2.Fn1319
func Fn1319(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1320 github.com/goccy/llamawasm2go/p2.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1321 github.com/goccy/llamawasm2go/p2.Fn1321
func Fn1321(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1322 github.com/goccy/llamawasm2go/p2.Fn1322
func Fn1322(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p2.Fn1324
func Fn1324(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1325 github.com/goccy/llamawasm2go/p2.Fn1325
func Fn1325(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1326 github.com/goccy/llamawasm2go/p2.Fn1326
func Fn1326(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1327 github.com/goccy/llamawasm2go/p2.Fn1327
func Fn1327(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1328 github.com/goccy/llamawasm2go/p2.Fn1328
func Fn1328(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1329 github.com/goccy/llamawasm2go/p2.Fn1329
func Fn1329(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1330 github.com/goccy/llamawasm2go/p2.Fn1330
func Fn1330(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1331 github.com/goccy/llamawasm2go/p2.Fn1331
func Fn1331(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1332 github.com/goccy/llamawasm2go/p2.Fn1332
func Fn1332(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1333 github.com/goccy/llamawasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1335 github.com/goccy/llamawasm2go/p2.Fn1335
func Fn1335(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1336 github.com/goccy/llamawasm2go/p2.Fn1336
func Fn1336(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1337 github.com/goccy/llamawasm2go/p2.Fn1337
func Fn1337(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1338 github.com/goccy/llamawasm2go/p2.Fn1338
func Fn1338(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1340 github.com/goccy/llamawasm2go/p2.Fn1340
func Fn1340(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1344 github.com/goccy/llamawasm2go/p2.Fn1344
func Fn1344(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1345 github.com/goccy/llamawasm2go/p2.Fn1345
func Fn1345(m *base.Module)

//go:linkname Fn1346 github.com/goccy/llamawasm2go/p2.Fn1346
func Fn1346(m *base.Module)

//go:linkname Fn1347 github.com/goccy/llamawasm2go/p0.Fn1347
func Fn1347(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1356 github.com/goccy/llamawasm2go/p2.Fn1356
func Fn1356(m *base.Module)

//go:linkname Fn1358 github.com/goccy/llamawasm2go/p2.Fn1358
func Fn1358(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1360 github.com/goccy/llamawasm2go/p0.Fn1360
func Fn1360(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1361 github.com/goccy/llamawasm2go/p2.Fn1361
func Fn1361(m *base.Module, l0 int64) int64

//go:linkname Fn1366 github.com/goccy/llamawasm2go/p2.Fn1366
func Fn1366(m *base.Module, l0 int64)

//go:linkname Fn1373 github.com/goccy/llamawasm2go/p2.Fn1373
func Fn1373(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1376 github.com/goccy/llamawasm2go/p2.Fn1376
func Fn1376(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1378 github.com/goccy/llamawasm2go/p2.Fn1378
func Fn1378(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1380 github.com/goccy/llamawasm2go/p2.Fn1380
func Fn1380(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1382 github.com/goccy/llamawasm2go/p2.Fn1382
func Fn1382(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1384 github.com/goccy/llamawasm2go/p2.Fn1384
func Fn1384(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1386 github.com/goccy/llamawasm2go/p2.Fn1386
func Fn1386(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p2.Fn1392
func Fn1392(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1393 github.com/goccy/llamawasm2go/p2.Fn1393
func Fn1393(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p2.Fn1394
func Fn1394(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p0.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p2.Fn1397
func Fn1397(m *base.Module, l0 int64)

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64) int64

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p2.Fn1401
func Fn1401(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p2.Fn1402
func Fn1402(m *base.Module, l0 int64)

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module, l0 int64)

//go:linkname Fn1404 github.com/goccy/llamawasm2go/p2.Fn1404
func Fn1404(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1410 github.com/goccy/llamawasm2go/p2.Fn1410
func Fn1410(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1418 github.com/goccy/llamawasm2go/p2.Fn1418
func Fn1418(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p0.Fn1420
func Fn1420(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1422 github.com/goccy/llamawasm2go/p2.Fn1422
func Fn1422(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1425 github.com/goccy/llamawasm2go/p0.Fn1425
func Fn1425(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p2.Fn1426
func Fn1426(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1430 github.com/goccy/llamawasm2go/p2.Fn1430
func Fn1430(m *base.Module, l0 int64)

//go:linkname Fn1433 github.com/goccy/llamawasm2go/p2.Fn1433
func Fn1433(m *base.Module, l0 int64)

//go:linkname Fn1436 github.com/goccy/llamawasm2go/p2.Fn1436
func Fn1436(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1437 github.com/goccy/llamawasm2go/p2.Fn1437
func Fn1437(m *base.Module, l0 int64) int64

//go:linkname Fn1438 github.com/goccy/llamawasm2go/p2.Fn1438
func Fn1438(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1439 github.com/goccy/llamawasm2go/p2.Fn1439
func Fn1439(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1440 github.com/goccy/llamawasm2go/p2.Fn1440
func Fn1440(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1441 github.com/goccy/llamawasm2go/p2.Fn1441
func Fn1441(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int64)

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1450 github.com/goccy/llamawasm2go/p2.Fn1450
func Fn1450(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1457 github.com/goccy/llamawasm2go/p2.Fn1457
func Fn1457(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1458 github.com/goccy/llamawasm2go/p2.Fn1458
func Fn1458(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1460 github.com/goccy/llamawasm2go/p2.Fn1460
func Fn1460(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1464 github.com/goccy/llamawasm2go/p2.Fn1464
func Fn1464(m *base.Module, l0 int64)

//go:linkname Fn1465 github.com/goccy/llamawasm2go/p2.Fn1465
func Fn1465(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1467 github.com/goccy/llamawasm2go/p2.Fn1467
func Fn1467(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1468 github.com/goccy/llamawasm2go/p2.Fn1468
func Fn1468(m *base.Module, l0 int64)

//go:linkname Fn1470 github.com/goccy/llamawasm2go/p2.Fn1470
func Fn1470(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1512 github.com/goccy/llamawasm2go/p2.Fn1512
func Fn1512(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1528 github.com/goccy/llamawasm2go/p2.Fn1528
func Fn1528(m *base.Module, l0 int64)

//go:linkname Fn1529 github.com/goccy/llamawasm2go/p2.Fn1529
func Fn1529(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1530 github.com/goccy/llamawasm2go/p2.Fn1530
func Fn1530(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1531 github.com/goccy/llamawasm2go/p2.Fn1531
func Fn1531(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1532 github.com/goccy/llamawasm2go/p2.Fn1532
func Fn1532(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1533 github.com/goccy/llamawasm2go/p2.Fn1533
func Fn1533(m *base.Module, l0 int64)

//go:linkname Fn1534 github.com/goccy/llamawasm2go/p2.Fn1534
func Fn1534(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1536 github.com/goccy/llamawasm2go/p2.Fn1536
func Fn1536(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1537 github.com/goccy/llamawasm2go/p2.Fn1537
func Fn1537(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1538 github.com/goccy/llamawasm2go/p2.Fn1538
func Fn1538(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1539 github.com/goccy/llamawasm2go/p2.Fn1539
func Fn1539(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1540 github.com/goccy/llamawasm2go/p2.Fn1540
func Fn1540(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1543 github.com/goccy/llamawasm2go/p2.Fn1543
func Fn1543(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1546 github.com/goccy/llamawasm2go/p2.Fn1546
func Fn1546(m *base.Module, l0 int64) int64

//go:linkname Fn1547 github.com/goccy/llamawasm2go/p2.Fn1547
func Fn1547(m *base.Module, l0 int64) int64

//go:linkname Fn1548 github.com/goccy/llamawasm2go/p2.Fn1548
func Fn1548(m *base.Module, l0 int64) int64

//go:linkname Fn1549 github.com/goccy/llamawasm2go/p2.Fn1549
func Fn1549(m *base.Module, l0 int64) int64

//go:linkname Fn1550 github.com/goccy/llamawasm2go/p2.Fn1550
func Fn1550(m *base.Module, l0 int64) int64

//go:linkname Fn1551 github.com/goccy/llamawasm2go/p2.Fn1551
func Fn1551(m *base.Module, l0 int64) int64

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64) int64

//go:linkname Fn1554 github.com/goccy/llamawasm2go/p2.Fn1554
func Fn1554(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1556 github.com/goccy/llamawasm2go/p2.Fn1556
func Fn1556(m *base.Module, l0 int64) int64

//go:linkname Fn1557 github.com/goccy/llamawasm2go/p2.Fn1557
func Fn1557(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1558 github.com/goccy/llamawasm2go/p2.Fn1558
func Fn1558(m *base.Module, l0 int64) int64

//go:linkname Fn1560 github.com/goccy/llamawasm2go/p2.Fn1560
func Fn1560(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1561 github.com/goccy/llamawasm2go/p2.Fn1561
func Fn1561(m *base.Module, l0 int64) int64

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p2.Fn1563
func Fn1563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1564 github.com/goccy/llamawasm2go/p2.Fn1564
func Fn1564(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1565 github.com/goccy/llamawasm2go/p2.Fn1565
func Fn1565(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1566 github.com/goccy/llamawasm2go/p2.Fn1566
func Fn1566(m *base.Module, l0 int64) int64

//go:linkname Fn1568 github.com/goccy/llamawasm2go/p2.Fn1568
func Fn1568(m *base.Module, l0 int64) int64

//go:linkname Fn1569 github.com/goccy/llamawasm2go/p2.Fn1569
func Fn1569(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1574 github.com/goccy/llamawasm2go/p2.Fn1574
func Fn1574(m *base.Module, l0 int64) int64

//go:linkname Fn1575 github.com/goccy/llamawasm2go/p2.Fn1575
func Fn1575(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1576 github.com/goccy/llamawasm2go/p2.Fn1576
func Fn1576(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1577 github.com/goccy/llamawasm2go/p2.Fn1577
func Fn1577(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1579 github.com/goccy/llamawasm2go/p2.Fn1579
func Fn1579(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1580 github.com/goccy/llamawasm2go/p2.Fn1580
func Fn1580(m *base.Module, l0 int64) int64

//go:linkname Fn1581 github.com/goccy/llamawasm2go/p2.Fn1581
func Fn1581(m *base.Module, l0 int64) int64

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p2.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1584 github.com/goccy/llamawasm2go/p2.Fn1584
func Fn1584(m *base.Module, l0 int64)

//go:linkname Fn1613 github.com/goccy/llamawasm2go/p2.Fn1613
func Fn1613(m *base.Module, l0 int64) int64

//go:linkname Fn1620 github.com/goccy/llamawasm2go/p2.Fn1620
func Fn1620(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1621 github.com/goccy/llamawasm2go/p2.Fn1621
func Fn1621(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1622 github.com/goccy/llamawasm2go/p2.Fn1622
func Fn1622(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1624 github.com/goccy/llamawasm2go/p2.Fn1624
func Fn1624(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1625 github.com/goccy/llamawasm2go/p2.Fn1625
func Fn1625(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1629 github.com/goccy/llamawasm2go/p2.Fn1629
func Fn1629(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1630 github.com/goccy/llamawasm2go/p2.Fn1630
func Fn1630(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1631 github.com/goccy/llamawasm2go/p2.Fn1631
func Fn1631(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1632 github.com/goccy/llamawasm2go/p2.Fn1632
func Fn1632(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1633 github.com/goccy/llamawasm2go/p2.Fn1633
func Fn1633(m *base.Module, l0 int64) int32

//go:linkname Fn1634 github.com/goccy/llamawasm2go/p2.Fn1634
func Fn1634(m *base.Module, l0 int64) int32

//go:linkname Fn1635 github.com/goccy/llamawasm2go/p2.Fn1635
func Fn1635(m *base.Module, l0 int64) int32

//go:linkname Fn1636 github.com/goccy/llamawasm2go/p2.Fn1636
func Fn1636(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1639 github.com/goccy/llamawasm2go/p2.Fn1639
func Fn1639(m *base.Module, l0 int64) int32

//go:linkname Fn1640 github.com/goccy/llamawasm2go/p2.Fn1640
func Fn1640(m *base.Module, l0 int64) int32

//go:linkname Fn1646 github.com/goccy/llamawasm2go/p2.Fn1646
func Fn1646(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1647 github.com/goccy/llamawasm2go/p2.Fn1647
func Fn1647(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1648 github.com/goccy/llamawasm2go/p2.Fn1648
func Fn1648(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1649 github.com/goccy/llamawasm2go/p2.Fn1649
func Fn1649(m *base.Module)

//go:linkname Fn1650 github.com/goccy/llamawasm2go/p2.Fn1650
func Fn1650(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1652 github.com/goccy/llamawasm2go/p2.Fn1652
func Fn1652(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1654 github.com/goccy/llamawasm2go/p2.Fn1654
func Fn1654(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1655 github.com/goccy/llamawasm2go/p2.Fn1655
func Fn1655(m *base.Module, l0 int64)

//go:linkname Fn1659 github.com/goccy/llamawasm2go/p2.Fn1659
func Fn1659(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1660 github.com/goccy/llamawasm2go/p2.Fn1660
func Fn1660(m *base.Module, l0 int64)

//go:linkname Fn1663 github.com/goccy/llamawasm2go/p2.Fn1663
func Fn1663(m *base.Module, l0 int64)

//go:linkname Fn1673 github.com/goccy/llamawasm2go/p2.Fn1673
func Fn1673(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1674 github.com/goccy/llamawasm2go/p2.Fn1674
func Fn1674(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1683 github.com/goccy/llamawasm2go/p2.Fn1683
func Fn1683(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1684 github.com/goccy/llamawasm2go/p0.Fn1684
func Fn1684(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1686 github.com/goccy/llamawasm2go/p2.Fn1686
func Fn1686(m *base.Module, l0 int64)

//go:linkname Fn1688 github.com/goccy/llamawasm2go/p2.Fn1688
func Fn1688(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1695 github.com/goccy/llamawasm2go/p2.Fn1695
func Fn1695(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1696 github.com/goccy/llamawasm2go/p2.Fn1696
func Fn1696(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1697 github.com/goccy/llamawasm2go/p2.Fn1697
func Fn1697(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1703 github.com/goccy/llamawasm2go/p2.Fn1703
func Fn1703(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1710 github.com/goccy/llamawasm2go/p2.Fn1710
func Fn1710(m *base.Module, l0 int64)

//go:linkname Fn1713 github.com/goccy/llamawasm2go/p2.Fn1713
func Fn1713(m *base.Module, l0 int64) int32

//go:linkname Fn1723 github.com/goccy/llamawasm2go/p2.Fn1723
func Fn1723(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1725 github.com/goccy/llamawasm2go/p2.Fn1725
func Fn1725(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1726 github.com/goccy/llamawasm2go/p2.Fn1726
func Fn1726(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1727 github.com/goccy/llamawasm2go/p2.Fn1727
func Fn1727(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1728 github.com/goccy/llamawasm2go/p2.Fn1728
func Fn1728(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1729 github.com/goccy/llamawasm2go/p2.Fn1729
func Fn1729(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1745 github.com/goccy/llamawasm2go/p2.Fn1745
func Fn1745(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1758 github.com/goccy/llamawasm2go/p2.Fn1758
func Fn1758(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1765 github.com/goccy/llamawasm2go/p2.Fn1765
func Fn1765(m *base.Module, l0 int64)

//go:linkname Fn1799 github.com/goccy/llamawasm2go/p2.Fn1799
func Fn1799(m *base.Module, l0 int64)

//go:linkname Fn1802 github.com/goccy/llamawasm2go/p2.Fn1802
func Fn1802(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1816 github.com/goccy/llamawasm2go/p2.Fn1816
func Fn1816(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1817 github.com/goccy/llamawasm2go/p2.Fn1817
func Fn1817(m *base.Module, l0 int64) int64

//go:linkname Fn1818 github.com/goccy/llamawasm2go/p2.Fn1818
func Fn1818(m *base.Module, l0 int64)

//go:linkname Fn1822 github.com/goccy/llamawasm2go/p0.Fn1822
func Fn1822(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1823 github.com/goccy/llamawasm2go/p2.Fn1823
func Fn1823(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1824 github.com/goccy/llamawasm2go/p2.Fn1824
func Fn1824(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1825 github.com/goccy/llamawasm2go/p2.Fn1825
func Fn1825(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1826 github.com/goccy/llamawasm2go/p2.Fn1826
func Fn1826(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1844 github.com/goccy/llamawasm2go/p2.Fn1844
func Fn1844(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1855 github.com/goccy/llamawasm2go/p2.Fn1855
func Fn1855(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1857 github.com/goccy/llamawasm2go/p2.Fn1857
func Fn1857(m *base.Module, l0 int64) int64

//go:linkname Fn1858 github.com/goccy/llamawasm2go/p2.Fn1858
func Fn1858(m *base.Module, l0 int64)

//go:linkname Fn1861 github.com/goccy/llamawasm2go/p0.Fn1861
func Fn1861(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1863 github.com/goccy/llamawasm2go/p2.Fn1863
func Fn1863(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1874 github.com/goccy/llamawasm2go/p2.Fn1874
func Fn1874(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1877 github.com/goccy/llamawasm2go/p2.Fn1877
func Fn1877(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1878 github.com/goccy/llamawasm2go/p2.Fn1878
func Fn1878(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1888 github.com/goccy/llamawasm2go/p2.Fn1888
func Fn1888(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1915 github.com/goccy/llamawasm2go/p2.Fn1915
func Fn1915(m *base.Module, l0 int64)

//go:linkname Fn1916 github.com/goccy/llamawasm2go/p2.Fn1916
func Fn1916(m *base.Module, l0 int64)

//go:linkname Fn1927 github.com/goccy/llamawasm2go/p2.Fn1927
func Fn1927(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1959 github.com/goccy/llamawasm2go/p2.Fn1959
func Fn1959(m *base.Module) int64

//go:linkname Fn1964 github.com/goccy/llamawasm2go/p2.Fn1964
func Fn1964(m *base.Module, l0 int64) int64

//go:linkname Fn1965 github.com/goccy/llamawasm2go/p2.Fn1965
func Fn1965(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1972 github.com/goccy/llamawasm2go/p2.Fn1972
func Fn1972(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1976 github.com/goccy/llamawasm2go/p2.Fn1976
func Fn1976(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1978 github.com/goccy/llamawasm2go/p2.Fn1978
func Fn1978(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1979 github.com/goccy/llamawasm2go/p2.Fn1979
func Fn1979(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1980 github.com/goccy/llamawasm2go/p2.Fn1980
func Fn1980(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1981 github.com/goccy/llamawasm2go/p2.Fn1981
func Fn1981(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1982 github.com/goccy/llamawasm2go/p2.Fn1982
func Fn1982(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1986 github.com/goccy/llamawasm2go/p2.Fn1986
func Fn1986(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1987 github.com/goccy/llamawasm2go/p2.Fn1987
func Fn1987(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1993 github.com/goccy/llamawasm2go/p2.Fn1993
func Fn1993(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1995 github.com/goccy/llamawasm2go/p2.Fn1995
func Fn1995(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1998 github.com/goccy/llamawasm2go/p2.Fn1998
func Fn1998(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1999 github.com/goccy/llamawasm2go/p2.Fn1999
func Fn1999(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2000 github.com/goccy/llamawasm2go/p2.Fn2000
func Fn2000(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn2003 github.com/goccy/llamawasm2go/p2.Fn2003
func Fn2003(m *base.Module, l0 int64)

//go:linkname Fn2013 github.com/goccy/llamawasm2go/p2.Fn2013
func Fn2013(m *base.Module, l0 int64)

//go:linkname Fn2015 github.com/goccy/llamawasm2go/p2.Fn2015
func Fn2015(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2016 github.com/goccy/llamawasm2go/p2.Fn2016
func Fn2016(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2029 github.com/goccy/llamawasm2go/p2.Fn2029
func Fn2029(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2030 github.com/goccy/llamawasm2go/p2.Fn2030
func Fn2030(m *base.Module, l0 int64) int64

//go:linkname Fn2031 github.com/goccy/llamawasm2go/p2.Fn2031
func Fn2031(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2033 github.com/goccy/llamawasm2go/p2.Fn2033
func Fn2033(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2047 github.com/goccy/llamawasm2go/p2.Fn2047
func Fn2047(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2048 github.com/goccy/llamawasm2go/p2.Fn2048
func Fn2048(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2050 github.com/goccy/llamawasm2go/p2.Fn2050
func Fn2050(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2051 github.com/goccy/llamawasm2go/p2.Fn2051
func Fn2051(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2052 github.com/goccy/llamawasm2go/p2.Fn2052
func Fn2052(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2053 github.com/goccy/llamawasm2go/p2.Fn2053
func Fn2053(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2054 github.com/goccy/llamawasm2go/p2.Fn2054
func Fn2054(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2055 github.com/goccy/llamawasm2go/p2.Fn2055
func Fn2055(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2074 github.com/goccy/llamawasm2go/p2.Fn2074
func Fn2074(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2075 github.com/goccy/llamawasm2go/p2.Fn2075
func Fn2075(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2076 github.com/goccy/llamawasm2go/p2.Fn2076
func Fn2076(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2078 github.com/goccy/llamawasm2go/p2.Fn2078
func Fn2078(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2079 github.com/goccy/llamawasm2go/p2.Fn2079
func Fn2079(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2080 github.com/goccy/llamawasm2go/p2.Fn2080
func Fn2080(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2081 github.com/goccy/llamawasm2go/p2.Fn2081
func Fn2081(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2082 github.com/goccy/llamawasm2go/p2.Fn2082
func Fn2082(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2084 github.com/goccy/llamawasm2go/p2.Fn2084
func Fn2084(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2086 github.com/goccy/llamawasm2go/p2.Fn2086
func Fn2086(m *base.Module, l0 int64)

//go:linkname Fn2103 github.com/goccy/llamawasm2go/p2.Fn2103
func Fn2103(m *base.Module, l0 int64)

//go:linkname Fn2104 github.com/goccy/llamawasm2go/p2.Fn2104
func Fn2104(m *base.Module, l0 int64)

//go:linkname Fn2105 github.com/goccy/llamawasm2go/p2.Fn2105
func Fn2105(m *base.Module, l0 int64)

//go:linkname Fn2107 github.com/goccy/llamawasm2go/p2.Fn2107
func Fn2107(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2108 github.com/goccy/llamawasm2go/p2.Fn2108
func Fn2108(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2147 github.com/goccy/llamawasm2go/p2.Fn2147
func Fn2147(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2156 github.com/goccy/llamawasm2go/p2.Fn2156
func Fn2156(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2158 github.com/goccy/llamawasm2go/p2.Fn2158
func Fn2158(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2160 github.com/goccy/llamawasm2go/p2.Fn2160
func Fn2160(m *base.Module, l0 int64) int64

//go:linkname Fn2163 github.com/goccy/llamawasm2go/p2.Fn2163
func Fn2163(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2167 github.com/goccy/llamawasm2go/p2.Fn2167
func Fn2167(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2172 github.com/goccy/llamawasm2go/p2.Fn2172
func Fn2172(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2186 github.com/goccy/llamawasm2go/p2.Fn2186
func Fn2186(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2188 github.com/goccy/llamawasm2go/p2.Fn2188
func Fn2188(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2189 github.com/goccy/llamawasm2go/p2.Fn2189
func Fn2189(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2203 github.com/goccy/llamawasm2go/p2.Fn2203
func Fn2203(m *base.Module, l0 int64) int64

//go:linkname Fn2205 github.com/goccy/llamawasm2go/p2.Fn2205
func Fn2205(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2208 github.com/goccy/llamawasm2go/p2.Fn2208
func Fn2208(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2212 github.com/goccy/llamawasm2go/p2.Fn2212
func Fn2212(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2213 github.com/goccy/llamawasm2go/p2.Fn2213
func Fn2213(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2214 github.com/goccy/llamawasm2go/p0.Fn2214
func Fn2214(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2216 github.com/goccy/llamawasm2go/p2.Fn2216
func Fn2216(m *base.Module, l0 int64) int64

//go:linkname Fn2222 github.com/goccy/llamawasm2go/p2.Fn2222
func Fn2222(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2226 github.com/goccy/llamawasm2go/p2.Fn2226
func Fn2226(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2227 github.com/goccy/llamawasm2go/p2.Fn2227
func Fn2227(m *base.Module, l0 int64)

//go:linkname Fn2228 github.com/goccy/llamawasm2go/p2.Fn2228
func Fn2228(m *base.Module, l0 int64)

//go:linkname Fn2229 github.com/goccy/llamawasm2go/p2.Fn2229
func Fn2229(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2231 github.com/goccy/llamawasm2go/p2.Fn2231
func Fn2231(m *base.Module, l0 int64)

//go:linkname Fn2232 github.com/goccy/llamawasm2go/p2.Fn2232
func Fn2232(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2233 github.com/goccy/llamawasm2go/p2.Fn2233
func Fn2233(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2240 github.com/goccy/llamawasm2go/p2.Fn2240
func Fn2240(m *base.Module, l0 int64) int64

//go:linkname Fn2241 github.com/goccy/llamawasm2go/p2.Fn2241
func Fn2241(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2242 github.com/goccy/llamawasm2go/p2.Fn2242
func Fn2242(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p2.Fn2249
func Fn2249(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2250 github.com/goccy/llamawasm2go/p2.Fn2250
func Fn2250(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2252 github.com/goccy/llamawasm2go/p2.Fn2252
func Fn2252(m *base.Module) int64

//go:linkname Fn2254 github.com/goccy/llamawasm2go/p2.Fn2254
func Fn2254(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2255 github.com/goccy/llamawasm2go/p2.Fn2255
func Fn2255(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2257 github.com/goccy/llamawasm2go/p2.Fn2257
func Fn2257(m *base.Module) int64

//go:linkname Fn2259 github.com/goccy/llamawasm2go/p2.Fn2259
func Fn2259(m *base.Module, l0 int32) int64

//go:linkname Fn2260 github.com/goccy/llamawasm2go/p2.Fn2260
func Fn2260(m *base.Module, l0 int32) int32

//go:linkname Fn2261 github.com/goccy/llamawasm2go/p2.Fn2261
func Fn2261(m *base.Module, l0 int32) int64

//go:linkname Fn2262 github.com/goccy/llamawasm2go/p2.Fn2262
func Fn2262(m *base.Module, l0 float32) int64

//go:linkname Fn2263 github.com/goccy/llamawasm2go/p2.Fn2263
func Fn2263(m *base.Module, l0 float32) int64

//go:linkname Fn2264 github.com/goccy/llamawasm2go/p2.Fn2264
func Fn2264(m *base.Module, l0 float32) int64

//go:linkname Fn2266 github.com/goccy/llamawasm2go/p2.Fn2266
func Fn2266(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2267 github.com/goccy/llamawasm2go/p2.Fn2267
func Fn2267(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2308 github.com/goccy/llamawasm2go/p2.Fn2308
func Fn2308(m *base.Module, l0 int64)

//go:linkname Fn2310 github.com/goccy/llamawasm2go/p2.Fn2310
func Fn2310(m *base.Module, l0 int64)

//go:linkname Fn2352 github.com/goccy/llamawasm2go/p2.Fn2352
func Fn2352(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2363 github.com/goccy/llamawasm2go/p2.Fn2363
func Fn2363(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2366 github.com/goccy/llamawasm2go/p2.Fn2366
func Fn2366(m *base.Module, l0 int64)

//go:linkname Fn2367 github.com/goccy/llamawasm2go/p2.Fn2367
func Fn2367(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2368 github.com/goccy/llamawasm2go/p2.Fn2368
func Fn2368(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2370 github.com/goccy/llamawasm2go/p2.Fn2370
func Fn2370(m *base.Module, l0 int64)

//go:linkname Fn2374 github.com/goccy/llamawasm2go/p2.Fn2374
func Fn2374(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2380 github.com/goccy/llamawasm2go/p2.Fn2380
func Fn2380(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2395 github.com/goccy/llamawasm2go/p2.Fn2395
func Fn2395(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2402 github.com/goccy/llamawasm2go/p0.Fn2402
func Fn2402(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2405 github.com/goccy/llamawasm2go/p2.Fn2405
func Fn2405(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2408 github.com/goccy/llamawasm2go/p2.Fn2408
func Fn2408(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2411 github.com/goccy/llamawasm2go/p2.Fn2411
func Fn2411(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2412 github.com/goccy/llamawasm2go/p2.Fn2412
func Fn2412(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2413 github.com/goccy/llamawasm2go/p2.Fn2413
func Fn2413(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2418 github.com/goccy/llamawasm2go/p2.Fn2418
func Fn2418(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2420 github.com/goccy/llamawasm2go/p2.Fn2420
func Fn2420(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2437 github.com/goccy/llamawasm2go/p2.Fn2437
func Fn2437(m *base.Module, l0 int64)

//go:linkname Fn2438 github.com/goccy/llamawasm2go/p2.Fn2438
func Fn2438(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2443 github.com/goccy/llamawasm2go/p2.Fn2443
func Fn2443(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2448 github.com/goccy/llamawasm2go/p2.Fn2448
func Fn2448(m *base.Module, l0 int64) int64

//go:linkname Fn2449 github.com/goccy/llamawasm2go/p2.Fn2449
func Fn2449(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2450 github.com/goccy/llamawasm2go/p0.Fn2450
func Fn2450(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2461 github.com/goccy/llamawasm2go/p2.Fn2461
func Fn2461(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2462 github.com/goccy/llamawasm2go/p2.Fn2462
func Fn2462(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2463 github.com/goccy/llamawasm2go/p2.Fn2463
func Fn2463(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2464 github.com/goccy/llamawasm2go/p2.Fn2464
func Fn2464(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2484 github.com/goccy/llamawasm2go/p2.Fn2484
func Fn2484(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2485 github.com/goccy/llamawasm2go/p2.Fn2485
func Fn2485(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2488 github.com/goccy/llamawasm2go/p2.Fn2488
func Fn2488(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2574 github.com/goccy/llamawasm2go/p2.Fn2574
func Fn2574(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2693 github.com/goccy/llamawasm2go/p2.Fn2693
func Fn2693(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2695 github.com/goccy/llamawasm2go/p2.Fn2695
func Fn2695(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2696 github.com/goccy/llamawasm2go/p2.Fn2696
func Fn2696(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2700 github.com/goccy/llamawasm2go/p2.Fn2700
func Fn2700(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2725 github.com/goccy/llamawasm2go/p2.Fn2725
func Fn2725(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2857 github.com/goccy/llamawasm2go/p2.Fn2857
func Fn2857(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2864 github.com/goccy/llamawasm2go/p2.Fn2864
func Fn2864(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2892 github.com/goccy/llamawasm2go/p2.Fn2892
func Fn2892(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2906 github.com/goccy/llamawasm2go/p2.Fn2906
func Fn2906(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2918 github.com/goccy/llamawasm2go/p2.Fn2918
func Fn2918(m *base.Module, l0 int32)

//go:linkname Fn2920 github.com/goccy/llamawasm2go/p2.Fn2920
func Fn2920(m *base.Module, l0 int64) int64

//go:linkname Fn2921 github.com/goccy/llamawasm2go/p2.Fn2921
func Fn2921(m *base.Module, l0 int64)

//go:linkname Fn2924 github.com/goccy/llamawasm2go/p2.Fn2924
func Fn2924(m *base.Module, l0 int64)

//go:linkname Fn2925 github.com/goccy/llamawasm2go/p2.Fn2925
func Fn2925(m *base.Module, l0 int64)

//go:linkname Fn2927 github.com/goccy/llamawasm2go/p2.Fn2927
func Fn2927(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2928 github.com/goccy/llamawasm2go/p2.Fn2928
func Fn2928(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2934 github.com/goccy/llamawasm2go/p2.Fn2934
func Fn2934(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2936 github.com/goccy/llamawasm2go/p2.Fn2936
func Fn2936(m *base.Module, l0 int64) int32

//go:linkname Fn2940 github.com/goccy/llamawasm2go/p2.Fn2940
func Fn2940(m *base.Module) int32

//go:linkname Fn2951 github.com/goccy/llamawasm2go/p2.Fn2951
func Fn2951(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2955 github.com/goccy/llamawasm2go/p2.Fn2955
func Fn2955(m *base.Module) int64

//go:linkname Fn2957 github.com/goccy/llamawasm2go/p2.Fn2957
func Fn2957(m *base.Module, l0 float64) float32

//go:linkname Fn2958 github.com/goccy/llamawasm2go/p2.Fn2958
func Fn2958(m *base.Module, l0 float64) float32

//go:linkname Fn2962 github.com/goccy/llamawasm2go/p2.Fn2962
func Fn2962(m *base.Module, l0 float64) float64

//go:linkname Fn2965 github.com/goccy/llamawasm2go/p2.Fn2965
func Fn2965(m *base.Module, l0 int32) float32

//go:linkname Fn2966 github.com/goccy/llamawasm2go/p2.Fn2966
func Fn2966(m *base.Module, l0 int32) float32

//go:linkname Fn2969 github.com/goccy/llamawasm2go/p2.Fn2969
func Fn2969(m *base.Module, l0 float32) float32

//go:linkname Fn2972 github.com/goccy/llamawasm2go/p2.Fn2972
func Fn2972(m *base.Module, l0 float64) float64

//go:linkname Fn2973 github.com/goccy/llamawasm2go/p2.Fn2973
func Fn2973(m *base.Module, l0 float64) float64

//go:linkname Fn2974 github.com/goccy/llamawasm2go/p2.Fn2974
func Fn2974(m *base.Module, l0 float32) float32

//go:linkname Fn2976 github.com/goccy/llamawasm2go/p2.Fn2976
func Fn2976(m *base.Module, l0 float32) float32

//go:linkname Fn2978 github.com/goccy/llamawasm2go/p2.Fn2978
func Fn2978(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn2979 github.com/goccy/llamawasm2go/p2.Fn2979
func Fn2979(m *base.Module, l0 float32) float32

//go:linkname Fn2996 github.com/goccy/llamawasm2go/p2.Fn2996
func Fn2996(m *base.Module, l0 int64) int32

//go:linkname Fn2997 github.com/goccy/llamawasm2go/p2.Fn2997
func Fn2997(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2999 github.com/goccy/llamawasm2go/p2.Fn2999
func Fn2999(m *base.Module, l0 int64)

//go:linkname Fn3000 github.com/goccy/llamawasm2go/p2.Fn3000
func Fn3000(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3001 github.com/goccy/llamawasm2go/p2.Fn3001
func Fn3001(m *base.Module, l0 int64) int32

//go:linkname Fn3008 github.com/goccy/llamawasm2go/p2.Fn3008
func Fn3008(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3010 github.com/goccy/llamawasm2go/p2.Fn3010
func Fn3010(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3016 github.com/goccy/llamawasm2go/p2.Fn3016
func Fn3016(m *base.Module, l0 int64) int32

//go:linkname Fn3019 github.com/goccy/llamawasm2go/p2.Fn3019
func Fn3019(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3022 github.com/goccy/llamawasm2go/p2.Fn3022
func Fn3022(m *base.Module, l0 int64) int32

//go:linkname Fn3024 github.com/goccy/llamawasm2go/p2.Fn3024
func Fn3024(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3026 github.com/goccy/llamawasm2go/p2.Fn3026
func Fn3026(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3027 github.com/goccy/llamawasm2go/p2.Fn3027
func Fn3027(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3030 github.com/goccy/llamawasm2go/p2.Fn3030
func Fn3030(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3033 github.com/goccy/llamawasm2go/p2.Fn3033
func Fn3033(m *base.Module, l0 int64) int64

//go:linkname Fn3037 github.com/goccy/llamawasm2go/p2.Fn3037
func Fn3037(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3047 github.com/goccy/llamawasm2go/p2.Fn3047
func Fn3047(m *base.Module)

//go:linkname Fn3048 github.com/goccy/llamawasm2go/p0.Fn3048
func Fn3048(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3051 github.com/goccy/llamawasm2go/p2.Fn3051
func Fn3051(m *base.Module)

//go:linkname Fn3053 github.com/goccy/llamawasm2go/p0.Fn3053
func Fn3053(m *base.Module, l0 int64) int64

//go:linkname Fn3055 github.com/goccy/llamawasm2go/p2.Fn3055
func Fn3055(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3059 github.com/goccy/llamawasm2go/p2.Fn3059
func Fn3059(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3128 github.com/goccy/llamawasm2go/p2.Fn3128
func Fn3128(m *base.Module, l0 int32)
