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

//go:linkname Fn335 github.com/goccy/llamawasm2go/p2.Fn335
func Fn335(m *base.Module, l0 int64) int32

//go:linkname Fn336 github.com/goccy/llamawasm2go/p2.Fn336
func Fn336(m *base.Module, l0 int64)

//go:linkname Fn337 github.com/goccy/llamawasm2go/p2.Fn337
func Fn337(m *base.Module, l0 int64)

//go:linkname Fn356 github.com/goccy/llamawasm2go/p2.Fn356
func Fn356(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn357 github.com/goccy/llamawasm2go/p2.Fn357
func Fn357(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn360 github.com/goccy/llamawasm2go/p2.Fn360
func Fn360(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn361 github.com/goccy/llamawasm2go/p2.Fn361
func Fn361(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn365 github.com/goccy/llamawasm2go/p2.Fn365
func Fn365(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn366 github.com/goccy/llamawasm2go/p2.Fn366
func Fn366(m *base.Module, l0 int64)

//go:linkname Fn367 github.com/goccy/llamawasm2go/p2.Fn367
func Fn367(m *base.Module, l0 int64)

//go:linkname Fn369 github.com/goccy/llamawasm2go/p2.Fn369
func Fn369(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn371 github.com/goccy/llamawasm2go/p2.Fn371
func Fn371(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn372 github.com/goccy/llamawasm2go/p2.Fn372
func Fn372(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn374 github.com/goccy/llamawasm2go/p2.Fn374
func Fn374(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn377 github.com/goccy/llamawasm2go/p0.Fn377
func Fn377(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int64)

//go:linkname Fn380 github.com/goccy/llamawasm2go/p2.Fn380
func Fn380(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn381 github.com/goccy/llamawasm2go/p2.Fn381
func Fn381(m *base.Module, l0 int64)

//go:linkname Fn384 github.com/goccy/llamawasm2go/p2.Fn384
func Fn384(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn385 github.com/goccy/llamawasm2go/p2.Fn385
func Fn385(m *base.Module, l0 int64)

//go:linkname Fn386 github.com/goccy/llamawasm2go/p2.Fn386
func Fn386(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn387 github.com/goccy/llamawasm2go/p2.Fn387
func Fn387(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn388 github.com/goccy/llamawasm2go/p2.Fn388
func Fn388(m *base.Module)

//go:linkname Fn389 github.com/goccy/llamawasm2go/p0.Fn389
func Fn389(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int32)

//go:linkname Fn393 github.com/goccy/llamawasm2go/p0.Fn393
func Fn393(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 int64)

//go:linkname Fn401 github.com/goccy/llamawasm2go/p2.Fn401
func Fn401(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32)

//go:linkname Fn407 github.com/goccy/llamawasm2go/p2.Fn407
func Fn407(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn409 github.com/goccy/llamawasm2go/p2.Fn409
func Fn409(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn410 github.com/goccy/llamawasm2go/p2.Fn410
func Fn410(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn431 github.com/goccy/llamawasm2go/p2.Fn431
func Fn431(m *base.Module, l0 int64) int32

//go:linkname Fn432 github.com/goccy/llamawasm2go/p2.Fn432
func Fn432(m *base.Module, l0 int64) int32

//go:linkname Fn443 github.com/goccy/llamawasm2go/p2.Fn443
func Fn443(m *base.Module, l0 int64) int64

//go:linkname Fn448 github.com/goccy/llamawasm2go/p2.Fn448
func Fn448(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn450 github.com/goccy/llamawasm2go/p2.Fn450
func Fn450(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn451 github.com/goccy/llamawasm2go/p2.Fn451
func Fn451(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn452 github.com/goccy/llamawasm2go/p2.Fn452
func Fn452(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn453 github.com/goccy/llamawasm2go/p2.Fn453
func Fn453(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn454 github.com/goccy/llamawasm2go/p2.Fn454
func Fn454(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn455 github.com/goccy/llamawasm2go/p2.Fn455
func Fn455(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn458 github.com/goccy/llamawasm2go/p2.Fn458
func Fn458(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn459 github.com/goccy/llamawasm2go/p2.Fn459
func Fn459(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn462 github.com/goccy/llamawasm2go/p2.Fn462
func Fn462(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn463 github.com/goccy/llamawasm2go/p2.Fn463
func Fn463(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn464 github.com/goccy/llamawasm2go/p2.Fn464
func Fn464(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn465 github.com/goccy/llamawasm2go/p2.Fn465
func Fn465(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn466 github.com/goccy/llamawasm2go/p2.Fn466
func Fn466(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn467 github.com/goccy/llamawasm2go/p2.Fn467
func Fn467(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn469 github.com/goccy/llamawasm2go/p2.Fn469
func Fn469(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn471 github.com/goccy/llamawasm2go/p2.Fn471
func Fn471(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn473 github.com/goccy/llamawasm2go/p2.Fn473
func Fn473(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn474 github.com/goccy/llamawasm2go/p2.Fn474
func Fn474(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn475 github.com/goccy/llamawasm2go/p2.Fn475
func Fn475(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn476 github.com/goccy/llamawasm2go/p2.Fn476
func Fn476(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn477 github.com/goccy/llamawasm2go/p2.Fn477
func Fn477(m *base.Module, l0 int64, l1 int64) int64

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

//go:linkname Fn484 github.com/goccy/llamawasm2go/p2.Fn484
func Fn484(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn486 github.com/goccy/llamawasm2go/p2.Fn486
func Fn486(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn487 github.com/goccy/llamawasm2go/p2.Fn487
func Fn487(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn488 github.com/goccy/llamawasm2go/p2.Fn488
func Fn488(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn489 github.com/goccy/llamawasm2go/p2.Fn489
func Fn489(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn490 github.com/goccy/llamawasm2go/p2.Fn490
func Fn490(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn491 github.com/goccy/llamawasm2go/p2.Fn491
func Fn491(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn492 github.com/goccy/llamawasm2go/p2.Fn492
func Fn492(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn493 github.com/goccy/llamawasm2go/p2.Fn493
func Fn493(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn494 github.com/goccy/llamawasm2go/p2.Fn494
func Fn494(m *base.Module, l0 int64)

//go:linkname Fn495 github.com/goccy/llamawasm2go/p2.Fn495
func Fn495(m *base.Module, l0 int64)

//go:linkname Fn496 github.com/goccy/llamawasm2go/p2.Fn496
func Fn496(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn497 github.com/goccy/llamawasm2go/p2.Fn497
func Fn497(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn498 github.com/goccy/llamawasm2go/p2.Fn498
func Fn498(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn499 github.com/goccy/llamawasm2go/p2.Fn499
func Fn499(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn500 github.com/goccy/llamawasm2go/p2.Fn500
func Fn500(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn501 github.com/goccy/llamawasm2go/p2.Fn501
func Fn501(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn502 github.com/goccy/llamawasm2go/p2.Fn502
func Fn502(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn503 github.com/goccy/llamawasm2go/p2.Fn503
func Fn503(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn504 github.com/goccy/llamawasm2go/p2.Fn504
func Fn504(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn505 github.com/goccy/llamawasm2go/p2.Fn505
func Fn505(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn507 github.com/goccy/llamawasm2go/p2.Fn507
func Fn507(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn508 github.com/goccy/llamawasm2go/p2.Fn508
func Fn508(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn509 github.com/goccy/llamawasm2go/p2.Fn509
func Fn509(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn510 github.com/goccy/llamawasm2go/p2.Fn510
func Fn510(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn511 github.com/goccy/llamawasm2go/p2.Fn511
func Fn511(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn512 github.com/goccy/llamawasm2go/p2.Fn512
func Fn512(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn513 github.com/goccy/llamawasm2go/p2.Fn513
func Fn513(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64) int64

//go:linkname Fn514 github.com/goccy/llamawasm2go/p2.Fn514
func Fn514(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn515 github.com/goccy/llamawasm2go/p2.Fn515
func Fn515(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn516 github.com/goccy/llamawasm2go/p2.Fn516
func Fn516(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn517 github.com/goccy/llamawasm2go/p2.Fn517
func Fn517(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn518 github.com/goccy/llamawasm2go/p2.Fn518
func Fn518(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn520 github.com/goccy/llamawasm2go/p2.Fn520
func Fn520(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn523 github.com/goccy/llamawasm2go/p2.Fn523
func Fn523(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32) int64

//go:linkname Fn524 github.com/goccy/llamawasm2go/p2.Fn524
func Fn524(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32, l7 float32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32) int64

//go:linkname Fn526 github.com/goccy/llamawasm2go/p2.Fn526
func Fn526(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn527 github.com/goccy/llamawasm2go/p2.Fn527
func Fn527(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn528 github.com/goccy/llamawasm2go/p2.Fn528
func Fn528(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn530 github.com/goccy/llamawasm2go/p2.Fn530
func Fn530(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn532 github.com/goccy/llamawasm2go/p2.Fn532
func Fn532(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn533 github.com/goccy/llamawasm2go/p2.Fn533
func Fn533(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn534 github.com/goccy/llamawasm2go/p2.Fn534
func Fn534(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn536 github.com/goccy/llamawasm2go/p2.Fn536
func Fn536(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 float32, l6 float32, l7 float32) int64

//go:linkname Fn537 github.com/goccy/llamawasm2go/p2.Fn537
func Fn537(m *base.Module, l0 int64)

//go:linkname Fn538 github.com/goccy/llamawasm2go/p2.Fn538
func Fn538(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn542 github.com/goccy/llamawasm2go/p2.Fn542
func Fn542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn543 github.com/goccy/llamawasm2go/p2.Fn543
func Fn543(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn546 github.com/goccy/llamawasm2go/p2.Fn546
func Fn546(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn548 github.com/goccy/llamawasm2go/p2.Fn548
func Fn548(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn550 github.com/goccy/llamawasm2go/p2.Fn550
func Fn550(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn556 github.com/goccy/llamawasm2go/p2.Fn556
func Fn556(m *base.Module, l0 int64)

//go:linkname Fn559 github.com/goccy/llamawasm2go/p2.Fn559
func Fn559(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn560 github.com/goccy/llamawasm2go/p2.Fn560
func Fn560(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn562 github.com/goccy/llamawasm2go/p2.Fn562
func Fn562(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn565 github.com/goccy/llamawasm2go/p2.Fn565
func Fn565(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn566 github.com/goccy/llamawasm2go/p2.Fn566
func Fn566(m *base.Module, l0 int64) int64

//go:linkname Fn567 github.com/goccy/llamawasm2go/p2.Fn567
func Fn567(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn568 github.com/goccy/llamawasm2go/p2.Fn568
func Fn568(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn569 github.com/goccy/llamawasm2go/p2.Fn569
func Fn569(m *base.Module, l0 int64) int64

//go:linkname Fn570 github.com/goccy/llamawasm2go/p2.Fn570
func Fn570(m *base.Module, l0 int64) int64

//go:linkname Fn571 github.com/goccy/llamawasm2go/p2.Fn571
func Fn571(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn573 github.com/goccy/llamawasm2go/p2.Fn573
func Fn573(m *base.Module, l0 int64) int64

//go:linkname Fn574 github.com/goccy/llamawasm2go/p2.Fn574
func Fn574(m *base.Module, l0 int64) int64

//go:linkname Fn576 github.com/goccy/llamawasm2go/p2.Fn576
func Fn576(m *base.Module, l0 int64)

//go:linkname Fn577 github.com/goccy/llamawasm2go/p2.Fn577
func Fn577(m *base.Module, l0 int64) int64

//go:linkname Fn578 github.com/goccy/llamawasm2go/p2.Fn578
func Fn578(m *base.Module, l0 int64) int64

//go:linkname Fn579 github.com/goccy/llamawasm2go/p2.Fn579
func Fn579(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn580 github.com/goccy/llamawasm2go/p2.Fn580
func Fn580(m *base.Module, l0 int64) int32

//go:linkname Fn581 github.com/goccy/llamawasm2go/p2.Fn581
func Fn581(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn583 github.com/goccy/llamawasm2go/p2.Fn583
func Fn583(m *base.Module, l0 int64)

//go:linkname Fn584 github.com/goccy/llamawasm2go/p2.Fn584
func Fn584(m *base.Module, l0 int64)

//go:linkname Fn585 github.com/goccy/llamawasm2go/p2.Fn585
func Fn585(m *base.Module, l0 int64) int64

//go:linkname Fn587 github.com/goccy/llamawasm2go/p2.Fn587
func Fn587(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn588 github.com/goccy/llamawasm2go/p2.Fn588
func Fn588(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn589 github.com/goccy/llamawasm2go/p2.Fn589
func Fn589(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn590 github.com/goccy/llamawasm2go/p2.Fn590
func Fn590(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn591 github.com/goccy/llamawasm2go/p2.Fn591
func Fn591(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn592 github.com/goccy/llamawasm2go/p2.Fn592
func Fn592(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn595 github.com/goccy/llamawasm2go/p2.Fn595
func Fn595(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn596 github.com/goccy/llamawasm2go/p2.Fn596
func Fn596(m *base.Module, l0 int64) int64

//go:linkname Fn597 github.com/goccy/llamawasm2go/p2.Fn597
func Fn597(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn600 github.com/goccy/llamawasm2go/p2.Fn600
func Fn600(m *base.Module, l0 int64) int64

//go:linkname Fn601 github.com/goccy/llamawasm2go/p2.Fn601
func Fn601(m *base.Module, l0 int64) int64

//go:linkname Fn605 github.com/goccy/llamawasm2go/p2.Fn605
func Fn605(m *base.Module, l0 int64) int64

//go:linkname Fn606 github.com/goccy/llamawasm2go/p2.Fn606
func Fn606(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn607 github.com/goccy/llamawasm2go/p2.Fn607
func Fn607(m *base.Module, l0 int64) int64

//go:linkname Fn611 github.com/goccy/llamawasm2go/p2.Fn611
func Fn611(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn612 github.com/goccy/llamawasm2go/p0.Fn612
func Fn612(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn617 github.com/goccy/llamawasm2go/p2.Fn617
func Fn617(m *base.Module, l0 int64)

//go:linkname Fn618 github.com/goccy/llamawasm2go/p2.Fn618
func Fn618(m *base.Module, l0 int64)

//go:linkname Fn619 github.com/goccy/llamawasm2go/p2.Fn619
func Fn619(m *base.Module, l0 int64)

//go:linkname Fn620 github.com/goccy/llamawasm2go/p2.Fn620
func Fn620(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn621 github.com/goccy/llamawasm2go/p2.Fn621
func Fn621(m *base.Module, l0 int64) int32

//go:linkname Fn622 github.com/goccy/llamawasm2go/p2.Fn622
func Fn622(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn623 github.com/goccy/llamawasm2go/p2.Fn623
func Fn623(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn624 github.com/goccy/llamawasm2go/p2.Fn624
func Fn624(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn625 github.com/goccy/llamawasm2go/p2.Fn625
func Fn625(m *base.Module, l0 int64) int32

//go:linkname Fn626 github.com/goccy/llamawasm2go/p2.Fn626
func Fn626(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn644 github.com/goccy/llamawasm2go/p2.Fn644
func Fn644(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn650 github.com/goccy/llamawasm2go/p2.Fn650
func Fn650(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn651 github.com/goccy/llamawasm2go/p2.Fn651
func Fn651(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn652 github.com/goccy/llamawasm2go/p2.Fn652
func Fn652(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn654 github.com/goccy/llamawasm2go/p2.Fn654
func Fn654(m *base.Module, l0 int64)

//go:linkname Fn656 github.com/goccy/llamawasm2go/p2.Fn656
func Fn656(m *base.Module, l0 int64)

//go:linkname Fn657 github.com/goccy/llamawasm2go/p0.Fn657
func Fn657(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn659 github.com/goccy/llamawasm2go/p2.Fn659
func Fn659(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn673 github.com/goccy/llamawasm2go/p2.Fn673
func Fn673(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn675 github.com/goccy/llamawasm2go/p2.Fn675
func Fn675(m *base.Module, l0 int64)

//go:linkname Fn676 github.com/goccy/llamawasm2go/p2.Fn676
func Fn676(m *base.Module, l0 int64) int64

//go:linkname Fn699 github.com/goccy/llamawasm2go/p2.Fn699
func Fn699(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn710 github.com/goccy/llamawasm2go/p2.Fn710
func Fn710(m *base.Module, l0 int64)

//go:linkname Fn729 github.com/goccy/llamawasm2go/p2.Fn729
func Fn729(m *base.Module, l0 int64)

//go:linkname Fn730 github.com/goccy/llamawasm2go/p2.Fn730
func Fn730(m *base.Module, l0 int64)

//go:linkname Fn798 github.com/goccy/llamawasm2go/p2.Fn798
func Fn798(m *base.Module, l0 int64)

//go:linkname Fn800 github.com/goccy/llamawasm2go/p2.Fn800
func Fn800(m *base.Module, l0 int64) int64

//go:linkname Fn848 github.com/goccy/llamawasm2go/p0.Fn848
func Fn848(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn861 github.com/goccy/llamawasm2go/p0.Fn861
func Fn861(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn908 github.com/goccy/llamawasm2go/p2.Fn908
func Fn908(m *base.Module) int64

//go:linkname Fn945 github.com/goccy/llamawasm2go/p2.Fn945
func Fn945(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32)

//go:linkname Fn955 github.com/goccy/llamawasm2go/p2.Fn955
func Fn955(m *base.Module, l0 int64)

//go:linkname Fn980 github.com/goccy/llamawasm2go/p2.Fn980
func Fn980(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn985 github.com/goccy/llamawasm2go/p2.Fn985
func Fn985(m *base.Module, l0 int64) int64

//go:linkname Fn995 github.com/goccy/llamawasm2go/p2.Fn995
func Fn995(m *base.Module, l0 int64)

//go:linkname Fn1000 github.com/goccy/llamawasm2go/p2.Fn1000
func Fn1000(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1001 github.com/goccy/llamawasm2go/p2.Fn1001
func Fn1001(m *base.Module)

//go:linkname Fn1010 github.com/goccy/llamawasm2go/p2.Fn1010
func Fn1010(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1018 github.com/goccy/llamawasm2go/p2.Fn1018
func Fn1018(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1019 github.com/goccy/llamawasm2go/p2.Fn1019
func Fn1019(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1020 github.com/goccy/llamawasm2go/p2.Fn1020
func Fn1020(m *base.Module, l0 int64, l1 int64, l2 int64) float32

//go:linkname Fn1022 github.com/goccy/llamawasm2go/p2.Fn1022
func Fn1022(m *base.Module, l0 int64, l1 int64, l2 int64) float64

//go:linkname Fn1024 github.com/goccy/llamawasm2go/p2.Fn1024
func Fn1024(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1027 github.com/goccy/llamawasm2go/p2.Fn1027
func Fn1027(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1030 github.com/goccy/llamawasm2go/p2.Fn1030
func Fn1030(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1036 github.com/goccy/llamawasm2go/p2.Fn1036
func Fn1036(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1037 github.com/goccy/llamawasm2go/p2.Fn1037
func Fn1037(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn1045 github.com/goccy/llamawasm2go/p2.Fn1045
func Fn1045(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1050 github.com/goccy/llamawasm2go/p2.Fn1050
func Fn1050(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1051 github.com/goccy/llamawasm2go/p2.Fn1051
func Fn1051(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1059 github.com/goccy/llamawasm2go/p2.Fn1059
func Fn1059(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn1075 github.com/goccy/llamawasm2go/p2.Fn1075
func Fn1075(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1083 github.com/goccy/llamawasm2go/p2.Fn1083
func Fn1083(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1099 github.com/goccy/llamawasm2go/p0.Fn1099
func Fn1099(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn1102 github.com/goccy/llamawasm2go/p2.Fn1102
func Fn1102(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1106 github.com/goccy/llamawasm2go/p2.Fn1106
func Fn1106(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn1117 github.com/goccy/llamawasm2go/p2.Fn1117
func Fn1117(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1118 github.com/goccy/llamawasm2go/p2.Fn1118
func Fn1118(m *base.Module, l0 int64)

//go:linkname Fn1210 github.com/goccy/llamawasm2go/p2.Fn1210
func Fn1210(m *base.Module, l0 int64)

//go:linkname Fn1232 github.com/goccy/llamawasm2go/p2.Fn1232
func Fn1232(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1242 github.com/goccy/llamawasm2go/p2.Fn1242
func Fn1242(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1248 github.com/goccy/llamawasm2go/p2.Fn1248
func Fn1248(m *base.Module)

//go:linkname Fn1252 github.com/goccy/llamawasm2go/p2.Fn1252
func Fn1252(m *base.Module, l0 int64) int64

//go:linkname Fn1280 github.com/goccy/llamawasm2go/p2.Fn1280
func Fn1280(m *base.Module, l0 int64)

//go:linkname Fn1284 github.com/goccy/llamawasm2go/p2.Fn1284
func Fn1284(m *base.Module, l0 int32) int64

//go:linkname Fn1295 github.com/goccy/llamawasm2go/p2.Fn1295
func Fn1295(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1296 github.com/goccy/llamawasm2go/p2.Fn1296
func Fn1296(m *base.Module, l0 int64)

//go:linkname Fn1298 github.com/goccy/llamawasm2go/p2.Fn1298
func Fn1298(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1320 github.com/goccy/llamawasm2go/p2.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1321 github.com/goccy/llamawasm2go/p2.Fn1321
func Fn1321(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1322 github.com/goccy/llamawasm2go/p2.Fn1322
func Fn1322(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p2.Fn1324
func Fn1324(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1325 github.com/goccy/llamawasm2go/p2.Fn1325
func Fn1325(m *base.Module, l0 int64) int64

//go:linkname Fn1328 github.com/goccy/llamawasm2go/p2.Fn1328
func Fn1328(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1332 github.com/goccy/llamawasm2go/p2.Fn1332
func Fn1332(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1333 github.com/goccy/llamawasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1335 github.com/goccy/llamawasm2go/p2.Fn1335
func Fn1335(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1336 github.com/goccy/llamawasm2go/p2.Fn1336
func Fn1336(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1337 github.com/goccy/llamawasm2go/p2.Fn1337
func Fn1337(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1338 github.com/goccy/llamawasm2go/p2.Fn1338
func Fn1338(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1339 github.com/goccy/llamawasm2go/p2.Fn1339
func Fn1339(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1340 github.com/goccy/llamawasm2go/p2.Fn1340
func Fn1340(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1341 github.com/goccy/llamawasm2go/p2.Fn1341
func Fn1341(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1342 github.com/goccy/llamawasm2go/p2.Fn1342
func Fn1342(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1343 github.com/goccy/llamawasm2go/p2.Fn1343
func Fn1343(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1344 github.com/goccy/llamawasm2go/p2.Fn1344
func Fn1344(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1345 github.com/goccy/llamawasm2go/p2.Fn1345
func Fn1345(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1346 github.com/goccy/llamawasm2go/p2.Fn1346
func Fn1346(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1347 github.com/goccy/llamawasm2go/p2.Fn1347
func Fn1347(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1348 github.com/goccy/llamawasm2go/p2.Fn1348
func Fn1348(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1349 github.com/goccy/llamawasm2go/p2.Fn1349
func Fn1349(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1350 github.com/goccy/llamawasm2go/p2.Fn1350
func Fn1350(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1351 github.com/goccy/llamawasm2go/p2.Fn1351
func Fn1351(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1352 github.com/goccy/llamawasm2go/p2.Fn1352
func Fn1352(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1353 github.com/goccy/llamawasm2go/p2.Fn1353
func Fn1353(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1354 github.com/goccy/llamawasm2go/p2.Fn1354
func Fn1354(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1355 github.com/goccy/llamawasm2go/p2.Fn1355
func Fn1355(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1356 github.com/goccy/llamawasm2go/p2.Fn1356
func Fn1356(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1357 github.com/goccy/llamawasm2go/p2.Fn1357
func Fn1357(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1358 github.com/goccy/llamawasm2go/p2.Fn1358
func Fn1358(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1360 github.com/goccy/llamawasm2go/p2.Fn1360
func Fn1360(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1364 github.com/goccy/llamawasm2go/p2.Fn1364
func Fn1364(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1365 github.com/goccy/llamawasm2go/p2.Fn1365
func Fn1365(m *base.Module)

//go:linkname Fn1366 github.com/goccy/llamawasm2go/p2.Fn1366
func Fn1366(m *base.Module)

//go:linkname Fn1367 github.com/goccy/llamawasm2go/p0.Fn1367
func Fn1367(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1375 github.com/goccy/llamawasm2go/p2.Fn1375
func Fn1375(m *base.Module)

//go:linkname Fn1377 github.com/goccy/llamawasm2go/p2.Fn1377
func Fn1377(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1379 github.com/goccy/llamawasm2go/p0.Fn1379
func Fn1379(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1380 github.com/goccy/llamawasm2go/p2.Fn1380
func Fn1380(m *base.Module, l0 int64) int64

//go:linkname Fn1385 github.com/goccy/llamawasm2go/p2.Fn1385
func Fn1385(m *base.Module, l0 int64)

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p2.Fn1392
func Fn1392(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p2.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p2.Fn1397
func Fn1397(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1399 github.com/goccy/llamawasm2go/p2.Fn1399
func Fn1399(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p2.Fn1401
func Fn1401(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1411 github.com/goccy/llamawasm2go/p2.Fn1411
func Fn1411(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1412 github.com/goccy/llamawasm2go/p2.Fn1412
func Fn1412(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p2.Fn1413
func Fn1413(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1414 github.com/goccy/llamawasm2go/p0.Fn1414
func Fn1414(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1416 github.com/goccy/llamawasm2go/p2.Fn1416
func Fn1416(m *base.Module, l0 int64)

//go:linkname Fn1417 github.com/goccy/llamawasm2go/p2.Fn1417
func Fn1417(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1419 github.com/goccy/llamawasm2go/p2.Fn1419
func Fn1419(m *base.Module, l0 int64) int64

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1421 github.com/goccy/llamawasm2go/p2.Fn1421
func Fn1421(m *base.Module, l0 int64)

//go:linkname Fn1422 github.com/goccy/llamawasm2go/p2.Fn1422
func Fn1422(m *base.Module, l0 int64)

//go:linkname Fn1423 github.com/goccy/llamawasm2go/p2.Fn1423
func Fn1423(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1427 github.com/goccy/llamawasm2go/p2.Fn1427
func Fn1427(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1435 github.com/goccy/llamawasm2go/p2.Fn1435
func Fn1435(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1437 github.com/goccy/llamawasm2go/p0.Fn1437
func Fn1437(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1439 github.com/goccy/llamawasm2go/p2.Fn1439
func Fn1439(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p0.Fn1442
func Fn1442(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int64)

//go:linkname Fn1450 github.com/goccy/llamawasm2go/p2.Fn1450
func Fn1450(m *base.Module, l0 int64)

//go:linkname Fn1453 github.com/goccy/llamawasm2go/p2.Fn1453
func Fn1453(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1454 github.com/goccy/llamawasm2go/p2.Fn1454
func Fn1454(m *base.Module, l0 int64) int64

//go:linkname Fn1455 github.com/goccy/llamawasm2go/p2.Fn1455
func Fn1455(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1456 github.com/goccy/llamawasm2go/p2.Fn1456
func Fn1456(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1457 github.com/goccy/llamawasm2go/p2.Fn1457
func Fn1457(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1458 github.com/goccy/llamawasm2go/p2.Fn1458
func Fn1458(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1459 github.com/goccy/llamawasm2go/p2.Fn1459
func Fn1459(m *base.Module, l0 int64)

//go:linkname Fn1462 github.com/goccy/llamawasm2go/p2.Fn1462
func Fn1462(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1464 github.com/goccy/llamawasm2go/p2.Fn1464
func Fn1464(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1467 github.com/goccy/llamawasm2go/p2.Fn1467
func Fn1467(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1474 github.com/goccy/llamawasm2go/p2.Fn1474
func Fn1474(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1475 github.com/goccy/llamawasm2go/p2.Fn1475
func Fn1475(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1477 github.com/goccy/llamawasm2go/p2.Fn1477
func Fn1477(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1481 github.com/goccy/llamawasm2go/p2.Fn1481
func Fn1481(m *base.Module, l0 int64)

//go:linkname Fn1482 github.com/goccy/llamawasm2go/p2.Fn1482
func Fn1482(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1484 github.com/goccy/llamawasm2go/p2.Fn1484
func Fn1484(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1485 github.com/goccy/llamawasm2go/p2.Fn1485
func Fn1485(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1486 github.com/goccy/llamawasm2go/p2.Fn1486
func Fn1486(m *base.Module, l0 int64)

//go:linkname Fn1488 github.com/goccy/llamawasm2go/p2.Fn1488
func Fn1488(m *base.Module, l0 int64) int32

//go:linkname Fn1489 github.com/goccy/llamawasm2go/p2.Fn1489
func Fn1489(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1531 github.com/goccy/llamawasm2go/p2.Fn1531
func Fn1531(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1547 github.com/goccy/llamawasm2go/p2.Fn1547
func Fn1547(m *base.Module, l0 int64)

//go:linkname Fn1548 github.com/goccy/llamawasm2go/p2.Fn1548
func Fn1548(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1549 github.com/goccy/llamawasm2go/p2.Fn1549
func Fn1549(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1550 github.com/goccy/llamawasm2go/p2.Fn1550
func Fn1550(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1551 github.com/goccy/llamawasm2go/p2.Fn1551
func Fn1551(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1552 github.com/goccy/llamawasm2go/p2.Fn1552
func Fn1552(m *base.Module, l0 int64)

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1555 github.com/goccy/llamawasm2go/p2.Fn1555
func Fn1555(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1556 github.com/goccy/llamawasm2go/p2.Fn1556
func Fn1556(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1557 github.com/goccy/llamawasm2go/p2.Fn1557
func Fn1557(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1558 github.com/goccy/llamawasm2go/p2.Fn1558
func Fn1558(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1559 github.com/goccy/llamawasm2go/p2.Fn1559
func Fn1559(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1562 github.com/goccy/llamawasm2go/p2.Fn1562
func Fn1562(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1564 github.com/goccy/llamawasm2go/p2.Fn1564
func Fn1564(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1565 github.com/goccy/llamawasm2go/p2.Fn1565
func Fn1565(m *base.Module, l0 int64) int64

//go:linkname Fn1566 github.com/goccy/llamawasm2go/p2.Fn1566
func Fn1566(m *base.Module, l0 int64) int64

//go:linkname Fn1567 github.com/goccy/llamawasm2go/p2.Fn1567
func Fn1567(m *base.Module, l0 int64) int64

//go:linkname Fn1568 github.com/goccy/llamawasm2go/p2.Fn1568
func Fn1568(m *base.Module, l0 int64) int64

//go:linkname Fn1569 github.com/goccy/llamawasm2go/p2.Fn1569
func Fn1569(m *base.Module, l0 int64) int64

//go:linkname Fn1570 github.com/goccy/llamawasm2go/p2.Fn1570
func Fn1570(m *base.Module, l0 int64) int64

//go:linkname Fn1572 github.com/goccy/llamawasm2go/p2.Fn1572
func Fn1572(m *base.Module, l0 int64) int64

//go:linkname Fn1573 github.com/goccy/llamawasm2go/p2.Fn1573
func Fn1573(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1575 github.com/goccy/llamawasm2go/p2.Fn1575
func Fn1575(m *base.Module, l0 int64) int64

//go:linkname Fn1576 github.com/goccy/llamawasm2go/p2.Fn1576
func Fn1576(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1577 github.com/goccy/llamawasm2go/p2.Fn1577
func Fn1577(m *base.Module, l0 int64) int64

//go:linkname Fn1579 github.com/goccy/llamawasm2go/p2.Fn1579
func Fn1579(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1580 github.com/goccy/llamawasm2go/p2.Fn1580
func Fn1580(m *base.Module, l0 int64) int64

//go:linkname Fn1582 github.com/goccy/llamawasm2go/p2.Fn1582
func Fn1582(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p2.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1584 github.com/goccy/llamawasm2go/p2.Fn1584
func Fn1584(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1585 github.com/goccy/llamawasm2go/p2.Fn1585
func Fn1585(m *base.Module, l0 int64) int64

//go:linkname Fn1587 github.com/goccy/llamawasm2go/p2.Fn1587
func Fn1587(m *base.Module, l0 int64) int64

//go:linkname Fn1593 github.com/goccy/llamawasm2go/p2.Fn1593
func Fn1593(m *base.Module, l0 int64) int64

//go:linkname Fn1595 github.com/goccy/llamawasm2go/p2.Fn1595
func Fn1595(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1596 github.com/goccy/llamawasm2go/p2.Fn1596
func Fn1596(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1598 github.com/goccy/llamawasm2go/p2.Fn1598
func Fn1598(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1599 github.com/goccy/llamawasm2go/p2.Fn1599
func Fn1599(m *base.Module, l0 int64) int64

//go:linkname Fn1600 github.com/goccy/llamawasm2go/p2.Fn1600
func Fn1600(m *base.Module, l0 int64) int64

//go:linkname Fn1601 github.com/goccy/llamawasm2go/p2.Fn1601
func Fn1601(m *base.Module, l0 int64) int64

//go:linkname Fn1602 github.com/goccy/llamawasm2go/p2.Fn1602
func Fn1602(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1603 github.com/goccy/llamawasm2go/p2.Fn1603
func Fn1603(m *base.Module, l0 int64)

//go:linkname Fn1632 github.com/goccy/llamawasm2go/p2.Fn1632
func Fn1632(m *base.Module, l0 int64) int64

//go:linkname Fn1639 github.com/goccy/llamawasm2go/p2.Fn1639
func Fn1639(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1640 github.com/goccy/llamawasm2go/p2.Fn1640
func Fn1640(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1641 github.com/goccy/llamawasm2go/p2.Fn1641
func Fn1641(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1643 github.com/goccy/llamawasm2go/p2.Fn1643
func Fn1643(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1644 github.com/goccy/llamawasm2go/p2.Fn1644
func Fn1644(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1648 github.com/goccy/llamawasm2go/p2.Fn1648
func Fn1648(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1649 github.com/goccy/llamawasm2go/p2.Fn1649
func Fn1649(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1650 github.com/goccy/llamawasm2go/p2.Fn1650
func Fn1650(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1651 github.com/goccy/llamawasm2go/p2.Fn1651
func Fn1651(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1652 github.com/goccy/llamawasm2go/p2.Fn1652
func Fn1652(m *base.Module, l0 int64) int32

//go:linkname Fn1653 github.com/goccy/llamawasm2go/p2.Fn1653
func Fn1653(m *base.Module, l0 int64) int32

//go:linkname Fn1654 github.com/goccy/llamawasm2go/p2.Fn1654
func Fn1654(m *base.Module, l0 int64) int32

//go:linkname Fn1655 github.com/goccy/llamawasm2go/p2.Fn1655
func Fn1655(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1658 github.com/goccy/llamawasm2go/p2.Fn1658
func Fn1658(m *base.Module, l0 int64) int32

//go:linkname Fn1659 github.com/goccy/llamawasm2go/p2.Fn1659
func Fn1659(m *base.Module, l0 int64) int32

//go:linkname Fn1665 github.com/goccy/llamawasm2go/p2.Fn1665
func Fn1665(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1666 github.com/goccy/llamawasm2go/p2.Fn1666
func Fn1666(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1667 github.com/goccy/llamawasm2go/p2.Fn1667
func Fn1667(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1668 github.com/goccy/llamawasm2go/p2.Fn1668
func Fn1668(m *base.Module)

//go:linkname Fn1669 github.com/goccy/llamawasm2go/p2.Fn1669
func Fn1669(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1671 github.com/goccy/llamawasm2go/p2.Fn1671
func Fn1671(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1673 github.com/goccy/llamawasm2go/p2.Fn1673
func Fn1673(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1674 github.com/goccy/llamawasm2go/p2.Fn1674
func Fn1674(m *base.Module, l0 int64)

//go:linkname Fn1678 github.com/goccy/llamawasm2go/p2.Fn1678
func Fn1678(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1679 github.com/goccy/llamawasm2go/p2.Fn1679
func Fn1679(m *base.Module, l0 int64)

//go:linkname Fn1682 github.com/goccy/llamawasm2go/p2.Fn1682
func Fn1682(m *base.Module, l0 int64)

//go:linkname Fn1692 github.com/goccy/llamawasm2go/p2.Fn1692
func Fn1692(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1693 github.com/goccy/llamawasm2go/p2.Fn1693
func Fn1693(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1702 github.com/goccy/llamawasm2go/p2.Fn1702
func Fn1702(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1703 github.com/goccy/llamawasm2go/p0.Fn1703
func Fn1703(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1705 github.com/goccy/llamawasm2go/p2.Fn1705
func Fn1705(m *base.Module, l0 int64)

//go:linkname Fn1707 github.com/goccy/llamawasm2go/p2.Fn1707
func Fn1707(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1714 github.com/goccy/llamawasm2go/p2.Fn1714
func Fn1714(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1715 github.com/goccy/llamawasm2go/p2.Fn1715
func Fn1715(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1722 github.com/goccy/llamawasm2go/p2.Fn1722
func Fn1722(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1729 github.com/goccy/llamawasm2go/p2.Fn1729
func Fn1729(m *base.Module, l0 int64)

//go:linkname Fn1732 github.com/goccy/llamawasm2go/p2.Fn1732
func Fn1732(m *base.Module, l0 int64) int32

//go:linkname Fn1742 github.com/goccy/llamawasm2go/p2.Fn1742
func Fn1742(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1744 github.com/goccy/llamawasm2go/p2.Fn1744
func Fn1744(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1745 github.com/goccy/llamawasm2go/p2.Fn1745
func Fn1745(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1746 github.com/goccy/llamawasm2go/p2.Fn1746
func Fn1746(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1748 github.com/goccy/llamawasm2go/p2.Fn1748
func Fn1748(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1764 github.com/goccy/llamawasm2go/p2.Fn1764
func Fn1764(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1777 github.com/goccy/llamawasm2go/p2.Fn1777
func Fn1777(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1784 github.com/goccy/llamawasm2go/p2.Fn1784
func Fn1784(m *base.Module, l0 int64)

//go:linkname Fn1818 github.com/goccy/llamawasm2go/p2.Fn1818
func Fn1818(m *base.Module, l0 int64)

//go:linkname Fn1821 github.com/goccy/llamawasm2go/p2.Fn1821
func Fn1821(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1835 github.com/goccy/llamawasm2go/p2.Fn1835
func Fn1835(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1836 github.com/goccy/llamawasm2go/p2.Fn1836
func Fn1836(m *base.Module, l0 int64) int64

//go:linkname Fn1837 github.com/goccy/llamawasm2go/p2.Fn1837
func Fn1837(m *base.Module, l0 int64)

//go:linkname Fn1841 github.com/goccy/llamawasm2go/p0.Fn1841
func Fn1841(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1842 github.com/goccy/llamawasm2go/p2.Fn1842
func Fn1842(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1843 github.com/goccy/llamawasm2go/p2.Fn1843
func Fn1843(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1844 github.com/goccy/llamawasm2go/p2.Fn1844
func Fn1844(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1845 github.com/goccy/llamawasm2go/p2.Fn1845
func Fn1845(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1863 github.com/goccy/llamawasm2go/p2.Fn1863
func Fn1863(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1874 github.com/goccy/llamawasm2go/p2.Fn1874
func Fn1874(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1876 github.com/goccy/llamawasm2go/p2.Fn1876
func Fn1876(m *base.Module, l0 int64) int64

//go:linkname Fn1877 github.com/goccy/llamawasm2go/p2.Fn1877
func Fn1877(m *base.Module, l0 int64)

//go:linkname Fn1880 github.com/goccy/llamawasm2go/p0.Fn1880
func Fn1880(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1882 github.com/goccy/llamawasm2go/p2.Fn1882
func Fn1882(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1893 github.com/goccy/llamawasm2go/p2.Fn1893
func Fn1893(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1896 github.com/goccy/llamawasm2go/p2.Fn1896
func Fn1896(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1897 github.com/goccy/llamawasm2go/p2.Fn1897
func Fn1897(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1907 github.com/goccy/llamawasm2go/p2.Fn1907
func Fn1907(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1934 github.com/goccy/llamawasm2go/p2.Fn1934
func Fn1934(m *base.Module, l0 int64)

//go:linkname Fn1935 github.com/goccy/llamawasm2go/p2.Fn1935
func Fn1935(m *base.Module, l0 int64)

//go:linkname Fn1946 github.com/goccy/llamawasm2go/p2.Fn1946
func Fn1946(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1978 github.com/goccy/llamawasm2go/p2.Fn1978
func Fn1978(m *base.Module) int64

//go:linkname Fn1983 github.com/goccy/llamawasm2go/p2.Fn1983
func Fn1983(m *base.Module, l0 int64) int64

//go:linkname Fn1984 github.com/goccy/llamawasm2go/p2.Fn1984
func Fn1984(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1991 github.com/goccy/llamawasm2go/p2.Fn1991
func Fn1991(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1998 github.com/goccy/llamawasm2go/p2.Fn1998
func Fn1998(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2005 github.com/goccy/llamawasm2go/p2.Fn2005
func Fn2005(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2012 github.com/goccy/llamawasm2go/p2.Fn2012
func Fn2012(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2014 github.com/goccy/llamawasm2go/p2.Fn2014
func Fn2014(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2017 github.com/goccy/llamawasm2go/p2.Fn2017
func Fn2017(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2018 github.com/goccy/llamawasm2go/p2.Fn2018
func Fn2018(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2019 github.com/goccy/llamawasm2go/p2.Fn2019
func Fn2019(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn2022 github.com/goccy/llamawasm2go/p2.Fn2022
func Fn2022(m *base.Module, l0 int64)

//go:linkname Fn2032 github.com/goccy/llamawasm2go/p2.Fn2032
func Fn2032(m *base.Module, l0 int64)

//go:linkname Fn2034 github.com/goccy/llamawasm2go/p2.Fn2034
func Fn2034(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2035 github.com/goccy/llamawasm2go/p2.Fn2035
func Fn2035(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2048 github.com/goccy/llamawasm2go/p2.Fn2048
func Fn2048(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2049 github.com/goccy/llamawasm2go/p2.Fn2049
func Fn2049(m *base.Module, l0 int64) int64

//go:linkname Fn2050 github.com/goccy/llamawasm2go/p2.Fn2050
func Fn2050(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2052 github.com/goccy/llamawasm2go/p2.Fn2052
func Fn2052(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2066 github.com/goccy/llamawasm2go/p2.Fn2066
func Fn2066(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2067 github.com/goccy/llamawasm2go/p2.Fn2067
func Fn2067(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2069 github.com/goccy/llamawasm2go/p2.Fn2069
func Fn2069(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2070 github.com/goccy/llamawasm2go/p2.Fn2070
func Fn2070(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2071 github.com/goccy/llamawasm2go/p2.Fn2071
func Fn2071(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2072 github.com/goccy/llamawasm2go/p2.Fn2072
func Fn2072(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2073 github.com/goccy/llamawasm2go/p2.Fn2073
func Fn2073(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2074 github.com/goccy/llamawasm2go/p2.Fn2074
func Fn2074(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2093 github.com/goccy/llamawasm2go/p2.Fn2093
func Fn2093(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2095 github.com/goccy/llamawasm2go/p2.Fn2095
func Fn2095(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2097 github.com/goccy/llamawasm2go/p2.Fn2097
func Fn2097(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2098 github.com/goccy/llamawasm2go/p2.Fn2098
func Fn2098(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2099 github.com/goccy/llamawasm2go/p2.Fn2099
func Fn2099(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2100 github.com/goccy/llamawasm2go/p2.Fn2100
func Fn2100(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2101 github.com/goccy/llamawasm2go/p2.Fn2101
func Fn2101(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2103 github.com/goccy/llamawasm2go/p2.Fn2103
func Fn2103(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2105 github.com/goccy/llamawasm2go/p2.Fn2105
func Fn2105(m *base.Module, l0 int64)

//go:linkname Fn2122 github.com/goccy/llamawasm2go/p2.Fn2122
func Fn2122(m *base.Module, l0 int64)

//go:linkname Fn2123 github.com/goccy/llamawasm2go/p2.Fn2123
func Fn2123(m *base.Module, l0 int64)

//go:linkname Fn2124 github.com/goccy/llamawasm2go/p2.Fn2124
func Fn2124(m *base.Module, l0 int64)

//go:linkname Fn2126 github.com/goccy/llamawasm2go/p2.Fn2126
func Fn2126(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2127 github.com/goccy/llamawasm2go/p2.Fn2127
func Fn2127(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2166 github.com/goccy/llamawasm2go/p2.Fn2166
func Fn2166(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2175 github.com/goccy/llamawasm2go/p2.Fn2175
func Fn2175(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2177 github.com/goccy/llamawasm2go/p2.Fn2177
func Fn2177(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2179 github.com/goccy/llamawasm2go/p2.Fn2179
func Fn2179(m *base.Module, l0 int64) int64

//go:linkname Fn2182 github.com/goccy/llamawasm2go/p2.Fn2182
func Fn2182(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2186 github.com/goccy/llamawasm2go/p2.Fn2186
func Fn2186(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2191 github.com/goccy/llamawasm2go/p2.Fn2191
func Fn2191(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2205 github.com/goccy/llamawasm2go/p2.Fn2205
func Fn2205(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2207 github.com/goccy/llamawasm2go/p2.Fn2207
func Fn2207(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2208 github.com/goccy/llamawasm2go/p2.Fn2208
func Fn2208(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2222 github.com/goccy/llamawasm2go/p2.Fn2222
func Fn2222(m *base.Module, l0 int64) int64

//go:linkname Fn2224 github.com/goccy/llamawasm2go/p2.Fn2224
func Fn2224(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2227 github.com/goccy/llamawasm2go/p2.Fn2227
func Fn2227(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2232 github.com/goccy/llamawasm2go/p2.Fn2232
func Fn2232(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2245 github.com/goccy/llamawasm2go/p2.Fn2245
func Fn2245(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2246 github.com/goccy/llamawasm2go/p2.Fn2246
func Fn2246(m *base.Module, l0 int64)

//go:linkname Fn2247 github.com/goccy/llamawasm2go/p2.Fn2247
func Fn2247(m *base.Module, l0 int64)

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p0.Fn2249
func Fn2249(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn2250 github.com/goccy/llamawasm2go/p2.Fn2250
func Fn2250(m *base.Module, l0 int64)

//go:linkname Fn2256 github.com/goccy/llamawasm2go/p2.Fn2256
func Fn2256(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2260 github.com/goccy/llamawasm2go/p2.Fn2260
func Fn2260(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2261 github.com/goccy/llamawasm2go/p2.Fn2261
func Fn2261(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2269 github.com/goccy/llamawasm2go/p2.Fn2269
func Fn2269(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2271 github.com/goccy/llamawasm2go/p2.Fn2271
func Fn2271(m *base.Module) int64

//go:linkname Fn2273 github.com/goccy/llamawasm2go/p2.Fn2273
func Fn2273(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2274 github.com/goccy/llamawasm2go/p2.Fn2274
func Fn2274(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2276 github.com/goccy/llamawasm2go/p2.Fn2276
func Fn2276(m *base.Module) int64

//go:linkname Fn2278 github.com/goccy/llamawasm2go/p2.Fn2278
func Fn2278(m *base.Module, l0 int32) int64

//go:linkname Fn2279 github.com/goccy/llamawasm2go/p2.Fn2279
func Fn2279(m *base.Module, l0 int32) int32

//go:linkname Fn2280 github.com/goccy/llamawasm2go/p2.Fn2280
func Fn2280(m *base.Module, l0 int32) int64

//go:linkname Fn2281 github.com/goccy/llamawasm2go/p2.Fn2281
func Fn2281(m *base.Module, l0 float32) int64

//go:linkname Fn2282 github.com/goccy/llamawasm2go/p2.Fn2282
func Fn2282(m *base.Module, l0 float32) int64

//go:linkname Fn2283 github.com/goccy/llamawasm2go/p2.Fn2283
func Fn2283(m *base.Module, l0 float32) int64

//go:linkname Fn2285 github.com/goccy/llamawasm2go/p2.Fn2285
func Fn2285(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2286 github.com/goccy/llamawasm2go/p2.Fn2286
func Fn2286(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2327 github.com/goccy/llamawasm2go/p2.Fn2327
func Fn2327(m *base.Module, l0 int64)

//go:linkname Fn2329 github.com/goccy/llamawasm2go/p2.Fn2329
func Fn2329(m *base.Module, l0 int64)

//go:linkname Fn2337 github.com/goccy/llamawasm2go/p2.Fn2337
func Fn2337(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2371 github.com/goccy/llamawasm2go/p2.Fn2371
func Fn2371(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2382 github.com/goccy/llamawasm2go/p2.Fn2382
func Fn2382(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2385 github.com/goccy/llamawasm2go/p2.Fn2385
func Fn2385(m *base.Module, l0 int64)

//go:linkname Fn2386 github.com/goccy/llamawasm2go/p2.Fn2386
func Fn2386(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2387 github.com/goccy/llamawasm2go/p2.Fn2387
func Fn2387(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2389 github.com/goccy/llamawasm2go/p2.Fn2389
func Fn2389(m *base.Module, l0 int64)

//go:linkname Fn2393 github.com/goccy/llamawasm2go/p2.Fn2393
func Fn2393(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2399 github.com/goccy/llamawasm2go/p2.Fn2399
func Fn2399(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2414 github.com/goccy/llamawasm2go/p2.Fn2414
func Fn2414(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2421 github.com/goccy/llamawasm2go/p0.Fn2421
func Fn2421(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2424 github.com/goccy/llamawasm2go/p2.Fn2424
func Fn2424(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2427 github.com/goccy/llamawasm2go/p2.Fn2427
func Fn2427(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2430 github.com/goccy/llamawasm2go/p2.Fn2430
func Fn2430(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2431 github.com/goccy/llamawasm2go/p2.Fn2431
func Fn2431(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2437 github.com/goccy/llamawasm2go/p2.Fn2437
func Fn2437(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2439 github.com/goccy/llamawasm2go/p2.Fn2439
func Fn2439(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2456 github.com/goccy/llamawasm2go/p2.Fn2456
func Fn2456(m *base.Module, l0 int64)

//go:linkname Fn2457 github.com/goccy/llamawasm2go/p2.Fn2457
func Fn2457(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2462 github.com/goccy/llamawasm2go/p2.Fn2462
func Fn2462(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2467 github.com/goccy/llamawasm2go/p2.Fn2467
func Fn2467(m *base.Module, l0 int64) int64

//go:linkname Fn2468 github.com/goccy/llamawasm2go/p2.Fn2468
func Fn2468(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2469 github.com/goccy/llamawasm2go/p0.Fn2469
func Fn2469(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2480 github.com/goccy/llamawasm2go/p2.Fn2480
func Fn2480(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2481 github.com/goccy/llamawasm2go/p2.Fn2481
func Fn2481(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2482 github.com/goccy/llamawasm2go/p2.Fn2482
func Fn2482(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2483 github.com/goccy/llamawasm2go/p2.Fn2483
func Fn2483(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2503 github.com/goccy/llamawasm2go/p2.Fn2503
func Fn2503(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2504 github.com/goccy/llamawasm2go/p2.Fn2504
func Fn2504(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2507 github.com/goccy/llamawasm2go/p2.Fn2507
func Fn2507(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2593 github.com/goccy/llamawasm2go/p2.Fn2593
func Fn2593(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2712 github.com/goccy/llamawasm2go/p2.Fn2712
func Fn2712(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2713 github.com/goccy/llamawasm2go/p0.Fn2713
func Fn2713(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2714 github.com/goccy/llamawasm2go/p2.Fn2714
func Fn2714(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2715 github.com/goccy/llamawasm2go/p2.Fn2715
func Fn2715(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2719 github.com/goccy/llamawasm2go/p2.Fn2719
func Fn2719(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2744 github.com/goccy/llamawasm2go/p2.Fn2744
func Fn2744(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2876 github.com/goccy/llamawasm2go/p2.Fn2876
func Fn2876(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2883 github.com/goccy/llamawasm2go/p2.Fn2883
func Fn2883(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2911 github.com/goccy/llamawasm2go/p2.Fn2911
func Fn2911(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2925 github.com/goccy/llamawasm2go/p2.Fn2925
func Fn2925(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2937 github.com/goccy/llamawasm2go/p2.Fn2937
func Fn2937(m *base.Module, l0 int32)

//go:linkname Fn2939 github.com/goccy/llamawasm2go/p2.Fn2939
func Fn2939(m *base.Module, l0 int64) int64

//go:linkname Fn2940 github.com/goccy/llamawasm2go/p2.Fn2940
func Fn2940(m *base.Module, l0 int64)

//go:linkname Fn2943 github.com/goccy/llamawasm2go/p2.Fn2943
func Fn2943(m *base.Module, l0 int64)

//go:linkname Fn2944 github.com/goccy/llamawasm2go/p2.Fn2944
func Fn2944(m *base.Module, l0 int64)

//go:linkname Fn2946 github.com/goccy/llamawasm2go/p2.Fn2946
func Fn2946(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2947 github.com/goccy/llamawasm2go/p2.Fn2947
func Fn2947(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2953 github.com/goccy/llamawasm2go/p2.Fn2953
func Fn2953(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2955 github.com/goccy/llamawasm2go/p2.Fn2955
func Fn2955(m *base.Module, l0 int64) int32

//go:linkname Fn2959 github.com/goccy/llamawasm2go/p2.Fn2959
func Fn2959(m *base.Module) int32

//go:linkname Fn2970 github.com/goccy/llamawasm2go/p2.Fn2970
func Fn2970(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2974 github.com/goccy/llamawasm2go/p2.Fn2974
func Fn2974(m *base.Module) int64

//go:linkname Fn2976 github.com/goccy/llamawasm2go/p2.Fn2976
func Fn2976(m *base.Module, l0 float64) float32

//go:linkname Fn2977 github.com/goccy/llamawasm2go/p2.Fn2977
func Fn2977(m *base.Module, l0 float64) float32

//go:linkname Fn2981 github.com/goccy/llamawasm2go/p2.Fn2981
func Fn2981(m *base.Module, l0 float64) float64

//go:linkname Fn2984 github.com/goccy/llamawasm2go/p2.Fn2984
func Fn2984(m *base.Module, l0 int32) float32

//go:linkname Fn2985 github.com/goccy/llamawasm2go/p2.Fn2985
func Fn2985(m *base.Module, l0 int32) float32

//go:linkname Fn2988 github.com/goccy/llamawasm2go/p2.Fn2988
func Fn2988(m *base.Module, l0 float32) float32

//go:linkname Fn2991 github.com/goccy/llamawasm2go/p2.Fn2991
func Fn2991(m *base.Module, l0 float64) float64

//go:linkname Fn2992 github.com/goccy/llamawasm2go/p2.Fn2992
func Fn2992(m *base.Module, l0 float64) float64

//go:linkname Fn2993 github.com/goccy/llamawasm2go/p2.Fn2993
func Fn2993(m *base.Module, l0 float32) float32

//go:linkname Fn2995 github.com/goccy/llamawasm2go/p2.Fn2995
func Fn2995(m *base.Module, l0 float32) float32

//go:linkname Fn2997 github.com/goccy/llamawasm2go/p2.Fn2997
func Fn2997(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn2998 github.com/goccy/llamawasm2go/p2.Fn2998
func Fn2998(m *base.Module, l0 float32) float32

//go:linkname Fn3015 github.com/goccy/llamawasm2go/p2.Fn3015
func Fn3015(m *base.Module, l0 int64) int32

//go:linkname Fn3016 github.com/goccy/llamawasm2go/p2.Fn3016
func Fn3016(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3018 github.com/goccy/llamawasm2go/p2.Fn3018
func Fn3018(m *base.Module, l0 int64)

//go:linkname Fn3019 github.com/goccy/llamawasm2go/p2.Fn3019
func Fn3019(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3020 github.com/goccy/llamawasm2go/p2.Fn3020
func Fn3020(m *base.Module, l0 int64) int32

//go:linkname Fn3027 github.com/goccy/llamawasm2go/p2.Fn3027
func Fn3027(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3029 github.com/goccy/llamawasm2go/p2.Fn3029
func Fn3029(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3035 github.com/goccy/llamawasm2go/p2.Fn3035
func Fn3035(m *base.Module, l0 int64) int32

//go:linkname Fn3038 github.com/goccy/llamawasm2go/p2.Fn3038
func Fn3038(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3041 github.com/goccy/llamawasm2go/p2.Fn3041
func Fn3041(m *base.Module, l0 int64) int32

//go:linkname Fn3043 github.com/goccy/llamawasm2go/p2.Fn3043
func Fn3043(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3045 github.com/goccy/llamawasm2go/p2.Fn3045
func Fn3045(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3046 github.com/goccy/llamawasm2go/p2.Fn3046
func Fn3046(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3049 github.com/goccy/llamawasm2go/p2.Fn3049
func Fn3049(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3052 github.com/goccy/llamawasm2go/p2.Fn3052
func Fn3052(m *base.Module, l0 int64) int64

//go:linkname Fn3056 github.com/goccy/llamawasm2go/p2.Fn3056
func Fn3056(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3066 github.com/goccy/llamawasm2go/p2.Fn3066
func Fn3066(m *base.Module)

//go:linkname Fn3067 github.com/goccy/llamawasm2go/p0.Fn3067
func Fn3067(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3070 github.com/goccy/llamawasm2go/p2.Fn3070
func Fn3070(m *base.Module)

//go:linkname Fn3072 github.com/goccy/llamawasm2go/p0.Fn3072
func Fn3072(m *base.Module, l0 int64) int64

//go:linkname Fn3074 github.com/goccy/llamawasm2go/p2.Fn3074
func Fn3074(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3078 github.com/goccy/llamawasm2go/p2.Fn3078
func Fn3078(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3147 github.com/goccy/llamawasm2go/p2.Fn3147
func Fn3147(m *base.Module, l0 int32)
