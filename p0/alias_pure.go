//go:build !arm64 && (!amd64 || !amd64.v2)

package p0

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn21 github.com/goccy/llamawasm2go/p2.Fn21
func Fn21(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn46 github.com/goccy/llamawasm2go/p2.Fn46
func Fn46(m *base.Module, l0 int64)

//go:linkname Fn52 github.com/goccy/llamawasm2go/p3.Fn52
func Fn52(m *base.Module, l0 int64) int64

//go:linkname Fn53 github.com/goccy/llamawasm2go/p3.Fn53
func Fn53(m *base.Module, l0 int64)

//go:linkname Fn55 github.com/goccy/llamawasm2go/p3.Fn55
func Fn55(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn56 github.com/goccy/llamawasm2go/p2.Fn56
func Fn56(m *base.Module, l0 int64) int64

//go:linkname Fn57 github.com/goccy/llamawasm2go/p2.Fn57
func Fn57(m *base.Module)

//go:linkname Fn64 github.com/goccy/llamawasm2go/p3.Fn64
func Fn64(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn69 github.com/goccy/llamawasm2go/p2.Fn69
func Fn69(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn74 github.com/goccy/llamawasm2go/p3.Fn74
func Fn74(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn75 github.com/goccy/llamawasm2go/p2.Fn75
func Fn75(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn76 github.com/goccy/llamawasm2go/p2.Fn76
func Fn76(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn78 github.com/goccy/llamawasm2go/p2.Fn78
func Fn78(m *base.Module, l0 int64)

//go:linkname Fn81 github.com/goccy/llamawasm2go/p3.Fn81
func Fn81(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn82 github.com/goccy/llamawasm2go/p3.Fn82
func Fn82(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn84 github.com/goccy/llamawasm2go/p3.Fn84
func Fn84(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn85 github.com/goccy/llamawasm2go/p2.Fn85
func Fn85(m *base.Module, l0 int64) int64

//go:linkname Fn88 github.com/goccy/llamawasm2go/p3.Fn88
func Fn88(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn89 github.com/goccy/llamawasm2go/p3.Fn89
func Fn89(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn91 github.com/goccy/llamawasm2go/p2.Fn91
func Fn91(m *base.Module, l0 int64) int64

//go:linkname Fn94 github.com/goccy/llamawasm2go/p3.Fn94
func Fn94(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn97 github.com/goccy/llamawasm2go/p2.Fn97
func Fn97(m *base.Module, l0 int64) int64

//go:linkname Fn107 github.com/goccy/llamawasm2go/p2.Fn107
func Fn107(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn109 github.com/goccy/llamawasm2go/p3.Fn109
func Fn109(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn117 github.com/goccy/llamawasm2go/p2.Fn117
func Fn117(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn118 github.com/goccy/llamawasm2go/p2.Fn118
func Fn118(m *base.Module, l0 int64) int64

//go:linkname Fn120 github.com/goccy/llamawasm2go/p3.Fn120
func Fn120(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn121 github.com/goccy/llamawasm2go/p2.Fn121
func Fn121(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn122 github.com/goccy/llamawasm2go/p3.Fn122
func Fn122(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn140 github.com/goccy/llamawasm2go/p3.Fn140
func Fn140(m *base.Module, l0 int64)

//go:linkname Fn147 github.com/goccy/llamawasm2go/p2.Fn147
func Fn147(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn240 github.com/goccy/llamawasm2go/p2.Fn240
func Fn240(m *base.Module)

//go:linkname Fn242 github.com/goccy/llamawasm2go/p3.Fn242
func Fn242(m *base.Module, l0 int64) int64

//go:linkname Fn243 github.com/goccy/llamawasm2go/p3.Fn243
func Fn243(m *base.Module, l0 int64) int64

//go:linkname Fn250 github.com/goccy/llamawasm2go/p3.Fn250
func Fn250(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn251 github.com/goccy/llamawasm2go/p3.Fn251
func Fn251(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn253 github.com/goccy/llamawasm2go/p3.Fn253
func Fn253(m *base.Module)

//go:linkname Fn254 github.com/goccy/llamawasm2go/p3.Fn254
func Fn254(m *base.Module)

//go:linkname Fn255 github.com/goccy/llamawasm2go/p3.Fn255
func Fn255(m *base.Module, l0 int64)

//go:linkname Fn260 github.com/goccy/llamawasm2go/p3.Fn260
func Fn260(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn261 github.com/goccy/llamawasm2go/p2.Fn261
func Fn261(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn263 github.com/goccy/llamawasm2go/p2.Fn263
func Fn263(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn264 github.com/goccy/llamawasm2go/p2.Fn264
func Fn264(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn265 github.com/goccy/llamawasm2go/p3.Fn265
func Fn265(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn267 github.com/goccy/llamawasm2go/p3.Fn267
func Fn267(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn268 github.com/goccy/llamawasm2go/p3.Fn268
func Fn268(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn269 github.com/goccy/llamawasm2go/p2.Fn269
func Fn269(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn270 github.com/goccy/llamawasm2go/p2.Fn270
func Fn270(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn272 github.com/goccy/llamawasm2go/p3.Fn272
func Fn272(m *base.Module)

//go:linkname Fn277 github.com/goccy/llamawasm2go/p2.Fn277
func Fn277(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn278 github.com/goccy/llamawasm2go/p2.Fn278
func Fn278(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn279 github.com/goccy/llamawasm2go/p2.Fn279
func Fn279(m *base.Module, l0 int64) int64

//go:linkname Fn281 github.com/goccy/llamawasm2go/p2.Fn281
func Fn281(m *base.Module, l0 int64) int64

//go:linkname Fn282 github.com/goccy/llamawasm2go/p2.Fn282
func Fn282(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn284 github.com/goccy/llamawasm2go/p2.Fn284
func Fn284(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn285 github.com/goccy/llamawasm2go/p3.Fn285
func Fn285(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn341 github.com/goccy/llamawasm2go/p2.Fn341
func Fn341(m *base.Module, l0 int64) int32

//go:linkname Fn342 github.com/goccy/llamawasm2go/p3.Fn342
func Fn342(m *base.Module, l0 int64)

//go:linkname Fn343 github.com/goccy/llamawasm2go/p3.Fn343
func Fn343(m *base.Module, l0 int64)

//go:linkname Fn350 github.com/goccy/llamawasm2go/p3.Fn350
func Fn350(m *base.Module) int64

//go:linkname Fn353 github.com/goccy/llamawasm2go/p3.Fn353
func Fn353(m *base.Module) int64

//go:linkname Fn354 github.com/goccy/llamawasm2go/p3.Fn354
func Fn354(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn360 github.com/goccy/llamawasm2go/p3.Fn360
func Fn360(m *base.Module) int64

//go:linkname Fn363 github.com/goccy/llamawasm2go/p2.Fn363
func Fn363(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn364 github.com/goccy/llamawasm2go/p3.Fn364
func Fn364(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn367 github.com/goccy/llamawasm2go/p2.Fn367
func Fn367(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn368 github.com/goccy/llamawasm2go/p2.Fn368
func Fn368(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn373 github.com/goccy/llamawasm2go/p2.Fn373
func Fn373(m *base.Module, l0 int64)

//go:linkname Fn374 github.com/goccy/llamawasm2go/p2.Fn374
func Fn374(m *base.Module, l0 int64)

//go:linkname Fn383 github.com/goccy/llamawasm2go/p2.Fn383
func Fn383(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn384 github.com/goccy/llamawasm2go/p3.Fn384
func Fn384(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn386 github.com/goccy/llamawasm2go/p2.Fn386
func Fn386(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn387 github.com/goccy/llamawasm2go/p2.Fn387
func Fn387(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn391 github.com/goccy/llamawasm2go/p1.Fn391
func Fn391(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn392 github.com/goccy/llamawasm2go/p1.Fn392
func Fn392(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn393 github.com/goccy/llamawasm2go/p2.Fn393
func Fn393(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn394 github.com/goccy/llamawasm2go/p2.Fn394
func Fn394(m *base.Module, l0 int64)

//go:linkname Fn397 github.com/goccy/llamawasm2go/p3.Fn397
func Fn397(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn398 github.com/goccy/llamawasm2go/p2.Fn398
func Fn398(m *base.Module, l0 int64)

//go:linkname Fn400 github.com/goccy/llamawasm2go/p2.Fn400
func Fn400(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn401 github.com/goccy/llamawasm2go/p3.Fn401
func Fn401(m *base.Module)

//go:linkname Fn404 github.com/goccy/llamawasm2go/p2.Fn404
func Fn404(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn407 github.com/goccy/llamawasm2go/p2.Fn407
func Fn407(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn408 github.com/goccy/llamawasm2go/p1.Fn408
func Fn408(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn414 github.com/goccy/llamawasm2go/p2.Fn414
func Fn414(m *base.Module, l0 int64)

//go:linkname Fn417 github.com/goccy/llamawasm2go/p2.Fn417
func Fn417(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn418 github.com/goccy/llamawasm2go/p3.Fn418
func Fn418(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn421 github.com/goccy/llamawasm2go/p2.Fn421
func Fn421(m *base.Module, l0 int64)

//go:linkname Fn438 github.com/goccy/llamawasm2go/p3.Fn438
func Fn438(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn439 github.com/goccy/llamawasm2go/p3.Fn439
func Fn439(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn445 github.com/goccy/llamawasm2go/p2.Fn445
func Fn445(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn451 github.com/goccy/llamawasm2go/p3.Fn451
func Fn451(m *base.Module, l0 int32) int64

//go:linkname Fn452 github.com/goccy/llamawasm2go/p3.Fn452
func Fn452(m *base.Module, l0 int64) int64

//go:linkname Fn453 github.com/goccy/llamawasm2go/p3.Fn453
func Fn453(m *base.Module, l0 int64) int64

//go:linkname Fn454 github.com/goccy/llamawasm2go/p3.Fn454
func Fn454(m *base.Module, l0 int64) int64

//go:linkname Fn456 github.com/goccy/llamawasm2go/p3.Fn456
func Fn456(m *base.Module, l0 int32) int64

//go:linkname Fn458 github.com/goccy/llamawasm2go/p3.Fn458
func Fn458(m *base.Module, l0 int32) int64

//go:linkname Fn461 github.com/goccy/llamawasm2go/p3.Fn461
func Fn461(m *base.Module, l0 int64) int32

//go:linkname Fn462 github.com/goccy/llamawasm2go/p3.Fn462
func Fn462(m *base.Module, l0 int64) int32

//go:linkname Fn463 github.com/goccy/llamawasm2go/p3.Fn463
func Fn463(m *base.Module, l0 int64) int64

//go:linkname Fn464 github.com/goccy/llamawasm2go/p3.Fn464
func Fn464(m *base.Module, l0 int64) int32

//go:linkname Fn466 github.com/goccy/llamawasm2go/p3.Fn466
func Fn466(m *base.Module, l0 int64) int32

//go:linkname Fn467 github.com/goccy/llamawasm2go/p3.Fn467
func Fn467(m *base.Module, l0 int64) int32

//go:linkname Fn469 github.com/goccy/llamawasm2go/p3.Fn469
func Fn469(m *base.Module, l0 int64) int32

//go:linkname Fn471 github.com/goccy/llamawasm2go/p3.Fn471
func Fn471(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn474 github.com/goccy/llamawasm2go/p2.Fn474
func Fn474(m *base.Module, l0 int64) int64

//go:linkname Fn480 github.com/goccy/llamawasm2go/p3.Fn480
func Fn480(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn481 github.com/goccy/llamawasm2go/p2.Fn481
func Fn481(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn483 github.com/goccy/llamawasm2go/p3.Fn483
func Fn483(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn484 github.com/goccy/llamawasm2go/p3.Fn484
func Fn484(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn485 github.com/goccy/llamawasm2go/p3.Fn485
func Fn485(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn486 github.com/goccy/llamawasm2go/p3.Fn486
func Fn486(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn487 github.com/goccy/llamawasm2go/p3.Fn487
func Fn487(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn490 github.com/goccy/llamawasm2go/p3.Fn490
func Fn490(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn491 github.com/goccy/llamawasm2go/p3.Fn491
func Fn491(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn492 github.com/goccy/llamawasm2go/p3.Fn492
func Fn492(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn496 github.com/goccy/llamawasm2go/p3.Fn496
func Fn496(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn497 github.com/goccy/llamawasm2go/p3.Fn497
func Fn497(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn501 github.com/goccy/llamawasm2go/p3.Fn501
func Fn501(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn504 github.com/goccy/llamawasm2go/p3.Fn504
func Fn504(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn507 github.com/goccy/llamawasm2go/p2.Fn507
func Fn507(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn508 github.com/goccy/llamawasm2go/p3.Fn508
func Fn508(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn509 github.com/goccy/llamawasm2go/p2.Fn509
func Fn509(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn512 github.com/goccy/llamawasm2go/p3.Fn512
func Fn512(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn513 github.com/goccy/llamawasm2go/p3.Fn513
func Fn513(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn514 github.com/goccy/llamawasm2go/p3.Fn514
func Fn514(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn518 github.com/goccy/llamawasm2go/p3.Fn518
func Fn518(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn526 github.com/goccy/llamawasm2go/p3.Fn526
func Fn526(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn527 github.com/goccy/llamawasm2go/p3.Fn527
func Fn527(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn528 github.com/goccy/llamawasm2go/p3.Fn528
func Fn528(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn531 github.com/goccy/llamawasm2go/p2.Fn531
func Fn531(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn532 github.com/goccy/llamawasm2go/p3.Fn532
func Fn532(m *base.Module, l0 int64)

//go:linkname Fn534 github.com/goccy/llamawasm2go/p3.Fn534
func Fn534(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn535 github.com/goccy/llamawasm2go/p3.Fn535
func Fn535(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn538 github.com/goccy/llamawasm2go/p3.Fn538
func Fn538(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn539 github.com/goccy/llamawasm2go/p3.Fn539
func Fn539(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn540 github.com/goccy/llamawasm2go/p3.Fn540
func Fn540(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn542 github.com/goccy/llamawasm2go/p3.Fn542
func Fn542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn544 github.com/goccy/llamawasm2go/p2.Fn544
func Fn544(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn545 github.com/goccy/llamawasm2go/p2.Fn545
func Fn545(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn546 github.com/goccy/llamawasm2go/p2.Fn546
func Fn546(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn547 github.com/goccy/llamawasm2go/p2.Fn547
func Fn547(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn549 github.com/goccy/llamawasm2go/p3.Fn549
func Fn549(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn550 github.com/goccy/llamawasm2go/p3.Fn550
func Fn550(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn551 github.com/goccy/llamawasm2go/p3.Fn551
func Fn551(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64) int64

//go:linkname Fn552 github.com/goccy/llamawasm2go/p2.Fn552
func Fn552(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn553 github.com/goccy/llamawasm2go/p3.Fn553
func Fn553(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn554 github.com/goccy/llamawasm2go/p2.Fn554
func Fn554(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn555 github.com/goccy/llamawasm2go/p2.Fn555
func Fn555(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn557 github.com/goccy/llamawasm2go/p3.Fn557
func Fn557(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn559 github.com/goccy/llamawasm2go/p3.Fn559
func Fn559(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn561 github.com/goccy/llamawasm2go/p2.Fn561
func Fn561(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32, l14 int32) int64

//go:linkname Fn562 github.com/goccy/llamawasm2go/p3.Fn562
func Fn562(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32) int64

//go:linkname Fn563 github.com/goccy/llamawasm2go/p3.Fn563
func Fn563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32, l7 float32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32) int64

//go:linkname Fn565 github.com/goccy/llamawasm2go/p3.Fn565
func Fn565(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn569 github.com/goccy/llamawasm2go/p3.Fn569
func Fn569(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32) int64

//go:linkname Fn570 github.com/goccy/llamawasm2go/p3.Fn570
func Fn570(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32) int64

//go:linkname Fn572 github.com/goccy/llamawasm2go/p2.Fn572
func Fn572(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn573 github.com/goccy/llamawasm2go/p3.Fn573
func Fn573(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn575 github.com/goccy/llamawasm2go/p3.Fn575
func Fn575(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn576 github.com/goccy/llamawasm2go/p3.Fn576
func Fn576(m *base.Module, l0 int64, l1 float32, l2 float32) int64

//go:linkname Fn582 github.com/goccy/llamawasm2go/p2.Fn582
func Fn582(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn590 github.com/goccy/llamawasm2go/p2.Fn590
func Fn590(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn592 github.com/goccy/llamawasm2go/p3.Fn592
func Fn592(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn593 github.com/goccy/llamawasm2go/p3.Fn593
func Fn593(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn595 github.com/goccy/llamawasm2go/p2.Fn595
func Fn595(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn597 github.com/goccy/llamawasm2go/p3.Fn597
func Fn597(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn598 github.com/goccy/llamawasm2go/p3.Fn598
func Fn598(m *base.Module, l0 int64)

//go:linkname Fn603 github.com/goccy/llamawasm2go/p1.Fn603
func Fn603(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn606 github.com/goccy/llamawasm2go/p3.Fn606
func Fn606(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn608 github.com/goccy/llamawasm2go/p3.Fn608
func Fn608(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn612 github.com/goccy/llamawasm2go/p2.Fn612
func Fn612(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn613 github.com/goccy/llamawasm2go/p1.Fn613
func Fn613(m *base.Module, l0 int64) int64

//go:linkname Fn615 github.com/goccy/llamawasm2go/p3.Fn615
func Fn615(m *base.Module, l0 int64) int64

//go:linkname Fn616 github.com/goccy/llamawasm2go/p3.Fn616
func Fn616(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn621 github.com/goccy/llamawasm2go/p3.Fn621
func Fn621(m *base.Module, l0 int64) int64

//go:linkname Fn625 github.com/goccy/llamawasm2go/p3.Fn625
func Fn625(m *base.Module, l0 int64) int64

//go:linkname Fn626 github.com/goccy/llamawasm2go/p3.Fn626
func Fn626(m *base.Module, l0 int64) int64

//go:linkname Fn627 github.com/goccy/llamawasm2go/p3.Fn627
func Fn627(m *base.Module, l0 int64) int64

//go:linkname Fn628 github.com/goccy/llamawasm2go/p3.Fn628
func Fn628(m *base.Module, l0 int64)

//go:linkname Fn629 github.com/goccy/llamawasm2go/p3.Fn629
func Fn629(m *base.Module, l0 int64) int64

//go:linkname Fn630 github.com/goccy/llamawasm2go/p3.Fn630
func Fn630(m *base.Module, l0 int64) int64

//go:linkname Fn631 github.com/goccy/llamawasm2go/p3.Fn631
func Fn631(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn632 github.com/goccy/llamawasm2go/p3.Fn632
func Fn632(m *base.Module, l0 int64) int32

//go:linkname Fn633 github.com/goccy/llamawasm2go/p2.Fn633
func Fn633(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn635 github.com/goccy/llamawasm2go/p3.Fn635
func Fn635(m *base.Module, l0 int64) int32

//go:linkname Fn637 github.com/goccy/llamawasm2go/p3.Fn637
func Fn637(m *base.Module, l0 int64)

//go:linkname Fn638 github.com/goccy/llamawasm2go/p3.Fn638
func Fn638(m *base.Module, l0 int64) int64

//go:linkname Fn639 github.com/goccy/llamawasm2go/p2.Fn639
func Fn639(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn643 github.com/goccy/llamawasm2go/p2.Fn643
func Fn643(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn648 github.com/goccy/llamawasm2go/p3.Fn648
func Fn648(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn651 github.com/goccy/llamawasm2go/p3.Fn651
func Fn651(m *base.Module, l0 int64) int64

//go:linkname Fn653 github.com/goccy/llamawasm2go/p2.Fn653
func Fn653(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn655 github.com/goccy/llamawasm2go/p3.Fn655
func Fn655(m *base.Module, l0 int64) int64

//go:linkname Fn657 github.com/goccy/llamawasm2go/p3.Fn657
func Fn657(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn658 github.com/goccy/llamawasm2go/p3.Fn658
func Fn658(m *base.Module, l0 int64) int32

//go:linkname Fn659 github.com/goccy/llamawasm2go/p3.Fn659
func Fn659(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn660 github.com/goccy/llamawasm2go/p3.Fn660
func Fn660(m *base.Module, l0 int64) int64

//go:linkname Fn661 github.com/goccy/llamawasm2go/p3.Fn661
func Fn661(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn662 github.com/goccy/llamawasm2go/p3.Fn662
func Fn662(m *base.Module, l0 int64) int64

//go:linkname Fn663 github.com/goccy/llamawasm2go/p3.Fn663
func Fn663(m *base.Module, l0 int64) int64

//go:linkname Fn664 github.com/goccy/llamawasm2go/p3.Fn664
func Fn664(m *base.Module, l0 int64) int64

//go:linkname Fn665 github.com/goccy/llamawasm2go/p3.Fn665
func Fn665(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn666 github.com/goccy/llamawasm2go/p3.Fn666
func Fn666(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn668 github.com/goccy/llamawasm2go/p2.Fn668
func Fn668(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn669 github.com/goccy/llamawasm2go/p2.Fn669
func Fn669(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn674 github.com/goccy/llamawasm2go/p3.Fn674
func Fn674(m *base.Module, l0 int64)

//go:linkname Fn675 github.com/goccy/llamawasm2go/p2.Fn675
func Fn675(m *base.Module, l0 int64)

//go:linkname Fn684 github.com/goccy/llamawasm2go/p2.Fn684
func Fn684(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn691 github.com/goccy/llamawasm2go/p2.Fn691
func Fn691(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn692 github.com/goccy/llamawasm2go/p2.Fn692
func Fn692(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn728 github.com/goccy/llamawasm2go/p3.Fn728
func Fn728(m *base.Module, l0 int64)

//go:linkname Fn735 github.com/goccy/llamawasm2go/p3.Fn735
func Fn735(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn742 github.com/goccy/llamawasm2go/p3.Fn742
func Fn742(m *base.Module, l0 int64)

//go:linkname Fn743 github.com/goccy/llamawasm2go/p3.Fn743
func Fn743(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn746 github.com/goccy/llamawasm2go/p1.Fn746
func Fn746(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn747 github.com/goccy/llamawasm2go/p3.Fn747
func Fn747(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn748 github.com/goccy/llamawasm2go/p2.Fn748
func Fn748(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn762 github.com/goccy/llamawasm2go/p3.Fn762
func Fn762(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn763 github.com/goccy/llamawasm2go/p3.Fn763
func Fn763(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn764 github.com/goccy/llamawasm2go/p2.Fn764
func Fn764(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn765 github.com/goccy/llamawasm2go/p2.Fn765
func Fn765(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn772 github.com/goccy/llamawasm2go/p3.Fn772
func Fn772(m *base.Module, l0 int64)

//go:linkname Fn791 github.com/goccy/llamawasm2go/p3.Fn791
func Fn791(m *base.Module, l0 int64)

//go:linkname Fn792 github.com/goccy/llamawasm2go/p3.Fn792
func Fn792(m *base.Module, l0 int64)

//go:linkname Fn800 github.com/goccy/llamawasm2go/p2.Fn800
func Fn800(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn802 github.com/goccy/llamawasm2go/p1.Fn802
func Fn802(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn803 github.com/goccy/llamawasm2go/p1.Fn803
func Fn803(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn817 github.com/goccy/llamawasm2go/p1.Fn817
func Fn817(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn819 github.com/goccy/llamawasm2go/p1.Fn819
func Fn819(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn823 github.com/goccy/llamawasm2go/p1.Fn823
func Fn823(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn824 github.com/goccy/llamawasm2go/p2.Fn824
func Fn824(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn839 github.com/goccy/llamawasm2go/p1.Fn839
func Fn839(m *base.Module, l0 int32)

//go:linkname Fn841 github.com/goccy/llamawasm2go/p1.Fn841
func Fn841(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn843 github.com/goccy/llamawasm2go/p1.Fn843
func Fn843(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn846 github.com/goccy/llamawasm2go/p1.Fn846
func Fn846(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 float32, l5 int64, l6 int64) int32

//go:linkname Fn849 github.com/goccy/llamawasm2go/p3.Fn849
func Fn849(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn851 github.com/goccy/llamawasm2go/p1.Fn851
func Fn851(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn855 github.com/goccy/llamawasm2go/p3.Fn855
func Fn855(m *base.Module, l0 int64)

//go:linkname Fn856 github.com/goccy/llamawasm2go/p3.Fn856
func Fn856(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn857 github.com/goccy/llamawasm2go/p3.Fn857
func Fn857(m *base.Module, l0 int64) int32

//go:linkname Fn859 github.com/goccy/llamawasm2go/p1.Fn859
func Fn859(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn860 github.com/goccy/llamawasm2go/p2.Fn860
func Fn860(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn905 github.com/goccy/llamawasm2go/p2.Fn905
func Fn905(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn977 github.com/goccy/llamawasm2go/p3.Fn977
func Fn977(m *base.Module) int64

//go:linkname Fn1009 github.com/goccy/llamawasm2go/p2.Fn1009
func Fn1009(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1011 github.com/goccy/llamawasm2go/p2.Fn1011
func Fn1011(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1012 github.com/goccy/llamawasm2go/p2.Fn1012
func Fn1012(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1013 github.com/goccy/llamawasm2go/p2.Fn1013
func Fn1013(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1014 github.com/goccy/llamawasm2go/p2.Fn1014
func Fn1014(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) float64

//go:linkname Fn1015 github.com/goccy/llamawasm2go/p2.Fn1015
func Fn1015(m *base.Module, l0 int32, l1 int64) float64

//go:linkname Fn1016 github.com/goccy/llamawasm2go/p2.Fn1016
func Fn1016(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) float64

//go:linkname Fn1018 github.com/goccy/llamawasm2go/p1.Fn1018
func Fn1018(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32)

//go:linkname Fn1020 github.com/goccy/llamawasm2go/p3.Fn1020
func Fn1020(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1021 github.com/goccy/llamawasm2go/p2.Fn1021
func Fn1021(m *base.Module, l0 int64, l1 int64, l2 float32)

//go:linkname Fn1022 github.com/goccy/llamawasm2go/p1.Fn1022
func Fn1022(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1023 github.com/goccy/llamawasm2go/p1.Fn1023
func Fn1023(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1024 github.com/goccy/llamawasm2go/p1.Fn1024
func Fn1024(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn1025 github.com/goccy/llamawasm2go/p1.Fn1025
func Fn1025(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn1028 github.com/goccy/llamawasm2go/p1.Fn1028
func Fn1028(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64, l6 int64, l7 int64)

//go:linkname Fn1050 github.com/goccy/llamawasm2go/p2.Fn1050
func Fn1050(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1057 github.com/goccy/llamawasm2go/p3.Fn1057
func Fn1057(m *base.Module, l0 int64) int64

//go:linkname Fn1059 github.com/goccy/llamawasm2go/p2.Fn1059
func Fn1059(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1066 github.com/goccy/llamawasm2go/p2.Fn1066
func Fn1066(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1067 github.com/goccy/llamawasm2go/p3.Fn1067
func Fn1067(m *base.Module, l0 int64)

//go:linkname Fn1068 github.com/goccy/llamawasm2go/p2.Fn1068
func Fn1068(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1072 github.com/goccy/llamawasm2go/p2.Fn1072
func Fn1072(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1169 github.com/goccy/llamawasm2go/p2.Fn1169
func Fn1169(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1185 github.com/goccy/llamawasm2go/p2.Fn1185
func Fn1185(m *base.Module, l0 int64)

//go:linkname Fn1190 github.com/goccy/llamawasm2go/p3.Fn1190
func Fn1190(m *base.Module, l0 int64)

//go:linkname Fn1192 github.com/goccy/llamawasm2go/p3.Fn1192
func Fn1192(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1282 github.com/goccy/llamawasm2go/p3.Fn1282
func Fn1282(m *base.Module, l0 int64)

//go:linkname Fn1304 github.com/goccy/llamawasm2go/p3.Fn1304
func Fn1304(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1314 github.com/goccy/llamawasm2go/p2.Fn1314
func Fn1314(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p3.Fn1324
func Fn1324(m *base.Module, l0 int64) int64

//go:linkname Fn1334 github.com/goccy/llamawasm2go/p2.Fn1334
func Fn1334(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1342 github.com/goccy/llamawasm2go/p2.Fn1342
func Fn1342(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1343 github.com/goccy/llamawasm2go/p3.Fn1343
func Fn1343(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1349 github.com/goccy/llamawasm2go/p2.Fn1349
func Fn1349(m *base.Module)

//go:linkname Fn1353 github.com/goccy/llamawasm2go/p3.Fn1353
func Fn1353(m *base.Module) int64

//go:linkname Fn1354 github.com/goccy/llamawasm2go/p3.Fn1354
func Fn1354(m *base.Module) int64

//go:linkname Fn1355 github.com/goccy/llamawasm2go/p2.Fn1355
func Fn1355(m *base.Module, l0 int64) int64

//go:linkname Fn1356 github.com/goccy/llamawasm2go/p3.Fn1356
func Fn1356(m *base.Module, l0 int32) int64

//go:linkname Fn1357 github.com/goccy/llamawasm2go/p2.Fn1357
func Fn1357(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1358 github.com/goccy/llamawasm2go/p2.Fn1358
func Fn1358(m *base.Module)

//go:linkname Fn1360 github.com/goccy/llamawasm2go/p2.Fn1360
func Fn1360(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1366 github.com/goccy/llamawasm2go/p2.Fn1366
func Fn1366(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1367 github.com/goccy/llamawasm2go/p3.Fn1367
func Fn1367(m *base.Module, l0 int64)

//go:linkname Fn1368 github.com/goccy/llamawasm2go/p2.Fn1368
func Fn1368(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1369 github.com/goccy/llamawasm2go/p2.Fn1369
func Fn1369(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1370 github.com/goccy/llamawasm2go/p2.Fn1370
func Fn1370(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1371 github.com/goccy/llamawasm2go/p2.Fn1371
func Fn1371(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1372 github.com/goccy/llamawasm2go/p1.Fn1372
func Fn1372(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1373 github.com/goccy/llamawasm2go/p1.Fn1373
func Fn1373(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1374 github.com/goccy/llamawasm2go/p1.Fn1374
func Fn1374(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1375 github.com/goccy/llamawasm2go/p1.Fn1375
func Fn1375(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1376 github.com/goccy/llamawasm2go/p1.Fn1376
func Fn1376(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1377 github.com/goccy/llamawasm2go/p1.Fn1377
func Fn1377(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1378 github.com/goccy/llamawasm2go/p1.Fn1378
func Fn1378(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1379 github.com/goccy/llamawasm2go/p1.Fn1379
func Fn1379(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1380 github.com/goccy/llamawasm2go/p1.Fn1380
func Fn1380(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1381 github.com/goccy/llamawasm2go/p1.Fn1381
func Fn1381(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1382 github.com/goccy/llamawasm2go/p1.Fn1382
func Fn1382(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1383 github.com/goccy/llamawasm2go/p1.Fn1383
func Fn1383(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64) int32

//go:linkname Fn1385 github.com/goccy/llamawasm2go/p3.Fn1385
func Fn1385(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1387 github.com/goccy/llamawasm2go/p2.Fn1387
func Fn1387(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1390 github.com/goccy/llamawasm2go/p3.Fn1390
func Fn1390(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1392 github.com/goccy/llamawasm2go/p3.Fn1392
func Fn1392(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p3.Fn1394
func Fn1394(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p3.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1396 github.com/goccy/llamawasm2go/p3.Fn1396
func Fn1396(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p3.Fn1397
func Fn1397(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1400 github.com/goccy/llamawasm2go/p2.Fn1400
func Fn1400(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1401 github.com/goccy/llamawasm2go/p2.Fn1401
func Fn1401(m *base.Module, l0 int64, l1 int64) float32

//go:linkname Fn1402 github.com/goccy/llamawasm2go/p2.Fn1402
func Fn1402(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1435 github.com/goccy/llamawasm2go/p2.Fn1435
func Fn1435(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1438 github.com/goccy/llamawasm2go/p2.Fn1438
func Fn1438(m *base.Module)

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p3.Fn1445
func Fn1445(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1446 github.com/goccy/llamawasm2go/p3.Fn1446
func Fn1446(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1447 github.com/goccy/llamawasm2go/p3.Fn1447
func Fn1447(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1448 github.com/goccy/llamawasm2go/p3.Fn1448
func Fn1448(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1449 github.com/goccy/llamawasm2go/p3.Fn1449
func Fn1449(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1454 github.com/goccy/llamawasm2go/p2.Fn1454
func Fn1454(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1456 github.com/goccy/llamawasm2go/p2.Fn1456
func Fn1456(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1457 github.com/goccy/llamawasm2go/p3.Fn1457
func Fn1457(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1458 github.com/goccy/llamawasm2go/p2.Fn1458
func Fn1458(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1459 github.com/goccy/llamawasm2go/p2.Fn1459
func Fn1459(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1460 github.com/goccy/llamawasm2go/p2.Fn1460
func Fn1460(m *base.Module, l0 int64)

//go:linkname Fn1461 github.com/goccy/llamawasm2go/p3.Fn1461
func Fn1461(m *base.Module, l0 int64)

//go:linkname Fn1462 github.com/goccy/llamawasm2go/p3.Fn1462
func Fn1462(m *base.Module, l0 int64)

//go:linkname Fn1476 github.com/goccy/llamawasm2go/p2.Fn1476
func Fn1476(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1480 github.com/goccy/llamawasm2go/p3.Fn1480
func Fn1480(m *base.Module, l0 int64) int32

//go:linkname Fn1483 github.com/goccy/llamawasm2go/p2.Fn1483
func Fn1483(m *base.Module)

//go:linkname Fn1484 github.com/goccy/llamawasm2go/p2.Fn1484
func Fn1484(m *base.Module, l0 int64, l1 int32, l2 int32) int64

//go:linkname Fn1486 github.com/goccy/llamawasm2go/p2.Fn1486
func Fn1486(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1487 github.com/goccy/llamawasm2go/p3.Fn1487
func Fn1487(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1489 github.com/goccy/llamawasm2go/p2.Fn1489
func Fn1489(m *base.Module, l0 int64)

//go:linkname Fn1490 github.com/goccy/llamawasm2go/p3.Fn1490
func Fn1490(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1491 github.com/goccy/llamawasm2go/p2.Fn1491
func Fn1491(m *base.Module, l0 int64)

//go:linkname Fn1492 github.com/goccy/llamawasm2go/p2.Fn1492
func Fn1492(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1493 github.com/goccy/llamawasm2go/p1.Fn1493
func Fn1493(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1494 github.com/goccy/llamawasm2go/p2.Fn1494
func Fn1494(m *base.Module, l0 int64) int64

//go:linkname Fn1495 github.com/goccy/llamawasm2go/p2.Fn1495
func Fn1495(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1496 github.com/goccy/llamawasm2go/p3.Fn1496
func Fn1496(m *base.Module, l0 int64)

//go:linkname Fn1498 github.com/goccy/llamawasm2go/p2.Fn1498
func Fn1498(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1516 github.com/goccy/llamawasm2go/p3.Fn1516
func Fn1516(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1520 github.com/goccy/llamawasm2go/p2.Fn1520
func Fn1520(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1521 github.com/goccy/llamawasm2go/p3.Fn1521
func Fn1521(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1522 github.com/goccy/llamawasm2go/p2.Fn1522
func Fn1522(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1526 github.com/goccy/llamawasm2go/p3.Fn1526
func Fn1526(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1527 github.com/goccy/llamawasm2go/p2.Fn1527
func Fn1527(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1528 github.com/goccy/llamawasm2go/p1.Fn1528
func Fn1528(m *base.Module, l0 int64)

//go:linkname Fn1529 github.com/goccy/llamawasm2go/p2.Fn1529
func Fn1529(m *base.Module, l0 int64)

//go:linkname Fn1530 github.com/goccy/llamawasm2go/p3.Fn1530
func Fn1530(m *base.Module, l0 int64)

//go:linkname Fn1531 github.com/goccy/llamawasm2go/p2.Fn1531
func Fn1531(m *base.Module, l0 int64) int64

//go:linkname Fn1533 github.com/goccy/llamawasm2go/p2.Fn1533
func Fn1533(m *base.Module, l0 int64)

//go:linkname Fn1538 github.com/goccy/llamawasm2go/p3.Fn1538
func Fn1538(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1539 github.com/goccy/llamawasm2go/p2.Fn1539
func Fn1539(m *base.Module, l0 int64) int64

//go:linkname Fn1540 github.com/goccy/llamawasm2go/p3.Fn1540
func Fn1540(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p3.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1543 github.com/goccy/llamawasm2go/p2.Fn1543
func Fn1543(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1544 github.com/goccy/llamawasm2go/p2.Fn1544
func Fn1544(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1546 github.com/goccy/llamawasm2go/p3.Fn1546
func Fn1546(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1547 github.com/goccy/llamawasm2go/p3.Fn1547
func Fn1547(m *base.Module, l0 int64)

//go:linkname Fn1549 github.com/goccy/llamawasm2go/p2.Fn1549
func Fn1549(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1554 github.com/goccy/llamawasm2go/p1.Fn1554
func Fn1554(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn1558 github.com/goccy/llamawasm2go/p1.Fn1558
func Fn1558(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1560 github.com/goccy/llamawasm2go/p2.Fn1560
func Fn1560(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1561 github.com/goccy/llamawasm2go/p2.Fn1561
func Fn1561(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64)

//go:linkname Fn1562 github.com/goccy/llamawasm2go/p2.Fn1562
func Fn1562(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int64)

//go:linkname Fn1579 github.com/goccy/llamawasm2go/p3.Fn1579
func Fn1579(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1580 github.com/goccy/llamawasm2go/p3.Fn1580
func Fn1580(m *base.Module, l0 int64)

//go:linkname Fn1581 github.com/goccy/llamawasm2go/p3.Fn1581
func Fn1581(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1582 github.com/goccy/llamawasm2go/p3.Fn1582
func Fn1582(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32)

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p3.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1584 github.com/goccy/llamawasm2go/p3.Fn1584
func Fn1584(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1633 github.com/goccy/llamawasm2go/p2.Fn1633
func Fn1633(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1648 github.com/goccy/llamawasm2go/p1.Fn1648
func Fn1648(m *base.Module, l0 int64)

//go:linkname Fn1651 github.com/goccy/llamawasm2go/p3.Fn1651
func Fn1651(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1652 github.com/goccy/llamawasm2go/p3.Fn1652
func Fn1652(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1653 github.com/goccy/llamawasm2go/p2.Fn1653
func Fn1653(m *base.Module, l0 int64)

//go:linkname Fn1655 github.com/goccy/llamawasm2go/p2.Fn1655
func Fn1655(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1656 github.com/goccy/llamawasm2go/p3.Fn1656
func Fn1656(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1659 github.com/goccy/llamawasm2go/p2.Fn1659
func Fn1659(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1661 github.com/goccy/llamawasm2go/p2.Fn1661
func Fn1661(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1664 github.com/goccy/llamawasm2go/p1.Fn1664
func Fn1664(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int32, l13 int32, l14 int32) int64

//go:linkname Fn1668 github.com/goccy/llamawasm2go/p2.Fn1668
func Fn1668(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1669 github.com/goccy/llamawasm2go/p2.Fn1669
func Fn1669(m *base.Module, l0 int64) int64

//go:linkname Fn1678 github.com/goccy/llamawasm2go/p1.Fn1678
func Fn1678(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1682 github.com/goccy/llamawasm2go/p2.Fn1682
func Fn1682(m *base.Module, l0 int64) int64

//go:linkname Fn1684 github.com/goccy/llamawasm2go/p2.Fn1684
func Fn1684(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 float32, l11 int32) int64

//go:linkname Fn1690 github.com/goccy/llamawasm2go/p2.Fn1690
func Fn1690(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1696 github.com/goccy/llamawasm2go/p2.Fn1696
func Fn1696(m *base.Module, l0 int64) int64

//go:linkname Fn1747 github.com/goccy/llamawasm2go/p3.Fn1747
func Fn1747(m *base.Module, l0 int64) int64

//go:linkname Fn1754 github.com/goccy/llamawasm2go/p3.Fn1754
func Fn1754(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1755 github.com/goccy/llamawasm2go/p3.Fn1755
func Fn1755(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1756 github.com/goccy/llamawasm2go/p3.Fn1756
func Fn1756(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1757 github.com/goccy/llamawasm2go/p3.Fn1757
func Fn1757(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1758 github.com/goccy/llamawasm2go/p3.Fn1758
func Fn1758(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1760 github.com/goccy/llamawasm2go/p3.Fn1760
func Fn1760(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1761 github.com/goccy/llamawasm2go/p3.Fn1761
func Fn1761(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1762 github.com/goccy/llamawasm2go/p3.Fn1762
func Fn1762(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1766 github.com/goccy/llamawasm2go/p3.Fn1766
func Fn1766(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1767 github.com/goccy/llamawasm2go/p3.Fn1767
func Fn1767(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1768 github.com/goccy/llamawasm2go/p3.Fn1768
func Fn1768(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1769 github.com/goccy/llamawasm2go/p3.Fn1769
func Fn1769(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1770 github.com/goccy/llamawasm2go/p2.Fn1770
func Fn1770(m *base.Module, l0 int64) int32

//go:linkname Fn1773 github.com/goccy/llamawasm2go/p3.Fn1773
func Fn1773(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1779 github.com/goccy/llamawasm2go/p3.Fn1779
func Fn1779(m *base.Module, l0 int64) int32

//go:linkname Fn1780 github.com/goccy/llamawasm2go/p3.Fn1780
func Fn1780(m *base.Module, l0 int64) int32

//go:linkname Fn1788 github.com/goccy/llamawasm2go/p3.Fn1788
func Fn1788(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1791 github.com/goccy/llamawasm2go/p2.Fn1791
func Fn1791(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1792 github.com/goccy/llamawasm2go/p2.Fn1792
func Fn1792(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1793 github.com/goccy/llamawasm2go/p3.Fn1793
func Fn1793(m *base.Module)

//go:linkname Fn1794 github.com/goccy/llamawasm2go/p2.Fn1794
func Fn1794(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1795 github.com/goccy/llamawasm2go/p2.Fn1795
func Fn1795(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1798 github.com/goccy/llamawasm2go/p2.Fn1798
func Fn1798(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1799 github.com/goccy/llamawasm2go/p2.Fn1799
func Fn1799(m *base.Module, l0 int64)

//go:linkname Fn1800 github.com/goccy/llamawasm2go/p2.Fn1800
func Fn1800(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1801 github.com/goccy/llamawasm2go/p3.Fn1801
func Fn1801(m *base.Module, l0 int64)

//go:linkname Fn1802 github.com/goccy/llamawasm2go/p2.Fn1802
func Fn1802(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1803 github.com/goccy/llamawasm2go/p3.Fn1803
func Fn1803(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1804 github.com/goccy/llamawasm2go/p2.Fn1804
func Fn1804(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1806 github.com/goccy/llamawasm2go/p2.Fn1806
func Fn1806(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1807 github.com/goccy/llamawasm2go/p2.Fn1807
func Fn1807(m *base.Module, l0 int64)

//go:linkname Fn1810 github.com/goccy/llamawasm2go/p3.Fn1810
func Fn1810(m *base.Module, l0 int64)

//go:linkname Fn1813 github.com/goccy/llamawasm2go/p2.Fn1813
func Fn1813(m *base.Module, l0 int64) int64

//go:linkname Fn1815 github.com/goccy/llamawasm2go/p3.Fn1815
func Fn1815(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1816 github.com/goccy/llamawasm2go/p3.Fn1816
func Fn1816(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1822 github.com/goccy/llamawasm2go/p2.Fn1822
func Fn1822(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn1834 github.com/goccy/llamawasm2go/p1.Fn1834
func Fn1834(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1835 github.com/goccy/llamawasm2go/p2.Fn1835
func Fn1835(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1836 github.com/goccy/llamawasm2go/p2.Fn1836
func Fn1836(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1838 github.com/goccy/llamawasm2go/p2.Fn1838
func Fn1838(m *base.Module, l0 int64) int64

//go:linkname Fn1839 github.com/goccy/llamawasm2go/p1.Fn1839
func Fn1839(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1843 github.com/goccy/llamawasm2go/p2.Fn1843
func Fn1843(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1849 github.com/goccy/llamawasm2go/p3.Fn1849
func Fn1849(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1855 github.com/goccy/llamawasm2go/p2.Fn1855
func Fn1855(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1877 github.com/goccy/llamawasm2go/p3.Fn1877
func Fn1877(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1879 github.com/goccy/llamawasm2go/p3.Fn1879
func Fn1879(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1880 github.com/goccy/llamawasm2go/p2.Fn1880
func Fn1880(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1895 github.com/goccy/llamawasm2go/p2.Fn1895
func Fn1895(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1897 github.com/goccy/llamawasm2go/p3.Fn1897
func Fn1897(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn1927 github.com/goccy/llamawasm2go/p2.Fn1927
func Fn1927(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64) int64

//go:linkname Fn1951 github.com/goccy/llamawasm2go/p2.Fn1951
func Fn1951(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64) int64

//go:linkname Fn1964 github.com/goccy/llamawasm2go/p2.Fn1964
func Fn1964(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int64, l12 int64, l13 int64) int64

//go:linkname Fn2006 github.com/goccy/llamawasm2go/p3.Fn2006
func Fn2006(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn2007 github.com/goccy/llamawasm2go/p3.Fn2007
func Fn2007(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn2008 github.com/goccy/llamawasm2go/p1.Fn2008
func Fn2008(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int64, l12 int64) int64

//go:linkname Fn2021 github.com/goccy/llamawasm2go/p2.Fn2021
func Fn2021(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2046 github.com/goccy/llamawasm2go/p3.Fn2046
func Fn2046(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2047 github.com/goccy/llamawasm2go/p2.Fn2047
func Fn2047(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn2049 github.com/goccy/llamawasm2go/p3.Fn2049
func Fn2049(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2050 github.com/goccy/llamawasm2go/p3.Fn2050
func Fn2050(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn2053 github.com/goccy/llamawasm2go/p2.Fn2053
func Fn2053(m *base.Module, l0 int64) int64

//go:linkname Fn2060 github.com/goccy/llamawasm2go/p2.Fn2060
func Fn2060(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2061 github.com/goccy/llamawasm2go/p2.Fn2061
func Fn2061(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn2062 github.com/goccy/llamawasm2go/p1.Fn2062
func Fn2062(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2063 github.com/goccy/llamawasm2go/p3.Fn2063
func Fn2063(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2064 github.com/goccy/llamawasm2go/p3.Fn2064
func Fn2064(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2089 github.com/goccy/llamawasm2go/p1.Fn2089
func Fn2089(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn2091 github.com/goccy/llamawasm2go/p3.Fn2091
func Fn2091(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 int64) int64

//go:linkname Fn2097 github.com/goccy/llamawasm2go/p2.Fn2097
func Fn2097(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64) int64

//go:linkname Fn2124 github.com/goccy/llamawasm2go/p1.Fn2124
func Fn2124(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64) int64

//go:linkname Fn2160 github.com/goccy/llamawasm2go/p2.Fn2160
func Fn2160(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2161 github.com/goccy/llamawasm2go/p3.Fn2161
func Fn2161(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2162 github.com/goccy/llamawasm2go/p2.Fn2162
func Fn2162(m *base.Module, l0 int64) int64

//go:linkname Fn2163 github.com/goccy/llamawasm2go/p3.Fn2163
func Fn2163(m *base.Module, l0 int64) int64

//go:linkname Fn2175 github.com/goccy/llamawasm2go/p3.Fn2175
func Fn2175(m *base.Module, l0 int64) int64

//go:linkname Fn2176 github.com/goccy/llamawasm2go/p3.Fn2176
func Fn2176(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2177 github.com/goccy/llamawasm2go/p3.Fn2177
func Fn2177(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2189 github.com/goccy/llamawasm2go/p2.Fn2189
func Fn2189(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2190 github.com/goccy/llamawasm2go/p2.Fn2190
func Fn2190(m *base.Module, l0 int32, l1 int64) int32

//go:linkname Fn2193 github.com/goccy/llamawasm2go/p3.Fn2193
func Fn2193(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2194 github.com/goccy/llamawasm2go/p2.Fn2194
func Fn2194(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2195 github.com/goccy/llamawasm2go/p2.Fn2195
func Fn2195(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2196 github.com/goccy/llamawasm2go/p2.Fn2196
func Fn2196(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2199 github.com/goccy/llamawasm2go/p2.Fn2199
func Fn2199(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn2203 github.com/goccy/llamawasm2go/p2.Fn2203
func Fn2203(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn2204 github.com/goccy/llamawasm2go/p3.Fn2204
func Fn2204(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2205 github.com/goccy/llamawasm2go/p3.Fn2205
func Fn2205(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2206 github.com/goccy/llamawasm2go/p2.Fn2206
func Fn2206(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2207 github.com/goccy/llamawasm2go/p2.Fn2207
func Fn2207(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2208 github.com/goccy/llamawasm2go/p3.Fn2208
func Fn2208(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2209 github.com/goccy/llamawasm2go/p1.Fn2209
func Fn2209(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2212 github.com/goccy/llamawasm2go/p3.Fn2212
func Fn2212(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2213 github.com/goccy/llamawasm2go/p2.Fn2213
func Fn2213(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn2216 github.com/goccy/llamawasm2go/p3.Fn2216
func Fn2216(m *base.Module, l0 int64)

//go:linkname Fn2218 github.com/goccy/llamawasm2go/p2.Fn2218
func Fn2218(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2220 github.com/goccy/llamawasm2go/p2.Fn2220
func Fn2220(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2222 github.com/goccy/llamawasm2go/p1.Fn2222
func Fn2222(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int32

//go:linkname Fn2224 github.com/goccy/llamawasm2go/p2.Fn2224
func Fn2224(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2334 github.com/goccy/llamawasm2go/p2.Fn2334
func Fn2334(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64) int64

//go:linkname Fn2410 github.com/goccy/llamawasm2go/p2.Fn2410
func Fn2410(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32) int64

//go:linkname Fn2412 github.com/goccy/llamawasm2go/p3.Fn2412
func Fn2412(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2420 github.com/goccy/llamawasm2go/p2.Fn2420
func Fn2420(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn2421 github.com/goccy/llamawasm2go/p2.Fn2421
func Fn2421(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2422 github.com/goccy/llamawasm2go/p2.Fn2422
func Fn2422(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2703 github.com/goccy/llamawasm2go/p2.Fn2703
func Fn2703(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2970 github.com/goccy/llamawasm2go/p2.Fn2970
func Fn2970(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2971 github.com/goccy/llamawasm2go/p1.Fn2971
func Fn2971(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64)

//go:linkname Fn2972 github.com/goccy/llamawasm2go/p2.Fn2972
func Fn2972(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2973 github.com/goccy/llamawasm2go/p3.Fn2973
func Fn2973(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2974 github.com/goccy/llamawasm2go/p3.Fn2974
func Fn2974(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2975 github.com/goccy/llamawasm2go/p3.Fn2975
func Fn2975(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2976 github.com/goccy/llamawasm2go/p2.Fn2976
func Fn2976(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2987 github.com/goccy/llamawasm2go/p1.Fn2987
func Fn2987(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2988 github.com/goccy/llamawasm2go/p1.Fn2988
func Fn2988(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2996 github.com/goccy/llamawasm2go/p3.Fn2996
func Fn2996(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2999 github.com/goccy/llamawasm2go/p2.Fn2999
func Fn2999(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3000 github.com/goccy/llamawasm2go/p3.Fn3000
func Fn3000(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3001 github.com/goccy/llamawasm2go/p2.Fn3001
func Fn3001(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3002 github.com/goccy/llamawasm2go/p1.Fn3002
func Fn3002(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3003 github.com/goccy/llamawasm2go/p1.Fn3003
func Fn3003(m *base.Module, l0 int64)

//go:linkname Fn3006 github.com/goccy/llamawasm2go/p3.Fn3006
func Fn3006(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn3007 github.com/goccy/llamawasm2go/p3.Fn3007
func Fn3007(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3012 github.com/goccy/llamawasm2go/p2.Fn3012
func Fn3012(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn3019 github.com/goccy/llamawasm2go/p3.Fn3019
func Fn3019(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3041 github.com/goccy/llamawasm2go/p1.Fn3041
func Fn3041(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3051 github.com/goccy/llamawasm2go/p2.Fn3051
func Fn3051(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn3053 github.com/goccy/llamawasm2go/p1.Fn3053
func Fn3053(m *base.Module)

//go:linkname Fn3055 github.com/goccy/llamawasm2go/p2.Fn3055
func Fn3055(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3056 github.com/goccy/llamawasm2go/p2.Fn3056
func Fn3056(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3059 github.com/goccy/llamawasm2go/p2.Fn3059
func Fn3059(m *base.Module, l0 int64) int64

//go:linkname Fn3060 github.com/goccy/llamawasm2go/p1.Fn3060
func Fn3060(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3061 github.com/goccy/llamawasm2go/p2.Fn3061
func Fn3061(m *base.Module, l0 int64) int64

//go:linkname Fn3062 github.com/goccy/llamawasm2go/p2.Fn3062
func Fn3062(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3074 github.com/goccy/llamawasm2go/p2.Fn3074
func Fn3074(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3093 github.com/goccy/llamawasm2go/p3.Fn3093
func Fn3093(m *base.Module, l0 int64)

//go:linkname Fn3094 github.com/goccy/llamawasm2go/p3.Fn3094
func Fn3094(m *base.Module) int64

//go:linkname Fn3095 github.com/goccy/llamawasm2go/p1.Fn3095
func Fn3095(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn3097 github.com/goccy/llamawasm2go/p3.Fn3097
func Fn3097(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3099 github.com/goccy/llamawasm2go/p3.Fn3099
func Fn3099(m *base.Module) int64

//go:linkname Fn3103 github.com/goccy/llamawasm2go/p3.Fn3103
func Fn3103(m *base.Module, l0 int64)

//go:linkname Fn3212 github.com/goccy/llamawasm2go/p2.Fn3212
func Fn3212(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3213 github.com/goccy/llamawasm2go/p3.Fn3213
func Fn3213(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3214 github.com/goccy/llamawasm2go/p2.Fn3214
func Fn3214(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn3216 github.com/goccy/llamawasm2go/p2.Fn3216
func Fn3216(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn3218 github.com/goccy/llamawasm2go/p1.Fn3218
func Fn3218(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn3219 github.com/goccy/llamawasm2go/p2.Fn3219
func Fn3219(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3221 github.com/goccy/llamawasm2go/p3.Fn3221
func Fn3221(m *base.Module, l0 int64)

//go:linkname Fn3222 github.com/goccy/llamawasm2go/p3.Fn3222
func Fn3222(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3224 github.com/goccy/llamawasm2go/p1.Fn3224
func Fn3224(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn3226 github.com/goccy/llamawasm2go/p2.Fn3226
func Fn3226(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn3227 github.com/goccy/llamawasm2go/p2.Fn3227
func Fn3227(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3228 github.com/goccy/llamawasm2go/p2.Fn3228
func Fn3228(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3232 github.com/goccy/llamawasm2go/p2.Fn3232
func Fn3232(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3236 github.com/goccy/llamawasm2go/p2.Fn3236
func Fn3236(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn3240 github.com/goccy/llamawasm2go/p2.Fn3240
func Fn3240(m *base.Module, l0 int64)

//go:linkname Fn3242 github.com/goccy/llamawasm2go/p1.Fn3242
func Fn3242(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3248 github.com/goccy/llamawasm2go/p2.Fn3248
func Fn3248(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn3249 github.com/goccy/llamawasm2go/p3.Fn3249
func Fn3249(m *base.Module, l0 int64)

//go:linkname Fn3258 github.com/goccy/llamawasm2go/p2.Fn3258
func Fn3258(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3259 github.com/goccy/llamawasm2go/p2.Fn3259
func Fn3259(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3260 github.com/goccy/llamawasm2go/p2.Fn3260
func Fn3260(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3268 github.com/goccy/llamawasm2go/p1.Fn3268
func Fn3268(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn3278 github.com/goccy/llamawasm2go/p2.Fn3278
func Fn3278(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3279 github.com/goccy/llamawasm2go/p1.Fn3279
func Fn3279(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn3280 github.com/goccy/llamawasm2go/p2.Fn3280
func Fn3280(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3281 github.com/goccy/llamawasm2go/p3.Fn3281
func Fn3281(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3282 github.com/goccy/llamawasm2go/p2.Fn3282
func Fn3282(m *base.Module)

//go:linkname Fn3284 github.com/goccy/llamawasm2go/p2.Fn3284
func Fn3284(m *base.Module)

//go:linkname Fn3291 github.com/goccy/llamawasm2go/p3.Fn3291
func Fn3291(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3294 github.com/goccy/llamawasm2go/p3.Fn3294
func Fn3294(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3297 github.com/goccy/llamawasm2go/p2.Fn3297
func Fn3297(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn3298 github.com/goccy/llamawasm2go/p3.Fn3298
func Fn3298(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3299 github.com/goccy/llamawasm2go/p3.Fn3299
func Fn3299(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3300 github.com/goccy/llamawasm2go/p3.Fn3300
func Fn3300(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3301 github.com/goccy/llamawasm2go/p1.Fn3301
func Fn3301(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3302 github.com/goccy/llamawasm2go/p1.Fn3302
func Fn3302(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3303 github.com/goccy/llamawasm2go/p2.Fn3303
func Fn3303(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3308 github.com/goccy/llamawasm2go/p1.Fn3308
func Fn3308(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3370 github.com/goccy/llamawasm2go/p3.Fn3370
func Fn3370(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3380 github.com/goccy/llamawasm2go/p3.Fn3380
func Fn3380(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn3381 github.com/goccy/llamawasm2go/p3.Fn3381
func Fn3381(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn3384 github.com/goccy/llamawasm2go/p2.Fn3384
func Fn3384(m *base.Module, l0 int64) int64

//go:linkname Fn3385 github.com/goccy/llamawasm2go/p3.Fn3385
func Fn3385(m *base.Module, l0 int64)

//go:linkname Fn3388 github.com/goccy/llamawasm2go/p2.Fn3388
func Fn3388(m *base.Module, l0 int64)

//go:linkname Fn3389 github.com/goccy/llamawasm2go/p3.Fn3389
func Fn3389(m *base.Module, l0 int64)

//go:linkname Fn3416 github.com/goccy/llamawasm2go/p2.Fn3416
func Fn3416(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3420 github.com/goccy/llamawasm2go/p2.Fn3420
func Fn3420(m *base.Module) int64

//go:linkname Fn3426 github.com/goccy/llamawasm2go/p3.Fn3426
func Fn3426(m *base.Module, l0 float64) float32

//go:linkname Fn3427 github.com/goccy/llamawasm2go/p3.Fn3427
func Fn3427(m *base.Module, l0 float64) float32

//go:linkname Fn3428 github.com/goccy/llamawasm2go/p2.Fn3428
func Fn3428(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn3430 github.com/goccy/llamawasm2go/p2.Fn3430
func Fn3430(m *base.Module, l0 float32) float32

//go:linkname Fn3431 github.com/goccy/llamawasm2go/p3.Fn3431
func Fn3431(m *base.Module, l0 float64) float64

//go:linkname Fn3434 github.com/goccy/llamawasm2go/p3.Fn3434
func Fn3434(m *base.Module, l0 int32) float32

//go:linkname Fn3435 github.com/goccy/llamawasm2go/p3.Fn3435
func Fn3435(m *base.Module, l0 int32) float32

//go:linkname Fn3436 github.com/goccy/llamawasm2go/p3.Fn3436
func Fn3436(m *base.Module, l0 float32) float32

//go:linkname Fn3437 github.com/goccy/llamawasm2go/p3.Fn3437
func Fn3437(m *base.Module, l0 float32) float32

//go:linkname Fn3438 github.com/goccy/llamawasm2go/p2.Fn3438
func Fn3438(m *base.Module, l0 float32) float32

//go:linkname Fn3441 github.com/goccy/llamawasm2go/p2.Fn3441
func Fn3441(m *base.Module, l0 float64) float64

//go:linkname Fn3442 github.com/goccy/llamawasm2go/p2.Fn3442
func Fn3442(m *base.Module, l0 float64) float64

//go:linkname Fn3443 github.com/goccy/llamawasm2go/p3.Fn3443
func Fn3443(m *base.Module, l0 float32) float32

//go:linkname Fn3445 github.com/goccy/llamawasm2go/p2.Fn3445
func Fn3445(m *base.Module, l0 float32) float32

//go:linkname Fn3447 github.com/goccy/llamawasm2go/p2.Fn3447
func Fn3447(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn3448 github.com/goccy/llamawasm2go/p3.Fn3448
func Fn3448(m *base.Module, l0 float32) float32

//go:linkname Fn3451 github.com/goccy/llamawasm2go/p3.Fn3451
func Fn3451(m *base.Module, l0 float32) float32

//go:linkname Fn3454 github.com/goccy/llamawasm2go/p2.Fn3454
func Fn3454(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3466 github.com/goccy/llamawasm2go/p3.Fn3466
func Fn3466(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3474 github.com/goccy/llamawasm2go/p3.Fn3474
func Fn3474(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3479 github.com/goccy/llamawasm2go/p3.Fn3479
func Fn3479(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3480 github.com/goccy/llamawasm2go/p1.Fn3480
func Fn3480(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3485 github.com/goccy/llamawasm2go/p3.Fn3485
func Fn3485(m *base.Module, l0 int64) int32

//go:linkname Fn3488 github.com/goccy/llamawasm2go/p2.Fn3488
func Fn3488(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3495 github.com/goccy/llamawasm2go/p2.Fn3495
func Fn3495(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3496 github.com/goccy/llamawasm2go/p3.Fn3496
func Fn3496(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3498 github.com/goccy/llamawasm2go/p2.Fn3498
func Fn3498(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3499 github.com/goccy/llamawasm2go/p3.Fn3499
func Fn3499(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3502 github.com/goccy/llamawasm2go/p2.Fn3502
func Fn3502(m *base.Module, l0 int64) int64

//go:linkname Fn3511 github.com/goccy/llamawasm2go/p3.Fn3511
func Fn3511(m *base.Module, l0 float64, l1 int64) float64

//go:linkname Fn3514 github.com/goccy/llamawasm2go/p2.Fn3514
func Fn3514(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn3515 github.com/goccy/llamawasm2go/p3.Fn3515
func Fn3515(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32)

//go:linkname Fn3520 github.com/goccy/llamawasm2go/p3.Fn3520
func Fn3520(m *base.Module)

//go:linkname Fn3523 github.com/goccy/llamawasm2go/p2.Fn3523
func Fn3523(m *base.Module, l0 int64)

//go:linkname Fn3524 github.com/goccy/llamawasm2go/p3.Fn3524
func Fn3524(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3525 github.com/goccy/llamawasm2go/p2.Fn3525
func Fn3525(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3601 github.com/goccy/llamawasm2go/p3.Fn3601
func Fn3601(m *base.Module, l0 int32)
