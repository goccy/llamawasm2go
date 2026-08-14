//go:build !arm64 && (!amd64 || !amd64.v2)

package p1

import (
	base "github.com/goccy/llamawasm2go/base"
	_ "unsafe"
)

//go:linkname Fn18 github.com/goccy/llamawasm2go/p2.Fn18
func Fn18(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn45 github.com/goccy/llamawasm2go/p2.Fn45
func Fn45(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn48 github.com/goccy/llamawasm2go/p2.Fn48
func Fn48(m *base.Module, l0 int32, l1 int64)

//go:linkname Fn52 github.com/goccy/llamawasm2go/p2.Fn52
func Fn52(m *base.Module, l0 int64) int64

//go:linkname Fn53 github.com/goccy/llamawasm2go/p2.Fn53
func Fn53(m *base.Module, l0 int64)

//go:linkname Fn55 github.com/goccy/llamawasm2go/p2.Fn55
func Fn55(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn56 github.com/goccy/llamawasm2go/p2.Fn56
func Fn56(m *base.Module, l0 int64) int64

//go:linkname Fn57 github.com/goccy/llamawasm2go/p2.Fn57
func Fn57(m *base.Module)

//go:linkname Fn63 github.com/goccy/llamawasm2go/p0.Fn63
func Fn63(m *base.Module, l0 int64) int64

//go:linkname Fn64 github.com/goccy/llamawasm2go/p2.Fn64
func Fn64(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn65 github.com/goccy/llamawasm2go/p0.Fn65
func Fn65(m *base.Module, l0 int64) int64

//go:linkname Fn81 github.com/goccy/llamawasm2go/p2.Fn81
func Fn81(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn126 github.com/goccy/llamawasm2go/p2.Fn126
func Fn126(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn132 github.com/goccy/llamawasm2go/p2.Fn132
func Fn132(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn165 github.com/goccy/llamawasm2go/p2.Fn165
func Fn165(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn239 github.com/goccy/llamawasm2go/p2.Fn239
func Fn239(m *base.Module)

//go:linkname Fn240 github.com/goccy/llamawasm2go/p2.Fn240
func Fn240(m *base.Module, l0 int64)

//go:linkname Fn241 github.com/goccy/llamawasm2go/p2.Fn241
func Fn241(m *base.Module, l0 int64) int64

//go:linkname Fn248 github.com/goccy/llamawasm2go/p2.Fn248
func Fn248(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn249 github.com/goccy/llamawasm2go/p2.Fn249
func Fn249(m *base.Module, l0 int64, l1 int64) int64

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

//go:linkname Fn263 github.com/goccy/llamawasm2go/p2.Fn263
func Fn263(m *base.Module, l0 int64, l1 int64)

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
func Fn271(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn272 github.com/goccy/llamawasm2go/p2.Fn272
func Fn272(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64)

//go:linkname Fn274 github.com/goccy/llamawasm2go/p2.Fn274
func Fn274(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn275 github.com/goccy/llamawasm2go/p2.Fn275
func Fn275(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn276 github.com/goccy/llamawasm2go/p2.Fn276
func Fn276(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn279 github.com/goccy/llamawasm2go/p2.Fn279
func Fn279(m *base.Module, l0 int64) int64

//go:linkname Fn280 github.com/goccy/llamawasm2go/p2.Fn280
func Fn280(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn283 github.com/goccy/llamawasm2go/p2.Fn283
func Fn283(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn289 github.com/goccy/llamawasm2go/p2.Fn289
func Fn289(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn290 github.com/goccy/llamawasm2go/p2.Fn290
func Fn290(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn345 github.com/goccy/llamawasm2go/p2.Fn345
func Fn345(m *base.Module) int64

//go:linkname Fn347 github.com/goccy/llamawasm2go/p2.Fn347
func Fn347(m *base.Module, l0 int64, l1 float64)

//go:linkname Fn348 github.com/goccy/llamawasm2go/p2.Fn348
func Fn348(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn352 github.com/goccy/llamawasm2go/p2.Fn352
func Fn352(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn353 github.com/goccy/llamawasm2go/p2.Fn353
func Fn353(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn356 github.com/goccy/llamawasm2go/p2.Fn356
func Fn356(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn357 github.com/goccy/llamawasm2go/p2.Fn357
func Fn357(m *base.Module, l0 int64)

//go:linkname Fn358 github.com/goccy/llamawasm2go/p2.Fn358
func Fn358(m *base.Module, l0 int64)

//go:linkname Fn360 github.com/goccy/llamawasm2go/p2.Fn360
func Fn360(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn362 github.com/goccy/llamawasm2go/p2.Fn362
func Fn362(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int64, l5 int64) int32

//go:linkname Fn363 github.com/goccy/llamawasm2go/p2.Fn363
func Fn363(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn365 github.com/goccy/llamawasm2go/p2.Fn365
func Fn365(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn371 github.com/goccy/llamawasm2go/p2.Fn371
func Fn371(m *base.Module, l0 int64)

//go:linkname Fn374 github.com/goccy/llamawasm2go/p2.Fn374
func Fn374(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn375 github.com/goccy/llamawasm2go/p2.Fn375
func Fn375(m *base.Module, l0 int64)

//go:linkname Fn376 github.com/goccy/llamawasm2go/p2.Fn376
func Fn376(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn377 github.com/goccy/llamawasm2go/p2.Fn377
func Fn377(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn378 github.com/goccy/llamawasm2go/p2.Fn378
func Fn378(m *base.Module)

//go:linkname Fn379 github.com/goccy/llamawasm2go/p0.Fn379
func Fn379(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int64, l5 int32, l6 int32)

//go:linkname Fn383 github.com/goccy/llamawasm2go/p0.Fn383
func Fn383(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 int64)

//go:linkname Fn392 github.com/goccy/llamawasm2go/p2.Fn392
func Fn392(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn399 github.com/goccy/llamawasm2go/p2.Fn399
func Fn399(m *base.Module, l0 int64)

//go:linkname Fn416 github.com/goccy/llamawasm2go/p2.Fn416
func Fn416(m *base.Module, l0 int64)

//go:linkname Fn417 github.com/goccy/llamawasm2go/p2.Fn417
func Fn417(m *base.Module, l0 int64)

//go:linkname Fn420 github.com/goccy/llamawasm2go/p2.Fn420
func Fn420(m *base.Module, l0 int64)

//go:linkname Fn423 github.com/goccy/llamawasm2go/p2.Fn423
func Fn423(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn424 github.com/goccy/llamawasm2go/p2.Fn424
func Fn424(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn426 github.com/goccy/llamawasm2go/p2.Fn426
func Fn426(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn429 github.com/goccy/llamawasm2go/p2.Fn429
func Fn429(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn434 github.com/goccy/llamawasm2go/p2.Fn434
func Fn434(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn440 github.com/goccy/llamawasm2go/p2.Fn440
func Fn440(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn442 github.com/goccy/llamawasm2go/p2.Fn442
func Fn442(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn444 github.com/goccy/llamawasm2go/p2.Fn444
func Fn444(m *base.Module, l0 int64)

//go:linkname Fn446 github.com/goccy/llamawasm2go/p2.Fn446
func Fn446(m *base.Module, l0 int64)

//go:linkname Fn447 github.com/goccy/llamawasm2go/p0.Fn447
func Fn447(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn449 github.com/goccy/llamawasm2go/p2.Fn449
func Fn449(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn463 github.com/goccy/llamawasm2go/p2.Fn463
func Fn463(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn465 github.com/goccy/llamawasm2go/p2.Fn465
func Fn465(m *base.Module, l0 int64)

//go:linkname Fn466 github.com/goccy/llamawasm2go/p2.Fn466
func Fn466(m *base.Module, l0 int64) int64

//go:linkname Fn480 github.com/goccy/llamawasm2go/p2.Fn480
func Fn480(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn489 github.com/goccy/llamawasm2go/p2.Fn489
func Fn489(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn494 github.com/goccy/llamawasm2go/p2.Fn494
func Fn494(m *base.Module, l0 int64) int64

//go:linkname Fn495 github.com/goccy/llamawasm2go/p2.Fn495
func Fn495(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn496 github.com/goccy/llamawasm2go/p2.Fn496
func Fn496(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn497 github.com/goccy/llamawasm2go/p2.Fn497
func Fn497(m *base.Module, l0 int64) int64

//go:linkname Fn498 github.com/goccy/llamawasm2go/p2.Fn498
func Fn498(m *base.Module, l0 int64) int64

//go:linkname Fn499 github.com/goccy/llamawasm2go/p2.Fn499
func Fn499(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn501 github.com/goccy/llamawasm2go/p2.Fn501
func Fn501(m *base.Module, l0 int64) int64

//go:linkname Fn502 github.com/goccy/llamawasm2go/p2.Fn502
func Fn502(m *base.Module, l0 int64) int64

//go:linkname Fn504 github.com/goccy/llamawasm2go/p2.Fn504
func Fn504(m *base.Module, l0 int64)

//go:linkname Fn505 github.com/goccy/llamawasm2go/p2.Fn505
func Fn505(m *base.Module, l0 int64) int64

//go:linkname Fn506 github.com/goccy/llamawasm2go/p2.Fn506
func Fn506(m *base.Module, l0 int64) int64

//go:linkname Fn507 github.com/goccy/llamawasm2go/p2.Fn507
func Fn507(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn508 github.com/goccy/llamawasm2go/p2.Fn508
func Fn508(m *base.Module, l0 int64) int32

//go:linkname Fn509 github.com/goccy/llamawasm2go/p2.Fn509
func Fn509(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn511 github.com/goccy/llamawasm2go/p2.Fn511
func Fn511(m *base.Module, l0 int64)

//go:linkname Fn512 github.com/goccy/llamawasm2go/p2.Fn512
func Fn512(m *base.Module, l0 int64)

//go:linkname Fn513 github.com/goccy/llamawasm2go/p2.Fn513
func Fn513(m *base.Module, l0 int64) int64

//go:linkname Fn515 github.com/goccy/llamawasm2go/p2.Fn515
func Fn515(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn516 github.com/goccy/llamawasm2go/p2.Fn516
func Fn516(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn517 github.com/goccy/llamawasm2go/p2.Fn517
func Fn517(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn518 github.com/goccy/llamawasm2go/p2.Fn518
func Fn518(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn519 github.com/goccy/llamawasm2go/p2.Fn519
func Fn519(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn520 github.com/goccy/llamawasm2go/p2.Fn520
func Fn520(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64)

//go:linkname Fn523 github.com/goccy/llamawasm2go/p2.Fn523
func Fn523(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn524 github.com/goccy/llamawasm2go/p2.Fn524
func Fn524(m *base.Module, l0 int64) int64

//go:linkname Fn525 github.com/goccy/llamawasm2go/p2.Fn525
func Fn525(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn528 github.com/goccy/llamawasm2go/p2.Fn528
func Fn528(m *base.Module, l0 int64) int64

//go:linkname Fn529 github.com/goccy/llamawasm2go/p2.Fn529
func Fn529(m *base.Module, l0 int64) int64

//go:linkname Fn533 github.com/goccy/llamawasm2go/p2.Fn533
func Fn533(m *base.Module, l0 int64) int64

//go:linkname Fn534 github.com/goccy/llamawasm2go/p2.Fn534
func Fn534(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn535 github.com/goccy/llamawasm2go/p2.Fn535
func Fn535(m *base.Module, l0 int64) int64

//go:linkname Fn539 github.com/goccy/llamawasm2go/p2.Fn539
func Fn539(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn540 github.com/goccy/llamawasm2go/p0.Fn540
func Fn540(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn543 github.com/goccy/llamawasm2go/p2.Fn543
func Fn543(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn545 github.com/goccy/llamawasm2go/p2.Fn545
func Fn545(m *base.Module, l0 int64)

//go:linkname Fn546 github.com/goccy/llamawasm2go/p2.Fn546
func Fn546(m *base.Module, l0 int64)

//go:linkname Fn547 github.com/goccy/llamawasm2go/p2.Fn547
func Fn547(m *base.Module, l0 int64)

//go:linkname Fn548 github.com/goccy/llamawasm2go/p2.Fn548
func Fn548(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn549 github.com/goccy/llamawasm2go/p2.Fn549
func Fn549(m *base.Module, l0 int64) int32

//go:linkname Fn550 github.com/goccy/llamawasm2go/p2.Fn550
func Fn550(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn551 github.com/goccy/llamawasm2go/p2.Fn551
func Fn551(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn552 github.com/goccy/llamawasm2go/p2.Fn552
func Fn552(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn553 github.com/goccy/llamawasm2go/p2.Fn553
func Fn553(m *base.Module, l0 int64) int32

//go:linkname Fn554 github.com/goccy/llamawasm2go/p2.Fn554
func Fn554(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn629 github.com/goccy/llamawasm2go/p2.Fn629
func Fn629(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64)

//go:linkname Fn630 github.com/goccy/llamawasm2go/p2.Fn630
func Fn630(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn651 github.com/goccy/llamawasm2go/p2.Fn651
func Fn651(m *base.Module, l0 int64) int32

//go:linkname Fn652 github.com/goccy/llamawasm2go/p2.Fn652
func Fn652(m *base.Module, l0 int64) int32

//go:linkname Fn663 github.com/goccy/llamawasm2go/p2.Fn663
func Fn663(m *base.Module, l0 int64) int64

//go:linkname Fn668 github.com/goccy/llamawasm2go/p2.Fn668
func Fn668(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn670 github.com/goccy/llamawasm2go/p2.Fn670
func Fn670(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn671 github.com/goccy/llamawasm2go/p2.Fn671
func Fn671(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn672 github.com/goccy/llamawasm2go/p2.Fn672
func Fn672(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64) int64

//go:linkname Fn673 github.com/goccy/llamawasm2go/p2.Fn673
func Fn673(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn674 github.com/goccy/llamawasm2go/p2.Fn674
func Fn674(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn675 github.com/goccy/llamawasm2go/p2.Fn675
func Fn675(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn678 github.com/goccy/llamawasm2go/p2.Fn678
func Fn678(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn679 github.com/goccy/llamawasm2go/p2.Fn679
func Fn679(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn682 github.com/goccy/llamawasm2go/p2.Fn682
func Fn682(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn683 github.com/goccy/llamawasm2go/p2.Fn683
func Fn683(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn684 github.com/goccy/llamawasm2go/p2.Fn684
func Fn684(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn685 github.com/goccy/llamawasm2go/p2.Fn685
func Fn685(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn686 github.com/goccy/llamawasm2go/p2.Fn686
func Fn686(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn687 github.com/goccy/llamawasm2go/p2.Fn687
func Fn687(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn689 github.com/goccy/llamawasm2go/p2.Fn689
func Fn689(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn691 github.com/goccy/llamawasm2go/p2.Fn691
func Fn691(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn693 github.com/goccy/llamawasm2go/p2.Fn693
func Fn693(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn694 github.com/goccy/llamawasm2go/p2.Fn694
func Fn694(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn695 github.com/goccy/llamawasm2go/p2.Fn695
func Fn695(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn696 github.com/goccy/llamawasm2go/p2.Fn696
func Fn696(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn697 github.com/goccy/llamawasm2go/p2.Fn697
func Fn697(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn699 github.com/goccy/llamawasm2go/p2.Fn699
func Fn699(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn700 github.com/goccy/llamawasm2go/p2.Fn700
func Fn700(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn701 github.com/goccy/llamawasm2go/p2.Fn701
func Fn701(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn702 github.com/goccy/llamawasm2go/p2.Fn702
func Fn702(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn703 github.com/goccy/llamawasm2go/p2.Fn703
func Fn703(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn704 github.com/goccy/llamawasm2go/p2.Fn704
func Fn704(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn706 github.com/goccy/llamawasm2go/p2.Fn706
func Fn706(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn707 github.com/goccy/llamawasm2go/p2.Fn707
func Fn707(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn708 github.com/goccy/llamawasm2go/p2.Fn708
func Fn708(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn709 github.com/goccy/llamawasm2go/p2.Fn709
func Fn709(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn710 github.com/goccy/llamawasm2go/p2.Fn710
func Fn710(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn711 github.com/goccy/llamawasm2go/p2.Fn711
func Fn711(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn712 github.com/goccy/llamawasm2go/p2.Fn712
func Fn712(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn713 github.com/goccy/llamawasm2go/p2.Fn713
func Fn713(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn714 github.com/goccy/llamawasm2go/p2.Fn714
func Fn714(m *base.Module, l0 int64)

//go:linkname Fn715 github.com/goccy/llamawasm2go/p2.Fn715
func Fn715(m *base.Module, l0 int64)

//go:linkname Fn716 github.com/goccy/llamawasm2go/p2.Fn716
func Fn716(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn717 github.com/goccy/llamawasm2go/p2.Fn717
func Fn717(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn718 github.com/goccy/llamawasm2go/p2.Fn718
func Fn718(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn719 github.com/goccy/llamawasm2go/p2.Fn719
func Fn719(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) int64

//go:linkname Fn720 github.com/goccy/llamawasm2go/p2.Fn720
func Fn720(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn721 github.com/goccy/llamawasm2go/p2.Fn721
func Fn721(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn722 github.com/goccy/llamawasm2go/p2.Fn722
func Fn722(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn723 github.com/goccy/llamawasm2go/p2.Fn723
func Fn723(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn724 github.com/goccy/llamawasm2go/p2.Fn724
func Fn724(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn725 github.com/goccy/llamawasm2go/p2.Fn725
func Fn725(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn727 github.com/goccy/llamawasm2go/p2.Fn727
func Fn727(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn728 github.com/goccy/llamawasm2go/p2.Fn728
func Fn728(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn729 github.com/goccy/llamawasm2go/p2.Fn729
func Fn729(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn730 github.com/goccy/llamawasm2go/p2.Fn730
func Fn730(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn731 github.com/goccy/llamawasm2go/p2.Fn731
func Fn731(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64) int64

//go:linkname Fn732 github.com/goccy/llamawasm2go/p2.Fn732
func Fn732(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64) int64

//go:linkname Fn733 github.com/goccy/llamawasm2go/p2.Fn733
func Fn733(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64) int64

//go:linkname Fn734 github.com/goccy/llamawasm2go/p2.Fn734
func Fn734(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn735 github.com/goccy/llamawasm2go/p2.Fn735
func Fn735(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn736 github.com/goccy/llamawasm2go/p2.Fn736
func Fn736(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn737 github.com/goccy/llamawasm2go/p2.Fn737
func Fn737(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn738 github.com/goccy/llamawasm2go/p2.Fn738
func Fn738(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn740 github.com/goccy/llamawasm2go/p2.Fn740
func Fn740(m *base.Module, l0 int64, l1 int64, l2 int64, l3 float32, l4 float32) int64

//go:linkname Fn743 github.com/goccy/llamawasm2go/p2.Fn743
func Fn743(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int64, l6 int32, l7 int32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32, l13 float32) int64

//go:linkname Fn744 github.com/goccy/llamawasm2go/p2.Fn744
func Fn744(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32, l7 float32, l8 float32, l9 float32, l10 float32, l11 float32, l12 float32) int64

//go:linkname Fn746 github.com/goccy/llamawasm2go/p2.Fn746
func Fn746(m *base.Module, l0 int64, l1 int64, l2 float32, l3 float32) int64

//go:linkname Fn747 github.com/goccy/llamawasm2go/p2.Fn747
func Fn747(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn748 github.com/goccy/llamawasm2go/p2.Fn748
func Fn748(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn750 github.com/goccy/llamawasm2go/p2.Fn750
func Fn750(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int64

//go:linkname Fn752 github.com/goccy/llamawasm2go/p2.Fn752
func Fn752(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn753 github.com/goccy/llamawasm2go/p2.Fn753
func Fn753(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn754 github.com/goccy/llamawasm2go/p2.Fn754
func Fn754(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn756 github.com/goccy/llamawasm2go/p2.Fn756
func Fn756(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 float32, l6 float32, l7 float32) int64

//go:linkname Fn757 github.com/goccy/llamawasm2go/p2.Fn757
func Fn757(m *base.Module, l0 int64)

//go:linkname Fn758 github.com/goccy/llamawasm2go/p2.Fn758
func Fn758(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn762 github.com/goccy/llamawasm2go/p2.Fn762
func Fn762(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn763 github.com/goccy/llamawasm2go/p2.Fn763
func Fn763(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn766 github.com/goccy/llamawasm2go/p2.Fn766
func Fn766(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn768 github.com/goccy/llamawasm2go/p2.Fn768
func Fn768(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn770 github.com/goccy/llamawasm2go/p2.Fn770
func Fn770(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn775 github.com/goccy/llamawasm2go/p2.Fn775
func Fn775(m *base.Module, l0 int64) int64

//go:linkname Fn785 github.com/goccy/llamawasm2go/p2.Fn785
func Fn785(m *base.Module, l0 int64)

//go:linkname Fn790 github.com/goccy/llamawasm2go/p2.Fn790
func Fn790(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn791 github.com/goccy/llamawasm2go/p2.Fn791
func Fn791(m *base.Module)

//go:linkname Fn800 github.com/goccy/llamawasm2go/p2.Fn800
func Fn800(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn808 github.com/goccy/llamawasm2go/p2.Fn808
func Fn808(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn809 github.com/goccy/llamawasm2go/p2.Fn809
func Fn809(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn814 github.com/goccy/llamawasm2go/p2.Fn814
func Fn814(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn817 github.com/goccy/llamawasm2go/p2.Fn817
func Fn817(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn823 github.com/goccy/llamawasm2go/p2.Fn823
func Fn823(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn824 github.com/goccy/llamawasm2go/p2.Fn824
func Fn824(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64) int32

//go:linkname Fn832 github.com/goccy/llamawasm2go/p2.Fn832
func Fn832(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn837 github.com/goccy/llamawasm2go/p2.Fn837
func Fn837(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn838 github.com/goccy/llamawasm2go/p2.Fn838
func Fn838(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn846 github.com/goccy/llamawasm2go/p2.Fn846
func Fn846(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32) int64

//go:linkname Fn862 github.com/goccy/llamawasm2go/p2.Fn862
func Fn862(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn870 github.com/goccy/llamawasm2go/p2.Fn870
func Fn870(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn887 github.com/goccy/llamawasm2go/p0.Fn887
func Fn887(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int32, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64) int32

//go:linkname Fn890 github.com/goccy/llamawasm2go/p2.Fn890
func Fn890(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn894 github.com/goccy/llamawasm2go/p2.Fn894
func Fn894(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64)

//go:linkname Fn903 github.com/goccy/llamawasm2go/p0.Fn903
func Fn903(m *base.Module, l0 int64) int64

//go:linkname Fn904 github.com/goccy/llamawasm2go/p2.Fn904
func Fn904(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn905 github.com/goccy/llamawasm2go/p2.Fn905
func Fn905(m *base.Module, l0 int64)

//go:linkname Fn907 github.com/goccy/llamawasm2go/p2.Fn907
func Fn907(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn998 github.com/goccy/llamawasm2go/p2.Fn998
func Fn998(m *base.Module, l0 int64)

//go:linkname Fn1020 github.com/goccy/llamawasm2go/p2.Fn1020
func Fn1020(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1030 github.com/goccy/llamawasm2go/p2.Fn1030
func Fn1030(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1036 github.com/goccy/llamawasm2go/p2.Fn1036
func Fn1036(m *base.Module)

//go:linkname Fn1040 github.com/goccy/llamawasm2go/p2.Fn1040
func Fn1040(m *base.Module, l0 int64) int64

//go:linkname Fn1212 github.com/goccy/llamawasm2go/p2.Fn1212
func Fn1212(m *base.Module) int64

//go:linkname Fn1246 github.com/goccy/llamawasm2go/p2.Fn1246
func Fn1246(m *base.Module, l0 int64)

//go:linkname Fn1250 github.com/goccy/llamawasm2go/p2.Fn1250
func Fn1250(m *base.Module, l0 int32) int64

//go:linkname Fn1261 github.com/goccy/llamawasm2go/p2.Fn1261
func Fn1261(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1262 github.com/goccy/llamawasm2go/p2.Fn1262
func Fn1262(m *base.Module, l0 int64)

//go:linkname Fn1264 github.com/goccy/llamawasm2go/p2.Fn1264
func Fn1264(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1285 github.com/goccy/llamawasm2go/p2.Fn1285
func Fn1285(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1286 github.com/goccy/llamawasm2go/p2.Fn1286
func Fn1286(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1287 github.com/goccy/llamawasm2go/p2.Fn1287
func Fn1287(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1288 github.com/goccy/llamawasm2go/p2.Fn1288
func Fn1288(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1289 github.com/goccy/llamawasm2go/p2.Fn1289
func Fn1289(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1290 github.com/goccy/llamawasm2go/p2.Fn1290
func Fn1290(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1291 github.com/goccy/llamawasm2go/p2.Fn1291
func Fn1291(m *base.Module, l0 int64) int64

//go:linkname Fn1294 github.com/goccy/llamawasm2go/p2.Fn1294
func Fn1294(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1298 github.com/goccy/llamawasm2go/p2.Fn1298
func Fn1298(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1299 github.com/goccy/llamawasm2go/p2.Fn1299
func Fn1299(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1300 github.com/goccy/llamawasm2go/p2.Fn1300
func Fn1300(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1301 github.com/goccy/llamawasm2go/p2.Fn1301
func Fn1301(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1302 github.com/goccy/llamawasm2go/p2.Fn1302
func Fn1302(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1303 github.com/goccy/llamawasm2go/p2.Fn1303
func Fn1303(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1304 github.com/goccy/llamawasm2go/p2.Fn1304
func Fn1304(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1305 github.com/goccy/llamawasm2go/p2.Fn1305
func Fn1305(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1306 github.com/goccy/llamawasm2go/p2.Fn1306
func Fn1306(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1307 github.com/goccy/llamawasm2go/p2.Fn1307
func Fn1307(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1308 github.com/goccy/llamawasm2go/p2.Fn1308
func Fn1308(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1309 github.com/goccy/llamawasm2go/p2.Fn1309
func Fn1309(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1310 github.com/goccy/llamawasm2go/p2.Fn1310
func Fn1310(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1311 github.com/goccy/llamawasm2go/p2.Fn1311
func Fn1311(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1312 github.com/goccy/llamawasm2go/p2.Fn1312
func Fn1312(m *base.Module, l0 int64, l1 int64, l2 float32) int64

//go:linkname Fn1313 github.com/goccy/llamawasm2go/p2.Fn1313
func Fn1313(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1314 github.com/goccy/llamawasm2go/p2.Fn1314
func Fn1314(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1315 github.com/goccy/llamawasm2go/p2.Fn1315
func Fn1315(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1316 github.com/goccy/llamawasm2go/p2.Fn1316
func Fn1316(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1317 github.com/goccy/llamawasm2go/p2.Fn1317
func Fn1317(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1318 github.com/goccy/llamawasm2go/p2.Fn1318
func Fn1318(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn1319 github.com/goccy/llamawasm2go/p2.Fn1319
func Fn1319(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1320 github.com/goccy/llamawasm2go/p2.Fn1320
func Fn1320(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1321 github.com/goccy/llamawasm2go/p2.Fn1321
func Fn1321(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1322 github.com/goccy/llamawasm2go/p2.Fn1322
func Fn1322(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1323 github.com/goccy/llamawasm2go/p2.Fn1323
func Fn1323(m *base.Module, l0 int64, l1 int64, l2 float64) int64

//go:linkname Fn1324 github.com/goccy/llamawasm2go/p2.Fn1324
func Fn1324(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1326 github.com/goccy/llamawasm2go/p2.Fn1326
func Fn1326(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1330 github.com/goccy/llamawasm2go/p2.Fn1330
func Fn1330(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1331 github.com/goccy/llamawasm2go/p2.Fn1331
func Fn1331(m *base.Module)

//go:linkname Fn1332 github.com/goccy/llamawasm2go/p2.Fn1332
func Fn1332(m *base.Module)

//go:linkname Fn1333 github.com/goccy/llamawasm2go/p0.Fn1333
func Fn1333(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1342 github.com/goccy/llamawasm2go/p2.Fn1342
func Fn1342(m *base.Module)

//go:linkname Fn1344 github.com/goccy/llamawasm2go/p2.Fn1344
func Fn1344(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1346 github.com/goccy/llamawasm2go/p0.Fn1346
func Fn1346(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1347 github.com/goccy/llamawasm2go/p2.Fn1347
func Fn1347(m *base.Module, l0 int64) int64

//go:linkname Fn1352 github.com/goccy/llamawasm2go/p2.Fn1352
func Fn1352(m *base.Module, l0 int64)

//go:linkname Fn1359 github.com/goccy/llamawasm2go/p2.Fn1359
func Fn1359(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1362 github.com/goccy/llamawasm2go/p2.Fn1362
func Fn1362(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1364 github.com/goccy/llamawasm2go/p2.Fn1364
func Fn1364(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1366 github.com/goccy/llamawasm2go/p2.Fn1366
func Fn1366(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1368 github.com/goccy/llamawasm2go/p2.Fn1368
func Fn1368(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1370 github.com/goccy/llamawasm2go/p2.Fn1370
func Fn1370(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1372 github.com/goccy/llamawasm2go/p2.Fn1372
func Fn1372(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1378 github.com/goccy/llamawasm2go/p2.Fn1378
func Fn1378(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1379 github.com/goccy/llamawasm2go/p2.Fn1379
func Fn1379(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1380 github.com/goccy/llamawasm2go/p2.Fn1380
func Fn1380(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1381 github.com/goccy/llamawasm2go/p0.Fn1381
func Fn1381(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32, l6 int32) int32

//go:linkname Fn1383 github.com/goccy/llamawasm2go/p2.Fn1383
func Fn1383(m *base.Module, l0 int64)

//go:linkname Fn1384 github.com/goccy/llamawasm2go/p2.Fn1384
func Fn1384(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1386 github.com/goccy/llamawasm2go/p2.Fn1386
func Fn1386(m *base.Module, l0 int64) int64

//go:linkname Fn1387 github.com/goccy/llamawasm2go/p2.Fn1387
func Fn1387(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1388 github.com/goccy/llamawasm2go/p2.Fn1388
func Fn1388(m *base.Module, l0 int64)

//go:linkname Fn1389 github.com/goccy/llamawasm2go/p2.Fn1389
func Fn1389(m *base.Module, l0 int64)

//go:linkname Fn1390 github.com/goccy/llamawasm2go/p2.Fn1390
func Fn1390(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn1391 github.com/goccy/llamawasm2go/p2.Fn1391
func Fn1391(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1395 github.com/goccy/llamawasm2go/p2.Fn1395
func Fn1395(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn1397 github.com/goccy/llamawasm2go/p2.Fn1397
func Fn1397(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1405 github.com/goccy/llamawasm2go/p2.Fn1405
func Fn1405(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1407 github.com/goccy/llamawasm2go/p0.Fn1407
func Fn1407(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1409 github.com/goccy/llamawasm2go/p2.Fn1409
func Fn1409(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1412 github.com/goccy/llamawasm2go/p0.Fn1412
func Fn1412(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1413 github.com/goccy/llamawasm2go/p2.Fn1413
func Fn1413(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1417 github.com/goccy/llamawasm2go/p2.Fn1417
func Fn1417(m *base.Module, l0 int64)

//go:linkname Fn1420 github.com/goccy/llamawasm2go/p2.Fn1420
func Fn1420(m *base.Module, l0 int64)

//go:linkname Fn1423 github.com/goccy/llamawasm2go/p2.Fn1423
func Fn1423(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1424 github.com/goccy/llamawasm2go/p2.Fn1424
func Fn1424(m *base.Module, l0 int64) int64

//go:linkname Fn1425 github.com/goccy/llamawasm2go/p2.Fn1425
func Fn1425(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1426 github.com/goccy/llamawasm2go/p2.Fn1426
func Fn1426(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1427 github.com/goccy/llamawasm2go/p2.Fn1427
func Fn1427(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1428 github.com/goccy/llamawasm2go/p2.Fn1428
func Fn1428(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32)

//go:linkname Fn1429 github.com/goccy/llamawasm2go/p2.Fn1429
func Fn1429(m *base.Module, l0 int64)

//go:linkname Fn1432 github.com/goccy/llamawasm2go/p2.Fn1432
func Fn1432(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1434 github.com/goccy/llamawasm2go/p2.Fn1434
func Fn1434(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1437 github.com/goccy/llamawasm2go/p2.Fn1437
func Fn1437(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1444 github.com/goccy/llamawasm2go/p2.Fn1444
func Fn1444(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1445 github.com/goccy/llamawasm2go/p2.Fn1445
func Fn1445(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1451 github.com/goccy/llamawasm2go/p2.Fn1451
func Fn1451(m *base.Module, l0 int64)

//go:linkname Fn1452 github.com/goccy/llamawasm2go/p2.Fn1452
func Fn1452(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1454 github.com/goccy/llamawasm2go/p2.Fn1454
func Fn1454(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn1455 github.com/goccy/llamawasm2go/p2.Fn1455
func Fn1455(m *base.Module, l0 int64)

//go:linkname Fn1457 github.com/goccy/llamawasm2go/p2.Fn1457
func Fn1457(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1499 github.com/goccy/llamawasm2go/p2.Fn1499
func Fn1499(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1515 github.com/goccy/llamawasm2go/p2.Fn1515
func Fn1515(m *base.Module, l0 int64)

//go:linkname Fn1516 github.com/goccy/llamawasm2go/p2.Fn1516
func Fn1516(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1517 github.com/goccy/llamawasm2go/p2.Fn1517
func Fn1517(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1518 github.com/goccy/llamawasm2go/p2.Fn1518
func Fn1518(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1519 github.com/goccy/llamawasm2go/p2.Fn1519
func Fn1519(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1520 github.com/goccy/llamawasm2go/p2.Fn1520
func Fn1520(m *base.Module, l0 int64)

//go:linkname Fn1521 github.com/goccy/llamawasm2go/p2.Fn1521
func Fn1521(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1523 github.com/goccy/llamawasm2go/p2.Fn1523
func Fn1523(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32)

//go:linkname Fn1524 github.com/goccy/llamawasm2go/p2.Fn1524
func Fn1524(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1525 github.com/goccy/llamawasm2go/p2.Fn1525
func Fn1525(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1526 github.com/goccy/llamawasm2go/p2.Fn1526
func Fn1526(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn1527 github.com/goccy/llamawasm2go/p2.Fn1527
func Fn1527(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn1530 github.com/goccy/llamawasm2go/p2.Fn1530
func Fn1530(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32, l10 int32, l11 float32, l12 int32, l13 int32, l14 int64, l15 int64, l16 int64, l17 int64, l18 int64, l19 int64) int64

//go:linkname Fn1532 github.com/goccy/llamawasm2go/p2.Fn1532
func Fn1532(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1533 github.com/goccy/llamawasm2go/p2.Fn1533
func Fn1533(m *base.Module, l0 int64) int64

//go:linkname Fn1534 github.com/goccy/llamawasm2go/p2.Fn1534
func Fn1534(m *base.Module, l0 int64) int64

//go:linkname Fn1535 github.com/goccy/llamawasm2go/p2.Fn1535
func Fn1535(m *base.Module, l0 int64) int64

//go:linkname Fn1536 github.com/goccy/llamawasm2go/p2.Fn1536
func Fn1536(m *base.Module, l0 int64) int64

//go:linkname Fn1537 github.com/goccy/llamawasm2go/p2.Fn1537
func Fn1537(m *base.Module, l0 int64) int64

//go:linkname Fn1538 github.com/goccy/llamawasm2go/p2.Fn1538
func Fn1538(m *base.Module, l0 int64) int64

//go:linkname Fn1540 github.com/goccy/llamawasm2go/p2.Fn1540
func Fn1540(m *base.Module, l0 int64) int64

//go:linkname Fn1541 github.com/goccy/llamawasm2go/p2.Fn1541
func Fn1541(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1543 github.com/goccy/llamawasm2go/p2.Fn1543
func Fn1543(m *base.Module, l0 int64) int64

//go:linkname Fn1544 github.com/goccy/llamawasm2go/p2.Fn1544
func Fn1544(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1545 github.com/goccy/llamawasm2go/p2.Fn1545
func Fn1545(m *base.Module, l0 int64) int64

//go:linkname Fn1547 github.com/goccy/llamawasm2go/p2.Fn1547
func Fn1547(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1548 github.com/goccy/llamawasm2go/p2.Fn1548
func Fn1548(m *base.Module, l0 int64) int64

//go:linkname Fn1550 github.com/goccy/llamawasm2go/p2.Fn1550
func Fn1550(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 float32, l9 int32) int64

//go:linkname Fn1551 github.com/goccy/llamawasm2go/p2.Fn1551
func Fn1551(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1552 github.com/goccy/llamawasm2go/p2.Fn1552
func Fn1552(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 float32, l10 int32) int64

//go:linkname Fn1553 github.com/goccy/llamawasm2go/p2.Fn1553
func Fn1553(m *base.Module, l0 int64) int64

//go:linkname Fn1555 github.com/goccy/llamawasm2go/p2.Fn1555
func Fn1555(m *base.Module, l0 int64) int64

//go:linkname Fn1556 github.com/goccy/llamawasm2go/p2.Fn1556
func Fn1556(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1561 github.com/goccy/llamawasm2go/p2.Fn1561
func Fn1561(m *base.Module, l0 int64) int64

//go:linkname Fn1562 github.com/goccy/llamawasm2go/p2.Fn1562
func Fn1562(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1563 github.com/goccy/llamawasm2go/p2.Fn1563
func Fn1563(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int64) int64

//go:linkname Fn1564 github.com/goccy/llamawasm2go/p2.Fn1564
func Fn1564(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1566 github.com/goccy/llamawasm2go/p2.Fn1566
func Fn1566(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1567 github.com/goccy/llamawasm2go/p2.Fn1567
func Fn1567(m *base.Module, l0 int64) int64

//go:linkname Fn1568 github.com/goccy/llamawasm2go/p2.Fn1568
func Fn1568(m *base.Module, l0 int64) int64

//go:linkname Fn1570 github.com/goccy/llamawasm2go/p2.Fn1570
func Fn1570(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1571 github.com/goccy/llamawasm2go/p2.Fn1571
func Fn1571(m *base.Module, l0 int64)

//go:linkname Fn1600 github.com/goccy/llamawasm2go/p2.Fn1600
func Fn1600(m *base.Module, l0 int64) int64

//go:linkname Fn1607 github.com/goccy/llamawasm2go/p2.Fn1607
func Fn1607(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1608 github.com/goccy/llamawasm2go/p2.Fn1608
func Fn1608(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1609 github.com/goccy/llamawasm2go/p2.Fn1609
func Fn1609(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1611 github.com/goccy/llamawasm2go/p2.Fn1611
func Fn1611(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1612 github.com/goccy/llamawasm2go/p2.Fn1612
func Fn1612(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1616 github.com/goccy/llamawasm2go/p2.Fn1616
func Fn1616(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1617 github.com/goccy/llamawasm2go/p2.Fn1617
func Fn1617(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1618 github.com/goccy/llamawasm2go/p2.Fn1618
func Fn1618(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1619 github.com/goccy/llamawasm2go/p2.Fn1619
func Fn1619(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1620 github.com/goccy/llamawasm2go/p2.Fn1620
func Fn1620(m *base.Module, l0 int64) int32

//go:linkname Fn1621 github.com/goccy/llamawasm2go/p2.Fn1621
func Fn1621(m *base.Module, l0 int64) int32

//go:linkname Fn1622 github.com/goccy/llamawasm2go/p2.Fn1622
func Fn1622(m *base.Module, l0 int64) int32

//go:linkname Fn1623 github.com/goccy/llamawasm2go/p2.Fn1623
func Fn1623(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn1626 github.com/goccy/llamawasm2go/p2.Fn1626
func Fn1626(m *base.Module, l0 int64) int32

//go:linkname Fn1627 github.com/goccy/llamawasm2go/p2.Fn1627
func Fn1627(m *base.Module, l0 int64) int32

//go:linkname Fn1633 github.com/goccy/llamawasm2go/p2.Fn1633
func Fn1633(m *base.Module, l0 int32, l1 int64, l2 int64)

//go:linkname Fn1634 github.com/goccy/llamawasm2go/p2.Fn1634
func Fn1634(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1635 github.com/goccy/llamawasm2go/p2.Fn1635
func Fn1635(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1636 github.com/goccy/llamawasm2go/p2.Fn1636
func Fn1636(m *base.Module)

//go:linkname Fn1637 github.com/goccy/llamawasm2go/p2.Fn1637
func Fn1637(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1639 github.com/goccy/llamawasm2go/p2.Fn1639
func Fn1639(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32)

//go:linkname Fn1641 github.com/goccy/llamawasm2go/p2.Fn1641
func Fn1641(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1642 github.com/goccy/llamawasm2go/p2.Fn1642
func Fn1642(m *base.Module, l0 int64)

//go:linkname Fn1646 github.com/goccy/llamawasm2go/p2.Fn1646
func Fn1646(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1647 github.com/goccy/llamawasm2go/p2.Fn1647
func Fn1647(m *base.Module, l0 int64)

//go:linkname Fn1650 github.com/goccy/llamawasm2go/p2.Fn1650
func Fn1650(m *base.Module, l0 int64)

//go:linkname Fn1660 github.com/goccy/llamawasm2go/p2.Fn1660
func Fn1660(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn1661 github.com/goccy/llamawasm2go/p2.Fn1661
func Fn1661(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn1670 github.com/goccy/llamawasm2go/p2.Fn1670
func Fn1670(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1671 github.com/goccy/llamawasm2go/p0.Fn1671
func Fn1671(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1673 github.com/goccy/llamawasm2go/p2.Fn1673
func Fn1673(m *base.Module, l0 int64)

//go:linkname Fn1675 github.com/goccy/llamawasm2go/p2.Fn1675
func Fn1675(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1682 github.com/goccy/llamawasm2go/p2.Fn1682
func Fn1682(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1683 github.com/goccy/llamawasm2go/p2.Fn1683
func Fn1683(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1684 github.com/goccy/llamawasm2go/p2.Fn1684
func Fn1684(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1690 github.com/goccy/llamawasm2go/p2.Fn1690
func Fn1690(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1697 github.com/goccy/llamawasm2go/p2.Fn1697
func Fn1697(m *base.Module, l0 int64)

//go:linkname Fn1700 github.com/goccy/llamawasm2go/p2.Fn1700
func Fn1700(m *base.Module, l0 int64) int32

//go:linkname Fn1710 github.com/goccy/llamawasm2go/p2.Fn1710
func Fn1710(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn1712 github.com/goccy/llamawasm2go/p2.Fn1712
func Fn1712(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1713 github.com/goccy/llamawasm2go/p2.Fn1713
func Fn1713(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int64

//go:linkname Fn1714 github.com/goccy/llamawasm2go/p2.Fn1714
func Fn1714(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1715 github.com/goccy/llamawasm2go/p2.Fn1715
func Fn1715(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1716 github.com/goccy/llamawasm2go/p2.Fn1716
func Fn1716(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1732 github.com/goccy/llamawasm2go/p2.Fn1732
func Fn1732(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int64, l14 int64, l15 int64, l16 int64) int64

//go:linkname Fn1745 github.com/goccy/llamawasm2go/p2.Fn1745
func Fn1745(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn1752 github.com/goccy/llamawasm2go/p2.Fn1752
func Fn1752(m *base.Module, l0 int64)

//go:linkname Fn1786 github.com/goccy/llamawasm2go/p2.Fn1786
func Fn1786(m *base.Module, l0 int64)

//go:linkname Fn1789 github.com/goccy/llamawasm2go/p2.Fn1789
func Fn1789(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1803 github.com/goccy/llamawasm2go/p2.Fn1803
func Fn1803(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64) int64

//go:linkname Fn1804 github.com/goccy/llamawasm2go/p2.Fn1804
func Fn1804(m *base.Module, l0 int64) int64

//go:linkname Fn1805 github.com/goccy/llamawasm2go/p2.Fn1805
func Fn1805(m *base.Module, l0 int64)

//go:linkname Fn1809 github.com/goccy/llamawasm2go/p0.Fn1809
func Fn1809(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1810 github.com/goccy/llamawasm2go/p2.Fn1810
func Fn1810(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1811 github.com/goccy/llamawasm2go/p2.Fn1811
func Fn1811(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1812 github.com/goccy/llamawasm2go/p2.Fn1812
func Fn1812(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1813 github.com/goccy/llamawasm2go/p2.Fn1813
func Fn1813(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn1831 github.com/goccy/llamawasm2go/p2.Fn1831
func Fn1831(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn1842 github.com/goccy/llamawasm2go/p2.Fn1842
func Fn1842(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1844 github.com/goccy/llamawasm2go/p2.Fn1844
func Fn1844(m *base.Module, l0 int64) int64

//go:linkname Fn1845 github.com/goccy/llamawasm2go/p2.Fn1845
func Fn1845(m *base.Module, l0 int64)

//go:linkname Fn1848 github.com/goccy/llamawasm2go/p0.Fn1848
func Fn1848(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32)

//go:linkname Fn1850 github.com/goccy/llamawasm2go/p2.Fn1850
func Fn1850(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1861 github.com/goccy/llamawasm2go/p2.Fn1861
func Fn1861(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1864 github.com/goccy/llamawasm2go/p2.Fn1864
func Fn1864(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1865 github.com/goccy/llamawasm2go/p2.Fn1865
func Fn1865(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1875 github.com/goccy/llamawasm2go/p2.Fn1875
func Fn1875(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn1902 github.com/goccy/llamawasm2go/p2.Fn1902
func Fn1902(m *base.Module, l0 int64)

//go:linkname Fn1903 github.com/goccy/llamawasm2go/p2.Fn1903
func Fn1903(m *base.Module, l0 int64)

//go:linkname Fn1914 github.com/goccy/llamawasm2go/p2.Fn1914
func Fn1914(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn1946 github.com/goccy/llamawasm2go/p2.Fn1946
func Fn1946(m *base.Module) int64

//go:linkname Fn1951 github.com/goccy/llamawasm2go/p2.Fn1951
func Fn1951(m *base.Module, l0 int64) int64

//go:linkname Fn1952 github.com/goccy/llamawasm2go/p2.Fn1952
func Fn1952(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn1959 github.com/goccy/llamawasm2go/p2.Fn1959
func Fn1959(m *base.Module, l0 int64, l1 int64, l2 int32)

//go:linkname Fn1963 github.com/goccy/llamawasm2go/p2.Fn1963
func Fn1963(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1965 github.com/goccy/llamawasm2go/p2.Fn1965
func Fn1965(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1966 github.com/goccy/llamawasm2go/p2.Fn1966
func Fn1966(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1967 github.com/goccy/llamawasm2go/p2.Fn1967
func Fn1967(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1968 github.com/goccy/llamawasm2go/p2.Fn1968
func Fn1968(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1969 github.com/goccy/llamawasm2go/p2.Fn1969
func Fn1969(m *base.Module, l0 int64, l1 int32, l2 int64, l3 int32) int32

//go:linkname Fn1973 github.com/goccy/llamawasm2go/p2.Fn1973
func Fn1973(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int32

//go:linkname Fn1980 github.com/goccy/llamawasm2go/p2.Fn1980
func Fn1980(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1982 github.com/goccy/llamawasm2go/p2.Fn1982
func Fn1982(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1985 github.com/goccy/llamawasm2go/p2.Fn1985
func Fn1985(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn1986 github.com/goccy/llamawasm2go/p2.Fn1986
func Fn1986(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn1987 github.com/goccy/llamawasm2go/p2.Fn1987
func Fn1987(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32) int32

//go:linkname Fn1990 github.com/goccy/llamawasm2go/p2.Fn1990
func Fn1990(m *base.Module, l0 int64)

//go:linkname Fn1999 github.com/goccy/llamawasm2go/p2.Fn1999
func Fn1999(m *base.Module, l0 int64)

//go:linkname Fn2001 github.com/goccy/llamawasm2go/p2.Fn2001
func Fn2001(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2002 github.com/goccy/llamawasm2go/p2.Fn2002
func Fn2002(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2015 github.com/goccy/llamawasm2go/p2.Fn2015
func Fn2015(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2016 github.com/goccy/llamawasm2go/p2.Fn2016
func Fn2016(m *base.Module, l0 int64) int64

//go:linkname Fn2017 github.com/goccy/llamawasm2go/p2.Fn2017
func Fn2017(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2019 github.com/goccy/llamawasm2go/p2.Fn2019
func Fn2019(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2033 github.com/goccy/llamawasm2go/p2.Fn2033
func Fn2033(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2034 github.com/goccy/llamawasm2go/p2.Fn2034
func Fn2034(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2036 github.com/goccy/llamawasm2go/p2.Fn2036
func Fn2036(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2037 github.com/goccy/llamawasm2go/p2.Fn2037
func Fn2037(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2038 github.com/goccy/llamawasm2go/p2.Fn2038
func Fn2038(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2039 github.com/goccy/llamawasm2go/p2.Fn2039
func Fn2039(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2040 github.com/goccy/llamawasm2go/p2.Fn2040
func Fn2040(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) int64

//go:linkname Fn2041 github.com/goccy/llamawasm2go/p2.Fn2041
func Fn2041(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2060 github.com/goccy/llamawasm2go/p2.Fn2060
func Fn2060(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2061 github.com/goccy/llamawasm2go/p2.Fn2061
func Fn2061(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2062 github.com/goccy/llamawasm2go/p2.Fn2062
func Fn2062(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2064 github.com/goccy/llamawasm2go/p2.Fn2064
func Fn2064(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2065 github.com/goccy/llamawasm2go/p2.Fn2065
func Fn2065(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2066 github.com/goccy/llamawasm2go/p2.Fn2066
func Fn2066(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2067 github.com/goccy/llamawasm2go/p2.Fn2067
func Fn2067(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2068 github.com/goccy/llamawasm2go/p2.Fn2068
func Fn2068(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2070 github.com/goccy/llamawasm2go/p2.Fn2070
func Fn2070(m *base.Module, l0 int64, l1 int32, l2 int32)

//go:linkname Fn2072 github.com/goccy/llamawasm2go/p2.Fn2072
func Fn2072(m *base.Module, l0 int64)

//go:linkname Fn2089 github.com/goccy/llamawasm2go/p2.Fn2089
func Fn2089(m *base.Module, l0 int64)

//go:linkname Fn2090 github.com/goccy/llamawasm2go/p2.Fn2090
func Fn2090(m *base.Module, l0 int64)

//go:linkname Fn2091 github.com/goccy/llamawasm2go/p2.Fn2091
func Fn2091(m *base.Module, l0 int64)

//go:linkname Fn2093 github.com/goccy/llamawasm2go/p2.Fn2093
func Fn2093(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2094 github.com/goccy/llamawasm2go/p2.Fn2094
func Fn2094(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2133 github.com/goccy/llamawasm2go/p2.Fn2133
func Fn2133(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2142 github.com/goccy/llamawasm2go/p2.Fn2142
func Fn2142(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2144 github.com/goccy/llamawasm2go/p2.Fn2144
func Fn2144(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2146 github.com/goccy/llamawasm2go/p2.Fn2146
func Fn2146(m *base.Module, l0 int64) int64

//go:linkname Fn2149 github.com/goccy/llamawasm2go/p2.Fn2149
func Fn2149(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2153 github.com/goccy/llamawasm2go/p2.Fn2153
func Fn2153(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2158 github.com/goccy/llamawasm2go/p2.Fn2158
func Fn2158(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2172 github.com/goccy/llamawasm2go/p2.Fn2172
func Fn2172(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2174 github.com/goccy/llamawasm2go/p2.Fn2174
func Fn2174(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2175 github.com/goccy/llamawasm2go/p2.Fn2175
func Fn2175(m *base.Module, l0 int64, l1 int64, l2 int32) float32

//go:linkname Fn2189 github.com/goccy/llamawasm2go/p2.Fn2189
func Fn2189(m *base.Module, l0 int64) int64

//go:linkname Fn2191 github.com/goccy/llamawasm2go/p2.Fn2191
func Fn2191(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int64, l4 int64, l5 int64, l6 int32)

//go:linkname Fn2194 github.com/goccy/llamawasm2go/p2.Fn2194
func Fn2194(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32)

//go:linkname Fn2198 github.com/goccy/llamawasm2go/p2.Fn2198
func Fn2198(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2199 github.com/goccy/llamawasm2go/p2.Fn2199
func Fn2199(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2200 github.com/goccy/llamawasm2go/p0.Fn2200
func Fn2200(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32) int64

//go:linkname Fn2202 github.com/goccy/llamawasm2go/p2.Fn2202
func Fn2202(m *base.Module, l0 int64) int64

//go:linkname Fn2208 github.com/goccy/llamawasm2go/p2.Fn2208
func Fn2208(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2212 github.com/goccy/llamawasm2go/p2.Fn2212
func Fn2212(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2213 github.com/goccy/llamawasm2go/p2.Fn2213
func Fn2213(m *base.Module, l0 int64)

//go:linkname Fn2214 github.com/goccy/llamawasm2go/p2.Fn2214
func Fn2214(m *base.Module, l0 int64)

//go:linkname Fn2215 github.com/goccy/llamawasm2go/p2.Fn2215
func Fn2215(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32

//go:linkname Fn2217 github.com/goccy/llamawasm2go/p2.Fn2217
func Fn2217(m *base.Module, l0 int64)

//go:linkname Fn2218 github.com/goccy/llamawasm2go/p2.Fn2218
func Fn2218(m *base.Module, l0 int64, l1 int64, l2 int32) int64

//go:linkname Fn2219 github.com/goccy/llamawasm2go/p2.Fn2219
func Fn2219(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2226 github.com/goccy/llamawasm2go/p2.Fn2226
func Fn2226(m *base.Module, l0 int64) int64

//go:linkname Fn2227 github.com/goccy/llamawasm2go/p2.Fn2227
func Fn2227(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2228 github.com/goccy/llamawasm2go/p2.Fn2228
func Fn2228(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2235 github.com/goccy/llamawasm2go/p2.Fn2235
func Fn2235(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2236 github.com/goccy/llamawasm2go/p2.Fn2236
func Fn2236(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int64

//go:linkname Fn2237 github.com/goccy/llamawasm2go/p2.Fn2237
func Fn2237(m *base.Module, l0 int64)

//go:linkname Fn2238 github.com/goccy/llamawasm2go/p2.Fn2238
func Fn2238(m *base.Module) int64

//go:linkname Fn2240 github.com/goccy/llamawasm2go/p2.Fn2240
func Fn2240(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2241 github.com/goccy/llamawasm2go/p2.Fn2241
func Fn2241(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2243 github.com/goccy/llamawasm2go/p2.Fn2243
func Fn2243(m *base.Module) int64

//go:linkname Fn2245 github.com/goccy/llamawasm2go/p2.Fn2245
func Fn2245(m *base.Module, l0 int32) int64

//go:linkname Fn2246 github.com/goccy/llamawasm2go/p2.Fn2246
func Fn2246(m *base.Module, l0 int32) int32

//go:linkname Fn2247 github.com/goccy/llamawasm2go/p2.Fn2247
func Fn2247(m *base.Module, l0 int32) int64

//go:linkname Fn2248 github.com/goccy/llamawasm2go/p2.Fn2248
func Fn2248(m *base.Module, l0 float32) int64

//go:linkname Fn2249 github.com/goccy/llamawasm2go/p2.Fn2249
func Fn2249(m *base.Module, l0 float32) int64

//go:linkname Fn2250 github.com/goccy/llamawasm2go/p2.Fn2250
func Fn2250(m *base.Module, l0 float32) int64

//go:linkname Fn2252 github.com/goccy/llamawasm2go/p2.Fn2252
func Fn2252(m *base.Module, l0 int32, l1 float32, l2 float32, l3 float32) int64

//go:linkname Fn2253 github.com/goccy/llamawasm2go/p2.Fn2253
func Fn2253(m *base.Module, l0 int32, l1 int32, l2 int64) int64

//go:linkname Fn2294 github.com/goccy/llamawasm2go/p2.Fn2294
func Fn2294(m *base.Module, l0 int64)

//go:linkname Fn2296 github.com/goccy/llamawasm2go/p2.Fn2296
func Fn2296(m *base.Module, l0 int64)

//go:linkname Fn2338 github.com/goccy/llamawasm2go/p2.Fn2338
func Fn2338(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2349 github.com/goccy/llamawasm2go/p2.Fn2349
func Fn2349(m *base.Module, l0 int64, l1 int32) int32

//go:linkname Fn2352 github.com/goccy/llamawasm2go/p2.Fn2352
func Fn2352(m *base.Module, l0 int64)

//go:linkname Fn2353 github.com/goccy/llamawasm2go/p2.Fn2353
func Fn2353(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn2354 github.com/goccy/llamawasm2go/p2.Fn2354
func Fn2354(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2356 github.com/goccy/llamawasm2go/p2.Fn2356
func Fn2356(m *base.Module, l0 int64)

//go:linkname Fn2360 github.com/goccy/llamawasm2go/p2.Fn2360
func Fn2360(m *base.Module, l0 int64, l1 int32) int64

//go:linkname Fn2366 github.com/goccy/llamawasm2go/p2.Fn2366
func Fn2366(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64)

//go:linkname Fn2381 github.com/goccy/llamawasm2go/p2.Fn2381
func Fn2381(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2388 github.com/goccy/llamawasm2go/p0.Fn2388
func Fn2388(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2391 github.com/goccy/llamawasm2go/p2.Fn2391
func Fn2391(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2394 github.com/goccy/llamawasm2go/p2.Fn2394
func Fn2394(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2397 github.com/goccy/llamawasm2go/p2.Fn2397
func Fn2397(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2398 github.com/goccy/llamawasm2go/p2.Fn2398
func Fn2398(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2404 github.com/goccy/llamawasm2go/p2.Fn2404
func Fn2404(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2406 github.com/goccy/llamawasm2go/p2.Fn2406
func Fn2406(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2421 github.com/goccy/llamawasm2go/p2.Fn2421
func Fn2421(m *base.Module, l0 int64)

//go:linkname Fn2422 github.com/goccy/llamawasm2go/p2.Fn2422
func Fn2422(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2427 github.com/goccy/llamawasm2go/p2.Fn2427
func Fn2427(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int32, l5 int32) int64

//go:linkname Fn2432 github.com/goccy/llamawasm2go/p2.Fn2432
func Fn2432(m *base.Module, l0 int64) int64

//go:linkname Fn2433 github.com/goccy/llamawasm2go/p2.Fn2433
func Fn2433(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32)

//go:linkname Fn2434 github.com/goccy/llamawasm2go/p0.Fn2434
func Fn2434(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2445 github.com/goccy/llamawasm2go/p2.Fn2445
func Fn2445(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2446 github.com/goccy/llamawasm2go/p2.Fn2446
func Fn2446(m *base.Module, l0 int64, l1 int32)

//go:linkname Fn2447 github.com/goccy/llamawasm2go/p2.Fn2447
func Fn2447(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2448 github.com/goccy/llamawasm2go/p2.Fn2448
func Fn2448(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2468 github.com/goccy/llamawasm2go/p2.Fn2468
func Fn2468(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2469 github.com/goccy/llamawasm2go/p2.Fn2469
func Fn2469(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64)

//go:linkname Fn2472 github.com/goccy/llamawasm2go/p2.Fn2472
func Fn2472(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int64

//go:linkname Fn2558 github.com/goccy/llamawasm2go/p2.Fn2558
func Fn2558(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) int64

//go:linkname Fn2677 github.com/goccy/llamawasm2go/p2.Fn2677
func Fn2677(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2678 github.com/goccy/llamawasm2go/p0.Fn2678
func Fn2678(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int32)

//go:linkname Fn2679 github.com/goccy/llamawasm2go/p2.Fn2679
func Fn2679(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int32) int64

//go:linkname Fn2680 github.com/goccy/llamawasm2go/p2.Fn2680
func Fn2680(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int32) int64

//go:linkname Fn2684 github.com/goccy/llamawasm2go/p2.Fn2684
func Fn2684(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64, l9 int64, l10 int64, l11 int64, l12 int64, l13 int64, l14 int64) int64

//go:linkname Fn2709 github.com/goccy/llamawasm2go/p2.Fn2709
func Fn2709(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2841 github.com/goccy/llamawasm2go/p2.Fn2841
func Fn2841(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2848 github.com/goccy/llamawasm2go/p2.Fn2848
func Fn2848(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2876 github.com/goccy/llamawasm2go/p2.Fn2876
func Fn2876(m *base.Module, l0 int64, l1 int64, l2 int64) int64

//go:linkname Fn2890 github.com/goccy/llamawasm2go/p2.Fn2890
func Fn2890(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2904 github.com/goccy/llamawasm2go/p2.Fn2904
func Fn2904(m *base.Module, l0 int64) int64

//go:linkname Fn2907 github.com/goccy/llamawasm2go/p2.Fn2907
func Fn2907(m *base.Module)

//go:linkname Fn2916 github.com/goccy/llamawasm2go/p2.Fn2916
func Fn2916(m *base.Module, l0 int64) int32

//go:linkname Fn2928 github.com/goccy/llamawasm2go/p2.Fn2928
func Fn2928(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn2932 github.com/goccy/llamawasm2go/p2.Fn2932
func Fn2932(m *base.Module) int64

//go:linkname Fn2933 github.com/goccy/llamawasm2go/p2.Fn2933
func Fn2933(m *base.Module, l0 int64, l1 int64) float32

//go:linkname Fn2934 github.com/goccy/llamawasm2go/p2.Fn2934
func Fn2934(m *base.Module, l0 int64, l1 int64) float64

//go:linkname Fn2935 github.com/goccy/llamawasm2go/p2.Fn2935
func Fn2935(m *base.Module)

//go:linkname Fn2937 github.com/goccy/llamawasm2go/p2.Fn2937
func Fn2937(m *base.Module, l0 float64) float32

//go:linkname Fn2938 github.com/goccy/llamawasm2go/p2.Fn2938
func Fn2938(m *base.Module, l0 float64) float32

//go:linkname Fn2942 github.com/goccy/llamawasm2go/p2.Fn2942
func Fn2942(m *base.Module, l0 float64) float64

//go:linkname Fn2945 github.com/goccy/llamawasm2go/p2.Fn2945
func Fn2945(m *base.Module, l0 int32) float32

//go:linkname Fn2946 github.com/goccy/llamawasm2go/p2.Fn2946
func Fn2946(m *base.Module, l0 int32) float32

//go:linkname Fn2949 github.com/goccy/llamawasm2go/p2.Fn2949
func Fn2949(m *base.Module, l0 float32) float32

//go:linkname Fn2952 github.com/goccy/llamawasm2go/p2.Fn2952
func Fn2952(m *base.Module, l0 float64) float64

//go:linkname Fn2953 github.com/goccy/llamawasm2go/p2.Fn2953
func Fn2953(m *base.Module, l0 float64) float64

//go:linkname Fn2954 github.com/goccy/llamawasm2go/p2.Fn2954
func Fn2954(m *base.Module, l0 float32) float32

//go:linkname Fn2956 github.com/goccy/llamawasm2go/p2.Fn2956
func Fn2956(m *base.Module, l0 float32) float32

//go:linkname Fn2958 github.com/goccy/llamawasm2go/p2.Fn2958
func Fn2958(m *base.Module, l0 float32, l1 float32) float32

//go:linkname Fn2959 github.com/goccy/llamawasm2go/p2.Fn2959
func Fn2959(m *base.Module, l0 float32) float32

//go:linkname Fn2976 github.com/goccy/llamawasm2go/p2.Fn2976
func Fn2976(m *base.Module, l0 int64) int32

//go:linkname Fn2977 github.com/goccy/llamawasm2go/p2.Fn2977
func Fn2977(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn2979 github.com/goccy/llamawasm2go/p2.Fn2979
func Fn2979(m *base.Module, l0 int64)

//go:linkname Fn2980 github.com/goccy/llamawasm2go/p2.Fn2980
func Fn2980(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn2981 github.com/goccy/llamawasm2go/p2.Fn2981
func Fn2981(m *base.Module, l0 int64) int32

//go:linkname Fn2987 github.com/goccy/llamawasm2go/p2.Fn2987
func Fn2987(m *base.Module, l0 int64, l1 int64)

//go:linkname Fn2989 github.com/goccy/llamawasm2go/p2.Fn2989
func Fn2989(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int64) int32

//go:linkname Fn2995 github.com/goccy/llamawasm2go/p2.Fn2995
func Fn2995(m *base.Module, l0 int64) int32

//go:linkname Fn2998 github.com/goccy/llamawasm2go/p2.Fn2998
func Fn2998(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3001 github.com/goccy/llamawasm2go/p2.Fn3001
func Fn3001(m *base.Module, l0 int64) int32

//go:linkname Fn3003 github.com/goccy/llamawasm2go/p2.Fn3003
func Fn3003(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3005 github.com/goccy/llamawasm2go/p2.Fn3005
func Fn3005(m *base.Module, l0 int64, l1 int32, l2 int64) int64

//go:linkname Fn3006 github.com/goccy/llamawasm2go/p2.Fn3006
func Fn3006(m *base.Module, l0 int64, l1 int64, l2 int64) int32

//go:linkname Fn3009 github.com/goccy/llamawasm2go/p2.Fn3009
func Fn3009(m *base.Module, l0 int64, l1 int64) int32

//go:linkname Fn3012 github.com/goccy/llamawasm2go/p2.Fn3012
func Fn3012(m *base.Module, l0 int64) int64

//go:linkname Fn3016 github.com/goccy/llamawasm2go/p2.Fn3016
func Fn3016(m *base.Module, l0 int64, l1 int32, l2 int32) int32

//go:linkname Fn3023 github.com/goccy/llamawasm2go/p2.Fn3023
func Fn3023(m *base.Module)

//go:linkname Fn3024 github.com/goccy/llamawasm2go/p0.Fn3024
func Fn3024(m *base.Module, l0 int64, l1 int32, l2 int32) float64

//go:linkname Fn3028 github.com/goccy/llamawasm2go/p0.Fn3028
func Fn3028(m *base.Module, l0 int64) int64

//go:linkname Fn3030 github.com/goccy/llamawasm2go/p2.Fn3030
func Fn3030(m *base.Module, l0 int64, l1 int64) int64

//go:linkname Fn3034 github.com/goccy/llamawasm2go/p2.Fn3034
func Fn3034(m *base.Module, l0 int64, l1 int64, l2 int64)

//go:linkname Fn3101 github.com/goccy/llamawasm2go/p2.Fn3101
func Fn3101(m *base.Module, l0 int32)
