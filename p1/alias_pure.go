//go:build !arm64 && (!amd64 || !amd64.v2)

package p1

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn21 github.com/goccy/llamawasm2go/p2.Fn21
func Fn21(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn47 github.com/goccy/llamawasm2go/p2.Fn47
func Fn47(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn50 github.com/goccy/llamawasm2go/p2.Fn50
func Fn50(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn53 github.com/goccy/llamawasm2go/p2.Fn53
func Fn53(m *base.Module, l0 int64) int64

//go:linkname Fn54 github.com/goccy/llamawasm2go/p2.Fn54
func Fn54(m *base.Module, l0 int64)

//go:linkname Fn56 github.com/goccy/llamawasm2go/p2.Fn56
func Fn56(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn57 github.com/goccy/llamawasm2go/p2.Fn57
func Fn57(m *base.Module, l0 int64) int64

//go:linkname Fn58 github.com/goccy/llamawasm2go/p2.Fn58
func Fn58(m *base.Module)

//go:linkname Fn64 github.com/goccy/llamawasm2go/p0.Fn64
func Fn64(m *base.Module, l0 int64) int64

//go:linkname Fn65 github.com/goccy/llamawasm2go/p2.Fn65
func Fn65(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn66 github.com/goccy/llamawasm2go/p0.Fn66
func Fn66(m *base.Module, l0 int64) int64

//go:linkname Fn240 github.com/goccy/llamawasm2go/p2.Fn240
func Fn240(m *base.Module)

//go:linkname Fn241 github.com/goccy/llamawasm2go/p2.Fn241
func Fn241(m *base.Module, l0 int64)

//go:linkname Fn242 github.com/goccy/llamawasm2go/p2.Fn242
func Fn242(m *base.Module, l0 int64) int64

//go:linkname Fn249 github.com/goccy/llamawasm2go/p2.Fn249
func Fn249(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn250 github.com/goccy/llamawasm2go/p2.Fn250
func Fn250(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn251 github.com/goccy/llamawasm2go/p2.Fn251
func Fn251(m *base.Module)

//go:linkname Fn252 github.com/goccy/llamawasm2go/p2.Fn252
func Fn252(m *base.Module)

//go:linkname Fn253 github.com/goccy/llamawasm2go/p2.Fn253
func Fn253(m *base.Module, l0 int64)

//go:linkname Fn257 github.com/goccy/llamawasm2go/p2.Fn257
func Fn257(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn258 github.com/goccy/llamawasm2go/p2.Fn258
func Fn258(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn259 github.com/goccy/llamawasm2go/p2.Fn259
func Fn259(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn261 github.com/goccy/llamawasm2go/p2.Fn261
func Fn261(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn262 github.com/goccy/llamawasm2go/p2.Fn262
func Fn262(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn264 github.com/goccy/llamawasm2go/p2.Fn264
func Fn264(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn265 github.com/goccy/llamawasm2go/p2.Fn265
func Fn265(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn266 github.com/goccy/llamawasm2go/p2.Fn266
func Fn266(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn267 github.com/goccy/llamawasm2go/p2.Fn267
func Fn267(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn268 github.com/goccy/llamawasm2go/p2.Fn268
func Fn268(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn270 github.com/goccy/llamawasm2go/p2.Fn270
func Fn270(m *base.Module)

//go:linkname Fn271 github.com/goccy/llamawasm2go/p2.Fn271
func Fn271(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn273 github.com/goccy/llamawasm2go/p2.Fn273
func Fn273(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn274 github.com/goccy/llamawasm2go/p2.Fn274
func Fn274(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn275 github.com/goccy/llamawasm2go/p2.Fn275
func Fn275(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn278 github.com/goccy/llamawasm2go/p2.Fn278
func Fn278(m *base.Module, l0 int64) int64

//go:linkname Fn279 github.com/goccy/llamawasm2go/p2.Fn279
func Fn279(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn282 github.com/goccy/llamawasm2go/p2.Fn282
func Fn282(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn283 github.com/goccy/llamawasm2go/p2.Fn283
func Fn283(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn289 github.com/goccy/llamawasm2go/p2.Fn289
func Fn289(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn290 github.com/goccy/llamawasm2go/p2.Fn290
func Fn290(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn330 github.com/goccy/llamawasm2go/p2.Fn330
func Fn330(m *base.Module, l0 int64) int32

//go:linkname Fn331 github.com/goccy/llamawasm2go/p2.Fn331
func Fn331(m *base.Module, l0 int64)

//go:linkname Fn332 github.com/goccy/llamawasm2go/p2.Fn332
func Fn332(m *base.Module, l0 int64)

//go:linkname Fn350 github.com/goccy/llamawasm2go/p2.Fn350
func Fn350(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn351 github.com/goccy/llamawasm2go/p2.Fn351
func Fn351(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn355 github.com/goccy/llamawasm2go/p2.Fn355
func Fn355(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn356 github.com/goccy/llamawasm2go/p2.Fn356
func Fn356(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn359 github.com/goccy/llamawasm2go/p2.Fn359
func Fn359(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn360 github.com/goccy/llamawasm2go/p2.Fn360
func Fn360(m *base.Module, l0 int64)

//go:linkname Fn361 github.com/goccy/llamawasm2go/p2.Fn361
func Fn361(m *base.Module, l0 int64)

//go:linkname Fn363 github.com/goccy/llamawasm2go/p2.Fn363
func Fn363(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn365 github.com/goccy/llamawasm2go/p2.Fn365
func Fn365(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn366 github.com/goccy/llamawasm2go/p2.Fn366
func Fn366(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn368 github.com/goccy/llamawasm2go/p2.Fn368
func Fn368(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn371 github.com/goccy/llamawasm2go/p0.Fn371
func Fn371(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int64)

//go:linkname Fn374 github.com/goccy/llamawasm2go/p2.Fn374
func Fn374(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn375 github.com/goccy/llamawasm2go/p2.Fn375
func Fn375(m *base.Module, l0 int64)

//go:linkname Fn378 github.com/goccy/llamawasm2go/p2.Fn378
func Fn378(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn379 github.com/goccy/llamawasm2go/p2.Fn379
func Fn379(m *base.Module, l0 int64)

//go:linkname Fn380 github.com/goccy/llamawasm2go/p2.Fn380
func Fn380(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn381 github.com/goccy/llamawasm2go/p2.Fn381
func Fn381(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn382 github.com/goccy/llamawasm2go/p2.Fn382
func Fn382(m *base.Module)

//go:linkname Fn384 github.com/goccy/llamawasm2go/p2.Fn384
func Fn384(m *base.Module, l0 int64) int32

//go:linkname Fn385 github.com/goccy/llamawasm2go/p2.Fn385
func Fn385(m *base.Module, l0 int64)

//go:linkname Fn387 github.com/goccy/llamawasm2go/p0.Fn387
func Fn387(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 int64)

//go:linkname Fn402 github.com/goccy/llamawasm2go/p2.Fn402
func Fn402(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn403 github.com/goccy/llamawasm2go/p2.Fn403
func Fn403(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn414 github.com/goccy/llamawasm2go/p2.Fn414
func Fn414(m *base.Module, l0 int64) int32

//go:linkname Fn415 github.com/goccy/llamawasm2go/p2.Fn415
func Fn415(m *base.Module, l0 int64) int32

//go:linkname Fn425 github.com/goccy/llamawasm2go/p2.Fn425
func Fn425(m *base.Module, l0 int64) int64

//go:linkname Fn430 github.com/goccy/llamawasm2go/p2.Fn430
func Fn430(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn431 github.com/goccy/llamawasm2go/p2.Fn431
func Fn431(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn432 github.com/goccy/llamawasm2go/p2.Fn432
func Fn432(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn433 github.com/goccy/llamawasm2go/p2.Fn433
func Fn433(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn434 github.com/goccy/llamawasm2go/p2.Fn434
func Fn434(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn435 github.com/goccy/llamawasm2go/p2.Fn435
func Fn435(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn437 github.com/goccy/llamawasm2go/p2.Fn437
func Fn437(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn438 github.com/goccy/llamawasm2go/p2.Fn438
func Fn438(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn439 github.com/goccy/llamawasm2go/p2.Fn439
func Fn439(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn440 github.com/goccy/llamawasm2go/p2.Fn440
func Fn440(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn441 github.com/goccy/llamawasm2go/p2.Fn441
func Fn441(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn442 github.com/goccy/llamawasm2go/p2.Fn442
func Fn442(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn443 github.com/goccy/llamawasm2go/p2.Fn443
func Fn443(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn444 github.com/goccy/llamawasm2go/p2.Fn444
func Fn444(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn446 github.com/goccy/llamawasm2go/p2.Fn446
func Fn446(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn448 github.com/goccy/llamawasm2go/p2.Fn448
func Fn448(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn450 github.com/goccy/llamawasm2go/p2.Fn450
func Fn450(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn451 github.com/goccy/llamawasm2go/p2.Fn451
func Fn451(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn452 github.com/goccy/llamawasm2go/p2.Fn452
func Fn452(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn453 github.com/goccy/llamawasm2go/p2.Fn453
func Fn453(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn454 github.com/goccy/llamawasm2go/p2.Fn454
func Fn454(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn456 github.com/goccy/llamawasm2go/p2.Fn456
func Fn456(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn457 github.com/goccy/llamawasm2go/p2.Fn457
func Fn457(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn458 github.com/goccy/llamawasm2go/p2.Fn458
func Fn458(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn459 github.com/goccy/llamawasm2go/p2.Fn459
func Fn459(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn460 github.com/goccy/llamawasm2go/p2.Fn460
func Fn460(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn461 github.com/goccy/llamawasm2go/p2.Fn461
func Fn461(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn462 github.com/goccy/llamawasm2go/p2.Fn462
func Fn462(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn463 github.com/goccy/llamawasm2go/p2.Fn463
func Fn463(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn464 github.com/goccy/llamawasm2go/p2.Fn464
func Fn464(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn465 github.com/goccy/llamawasm2go/p2.Fn465
func Fn465(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn466 github.com/goccy/llamawasm2go/p2.Fn466
func Fn466(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn467 github.com/goccy/llamawasm2go/p2.Fn467
func Fn467(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn468 github.com/goccy/llamawasm2go/p2.Fn468
func Fn468(m *base.Module, l0 int64)

//go:linkname Fn469 github.com/goccy/llamawasm2go/p2.Fn469
func Fn469(m *base.Module, l0 int64)

//go:linkname Fn470 github.com/goccy/llamawasm2go/p2.Fn470
func Fn470(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn471 github.com/goccy/llamawasm2go/p2.Fn471
func Fn471(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn472 github.com/goccy/llamawasm2go/p2.Fn472
func Fn472(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn473 github.com/goccy/llamawasm2go/p2.Fn473
func Fn473(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn474 github.com/goccy/llamawasm2go/p2.Fn474
func Fn474(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn475 github.com/goccy/llamawasm2go/p2.Fn475
func Fn475(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn476 github.com/goccy/llamawasm2go/p2.Fn476
func Fn476(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn477 github.com/goccy/llamawasm2go/p2.Fn477
func Fn477(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn478 github.com/goccy/llamawasm2go/p2.Fn478
func Fn478(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn479 github.com/goccy/llamawasm2go/p2.Fn479
func Fn479(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn481 github.com/goccy/llamawasm2go/p2.Fn481
func Fn481(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn482 github.com/goccy/llamawasm2go/p2.Fn482
func Fn482(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn483 github.com/goccy/llamawasm2go/p2.Fn483
func Fn483(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn484 github.com/goccy/llamawasm2go/p2.Fn484
func Fn484(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn485 github.com/goccy/llamawasm2go/p2.Fn485
func Fn485(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn486 github.com/goccy/llamawasm2go/p2.Fn486
func Fn486(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn487 github.com/goccy/llamawasm2go/p2.Fn487
func Fn487(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64) int64

//go:linkname Fn488 github.com/goccy/llamawasm2go/p2.Fn488
func Fn488(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn489 github.com/goccy/llamawasm2go/p2.Fn489
func Fn489(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn490 github.com/goccy/llamawasm2go/p2.Fn490
func Fn490(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn491 github.com/goccy/llamawasm2go/p2.Fn491
func Fn491(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn492 github.com/goccy/llamawasm2go/p2.Fn492
func Fn492(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn493 github.com/goccy/llamawasm2go/p2.Fn493
func Fn493(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32, l14 int32) int64

//go:linkname Fn495 github.com/goccy/llamawasm2go/p2.Fn495
func Fn495(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn496 github.com/goccy/llamawasm2go/p2.Fn496
func Fn496(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn497 github.com/goccy/llamawasm2go/p2.Fn497
func Fn497(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn499 github.com/goccy/llamawasm2go/p2.Fn499
func Fn499(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn501 github.com/goccy/llamawasm2go/p2.Fn501
func Fn501(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn502 github.com/goccy/llamawasm2go/p2.Fn502
func Fn502(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn503 github.com/goccy/llamawasm2go/p2.Fn503
func Fn503(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn505 github.com/goccy/llamawasm2go/p2.Fn505
func Fn505(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 float32, l6 float32, l7 float32) int64

//go:linkname Fn506 github.com/goccy/llamawasm2go/p2.Fn506
func Fn506(m *base.Module, l0 int64)

//go:linkname Fn507 github.com/goccy/llamawasm2go/p2.Fn507
func Fn507(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn511 github.com/goccy/llamawasm2go/p2.Fn511
func Fn511(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn512 github.com/goccy/llamawasm2go/p2.Fn512
func Fn512(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn514 github.com/goccy/llamawasm2go/p2.Fn514
func Fn514(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn516 github.com/goccy/llamawasm2go/p2.Fn516
func Fn516(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn518 github.com/goccy/llamawasm2go/p2.Fn518
func Fn518(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn522 github.com/goccy/llamawasm2go/p2.Fn522
func Fn522(m *base.Module, l0 int64)

//go:linkname Fn525 github.com/goccy/llamawasm2go/p2.Fn525
func Fn525(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn527 github.com/goccy/llamawasm2go/p2.Fn527
func Fn527(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn530 github.com/goccy/llamawasm2go/p2.Fn530
func Fn530(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn531 github.com/goccy/llamawasm2go/p2.Fn531
func Fn531(m *base.Module, l0 int64) int64

//go:linkname Fn532 github.com/goccy/llamawasm2go/p2.Fn532
func Fn532(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn533 github.com/goccy/llamawasm2go/p2.Fn533
func Fn533(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn534 github.com/goccy/llamawasm2go/p2.Fn534
func Fn534(m *base.Module, l0 int64) int64

//go:linkname Fn535 github.com/goccy/llamawasm2go/p2.Fn535
func Fn535(m *base.Module, l0 int64) int64

//go:linkname Fn536 github.com/goccy/llamawasm2go/p2.Fn536
func Fn536(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn538 github.com/goccy/llamawasm2go/p2.Fn538
func Fn538(m *base.Module, l0 int64) int64

//go:linkname Fn539 github.com/goccy/llamawasm2go/p2.Fn539
func Fn539(m *base.Module, l0 int64) int64

//go:linkname Fn541 github.com/goccy/llamawasm2go/p2.Fn541
func Fn541(m *base.Module, l0 int64)

//go:linkname Fn542 github.com/goccy/llamawasm2go/p2.Fn542
func Fn542(m *base.Module, l0 int64) int64

//go:linkname Fn543 github.com/goccy/llamawasm2go/p2.Fn543
func Fn543(m *base.Module, l0 int64) int64

//go:linkname Fn544 github.com/goccy/llamawasm2go/p2.Fn544
func Fn544(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn545 github.com/goccy/llamawasm2go/p2.Fn545
func Fn545(m *base.Module, l0 int64) int32

//go:linkname Fn546 github.com/goccy/llamawasm2go/p2.Fn546
func Fn546(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn548 github.com/goccy/llamawasm2go/p2.Fn548
func Fn548(m *base.Module, l0 int64)

//go:linkname Fn549 github.com/goccy/llamawasm2go/p2.Fn549
func Fn549(m *base.Module, l0 int64) int64

//go:linkname Fn551 github.com/goccy/llamawasm2go/p2.Fn551
func Fn551(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn552 github.com/goccy/llamawasm2go/p2.Fn552
func Fn552(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn553 github.com/goccy/llamawasm2go/p2.Fn553
func Fn553(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn554 github.com/goccy/llamawasm2go/p2.Fn554
func Fn554(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn555 github.com/goccy/llamawasm2go/p2.Fn555
func Fn555(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn556 github.com/goccy/llamawasm2go/p2.Fn556
func Fn556(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn559 github.com/goccy/llamawasm2go/p2.Fn559
func Fn559(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn560 github.com/goccy/llamawasm2go/p2.Fn560
func Fn560(m *base.Module, l0 int64) int64

//go:linkname Fn561 github.com/goccy/llamawasm2go/p2.Fn561
func Fn561(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn564 github.com/goccy/llamawasm2go/p2.Fn564
func Fn564(m *base.Module, l0 int64) int64

//go:linkname Fn565 github.com/goccy/llamawasm2go/p2.Fn565
func Fn565(m *base.Module, l0 int64) int64

//go:linkname Fn569 github.com/goccy/llamawasm2go/p2.Fn569
func Fn569(m *base.Module, l0 int64) int64

//go:linkname Fn570 github.com/goccy/llamawasm2go/p2.Fn570
func Fn570(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn571 github.com/goccy/llamawasm2go/p2.Fn571
func Fn571(m *base.Module, l0 int64) int64

//go:linkname Fn575 github.com/goccy/llamawasm2go/p2.Fn575
func Fn575(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn576 github.com/goccy/llamawasm2go/p0.Fn576
func Fn576(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn581 github.com/goccy/llamawasm2go/p2.Fn581
func Fn581(m *base.Module, l0 int64)

//go:linkname Fn582 github.com/goccy/llamawasm2go/p2.Fn582
func Fn582(m *base.Module, l0 int64)

//go:linkname Fn583 github.com/goccy/llamawasm2go/p2.Fn583
func Fn583(m *base.Module, l0 int64)

//go:linkname Fn584 github.com/goccy/llamawasm2go/p2.Fn584
func Fn584(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn585 github.com/goccy/llamawasm2go/p2.Fn585
func Fn585(m *base.Module, l0 int64) int32

//go:linkname Fn586 github.com/goccy/llamawasm2go/p2.Fn586
func Fn586(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn587 github.com/goccy/llamawasm2go/p2.Fn587
func Fn587(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn588 github.com/goccy/llamawasm2go/p2.Fn588
func Fn588(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn589 github.com/goccy/llamawasm2go/p2.Fn589
func Fn589(m *base.Module, l0 int64) int32

//go:linkname Fn590 github.com/goccy/llamawasm2go/p2.Fn590
func Fn590(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn608 github.com/goccy/llamawasm2go/p2.Fn608
func Fn608(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn613 github.com/goccy/llamawasm2go/p2.Fn613
func Fn613(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn614 github.com/goccy/llamawasm2go/p2.Fn614
func Fn614(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn615 github.com/goccy/llamawasm2go/p2.Fn615
func Fn615(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn617 github.com/goccy/llamawasm2go/p2.Fn617
func Fn617(m *base.Module, l0 int64)

//go:linkname Fn619 github.com/goccy/llamawasm2go/p2.Fn619
func Fn619(m *base.Module, l0 int64)

//go:linkname Fn620 github.com/goccy/llamawasm2go/p0.Fn620
func Fn620(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn622 github.com/goccy/llamawasm2go/p2.Fn622
func Fn622(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn636 github.com/goccy/llamawasm2go/p2.Fn636
func Fn636(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn638 github.com/goccy/llamawasm2go/p2.Fn638
func Fn638(m *base.Module, l0 int64)

//go:linkname Fn639 github.com/goccy/llamawasm2go/p2.Fn639
func Fn639(m *base.Module, l0 int64) int64

//go:linkname Fn662 github.com/goccy/llamawasm2go/p2.Fn662
func Fn662(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn689 github.com/goccy/llamawasm2go/p2.Fn689
func Fn689(m *base.Module, l0 int64)

//go:linkname Fn690 github.com/goccy/llamawasm2go/p2.Fn690
func Fn690(m *base.Module, l0 int64)

//go:linkname Fn755 github.com/goccy/llamawasm2go/p2.Fn755
func Fn755(m *base.Module, l0 int64)

//go:linkname Fn758 github.com/goccy/llamawasm2go/p2.Fn758
func Fn758(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn800 github.com/goccy/llamawasm2go/p2.Fn800
func Fn800(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn859 github.com/goccy/llamawasm2go/p2.Fn859
func Fn859(m *base.Module) int64

//go:linkname Fn896 github.com/goccy/llamawasm2go/p2.Fn896
func Fn896(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32)

//go:linkname Fn905 github.com/goccy/llamawasm2go/p2.Fn905
func Fn905(m *base.Module, l0 int64)

//go:linkname Fn912 github.com/goccy/llamawasm2go/p2.Fn912
func Fn912(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn914 github.com/goccy/llamawasm2go/p2.Fn914
func Fn914(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn929 github.com/goccy/llamawasm2go/p2.Fn929
func Fn929(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn940 github.com/goccy/llamawasm2go/p2.Fn940
func Fn940(m *base.Module, l0 int64)

//go:linkname Fn945 github.com/goccy/llamawasm2go/p2.Fn945
func Fn945(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn990 github.com/goccy/llamawasm2go/p0.Fn990
func Fn990(m *base.Module, l0 int64) int64

//go:linkname Fn992 github.com/goccy/llamawasm2go/p2.Fn992
func Fn992(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1069 github.com/goccy/llamawasm2go/p2.Fn1069
func Fn1069(m *base.Module, l0 int64)

//go:linkname Fn1091 github.com/goccy/llamawasm2go/p2.Fn1091
func Fn1091(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1099 github.com/goccy/llamawasm2go/p2.Fn1099
func Fn1099(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1109 github.com/goccy/llamawasm2go/p2.Fn1109
func Fn1109(m *base.Module, l0 int64) int64

//go:linkname Fn1139 github.com/goccy/llamawasm2go/p2.Fn1139
func Fn1139(m *base.Module, l0 int32) int64

//go:linkname Fn1150 github.com/goccy/llamawasm2go/p2.Fn1150
func Fn1150(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1151 github.com/goccy/llamawasm2go/p2.Fn1151
func Fn1151(m *base.Module, l0 int64)

//go:linkname Fn1153 github.com/goccy/llamawasm2go/p2.Fn1153
func Fn1153(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1173 github.com/goccy/llamawasm2go/p2.Fn1173
func Fn1173(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1174 github.com/goccy/llamawasm2go/p2.Fn1174
func Fn1174(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1175 github.com/goccy/llamawasm2go/p2.Fn1175
func Fn1175(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1176 github.com/goccy/llamawasm2go/p2.Fn1176
func Fn1176(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1177 github.com/goccy/llamawasm2go/p2.Fn1177
func Fn1177(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1178 github.com/goccy/llamawasm2go/p2.Fn1178
func Fn1178(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1179 github.com/goccy/llamawasm2go/p2.Fn1179
func Fn1179(m *base.Module, l0 int64) int64

//go:linkname Fn1182 github.com/goccy/llamawasm2go/p2.Fn1182
func Fn1182(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1186 github.com/goccy/llamawasm2go/p2.Fn1186
func Fn1186(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1187 github.com/goccy/llamawasm2go/p2.Fn1187
func Fn1187(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1188 github.com/goccy/llamawasm2go/p2.Fn1188
func Fn1188(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1189 github.com/goccy/llamawasm2go/p2.Fn1189
func Fn1189(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1190 github.com/goccy/llamawasm2go/p2.Fn1190
func Fn1190(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1191 github.com/goccy/llamawasm2go/p2.Fn1191
func Fn1191(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1192 github.com/goccy/llamawasm2go/p2.Fn1192
func Fn1192(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1193 github.com/goccy/llamawasm2go/p2.Fn1193
func Fn1193(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1194 github.com/goccy/llamawasm2go/p2.Fn1194
func Fn1194(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1195 github.com/goccy/llamawasm2go/p2.Fn1195
func Fn1195(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1196 github.com/goccy/llamawasm2go/p2.Fn1196
func Fn1196(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1197 github.com/goccy/llamawasm2go/p2.Fn1197
func Fn1197(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1198 github.com/goccy/llamawasm2go/p2.Fn1198
func Fn1198(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1199 github.com/goccy/llamawasm2go/p2.Fn1199
func Fn1199(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1200 github.com/goccy/llamawasm2go/p2.Fn1200
func Fn1200(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1201 github.com/goccy/llamawasm2go/p2.Fn1201
func Fn1201(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1202 github.com/goccy/llamawasm2go/p2.Fn1202
func Fn1202(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1203 github.com/goccy/llamawasm2go/p2.Fn1203
func Fn1203(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1204 github.com/goccy/llamawasm2go/p2.Fn1204
func Fn1204(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1206 github.com/goccy/llamawasm2go/p2.Fn1206
func Fn1206(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1207 github.com/goccy/llamawasm2go/p2.Fn1207
func Fn1207(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1208 github.com/goccy/llamawasm2go/p2.Fn1208
func Fn1208(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1209 github.com/goccy/llamawasm2go/p2.Fn1209
func Fn1209(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1210 github.com/goccy/llamawasm2go/p2.Fn1210
func Fn1210(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1211 github.com/goccy/llamawasm2go/p2.Fn1211
func Fn1211(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1212 github.com/goccy/llamawasm2go/p2.Fn1212
func Fn1212(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1214 github.com/goccy/llamawasm2go/p2.Fn1214
func Fn1214(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1218 github.com/goccy/llamawasm2go/p2.Fn1218
func Fn1218(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1219 github.com/goccy/llamawasm2go/p2.Fn1219
func Fn1219(m *base.Module)

//go:linkname Fn1220 github.com/goccy/llamawasm2go/p2.Fn1220
func Fn1220(m *base.Module)

//go:linkname Fn1221 github.com/goccy/llamawasm2go/p0.Fn1221
func Fn1221(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1229 github.com/goccy/llamawasm2go/p2.Fn1229
func Fn1229(m *base.Module)

//go:linkname Fn1231 github.com/goccy/llamawasm2go/p2.Fn1231
func Fn1231(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1239 github.com/goccy/llamawasm2go/p2.Fn1239
func Fn1239(m *base.Module, l0 int64)

//go:linkname Fn1246 github.com/goccy/llamawasm2go/p2.Fn1246
func Fn1246(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1249 github.com/goccy/llamawasm2go/p2.Fn1249
func Fn1249(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1251 github.com/goccy/llamawasm2go/p2.Fn1251
func Fn1251(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1253 github.com/goccy/llamawasm2go/p2.Fn1253
func Fn1253(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1255 github.com/goccy/llamawasm2go/p2.Fn1255
func Fn1255(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1256 github.com/goccy/llamawasm2go/p2.Fn1256
func Fn1256(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1257 github.com/goccy/llamawasm2go/p2.Fn1257
func Fn1257(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1262 github.com/goccy/llamawasm2go/p2.Fn1262
func Fn1262(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1263 github.com/goccy/llamawasm2go/p2.Fn1263
func Fn1263(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1264 github.com/goccy/llamawasm2go/p2.Fn1264
func Fn1264(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1265 github.com/goccy/llamawasm2go/p0.Fn1265
func Fn1265(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1267 github.com/goccy/llamawasm2go/p2.Fn1267
func Fn1267(m *base.Module, l0 int64)

//go:linkname Fn1268 github.com/goccy/llamawasm2go/p2.Fn1268
func Fn1268(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1270 github.com/goccy/llamawasm2go/p2.Fn1270
func Fn1270(m *base.Module, l0 int64) int64

//go:linkname Fn1271 github.com/goccy/llamawasm2go/p2.Fn1271
func Fn1271(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1272 github.com/goccy/llamawasm2go/p2.Fn1272
func Fn1272(m *base.Module, l0 int64)

//go:linkname Fn1273 github.com/goccy/llamawasm2go/p2.Fn1273
func Fn1273(m *base.Module, l0 int64)

//go:linkname Fn1274 github.com/goccy/llamawasm2go/p2.Fn1274
func Fn1274(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1275 github.com/goccy/llamawasm2go/p2.Fn1275
func Fn1275(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1278 github.com/goccy/llamawasm2go/p2.Fn1278
func Fn1278(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1285 github.com/goccy/llamawasm2go/p2.Fn1285
func Fn1285(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1287 github.com/goccy/llamawasm2go/p0.Fn1287
func Fn1287(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1289 github.com/goccy/llamawasm2go/p2.Fn1289
func Fn1289(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1290 github.com/goccy/llamawasm2go/p2.Fn1290
func Fn1290(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1292 github.com/goccy/llamawasm2go/p0.Fn1292
func Fn1292(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1293 github.com/goccy/llamawasm2go/p2.Fn1293
func Fn1293(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1297 github.com/goccy/llamawasm2go/p2.Fn1297
func Fn1297(m *base.Module, l0 int64)

//go:linkname Fn1300 github.com/goccy/llamawasm2go/p2.Fn1300
func Fn1300(m *base.Module, l0 int64)

//go:linkname Fn1302 github.com/goccy/llamawasm2go/p2.Fn1302
func Fn1302(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1303 github.com/goccy/llamawasm2go/p2.Fn1303
func Fn1303(m *base.Module, l0 int64) int64

//go:linkname Fn1304 github.com/goccy/llamawasm2go/p2.Fn1304
func Fn1304(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1305 github.com/goccy/llamawasm2go/p2.Fn1305
func Fn1305(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1306 github.com/goccy/llamawasm2go/p2.Fn1306
func Fn1306(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1307 github.com/goccy/llamawasm2go/p2.Fn1307
func Fn1307(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1308 github.com/goccy/llamawasm2go/p2.Fn1308
func Fn1308(m *base.Module, l0 int64)

//go:linkname Fn1311 github.com/goccy/llamawasm2go/p2.Fn1311
func Fn1311(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1313 github.com/goccy/llamawasm2go/p2.Fn1313
func Fn1313(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1316 github.com/goccy/llamawasm2go/p2.Fn1316
func Fn1316(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p2.Fn1324
func Fn1324(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1330 github.com/goccy/llamawasm2go/p2.Fn1330
func Fn1330(m *base.Module, l0 int64)

//go:linkname Fn1331 github.com/goccy/llamawasm2go/p2.Fn1331
func Fn1331(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1333 github.com/goccy/llamawasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1335 github.com/goccy/llamawasm2go/p2.Fn1335
func Fn1335(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1377 github.com/goccy/llamawasm2go/p2.Fn1377
func Fn1377(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1393 github.com/goccy/llamawasm2go/p2.Fn1393
func Fn1393(m *base.Module, l0 int64)

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p2.Fn1394
func Fn1394(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p2.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1396 github.com/goccy/llamawasm2go/p2.Fn1396
func Fn1396(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p2.Fn1397
func Fn1397(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64)

//go:linkname Fn1399 github.com/goccy/llamawasm2go/p2.Fn1399
func Fn1399(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p2.Fn1401
func Fn1401(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p2.Fn1402
func Fn1402(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1404 github.com/goccy/llamawasm2go/p2.Fn1404
func Fn1404(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p2.Fn1407
func Fn1407(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1409 github.com/goccy/llamawasm2go/p2.Fn1409
func Fn1409(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1410 github.com/goccy/llamawasm2go/p2.Fn1410
func Fn1410(m *base.Module, l0 int64) int64

//go:linkname Fn1411 github.com/goccy/llamawasm2go/p2.Fn1411
func Fn1411(m *base.Module, l0 int64) int64

//go:linkname Fn1412 github.com/goccy/llamawasm2go/p2.Fn1412
func Fn1412(m *base.Module, l0 int64) int64

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p2.Fn1413
func Fn1413(m *base.Module, l0 int64) int64

//go:linkname Fn1414 github.com/goccy/llamawasm2go/p2.Fn1414
func Fn1414(m *base.Module, l0 int64) int64

//go:linkname Fn1415 github.com/goccy/llamawasm2go/p2.Fn1415
func Fn1415(m *base.Module, l0 int64) int64

//go:linkname Fn1417 github.com/goccy/llamawasm2go/p2.Fn1417
func Fn1417(m *base.Module, l0 int64) int64

//go:linkname Fn1418 github.com/goccy/llamawasm2go/p2.Fn1418
func Fn1418(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64) int64

//go:linkname Fn1421 github.com/goccy/llamawasm2go/p2.Fn1421
func Fn1421(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1422 github.com/goccy/llamawasm2go/p2.Fn1422
func Fn1422(m *base.Module, l0 int64) int64

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1425 github.com/goccy/llamawasm2go/p2.Fn1425
func Fn1425(m *base.Module, l0 int64) int64

//go:linkname Fn1427 github.com/goccy/llamawasm2go/p2.Fn1427
func Fn1427(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1428 github.com/goccy/llamawasm2go/p2.Fn1428
func Fn1428(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1429 github.com/goccy/llamawasm2go/p2.Fn1429
func Fn1429(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1430 github.com/goccy/llamawasm2go/p2.Fn1430
func Fn1430(m *base.Module, l0 int64) int64

//go:linkname Fn1431 github.com/goccy/llamawasm2go/p2.Fn1431
func Fn1431(m *base.Module, l0 int64) int64

//go:linkname Fn1432 github.com/goccy/llamawasm2go/p2.Fn1432
func Fn1432(m *base.Module, l0 int64) int64

//go:linkname Fn1438 github.com/goccy/llamawasm2go/p2.Fn1438
func Fn1438(m *base.Module, l0 int64) int64

//go:linkname Fn1440 github.com/goccy/llamawasm2go/p2.Fn1440
func Fn1440(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1441 github.com/goccy/llamawasm2go/p2.Fn1441
func Fn1441(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64) int64

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int64) int64

//go:linkname Fn1446 github.com/goccy/llamawasm2go/p2.Fn1446
func Fn1446(m *base.Module, l0 int64) int64

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1448 github.com/goccy/llamawasm2go/p2.Fn1448
func Fn1448(m *base.Module, l0 int64)

//go:linkname Fn1477 github.com/goccy/llamawasm2go/p2.Fn1477
func Fn1477(m *base.Module, l0 int64) int64

//go:linkname Fn1483 github.com/goccy/llamawasm2go/p2.Fn1483
func Fn1483(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1484 github.com/goccy/llamawasm2go/p2.Fn1484
func Fn1484(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1485 github.com/goccy/llamawasm2go/p2.Fn1485
func Fn1485(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1487 github.com/goccy/llamawasm2go/p2.Fn1487
func Fn1487(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1488 github.com/goccy/llamawasm2go/p2.Fn1488
func Fn1488(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1489 github.com/goccy/llamawasm2go/p2.Fn1489
func Fn1489(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1490 github.com/goccy/llamawasm2go/p2.Fn1490
func Fn1490(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1491 github.com/goccy/llamawasm2go/p2.Fn1491
func Fn1491(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1492 github.com/goccy/llamawasm2go/p2.Fn1492
func Fn1492(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1494 github.com/goccy/llamawasm2go/p2.Fn1494
func Fn1494(m *base.Module, l0 int64) int32

//go:linkname Fn1495 github.com/goccy/llamawasm2go/p2.Fn1495
func Fn1495(m *base.Module, l0 int64) int32

//go:linkname Fn1496 github.com/goccy/llamawasm2go/p2.Fn1496
func Fn1496(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1497 github.com/goccy/llamawasm2go/p2.Fn1497
func Fn1497(m *base.Module, l0 int64) int32

//go:linkname Fn1498 github.com/goccy/llamawasm2go/p2.Fn1498
func Fn1498(m *base.Module, l0 int64) int32

//go:linkname Fn1501 github.com/goccy/llamawasm2go/p2.Fn1501
func Fn1501(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1502 github.com/goccy/llamawasm2go/p2.Fn1502
func Fn1502(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1503 github.com/goccy/llamawasm2go/p2.Fn1503
func Fn1503(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1504 github.com/goccy/llamawasm2go/p2.Fn1504
func Fn1504(m *base.Module)

//go:linkname Fn1505 github.com/goccy/llamawasm2go/p2.Fn1505
func Fn1505(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1507 github.com/goccy/llamawasm2go/p2.Fn1507
func Fn1507(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1509 github.com/goccy/llamawasm2go/p2.Fn1509
func Fn1509(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1510 github.com/goccy/llamawasm2go/p2.Fn1510
func Fn1510(m *base.Module, l0 int64)

//go:linkname Fn1514 github.com/goccy/llamawasm2go/p2.Fn1514
func Fn1514(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1515 github.com/goccy/llamawasm2go/p2.Fn1515
func Fn1515(m *base.Module, l0 int64)

//go:linkname Fn1518 github.com/goccy/llamawasm2go/p2.Fn1518
func Fn1518(m *base.Module, l0 int64)

//go:linkname Fn1527 github.com/goccy/llamawasm2go/p2.Fn1527
func Fn1527(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1528 github.com/goccy/llamawasm2go/p2.Fn1528
func Fn1528(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1537 github.com/goccy/llamawasm2go/p2.Fn1537
func Fn1537(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1538 github.com/goccy/llamawasm2go/p0.Fn1538
func Fn1538(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1540 github.com/goccy/llamawasm2go/p2.Fn1540
func Fn1540(m *base.Module, l0 int64)

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p2.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1549 github.com/goccy/llamawasm2go/p2.Fn1549
func Fn1549(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1550 github.com/goccy/llamawasm2go/p2.Fn1550
func Fn1550(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1557 github.com/goccy/llamawasm2go/p2.Fn1557
func Fn1557(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1564 github.com/goccy/llamawasm2go/p2.Fn1564
func Fn1564(m *base.Module, l0 int64)

//go:linkname Fn1567 github.com/goccy/llamawasm2go/p2.Fn1567
func Fn1567(m *base.Module, l0 int64) int32

//go:linkname Fn1577 github.com/goccy/llamawasm2go/p2.Fn1577
func Fn1577(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1579 github.com/goccy/llamawasm2go/p2.Fn1579
func Fn1579(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1580 github.com/goccy/llamawasm2go/p2.Fn1580
func Fn1580(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1599 github.com/goccy/llamawasm2go/p2.Fn1599
func Fn1599(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1612 github.com/goccy/llamawasm2go/p2.Fn1612
func Fn1612(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1619 github.com/goccy/llamawasm2go/p2.Fn1619
func Fn1619(m *base.Module, l0 int64)

//go:linkname Fn1653 github.com/goccy/llamawasm2go/p2.Fn1653
func Fn1653(m *base.Module, l0 int64)

//go:linkname Fn1656 github.com/goccy/llamawasm2go/p2.Fn1656
func Fn1656(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1670 github.com/goccy/llamawasm2go/p2.Fn1670
func Fn1670(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1671 github.com/goccy/llamawasm2go/p2.Fn1671
func Fn1671(m *base.Module, l0 int64) int64

//go:linkname Fn1672 github.com/goccy/llamawasm2go/p2.Fn1672
func Fn1672(m *base.Module, l0 int64)

//go:linkname Fn1676 github.com/goccy/llamawasm2go/p0.Fn1676
func Fn1676(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1677 github.com/goccy/llamawasm2go/p2.Fn1677
func Fn1677(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1678 github.com/goccy/llamawasm2go/p2.Fn1678
func Fn1678(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1679 github.com/goccy/llamawasm2go/p2.Fn1679
func Fn1679(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1680 github.com/goccy/llamawasm2go/p2.Fn1680
func Fn1680(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1698 github.com/goccy/llamawasm2go/p2.Fn1698
func Fn1698(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1708 github.com/goccy/llamawasm2go/p2.Fn1708
func Fn1708(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1710 github.com/goccy/llamawasm2go/p2.Fn1710
func Fn1710(m *base.Module, l0 int64) int64

//go:linkname Fn1711 github.com/goccy/llamawasm2go/p2.Fn1711
func Fn1711(m *base.Module, l0 int64)

//go:linkname Fn1714 github.com/goccy/llamawasm2go/p0.Fn1714
func Fn1714(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1716 github.com/goccy/llamawasm2go/p2.Fn1716
func Fn1716(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1727 github.com/goccy/llamawasm2go/p2.Fn1727
func Fn1727(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1730 github.com/goccy/llamawasm2go/p2.Fn1730
func Fn1730(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1731 github.com/goccy/llamawasm2go/p2.Fn1731
func Fn1731(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1741 github.com/goccy/llamawasm2go/p2.Fn1741
func Fn1741(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1767 github.com/goccy/llamawasm2go/p2.Fn1767
func Fn1767(m *base.Module, l0 int64)

//go:linkname Fn1768 github.com/goccy/llamawasm2go/p2.Fn1768
func Fn1768(m *base.Module, l0 int64)

//go:linkname Fn1779 github.com/goccy/llamawasm2go/p2.Fn1779
func Fn1779(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1804 github.com/goccy/llamawasm2go/p2.Fn1804
func Fn1804(m *base.Module) int64

//go:linkname Fn1807 github.com/goccy/llamawasm2go/p2.Fn1807
func Fn1807(m *base.Module, l0 int64) int64

//go:linkname Fn1808 github.com/goccy/llamawasm2go/p2.Fn1808
func Fn1808(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1815 github.com/goccy/llamawasm2go/p2.Fn1815
func Fn1815(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1819 github.com/goccy/llamawasm2go/p2.Fn1819
func Fn1819(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1821 github.com/goccy/llamawasm2go/p2.Fn1821
func Fn1821(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1822 github.com/goccy/llamawasm2go/p2.Fn1822
func Fn1822(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1823 github.com/goccy/llamawasm2go/p2.Fn1823
func Fn1823(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1824 github.com/goccy/llamawasm2go/p2.Fn1824
func Fn1824(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1825 github.com/goccy/llamawasm2go/p2.Fn1825
func Fn1825(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1830 github.com/goccy/llamawasm2go/p2.Fn1830
func Fn1830(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1836 github.com/goccy/llamawasm2go/p2.Fn1836
func Fn1836(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1838 github.com/goccy/llamawasm2go/p2.Fn1838
func Fn1838(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1841 github.com/goccy/llamawasm2go/p2.Fn1841
func Fn1841(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1842 github.com/goccy/llamawasm2go/p2.Fn1842
func Fn1842(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1843 github.com/goccy/llamawasm2go/p2.Fn1843
func Fn1843(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1846 github.com/goccy/llamawasm2go/p2.Fn1846
func Fn1846(m *base.Module, l0 int64)

//go:linkname Fn1856 github.com/goccy/llamawasm2go/p2.Fn1856
func Fn1856(m *base.Module, l0 int64)

//go:linkname Fn1858 github.com/goccy/llamawasm2go/p2.Fn1858
func Fn1858(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1859 github.com/goccy/llamawasm2go/p2.Fn1859
func Fn1859(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1872 github.com/goccy/llamawasm2go/p2.Fn1872
func Fn1872(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1873 github.com/goccy/llamawasm2go/p2.Fn1873
func Fn1873(m *base.Module, l0 int64) int64

//go:linkname Fn1874 github.com/goccy/llamawasm2go/p2.Fn1874
func Fn1874(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1876 github.com/goccy/llamawasm2go/p2.Fn1876
func Fn1876(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1888 github.com/goccy/llamawasm2go/p2.Fn1888
func Fn1888(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1890 github.com/goccy/llamawasm2go/p2.Fn1890
func Fn1890(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1891 github.com/goccy/llamawasm2go/p2.Fn1891
func Fn1891(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1893 github.com/goccy/llamawasm2go/p2.Fn1893
func Fn1893(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1894 github.com/goccy/llamawasm2go/p2.Fn1894
func Fn1894(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1895 github.com/goccy/llamawasm2go/p2.Fn1895
func Fn1895(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1896 github.com/goccy/llamawasm2go/p2.Fn1896
func Fn1896(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1897 github.com/goccy/llamawasm2go/p2.Fn1897
func Fn1897(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn1898 github.com/goccy/llamawasm2go/p2.Fn1898
func Fn1898(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1917 github.com/goccy/llamawasm2go/p2.Fn1917
func Fn1917(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1918 github.com/goccy/llamawasm2go/p2.Fn1918
func Fn1918(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1919 github.com/goccy/llamawasm2go/p2.Fn1919
func Fn1919(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1921 github.com/goccy/llamawasm2go/p2.Fn1921
func Fn1921(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1922 github.com/goccy/llamawasm2go/p2.Fn1922
func Fn1922(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1923 github.com/goccy/llamawasm2go/p2.Fn1923
func Fn1923(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1924 github.com/goccy/llamawasm2go/p2.Fn1924
func Fn1924(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1925 github.com/goccy/llamawasm2go/p2.Fn1925
func Fn1925(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1927 github.com/goccy/llamawasm2go/p2.Fn1927
func Fn1927(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1929 github.com/goccy/llamawasm2go/p2.Fn1929
func Fn1929(m *base.Module, l0 int64)

//go:linkname Fn1946 github.com/goccy/llamawasm2go/p2.Fn1946
func Fn1946(m *base.Module, l0 int64)

//go:linkname Fn1947 github.com/goccy/llamawasm2go/p2.Fn1947
func Fn1947(m *base.Module, l0 int64)

//go:linkname Fn1948 github.com/goccy/llamawasm2go/p2.Fn1948
func Fn1948(m *base.Module, l0 int64)

//go:linkname Fn1950 github.com/goccy/llamawasm2go/p2.Fn1950
func Fn1950(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1951 github.com/goccy/llamawasm2go/p2.Fn1951
func Fn1951(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1990 github.com/goccy/llamawasm2go/p2.Fn1990
func Fn1990(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2002 github.com/goccy/llamawasm2go/p2.Fn2002
func Fn2002(m *base.Module, l0 int64) int64

//go:linkname Fn2005 github.com/goccy/llamawasm2go/p2.Fn2005
func Fn2005(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2009 github.com/goccy/llamawasm2go/p2.Fn2009
func Fn2009(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2013 github.com/goccy/llamawasm2go/p2.Fn2013
func Fn2013(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2026 github.com/goccy/llamawasm2go/p2.Fn2026
func Fn2026(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2027 github.com/goccy/llamawasm2go/p2.Fn2027
func Fn2027(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2028 github.com/goccy/llamawasm2go/p2.Fn2028
func Fn2028(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2040 github.com/goccy/llamawasm2go/p2.Fn2040
func Fn2040(m *base.Module, l0 int64) int64

//go:linkname Fn2041 github.com/goccy/llamawasm2go/p2.Fn2041
func Fn2041(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2044 github.com/goccy/llamawasm2go/p2.Fn2044
func Fn2044(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2046 github.com/goccy/llamawasm2go/p2.Fn2046
func Fn2046(m *base.Module)

//go:linkname Fn2049 github.com/goccy/llamawasm2go/p2.Fn2049
func Fn2049(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2058 github.com/goccy/llamawasm2go/p2.Fn2058
func Fn2058(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2062 github.com/goccy/llamawasm2go/p2.Fn2062
func Fn2062(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2063 github.com/goccy/llamawasm2go/p2.Fn2063
func Fn2063(m *base.Module, l0 int64)

//go:linkname Fn2064 github.com/goccy/llamawasm2go/p2.Fn2064
func Fn2064(m *base.Module, l0 int64)

//go:linkname Fn2065 github.com/goccy/llamawasm2go/p2.Fn2065
func Fn2065(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2067 github.com/goccy/llamawasm2go/p2.Fn2067
func Fn2067(m *base.Module, l0 int64)

//go:linkname Fn2068 github.com/goccy/llamawasm2go/p2.Fn2068
func Fn2068(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2069 github.com/goccy/llamawasm2go/p2.Fn2069
func Fn2069(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2073 github.com/goccy/llamawasm2go/p2.Fn2073
func Fn2073(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2076 github.com/goccy/llamawasm2go/p2.Fn2076
func Fn2076(m *base.Module, l0 int64) int64

//go:linkname Fn2077 github.com/goccy/llamawasm2go/p2.Fn2077
func Fn2077(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2078 github.com/goccy/llamawasm2go/p2.Fn2078
func Fn2078(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2085 github.com/goccy/llamawasm2go/p2.Fn2085
func Fn2085(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2086 github.com/goccy/llamawasm2go/p2.Fn2086
func Fn2086(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2088 github.com/goccy/llamawasm2go/p2.Fn2088
func Fn2088(m *base.Module) int64

//go:linkname Fn2090 github.com/goccy/llamawasm2go/p2.Fn2090
func Fn2090(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2091 github.com/goccy/llamawasm2go/p2.Fn2091
func Fn2091(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2092 github.com/goccy/llamawasm2go/p2.Fn2092
func Fn2092(m *base.Module) int64

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 int32) int64

//go:linkname Fn2095 github.com/goccy/llamawasm2go/p2.Fn2095
func Fn2095(m *base.Module, l0 int32) int32

//go:linkname Fn2096 github.com/goccy/llamawasm2go/p2.Fn2096
func Fn2096(m *base.Module, l0 int32) int64

//go:linkname Fn2097 github.com/goccy/llamawasm2go/p2.Fn2097
func Fn2097(m *base.Module, l0 float32) int64

//go:linkname Fn2098 github.com/goccy/llamawasm2go/p2.Fn2098
func Fn2098(m *base.Module, l0 float32) int64

//go:linkname Fn2099 github.com/goccy/llamawasm2go/p2.Fn2099
func Fn2099(m *base.Module, l0 float32) int64

//go:linkname Fn2101 github.com/goccy/llamawasm2go/p2.Fn2101
func Fn2101(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2102 github.com/goccy/llamawasm2go/p2.Fn2102
func Fn2102(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2143 github.com/goccy/llamawasm2go/p2.Fn2143
func Fn2143(m *base.Module, l0 int64)

//go:linkname Fn2145 github.com/goccy/llamawasm2go/p2.Fn2145
func Fn2145(m *base.Module, l0 int64)

//go:linkname Fn2153 github.com/goccy/llamawasm2go/p2.Fn2153
func Fn2153(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2187 github.com/goccy/llamawasm2go/p2.Fn2187
func Fn2187(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2198 github.com/goccy/llamawasm2go/p2.Fn2198
func Fn2198(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2201 github.com/goccy/llamawasm2go/p2.Fn2201
func Fn2201(m *base.Module, l0 int64)

//go:linkname Fn2202 github.com/goccy/llamawasm2go/p2.Fn2202
func Fn2202(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2203 github.com/goccy/llamawasm2go/p2.Fn2203
func Fn2203(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2204 github.com/goccy/llamawasm2go/p2.Fn2204
func Fn2204(m *base.Module, l0 int64)

//go:linkname Fn2207 github.com/goccy/llamawasm2go/p2.Fn2207
func Fn2207(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2212 github.com/goccy/llamawasm2go/p2.Fn2212
func Fn2212(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2213 github.com/goccy/llamawasm2go/p2.Fn2213
func Fn2213(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2227 github.com/goccy/llamawasm2go/p2.Fn2227
func Fn2227(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2237 github.com/goccy/llamawasm2go/p2.Fn2237
func Fn2237(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2240 github.com/goccy/llamawasm2go/p2.Fn2240
func Fn2240(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2242 github.com/goccy/llamawasm2go/p2.Fn2242
func Fn2242(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2243 github.com/goccy/llamawasm2go/p2.Fn2243
func Fn2243(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2244 github.com/goccy/llamawasm2go/p2.Fn2244
func Fn2244(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2247 github.com/goccy/llamawasm2go/p2.Fn2247
func Fn2247(m *base.Module)

//go:linkname Fn2251 github.com/goccy/llamawasm2go/p2.Fn2251
func Fn2251(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2253 github.com/goccy/llamawasm2go/p0.Fn2253
func Fn2253(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 int64)

//go:linkname Fn2269 github.com/goccy/llamawasm2go/p2.Fn2269
func Fn2269(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2274 github.com/goccy/llamawasm2go/p2.Fn2274
func Fn2274(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2279 github.com/goccy/llamawasm2go/p2.Fn2279
func Fn2279(m *base.Module, l0 int64) int64

//go:linkname Fn2280 github.com/goccy/llamawasm2go/p2.Fn2280
func Fn2280(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2281 github.com/goccy/llamawasm2go/p0.Fn2281
func Fn2281(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2292 github.com/goccy/llamawasm2go/p2.Fn2292
func Fn2292(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2293 github.com/goccy/llamawasm2go/p2.Fn2293
func Fn2293(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2294 github.com/goccy/llamawasm2go/p2.Fn2294
func Fn2294(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2295 github.com/goccy/llamawasm2go/p2.Fn2295
func Fn2295(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2315 github.com/goccy/llamawasm2go/p2.Fn2315
func Fn2315(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2316 github.com/goccy/llamawasm2go/p2.Fn2316
func Fn2316(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2319 github.com/goccy/llamawasm2go/p2.Fn2319
func Fn2319(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2405 github.com/goccy/llamawasm2go/p2.Fn2405
func Fn2405(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2524 github.com/goccy/llamawasm2go/p2.Fn2524
func Fn2524(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2525 github.com/goccy/llamawasm2go/p0.Fn2525
func Fn2525(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2526 github.com/goccy/llamawasm2go/p2.Fn2526
func Fn2526(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2527 github.com/goccy/llamawasm2go/p2.Fn2527
func Fn2527(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2531 github.com/goccy/llamawasm2go/p2.Fn2531
func Fn2531(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2556 github.com/goccy/llamawasm2go/p2.Fn2556
func Fn2556(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2688 github.com/goccy/llamawasm2go/p2.Fn2688
func Fn2688(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2723 github.com/goccy/llamawasm2go/p2.Fn2723
func Fn2723(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2737 github.com/goccy/llamawasm2go/p2.Fn2737
func Fn2737(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2747 github.com/goccy/llamawasm2go/p2.Fn2747
func Fn2747(m *base.Module, l0 int32)

//go:linkname Fn2749 github.com/goccy/llamawasm2go/p2.Fn2749
func Fn2749(m *base.Module, l0 int64) int64

//go:linkname Fn2750 github.com/goccy/llamawasm2go/p2.Fn2750
func Fn2750(m *base.Module, l0 int64)

//go:linkname Fn2753 github.com/goccy/llamawasm2go/p2.Fn2753
func Fn2753(m *base.Module, l0 int64)

//go:linkname Fn2754 github.com/goccy/llamawasm2go/p2.Fn2754
func Fn2754(m *base.Module, l0 int64)

//go:linkname Fn2756 github.com/goccy/llamawasm2go/p2.Fn2756
func Fn2756(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2761 github.com/goccy/llamawasm2go/p2.Fn2761
func Fn2761(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2767 github.com/goccy/llamawasm2go/p2.Fn2767
func Fn2767(m *base.Module) int32

//go:linkname Fn2775 github.com/goccy/llamawasm2go/p2.Fn2775
func Fn2775(m *base.Module, l0 float64) float32

//go:linkname Fn2776 github.com/goccy/llamawasm2go/p2.Fn2776
func Fn2776(m *base.Module, l0 float64) float32

//go:linkname Fn2783 github.com/goccy/llamawasm2go/p2.Fn2783
func Fn2783(m *base.Module, l0 float32) float32

//go:linkname Fn2787 github.com/goccy/llamawasm2go/p2.Fn2787
func Fn2787(m *base.Module, l0 float32) float32

//go:linkname Fn2790 github.com/goccy/llamawasm2go/p2.Fn2790
func Fn2790(m *base.Module, l0 float32) float32

//go:linkname Fn2805 github.com/goccy/llamawasm2go/p2.Fn2805
func Fn2805(m *base.Module, l0 int64) int32

//go:linkname Fn2806 github.com/goccy/llamawasm2go/p2.Fn2806
func Fn2806(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2808 github.com/goccy/llamawasm2go/p2.Fn2808
func Fn2808(m *base.Module, l0 int64)

//go:linkname Fn2809 github.com/goccy/llamawasm2go/p2.Fn2809
func Fn2809(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2810 github.com/goccy/llamawasm2go/p2.Fn2810
func Fn2810(m *base.Module, l0 int64) int32

//go:linkname Fn2817 github.com/goccy/llamawasm2go/p2.Fn2817
func Fn2817(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2819 github.com/goccy/llamawasm2go/p2.Fn2819
func Fn2819(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2825 github.com/goccy/llamawasm2go/p2.Fn2825
func Fn2825(m *base.Module, l0 int64) int32

//go:linkname Fn2826 github.com/goccy/llamawasm2go/p2.Fn2826
func Fn2826(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2831 github.com/goccy/llamawasm2go/p2.Fn2831
func Fn2831(m *base.Module, l0 int64) int32

//go:linkname Fn2835 github.com/goccy/llamawasm2go/p2.Fn2835
func Fn2835(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn2836 github.com/goccy/llamawasm2go/p2.Fn2836
func Fn2836(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2839 github.com/goccy/llamawasm2go/p2.Fn2839
func Fn2839(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2842 github.com/goccy/llamawasm2go/p2.Fn2842
func Fn2842(m *base.Module, l0 int64) int64

//go:linkname Fn2846 github.com/goccy/llamawasm2go/p2.Fn2846
func Fn2846(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn2856 github.com/goccy/llamawasm2go/p2.Fn2856
func Fn2856(m *base.Module)

//go:linkname Fn2857 github.com/goccy/llamawasm2go/p0.Fn2857
func Fn2857(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2859 github.com/goccy/llamawasm2go/p2.Fn2859
func Fn2859(m *base.Module)

//go:linkname Fn2861 github.com/goccy/llamawasm2go/p0.Fn2861
func Fn2861(m *base.Module, l0 int64) int64

//go:linkname Fn2863 github.com/goccy/llamawasm2go/p2.Fn2863
func Fn2863(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2864 github.com/goccy/llamawasm2go/p2.Fn2864
func Fn2864(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2867 github.com/goccy/llamawasm2go/p2.Fn2867
func Fn2867(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2896 github.com/goccy/llamawasm2go/p2.Fn2896
func Fn2896(m *base.Module, l0 int32)
