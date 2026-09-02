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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:122256], wasm2goData_data_bin[122256:142249], wasm2goData_data_bin[142249:142530], wasm2goData_data_bin[142530:142701], wasm2goData_data_bin[142701:142704], wasm2goData_data_bin[142704:143737], wasm2goData_data_bin[143737:143762], wasm2goData_data_bin[143762:143787], wasm2goData_data_bin[143787:143812], wasm2goData_data_bin[143812:143837], wasm2goData_data_bin[143837:143952], wasm2goData_data_bin[143952:143955], wasm2goData_data_bin[143955:143958], wasm2goData_data_bin[143958:144073], wasm2goData_data_bin[144073:144076], wasm2goData_data_bin[144076:144079], wasm2goData_data_bin[144079:144353], wasm2goData_data_bin[144353:144627], wasm2goData_data_bin[144627:144717], wasm2goData_data_bin[144717:144759], wasm2goData_data_bin[144759:144761], wasm2goData_data_bin[144761:144763], wasm2goData_data_bin[144763:188620], wasm2goData_data_bin[188620:188910], wasm2goData_data_bin[188910:189151], wasm2goData_data_bin[189151:189153], wasm2goData_data_bin[189153:189202], wasm2goData_data_bin[189202:189259], wasm2goData_data_bin[189259:226021], wasm2goData_data_bin[226021:226055], wasm2goData_data_bin[226055:226137], wasm2goData_data_bin[226137:232042], wasm2goData_data_bin[232042:257143], wasm2goData_data_bin[257143:257751], wasm2goData_data_bin[257751:260321], wasm2goData_data_bin[260321:261494], wasm2goData_data_bin[261494:261660], wasm2goData_data_bin[261660:262593], wasm2goData_data_bin[262593:262722], wasm2goData_data_bin[262722:262851], wasm2goData_data_bin[262851:263304], wasm2goData_data_bin[263304:263446], wasm2goData_data_bin[263446:263920], wasm2goData_data_bin[263920:264052], wasm2goData_data_bin[264052:271457], wasm2goData_data_bin[271457:271515], wasm2goData_data_bin[271515:271574], wasm2goData_data_bin[271574:271633], wasm2goData_data_bin[271633:271692], wasm2goData_data_bin[271692:271751], wasm2goData_data_bin[271751:271810], wasm2goData_data_bin[271810:271869], wasm2goData_data_bin[271869:271970], wasm2goData_data_bin[271970:272071], wasm2goData_data_bin[272071:272172], wasm2goData_data_bin[272172:272273], wasm2goData_data_bin[272273:272374], wasm2goData_data_bin[272374:272475], wasm2goData_data_bin[272475:272576], wasm2goData_data_bin[272576:272677], wasm2goData_data_bin[272677:272778], wasm2goData_data_bin[272778:272879], wasm2goData_data_bin[272879:272981], wasm2goData_data_bin[272981:273083], wasm2goData_data_bin[273083:353080], wasm2goData_data_bin[353080:377027], wasm2goData_data_bin[377027:377147], wasm2goData_data_bin[377147:377588], wasm2goData_data_bin[377588:377591], wasm2goData_data_bin[377591:392463], wasm2goData_data_bin[392463:392465], wasm2goData_data_bin[392465:395038], wasm2goData_data_bin[395038:395071], wasm2goData_data_bin[395071:395104], wasm2goData_data_bin[395104:395146], wasm2goData_data_bin[395146:395160], wasm2goData_data_bin[395160:395193], wasm2goData_data_bin[395193:395292], wasm2goData_data_bin[395292:395454], wasm2goData_data_bin[395454:396312], wasm2goData_data_bin[396312:396434], wasm2goData_data_bin[396434:396476], wasm2goData_data_bin[396476:396518], wasm2goData_data_bin[396518:396720], wasm2goData_data_bin[396720:396786], wasm2goData_data_bin[396786:396805], wasm2goData_data_bin[396805:396833], wasm2goData_data_bin[396833:396867], wasm2goData_data_bin[396867:396940], wasm2goData_data_bin[396940:396948]}
	m.ThreadStart64 = Fn3046
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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:122256], wasm2goData_data_bin[122256:142249], wasm2goData_data_bin[142249:142530], wasm2goData_data_bin[142530:142701], wasm2goData_data_bin[142701:142704], wasm2goData_data_bin[142704:143737], wasm2goData_data_bin[143737:143762], wasm2goData_data_bin[143762:143787], wasm2goData_data_bin[143787:143812], wasm2goData_data_bin[143812:143837], wasm2goData_data_bin[143837:143952], wasm2goData_data_bin[143952:143955], wasm2goData_data_bin[143955:143958], wasm2goData_data_bin[143958:144073], wasm2goData_data_bin[144073:144076], wasm2goData_data_bin[144076:144079], wasm2goData_data_bin[144079:144353], wasm2goData_data_bin[144353:144627], wasm2goData_data_bin[144627:144717], wasm2goData_data_bin[144717:144759], wasm2goData_data_bin[144759:144761], wasm2goData_data_bin[144761:144763], wasm2goData_data_bin[144763:188620], wasm2goData_data_bin[188620:188910], wasm2goData_data_bin[188910:189151], wasm2goData_data_bin[189151:189153], wasm2goData_data_bin[189153:189202], wasm2goData_data_bin[189202:189259], wasm2goData_data_bin[189259:226021], wasm2goData_data_bin[226021:226055], wasm2goData_data_bin[226055:226137], wasm2goData_data_bin[226137:232042], wasm2goData_data_bin[232042:257143], wasm2goData_data_bin[257143:257751], wasm2goData_data_bin[257751:260321], wasm2goData_data_bin[260321:261494], wasm2goData_data_bin[261494:261660], wasm2goData_data_bin[261660:262593], wasm2goData_data_bin[262593:262722], wasm2goData_data_bin[262722:262851], wasm2goData_data_bin[262851:263304], wasm2goData_data_bin[263304:263446], wasm2goData_data_bin[263446:263920], wasm2goData_data_bin[263920:264052], wasm2goData_data_bin[264052:271457], wasm2goData_data_bin[271457:271515], wasm2goData_data_bin[271515:271574], wasm2goData_data_bin[271574:271633], wasm2goData_data_bin[271633:271692], wasm2goData_data_bin[271692:271751], wasm2goData_data_bin[271751:271810], wasm2goData_data_bin[271810:271869], wasm2goData_data_bin[271869:271970], wasm2goData_data_bin[271970:272071], wasm2goData_data_bin[272071:272172], wasm2goData_data_bin[272172:272273], wasm2goData_data_bin[272273:272374], wasm2goData_data_bin[272374:272475], wasm2goData_data_bin[272475:272576], wasm2goData_data_bin[272576:272677], wasm2goData_data_bin[272677:272778], wasm2goData_data_bin[272778:272879], wasm2goData_data_bin[272879:272981], wasm2goData_data_bin[272981:273083], wasm2goData_data_bin[273083:353080], wasm2goData_data_bin[353080:377027], wasm2goData_data_bin[377027:377147], wasm2goData_data_bin[377147:377588], wasm2goData_data_bin[377588:377591], wasm2goData_data_bin[377591:392463], wasm2goData_data_bin[392463:392465], wasm2goData_data_bin[392465:395038], wasm2goData_data_bin[395038:395071], wasm2goData_data_bin[395071:395104], wasm2goData_data_bin[395104:395146], wasm2goData_data_bin[395146:395160], wasm2goData_data_bin[395160:395193], wasm2goData_data_bin[395193:395292], wasm2goData_data_bin[395292:395454], wasm2goData_data_bin[395454:396312], wasm2goData_data_bin[396312:396434], wasm2goData_data_bin[396434:396476], wasm2goData_data_bin[396476:396518], wasm2goData_data_bin[396518:396720], wasm2goData_data_bin[396720:396786], wasm2goData_data_bin[396786:396805], wasm2goData_data_bin[396805:396833], wasm2goData_data_bin[396833:396867], wasm2goData_data_bin[396867:396940], wasm2goData_data_bin[396940:396948]}
	m.ThreadStart64 = Fn3046
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
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:122256], wasm2goData_data_bin[122256:142249], wasm2goData_data_bin[142249:142530], wasm2goData_data_bin[142530:142701], wasm2goData_data_bin[142701:142704], wasm2goData_data_bin[142704:143737], wasm2goData_data_bin[143737:143762], wasm2goData_data_bin[143762:143787], wasm2goData_data_bin[143787:143812], wasm2goData_data_bin[143812:143837], wasm2goData_data_bin[143837:143952], wasm2goData_data_bin[143952:143955], wasm2goData_data_bin[143955:143958], wasm2goData_data_bin[143958:144073], wasm2goData_data_bin[144073:144076], wasm2goData_data_bin[144076:144079], wasm2goData_data_bin[144079:144353], wasm2goData_data_bin[144353:144627], wasm2goData_data_bin[144627:144717], wasm2goData_data_bin[144717:144759], wasm2goData_data_bin[144759:144761], wasm2goData_data_bin[144761:144763], wasm2goData_data_bin[144763:188620], wasm2goData_data_bin[188620:188910], wasm2goData_data_bin[188910:189151], wasm2goData_data_bin[189151:189153], wasm2goData_data_bin[189153:189202], wasm2goData_data_bin[189202:189259], wasm2goData_data_bin[189259:226021], wasm2goData_data_bin[226021:226055], wasm2goData_data_bin[226055:226137], wasm2goData_data_bin[226137:232042], wasm2goData_data_bin[232042:257143], wasm2goData_data_bin[257143:257751], wasm2goData_data_bin[257751:260321], wasm2goData_data_bin[260321:261494], wasm2goData_data_bin[261494:261660], wasm2goData_data_bin[261660:262593], wasm2goData_data_bin[262593:262722], wasm2goData_data_bin[262722:262851], wasm2goData_data_bin[262851:263304], wasm2goData_data_bin[263304:263446], wasm2goData_data_bin[263446:263920], wasm2goData_data_bin[263920:264052], wasm2goData_data_bin[264052:271457], wasm2goData_data_bin[271457:271515], wasm2goData_data_bin[271515:271574], wasm2goData_data_bin[271574:271633], wasm2goData_data_bin[271633:271692], wasm2goData_data_bin[271692:271751], wasm2goData_data_bin[271751:271810], wasm2goData_data_bin[271810:271869], wasm2goData_data_bin[271869:271970], wasm2goData_data_bin[271970:272071], wasm2goData_data_bin[272071:272172], wasm2goData_data_bin[272172:272273], wasm2goData_data_bin[272273:272374], wasm2goData_data_bin[272374:272475], wasm2goData_data_bin[272475:272576], wasm2goData_data_bin[272576:272677], wasm2goData_data_bin[272677:272778], wasm2goData_data_bin[272778:272879], wasm2goData_data_bin[272879:272981], wasm2goData_data_bin[272981:273083], wasm2goData_data_bin[273083:353080], wasm2goData_data_bin[353080:377027], wasm2goData_data_bin[377027:377147], wasm2goData_data_bin[377147:377588], wasm2goData_data_bin[377588:377591], wasm2goData_data_bin[377591:392463], wasm2goData_data_bin[392463:392465], wasm2goData_data_bin[392465:395038], wasm2goData_data_bin[395038:395071], wasm2goData_data_bin[395071:395104], wasm2goData_data_bin[395104:395146], wasm2goData_data_bin[395146:395160], wasm2goData_data_bin[395160:395193], wasm2goData_data_bin[395193:395292], wasm2goData_data_bin[395292:395454], wasm2goData_data_bin[395454:396312], wasm2goData_data_bin[396312:396434], wasm2goData_data_bin[396434:396476], wasm2goData_data_bin[396476:396518], wasm2goData_data_bin[396518:396720], wasm2goData_data_bin[396720:396786], wasm2goData_data_bin[396786:396805], wasm2goData_data_bin[396805:396833], wasm2goData_data_bin[396833:396867], wasm2goData_data_bin[396867:396940], wasm2goData_data_bin[396940:396948]}
	m.ThreadStart64 = Fn3046
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
	return Fn329(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn330(m)
}
func WasmShutdown(m *base.Module) {
	Fn331(m)
}
func DbgGemvQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn969(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn970(m, l0, l1, l2, l3, l4, l5, l6)
}
func WasiThreadStart(m *base.Module, l0 int32, l1 int64) {
	Fn3046(m, l0, l1)
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
	packed = Fn325(m, l0, l1)
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
	packed = Fn327(m, l0, l1)
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
	packed = Fn328(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
