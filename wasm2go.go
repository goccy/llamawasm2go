package wasm2go

import (
	base "github.com/goccy/llamawasm2go/base"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "github.com/goccy/llamawasm2go/p2"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	__memcap := reserveBytes
	if __memcap < 9371648 {
		__memcap = 9371648
	}
	if __memcap < 8589934592 {
		__memcap = 8589934592
	}
	m.Memory = make([]byte, __memcap, __memcap)
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(9371648)
	m.MemShared = true
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 8589934592
	m.T0 = make([]any, 2243)
	m.G0 = int64(8388608)
	m.G1 = int64(0)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:124208], wasm2goData_data_bin[124208:144313], wasm2goData_data_bin[144313:144594], wasm2goData_data_bin[144594:144765], wasm2goData_data_bin[144765:144768], wasm2goData_data_bin[144768:145801], wasm2goData_data_bin[145801:145826], wasm2goData_data_bin[145826:145851], wasm2goData_data_bin[145851:145876], wasm2goData_data_bin[145876:145901], wasm2goData_data_bin[145901:146016], wasm2goData_data_bin[146016:146019], wasm2goData_data_bin[146019:146022], wasm2goData_data_bin[146022:146137], wasm2goData_data_bin[146137:146140], wasm2goData_data_bin[146140:146143], wasm2goData_data_bin[146143:146417], wasm2goData_data_bin[146417:146691], wasm2goData_data_bin[146691:146781], wasm2goData_data_bin[146781:146823], wasm2goData_data_bin[146823:146825], wasm2goData_data_bin[146825:146827], wasm2goData_data_bin[146827:190684], wasm2goData_data_bin[190684:190974], wasm2goData_data_bin[190974:191215], wasm2goData_data_bin[191215:191217], wasm2goData_data_bin[191217:191266], wasm2goData_data_bin[191266:191323], wasm2goData_data_bin[191323:228237], wasm2goData_data_bin[228237:228271], wasm2goData_data_bin[228271:228353], wasm2goData_data_bin[228353:234250], wasm2goData_data_bin[234250:259351], wasm2goData_data_bin[259351:259959], wasm2goData_data_bin[259959:262529], wasm2goData_data_bin[262529:263702], wasm2goData_data_bin[263702:263868], wasm2goData_data_bin[263868:264801], wasm2goData_data_bin[264801:264930], wasm2goData_data_bin[264930:265059], wasm2goData_data_bin[265059:265512], wasm2goData_data_bin[265512:265654], wasm2goData_data_bin[265654:266127], wasm2goData_data_bin[266127:266259], wasm2goData_data_bin[266259:273664], wasm2goData_data_bin[273664:273723], wasm2goData_data_bin[273723:273782], wasm2goData_data_bin[273782:273841], wasm2goData_data_bin[273841:273900], wasm2goData_data_bin[273900:273958], wasm2goData_data_bin[273958:274017], wasm2goData_data_bin[274017:274076], wasm2goData_data_bin[274076:274177], wasm2goData_data_bin[274177:274278], wasm2goData_data_bin[274278:274379], wasm2goData_data_bin[274379:274480], wasm2goData_data_bin[274480:274581], wasm2goData_data_bin[274581:274682], wasm2goData_data_bin[274682:274783], wasm2goData_data_bin[274783:274884], wasm2goData_data_bin[274884:274985], wasm2goData_data_bin[274985:275086], wasm2goData_data_bin[275086:275188], wasm2goData_data_bin[275188:275290], wasm2goData_data_bin[275290:355287], wasm2goData_data_bin[355287:379233], wasm2goData_data_bin[379233:379353], wasm2goData_data_bin[379353:379794], wasm2goData_data_bin[379794:379797], wasm2goData_data_bin[379797:394669], wasm2goData_data_bin[394669:394671], wasm2goData_data_bin[394671:397244], wasm2goData_data_bin[397244:397277], wasm2goData_data_bin[397277:397310], wasm2goData_data_bin[397310:397352], wasm2goData_data_bin[397352:397366], wasm2goData_data_bin[397366:397399], wasm2goData_data_bin[397399:397498], wasm2goData_data_bin[397498:397660], wasm2goData_data_bin[397660:398518], wasm2goData_data_bin[398518:398640], wasm2goData_data_bin[398640:398682], wasm2goData_data_bin[398682:398724], wasm2goData_data_bin[398724:398926], wasm2goData_data_bin[398926:398992], wasm2goData_data_bin[398992:399011], wasm2goData_data_bin[399011:399039], wasm2goData_data_bin[399039:399073], wasm2goData_data_bin[399073:399146], wasm2goData_data_bin[399146:399154]}
	m.ThreadStart64 = Fn3080
	Fn19(m)
	return m
}

// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, env, wasmify, 11714560)
}

// New constructs a *Module using DefaultWASI() for the
// wasi_snapshot_preview1 import. Use NewWithWASI to plug in a
// custom implementation (sandboxed FS, captured stdout, ...).
func New(env base.EnvImports, wasmify base.WasmifyImports) *base.Module {
	return NewWithWASI(base.DefaultWASI(), env, wasmify)
}

const InitialMemoryBytes = 9371648

func NewWithMemory(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, memory []byte, memSize uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 281474976710656 {
		panic("wasm2go: memory size exceeds the implementation limit (281474976710656 bytes)")
	}
	m.MemSize.Store(memSize)
	m.MemShared = true
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 2243)
	m.G0 = int64(8388608)
	m.G1 = int64(0)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:124208], wasm2goData_data_bin[124208:144313], wasm2goData_data_bin[144313:144594], wasm2goData_data_bin[144594:144765], wasm2goData_data_bin[144765:144768], wasm2goData_data_bin[144768:145801], wasm2goData_data_bin[145801:145826], wasm2goData_data_bin[145826:145851], wasm2goData_data_bin[145851:145876], wasm2goData_data_bin[145876:145901], wasm2goData_data_bin[145901:146016], wasm2goData_data_bin[146016:146019], wasm2goData_data_bin[146019:146022], wasm2goData_data_bin[146022:146137], wasm2goData_data_bin[146137:146140], wasm2goData_data_bin[146140:146143], wasm2goData_data_bin[146143:146417], wasm2goData_data_bin[146417:146691], wasm2goData_data_bin[146691:146781], wasm2goData_data_bin[146781:146823], wasm2goData_data_bin[146823:146825], wasm2goData_data_bin[146825:146827], wasm2goData_data_bin[146827:190684], wasm2goData_data_bin[190684:190974], wasm2goData_data_bin[190974:191215], wasm2goData_data_bin[191215:191217], wasm2goData_data_bin[191217:191266], wasm2goData_data_bin[191266:191323], wasm2goData_data_bin[191323:228237], wasm2goData_data_bin[228237:228271], wasm2goData_data_bin[228271:228353], wasm2goData_data_bin[228353:234250], wasm2goData_data_bin[234250:259351], wasm2goData_data_bin[259351:259959], wasm2goData_data_bin[259959:262529], wasm2goData_data_bin[262529:263702], wasm2goData_data_bin[263702:263868], wasm2goData_data_bin[263868:264801], wasm2goData_data_bin[264801:264930], wasm2goData_data_bin[264930:265059], wasm2goData_data_bin[265059:265512], wasm2goData_data_bin[265512:265654], wasm2goData_data_bin[265654:266127], wasm2goData_data_bin[266127:266259], wasm2goData_data_bin[266259:273664], wasm2goData_data_bin[273664:273723], wasm2goData_data_bin[273723:273782], wasm2goData_data_bin[273782:273841], wasm2goData_data_bin[273841:273900], wasm2goData_data_bin[273900:273958], wasm2goData_data_bin[273958:274017], wasm2goData_data_bin[274017:274076], wasm2goData_data_bin[274076:274177], wasm2goData_data_bin[274177:274278], wasm2goData_data_bin[274278:274379], wasm2goData_data_bin[274379:274480], wasm2goData_data_bin[274480:274581], wasm2goData_data_bin[274581:274682], wasm2goData_data_bin[274682:274783], wasm2goData_data_bin[274783:274884], wasm2goData_data_bin[274884:274985], wasm2goData_data_bin[274985:275086], wasm2goData_data_bin[275086:275188], wasm2goData_data_bin[275188:275290], wasm2goData_data_bin[275290:355287], wasm2goData_data_bin[355287:379233], wasm2goData_data_bin[379233:379353], wasm2goData_data_bin[379353:379794], wasm2goData_data_bin[379794:379797], wasm2goData_data_bin[379797:394669], wasm2goData_data_bin[394669:394671], wasm2goData_data_bin[394671:397244], wasm2goData_data_bin[397244:397277], wasm2goData_data_bin[397277:397310], wasm2goData_data_bin[397310:397352], wasm2goData_data_bin[397352:397366], wasm2goData_data_bin[397366:397399], wasm2goData_data_bin[397399:397498], wasm2goData_data_bin[397498:397660], wasm2goData_data_bin[397660:398518], wasm2goData_data_bin[398518:398640], wasm2goData_data_bin[398640:398682], wasm2goData_data_bin[398682:398724], wasm2goData_data_bin[398724:398926], wasm2goData_data_bin[398926:398992], wasm2goData_data_bin[398992:399011], wasm2goData_data_bin[399011:399039], wasm2goData_data_bin[399039:399073], wasm2goData_data_bin[399073:399146], wasm2goData_data_bin[399146:399154]}
	m.ThreadStart64 = Fn3080
	Fn19(m)
	return m
}
func NewFromSnapshot(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, memory []byte, memSize uint64, globals []uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 281474976710656 {
		panic("wasm2go: memory size exceeds the implementation limit (281474976710656 bytes)")
	}
	m.MemSize.Store(memSize)
	m.MemShared = true
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 2243)
	m.G0 = int64(8388608)
	m.G1 = int64(0)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_1_1(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:124208], wasm2goData_data_bin[124208:144313], wasm2goData_data_bin[144313:144594], wasm2goData_data_bin[144594:144765], wasm2goData_data_bin[144765:144768], wasm2goData_data_bin[144768:145801], wasm2goData_data_bin[145801:145826], wasm2goData_data_bin[145826:145851], wasm2goData_data_bin[145851:145876], wasm2goData_data_bin[145876:145901], wasm2goData_data_bin[145901:146016], wasm2goData_data_bin[146016:146019], wasm2goData_data_bin[146019:146022], wasm2goData_data_bin[146022:146137], wasm2goData_data_bin[146137:146140], wasm2goData_data_bin[146140:146143], wasm2goData_data_bin[146143:146417], wasm2goData_data_bin[146417:146691], wasm2goData_data_bin[146691:146781], wasm2goData_data_bin[146781:146823], wasm2goData_data_bin[146823:146825], wasm2goData_data_bin[146825:146827], wasm2goData_data_bin[146827:190684], wasm2goData_data_bin[190684:190974], wasm2goData_data_bin[190974:191215], wasm2goData_data_bin[191215:191217], wasm2goData_data_bin[191217:191266], wasm2goData_data_bin[191266:191323], wasm2goData_data_bin[191323:228237], wasm2goData_data_bin[228237:228271], wasm2goData_data_bin[228271:228353], wasm2goData_data_bin[228353:234250], wasm2goData_data_bin[234250:259351], wasm2goData_data_bin[259351:259959], wasm2goData_data_bin[259959:262529], wasm2goData_data_bin[262529:263702], wasm2goData_data_bin[263702:263868], wasm2goData_data_bin[263868:264801], wasm2goData_data_bin[264801:264930], wasm2goData_data_bin[264930:265059], wasm2goData_data_bin[265059:265512], wasm2goData_data_bin[265512:265654], wasm2goData_data_bin[265654:266127], wasm2goData_data_bin[266127:266259], wasm2goData_data_bin[266259:273664], wasm2goData_data_bin[273664:273723], wasm2goData_data_bin[273723:273782], wasm2goData_data_bin[273782:273841], wasm2goData_data_bin[273841:273900], wasm2goData_data_bin[273900:273958], wasm2goData_data_bin[273958:274017], wasm2goData_data_bin[274017:274076], wasm2goData_data_bin[274076:274177], wasm2goData_data_bin[274177:274278], wasm2goData_data_bin[274278:274379], wasm2goData_data_bin[274379:274480], wasm2goData_data_bin[274480:274581], wasm2goData_data_bin[274581:274682], wasm2goData_data_bin[274682:274783], wasm2goData_data_bin[274783:274884], wasm2goData_data_bin[274884:274985], wasm2goData_data_bin[274985:275086], wasm2goData_data_bin[275086:275188], wasm2goData_data_bin[275188:275290], wasm2goData_data_bin[275290:355287], wasm2goData_data_bin[355287:379233], wasm2goData_data_bin[379233:379353], wasm2goData_data_bin[379353:379794], wasm2goData_data_bin[379794:379797], wasm2goData_data_bin[379797:394669], wasm2goData_data_bin[394669:394671], wasm2goData_data_bin[394671:397244], wasm2goData_data_bin[397244:397277], wasm2goData_data_bin[397277:397310], wasm2goData_data_bin[397310:397352], wasm2goData_data_bin[397352:397366], wasm2goData_data_bin[397366:397399], wasm2goData_data_bin[397399:397498], wasm2goData_data_bin[397498:397660], wasm2goData_data_bin[397660:398518], wasm2goData_data_bin[398518:398640], wasm2goData_data_bin[398640:398682], wasm2goData_data_bin[398682:398724], wasm2goData_data_bin[398724:398926], wasm2goData_data_bin[398926:398992], wasm2goData_data_bin[398992:399011], wasm2goData_data_bin[399011:399039], wasm2goData_data_bin[399039:399073], wasm2goData_data_bin[399073:399146], wasm2goData_data_bin[399146:399154]}
	m.ThreadStart64 = Fn3080
	base.RestoreGlobals(m, globals)
	return m
}
func Initialize(m *base.Module) {
	Fn20(m)
}
func WasmAlloc(m *base.Module, l0 int64) int64 {
	return Fn291(m, l0)
}
func WasmFree(m *base.Module, l0 int64) {
	Fn26(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int64, l1 int64) int64 {
	return Fn335(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn336(m)
}
func WasmShutdown(m *base.Module) {
	Fn337(m)
}
func DbgKernelInit(m *base.Module) {
	Fn360(m)
}
func DbgVecDotF16(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn967(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn993(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4_1_q8_1(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn994(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn995(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_1_q8_1(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn996(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ8_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn997(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ2KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn998(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ3KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn999(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1000(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1001(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ6KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1002(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq2XxsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn855(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq2XsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn856(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq3XxsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn858(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq1SQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn860(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq4NlQ8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn862(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq3SQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn859(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq2SQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn857(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq4XsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn863(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq1MQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn861(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotTq1_0_q8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn853(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotTq2_0_q8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn854(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotMxfp4Q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn851(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotNvfp4Q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn852(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ1_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn849(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ2_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn850(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgQuantizeMatQ8K4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn867(m, l0, l1, l2)
}
func DbgGemvQ4_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn870(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn872(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ5K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn875(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ6K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn877(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvIq4Nl8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn879(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvMxfp48x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn881(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ4_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn883(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn884(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ5K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn885(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ6K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn886(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmIq4Nl8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn887(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmMxfp48x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn888(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgQuantizeMatQ8_0_4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn1004(m, l0, l1, l2)
}
func DbgGemvQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1007(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1005(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1008(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1006(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgVecSwigluF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	Fn968(m, l0, l1, l2, l3)
}
func DbgVecSoftMaxF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) float64 {
	return Fn969(m, l0, l1, l2, l3)
}
func DbgVecMadF16F32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) {
	Fn970(m, l0, l1, l2, l3)
}
func DbgSimdGemmF32(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) {
	Fn971(m, l0, l1, l2, l3, l4, l5)
}
func DbgFlashAttnKvF16(m *base.Module, l0 int64) {
	Fn980(m, l0)
}
func WasiThreadStart(m *base.Module, l0 int32, l1 int64) {
	Fn3080(m, l0, l1)
}
func Inv_0_0(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn292(m, l0, l1)
	return
}
func Inv_0_1(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn295(m, l0, l1)
	return
}
func Inv_0_2(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn296(m, l0, l1)
	return
}
func Inv_0_3(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn297(m, l0, l1)
	return
}
func Inv_0_4(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn298(m, l0, l1)
	return
}
func Inv_0_5(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn299(m, l0, l1)
	return
}
func Inv_0_6(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn300(m, l0, l1)
	return
}
func Inv_0_7(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn302(m, l0, l1)
	return
}
func Inv_0_8(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn303(m, l0, l1)
	return
}
func Inv_0_9(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn305(m, l0, l1)
	return
}
func Inv_0_10(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn306(m, l0, l1)
	return
}
func Inv_0_11(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn307(m, l0, l1)
	return
}
func Inv_0_12(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn308(m, l0, l1)
	return
}
func Inv_0_13(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn309(m, l0, l1)
	return
}
func Inv_0_14(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn310(m, l0, l1)
	return
}
func Inv_0_15(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn311(m, l0, l1)
	return
}
func Inv_0_16(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn312(m, l0, l1)
	return
}
func Inv_0_17(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn313(m, l0, l1)
	return
}
func Inv_0_18(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn314(m, l0, l1)
	return
}
func Inv_0_19(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn315(m, l0, l1)
	return
}
func Inv_0_20(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn316(m, l0, l1)
	return
}
func Inv_0_21(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn317(m, l0, l1)
	return
}
func Inv_0_22(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn318(m, l0, l1)
	return
}
func Inv_0_23(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn319(m, l0, l1)
	return
}
func Inv_0_24(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn320(m, l0, l1)
	return
}
func Inv_0_25(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn321(m, l0, l1)
	return
}
func Inv_0_26(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn322(m, l0, l1)
	return
}
func Inv_0_27(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn323(m, l0, l1)
	return
}
func Inv_0_28(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn324(m, l0, l1)
	return
}
func Inv_0_29(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn325(m, l0, l1)
	return
}
func Inv_0_30(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn326(m, l0, l1)
	return
}
func Inv_0_31(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn327(m, l0, l1)
	return
}
func Inv_0_32(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn328(m, l0, l1)
	return
}
func Inv_0_33(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn329(m, l0, l1)
	return
}
func Inv_0_34(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn330(m, l0, l1)
	return
}
func Inv_1_0(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn331(m, l0, l1)
	return
}
func Inv_1_1(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn333(m, l0, l1)
	return
}
func Inv_1_2(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn334(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
