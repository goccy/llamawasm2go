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

//go:linkname Fn858 github.com/goccy/llamawasm2go/p2.Fn858
func Fn858(m *base.Module) int64

//go:linkname Fn895 github.com/goccy/llamawasm2go/p2.Fn895
func Fn895(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32)

//go:linkname Fn904 github.com/goccy/llamawasm2go/p2.Fn904
func Fn904(m *base.Module, l0 int64)

//go:linkname Fn911 github.com/goccy/llamawasm2go/p2.Fn911
func Fn911(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn913 github.com/goccy/llamawasm2go/p2.Fn913
func Fn913(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn928 github.com/goccy/llamawasm2go/p2.Fn928
func Fn928(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn939 github.com/goccy/llamawasm2go/p2.Fn939
func Fn939(m *base.Module, l0 int64)

//go:linkname Fn944 github.com/goccy/llamawasm2go/p2.Fn944
func Fn944(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn989 github.com/goccy/llamawasm2go/p0.Fn989
func Fn989(m *base.Module, l0 int64) int64

//go:linkname Fn991 github.com/goccy/llamawasm2go/p2.Fn991
func Fn991(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1068 github.com/goccy/llamawasm2go/p2.Fn1068
func Fn1068(m *base.Module, l0 int64)

//go:linkname Fn1090 github.com/goccy/llamawasm2go/p2.Fn1090
func Fn1090(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1098 github.com/goccy/llamawasm2go/p2.Fn1098
func Fn1098(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1108 github.com/goccy/llamawasm2go/p2.Fn1108
func Fn1108(m *base.Module, l0 int64) int64

//go:linkname Fn1138 github.com/goccy/llamawasm2go/p2.Fn1138
func Fn1138(m *base.Module, l0 int32) int64

//go:linkname Fn1149 github.com/goccy/llamawasm2go/p2.Fn1149
func Fn1149(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1150 github.com/goccy/llamawasm2go/p2.Fn1150
func Fn1150(m *base.Module, l0 int64)

//go:linkname Fn1152 github.com/goccy/llamawasm2go/p2.Fn1152
func Fn1152(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1172 github.com/goccy/llamawasm2go/p2.Fn1172
func Fn1172(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1173 github.com/goccy/llamawasm2go/p2.Fn1173
func Fn1173(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1174 github.com/goccy/llamawasm2go/p2.Fn1174
func Fn1174(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1175 github.com/goccy/llamawasm2go/p2.Fn1175
func Fn1175(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1176 github.com/goccy/llamawasm2go/p2.Fn1176
func Fn1176(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1177 github.com/goccy/llamawasm2go/p2.Fn1177
func Fn1177(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1178 github.com/goccy/llamawasm2go/p2.Fn1178
func Fn1178(m *base.Module, l0 int64) int64

//go:linkname Fn1181 github.com/goccy/llamawasm2go/p2.Fn1181
func Fn1181(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1185 github.com/goccy/llamawasm2go/p2.Fn1185
func Fn1185(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1186 github.com/goccy/llamawasm2go/p2.Fn1186
func Fn1186(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1187 github.com/goccy/llamawasm2go/p2.Fn1187
func Fn1187(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1188 github.com/goccy/llamawasm2go/p2.Fn1188
func Fn1188(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1189 github.com/goccy/llamawasm2go/p2.Fn1189
func Fn1189(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1190 github.com/goccy/llamawasm2go/p2.Fn1190
func Fn1190(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1191 github.com/goccy/llamawasm2go/p2.Fn1191
func Fn1191(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1192 github.com/goccy/llamawasm2go/p2.Fn1192
func Fn1192(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1193 github.com/goccy/llamawasm2go/p2.Fn1193
func Fn1193(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1194 github.com/goccy/llamawasm2go/p2.Fn1194
func Fn1194(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1195 github.com/goccy/llamawasm2go/p2.Fn1195
func Fn1195(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1196 github.com/goccy/llamawasm2go/p2.Fn1196
func Fn1196(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1197 github.com/goccy/llamawasm2go/p2.Fn1197
func Fn1197(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1198 github.com/goccy/llamawasm2go/p2.Fn1198
func Fn1198(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1199 github.com/goccy/llamawasm2go/p2.Fn1199
func Fn1199(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1200 github.com/goccy/llamawasm2go/p2.Fn1200
func Fn1200(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1201 github.com/goccy/llamawasm2go/p2.Fn1201
func Fn1201(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1202 github.com/goccy/llamawasm2go/p2.Fn1202
func Fn1202(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1203 github.com/goccy/llamawasm2go/p2.Fn1203
func Fn1203(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1204 github.com/goccy/llamawasm2go/p2.Fn1204
func Fn1204(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1205 github.com/goccy/llamawasm2go/p2.Fn1205
func Fn1205(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1206 github.com/goccy/llamawasm2go/p2.Fn1206
func Fn1206(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1207 github.com/goccy/llamawasm2go/p2.Fn1207
func Fn1207(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1208 github.com/goccy/llamawasm2go/p2.Fn1208
func Fn1208(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1209 github.com/goccy/llamawasm2go/p2.Fn1209
func Fn1209(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1210 github.com/goccy/llamawasm2go/p2.Fn1210
func Fn1210(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1211 github.com/goccy/llamawasm2go/p2.Fn1211
func Fn1211(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1213 github.com/goccy/llamawasm2go/p2.Fn1213
func Fn1213(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1217 github.com/goccy/llamawasm2go/p2.Fn1217
func Fn1217(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1218 github.com/goccy/llamawasm2go/p2.Fn1218
func Fn1218(m *base.Module)

//go:linkname Fn1219 github.com/goccy/llamawasm2go/p2.Fn1219
func Fn1219(m *base.Module)

//go:linkname Fn1220 github.com/goccy/llamawasm2go/p0.Fn1220
func Fn1220(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1228 github.com/goccy/llamawasm2go/p2.Fn1228
func Fn1228(m *base.Module)

//go:linkname Fn1230 github.com/goccy/llamawasm2go/p2.Fn1230
func Fn1230(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1238 github.com/goccy/llamawasm2go/p2.Fn1238
func Fn1238(m *base.Module, l0 int64)

//go:linkname Fn1245 github.com/goccy/llamawasm2go/p2.Fn1245
func Fn1245(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1248 github.com/goccy/llamawasm2go/p2.Fn1248
func Fn1248(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1250 github.com/goccy/llamawasm2go/p2.Fn1250
func Fn1250(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1252 github.com/goccy/llamawasm2go/p2.Fn1252
func Fn1252(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1254 github.com/goccy/llamawasm2go/p2.Fn1254
func Fn1254(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1255 github.com/goccy/llamawasm2go/p2.Fn1255
func Fn1255(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1256 github.com/goccy/llamawasm2go/p2.Fn1256
func Fn1256(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1261 github.com/goccy/llamawasm2go/p2.Fn1261
func Fn1261(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1262 github.com/goccy/llamawasm2go/p2.Fn1262
func Fn1262(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1263 github.com/goccy/llamawasm2go/p2.Fn1263
func Fn1263(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1264 github.com/goccy/llamawasm2go/p0.Fn1264
func Fn1264(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1266 github.com/goccy/llamawasm2go/p2.Fn1266
func Fn1266(m *base.Module, l0 int64)

//go:linkname Fn1267 github.com/goccy/llamawasm2go/p2.Fn1267
func Fn1267(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1269 github.com/goccy/llamawasm2go/p2.Fn1269
func Fn1269(m *base.Module, l0 int64) int64

//go:linkname Fn1270 github.com/goccy/llamawasm2go/p2.Fn1270
func Fn1270(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1271 github.com/goccy/llamawasm2go/p2.Fn1271
func Fn1271(m *base.Module, l0 int64)

//go:linkname Fn1272 github.com/goccy/llamawasm2go/p2.Fn1272
func Fn1272(m *base.Module, l0 int64)

//go:linkname Fn1273 github.com/goccy/llamawasm2go/p2.Fn1273
func Fn1273(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1274 github.com/goccy/llamawasm2go/p2.Fn1274
func Fn1274(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1277 github.com/goccy/llamawasm2go/p2.Fn1277
func Fn1277(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1284 github.com/goccy/llamawasm2go/p2.Fn1284
func Fn1284(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1286 github.com/goccy/llamawasm2go/p0.Fn1286
func Fn1286(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1288 github.com/goccy/llamawasm2go/p2.Fn1288
func Fn1288(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1289 github.com/goccy/llamawasm2go/p2.Fn1289
func Fn1289(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1291 github.com/goccy/llamawasm2go/p0.Fn1291
func Fn1291(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1292 github.com/goccy/llamawasm2go/p2.Fn1292
func Fn1292(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1296 github.com/goccy/llamawasm2go/p2.Fn1296
func Fn1296(m *base.Module, l0 int64)

//go:linkname Fn1299 github.com/goccy/llamawasm2go/p2.Fn1299
func Fn1299(m *base.Module, l0 int64)

//go:linkname Fn1301 github.com/goccy/llamawasm2go/p2.Fn1301
func Fn1301(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1302 github.com/goccy/llamawasm2go/p2.Fn1302
func Fn1302(m *base.Module, l0 int64) int64

//go:linkname Fn1303 github.com/goccy/llamawasm2go/p2.Fn1303
func Fn1303(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1304 github.com/goccy/llamawasm2go/p2.Fn1304
func Fn1304(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1305 github.com/goccy/llamawasm2go/p2.Fn1305
func Fn1305(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1306 github.com/goccy/llamawasm2go/p2.Fn1306
func Fn1306(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1307 github.com/goccy/llamawasm2go/p2.Fn1307
func Fn1307(m *base.Module, l0 int64)

//go:linkname Fn1310 github.com/goccy/llamawasm2go/p2.Fn1310
func Fn1310(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1312 github.com/goccy/llamawasm2go/p2.Fn1312
func Fn1312(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1315 github.com/goccy/llamawasm2go/p2.Fn1315
func Fn1315(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1322 github.com/goccy/llamawasm2go/p2.Fn1322
func Fn1322(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1329 github.com/goccy/llamawasm2go/p2.Fn1329
func Fn1329(m *base.Module, l0 int64)

//go:linkname Fn1330 github.com/goccy/llamawasm2go/p2.Fn1330
func Fn1330(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1332 github.com/goccy/llamawasm2go/p2.Fn1332
func Fn1332(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1333 github.com/goccy/llamawasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1376 github.com/goccy/llamawasm2go/p2.Fn1376
func Fn1376(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p2.Fn1392
func Fn1392(m *base.Module, l0 int64)

//go:linkname Fn1393 github.com/goccy/llamawasm2go/p2.Fn1393
func Fn1393(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p2.Fn1394
func Fn1394(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p2.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1396 github.com/goccy/llamawasm2go/p2.Fn1396
func Fn1396(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p2.Fn1397
func Fn1397(m *base.Module, l0 int64)

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p2.Fn1401
func Fn1401(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p2.Fn1402
func Fn1402(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p2.Fn1406
func Fn1406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p2.Fn1408
func Fn1408(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1409 github.com/goccy/llamawasm2go/p2.Fn1409
func Fn1409(m *base.Module, l0 int64) int64

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

//go:linkname Fn1416 github.com/goccy/llamawasm2go/p2.Fn1416
func Fn1416(m *base.Module, l0 int64) int64

//go:linkname Fn1417 github.com/goccy/llamawasm2go/p2.Fn1417
func Fn1417(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1419 github.com/goccy/llamawasm2go/p2.Fn1419
func Fn1419(m *base.Module, l0 int64) int64

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1421 github.com/goccy/llamawasm2go/p2.Fn1421
func Fn1421(m *base.Module, l0 int64) int64

//go:linkname Fn1423 github.com/goccy/llamawasm2go/p2.Fn1423
func Fn1423(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64) int64

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p2.Fn1426
func Fn1426(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1427 github.com/goccy/llamawasm2go/p2.Fn1427
func Fn1427(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1428 github.com/goccy/llamawasm2go/p2.Fn1428
func Fn1428(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1429 github.com/goccy/llamawasm2go/p2.Fn1429
func Fn1429(m *base.Module, l0 int64) int64

//go:linkname Fn1430 github.com/goccy/llamawasm2go/p2.Fn1430
func Fn1430(m *base.Module, l0 int64) int64

//go:linkname Fn1431 github.com/goccy/llamawasm2go/p2.Fn1431
func Fn1431(m *base.Module, l0 int64) int64

//go:linkname Fn1437 github.com/goccy/llamawasm2go/p2.Fn1437
func Fn1437(m *base.Module, l0 int64) int64

//go:linkname Fn1439 github.com/goccy/llamawasm2go/p2.Fn1439
func Fn1439(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1440 github.com/goccy/llamawasm2go/p2.Fn1440
func Fn1440(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64) int64

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64) int64

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int64) int64

//go:linkname Fn1446 github.com/goccy/llamawasm2go/p2.Fn1446
func Fn1446(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int64)

//go:linkname Fn1476 github.com/goccy/llamawasm2go/p2.Fn1476
func Fn1476(m *base.Module, l0 int64) int64

//go:linkname Fn1482 github.com/goccy/llamawasm2go/p2.Fn1482
func Fn1482(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1483 github.com/goccy/llamawasm2go/p2.Fn1483
func Fn1483(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1484 github.com/goccy/llamawasm2go/p2.Fn1484
func Fn1484(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1486 github.com/goccy/llamawasm2go/p2.Fn1486
func Fn1486(m *base.Module, l0 int64, l1 int32) int32

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

//go:linkname Fn1493 github.com/goccy/llamawasm2go/p2.Fn1493
func Fn1493(m *base.Module, l0 int64) int32

//go:linkname Fn1494 github.com/goccy/llamawasm2go/p2.Fn1494
func Fn1494(m *base.Module, l0 int64) int32

//go:linkname Fn1495 github.com/goccy/llamawasm2go/p2.Fn1495
func Fn1495(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1496 github.com/goccy/llamawasm2go/p2.Fn1496
func Fn1496(m *base.Module, l0 int64) int32

//go:linkname Fn1497 github.com/goccy/llamawasm2go/p2.Fn1497
func Fn1497(m *base.Module, l0 int64) int32

//go:linkname Fn1500 github.com/goccy/llamawasm2go/p2.Fn1500
func Fn1500(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1501 github.com/goccy/llamawasm2go/p2.Fn1501
func Fn1501(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1502 github.com/goccy/llamawasm2go/p2.Fn1502
func Fn1502(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1503 github.com/goccy/llamawasm2go/p2.Fn1503
func Fn1503(m *base.Module)

//go:linkname Fn1504 github.com/goccy/llamawasm2go/p2.Fn1504
func Fn1504(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1506 github.com/goccy/llamawasm2go/p2.Fn1506
func Fn1506(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1508 github.com/goccy/llamawasm2go/p2.Fn1508
func Fn1508(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1509 github.com/goccy/llamawasm2go/p2.Fn1509
func Fn1509(m *base.Module, l0 int64)

//go:linkname Fn1513 github.com/goccy/llamawasm2go/p2.Fn1513
func Fn1513(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1514 github.com/goccy/llamawasm2go/p2.Fn1514
func Fn1514(m *base.Module, l0 int64)

//go:linkname Fn1517 github.com/goccy/llamawasm2go/p2.Fn1517
func Fn1517(m *base.Module, l0 int64)

//go:linkname Fn1526 github.com/goccy/llamawasm2go/p2.Fn1526
func Fn1526(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1527 github.com/goccy/llamawasm2go/p2.Fn1527
func Fn1527(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1536 github.com/goccy/llamawasm2go/p2.Fn1536
func Fn1536(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1537 github.com/goccy/llamawasm2go/p0.Fn1537
func Fn1537(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1539 github.com/goccy/llamawasm2go/p2.Fn1539
func Fn1539(m *base.Module, l0 int64)

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p2.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1548 github.com/goccy/llamawasm2go/p2.Fn1548
func Fn1548(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1549 github.com/goccy/llamawasm2go/p2.Fn1549
func Fn1549(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1556 github.com/goccy/llamawasm2go/p2.Fn1556
func Fn1556(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p2.Fn1563
func Fn1563(m *base.Module, l0 int64)

//go:linkname Fn1566 github.com/goccy/llamawasm2go/p2.Fn1566
func Fn1566(m *base.Module, l0 int64) int32

//go:linkname Fn1576 github.com/goccy/llamawasm2go/p2.Fn1576
func Fn1576(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1578 github.com/goccy/llamawasm2go/p2.Fn1578
func Fn1578(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1579 github.com/goccy/llamawasm2go/p2.Fn1579
func Fn1579(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1598 github.com/goccy/llamawasm2go/p2.Fn1598
func Fn1598(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1611 github.com/goccy/llamawasm2go/p2.Fn1611
func Fn1611(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1618 github.com/goccy/llamawasm2go/p2.Fn1618
func Fn1618(m *base.Module, l0 int64)

//go:linkname Fn1652 github.com/goccy/llamawasm2go/p2.Fn1652
func Fn1652(m *base.Module, l0 int64)

//go:linkname Fn1655 github.com/goccy/llamawasm2go/p2.Fn1655
func Fn1655(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1669 github.com/goccy/llamawasm2go/p2.Fn1669
func Fn1669(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1670 github.com/goccy/llamawasm2go/p2.Fn1670
func Fn1670(m *base.Module, l0 int64) int64

//go:linkname Fn1671 github.com/goccy/llamawasm2go/p2.Fn1671
func Fn1671(m *base.Module, l0 int64)

//go:linkname Fn1675 github.com/goccy/llamawasm2go/p0.Fn1675
func Fn1675(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1676 github.com/goccy/llamawasm2go/p2.Fn1676
func Fn1676(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1677 github.com/goccy/llamawasm2go/p2.Fn1677
func Fn1677(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1678 github.com/goccy/llamawasm2go/p2.Fn1678
func Fn1678(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1679 github.com/goccy/llamawasm2go/p2.Fn1679
func Fn1679(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1697 github.com/goccy/llamawasm2go/p2.Fn1697
func Fn1697(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1707 github.com/goccy/llamawasm2go/p2.Fn1707
func Fn1707(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1709 github.com/goccy/llamawasm2go/p2.Fn1709
func Fn1709(m *base.Module, l0 int64) int64

//go:linkname Fn1710 github.com/goccy/llamawasm2go/p2.Fn1710
func Fn1710(m *base.Module, l0 int64)

//go:linkname Fn1713 github.com/goccy/llamawasm2go/p0.Fn1713
func Fn1713(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1715 github.com/goccy/llamawasm2go/p2.Fn1715
func Fn1715(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1726 github.com/goccy/llamawasm2go/p2.Fn1726
func Fn1726(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1729 github.com/goccy/llamawasm2go/p2.Fn1729
func Fn1729(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1730 github.com/goccy/llamawasm2go/p2.Fn1730
func Fn1730(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1740 github.com/goccy/llamawasm2go/p2.Fn1740
func Fn1740(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1766 github.com/goccy/llamawasm2go/p2.Fn1766
func Fn1766(m *base.Module, l0 int64)

//go:linkname Fn1767 github.com/goccy/llamawasm2go/p2.Fn1767
func Fn1767(m *base.Module, l0 int64)

//go:linkname Fn1778 github.com/goccy/llamawasm2go/p2.Fn1778
func Fn1778(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1803 github.com/goccy/llamawasm2go/p2.Fn1803
func Fn1803(m *base.Module) int64

//go:linkname Fn1806 github.com/goccy/llamawasm2go/p2.Fn1806
func Fn1806(m *base.Module, l0 int64) int64

//go:linkname Fn1807 github.com/goccy/llamawasm2go/p2.Fn1807
func Fn1807(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1814 github.com/goccy/llamawasm2go/p2.Fn1814
func Fn1814(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1818 github.com/goccy/llamawasm2go/p2.Fn1818
func Fn1818(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1820 github.com/goccy/llamawasm2go/p2.Fn1820
func Fn1820(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1821 github.com/goccy/llamawasm2go/p2.Fn1821
func Fn1821(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1822 github.com/goccy/llamawasm2go/p2.Fn1822
func Fn1822(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1823 github.com/goccy/llamawasm2go/p2.Fn1823
func Fn1823(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1824 github.com/goccy/llamawasm2go/p2.Fn1824
func Fn1824(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1829 github.com/goccy/llamawasm2go/p2.Fn1829
func Fn1829(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1835 github.com/goccy/llamawasm2go/p2.Fn1835
func Fn1835(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1837 github.com/goccy/llamawasm2go/p2.Fn1837
func Fn1837(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1840 github.com/goccy/llamawasm2go/p2.Fn1840
func Fn1840(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1841 github.com/goccy/llamawasm2go/p2.Fn1841
func Fn1841(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1842 github.com/goccy/llamawasm2go/p2.Fn1842
func Fn1842(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1845 github.com/goccy/llamawasm2go/p2.Fn1845
func Fn1845(m *base.Module, l0 int64)

//go:linkname Fn1855 github.com/goccy/llamawasm2go/p2.Fn1855
func Fn1855(m *base.Module, l0 int64)

//go:linkname Fn1857 github.com/goccy/llamawasm2go/p2.Fn1857
func Fn1857(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1858 github.com/goccy/llamawasm2go/p2.Fn1858
func Fn1858(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1871 github.com/goccy/llamawasm2go/p2.Fn1871
func Fn1871(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1872 github.com/goccy/llamawasm2go/p2.Fn1872
func Fn1872(m *base.Module, l0 int64) int64

//go:linkname Fn1873 github.com/goccy/llamawasm2go/p2.Fn1873
func Fn1873(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1875 github.com/goccy/llamawasm2go/p2.Fn1875
func Fn1875(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1887 github.com/goccy/llamawasm2go/p2.Fn1887
func Fn1887(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1889 github.com/goccy/llamawasm2go/p2.Fn1889
func Fn1889(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1890 github.com/goccy/llamawasm2go/p2.Fn1890
func Fn1890(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1892 github.com/goccy/llamawasm2go/p2.Fn1892
func Fn1892(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1893 github.com/goccy/llamawasm2go/p2.Fn1893
func Fn1893(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1894 github.com/goccy/llamawasm2go/p2.Fn1894
func Fn1894(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1895 github.com/goccy/llamawasm2go/p2.Fn1895
func Fn1895(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1896 github.com/goccy/llamawasm2go/p2.Fn1896
func Fn1896(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn1897 github.com/goccy/llamawasm2go/p2.Fn1897
func Fn1897(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1916 github.com/goccy/llamawasm2go/p2.Fn1916
func Fn1916(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1917 github.com/goccy/llamawasm2go/p2.Fn1917
func Fn1917(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1918 github.com/goccy/llamawasm2go/p2.Fn1918
func Fn1918(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1920 github.com/goccy/llamawasm2go/p2.Fn1920
func Fn1920(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1921 github.com/goccy/llamawasm2go/p2.Fn1921
func Fn1921(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1922 github.com/goccy/llamawasm2go/p2.Fn1922
func Fn1922(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1923 github.com/goccy/llamawasm2go/p2.Fn1923
func Fn1923(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1924 github.com/goccy/llamawasm2go/p2.Fn1924
func Fn1924(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1926 github.com/goccy/llamawasm2go/p2.Fn1926
func Fn1926(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1928 github.com/goccy/llamawasm2go/p2.Fn1928
func Fn1928(m *base.Module, l0 int64)

//go:linkname Fn1945 github.com/goccy/llamawasm2go/p2.Fn1945
func Fn1945(m *base.Module, l0 int64)

//go:linkname Fn1946 github.com/goccy/llamawasm2go/p2.Fn1946
func Fn1946(m *base.Module, l0 int64)

//go:linkname Fn1947 github.com/goccy/llamawasm2go/p2.Fn1947
func Fn1947(m *base.Module, l0 int64)

//go:linkname Fn1949 github.com/goccy/llamawasm2go/p2.Fn1949
func Fn1949(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1950 github.com/goccy/llamawasm2go/p2.Fn1950
func Fn1950(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1989 github.com/goccy/llamawasm2go/p2.Fn1989
func Fn1989(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2001 github.com/goccy/llamawasm2go/p2.Fn2001
func Fn2001(m *base.Module, l0 int64) int64

//go:linkname Fn2004 github.com/goccy/llamawasm2go/p2.Fn2004
func Fn2004(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2008 github.com/goccy/llamawasm2go/p2.Fn2008
func Fn2008(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2012 github.com/goccy/llamawasm2go/p2.Fn2012
func Fn2012(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2025 github.com/goccy/llamawasm2go/p2.Fn2025
func Fn2025(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2026 github.com/goccy/llamawasm2go/p2.Fn2026
func Fn2026(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2027 github.com/goccy/llamawasm2go/p2.Fn2027
func Fn2027(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2039 github.com/goccy/llamawasm2go/p2.Fn2039
func Fn2039(m *base.Module, l0 int64) int64

//go:linkname Fn2040 github.com/goccy/llamawasm2go/p2.Fn2040
func Fn2040(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2043 github.com/goccy/llamawasm2go/p2.Fn2043
func Fn2043(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2045 github.com/goccy/llamawasm2go/p2.Fn2045
func Fn2045(m *base.Module)

//go:linkname Fn2048 github.com/goccy/llamawasm2go/p2.Fn2048
func Fn2048(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2057 github.com/goccy/llamawasm2go/p2.Fn2057
func Fn2057(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2061 github.com/goccy/llamawasm2go/p2.Fn2061
func Fn2061(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2062 github.com/goccy/llamawasm2go/p2.Fn2062
func Fn2062(m *base.Module, l0 int64)

//go:linkname Fn2063 github.com/goccy/llamawasm2go/p2.Fn2063
func Fn2063(m *base.Module, l0 int64)

//go:linkname Fn2064 github.com/goccy/llamawasm2go/p2.Fn2064
func Fn2064(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2066 github.com/goccy/llamawasm2go/p2.Fn2066
func Fn2066(m *base.Module, l0 int64)

//go:linkname Fn2067 github.com/goccy/llamawasm2go/p2.Fn2067
func Fn2067(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2068 github.com/goccy/llamawasm2go/p2.Fn2068
func Fn2068(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2072 github.com/goccy/llamawasm2go/p2.Fn2072
func Fn2072(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2075 github.com/goccy/llamawasm2go/p2.Fn2075
func Fn2075(m *base.Module, l0 int64) int64

//go:linkname Fn2076 github.com/goccy/llamawasm2go/p2.Fn2076
func Fn2076(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2077 github.com/goccy/llamawasm2go/p2.Fn2077
func Fn2077(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2084 github.com/goccy/llamawasm2go/p2.Fn2084
func Fn2084(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2085 github.com/goccy/llamawasm2go/p2.Fn2085
func Fn2085(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2087 github.com/goccy/llamawasm2go/p2.Fn2087
func Fn2087(m *base.Module) int64

//go:linkname Fn2089 github.com/goccy/llamawasm2go/p2.Fn2089
func Fn2089(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2090 github.com/goccy/llamawasm2go/p2.Fn2090
func Fn2090(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2091 github.com/goccy/llamawasm2go/p2.Fn2091
func Fn2091(m *base.Module) int64

//go:linkname Fn2093 github.com/goccy/llamawasm2go/p2.Fn2093
func Fn2093(m *base.Module, l0 int32) int64

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 int32) int32

//go:linkname Fn2095 github.com/goccy/llamawasm2go/p2.Fn2095
func Fn2095(m *base.Module, l0 int32) int64

//go:linkname Fn2096 github.com/goccy/llamawasm2go/p2.Fn2096
func Fn2096(m *base.Module, l0 float32) int64

//go:linkname Fn2097 github.com/goccy/llamawasm2go/p2.Fn2097
func Fn2097(m *base.Module, l0 float32) int64

//go:linkname Fn2098 github.com/goccy/llamawasm2go/p2.Fn2098
func Fn2098(m *base.Module, l0 float32) int64

//go:linkname Fn2100 github.com/goccy/llamawasm2go/p2.Fn2100
func Fn2100(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2101 github.com/goccy/llamawasm2go/p2.Fn2101
func Fn2101(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2142 github.com/goccy/llamawasm2go/p2.Fn2142
func Fn2142(m *base.Module, l0 int64)

//go:linkname Fn2144 github.com/goccy/llamawasm2go/p2.Fn2144
func Fn2144(m *base.Module, l0 int64)

//go:linkname Fn2152 github.com/goccy/llamawasm2go/p2.Fn2152
func Fn2152(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2186 github.com/goccy/llamawasm2go/p2.Fn2186
func Fn2186(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2197 github.com/goccy/llamawasm2go/p2.Fn2197
func Fn2197(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2200 github.com/goccy/llamawasm2go/p2.Fn2200
func Fn2200(m *base.Module, l0 int64)

//go:linkname Fn2201 github.com/goccy/llamawasm2go/p2.Fn2201
func Fn2201(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2202 github.com/goccy/llamawasm2go/p2.Fn2202
func Fn2202(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2203 github.com/goccy/llamawasm2go/p2.Fn2203
func Fn2203(m *base.Module, l0 int64)

//go:linkname Fn2206 github.com/goccy/llamawasm2go/p2.Fn2206
func Fn2206(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2211 github.com/goccy/llamawasm2go/p2.Fn2211
func Fn2211(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2212 github.com/goccy/llamawasm2go/p2.Fn2212
func Fn2212(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2226 github.com/goccy/llamawasm2go/p2.Fn2226
func Fn2226(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2236 github.com/goccy/llamawasm2go/p2.Fn2236
func Fn2236(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2239 github.com/goccy/llamawasm2go/p2.Fn2239
func Fn2239(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2241 github.com/goccy/llamawasm2go/p2.Fn2241
func Fn2241(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2242 github.com/goccy/llamawasm2go/p2.Fn2242
func Fn2242(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2243 github.com/goccy/llamawasm2go/p2.Fn2243
func Fn2243(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2246 github.com/goccy/llamawasm2go/p2.Fn2246
func Fn2246(m *base.Module)

//go:linkname Fn2250 github.com/goccy/llamawasm2go/p2.Fn2250
func Fn2250(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2252 github.com/goccy/llamawasm2go/p0.Fn2252
func Fn2252(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2267 github.com/goccy/llamawasm2go/p2.Fn2267
func Fn2267(m *base.Module, l0 int64)

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2273 github.com/goccy/llamawasm2go/p2.Fn2273
func Fn2273(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2278 github.com/goccy/llamawasm2go/p2.Fn2278
func Fn2278(m *base.Module, l0 int64) int64

//go:linkname Fn2279 github.com/goccy/llamawasm2go/p2.Fn2279
func Fn2279(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2280 github.com/goccy/llamawasm2go/p0.Fn2280
func Fn2280(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2291 github.com/goccy/llamawasm2go/p2.Fn2291
func Fn2291(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2292 github.com/goccy/llamawasm2go/p2.Fn2292
func Fn2292(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2293 github.com/goccy/llamawasm2go/p2.Fn2293
func Fn2293(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2294 github.com/goccy/llamawasm2go/p2.Fn2294
func Fn2294(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2314 github.com/goccy/llamawasm2go/p2.Fn2314
func Fn2314(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2315 github.com/goccy/llamawasm2go/p2.Fn2315
func Fn2315(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2318 github.com/goccy/llamawasm2go/p2.Fn2318
func Fn2318(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2404 github.com/goccy/llamawasm2go/p2.Fn2404
func Fn2404(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2523 github.com/goccy/llamawasm2go/p2.Fn2523
func Fn2523(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2524 github.com/goccy/llamawasm2go/p0.Fn2524
func Fn2524(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2525 github.com/goccy/llamawasm2go/p2.Fn2525
func Fn2525(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2526 github.com/goccy/llamawasm2go/p2.Fn2526
func Fn2526(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2530 github.com/goccy/llamawasm2go/p2.Fn2530
func Fn2530(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2555 github.com/goccy/llamawasm2go/p2.Fn2555
func Fn2555(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2687 github.com/goccy/llamawasm2go/p2.Fn2687
func Fn2687(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2722 github.com/goccy/llamawasm2go/p2.Fn2722
func Fn2722(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2736 github.com/goccy/llamawasm2go/p2.Fn2736
func Fn2736(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2746 github.com/goccy/llamawasm2go/p2.Fn2746
func Fn2746(m *base.Module, l0 int32)

//go:linkname Fn2748 github.com/goccy/llamawasm2go/p2.Fn2748
func Fn2748(m *base.Module, l0 int64) int64

//go:linkname Fn2749 github.com/goccy/llamawasm2go/p2.Fn2749
func Fn2749(m *base.Module, l0 int64)

//go:linkname Fn2752 github.com/goccy/llamawasm2go/p2.Fn2752
func Fn2752(m *base.Module, l0 int64)

//go:linkname Fn2753 github.com/goccy/llamawasm2go/p2.Fn2753
func Fn2753(m *base.Module, l0 int64)

//go:linkname Fn2755 github.com/goccy/llamawasm2go/p2.Fn2755
func Fn2755(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2760 github.com/goccy/llamawasm2go/p2.Fn2760
func Fn2760(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2762 github.com/goccy/llamawasm2go/p2.Fn2762
func Fn2762(m *base.Module, l0 int64) int32

//go:linkname Fn2766 github.com/goccy/llamawasm2go/p2.Fn2766
func Fn2766(m *base.Module) int32

//go:linkname Fn2774 github.com/goccy/llamawasm2go/p2.Fn2774
func Fn2774(m *base.Module, l0 float64) float32

//go:linkname Fn2775 github.com/goccy/llamawasm2go/p2.Fn2775
func Fn2775(m *base.Module, l0 float64) float32

//go:linkname Fn2782 github.com/goccy/llamawasm2go/p2.Fn2782
func Fn2782(m *base.Module, l0 float32) float32

//go:linkname Fn2786 github.com/goccy/llamawasm2go/p2.Fn2786
func Fn2786(m *base.Module, l0 float32) float32

//go:linkname Fn2789 github.com/goccy/llamawasm2go/p2.Fn2789
func Fn2789(m *base.Module, l0 float32) float32

//go:linkname Fn2804 github.com/goccy/llamawasm2go/p2.Fn2804
func Fn2804(m *base.Module, l0 int64) int32

//go:linkname Fn2805 github.com/goccy/llamawasm2go/p2.Fn2805
func Fn2805(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2807 github.com/goccy/llamawasm2go/p2.Fn2807
func Fn2807(m *base.Module, l0 int64)

//go:linkname Fn2808 github.com/goccy/llamawasm2go/p2.Fn2808
func Fn2808(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2809 github.com/goccy/llamawasm2go/p2.Fn2809
func Fn2809(m *base.Module, l0 int64) int32

//go:linkname Fn2816 github.com/goccy/llamawasm2go/p2.Fn2816
func Fn2816(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2817 github.com/goccy/llamawasm2go/p2.Fn2817
func Fn2817(m *base.Module)

//go:linkname Fn2818 github.com/goccy/llamawasm2go/p2.Fn2818
func Fn2818(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2824 github.com/goccy/llamawasm2go/p2.Fn2824
func Fn2824(m *base.Module, l0 int64) int32

//go:linkname Fn2825 github.com/goccy/llamawasm2go/p2.Fn2825
func Fn2825(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2830 github.com/goccy/llamawasm2go/p2.Fn2830
func Fn2830(m *base.Module, l0 int64) int32

//go:linkname Fn2834 github.com/goccy/llamawasm2go/p2.Fn2834
func Fn2834(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn2835 github.com/goccy/llamawasm2go/p2.Fn2835
func Fn2835(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2838 github.com/goccy/llamawasm2go/p2.Fn2838
func Fn2838(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2841 github.com/goccy/llamawasm2go/p2.Fn2841
func Fn2841(m *base.Module, l0 int64) int64

//go:linkname Fn2845 github.com/goccy/llamawasm2go/p2.Fn2845
func Fn2845(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn2855 github.com/goccy/llamawasm2go/p2.Fn2855
func Fn2855(m *base.Module)

//go:linkname Fn2856 github.com/goccy/llamawasm2go/p0.Fn2856
func Fn2856(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2858 github.com/goccy/llamawasm2go/p2.Fn2858
func Fn2858(m *base.Module)

//go:linkname Fn2860 github.com/goccy/llamawasm2go/p0.Fn2860
func Fn2860(m *base.Module, l0 int64) int64

//go:linkname Fn2862 github.com/goccy/llamawasm2go/p2.Fn2862
func Fn2862(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2863 github.com/goccy/llamawasm2go/p2.Fn2863
func Fn2863(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2866 github.com/goccy/llamawasm2go/p2.Fn2866
func Fn2866(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2895 github.com/goccy/llamawasm2go/p2.Fn2895
func Fn2895(m *base.Module, l0 int32)
