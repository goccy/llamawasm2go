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
	m.T0 = make([]any, 2244)
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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:124280], wasm2goData_data_bin[124280:144377], wasm2goData_data_bin[144377:144658], wasm2goData_data_bin[144658:144829], wasm2goData_data_bin[144829:144832], wasm2goData_data_bin[144832:145865], wasm2goData_data_bin[145865:145890], wasm2goData_data_bin[145890:145915], wasm2goData_data_bin[145915:145940], wasm2goData_data_bin[145940:145965], wasm2goData_data_bin[145965:146080], wasm2goData_data_bin[146080:146083], wasm2goData_data_bin[146083:146086], wasm2goData_data_bin[146086:146201], wasm2goData_data_bin[146201:146204], wasm2goData_data_bin[146204:146207], wasm2goData_data_bin[146207:146481], wasm2goData_data_bin[146481:146755], wasm2goData_data_bin[146755:146845], wasm2goData_data_bin[146845:146887], wasm2goData_data_bin[146887:146889], wasm2goData_data_bin[146889:146891], wasm2goData_data_bin[146891:190748], wasm2goData_data_bin[190748:191038], wasm2goData_data_bin[191038:191279], wasm2goData_data_bin[191279:191281], wasm2goData_data_bin[191281:191330], wasm2goData_data_bin[191330:191387], wasm2goData_data_bin[191387:228301], wasm2goData_data_bin[228301:228335], wasm2goData_data_bin[228335:228417], wasm2goData_data_bin[228417:234314], wasm2goData_data_bin[234314:259415], wasm2goData_data_bin[259415:260023], wasm2goData_data_bin[260023:262593], wasm2goData_data_bin[262593:263766], wasm2goData_data_bin[263766:263932], wasm2goData_data_bin[263932:264865], wasm2goData_data_bin[264865:264994], wasm2goData_data_bin[264994:265123], wasm2goData_data_bin[265123:265576], wasm2goData_data_bin[265576:265718], wasm2goData_data_bin[265718:266192], wasm2goData_data_bin[266192:266324], wasm2goData_data_bin[266324:273729], wasm2goData_data_bin[273729:273788], wasm2goData_data_bin[273788:273847], wasm2goData_data_bin[273847:273906], wasm2goData_data_bin[273906:273965], wasm2goData_data_bin[273965:274024], wasm2goData_data_bin[274024:274083], wasm2goData_data_bin[274083:274142], wasm2goData_data_bin[274142:274243], wasm2goData_data_bin[274243:274344], wasm2goData_data_bin[274344:274445], wasm2goData_data_bin[274445:274546], wasm2goData_data_bin[274546:274647], wasm2goData_data_bin[274647:274748], wasm2goData_data_bin[274748:274849], wasm2goData_data_bin[274849:274950], wasm2goData_data_bin[274950:275051], wasm2goData_data_bin[275051:275152], wasm2goData_data_bin[275152:275254], wasm2goData_data_bin[275254:275356], wasm2goData_data_bin[275356:355353], wasm2goData_data_bin[355353:379300], wasm2goData_data_bin[379300:379420], wasm2goData_data_bin[379420:379861], wasm2goData_data_bin[379861:379864], wasm2goData_data_bin[379864:394736], wasm2goData_data_bin[394736:394738], wasm2goData_data_bin[394738:397311], wasm2goData_data_bin[397311:397344], wasm2goData_data_bin[397344:397377], wasm2goData_data_bin[397377:397419], wasm2goData_data_bin[397419:397433], wasm2goData_data_bin[397433:397466], wasm2goData_data_bin[397466:397565], wasm2goData_data_bin[397565:397727], wasm2goData_data_bin[397727:398585], wasm2goData_data_bin[398585:398707], wasm2goData_data_bin[398707:398749], wasm2goData_data_bin[398749:398791], wasm2goData_data_bin[398791:398993], wasm2goData_data_bin[398993:399059], wasm2goData_data_bin[399059:399078], wasm2goData_data_bin[399078:399106], wasm2goData_data_bin[399106:399140], wasm2goData_data_bin[399140:399213], wasm2goData_data_bin[399213:399221]}
	m.ThreadStart64 = Fn3088
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
	m.T0 = make([]any, 2244)
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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:124280], wasm2goData_data_bin[124280:144377], wasm2goData_data_bin[144377:144658], wasm2goData_data_bin[144658:144829], wasm2goData_data_bin[144829:144832], wasm2goData_data_bin[144832:145865], wasm2goData_data_bin[145865:145890], wasm2goData_data_bin[145890:145915], wasm2goData_data_bin[145915:145940], wasm2goData_data_bin[145940:145965], wasm2goData_data_bin[145965:146080], wasm2goData_data_bin[146080:146083], wasm2goData_data_bin[146083:146086], wasm2goData_data_bin[146086:146201], wasm2goData_data_bin[146201:146204], wasm2goData_data_bin[146204:146207], wasm2goData_data_bin[146207:146481], wasm2goData_data_bin[146481:146755], wasm2goData_data_bin[146755:146845], wasm2goData_data_bin[146845:146887], wasm2goData_data_bin[146887:146889], wasm2goData_data_bin[146889:146891], wasm2goData_data_bin[146891:190748], wasm2goData_data_bin[190748:191038], wasm2goData_data_bin[191038:191279], wasm2goData_data_bin[191279:191281], wasm2goData_data_bin[191281:191330], wasm2goData_data_bin[191330:191387], wasm2goData_data_bin[191387:228301], wasm2goData_data_bin[228301:228335], wasm2goData_data_bin[228335:228417], wasm2goData_data_bin[228417:234314], wasm2goData_data_bin[234314:259415], wasm2goData_data_bin[259415:260023], wasm2goData_data_bin[260023:262593], wasm2goData_data_bin[262593:263766], wasm2goData_data_bin[263766:263932], wasm2goData_data_bin[263932:264865], wasm2goData_data_bin[264865:264994], wasm2goData_data_bin[264994:265123], wasm2goData_data_bin[265123:265576], wasm2goData_data_bin[265576:265718], wasm2goData_data_bin[265718:266192], wasm2goData_data_bin[266192:266324], wasm2goData_data_bin[266324:273729], wasm2goData_data_bin[273729:273788], wasm2goData_data_bin[273788:273847], wasm2goData_data_bin[273847:273906], wasm2goData_data_bin[273906:273965], wasm2goData_data_bin[273965:274024], wasm2goData_data_bin[274024:274083], wasm2goData_data_bin[274083:274142], wasm2goData_data_bin[274142:274243], wasm2goData_data_bin[274243:274344], wasm2goData_data_bin[274344:274445], wasm2goData_data_bin[274445:274546], wasm2goData_data_bin[274546:274647], wasm2goData_data_bin[274647:274748], wasm2goData_data_bin[274748:274849], wasm2goData_data_bin[274849:274950], wasm2goData_data_bin[274950:275051], wasm2goData_data_bin[275051:275152], wasm2goData_data_bin[275152:275254], wasm2goData_data_bin[275254:275356], wasm2goData_data_bin[275356:355353], wasm2goData_data_bin[355353:379300], wasm2goData_data_bin[379300:379420], wasm2goData_data_bin[379420:379861], wasm2goData_data_bin[379861:379864], wasm2goData_data_bin[379864:394736], wasm2goData_data_bin[394736:394738], wasm2goData_data_bin[394738:397311], wasm2goData_data_bin[397311:397344], wasm2goData_data_bin[397344:397377], wasm2goData_data_bin[397377:397419], wasm2goData_data_bin[397419:397433], wasm2goData_data_bin[397433:397466], wasm2goData_data_bin[397466:397565], wasm2goData_data_bin[397565:397727], wasm2goData_data_bin[397727:398585], wasm2goData_data_bin[398585:398707], wasm2goData_data_bin[398707:398749], wasm2goData_data_bin[398749:398791], wasm2goData_data_bin[398791:398993], wasm2goData_data_bin[398993:399059], wasm2goData_data_bin[399059:399078], wasm2goData_data_bin[399078:399106], wasm2goData_data_bin[399106:399140], wasm2goData_data_bin[399140:399213], wasm2goData_data_bin[399213:399221]}
	m.ThreadStart64 = Fn3088
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
	m.T0 = make([]any, 2244)
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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:124280], wasm2goData_data_bin[124280:144377], wasm2goData_data_bin[144377:144658], wasm2goData_data_bin[144658:144829], wasm2goData_data_bin[144829:144832], wasm2goData_data_bin[144832:145865], wasm2goData_data_bin[145865:145890], wasm2goData_data_bin[145890:145915], wasm2goData_data_bin[145915:145940], wasm2goData_data_bin[145940:145965], wasm2goData_data_bin[145965:146080], wasm2goData_data_bin[146080:146083], wasm2goData_data_bin[146083:146086], wasm2goData_data_bin[146086:146201], wasm2goData_data_bin[146201:146204], wasm2goData_data_bin[146204:146207], wasm2goData_data_bin[146207:146481], wasm2goData_data_bin[146481:146755], wasm2goData_data_bin[146755:146845], wasm2goData_data_bin[146845:146887], wasm2goData_data_bin[146887:146889], wasm2goData_data_bin[146889:146891], wasm2goData_data_bin[146891:190748], wasm2goData_data_bin[190748:191038], wasm2goData_data_bin[191038:191279], wasm2goData_data_bin[191279:191281], wasm2goData_data_bin[191281:191330], wasm2goData_data_bin[191330:191387], wasm2goData_data_bin[191387:228301], wasm2goData_data_bin[228301:228335], wasm2goData_data_bin[228335:228417], wasm2goData_data_bin[228417:234314], wasm2goData_data_bin[234314:259415], wasm2goData_data_bin[259415:260023], wasm2goData_data_bin[260023:262593], wasm2goData_data_bin[262593:263766], wasm2goData_data_bin[263766:263932], wasm2goData_data_bin[263932:264865], wasm2goData_data_bin[264865:264994], wasm2goData_data_bin[264994:265123], wasm2goData_data_bin[265123:265576], wasm2goData_data_bin[265576:265718], wasm2goData_data_bin[265718:266192], wasm2goData_data_bin[266192:266324], wasm2goData_data_bin[266324:273729], wasm2goData_data_bin[273729:273788], wasm2goData_data_bin[273788:273847], wasm2goData_data_bin[273847:273906], wasm2goData_data_bin[273906:273965], wasm2goData_data_bin[273965:274024], wasm2goData_data_bin[274024:274083], wasm2goData_data_bin[274083:274142], wasm2goData_data_bin[274142:274243], wasm2goData_data_bin[274243:274344], wasm2goData_data_bin[274344:274445], wasm2goData_data_bin[274445:274546], wasm2goData_data_bin[274546:274647], wasm2goData_data_bin[274647:274748], wasm2goData_data_bin[274748:274849], wasm2goData_data_bin[274849:274950], wasm2goData_data_bin[274950:275051], wasm2goData_data_bin[275051:275152], wasm2goData_data_bin[275152:275254], wasm2goData_data_bin[275254:275356], wasm2goData_data_bin[275356:355353], wasm2goData_data_bin[355353:379300], wasm2goData_data_bin[379300:379420], wasm2goData_data_bin[379420:379861], wasm2goData_data_bin[379861:379864], wasm2goData_data_bin[379864:394736], wasm2goData_data_bin[394736:394738], wasm2goData_data_bin[394738:397311], wasm2goData_data_bin[397311:397344], wasm2goData_data_bin[397344:397377], wasm2goData_data_bin[397377:397419], wasm2goData_data_bin[397419:397433], wasm2goData_data_bin[397433:397466], wasm2goData_data_bin[397466:397565], wasm2goData_data_bin[397565:397727], wasm2goData_data_bin[397727:398585], wasm2goData_data_bin[398585:398707], wasm2goData_data_bin[398707:398749], wasm2goData_data_bin[398749:398791], wasm2goData_data_bin[398791:398993], wasm2goData_data_bin[398993:399059], wasm2goData_data_bin[399059:399078], wasm2goData_data_bin[399078:399106], wasm2goData_data_bin[399106:399140], wasm2goData_data_bin[399140:399213], wasm2goData_data_bin[399213:399221]}
	m.ThreadStart64 = Fn3088
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
	return Fn337(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn338(m)
}
func WasmShutdown(m *base.Module) {
	Fn339(m)
}
func DbgKernelInit(m *base.Module) {
	Fn362(m)
}
func DbgVecDotF16(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn972(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn998(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4_1_q8_1(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn999(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1000(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_1_q8_1(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1001(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ8_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1002(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ2KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1003(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ3KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1004(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1005(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1006(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ6KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn1007(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq2XxsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn860(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq2XsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn861(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq3XxsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn863(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq1SQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn865(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq4NlQ8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn867(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq3SQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn864(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq2SQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn862(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq4XsQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn868(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotIq1MQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn866(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotTq1_0_q8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn858(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotTq2_0_q8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn859(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotMxfp4Q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn856(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotNvfp4Q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn857(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ1_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn854(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ2_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn855(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgQuantizeMatQ8K4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn872(m, l0, l1, l2)
}
func DbgGemvQ4_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn875(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn877(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ5K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn880(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ6K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn882(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvIq4Nl8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn884(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvMxfp48x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn886(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ4_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn888(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn889(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ5K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn890(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ6K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn891(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmIq4Nl8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn892(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmMxfp48x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn893(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgQuantizeMatQ8_0_4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn1009(m, l0, l1, l2)
}
func DbgGemvQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1012(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1010(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1013(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn1011(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgVecSwigluF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	Fn973(m, l0, l1, l2, l3)
}
func DbgVecSoftMaxF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) float64 {
	return Fn974(m, l0, l1, l2, l3)
}
func DbgVecMadF16F32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) {
	Fn975(m, l0, l1, l2, l3)
}
func DbgSimdGemmF32(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) {
	Fn976(m, l0, l1, l2, l3, l4, l5)
}
func DbgFlashAttnKvF16(m *base.Module, l0 int64) {
	Fn985(m, l0)
}
func WasiThreadStart(m *base.Module, l0 int32, l1 int64) {
	Fn3088(m, l0, l1)
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
	packed = Fn301(m, l0, l1)
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
	packed = Fn302(m, l0, l1)
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
	packed = Fn304(m, l0, l1)
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
	packed = Fn305(m, l0, l1)
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
	packed = Fn331(m, l0, l1)
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
	packed = Fn332(m, l0, l1)
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
	packed = Fn333(m, l0, l1)
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
	packed = Fn335(m, l0, l1)
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
	packed = Fn336(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
