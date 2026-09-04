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

//go:linkname Fn429 github.com/goccy/llamawasm2go/p2.Fn429
func Fn429(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

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

//go:linkname Fn861 github.com/goccy/llamawasm2go/p2.Fn861
func Fn861(m *base.Module) int64

//go:linkname Fn898 github.com/goccy/llamawasm2go/p2.Fn898
func Fn898(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32)

//go:linkname Fn907 github.com/goccy/llamawasm2go/p2.Fn907
func Fn907(m *base.Module, l0 int64)

//go:linkname Fn914 github.com/goccy/llamawasm2go/p2.Fn914
func Fn914(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn916 github.com/goccy/llamawasm2go/p2.Fn916
func Fn916(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn931 github.com/goccy/llamawasm2go/p2.Fn931
func Fn931(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn944 github.com/goccy/llamawasm2go/p2.Fn944
func Fn944(m *base.Module, l0 int64)

//go:linkname Fn949 github.com/goccy/llamawasm2go/p2.Fn949
func Fn949(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn994 github.com/goccy/llamawasm2go/p0.Fn994
func Fn994(m *base.Module, l0 int64) int64

//go:linkname Fn996 github.com/goccy/llamawasm2go/p2.Fn996
func Fn996(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1073 github.com/goccy/llamawasm2go/p2.Fn1073
func Fn1073(m *base.Module, l0 int64)

//go:linkname Fn1095 github.com/goccy/llamawasm2go/p2.Fn1095
func Fn1095(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1103 github.com/goccy/llamawasm2go/p2.Fn1103
func Fn1103(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1113 github.com/goccy/llamawasm2go/p2.Fn1113
func Fn1113(m *base.Module, l0 int64) int64

//go:linkname Fn1143 github.com/goccy/llamawasm2go/p2.Fn1143
func Fn1143(m *base.Module, l0 int32) int64

//go:linkname Fn1154 github.com/goccy/llamawasm2go/p2.Fn1154
func Fn1154(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1155 github.com/goccy/llamawasm2go/p2.Fn1155
func Fn1155(m *base.Module, l0 int64)

//go:linkname Fn1157 github.com/goccy/llamawasm2go/p2.Fn1157
func Fn1157(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1177 github.com/goccy/llamawasm2go/p2.Fn1177
func Fn1177(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1178 github.com/goccy/llamawasm2go/p2.Fn1178
func Fn1178(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1179 github.com/goccy/llamawasm2go/p2.Fn1179
func Fn1179(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1180 github.com/goccy/llamawasm2go/p2.Fn1180
func Fn1180(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1181 github.com/goccy/llamawasm2go/p2.Fn1181
func Fn1181(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1182 github.com/goccy/llamawasm2go/p2.Fn1182
func Fn1182(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1183 github.com/goccy/llamawasm2go/p2.Fn1183
func Fn1183(m *base.Module, l0 int64) int64

//go:linkname Fn1186 github.com/goccy/llamawasm2go/p2.Fn1186
func Fn1186(m *base.Module, l0 int64, l1 int64) int64

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
func Fn1199(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1200 github.com/goccy/llamawasm2go/p2.Fn1200
func Fn1200(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1201 github.com/goccy/llamawasm2go/p2.Fn1201
func Fn1201(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1202 github.com/goccy/llamawasm2go/p2.Fn1202
func Fn1202(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1203 github.com/goccy/llamawasm2go/p2.Fn1203
func Fn1203(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1204 github.com/goccy/llamawasm2go/p2.Fn1204
func Fn1204(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1205 github.com/goccy/llamawasm2go/p2.Fn1205
func Fn1205(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1206 github.com/goccy/llamawasm2go/p2.Fn1206
func Fn1206(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1207 github.com/goccy/llamawasm2go/p2.Fn1207
func Fn1207(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1208 github.com/goccy/llamawasm2go/p2.Fn1208
func Fn1208(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1209 github.com/goccy/llamawasm2go/p2.Fn1209
func Fn1209(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1210 github.com/goccy/llamawasm2go/p2.Fn1210
func Fn1210(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1211 github.com/goccy/llamawasm2go/p2.Fn1211
func Fn1211(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1212 github.com/goccy/llamawasm2go/p2.Fn1212
func Fn1212(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1213 github.com/goccy/llamawasm2go/p2.Fn1213
func Fn1213(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1214 github.com/goccy/llamawasm2go/p2.Fn1214
func Fn1214(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1215 github.com/goccy/llamawasm2go/p2.Fn1215
func Fn1215(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1216 github.com/goccy/llamawasm2go/p2.Fn1216
func Fn1216(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1218 github.com/goccy/llamawasm2go/p2.Fn1218
func Fn1218(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1222 github.com/goccy/llamawasm2go/p2.Fn1222
func Fn1222(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1223 github.com/goccy/llamawasm2go/p2.Fn1223
func Fn1223(m *base.Module)

//go:linkname Fn1224 github.com/goccy/llamawasm2go/p2.Fn1224
func Fn1224(m *base.Module)

//go:linkname Fn1225 github.com/goccy/llamawasm2go/p0.Fn1225
func Fn1225(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1233 github.com/goccy/llamawasm2go/p2.Fn1233
func Fn1233(m *base.Module)

//go:linkname Fn1235 github.com/goccy/llamawasm2go/p2.Fn1235
func Fn1235(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1243 github.com/goccy/llamawasm2go/p2.Fn1243
func Fn1243(m *base.Module, l0 int64)

//go:linkname Fn1250 github.com/goccy/llamawasm2go/p2.Fn1250
func Fn1250(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1253 github.com/goccy/llamawasm2go/p2.Fn1253
func Fn1253(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1255 github.com/goccy/llamawasm2go/p2.Fn1255
func Fn1255(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1257 github.com/goccy/llamawasm2go/p2.Fn1257
func Fn1257(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1259 github.com/goccy/llamawasm2go/p2.Fn1259
func Fn1259(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1260 github.com/goccy/llamawasm2go/p2.Fn1260
func Fn1260(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1261 github.com/goccy/llamawasm2go/p2.Fn1261
func Fn1261(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1266 github.com/goccy/llamawasm2go/p2.Fn1266
func Fn1266(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1267 github.com/goccy/llamawasm2go/p2.Fn1267
func Fn1267(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1268 github.com/goccy/llamawasm2go/p2.Fn1268
func Fn1268(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1269 github.com/goccy/llamawasm2go/p0.Fn1269
func Fn1269(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1271 github.com/goccy/llamawasm2go/p2.Fn1271
func Fn1271(m *base.Module, l0 int64)

//go:linkname Fn1272 github.com/goccy/llamawasm2go/p2.Fn1272
func Fn1272(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1274 github.com/goccy/llamawasm2go/p2.Fn1274
func Fn1274(m *base.Module, l0 int64) int64

//go:linkname Fn1275 github.com/goccy/llamawasm2go/p2.Fn1275
func Fn1275(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1276 github.com/goccy/llamawasm2go/p2.Fn1276
func Fn1276(m *base.Module, l0 int64)

//go:linkname Fn1277 github.com/goccy/llamawasm2go/p2.Fn1277
func Fn1277(m *base.Module, l0 int64)

//go:linkname Fn1278 github.com/goccy/llamawasm2go/p2.Fn1278
func Fn1278(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1279 github.com/goccy/llamawasm2go/p2.Fn1279
func Fn1279(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1282 github.com/goccy/llamawasm2go/p2.Fn1282
func Fn1282(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1289 github.com/goccy/llamawasm2go/p2.Fn1289
func Fn1289(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1291 github.com/goccy/llamawasm2go/p0.Fn1291
func Fn1291(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1293 github.com/goccy/llamawasm2go/p2.Fn1293
func Fn1293(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1294 github.com/goccy/llamawasm2go/p2.Fn1294
func Fn1294(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1296 github.com/goccy/llamawasm2go/p0.Fn1296
func Fn1296(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1297 github.com/goccy/llamawasm2go/p2.Fn1297
func Fn1297(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1301 github.com/goccy/llamawasm2go/p2.Fn1301
func Fn1301(m *base.Module, l0 int64)

//go:linkname Fn1304 github.com/goccy/llamawasm2go/p2.Fn1304
func Fn1304(m *base.Module, l0 int64)

//go:linkname Fn1306 github.com/goccy/llamawasm2go/p2.Fn1306
func Fn1306(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1307 github.com/goccy/llamawasm2go/p2.Fn1307
func Fn1307(m *base.Module, l0 int64) int64

//go:linkname Fn1308 github.com/goccy/llamawasm2go/p2.Fn1308
func Fn1308(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1309 github.com/goccy/llamawasm2go/p2.Fn1309
func Fn1309(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1310 github.com/goccy/llamawasm2go/p2.Fn1310
func Fn1310(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1311 github.com/goccy/llamawasm2go/p2.Fn1311
func Fn1311(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1312 github.com/goccy/llamawasm2go/p2.Fn1312
func Fn1312(m *base.Module, l0 int64)

//go:linkname Fn1315 github.com/goccy/llamawasm2go/p2.Fn1315
func Fn1315(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1317 github.com/goccy/llamawasm2go/p2.Fn1317
func Fn1317(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1320 github.com/goccy/llamawasm2go/p2.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1327 github.com/goccy/llamawasm2go/p2.Fn1327
func Fn1327(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1328 github.com/goccy/llamawasm2go/p2.Fn1328
func Fn1328(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64)

//go:linkname Fn1335 github.com/goccy/llamawasm2go/p2.Fn1335
func Fn1335(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1337 github.com/goccy/llamawasm2go/p2.Fn1337
func Fn1337(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1338 github.com/goccy/llamawasm2go/p2.Fn1338
func Fn1338(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1339 github.com/goccy/llamawasm2go/p2.Fn1339
func Fn1339(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1381 github.com/goccy/llamawasm2go/p2.Fn1381
func Fn1381(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p2.Fn1394
func Fn1394(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1399 github.com/goccy/llamawasm2go/p2.Fn1399
func Fn1399(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p2.Fn1401
func Fn1401(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p2.Fn1402
func Fn1402(m *base.Module, l0 int64)

//go:linkname Fn1403 github.com/goccy/llamawasm2go/p2.Fn1403
func Fn1403(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p2.Fn1406
func Fn1406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p2.Fn1407
func Fn1407(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p2.Fn1408
func Fn1408(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1409 github.com/goccy/llamawasm2go/p2.Fn1409
func Fn1409(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1411 github.com/goccy/llamawasm2go/p2.Fn1411
func Fn1411(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p2.Fn1413
func Fn1413(m *base.Module, l0 int64, l1 int64) int64

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

//go:linkname Fn1419 github.com/goccy/llamawasm2go/p2.Fn1419
func Fn1419(m *base.Module, l0 int64) int64

//go:linkname Fn1421 github.com/goccy/llamawasm2go/p2.Fn1421
func Fn1421(m *base.Module, l0 int64) int64

//go:linkname Fn1422 github.com/goccy/llamawasm2go/p2.Fn1422
func Fn1422(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64) int64

//go:linkname Fn1425 github.com/goccy/llamawasm2go/p2.Fn1425
func Fn1425(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p2.Fn1426
func Fn1426(m *base.Module, l0 int64) int64

//go:linkname Fn1428 github.com/goccy/llamawasm2go/p2.Fn1428
func Fn1428(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1429 github.com/goccy/llamawasm2go/p2.Fn1429
func Fn1429(m *base.Module, l0 int64) int64

//go:linkname Fn1431 github.com/goccy/llamawasm2go/p2.Fn1431
func Fn1431(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1432 github.com/goccy/llamawasm2go/p2.Fn1432
func Fn1432(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1433 github.com/goccy/llamawasm2go/p2.Fn1433
func Fn1433(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1434 github.com/goccy/llamawasm2go/p2.Fn1434
func Fn1434(m *base.Module, l0 int64) int64

//go:linkname Fn1435 github.com/goccy/llamawasm2go/p2.Fn1435
func Fn1435(m *base.Module, l0 int64) int64

//go:linkname Fn1436 github.com/goccy/llamawasm2go/p2.Fn1436
func Fn1436(m *base.Module, l0 int64) int64

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int64) int64

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p2.Fn1447
func Fn1447(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1448 github.com/goccy/llamawasm2go/p2.Fn1448
func Fn1448(m *base.Module, l0 int64) int64

//go:linkname Fn1449 github.com/goccy/llamawasm2go/p2.Fn1449
func Fn1449(m *base.Module, l0 int64) int64

//go:linkname Fn1450 github.com/goccy/llamawasm2go/p2.Fn1450
func Fn1450(m *base.Module, l0 int64) int64

//go:linkname Fn1451 github.com/goccy/llamawasm2go/p2.Fn1451
func Fn1451(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1452 github.com/goccy/llamawasm2go/p2.Fn1452
func Fn1452(m *base.Module, l0 int64)

//go:linkname Fn1481 github.com/goccy/llamawasm2go/p2.Fn1481
func Fn1481(m *base.Module, l0 int64) int64

//go:linkname Fn1484 github.com/goccy/llamawasm2go/p2.Fn1484
func Fn1484(m *base.Module)

//go:linkname Fn1487 github.com/goccy/llamawasm2go/p2.Fn1487
func Fn1487(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1488 github.com/goccy/llamawasm2go/p2.Fn1488
func Fn1488(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1489 github.com/goccy/llamawasm2go/p2.Fn1489
func Fn1489(m *base.Module, l0 int64, l1 int32) int32

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

//go:linkname Fn1496 github.com/goccy/llamawasm2go/p2.Fn1496
func Fn1496(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1498 github.com/goccy/llamawasm2go/p2.Fn1498
func Fn1498(m *base.Module, l0 int64) int32

//go:linkname Fn1499 github.com/goccy/llamawasm2go/p2.Fn1499
func Fn1499(m *base.Module, l0 int64) int32

//go:linkname Fn1500 github.com/goccy/llamawasm2go/p2.Fn1500
func Fn1500(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1501 github.com/goccy/llamawasm2go/p2.Fn1501
func Fn1501(m *base.Module, l0 int64) int32

//go:linkname Fn1502 github.com/goccy/llamawasm2go/p2.Fn1502
func Fn1502(m *base.Module, l0 int64) int32

//go:linkname Fn1505 github.com/goccy/llamawasm2go/p2.Fn1505
func Fn1505(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1506 github.com/goccy/llamawasm2go/p2.Fn1506
func Fn1506(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1507 github.com/goccy/llamawasm2go/p2.Fn1507
func Fn1507(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1508 github.com/goccy/llamawasm2go/p2.Fn1508
func Fn1508(m *base.Module)

//go:linkname Fn1509 github.com/goccy/llamawasm2go/p2.Fn1509
func Fn1509(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1511 github.com/goccy/llamawasm2go/p2.Fn1511
func Fn1511(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1513 github.com/goccy/llamawasm2go/p2.Fn1513
func Fn1513(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1514 github.com/goccy/llamawasm2go/p2.Fn1514
func Fn1514(m *base.Module, l0 int64)

//go:linkname Fn1518 github.com/goccy/llamawasm2go/p2.Fn1518
func Fn1518(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1519 github.com/goccy/llamawasm2go/p2.Fn1519
func Fn1519(m *base.Module, l0 int64)

//go:linkname Fn1522 github.com/goccy/llamawasm2go/p2.Fn1522
func Fn1522(m *base.Module, l0 int64)

//go:linkname Fn1531 github.com/goccy/llamawasm2go/p2.Fn1531
func Fn1531(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1532 github.com/goccy/llamawasm2go/p2.Fn1532
func Fn1532(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p2.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p0.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1544 github.com/goccy/llamawasm2go/p2.Fn1544
func Fn1544(m *base.Module, l0 int64)

//go:linkname Fn1546 github.com/goccy/llamawasm2go/p2.Fn1546
func Fn1546(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1554 github.com/goccy/llamawasm2go/p2.Fn1554
func Fn1554(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1561 github.com/goccy/llamawasm2go/p2.Fn1561
func Fn1561(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1568 github.com/goccy/llamawasm2go/p2.Fn1568
func Fn1568(m *base.Module, l0 int64)

//go:linkname Fn1571 github.com/goccy/llamawasm2go/p2.Fn1571
func Fn1571(m *base.Module, l0 int64) int32

//go:linkname Fn1581 github.com/goccy/llamawasm2go/p2.Fn1581
func Fn1581(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p2.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1584 github.com/goccy/llamawasm2go/p2.Fn1584
func Fn1584(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1603 github.com/goccy/llamawasm2go/p2.Fn1603
func Fn1603(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1616 github.com/goccy/llamawasm2go/p2.Fn1616
func Fn1616(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1623 github.com/goccy/llamawasm2go/p2.Fn1623
func Fn1623(m *base.Module, l0 int64)

//go:linkname Fn1657 github.com/goccy/llamawasm2go/p2.Fn1657
func Fn1657(m *base.Module, l0 int64)

//go:linkname Fn1660 github.com/goccy/llamawasm2go/p2.Fn1660
func Fn1660(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1674 github.com/goccy/llamawasm2go/p2.Fn1674
func Fn1674(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1675 github.com/goccy/llamawasm2go/p2.Fn1675
func Fn1675(m *base.Module, l0 int64) int64

//go:linkname Fn1676 github.com/goccy/llamawasm2go/p2.Fn1676
func Fn1676(m *base.Module, l0 int64)

//go:linkname Fn1680 github.com/goccy/llamawasm2go/p0.Fn1680
func Fn1680(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1681 github.com/goccy/llamawasm2go/p2.Fn1681
func Fn1681(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1682 github.com/goccy/llamawasm2go/p2.Fn1682
func Fn1682(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1683 github.com/goccy/llamawasm2go/p2.Fn1683
func Fn1683(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1684 github.com/goccy/llamawasm2go/p2.Fn1684
func Fn1684(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1702 github.com/goccy/llamawasm2go/p2.Fn1702
func Fn1702(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1712 github.com/goccy/llamawasm2go/p2.Fn1712
func Fn1712(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1714 github.com/goccy/llamawasm2go/p2.Fn1714
func Fn1714(m *base.Module, l0 int64) int64

//go:linkname Fn1715 github.com/goccy/llamawasm2go/p2.Fn1715
func Fn1715(m *base.Module, l0 int64)

//go:linkname Fn1718 github.com/goccy/llamawasm2go/p0.Fn1718
func Fn1718(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1720 github.com/goccy/llamawasm2go/p2.Fn1720
func Fn1720(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1731 github.com/goccy/llamawasm2go/p2.Fn1731
func Fn1731(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1734 github.com/goccy/llamawasm2go/p2.Fn1734
func Fn1734(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1735 github.com/goccy/llamawasm2go/p2.Fn1735
func Fn1735(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1745 github.com/goccy/llamawasm2go/p2.Fn1745
func Fn1745(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1771 github.com/goccy/llamawasm2go/p2.Fn1771
func Fn1771(m *base.Module, l0 int64)

//go:linkname Fn1772 github.com/goccy/llamawasm2go/p2.Fn1772
func Fn1772(m *base.Module, l0 int64)

//go:linkname Fn1783 github.com/goccy/llamawasm2go/p2.Fn1783
func Fn1783(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1819 github.com/goccy/llamawasm2go/p2.Fn1819
func Fn1819(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1823 github.com/goccy/llamawasm2go/p2.Fn1823
func Fn1823(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1825 github.com/goccy/llamawasm2go/p2.Fn1825
func Fn1825(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1826 github.com/goccy/llamawasm2go/p2.Fn1826
func Fn1826(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1827 github.com/goccy/llamawasm2go/p2.Fn1827
func Fn1827(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1828 github.com/goccy/llamawasm2go/p2.Fn1828
func Fn1828(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1829 github.com/goccy/llamawasm2go/p2.Fn1829
func Fn1829(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1834 github.com/goccy/llamawasm2go/p2.Fn1834
func Fn1834(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn1840 github.com/goccy/llamawasm2go/p2.Fn1840
func Fn1840(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1842 github.com/goccy/llamawasm2go/p2.Fn1842
func Fn1842(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1845 github.com/goccy/llamawasm2go/p2.Fn1845
func Fn1845(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1846 github.com/goccy/llamawasm2go/p2.Fn1846
func Fn1846(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1847 github.com/goccy/llamawasm2go/p2.Fn1847
func Fn1847(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1850 github.com/goccy/llamawasm2go/p2.Fn1850
func Fn1850(m *base.Module, l0 int64)

//go:linkname Fn1860 github.com/goccy/llamawasm2go/p2.Fn1860
func Fn1860(m *base.Module, l0 int64)

//go:linkname Fn1862 github.com/goccy/llamawasm2go/p2.Fn1862
func Fn1862(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1863 github.com/goccy/llamawasm2go/p2.Fn1863
func Fn1863(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1876 github.com/goccy/llamawasm2go/p2.Fn1876
func Fn1876(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1877 github.com/goccy/llamawasm2go/p2.Fn1877
func Fn1877(m *base.Module, l0 int64) int64

//go:linkname Fn1878 github.com/goccy/llamawasm2go/p2.Fn1878
func Fn1878(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1880 github.com/goccy/llamawasm2go/p2.Fn1880
func Fn1880(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1892 github.com/goccy/llamawasm2go/p2.Fn1892
func Fn1892(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1894 github.com/goccy/llamawasm2go/p2.Fn1894
func Fn1894(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1895 github.com/goccy/llamawasm2go/p2.Fn1895
func Fn1895(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1897 github.com/goccy/llamawasm2go/p2.Fn1897
func Fn1897(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1898 github.com/goccy/llamawasm2go/p2.Fn1898
func Fn1898(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1899 github.com/goccy/llamawasm2go/p2.Fn1899
func Fn1899(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1900 github.com/goccy/llamawasm2go/p2.Fn1900
func Fn1900(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1901 github.com/goccy/llamawasm2go/p2.Fn1901
func Fn1901(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn1902 github.com/goccy/llamawasm2go/p2.Fn1902
func Fn1902(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1921 github.com/goccy/llamawasm2go/p2.Fn1921
func Fn1921(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1922 github.com/goccy/llamawasm2go/p2.Fn1922
func Fn1922(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1923 github.com/goccy/llamawasm2go/p2.Fn1923
func Fn1923(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1925 github.com/goccy/llamawasm2go/p2.Fn1925
func Fn1925(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1926 github.com/goccy/llamawasm2go/p2.Fn1926
func Fn1926(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1927 github.com/goccy/llamawasm2go/p2.Fn1927
func Fn1927(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1928 github.com/goccy/llamawasm2go/p2.Fn1928
func Fn1928(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1929 github.com/goccy/llamawasm2go/p2.Fn1929
func Fn1929(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1931 github.com/goccy/llamawasm2go/p2.Fn1931
func Fn1931(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1933 github.com/goccy/llamawasm2go/p2.Fn1933
func Fn1933(m *base.Module, l0 int64)

//go:linkname Fn1950 github.com/goccy/llamawasm2go/p2.Fn1950
func Fn1950(m *base.Module, l0 int64)

//go:linkname Fn1951 github.com/goccy/llamawasm2go/p2.Fn1951
func Fn1951(m *base.Module, l0 int64)

//go:linkname Fn1952 github.com/goccy/llamawasm2go/p2.Fn1952
func Fn1952(m *base.Module, l0 int64)

//go:linkname Fn1954 github.com/goccy/llamawasm2go/p2.Fn1954
func Fn1954(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1955 github.com/goccy/llamawasm2go/p2.Fn1955
func Fn1955(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1994 github.com/goccy/llamawasm2go/p2.Fn1994
func Fn1994(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2006 github.com/goccy/llamawasm2go/p2.Fn2006
func Fn2006(m *base.Module, l0 int64) int64

//go:linkname Fn2009 github.com/goccy/llamawasm2go/p2.Fn2009
func Fn2009(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2013 github.com/goccy/llamawasm2go/p2.Fn2013
func Fn2013(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2017 github.com/goccy/llamawasm2go/p2.Fn2017
func Fn2017(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2030 github.com/goccy/llamawasm2go/p2.Fn2030
func Fn2030(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2031 github.com/goccy/llamawasm2go/p2.Fn2031
func Fn2031(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2032 github.com/goccy/llamawasm2go/p2.Fn2032
func Fn2032(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2044 github.com/goccy/llamawasm2go/p2.Fn2044
func Fn2044(m *base.Module, l0 int64) int64

//go:linkname Fn2045 github.com/goccy/llamawasm2go/p2.Fn2045
func Fn2045(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2048 github.com/goccy/llamawasm2go/p2.Fn2048
func Fn2048(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2050 github.com/goccy/llamawasm2go/p2.Fn2050
func Fn2050(m *base.Module)

//go:linkname Fn2053 github.com/goccy/llamawasm2go/p2.Fn2053
func Fn2053(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2062 github.com/goccy/llamawasm2go/p2.Fn2062
func Fn2062(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2066 github.com/goccy/llamawasm2go/p2.Fn2066
func Fn2066(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2067 github.com/goccy/llamawasm2go/p2.Fn2067
func Fn2067(m *base.Module, l0 int64)

//go:linkname Fn2068 github.com/goccy/llamawasm2go/p2.Fn2068
func Fn2068(m *base.Module, l0 int64)

//go:linkname Fn2069 github.com/goccy/llamawasm2go/p2.Fn2069
func Fn2069(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2071 github.com/goccy/llamawasm2go/p2.Fn2071
func Fn2071(m *base.Module, l0 int64)

//go:linkname Fn2072 github.com/goccy/llamawasm2go/p2.Fn2072
func Fn2072(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2073 github.com/goccy/llamawasm2go/p2.Fn2073
func Fn2073(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2077 github.com/goccy/llamawasm2go/p2.Fn2077
func Fn2077(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2080 github.com/goccy/llamawasm2go/p2.Fn2080
func Fn2080(m *base.Module, l0 int64) int64

//go:linkname Fn2081 github.com/goccy/llamawasm2go/p2.Fn2081
func Fn2081(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2082 github.com/goccy/llamawasm2go/p2.Fn2082
func Fn2082(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2089 github.com/goccy/llamawasm2go/p2.Fn2089
func Fn2089(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2090 github.com/goccy/llamawasm2go/p2.Fn2090
func Fn2090(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2092 github.com/goccy/llamawasm2go/p2.Fn2092
func Fn2092(m *base.Module) int64

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2095 github.com/goccy/llamawasm2go/p2.Fn2095
func Fn2095(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2096 github.com/goccy/llamawasm2go/p2.Fn2096
func Fn2096(m *base.Module) int64

//go:linkname Fn2098 github.com/goccy/llamawasm2go/p2.Fn2098
func Fn2098(m *base.Module, l0 int32) int64

//go:linkname Fn2099 github.com/goccy/llamawasm2go/p2.Fn2099
func Fn2099(m *base.Module, l0 int32) int32

//go:linkname Fn2100 github.com/goccy/llamawasm2go/p2.Fn2100
func Fn2100(m *base.Module, l0 int32) int64

//go:linkname Fn2101 github.com/goccy/llamawasm2go/p2.Fn2101
func Fn2101(m *base.Module, l0 float32) int64

//go:linkname Fn2102 github.com/goccy/llamawasm2go/p2.Fn2102
func Fn2102(m *base.Module, l0 float32) int64

//go:linkname Fn2103 github.com/goccy/llamawasm2go/p2.Fn2103
func Fn2103(m *base.Module, l0 float32) int64

//go:linkname Fn2105 github.com/goccy/llamawasm2go/p2.Fn2105
func Fn2105(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2106 github.com/goccy/llamawasm2go/p2.Fn2106
func Fn2106(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2147 github.com/goccy/llamawasm2go/p2.Fn2147
func Fn2147(m *base.Module, l0 int64)

//go:linkname Fn2149 github.com/goccy/llamawasm2go/p2.Fn2149
func Fn2149(m *base.Module, l0 int64)

//go:linkname Fn2157 github.com/goccy/llamawasm2go/p2.Fn2157
func Fn2157(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2191 github.com/goccy/llamawasm2go/p2.Fn2191
func Fn2191(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2202 github.com/goccy/llamawasm2go/p2.Fn2202
func Fn2202(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2205 github.com/goccy/llamawasm2go/p2.Fn2205
func Fn2205(m *base.Module, l0 int64)

//go:linkname Fn2206 github.com/goccy/llamawasm2go/p2.Fn2206
func Fn2206(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2207 github.com/goccy/llamawasm2go/p2.Fn2207
func Fn2207(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2208 github.com/goccy/llamawasm2go/p2.Fn2208
func Fn2208(m *base.Module, l0 int64)

//go:linkname Fn2211 github.com/goccy/llamawasm2go/p2.Fn2211
func Fn2211(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2216 github.com/goccy/llamawasm2go/p2.Fn2216
func Fn2216(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2217 github.com/goccy/llamawasm2go/p2.Fn2217
func Fn2217(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2230 github.com/goccy/llamawasm2go/p2.Fn2230
func Fn2230(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2231 github.com/goccy/llamawasm2go/p2.Fn2231
func Fn2231(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2241 github.com/goccy/llamawasm2go/p2.Fn2241
func Fn2241(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2244 github.com/goccy/llamawasm2go/p2.Fn2244
func Fn2244(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2246 github.com/goccy/llamawasm2go/p2.Fn2246
func Fn2246(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2247 github.com/goccy/llamawasm2go/p2.Fn2247
func Fn2247(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2248 github.com/goccy/llamawasm2go/p2.Fn2248
func Fn2248(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2251 github.com/goccy/llamawasm2go/p2.Fn2251
func Fn2251(m *base.Module)

//go:linkname Fn2255 github.com/goccy/llamawasm2go/p2.Fn2255
func Fn2255(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2257 github.com/goccy/llamawasm2go/p0.Fn2257
func Fn2257(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2272 github.com/goccy/llamawasm2go/p2.Fn2272
func Fn2272(m *base.Module, l0 int64)

//go:linkname Fn2273 github.com/goccy/llamawasm2go/p2.Fn2273
func Fn2273(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2278 github.com/goccy/llamawasm2go/p2.Fn2278
func Fn2278(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2283 github.com/goccy/llamawasm2go/p2.Fn2283
func Fn2283(m *base.Module, l0 int64) int64

//go:linkname Fn2284 github.com/goccy/llamawasm2go/p2.Fn2284
func Fn2284(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2285 github.com/goccy/llamawasm2go/p0.Fn2285
func Fn2285(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2296 github.com/goccy/llamawasm2go/p2.Fn2296
func Fn2296(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2297 github.com/goccy/llamawasm2go/p2.Fn2297
func Fn2297(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2298 github.com/goccy/llamawasm2go/p2.Fn2298
func Fn2298(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2299 github.com/goccy/llamawasm2go/p2.Fn2299
func Fn2299(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2319 github.com/goccy/llamawasm2go/p2.Fn2319
func Fn2319(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2320 github.com/goccy/llamawasm2go/p2.Fn2320
func Fn2320(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2323 github.com/goccy/llamawasm2go/p2.Fn2323
func Fn2323(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2409 github.com/goccy/llamawasm2go/p2.Fn2409
func Fn2409(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2528 github.com/goccy/llamawasm2go/p2.Fn2528
func Fn2528(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2529 github.com/goccy/llamawasm2go/p0.Fn2529
func Fn2529(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2530 github.com/goccy/llamawasm2go/p2.Fn2530
func Fn2530(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2531 github.com/goccy/llamawasm2go/p2.Fn2531
func Fn2531(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2535 github.com/goccy/llamawasm2go/p2.Fn2535
func Fn2535(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2560 github.com/goccy/llamawasm2go/p2.Fn2560
func Fn2560(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2692 github.com/goccy/llamawasm2go/p2.Fn2692
func Fn2692(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2727 github.com/goccy/llamawasm2go/p2.Fn2727
func Fn2727(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2741 github.com/goccy/llamawasm2go/p2.Fn2741
func Fn2741(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2751 github.com/goccy/llamawasm2go/p2.Fn2751
func Fn2751(m *base.Module, l0 int32)

//go:linkname Fn2753 github.com/goccy/llamawasm2go/p2.Fn2753
func Fn2753(m *base.Module, l0 int64) int64

//go:linkname Fn2754 github.com/goccy/llamawasm2go/p2.Fn2754
func Fn2754(m *base.Module, l0 int64)

//go:linkname Fn2757 github.com/goccy/llamawasm2go/p2.Fn2757
func Fn2757(m *base.Module, l0 int64)

//go:linkname Fn2758 github.com/goccy/llamawasm2go/p2.Fn2758
func Fn2758(m *base.Module, l0 int64)

//go:linkname Fn2760 github.com/goccy/llamawasm2go/p2.Fn2760
func Fn2760(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2765 github.com/goccy/llamawasm2go/p2.Fn2765
func Fn2765(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2771 github.com/goccy/llamawasm2go/p2.Fn2771
func Fn2771(m *base.Module) int32

//go:linkname Fn2779 github.com/goccy/llamawasm2go/p2.Fn2779
func Fn2779(m *base.Module, l0 float64) float32

//go:linkname Fn2780 github.com/goccy/llamawasm2go/p2.Fn2780
func Fn2780(m *base.Module, l0 float64) float32

//go:linkname Fn2787 github.com/goccy/llamawasm2go/p2.Fn2787
func Fn2787(m *base.Module, l0 float32) float32

//go:linkname Fn2791 github.com/goccy/llamawasm2go/p2.Fn2791
func Fn2791(m *base.Module, l0 float32) float32

//go:linkname Fn2794 github.com/goccy/llamawasm2go/p2.Fn2794
func Fn2794(m *base.Module, l0 float32) float32

//go:linkname Fn2809 github.com/goccy/llamawasm2go/p2.Fn2809
func Fn2809(m *base.Module, l0 int64) int32

//go:linkname Fn2810 github.com/goccy/llamawasm2go/p2.Fn2810
func Fn2810(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2812 github.com/goccy/llamawasm2go/p2.Fn2812
func Fn2812(m *base.Module, l0 int64)

//go:linkname Fn2813 github.com/goccy/llamawasm2go/p2.Fn2813
func Fn2813(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2814 github.com/goccy/llamawasm2go/p2.Fn2814
func Fn2814(m *base.Module, l0 int64) int32

//go:linkname Fn2821 github.com/goccy/llamawasm2go/p2.Fn2821
func Fn2821(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2823 github.com/goccy/llamawasm2go/p2.Fn2823
func Fn2823(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2829 github.com/goccy/llamawasm2go/p2.Fn2829
func Fn2829(m *base.Module, l0 int64) int32

//go:linkname Fn2830 github.com/goccy/llamawasm2go/p2.Fn2830
func Fn2830(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2835 github.com/goccy/llamawasm2go/p2.Fn2835
func Fn2835(m *base.Module, l0 int64) int32

//go:linkname Fn2839 github.com/goccy/llamawasm2go/p2.Fn2839
func Fn2839(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn2840 github.com/goccy/llamawasm2go/p2.Fn2840
func Fn2840(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2843 github.com/goccy/llamawasm2go/p2.Fn2843
func Fn2843(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2846 github.com/goccy/llamawasm2go/p2.Fn2846
func Fn2846(m *base.Module, l0 int64) int64

//go:linkname Fn2850 github.com/goccy/llamawasm2go/p2.Fn2850
func Fn2850(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn2860 github.com/goccy/llamawasm2go/p2.Fn2860
func Fn2860(m *base.Module)

//go:linkname Fn2861 github.com/goccy/llamawasm2go/p0.Fn2861
func Fn2861(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn2863 github.com/goccy/llamawasm2go/p2.Fn2863
func Fn2863(m *base.Module)

//go:linkname Fn2865 github.com/goccy/llamawasm2go/p0.Fn2865
func Fn2865(m *base.Module, l0 int64) int64

//go:linkname Fn2867 github.com/goccy/llamawasm2go/p2.Fn2867
func Fn2867(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2868 github.com/goccy/llamawasm2go/p2.Fn2868
func Fn2868(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2871 github.com/goccy/llamawasm2go/p2.Fn2871
func Fn2871(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2900 github.com/goccy/llamawasm2go/p2.Fn2900
func Fn2900(m *base.Module, l0 int32)
