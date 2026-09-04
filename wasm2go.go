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
	m.T0 = make([]any, 2238)
	m.G0 = int64(8388608)
	m.G1 = int64(0)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:127632], wasm2goData_data_bin[127632:147625], wasm2goData_data_bin[147625:147906], wasm2goData_data_bin[147906:148077], wasm2goData_data_bin[148077:148080], wasm2goData_data_bin[148080:149113], wasm2goData_data_bin[149113:149138], wasm2goData_data_bin[149138:149163], wasm2goData_data_bin[149163:149188], wasm2goData_data_bin[149188:149213], wasm2goData_data_bin[149213:149328], wasm2goData_data_bin[149328:149331], wasm2goData_data_bin[149331:149334], wasm2goData_data_bin[149334:149449], wasm2goData_data_bin[149449:149452], wasm2goData_data_bin[149452:149455], wasm2goData_data_bin[149455:149729], wasm2goData_data_bin[149729:150003], wasm2goData_data_bin[150003:150093], wasm2goData_data_bin[150093:150135], wasm2goData_data_bin[150135:150137], wasm2goData_data_bin[150137:150139], wasm2goData_data_bin[150139:193996], wasm2goData_data_bin[193996:194286], wasm2goData_data_bin[194286:194527], wasm2goData_data_bin[194527:194529], wasm2goData_data_bin[194529:194578], wasm2goData_data_bin[194578:194635], wasm2goData_data_bin[194635:231397], wasm2goData_data_bin[231397:231431], wasm2goData_data_bin[231431:231513], wasm2goData_data_bin[231513:237418], wasm2goData_data_bin[237418:262519], wasm2goData_data_bin[262519:263127], wasm2goData_data_bin[263127:265697], wasm2goData_data_bin[265697:266870], wasm2goData_data_bin[266870:267036], wasm2goData_data_bin[267036:267969], wasm2goData_data_bin[267969:268098], wasm2goData_data_bin[268098:268227], wasm2goData_data_bin[268227:268680], wasm2goData_data_bin[268680:268822], wasm2goData_data_bin[268822:269296], wasm2goData_data_bin[269296:269428], wasm2goData_data_bin[269428:276833], wasm2goData_data_bin[276833:276891], wasm2goData_data_bin[276891:276950], wasm2goData_data_bin[276950:277009], wasm2goData_data_bin[277009:277068], wasm2goData_data_bin[277068:277127], wasm2goData_data_bin[277127:277186], wasm2goData_data_bin[277186:277245], wasm2goData_data_bin[277245:277346], wasm2goData_data_bin[277346:277447], wasm2goData_data_bin[277447:277548], wasm2goData_data_bin[277548:277649], wasm2goData_data_bin[277649:277750], wasm2goData_data_bin[277750:277851], wasm2goData_data_bin[277851:277952], wasm2goData_data_bin[277952:278053], wasm2goData_data_bin[278053:278154], wasm2goData_data_bin[278154:278255], wasm2goData_data_bin[278255:278357], wasm2goData_data_bin[278357:278459], wasm2goData_data_bin[278459:358456], wasm2goData_data_bin[358456:382403], wasm2goData_data_bin[382403:382523], wasm2goData_data_bin[382523:382964], wasm2goData_data_bin[382964:382967], wasm2goData_data_bin[382967:397839], wasm2goData_data_bin[397839:397841], wasm2goData_data_bin[397841:400414], wasm2goData_data_bin[400414:400447], wasm2goData_data_bin[400447:400480], wasm2goData_data_bin[400480:400522], wasm2goData_data_bin[400522:400536], wasm2goData_data_bin[400536:400569], wasm2goData_data_bin[400569:400668], wasm2goData_data_bin[400668:400830], wasm2goData_data_bin[400830:401688], wasm2goData_data_bin[401688:401810], wasm2goData_data_bin[401810:401852], wasm2goData_data_bin[401852:401894], wasm2goData_data_bin[401894:402096], wasm2goData_data_bin[402096:402162], wasm2goData_data_bin[402162:402181], wasm2goData_data_bin[402181:402209], wasm2goData_data_bin[402209:402243], wasm2goData_data_bin[402243:402316], wasm2goData_data_bin[402316:402324]}
	m.ThreadStart64 = Fn2848
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
	m.T0 = make([]any, 2238)
	m.G0 = int64(8388608)
	m.G1 = int64(0)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:127632], wasm2goData_data_bin[127632:147625], wasm2goData_data_bin[147625:147906], wasm2goData_data_bin[147906:148077], wasm2goData_data_bin[148077:148080], wasm2goData_data_bin[148080:149113], wasm2goData_data_bin[149113:149138], wasm2goData_data_bin[149138:149163], wasm2goData_data_bin[149163:149188], wasm2goData_data_bin[149188:149213], wasm2goData_data_bin[149213:149328], wasm2goData_data_bin[149328:149331], wasm2goData_data_bin[149331:149334], wasm2goData_data_bin[149334:149449], wasm2goData_data_bin[149449:149452], wasm2goData_data_bin[149452:149455], wasm2goData_data_bin[149455:149729], wasm2goData_data_bin[149729:150003], wasm2goData_data_bin[150003:150093], wasm2goData_data_bin[150093:150135], wasm2goData_data_bin[150135:150137], wasm2goData_data_bin[150137:150139], wasm2goData_data_bin[150139:193996], wasm2goData_data_bin[193996:194286], wasm2goData_data_bin[194286:194527], wasm2goData_data_bin[194527:194529], wasm2goData_data_bin[194529:194578], wasm2goData_data_bin[194578:194635], wasm2goData_data_bin[194635:231397], wasm2goData_data_bin[231397:231431], wasm2goData_data_bin[231431:231513], wasm2goData_data_bin[231513:237418], wasm2goData_data_bin[237418:262519], wasm2goData_data_bin[262519:263127], wasm2goData_data_bin[263127:265697], wasm2goData_data_bin[265697:266870], wasm2goData_data_bin[266870:267036], wasm2goData_data_bin[267036:267969], wasm2goData_data_bin[267969:268098], wasm2goData_data_bin[268098:268227], wasm2goData_data_bin[268227:268680], wasm2goData_data_bin[268680:268822], wasm2goData_data_bin[268822:269296], wasm2goData_data_bin[269296:269428], wasm2goData_data_bin[269428:276833], wasm2goData_data_bin[276833:276891], wasm2goData_data_bin[276891:276950], wasm2goData_data_bin[276950:277009], wasm2goData_data_bin[277009:277068], wasm2goData_data_bin[277068:277127], wasm2goData_data_bin[277127:277186], wasm2goData_data_bin[277186:277245], wasm2goData_data_bin[277245:277346], wasm2goData_data_bin[277346:277447], wasm2goData_data_bin[277447:277548], wasm2goData_data_bin[277548:277649], wasm2goData_data_bin[277649:277750], wasm2goData_data_bin[277750:277851], wasm2goData_data_bin[277851:277952], wasm2goData_data_bin[277952:278053], wasm2goData_data_bin[278053:278154], wasm2goData_data_bin[278154:278255], wasm2goData_data_bin[278255:278357], wasm2goData_data_bin[278357:278459], wasm2goData_data_bin[278459:358456], wasm2goData_data_bin[358456:382403], wasm2goData_data_bin[382403:382523], wasm2goData_data_bin[382523:382964], wasm2goData_data_bin[382964:382967], wasm2goData_data_bin[382967:397839], wasm2goData_data_bin[397839:397841], wasm2goData_data_bin[397841:400414], wasm2goData_data_bin[400414:400447], wasm2goData_data_bin[400447:400480], wasm2goData_data_bin[400480:400522], wasm2goData_data_bin[400522:400536], wasm2goData_data_bin[400536:400569], wasm2goData_data_bin[400569:400668], wasm2goData_data_bin[400668:400830], wasm2goData_data_bin[400830:401688], wasm2goData_data_bin[401688:401810], wasm2goData_data_bin[401810:401852], wasm2goData_data_bin[401852:401894], wasm2goData_data_bin[401894:402096], wasm2goData_data_bin[402096:402162], wasm2goData_data_bin[402162:402181], wasm2goData_data_bin[402181:402209], wasm2goData_data_bin[402209:402243], wasm2goData_data_bin[402243:402316], wasm2goData_data_bin[402316:402324]}
	m.ThreadStart64 = Fn2848
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
	m.T0 = make([]any, 2238)
	m.G0 = int64(8388608)
	m.G1 = int64(0)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:127632], wasm2goData_data_bin[127632:147625], wasm2goData_data_bin[147625:147906], wasm2goData_data_bin[147906:148077], wasm2goData_data_bin[148077:148080], wasm2goData_data_bin[148080:149113], wasm2goData_data_bin[149113:149138], wasm2goData_data_bin[149138:149163], wasm2goData_data_bin[149163:149188], wasm2goData_data_bin[149188:149213], wasm2goData_data_bin[149213:149328], wasm2goData_data_bin[149328:149331], wasm2goData_data_bin[149331:149334], wasm2goData_data_bin[149334:149449], wasm2goData_data_bin[149449:149452], wasm2goData_data_bin[149452:149455], wasm2goData_data_bin[149455:149729], wasm2goData_data_bin[149729:150003], wasm2goData_data_bin[150003:150093], wasm2goData_data_bin[150093:150135], wasm2goData_data_bin[150135:150137], wasm2goData_data_bin[150137:150139], wasm2goData_data_bin[150139:193996], wasm2goData_data_bin[193996:194286], wasm2goData_data_bin[194286:194527], wasm2goData_data_bin[194527:194529], wasm2goData_data_bin[194529:194578], wasm2goData_data_bin[194578:194635], wasm2goData_data_bin[194635:231397], wasm2goData_data_bin[231397:231431], wasm2goData_data_bin[231431:231513], wasm2goData_data_bin[231513:237418], wasm2goData_data_bin[237418:262519], wasm2goData_data_bin[262519:263127], wasm2goData_data_bin[263127:265697], wasm2goData_data_bin[265697:266870], wasm2goData_data_bin[266870:267036], wasm2goData_data_bin[267036:267969], wasm2goData_data_bin[267969:268098], wasm2goData_data_bin[268098:268227], wasm2goData_data_bin[268227:268680], wasm2goData_data_bin[268680:268822], wasm2goData_data_bin[268822:269296], wasm2goData_data_bin[269296:269428], wasm2goData_data_bin[269428:276833], wasm2goData_data_bin[276833:276891], wasm2goData_data_bin[276891:276950], wasm2goData_data_bin[276950:277009], wasm2goData_data_bin[277009:277068], wasm2goData_data_bin[277068:277127], wasm2goData_data_bin[277127:277186], wasm2goData_data_bin[277186:277245], wasm2goData_data_bin[277245:277346], wasm2goData_data_bin[277346:277447], wasm2goData_data_bin[277447:277548], wasm2goData_data_bin[277548:277649], wasm2goData_data_bin[277649:277750], wasm2goData_data_bin[277750:277851], wasm2goData_data_bin[277851:277952], wasm2goData_data_bin[277952:278053], wasm2goData_data_bin[278053:278154], wasm2goData_data_bin[278154:278255], wasm2goData_data_bin[278255:278357], wasm2goData_data_bin[278357:278459], wasm2goData_data_bin[278459:358456], wasm2goData_data_bin[358456:382403], wasm2goData_data_bin[382403:382523], wasm2goData_data_bin[382523:382964], wasm2goData_data_bin[382964:382967], wasm2goData_data_bin[382967:397839], wasm2goData_data_bin[397839:397841], wasm2goData_data_bin[397841:400414], wasm2goData_data_bin[400414:400447], wasm2goData_data_bin[400447:400480], wasm2goData_data_bin[400480:400522], wasm2goData_data_bin[400522:400536], wasm2goData_data_bin[400536:400569], wasm2goData_data_bin[400569:400668], wasm2goData_data_bin[400668:400830], wasm2goData_data_bin[400830:401688], wasm2goData_data_bin[401688:401810], wasm2goData_data_bin[401810:401852], wasm2goData_data_bin[401852:401894], wasm2goData_data_bin[401894:402096], wasm2goData_data_bin[402096:402162], wasm2goData_data_bin[402162:402181], wasm2goData_data_bin[402181:402209], wasm2goData_data_bin[402209:402243], wasm2goData_data_bin[402243:402316], wasm2goData_data_bin[402316:402324]}
	m.ThreadStart64 = Fn2848
	base.RestoreGlobals(m, globals)
	return m
}
func Initialize(m *base.Module) {
	Fn20(m)
}
func WasmAlloc(m *base.Module, l0 int64) int64 {
	return Fn287(m, l0)
}
func WasmFree(m *base.Module, l0 int64) {
	Fn26(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int64, l1 int64) int64 {
	return Fn325(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn326(m)
}
func WasmShutdown(m *base.Module) {
	Fn327(m)
}
func DbgVecDotF16(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn892(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn919(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn924(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ6KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn926(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgGemvQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn928(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn929(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgVecSwigluF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	Fn893(m, l0, l1, l2, l3)
}
func DbgVecSoftMaxF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) float64 {
	return Fn894(m, l0, l1, l2, l3)
}
func DbgVecMadF16F32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) {
	Fn895(m, l0, l1, l2, l3)
}
func DbgSimdGemmF32(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) {
	Fn896(m, l0, l1, l2, l3, l4, l5)
}
func DbgFlashAttnKvF16(m *base.Module, l0 int64) {
	Fn904(m, l0)
}
func WasiThreadStart(m *base.Module, l0 int32, l1 int64) {
	Fn2848(m, l0, l1)
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
	packed = Fn288(m, l0, l1)
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
	packed = Fn291(m, l0, l1)
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
	packed = Fn292(m, l0, l1)
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
	packed = Fn293(m, l0, l1)
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
	packed = Fn294(m, l0, l1)
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
	packed = Fn295(m, l0, l1)
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
	packed = Fn296(m, l0, l1)
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
	packed = Fn298(m, l0, l1)
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
	packed = Fn299(m, l0, l1)
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
	packed = Fn301(m, l0, l1)
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
	packed = Fn302(m, l0, l1)
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
	packed = Fn303(m, l0, l1)
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
	packed = Fn304(m, l0, l1)
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
	packed = Fn305(m, l0, l1)
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
	packed = Fn306(m, l0, l1)
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
	packed = Fn307(m, l0, l1)
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
	packed = Fn308(m, l0, l1)
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
	packed = Fn309(m, l0, l1)
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
	packed = Fn310(m, l0, l1)
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
	packed = Fn311(m, l0, l1)
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
	packed = Fn312(m, l0, l1)
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
	packed = Fn313(m, l0, l1)
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
	packed = Fn314(m, l0, l1)
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
	packed = Fn315(m, l0, l1)
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
	packed = Fn316(m, l0, l1)
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
	packed = Fn317(m, l0, l1)
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
	packed = Fn318(m, l0, l1)
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
	packed = Fn319(m, l0, l1)
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
	packed = Fn320(m, l0, l1)
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
	packed = Fn321(m, l0, l1)
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
	packed = Fn323(m, l0, l1)
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
	packed = Fn324(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
