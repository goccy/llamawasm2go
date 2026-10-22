//go:build !arm64 && (!amd64 || !amd64.v2)

package p1

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn22 github.com/goccy/llamawasm2go/p2.Fn22
func Fn22(m *base.Module, l0 int64) int64

//go:linkname Fn32 github.com/goccy/llamawasm2go/p2.Fn32
func Fn32(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn46 github.com/goccy/llamawasm2go/p2.Fn46
func Fn46(m *base.Module, l0 int64)

//go:linkname Fn47 github.com/goccy/llamawasm2go/p3.Fn47
func Fn47(m *base.Module, l0 int64)

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

//go:linkname Fn59 github.com/goccy/llamawasm2go/p2.Fn59
func Fn59(m *base.Module)

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

//go:linkname Fn259 github.com/goccy/llamawasm2go/p2.Fn259
func Fn259(m *base.Module, l0 int64, l1 int64, l2 int32)

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

//go:linkname Fn266 github.com/goccy/llamawasm2go/p3.Fn266
func Fn266(m *base.Module, l0 int64, l1 int64)

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

//go:linkname Fn273 github.com/goccy/llamawasm2go/p2.Fn273
func Fn273(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn275 github.com/goccy/llamawasm2go/p3.Fn275
func Fn275(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn277 github.com/goccy/llamawasm2go/p2.Fn277
func Fn277(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn278 github.com/goccy/llamawasm2go/p2.Fn278
func Fn278(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn281 github.com/goccy/llamawasm2go/p2.Fn281
func Fn281(m *base.Module, l0 int64) int64

//go:linkname Fn282 github.com/goccy/llamawasm2go/p2.Fn282
func Fn282(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn285 github.com/goccy/llamawasm2go/p3.Fn285
func Fn285(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn286 github.com/goccy/llamawasm2go/p2.Fn286
func Fn286(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn292 github.com/goccy/llamawasm2go/p2.Fn292
func Fn292(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn293 github.com/goccy/llamawasm2go/p2.Fn293
func Fn293(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn341 github.com/goccy/llamawasm2go/p2.Fn341
func Fn341(m *base.Module, l0 int64) int32

//go:linkname Fn342 github.com/goccy/llamawasm2go/p3.Fn342
func Fn342(m *base.Module, l0 int64)

//go:linkname Fn343 github.com/goccy/llamawasm2go/p3.Fn343
func Fn343(m *base.Module, l0 int64)

//go:linkname Fn359 github.com/goccy/llamawasm2go/p3.Fn359
func Fn359(m *base.Module, l0 int32, l1 int64)

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

//go:linkname Fn372 github.com/goccy/llamawasm2go/p2.Fn372
func Fn372(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn373 github.com/goccy/llamawasm2go/p2.Fn373
func Fn373(m *base.Module, l0 int64)

//go:linkname Fn374 github.com/goccy/llamawasm2go/p2.Fn374
func Fn374(m *base.Module, l0 int64)

//go:linkname Fn376 github.com/goccy/llamawasm2go/p2.Fn376
func Fn376(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn377 github.com/goccy/llamawasm2go/p3.Fn377
func Fn377(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn383 github.com/goccy/llamawasm2go/p2.Fn383
func Fn383(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn384 github.com/goccy/llamawasm2go/p3.Fn384
func Fn384(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn386 github.com/goccy/llamawasm2go/p2.Fn386
func Fn386(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn387 github.com/goccy/llamawasm2go/p2.Fn387
func Fn387(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn393 github.com/goccy/llamawasm2go/p2.Fn393
func Fn393(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn394 github.com/goccy/llamawasm2go/p2.Fn394
func Fn394(m *base.Module, l0 int64)

//go:linkname Fn397 github.com/goccy/llamawasm2go/p3.Fn397
func Fn397(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn398 github.com/goccy/llamawasm2go/p2.Fn398
func Fn398(m *base.Module, l0 int64)

//go:linkname Fn399 github.com/goccy/llamawasm2go/p2.Fn399
func Fn399(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn400 github.com/goccy/llamawasm2go/p2.Fn400
func Fn400(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn401 github.com/goccy/llamawasm2go/p3.Fn401
func Fn401(m *base.Module)

//go:linkname Fn404 github.com/goccy/llamawasm2go/p2.Fn404
func Fn404(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn405 github.com/goccy/llamawasm2go/p2.Fn405
func Fn405(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn413 github.com/goccy/llamawasm2go/p2.Fn413
func Fn413(m *base.Module, l0 int64) int32

//go:linkname Fn414 github.com/goccy/llamawasm2go/p2.Fn414
func Fn414(m *base.Module, l0 int64)

//go:linkname Fn433 github.com/goccy/llamawasm2go/p2.Fn433
func Fn433(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn434 github.com/goccy/llamawasm2go/p2.Fn434
func Fn434(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn435 github.com/goccy/llamawasm2go/p2.Fn435
func Fn435(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn436 github.com/goccy/llamawasm2go/p2.Fn436
func Fn436(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn438 github.com/goccy/llamawasm2go/p3.Fn438
func Fn438(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn439 github.com/goccy/llamawasm2go/p3.Fn439
func Fn439(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn449 github.com/goccy/llamawasm2go/p3.Fn449
func Fn449(m *base.Module) int64

//go:linkname Fn461 github.com/goccy/llamawasm2go/p3.Fn461
func Fn461(m *base.Module, l0 int64) int32

//go:linkname Fn462 github.com/goccy/llamawasm2go/p3.Fn462
func Fn462(m *base.Module, l0 int64) int32

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

//go:linkname Fn492 github.com/goccy/llamawasm2go/p3.Fn492
func Fn492(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn494 github.com/goccy/llamawasm2go/p3.Fn494
func Fn494(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn495 github.com/goccy/llamawasm2go/p2.Fn495
func Fn495(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn496 github.com/goccy/llamawasm2go/p3.Fn496
func Fn496(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn497 github.com/goccy/llamawasm2go/p3.Fn497
func Fn497(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn498 github.com/goccy/llamawasm2go/p3.Fn498
func Fn498(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn499 github.com/goccy/llamawasm2go/p3.Fn499
func Fn499(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn500 github.com/goccy/llamawasm2go/p3.Fn500
func Fn500(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn502 github.com/goccy/llamawasm2go/p3.Fn502
func Fn502(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn504 github.com/goccy/llamawasm2go/p3.Fn504
func Fn504(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn505 github.com/goccy/llamawasm2go/p3.Fn505
func Fn505(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn507 github.com/goccy/llamawasm2go/p2.Fn507
func Fn507(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn508 github.com/goccy/llamawasm2go/p3.Fn508
func Fn508(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn509 github.com/goccy/llamawasm2go/p2.Fn509
func Fn509(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn510 github.com/goccy/llamawasm2go/p3.Fn510
func Fn510(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn511 github.com/goccy/llamawasm2go/p3.Fn511
func Fn511(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn513 github.com/goccy/llamawasm2go/p3.Fn513
func Fn513(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn514 github.com/goccy/llamawasm2go/p3.Fn514
func Fn514(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn515 github.com/goccy/llamawasm2go/p3.Fn515
func Fn515(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn516 github.com/goccy/llamawasm2go/p3.Fn516
func Fn516(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn517 github.com/goccy/llamawasm2go/p3.Fn517
func Fn517(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn518 github.com/goccy/llamawasm2go/p3.Fn518
func Fn518(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn519 github.com/goccy/llamawasm2go/p3.Fn519
func Fn519(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn520 github.com/goccy/llamawasm2go/p2.Fn520
func Fn520(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn521 github.com/goccy/llamawasm2go/p3.Fn521
func Fn521(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn522 github.com/goccy/llamawasm2go/p3.Fn522
func Fn522(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn523 github.com/goccy/llamawasm2go/p3.Fn523
func Fn523(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn524 github.com/goccy/llamawasm2go/p3.Fn524
func Fn524(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn525 github.com/goccy/llamawasm2go/p3.Fn525
func Fn525(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32) int64

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

//go:linkname Fn533 github.com/goccy/llamawasm2go/p2.Fn533
func Fn533(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn534 github.com/goccy/llamawasm2go/p3.Fn534
func Fn534(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn535 github.com/goccy/llamawasm2go/p3.Fn535
func Fn535(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn536 github.com/goccy/llamawasm2go/p2.Fn536
func Fn536(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn537 github.com/goccy/llamawasm2go/p2.Fn537
func Fn537(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn538 github.com/goccy/llamawasm2go/p3.Fn538
func Fn538(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn539 github.com/goccy/llamawasm2go/p3.Fn539
func Fn539(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn541 github.com/goccy/llamawasm2go/p3.Fn541
func Fn541(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn542 github.com/goccy/llamawasm2go/p3.Fn542
func Fn542(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn543 github.com/goccy/llamawasm2go/p3.Fn543
func Fn543(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn544 github.com/goccy/llamawasm2go/p2.Fn544
func Fn544(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn545 github.com/goccy/llamawasm2go/p2.Fn545
func Fn545(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn546 github.com/goccy/llamawasm2go/p2.Fn546
func Fn546(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn547 github.com/goccy/llamawasm2go/p2.Fn547
func Fn547(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn548 github.com/goccy/llamawasm2go/p3.Fn548
func Fn548(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

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

//go:linkname Fn556 github.com/goccy/llamawasm2go/p3.Fn556
func Fn556(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn557 github.com/goccy/llamawasm2go/p3.Fn557
func Fn557(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn559 github.com/goccy/llamawasm2go/p3.Fn559
func Fn559(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn562 github.com/goccy/llamawasm2go/p3.Fn562
func Fn562(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32) int64

//go:linkname Fn563 github.com/goccy/llamawasm2go/p3.Fn563
func Fn563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32, l7 float32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32) int64

//go:linkname Fn565 github.com/goccy/llamawasm2go/p3.Fn565
func Fn565(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn568 github.com/goccy/llamawasm2go/p2.Fn568
func Fn568(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn569 github.com/goccy/llamawasm2go/p3.Fn569
func Fn569(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32) int64

//go:linkname Fn570 github.com/goccy/llamawasm2go/p3.Fn570
func Fn570(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32) int64

//go:linkname Fn571 github.com/goccy/llamawasm2go/p2.Fn571
func Fn571(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn572 github.com/goccy/llamawasm2go/p2.Fn572
func Fn572(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn573 github.com/goccy/llamawasm2go/p3.Fn573
func Fn573(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn574 github.com/goccy/llamawasm2go/p2.Fn574
func Fn574(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn575 github.com/goccy/llamawasm2go/p3.Fn575
func Fn575(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn576 github.com/goccy/llamawasm2go/p3.Fn576
func Fn576(m *base.Module, l0 int64, l1 float32, l2 float32) int64

//go:linkname Fn577 github.com/goccy/llamawasm2go/p2.Fn577
func Fn577(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 float32, l6 float32, l7 float32) int64

//go:linkname Fn578 github.com/goccy/llamawasm2go/p2.Fn578
func Fn578(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn580 github.com/goccy/llamawasm2go/p2.Fn580
func Fn580(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn581 github.com/goccy/llamawasm2go/p2.Fn581
func Fn581(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn582 github.com/goccy/llamawasm2go/p2.Fn582
func Fn582(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn587 github.com/goccy/llamawasm2go/p2.Fn587
func Fn587(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn590 github.com/goccy/llamawasm2go/p2.Fn590
func Fn590(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn592 github.com/goccy/llamawasm2go/p3.Fn592
func Fn592(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn595 github.com/goccy/llamawasm2go/p2.Fn595
func Fn595(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn597 github.com/goccy/llamawasm2go/p3.Fn597
func Fn597(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn598 github.com/goccy/llamawasm2go/p3.Fn598
func Fn598(m *base.Module, l0 int64)

//go:linkname Fn602 github.com/goccy/llamawasm2go/p3.Fn602
func Fn602(m *base.Module, l0 int64)

//go:linkname Fn605 github.com/goccy/llamawasm2go/p2.Fn605
func Fn605(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn608 github.com/goccy/llamawasm2go/p3.Fn608
func Fn608(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn610 github.com/goccy/llamawasm2go/p2.Fn610
func Fn610(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn612 github.com/goccy/llamawasm2go/p2.Fn612
func Fn612(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn614 github.com/goccy/llamawasm2go/p3.Fn614
func Fn614(m *base.Module)

//go:linkname Fn615 github.com/goccy/llamawasm2go/p3.Fn615
func Fn615(m *base.Module, l0 int64) int64

//go:linkname Fn616 github.com/goccy/llamawasm2go/p3.Fn616
func Fn616(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn620 github.com/goccy/llamawasm2go/p2.Fn620
func Fn620(m *base.Module, l0 int64) int32

//go:linkname Fn621 github.com/goccy/llamawasm2go/p3.Fn621
func Fn621(m *base.Module, l0 int64) int64

//go:linkname Fn622 github.com/goccy/llamawasm2go/p3.Fn622
func Fn622(m *base.Module, l0 int64) int64

//go:linkname Fn623 github.com/goccy/llamawasm2go/p3.Fn623
func Fn623(m *base.Module, l0 int64, l1 int64) int64

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

//go:linkname Fn636 github.com/goccy/llamawasm2go/p3.Fn636
func Fn636(m *base.Module, l0 int64)

//go:linkname Fn637 github.com/goccy/llamawasm2go/p3.Fn637
func Fn637(m *base.Module, l0 int64)

//go:linkname Fn638 github.com/goccy/llamawasm2go/p3.Fn638
func Fn638(m *base.Module, l0 int64) int64

//go:linkname Fn640 github.com/goccy/llamawasm2go/p2.Fn640
func Fn640(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn641 github.com/goccy/llamawasm2go/p3.Fn641
func Fn641(m *base.Module, l0 int64)

//go:linkname Fn642 github.com/goccy/llamawasm2go/p2.Fn642
func Fn642(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn643 github.com/goccy/llamawasm2go/p2.Fn643
func Fn643(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn644 github.com/goccy/llamawasm2go/p2.Fn644
func Fn644(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn645 github.com/goccy/llamawasm2go/p2.Fn645
func Fn645(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn646 github.com/goccy/llamawasm2go/p2.Fn646
func Fn646(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn647 github.com/goccy/llamawasm2go/p2.Fn647
func Fn647(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn649 github.com/goccy/llamawasm2go/p3.Fn649
func Fn649(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn651 github.com/goccy/llamawasm2go/p3.Fn651
func Fn651(m *base.Module, l0 int64) int64

//go:linkname Fn653 github.com/goccy/llamawasm2go/p2.Fn653
func Fn653(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn654 github.com/goccy/llamawasm2go/p3.Fn654
func Fn654(m *base.Module, l0 int64)

//go:linkname Fn655 github.com/goccy/llamawasm2go/p3.Fn655
func Fn655(m *base.Module, l0 int64) int64

//go:linkname Fn656 github.com/goccy/llamawasm2go/p3.Fn656
func Fn656(m *base.Module, l0 int64) int64

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

//go:linkname Fn666 github.com/goccy/llamawasm2go/p3.Fn666
func Fn666(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn674 github.com/goccy/llamawasm2go/p3.Fn674
func Fn674(m *base.Module, l0 int64)

//go:linkname Fn675 github.com/goccy/llamawasm2go/p2.Fn675
func Fn675(m *base.Module, l0 int64)

//go:linkname Fn676 github.com/goccy/llamawasm2go/p3.Fn676
func Fn676(m *base.Module, l0 int64)

//go:linkname Fn677 github.com/goccy/llamawasm2go/p2.Fn677
func Fn677(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn679 github.com/goccy/llamawasm2go/p3.Fn679
func Fn679(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn680 github.com/goccy/llamawasm2go/p3.Fn680
func Fn680(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn681 github.com/goccy/llamawasm2go/p3.Fn681
func Fn681(m *base.Module, l0 int64) int32

//go:linkname Fn682 github.com/goccy/llamawasm2go/p2.Fn682
func Fn682(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn683 github.com/goccy/llamawasm2go/p2.Fn683
func Fn683(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn684 github.com/goccy/llamawasm2go/p2.Fn684
func Fn684(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn685 github.com/goccy/llamawasm2go/p2.Fn685
func Fn685(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn692 github.com/goccy/llamawasm2go/p2.Fn692
func Fn692(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn706 github.com/goccy/llamawasm2go/p2.Fn706
func Fn706(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn725 github.com/goccy/llamawasm2go/p2.Fn725
func Fn725(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn726 github.com/goccy/llamawasm2go/p3.Fn726
func Fn726(m *base.Module, l0 int64)

//go:linkname Fn728 github.com/goccy/llamawasm2go/p3.Fn728
func Fn728(m *base.Module, l0 int64)

//go:linkname Fn729 github.com/goccy/llamawasm2go/p2.Fn729
func Fn729(m *base.Module, l0 int64) int64

//go:linkname Fn744 github.com/goccy/llamawasm2go/p0.Fn744
func Fn744(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn745 github.com/goccy/llamawasm2go/p2.Fn745
func Fn745(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn762 github.com/goccy/llamawasm2go/p3.Fn762
func Fn762(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn770 github.com/goccy/llamawasm2go/p3.Fn770
func Fn770(m *base.Module, l0 int64)

//go:linkname Fn772 github.com/goccy/llamawasm2go/p3.Fn772
func Fn772(m *base.Module, l0 int64)

//go:linkname Fn778 github.com/goccy/llamawasm2go/p2.Fn778
func Fn778(m *base.Module, l0 int32)

//go:linkname Fn784 github.com/goccy/llamawasm2go/p3.Fn784
func Fn784(m *base.Module, l0 int64)

//go:linkname Fn785 github.com/goccy/llamawasm2go/p3.Fn785
func Fn785(m *base.Module, l0 int64)

//go:linkname Fn787 github.com/goccy/llamawasm2go/p3.Fn787
func Fn787(m *base.Module, l0 int64) int64

//go:linkname Fn790 github.com/goccy/llamawasm2go/p3.Fn790
func Fn790(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn791 github.com/goccy/llamawasm2go/p3.Fn791
func Fn791(m *base.Module, l0 int64)

//go:linkname Fn792 github.com/goccy/llamawasm2go/p3.Fn792
func Fn792(m *base.Module, l0 int64)

//go:linkname Fn838 github.com/goccy/llamawasm2go/p2.Fn838
func Fn838(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 float32, l5 int64)

//go:linkname Fn860 github.com/goccy/llamawasm2go/p2.Fn860
func Fn860(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn863 github.com/goccy/llamawasm2go/p3.Fn863
func Fn863(m *base.Module, l0 int64) int64

//go:linkname Fn905 github.com/goccy/llamawasm2go/p2.Fn905
func Fn905(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn977 github.com/goccy/llamawasm2go/p3.Fn977
func Fn977(m *base.Module) int64

//go:linkname Fn1017 github.com/goccy/llamawasm2go/p2.Fn1017
func Fn1017(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32)

//go:linkname Fn1027 github.com/goccy/llamawasm2go/p2.Fn1027
func Fn1027(m *base.Module, l0 int64)

//go:linkname Fn1034 github.com/goccy/llamawasm2go/p2.Fn1034
func Fn1034(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1036 github.com/goccy/llamawasm2go/p2.Fn1036
func Fn1036(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1050 github.com/goccy/llamawasm2go/p2.Fn1050
func Fn1050(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1051 github.com/goccy/llamawasm2go/p2.Fn1051
func Fn1051(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1052 github.com/goccy/llamawasm2go/p2.Fn1052
func Fn1052(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn1054 github.com/goccy/llamawasm2go/p2.Fn1054
func Fn1054(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32)

//go:linkname Fn1067 github.com/goccy/llamawasm2go/p3.Fn1067
func Fn1067(m *base.Module, l0 int64)

//go:linkname Fn1072 github.com/goccy/llamawasm2go/p2.Fn1072
func Fn1072(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1073 github.com/goccy/llamawasm2go/p3.Fn1073
func Fn1073(m *base.Module)

//go:linkname Fn1082 github.com/goccy/llamawasm2go/p2.Fn1082
func Fn1082(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1099 github.com/goccy/llamawasm2go/p2.Fn1099
func Fn1099(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn1102 github.com/goccy/llamawasm2go/p3.Fn1102
func Fn1102(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1140 github.com/goccy/llamawasm2go/p2.Fn1140
func Fn1140(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1147 github.com/goccy/llamawasm2go/p2.Fn1147
func Fn1147(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1149 github.com/goccy/llamawasm2go/p2.Fn1149
func Fn1149(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn1155 github.com/goccy/llamawasm2go/p2.Fn1155
func Fn1155(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1188 github.com/goccy/llamawasm2go/p0.Fn1188
func Fn1188(m *base.Module, l0 int64) int64

//go:linkname Fn1189 github.com/goccy/llamawasm2go/p3.Fn1189
func Fn1189(m *base.Module, l0 int64, l1 int64) int64

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

//go:linkname Fn1320 github.com/goccy/llamawasm2go/p3.Fn1320
func Fn1320(m *base.Module)

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p3.Fn1324
func Fn1324(m *base.Module, l0 int64) int64

//go:linkname Fn1353 github.com/goccy/llamawasm2go/p3.Fn1353
func Fn1353(m *base.Module) int64

//go:linkname Fn1354 github.com/goccy/llamawasm2go/p3.Fn1354
func Fn1354(m *base.Module) int64

//go:linkname Fn1355 github.com/goccy/llamawasm2go/p2.Fn1355
func Fn1355(m *base.Module, l0 int64) int64

//go:linkname Fn1356 github.com/goccy/llamawasm2go/p3.Fn1356
func Fn1356(m *base.Module, l0 int32) int64

//go:linkname Fn1366 github.com/goccy/llamawasm2go/p2.Fn1366
func Fn1366(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1367 github.com/goccy/llamawasm2go/p3.Fn1367
func Fn1367(m *base.Module, l0 int64)

//go:linkname Fn1369 github.com/goccy/llamawasm2go/p2.Fn1369
func Fn1369(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1370 github.com/goccy/llamawasm2go/p2.Fn1370
func Fn1370(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1385 github.com/goccy/llamawasm2go/p3.Fn1385
func Fn1385(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1394 github.com/goccy/llamawasm2go/p3.Fn1394
func Fn1394(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p3.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1396 github.com/goccy/llamawasm2go/p3.Fn1396
func Fn1396(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1398 github.com/goccy/llamawasm2go/p2.Fn1398
func Fn1398(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1406 github.com/goccy/llamawasm2go/p2.Fn1406
func Fn1406(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p2.Fn1407
func Fn1407(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1408 github.com/goccy/llamawasm2go/p3.Fn1408
func Fn1408(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1409 github.com/goccy/llamawasm2go/p2.Fn1409
func Fn1409(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1410 github.com/goccy/llamawasm2go/p2.Fn1410
func Fn1410(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1411 github.com/goccy/llamawasm2go/p3.Fn1411
func Fn1411(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1412 github.com/goccy/llamawasm2go/p3.Fn1412
func Fn1412(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p3.Fn1413
func Fn1413(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1414 github.com/goccy/llamawasm2go/p3.Fn1414
func Fn1414(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1415 github.com/goccy/llamawasm2go/p2.Fn1415
func Fn1415(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1416 github.com/goccy/llamawasm2go/p3.Fn1416
func Fn1416(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1417 github.com/goccy/llamawasm2go/p3.Fn1417
func Fn1417(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1418 github.com/goccy/llamawasm2go/p3.Fn1418
func Fn1418(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1419 github.com/goccy/llamawasm2go/p3.Fn1419
func Fn1419(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1421 github.com/goccy/llamawasm2go/p2.Fn1421
func Fn1421(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1422 github.com/goccy/llamawasm2go/p3.Fn1422
func Fn1422(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1423 github.com/goccy/llamawasm2go/p2.Fn1423
func Fn1423(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1425 github.com/goccy/llamawasm2go/p2.Fn1425
func Fn1425(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p2.Fn1426
func Fn1426(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1427 github.com/goccy/llamawasm2go/p3.Fn1427
func Fn1427(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1428 github.com/goccy/llamawasm2go/p3.Fn1428
func Fn1428(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1429 github.com/goccy/llamawasm2go/p3.Fn1429
func Fn1429(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1430 github.com/goccy/llamawasm2go/p3.Fn1430
func Fn1430(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1431 github.com/goccy/llamawasm2go/p2.Fn1431
func Fn1431(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1432 github.com/goccy/llamawasm2go/p2.Fn1432
func Fn1432(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1435 github.com/goccy/llamawasm2go/p2.Fn1435
func Fn1435(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1437 github.com/goccy/llamawasm2go/p2.Fn1437
func Fn1437(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1442 github.com/goccy/llamawasm2go/p2.Fn1442
func Fn1442(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1443 github.com/goccy/llamawasm2go/p2.Fn1443
func Fn1443(m *base.Module, l0 int64)

//go:linkname Fn1450 github.com/goccy/llamawasm2go/p2.Fn1450
func Fn1450(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1454 github.com/goccy/llamawasm2go/p2.Fn1454
func Fn1454(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1460 github.com/goccy/llamawasm2go/p2.Fn1460
func Fn1460(m *base.Module, l0 int64)

//go:linkname Fn1461 github.com/goccy/llamawasm2go/p3.Fn1461
func Fn1461(m *base.Module, l0 int64)

//go:linkname Fn1466 github.com/goccy/llamawasm2go/p2.Fn1466
func Fn1466(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1476 github.com/goccy/llamawasm2go/p2.Fn1476
func Fn1476(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1478 github.com/goccy/llamawasm2go/p2.Fn1478
func Fn1478(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1484 github.com/goccy/llamawasm2go/p2.Fn1484
func Fn1484(m *base.Module, l0 int64, l1 int32, l2 int32) int64

//go:linkname Fn1485 github.com/goccy/llamawasm2go/p2.Fn1485
func Fn1485(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1486 github.com/goccy/llamawasm2go/p2.Fn1486
func Fn1486(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1487 github.com/goccy/llamawasm2go/p3.Fn1487
func Fn1487(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1488 github.com/goccy/llamawasm2go/p0.Fn1488
func Fn1488(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1490 github.com/goccy/llamawasm2go/p3.Fn1490
func Fn1490(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1491 github.com/goccy/llamawasm2go/p2.Fn1491
func Fn1491(m *base.Module, l0 int64)

//go:linkname Fn1492 github.com/goccy/llamawasm2go/p2.Fn1492
func Fn1492(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1494 github.com/goccy/llamawasm2go/p2.Fn1494
func Fn1494(m *base.Module, l0 int64) int64

//go:linkname Fn1495 github.com/goccy/llamawasm2go/p2.Fn1495
func Fn1495(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1496 github.com/goccy/llamawasm2go/p3.Fn1496
func Fn1496(m *base.Module, l0 int64)

//go:linkname Fn1497 github.com/goccy/llamawasm2go/p3.Fn1497
func Fn1497(m *base.Module, l0 int64)

//go:linkname Fn1498 github.com/goccy/llamawasm2go/p2.Fn1498
func Fn1498(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1499 github.com/goccy/llamawasm2go/p2.Fn1499
func Fn1499(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1502 github.com/goccy/llamawasm2go/p2.Fn1502
func Fn1502(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1506 github.com/goccy/llamawasm2go/p2.Fn1506
func Fn1506(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1507 github.com/goccy/llamawasm2go/p2.Fn1507
func Fn1507(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1508 github.com/goccy/llamawasm2go/p2.Fn1508
func Fn1508(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1509 github.com/goccy/llamawasm2go/p2.Fn1509
func Fn1509(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1519 github.com/goccy/llamawasm2go/p0.Fn1519
func Fn1519(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1521 github.com/goccy/llamawasm2go/p3.Fn1521
func Fn1521(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1522 github.com/goccy/llamawasm2go/p2.Fn1522
func Fn1522(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1524 github.com/goccy/llamawasm2go/p3.Fn1524
func Fn1524(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1526 github.com/goccy/llamawasm2go/p3.Fn1526
func Fn1526(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1527 github.com/goccy/llamawasm2go/p2.Fn1527
func Fn1527(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1529 github.com/goccy/llamawasm2go/p2.Fn1529
func Fn1529(m *base.Module, l0 int64)

//go:linkname Fn1533 github.com/goccy/llamawasm2go/p2.Fn1533
func Fn1533(m *base.Module, l0 int64)

//go:linkname Fn1537 github.com/goccy/llamawasm2go/p0.Fn1537
func Fn1537(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int64) int64

//go:linkname Fn1539 github.com/goccy/llamawasm2go/p2.Fn1539
func Fn1539(m *base.Module, l0 int64) int64

//go:linkname Fn1540 github.com/goccy/llamawasm2go/p3.Fn1540
func Fn1540(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p3.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1542 github.com/goccy/llamawasm2go/p2.Fn1542
func Fn1542(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1543 github.com/goccy/llamawasm2go/p2.Fn1543
func Fn1543(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1544 github.com/goccy/llamawasm2go/p2.Fn1544
func Fn1544(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1547 github.com/goccy/llamawasm2go/p3.Fn1547
func Fn1547(m *base.Module, l0 int64)

//go:linkname Fn1551 github.com/goccy/llamawasm2go/p2.Fn1551
func Fn1551(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1555 github.com/goccy/llamawasm2go/p2.Fn1555
func Fn1555(m *base.Module, l0 int64, l1 int64, l2 int32) int32

//go:linkname Fn1556 github.com/goccy/llamawasm2go/p2.Fn1556
func Fn1556(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p2.Fn1563
func Fn1563(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1564 github.com/goccy/llamawasm2go/p2.Fn1564
func Fn1564(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1571 github.com/goccy/llamawasm2go/p3.Fn1571
func Fn1571(m *base.Module, l0 int64)

//go:linkname Fn1572 github.com/goccy/llamawasm2go/p3.Fn1572
func Fn1572(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1573 github.com/goccy/llamawasm2go/p2.Fn1573
func Fn1573(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1576 github.com/goccy/llamawasm2go/p3.Fn1576
func Fn1576(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1578 github.com/goccy/llamawasm2go/p3.Fn1578
func Fn1578(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1579 github.com/goccy/llamawasm2go/p3.Fn1579
func Fn1579(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1580 github.com/goccy/llamawasm2go/p3.Fn1580
func Fn1580(m *base.Module, l0 int64)

//go:linkname Fn1581 github.com/goccy/llamawasm2go/p3.Fn1581
func Fn1581(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1583 github.com/goccy/llamawasm2go/p3.Fn1583
func Fn1583(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1584 github.com/goccy/llamawasm2go/p3.Fn1584
func Fn1584(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1633 github.com/goccy/llamawasm2go/p2.Fn1633
func Fn1633(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1646 github.com/goccy/llamawasm2go/p2.Fn1646
func Fn1646(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1649 github.com/goccy/llamawasm2go/p3.Fn1649
func Fn1649(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1650 github.com/goccy/llamawasm2go/p3.Fn1650
func Fn1650(m *base.Module, l0 int64, l1 int64)

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

//go:linkname Fn1657 github.com/goccy/llamawasm2go/p3.Fn1657
func Fn1657(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1659 github.com/goccy/llamawasm2go/p2.Fn1659
func Fn1659(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1660 github.com/goccy/llamawasm2go/p2.Fn1660
func Fn1660(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn1661 github.com/goccy/llamawasm2go/p2.Fn1661
func Fn1661(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1662 github.com/goccy/llamawasm2go/p3.Fn1662
func Fn1662(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32)

//go:linkname Fn1663 github.com/goccy/llamawasm2go/p2.Fn1663
func Fn1663(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int32, l11 int32)

//go:linkname Fn1665 github.com/goccy/llamawasm2go/p3.Fn1665
func Fn1665(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64, l20 int64, l21 float32) int64

//go:linkname Fn1667 github.com/goccy/llamawasm2go/p0.Fn1667
func Fn1667(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1669 github.com/goccy/llamawasm2go/p2.Fn1669
func Fn1669(m *base.Module, l0 int64) int64

//go:linkname Fn1670 github.com/goccy/llamawasm2go/p2.Fn1670
func Fn1670(m *base.Module, l0 int64) int64

//go:linkname Fn1671 github.com/goccy/llamawasm2go/p2.Fn1671
func Fn1671(m *base.Module, l0 int64) int64

//go:linkname Fn1672 github.com/goccy/llamawasm2go/p2.Fn1672
func Fn1672(m *base.Module, l0 int64) int64

//go:linkname Fn1673 github.com/goccy/llamawasm2go/p2.Fn1673
func Fn1673(m *base.Module, l0 int64) int64

//go:linkname Fn1680 github.com/goccy/llamawasm2go/p2.Fn1680
func Fn1680(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1681 github.com/goccy/llamawasm2go/p2.Fn1681
func Fn1681(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1682 github.com/goccy/llamawasm2go/p2.Fn1682
func Fn1682(m *base.Module, l0 int64) int64

//go:linkname Fn1684 github.com/goccy/llamawasm2go/p2.Fn1684
func Fn1684(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 float32, l11 int32) int64

//go:linkname Fn1685 github.com/goccy/llamawasm2go/p2.Fn1685
func Fn1685(m *base.Module, l0 int64) int64

//go:linkname Fn1688 github.com/goccy/llamawasm2go/p2.Fn1688
func Fn1688(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1689 github.com/goccy/llamawasm2go/p2.Fn1689
func Fn1689(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 float32, l11 int32) int64

//go:linkname Fn1690 github.com/goccy/llamawasm2go/p2.Fn1690
func Fn1690(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1692 github.com/goccy/llamawasm2go/p2.Fn1692
func Fn1692(m *base.Module, l0 int64) int64

//go:linkname Fn1694 github.com/goccy/llamawasm2go/p2.Fn1694
func Fn1694(m *base.Module, l0 int64) int64

//go:linkname Fn1695 github.com/goccy/llamawasm2go/p2.Fn1695
func Fn1695(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1696 github.com/goccy/llamawasm2go/p2.Fn1696
func Fn1696(m *base.Module, l0 int64) int64

//go:linkname Fn1698 github.com/goccy/llamawasm2go/p2.Fn1698
func Fn1698(m *base.Module, l0 int64) int64

//go:linkname Fn1705 github.com/goccy/llamawasm2go/p3.Fn1705
func Fn1705(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1709 github.com/goccy/llamawasm2go/p2.Fn1709
func Fn1709(m *base.Module, l0 int64) int64

//go:linkname Fn1710 github.com/goccy/llamawasm2go/p2.Fn1710
func Fn1710(m *base.Module, l0 int64) int64

//go:linkname Fn1712 github.com/goccy/llamawasm2go/p2.Fn1712
func Fn1712(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1714 github.com/goccy/llamawasm2go/p3.Fn1714
func Fn1714(m *base.Module, l0 int64, l1 int64)

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

//go:linkname Fn1771 github.com/goccy/llamawasm2go/p3.Fn1771
func Fn1771(m *base.Module, l0 int64) int32

//go:linkname Fn1772 github.com/goccy/llamawasm2go/p3.Fn1772
func Fn1772(m *base.Module, l0 int64) int32

//go:linkname Fn1773 github.com/goccy/llamawasm2go/p3.Fn1773
func Fn1773(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1775 github.com/goccy/llamawasm2go/p3.Fn1775
func Fn1775(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1778 github.com/goccy/llamawasm2go/p3.Fn1778
func Fn1778(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1779 github.com/goccy/llamawasm2go/p3.Fn1779
func Fn1779(m *base.Module, l0 int64) int32

//go:linkname Fn1780 github.com/goccy/llamawasm2go/p3.Fn1780
func Fn1780(m *base.Module, l0 int64) int32

//go:linkname Fn1782 github.com/goccy/llamawasm2go/p3.Fn1782
func Fn1782(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1784 github.com/goccy/llamawasm2go/p2.Fn1784
func Fn1784(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1786 github.com/goccy/llamawasm2go/p3.Fn1786
func Fn1786(m *base.Module, l0 int64)

//go:linkname Fn1788 github.com/goccy/llamawasm2go/p3.Fn1788
func Fn1788(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1789 github.com/goccy/llamawasm2go/p2.Fn1789
func Fn1789(m *base.Module, l0 int64, l1 int64, l2 int64)

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

//go:linkname Fn1803 github.com/goccy/llamawasm2go/p3.Fn1803
func Fn1803(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1804 github.com/goccy/llamawasm2go/p2.Fn1804
func Fn1804(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1806 github.com/goccy/llamawasm2go/p2.Fn1806
func Fn1806(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1810 github.com/goccy/llamawasm2go/p3.Fn1810
func Fn1810(m *base.Module, l0 int64)

//go:linkname Fn1820 github.com/goccy/llamawasm2go/p2.Fn1820
func Fn1820(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1830 github.com/goccy/llamawasm2go/p2.Fn1830
func Fn1830(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1831 github.com/goccy/llamawasm2go/p0.Fn1831
func Fn1831(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1833 github.com/goccy/llamawasm2go/p3.Fn1833
func Fn1833(m *base.Module, l0 int64)

//go:linkname Fn1835 github.com/goccy/llamawasm2go/p2.Fn1835
func Fn1835(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1836 github.com/goccy/llamawasm2go/p2.Fn1836
func Fn1836(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1843 github.com/goccy/llamawasm2go/p2.Fn1843
func Fn1843(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1844 github.com/goccy/llamawasm2go/p2.Fn1844
func Fn1844(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1849 github.com/goccy/llamawasm2go/p3.Fn1849
func Fn1849(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1855 github.com/goccy/llamawasm2go/p2.Fn1855
func Fn1855(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1859 github.com/goccy/llamawasm2go/p2.Fn1859
func Fn1859(m *base.Module, l0 int64)

//go:linkname Fn1863 github.com/goccy/llamawasm2go/p2.Fn1863
func Fn1863(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1864 github.com/goccy/llamawasm2go/p2.Fn1864
func Fn1864(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64)

//go:linkname Fn1866 github.com/goccy/llamawasm2go/p2.Fn1866
func Fn1866(m *base.Module, l0 int64)

//go:linkname Fn1877 github.com/goccy/llamawasm2go/p3.Fn1877
func Fn1877(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1878 github.com/goccy/llamawasm2go/p2.Fn1878
func Fn1878(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1879 github.com/goccy/llamawasm2go/p3.Fn1879
func Fn1879(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1880 github.com/goccy/llamawasm2go/p2.Fn1880
func Fn1880(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1898 github.com/goccy/llamawasm2go/p2.Fn1898
func Fn1898(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1918 github.com/goccy/llamawasm2go/p3.Fn1918
func Fn1918(m *base.Module, l0 int64)

//go:linkname Fn1922 github.com/goccy/llamawasm2go/p3.Fn1922
func Fn1922(m *base.Module, l0 int64) int32

//go:linkname Fn1989 github.com/goccy/llamawasm2go/p3.Fn1989
func Fn1989(m *base.Module, l0 int64) int32

//go:linkname Fn1990 github.com/goccy/llamawasm2go/p2.Fn1990
func Fn1990(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1994 github.com/goccy/llamawasm2go/p3.Fn1994
func Fn1994(m *base.Module, l0 int64)

//go:linkname Fn1998 github.com/goccy/llamawasm2go/p2.Fn1998
func Fn1998(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2012 github.com/goccy/llamawasm2go/p2.Fn2012
func Fn2012(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn2013 github.com/goccy/llamawasm2go/p2.Fn2013
func Fn2013(m *base.Module, l0 int64) int64

//go:linkname Fn2014 github.com/goccy/llamawasm2go/p3.Fn2014
func Fn2014(m *base.Module, l0 int64)

//go:linkname Fn2019 github.com/goccy/llamawasm2go/p2.Fn2019
func Fn2019(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2020 github.com/goccy/llamawasm2go/p2.Fn2020
func Fn2020(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2021 github.com/goccy/llamawasm2go/p2.Fn2021
func Fn2021(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2039 github.com/goccy/llamawasm2go/p2.Fn2039
func Fn2039(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2041 github.com/goccy/llamawasm2go/p2.Fn2041
func Fn2041(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64) int64

//go:linkname Fn2053 github.com/goccy/llamawasm2go/p2.Fn2053
func Fn2053(m *base.Module, l0 int64) int64

//go:linkname Fn2054 github.com/goccy/llamawasm2go/p3.Fn2054
func Fn2054(m *base.Module, l0 int64)

//go:linkname Fn2057 github.com/goccy/llamawasm2go/p0.Fn2057
func Fn2057(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int64)

//go:linkname Fn2058 github.com/goccy/llamawasm2go/p2.Fn2058
func Fn2058(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2059 github.com/goccy/llamawasm2go/p2.Fn2059
func Fn2059(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2084 github.com/goccy/llamawasm2go/p2.Fn2084
func Fn2084(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2101 github.com/goccy/llamawasm2go/p2.Fn2101
func Fn2101(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2114 github.com/goccy/llamawasm2go/p3.Fn2114
func Fn2114(m *base.Module, l0 int64) int64

//go:linkname Fn2117 github.com/goccy/llamawasm2go/p2.Fn2117
func Fn2117(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2125 github.com/goccy/llamawasm2go/p3.Fn2125
func Fn2125(m *base.Module, l0 int64)

//go:linkname Fn2126 github.com/goccy/llamawasm2go/p3.Fn2126
func Fn2126(m *base.Module, l0 int64)

//go:linkname Fn2137 github.com/goccy/llamawasm2go/p2.Fn2137
func Fn2137(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2143 github.com/goccy/llamawasm2go/p2.Fn2143
func Fn2143(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32

//go:linkname Fn2166 github.com/goccy/llamawasm2go/p3.Fn2166
func Fn2166(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2167 github.com/goccy/llamawasm2go/p2.Fn2167
func Fn2167(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2168 github.com/goccy/llamawasm2go/p2.Fn2168
func Fn2168(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2174 github.com/goccy/llamawasm2go/p3.Fn2174
func Fn2174(m *base.Module)

//go:linkname Fn2177 github.com/goccy/llamawasm2go/p3.Fn2177
func Fn2177(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2182 github.com/goccy/llamawasm2go/p2.Fn2182
func Fn2182(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn2183 github.com/goccy/llamawasm2go/p2.Fn2183
func Fn2183(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2184 github.com/goccy/llamawasm2go/p2.Fn2184
func Fn2184(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2185 github.com/goccy/llamawasm2go/p3.Fn2185
func Fn2185(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2187 github.com/goccy/llamawasm2go/p2.Fn2187
func Fn2187(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2188 github.com/goccy/llamawasm2go/p2.Fn2188
func Fn2188(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2189 github.com/goccy/llamawasm2go/p2.Fn2189
func Fn2189(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2191 github.com/goccy/llamawasm2go/p3.Fn2191
func Fn2191(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2193 github.com/goccy/llamawasm2go/p3.Fn2193
func Fn2193(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2195 github.com/goccy/llamawasm2go/p2.Fn2195
func Fn2195(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn2197 github.com/goccy/llamawasm2go/p3.Fn2197
func Fn2197(m *base.Module, l0 int64, l1 int32, l2 int64)

//go:linkname Fn2198 github.com/goccy/llamawasm2go/p2.Fn2198
func Fn2198(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn2200 github.com/goccy/llamawasm2go/p2.Fn2200
func Fn2200(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2201 github.com/goccy/llamawasm2go/p2.Fn2201
func Fn2201(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32, l4 int32) int32

//go:linkname Fn2202 github.com/goccy/llamawasm2go/p0.Fn2202
func Fn2202(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int64, l8 int64) int64

//go:linkname Fn2206 github.com/goccy/llamawasm2go/p2.Fn2206
func Fn2206(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2208 github.com/goccy/llamawasm2go/p3.Fn2208
func Fn2208(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2211 github.com/goccy/llamawasm2go/p2.Fn2211
func Fn2211(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2212 github.com/goccy/llamawasm2go/p3.Fn2212
func Fn2212(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2213 github.com/goccy/llamawasm2go/p2.Fn2213
func Fn2213(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn2214 github.com/goccy/llamawasm2go/p2.Fn2214
func Fn2214(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64) int64

//go:linkname Fn2216 github.com/goccy/llamawasm2go/p3.Fn2216
func Fn2216(m *base.Module, l0 int64)

//go:linkname Fn2224 github.com/goccy/llamawasm2go/p2.Fn2224
func Fn2224(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2226 github.com/goccy/llamawasm2go/p3.Fn2226
func Fn2226(m *base.Module, l0 int64)

//go:linkname Fn2227 github.com/goccy/llamawasm2go/p3.Fn2227
func Fn2227(m *base.Module, l0 int64)

//go:linkname Fn2228 github.com/goccy/llamawasm2go/p3.Fn2228
func Fn2228(m *base.Module, l0 int64)

//go:linkname Fn2229 github.com/goccy/llamawasm2go/p2.Fn2229
func Fn2229(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2231 github.com/goccy/llamawasm2go/p3.Fn2231
func Fn2231(m *base.Module, l0 int64)

//go:linkname Fn2233 github.com/goccy/llamawasm2go/p2.Fn2233
func Fn2233(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2234 github.com/goccy/llamawasm2go/p2.Fn2234
func Fn2234(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2247 github.com/goccy/llamawasm2go/p2.Fn2247
func Fn2247(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2248 github.com/goccy/llamawasm2go/p3.Fn2248
func Fn2248(m *base.Module, l0 int64) int64

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p3.Fn2249
func Fn2249(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2250 github.com/goccy/llamawasm2go/p2.Fn2250
func Fn2250(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2251 github.com/goccy/llamawasm2go/p3.Fn2251
func Fn2251(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2263 github.com/goccy/llamawasm2go/p2.Fn2263
func Fn2263(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2265 github.com/goccy/llamawasm2go/p2.Fn2265
func Fn2265(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2266 github.com/goccy/llamawasm2go/p2.Fn2266
func Fn2266(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2268 github.com/goccy/llamawasm2go/p2.Fn2268
func Fn2268(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2269 github.com/goccy/llamawasm2go/p3.Fn2269
func Fn2269(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2270 github.com/goccy/llamawasm2go/p2.Fn2270
func Fn2270(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2271 github.com/goccy/llamawasm2go/p2.Fn2271
func Fn2271(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2292 github.com/goccy/llamawasm2go/p2.Fn2292
func Fn2292(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2293 github.com/goccy/llamawasm2go/p2.Fn2293
func Fn2293(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2294 github.com/goccy/llamawasm2go/p2.Fn2294
func Fn2294(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2304 github.com/goccy/llamawasm2go/p2.Fn2304
func Fn2304(m *base.Module, l0 int64)

//go:linkname Fn2321 github.com/goccy/llamawasm2go/p2.Fn2321
func Fn2321(m *base.Module, l0 int64)

//go:linkname Fn2322 github.com/goccy/llamawasm2go/p2.Fn2322
func Fn2322(m *base.Module, l0 int64)

//go:linkname Fn2323 github.com/goccy/llamawasm2go/p2.Fn2323
func Fn2323(m *base.Module, l0 int64)

//go:linkname Fn2325 github.com/goccy/llamawasm2go/p2.Fn2325
func Fn2325(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2326 github.com/goccy/llamawasm2go/p2.Fn2326
func Fn2326(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2327 github.com/goccy/llamawasm2go/p2.Fn2327
func Fn2327(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2356 github.com/goccy/llamawasm2go/p3.Fn2356
func Fn2356(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2357 github.com/goccy/llamawasm2go/p3.Fn2357
func Fn2357(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2358 github.com/goccy/llamawasm2go/p3.Fn2358
func Fn2358(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2359 github.com/goccy/llamawasm2go/p3.Fn2359
func Fn2359(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64)

//go:linkname Fn2404 github.com/goccy/llamawasm2go/p3.Fn2404
func Fn2404(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2407 github.com/goccy/llamawasm2go/p2.Fn2407
func Fn2407(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn2408 github.com/goccy/llamawasm2go/p3.Fn2408
func Fn2408(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2409 github.com/goccy/llamawasm2go/p3.Fn2409
func Fn2409(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2411 github.com/goccy/llamawasm2go/p0.Fn2411
func Fn2411(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2413 github.com/goccy/llamawasm2go/p3.Fn2413
func Fn2413(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2415 github.com/goccy/llamawasm2go/p2.Fn2415
func Fn2415(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2416 github.com/goccy/llamawasm2go/p2.Fn2416
func Fn2416(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2421 github.com/goccy/llamawasm2go/p2.Fn2421
func Fn2421(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2469 github.com/goccy/llamawasm2go/p2.Fn2469
func Fn2469(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2470 github.com/goccy/llamawasm2go/p2.Fn2470
func Fn2470(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2474 github.com/goccy/llamawasm2go/p2.Fn2474
func Fn2474(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2475 github.com/goccy/llamawasm2go/p2.Fn2475
func Fn2475(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn2476 github.com/goccy/llamawasm2go/p2.Fn2476
func Fn2476(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn2477 github.com/goccy/llamawasm2go/p2.Fn2477
func Fn2477(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2479 github.com/goccy/llamawasm2go/p2.Fn2479
func Fn2479(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2480 github.com/goccy/llamawasm2go/p2.Fn2480
func Fn2480(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2573 github.com/goccy/llamawasm2go/p2.Fn2573
func Fn2573(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64, l15 int32, l16 int32) int64

//go:linkname Fn2591 github.com/goccy/llamawasm2go/p2.Fn2591
func Fn2591(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32)

//go:linkname Fn2592 github.com/goccy/llamawasm2go/p2.Fn2592
func Fn2592(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32)

//go:linkname Fn2594 github.com/goccy/llamawasm2go/p3.Fn2594
func Fn2594(m *base.Module, l0 int64) int64

//go:linkname Fn2595 github.com/goccy/llamawasm2go/p2.Fn2595
func Fn2595(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2602 github.com/goccy/llamawasm2go/p2.Fn2602
func Fn2602(m *base.Module, l0 int64)

//go:linkname Fn2620 github.com/goccy/llamawasm2go/p2.Fn2620
func Fn2620(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2621 github.com/goccy/llamawasm2go/p3.Fn2621
func Fn2621(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2622 github.com/goccy/llamawasm2go/p2.Fn2622
func Fn2622(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2644 github.com/goccy/llamawasm2go/p2.Fn2644
func Fn2644(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn2645 github.com/goccy/llamawasm2go/p3.Fn2645
func Fn2645(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn2646 github.com/goccy/llamawasm2go/p2.Fn2646
func Fn2646(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn2650 github.com/goccy/llamawasm2go/p2.Fn2650
func Fn2650(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2683 github.com/goccy/llamawasm2go/p2.Fn2683
func Fn2683(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2685 github.com/goccy/llamawasm2go/p2.Fn2685
func Fn2685(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2686 github.com/goccy/llamawasm2go/p2.Fn2686
func Fn2686(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2687 github.com/goccy/llamawasm2go/p2.Fn2687
func Fn2687(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn2688 github.com/goccy/llamawasm2go/p3.Fn2688
func Fn2688(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn2749 github.com/goccy/llamawasm2go/p3.Fn2749
func Fn2749(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2832 github.com/goccy/llamawasm2go/p2.Fn2832
func Fn2832(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 float32, l7 int32) int64

//go:linkname Fn2927 github.com/goccy/llamawasm2go/p0.Fn2927
func Fn2927(m *base.Module, l0 int32, l1 int64) int64

//go:linkname Fn2970 github.com/goccy/llamawasm2go/p2.Fn2970
func Fn2970(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2979 github.com/goccy/llamawasm2go/p3.Fn2979
func Fn2979(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2981 github.com/goccy/llamawasm2go/p3.Fn2981
func Fn2981(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2983 github.com/goccy/llamawasm2go/p3.Fn2983
func Fn2983(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2985 github.com/goccy/llamawasm2go/p3.Fn2985
func Fn2985(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2989 github.com/goccy/llamawasm2go/p2.Fn2989
func Fn2989(m *base.Module, l0 int64)

//go:linkname Fn2990 github.com/goccy/llamawasm2go/p2.Fn2990
func Fn2990(m *base.Module, l0 int64) int64

//go:linkname Fn2995 github.com/goccy/llamawasm2go/p2.Fn2995
func Fn2995(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn3000 github.com/goccy/llamawasm2go/p3.Fn3000
func Fn3000(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3001 github.com/goccy/llamawasm2go/p2.Fn3001
func Fn3001(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3006 github.com/goccy/llamawasm2go/p3.Fn3006
func Fn3006(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn3008 github.com/goccy/llamawasm2go/p3.Fn3008
func Fn3008(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3010 github.com/goccy/llamawasm2go/p0.Fn3010
func Fn3010(m *base.Module, l0 int64)

//go:linkname Fn3019 github.com/goccy/llamawasm2go/p3.Fn3019
func Fn3019(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3021 github.com/goccy/llamawasm2go/p3.Fn3021
func Fn3021(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn3022 github.com/goccy/llamawasm2go/p3.Fn3022
func Fn3022(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn3043 github.com/goccy/llamawasm2go/p3.Fn3043
func Fn3043(m *base.Module, l0 int64)

//go:linkname Fn3044 github.com/goccy/llamawasm2go/p3.Fn3044
func Fn3044(m *base.Module, l0 int64) int64

//go:linkname Fn3046 github.com/goccy/llamawasm2go/p2.Fn3046
func Fn3046(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn3048 github.com/goccy/llamawasm2go/p2.Fn3048
func Fn3048(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn3051 github.com/goccy/llamawasm2go/p2.Fn3051
func Fn3051(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn3056 github.com/goccy/llamawasm2go/p2.Fn3056
func Fn3056(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3065 github.com/goccy/llamawasm2go/p2.Fn3065
func Fn3065(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3069 github.com/goccy/llamawasm2go/p2.Fn3069
func Fn3069(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3070 github.com/goccy/llamawasm2go/p3.Fn3070
func Fn3070(m *base.Module, l0 int64)

//go:linkname Fn3071 github.com/goccy/llamawasm2go/p2.Fn3071
func Fn3071(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn3073 github.com/goccy/llamawasm2go/p3.Fn3073
func Fn3073(m *base.Module, l0 int64)

//go:linkname Fn3074 github.com/goccy/llamawasm2go/p2.Fn3074
func Fn3074(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn3075 github.com/goccy/llamawasm2go/p3.Fn3075
func Fn3075(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3079 github.com/goccy/llamawasm2go/p2.Fn3079
func Fn3079(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn3082 github.com/goccy/llamawasm2go/p2.Fn3082
func Fn3082(m *base.Module, l0 int64) int64

//go:linkname Fn3083 github.com/goccy/llamawasm2go/p2.Fn3083
func Fn3083(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3084 github.com/goccy/llamawasm2go/p2.Fn3084
func Fn3084(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3090 github.com/goccy/llamawasm2go/p2.Fn3090
func Fn3090(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3091 github.com/goccy/llamawasm2go/p2.Fn3091
func Fn3091(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3093 github.com/goccy/llamawasm2go/p3.Fn3093
func Fn3093(m *base.Module, l0 int64)

//go:linkname Fn3094 github.com/goccy/llamawasm2go/p3.Fn3094
func Fn3094(m *base.Module) int64

//go:linkname Fn3096 github.com/goccy/llamawasm2go/p3.Fn3096
func Fn3096(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3097 github.com/goccy/llamawasm2go/p3.Fn3097
func Fn3097(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3099 github.com/goccy/llamawasm2go/p3.Fn3099
func Fn3099(m *base.Module) int64

//go:linkname Fn3101 github.com/goccy/llamawasm2go/p2.Fn3101
func Fn3101(m *base.Module, l0 int32) int64

//go:linkname Fn3102 github.com/goccy/llamawasm2go/p2.Fn3102
func Fn3102(m *base.Module, l0 int32) int32

//go:linkname Fn3104 github.com/goccy/llamawasm2go/p3.Fn3104
func Fn3104(m *base.Module, l0 int32) int64

//go:linkname Fn3105 github.com/goccy/llamawasm2go/p3.Fn3105
func Fn3105(m *base.Module, l0 float32) int64

//go:linkname Fn3106 github.com/goccy/llamawasm2go/p3.Fn3106
func Fn3106(m *base.Module, l0 float32) int64

//go:linkname Fn3107 github.com/goccy/llamawasm2go/p3.Fn3107
func Fn3107(m *base.Module, l0 float32) int64

//go:linkname Fn3109 github.com/goccy/llamawasm2go/p3.Fn3109
func Fn3109(m *base.Module, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32) int64

//go:linkname Fn3111 github.com/goccy/llamawasm2go/p2.Fn3111
func Fn3111(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn3156 github.com/goccy/llamawasm2go/p2.Fn3156
func Fn3156(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3170 github.com/goccy/llamawasm2go/p2.Fn3170
func Fn3170(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3202 github.com/goccy/llamawasm2go/p2.Fn3202
func Fn3202(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3213 github.com/goccy/llamawasm2go/p3.Fn3213
func Fn3213(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3225 github.com/goccy/llamawasm2go/p2.Fn3225
func Fn3225(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn3230 github.com/goccy/llamawasm2go/p3.Fn3230
func Fn3230(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3231 github.com/goccy/llamawasm2go/p3.Fn3231
func Fn3231(m *base.Module, l0 int64)

//go:linkname Fn3232 github.com/goccy/llamawasm2go/p2.Fn3232
func Fn3232(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3233 github.com/goccy/llamawasm2go/p2.Fn3233
func Fn3233(m *base.Module, l0 int64)

//go:linkname Fn3237 github.com/goccy/llamawasm2go/p3.Fn3237
func Fn3237(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3244 github.com/goccy/llamawasm2go/p2.Fn3244
func Fn3244(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn3245 github.com/goccy/llamawasm2go/p2.Fn3245
func Fn3245(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3266 github.com/goccy/llamawasm2go/p0.Fn3266
func Fn3266(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3270 github.com/goccy/llamawasm2go/p2.Fn3270
func Fn3270(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3278 github.com/goccy/llamawasm2go/p2.Fn3278
func Fn3278(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3280 github.com/goccy/llamawasm2go/p2.Fn3280
func Fn3280(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3281 github.com/goccy/llamawasm2go/p3.Fn3281
func Fn3281(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3282 github.com/goccy/llamawasm2go/p2.Fn3282
func Fn3282(m *base.Module)

//go:linkname Fn3288 github.com/goccy/llamawasm2go/p2.Fn3288
func Fn3288(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3306 github.com/goccy/llamawasm2go/p2.Fn3306
func Fn3306(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn3309 github.com/goccy/llamawasm2go/p3.Fn3309
func Fn3309(m *base.Module, l0 int64)

//go:linkname Fn3310 github.com/goccy/llamawasm2go/p2.Fn3310
func Fn3310(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3311 github.com/goccy/llamawasm2go/p2.Fn3311
func Fn3311(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3315 github.com/goccy/llamawasm2go/p2.Fn3315
func Fn3315(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn3320 github.com/goccy/llamawasm2go/p3.Fn3320
func Fn3320(m *base.Module, l0 int64) int64

//go:linkname Fn3321 github.com/goccy/llamawasm2go/p3.Fn3321
func Fn3321(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn3333 github.com/goccy/llamawasm2go/p2.Fn3333
func Fn3333(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3334 github.com/goccy/llamawasm2go/p3.Fn3334
func Fn3334(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3335 github.com/goccy/llamawasm2go/p2.Fn3335
func Fn3335(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn3336 github.com/goccy/llamawasm2go/p2.Fn3336
func Fn3336(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3337 github.com/goccy/llamawasm2go/p2.Fn3337
func Fn3337(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn3338 github.com/goccy/llamawasm2go/p2.Fn3338
func Fn3338(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3355 github.com/goccy/llamawasm2go/p2.Fn3355
func Fn3355(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3356 github.com/goccy/llamawasm2go/p2.Fn3356
func Fn3356(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3357 github.com/goccy/llamawasm2go/p2.Fn3357
func Fn3357(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn3358 github.com/goccy/llamawasm2go/p2.Fn3358
func Fn3358(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3359 github.com/goccy/llamawasm2go/p2.Fn3359
func Fn3359(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn3360 github.com/goccy/llamawasm2go/p2.Fn3360
func Fn3360(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn3361 github.com/goccy/llamawasm2go/p2.Fn3361
func Fn3361(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn3370 github.com/goccy/llamawasm2go/p3.Fn3370
func Fn3370(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3384 github.com/goccy/llamawasm2go/p2.Fn3384
func Fn3384(m *base.Module, l0 int64) int64

//go:linkname Fn3385 github.com/goccy/llamawasm2go/p3.Fn3385
func Fn3385(m *base.Module, l0 int64)

//go:linkname Fn3388 github.com/goccy/llamawasm2go/p2.Fn3388
func Fn3388(m *base.Module, l0 int64)

//go:linkname Fn3389 github.com/goccy/llamawasm2go/p3.Fn3389
func Fn3389(m *base.Module, l0 int64)

//go:linkname Fn3391 github.com/goccy/llamawasm2go/p3.Fn3391
func Fn3391(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3392 github.com/goccy/llamawasm2go/p3.Fn3392
func Fn3392(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn3416 github.com/goccy/llamawasm2go/p2.Fn3416
func Fn3416(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3424 github.com/goccy/llamawasm2go/p2.Fn3424
func Fn3424(m *base.Module, l0 float64, l1 int64) int32

//go:linkname Fn3426 github.com/goccy/llamawasm2go/p3.Fn3426
func Fn3426(m *base.Module, l0 float64) float32

//go:linkname Fn3427 github.com/goccy/llamawasm2go/p3.Fn3427
func Fn3427(m *base.Module, l0 float64) float32

//go:linkname Fn3428 github.com/goccy/llamawasm2go/p2.Fn3428
func Fn3428(m *base.Module, l0 float32, l1 int64) int32

//go:linkname Fn3431 github.com/goccy/llamawasm2go/p3.Fn3431
func Fn3431(m *base.Module, l0 float64) float64

//go:linkname Fn3434 github.com/goccy/llamawasm2go/p3.Fn3434
func Fn3434(m *base.Module, l0 int32) float32

//go:linkname Fn3435 github.com/goccy/llamawasm2go/p3.Fn3435
func Fn3435(m *base.Module, l0 int32) float32

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

//go:linkname Fn3466 github.com/goccy/llamawasm2go/p3.Fn3466
func Fn3466(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3468 github.com/goccy/llamawasm2go/p3.Fn3468
func Fn3468(m *base.Module, l0 int64)

//go:linkname Fn3469 github.com/goccy/llamawasm2go/p3.Fn3469
func Fn3469(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3470 github.com/goccy/llamawasm2go/p3.Fn3470
func Fn3470(m *base.Module, l0 int64) int32

//go:linkname Fn3477 github.com/goccy/llamawasm2go/p3.Fn3477
func Fn3477(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn3478 github.com/goccy/llamawasm2go/p2.Fn3478
func Fn3478(m *base.Module)

//go:linkname Fn3479 github.com/goccy/llamawasm2go/p3.Fn3479
func Fn3479(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn3485 github.com/goccy/llamawasm2go/p3.Fn3485
func Fn3485(m *base.Module, l0 int64) int32

//go:linkname Fn3488 github.com/goccy/llamawasm2go/p2.Fn3488
func Fn3488(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3491 github.com/goccy/llamawasm2go/p2.Fn3491
func Fn3491(m *base.Module, l0 int64) int32

//go:linkname Fn3492 github.com/goccy/llamawasm2go/p2.Fn3492
func Fn3492(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int64) int64

//go:linkname Fn3493 github.com/goccy/llamawasm2go/p3.Fn3493
func Fn3493(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3495 github.com/goccy/llamawasm2go/p2.Fn3495
func Fn3495(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3496 github.com/goccy/llamawasm2go/p3.Fn3496
func Fn3496(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3499 github.com/goccy/llamawasm2go/p3.Fn3499
func Fn3499(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3502 github.com/goccy/llamawasm2go/p2.Fn3502
func Fn3502(m *base.Module, l0 int64) int64

//go:linkname Fn3507 github.com/goccy/llamawasm2go/p2.Fn3507
func Fn3507(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3516 github.com/goccy/llamawasm2go/p3.Fn3516
func Fn3516(m *base.Module)

//go:linkname Fn3518 github.com/goccy/llamawasm2go/p2.Fn3518
func Fn3518(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn3522 github.com/goccy/llamawasm2go/p0.Fn3522
func Fn3522(m *base.Module, l0 int64) int64

//go:linkname Fn3523 github.com/goccy/llamawasm2go/p2.Fn3523
func Fn3523(m *base.Module, l0 int64)

//go:linkname Fn3524 github.com/goccy/llamawasm2go/p3.Fn3524
func Fn3524(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3525 github.com/goccy/llamawasm2go/p2.Fn3525
func Fn3525(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3528 github.com/goccy/llamawasm2go/p3.Fn3528
func Fn3528(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3601 github.com/goccy/llamawasm2go/p3.Fn3601
func Fn3601(m *base.Module, l0 int32)
