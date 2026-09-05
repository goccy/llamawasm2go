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
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:122496], wasm2goData_data_bin[122496:142489], wasm2goData_data_bin[142489:142770], wasm2goData_data_bin[142770:142941], wasm2goData_data_bin[142941:142944], wasm2goData_data_bin[142944:143977], wasm2goData_data_bin[143977:144002], wasm2goData_data_bin[144002:144027], wasm2goData_data_bin[144027:144052], wasm2goData_data_bin[144052:144077], wasm2goData_data_bin[144077:144192], wasm2goData_data_bin[144192:144195], wasm2goData_data_bin[144195:144198], wasm2goData_data_bin[144198:144313], wasm2goData_data_bin[144313:144316], wasm2goData_data_bin[144316:144319], wasm2goData_data_bin[144319:144593], wasm2goData_data_bin[144593:144867], wasm2goData_data_bin[144867:144957], wasm2goData_data_bin[144957:144999], wasm2goData_data_bin[144999:145001], wasm2goData_data_bin[145001:145003], wasm2goData_data_bin[145003:188860], wasm2goData_data_bin[188860:189150], wasm2goData_data_bin[189150:189391], wasm2goData_data_bin[189391:189393], wasm2goData_data_bin[189393:189442], wasm2goData_data_bin[189442:189499], wasm2goData_data_bin[189499:226413], wasm2goData_data_bin[226413:226447], wasm2goData_data_bin[226447:226529], wasm2goData_data_bin[226529:232425], wasm2goData_data_bin[232425:257526], wasm2goData_data_bin[257526:258134], wasm2goData_data_bin[258134:260704], wasm2goData_data_bin[260704:261877], wasm2goData_data_bin[261877:262043], wasm2goData_data_bin[262043:262976], wasm2goData_data_bin[262976:263105], wasm2goData_data_bin[263105:263234], wasm2goData_data_bin[263234:263687], wasm2goData_data_bin[263687:263829], wasm2goData_data_bin[263829:264303], wasm2goData_data_bin[264303:264435], wasm2goData_data_bin[264435:271840], wasm2goData_data_bin[271840:271899], wasm2goData_data_bin[271899:271958], wasm2goData_data_bin[271958:272017], wasm2goData_data_bin[272017:272076], wasm2goData_data_bin[272076:272135], wasm2goData_data_bin[272135:272194], wasm2goData_data_bin[272194:272253], wasm2goData_data_bin[272253:272354], wasm2goData_data_bin[272354:272455], wasm2goData_data_bin[272455:272556], wasm2goData_data_bin[272556:272657], wasm2goData_data_bin[272657:272758], wasm2goData_data_bin[272758:272859], wasm2goData_data_bin[272859:272960], wasm2goData_data_bin[272960:273061], wasm2goData_data_bin[273061:273162], wasm2goData_data_bin[273162:273263], wasm2goData_data_bin[273263:273365], wasm2goData_data_bin[273365:273467], wasm2goData_data_bin[273467:353464], wasm2goData_data_bin[353464:377411], wasm2goData_data_bin[377411:377531], wasm2goData_data_bin[377531:377972], wasm2goData_data_bin[377972:377975], wasm2goData_data_bin[377975:392847], wasm2goData_data_bin[392847:392849], wasm2goData_data_bin[392849:395422], wasm2goData_data_bin[395422:395455], wasm2goData_data_bin[395455:395488], wasm2goData_data_bin[395488:395530], wasm2goData_data_bin[395530:395544], wasm2goData_data_bin[395544:395577], wasm2goData_data_bin[395577:395676], wasm2goData_data_bin[395676:395838], wasm2goData_data_bin[395838:396696], wasm2goData_data_bin[396696:396818], wasm2goData_data_bin[396818:396860], wasm2goData_data_bin[396860:396902], wasm2goData_data_bin[396902:397104], wasm2goData_data_bin[397104:397170], wasm2goData_data_bin[397170:397189], wasm2goData_data_bin[397189:397217], wasm2goData_data_bin[397217:397251], wasm2goData_data_bin[397251:397324], wasm2goData_data_bin[397324:397332]}
	m.ThreadStart64 = Fn3059
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
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:122496], wasm2goData_data_bin[122496:142489], wasm2goData_data_bin[142489:142770], wasm2goData_data_bin[142770:142941], wasm2goData_data_bin[142941:142944], wasm2goData_data_bin[142944:143977], wasm2goData_data_bin[143977:144002], wasm2goData_data_bin[144002:144027], wasm2goData_data_bin[144027:144052], wasm2goData_data_bin[144052:144077], wasm2goData_data_bin[144077:144192], wasm2goData_data_bin[144192:144195], wasm2goData_data_bin[144195:144198], wasm2goData_data_bin[144198:144313], wasm2goData_data_bin[144313:144316], wasm2goData_data_bin[144316:144319], wasm2goData_data_bin[144319:144593], wasm2goData_data_bin[144593:144867], wasm2goData_data_bin[144867:144957], wasm2goData_data_bin[144957:144999], wasm2goData_data_bin[144999:145001], wasm2goData_data_bin[145001:145003], wasm2goData_data_bin[145003:188860], wasm2goData_data_bin[188860:189150], wasm2goData_data_bin[189150:189391], wasm2goData_data_bin[189391:189393], wasm2goData_data_bin[189393:189442], wasm2goData_data_bin[189442:189499], wasm2goData_data_bin[189499:226413], wasm2goData_data_bin[226413:226447], wasm2goData_data_bin[226447:226529], wasm2goData_data_bin[226529:232425], wasm2goData_data_bin[232425:257526], wasm2goData_data_bin[257526:258134], wasm2goData_data_bin[258134:260704], wasm2goData_data_bin[260704:261877], wasm2goData_data_bin[261877:262043], wasm2goData_data_bin[262043:262976], wasm2goData_data_bin[262976:263105], wasm2goData_data_bin[263105:263234], wasm2goData_data_bin[263234:263687], wasm2goData_data_bin[263687:263829], wasm2goData_data_bin[263829:264303], wasm2goData_data_bin[264303:264435], wasm2goData_data_bin[264435:271840], wasm2goData_data_bin[271840:271899], wasm2goData_data_bin[271899:271958], wasm2goData_data_bin[271958:272017], wasm2goData_data_bin[272017:272076], wasm2goData_data_bin[272076:272135], wasm2goData_data_bin[272135:272194], wasm2goData_data_bin[272194:272253], wasm2goData_data_bin[272253:272354], wasm2goData_data_bin[272354:272455], wasm2goData_data_bin[272455:272556], wasm2goData_data_bin[272556:272657], wasm2goData_data_bin[272657:272758], wasm2goData_data_bin[272758:272859], wasm2goData_data_bin[272859:272960], wasm2goData_data_bin[272960:273061], wasm2goData_data_bin[273061:273162], wasm2goData_data_bin[273162:273263], wasm2goData_data_bin[273263:273365], wasm2goData_data_bin[273365:273467], wasm2goData_data_bin[273467:353464], wasm2goData_data_bin[353464:377411], wasm2goData_data_bin[377411:377531], wasm2goData_data_bin[377531:377972], wasm2goData_data_bin[377972:377975], wasm2goData_data_bin[377975:392847], wasm2goData_data_bin[392847:392849], wasm2goData_data_bin[392849:395422], wasm2goData_data_bin[395422:395455], wasm2goData_data_bin[395455:395488], wasm2goData_data_bin[395488:395530], wasm2goData_data_bin[395530:395544], wasm2goData_data_bin[395544:395577], wasm2goData_data_bin[395577:395676], wasm2goData_data_bin[395676:395838], wasm2goData_data_bin[395838:396696], wasm2goData_data_bin[396696:396818], wasm2goData_data_bin[396818:396860], wasm2goData_data_bin[396860:396902], wasm2goData_data_bin[396902:397104], wasm2goData_data_bin[397104:397170], wasm2goData_data_bin[397170:397189], wasm2goData_data_bin[397189:397217], wasm2goData_data_bin[397217:397251], wasm2goData_data_bin[397251:397324], wasm2goData_data_bin[397324:397332]}
	m.ThreadStart64 = Fn3059
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
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:122496], wasm2goData_data_bin[122496:142489], wasm2goData_data_bin[142489:142770], wasm2goData_data_bin[142770:142941], wasm2goData_data_bin[142941:142944], wasm2goData_data_bin[142944:143977], wasm2goData_data_bin[143977:144002], wasm2goData_data_bin[144002:144027], wasm2goData_data_bin[144027:144052], wasm2goData_data_bin[144052:144077], wasm2goData_data_bin[144077:144192], wasm2goData_data_bin[144192:144195], wasm2goData_data_bin[144195:144198], wasm2goData_data_bin[144198:144313], wasm2goData_data_bin[144313:144316], wasm2goData_data_bin[144316:144319], wasm2goData_data_bin[144319:144593], wasm2goData_data_bin[144593:144867], wasm2goData_data_bin[144867:144957], wasm2goData_data_bin[144957:144999], wasm2goData_data_bin[144999:145001], wasm2goData_data_bin[145001:145003], wasm2goData_data_bin[145003:188860], wasm2goData_data_bin[188860:189150], wasm2goData_data_bin[189150:189391], wasm2goData_data_bin[189391:189393], wasm2goData_data_bin[189393:189442], wasm2goData_data_bin[189442:189499], wasm2goData_data_bin[189499:226413], wasm2goData_data_bin[226413:226447], wasm2goData_data_bin[226447:226529], wasm2goData_data_bin[226529:232425], wasm2goData_data_bin[232425:257526], wasm2goData_data_bin[257526:258134], wasm2goData_data_bin[258134:260704], wasm2goData_data_bin[260704:261877], wasm2goData_data_bin[261877:262043], wasm2goData_data_bin[262043:262976], wasm2goData_data_bin[262976:263105], wasm2goData_data_bin[263105:263234], wasm2goData_data_bin[263234:263687], wasm2goData_data_bin[263687:263829], wasm2goData_data_bin[263829:264303], wasm2goData_data_bin[264303:264435], wasm2goData_data_bin[264435:271840], wasm2goData_data_bin[271840:271899], wasm2goData_data_bin[271899:271958], wasm2goData_data_bin[271958:272017], wasm2goData_data_bin[272017:272076], wasm2goData_data_bin[272076:272135], wasm2goData_data_bin[272135:272194], wasm2goData_data_bin[272194:272253], wasm2goData_data_bin[272253:272354], wasm2goData_data_bin[272354:272455], wasm2goData_data_bin[272455:272556], wasm2goData_data_bin[272556:272657], wasm2goData_data_bin[272657:272758], wasm2goData_data_bin[272758:272859], wasm2goData_data_bin[272859:272960], wasm2goData_data_bin[272960:273061], wasm2goData_data_bin[273061:273162], wasm2goData_data_bin[273162:273263], wasm2goData_data_bin[273263:273365], wasm2goData_data_bin[273365:273467], wasm2goData_data_bin[273467:353464], wasm2goData_data_bin[353464:377411], wasm2goData_data_bin[377411:377531], wasm2goData_data_bin[377531:377972], wasm2goData_data_bin[377972:377975], wasm2goData_data_bin[377975:392847], wasm2goData_data_bin[392847:392849], wasm2goData_data_bin[392849:395422], wasm2goData_data_bin[395422:395455], wasm2goData_data_bin[395455:395488], wasm2goData_data_bin[395488:395530], wasm2goData_data_bin[395530:395544], wasm2goData_data_bin[395544:395577], wasm2goData_data_bin[395577:395676], wasm2goData_data_bin[395676:395838], wasm2goData_data_bin[395838:396696], wasm2goData_data_bin[396696:396818], wasm2goData_data_bin[396818:396860], wasm2goData_data_bin[396860:396902], wasm2goData_data_bin[396902:397104], wasm2goData_data_bin[397104:397170], wasm2goData_data_bin[397170:397189], wasm2goData_data_bin[397189:397217], wasm2goData_data_bin[397217:397251], wasm2goData_data_bin[397251:397324], wasm2goData_data_bin[397324:397332]}
	m.ThreadStart64 = Fn3059
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
	return Fn330(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn331(m)
}
func WasmShutdown(m *base.Module) {
	Fn332(m)
}
func DbgVecDotF16(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn942(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn968(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4_1_q8_1(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn969(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn970(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_1_q8_1(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn971(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ8_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn972(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ2KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn973(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ3KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn974(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn975(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn976(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ6KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn977(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq4NlQ8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn838(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgQuantizeMatQ8K4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn843(m, l0, l1, l2)
}
func DbgGemvQ4_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn846(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn848(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ5K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn851(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ6K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn853(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvIq4Nl8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn855(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ4_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn859(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn860(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ5K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn861(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ6K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn862(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmIq4Nl8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn863(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgQuantizeMatQ8_0_4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn979(m, l0, l1, l2)
}
func DbgGemvQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn982(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn980(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn983(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn981(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgVecSwigluF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	Fn943(m, l0, l1, l2, l3)
}
func DbgVecSoftMaxF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) float64 {
	return Fn944(m, l0, l1, l2, l3)
}
func DbgVecMadF16F32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) {
	Fn945(m, l0, l1, l2, l3)
}
func DbgSimdGemmF32(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) {
	Fn946(m, l0, l1, l2, l3, l4, l5)
}
func DbgFlashAttnKvF16(m *base.Module, l0 int64) {
	Fn955(m, l0)
}
func WasiThreadStart(m *base.Module, l0 int32, l1 int64) {
	Fn3059(m, l0, l1)
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
	packed = Fn326(m, l0, l1)
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
	packed = Fn328(m, l0, l1)
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
	packed = Fn329(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
