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
	InitElemSeg_1_1(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:121576], wasm2goData_data_bin[121576:141529], wasm2goData_data_bin[141529:141810], wasm2goData_data_bin[141810:141981], wasm2goData_data_bin[141981:141984], wasm2goData_data_bin[141984:143017], wasm2goData_data_bin[143017:143042], wasm2goData_data_bin[143042:143067], wasm2goData_data_bin[143067:143092], wasm2goData_data_bin[143092:143117], wasm2goData_data_bin[143117:143232], wasm2goData_data_bin[143232:143235], wasm2goData_data_bin[143235:143238], wasm2goData_data_bin[143238:143353], wasm2goData_data_bin[143353:143356], wasm2goData_data_bin[143356:143359], wasm2goData_data_bin[143359:143633], wasm2goData_data_bin[143633:143907], wasm2goData_data_bin[143907:143997], wasm2goData_data_bin[143997:144039], wasm2goData_data_bin[144039:144041], wasm2goData_data_bin[144041:144043], wasm2goData_data_bin[144043:187900], wasm2goData_data_bin[187900:188190], wasm2goData_data_bin[188190:188431], wasm2goData_data_bin[188431:188433], wasm2goData_data_bin[188433:188482], wasm2goData_data_bin[188482:188539], wasm2goData_data_bin[188539:225301], wasm2goData_data_bin[225301:225335], wasm2goData_data_bin[225335:225417], wasm2goData_data_bin[225417:231322], wasm2goData_data_bin[231322:256423], wasm2goData_data_bin[256423:257031], wasm2goData_data_bin[257031:259601], wasm2goData_data_bin[259601:260774], wasm2goData_data_bin[260774:260940], wasm2goData_data_bin[260940:261873], wasm2goData_data_bin[261873:262002], wasm2goData_data_bin[262002:262131], wasm2goData_data_bin[262131:262584], wasm2goData_data_bin[262584:262726], wasm2goData_data_bin[262726:263200], wasm2goData_data_bin[263200:263332], wasm2goData_data_bin[263332:270737], wasm2goData_data_bin[270737:270796], wasm2goData_data_bin[270796:270855], wasm2goData_data_bin[270855:270914], wasm2goData_data_bin[270914:270973], wasm2goData_data_bin[270973:271032], wasm2goData_data_bin[271032:271091], wasm2goData_data_bin[271091:271150], wasm2goData_data_bin[271150:271251], wasm2goData_data_bin[271251:271352], wasm2goData_data_bin[271352:271453], wasm2goData_data_bin[271453:271554], wasm2goData_data_bin[271554:271655], wasm2goData_data_bin[271655:271756], wasm2goData_data_bin[271756:271857], wasm2goData_data_bin[271857:271958], wasm2goData_data_bin[271958:272059], wasm2goData_data_bin[272059:272160], wasm2goData_data_bin[272160:272262], wasm2goData_data_bin[272262:272364], wasm2goData_data_bin[272364:352361], wasm2goData_data_bin[352361:376308], wasm2goData_data_bin[376308:376428], wasm2goData_data_bin[376428:376869], wasm2goData_data_bin[376869:376872], wasm2goData_data_bin[376872:391744], wasm2goData_data_bin[391744:391746], wasm2goData_data_bin[391746:394319], wasm2goData_data_bin[394319:394352], wasm2goData_data_bin[394352:394385], wasm2goData_data_bin[394385:394427], wasm2goData_data_bin[394427:394441], wasm2goData_data_bin[394441:394474], wasm2goData_data_bin[394474:394573], wasm2goData_data_bin[394573:394735], wasm2goData_data_bin[394735:395593], wasm2goData_data_bin[395593:395715], wasm2goData_data_bin[395715:395757], wasm2goData_data_bin[395757:395799], wasm2goData_data_bin[395799:396001], wasm2goData_data_bin[396001:396067], wasm2goData_data_bin[396067:396086], wasm2goData_data_bin[396086:396114], wasm2goData_data_bin[396114:396148], wasm2goData_data_bin[396148:396221], wasm2goData_data_bin[396221:396229]}
	m.ThreadStart64 = Fn3041
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
	InitElemSeg_1_1(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:121576], wasm2goData_data_bin[121576:141529], wasm2goData_data_bin[141529:141810], wasm2goData_data_bin[141810:141981], wasm2goData_data_bin[141981:141984], wasm2goData_data_bin[141984:143017], wasm2goData_data_bin[143017:143042], wasm2goData_data_bin[143042:143067], wasm2goData_data_bin[143067:143092], wasm2goData_data_bin[143092:143117], wasm2goData_data_bin[143117:143232], wasm2goData_data_bin[143232:143235], wasm2goData_data_bin[143235:143238], wasm2goData_data_bin[143238:143353], wasm2goData_data_bin[143353:143356], wasm2goData_data_bin[143356:143359], wasm2goData_data_bin[143359:143633], wasm2goData_data_bin[143633:143907], wasm2goData_data_bin[143907:143997], wasm2goData_data_bin[143997:144039], wasm2goData_data_bin[144039:144041], wasm2goData_data_bin[144041:144043], wasm2goData_data_bin[144043:187900], wasm2goData_data_bin[187900:188190], wasm2goData_data_bin[188190:188431], wasm2goData_data_bin[188431:188433], wasm2goData_data_bin[188433:188482], wasm2goData_data_bin[188482:188539], wasm2goData_data_bin[188539:225301], wasm2goData_data_bin[225301:225335], wasm2goData_data_bin[225335:225417], wasm2goData_data_bin[225417:231322], wasm2goData_data_bin[231322:256423], wasm2goData_data_bin[256423:257031], wasm2goData_data_bin[257031:259601], wasm2goData_data_bin[259601:260774], wasm2goData_data_bin[260774:260940], wasm2goData_data_bin[260940:261873], wasm2goData_data_bin[261873:262002], wasm2goData_data_bin[262002:262131], wasm2goData_data_bin[262131:262584], wasm2goData_data_bin[262584:262726], wasm2goData_data_bin[262726:263200], wasm2goData_data_bin[263200:263332], wasm2goData_data_bin[263332:270737], wasm2goData_data_bin[270737:270796], wasm2goData_data_bin[270796:270855], wasm2goData_data_bin[270855:270914], wasm2goData_data_bin[270914:270973], wasm2goData_data_bin[270973:271032], wasm2goData_data_bin[271032:271091], wasm2goData_data_bin[271091:271150], wasm2goData_data_bin[271150:271251], wasm2goData_data_bin[271251:271352], wasm2goData_data_bin[271352:271453], wasm2goData_data_bin[271453:271554], wasm2goData_data_bin[271554:271655], wasm2goData_data_bin[271655:271756], wasm2goData_data_bin[271756:271857], wasm2goData_data_bin[271857:271958], wasm2goData_data_bin[271958:272059], wasm2goData_data_bin[272059:272160], wasm2goData_data_bin[272160:272262], wasm2goData_data_bin[272262:272364], wasm2goData_data_bin[272364:352361], wasm2goData_data_bin[352361:376308], wasm2goData_data_bin[376308:376428], wasm2goData_data_bin[376428:376869], wasm2goData_data_bin[376869:376872], wasm2goData_data_bin[376872:391744], wasm2goData_data_bin[391744:391746], wasm2goData_data_bin[391746:394319], wasm2goData_data_bin[394319:394352], wasm2goData_data_bin[394352:394385], wasm2goData_data_bin[394385:394427], wasm2goData_data_bin[394427:394441], wasm2goData_data_bin[394441:394474], wasm2goData_data_bin[394474:394573], wasm2goData_data_bin[394573:394735], wasm2goData_data_bin[394735:395593], wasm2goData_data_bin[395593:395715], wasm2goData_data_bin[395715:395757], wasm2goData_data_bin[395757:395799], wasm2goData_data_bin[395799:396001], wasm2goData_data_bin[396001:396067], wasm2goData_data_bin[396067:396086], wasm2goData_data_bin[396086:396114], wasm2goData_data_bin[396114:396148], wasm2goData_data_bin[396148:396221], wasm2goData_data_bin[396221:396229]}
	m.ThreadStart64 = Fn3041
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
	InitElemSeg_1_1(m)
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:121576], wasm2goData_data_bin[121576:141529], wasm2goData_data_bin[141529:141810], wasm2goData_data_bin[141810:141981], wasm2goData_data_bin[141981:141984], wasm2goData_data_bin[141984:143017], wasm2goData_data_bin[143017:143042], wasm2goData_data_bin[143042:143067], wasm2goData_data_bin[143067:143092], wasm2goData_data_bin[143092:143117], wasm2goData_data_bin[143117:143232], wasm2goData_data_bin[143232:143235], wasm2goData_data_bin[143235:143238], wasm2goData_data_bin[143238:143353], wasm2goData_data_bin[143353:143356], wasm2goData_data_bin[143356:143359], wasm2goData_data_bin[143359:143633], wasm2goData_data_bin[143633:143907], wasm2goData_data_bin[143907:143997], wasm2goData_data_bin[143997:144039], wasm2goData_data_bin[144039:144041], wasm2goData_data_bin[144041:144043], wasm2goData_data_bin[144043:187900], wasm2goData_data_bin[187900:188190], wasm2goData_data_bin[188190:188431], wasm2goData_data_bin[188431:188433], wasm2goData_data_bin[188433:188482], wasm2goData_data_bin[188482:188539], wasm2goData_data_bin[188539:225301], wasm2goData_data_bin[225301:225335], wasm2goData_data_bin[225335:225417], wasm2goData_data_bin[225417:231322], wasm2goData_data_bin[231322:256423], wasm2goData_data_bin[256423:257031], wasm2goData_data_bin[257031:259601], wasm2goData_data_bin[259601:260774], wasm2goData_data_bin[260774:260940], wasm2goData_data_bin[260940:261873], wasm2goData_data_bin[261873:262002], wasm2goData_data_bin[262002:262131], wasm2goData_data_bin[262131:262584], wasm2goData_data_bin[262584:262726], wasm2goData_data_bin[262726:263200], wasm2goData_data_bin[263200:263332], wasm2goData_data_bin[263332:270737], wasm2goData_data_bin[270737:270796], wasm2goData_data_bin[270796:270855], wasm2goData_data_bin[270855:270914], wasm2goData_data_bin[270914:270973], wasm2goData_data_bin[270973:271032], wasm2goData_data_bin[271032:271091], wasm2goData_data_bin[271091:271150], wasm2goData_data_bin[271150:271251], wasm2goData_data_bin[271251:271352], wasm2goData_data_bin[271352:271453], wasm2goData_data_bin[271453:271554], wasm2goData_data_bin[271554:271655], wasm2goData_data_bin[271655:271756], wasm2goData_data_bin[271756:271857], wasm2goData_data_bin[271857:271958], wasm2goData_data_bin[271958:272059], wasm2goData_data_bin[272059:272160], wasm2goData_data_bin[272160:272262], wasm2goData_data_bin[272262:272364], wasm2goData_data_bin[272364:352361], wasm2goData_data_bin[352361:376308], wasm2goData_data_bin[376308:376428], wasm2goData_data_bin[376428:376869], wasm2goData_data_bin[376869:376872], wasm2goData_data_bin[376872:391744], wasm2goData_data_bin[391744:391746], wasm2goData_data_bin[391746:394319], wasm2goData_data_bin[394319:394352], wasm2goData_data_bin[394352:394385], wasm2goData_data_bin[394385:394427], wasm2goData_data_bin[394427:394441], wasm2goData_data_bin[394441:394474], wasm2goData_data_bin[394474:394573], wasm2goData_data_bin[394573:394735], wasm2goData_data_bin[394735:395593], wasm2goData_data_bin[395593:395715], wasm2goData_data_bin[395715:395757], wasm2goData_data_bin[395757:395799], wasm2goData_data_bin[395799:396001], wasm2goData_data_bin[396001:396067], wasm2goData_data_bin[396067:396086], wasm2goData_data_bin[396086:396114], wasm2goData_data_bin[396114:396148], wasm2goData_data_bin[396148:396221], wasm2goData_data_bin[396221:396229]}
	m.ThreadStart64 = Fn3041
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
	return Fn327(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn328(m)
}
func WasmShutdown(m *base.Module) {
	Fn329(m)
}
func DbgGemvQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn963(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn964(m, l0, l1, l2, l3, l4, l5, l6)
}
func WasiThreadStart(m *base.Module, l0 int32, l1 int64) {
	Fn3041(m, l0, l1)
}
func Inv_0_0(m *base.Module, l0, l1 int64) (packed int64, err error) {
	savedG0 := m.G0
	savedG1 := m.G1
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			m.G1 = savedG1
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn322(m, l0, l1)
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
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn323(m, l0, l1)
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
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn325(m, l0, l1)
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
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn326(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
