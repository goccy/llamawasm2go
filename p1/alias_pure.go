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

//go:linkname Fn798 github.com/goccy/llamawasm2go/p2.Fn798
func Fn798(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn864 github.com/goccy/llamawasm2go/p2.Fn864
func Fn864(m *base.Module) int64

//go:linkname Fn915 github.com/goccy/llamawasm2go/p2.Fn915
func Fn915(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn917 github.com/goccy/llamawasm2go/p2.Fn917
func Fn917(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn942 github.com/goccy/llamawasm2go/p2.Fn942
func Fn942(m *base.Module, l0 int64)

//go:linkname Fn947 github.com/goccy/llamawasm2go/p2.Fn947
func Fn947(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn992 github.com/goccy/llamawasm2go/p0.Fn992
func Fn992(m *base.Module, l0 int64) int64

//go:linkname Fn994 github.com/goccy/llamawasm2go/p2.Fn994
func Fn994(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1072 github.com/goccy/llamawasm2go/p2.Fn1072
func Fn1072(m *base.Module, l0 int64)

//go:linkname Fn1094 github.com/goccy/llamawasm2go/p2.Fn1094
func Fn1094(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1102 github.com/goccy/llamawasm2go/p2.Fn1102
func Fn1102(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1112 github.com/goccy/llamawasm2go/p2.Fn1112
func Fn1112(m *base.Module, l0 int64) int64

//go:linkname Fn1142 github.com/goccy/llamawasm2go/p2.Fn1142
func Fn1142(m *base.Module, l0 int32) int64

//go:linkname Fn1153 github.com/goccy/llamawasm2go/p2.Fn1153
func Fn1153(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1154 github.com/goccy/llamawasm2go/p2.Fn1154
func Fn1154(m *base.Module, l0 int64)

//go:linkname Fn1156 github.com/goccy/llamawasm2go/p2.Fn1156
func Fn1156(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1176 github.com/goccy/llamawasm2go/p2.Fn1176
func Fn1176(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1177 github.com/goccy/llamawasm2go/p2.Fn1177
func Fn1177(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1178 github.com/goccy/llamawasm2go/p2.Fn1178
func Fn1178(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1179 github.com/goccy/llamawasm2go/p2.Fn1179
func Fn1179(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1180 github.com/goccy/llamawasm2go/p2.Fn1180
func Fn1180(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1181 github.com/goccy/llamawasm2go/p2.Fn1181
func Fn1181(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1182 github.com/goccy/llamawasm2go/p2.Fn1182
func Fn1182(m *base.Module, l0 int64) int64

//go:linkname Fn1185 github.com/goccy/llamawasm2go/p2.Fn1185
func Fn1185(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1189 github.com/goccy/llamawasm2go/p2.Fn1189
func Fn1189(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1190 github.com/goccy/llamawasm2go/p2.Fn1190
func Fn1190(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1191 github.com/goccy/llamawasm2go/p2.Fn1191
func Fn1191(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1192 github.com/goccy/llamawasm2go/p2.Fn1192
func Fn1192(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1193 github.com/goccy/llamawasm2go/p2.Fn1193
func Fn1193(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1194 github.com/goccy/llamawasm2go/p2.Fn1194
func Fn1194(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1195 github.com/goccy/llamawasm2go/p2.Fn1195
func Fn1195(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1196 github.com/goccy/llamawasm2go/p2.Fn1196
func Fn1196(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1197 github.com/goccy/llamawasm2go/p2.Fn1197
func Fn1197(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1198 github.com/goccy/llamawasm2go/p2.Fn1198
func Fn1198(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1199 github.com/goccy/llamawasm2go/p2.Fn1199
func Fn1199(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1200 github.com/goccy/llamawasm2go/p2.Fn1200
func Fn1200(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1201 github.com/goccy/llamawasm2go/p2.Fn1201
func Fn1201(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1202 github.com/goccy/llamawasm2go/p2.Fn1202
func Fn1202(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1203 github.com/goccy/llamawasm2go/p2.Fn1203
func Fn1203(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1204 github.com/goccy/llamawasm2go/p2.Fn1204
func Fn1204(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1205 github.com/goccy/llamawasm2go/p2.Fn1205
func Fn1205(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1206 github.com/goccy/llamawasm2go/p2.Fn1206
func Fn1206(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1207 github.com/goccy/llamawasm2go/p2.Fn1207
func Fn1207(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1208 github.com/goccy/llamawasm2go/p2.Fn1208
func Fn1208(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1209 github.com/goccy/llamawasm2go/p2.Fn1209
func Fn1209(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1210 github.com/goccy/llamawasm2go/p2.Fn1210
func Fn1210(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1211 github.com/goccy/llamawasm2go/p2.Fn1211
func Fn1211(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1212 github.com/goccy/llamawasm2go/p2.Fn1212
func Fn1212(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1213 github.com/goccy/llamawasm2go/p2.Fn1213
func Fn1213(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1214 github.com/goccy/llamawasm2go/p2.Fn1214
func Fn1214(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1215 github.com/goccy/llamawasm2go/p2.Fn1215
func Fn1215(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1217 github.com/goccy/llamawasm2go/p2.Fn1217
func Fn1217(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1221 github.com/goccy/llamawasm2go/p2.Fn1221
func Fn1221(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1222 github.com/goccy/llamawasm2go/p2.Fn1222
func Fn1222(m *base.Module)

//go:linkname Fn1223 github.com/goccy/llamawasm2go/p2.Fn1223
func Fn1223(m *base.Module)

//go:linkname Fn1224 github.com/goccy/llamawasm2go/p0.Fn1224
func Fn1224(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1232 github.com/goccy/llamawasm2go/p2.Fn1232
func Fn1232(m *base.Module)

//go:linkname Fn1234 github.com/goccy/llamawasm2go/p2.Fn1234
func Fn1234(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1242 github.com/goccy/llamawasm2go/p2.Fn1242
func Fn1242(m *base.Module, l0 int64)

//go:linkname Fn1249 github.com/goccy/llamawasm2go/p2.Fn1249
func Fn1249(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1252 github.com/goccy/llamawasm2go/p2.Fn1252
func Fn1252(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1254 github.com/goccy/llamawasm2go/p2.Fn1254
func Fn1254(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1256 github.com/goccy/llamawasm2go/p2.Fn1256
func Fn1256(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1258 github.com/goccy/llamawasm2go/p2.Fn1258
func Fn1258(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1259 github.com/goccy/llamawasm2go/p2.Fn1259
func Fn1259(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1260 github.com/goccy/llamawasm2go/p2.Fn1260
func Fn1260(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1265 github.com/goccy/llamawasm2go/p2.Fn1265
func Fn1265(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1266 github.com/goccy/llamawasm2go/p2.Fn1266
func Fn1266(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1267 github.com/goccy/llamawasm2go/p2.Fn1267
func Fn1267(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1268 github.com/goccy/llamawasm2go/p0.Fn1268
func Fn1268(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1270 github.com/goccy/llamawasm2go/p2.Fn1270
func Fn1270(m *base.Module, l0 int64)

//go:linkname Fn1271 github.com/goccy/llamawasm2go/p2.Fn1271
func Fn1271(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1273 github.com/goccy/llamawasm2go/p2.Fn1273
func Fn1273(m *base.Module, l0 int64) int64

//go:linkname Fn1274 github.com/goccy/llamawasm2go/p2.Fn1274
func Fn1274(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1275 github.com/goccy/llamawasm2go/p2.Fn1275
func Fn1275(m *base.Module, l0 int64)

//go:linkname Fn1276 github.com/goccy/llamawasm2go/p2.Fn1276
func Fn1276(m *base.Module, l0 int64)

//go:linkname Fn1277 github.com/goccy/llamawasm2go/p2.Fn1277
func Fn1277(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1278 github.com/goccy/llamawasm2go/p2.Fn1278
func Fn1278(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1281 github.com/goccy/llamawasm2go/p2.Fn1281
func Fn1281(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1288 github.com/goccy/llamawasm2go/p2.Fn1288
func Fn1288(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1290 github.com/goccy/llamawasm2go/p0.Fn1290
func Fn1290(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1292 github.com/goccy/llamawasm2go/p2.Fn1292
func Fn1292(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1293 github.com/goccy/llamawasm2go/p2.Fn1293
func Fn1293(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1295 github.com/goccy/llamawasm2go/p0.Fn1295
func Fn1295(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1296 github.com/goccy/llamawasm2go/p2.Fn1296
func Fn1296(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1300 github.com/goccy/llamawasm2go/p2.Fn1300
func Fn1300(m *base.Module, l0 int64)

//go:linkname Fn1303 github.com/goccy/llamawasm2go/p2.Fn1303
func Fn1303(m *base.Module, l0 int64)

//go:linkname Fn1305 github.com/goccy/llamawasm2go/p2.Fn1305
func Fn1305(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1306 github.com/goccy/llamawasm2go/p2.Fn1306
func Fn1306(m *base.Module, l0 int64) int64

//go:linkname Fn1307 github.com/goccy/llamawasm2go/p2.Fn1307
func Fn1307(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1308 github.com/goccy/llamawasm2go/p2.Fn1308
func Fn1308(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1309 github.com/goccy/llamawasm2go/p2.Fn1309
func Fn1309(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1310 github.com/goccy/llamawasm2go/p2.Fn1310
func Fn1310(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1311 github.com/goccy/llamawasm2go/p2.Fn1311
func Fn1311(m *base.Module, l0 int64)

//go:linkname Fn1314 github.com/goccy/llamawasm2go/p2.Fn1314
func Fn1314(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1316 github.com/goccy/llamawasm2go/p2.Fn1316
func Fn1316(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1319 github.com/goccy/llamawasm2go/p2.Fn1319
func Fn1319(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1326 github.com/goccy/llamawasm2go/p2.Fn1326
func Fn1326(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1327 github.com/goccy/llamawasm2go/p2.Fn1327
func Fn1327(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1333 github.com/goccy/llamawasm2go/p2.Fn1333
func Fn1333(m *base.Module, l0 int64)

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1336 github.com/goccy/llamawasm2go/p2.Fn1336
func Fn1336(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1337 github.com/goccy/llamawasm2go/p2.Fn1337
func Fn1337(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1338 github.com/goccy/llamawasm2go/p2.Fn1338
func Fn1338(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1380 github.com/goccy/llamawasm2go/p2.Fn1380
func Fn1380(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1393 github.com/goccy/llamawasm2go/p2.Fn1393
func Fn1393(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p2.Fn1397
func Fn1397(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1399 github.com/goccy/llamawasm2go/p2.Fn1399
func Fn1399(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p2.Fn1401
func Fn1401(m *base.Module, l0 int64)

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p2.Fn1402
func Fn1402(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1404 github.com/goccy/llamawasm2go/p2.Fn1404
func Fn1404(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p2.Fn1406
func Fn1406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p2.Fn1407
func Fn1407(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p2.Fn1408
func Fn1408(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1410 github.com/goccy/llamawasm2go/p2.Fn1410
func Fn1410(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1412 github.com/goccy/llamawasm2go/p2.Fn1412
func Fn1412(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p2.Fn1413
func Fn1413(m *base.Module, l0 int64) int64

//go:linkname Fn1414 github.com/goccy/llamawasm2go/p2.Fn1414
func Fn1414(m *base.Module, l0 int64) int64

//go:linkname Fn1415 github.com/goccy/llamawasm2go/p2.Fn1415
func Fn1415(m *base.Module, l0 int64) int64

//go:linkname Fn1416 github.com/goccy/llamawasm2go/p2.Fn1416
func Fn1416(m *base.Module, l0 int64) int64

//go:linkname Fn1417 github.com/goccy/llamawasm2go/p2.Fn1417
func Fn1417(m *base.Module, l0 int64) int64

//go:linkname Fn1418 github.com/goccy/llamawasm2go/p2.Fn1418
func Fn1418(m *base.Module, l0 int64) int64

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64) int64

//go:linkname Fn1421 github.com/goccy/llamawasm2go/p2.Fn1421
func Fn1421(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1423 github.com/goccy/llamawasm2go/p2.Fn1423
func Fn1423(m *base.Module, l0 int64) int64

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1425 github.com/goccy/llamawasm2go/p2.Fn1425
func Fn1425(m *base.Module, l0 int64) int64

//go:linkname Fn1427 github.com/goccy/llamawasm2go/p2.Fn1427
func Fn1427(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1428 github.com/goccy/llamawasm2go/p2.Fn1428
func Fn1428(m *base.Module, l0 int64) int64

//go:linkname Fn1430 github.com/goccy/llamawasm2go/p2.Fn1430
func Fn1430(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1431 github.com/goccy/llamawasm2go/p2.Fn1431
func Fn1431(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1432 github.com/goccy/llamawasm2go/p2.Fn1432
func Fn1432(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1433 github.com/goccy/llamawasm2go/p2.Fn1433
func Fn1433(m *base.Module, l0 int64) int64

//go:linkname Fn1434 github.com/goccy/llamawasm2go/p2.Fn1434
func Fn1434(m *base.Module, l0 int64) int64

//go:linkname Fn1435 github.com/goccy/llamawasm2go/p2.Fn1435
func Fn1435(m *base.Module, l0 int64) int64

//go:linkname Fn1441 github.com/goccy/llamawasm2go/p2.Fn1441
func Fn1441(m *base.Module, l0 int64) int64

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1446 github.com/goccy/llamawasm2go/p2.Fn1446
func Fn1446(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int64) int64

//go:linkname Fn1448 github.com/goccy/llamawasm2go/p2.Fn1448
func Fn1448(m *base.Module, l0 int64) int64

//go:linkname Fn1449 github.com/goccy/llamawasm2go/p2.Fn1449
func Fn1449(m *base.Module, l0 int64) int64

//go:linkname Fn1450 github.com/goccy/llamawasm2go/p2.Fn1450
func Fn1450(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1451 github.com/goccy/llamawasm2go/p2.Fn1451
func Fn1451(m *base.Module, l0 int64)

//go:linkname Fn1480 github.com/goccy/llamawasm2go/p2.Fn1480
func Fn1480(m *base.Module, l0 int64) int64

//go:linkname Fn1486 github.com/goccy/llamawasm2go/p2.Fn1486
func Fn1486(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1487 github.com/goccy/llamawasm2go/p2.Fn1487
func Fn1487(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1488 github.com/goccy/llamawasm2go/p2.Fn1488
func Fn1488(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1490 github.com/goccy/llamawasm2go/p2.Fn1490
func Fn1490(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1491 github.com/goccy/llamawasm2go/p2.Fn1491
func Fn1491(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1492 github.com/goccy/llamawasm2go/p2.Fn1492
func Fn1492(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1493 github.com/goccy/llamawasm2go/p2.Fn1493
func Fn1493(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1494 github.com/goccy/llamawasm2go/p2.Fn1494
func Fn1494(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1495 github.com/goccy/llamawasm2go/p2.Fn1495
func Fn1495(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1497 github.com/goccy/llamawasm2go/p2.Fn1497
func Fn1497(m *base.Module, l0 int64) int32

//go:linkname Fn1498 github.com/goccy/llamawasm2go/p2.Fn1498
func Fn1498(m *base.Module, l0 int64) int32

//go:linkname Fn1499 github.com/goccy/llamawasm2go/p2.Fn1499
func Fn1499(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1500 github.com/goccy/llamawasm2go/p2.Fn1500
func Fn1500(m *base.Module, l0 int64) int32

//go:linkname Fn1501 github.com/goccy/llamawasm2go/p2.Fn1501
func Fn1501(m *base.Module, l0 int64) int32

//go:linkname Fn1504 github.com/goccy/llamawasm2go/p2.Fn1504
func Fn1504(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1505 github.com/goccy/llamawasm2go/p2.Fn1505
func Fn1505(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1506 github.com/goccy/llamawasm2go/p2.Fn1506
func Fn1506(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1507 github.com/goccy/llamawasm2go/p2.Fn1507
func Fn1507(m *base.Module)

//go:linkname Fn1508 github.com/goccy/llamawasm2go/p2.Fn1508
func Fn1508(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1510 github.com/goccy/llamawasm2go/p2.Fn1510
func Fn1510(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1512 github.com/goccy/llamawasm2go/p2.Fn1512
func Fn1512(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1513 github.com/goccy/llamawasm2go/p2.Fn1513
func Fn1513(m *base.Module, l0 int64)

//go:linkname Fn1517 github.com/goccy/llamawasm2go/p2.Fn1517
func Fn1517(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1518 github.com/goccy/llamawasm2go/p2.Fn1518
func Fn1518(m *base.Module, l0 int64)

//go:linkname Fn1521 github.com/goccy/llamawasm2go/p2.Fn1521
func Fn1521(m *base.Module, l0 int64)

//go:linkname Fn1530 github.com/goccy/llamawasm2go/p2.Fn1530
func Fn1530(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1531 github.com/goccy/llamawasm2go/p2.Fn1531
func Fn1531(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1540 github.com/goccy/llamawasm2go/p2.Fn1540
func Fn1540(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p0.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1543 github.com/goccy/llamawasm2go/p2.Fn1543
func Fn1543(m *base.Module, l0 int64)

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1552 github.com/goccy/llamawasm2go/p2.Fn1552
func Fn1552(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1560 github.com/goccy/llamawasm2go/p2.Fn1560
func Fn1560(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1567 github.com/goccy/llamawasm2go/p2.Fn1567
func Fn1567(m *base.Module, l0 int64)

//go:linkname Fn1570 github.com/goccy/llamawasm2go/p2.Fn1570
func Fn1570(m *base.Module, l0 int64) int32

//go:linkname Fn1580 github.com/goccy/llamawasm2go/p2.Fn1580
func Fn1580(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1582 github.com/goccy/llamawasm2go/p2.Fn1582
func Fn1582(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p2.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1602 github.com/goccy/llamawasm2go/p2.Fn1602
func Fn1602(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1615 github.com/goccy/llamawasm2go/p2.Fn1615
func Fn1615(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1624 github.com/goccy/llamawasm2go/p2.Fn1624
func Fn1624(m *base.Module, l0 int64)

//go:linkname Fn1658 github.com/goccy/llamawasm2go/p2.Fn1658
func Fn1658(m *base.Module, l0 int64)

//go:linkname Fn1661 github.com/goccy/llamawasm2go/p2.Fn1661
func Fn1661(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1675 github.com/goccy/llamawasm2go/p2.Fn1675
func Fn1675(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1676 github.com/goccy/llamawasm2go/p2.Fn1676
func Fn1676(m *base.Module, l0 int64) int64

//go:linkname Fn1677 github.com/goccy/llamawasm2go/p2.Fn1677
func Fn1677(m *base.Module, l0 int64)

//go:linkname Fn1681 github.com/goccy/llamawasm2go/p0.Fn1681
func Fn1681(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1682 github.com/goccy/llamawasm2go/p2.Fn1682
func Fn1682(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1683 github.com/goccy/llamawasm2go/p2.Fn1683
func Fn1683(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1684 github.com/goccy/llamawasm2go/p2.Fn1684
func Fn1684(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1685 github.com/goccy/llamawasm2go/p2.Fn1685
func Fn1685(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1703 github.com/goccy/llamawasm2go/p2.Fn1703
func Fn1703(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1713 github.com/goccy/llamawasm2go/p2.Fn1713
func Fn1713(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1715 github.com/goccy/llamawasm2go/p2.Fn1715
func Fn1715(m *base.Module, l0 int64) int64

//go:linkname Fn1716 github.com/goccy/llamawasm2go/p2.Fn1716
func Fn1716(m *base.Module, l0 int64)

//go:linkname Fn1719 github.com/goccy/llamawasm2go/p0.Fn1719
func Fn1719(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1721 github.com/goccy/llamawasm2go/p2.Fn1721
func Fn1721(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1732 github.com/goccy/llamawasm2go/p2.Fn1732
func Fn1732(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1735 github.com/goccy/llamawasm2go/p2.Fn1735
func Fn1735(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1736 github.com/goccy/llamawasm2go/p2.Fn1736
func Fn1736(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1746 github.com/goccy/llamawasm2go/p2.Fn1746
func Fn1746(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1776 github.com/goccy/llamawasm2go/p2.Fn1776
func Fn1776(m *base.Module, l0 int64)

//go:linkname Fn1777 github.com/goccy/llamawasm2go/p2.Fn1777
func Fn1777(m *base.Module, l0 int64)

//go:linkname Fn1788 github.com/goccy/llamawasm2go/p2.Fn1788
func Fn1788(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1825 github.com/goccy/llamawasm2go/p2.Fn1825
func Fn1825(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1829 github.com/goccy/llamawasm2go/p2.Fn1829
func Fn1829(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1831 github.com/goccy/llamawasm2go/p2.Fn1831
func Fn1831(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1832 github.com/goccy/llamawasm2go/p2.Fn1832
func Fn1832(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1833 github.com/goccy/llamawasm2go/p2.Fn1833
func Fn1833(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1834 github.com/goccy/llamawasm2go/p2.Fn1834
func Fn1834(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1835 github.com/goccy/llamawasm2go/p2.Fn1835
func Fn1835(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1840 github.com/goccy/llamawasm2go/p2.Fn1840
func Fn1840(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1846 github.com/goccy/llamawasm2go/p2.Fn1846
func Fn1846(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1848 github.com/goccy/llamawasm2go/p2.Fn1848
func Fn1848(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1851 github.com/goccy/llamawasm2go/p2.Fn1851
func Fn1851(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1852 github.com/goccy/llamawasm2go/p2.Fn1852
func Fn1852(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1853 github.com/goccy/llamawasm2go/p2.Fn1853
func Fn1853(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1856 github.com/goccy/llamawasm2go/p2.Fn1856
func Fn1856(m *base.Module, l0 int64)

//go:linkname Fn1867 github.com/goccy/llamawasm2go/p2.Fn1867
func Fn1867(m *base.Module, l0 int64)

//go:linkname Fn1869 github.com/goccy/llamawasm2go/p2.Fn1869
func Fn1869(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1870 github.com/goccy/llamawasm2go/p2.Fn1870
func Fn1870(m *base.Module)

//go:linkname Fn1871 github.com/goccy/llamawasm2go/p2.Fn1871
func Fn1871(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1876 github.com/goccy/llamawasm2go/p2.Fn1876
func Fn1876(m *base.Module)

//go:linkname Fn1885 github.com/goccy/llamawasm2go/p2.Fn1885
func Fn1885(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1886 github.com/goccy/llamawasm2go/p2.Fn1886
func Fn1886(m *base.Module, l0 int64) int64

//go:linkname Fn1887 github.com/goccy/llamawasm2go/p2.Fn1887
func Fn1887(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1888 github.com/goccy/llamawasm2go/p2.Fn1888
func Fn1888(m *base.Module)

//go:linkname Fn1890 github.com/goccy/llamawasm2go/p2.Fn1890
func Fn1890(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1891 github.com/goccy/llamawasm2go/p2.Fn1891
func Fn1891(m *base.Module)

//go:linkname Fn1905 github.com/goccy/llamawasm2go/p2.Fn1905
func Fn1905(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1907 github.com/goccy/llamawasm2go/p2.Fn1907
func Fn1907(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1908 github.com/goccy/llamawasm2go/p2.Fn1908
func Fn1908(m *base.Module)

//go:linkname Fn1909 github.com/goccy/llamawasm2go/p2.Fn1909
func Fn1909(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1910 github.com/goccy/llamawasm2go/p2.Fn1910
func Fn1910(m *base.Module)

//go:linkname Fn1912 github.com/goccy/llamawasm2go/p2.Fn1912
func Fn1912(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1913 github.com/goccy/llamawasm2go/p2.Fn1913
func Fn1913(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1914 github.com/goccy/llamawasm2go/p2.Fn1914
func Fn1914(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1915 github.com/goccy/llamawasm2go/p2.Fn1915
func Fn1915(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1916 github.com/goccy/llamawasm2go/p2.Fn1916
func Fn1916(m *base.Module)

//go:linkname Fn1917 github.com/goccy/llamawasm2go/p2.Fn1917
func Fn1917(m *base.Module)

//go:linkname Fn1918 github.com/goccy/llamawasm2go/p2.Fn1918
func Fn1918(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn1919 github.com/goccy/llamawasm2go/p2.Fn1919
func Fn1919(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1920 github.com/goccy/llamawasm2go/p2.Fn1920
func Fn1920(m *base.Module)

//go:linkname Fn1939 github.com/goccy/llamawasm2go/p2.Fn1939
func Fn1939(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1940 github.com/goccy/llamawasm2go/p2.Fn1940
func Fn1940(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1941 github.com/goccy/llamawasm2go/p2.Fn1941
func Fn1941(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1943 github.com/goccy/llamawasm2go/p2.Fn1943
func Fn1943(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1944 github.com/goccy/llamawasm2go/p2.Fn1944
func Fn1944(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1945 github.com/goccy/llamawasm2go/p2.Fn1945
func Fn1945(m *base.Module)

//go:linkname Fn1946 github.com/goccy/llamawasm2go/p2.Fn1946
func Fn1946(m *base.Module)

//go:linkname Fn1947 github.com/goccy/llamawasm2go/p2.Fn1947
func Fn1947(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1948 github.com/goccy/llamawasm2go/p2.Fn1948
func Fn1948(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1949 github.com/goccy/llamawasm2go/p2.Fn1949
func Fn1949(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1951 github.com/goccy/llamawasm2go/p2.Fn1951
func Fn1951(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1953 github.com/goccy/llamawasm2go/p2.Fn1953
func Fn1953(m *base.Module, l0 int64)

//go:linkname Fn1954 github.com/goccy/llamawasm2go/p2.Fn1954
func Fn1954(m *base.Module)

//go:linkname Fn1971 github.com/goccy/llamawasm2go/p2.Fn1971
func Fn1971(m *base.Module, l0 int64)

//go:linkname Fn1972 github.com/goccy/llamawasm2go/p2.Fn1972
func Fn1972(m *base.Module, l0 int64)

//go:linkname Fn1973 github.com/goccy/llamawasm2go/p2.Fn1973
func Fn1973(m *base.Module, l0 int64)

//go:linkname Fn1975 github.com/goccy/llamawasm2go/p2.Fn1975
func Fn1975(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1976 github.com/goccy/llamawasm2go/p2.Fn1976
func Fn1976(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2015 github.com/goccy/llamawasm2go/p2.Fn2015
func Fn2015(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2027 github.com/goccy/llamawasm2go/p2.Fn2027
func Fn2027(m *base.Module, l0 int64) int64

//go:linkname Fn2030 github.com/goccy/llamawasm2go/p2.Fn2030
func Fn2030(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2034 github.com/goccy/llamawasm2go/p2.Fn2034
func Fn2034(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2038 github.com/goccy/llamawasm2go/p2.Fn2038
func Fn2038(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2051 github.com/goccy/llamawasm2go/p2.Fn2051
func Fn2051(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2052 github.com/goccy/llamawasm2go/p2.Fn2052
func Fn2052(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2053 github.com/goccy/llamawasm2go/p2.Fn2053
func Fn2053(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2065 github.com/goccy/llamawasm2go/p2.Fn2065
func Fn2065(m *base.Module, l0 int64) int64

//go:linkname Fn2066 github.com/goccy/llamawasm2go/p2.Fn2066
func Fn2066(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2069 github.com/goccy/llamawasm2go/p2.Fn2069
func Fn2069(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2071 github.com/goccy/llamawasm2go/p2.Fn2071
func Fn2071(m *base.Module)

//go:linkname Fn2074 github.com/goccy/llamawasm2go/p2.Fn2074
func Fn2074(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2083 github.com/goccy/llamawasm2go/p2.Fn2083
func Fn2083(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2087 github.com/goccy/llamawasm2go/p2.Fn2087
func Fn2087(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2088 github.com/goccy/llamawasm2go/p2.Fn2088
func Fn2088(m *base.Module, l0 int64)

//go:linkname Fn2089 github.com/goccy/llamawasm2go/p2.Fn2089
func Fn2089(m *base.Module, l0 int64)

//go:linkname Fn2090 github.com/goccy/llamawasm2go/p2.Fn2090
func Fn2090(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2092 github.com/goccy/llamawasm2go/p2.Fn2092
func Fn2092(m *base.Module, l0 int64)

//go:linkname Fn2093 github.com/goccy/llamawasm2go/p2.Fn2093
func Fn2093(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2098 github.com/goccy/llamawasm2go/p2.Fn2098
func Fn2098(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2101 github.com/goccy/llamawasm2go/p2.Fn2101
func Fn2101(m *base.Module, l0 int64) int64

//go:linkname Fn2102 github.com/goccy/llamawasm2go/p2.Fn2102
func Fn2102(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2103 github.com/goccy/llamawasm2go/p2.Fn2103
func Fn2103(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2110 github.com/goccy/llamawasm2go/p2.Fn2110
func Fn2110(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2111 github.com/goccy/llamawasm2go/p2.Fn2111
func Fn2111(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2113 github.com/goccy/llamawasm2go/p2.Fn2113
func Fn2113(m *base.Module) int64

//go:linkname Fn2115 github.com/goccy/llamawasm2go/p2.Fn2115
func Fn2115(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2116 github.com/goccy/llamawasm2go/p2.Fn2116
func Fn2116(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2117 github.com/goccy/llamawasm2go/p2.Fn2117
func Fn2117(m *base.Module) int64

//go:linkname Fn2119 github.com/goccy/llamawasm2go/p2.Fn2119
func Fn2119(m *base.Module, l0 int32) int64

//go:linkname Fn2120 github.com/goccy/llamawasm2go/p2.Fn2120
func Fn2120(m *base.Module, l0 int32) int32

//go:linkname Fn2121 github.com/goccy/llamawasm2go/p2.Fn2121
func Fn2121(m *base.Module, l0 int32) int64

//go:linkname Fn2122 github.com/goccy/llamawasm2go/p2.Fn2122
func Fn2122(m *base.Module, l0 float32) int64

//go:linkname Fn2123 github.com/goccy/llamawasm2go/p2.Fn2123
func Fn2123(m *base.Module, l0 float32) int64

//go:linkname Fn2124 github.com/goccy/llamawasm2go/p2.Fn2124
func Fn2124(m *base.Module, l0 float32) int64

//go:linkname Fn2126 github.com/goccy/llamawasm2go/p2.Fn2126
func Fn2126(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2127 github.com/goccy/llamawasm2go/p2.Fn2127
func Fn2127(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2168 github.com/goccy/llamawasm2go/p2.Fn2168
func Fn2168(m *base.Module, l0 int64)

//go:linkname Fn2170 github.com/goccy/llamawasm2go/p2.Fn2170
func Fn2170(m *base.Module, l0 int64)

//go:linkname Fn2178 github.com/goccy/llamawasm2go/p2.Fn2178
func Fn2178(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2212 github.com/goccy/llamawasm2go/p2.Fn2212
func Fn2212(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2223 github.com/goccy/llamawasm2go/p2.Fn2223
func Fn2223(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2226 github.com/goccy/llamawasm2go/p2.Fn2226
func Fn2226(m *base.Module, l0 int64)

//go:linkname Fn2227 github.com/goccy/llamawasm2go/p2.Fn2227
func Fn2227(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2228 github.com/goccy/llamawasm2go/p2.Fn2228
func Fn2228(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2229 github.com/goccy/llamawasm2go/p2.Fn2229
func Fn2229(m *base.Module, l0 int64)

//go:linkname Fn2232 github.com/goccy/llamawasm2go/p2.Fn2232
func Fn2232(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2237 github.com/goccy/llamawasm2go/p2.Fn2237
func Fn2237(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2238 github.com/goccy/llamawasm2go/p2.Fn2238
func Fn2238(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2251 github.com/goccy/llamawasm2go/p2.Fn2251
func Fn2251(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2252 github.com/goccy/llamawasm2go/p2.Fn2252
func Fn2252(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2262 github.com/goccy/llamawasm2go/p2.Fn2262
func Fn2262(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2265 github.com/goccy/llamawasm2go/p2.Fn2265
func Fn2265(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2267 github.com/goccy/llamawasm2go/p2.Fn2267
func Fn2267(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2269 github.com/goccy/llamawasm2go/p2.Fn2269
func Fn2269(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2272 github.com/goccy/llamawasm2go/p2.Fn2272
func Fn2272(m *base.Module)

//go:linkname Fn2276 github.com/goccy/llamawasm2go/p2.Fn2276
func Fn2276(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2278 github.com/goccy/llamawasm2go/p0.Fn2278
func Fn2278(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2293 github.com/goccy/llamawasm2go/p2.Fn2293
func Fn2293(m *base.Module, l0 int64)

//go:linkname Fn2294 github.com/goccy/llamawasm2go/p2.Fn2294
func Fn2294(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2299 github.com/goccy/llamawasm2go/p2.Fn2299
func Fn2299(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2304 github.com/goccy/llamawasm2go/p2.Fn2304
func Fn2304(m *base.Module, l0 int64) int64

//go:linkname Fn2305 github.com/goccy/llamawasm2go/p2.Fn2305
func Fn2305(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2317 github.com/goccy/llamawasm2go/p2.Fn2317
func Fn2317(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2318 github.com/goccy/llamawasm2go/p2.Fn2318
func Fn2318(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2319 github.com/goccy/llamawasm2go/p2.Fn2319
func Fn2319(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2320 github.com/goccy/llamawasm2go/p2.Fn2320
func Fn2320(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2321 github.com/goccy/llamawasm2go/p2.Fn2321
func Fn2321(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2322 github.com/goccy/llamawasm2go/p2.Fn2322
func Fn2322(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2340 github.com/goccy/llamawasm2go/p2.Fn2340
func Fn2340(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2341 github.com/goccy/llamawasm2go/p2.Fn2341
func Fn2341(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2342 github.com/goccy/llamawasm2go/p2.Fn2342
func Fn2342(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2343 github.com/goccy/llamawasm2go/p2.Fn2343
func Fn2343(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2344 github.com/goccy/llamawasm2go/p2.Fn2344
func Fn2344(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2345 github.com/goccy/llamawasm2go/p2.Fn2345
func Fn2345(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2430 github.com/goccy/llamawasm2go/p2.Fn2430
func Fn2430(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2549 github.com/goccy/llamawasm2go/p2.Fn2549
func Fn2549(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2550 github.com/goccy/llamawasm2go/p0.Fn2550
func Fn2550(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2551 github.com/goccy/llamawasm2go/p2.Fn2551
func Fn2551(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2552 github.com/goccy/llamawasm2go/p2.Fn2552
func Fn2552(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2556 github.com/goccy/llamawasm2go/p2.Fn2556
func Fn2556(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2581 github.com/goccy/llamawasm2go/p2.Fn2581
func Fn2581(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2713 github.com/goccy/llamawasm2go/p2.Fn2713
func Fn2713(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2748 github.com/goccy/llamawasm2go/p2.Fn2748
func Fn2748(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2762 github.com/goccy/llamawasm2go/p2.Fn2762
func Fn2762(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2772 github.com/goccy/llamawasm2go/p2.Fn2772
func Fn2772(m *base.Module, l0 int32)

//go:linkname Fn2774 github.com/goccy/llamawasm2go/p2.Fn2774
func Fn2774(m *base.Module, l0 int64) int64

//go:linkname Fn2775 github.com/goccy/llamawasm2go/p2.Fn2775
func Fn2775(m *base.Module, l0 int64)

//go:linkname Fn2778 github.com/goccy/llamawasm2go/p2.Fn2778
func Fn2778(m *base.Module, l0 int64)

//go:linkname Fn2779 github.com/goccy/llamawasm2go/p2.Fn2779
func Fn2779(m *base.Module, l0 int64)

//go:linkname Fn2781 github.com/goccy/llamawasm2go/p2.Fn2781
func Fn2781(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2786 github.com/goccy/llamawasm2go/p2.Fn2786
func Fn2786(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2792 github.com/goccy/llamawasm2go/p2.Fn2792
func Fn2792(m *base.Module) int32

//go:linkname Fn2800 github.com/goccy/llamawasm2go/p2.Fn2800
func Fn2800(m *base.Module, l0 float64) float32

//go:linkname Fn2801 github.com/goccy/llamawasm2go/p2.Fn2801
func Fn2801(m *base.Module, l0 float64) float32

//go:linkname Fn2808 github.com/goccy/llamawasm2go/p2.Fn2808
func Fn2808(m *base.Module, l0 float32) float32

//go:linkname Fn2812 github.com/goccy/llamawasm2go/p2.Fn2812
func Fn2812(m *base.Module, l0 float32) float32

//go:linkname Fn2815 github.com/goccy/llamawasm2go/p2.Fn2815
func Fn2815(m *base.Module, l0 float32) float32

//go:linkname Fn2830 github.com/goccy/llamawasm2go/p2.Fn2830
func Fn2830(m *base.Module, l0 int64) int32

//go:linkname Fn2831 github.com/goccy/llamawasm2go/p2.Fn2831
func Fn2831(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2833 github.com/goccy/llamawasm2go/p2.Fn2833
func Fn2833(m *base.Module, l0 int64)

//go:linkname Fn2834 github.com/goccy/llamawasm2go/p2.Fn2834
func Fn2834(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2835 github.com/goccy/llamawasm2go/p2.Fn2835
func Fn2835(m *base.Module, l0 int64) int32

//go:linkname Fn2842 github.com/goccy/llamawasm2go/p2.Fn2842
func Fn2842(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2843 github.com/goccy/llamawasm2go/p2.Fn2843
func Fn2843(m *base.Module)

//go:linkname Fn2844 github.com/goccy/llamawasm2go/p2.Fn2844
func Fn2844(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2850 github.com/goccy/llamawasm2go/p2.Fn2850
func Fn2850(m *base.Module, l0 int64) int32

//go:linkname Fn2851 github.com/goccy/llamawasm2go/p2.Fn2851
func Fn2851(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2856 github.com/goccy/llamawasm2go/p2.Fn2856
func Fn2856(m *base.Module, l0 int64) int32

//go:linkname Fn2860 github.com/goccy/llamawasm2go/p2.Fn2860
func Fn2860(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn2861 github.com/goccy/llamawasm2go/p2.Fn2861
func Fn2861(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2864 github.com/goccy/llamawasm2go/p2.Fn2864
func Fn2864(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2867 github.com/goccy/llamawasm2go/p2.Fn2867
func Fn2867(m *base.Module, l0 int64) int64

//go:linkname Fn2871 github.com/goccy/llamawasm2go/p2.Fn2871
func Fn2871(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn2881 github.com/goccy/llamawasm2go/p2.Fn2881
func Fn2881(m *base.Module)

//go:linkname Fn2882 github.com/goccy/llamawasm2go/p0.Fn2882
func Fn2882(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2884 github.com/goccy/llamawasm2go/p2.Fn2884
func Fn2884(m *base.Module)

//go:linkname Fn2886 github.com/goccy/llamawasm2go/p0.Fn2886
func Fn2886(m *base.Module, l0 int64) int64

//go:linkname Fn2888 github.com/goccy/llamawasm2go/p2.Fn2888
func Fn2888(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2889 github.com/goccy/llamawasm2go/p2.Fn2889
func Fn2889(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2892 github.com/goccy/llamawasm2go/p2.Fn2892
func Fn2892(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn925rows github.com/goccy/llamawasm2go/p2.Fn925rows
func Fn925rows(m *base.Module)
