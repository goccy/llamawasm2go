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

//go:linkname Fn272 github.com/goccy/llamawasm2go/p2.Fn272
func Fn272(m *base.Module, l0 int64, l1 int64, l2 int64)

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

//go:linkname Fn797 github.com/goccy/llamawasm2go/p2.Fn797
func Fn797(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn858 github.com/goccy/llamawasm2go/p2.Fn858
func Fn858(m *base.Module) int64

//go:linkname Fn909 github.com/goccy/llamawasm2go/p2.Fn909
func Fn909(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn911 github.com/goccy/llamawasm2go/p2.Fn911
func Fn911(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn936 github.com/goccy/llamawasm2go/p2.Fn936
func Fn936(m *base.Module, l0 int64)

//go:linkname Fn941 github.com/goccy/llamawasm2go/p2.Fn941
func Fn941(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn986 github.com/goccy/llamawasm2go/p0.Fn986
func Fn986(m *base.Module, l0 int64) int64

//go:linkname Fn988 github.com/goccy/llamawasm2go/p2.Fn988
func Fn988(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1065 github.com/goccy/llamawasm2go/p2.Fn1065
func Fn1065(m *base.Module, l0 int64)

//go:linkname Fn1087 github.com/goccy/llamawasm2go/p2.Fn1087
func Fn1087(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1095 github.com/goccy/llamawasm2go/p2.Fn1095
func Fn1095(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1105 github.com/goccy/llamawasm2go/p2.Fn1105
func Fn1105(m *base.Module, l0 int64) int64

//go:linkname Fn1135 github.com/goccy/llamawasm2go/p2.Fn1135
func Fn1135(m *base.Module, l0 int32) int64

//go:linkname Fn1146 github.com/goccy/llamawasm2go/p2.Fn1146
func Fn1146(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1147 github.com/goccy/llamawasm2go/p2.Fn1147
func Fn1147(m *base.Module, l0 int64)

//go:linkname Fn1149 github.com/goccy/llamawasm2go/p2.Fn1149
func Fn1149(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1169 github.com/goccy/llamawasm2go/p2.Fn1169
func Fn1169(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1170 github.com/goccy/llamawasm2go/p2.Fn1170
func Fn1170(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1171 github.com/goccy/llamawasm2go/p2.Fn1171
func Fn1171(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1172 github.com/goccy/llamawasm2go/p2.Fn1172
func Fn1172(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1173 github.com/goccy/llamawasm2go/p2.Fn1173
func Fn1173(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1174 github.com/goccy/llamawasm2go/p2.Fn1174
func Fn1174(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1175 github.com/goccy/llamawasm2go/p2.Fn1175
func Fn1175(m *base.Module, l0 int64) int64

//go:linkname Fn1178 github.com/goccy/llamawasm2go/p2.Fn1178
func Fn1178(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1182 github.com/goccy/llamawasm2go/p2.Fn1182
func Fn1182(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1183 github.com/goccy/llamawasm2go/p2.Fn1183
func Fn1183(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1184 github.com/goccy/llamawasm2go/p2.Fn1184
func Fn1184(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1185 github.com/goccy/llamawasm2go/p2.Fn1185
func Fn1185(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1186 github.com/goccy/llamawasm2go/p2.Fn1186
func Fn1186(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1187 github.com/goccy/llamawasm2go/p2.Fn1187
func Fn1187(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1188 github.com/goccy/llamawasm2go/p2.Fn1188
func Fn1188(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1189 github.com/goccy/llamawasm2go/p2.Fn1189
func Fn1189(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1190 github.com/goccy/llamawasm2go/p2.Fn1190
func Fn1190(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1191 github.com/goccy/llamawasm2go/p2.Fn1191
func Fn1191(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1192 github.com/goccy/llamawasm2go/p2.Fn1192
func Fn1192(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1193 github.com/goccy/llamawasm2go/p2.Fn1193
func Fn1193(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1194 github.com/goccy/llamawasm2go/p2.Fn1194
func Fn1194(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1195 github.com/goccy/llamawasm2go/p2.Fn1195
func Fn1195(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1196 github.com/goccy/llamawasm2go/p2.Fn1196
func Fn1196(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1197 github.com/goccy/llamawasm2go/p2.Fn1197
func Fn1197(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1198 github.com/goccy/llamawasm2go/p2.Fn1198
func Fn1198(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1199 github.com/goccy/llamawasm2go/p2.Fn1199
func Fn1199(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1200 github.com/goccy/llamawasm2go/p2.Fn1200
func Fn1200(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1201 github.com/goccy/llamawasm2go/p2.Fn1201
func Fn1201(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1202 github.com/goccy/llamawasm2go/p2.Fn1202
func Fn1202(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1203 github.com/goccy/llamawasm2go/p2.Fn1203
func Fn1203(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1204 github.com/goccy/llamawasm2go/p2.Fn1204
func Fn1204(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1205 github.com/goccy/llamawasm2go/p2.Fn1205
func Fn1205(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1206 github.com/goccy/llamawasm2go/p2.Fn1206
func Fn1206(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1207 github.com/goccy/llamawasm2go/p2.Fn1207
func Fn1207(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1208 github.com/goccy/llamawasm2go/p2.Fn1208
func Fn1208(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1210 github.com/goccy/llamawasm2go/p2.Fn1210
func Fn1210(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1214 github.com/goccy/llamawasm2go/p2.Fn1214
func Fn1214(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1215 github.com/goccy/llamawasm2go/p2.Fn1215
func Fn1215(m *base.Module)

//go:linkname Fn1216 github.com/goccy/llamawasm2go/p2.Fn1216
func Fn1216(m *base.Module)

//go:linkname Fn1217 github.com/goccy/llamawasm2go/p0.Fn1217
func Fn1217(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1225 github.com/goccy/llamawasm2go/p2.Fn1225
func Fn1225(m *base.Module)

//go:linkname Fn1227 github.com/goccy/llamawasm2go/p2.Fn1227
func Fn1227(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1235 github.com/goccy/llamawasm2go/p2.Fn1235
func Fn1235(m *base.Module, l0 int64)

//go:linkname Fn1242 github.com/goccy/llamawasm2go/p2.Fn1242
func Fn1242(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1245 github.com/goccy/llamawasm2go/p2.Fn1245
func Fn1245(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1247 github.com/goccy/llamawasm2go/p2.Fn1247
func Fn1247(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1249 github.com/goccy/llamawasm2go/p2.Fn1249
func Fn1249(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1251 github.com/goccy/llamawasm2go/p2.Fn1251
func Fn1251(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1252 github.com/goccy/llamawasm2go/p2.Fn1252
func Fn1252(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1253 github.com/goccy/llamawasm2go/p2.Fn1253
func Fn1253(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1258 github.com/goccy/llamawasm2go/p2.Fn1258
func Fn1258(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1259 github.com/goccy/llamawasm2go/p2.Fn1259
func Fn1259(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1260 github.com/goccy/llamawasm2go/p2.Fn1260
func Fn1260(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1261 github.com/goccy/llamawasm2go/p0.Fn1261
func Fn1261(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1263 github.com/goccy/llamawasm2go/p2.Fn1263
func Fn1263(m *base.Module, l0 int64)

//go:linkname Fn1264 github.com/goccy/llamawasm2go/p2.Fn1264
func Fn1264(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1266 github.com/goccy/llamawasm2go/p2.Fn1266
func Fn1266(m *base.Module, l0 int64) int64

//go:linkname Fn1267 github.com/goccy/llamawasm2go/p2.Fn1267
func Fn1267(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1268 github.com/goccy/llamawasm2go/p2.Fn1268
func Fn1268(m *base.Module, l0 int64)

//go:linkname Fn1269 github.com/goccy/llamawasm2go/p2.Fn1269
func Fn1269(m *base.Module, l0 int64)

//go:linkname Fn1270 github.com/goccy/llamawasm2go/p2.Fn1270
func Fn1270(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1271 github.com/goccy/llamawasm2go/p2.Fn1271
func Fn1271(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1274 github.com/goccy/llamawasm2go/p2.Fn1274
func Fn1274(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1281 github.com/goccy/llamawasm2go/p2.Fn1281
func Fn1281(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1283 github.com/goccy/llamawasm2go/p0.Fn1283
func Fn1283(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1285 github.com/goccy/llamawasm2go/p2.Fn1285
func Fn1285(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1286 github.com/goccy/llamawasm2go/p2.Fn1286
func Fn1286(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1288 github.com/goccy/llamawasm2go/p0.Fn1288
func Fn1288(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1289 github.com/goccy/llamawasm2go/p2.Fn1289
func Fn1289(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1293 github.com/goccy/llamawasm2go/p2.Fn1293
func Fn1293(m *base.Module, l0 int64)

//go:linkname Fn1296 github.com/goccy/llamawasm2go/p2.Fn1296
func Fn1296(m *base.Module, l0 int64)

//go:linkname Fn1298 github.com/goccy/llamawasm2go/p2.Fn1298
func Fn1298(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1299 github.com/goccy/llamawasm2go/p2.Fn1299
func Fn1299(m *base.Module, l0 int64) int64

//go:linkname Fn1300 github.com/goccy/llamawasm2go/p2.Fn1300
func Fn1300(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1301 github.com/goccy/llamawasm2go/p2.Fn1301
func Fn1301(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1302 github.com/goccy/llamawasm2go/p2.Fn1302
func Fn1302(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1303 github.com/goccy/llamawasm2go/p2.Fn1303
func Fn1303(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1304 github.com/goccy/llamawasm2go/p2.Fn1304
func Fn1304(m *base.Module, l0 int64)

//go:linkname Fn1307 github.com/goccy/llamawasm2go/p2.Fn1307
func Fn1307(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1309 github.com/goccy/llamawasm2go/p2.Fn1309
func Fn1309(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1312 github.com/goccy/llamawasm2go/p2.Fn1312
func Fn1312(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1319 github.com/goccy/llamawasm2go/p2.Fn1319
func Fn1319(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1320 github.com/goccy/llamawasm2go/p2.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1326 github.com/goccy/llamawasm2go/p2.Fn1326
func Fn1326(m *base.Module, l0 int64)

//go:linkname Fn1327 github.com/goccy/llamawasm2go/p2.Fn1327
func Fn1327(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1329 github.com/goccy/llamawasm2go/p2.Fn1329
func Fn1329(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1330 github.com/goccy/llamawasm2go/p2.Fn1330
func Fn1330(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1331 github.com/goccy/llamawasm2go/p2.Fn1331
func Fn1331(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1373 github.com/goccy/llamawasm2go/p2.Fn1373
func Fn1373(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1389 github.com/goccy/llamawasm2go/p2.Fn1389
func Fn1389(m *base.Module, l0 int64)

//go:linkname Fn1390 github.com/goccy/llamawasm2go/p2.Fn1390
func Fn1390(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1391 github.com/goccy/llamawasm2go/p2.Fn1391
func Fn1391(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p2.Fn1392
func Fn1392(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1393 github.com/goccy/llamawasm2go/p2.Fn1393
func Fn1393(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p2.Fn1394
func Fn1394(m *base.Module, l0 int64)

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p2.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p2.Fn1397
func Fn1397(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1399 github.com/goccy/llamawasm2go/p2.Fn1399
func Fn1399(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p2.Fn1406
func Fn1406(m *base.Module, l0 int64) int64

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p2.Fn1407
func Fn1407(m *base.Module, l0 int64) int64

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p2.Fn1408
func Fn1408(m *base.Module, l0 int64) int64

//go:linkname Fn1409 github.com/goccy/llamawasm2go/p2.Fn1409
func Fn1409(m *base.Module, l0 int64) int64

//go:linkname Fn1410 github.com/goccy/llamawasm2go/p2.Fn1410
func Fn1410(m *base.Module, l0 int64) int64

//go:linkname Fn1411 github.com/goccy/llamawasm2go/p2.Fn1411
func Fn1411(m *base.Module, l0 int64) int64

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p2.Fn1413
func Fn1413(m *base.Module, l0 int64) int64

//go:linkname Fn1414 github.com/goccy/llamawasm2go/p2.Fn1414
func Fn1414(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1416 github.com/goccy/llamawasm2go/p2.Fn1416
func Fn1416(m *base.Module, l0 int64) int64

//go:linkname Fn1417 github.com/goccy/llamawasm2go/p2.Fn1417
func Fn1417(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1418 github.com/goccy/llamawasm2go/p2.Fn1418
func Fn1418(m *base.Module, l0 int64) int64

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1421 github.com/goccy/llamawasm2go/p2.Fn1421
func Fn1421(m *base.Module, l0 int64) int64

//go:linkname Fn1423 github.com/goccy/llamawasm2go/p2.Fn1423
func Fn1423(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1425 github.com/goccy/llamawasm2go/p2.Fn1425
func Fn1425(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p2.Fn1426
func Fn1426(m *base.Module, l0 int64) int64

//go:linkname Fn1427 github.com/goccy/llamawasm2go/p2.Fn1427
func Fn1427(m *base.Module, l0 int64) int64

//go:linkname Fn1428 github.com/goccy/llamawasm2go/p2.Fn1428
func Fn1428(m *base.Module, l0 int64) int64

//go:linkname Fn1434 github.com/goccy/llamawasm2go/p2.Fn1434
func Fn1434(m *base.Module, l0 int64) int64

//go:linkname Fn1436 github.com/goccy/llamawasm2go/p2.Fn1436
func Fn1436(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1437 github.com/goccy/llamawasm2go/p2.Fn1437
func Fn1437(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1439 github.com/goccy/llamawasm2go/p2.Fn1439
func Fn1439(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1440 github.com/goccy/llamawasm2go/p2.Fn1440
func Fn1440(m *base.Module, l0 int64) int64

//go:linkname Fn1441 github.com/goccy/llamawasm2go/p2.Fn1441
func Fn1441(m *base.Module, l0 int64) int64

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int64) int64

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64)

//go:linkname Fn1473 github.com/goccy/llamawasm2go/p2.Fn1473
func Fn1473(m *base.Module, l0 int64) int64

//go:linkname Fn1479 github.com/goccy/llamawasm2go/p2.Fn1479
func Fn1479(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1480 github.com/goccy/llamawasm2go/p2.Fn1480
func Fn1480(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1481 github.com/goccy/llamawasm2go/p2.Fn1481
func Fn1481(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1483 github.com/goccy/llamawasm2go/p2.Fn1483
func Fn1483(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1484 github.com/goccy/llamawasm2go/p2.Fn1484
func Fn1484(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1485 github.com/goccy/llamawasm2go/p2.Fn1485
func Fn1485(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1486 github.com/goccy/llamawasm2go/p2.Fn1486
func Fn1486(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1487 github.com/goccy/llamawasm2go/p2.Fn1487
func Fn1487(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1488 github.com/goccy/llamawasm2go/p2.Fn1488
func Fn1488(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1490 github.com/goccy/llamawasm2go/p2.Fn1490
func Fn1490(m *base.Module, l0 int64) int32

//go:linkname Fn1491 github.com/goccy/llamawasm2go/p2.Fn1491
func Fn1491(m *base.Module, l0 int64) int32

//go:linkname Fn1492 github.com/goccy/llamawasm2go/p2.Fn1492
func Fn1492(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1493 github.com/goccy/llamawasm2go/p2.Fn1493
func Fn1493(m *base.Module, l0 int64) int32

//go:linkname Fn1494 github.com/goccy/llamawasm2go/p2.Fn1494
func Fn1494(m *base.Module, l0 int64) int32

//go:linkname Fn1497 github.com/goccy/llamawasm2go/p2.Fn1497
func Fn1497(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1498 github.com/goccy/llamawasm2go/p2.Fn1498
func Fn1498(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1499 github.com/goccy/llamawasm2go/p2.Fn1499
func Fn1499(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1500 github.com/goccy/llamawasm2go/p2.Fn1500
func Fn1500(m *base.Module)

//go:linkname Fn1501 github.com/goccy/llamawasm2go/p2.Fn1501
func Fn1501(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1503 github.com/goccy/llamawasm2go/p2.Fn1503
func Fn1503(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1505 github.com/goccy/llamawasm2go/p2.Fn1505
func Fn1505(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1506 github.com/goccy/llamawasm2go/p2.Fn1506
func Fn1506(m *base.Module, l0 int64)

//go:linkname Fn1510 github.com/goccy/llamawasm2go/p2.Fn1510
func Fn1510(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1511 github.com/goccy/llamawasm2go/p2.Fn1511
func Fn1511(m *base.Module, l0 int64)

//go:linkname Fn1514 github.com/goccy/llamawasm2go/p2.Fn1514
func Fn1514(m *base.Module, l0 int64)

//go:linkname Fn1523 github.com/goccy/llamawasm2go/p2.Fn1523
func Fn1523(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1524 github.com/goccy/llamawasm2go/p2.Fn1524
func Fn1524(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1533 github.com/goccy/llamawasm2go/p2.Fn1533
func Fn1533(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1534 github.com/goccy/llamawasm2go/p0.Fn1534
func Fn1534(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1536 github.com/goccy/llamawasm2go/p2.Fn1536
func Fn1536(m *base.Module, l0 int64)

//go:linkname Fn1538 github.com/goccy/llamawasm2go/p2.Fn1538
func Fn1538(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1546 github.com/goccy/llamawasm2go/p2.Fn1546
func Fn1546(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1560 github.com/goccy/llamawasm2go/p2.Fn1560
func Fn1560(m *base.Module, l0 int64)

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p2.Fn1563
func Fn1563(m *base.Module, l0 int64) int32

//go:linkname Fn1573 github.com/goccy/llamawasm2go/p2.Fn1573
func Fn1573(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1575 github.com/goccy/llamawasm2go/p2.Fn1575
func Fn1575(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1576 github.com/goccy/llamawasm2go/p2.Fn1576
func Fn1576(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1595 github.com/goccy/llamawasm2go/p2.Fn1595
func Fn1595(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1608 github.com/goccy/llamawasm2go/p2.Fn1608
func Fn1608(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1615 github.com/goccy/llamawasm2go/p2.Fn1615
func Fn1615(m *base.Module, l0 int64)

//go:linkname Fn1649 github.com/goccy/llamawasm2go/p2.Fn1649
func Fn1649(m *base.Module, l0 int64)

//go:linkname Fn1652 github.com/goccy/llamawasm2go/p2.Fn1652
func Fn1652(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1666 github.com/goccy/llamawasm2go/p2.Fn1666
func Fn1666(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1667 github.com/goccy/llamawasm2go/p2.Fn1667
func Fn1667(m *base.Module, l0 int64) int64

//go:linkname Fn1668 github.com/goccy/llamawasm2go/p2.Fn1668
func Fn1668(m *base.Module, l0 int64)

//go:linkname Fn1672 github.com/goccy/llamawasm2go/p0.Fn1672
func Fn1672(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1673 github.com/goccy/llamawasm2go/p2.Fn1673
func Fn1673(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1674 github.com/goccy/llamawasm2go/p2.Fn1674
func Fn1674(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1675 github.com/goccy/llamawasm2go/p2.Fn1675
func Fn1675(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1676 github.com/goccy/llamawasm2go/p2.Fn1676
func Fn1676(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1694 github.com/goccy/llamawasm2go/p2.Fn1694
func Fn1694(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1704 github.com/goccy/llamawasm2go/p2.Fn1704
func Fn1704(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1706 github.com/goccy/llamawasm2go/p2.Fn1706
func Fn1706(m *base.Module, l0 int64) int64

//go:linkname Fn1707 github.com/goccy/llamawasm2go/p2.Fn1707
func Fn1707(m *base.Module, l0 int64)

//go:linkname Fn1710 github.com/goccy/llamawasm2go/p0.Fn1710
func Fn1710(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1712 github.com/goccy/llamawasm2go/p2.Fn1712
func Fn1712(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1723 github.com/goccy/llamawasm2go/p2.Fn1723
func Fn1723(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1726 github.com/goccy/llamawasm2go/p2.Fn1726
func Fn1726(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1727 github.com/goccy/llamawasm2go/p2.Fn1727
func Fn1727(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1737 github.com/goccy/llamawasm2go/p2.Fn1737
func Fn1737(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1763 github.com/goccy/llamawasm2go/p2.Fn1763
func Fn1763(m *base.Module, l0 int64)

//go:linkname Fn1764 github.com/goccy/llamawasm2go/p2.Fn1764
func Fn1764(m *base.Module, l0 int64)

//go:linkname Fn1775 github.com/goccy/llamawasm2go/p2.Fn1775
func Fn1775(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1800 github.com/goccy/llamawasm2go/p2.Fn1800
func Fn1800(m *base.Module) int64

//go:linkname Fn1803 github.com/goccy/llamawasm2go/p2.Fn1803
func Fn1803(m *base.Module, l0 int64) int64

//go:linkname Fn1804 github.com/goccy/llamawasm2go/p2.Fn1804
func Fn1804(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1811 github.com/goccy/llamawasm2go/p2.Fn1811
func Fn1811(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1815 github.com/goccy/llamawasm2go/p2.Fn1815
func Fn1815(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1817 github.com/goccy/llamawasm2go/p2.Fn1817
func Fn1817(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1818 github.com/goccy/llamawasm2go/p2.Fn1818
func Fn1818(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1819 github.com/goccy/llamawasm2go/p2.Fn1819
func Fn1819(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1820 github.com/goccy/llamawasm2go/p2.Fn1820
func Fn1820(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1821 github.com/goccy/llamawasm2go/p2.Fn1821
func Fn1821(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1826 github.com/goccy/llamawasm2go/p2.Fn1826
func Fn1826(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1832 github.com/goccy/llamawasm2go/p2.Fn1832
func Fn1832(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1834 github.com/goccy/llamawasm2go/p2.Fn1834
func Fn1834(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1837 github.com/goccy/llamawasm2go/p2.Fn1837
func Fn1837(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1838 github.com/goccy/llamawasm2go/p2.Fn1838
func Fn1838(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1839 github.com/goccy/llamawasm2go/p2.Fn1839
func Fn1839(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1842 github.com/goccy/llamawasm2go/p2.Fn1842
func Fn1842(m *base.Module, l0 int64)

//go:linkname Fn1852 github.com/goccy/llamawasm2go/p2.Fn1852
func Fn1852(m *base.Module, l0 int64)

//go:linkname Fn1854 github.com/goccy/llamawasm2go/p2.Fn1854
func Fn1854(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1855 github.com/goccy/llamawasm2go/p2.Fn1855
func Fn1855(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1868 github.com/goccy/llamawasm2go/p2.Fn1868
func Fn1868(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1869 github.com/goccy/llamawasm2go/p2.Fn1869
func Fn1869(m *base.Module, l0 int64) int64

//go:linkname Fn1870 github.com/goccy/llamawasm2go/p2.Fn1870
func Fn1870(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1872 github.com/goccy/llamawasm2go/p2.Fn1872
func Fn1872(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1884 github.com/goccy/llamawasm2go/p2.Fn1884
func Fn1884(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1886 github.com/goccy/llamawasm2go/p2.Fn1886
func Fn1886(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1887 github.com/goccy/llamawasm2go/p2.Fn1887
func Fn1887(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1889 github.com/goccy/llamawasm2go/p2.Fn1889
func Fn1889(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1890 github.com/goccy/llamawasm2go/p0.Fn1890
func Fn1890(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1891 github.com/goccy/llamawasm2go/p2.Fn1891
func Fn1891(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1892 github.com/goccy/llamawasm2go/p2.Fn1892
func Fn1892(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1893 github.com/goccy/llamawasm2go/p2.Fn1893
func Fn1893(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn1894 github.com/goccy/llamawasm2go/p2.Fn1894
func Fn1894(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1913 github.com/goccy/llamawasm2go/p2.Fn1913
func Fn1913(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1914 github.com/goccy/llamawasm2go/p2.Fn1914
func Fn1914(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1915 github.com/goccy/llamawasm2go/p2.Fn1915
func Fn1915(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1917 github.com/goccy/llamawasm2go/p2.Fn1917
func Fn1917(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1918 github.com/goccy/llamawasm2go/p2.Fn1918
func Fn1918(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1919 github.com/goccy/llamawasm2go/p2.Fn1919
func Fn1919(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1920 github.com/goccy/llamawasm2go/p2.Fn1920
func Fn1920(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1921 github.com/goccy/llamawasm2go/p2.Fn1921
func Fn1921(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1923 github.com/goccy/llamawasm2go/p2.Fn1923
func Fn1923(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1925 github.com/goccy/llamawasm2go/p2.Fn1925
func Fn1925(m *base.Module, l0 int64)

//go:linkname Fn1942 github.com/goccy/llamawasm2go/p2.Fn1942
func Fn1942(m *base.Module, l0 int64)

//go:linkname Fn1943 github.com/goccy/llamawasm2go/p2.Fn1943
func Fn1943(m *base.Module, l0 int64)

//go:linkname Fn1944 github.com/goccy/llamawasm2go/p2.Fn1944
func Fn1944(m *base.Module, l0 int64)

//go:linkname Fn1946 github.com/goccy/llamawasm2go/p2.Fn1946
func Fn1946(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1947 github.com/goccy/llamawasm2go/p2.Fn1947
func Fn1947(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1986 github.com/goccy/llamawasm2go/p2.Fn1986
func Fn1986(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn1998 github.com/goccy/llamawasm2go/p2.Fn1998
func Fn1998(m *base.Module, l0 int64) int64

//go:linkname Fn2001 github.com/goccy/llamawasm2go/p2.Fn2001
func Fn2001(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2005 github.com/goccy/llamawasm2go/p2.Fn2005
func Fn2005(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2009 github.com/goccy/llamawasm2go/p2.Fn2009
func Fn2009(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2022 github.com/goccy/llamawasm2go/p2.Fn2022
func Fn2022(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2023 github.com/goccy/llamawasm2go/p2.Fn2023
func Fn2023(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2024 github.com/goccy/llamawasm2go/p2.Fn2024
func Fn2024(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2036 github.com/goccy/llamawasm2go/p2.Fn2036
func Fn2036(m *base.Module, l0 int64) int64

//go:linkname Fn2037 github.com/goccy/llamawasm2go/p2.Fn2037
func Fn2037(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2040 github.com/goccy/llamawasm2go/p2.Fn2040
func Fn2040(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2042 github.com/goccy/llamawasm2go/p2.Fn2042
func Fn2042(m *base.Module)

//go:linkname Fn2045 github.com/goccy/llamawasm2go/p2.Fn2045
func Fn2045(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2054 github.com/goccy/llamawasm2go/p2.Fn2054
func Fn2054(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2058 github.com/goccy/llamawasm2go/p2.Fn2058
func Fn2058(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2059 github.com/goccy/llamawasm2go/p2.Fn2059
func Fn2059(m *base.Module, l0 int64)

//go:linkname Fn2060 github.com/goccy/llamawasm2go/p2.Fn2060
func Fn2060(m *base.Module, l0 int64)

//go:linkname Fn2061 github.com/goccy/llamawasm2go/p2.Fn2061
func Fn2061(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2063 github.com/goccy/llamawasm2go/p2.Fn2063
func Fn2063(m *base.Module, l0 int64)

//go:linkname Fn2064 github.com/goccy/llamawasm2go/p2.Fn2064
func Fn2064(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2065 github.com/goccy/llamawasm2go/p2.Fn2065
func Fn2065(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2069 github.com/goccy/llamawasm2go/p2.Fn2069
func Fn2069(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2072 github.com/goccy/llamawasm2go/p2.Fn2072
func Fn2072(m *base.Module, l0 int64) int64

//go:linkname Fn2073 github.com/goccy/llamawasm2go/p2.Fn2073
func Fn2073(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2074 github.com/goccy/llamawasm2go/p2.Fn2074
func Fn2074(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2081 github.com/goccy/llamawasm2go/p2.Fn2081
func Fn2081(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2082 github.com/goccy/llamawasm2go/p2.Fn2082
func Fn2082(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2084 github.com/goccy/llamawasm2go/p2.Fn2084
func Fn2084(m *base.Module) int64

//go:linkname Fn2086 github.com/goccy/llamawasm2go/p2.Fn2086
func Fn2086(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2087 github.com/goccy/llamawasm2go/p2.Fn2087
func Fn2087(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2088 github.com/goccy/llamawasm2go/p2.Fn2088
func Fn2088(m *base.Module) int64

//go:linkname Fn2090 github.com/goccy/llamawasm2go/p2.Fn2090
func Fn2090(m *base.Module, l0 int32) int64

//go:linkname Fn2091 github.com/goccy/llamawasm2go/p2.Fn2091
func Fn2091(m *base.Module, l0 int32) int32

//go:linkname Fn2092 github.com/goccy/llamawasm2go/p2.Fn2092
func Fn2092(m *base.Module, l0 int32) int64

//go:linkname Fn2093 github.com/goccy/llamawasm2go/p2.Fn2093
func Fn2093(m *base.Module, l0 float32) int64

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 float32) int64

//go:linkname Fn2095 github.com/goccy/llamawasm2go/p2.Fn2095
func Fn2095(m *base.Module, l0 float32) int64

//go:linkname Fn2097 github.com/goccy/llamawasm2go/p2.Fn2097
func Fn2097(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2098 github.com/goccy/llamawasm2go/p2.Fn2098
func Fn2098(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2139 github.com/goccy/llamawasm2go/p2.Fn2139
func Fn2139(m *base.Module, l0 int64)

//go:linkname Fn2141 github.com/goccy/llamawasm2go/p2.Fn2141
func Fn2141(m *base.Module, l0 int64)

//go:linkname Fn2149 github.com/goccy/llamawasm2go/p2.Fn2149
func Fn2149(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2183 github.com/goccy/llamawasm2go/p2.Fn2183
func Fn2183(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2194 github.com/goccy/llamawasm2go/p2.Fn2194
func Fn2194(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2197 github.com/goccy/llamawasm2go/p2.Fn2197
func Fn2197(m *base.Module, l0 int64)

//go:linkname Fn2198 github.com/goccy/llamawasm2go/p2.Fn2198
func Fn2198(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2199 github.com/goccy/llamawasm2go/p2.Fn2199
func Fn2199(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2200 github.com/goccy/llamawasm2go/p2.Fn2200
func Fn2200(m *base.Module, l0 int64)

//go:linkname Fn2203 github.com/goccy/llamawasm2go/p2.Fn2203
func Fn2203(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2208 github.com/goccy/llamawasm2go/p2.Fn2208
func Fn2208(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2209 github.com/goccy/llamawasm2go/p2.Fn2209
func Fn2209(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2223 github.com/goccy/llamawasm2go/p2.Fn2223
func Fn2223(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2233 github.com/goccy/llamawasm2go/p2.Fn2233
func Fn2233(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2236 github.com/goccy/llamawasm2go/p2.Fn2236
func Fn2236(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2238 github.com/goccy/llamawasm2go/p2.Fn2238
func Fn2238(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2239 github.com/goccy/llamawasm2go/p2.Fn2239
func Fn2239(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2240 github.com/goccy/llamawasm2go/p2.Fn2240
func Fn2240(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2243 github.com/goccy/llamawasm2go/p2.Fn2243
func Fn2243(m *base.Module)

//go:linkname Fn2247 github.com/goccy/llamawasm2go/p2.Fn2247
func Fn2247(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p0.Fn2249
func Fn2249(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2264 github.com/goccy/llamawasm2go/p2.Fn2264
func Fn2264(m *base.Module, l0 int64)

//go:linkname Fn2265 github.com/goccy/llamawasm2go/p2.Fn2265
func Fn2265(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2270 github.com/goccy/llamawasm2go/p2.Fn2270
func Fn2270(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2275 github.com/goccy/llamawasm2go/p2.Fn2275
func Fn2275(m *base.Module, l0 int64) int64

//go:linkname Fn2276 github.com/goccy/llamawasm2go/p2.Fn2276
func Fn2276(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2288 github.com/goccy/llamawasm2go/p2.Fn2288
func Fn2288(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2289 github.com/goccy/llamawasm2go/p2.Fn2289
func Fn2289(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2290 github.com/goccy/llamawasm2go/p2.Fn2290
func Fn2290(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2291 github.com/goccy/llamawasm2go/p2.Fn2291
func Fn2291(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2292 github.com/goccy/llamawasm2go/p2.Fn2292
func Fn2292(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2293 github.com/goccy/llamawasm2go/p2.Fn2293
func Fn2293(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2311 github.com/goccy/llamawasm2go/p2.Fn2311
func Fn2311(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2312 github.com/goccy/llamawasm2go/p2.Fn2312
func Fn2312(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2313 github.com/goccy/llamawasm2go/p2.Fn2313
func Fn2313(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2314 github.com/goccy/llamawasm2go/p2.Fn2314
func Fn2314(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2315 github.com/goccy/llamawasm2go/p2.Fn2315
func Fn2315(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2316 github.com/goccy/llamawasm2go/p2.Fn2316
func Fn2316(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2401 github.com/goccy/llamawasm2go/p2.Fn2401
func Fn2401(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2520 github.com/goccy/llamawasm2go/p2.Fn2520
func Fn2520(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2521 github.com/goccy/llamawasm2go/p0.Fn2521
func Fn2521(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2522 github.com/goccy/llamawasm2go/p2.Fn2522
func Fn2522(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2523 github.com/goccy/llamawasm2go/p2.Fn2523
func Fn2523(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2527 github.com/goccy/llamawasm2go/p2.Fn2527
func Fn2527(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2552 github.com/goccy/llamawasm2go/p2.Fn2552
func Fn2552(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2684 github.com/goccy/llamawasm2go/p2.Fn2684
func Fn2684(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2719 github.com/goccy/llamawasm2go/p2.Fn2719
func Fn2719(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2733 github.com/goccy/llamawasm2go/p2.Fn2733
func Fn2733(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2743 github.com/goccy/llamawasm2go/p2.Fn2743
func Fn2743(m *base.Module, l0 int32)

//go:linkname Fn2745 github.com/goccy/llamawasm2go/p2.Fn2745
func Fn2745(m *base.Module, l0 int64) int64

//go:linkname Fn2746 github.com/goccy/llamawasm2go/p2.Fn2746
func Fn2746(m *base.Module, l0 int64)

//go:linkname Fn2749 github.com/goccy/llamawasm2go/p2.Fn2749
func Fn2749(m *base.Module, l0 int64)

//go:linkname Fn2750 github.com/goccy/llamawasm2go/p2.Fn2750
func Fn2750(m *base.Module, l0 int64)

//go:linkname Fn2752 github.com/goccy/llamawasm2go/p2.Fn2752
func Fn2752(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2757 github.com/goccy/llamawasm2go/p2.Fn2757
func Fn2757(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2763 github.com/goccy/llamawasm2go/p2.Fn2763
func Fn2763(m *base.Module) int32

//go:linkname Fn2771 github.com/goccy/llamawasm2go/p2.Fn2771
func Fn2771(m *base.Module, l0 float64) float32

//go:linkname Fn2772 github.com/goccy/llamawasm2go/p2.Fn2772
func Fn2772(m *base.Module, l0 float64) float32

//go:linkname Fn2779 github.com/goccy/llamawasm2go/p2.Fn2779
func Fn2779(m *base.Module, l0 float32) float32

//go:linkname Fn2783 github.com/goccy/llamawasm2go/p2.Fn2783
func Fn2783(m *base.Module, l0 float32) float32

//go:linkname Fn2786 github.com/goccy/llamawasm2go/p2.Fn2786
func Fn2786(m *base.Module, l0 float32) float32

//go:linkname Fn2801 github.com/goccy/llamawasm2go/p2.Fn2801
func Fn2801(m *base.Module, l0 int64) int32

//go:linkname Fn2802 github.com/goccy/llamawasm2go/p2.Fn2802
func Fn2802(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2804 github.com/goccy/llamawasm2go/p2.Fn2804
func Fn2804(m *base.Module, l0 int64)

//go:linkname Fn2805 github.com/goccy/llamawasm2go/p2.Fn2805
func Fn2805(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2806 github.com/goccy/llamawasm2go/p2.Fn2806
func Fn2806(m *base.Module, l0 int64) int32

//go:linkname Fn2813 github.com/goccy/llamawasm2go/p2.Fn2813
func Fn2813(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2814 github.com/goccy/llamawasm2go/p2.Fn2814
func Fn2814(m *base.Module)

//go:linkname Fn2815 github.com/goccy/llamawasm2go/p2.Fn2815
func Fn2815(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2821 github.com/goccy/llamawasm2go/p2.Fn2821
func Fn2821(m *base.Module, l0 int64) int32

//go:linkname Fn2822 github.com/goccy/llamawasm2go/p2.Fn2822
func Fn2822(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2827 github.com/goccy/llamawasm2go/p2.Fn2827
func Fn2827(m *base.Module, l0 int64) int32

//go:linkname Fn2831 github.com/goccy/llamawasm2go/p2.Fn2831
func Fn2831(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn2832 github.com/goccy/llamawasm2go/p2.Fn2832
func Fn2832(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2835 github.com/goccy/llamawasm2go/p2.Fn2835
func Fn2835(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2838 github.com/goccy/llamawasm2go/p2.Fn2838
func Fn2838(m *base.Module, l0 int64) int64

//go:linkname Fn2842 github.com/goccy/llamawasm2go/p2.Fn2842
func Fn2842(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn2852 github.com/goccy/llamawasm2go/p2.Fn2852
func Fn2852(m *base.Module)

//go:linkname Fn2853 github.com/goccy/llamawasm2go/p0.Fn2853
func Fn2853(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2855 github.com/goccy/llamawasm2go/p2.Fn2855
func Fn2855(m *base.Module)

//go:linkname Fn2857 github.com/goccy/llamawasm2go/p0.Fn2857
func Fn2857(m *base.Module, l0 int64) int64

//go:linkname Fn2859 github.com/goccy/llamawasm2go/p2.Fn2859
func Fn2859(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2860 github.com/goccy/llamawasm2go/p2.Fn2860
func Fn2860(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2863 github.com/goccy/llamawasm2go/p2.Fn2863
func Fn2863(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2892 github.com/goccy/llamawasm2go/p2.Fn2892
func Fn2892(m *base.Module, l0 int32)

//go:linkname Fn919rows github.com/goccy/llamawasm2go/p2.Fn919rows
func Fn919rows(m *base.Module)
