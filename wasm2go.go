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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:140360], wasm2goData_data_bin[140360:160345], wasm2goData_data_bin[160345:160626], wasm2goData_data_bin[160626:160797], wasm2goData_data_bin[160797:160800], wasm2goData_data_bin[160800:161833], wasm2goData_data_bin[161833:161858], wasm2goData_data_bin[161858:161883], wasm2goData_data_bin[161883:161908], wasm2goData_data_bin[161908:161933], wasm2goData_data_bin[161933:162048], wasm2goData_data_bin[162048:162051], wasm2goData_data_bin[162051:162054], wasm2goData_data_bin[162054:162169], wasm2goData_data_bin[162169:162172], wasm2goData_data_bin[162172:162175], wasm2goData_data_bin[162175:162449], wasm2goData_data_bin[162449:162723], wasm2goData_data_bin[162723:162813], wasm2goData_data_bin[162813:162855], wasm2goData_data_bin[162855:162857], wasm2goData_data_bin[162857:162859], wasm2goData_data_bin[162859:206716], wasm2goData_data_bin[206716:207006], wasm2goData_data_bin[207006:207247], wasm2goData_data_bin[207247:207249], wasm2goData_data_bin[207249:207298], wasm2goData_data_bin[207298:207355], wasm2goData_data_bin[207355:244117], wasm2goData_data_bin[244117:244151], wasm2goData_data_bin[244151:244233], wasm2goData_data_bin[244233:250138], wasm2goData_data_bin[250138:275239], wasm2goData_data_bin[275239:275847], wasm2goData_data_bin[275847:278417], wasm2goData_data_bin[278417:279590], wasm2goData_data_bin[279590:279756], wasm2goData_data_bin[279756:280689], wasm2goData_data_bin[280689:280818], wasm2goData_data_bin[280818:280947], wasm2goData_data_bin[280947:281400], wasm2goData_data_bin[281400:281542], wasm2goData_data_bin[281542:282016], wasm2goData_data_bin[282016:282148], wasm2goData_data_bin[282148:289553], wasm2goData_data_bin[289553:289612], wasm2goData_data_bin[289612:289671], wasm2goData_data_bin[289671:289730], wasm2goData_data_bin[289730:289789], wasm2goData_data_bin[289789:289848], wasm2goData_data_bin[289848:289907], wasm2goData_data_bin[289907:289966], wasm2goData_data_bin[289966:290067], wasm2goData_data_bin[290067:290168], wasm2goData_data_bin[290168:290269], wasm2goData_data_bin[290269:290370], wasm2goData_data_bin[290370:290471], wasm2goData_data_bin[290471:290572], wasm2goData_data_bin[290572:290673], wasm2goData_data_bin[290673:290774], wasm2goData_data_bin[290774:290875], wasm2goData_data_bin[290875:290976], wasm2goData_data_bin[290976:291078], wasm2goData_data_bin[291078:291180], wasm2goData_data_bin[291180:371177], wasm2goData_data_bin[371177:395124], wasm2goData_data_bin[395124:395244], wasm2goData_data_bin[395244:395685], wasm2goData_data_bin[395685:395688], wasm2goData_data_bin[395688:410560], wasm2goData_data_bin[410560:410562], wasm2goData_data_bin[410562:413135], wasm2goData_data_bin[413135:413168], wasm2goData_data_bin[413168:413201], wasm2goData_data_bin[413201:413243], wasm2goData_data_bin[413243:413257], wasm2goData_data_bin[413257:413290], wasm2goData_data_bin[413290:413389], wasm2goData_data_bin[413389:413551], wasm2goData_data_bin[413551:414409], wasm2goData_data_bin[414409:414531], wasm2goData_data_bin[414531:414573], wasm2goData_data_bin[414573:414615], wasm2goData_data_bin[414615:414817], wasm2goData_data_bin[414817:414883], wasm2goData_data_bin[414883:414902], wasm2goData_data_bin[414902:414930], wasm2goData_data_bin[414930:414964], wasm2goData_data_bin[414964:415037], wasm2goData_data_bin[415037:415045]}
	m.ThreadStart64 = Fn2874
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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:140360], wasm2goData_data_bin[140360:160345], wasm2goData_data_bin[160345:160626], wasm2goData_data_bin[160626:160797], wasm2goData_data_bin[160797:160800], wasm2goData_data_bin[160800:161833], wasm2goData_data_bin[161833:161858], wasm2goData_data_bin[161858:161883], wasm2goData_data_bin[161883:161908], wasm2goData_data_bin[161908:161933], wasm2goData_data_bin[161933:162048], wasm2goData_data_bin[162048:162051], wasm2goData_data_bin[162051:162054], wasm2goData_data_bin[162054:162169], wasm2goData_data_bin[162169:162172], wasm2goData_data_bin[162172:162175], wasm2goData_data_bin[162175:162449], wasm2goData_data_bin[162449:162723], wasm2goData_data_bin[162723:162813], wasm2goData_data_bin[162813:162855], wasm2goData_data_bin[162855:162857], wasm2goData_data_bin[162857:162859], wasm2goData_data_bin[162859:206716], wasm2goData_data_bin[206716:207006], wasm2goData_data_bin[207006:207247], wasm2goData_data_bin[207247:207249], wasm2goData_data_bin[207249:207298], wasm2goData_data_bin[207298:207355], wasm2goData_data_bin[207355:244117], wasm2goData_data_bin[244117:244151], wasm2goData_data_bin[244151:244233], wasm2goData_data_bin[244233:250138], wasm2goData_data_bin[250138:275239], wasm2goData_data_bin[275239:275847], wasm2goData_data_bin[275847:278417], wasm2goData_data_bin[278417:279590], wasm2goData_data_bin[279590:279756], wasm2goData_data_bin[279756:280689], wasm2goData_data_bin[280689:280818], wasm2goData_data_bin[280818:280947], wasm2goData_data_bin[280947:281400], wasm2goData_data_bin[281400:281542], wasm2goData_data_bin[281542:282016], wasm2goData_data_bin[282016:282148], wasm2goData_data_bin[282148:289553], wasm2goData_data_bin[289553:289612], wasm2goData_data_bin[289612:289671], wasm2goData_data_bin[289671:289730], wasm2goData_data_bin[289730:289789], wasm2goData_data_bin[289789:289848], wasm2goData_data_bin[289848:289907], wasm2goData_data_bin[289907:289966], wasm2goData_data_bin[289966:290067], wasm2goData_data_bin[290067:290168], wasm2goData_data_bin[290168:290269], wasm2goData_data_bin[290269:290370], wasm2goData_data_bin[290370:290471], wasm2goData_data_bin[290471:290572], wasm2goData_data_bin[290572:290673], wasm2goData_data_bin[290673:290774], wasm2goData_data_bin[290774:290875], wasm2goData_data_bin[290875:290976], wasm2goData_data_bin[290976:291078], wasm2goData_data_bin[291078:291180], wasm2goData_data_bin[291180:371177], wasm2goData_data_bin[371177:395124], wasm2goData_data_bin[395124:395244], wasm2goData_data_bin[395244:395685], wasm2goData_data_bin[395685:395688], wasm2goData_data_bin[395688:410560], wasm2goData_data_bin[410560:410562], wasm2goData_data_bin[410562:413135], wasm2goData_data_bin[413135:413168], wasm2goData_data_bin[413168:413201], wasm2goData_data_bin[413201:413243], wasm2goData_data_bin[413243:413257], wasm2goData_data_bin[413257:413290], wasm2goData_data_bin[413290:413389], wasm2goData_data_bin[413389:413551], wasm2goData_data_bin[413551:414409], wasm2goData_data_bin[414409:414531], wasm2goData_data_bin[414531:414573], wasm2goData_data_bin[414573:414615], wasm2goData_data_bin[414615:414817], wasm2goData_data_bin[414817:414883], wasm2goData_data_bin[414883:414902], wasm2goData_data_bin[414902:414930], wasm2goData_data_bin[414930:414964], wasm2goData_data_bin[414964:415037], wasm2goData_data_bin[415037:415045]}
	m.ThreadStart64 = Fn2874
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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:140360], wasm2goData_data_bin[140360:160345], wasm2goData_data_bin[160345:160626], wasm2goData_data_bin[160626:160797], wasm2goData_data_bin[160797:160800], wasm2goData_data_bin[160800:161833], wasm2goData_data_bin[161833:161858], wasm2goData_data_bin[161858:161883], wasm2goData_data_bin[161883:161908], wasm2goData_data_bin[161908:161933], wasm2goData_data_bin[161933:162048], wasm2goData_data_bin[162048:162051], wasm2goData_data_bin[162051:162054], wasm2goData_data_bin[162054:162169], wasm2goData_data_bin[162169:162172], wasm2goData_data_bin[162172:162175], wasm2goData_data_bin[162175:162449], wasm2goData_data_bin[162449:162723], wasm2goData_data_bin[162723:162813], wasm2goData_data_bin[162813:162855], wasm2goData_data_bin[162855:162857], wasm2goData_data_bin[162857:162859], wasm2goData_data_bin[162859:206716], wasm2goData_data_bin[206716:207006], wasm2goData_data_bin[207006:207247], wasm2goData_data_bin[207247:207249], wasm2goData_data_bin[207249:207298], wasm2goData_data_bin[207298:207355], wasm2goData_data_bin[207355:244117], wasm2goData_data_bin[244117:244151], wasm2goData_data_bin[244151:244233], wasm2goData_data_bin[244233:250138], wasm2goData_data_bin[250138:275239], wasm2goData_data_bin[275239:275847], wasm2goData_data_bin[275847:278417], wasm2goData_data_bin[278417:279590], wasm2goData_data_bin[279590:279756], wasm2goData_data_bin[279756:280689], wasm2goData_data_bin[280689:280818], wasm2goData_data_bin[280818:280947], wasm2goData_data_bin[280947:281400], wasm2goData_data_bin[281400:281542], wasm2goData_data_bin[281542:282016], wasm2goData_data_bin[282016:282148], wasm2goData_data_bin[282148:289553], wasm2goData_data_bin[289553:289612], wasm2goData_data_bin[289612:289671], wasm2goData_data_bin[289671:289730], wasm2goData_data_bin[289730:289789], wasm2goData_data_bin[289789:289848], wasm2goData_data_bin[289848:289907], wasm2goData_data_bin[289907:289966], wasm2goData_data_bin[289966:290067], wasm2goData_data_bin[290067:290168], wasm2goData_data_bin[290168:290269], wasm2goData_data_bin[290269:290370], wasm2goData_data_bin[290370:290471], wasm2goData_data_bin[290471:290572], wasm2goData_data_bin[290572:290673], wasm2goData_data_bin[290673:290774], wasm2goData_data_bin[290774:290875], wasm2goData_data_bin[290875:290976], wasm2goData_data_bin[290976:291078], wasm2goData_data_bin[291078:291180], wasm2goData_data_bin[291180:371177], wasm2goData_data_bin[371177:395124], wasm2goData_data_bin[395124:395244], wasm2goData_data_bin[395244:395685], wasm2goData_data_bin[395685:395688], wasm2goData_data_bin[395688:410560], wasm2goData_data_bin[410560:410562], wasm2goData_data_bin[410562:413135], wasm2goData_data_bin[413135:413168], wasm2goData_data_bin[413168:413201], wasm2goData_data_bin[413201:413243], wasm2goData_data_bin[413243:413257], wasm2goData_data_bin[413257:413290], wasm2goData_data_bin[413290:413389], wasm2goData_data_bin[413389:413551], wasm2goData_data_bin[413551:414409], wasm2goData_data_bin[414409:414531], wasm2goData_data_bin[414531:414573], wasm2goData_data_bin[414573:414615], wasm2goData_data_bin[414615:414817], wasm2goData_data_bin[414817:414883], wasm2goData_data_bin[414883:414902], wasm2goData_data_bin[414902:414930], wasm2goData_data_bin[414930:414964], wasm2goData_data_bin[414964:415037], wasm2goData_data_bin[415037:415045]}
	m.ThreadStart64 = Fn2874
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
func DbgGemvQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn931(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn932(m, l0, l1, l2, l3, l4, l5, l6)
}
func WasiThreadStart(m *base.Module, l0 int32, l1 int64) {
	Fn2874(m, l0, l1)
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
