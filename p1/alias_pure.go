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

//go:linkname Fn471 github.com/goccy/llamawasm2go/p2.Fn471
func Fn471(m *base.Module, l0 int64, l1 int64) int64

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

//go:linkname Fn530 github.com/goccy/llamawasm2go/p2.Fn530
func Fn530(m *base.Module, l0 int64, l1 int64, l2 int32) int64

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

//go:linkname Fn972 github.com/goccy/llamawasm2go/p2.Fn972
func Fn972(m *base.Module, l0 int64) int64

//go:linkname Fn982 github.com/goccy/llamawasm2go/p2.Fn982
func Fn982(m *base.Module, l0 int64)

//go:linkname Fn987 github.com/goccy/llamawasm2go/p2.Fn987
func Fn987(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn988 github.com/goccy/llamawasm2go/p2.Fn988
func Fn988(m *base.Module)

//go:linkname Fn997 github.com/goccy/llamawasm2go/p2.Fn997
func Fn997(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1005 github.com/goccy/llamawasm2go/p2.Fn1005
func Fn1005(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1006 github.com/goccy/llamawasm2go/p2.Fn1006
func Fn1006(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1007 github.com/goccy/llamawasm2go/p2.Fn1007
func Fn1007(m *base.Module, l0 int64, l1 int64, l2 int64) float32

//go:linkname Fn1009 github.com/goccy/llamawasm2go/p2.Fn1009
func Fn1009(m *base.Module, l0 int64, l1 int64, l2 int64) float64

//go:linkname Fn1011 github.com/goccy/llamawasm2go/p2.Fn1011
func Fn1011(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1014 github.com/goccy/llamawasm2go/p2.Fn1014
func Fn1014(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1017 github.com/goccy/llamawasm2go/p2.Fn1017
func Fn1017(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1023 github.com/goccy/llamawasm2go/p2.Fn1023
func Fn1023(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1024 github.com/goccy/llamawasm2go/p2.Fn1024
func Fn1024(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1032 github.com/goccy/llamawasm2go/p2.Fn1032
func Fn1032(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1037 github.com/goccy/llamawasm2go/p2.Fn1037
func Fn1037(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1038 github.com/goccy/llamawasm2go/p2.Fn1038
func Fn1038(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1046 github.com/goccy/llamawasm2go/p2.Fn1046
func Fn1046(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1062 github.com/goccy/llamawasm2go/p2.Fn1062
func Fn1062(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1070 github.com/goccy/llamawasm2go/p2.Fn1070
func Fn1070(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1086 github.com/goccy/llamawasm2go/p0.Fn1086
func Fn1086(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1089 github.com/goccy/llamawasm2go/p2.Fn1089
func Fn1089(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1093 github.com/goccy/llamawasm2go/p2.Fn1093
func Fn1093(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1103 github.com/goccy/llamawasm2go/p0.Fn1103
func Fn1103(m *base.Module, l0 int64) int64

//go:linkname Fn1104 github.com/goccy/llamawasm2go/p2.Fn1104
func Fn1104(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1105 github.com/goccy/llamawasm2go/p2.Fn1105
func Fn1105(m *base.Module, l0 int64)

//go:linkname Fn1107 github.com/goccy/llamawasm2go/p2.Fn1107
func Fn1107(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1197 github.com/goccy/llamawasm2go/p2.Fn1197
func Fn1197(m *base.Module, l0 int64)

//go:linkname Fn1219 github.com/goccy/llamawasm2go/p2.Fn1219
func Fn1219(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1229 github.com/goccy/llamawasm2go/p2.Fn1229
func Fn1229(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1235 github.com/goccy/llamawasm2go/p2.Fn1235
func Fn1235(m *base.Module)

//go:linkname Fn1239 github.com/goccy/llamawasm2go/p2.Fn1239
func Fn1239(m *base.Module, l0 int64) int64

//go:linkname Fn1267 github.com/goccy/llamawasm2go/p2.Fn1267
func Fn1267(m *base.Module, l0 int64)

//go:linkname Fn1271 github.com/goccy/llamawasm2go/p2.Fn1271
func Fn1271(m *base.Module, l0 int32) int64

//go:linkname Fn1282 github.com/goccy/llamawasm2go/p2.Fn1282
func Fn1282(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1283 github.com/goccy/llamawasm2go/p2.Fn1283
func Fn1283(m *base.Module, l0 int64)

//go:linkname Fn1285 github.com/goccy/llamawasm2go/p2.Fn1285
func Fn1285(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1306 github.com/goccy/llamawasm2go/p2.Fn1306
func Fn1306(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1307 github.com/goccy/llamawasm2go/p2.Fn1307
func Fn1307(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1308 github.com/goccy/llamawasm2go/p2.Fn1308
func Fn1308(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1309 github.com/goccy/llamawasm2go/p2.Fn1309
func Fn1309(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1310 github.com/goccy/llamawasm2go/p2.Fn1310
func Fn1310(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1311 github.com/goccy/llamawasm2go/p2.Fn1311
func Fn1311(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1312 github.com/goccy/llamawasm2go/p2.Fn1312
func Fn1312(m *base.Module, l0 int64) int64

//go:linkname Fn1315 github.com/goccy/llamawasm2go/p2.Fn1315
func Fn1315(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1319 github.com/goccy/llamawasm2go/p2.Fn1319
func Fn1319(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1320 github.com/goccy/llamawasm2go/p2.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1321 github.com/goccy/llamawasm2go/p2.Fn1321
func Fn1321(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1322 github.com/goccy/llamawasm2go/p2.Fn1322
func Fn1322(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p2.Fn1324
func Fn1324(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1325 github.com/goccy/llamawasm2go/p2.Fn1325
func Fn1325(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1326 github.com/goccy/llamawasm2go/p2.Fn1326
func Fn1326(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1327 github.com/goccy/llamawasm2go/p2.Fn1327
func Fn1327(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1328 github.com/goccy/llamawasm2go/p2.Fn1328
func Fn1328(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1329 github.com/goccy/llamawasm2go/p2.Fn1329
func Fn1329(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1330 github.com/goccy/llamawasm2go/p2.Fn1330
func Fn1330(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1331 github.com/goccy/llamawasm2go/p2.Fn1331
func Fn1331(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1332 github.com/goccy/llamawasm2go/p2.Fn1332
func Fn1332(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1333 github.com/goccy/llamawasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1335 github.com/goccy/llamawasm2go/p2.Fn1335
func Fn1335(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1336 github.com/goccy/llamawasm2go/p2.Fn1336
func Fn1336(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1337 github.com/goccy/llamawasm2go/p2.Fn1337
func Fn1337(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1338 github.com/goccy/llamawasm2go/p2.Fn1338
func Fn1338(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1339 github.com/goccy/llamawasm2go/p2.Fn1339
func Fn1339(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1340 github.com/goccy/llamawasm2go/p2.Fn1340
func Fn1340(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1341 github.com/goccy/llamawasm2go/p2.Fn1341
func Fn1341(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1342 github.com/goccy/llamawasm2go/p2.Fn1342
func Fn1342(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1343 github.com/goccy/llamawasm2go/p2.Fn1343
func Fn1343(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1344 github.com/goccy/llamawasm2go/p2.Fn1344
func Fn1344(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1345 github.com/goccy/llamawasm2go/p2.Fn1345
func Fn1345(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1347 github.com/goccy/llamawasm2go/p2.Fn1347
func Fn1347(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1351 github.com/goccy/llamawasm2go/p2.Fn1351
func Fn1351(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1352 github.com/goccy/llamawasm2go/p2.Fn1352
func Fn1352(m *base.Module)

//go:linkname Fn1353 github.com/goccy/llamawasm2go/p2.Fn1353
func Fn1353(m *base.Module)

//go:linkname Fn1354 github.com/goccy/llamawasm2go/p0.Fn1354
func Fn1354(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1362 github.com/goccy/llamawasm2go/p2.Fn1362
func Fn1362(m *base.Module)

//go:linkname Fn1364 github.com/goccy/llamawasm2go/p2.Fn1364
func Fn1364(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1366 github.com/goccy/llamawasm2go/p0.Fn1366
func Fn1366(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1367 github.com/goccy/llamawasm2go/p2.Fn1367
func Fn1367(m *base.Module, l0 int64) int64

//go:linkname Fn1372 github.com/goccy/llamawasm2go/p2.Fn1372
func Fn1372(m *base.Module, l0 int64)

//go:linkname Fn1379 github.com/goccy/llamawasm2go/p2.Fn1379
func Fn1379(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1382 github.com/goccy/llamawasm2go/p2.Fn1382
func Fn1382(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1384 github.com/goccy/llamawasm2go/p2.Fn1384
func Fn1384(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1386 github.com/goccy/llamawasm2go/p2.Fn1386
func Fn1386(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1388 github.com/goccy/llamawasm2go/p2.Fn1388
func Fn1388(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1390 github.com/goccy/llamawasm2go/p2.Fn1390
func Fn1390(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p2.Fn1392
func Fn1392(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1399 github.com/goccy/llamawasm2go/p2.Fn1399
func Fn1399(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p0.Fn1401
func Fn1401(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module, l0 int64)

//go:linkname Fn1404 github.com/goccy/llamawasm2go/p2.Fn1404
func Fn1404(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p2.Fn1406
func Fn1406(m *base.Module, l0 int64) int64

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p2.Fn1407
func Fn1407(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p2.Fn1408
func Fn1408(m *base.Module, l0 int64)

//go:linkname Fn1409 github.com/goccy/llamawasm2go/p2.Fn1409
func Fn1409(m *base.Module, l0 int64)

//go:linkname Fn1410 github.com/goccy/llamawasm2go/p2.Fn1410
func Fn1410(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1411 github.com/goccy/llamawasm2go/p2.Fn1411
func Fn1411(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1414 github.com/goccy/llamawasm2go/p2.Fn1414
func Fn1414(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1422 github.com/goccy/llamawasm2go/p2.Fn1422
func Fn1422(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p0.Fn1424
func Fn1424(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p2.Fn1426
func Fn1426(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1429 github.com/goccy/llamawasm2go/p0.Fn1429
func Fn1429(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1430 github.com/goccy/llamawasm2go/p2.Fn1430
func Fn1430(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1434 github.com/goccy/llamawasm2go/p2.Fn1434
func Fn1434(m *base.Module, l0 int64)

//go:linkname Fn1437 github.com/goccy/llamawasm2go/p2.Fn1437
func Fn1437(m *base.Module, l0 int64)

//go:linkname Fn1440 github.com/goccy/llamawasm2go/p2.Fn1440
func Fn1440(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1441 github.com/goccy/llamawasm2go/p2.Fn1441
func Fn1441(m *base.Module, l0 int64) int64

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1446 github.com/goccy/llamawasm2go/p2.Fn1446
func Fn1446(m *base.Module, l0 int64)

//go:linkname Fn1449 github.com/goccy/llamawasm2go/p2.Fn1449
func Fn1449(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1451 github.com/goccy/llamawasm2go/p2.Fn1451
func Fn1451(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1454 github.com/goccy/llamawasm2go/p2.Fn1454
func Fn1454(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1461 github.com/goccy/llamawasm2go/p2.Fn1461
func Fn1461(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1462 github.com/goccy/llamawasm2go/p2.Fn1462
func Fn1462(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1464 github.com/goccy/llamawasm2go/p2.Fn1464
func Fn1464(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1468 github.com/goccy/llamawasm2go/p2.Fn1468
func Fn1468(m *base.Module, l0 int64)

//go:linkname Fn1469 github.com/goccy/llamawasm2go/p2.Fn1469
func Fn1469(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1471 github.com/goccy/llamawasm2go/p2.Fn1471
func Fn1471(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1472 github.com/goccy/llamawasm2go/p2.Fn1472
func Fn1472(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1473 github.com/goccy/llamawasm2go/p2.Fn1473
func Fn1473(m *base.Module, l0 int64)

//go:linkname Fn1475 github.com/goccy/llamawasm2go/p2.Fn1475
func Fn1475(m *base.Module, l0 int64) int32

//go:linkname Fn1476 github.com/goccy/llamawasm2go/p2.Fn1476
func Fn1476(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1518 github.com/goccy/llamawasm2go/p2.Fn1518
func Fn1518(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1534 github.com/goccy/llamawasm2go/p2.Fn1534
func Fn1534(m *base.Module, l0 int64)

//go:linkname Fn1535 github.com/goccy/llamawasm2go/p2.Fn1535
func Fn1535(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1536 github.com/goccy/llamawasm2go/p2.Fn1536
func Fn1536(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1537 github.com/goccy/llamawasm2go/p2.Fn1537
func Fn1537(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1538 github.com/goccy/llamawasm2go/p2.Fn1538
func Fn1538(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1539 github.com/goccy/llamawasm2go/p2.Fn1539
func Fn1539(m *base.Module, l0 int64)

//go:linkname Fn1540 github.com/goccy/llamawasm2go/p2.Fn1540
func Fn1540(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p2.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1543 github.com/goccy/llamawasm2go/p2.Fn1543
func Fn1543(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1544 github.com/goccy/llamawasm2go/p2.Fn1544
func Fn1544(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1546 github.com/goccy/llamawasm2go/p2.Fn1546
func Fn1546(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1549 github.com/goccy/llamawasm2go/p2.Fn1549
func Fn1549(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1551 github.com/goccy/llamawasm2go/p2.Fn1551
func Fn1551(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1552 github.com/goccy/llamawasm2go/p2.Fn1552
func Fn1552(m *base.Module, l0 int64) int64

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64) int64

//go:linkname Fn1554 github.com/goccy/llamawasm2go/p2.Fn1554
func Fn1554(m *base.Module, l0 int64) int64

//go:linkname Fn1555 github.com/goccy/llamawasm2go/p2.Fn1555
func Fn1555(m *base.Module, l0 int64) int64

//go:linkname Fn1556 github.com/goccy/llamawasm2go/p2.Fn1556
func Fn1556(m *base.Module, l0 int64) int64

//go:linkname Fn1557 github.com/goccy/llamawasm2go/p2.Fn1557
func Fn1557(m *base.Module, l0 int64) int64

//go:linkname Fn1559 github.com/goccy/llamawasm2go/p2.Fn1559
func Fn1559(m *base.Module, l0 int64) int64

//go:linkname Fn1560 github.com/goccy/llamawasm2go/p2.Fn1560
func Fn1560(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1562 github.com/goccy/llamawasm2go/p2.Fn1562
func Fn1562(m *base.Module, l0 int64) int64

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p2.Fn1563
func Fn1563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1564 github.com/goccy/llamawasm2go/p2.Fn1564
func Fn1564(m *base.Module, l0 int64) int64

//go:linkname Fn1566 github.com/goccy/llamawasm2go/p2.Fn1566
func Fn1566(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1567 github.com/goccy/llamawasm2go/p2.Fn1567
func Fn1567(m *base.Module, l0 int64) int64

//go:linkname Fn1569 github.com/goccy/llamawasm2go/p2.Fn1569
func Fn1569(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1570 github.com/goccy/llamawasm2go/p2.Fn1570
func Fn1570(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1571 github.com/goccy/llamawasm2go/p2.Fn1571
func Fn1571(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1572 github.com/goccy/llamawasm2go/p2.Fn1572
func Fn1572(m *base.Module, l0 int64) int64

//go:linkname Fn1574 github.com/goccy/llamawasm2go/p2.Fn1574
func Fn1574(m *base.Module, l0 int64) int64

//go:linkname Fn1575 github.com/goccy/llamawasm2go/p2.Fn1575
func Fn1575(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1580 github.com/goccy/llamawasm2go/p2.Fn1580
func Fn1580(m *base.Module, l0 int64) int64

//go:linkname Fn1581 github.com/goccy/llamawasm2go/p2.Fn1581
func Fn1581(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1582 github.com/goccy/llamawasm2go/p2.Fn1582
func Fn1582(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p2.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1585 github.com/goccy/llamawasm2go/p2.Fn1585
func Fn1585(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1586 github.com/goccy/llamawasm2go/p2.Fn1586
func Fn1586(m *base.Module, l0 int64) int64

//go:linkname Fn1587 github.com/goccy/llamawasm2go/p2.Fn1587
func Fn1587(m *base.Module, l0 int64) int64

//go:linkname Fn1589 github.com/goccy/llamawasm2go/p2.Fn1589
func Fn1589(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1590 github.com/goccy/llamawasm2go/p2.Fn1590
func Fn1590(m *base.Module, l0 int64)

//go:linkname Fn1619 github.com/goccy/llamawasm2go/p2.Fn1619
func Fn1619(m *base.Module, l0 int64) int64

//go:linkname Fn1626 github.com/goccy/llamawasm2go/p2.Fn1626
func Fn1626(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1627 github.com/goccy/llamawasm2go/p2.Fn1627
func Fn1627(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1628 github.com/goccy/llamawasm2go/p2.Fn1628
func Fn1628(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1630 github.com/goccy/llamawasm2go/p2.Fn1630
func Fn1630(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1631 github.com/goccy/llamawasm2go/p2.Fn1631
func Fn1631(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1635 github.com/goccy/llamawasm2go/p2.Fn1635
func Fn1635(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1636 github.com/goccy/llamawasm2go/p2.Fn1636
func Fn1636(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1637 github.com/goccy/llamawasm2go/p2.Fn1637
func Fn1637(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1638 github.com/goccy/llamawasm2go/p2.Fn1638
func Fn1638(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1639 github.com/goccy/llamawasm2go/p2.Fn1639
func Fn1639(m *base.Module, l0 int64) int32

//go:linkname Fn1640 github.com/goccy/llamawasm2go/p2.Fn1640
func Fn1640(m *base.Module, l0 int64) int32

//go:linkname Fn1641 github.com/goccy/llamawasm2go/p2.Fn1641
func Fn1641(m *base.Module, l0 int64) int32

//go:linkname Fn1642 github.com/goccy/llamawasm2go/p2.Fn1642
func Fn1642(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1645 github.com/goccy/llamawasm2go/p2.Fn1645
func Fn1645(m *base.Module, l0 int64) int32

//go:linkname Fn1646 github.com/goccy/llamawasm2go/p2.Fn1646
func Fn1646(m *base.Module, l0 int64) int32

//go:linkname Fn1652 github.com/goccy/llamawasm2go/p2.Fn1652
func Fn1652(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1653 github.com/goccy/llamawasm2go/p2.Fn1653
func Fn1653(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1654 github.com/goccy/llamawasm2go/p2.Fn1654
func Fn1654(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1655 github.com/goccy/llamawasm2go/p2.Fn1655
func Fn1655(m *base.Module)

//go:linkname Fn1656 github.com/goccy/llamawasm2go/p2.Fn1656
func Fn1656(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1658 github.com/goccy/llamawasm2go/p2.Fn1658
func Fn1658(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1660 github.com/goccy/llamawasm2go/p2.Fn1660
func Fn1660(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1661 github.com/goccy/llamawasm2go/p2.Fn1661
func Fn1661(m *base.Module, l0 int64)

//go:linkname Fn1665 github.com/goccy/llamawasm2go/p2.Fn1665
func Fn1665(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1666 github.com/goccy/llamawasm2go/p2.Fn1666
func Fn1666(m *base.Module, l0 int64)

//go:linkname Fn1669 github.com/goccy/llamawasm2go/p2.Fn1669
func Fn1669(m *base.Module, l0 int64)

//go:linkname Fn1679 github.com/goccy/llamawasm2go/p2.Fn1679
func Fn1679(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1680 github.com/goccy/llamawasm2go/p2.Fn1680
func Fn1680(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1689 github.com/goccy/llamawasm2go/p2.Fn1689
func Fn1689(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1690 github.com/goccy/llamawasm2go/p0.Fn1690
func Fn1690(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1692 github.com/goccy/llamawasm2go/p2.Fn1692
func Fn1692(m *base.Module, l0 int64)

//go:linkname Fn1694 github.com/goccy/llamawasm2go/p2.Fn1694
func Fn1694(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1701 github.com/goccy/llamawasm2go/p2.Fn1701
func Fn1701(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1702 github.com/goccy/llamawasm2go/p2.Fn1702
func Fn1702(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1709 github.com/goccy/llamawasm2go/p2.Fn1709
func Fn1709(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1716 github.com/goccy/llamawasm2go/p2.Fn1716
func Fn1716(m *base.Module, l0 int64)

//go:linkname Fn1719 github.com/goccy/llamawasm2go/p2.Fn1719
func Fn1719(m *base.Module, l0 int64) int32

//go:linkname Fn1729 github.com/goccy/llamawasm2go/p2.Fn1729
func Fn1729(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1731 github.com/goccy/llamawasm2go/p2.Fn1731
func Fn1731(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1732 github.com/goccy/llamawasm2go/p2.Fn1732
func Fn1732(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1733 github.com/goccy/llamawasm2go/p2.Fn1733
func Fn1733(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1734 github.com/goccy/llamawasm2go/p2.Fn1734
func Fn1734(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1735 github.com/goccy/llamawasm2go/p2.Fn1735
func Fn1735(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1751 github.com/goccy/llamawasm2go/p2.Fn1751
func Fn1751(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1764 github.com/goccy/llamawasm2go/p2.Fn1764
func Fn1764(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1771 github.com/goccy/llamawasm2go/p2.Fn1771
func Fn1771(m *base.Module, l0 int64)

//go:linkname Fn1805 github.com/goccy/llamawasm2go/p2.Fn1805
func Fn1805(m *base.Module, l0 int64)

//go:linkname Fn1808 github.com/goccy/llamawasm2go/p2.Fn1808
func Fn1808(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1822 github.com/goccy/llamawasm2go/p2.Fn1822
func Fn1822(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1823 github.com/goccy/llamawasm2go/p2.Fn1823
func Fn1823(m *base.Module, l0 int64) int64

//go:linkname Fn1824 github.com/goccy/llamawasm2go/p2.Fn1824
func Fn1824(m *base.Module, l0 int64)

//go:linkname Fn1828 github.com/goccy/llamawasm2go/p0.Fn1828
func Fn1828(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1829 github.com/goccy/llamawasm2go/p2.Fn1829
func Fn1829(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1830 github.com/goccy/llamawasm2go/p2.Fn1830
func Fn1830(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1831 github.com/goccy/llamawasm2go/p2.Fn1831
func Fn1831(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1832 github.com/goccy/llamawasm2go/p2.Fn1832
func Fn1832(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1850 github.com/goccy/llamawasm2go/p2.Fn1850
func Fn1850(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1861 github.com/goccy/llamawasm2go/p2.Fn1861
func Fn1861(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1863 github.com/goccy/llamawasm2go/p2.Fn1863
func Fn1863(m *base.Module, l0 int64) int64

//go:linkname Fn1864 github.com/goccy/llamawasm2go/p2.Fn1864
func Fn1864(m *base.Module, l0 int64)

//go:linkname Fn1867 github.com/goccy/llamawasm2go/p0.Fn1867
func Fn1867(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1869 github.com/goccy/llamawasm2go/p2.Fn1869
func Fn1869(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1880 github.com/goccy/llamawasm2go/p2.Fn1880
func Fn1880(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1883 github.com/goccy/llamawasm2go/p2.Fn1883
func Fn1883(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1884 github.com/goccy/llamawasm2go/p2.Fn1884
func Fn1884(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1894 github.com/goccy/llamawasm2go/p2.Fn1894
func Fn1894(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1921 github.com/goccy/llamawasm2go/p2.Fn1921
func Fn1921(m *base.Module, l0 int64)

//go:linkname Fn1922 github.com/goccy/llamawasm2go/p2.Fn1922
func Fn1922(m *base.Module, l0 int64)

//go:linkname Fn1933 github.com/goccy/llamawasm2go/p2.Fn1933
func Fn1933(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1965 github.com/goccy/llamawasm2go/p2.Fn1965
func Fn1965(m *base.Module) int64

//go:linkname Fn1970 github.com/goccy/llamawasm2go/p2.Fn1970
func Fn1970(m *base.Module, l0 int64) int64

//go:linkname Fn1971 github.com/goccy/llamawasm2go/p2.Fn1971
func Fn1971(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1978 github.com/goccy/llamawasm2go/p2.Fn1978
func Fn1978(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1982 github.com/goccy/llamawasm2go/p2.Fn1982
func Fn1982(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1984 github.com/goccy/llamawasm2go/p2.Fn1984
func Fn1984(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1985 github.com/goccy/llamawasm2go/p2.Fn1985
func Fn1985(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1986 github.com/goccy/llamawasm2go/p2.Fn1986
func Fn1986(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1987 github.com/goccy/llamawasm2go/p2.Fn1987
func Fn1987(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1988 github.com/goccy/llamawasm2go/p2.Fn1988
func Fn1988(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1992 github.com/goccy/llamawasm2go/p2.Fn1992
func Fn1992(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1993 github.com/goccy/llamawasm2go/p2.Fn1993
func Fn1993(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1999 github.com/goccy/llamawasm2go/p2.Fn1999
func Fn1999(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2001 github.com/goccy/llamawasm2go/p2.Fn2001
func Fn2001(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2004 github.com/goccy/llamawasm2go/p2.Fn2004
func Fn2004(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2005 github.com/goccy/llamawasm2go/p2.Fn2005
func Fn2005(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2006 github.com/goccy/llamawasm2go/p2.Fn2006
func Fn2006(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn2009 github.com/goccy/llamawasm2go/p2.Fn2009
func Fn2009(m *base.Module, l0 int64)

//go:linkname Fn2019 github.com/goccy/llamawasm2go/p2.Fn2019
func Fn2019(m *base.Module, l0 int64)

//go:linkname Fn2021 github.com/goccy/llamawasm2go/p2.Fn2021
func Fn2021(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2022 github.com/goccy/llamawasm2go/p2.Fn2022
func Fn2022(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2035 github.com/goccy/llamawasm2go/p2.Fn2035
func Fn2035(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2036 github.com/goccy/llamawasm2go/p2.Fn2036
func Fn2036(m *base.Module, l0 int64) int64

//go:linkname Fn2037 github.com/goccy/llamawasm2go/p2.Fn2037
func Fn2037(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2039 github.com/goccy/llamawasm2go/p2.Fn2039
func Fn2039(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2053 github.com/goccy/llamawasm2go/p2.Fn2053
func Fn2053(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2054 github.com/goccy/llamawasm2go/p2.Fn2054
func Fn2054(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2056 github.com/goccy/llamawasm2go/p2.Fn2056
func Fn2056(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2057 github.com/goccy/llamawasm2go/p2.Fn2057
func Fn2057(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2058 github.com/goccy/llamawasm2go/p2.Fn2058
func Fn2058(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2059 github.com/goccy/llamawasm2go/p2.Fn2059
func Fn2059(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2060 github.com/goccy/llamawasm2go/p2.Fn2060
func Fn2060(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2061 github.com/goccy/llamawasm2go/p2.Fn2061
func Fn2061(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2080 github.com/goccy/llamawasm2go/p2.Fn2080
func Fn2080(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2081 github.com/goccy/llamawasm2go/p2.Fn2081
func Fn2081(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2082 github.com/goccy/llamawasm2go/p2.Fn2082
func Fn2082(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2084 github.com/goccy/llamawasm2go/p2.Fn2084
func Fn2084(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2085 github.com/goccy/llamawasm2go/p2.Fn2085
func Fn2085(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2086 github.com/goccy/llamawasm2go/p2.Fn2086
func Fn2086(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2087 github.com/goccy/llamawasm2go/p2.Fn2087
func Fn2087(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2088 github.com/goccy/llamawasm2go/p2.Fn2088
func Fn2088(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2090 github.com/goccy/llamawasm2go/p2.Fn2090
func Fn2090(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2092 github.com/goccy/llamawasm2go/p2.Fn2092
func Fn2092(m *base.Module, l0 int64)

//go:linkname Fn2109 github.com/goccy/llamawasm2go/p2.Fn2109
func Fn2109(m *base.Module, l0 int64)

//go:linkname Fn2110 github.com/goccy/llamawasm2go/p2.Fn2110
func Fn2110(m *base.Module, l0 int64)

//go:linkname Fn2111 github.com/goccy/llamawasm2go/p2.Fn2111
func Fn2111(m *base.Module, l0 int64)

//go:linkname Fn2113 github.com/goccy/llamawasm2go/p2.Fn2113
func Fn2113(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2114 github.com/goccy/llamawasm2go/p2.Fn2114
func Fn2114(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2153 github.com/goccy/llamawasm2go/p2.Fn2153
func Fn2153(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2162 github.com/goccy/llamawasm2go/p2.Fn2162
func Fn2162(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2164 github.com/goccy/llamawasm2go/p2.Fn2164
func Fn2164(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2166 github.com/goccy/llamawasm2go/p2.Fn2166
func Fn2166(m *base.Module, l0 int64) int64

//go:linkname Fn2169 github.com/goccy/llamawasm2go/p2.Fn2169
func Fn2169(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2173 github.com/goccy/llamawasm2go/p2.Fn2173
func Fn2173(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2178 github.com/goccy/llamawasm2go/p2.Fn2178
func Fn2178(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2192 github.com/goccy/llamawasm2go/p2.Fn2192
func Fn2192(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2194 github.com/goccy/llamawasm2go/p2.Fn2194
func Fn2194(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2195 github.com/goccy/llamawasm2go/p2.Fn2195
func Fn2195(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2209 github.com/goccy/llamawasm2go/p2.Fn2209
func Fn2209(m *base.Module, l0 int64) int64

//go:linkname Fn2211 github.com/goccy/llamawasm2go/p2.Fn2211
func Fn2211(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2214 github.com/goccy/llamawasm2go/p2.Fn2214
func Fn2214(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2218 github.com/goccy/llamawasm2go/p2.Fn2218
func Fn2218(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2219 github.com/goccy/llamawasm2go/p2.Fn2219
func Fn2219(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2220 github.com/goccy/llamawasm2go/p0.Fn2220
func Fn2220(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2222 github.com/goccy/llamawasm2go/p2.Fn2222
func Fn2222(m *base.Module, l0 int64) int64

//go:linkname Fn2226 github.com/goccy/llamawasm2go/p2.Fn2226
func Fn2226(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2228 github.com/goccy/llamawasm2go/p2.Fn2228
func Fn2228(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2232 github.com/goccy/llamawasm2go/p2.Fn2232
func Fn2232(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2233 github.com/goccy/llamawasm2go/p2.Fn2233
func Fn2233(m *base.Module, l0 int64)

//go:linkname Fn2234 github.com/goccy/llamawasm2go/p2.Fn2234
func Fn2234(m *base.Module, l0 int64)

//go:linkname Fn2235 github.com/goccy/llamawasm2go/p2.Fn2235
func Fn2235(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2237 github.com/goccy/llamawasm2go/p2.Fn2237
func Fn2237(m *base.Module, l0 int64)

//go:linkname Fn2238 github.com/goccy/llamawasm2go/p2.Fn2238
func Fn2238(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2239 github.com/goccy/llamawasm2go/p2.Fn2239
func Fn2239(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2243 github.com/goccy/llamawasm2go/p2.Fn2243
func Fn2243(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2246 github.com/goccy/llamawasm2go/p2.Fn2246
func Fn2246(m *base.Module, l0 int64) int64

//go:linkname Fn2247 github.com/goccy/llamawasm2go/p2.Fn2247
func Fn2247(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2248 github.com/goccy/llamawasm2go/p2.Fn2248
func Fn2248(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2255 github.com/goccy/llamawasm2go/p2.Fn2255
func Fn2255(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2256 github.com/goccy/llamawasm2go/p2.Fn2256
func Fn2256(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2258 github.com/goccy/llamawasm2go/p2.Fn2258
func Fn2258(m *base.Module) int64

//go:linkname Fn2260 github.com/goccy/llamawasm2go/p2.Fn2260
func Fn2260(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2261 github.com/goccy/llamawasm2go/p2.Fn2261
func Fn2261(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2263 github.com/goccy/llamawasm2go/p2.Fn2263
func Fn2263(m *base.Module) int64

//go:linkname Fn2265 github.com/goccy/llamawasm2go/p2.Fn2265
func Fn2265(m *base.Module, l0 int32) int64

//go:linkname Fn2266 github.com/goccy/llamawasm2go/p2.Fn2266
func Fn2266(m *base.Module, l0 int32) int32

//go:linkname Fn2267 github.com/goccy/llamawasm2go/p2.Fn2267
func Fn2267(m *base.Module, l0 int32) int64

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 float32) int64

//go:linkname Fn2269 github.com/goccy/llamawasm2go/p2.Fn2269
func Fn2269(m *base.Module, l0 float32) int64

//go:linkname Fn2270 github.com/goccy/llamawasm2go/p2.Fn2270
func Fn2270(m *base.Module, l0 float32) int64

//go:linkname Fn2272 github.com/goccy/llamawasm2go/p2.Fn2272
func Fn2272(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2273 github.com/goccy/llamawasm2go/p2.Fn2273
func Fn2273(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2314 github.com/goccy/llamawasm2go/p2.Fn2314
func Fn2314(m *base.Module, l0 int64)

//go:linkname Fn2316 github.com/goccy/llamawasm2go/p2.Fn2316
func Fn2316(m *base.Module, l0 int64)

//go:linkname Fn2358 github.com/goccy/llamawasm2go/p2.Fn2358
func Fn2358(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2369 github.com/goccy/llamawasm2go/p2.Fn2369
func Fn2369(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2372 github.com/goccy/llamawasm2go/p2.Fn2372
func Fn2372(m *base.Module, l0 int64)

//go:linkname Fn2373 github.com/goccy/llamawasm2go/p2.Fn2373
func Fn2373(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2374 github.com/goccy/llamawasm2go/p2.Fn2374
func Fn2374(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2376 github.com/goccy/llamawasm2go/p2.Fn2376
func Fn2376(m *base.Module, l0 int64)

//go:linkname Fn2380 github.com/goccy/llamawasm2go/p2.Fn2380
func Fn2380(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2386 github.com/goccy/llamawasm2go/p2.Fn2386
func Fn2386(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2401 github.com/goccy/llamawasm2go/p2.Fn2401
func Fn2401(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2408 github.com/goccy/llamawasm2go/p0.Fn2408
func Fn2408(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2411 github.com/goccy/llamawasm2go/p2.Fn2411
func Fn2411(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2414 github.com/goccy/llamawasm2go/p2.Fn2414
func Fn2414(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2417 github.com/goccy/llamawasm2go/p2.Fn2417
func Fn2417(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2418 github.com/goccy/llamawasm2go/p2.Fn2418
func Fn2418(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2419 github.com/goccy/llamawasm2go/p2.Fn2419
func Fn2419(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2424 github.com/goccy/llamawasm2go/p2.Fn2424
func Fn2424(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2426 github.com/goccy/llamawasm2go/p2.Fn2426
func Fn2426(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2443 github.com/goccy/llamawasm2go/p2.Fn2443
func Fn2443(m *base.Module, l0 int64)

//go:linkname Fn2444 github.com/goccy/llamawasm2go/p2.Fn2444
func Fn2444(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2449 github.com/goccy/llamawasm2go/p2.Fn2449
func Fn2449(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2454 github.com/goccy/llamawasm2go/p2.Fn2454
func Fn2454(m *base.Module, l0 int64) int64

//go:linkname Fn2455 github.com/goccy/llamawasm2go/p2.Fn2455
func Fn2455(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2456 github.com/goccy/llamawasm2go/p0.Fn2456
func Fn2456(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2467 github.com/goccy/llamawasm2go/p2.Fn2467
func Fn2467(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2468 github.com/goccy/llamawasm2go/p2.Fn2468
func Fn2468(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2469 github.com/goccy/llamawasm2go/p2.Fn2469
func Fn2469(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2470 github.com/goccy/llamawasm2go/p2.Fn2470
func Fn2470(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2490 github.com/goccy/llamawasm2go/p2.Fn2490
func Fn2490(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2491 github.com/goccy/llamawasm2go/p2.Fn2491
func Fn2491(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2494 github.com/goccy/llamawasm2go/p2.Fn2494
func Fn2494(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2580 github.com/goccy/llamawasm2go/p2.Fn2580
func Fn2580(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2699 github.com/goccy/llamawasm2go/p2.Fn2699
func Fn2699(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2701 github.com/goccy/llamawasm2go/p2.Fn2701
func Fn2701(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2702 github.com/goccy/llamawasm2go/p2.Fn2702
func Fn2702(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2706 github.com/goccy/llamawasm2go/p2.Fn2706
func Fn2706(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2731 github.com/goccy/llamawasm2go/p2.Fn2731
func Fn2731(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2863 github.com/goccy/llamawasm2go/p2.Fn2863
func Fn2863(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2870 github.com/goccy/llamawasm2go/p2.Fn2870
func Fn2870(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2898 github.com/goccy/llamawasm2go/p2.Fn2898
func Fn2898(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2912 github.com/goccy/llamawasm2go/p2.Fn2912
func Fn2912(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2924 github.com/goccy/llamawasm2go/p2.Fn2924
func Fn2924(m *base.Module, l0 int32)

//go:linkname Fn2926 github.com/goccy/llamawasm2go/p2.Fn2926
func Fn2926(m *base.Module, l0 int64) int64

//go:linkname Fn2927 github.com/goccy/llamawasm2go/p2.Fn2927
func Fn2927(m *base.Module, l0 int64)

//go:linkname Fn2930 github.com/goccy/llamawasm2go/p2.Fn2930
func Fn2930(m *base.Module, l0 int64)

//go:linkname Fn2931 github.com/goccy/llamawasm2go/p2.Fn2931
func Fn2931(m *base.Module, l0 int64)

//go:linkname Fn2933 github.com/goccy/llamawasm2go/p2.Fn2933
func Fn2933(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2934 github.com/goccy/llamawasm2go/p2.Fn2934
func Fn2934(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2940 github.com/goccy/llamawasm2go/p2.Fn2940
func Fn2940(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2942 github.com/goccy/llamawasm2go/p2.Fn2942
func Fn2942(m *base.Module, l0 int64) int32

//go:linkname Fn2946 github.com/goccy/llamawasm2go/p2.Fn2946
func Fn2946(m *base.Module) int32

//go:linkname Fn2957 github.com/goccy/llamawasm2go/p2.Fn2957
func Fn2957(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2961 github.com/goccy/llamawasm2go/p2.Fn2961
func Fn2961(m *base.Module) int64

//go:linkname Fn2963 github.com/goccy/llamawasm2go/p2.Fn2963
func Fn2963(m *base.Module, l0 float64) float32

//go:linkname Fn2964 github.com/goccy/llamawasm2go/p2.Fn2964
func Fn2964(m *base.Module, l0 float64) float32

//go:linkname Fn2968 github.com/goccy/llamawasm2go/p2.Fn2968
func Fn2968(m *base.Module, l0 float64) float64

//go:linkname Fn2971 github.com/goccy/llamawasm2go/p2.Fn2971
func Fn2971(m *base.Module, l0 int32) float32

//go:linkname Fn2972 github.com/goccy/llamawasm2go/p2.Fn2972
func Fn2972(m *base.Module, l0 int32) float32

//go:linkname Fn2975 github.com/goccy/llamawasm2go/p2.Fn2975
func Fn2975(m *base.Module, l0 float32) float32

//go:linkname Fn2978 github.com/goccy/llamawasm2go/p2.Fn2978
func Fn2978(m *base.Module, l0 float64) float64

//go:linkname Fn2979 github.com/goccy/llamawasm2go/p2.Fn2979
func Fn2979(m *base.Module, l0 float64) float64

//go:linkname Fn2980 github.com/goccy/llamawasm2go/p2.Fn2980
func Fn2980(m *base.Module, l0 float32) float32

//go:linkname Fn2982 github.com/goccy/llamawasm2go/p2.Fn2982
func Fn2982(m *base.Module, l0 float32) float32

//go:linkname Fn2984 github.com/goccy/llamawasm2go/p2.Fn2984
func Fn2984(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn2985 github.com/goccy/llamawasm2go/p2.Fn2985
func Fn2985(m *base.Module, l0 float32) float32

//go:linkname Fn3002 github.com/goccy/llamawasm2go/p2.Fn3002
func Fn3002(m *base.Module, l0 int64) int32

//go:linkname Fn3003 github.com/goccy/llamawasm2go/p2.Fn3003
func Fn3003(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3005 github.com/goccy/llamawasm2go/p2.Fn3005
func Fn3005(m *base.Module, l0 int64)

//go:linkname Fn3006 github.com/goccy/llamawasm2go/p2.Fn3006
func Fn3006(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3007 github.com/goccy/llamawasm2go/p2.Fn3007
func Fn3007(m *base.Module, l0 int64) int32

//go:linkname Fn3014 github.com/goccy/llamawasm2go/p2.Fn3014
func Fn3014(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3016 github.com/goccy/llamawasm2go/p2.Fn3016
func Fn3016(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3022 github.com/goccy/llamawasm2go/p2.Fn3022
func Fn3022(m *base.Module, l0 int64) int32

//go:linkname Fn3025 github.com/goccy/llamawasm2go/p2.Fn3025
func Fn3025(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3028 github.com/goccy/llamawasm2go/p2.Fn3028
func Fn3028(m *base.Module, l0 int64) int32

//go:linkname Fn3030 github.com/goccy/llamawasm2go/p2.Fn3030
func Fn3030(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3032 github.com/goccy/llamawasm2go/p2.Fn3032
func Fn3032(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3033 github.com/goccy/llamawasm2go/p2.Fn3033
func Fn3033(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3036 github.com/goccy/llamawasm2go/p2.Fn3036
func Fn3036(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3039 github.com/goccy/llamawasm2go/p2.Fn3039
func Fn3039(m *base.Module, l0 int64) int64

//go:linkname Fn3043 github.com/goccy/llamawasm2go/p2.Fn3043
func Fn3043(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3053 github.com/goccy/llamawasm2go/p2.Fn3053
func Fn3053(m *base.Module)

//go:linkname Fn3054 github.com/goccy/llamawasm2go/p0.Fn3054
func Fn3054(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3057 github.com/goccy/llamawasm2go/p2.Fn3057
func Fn3057(m *base.Module)

//go:linkname Fn3059 github.com/goccy/llamawasm2go/p0.Fn3059
func Fn3059(m *base.Module, l0 int64) int64

//go:linkname Fn3061 github.com/goccy/llamawasm2go/p2.Fn3061
func Fn3061(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3065 github.com/goccy/llamawasm2go/p2.Fn3065
func Fn3065(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3134 github.com/goccy/llamawasm2go/p2.Fn3134
func Fn3134(m *base.Module, l0 int32)

//go:linkname Fn963rows github.com/goccy/llamawasm2go/p2.Fn963rows
func Fn963rows(m *base.Module)
