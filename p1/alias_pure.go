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

//go:linkname Fn334 github.com/goccy/llamawasm2go/p2.Fn334
func Fn334(m *base.Module, l0 int64) int32

//go:linkname Fn335 github.com/goccy/llamawasm2go/p2.Fn335
func Fn335(m *base.Module, l0 int64)

//go:linkname Fn336 github.com/goccy/llamawasm2go/p2.Fn336
func Fn336(m *base.Module, l0 int64)

//go:linkname Fn355 github.com/goccy/llamawasm2go/p2.Fn355
func Fn355(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn356 github.com/goccy/llamawasm2go/p2.Fn356
func Fn356(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn360 github.com/goccy/llamawasm2go/p2.Fn360
func Fn360(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn361 github.com/goccy/llamawasm2go/p2.Fn361
func Fn361(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn364 github.com/goccy/llamawasm2go/p2.Fn364
func Fn364(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn365 github.com/goccy/llamawasm2go/p2.Fn365
func Fn365(m *base.Module, l0 int64)

//go:linkname Fn366 github.com/goccy/llamawasm2go/p2.Fn366
func Fn366(m *base.Module, l0 int64)

//go:linkname Fn368 github.com/goccy/llamawasm2go/p2.Fn368
func Fn368(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn370 github.com/goccy/llamawasm2go/p2.Fn370
func Fn370(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn371 github.com/goccy/llamawasm2go/p2.Fn371
func Fn371(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn373 github.com/goccy/llamawasm2go/p2.Fn373
func Fn373(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn376 github.com/goccy/llamawasm2go/p0.Fn376
func Fn376(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int64)

//go:linkname Fn379 github.com/goccy/llamawasm2go/p2.Fn379
func Fn379(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn380 github.com/goccy/llamawasm2go/p2.Fn380
func Fn380(m *base.Module, l0 int64)

//go:linkname Fn383 github.com/goccy/llamawasm2go/p2.Fn383
func Fn383(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn384 github.com/goccy/llamawasm2go/p2.Fn384
func Fn384(m *base.Module, l0 int64)

//go:linkname Fn385 github.com/goccy/llamawasm2go/p2.Fn385
func Fn385(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn386 github.com/goccy/llamawasm2go/p2.Fn386
func Fn386(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn387 github.com/goccy/llamawasm2go/p2.Fn387
func Fn387(m *base.Module)

//go:linkname Fn388 github.com/goccy/llamawasm2go/p0.Fn388
func Fn388(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int32)

//go:linkname Fn392 github.com/goccy/llamawasm2go/p0.Fn392
func Fn392(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 int64)

//go:linkname Fn400 github.com/goccy/llamawasm2go/p2.Fn400
func Fn400(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn406 github.com/goccy/llamawasm2go/p2.Fn406
func Fn406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn408 github.com/goccy/llamawasm2go/p2.Fn408
func Fn408(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn409 github.com/goccy/llamawasm2go/p2.Fn409
func Fn409(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn430 github.com/goccy/llamawasm2go/p2.Fn430
func Fn430(m *base.Module, l0 int64) int32

//go:linkname Fn431 github.com/goccy/llamawasm2go/p2.Fn431
func Fn431(m *base.Module, l0 int64) int32

//go:linkname Fn442 github.com/goccy/llamawasm2go/p2.Fn442
func Fn442(m *base.Module, l0 int64) int64

//go:linkname Fn447 github.com/goccy/llamawasm2go/p2.Fn447
func Fn447(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn449 github.com/goccy/llamawasm2go/p2.Fn449
func Fn449(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn450 github.com/goccy/llamawasm2go/p2.Fn450
func Fn450(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn451 github.com/goccy/llamawasm2go/p2.Fn451
func Fn451(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn452 github.com/goccy/llamawasm2go/p2.Fn452
func Fn452(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn453 github.com/goccy/llamawasm2go/p2.Fn453
func Fn453(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn454 github.com/goccy/llamawasm2go/p2.Fn454
func Fn454(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn457 github.com/goccy/llamawasm2go/p2.Fn457
func Fn457(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn458 github.com/goccy/llamawasm2go/p2.Fn458
func Fn458(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn461 github.com/goccy/llamawasm2go/p2.Fn461
func Fn461(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn462 github.com/goccy/llamawasm2go/p2.Fn462
func Fn462(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn463 github.com/goccy/llamawasm2go/p2.Fn463
func Fn463(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn464 github.com/goccy/llamawasm2go/p2.Fn464
func Fn464(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn465 github.com/goccy/llamawasm2go/p2.Fn465
func Fn465(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn466 github.com/goccy/llamawasm2go/p2.Fn466
func Fn466(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn468 github.com/goccy/llamawasm2go/p2.Fn468
func Fn468(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn470 github.com/goccy/llamawasm2go/p2.Fn470
func Fn470(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn472 github.com/goccy/llamawasm2go/p2.Fn472
func Fn472(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn473 github.com/goccy/llamawasm2go/p2.Fn473
func Fn473(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn474 github.com/goccy/llamawasm2go/p2.Fn474
func Fn474(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn475 github.com/goccy/llamawasm2go/p2.Fn475
func Fn475(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn476 github.com/goccy/llamawasm2go/p2.Fn476
func Fn476(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn478 github.com/goccy/llamawasm2go/p2.Fn478
func Fn478(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn479 github.com/goccy/llamawasm2go/p2.Fn479
func Fn479(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn480 github.com/goccy/llamawasm2go/p2.Fn480
func Fn480(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn481 github.com/goccy/llamawasm2go/p2.Fn481
func Fn481(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn482 github.com/goccy/llamawasm2go/p2.Fn482
func Fn482(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn483 github.com/goccy/llamawasm2go/p2.Fn483
func Fn483(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn485 github.com/goccy/llamawasm2go/p2.Fn485
func Fn485(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn486 github.com/goccy/llamawasm2go/p2.Fn486
func Fn486(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn487 github.com/goccy/llamawasm2go/p2.Fn487
func Fn487(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn488 github.com/goccy/llamawasm2go/p2.Fn488
func Fn488(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn489 github.com/goccy/llamawasm2go/p2.Fn489
func Fn489(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn490 github.com/goccy/llamawasm2go/p2.Fn490
func Fn490(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn491 github.com/goccy/llamawasm2go/p2.Fn491
func Fn491(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn492 github.com/goccy/llamawasm2go/p2.Fn492
func Fn492(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn493 github.com/goccy/llamawasm2go/p2.Fn493
func Fn493(m *base.Module, l0 int64)

//go:linkname Fn494 github.com/goccy/llamawasm2go/p2.Fn494
func Fn494(m *base.Module, l0 int64)

//go:linkname Fn495 github.com/goccy/llamawasm2go/p2.Fn495
func Fn495(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn496 github.com/goccy/llamawasm2go/p2.Fn496
func Fn496(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn497 github.com/goccy/llamawasm2go/p2.Fn497
func Fn497(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn498 github.com/goccy/llamawasm2go/p2.Fn498
func Fn498(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn499 github.com/goccy/llamawasm2go/p2.Fn499
func Fn499(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn500 github.com/goccy/llamawasm2go/p2.Fn500
func Fn500(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn501 github.com/goccy/llamawasm2go/p2.Fn501
func Fn501(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn502 github.com/goccy/llamawasm2go/p2.Fn502
func Fn502(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn503 github.com/goccy/llamawasm2go/p2.Fn503
func Fn503(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn504 github.com/goccy/llamawasm2go/p2.Fn504
func Fn504(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn506 github.com/goccy/llamawasm2go/p2.Fn506
func Fn506(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn507 github.com/goccy/llamawasm2go/p2.Fn507
func Fn507(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn508 github.com/goccy/llamawasm2go/p2.Fn508
func Fn508(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn509 github.com/goccy/llamawasm2go/p2.Fn509
func Fn509(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn510 github.com/goccy/llamawasm2go/p2.Fn510
func Fn510(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn511 github.com/goccy/llamawasm2go/p2.Fn511
func Fn511(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn512 github.com/goccy/llamawasm2go/p2.Fn512
func Fn512(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64) int64

//go:linkname Fn513 github.com/goccy/llamawasm2go/p2.Fn513
func Fn513(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn514 github.com/goccy/llamawasm2go/p2.Fn514
func Fn514(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn515 github.com/goccy/llamawasm2go/p2.Fn515
func Fn515(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn516 github.com/goccy/llamawasm2go/p2.Fn516
func Fn516(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn517 github.com/goccy/llamawasm2go/p2.Fn517
func Fn517(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn519 github.com/goccy/llamawasm2go/p2.Fn519
func Fn519(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn522 github.com/goccy/llamawasm2go/p2.Fn522
func Fn522(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32) int64

//go:linkname Fn523 github.com/goccy/llamawasm2go/p2.Fn523
func Fn523(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32, l7 float32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32) int64

//go:linkname Fn525 github.com/goccy/llamawasm2go/p2.Fn525
func Fn525(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn526 github.com/goccy/llamawasm2go/p2.Fn526
func Fn526(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn527 github.com/goccy/llamawasm2go/p2.Fn527
func Fn527(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn529 github.com/goccy/llamawasm2go/p2.Fn529
func Fn529(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn531 github.com/goccy/llamawasm2go/p2.Fn531
func Fn531(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn532 github.com/goccy/llamawasm2go/p2.Fn532
func Fn532(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn533 github.com/goccy/llamawasm2go/p2.Fn533
func Fn533(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn535 github.com/goccy/llamawasm2go/p2.Fn535
func Fn535(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 float32, l6 float32, l7 float32) int64

//go:linkname Fn536 github.com/goccy/llamawasm2go/p2.Fn536
func Fn536(m *base.Module, l0 int64)

//go:linkname Fn537 github.com/goccy/llamawasm2go/p2.Fn537
func Fn537(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn541 github.com/goccy/llamawasm2go/p2.Fn541
func Fn541(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn542 github.com/goccy/llamawasm2go/p2.Fn542
func Fn542(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn545 github.com/goccy/llamawasm2go/p2.Fn545
func Fn545(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn547 github.com/goccy/llamawasm2go/p2.Fn547
func Fn547(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn549 github.com/goccy/llamawasm2go/p2.Fn549
func Fn549(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn555 github.com/goccy/llamawasm2go/p2.Fn555
func Fn555(m *base.Module, l0 int64)

//go:linkname Fn558 github.com/goccy/llamawasm2go/p2.Fn558
func Fn558(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn559 github.com/goccy/llamawasm2go/p2.Fn559
func Fn559(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn561 github.com/goccy/llamawasm2go/p2.Fn561
func Fn561(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn564 github.com/goccy/llamawasm2go/p2.Fn564
func Fn564(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn565 github.com/goccy/llamawasm2go/p2.Fn565
func Fn565(m *base.Module, l0 int64) int64

//go:linkname Fn566 github.com/goccy/llamawasm2go/p2.Fn566
func Fn566(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn567 github.com/goccy/llamawasm2go/p2.Fn567
func Fn567(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn568 github.com/goccy/llamawasm2go/p2.Fn568
func Fn568(m *base.Module, l0 int64) int64

//go:linkname Fn569 github.com/goccy/llamawasm2go/p2.Fn569
func Fn569(m *base.Module, l0 int64) int64

//go:linkname Fn570 github.com/goccy/llamawasm2go/p2.Fn570
func Fn570(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn572 github.com/goccy/llamawasm2go/p2.Fn572
func Fn572(m *base.Module, l0 int64) int64

//go:linkname Fn573 github.com/goccy/llamawasm2go/p2.Fn573
func Fn573(m *base.Module, l0 int64) int64

//go:linkname Fn575 github.com/goccy/llamawasm2go/p2.Fn575
func Fn575(m *base.Module, l0 int64)

//go:linkname Fn576 github.com/goccy/llamawasm2go/p2.Fn576
func Fn576(m *base.Module, l0 int64) int64

//go:linkname Fn577 github.com/goccy/llamawasm2go/p2.Fn577
func Fn577(m *base.Module, l0 int64) int64

//go:linkname Fn578 github.com/goccy/llamawasm2go/p2.Fn578
func Fn578(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn579 github.com/goccy/llamawasm2go/p2.Fn579
func Fn579(m *base.Module, l0 int64) int32

//go:linkname Fn580 github.com/goccy/llamawasm2go/p2.Fn580
func Fn580(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn582 github.com/goccy/llamawasm2go/p2.Fn582
func Fn582(m *base.Module, l0 int64)

//go:linkname Fn583 github.com/goccy/llamawasm2go/p2.Fn583
func Fn583(m *base.Module, l0 int64)

//go:linkname Fn584 github.com/goccy/llamawasm2go/p2.Fn584
func Fn584(m *base.Module, l0 int64) int64

//go:linkname Fn586 github.com/goccy/llamawasm2go/p2.Fn586
func Fn586(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn587 github.com/goccy/llamawasm2go/p2.Fn587
func Fn587(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn588 github.com/goccy/llamawasm2go/p2.Fn588
func Fn588(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn589 github.com/goccy/llamawasm2go/p2.Fn589
func Fn589(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn590 github.com/goccy/llamawasm2go/p2.Fn590
func Fn590(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn591 github.com/goccy/llamawasm2go/p2.Fn591
func Fn591(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn594 github.com/goccy/llamawasm2go/p2.Fn594
func Fn594(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn595 github.com/goccy/llamawasm2go/p2.Fn595
func Fn595(m *base.Module, l0 int64) int64

//go:linkname Fn596 github.com/goccy/llamawasm2go/p2.Fn596
func Fn596(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn599 github.com/goccy/llamawasm2go/p2.Fn599
func Fn599(m *base.Module, l0 int64) int64

//go:linkname Fn600 github.com/goccy/llamawasm2go/p2.Fn600
func Fn600(m *base.Module, l0 int64) int64

//go:linkname Fn604 github.com/goccy/llamawasm2go/p2.Fn604
func Fn604(m *base.Module, l0 int64) int64

//go:linkname Fn605 github.com/goccy/llamawasm2go/p2.Fn605
func Fn605(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn606 github.com/goccy/llamawasm2go/p2.Fn606
func Fn606(m *base.Module, l0 int64) int64

//go:linkname Fn610 github.com/goccy/llamawasm2go/p2.Fn610
func Fn610(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn611 github.com/goccy/llamawasm2go/p0.Fn611
func Fn611(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn616 github.com/goccy/llamawasm2go/p2.Fn616
func Fn616(m *base.Module, l0 int64)

//go:linkname Fn617 github.com/goccy/llamawasm2go/p2.Fn617
func Fn617(m *base.Module, l0 int64)

//go:linkname Fn618 github.com/goccy/llamawasm2go/p2.Fn618
func Fn618(m *base.Module, l0 int64)

//go:linkname Fn619 github.com/goccy/llamawasm2go/p2.Fn619
func Fn619(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn620 github.com/goccy/llamawasm2go/p2.Fn620
func Fn620(m *base.Module, l0 int64) int32

//go:linkname Fn621 github.com/goccy/llamawasm2go/p2.Fn621
func Fn621(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn622 github.com/goccy/llamawasm2go/p2.Fn622
func Fn622(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn623 github.com/goccy/llamawasm2go/p2.Fn623
func Fn623(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn624 github.com/goccy/llamawasm2go/p2.Fn624
func Fn624(m *base.Module, l0 int64) int32

//go:linkname Fn625 github.com/goccy/llamawasm2go/p2.Fn625
func Fn625(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn643 github.com/goccy/llamawasm2go/p2.Fn643
func Fn643(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn649 github.com/goccy/llamawasm2go/p2.Fn649
func Fn649(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn650 github.com/goccy/llamawasm2go/p2.Fn650
func Fn650(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn651 github.com/goccy/llamawasm2go/p2.Fn651
func Fn651(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn653 github.com/goccy/llamawasm2go/p2.Fn653
func Fn653(m *base.Module, l0 int64)

//go:linkname Fn655 github.com/goccy/llamawasm2go/p2.Fn655
func Fn655(m *base.Module, l0 int64)

//go:linkname Fn656 github.com/goccy/llamawasm2go/p0.Fn656
func Fn656(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn658 github.com/goccy/llamawasm2go/p2.Fn658
func Fn658(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn672 github.com/goccy/llamawasm2go/p2.Fn672
func Fn672(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn674 github.com/goccy/llamawasm2go/p2.Fn674
func Fn674(m *base.Module, l0 int64)

//go:linkname Fn675 github.com/goccy/llamawasm2go/p2.Fn675
func Fn675(m *base.Module, l0 int64) int64

//go:linkname Fn698 github.com/goccy/llamawasm2go/p2.Fn698
func Fn698(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn709 github.com/goccy/llamawasm2go/p2.Fn709
func Fn709(m *base.Module, l0 int64)

//go:linkname Fn728 github.com/goccy/llamawasm2go/p2.Fn728
func Fn728(m *base.Module, l0 int64)

//go:linkname Fn729 github.com/goccy/llamawasm2go/p2.Fn729
func Fn729(m *base.Module, l0 int64)

//go:linkname Fn797 github.com/goccy/llamawasm2go/p2.Fn797
func Fn797(m *base.Module, l0 int64)

//go:linkname Fn799 github.com/goccy/llamawasm2go/p2.Fn799
func Fn799(m *base.Module, l0 int64) int64

//go:linkname Fn901 github.com/goccy/llamawasm2go/p2.Fn901
func Fn901(m *base.Module) int64

//go:linkname Fn938 github.com/goccy/llamawasm2go/p2.Fn938
func Fn938(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32)

//go:linkname Fn971 github.com/goccy/llamawasm2go/p2.Fn971
func Fn971(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn974 github.com/goccy/llamawasm2go/p2.Fn974
func Fn974(m *base.Module, l0 int64) int64

//go:linkname Fn984 github.com/goccy/llamawasm2go/p2.Fn984
func Fn984(m *base.Module, l0 int64)

//go:linkname Fn989 github.com/goccy/llamawasm2go/p2.Fn989
func Fn989(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn990 github.com/goccy/llamawasm2go/p2.Fn990
func Fn990(m *base.Module)

//go:linkname Fn999 github.com/goccy/llamawasm2go/p2.Fn999
func Fn999(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1007 github.com/goccy/llamawasm2go/p2.Fn1007
func Fn1007(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1008 github.com/goccy/llamawasm2go/p2.Fn1008
func Fn1008(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1009 github.com/goccy/llamawasm2go/p2.Fn1009
func Fn1009(m *base.Module, l0 int64, l1 int64, l2 int64) float32

//go:linkname Fn1011 github.com/goccy/llamawasm2go/p2.Fn1011
func Fn1011(m *base.Module, l0 int64, l1 int64, l2 int64) float64

//go:linkname Fn1013 github.com/goccy/llamawasm2go/p2.Fn1013
func Fn1013(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1016 github.com/goccy/llamawasm2go/p2.Fn1016
func Fn1016(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1019 github.com/goccy/llamawasm2go/p2.Fn1019
func Fn1019(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1025 github.com/goccy/llamawasm2go/p2.Fn1025
func Fn1025(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1026 github.com/goccy/llamawasm2go/p2.Fn1026
func Fn1026(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1034 github.com/goccy/llamawasm2go/p2.Fn1034
func Fn1034(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1039 github.com/goccy/llamawasm2go/p2.Fn1039
func Fn1039(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1040 github.com/goccy/llamawasm2go/p2.Fn1040
func Fn1040(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1048 github.com/goccy/llamawasm2go/p2.Fn1048
func Fn1048(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1064 github.com/goccy/llamawasm2go/p2.Fn1064
func Fn1064(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1072 github.com/goccy/llamawasm2go/p2.Fn1072
func Fn1072(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1088 github.com/goccy/llamawasm2go/p0.Fn1088
func Fn1088(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1095 github.com/goccy/llamawasm2go/p2.Fn1095
func Fn1095(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1105 github.com/goccy/llamawasm2go/p0.Fn1105
func Fn1105(m *base.Module, l0 int64) int64

//go:linkname Fn1106 github.com/goccy/llamawasm2go/p2.Fn1106
func Fn1106(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1107 github.com/goccy/llamawasm2go/p2.Fn1107
func Fn1107(m *base.Module, l0 int64)

//go:linkname Fn1109 github.com/goccy/llamawasm2go/p2.Fn1109
func Fn1109(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1199 github.com/goccy/llamawasm2go/p2.Fn1199
func Fn1199(m *base.Module, l0 int64)

//go:linkname Fn1221 github.com/goccy/llamawasm2go/p2.Fn1221
func Fn1221(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1231 github.com/goccy/llamawasm2go/p2.Fn1231
func Fn1231(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1237 github.com/goccy/llamawasm2go/p2.Fn1237
func Fn1237(m *base.Module)

//go:linkname Fn1241 github.com/goccy/llamawasm2go/p2.Fn1241
func Fn1241(m *base.Module, l0 int64) int64

//go:linkname Fn1269 github.com/goccy/llamawasm2go/p2.Fn1269
func Fn1269(m *base.Module, l0 int64)

//go:linkname Fn1273 github.com/goccy/llamawasm2go/p2.Fn1273
func Fn1273(m *base.Module, l0 int32) int64

//go:linkname Fn1284 github.com/goccy/llamawasm2go/p2.Fn1284
func Fn1284(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1285 github.com/goccy/llamawasm2go/p2.Fn1285
func Fn1285(m *base.Module, l0 int64)

//go:linkname Fn1287 github.com/goccy/llamawasm2go/p2.Fn1287
func Fn1287(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1308 github.com/goccy/llamawasm2go/p2.Fn1308
func Fn1308(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1309 github.com/goccy/llamawasm2go/p2.Fn1309
func Fn1309(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1310 github.com/goccy/llamawasm2go/p2.Fn1310
func Fn1310(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1311 github.com/goccy/llamawasm2go/p2.Fn1311
func Fn1311(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1312 github.com/goccy/llamawasm2go/p2.Fn1312
func Fn1312(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1313 github.com/goccy/llamawasm2go/p2.Fn1313
func Fn1313(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1314 github.com/goccy/llamawasm2go/p2.Fn1314
func Fn1314(m *base.Module, l0 int64) int64

//go:linkname Fn1317 github.com/goccy/llamawasm2go/p2.Fn1317
func Fn1317(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1321 github.com/goccy/llamawasm2go/p2.Fn1321
func Fn1321(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1322 github.com/goccy/llamawasm2go/p2.Fn1322
func Fn1322(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p2.Fn1324
func Fn1324(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1325 github.com/goccy/llamawasm2go/p2.Fn1325
func Fn1325(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1326 github.com/goccy/llamawasm2go/p2.Fn1326
func Fn1326(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1327 github.com/goccy/llamawasm2go/p2.Fn1327
func Fn1327(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1328 github.com/goccy/llamawasm2go/p2.Fn1328
func Fn1328(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1329 github.com/goccy/llamawasm2go/p2.Fn1329
func Fn1329(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1330 github.com/goccy/llamawasm2go/p2.Fn1330
func Fn1330(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1331 github.com/goccy/llamawasm2go/p2.Fn1331
func Fn1331(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1332 github.com/goccy/llamawasm2go/p2.Fn1332
func Fn1332(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1333 github.com/goccy/llamawasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1335 github.com/goccy/llamawasm2go/p2.Fn1335
func Fn1335(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1336 github.com/goccy/llamawasm2go/p2.Fn1336
func Fn1336(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1337 github.com/goccy/llamawasm2go/p2.Fn1337
func Fn1337(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1338 github.com/goccy/llamawasm2go/p2.Fn1338
func Fn1338(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1339 github.com/goccy/llamawasm2go/p2.Fn1339
func Fn1339(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1340 github.com/goccy/llamawasm2go/p2.Fn1340
func Fn1340(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1341 github.com/goccy/llamawasm2go/p2.Fn1341
func Fn1341(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1342 github.com/goccy/llamawasm2go/p2.Fn1342
func Fn1342(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1343 github.com/goccy/llamawasm2go/p2.Fn1343
func Fn1343(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1344 github.com/goccy/llamawasm2go/p2.Fn1344
func Fn1344(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1345 github.com/goccy/llamawasm2go/p2.Fn1345
func Fn1345(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1346 github.com/goccy/llamawasm2go/p2.Fn1346
func Fn1346(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1347 github.com/goccy/llamawasm2go/p2.Fn1347
func Fn1347(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1349 github.com/goccy/llamawasm2go/p2.Fn1349
func Fn1349(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1353 github.com/goccy/llamawasm2go/p2.Fn1353
func Fn1353(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1354 github.com/goccy/llamawasm2go/p2.Fn1354
func Fn1354(m *base.Module)

//go:linkname Fn1355 github.com/goccy/llamawasm2go/p2.Fn1355
func Fn1355(m *base.Module)

//go:linkname Fn1356 github.com/goccy/llamawasm2go/p0.Fn1356
func Fn1356(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1364 github.com/goccy/llamawasm2go/p2.Fn1364
func Fn1364(m *base.Module)

//go:linkname Fn1366 github.com/goccy/llamawasm2go/p2.Fn1366
func Fn1366(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1368 github.com/goccy/llamawasm2go/p0.Fn1368
func Fn1368(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1369 github.com/goccy/llamawasm2go/p2.Fn1369
func Fn1369(m *base.Module, l0 int64) int64

//go:linkname Fn1374 github.com/goccy/llamawasm2go/p2.Fn1374
func Fn1374(m *base.Module, l0 int64)

//go:linkname Fn1381 github.com/goccy/llamawasm2go/p2.Fn1381
func Fn1381(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1384 github.com/goccy/llamawasm2go/p2.Fn1384
func Fn1384(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1386 github.com/goccy/llamawasm2go/p2.Fn1386
func Fn1386(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1388 github.com/goccy/llamawasm2go/p2.Fn1388
func Fn1388(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1390 github.com/goccy/llamawasm2go/p2.Fn1390
func Fn1390(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p2.Fn1392
func Fn1392(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p2.Fn1394
func Fn1394(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p2.Fn1401
func Fn1401(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p2.Fn1402
func Fn1402(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p0.Fn1403
func Fn1403(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int64)

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p2.Fn1406
func Fn1406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p2.Fn1408
func Fn1408(m *base.Module, l0 int64) int64

//go:linkname Fn1409 github.com/goccy/llamawasm2go/p2.Fn1409
func Fn1409(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1410 github.com/goccy/llamawasm2go/p2.Fn1410
func Fn1410(m *base.Module, l0 int64)

//go:linkname Fn1411 github.com/goccy/llamawasm2go/p2.Fn1411
func Fn1411(m *base.Module, l0 int64)

//go:linkname Fn1412 github.com/goccy/llamawasm2go/p2.Fn1412
func Fn1412(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p2.Fn1413
func Fn1413(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1416 github.com/goccy/llamawasm2go/p2.Fn1416
func Fn1416(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p0.Fn1426
func Fn1426(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1428 github.com/goccy/llamawasm2go/p2.Fn1428
func Fn1428(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1431 github.com/goccy/llamawasm2go/p0.Fn1431
func Fn1431(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1432 github.com/goccy/llamawasm2go/p2.Fn1432
func Fn1432(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1436 github.com/goccy/llamawasm2go/p2.Fn1436
func Fn1436(m *base.Module, l0 int64)

//go:linkname Fn1439 github.com/goccy/llamawasm2go/p2.Fn1439
func Fn1439(m *base.Module, l0 int64)

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64) int64

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1446 github.com/goccy/llamawasm2go/p2.Fn1446
func Fn1446(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1448 github.com/goccy/llamawasm2go/p2.Fn1448
func Fn1448(m *base.Module, l0 int64)

//go:linkname Fn1451 github.com/goccy/llamawasm2go/p2.Fn1451
func Fn1451(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1453 github.com/goccy/llamawasm2go/p2.Fn1453
func Fn1453(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1456 github.com/goccy/llamawasm2go/p2.Fn1456
func Fn1456(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1463 github.com/goccy/llamawasm2go/p2.Fn1463
func Fn1463(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1464 github.com/goccy/llamawasm2go/p2.Fn1464
func Fn1464(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1466 github.com/goccy/llamawasm2go/p2.Fn1466
func Fn1466(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1470 github.com/goccy/llamawasm2go/p2.Fn1470
func Fn1470(m *base.Module, l0 int64)

//go:linkname Fn1471 github.com/goccy/llamawasm2go/p2.Fn1471
func Fn1471(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1473 github.com/goccy/llamawasm2go/p2.Fn1473
func Fn1473(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1474 github.com/goccy/llamawasm2go/p2.Fn1474
func Fn1474(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1475 github.com/goccy/llamawasm2go/p2.Fn1475
func Fn1475(m *base.Module, l0 int64)

//go:linkname Fn1477 github.com/goccy/llamawasm2go/p2.Fn1477
func Fn1477(m *base.Module, l0 int64) int32

//go:linkname Fn1478 github.com/goccy/llamawasm2go/p2.Fn1478
func Fn1478(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1520 github.com/goccy/llamawasm2go/p2.Fn1520
func Fn1520(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1536 github.com/goccy/llamawasm2go/p2.Fn1536
func Fn1536(m *base.Module, l0 int64)

//go:linkname Fn1537 github.com/goccy/llamawasm2go/p2.Fn1537
func Fn1537(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1538 github.com/goccy/llamawasm2go/p2.Fn1538
func Fn1538(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1539 github.com/goccy/llamawasm2go/p2.Fn1539
func Fn1539(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1540 github.com/goccy/llamawasm2go/p2.Fn1540
func Fn1540(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p2.Fn1541
func Fn1541(m *base.Module, l0 int64)

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p2.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1544 github.com/goccy/llamawasm2go/p2.Fn1544
func Fn1544(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1546 github.com/goccy/llamawasm2go/p2.Fn1546
func Fn1546(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1547 github.com/goccy/llamawasm2go/p2.Fn1547
func Fn1547(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1548 github.com/goccy/llamawasm2go/p2.Fn1548
func Fn1548(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1551 github.com/goccy/llamawasm2go/p2.Fn1551
func Fn1551(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1554 github.com/goccy/llamawasm2go/p2.Fn1554
func Fn1554(m *base.Module, l0 int64) int64

//go:linkname Fn1555 github.com/goccy/llamawasm2go/p2.Fn1555
func Fn1555(m *base.Module, l0 int64) int64

//go:linkname Fn1556 github.com/goccy/llamawasm2go/p2.Fn1556
func Fn1556(m *base.Module, l0 int64) int64

//go:linkname Fn1557 github.com/goccy/llamawasm2go/p2.Fn1557
func Fn1557(m *base.Module, l0 int64) int64

//go:linkname Fn1558 github.com/goccy/llamawasm2go/p2.Fn1558
func Fn1558(m *base.Module, l0 int64) int64

//go:linkname Fn1559 github.com/goccy/llamawasm2go/p2.Fn1559
func Fn1559(m *base.Module, l0 int64) int64

//go:linkname Fn1561 github.com/goccy/llamawasm2go/p2.Fn1561
func Fn1561(m *base.Module, l0 int64) int64

//go:linkname Fn1562 github.com/goccy/llamawasm2go/p2.Fn1562
func Fn1562(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1564 github.com/goccy/llamawasm2go/p2.Fn1564
func Fn1564(m *base.Module, l0 int64) int64

//go:linkname Fn1565 github.com/goccy/llamawasm2go/p2.Fn1565
func Fn1565(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1566 github.com/goccy/llamawasm2go/p2.Fn1566
func Fn1566(m *base.Module, l0 int64) int64

//go:linkname Fn1568 github.com/goccy/llamawasm2go/p2.Fn1568
func Fn1568(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1569 github.com/goccy/llamawasm2go/p2.Fn1569
func Fn1569(m *base.Module, l0 int64) int64

//go:linkname Fn1571 github.com/goccy/llamawasm2go/p2.Fn1571
func Fn1571(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1572 github.com/goccy/llamawasm2go/p2.Fn1572
func Fn1572(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1573 github.com/goccy/llamawasm2go/p2.Fn1573
func Fn1573(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1574 github.com/goccy/llamawasm2go/p2.Fn1574
func Fn1574(m *base.Module, l0 int64) int64

//go:linkname Fn1576 github.com/goccy/llamawasm2go/p2.Fn1576
func Fn1576(m *base.Module, l0 int64) int64

//go:linkname Fn1577 github.com/goccy/llamawasm2go/p2.Fn1577
func Fn1577(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1582 github.com/goccy/llamawasm2go/p2.Fn1582
func Fn1582(m *base.Module, l0 int64) int64

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p2.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1584 github.com/goccy/llamawasm2go/p2.Fn1584
func Fn1584(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1585 github.com/goccy/llamawasm2go/p2.Fn1585
func Fn1585(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1587 github.com/goccy/llamawasm2go/p2.Fn1587
func Fn1587(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1588 github.com/goccy/llamawasm2go/p2.Fn1588
func Fn1588(m *base.Module, l0 int64) int64

//go:linkname Fn1589 github.com/goccy/llamawasm2go/p2.Fn1589
func Fn1589(m *base.Module, l0 int64) int64

//go:linkname Fn1591 github.com/goccy/llamawasm2go/p2.Fn1591
func Fn1591(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1592 github.com/goccy/llamawasm2go/p2.Fn1592
func Fn1592(m *base.Module, l0 int64)

//go:linkname Fn1621 github.com/goccy/llamawasm2go/p2.Fn1621
func Fn1621(m *base.Module, l0 int64) int64

//go:linkname Fn1628 github.com/goccy/llamawasm2go/p2.Fn1628
func Fn1628(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1629 github.com/goccy/llamawasm2go/p2.Fn1629
func Fn1629(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1630 github.com/goccy/llamawasm2go/p2.Fn1630
func Fn1630(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1632 github.com/goccy/llamawasm2go/p2.Fn1632
func Fn1632(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1633 github.com/goccy/llamawasm2go/p2.Fn1633
func Fn1633(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1637 github.com/goccy/llamawasm2go/p2.Fn1637
func Fn1637(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1638 github.com/goccy/llamawasm2go/p2.Fn1638
func Fn1638(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1639 github.com/goccy/llamawasm2go/p2.Fn1639
func Fn1639(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1640 github.com/goccy/llamawasm2go/p2.Fn1640
func Fn1640(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1641 github.com/goccy/llamawasm2go/p2.Fn1641
func Fn1641(m *base.Module, l0 int64) int32

//go:linkname Fn1642 github.com/goccy/llamawasm2go/p2.Fn1642
func Fn1642(m *base.Module, l0 int64) int32

//go:linkname Fn1643 github.com/goccy/llamawasm2go/p2.Fn1643
func Fn1643(m *base.Module, l0 int64) int32

//go:linkname Fn1644 github.com/goccy/llamawasm2go/p2.Fn1644
func Fn1644(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1647 github.com/goccy/llamawasm2go/p2.Fn1647
func Fn1647(m *base.Module, l0 int64) int32

//go:linkname Fn1648 github.com/goccy/llamawasm2go/p2.Fn1648
func Fn1648(m *base.Module, l0 int64) int32

//go:linkname Fn1654 github.com/goccy/llamawasm2go/p2.Fn1654
func Fn1654(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1655 github.com/goccy/llamawasm2go/p2.Fn1655
func Fn1655(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1656 github.com/goccy/llamawasm2go/p2.Fn1656
func Fn1656(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1657 github.com/goccy/llamawasm2go/p2.Fn1657
func Fn1657(m *base.Module)

//go:linkname Fn1658 github.com/goccy/llamawasm2go/p2.Fn1658
func Fn1658(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1660 github.com/goccy/llamawasm2go/p2.Fn1660
func Fn1660(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1662 github.com/goccy/llamawasm2go/p2.Fn1662
func Fn1662(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1663 github.com/goccy/llamawasm2go/p2.Fn1663
func Fn1663(m *base.Module, l0 int64)

//go:linkname Fn1667 github.com/goccy/llamawasm2go/p2.Fn1667
func Fn1667(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1668 github.com/goccy/llamawasm2go/p2.Fn1668
func Fn1668(m *base.Module, l0 int64)

//go:linkname Fn1671 github.com/goccy/llamawasm2go/p2.Fn1671
func Fn1671(m *base.Module, l0 int64)

//go:linkname Fn1681 github.com/goccy/llamawasm2go/p2.Fn1681
func Fn1681(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1682 github.com/goccy/llamawasm2go/p2.Fn1682
func Fn1682(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1691 github.com/goccy/llamawasm2go/p2.Fn1691
func Fn1691(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1692 github.com/goccy/llamawasm2go/p0.Fn1692
func Fn1692(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1694 github.com/goccy/llamawasm2go/p2.Fn1694
func Fn1694(m *base.Module, l0 int64)

//go:linkname Fn1696 github.com/goccy/llamawasm2go/p2.Fn1696
func Fn1696(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1703 github.com/goccy/llamawasm2go/p2.Fn1703
func Fn1703(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1704 github.com/goccy/llamawasm2go/p2.Fn1704
func Fn1704(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1711 github.com/goccy/llamawasm2go/p2.Fn1711
func Fn1711(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1718 github.com/goccy/llamawasm2go/p2.Fn1718
func Fn1718(m *base.Module, l0 int64)

//go:linkname Fn1721 github.com/goccy/llamawasm2go/p2.Fn1721
func Fn1721(m *base.Module, l0 int64) int32

//go:linkname Fn1731 github.com/goccy/llamawasm2go/p2.Fn1731
func Fn1731(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1733 github.com/goccy/llamawasm2go/p2.Fn1733
func Fn1733(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1734 github.com/goccy/llamawasm2go/p2.Fn1734
func Fn1734(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1735 github.com/goccy/llamawasm2go/p2.Fn1735
func Fn1735(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1736 github.com/goccy/llamawasm2go/p2.Fn1736
func Fn1736(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1737 github.com/goccy/llamawasm2go/p2.Fn1737
func Fn1737(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1753 github.com/goccy/llamawasm2go/p2.Fn1753
func Fn1753(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1766 github.com/goccy/llamawasm2go/p2.Fn1766
func Fn1766(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1773 github.com/goccy/llamawasm2go/p2.Fn1773
func Fn1773(m *base.Module, l0 int64)

//go:linkname Fn1807 github.com/goccy/llamawasm2go/p2.Fn1807
func Fn1807(m *base.Module, l0 int64)

//go:linkname Fn1810 github.com/goccy/llamawasm2go/p2.Fn1810
func Fn1810(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1824 github.com/goccy/llamawasm2go/p2.Fn1824
func Fn1824(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1825 github.com/goccy/llamawasm2go/p2.Fn1825
func Fn1825(m *base.Module, l0 int64) int64

//go:linkname Fn1826 github.com/goccy/llamawasm2go/p2.Fn1826
func Fn1826(m *base.Module, l0 int64)

//go:linkname Fn1830 github.com/goccy/llamawasm2go/p0.Fn1830
func Fn1830(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1831 github.com/goccy/llamawasm2go/p2.Fn1831
func Fn1831(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1832 github.com/goccy/llamawasm2go/p2.Fn1832
func Fn1832(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1833 github.com/goccy/llamawasm2go/p2.Fn1833
func Fn1833(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1834 github.com/goccy/llamawasm2go/p2.Fn1834
func Fn1834(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1852 github.com/goccy/llamawasm2go/p2.Fn1852
func Fn1852(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1863 github.com/goccy/llamawasm2go/p2.Fn1863
func Fn1863(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1865 github.com/goccy/llamawasm2go/p2.Fn1865
func Fn1865(m *base.Module, l0 int64) int64

//go:linkname Fn1866 github.com/goccy/llamawasm2go/p2.Fn1866
func Fn1866(m *base.Module, l0 int64)

//go:linkname Fn1869 github.com/goccy/llamawasm2go/p0.Fn1869
func Fn1869(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1871 github.com/goccy/llamawasm2go/p2.Fn1871
func Fn1871(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1882 github.com/goccy/llamawasm2go/p2.Fn1882
func Fn1882(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1885 github.com/goccy/llamawasm2go/p2.Fn1885
func Fn1885(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1886 github.com/goccy/llamawasm2go/p2.Fn1886
func Fn1886(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1896 github.com/goccy/llamawasm2go/p2.Fn1896
func Fn1896(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1923 github.com/goccy/llamawasm2go/p2.Fn1923
func Fn1923(m *base.Module, l0 int64)

//go:linkname Fn1924 github.com/goccy/llamawasm2go/p2.Fn1924
func Fn1924(m *base.Module, l0 int64)

//go:linkname Fn1935 github.com/goccy/llamawasm2go/p2.Fn1935
func Fn1935(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1967 github.com/goccy/llamawasm2go/p2.Fn1967
func Fn1967(m *base.Module) int64

//go:linkname Fn1972 github.com/goccy/llamawasm2go/p2.Fn1972
func Fn1972(m *base.Module, l0 int64) int64

//go:linkname Fn1973 github.com/goccy/llamawasm2go/p2.Fn1973
func Fn1973(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1980 github.com/goccy/llamawasm2go/p2.Fn1980
func Fn1980(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1984 github.com/goccy/llamawasm2go/p2.Fn1984
func Fn1984(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1986 github.com/goccy/llamawasm2go/p2.Fn1986
func Fn1986(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1987 github.com/goccy/llamawasm2go/p2.Fn1987
func Fn1987(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1988 github.com/goccy/llamawasm2go/p2.Fn1988
func Fn1988(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1989 github.com/goccy/llamawasm2go/p2.Fn1989
func Fn1989(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1990 github.com/goccy/llamawasm2go/p2.Fn1990
func Fn1990(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1994 github.com/goccy/llamawasm2go/p2.Fn1994
func Fn1994(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1995 github.com/goccy/llamawasm2go/p2.Fn1995
func Fn1995(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2001 github.com/goccy/llamawasm2go/p2.Fn2001
func Fn2001(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2003 github.com/goccy/llamawasm2go/p2.Fn2003
func Fn2003(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2006 github.com/goccy/llamawasm2go/p2.Fn2006
func Fn2006(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2007 github.com/goccy/llamawasm2go/p2.Fn2007
func Fn2007(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2008 github.com/goccy/llamawasm2go/p2.Fn2008
func Fn2008(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn2011 github.com/goccy/llamawasm2go/p2.Fn2011
func Fn2011(m *base.Module, l0 int64)

//go:linkname Fn2021 github.com/goccy/llamawasm2go/p2.Fn2021
func Fn2021(m *base.Module, l0 int64)

//go:linkname Fn2023 github.com/goccy/llamawasm2go/p2.Fn2023
func Fn2023(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2024 github.com/goccy/llamawasm2go/p2.Fn2024
func Fn2024(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2037 github.com/goccy/llamawasm2go/p2.Fn2037
func Fn2037(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2038 github.com/goccy/llamawasm2go/p2.Fn2038
func Fn2038(m *base.Module, l0 int64) int64

//go:linkname Fn2039 github.com/goccy/llamawasm2go/p2.Fn2039
func Fn2039(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2041 github.com/goccy/llamawasm2go/p2.Fn2041
func Fn2041(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2055 github.com/goccy/llamawasm2go/p2.Fn2055
func Fn2055(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2056 github.com/goccy/llamawasm2go/p2.Fn2056
func Fn2056(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2058 github.com/goccy/llamawasm2go/p2.Fn2058
func Fn2058(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2059 github.com/goccy/llamawasm2go/p2.Fn2059
func Fn2059(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2060 github.com/goccy/llamawasm2go/p2.Fn2060
func Fn2060(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2061 github.com/goccy/llamawasm2go/p2.Fn2061
func Fn2061(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2062 github.com/goccy/llamawasm2go/p2.Fn2062
func Fn2062(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2063 github.com/goccy/llamawasm2go/p2.Fn2063
func Fn2063(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2082 github.com/goccy/llamawasm2go/p2.Fn2082
func Fn2082(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2083 github.com/goccy/llamawasm2go/p2.Fn2083
func Fn2083(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2084 github.com/goccy/llamawasm2go/p2.Fn2084
func Fn2084(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2086 github.com/goccy/llamawasm2go/p2.Fn2086
func Fn2086(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2087 github.com/goccy/llamawasm2go/p2.Fn2087
func Fn2087(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2088 github.com/goccy/llamawasm2go/p2.Fn2088
func Fn2088(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2089 github.com/goccy/llamawasm2go/p2.Fn2089
func Fn2089(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2090 github.com/goccy/llamawasm2go/p2.Fn2090
func Fn2090(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2092 github.com/goccy/llamawasm2go/p2.Fn2092
func Fn2092(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 int64)

//go:linkname Fn2111 github.com/goccy/llamawasm2go/p2.Fn2111
func Fn2111(m *base.Module, l0 int64)

//go:linkname Fn2112 github.com/goccy/llamawasm2go/p2.Fn2112
func Fn2112(m *base.Module, l0 int64)

//go:linkname Fn2113 github.com/goccy/llamawasm2go/p2.Fn2113
func Fn2113(m *base.Module, l0 int64)

//go:linkname Fn2115 github.com/goccy/llamawasm2go/p2.Fn2115
func Fn2115(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2116 github.com/goccy/llamawasm2go/p2.Fn2116
func Fn2116(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2155 github.com/goccy/llamawasm2go/p2.Fn2155
func Fn2155(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2164 github.com/goccy/llamawasm2go/p2.Fn2164
func Fn2164(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2166 github.com/goccy/llamawasm2go/p2.Fn2166
func Fn2166(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2168 github.com/goccy/llamawasm2go/p2.Fn2168
func Fn2168(m *base.Module, l0 int64) int64

//go:linkname Fn2171 github.com/goccy/llamawasm2go/p2.Fn2171
func Fn2171(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2175 github.com/goccy/llamawasm2go/p2.Fn2175
func Fn2175(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2180 github.com/goccy/llamawasm2go/p2.Fn2180
func Fn2180(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2194 github.com/goccy/llamawasm2go/p2.Fn2194
func Fn2194(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2196 github.com/goccy/llamawasm2go/p2.Fn2196
func Fn2196(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2197 github.com/goccy/llamawasm2go/p2.Fn2197
func Fn2197(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2211 github.com/goccy/llamawasm2go/p2.Fn2211
func Fn2211(m *base.Module, l0 int64) int64

//go:linkname Fn2213 github.com/goccy/llamawasm2go/p2.Fn2213
func Fn2213(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2216 github.com/goccy/llamawasm2go/p2.Fn2216
func Fn2216(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2220 github.com/goccy/llamawasm2go/p2.Fn2220
func Fn2220(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2221 github.com/goccy/llamawasm2go/p2.Fn2221
func Fn2221(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2222 github.com/goccy/llamawasm2go/p0.Fn2222
func Fn2222(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2224 github.com/goccy/llamawasm2go/p2.Fn2224
func Fn2224(m *base.Module, l0 int64) int64

//go:linkname Fn2230 github.com/goccy/llamawasm2go/p2.Fn2230
func Fn2230(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2234 github.com/goccy/llamawasm2go/p2.Fn2234
func Fn2234(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2235 github.com/goccy/llamawasm2go/p2.Fn2235
func Fn2235(m *base.Module, l0 int64)

//go:linkname Fn2236 github.com/goccy/llamawasm2go/p2.Fn2236
func Fn2236(m *base.Module, l0 int64)

//go:linkname Fn2237 github.com/goccy/llamawasm2go/p2.Fn2237
func Fn2237(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2239 github.com/goccy/llamawasm2go/p2.Fn2239
func Fn2239(m *base.Module, l0 int64)

//go:linkname Fn2240 github.com/goccy/llamawasm2go/p2.Fn2240
func Fn2240(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2241 github.com/goccy/llamawasm2go/p2.Fn2241
func Fn2241(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2245 github.com/goccy/llamawasm2go/p2.Fn2245
func Fn2245(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2248 github.com/goccy/llamawasm2go/p2.Fn2248
func Fn2248(m *base.Module, l0 int64) int64

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p2.Fn2249
func Fn2249(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2250 github.com/goccy/llamawasm2go/p2.Fn2250
func Fn2250(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2257 github.com/goccy/llamawasm2go/p2.Fn2257
func Fn2257(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2258 github.com/goccy/llamawasm2go/p2.Fn2258
func Fn2258(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2260 github.com/goccy/llamawasm2go/p2.Fn2260
func Fn2260(m *base.Module) int64

//go:linkname Fn2262 github.com/goccy/llamawasm2go/p2.Fn2262
func Fn2262(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2263 github.com/goccy/llamawasm2go/p2.Fn2263
func Fn2263(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2265 github.com/goccy/llamawasm2go/p2.Fn2265
func Fn2265(m *base.Module) int64

//go:linkname Fn2267 github.com/goccy/llamawasm2go/p2.Fn2267
func Fn2267(m *base.Module, l0 int32) int64

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 int32) int32

//go:linkname Fn2269 github.com/goccy/llamawasm2go/p2.Fn2269
func Fn2269(m *base.Module, l0 int32) int64

//go:linkname Fn2270 github.com/goccy/llamawasm2go/p2.Fn2270
func Fn2270(m *base.Module, l0 float32) int64

//go:linkname Fn2271 github.com/goccy/llamawasm2go/p2.Fn2271
func Fn2271(m *base.Module, l0 float32) int64

//go:linkname Fn2272 github.com/goccy/llamawasm2go/p2.Fn2272
func Fn2272(m *base.Module, l0 float32) int64

//go:linkname Fn2274 github.com/goccy/llamawasm2go/p2.Fn2274
func Fn2274(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2275 github.com/goccy/llamawasm2go/p2.Fn2275
func Fn2275(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2316 github.com/goccy/llamawasm2go/p2.Fn2316
func Fn2316(m *base.Module, l0 int64)

//go:linkname Fn2318 github.com/goccy/llamawasm2go/p2.Fn2318
func Fn2318(m *base.Module, l0 int64)

//go:linkname Fn2360 github.com/goccy/llamawasm2go/p2.Fn2360
func Fn2360(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2371 github.com/goccy/llamawasm2go/p2.Fn2371
func Fn2371(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2374 github.com/goccy/llamawasm2go/p2.Fn2374
func Fn2374(m *base.Module, l0 int64)

//go:linkname Fn2375 github.com/goccy/llamawasm2go/p2.Fn2375
func Fn2375(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2376 github.com/goccy/llamawasm2go/p2.Fn2376
func Fn2376(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2378 github.com/goccy/llamawasm2go/p2.Fn2378
func Fn2378(m *base.Module, l0 int64)

//go:linkname Fn2382 github.com/goccy/llamawasm2go/p2.Fn2382
func Fn2382(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2388 github.com/goccy/llamawasm2go/p2.Fn2388
func Fn2388(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2403 github.com/goccy/llamawasm2go/p2.Fn2403
func Fn2403(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2410 github.com/goccy/llamawasm2go/p0.Fn2410
func Fn2410(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2413 github.com/goccy/llamawasm2go/p2.Fn2413
func Fn2413(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2416 github.com/goccy/llamawasm2go/p2.Fn2416
func Fn2416(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2419 github.com/goccy/llamawasm2go/p2.Fn2419
func Fn2419(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2420 github.com/goccy/llamawasm2go/p2.Fn2420
func Fn2420(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2421 github.com/goccy/llamawasm2go/p2.Fn2421
func Fn2421(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2426 github.com/goccy/llamawasm2go/p2.Fn2426
func Fn2426(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2428 github.com/goccy/llamawasm2go/p2.Fn2428
func Fn2428(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2445 github.com/goccy/llamawasm2go/p2.Fn2445
func Fn2445(m *base.Module, l0 int64)

//go:linkname Fn2446 github.com/goccy/llamawasm2go/p2.Fn2446
func Fn2446(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2451 github.com/goccy/llamawasm2go/p2.Fn2451
func Fn2451(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2456 github.com/goccy/llamawasm2go/p2.Fn2456
func Fn2456(m *base.Module, l0 int64) int64

//go:linkname Fn2457 github.com/goccy/llamawasm2go/p2.Fn2457
func Fn2457(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2458 github.com/goccy/llamawasm2go/p0.Fn2458
func Fn2458(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2469 github.com/goccy/llamawasm2go/p2.Fn2469
func Fn2469(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2470 github.com/goccy/llamawasm2go/p2.Fn2470
func Fn2470(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2471 github.com/goccy/llamawasm2go/p2.Fn2471
func Fn2471(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2472 github.com/goccy/llamawasm2go/p2.Fn2472
func Fn2472(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2492 github.com/goccy/llamawasm2go/p2.Fn2492
func Fn2492(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2493 github.com/goccy/llamawasm2go/p2.Fn2493
func Fn2493(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2496 github.com/goccy/llamawasm2go/p2.Fn2496
func Fn2496(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2582 github.com/goccy/llamawasm2go/p2.Fn2582
func Fn2582(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2701 github.com/goccy/llamawasm2go/p2.Fn2701
func Fn2701(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2702 github.com/goccy/llamawasm2go/p0.Fn2702
func Fn2702(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2703 github.com/goccy/llamawasm2go/p2.Fn2703
func Fn2703(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2704 github.com/goccy/llamawasm2go/p2.Fn2704
func Fn2704(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2708 github.com/goccy/llamawasm2go/p2.Fn2708
func Fn2708(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2733 github.com/goccy/llamawasm2go/p2.Fn2733
func Fn2733(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2865 github.com/goccy/llamawasm2go/p2.Fn2865
func Fn2865(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2872 github.com/goccy/llamawasm2go/p2.Fn2872
func Fn2872(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2900 github.com/goccy/llamawasm2go/p2.Fn2900
func Fn2900(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2914 github.com/goccy/llamawasm2go/p2.Fn2914
func Fn2914(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2926 github.com/goccy/llamawasm2go/p2.Fn2926
func Fn2926(m *base.Module, l0 int32)

//go:linkname Fn2928 github.com/goccy/llamawasm2go/p2.Fn2928
func Fn2928(m *base.Module, l0 int64) int64

//go:linkname Fn2929 github.com/goccy/llamawasm2go/p2.Fn2929
func Fn2929(m *base.Module, l0 int64)

//go:linkname Fn2932 github.com/goccy/llamawasm2go/p2.Fn2932
func Fn2932(m *base.Module, l0 int64)

//go:linkname Fn2933 github.com/goccy/llamawasm2go/p2.Fn2933
func Fn2933(m *base.Module, l0 int64)

//go:linkname Fn2935 github.com/goccy/llamawasm2go/p2.Fn2935
func Fn2935(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2936 github.com/goccy/llamawasm2go/p2.Fn2936
func Fn2936(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2942 github.com/goccy/llamawasm2go/p2.Fn2942
func Fn2942(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2944 github.com/goccy/llamawasm2go/p2.Fn2944
func Fn2944(m *base.Module, l0 int64) int32

//go:linkname Fn2948 github.com/goccy/llamawasm2go/p2.Fn2948
func Fn2948(m *base.Module) int32

//go:linkname Fn2959 github.com/goccy/llamawasm2go/p2.Fn2959
func Fn2959(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2963 github.com/goccy/llamawasm2go/p2.Fn2963
func Fn2963(m *base.Module) int64

//go:linkname Fn2965 github.com/goccy/llamawasm2go/p2.Fn2965
func Fn2965(m *base.Module, l0 float64) float32

//go:linkname Fn2966 github.com/goccy/llamawasm2go/p2.Fn2966
func Fn2966(m *base.Module, l0 float64) float32

//go:linkname Fn2970 github.com/goccy/llamawasm2go/p2.Fn2970
func Fn2970(m *base.Module, l0 float64) float64

//go:linkname Fn2973 github.com/goccy/llamawasm2go/p2.Fn2973
func Fn2973(m *base.Module, l0 int32) float32

//go:linkname Fn2974 github.com/goccy/llamawasm2go/p2.Fn2974
func Fn2974(m *base.Module, l0 int32) float32

//go:linkname Fn2977 github.com/goccy/llamawasm2go/p2.Fn2977
func Fn2977(m *base.Module, l0 float32) float32

//go:linkname Fn2980 github.com/goccy/llamawasm2go/p2.Fn2980
func Fn2980(m *base.Module, l0 float64) float64

//go:linkname Fn2981 github.com/goccy/llamawasm2go/p2.Fn2981
func Fn2981(m *base.Module, l0 float64) float64

//go:linkname Fn2982 github.com/goccy/llamawasm2go/p2.Fn2982
func Fn2982(m *base.Module, l0 float32) float32

//go:linkname Fn2984 github.com/goccy/llamawasm2go/p2.Fn2984
func Fn2984(m *base.Module, l0 float32) float32

//go:linkname Fn2986 github.com/goccy/llamawasm2go/p2.Fn2986
func Fn2986(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn2987 github.com/goccy/llamawasm2go/p2.Fn2987
func Fn2987(m *base.Module, l0 float32) float32

//go:linkname Fn3004 github.com/goccy/llamawasm2go/p2.Fn3004
func Fn3004(m *base.Module, l0 int64) int32

//go:linkname Fn3005 github.com/goccy/llamawasm2go/p2.Fn3005
func Fn3005(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3007 github.com/goccy/llamawasm2go/p2.Fn3007
func Fn3007(m *base.Module, l0 int64)

//go:linkname Fn3008 github.com/goccy/llamawasm2go/p2.Fn3008
func Fn3008(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3009 github.com/goccy/llamawasm2go/p2.Fn3009
func Fn3009(m *base.Module, l0 int64) int32

//go:linkname Fn3016 github.com/goccy/llamawasm2go/p2.Fn3016
func Fn3016(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3018 github.com/goccy/llamawasm2go/p2.Fn3018
func Fn3018(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3024 github.com/goccy/llamawasm2go/p2.Fn3024
func Fn3024(m *base.Module, l0 int64) int32

//go:linkname Fn3027 github.com/goccy/llamawasm2go/p2.Fn3027
func Fn3027(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3030 github.com/goccy/llamawasm2go/p2.Fn3030
func Fn3030(m *base.Module, l0 int64) int32

//go:linkname Fn3032 github.com/goccy/llamawasm2go/p2.Fn3032
func Fn3032(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3034 github.com/goccy/llamawasm2go/p2.Fn3034
func Fn3034(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3035 github.com/goccy/llamawasm2go/p2.Fn3035
func Fn3035(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3038 github.com/goccy/llamawasm2go/p2.Fn3038
func Fn3038(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3041 github.com/goccy/llamawasm2go/p2.Fn3041
func Fn3041(m *base.Module, l0 int64) int64

//go:linkname Fn3045 github.com/goccy/llamawasm2go/p2.Fn3045
func Fn3045(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3055 github.com/goccy/llamawasm2go/p2.Fn3055
func Fn3055(m *base.Module)

//go:linkname Fn3056 github.com/goccy/llamawasm2go/p0.Fn3056
func Fn3056(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3059 github.com/goccy/llamawasm2go/p2.Fn3059
func Fn3059(m *base.Module)

//go:linkname Fn3061 github.com/goccy/llamawasm2go/p0.Fn3061
func Fn3061(m *base.Module, l0 int64) int64

//go:linkname Fn3063 github.com/goccy/llamawasm2go/p2.Fn3063
func Fn3063(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3067 github.com/goccy/llamawasm2go/p2.Fn3067
func Fn3067(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3136 github.com/goccy/llamawasm2go/p2.Fn3136
func Fn3136(m *base.Module, l0 int32)
