package wasm2go

import (
	base "github.com/goccy/llamawasm2go/base"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "github.com/goccy/llamawasm2go/p3"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env, Wasmify: wasmify}
	__memcap := reserveBytes
	if __memcap < 9437184 {
		__memcap = 9437184
	}
	if __memcap < 8589934592 {
		__memcap = 8589934592
	}
	m.Memory = make([]byte, __memcap, __memcap)
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(9437184)
	m.MemShared = true
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 8589934592
	m.T0 = make([]any, 2539)
	m.G0 = int64(8388608)
	m.G1 = int64(0)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_3_2(m)
	InitElemSeg_3_3(m)
	InitElemSeg_3_4(m)
	InitElemSeg_3_5(m)
	InitElemSeg_3_6(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:148720], wasm2goData_data_bin[148720:168841], wasm2goData_data_bin[168841:169122], wasm2goData_data_bin[169122:169293], wasm2goData_data_bin[169293:169296], wasm2goData_data_bin[169296:170329], wasm2goData_data_bin[170329:170354], wasm2goData_data_bin[170354:170379], wasm2goData_data_bin[170379:170404], wasm2goData_data_bin[170404:170429], wasm2goData_data_bin[170429:170544], wasm2goData_data_bin[170544:170547], wasm2goData_data_bin[170547:170550], wasm2goData_data_bin[170550:170665], wasm2goData_data_bin[170665:170668], wasm2goData_data_bin[170668:170671], wasm2goData_data_bin[170671:171433], wasm2goData_data_bin[171433:171715], wasm2goData_data_bin[171715:171821], wasm2goData_data_bin[171821:171863], wasm2goData_data_bin[171863:171865], wasm2goData_data_bin[171865:171867], wasm2goData_data_bin[171867:215636], wasm2goData_data_bin[215636:215926], wasm2goData_data_bin[215926:216167], wasm2goData_data_bin[216167:216169], wasm2goData_data_bin[216169:216218], wasm2goData_data_bin[216218:216275], wasm2goData_data_bin[216275:253493], wasm2goData_data_bin[253493:253527], wasm2goData_data_bin[253527:253609], wasm2goData_data_bin[253609:259506], wasm2goData_data_bin[259506:287879], wasm2goData_data_bin[287879:287883], wasm2goData_data_bin[287883:288499], wasm2goData_data_bin[288499:292432], wasm2goData_data_bin[292432:292598], wasm2goData_data_bin[292598:293415], wasm2goData_data_bin[293415:293551], wasm2goData_data_bin[293551:294460], wasm2goData_data_bin[294460:294590], wasm2goData_data_bin[294590:294720], wasm2goData_data_bin[294720:295173], wasm2goData_data_bin[295173:295315], wasm2goData_data_bin[295315:295789], wasm2goData_data_bin[295789:295921], wasm2goData_data_bin[295921:327082], wasm2goData_data_bin[327082:344005], wasm2goData_data_bin[344005:344064], wasm2goData_data_bin[344064:344123], wasm2goData_data_bin[344123:344182], wasm2goData_data_bin[344182:344241], wasm2goData_data_bin[344241:344300], wasm2goData_data_bin[344300:344359], wasm2goData_data_bin[344359:344418], wasm2goData_data_bin[344418:344519], wasm2goData_data_bin[344519:344620], wasm2goData_data_bin[344620:344721], wasm2goData_data_bin[344721:344822], wasm2goData_data_bin[344822:344923], wasm2goData_data_bin[344923:345024], wasm2goData_data_bin[345024:345125], wasm2goData_data_bin[345125:345226], wasm2goData_data_bin[345226:345327], wasm2goData_data_bin[345327:345428], wasm2goData_data_bin[345428:345530], wasm2goData_data_bin[345530:345632], wasm2goData_data_bin[345632:345734], wasm2goData_data_bin[345734:345836], wasm2goData_data_bin[345836:345938], wasm2goData_data_bin[345938:346040], wasm2goData_data_bin[346040:346142], wasm2goData_data_bin[346142:346244], wasm2goData_data_bin[346244:346346], wasm2goData_data_bin[346346:346448], wasm2goData_data_bin[346448:346550], wasm2goData_data_bin[346550:346652], wasm2goData_data_bin[346652:346754], wasm2goData_data_bin[346754:346856], wasm2goData_data_bin[346856:346958], wasm2goData_data_bin[346958:347060], wasm2goData_data_bin[347060:347162], wasm2goData_data_bin[347162:347264], wasm2goData_data_bin[347264:347366], wasm2goData_data_bin[347366:441337], wasm2goData_data_bin[441337:441457], wasm2goData_data_bin[441457:441898], wasm2goData_data_bin[441898:441901], wasm2goData_data_bin[441901:456773], wasm2goData_data_bin[456773:456775], wasm2goData_data_bin[456775:459348], wasm2goData_data_bin[459348:459381], wasm2goData_data_bin[459381:459414], wasm2goData_data_bin[459414:459456], wasm2goData_data_bin[459456:459470], wasm2goData_data_bin[459470:459503], wasm2goData_data_bin[459503:459602], wasm2goData_data_bin[459602:459684], wasm2goData_data_bin[459684:459734], wasm2goData_data_bin[459734:459736], wasm2goData_data_bin[459736:460706], wasm2goData_data_bin[460706:460844], wasm2goData_data_bin[460844:460886], wasm2goData_data_bin[460886:461216], wasm2goData_data_bin[461216:461290], wasm2goData_data_bin[461290:461309], wasm2goData_data_bin[461309:461337], wasm2goData_data_bin[461337:461371], wasm2goData_data_bin[461371:461444], wasm2goData_data_bin[461444:461452]}
	m.ThreadStart64 = Fn3509
	Fn19(m)
	return m
}

// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, wasmify base.WasmifyImports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, env, wasmify, 11796480)
}

// New constructs a *Module using DefaultWASI() for the
// wasi_snapshot_preview1 import. Use NewWithWASI to plug in a
// custom implementation (sandboxed FS, captured stdout, ...).
func New(env base.EnvImports, wasmify base.WasmifyImports) *base.Module {
	return NewWithWASI(base.DefaultWASI(), env, wasmify)
}

const InitialMemoryBytes = 9437184

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
	m.T0 = make([]any, 2539)
	m.G0 = int64(8388608)
	m.G1 = int64(0)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_3_2(m)
	InitElemSeg_3_3(m)
	InitElemSeg_3_4(m)
	InitElemSeg_3_5(m)
	InitElemSeg_3_6(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:148720], wasm2goData_data_bin[148720:168841], wasm2goData_data_bin[168841:169122], wasm2goData_data_bin[169122:169293], wasm2goData_data_bin[169293:169296], wasm2goData_data_bin[169296:170329], wasm2goData_data_bin[170329:170354], wasm2goData_data_bin[170354:170379], wasm2goData_data_bin[170379:170404], wasm2goData_data_bin[170404:170429], wasm2goData_data_bin[170429:170544], wasm2goData_data_bin[170544:170547], wasm2goData_data_bin[170547:170550], wasm2goData_data_bin[170550:170665], wasm2goData_data_bin[170665:170668], wasm2goData_data_bin[170668:170671], wasm2goData_data_bin[170671:171433], wasm2goData_data_bin[171433:171715], wasm2goData_data_bin[171715:171821], wasm2goData_data_bin[171821:171863], wasm2goData_data_bin[171863:171865], wasm2goData_data_bin[171865:171867], wasm2goData_data_bin[171867:215636], wasm2goData_data_bin[215636:215926], wasm2goData_data_bin[215926:216167], wasm2goData_data_bin[216167:216169], wasm2goData_data_bin[216169:216218], wasm2goData_data_bin[216218:216275], wasm2goData_data_bin[216275:253493], wasm2goData_data_bin[253493:253527], wasm2goData_data_bin[253527:253609], wasm2goData_data_bin[253609:259506], wasm2goData_data_bin[259506:287879], wasm2goData_data_bin[287879:287883], wasm2goData_data_bin[287883:288499], wasm2goData_data_bin[288499:292432], wasm2goData_data_bin[292432:292598], wasm2goData_data_bin[292598:293415], wasm2goData_data_bin[293415:293551], wasm2goData_data_bin[293551:294460], wasm2goData_data_bin[294460:294590], wasm2goData_data_bin[294590:294720], wasm2goData_data_bin[294720:295173], wasm2goData_data_bin[295173:295315], wasm2goData_data_bin[295315:295789], wasm2goData_data_bin[295789:295921], wasm2goData_data_bin[295921:327082], wasm2goData_data_bin[327082:344005], wasm2goData_data_bin[344005:344064], wasm2goData_data_bin[344064:344123], wasm2goData_data_bin[344123:344182], wasm2goData_data_bin[344182:344241], wasm2goData_data_bin[344241:344300], wasm2goData_data_bin[344300:344359], wasm2goData_data_bin[344359:344418], wasm2goData_data_bin[344418:344519], wasm2goData_data_bin[344519:344620], wasm2goData_data_bin[344620:344721], wasm2goData_data_bin[344721:344822], wasm2goData_data_bin[344822:344923], wasm2goData_data_bin[344923:345024], wasm2goData_data_bin[345024:345125], wasm2goData_data_bin[345125:345226], wasm2goData_data_bin[345226:345327], wasm2goData_data_bin[345327:345428], wasm2goData_data_bin[345428:345530], wasm2goData_data_bin[345530:345632], wasm2goData_data_bin[345632:345734], wasm2goData_data_bin[345734:345836], wasm2goData_data_bin[345836:345938], wasm2goData_data_bin[345938:346040], wasm2goData_data_bin[346040:346142], wasm2goData_data_bin[346142:346244], wasm2goData_data_bin[346244:346346], wasm2goData_data_bin[346346:346448], wasm2goData_data_bin[346448:346550], wasm2goData_data_bin[346550:346652], wasm2goData_data_bin[346652:346754], wasm2goData_data_bin[346754:346856], wasm2goData_data_bin[346856:346958], wasm2goData_data_bin[346958:347060], wasm2goData_data_bin[347060:347162], wasm2goData_data_bin[347162:347264], wasm2goData_data_bin[347264:347366], wasm2goData_data_bin[347366:441337], wasm2goData_data_bin[441337:441457], wasm2goData_data_bin[441457:441898], wasm2goData_data_bin[441898:441901], wasm2goData_data_bin[441901:456773], wasm2goData_data_bin[456773:456775], wasm2goData_data_bin[456775:459348], wasm2goData_data_bin[459348:459381], wasm2goData_data_bin[459381:459414], wasm2goData_data_bin[459414:459456], wasm2goData_data_bin[459456:459470], wasm2goData_data_bin[459470:459503], wasm2goData_data_bin[459503:459602], wasm2goData_data_bin[459602:459684], wasm2goData_data_bin[459684:459734], wasm2goData_data_bin[459734:459736], wasm2goData_data_bin[459736:460706], wasm2goData_data_bin[460706:460844], wasm2goData_data_bin[460844:460886], wasm2goData_data_bin[460886:461216], wasm2goData_data_bin[461216:461290], wasm2goData_data_bin[461290:461309], wasm2goData_data_bin[461309:461337], wasm2goData_data_bin[461337:461371], wasm2goData_data_bin[461371:461444], wasm2goData_data_bin[461444:461452]}
	m.ThreadStart64 = Fn3509
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
	m.T0 = make([]any, 2539)
	m.G0 = int64(8388608)
	m.G1 = int64(0)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_3_2(m)
	InitElemSeg_3_3(m)
	InitElemSeg_3_4(m)
	InitElemSeg_3_5(m)
	InitElemSeg_3_6(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:148720], wasm2goData_data_bin[148720:168841], wasm2goData_data_bin[168841:169122], wasm2goData_data_bin[169122:169293], wasm2goData_data_bin[169293:169296], wasm2goData_data_bin[169296:170329], wasm2goData_data_bin[170329:170354], wasm2goData_data_bin[170354:170379], wasm2goData_data_bin[170379:170404], wasm2goData_data_bin[170404:170429], wasm2goData_data_bin[170429:170544], wasm2goData_data_bin[170544:170547], wasm2goData_data_bin[170547:170550], wasm2goData_data_bin[170550:170665], wasm2goData_data_bin[170665:170668], wasm2goData_data_bin[170668:170671], wasm2goData_data_bin[170671:171433], wasm2goData_data_bin[171433:171715], wasm2goData_data_bin[171715:171821], wasm2goData_data_bin[171821:171863], wasm2goData_data_bin[171863:171865], wasm2goData_data_bin[171865:171867], wasm2goData_data_bin[171867:215636], wasm2goData_data_bin[215636:215926], wasm2goData_data_bin[215926:216167], wasm2goData_data_bin[216167:216169], wasm2goData_data_bin[216169:216218], wasm2goData_data_bin[216218:216275], wasm2goData_data_bin[216275:253493], wasm2goData_data_bin[253493:253527], wasm2goData_data_bin[253527:253609], wasm2goData_data_bin[253609:259506], wasm2goData_data_bin[259506:287879], wasm2goData_data_bin[287879:287883], wasm2goData_data_bin[287883:288499], wasm2goData_data_bin[288499:292432], wasm2goData_data_bin[292432:292598], wasm2goData_data_bin[292598:293415], wasm2goData_data_bin[293415:293551], wasm2goData_data_bin[293551:294460], wasm2goData_data_bin[294460:294590], wasm2goData_data_bin[294590:294720], wasm2goData_data_bin[294720:295173], wasm2goData_data_bin[295173:295315], wasm2goData_data_bin[295315:295789], wasm2goData_data_bin[295789:295921], wasm2goData_data_bin[295921:327082], wasm2goData_data_bin[327082:344005], wasm2goData_data_bin[344005:344064], wasm2goData_data_bin[344064:344123], wasm2goData_data_bin[344123:344182], wasm2goData_data_bin[344182:344241], wasm2goData_data_bin[344241:344300], wasm2goData_data_bin[344300:344359], wasm2goData_data_bin[344359:344418], wasm2goData_data_bin[344418:344519], wasm2goData_data_bin[344519:344620], wasm2goData_data_bin[344620:344721], wasm2goData_data_bin[344721:344822], wasm2goData_data_bin[344822:344923], wasm2goData_data_bin[344923:345024], wasm2goData_data_bin[345024:345125], wasm2goData_data_bin[345125:345226], wasm2goData_data_bin[345226:345327], wasm2goData_data_bin[345327:345428], wasm2goData_data_bin[345428:345530], wasm2goData_data_bin[345530:345632], wasm2goData_data_bin[345632:345734], wasm2goData_data_bin[345734:345836], wasm2goData_data_bin[345836:345938], wasm2goData_data_bin[345938:346040], wasm2goData_data_bin[346040:346142], wasm2goData_data_bin[346142:346244], wasm2goData_data_bin[346244:346346], wasm2goData_data_bin[346346:346448], wasm2goData_data_bin[346448:346550], wasm2goData_data_bin[346550:346652], wasm2goData_data_bin[346652:346754], wasm2goData_data_bin[346754:346856], wasm2goData_data_bin[346856:346958], wasm2goData_data_bin[346958:347060], wasm2goData_data_bin[347060:347162], wasm2goData_data_bin[347162:347264], wasm2goData_data_bin[347264:347366], wasm2goData_data_bin[347366:441337], wasm2goData_data_bin[441337:441457], wasm2goData_data_bin[441457:441898], wasm2goData_data_bin[441898:441901], wasm2goData_data_bin[441901:456773], wasm2goData_data_bin[456773:456775], wasm2goData_data_bin[456775:459348], wasm2goData_data_bin[459348:459381], wasm2goData_data_bin[459381:459414], wasm2goData_data_bin[459414:459456], wasm2goData_data_bin[459456:459470], wasm2goData_data_bin[459470:459503], wasm2goData_data_bin[459503:459602], wasm2goData_data_bin[459602:459684], wasm2goData_data_bin[459684:459734], wasm2goData_data_bin[459734:459736], wasm2goData_data_bin[459736:460706], wasm2goData_data_bin[460706:460844], wasm2goData_data_bin[460844:460886], wasm2goData_data_bin[460886:461216], wasm2goData_data_bin[461216:461290], wasm2goData_data_bin[461290:461309], wasm2goData_data_bin[461309:461337], wasm2goData_data_bin[461337:461371], wasm2goData_data_bin[461371:461444], wasm2goData_data_bin[461444:461452]}
	m.ThreadStart64 = Fn3509
	base.RestoreGlobals(m, globals)
	return m
}
func Initialize(m *base.Module) {
	Fn20(m)
}
func WasmAlloc(m *base.Module, l0 int64) int64 {
	return Fn290(m, l0)
}
func WasmFree(m *base.Module, l0 int64) {
	Fn26(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int64, l1 int64) int64 {
	return Fn336(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn337(m)
}
func WasmShutdown(m *base.Module) {
	Fn338(m)
}
func DbgKernelInit(m *base.Module) {
	Fn361(m)
}
func DbgCpuFp16ToFp32(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn860(m, l0, l1, l2)
}
func DbgCpuFp32ToFp16(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn445(m, l0, l1, l2)
}
func DbgVecDotF16(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1011(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1040(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4_1_q8_1(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1041(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1042(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_1_q8_1(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1043(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ8_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1044(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ2KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1045(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ3KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1046(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1047(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1048(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ6KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1049(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq2XxsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn893(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq2XsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn894(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq3XxsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn896(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq1SQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn898(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq4NlQ8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn900(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq3SQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn897(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq2SQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn895(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq4XsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn901(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq1MQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn899(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotTq1_0_q8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn891(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotTq2_0_q8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn892(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotMxfp4Q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn889(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotNvfp4Q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn890(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ1_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn887(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ2_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn888(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgQuantizeMatQ8K4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn905(m, l0, l1, l2)
}
func DbgGemvQ4_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn908(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn910(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ5K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn913(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ6K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn915(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvIq4Nl8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn917(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvMxfp48x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn919(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ4_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn923(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn924(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ5K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn925(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ6K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn926(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmIq4Nl8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn927(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmMxfp48x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn928(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgQuantizeMatQ8_0_4x4(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn1050(m, l0, l1, l2)
}
func DbgQuantizeMatQ8_0_4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn1051(m, l0, l1, l2)
}
func DbgGemvQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1054(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1052(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1055(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1053(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgVecGegluErfF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	Fn1012(m, l0, l1, l2, l3)
}
func DbgVecSwigluF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	Fn1013(m, l0, l1, l2, l3)
}
func DbgVecCvarF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) float64 {
	return Fn1014(m, l0, l1, l2, l3)
}
func DbgVecSumF32(m *base.Module, l0 int32, l1 int64) float64 {
	return Fn1015(m, l0, l1)
}
func DbgVecSoftMaxF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) float64 {
	return Fn1016(m, l0, l1, l2, l3)
}
func DbgVecMadF16F32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) {
	Fn1017(m, l0, l1, l2, l3)
}
func DbgSimdGemmF32(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) {
	Fn1018(m, l0, l1, l2, l3, l4, l5)
}
func DbgFlashAttnKvF16(m *base.Module, l0 int64) {
	Fn1027(m, l0)
}
func WasiThreadStart(m *base.Module, l0 int32, l1 int64) {
	Fn3509(m, l0, l1)
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
	packed = Fn291(m, l0, l1)
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
	packed = Fn294(m, l0, l1)
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
	packed = Fn295(m, l0, l1)
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
	packed = Fn296(m, l0, l1)
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
	packed = Fn297(m, l0, l1)
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
	packed = Fn298(m, l0, l1)
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
	packed = Fn299(m, l0, l1)
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
	packed = Fn300(m, l0, l1)
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
	packed = Fn301(m, l0, l1)
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
	packed = Fn303(m, l0, l1)
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
	packed = Fn304(m, l0, l1)
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
	packed = Fn306(m, l0, l1)
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
	packed = Fn307(m, l0, l1)
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
	packed = Fn308(m, l0, l1)
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
	packed = Fn309(m, l0, l1)
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
	packed = Fn310(m, l0, l1)
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
	packed = Fn311(m, l0, l1)
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
	packed = Fn312(m, l0, l1)
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
	packed = Fn313(m, l0, l1)
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
	packed = Fn314(m, l0, l1)
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
	packed = Fn315(m, l0, l1)
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
	packed = Fn316(m, l0, l1)
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
	packed = Fn317(m, l0, l1)
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
	packed = Fn318(m, l0, l1)
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
	packed = Fn319(m, l0, l1)
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
	packed = Fn320(m, l0, l1)
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
	packed = Fn321(m, l0, l1)
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
	packed = Fn322(m, l0, l1)
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
	packed = Fn323(m, l0, l1)
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
	packed = Fn324(m, l0, l1)
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
	packed = Fn325(m, l0, l1)
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
	packed = Fn326(m, l0, l1)
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
	packed = Fn327(m, l0, l1)
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
	packed = Fn328(m, l0, l1)
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
	packed = Fn329(m, l0, l1)
	return
}
func Inv_0_35(m *base.Module, l0, l1 int64) (packed int64, err error) {
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
func Inv_0_36(m *base.Module, l0, l1 int64) (packed int64, err error) {
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
	packed = Fn332(m, l0, l1)
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
	packed = Fn334(m, l0, l1)
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
	packed = Fn335(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
