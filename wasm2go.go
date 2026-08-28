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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:121680], wasm2goData_data_bin[121680:141641], wasm2goData_data_bin[141641:141922], wasm2goData_data_bin[141922:142093], wasm2goData_data_bin[142093:142096], wasm2goData_data_bin[142096:143129], wasm2goData_data_bin[143129:143154], wasm2goData_data_bin[143154:143179], wasm2goData_data_bin[143179:143204], wasm2goData_data_bin[143204:143229], wasm2goData_data_bin[143229:143344], wasm2goData_data_bin[143344:143347], wasm2goData_data_bin[143347:143350], wasm2goData_data_bin[143350:143465], wasm2goData_data_bin[143465:143468], wasm2goData_data_bin[143468:143471], wasm2goData_data_bin[143471:143745], wasm2goData_data_bin[143745:144019], wasm2goData_data_bin[144019:144109], wasm2goData_data_bin[144109:144151], wasm2goData_data_bin[144151:144153], wasm2goData_data_bin[144153:144155], wasm2goData_data_bin[144155:188012], wasm2goData_data_bin[188012:188302], wasm2goData_data_bin[188302:188543], wasm2goData_data_bin[188543:188545], wasm2goData_data_bin[188545:188594], wasm2goData_data_bin[188594:188651], wasm2goData_data_bin[188651:225413], wasm2goData_data_bin[225413:225447], wasm2goData_data_bin[225447:225529], wasm2goData_data_bin[225529:231434], wasm2goData_data_bin[231434:256535], wasm2goData_data_bin[256535:257143], wasm2goData_data_bin[257143:259713], wasm2goData_data_bin[259713:260886], wasm2goData_data_bin[260886:261052], wasm2goData_data_bin[261052:261985], wasm2goData_data_bin[261985:262114], wasm2goData_data_bin[262114:262243], wasm2goData_data_bin[262243:262696], wasm2goData_data_bin[262696:262838], wasm2goData_data_bin[262838:263311], wasm2goData_data_bin[263311:263443], wasm2goData_data_bin[263443:270848], wasm2goData_data_bin[270848:270907], wasm2goData_data_bin[270907:270966], wasm2goData_data_bin[270966:271025], wasm2goData_data_bin[271025:271084], wasm2goData_data_bin[271084:271142], wasm2goData_data_bin[271142:271201], wasm2goData_data_bin[271201:271260], wasm2goData_data_bin[271260:271361], wasm2goData_data_bin[271361:271462], wasm2goData_data_bin[271462:271563], wasm2goData_data_bin[271563:271664], wasm2goData_data_bin[271664:271765], wasm2goData_data_bin[271765:271866], wasm2goData_data_bin[271866:271967], wasm2goData_data_bin[271967:272068], wasm2goData_data_bin[272068:272169], wasm2goData_data_bin[272169:272270], wasm2goData_data_bin[272270:272372], wasm2goData_data_bin[272372:272474], wasm2goData_data_bin[272474:352471], wasm2goData_data_bin[352471:376417], wasm2goData_data_bin[376417:376537], wasm2goData_data_bin[376537:376978], wasm2goData_data_bin[376978:376981], wasm2goData_data_bin[376981:391853], wasm2goData_data_bin[391853:391855], wasm2goData_data_bin[391855:394428], wasm2goData_data_bin[394428:394461], wasm2goData_data_bin[394461:394494], wasm2goData_data_bin[394494:394536], wasm2goData_data_bin[394536:394550], wasm2goData_data_bin[394550:394583], wasm2goData_data_bin[394583:394682], wasm2goData_data_bin[394682:394844], wasm2goData_data_bin[394844:395702], wasm2goData_data_bin[395702:395824], wasm2goData_data_bin[395824:395866], wasm2goData_data_bin[395866:395908], wasm2goData_data_bin[395908:396110], wasm2goData_data_bin[396110:396176], wasm2goData_data_bin[396176:396195], wasm2goData_data_bin[396195:396223], wasm2goData_data_bin[396223:396257], wasm2goData_data_bin[396257:396330], wasm2goData_data_bin[396330:396338]}
	m.ThreadStart64 = Fn3042
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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:121680], wasm2goData_data_bin[121680:141641], wasm2goData_data_bin[141641:141922], wasm2goData_data_bin[141922:142093], wasm2goData_data_bin[142093:142096], wasm2goData_data_bin[142096:143129], wasm2goData_data_bin[143129:143154], wasm2goData_data_bin[143154:143179], wasm2goData_data_bin[143179:143204], wasm2goData_data_bin[143204:143229], wasm2goData_data_bin[143229:143344], wasm2goData_data_bin[143344:143347], wasm2goData_data_bin[143347:143350], wasm2goData_data_bin[143350:143465], wasm2goData_data_bin[143465:143468], wasm2goData_data_bin[143468:143471], wasm2goData_data_bin[143471:143745], wasm2goData_data_bin[143745:144019], wasm2goData_data_bin[144019:144109], wasm2goData_data_bin[144109:144151], wasm2goData_data_bin[144151:144153], wasm2goData_data_bin[144153:144155], wasm2goData_data_bin[144155:188012], wasm2goData_data_bin[188012:188302], wasm2goData_data_bin[188302:188543], wasm2goData_data_bin[188543:188545], wasm2goData_data_bin[188545:188594], wasm2goData_data_bin[188594:188651], wasm2goData_data_bin[188651:225413], wasm2goData_data_bin[225413:225447], wasm2goData_data_bin[225447:225529], wasm2goData_data_bin[225529:231434], wasm2goData_data_bin[231434:256535], wasm2goData_data_bin[256535:257143], wasm2goData_data_bin[257143:259713], wasm2goData_data_bin[259713:260886], wasm2goData_data_bin[260886:261052], wasm2goData_data_bin[261052:261985], wasm2goData_data_bin[261985:262114], wasm2goData_data_bin[262114:262243], wasm2goData_data_bin[262243:262696], wasm2goData_data_bin[262696:262838], wasm2goData_data_bin[262838:263311], wasm2goData_data_bin[263311:263443], wasm2goData_data_bin[263443:270848], wasm2goData_data_bin[270848:270907], wasm2goData_data_bin[270907:270966], wasm2goData_data_bin[270966:271025], wasm2goData_data_bin[271025:271084], wasm2goData_data_bin[271084:271142], wasm2goData_data_bin[271142:271201], wasm2goData_data_bin[271201:271260], wasm2goData_data_bin[271260:271361], wasm2goData_data_bin[271361:271462], wasm2goData_data_bin[271462:271563], wasm2goData_data_bin[271563:271664], wasm2goData_data_bin[271664:271765], wasm2goData_data_bin[271765:271866], wasm2goData_data_bin[271866:271967], wasm2goData_data_bin[271967:272068], wasm2goData_data_bin[272068:272169], wasm2goData_data_bin[272169:272270], wasm2goData_data_bin[272270:272372], wasm2goData_data_bin[272372:272474], wasm2goData_data_bin[272474:352471], wasm2goData_data_bin[352471:376417], wasm2goData_data_bin[376417:376537], wasm2goData_data_bin[376537:376978], wasm2goData_data_bin[376978:376981], wasm2goData_data_bin[376981:391853], wasm2goData_data_bin[391853:391855], wasm2goData_data_bin[391855:394428], wasm2goData_data_bin[394428:394461], wasm2goData_data_bin[394461:394494], wasm2goData_data_bin[394494:394536], wasm2goData_data_bin[394536:394550], wasm2goData_data_bin[394550:394583], wasm2goData_data_bin[394583:394682], wasm2goData_data_bin[394682:394844], wasm2goData_data_bin[394844:395702], wasm2goData_data_bin[395702:395824], wasm2goData_data_bin[395824:395866], wasm2goData_data_bin[395866:395908], wasm2goData_data_bin[395908:396110], wasm2goData_data_bin[396110:396176], wasm2goData_data_bin[396176:396195], wasm2goData_data_bin[396195:396223], wasm2goData_data_bin[396223:396257], wasm2goData_data_bin[396257:396330], wasm2goData_data_bin[396330:396338]}
	m.ThreadStart64 = Fn3042
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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:121680], wasm2goData_data_bin[121680:141641], wasm2goData_data_bin[141641:141922], wasm2goData_data_bin[141922:142093], wasm2goData_data_bin[142093:142096], wasm2goData_data_bin[142096:143129], wasm2goData_data_bin[143129:143154], wasm2goData_data_bin[143154:143179], wasm2goData_data_bin[143179:143204], wasm2goData_data_bin[143204:143229], wasm2goData_data_bin[143229:143344], wasm2goData_data_bin[143344:143347], wasm2goData_data_bin[143347:143350], wasm2goData_data_bin[143350:143465], wasm2goData_data_bin[143465:143468], wasm2goData_data_bin[143468:143471], wasm2goData_data_bin[143471:143745], wasm2goData_data_bin[143745:144019], wasm2goData_data_bin[144019:144109], wasm2goData_data_bin[144109:144151], wasm2goData_data_bin[144151:144153], wasm2goData_data_bin[144153:144155], wasm2goData_data_bin[144155:188012], wasm2goData_data_bin[188012:188302], wasm2goData_data_bin[188302:188543], wasm2goData_data_bin[188543:188545], wasm2goData_data_bin[188545:188594], wasm2goData_data_bin[188594:188651], wasm2goData_data_bin[188651:225413], wasm2goData_data_bin[225413:225447], wasm2goData_data_bin[225447:225529], wasm2goData_data_bin[225529:231434], wasm2goData_data_bin[231434:256535], wasm2goData_data_bin[256535:257143], wasm2goData_data_bin[257143:259713], wasm2goData_data_bin[259713:260886], wasm2goData_data_bin[260886:261052], wasm2goData_data_bin[261052:261985], wasm2goData_data_bin[261985:262114], wasm2goData_data_bin[262114:262243], wasm2goData_data_bin[262243:262696], wasm2goData_data_bin[262696:262838], wasm2goData_data_bin[262838:263311], wasm2goData_data_bin[263311:263443], wasm2goData_data_bin[263443:270848], wasm2goData_data_bin[270848:270907], wasm2goData_data_bin[270907:270966], wasm2goData_data_bin[270966:271025], wasm2goData_data_bin[271025:271084], wasm2goData_data_bin[271084:271142], wasm2goData_data_bin[271142:271201], wasm2goData_data_bin[271201:271260], wasm2goData_data_bin[271260:271361], wasm2goData_data_bin[271361:271462], wasm2goData_data_bin[271462:271563], wasm2goData_data_bin[271563:271664], wasm2goData_data_bin[271664:271765], wasm2goData_data_bin[271765:271866], wasm2goData_data_bin[271866:271967], wasm2goData_data_bin[271967:272068], wasm2goData_data_bin[272068:272169], wasm2goData_data_bin[272169:272270], wasm2goData_data_bin[272270:272372], wasm2goData_data_bin[272372:272474], wasm2goData_data_bin[272474:352471], wasm2goData_data_bin[352471:376417], wasm2goData_data_bin[376417:376537], wasm2goData_data_bin[376537:376978], wasm2goData_data_bin[376978:376981], wasm2goData_data_bin[376981:391853], wasm2goData_data_bin[391853:391855], wasm2goData_data_bin[391855:394428], wasm2goData_data_bin[394428:394461], wasm2goData_data_bin[394461:394494], wasm2goData_data_bin[394494:394536], wasm2goData_data_bin[394536:394550], wasm2goData_data_bin[394550:394583], wasm2goData_data_bin[394583:394682], wasm2goData_data_bin[394682:394844], wasm2goData_data_bin[394844:395702], wasm2goData_data_bin[395702:395824], wasm2goData_data_bin[395824:395866], wasm2goData_data_bin[395866:395908], wasm2goData_data_bin[395908:396110], wasm2goData_data_bin[396110:396176], wasm2goData_data_bin[396176:396195], wasm2goData_data_bin[396195:396223], wasm2goData_data_bin[396223:396257], wasm2goData_data_bin[396257:396330], wasm2goData_data_bin[396330:396338]}
	m.ThreadStart64 = Fn3042
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
	Fn3042(m, l0, l1)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
			err = fmt.Errorf("wasm trap: %v", r)
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
