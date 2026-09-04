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
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:127704], wasm2goData_data_bin[127704:147689], wasm2goData_data_bin[147689:147970], wasm2goData_data_bin[147970:148141], wasm2goData_data_bin[148141:148144], wasm2goData_data_bin[148144:149177], wasm2goData_data_bin[149177:149202], wasm2goData_data_bin[149202:149227], wasm2goData_data_bin[149227:149252], wasm2goData_data_bin[149252:149277], wasm2goData_data_bin[149277:149392], wasm2goData_data_bin[149392:149395], wasm2goData_data_bin[149395:149398], wasm2goData_data_bin[149398:149513], wasm2goData_data_bin[149513:149516], wasm2goData_data_bin[149516:149519], wasm2goData_data_bin[149519:149793], wasm2goData_data_bin[149793:150067], wasm2goData_data_bin[150067:150157], wasm2goData_data_bin[150157:150199], wasm2goData_data_bin[150199:150201], wasm2goData_data_bin[150201:150203], wasm2goData_data_bin[150203:194060], wasm2goData_data_bin[194060:194350], wasm2goData_data_bin[194350:194591], wasm2goData_data_bin[194591:194593], wasm2goData_data_bin[194593:194642], wasm2goData_data_bin[194642:194699], wasm2goData_data_bin[194699:231613], wasm2goData_data_bin[231613:231647], wasm2goData_data_bin[231647:231729], wasm2goData_data_bin[231729:237626], wasm2goData_data_bin[237626:262727], wasm2goData_data_bin[262727:263335], wasm2goData_data_bin[263335:265905], wasm2goData_data_bin[265905:267078], wasm2goData_data_bin[267078:267244], wasm2goData_data_bin[267244:268177], wasm2goData_data_bin[268177:268306], wasm2goData_data_bin[268306:268435], wasm2goData_data_bin[268435:268888], wasm2goData_data_bin[268888:269030], wasm2goData_data_bin[269030:269504], wasm2goData_data_bin[269504:269636], wasm2goData_data_bin[269636:277041], wasm2goData_data_bin[277041:277100], wasm2goData_data_bin[277100:277159], wasm2goData_data_bin[277159:277218], wasm2goData_data_bin[277218:277277], wasm2goData_data_bin[277277:277336], wasm2goData_data_bin[277336:277395], wasm2goData_data_bin[277395:277454], wasm2goData_data_bin[277454:277555], wasm2goData_data_bin[277555:277656], wasm2goData_data_bin[277656:277757], wasm2goData_data_bin[277757:277858], wasm2goData_data_bin[277858:277959], wasm2goData_data_bin[277959:278060], wasm2goData_data_bin[278060:278161], wasm2goData_data_bin[278161:278262], wasm2goData_data_bin[278262:278363], wasm2goData_data_bin[278363:278464], wasm2goData_data_bin[278464:278566], wasm2goData_data_bin[278566:278668], wasm2goData_data_bin[278668:358665], wasm2goData_data_bin[358665:382612], wasm2goData_data_bin[382612:382732], wasm2goData_data_bin[382732:383173], wasm2goData_data_bin[383173:383176], wasm2goData_data_bin[383176:398048], wasm2goData_data_bin[398048:398050], wasm2goData_data_bin[398050:400623], wasm2goData_data_bin[400623:400656], wasm2goData_data_bin[400656:400689], wasm2goData_data_bin[400689:400731], wasm2goData_data_bin[400731:400745], wasm2goData_data_bin[400745:400778], wasm2goData_data_bin[400778:400877], wasm2goData_data_bin[400877:401039], wasm2goData_data_bin[401039:401897], wasm2goData_data_bin[401897:402019], wasm2goData_data_bin[402019:402061], wasm2goData_data_bin[402061:402103], wasm2goData_data_bin[402103:402305], wasm2goData_data_bin[402305:402371], wasm2goData_data_bin[402371:402390], wasm2goData_data_bin[402390:402418], wasm2goData_data_bin[402418:402452], wasm2goData_data_bin[402452:402525], wasm2goData_data_bin[402525:402533]}
	m.ThreadStart64 = Fn2853
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
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:127704], wasm2goData_data_bin[127704:147689], wasm2goData_data_bin[147689:147970], wasm2goData_data_bin[147970:148141], wasm2goData_data_bin[148141:148144], wasm2goData_data_bin[148144:149177], wasm2goData_data_bin[149177:149202], wasm2goData_data_bin[149202:149227], wasm2goData_data_bin[149227:149252], wasm2goData_data_bin[149252:149277], wasm2goData_data_bin[149277:149392], wasm2goData_data_bin[149392:149395], wasm2goData_data_bin[149395:149398], wasm2goData_data_bin[149398:149513], wasm2goData_data_bin[149513:149516], wasm2goData_data_bin[149516:149519], wasm2goData_data_bin[149519:149793], wasm2goData_data_bin[149793:150067], wasm2goData_data_bin[150067:150157], wasm2goData_data_bin[150157:150199], wasm2goData_data_bin[150199:150201], wasm2goData_data_bin[150201:150203], wasm2goData_data_bin[150203:194060], wasm2goData_data_bin[194060:194350], wasm2goData_data_bin[194350:194591], wasm2goData_data_bin[194591:194593], wasm2goData_data_bin[194593:194642], wasm2goData_data_bin[194642:194699], wasm2goData_data_bin[194699:231613], wasm2goData_data_bin[231613:231647], wasm2goData_data_bin[231647:231729], wasm2goData_data_bin[231729:237626], wasm2goData_data_bin[237626:262727], wasm2goData_data_bin[262727:263335], wasm2goData_data_bin[263335:265905], wasm2goData_data_bin[265905:267078], wasm2goData_data_bin[267078:267244], wasm2goData_data_bin[267244:268177], wasm2goData_data_bin[268177:268306], wasm2goData_data_bin[268306:268435], wasm2goData_data_bin[268435:268888], wasm2goData_data_bin[268888:269030], wasm2goData_data_bin[269030:269504], wasm2goData_data_bin[269504:269636], wasm2goData_data_bin[269636:277041], wasm2goData_data_bin[277041:277100], wasm2goData_data_bin[277100:277159], wasm2goData_data_bin[277159:277218], wasm2goData_data_bin[277218:277277], wasm2goData_data_bin[277277:277336], wasm2goData_data_bin[277336:277395], wasm2goData_data_bin[277395:277454], wasm2goData_data_bin[277454:277555], wasm2goData_data_bin[277555:277656], wasm2goData_data_bin[277656:277757], wasm2goData_data_bin[277757:277858], wasm2goData_data_bin[277858:277959], wasm2goData_data_bin[277959:278060], wasm2goData_data_bin[278060:278161], wasm2goData_data_bin[278161:278262], wasm2goData_data_bin[278262:278363], wasm2goData_data_bin[278363:278464], wasm2goData_data_bin[278464:278566], wasm2goData_data_bin[278566:278668], wasm2goData_data_bin[278668:358665], wasm2goData_data_bin[358665:382612], wasm2goData_data_bin[382612:382732], wasm2goData_data_bin[382732:383173], wasm2goData_data_bin[383173:383176], wasm2goData_data_bin[383176:398048], wasm2goData_data_bin[398048:398050], wasm2goData_data_bin[398050:400623], wasm2goData_data_bin[400623:400656], wasm2goData_data_bin[400656:400689], wasm2goData_data_bin[400689:400731], wasm2goData_data_bin[400731:400745], wasm2goData_data_bin[400745:400778], wasm2goData_data_bin[400778:400877], wasm2goData_data_bin[400877:401039], wasm2goData_data_bin[401039:401897], wasm2goData_data_bin[401897:402019], wasm2goData_data_bin[402019:402061], wasm2goData_data_bin[402061:402103], wasm2goData_data_bin[402103:402305], wasm2goData_data_bin[402305:402371], wasm2goData_data_bin[402371:402390], wasm2goData_data_bin[402390:402418], wasm2goData_data_bin[402418:402452], wasm2goData_data_bin[402452:402525], wasm2goData_data_bin[402525:402533]}
	m.ThreadStart64 = Fn2853
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
	InitElemSeg_2_0(m)
	InitElemSeg_2_1(m)
	InitElemSeg_2_2(m)
	InitElemSeg_2_3(m)
	InitElemSeg_2_4(m)
	InitElemSeg_2_5(m)
	InitElemSeg_2_6(m)
	InitElemSeg_2_7(m)
	m.DataSegs = [][]byte{wasm2goData_data_bin[0:127704], wasm2goData_data_bin[127704:147689], wasm2goData_data_bin[147689:147970], wasm2goData_data_bin[147970:148141], wasm2goData_data_bin[148141:148144], wasm2goData_data_bin[148144:149177], wasm2goData_data_bin[149177:149202], wasm2goData_data_bin[149202:149227], wasm2goData_data_bin[149227:149252], wasm2goData_data_bin[149252:149277], wasm2goData_data_bin[149277:149392], wasm2goData_data_bin[149392:149395], wasm2goData_data_bin[149395:149398], wasm2goData_data_bin[149398:149513], wasm2goData_data_bin[149513:149516], wasm2goData_data_bin[149516:149519], wasm2goData_data_bin[149519:149793], wasm2goData_data_bin[149793:150067], wasm2goData_data_bin[150067:150157], wasm2goData_data_bin[150157:150199], wasm2goData_data_bin[150199:150201], wasm2goData_data_bin[150201:150203], wasm2goData_data_bin[150203:194060], wasm2goData_data_bin[194060:194350], wasm2goData_data_bin[194350:194591], wasm2goData_data_bin[194591:194593], wasm2goData_data_bin[194593:194642], wasm2goData_data_bin[194642:194699], wasm2goData_data_bin[194699:231613], wasm2goData_data_bin[231613:231647], wasm2goData_data_bin[231647:231729], wasm2goData_data_bin[231729:237626], wasm2goData_data_bin[237626:262727], wasm2goData_data_bin[262727:263335], wasm2goData_data_bin[263335:265905], wasm2goData_data_bin[265905:267078], wasm2goData_data_bin[267078:267244], wasm2goData_data_bin[267244:268177], wasm2goData_data_bin[268177:268306], wasm2goData_data_bin[268306:268435], wasm2goData_data_bin[268435:268888], wasm2goData_data_bin[268888:269030], wasm2goData_data_bin[269030:269504], wasm2goData_data_bin[269504:269636], wasm2goData_data_bin[269636:277041], wasm2goData_data_bin[277041:277100], wasm2goData_data_bin[277100:277159], wasm2goData_data_bin[277159:277218], wasm2goData_data_bin[277218:277277], wasm2goData_data_bin[277277:277336], wasm2goData_data_bin[277336:277395], wasm2goData_data_bin[277395:277454], wasm2goData_data_bin[277454:277555], wasm2goData_data_bin[277555:277656], wasm2goData_data_bin[277656:277757], wasm2goData_data_bin[277757:277858], wasm2goData_data_bin[277858:277959], wasm2goData_data_bin[277959:278060], wasm2goData_data_bin[278060:278161], wasm2goData_data_bin[278161:278262], wasm2goData_data_bin[278262:278363], wasm2goData_data_bin[278363:278464], wasm2goData_data_bin[278464:278566], wasm2goData_data_bin[278566:278668], wasm2goData_data_bin[278668:358665], wasm2goData_data_bin[358665:382612], wasm2goData_data_bin[382612:382732], wasm2goData_data_bin[382732:383173], wasm2goData_data_bin[383173:383176], wasm2goData_data_bin[383176:398048], wasm2goData_data_bin[398048:398050], wasm2goData_data_bin[398050:400623], wasm2goData_data_bin[400623:400656], wasm2goData_data_bin[400656:400689], wasm2goData_data_bin[400689:400731], wasm2goData_data_bin[400731:400745], wasm2goData_data_bin[400745:400778], wasm2goData_data_bin[400778:400877], wasm2goData_data_bin[400877:401039], wasm2goData_data_bin[401039:401897], wasm2goData_data_bin[401897:402019], wasm2goData_data_bin[402019:402061], wasm2goData_data_bin[402061:402103], wasm2goData_data_bin[402103:402305], wasm2goData_data_bin[402305:402371], wasm2goData_data_bin[402371:402390], wasm2goData_data_bin[402390:402418], wasm2goData_data_bin[402418:402452], wasm2goData_data_bin[402452:402525], wasm2goData_data_bin[402525:402533]}
	m.ThreadStart64 = Fn2853
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
	Fn895(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ5_0_q8_0(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn922(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ4KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn927(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgVecDotQ6KQ8K(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int32) {
	Fn929(m, l0, l1, l2, l3, l4, l5, l6, l7)
}
func DbgQuantizeMatQ8_0_4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn798(m, l0, l1, l2)
}
func DbgQuantizeMatQ8K4x8(m *base.Module, l0 int64, l1 int64, l2 int64) {
	Fn800(m, l0, l1, l2)
}
func DbgGemvQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn805(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ4K8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn816(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn933(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemvQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn931(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ5_0_8x8(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn934(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgGemmQ8_0_4x4(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int32, l6 int32) {
	Fn932(m, l0, l1, l2, l3, l4, l5, l6)
}
func DbgVecSwigluF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	Fn896(m, l0, l1, l2, l3)
}
func DbgVecSoftMaxF32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) float64 {
	return Fn897(m, l0, l1, l2, l3)
}
func DbgVecMadF16F32(m *base.Module, l0 int32, l1 int64, l2 int64, l3 float32) {
	Fn898(m, l0, l1, l2, l3)
}
func DbgSimdGemmF32(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32, l4 int32, l5 int32) {
	Fn899(m, l0, l1, l2, l3, l4, l5)
}
func DbgFlashAttnKvF16(m *base.Module, l0 int64) {
	Fn907(m, l0)
}
func WasiThreadStart(m *base.Module, l0 int32, l1 int64) {
	Fn2853(m, l0, l1)
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
