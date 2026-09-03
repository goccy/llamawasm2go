package base

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"math"
	"math/bits"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

type Wasi_snapshot_preview1Imports interface {
	Environ_get64(m *Module, l0 int64, l1 int64) int32
	Environ_sizes_get64(m *Module, l0 int64, l1 int64) int32
	Clock_time_get64(m *Module, l0 int64, l1 int64, l2 int64) int32
	Fd_close64(m *Module, l0 int64) int32
	Fd_fdstat_get64(m *Module, l0 int64, l1 int64) int32
	Fd_fdstat_set_flags64(m *Module, l0 int64, l1 int64) int32
	Fd_prestat_get64(m *Module, l0 int64, l1 int64) int32
	Fd_prestat_dir_name64(m *Module, l0 int64, l1 int64, l2 int64) int32
	Fd_read64(m *Module, l0 int64, l1 int64, l2 int64, l3 int64) int32
	Fd_readdir64(m *Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32
	Fd_seek64(m *Module, l0 int64, l1 int64, l2 int64, l3 int64) int32
	Fd_write64(m *Module, l0 int64, l1 int64, l2 int64, l3 int64) int32
	Path_filestat_get64(m *Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64) int32
	Path_open64(m *Module, l0 int64, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64) int32
	Proc_exit64(m *Module, l0 int64)
	Sched_yield64(m *Module) int32
	Random_get64(m *Module, l0 int64, l1 int64) int32
}
type EnvImports interface {
}
type WasmifyImports interface {
	Callback_invoke(m *Module, l0 int32, l1 int32, l2 int64, l3 int64) int64
}
type Module struct {
	Memory                 []byte
	MaxMem                 uint64
	M                      unsafe.Pointer
	OutlinePack            [128]uint64
	ExcPending             int32
	ExcTag                 uint32
	ExcVals                [1]uint64
	T0                     []any
	G0                     int64
	G1                     int64
	Wasi_snapshot_preview1 Wasi_snapshot_preview1Imports
	Env                    EnvImports
	Wasmify                WasmifyImports
	MemMu                  *sync.Mutex
	MemSize                *atomic.Uint64
	DataSegs               [][]byte
	DataEnd                uint32
	MemShared              bool
	Threads                *ThreadPool
	ThreadStart64          func(*Module, int32, int64)
}

func I32(x int32) int32 { return x }

func I64(x int64) int64 { return x }

// ui32 / ui64 reinterpret a signed integer as its unsigned bit
// equivalent at runtime. Used for the operands of wasm unsigned
// comparisons (i32.lt_u etc.) — emitting `uint32(int32(-N))` directly
// fails Go's compile-time constant rule because the negative typed
// constant isn't representable in uint32; routing through these
// function-call boundaries forces runtime conversion.
func Ui32(x int32) uint32 { return uint32(x) }

func Ui64(x int64) uint64 { return uint64(x) }

// b2i32 materialises a wasm comparison result — an i32 that is 0 or 1 — from
// the Go bool the comparison expression evaluates to.
//
// It exists as a named helper rather than an inline `func() int32 { ... }()`
// because the gcasm backend requires every direct call left in the compiled
// output to be either a package-local FnN or something the Go inliner removed.
// A func literal is normally inlined at its call site, but the inliner gives up
// once the ENCLOSING function grows past its budget — and a single wasm function
// can translate to tens of thousands of lines of Go, as an interpreter's
// bytecode dispatch loop does. The literal is then outlined into a real closure
// symbol (FnN.funcA.funcB), which reaches the assembler as a direct call gcasm
// cannot marshal. A named helper this small is always inlined, and if it ever
// were not, it would fail loudly at its own symbol rather than as a nested
// closure.
func B2i32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func F32(x float32) float32 { runtime.KeepAlive(&x); return x }

func F64(x float64) float64 { runtime.KeepAlive(&x); return x }

//go:noinline
func Wasm_trap_div_zero() { panic("wasm: integer divide by zero") }

//go:noinline
func Wasm_trap_int_overflow() { panic("wasm: integer overflow") }

//go:noinline
func Wasm_trap_invalid_conv() { panic("wasm: invalid conversion to integer") }

//go:noinline
func Wasm_trap_unreachable() { panic("wasm: unreachable") }

//go:noinline
func Wasm_trap_memfill_oob() { panic("wasm: memory.fill out of bounds") }

//go:noinline
func Wasm_trap_memcopy_oob() { panic("wasm: memory.copy out of bounds") }

//go:noinline
func Wasm_trap_meminit_oob() { panic("wasm: memory.init out of bounds") }

func I32_div_s(x, y int32) int32 {
	if y == -1 && x == math.MinInt32 {
		Wasm_trap_int_overflow()
	}
	if y == 0 {
		Wasm_trap_div_zero()
	}
	return x / y
}

func I64_div_s(x, y int64) int64 {
	if y == -1 && x == math.MinInt64 {
		Wasm_trap_int_overflow()
	}
	if y == 0 {
		Wasm_trap_div_zero()
	}
	return x / y
}

func I32_div_u(x, y uint32) uint32 {
	if y == 0 {
		Wasm_trap_div_zero()
	}
	return x / y
}

func I64_div_u(x, y uint64) uint64 {
	if y == 0 {
		Wasm_trap_div_zero()
	}
	return x / y
}

func I32_rem_s(x, y int32) int32 {
	if y == 0 {
		Wasm_trap_div_zero()
	}
	if y == -1 {

		return 0
	}
	return x % y
}

func I64_rem_s(x, y int64) int64 {
	if y == 0 {
		Wasm_trap_div_zero()
	}
	if y == -1 {
		return 0
	}
	return x % y
}

func I32_rem_u(x, y uint32) uint32 {
	if y == 0 {
		Wasm_trap_div_zero()
	}
	return x % y
}

func I64_rem_u(x, y uint64) uint64 {
	if y == 0 {
		Wasm_trap_div_zero()
	}
	return x % y
}

func I32_rotl(x, y int32) int32 { return int32(bits.RotateLeft32(uint32(x), int(y&31))) }

func I32_rotr(x, y int32) int32 { return int32(bits.RotateLeft32(uint32(x), -int(y&31))) }

func I64_rotl(x, y int64) int64 { return int64(bits.RotateLeft64(uint64(x), int(y&63))) }

func I64_rotr(x, y int64) int64 { return int64(bits.RotateLeft64(uint64(x), -int(y&63))) }

func F32_min(x, y float32) float32 {
	if x != x || y != y {
		return float32(math.NaN())
	}
	if x < y {
		return x
	}
	if y < x {
		return y
	}

	if x == 0 {
		if math.Signbit(float64(x)) {
			return x
		}
		return y
	}
	return x
}

func F32_max(x, y float32) float32 {
	if x != x || y != y {
		return float32(math.NaN())
	}
	if x > y {
		return x
	}
	if y > x {
		return y
	}
	if x == 0 {
		if math.Signbit(float64(x)) {
			return y
		}
		return x
	}
	return x
}

func F32_abs(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}

func F64_abs(x float64) float64 {
	return math.Float64frombits(math.Float64bits(x) &^ (1 << 63))
}

func F32_neg(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) ^ (1 << 31))
}

func F64_neg(x float64) float64 {
	return math.Float64frombits(math.Float64bits(x) ^ (1 << 63))
}

func F32_copysign(x, y float32) float32 {
	return float32(math.Copysign(float64(x), float64(y)))
}

func F64_copysign(x, y float64) float64 { return math.Copysign(x, y) }

func I32_trunc_sat_f32_s(x float32) int32 {
	if x != x {
		return 0
	}
	if x <= -2147483648.0 {
		return math.MinInt32
	}
	if x >= 2147483648.0 {
		return math.MaxInt32
	}
	return int32(x)
}

func I32_trunc_sat_f32_u(x float32) int32 {
	if x != x || x <= 0 {
		return 0
	}
	if x >= 4294967296.0 {
		return -1
	}
	return int32(uint32(x))
}

func I32_trunc_sat_f64_s(x float64) int32 {
	if x != x {
		return 0
	}
	if x <= -2147483648.0 {
		return math.MinInt32
	}
	if x >= 2147483648.0 {
		return math.MaxInt32
	}
	return int32(x)
}

func I32_trunc_sat_f64_u(x float64) int32 {
	if x != x || x <= 0 {
		return 0
	}
	if x >= 4294967296.0 {
		return -1
	}
	return int32(uint32(x))
}

func I64_trunc_sat_f32_s(x float32) int64 {
	if x != x {
		return 0
	}
	if float64(x) <= -9223372036854775808.0 {
		return math.MinInt64
	}
	if float64(x) >= 9223372036854775808.0 {
		return math.MaxInt64
	}
	return int64(x)
}

func I64_trunc_sat_f32_u(x float32) int64 {
	if x != x || x <= 0 {
		return 0
	}
	if float64(x) >= 18446744073709551616.0 {
		return -1
	}
	return int64(uint64(x))
}

func I64_trunc_sat_f64_s(x float64) int64 {
	if x != x {
		return 0
	}
	if x <= -9223372036854775808.0 {
		return math.MinInt64
	}
	if x >= 9223372036854775808.0 {
		return math.MaxInt64
	}
	return int64(x)
}

func I64_trunc_sat_f64_u(x float64) int64 {
	if x != x || x <= 0 {
		return 0
	}
	if x >= 18446744073709551616.0 {
		return -1
	}
	return int64(uint64(x))
}

// accessMemory runs f with the module's current linear memory while
// holding the same lock memoryGrow takes to mutate the memory slice
// header or relocate its backing array. It is the ONE safe way to
// touch linear memory from OUTSIDE the module's execution goroutine —
// e.g. a watchdog goroutine raising CPython's eval-breaker bit while
// an evaluation is running. For the duration of f the memory can
// neither be resliced nor relocated, so f's writes land in the array
// the guest observes; a grow that raced in just before blocks until f
// returns and then copies f's writes forward with the rest of the
// contents. Determinism notes for callers:
//
//   - f MUST NOT call back into the module or into memoryGrow — that
//     would self-deadlock.
//   - f should be short: a running guest blocks inside memory.grow
//     until f returns (ordinary guest loads/stores do not block).
//   - Bytes the guest reads or writes concurrently with f (that is
//     the point of an eval-breaker-style flag) are exchanged with
//     plain single-word accesses; keep such shared words
//     word-aligned and word-sized.
func AccessMemory(m *Module, f func(mem []byte)) {
	m.MemMu.Lock()
	defer m.MemMu.Unlock()
	f(m.Memory)
}

func I32_div_u_s(x, y int32) int32 { return int32(I32_div_u(uint32(x), uint32(y))) }
func I32_rem_u_s(x, y int32) int32 { return int32(I32_rem_u(uint32(x), uint32(y))) }
func I64_div_u_s(x, y int64) int64 { return int64(I64_div_u(uint64(x), uint64(y))) }
func I64_rem_u_s(x, y int64) int64 { return int64(I64_rem_u(uint64(x), uint64(y))) }

// The explicit same-type conversions are NOT redundant: they are
// rounding points. Once these helpers inline, gc is free to fuse a
// multiply feeding an add into a single FMA — legal Go, but wasm
// requires every operation individually rounded, and a fused result
// diverges from every wasm runtime (bitwise, and observably in greedy
// sampling). A float conversion forces the intermediate rounding and
// forbids the fusion (spec: Conversions, "rounds to the precision of
// the target type"; the same rule math.FMA documents).
func F32_add(x, y float32) float32 { return float32(x + y) }
func F32_sub(x, y float32) float32 { return float32(x - y) }
func F32_mul(x, y float32) float32 { return float32(x * y) }
func F32_div(x, y float32) float32 { return float32(x / y) }
func F64_add(x, y float64) float64 { return float64(x + y) }
func F64_sub(x, y float64) float64 { return float64(x - y) }
func F64_mul(x, y float64) float64 { return float64(x * y) }
func F64_div(x, y float64) float64 { return float64(x / y) }

func I32_clz(x int32) int32    { return int32(bits.LeadingZeros32(uint32(x))) }
func I32_ctz(x int32) int32    { return int32(bits.TrailingZeros32(uint32(x))) }
func I32_popcnt(x int32) int32 { return int32(bits.OnesCount32(uint32(x))) }

func I64_clz(x int64) int64    { return int64(bits.LeadingZeros64(uint64(x))) }
func I64_ctz(x int64) int64    { return int64(bits.TrailingZeros64(uint64(x))) }
func I64_popcnt(x int64) int64 { return int64(bits.OnesCount64(uint64(x))) }

func F32_ceil(x float32) float32  { return float32(math.Ceil(float64(x))) }
func F64_ceil(x float64) float64  { return math.Ceil(x) }
func F32_floor(x float32) float32 { return float32(math.Floor(float64(x))) }
func F64_floor(x float64) float64 { return math.Floor(x) }
func F32_trunc(x float32) float32 { return float32(math.Trunc(float64(x))) }
func F64_trunc(x float64) float64 { return math.Trunc(x) }
func F32_sqrt(x float32) float32  { return float32(math.Sqrt(float64(x))) }
func F64_sqrt(x float64) float64  { return math.Sqrt(x) }

func F32_eq(x, y float32) int32 {
	if x == y {
		return 1
	}
	return 0
}
func F32_ne(x, y float32) int32 {
	if x != y {
		return 1
	}
	return 0
}
func F32_lt(x, y float32) int32 {
	if x < y {
		return 1
	}
	return 0
}
func F32_gt(x, y float32) int32 {
	if x > y {
		return 1
	}
	return 0
}
func F32_le(x, y float32) int32 {
	if x <= y {
		return 1
	}
	return 0
}
func F32_ge(x, y float32) int32 {
	if x >= y {
		return 1
	}
	return 0
}

func F64_eq(x, y float64) int32 {
	if x == y {
		return 1
	}
	return 0
}
func F64_ne(x, y float64) int32 {
	if x != y {
		return 1
	}
	return 0
}
func F64_lt(x, y float64) int32 {
	if x < y {
		return 1
	}
	return 0
}
func F64_gt(x, y float64) int32 {
	if x > y {
		return 1
	}
	return 0
}
func F64_le(x, y float64) int32 {
	if x <= y {
		return 1
	}
	return 0
}
func F64_ge(x, y float64) int32 {
	if x >= y {
		return 1
	}
	return 0
}

func I32_wrap_i64(x int64) int32       { return int32(x) }
func I64_extend_i32_s(x int32) int64   { return int64(x) }
func I64_extend_i32_u(x int32) int64   { return int64(uint32(x)) }
func F32_demote_f64(x float64) float32 { return float32(x) }
func F64_promote_f32(x float32) float64 {

	if math.IsNaN(float64(x)) {

		return float64(x)
	}
	return float64(x)
}

func F32_convert_i32_s(x int32) float32 { return float32(x) }
func F32_convert_i32_u(x int32) float32 { return float32(uint32(x)) }
func F32_convert_i64_s(x int64) float32 { return float32(x) }
func F32_convert_i64_u(x int64) float32 { return float32(uint64(x)) }
func F64_convert_i32_s(x int32) float64 { return float64(x) }
func F64_convert_i32_u(x int32) float64 { return float64(uint32(x)) }
func F64_convert_i64_s(x int64) float64 { return float64(x) }
func F64_convert_i64_u(x int64) float64 { return float64(uint64(x)) }

func I32_reinterpret_f32(x float32) int32 { return int32(math.Float32bits(x)) }
func I64_reinterpret_f64(x float64) int64 { return int64(math.Float64bits(x)) }
func F32_reinterpret_i32(x int32) float32 { return math.Float32frombits(uint32(x)) }
func F64_reinterpret_i64(x int64) float64 { return math.Float64frombits(uint64(x)) }

func I32_extend8_s(x int32) int32  { return int32(int8(x)) }
func I32_extend16_s(x int32) int32 { return int32(int16(x)) }
func I64_extend8_s(x int64) int64  { return int64(int8(x)) }
func I64_extend16_s(x int64) int64 { return int64(int16(x)) }
func I64_extend32_s(x int64) int64 { return int64(int32(x)) }

// dataDrop implements data.drop: discard passive segment seg. A later
// memory.init naming it traps (nil view); double-drop is a no-op per spec.
// dataDrop stays out of line: inlined into a gcasm-transformed function, the
// pointer write (a nil store into dataSegs) would drag runtime.gcWriteBarrier
// into the asm body, which the transformer rejects.
//
//go:noinline
func DataDrop(m *Module, seg int) {
	m.DataSegs[seg] = nil
}

//go:noinline
func Wasm_trap_atomic_oob() { panic("wasm: atomic access out of bounds") }

//go:noinline
func Wasm_trap_atomic_unaligned() { panic("wasm: unaligned atomic access") }

//go:noinline
func Wasm_trap_atomic_wait_forever() {
	panic("wasm: blocking atomic wait with no other agents (wasi-threads not enabled)")
}

// atomicEA64 is atomicEA for a memory64 memory: i64 address and offset with
// overflow-safe u64 effective-address arithmetic (mirroring simdEA64 — a
// negative addr or offset becomes a huge u64 that either wraps the addition,
// caught by the wrap checks, or fails the bounds check), then the same
// natural-alignment check the proposal requires.
//
//go:noinline
func AtomicEA64(m *Module, addr int64, offset int64, size uint64) uint64 {
	ea := uint64(addr) + uint64(offset)
	end := ea + size
	if ea < uint64(addr) || end < ea || end > m.MemSize.Load() {
		Wasm_trap_atomic_oob()
	}
	if ea&(size-1) != 0 {
		Wasm_trap_atomic_unaligned()
	}
	return ea
}

// atomicPtr32At / atomicPtr64At turn a CHECKED effective address into a
// pointer into linear memory. The caller went through atomicEA/atomicEA64,
// which already bounds-checked ea against memSize, so index off the raw
// base pointer to skip Go's redundant slice bounds check — the same deal
// the plain load/store path gets. m.M tracks m.memory's data pointer (New
// sets it; a shared memory never relocates, and the non-shared reallocate
// path refreshes it).
//
//go:noinline
func AtomicPtr32At(m *Module, ea uint64) *uint32 {
	return (*uint32)(unsafe.Add(m.M, uintptr(ea)))
}

//go:noinline
func AtomicPtr64At(m *Module, ea uint64) *uint64 {
	return (*uint64)(unsafe.Add(m.M, uintptr(ea)))
}

// atomicsContended reports whether more than the main agent can touch the
// memory — i.e. at least one wasi thread has been spawned. Until that happens
// the engine's own atomic ops (interrupt-flag reads, GC bookkeeping) have no
// peer to race, so store/RMW helpers take an ordinary read-modify-write
// instead of a LOCKed one. The 0->1 transition happens inside threadSpawn on
// the sole agent, and the `go` statement that starts the child publishes
// every prior non-atomic write to it, so the fast path is race-free.
func AtomicsContended(m *Module) bool {
	return m.Threads != nil && m.Threads.nextTID.Load() != 0
}

// atomicSubword32 runs op on the byte lanes [shift, shift+bits) of the
// aligned 32-bit word containing ea, via a CAS loop; returns the OLD lane
// value zero-extended. Little-endian lane math.
//
//go:noinline
func AtomicSubword32(m *Module, ea uint64, bits uint, op func(old uint32) uint32) uint32 {
	word := (*uint32)(unsafe.Add(m.M, uintptr(ea&^3)))
	shift := uint(ea&3) * 8
	mask := uint32(1)<<bits - 1
	if !AtomicsContended(m) {
		cur := *word
		lane := (cur >> shift) & mask
		*word = (cur &^ (mask << shift)) | ((op(lane) & mask) << shift)
		return lane
	}
	for {
		cur := atomic.LoadUint32(word)
		lane := (cur >> shift) & mask
		next := (cur &^ (mask << shift)) | ((op(lane) & mask) << shift)
		if atomic.CompareAndSwapUint32(word, cur, next) {
			return lane
		}
	}
}

//go:noinline
func AtomicRmwAdd32At(m *Module, ea uint64, v int32) int32 {
	p := AtomicPtr32At(m, ea)
	if !AtomicsContended(m) {
		old := *p
		*p = old + uint32(v)
		return int32(old)
	}
	return int32(atomic.AddUint32(p, uint32(v)) - uint32(v))
}

//go:noinline
func AtomicRmwXchg32At(m *Module, ea uint64, v int32) int32 {
	p := AtomicPtr32At(m, ea)
	if !AtomicsContended(m) {
		old := *p
		*p = uint32(v)
		return int32(old)
	}
	return int32(atomic.SwapUint32(p, uint32(v)))
}

//go:noinline
func AtomicRmwCmpxchg32At(m *Module, ea uint64, expected, replacement int32) int32 {
	p := AtomicPtr32At(m, ea)
	if !AtomicsContended(m) {
		cur := *p
		if cur == uint32(expected) {
			*p = uint32(replacement)
		}
		return int32(cur)
	}
	for {
		cur := atomic.LoadUint32(p)
		if cur != uint32(expected) {
			return int32(cur)
		}
		if atomic.CompareAndSwapUint32(p, cur, uint32(replacement)) {
			return int32(cur)
		}
	}
}

//go:noinline
func AtomicRmwAdd64At(m *Module, ea uint64, v int64) int64 {
	p := AtomicPtr64At(m, ea)
	if !AtomicsContended(m) {
		old := *p
		*p = old + uint64(v)
		return int64(old)
	}
	return int64(atomic.AddUint64(p, uint64(v)) - uint64(v))
}

// spinRelax is the cold half of the preemption guard the emitters
// plant in bare atomic spin loops (a loop that waits on an inline
// atomic load and makes no other call — see spinguard.go). Such a loop
// is fine as Go, but once the gcasm bundler captures the compiled
// function into a .s TEXT the runtime can no longer async-preempt it,
// and a goroutine spinning there blocks every stop-the-world — a
// livelock when the store it waits for comes from a goroutine the GC
// already parked.
//
// The generated hot path is a counter increment and a not-taken
// branch; every 2^k-th iteration reaches this call, with k derived at
// emission from the loop body's size so the interval is a roughly
// constant TIME budget (see spinGuardMask). The call itself
// is the fix — it must survive to machine code (hence //go:noinline),
// and its prologue's stack check is the preemption point, so a
// stop-the-world waits at most tens-to-low-hundreds of microseconds
// of spinning. The Gosched additionally donates the core when a
// wait is genuinely long. Calling on every iteration instead measured
// ~40% decode overhead at n_threads=8: eight workers reaching
// runtime.Gosched at spin rate serialize on sched.lock, and the call
// round-trip alone showed ~15%.
//
// The Gosched is rate-limited across every spinning worker (the
// spinRelaxColdCalls counter lives in the runtime template — helper
// extraction carries function decls only): the preemption point is the
// spinRelax call itself (its prologue's stack check), but yielding on
// every cold call still measured double-digit scheduler churn
// (pthread_cond_signal — wakep — at 30% of the profile) on
// barrier-heavy workloads, where waiting IS most of a worker's time.
// One yield per 64 cold calls keeps donation proportional to
// aggregate spin time — rare on an uncontended box, automatically
// more frequent when oversubscription makes the spins long.
//
//go:noinline
func SpinRelax() {
	if atomic.AddUint32(&spinRelaxColdCalls, 1)&63 == 0 {
		runtime.Gosched()
	}
}

// atomicWait32/64 implement memory.atomic.wait: compare-and-park. The compare
// happens under the parking-lot lock a notifier must also take, so a notify
// that lands between the compare and the park cannot be missed.
//
// Returns 0 = woken, 1 = not-equal, 2 = timed-out. A negative timeout means
// wait forever; with no other agent able to notify, that is a guaranteed
// deadlock, so it traps rather than hanging the process.
//
//go:noinline
func AtomicWait32At(m *Module, ea uint64, expected int32, timeout int64) int32 {
	p := AtomicPtr32At(m, ea)
	return AtomicWait(m, ea, timeout, func() bool {
		return int32(atomic.LoadUint32(p)) == expected
	})
}

//go:noinline
func AtomicWait(m *Module, ea uint64, timeout int64, stillEqual func() bool) int32 {
	if !m.MemShared {

		return 1
	}
	m.Threads.parkMu.Lock()
	if !stillEqual() {
		m.Threads.parkMu.Unlock()
		return 1
	}
	ch := make(chan struct{})
	if m.Threads.parked == nil {
		m.Threads.parked = make(map[uint64][]chan struct{})
	}
	m.Threads.parked[ea] = append(m.Threads.parked[ea], ch)
	m.Threads.parkMu.Unlock()

	unpark := func() {
		m.Threads.parkMu.Lock()
		defer m.Threads.parkMu.Unlock()
		waiters := m.Threads.parked[ea]
		for i, c := range waiters {
			if c == ch {
				m.Threads.parked[ea] = append(waiters[:i], waiters[i+1:]...)
				break
			}
		}
		if len(m.Threads.parked[ea]) == 0 {
			delete(m.Threads.parked, ea)
		}
	}

	if timeout < 0 {
		if m.Threads.nextTID.Load() == 0 {

			unpark()
			Wasm_trap_atomic_wait_forever()
		}
		<-ch
		return 0
	}

	timer := time.NewTimer(time.Duration(timeout) + time.Millisecond)
	defer timer.Stop()
	select {
	case <-ch:
		return 0
	case <-timer.C:
		unpark()
		return 2
	}
}

//go:noinline
func AtomicLoad32_8u_m64(m *Module, addr int64, offset int64) int32 {
	return int32(AtomicSubword32(m, AtomicEA64(m, addr, offset, 1), 8, func(old uint32) uint32 { return old }))
}

//go:noinline
func AtomicStore32_8_m64(m *Module, addr int64, offset int64, v int32) int32 {
	AtomicSubword32(m, AtomicEA64(m, addr, offset, 1), 8, func(uint32) uint32 { return uint32(v) })
	return 0
}

//go:noinline
func AtomicRmwAdd32_m64(m *Module, addr, offset int64, v int32) int32 {
	return AtomicRmwAdd32At(m, AtomicEA64(m, addr, offset, 4), v)
}

//go:noinline
func AtomicRmwXchg32_m64(m *Module, addr, offset int64, v int32) int32 {
	return AtomicRmwXchg32At(m, AtomicEA64(m, addr, offset, 4), v)
}

//go:noinline
func AtomicRmwCmpxchg32_m64(m *Module, addr, offset int64, expected, replacement int32) int32 {
	return AtomicRmwCmpxchg32At(m, AtomicEA64(m, addr, offset, 4), expected, replacement)
}

//go:noinline
func AtomicRmwAdd64_m64(m *Module, addr, offset int64, v int64) int64 {
	return AtomicRmwAdd64At(m, AtomicEA64(m, addr, offset, 8), v)
}

//go:noinline
func AtomicNotify_m64(m *Module, addr int64, offset int64, count int32) int32 {
	return m.Threads.wake(AtomicEA64(m, addr, offset, 4), count)
}

//go:noinline
func AtomicWait32_m64(m *Module, addr int64, offset int64, expected int32, timeout int64) int32 {
	return AtomicWait32At(m, AtomicEA64(m, addr, offset, 4), expected, timeout)
}

// threadLaunch allocates a TID and runs body — the guest's thread entry
// already bound to its start argument — on a fresh goroutine-agent.
//
//go:noinline
func ThreadLaunch(m *Module, body func(child *Module, tid int32)) int32 {
	tid := m.Threads.nextTID.Add(1)
	m.Threads.wg.Add(1)

	child := new(Module)
	*child = *m
	go func() {
		defer m.Threads.wg.Done()

		defer func() {
			if r := recover(); r != nil {
				println("wasm2go: wasi thread", tid, "trapped:")
				switch v := r.(type) {
				case error:
					println("  ", v.Error())
				case string:
					println("  ", v)
				}
				panic(r)
			}
		}()
		body(child, tid)
	}()
	return tid
}

// threadSpawn_m64 is threadSpawn for a memory64 module: the guest's
// start_arg is a linear-memory pointer and therefore an i64. The TID
// stays an i32 — wasi_thread_spawn returns i32 in both widths.
//
//go:noinline
func ThreadSpawn_m64(m *Module, arg int64) int32 {
	start := m.ThreadStart64
	if start == nil {
		return -1
	}
	return ThreadLaunch(m, func(child *Module, tid int32) { start(child, tid, arg) })
}

//go:noinline
func Wasm_trap_simd_oob() { panic("wasm: v128 memory access out of bounds") }

// gcasmMemProbe anchors the Module field offsets the gcasm memory-op
// splices hardcode. The splices read m.M and m.memSize straight off the
// receiver in generated assembly, and the offsets of those fields
// depend on the module (the import-interface fields between them vary).
// Rather than re-deriving Go's struct layout, gcasm extracts the two
// offsets from THIS function's captured assembly — two loads off R0/AX,
// M first — so they always come from the same compile that produced the
// code being spliced. Never called at run time.
//
//go:noinline
func GcasmMemProbe(m *Module) (unsafe.Pointer, *atomic.Uint64) {
	return m.M, m.MemSize
}

//go:noinline
func Simd_scalar_i32_load16_u(m *Module, addr int32) int32 {
	return int32(*(*uint16)(unsafe.Add(m.M, uintptr(uint32(addr)))))
}

//go:noinline
func Simd_scalar_f32_load(m *Module, addr int32) float32 {
	return *(*float32)(unsafe.Add(m.M, uintptr(uint32(addr))))
}

//go:noinline
func Simd_scalar_i32_shl(v int32, s int32) int32 { return v << (uint(s) % 32) }

//go:noinline
func Simd_scalar_i32_add(a int32, b int32) int32 { return a + b }

//go:noinline
func Simd_scalar_f32_mul(a float32, b float32) float32 { return a * b }

//go:noinline
func Simd_m64_scalar_i32_load16_u(m *Module, addr int64) int64 {
	return int64(*(*uint16)(unsafe.Add(m.M, uintptr(uint64(addr)))))
}

//go:noinline
func Simd_m64_scalar_f32_load(m *Module, addr int64) float32 {
	return *(*float32)(unsafe.Add(m.M, uintptr(uint64(addr))))
}

//go:noinline
func Simd_m64_scalar_i32_shl(v int64, s int64) int64 { return v << (uint(s) % 64) }

//go:noinline
func Simd_m64_scalar_i32_add(a int64, b int64) int64 { return a + b }

// f16BitsToF32Bits is the IEEE binary16 -> binary32 conversion,
// bit-exact including subnormals, infinities and NaN payloads.
func F16BitsToF32Bits(h uint16) uint32 {
	sign := uint32(h>>15) << 31
	exp := uint32(h>>10) & 0x1F
	man := uint32(h) & 0x3FF
	switch exp {
	case 0:
		if man == 0 {
			return sign
		}
		e := uint32(113)
		for man&0x400 == 0 {
			man <<= 1
			e--
		}
		return sign | e<<23 | (man&0x3FF)<<13
	case 0x1F:
		return sign | 0xFF<<23 | man<<13
	}
	return sign | (exp+112)<<23 | man<<13
}

// simd_f16x4_cvt converts four f16 values (widened to the low 16 bits
// of each i32x4 lane, the v128_load16x4_u result shape) to f32 lanes.
// Emitted only after the transpiler verified the module's conversion
// table is the IEEE map, so this computed conversion is bit-identical
// to the table reads it replaces (and to XTN+FCVTL on arm64).
//
//go:noinline
func Simd_f16x4_cvt(v [2]uint64) [2]uint64 {
	var out [2]uint64
	for i := 0; i < 4; i++ {
		bits := uint16(v[i/2] >> (32 * uint(i) % 64))
		out[i/2] |= uint64(F16BitsToF32Bits(bits)) << (32 * uint(i) % 64)
	}
	return out
}

func Simd_p_pack(lo, hi uint64) [2]uint64 { return [2]uint64{lo, hi} }

// mem64HardCap bounds a memory64's linear memory. Far above any real
// host, it exists so effective-address arithmetic in the checked
// helpers keeps a comfortable non-overflow margin below 2^64.
func Mem64HardCap() uint64 { return 1 << 48 }

// memorySize64 returns the current size in wasm pages as i64.
func MemorySize64(m *Module) int64 {
	return int64(m.MemSize.Load() >> 16)
}

// memoryGrow64 is memoryGrow for a memory64: i64 page delta, i64
// previous-page-count result, -1 on failure. Growth uses the same
// geometric-capacity scheme; the wasm32 hard cap does not apply.
func MemoryGrow64(m *Module, n int64) int64 {
	m.MemMu.Lock()
	defer m.MemMu.Unlock()
	cur := m.MemSize.Load()
	prev := int64(cur >> 16)
	if n == 0 {
		return prev
	}
	if n < 0 {
		return -1
	}
	if uint64(n) > Mem64HardCap()>>16 {
		return -1
	}
	want := cur + uint64(n)*65536
	if want < cur || want > Mem64HardCap() {
		return -1
	}
	if m.MaxMem != 0 && want > m.MaxMem {
		return -1
	}
	if m.MemShared {

		if want > uint64(len(m.Memory)) {
			return -1
		}
		m.MemSize.Store(want)
		return prev
	}
	if want <= uint64(cap(m.Memory)) {
		m.Memory = m.Memory[:want]
		m.MemSize.Store(want)
		return prev
	}
	newCap := uint64(cap(m.Memory)) * 2
	if newCap < want {
		newCap = want
	}
	if m.MaxMem != 0 && newCap > m.MaxMem {
		newCap = m.MaxMem
	}
	if newCap > Mem64HardCap() {
		newCap = Mem64HardCap()
	}
	grown := make([]byte, want, newCap)
	copy(grown, m.Memory)
	m.Memory = grown
	m.MemSize.Store(want)

	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	return prev
}

//go:noinline
func MemoryFill64(m *Module, dst int64, val int32, n int64) {
	if n == 0 {
		return
	}
	end := uint64(dst) + uint64(n)
	if n < 0 || end < uint64(dst) || end > m.MemSize.Load() {
		Wasm_trap_memfill_oob()
	}
	b := m.Memory[uint64(dst):end]
	v := byte(val)
	if v == 0 {
		for k := range b {
			b[k] = 0
		}
		return
	}
	b[0] = v
	for filled := 1; filled < len(b); filled *= 2 {
		copy(b[filled:], b[:filled])
	}
}

//go:noinline
func MemoryCopy64(m *Module, dst int64, src int64, n int64) {
	if n == 0 {
		return
	}
	srcEnd := uint64(src) + uint64(n)
	dstEnd := uint64(dst) + uint64(n)
	size := m.MemSize.Load()
	if n < 0 || srcEnd < uint64(src) || dstEnd < uint64(dst) || srcEnd > size || dstEnd > size {
		Wasm_trap_memcopy_oob()
	}
	copy(m.Memory[uint64(dst):dstEnd], m.Memory[uint64(src):srcEnd])
}

//go:noinline
func MemoryInit64(m *Module, seg int, dst int64, src int32, n int32) {
	data := m.DataSegs[seg]
	if n == 0 {
		return
	}
	dstEnd := uint64(dst) + uint64(uint32(n))
	if data == nil || n < 0 ||
		uint64(uint32(src))+uint64(uint32(n)) > uint64(len(data)) ||
		dstEnd < uint64(dst) || dstEnd > m.MemSize.Load() {
		Wasm_trap_meminit_oob()
	}
	if dstEnd > uint64(m.DataEnd) {

		if dstEnd >= uint64(^uint32(0)) {
			m.DataEnd = ^uint32(0)
		} else {
			m.DataEnd = uint32(dstEnd)
		}
	}
	d := m.Memory[uint64(dst):dstEnd]
	s := data[uint32(src) : uint32(src)+uint32(n)]

	if bytes.Equal(d, s) {
		return
	}
	copy(d, s)
}

// simdEA64 is simdEA for 64-bit addresses: u64 effective address with
// an overflow-safe range check.
func SimdEA64(m *Module, addr int64, offset int64, size uint64) uint64 {
	ea := uint64(addr) + uint64(offset)
	end := ea + size
	if ea < uint64(addr) || end < ea || end > m.MemSize.Load() {
		Wasm_trap_simd_oob()
	}
	return ea
}

//go:noinline
func Simd_m64_v128_load(m *Module, addr int64, offset int64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 16)
	p := unsafe.Add(m.M, uintptr(ea))
	return [2]uint64{*(*uint64)(p), *(*uint64)(unsafe.Add(p, 8))}
}

// simd_m64_v128_load_rng / _nc are the bounds-coalesced memory64 loads
// (see internal/ssa/pass CoalesceSimdBounds): the group leader carries
// one full-width, overflow-safe range check covering the window, the
// rest load unchecked. rlo/span are the window; all four operands ride
// at pointer width.
//
//go:noinline
func Simd_m64_v128_load_rng(m *Module, addr int64, offset int64, rlo int64, span int64) [2]uint64 {

	start := addr + rlo
	if start < 0 || uint64(start)+uint64(span) > m.MemSize.Load() {
		Wasm_trap_simd_oob()
	}
	ea := uint64(addr) + uint64(offset)
	p := unsafe.Add(m.M, uintptr(ea))
	return [2]uint64{*(*uint64)(p), *(*uint64)(unsafe.Add(p, 8))}
}

//go:noinline
func Simd_m64_v128_load_nc(m *Module, addr int64, offset int64) [2]uint64 {
	ea := uint64(addr) + uint64(offset)
	p := unsafe.Add(m.M, uintptr(ea))
	return [2]uint64{*(*uint64)(p), *(*uint64)(unsafe.Add(p, 8))}
}

//go:noinline
func Simd_m64_v128_store(m *Module, addr int64, offset int64, v [2]uint64) int32 {
	ea := SimdEA64(m, addr, offset, 16)
	p := unsafe.Add(m.M, uintptr(ea))
	*(*uint64)(p) = v[0]
	*(*uint64)(unsafe.Add(p, 8)) = v[1]
	return 0
}

//go:noinline
func Simd_m64_v128_load32_zero(m *Module, addr int64, offset int64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 4)
	return [2]uint64{uint64(*(*uint32)(unsafe.Add(m.M, uintptr(ea)))), 0}
}

//go:noinline
func Simd_m64_v128_load64_zero(m *Module, addr int64, offset int64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 8)
	return [2]uint64{*(*uint64)(unsafe.Add(m.M, uintptr(ea))), 0}
}

//go:noinline
func Simd_m64_v128_load8_splat(m *Module, addr int64, offset int64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 1)
	b := uint64(*(*uint8)(unsafe.Add(m.M, uintptr(ea))))
	w := b * 0x0101010101010101
	return [2]uint64{w, w}
}

//go:noinline
func Simd_m64_v128_load16_splat(m *Module, addr int64, offset int64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 2)
	h := uint64(*(*uint16)(unsafe.Add(m.M, uintptr(ea))))
	w := h * 0x0001000100010001
	return [2]uint64{w, w}
}

//go:noinline
func Simd_m64_v128_load32_splat(m *Module, addr int64, offset int64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 4)
	x := uint64(*(*uint32)(unsafe.Add(m.M, uintptr(ea))))
	w := x | x<<32
	return [2]uint64{w, w}
}

// widenLoad64 reads 8 bytes and widens each source lane to the next
// width, signed or unsigned — the shared body of the load8x8/16x4/32x2
// family.
func Simd_m64_widen(m *Module, addr int64, offset int64, elemBits int, signed bool) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 8)
	x := *(*uint64)(unsafe.Add(m.M, uintptr(ea)))
	var v [2]uint64
	switch elemBits {
	case 8:
		for i := 0; i < 8; i++ {
			b := byte(x >> (8 * uint(i)))
			var w uint16
			if signed {
				w = uint16(int16(int8(b)))
			} else {
				w = uint16(b)
			}
			v[i/4] |= uint64(w) << (16 * uint(i%4))
		}
	case 16:
		for i := 0; i < 4; i++ {
			h := uint16(x >> (16 * uint(i)))
			var w uint32
			if signed {
				w = uint32(int32(int16(h)))
			} else {
				w = uint32(h)
			}
			v[i/2] |= uint64(w) << (32 * uint(i%2))
		}
	default:
		for i := 0; i < 2; i++ {
			q := uint32(x >> (32 * uint(i)))
			var w uint64
			if signed {
				w = uint64(int64(int32(q)))
			} else {
				w = uint64(q)
			}
			v[i] = w
		}
	}
	return v
}

//go:noinline
func Simd_m64_v128_load8x8_s(m *Module, addr int64, offset int64) [2]uint64 {
	return Simd_m64_widen(m, addr, offset, 8, true)
}

//go:noinline
func Simd_m64_v128_load16x4_u(m *Module, addr int64, offset int64) [2]uint64 {
	return Simd_m64_widen(m, addr, offset, 16, false)
}

//go:noinline
func Simd_m64_v128_load32x2_u(m *Module, addr int64, offset int64) [2]uint64 {
	return Simd_m64_widen(m, addr, offset, 32, false)
}

//go:noinline
func Simd_m64_v128_load8_lane(m *Module, addr int64, offset int64, lane int32, v [2]uint64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 1)
	x := *(*uint8)(unsafe.Add(m.M, uintptr(ea)))
	sh := 8 * uint(lane) % 64
	i := int(lane) * 8 / 64
	v[i] = v[i]&^(uint64(0xff)<<sh) | uint64(x)<<sh
	return v
}

//go:noinline
func Simd_m64_v128_load16_lane(m *Module, addr int64, offset int64, lane int32, v [2]uint64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 2)
	x := *(*uint16)(unsafe.Add(m.M, uintptr(ea)))
	sh := 16 * uint(lane) % 64
	i := int(lane) * 16 / 64
	v[i] = v[i]&^(uint64(0xffff)<<sh) | uint64(x)<<sh
	return v
}

//go:noinline
func Simd_m64_v128_load32_lane(m *Module, addr int64, offset int64, lane int32, v [2]uint64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 4)
	x := *(*uint32)(unsafe.Add(m.M, uintptr(ea)))
	sh := 32 * uint(lane) % 64
	i := int(lane) * 32 / 64
	v[i] = v[i]&^(uint64(^uint32(0))<<sh) | uint64(x)<<sh
	return v
}

//go:noinline
func Simd_m64_v128_load64_lane(m *Module, addr int64, offset int64, lane int32, v [2]uint64) [2]uint64 {
	ea := SimdEA64(m, addr, offset, 8)
	v[lane] = *(*uint64)(unsafe.Add(m.M, uintptr(ea)))
	return v
}

//go:noinline
func Simd_m64_v128_store8_lane(m *Module, addr int64, offset int64, lane int32, v [2]uint64) int32 {
	ea := SimdEA64(m, addr, offset, 1)
	sh := 8 * uint(lane) % 64
	*(*uint8)(unsafe.Add(m.M, uintptr(ea))) = uint8(v[int(lane)*8/64] >> sh)
	return 0
}

//go:noinline
func Simd_m64_v128_store16_lane(m *Module, addr int64, offset int64, lane int32, v [2]uint64) int32 {
	ea := SimdEA64(m, addr, offset, 2)
	sh := 16 * uint(lane) % 64
	*(*uint16)(unsafe.Add(m.M, uintptr(ea))) = uint16(v[int(lane)*16/64] >> sh)
	return 0
}

//go:noinline
func Simd_m64_v128_store32_lane(m *Module, addr int64, offset int64, lane int32, v [2]uint64) int32 {
	ea := SimdEA64(m, addr, offset, 4)
	sh := 32 * uint(lane) % 64
	*(*uint32)(unsafe.Add(m.M, uintptr(ea))) = uint32(v[int(lane)*32/64] >> sh)
	return 0
}

//go:noinline
func Simd_m64_v128_store64_lane(m *Module, addr int64, offset int64, lane int32, v [2]uint64) int32 {
	ea := SimdEA64(m, addr, offset, 8)
	*(*uint64)(unsafe.Add(m.M, uintptr(ea))) = v[lane]
	return 0
}

//go:noinline
func Simd_p_m64_v128_load(m *Module, addr int64, offset int64) (uint64, uint64) {
	r := Simd_m64_v128_load(m, addr, offset)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load_rng(m *Module, addr int64, offset int64, rlo int64, span int64) (uint64, uint64) {
	r := Simd_m64_v128_load_rng(m, addr, offset, rlo, span)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load_nc(m *Module, addr int64, offset int64) (uint64, uint64) {
	r := Simd_m64_v128_load_nc(m, addr, offset)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_store(m *Module, addr int64, offset int64, v0, v1 uint64) int32 {
	return Simd_m64_v128_store(m, addr, offset, [2]uint64{v0, v1})
}

//go:noinline
func Simd_p_m64_v128_load32_zero(m *Module, addr int64, offset int64) (uint64, uint64) {
	r := Simd_m64_v128_load32_zero(m, addr, offset)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load64_zero(m *Module, addr int64, offset int64) (uint64, uint64) {
	r := Simd_m64_v128_load64_zero(m, addr, offset)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load8_splat(m *Module, addr int64, offset int64) (uint64, uint64) {
	r := Simd_m64_v128_load8_splat(m, addr, offset)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load16_splat(m *Module, addr int64, offset int64) (uint64, uint64) {
	r := Simd_m64_v128_load16_splat(m, addr, offset)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load32_splat(m *Module, addr int64, offset int64) (uint64, uint64) {
	r := Simd_m64_v128_load32_splat(m, addr, offset)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load8x8_s(m *Module, addr int64, offset int64) (uint64, uint64) {
	r := Simd_m64_v128_load8x8_s(m, addr, offset)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load32x2_u(m *Module, addr int64, offset int64) (uint64, uint64) {
	r := Simd_m64_v128_load32x2_u(m, addr, offset)
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load8_lane(m *Module, addr int64, offset int64, lane int32, v0, v1 uint64) (uint64, uint64) {
	r := Simd_m64_v128_load8_lane(m, addr, offset, lane, [2]uint64{v0, v1})
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load16_lane(m *Module, addr int64, offset int64, lane int32, v0, v1 uint64) (uint64, uint64) {
	r := Simd_m64_v128_load16_lane(m, addr, offset, lane, [2]uint64{v0, v1})
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load32_lane(m *Module, addr int64, offset int64, lane int32, v0, v1 uint64) (uint64, uint64) {
	r := Simd_m64_v128_load32_lane(m, addr, offset, lane, [2]uint64{v0, v1})
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_load64_lane(m *Module, addr int64, offset int64, lane int32, v0, v1 uint64) (uint64, uint64) {
	r := Simd_m64_v128_load64_lane(m, addr, offset, lane, [2]uint64{v0, v1})
	return r[0], r[1]
}

//go:noinline
func Simd_p_m64_v128_store8_lane(m *Module, addr int64, offset int64, lane int32, v0, v1 uint64) int32 {
	return Simd_m64_v128_store8_lane(m, addr, offset, lane, [2]uint64{v0, v1})
}

//go:noinline
func Simd_p_m64_v128_store16_lane(m *Module, addr int64, offset int64, lane int32, v0, v1 uint64) int32 {
	return Simd_m64_v128_store16_lane(m, addr, offset, lane, [2]uint64{v0, v1})
}

//go:noinline
func Simd_p_m64_v128_store32_lane(m *Module, addr int64, offset int64, lane int32, v0, v1 uint64) int32 {
	return Simd_m64_v128_store32_lane(m, addr, offset, lane, [2]uint64{v0, v1})
}

//go:noinline
func Simd_p_m64_v128_store64_lane(m *Module, addr int64, offset int64, lane int32, v0, v1 uint64) int32 {
	return Simd_m64_v128_store64_lane(m, addr, offset, lane, [2]uint64{v0, v1})
}

//go:noinline
func Simd_p_fx0(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 440, n0)
	return
}

//go:noinline
func Simd_p_fx1(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 576, n0)
	return
}

//go:noinline
func Simd_p_fx2(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 320)
	_ = Simd_m64_v128_store(m, s0, 32, n0)
	n2 := Simd_m64_v128_load(m, s0, 304)
	_ = Simd_m64_v128_store(m, s0, 16, n2)
	n4 := Simd_m64_v128_load(m, s0, 288)
	_ = Simd_m64_v128_store(m, s0, 0, n4)
	return
}

//go:noinline
func Simd_p_fx3(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx4(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 264)
	_ = Simd_m64_v128_store(m, s0, 88, n0)
	n2 := Simd_m64_v128_load(m, s0, 248)
	_ = Simd_m64_v128_store(m, s0, 72, n2)
	n4 := Simd_m64_v128_load(m, s0, 232)
	_ = Simd_m64_v128_store(m, s0, 56, n4)
	return
}

//go:noinline
func Simd_p_fx5(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 160, n0)
	return
}

//go:noinline
func Simd_p_fx6(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 184, n0)
	return
}

//go:noinline
func Simd_p_fx7(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 208, n0)
	return
}

//go:noinline
func Simd_p_fx8(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 200, n0)
	return
}

//go:noinline
func Simd_p_fx9(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 8, n0)
	n2 := Simd_m64_v128_load(m, s2, 0)
	_ = Simd_m64_v128_store(m, s1, 24, n2)
	return
}

//go:noinline
func Simd_p_fx10(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	n2 := Simd_m64_v128_load(m, s2, 0)
	_ = Simd_m64_v128_store(m, s1, 16, n2)
	return
}

//go:noinline
func Simd_p_fx11(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 8, n0)
	return
}

//go:noinline
func Simd_p_fx12(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 200)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx13(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 80)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx14(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 56)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx15(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 56, n0)
	return
}

//go:noinline
func Simd_p_fx16(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 864, n0)
	return
}

//go:noinline
func Simd_p_fx17(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	n2 := Simd_m64_v128_load(m, s2, 0)
	_ = Simd_m64_v128_store(m, s1, 16, n2)
	n4 := Simd_m64_v128_load(m, s3, 0)
	_ = Simd_m64_v128_store(m, s1, 32, n4)
	return
}

//go:noinline
func Simd_p_fx18(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 1136, n0)
	return
}

//go:noinline
func Simd_p_fx19(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1168)
	_ = Simd_m64_v128_store(m, s0, 488, n0)
	n2 := Simd_m64_v128_load(m, s0, 1152)
	_ = Simd_m64_v128_store(m, s0, 472, n2)
	n4 := Simd_m64_v128_load(m, s0, 1136)
	_ = Simd_m64_v128_store(m, s0, 456, n4)
	return
}

//go:noinline
func Simd_p_fx20(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1168)
	_ = Simd_m64_v128_store(m, s0, 432, n0)
	n2 := Simd_m64_v128_load(m, s0, 1152)
	_ = Simd_m64_v128_store(m, s0, 416, n2)
	n4 := Simd_m64_v128_load(m, s0, 1136)
	_ = Simd_m64_v128_store(m, s0, 400, n4)
	return
}

//go:noinline
func Simd_p_fx21(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1080)
	_ = Simd_m64_v128_store(m, s0, 376, n0)
	n2 := Simd_m64_v128_load(m, s0, 1064)
	_ = Simd_m64_v128_store(m, s0, 360, n2)
	n4 := Simd_m64_v128_load(m, s0, 1048)
	_ = Simd_m64_v128_store(m, s0, 344, n4)
	return
}

//go:noinline
func Simd_p_fx22(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1168)
	_ = Simd_m64_v128_store(m, s0, 1080, n0)
	n2 := Simd_m64_v128_load(m, s0, 1152)
	_ = Simd_m64_v128_store(m, s0, 1064, n2)
	n4 := Simd_m64_v128_load(m, s0, 1136)
	_ = Simd_m64_v128_store(m, s0, 1048, n4)
	return
}

//go:noinline
func Simd_p_fx23(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1080)
	_ = Simd_m64_v128_store(m, s0, 40, n0)
	n2 := Simd_m64_v128_load(m, s0, 1064)
	_ = Simd_m64_v128_store(m, s0, 24, n2)
	n4 := Simd_m64_v128_load(m, s0, 1048)
	_ = Simd_m64_v128_store(m, s0, 8, n4)
	return
}

//go:noinline
func Simd_p_fx24(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 1192, n0)
	return
}

//go:noinline
func Simd_p_fx25(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 720)
	_ = Simd_m64_v128_store(m, s0, 320, n0)
	n2 := Simd_m64_v128_load(m, s0, 704)
	_ = Simd_m64_v128_store(m, s0, 304, n2)
	n4 := Simd_m64_v128_load(m, s0, 688)
	_ = Simd_m64_v128_store(m, s0, 288, n4)
	return
}

//go:noinline
func Simd_p_fx26(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1080)
	_ = Simd_m64_v128_store(m, s0, 264, n0)
	n2 := Simd_m64_v128_load(m, s0, 1064)
	_ = Simd_m64_v128_store(m, s0, 248, n2)
	n4 := Simd_m64_v128_load(m, s0, 1048)
	_ = Simd_m64_v128_store(m, s0, 232, n4)
	return
}

//go:noinline
func Simd_p_fx27(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 640)
	_ = Simd_m64_v128_store(m, s0, 208, n0)
	n2 := Simd_m64_v128_load(m, s0, 624)
	_ = Simd_m64_v128_store(m, s0, 192, n2)
	n4 := Simd_m64_v128_load(m, s0, 608)
	_ = Simd_m64_v128_store(m, s0, 176, n4)
	return
}

//go:noinline
func Simd_p_fx28(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1080)
	_ = Simd_m64_v128_store(m, s0, 152, n0)
	n2 := Simd_m64_v128_load(m, s0, 1064)
	_ = Simd_m64_v128_store(m, s0, 136, n2)
	n4 := Simd_m64_v128_load(m, s0, 1048)
	_ = Simd_m64_v128_store(m, s0, 120, n4)
	return
}

//go:noinline
func Simd_p_fx29(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1080)
	_ = Simd_m64_v128_store(m, s0, 96, n0)
	n2 := Simd_m64_v128_load(m, s0, 1064)
	_ = Simd_m64_v128_store(m, s0, 80, n2)
	n4 := Simd_m64_v128_load(m, s0, 1048)
	_ = Simd_m64_v128_store(m, s0, 64, n4)
	return
}

//go:noinline
func Simd_p_fx30(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 560, n0)
	return
}

//go:noinline
func Simd_p_fx31(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 584, n0)
	return
}

//go:noinline
func Simd_p_fx32(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 664, n0)
	return
}

//go:noinline
func Simd_p_fx33(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 400, n0)
	return
}

//go:noinline
func Simd_p_fx34(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 456)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx35(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 432)
	_ = Simd_m64_v128_store(m, s0, 208, n0)
	n2 := Simd_m64_v128_load(m, s0, 416)
	_ = Simd_m64_v128_store(m, s0, 192, n2)
	n4 := Simd_m64_v128_load(m, s0, 400)
	_ = Simd_m64_v128_store(m, s0, 176, n4)
	return
}

//go:noinline
func Simd_p_fx36(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 432)
	_ = Simd_m64_v128_store(m, s0, 152, n0)
	n2 := Simd_m64_v128_load(m, s0, 416)
	_ = Simd_m64_v128_store(m, s0, 136, n2)
	n4 := Simd_m64_v128_load(m, s0, 400)
	_ = Simd_m64_v128_store(m, s0, 120, n4)
	return
}

//go:noinline
func Simd_p_fx37(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 304, n0)
	return
}

//go:noinline
func Simd_p_fx38(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 328, n0)
	return
}

//go:noinline
func Simd_p_fx39(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 432)
	_ = Simd_m64_v128_store(m, s0, 96, n0)
	n2 := Simd_m64_v128_load(m, s0, 416)
	_ = Simd_m64_v128_store(m, s0, 80, n2)
	n4 := Simd_m64_v128_load(m, s0, 400)
	_ = Simd_m64_v128_store(m, s0, 64, n4)
	return
}

//go:noinline
func Simd_p_fx40(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 432)
	_ = Simd_m64_v128_store(m, s0, 40, n0)
	n2 := Simd_m64_v128_load(m, s0, 416)
	_ = Simd_m64_v128_store(m, s0, 24, n2)
	n4 := Simd_m64_v128_load(m, s0, 400)
	_ = Simd_m64_v128_store(m, s0, 8, n4)
	return
}

//go:noinline
func Simd_p_fx41(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 280, n0)
	return
}

//go:noinline
func Simd_p_fx42(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 800)
	_ = Simd_m64_v128_store(m, s0, 680, n0)
	return
}

//go:noinline
func Simd_p_fx43(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 312, 312, 32)
	n1 := Simd_m64_v128_load_nc(m, s0, 328)
	_ = Simd_m64_v128_store(m, s0, 312, n1)
	_ = Simd_m64_v128_store(m, s0, 328, n0)
	return
}

//go:noinline
func Simd_p_fx44(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 48)
	_ = Simd_m64_v128_store(m, s1, 48, n0)
	n2 := Simd_m64_v128_load(m, s0, 64)
	_ = Simd_m64_v128_store(m, s1, 64, n2)
	return
}

//go:noinline
func Simd_p_fx45(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{1084818905618843912, 506097522914230528})
	n1 := Simd_i64x2_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx46(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 504, n0)
	return
}

//go:noinline
func Simd_p_fx47(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 4856, n0)
	return
}

//go:noinline
func Simd_p_fx48(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 360, n0)
	return
}

//go:noinline
func Simd_p_fx49(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 384, n0)
	return
}

//go:noinline
func Simd_p_fx50(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 408, n0)
	return
}

//go:noinline
func Simd_p_fx51(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 432, n0)
	return
}

//go:noinline
func Simd_p_fx52(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 456, n0)
	return
}

//go:noinline
func Simd_p_fx53(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 480, n0)
	return
}

//go:noinline
func Simd_p_fx54(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 184)
	_ = Simd_m64_v128_store(m, s0, 120, n0)
	return
}

//go:noinline
func Simd_p_fx55(m *Module, s0 int64, s1 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_m64_v128_load32_zero(m, s0, 0)
	n2 := Simd_i16x8_extmul_low_i8x16_u(n1, [2]uint64{p0, p0h})
	n3 := Simd_i8x16_shuffle(n2, n0, [2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_low_i8x16_u(n3)
	n5 := Simd_i32x4_extend_low_i16x8_u(n4)
	n6 := Simd_i32x4_mul(n5, [2]uint64{p2, p2h})
	n7 := Simd_i32x4_shr_u(n6, 8)
	n8 := Simd_i32x4_add(n7, [2]uint64{p3, p3h})
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul(n0, n9)
	_ = Simd_m64_v128_store(m, s1+128, 0, n10)
	n12 := Simd_m64_v128_load32_zero(m, s0, 4)
	n13 := Simd_i16x8_extmul_low_i8x16_u(n12, [2]uint64{p0, p0h})
	n14 := Simd_i8x16_shuffle(n13, n0, [2]uint64{p1, p1h})
	n15 := Simd_i16x8_extend_low_i8x16_u(n14)
	n16 := Simd_i32x4_extend_low_i16x8_u(n15)
	n17 := Simd_i32x4_mul(n16, [2]uint64{p2, p2h})
	n18 := Simd_i32x4_shr_u(n17, 8)
	n19 := Simd_i32x4_add(n18, [2]uint64{p3, p3h})
	n20 := Simd_f32x4_convert_i32x4_s(n19)
	n21 := Simd_f32x4_mul(n0, n20)
	_ = Simd_m64_v128_store(m, s1+144, 0, n21)
	n23 := Simd_m64_v128_load32_zero(m, s0, 8)
	n24 := Simd_i16x8_extmul_low_i8x16_u(n23, [2]uint64{p0, p0h})
	n25 := Simd_i8x16_shuffle(n24, n0, [2]uint64{p1, p1h})
	n26 := Simd_i16x8_extend_low_i8x16_u(n25)
	n27 := Simd_i32x4_extend_low_i16x8_u(n26)
	n28 := Simd_i32x4_mul(n27, [2]uint64{p2, p2h})
	n29 := Simd_i32x4_shr_u(n28, 8)
	n30 := Simd_i32x4_add(n29, [2]uint64{p3, p3h})
	n31 := Simd_f32x4_convert_i32x4_s(n30)
	n32 := Simd_f32x4_mul(n0, n31)
	_ = Simd_m64_v128_store(m, s1+160, 0, n32)
	n34 := Simd_m64_v128_load32_zero(m, s0, 12)
	n35 := Simd_i16x8_extmul_low_i8x16_u(n34, [2]uint64{p0, p0h})
	n36 := Simd_i8x16_shuffle(n35, n0, [2]uint64{p1, p1h})
	n37 := Simd_i16x8_extend_low_i8x16_u(n36)
	n38 := Simd_i32x4_extend_low_i16x8_u(n37)
	n39 := Simd_i32x4_mul(n38, [2]uint64{p2, p2h})
	n40 := Simd_i32x4_shr_u(n39, 8)
	n41 := Simd_i32x4_add(n40, [2]uint64{p3, p3h})
	n42 := Simd_f32x4_convert_i32x4_s(n41)
	n43 := Simd_f32x4_mul(n0, n42)
	_ = Simd_m64_v128_store(m, s1+176, 0, n43)
	n45 := Simd_m64_v128_load32_zero(m, s0, 16)
	n46 := Simd_i16x8_extmul_low_i8x16_u(n45, [2]uint64{p0, p0h})
	n47 := Simd_i8x16_shuffle(n46, n0, [2]uint64{p1, p1h})
	n48 := Simd_i16x8_extend_low_i8x16_u(n47)
	n49 := Simd_i32x4_extend_low_i16x8_u(n48)
	n50 := Simd_i32x4_mul(n49, [2]uint64{p2, p2h})
	n51 := Simd_i32x4_shr_u(n50, 8)
	n52 := Simd_i32x4_add(n51, [2]uint64{p3, p3h})
	n53 := Simd_f32x4_convert_i32x4_s(n52)
	n54 := Simd_f32x4_mul(n0, n53)
	_ = Simd_m64_v128_store(m, s1+192, 0, n54)
	n56 := Simd_m64_v128_load32_zero(m, s0, 20)
	n57 := Simd_i16x8_extmul_low_i8x16_u(n56, [2]uint64{p0, p0h})
	n58 := Simd_i8x16_shuffle(n57, n0, [2]uint64{p1, p1h})
	n59 := Simd_i16x8_extend_low_i8x16_u(n58)
	n60 := Simd_i32x4_extend_low_i16x8_u(n59)
	n61 := Simd_i32x4_mul(n60, [2]uint64{p2, p2h})
	n62 := Simd_i32x4_shr_u(n61, 8)
	n63 := Simd_i32x4_add(n62, [2]uint64{p3, p3h})
	n64 := Simd_f32x4_convert_i32x4_s(n63)
	n65 := Simd_f32x4_mul(n0, n64)
	_ = Simd_m64_v128_store(m, s1+208, 0, n65)
	n67 := Simd_m64_v128_load32_zero(m, s0, 24)
	n68 := Simd_i16x8_extmul_low_i8x16_u(n67, [2]uint64{p0, p0h})
	n69 := Simd_i8x16_shuffle(n68, n0, [2]uint64{p1, p1h})
	n70 := Simd_i16x8_extend_low_i8x16_u(n69)
	n71 := Simd_i32x4_extend_low_i16x8_u(n70)
	n72 := Simd_i32x4_mul(n71, [2]uint64{p2, p2h})
	n73 := Simd_i32x4_shr_u(n72, 8)
	n74 := Simd_i32x4_add(n73, [2]uint64{p3, p3h})
	n75 := Simd_f32x4_convert_i32x4_s(n74)
	n76 := Simd_f32x4_mul(n0, n75)
	_ = Simd_m64_v128_store(m, s1+224, 0, n76)
	n78 := Simd_m64_v128_load32_zero(m, s0, 28)
	n79 := Simd_i16x8_extmul_low_i8x16_u(n78, [2]uint64{p0, p0h})
	n80 := Simd_i8x16_shuffle(n79, n0, [2]uint64{p1, p1h})
	n81 := Simd_i16x8_extend_low_i8x16_u(n80)
	n82 := Simd_i32x4_extend_low_i16x8_u(n81)
	n83 := Simd_i32x4_mul(n82, [2]uint64{p2, p2h})
	n84 := Simd_i32x4_shr_u(n83, 8)
	n85 := Simd_i32x4_add(n84, [2]uint64{p3, p3h})
	n86 := Simd_f32x4_convert_i32x4_s(n85)
	n87 := Simd_f32x4_mul(n0, n86)
	_ = Simd_m64_v128_store(m, s1+240, 0, n87)
	return
}

//go:noinline
func Simd_p_fx56(m *Module, s0 int64, s1 int64, s2 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_m64_v128_load32_zero(m, s0, 0)
	n2 := Simd_i16x8_extmul_low_i8x16_u(n1, [2]uint64{p0, p0h})
	n3 := Simd_i8x16_shuffle(n2, n0, [2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_low_i8x16_u(n3)
	n5 := Simd_i32x4_extend_low_i16x8_u(n4)
	n6 := Simd_i32x4_mul(n5, [2]uint64{p2, p2h})
	n7 := Simd_i32x4_shr_u(n6, 8)
	n8 := Simd_i32x4_add(n7, [2]uint64{p3, p3h})
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul(n0, n9)
	_ = Simd_m64_v128_store(m, s1, 0, n10)
	n12 := Simd_m64_v128_load32_zero(m, s0, 4)
	n13 := Simd_i16x8_extmul_low_i8x16_u(n12, [2]uint64{p0, p0h})
	n14 := Simd_i8x16_shuffle(n13, n0, [2]uint64{p1, p1h})
	n15 := Simd_i16x8_extend_low_i8x16_u(n14)
	n16 := Simd_i32x4_extend_low_i16x8_u(n15)
	n17 := Simd_i32x4_mul(n16, [2]uint64{p2, p2h})
	n18 := Simd_i32x4_shr_u(n17, 8)
	n19 := Simd_i32x4_add(n18, [2]uint64{p3, p3h})
	n20 := Simd_f32x4_convert_i32x4_s(n19)
	n21 := Simd_f32x4_mul(n0, n20)
	_ = Simd_m64_v128_store(m, s1+16, 0, n21)
	n23 := Simd_m64_v128_load32_zero(m, s0, 8)
	n24 := Simd_i16x8_extmul_low_i8x16_u(n23, [2]uint64{p0, p0h})
	n25 := Simd_i8x16_shuffle(n24, n0, [2]uint64{p1, p1h})
	n26 := Simd_i16x8_extend_low_i8x16_u(n25)
	n27 := Simd_i32x4_extend_low_i16x8_u(n26)
	n28 := Simd_i32x4_mul(n27, [2]uint64{p2, p2h})
	n29 := Simd_i32x4_shr_u(n28, 8)
	n30 := Simd_i32x4_add(n29, [2]uint64{p3, p3h})
	n31 := Simd_f32x4_convert_i32x4_s(n30)
	n32 := Simd_f32x4_mul(n0, n31)
	_ = Simd_m64_v128_store(m, s1+32, 0, n32)
	n34 := Simd_m64_v128_load32_zero(m, s0, 12)
	n35 := Simd_i16x8_extmul_low_i8x16_u(n34, [2]uint64{p0, p0h})
	n36 := Simd_i8x16_shuffle(n35, n0, [2]uint64{p1, p1h})
	n37 := Simd_i16x8_extend_low_i8x16_u(n36)
	n38 := Simd_i32x4_extend_low_i16x8_u(n37)
	n39 := Simd_i32x4_mul(n38, [2]uint64{p2, p2h})
	n40 := Simd_i32x4_shr_u(n39, 8)
	n41 := Simd_i32x4_add(n40, [2]uint64{p3, p3h})
	n42 := Simd_f32x4_convert_i32x4_s(n41)
	n43 := Simd_f32x4_mul(n0, n42)
	_ = Simd_m64_v128_store(m, s1+48, 0, n43)
	n45 := Simd_m64_v128_load32_zero(m, s0, 16)
	n46 := Simd_i16x8_extmul_low_i8x16_u(n45, [2]uint64{p0, p0h})
	n47 := Simd_i8x16_shuffle(n46, n0, [2]uint64{p1, p1h})
	n48 := Simd_i16x8_extend_low_i8x16_u(n47)
	n49 := Simd_i32x4_extend_low_i16x8_u(n48)
	n50 := Simd_i32x4_mul(n49, [2]uint64{p2, p2h})
	n51 := Simd_i32x4_shr_u(n50, 8)
	n52 := Simd_i32x4_add(n51, [2]uint64{p3, p3h})
	n53 := Simd_f32x4_convert_i32x4_s(n52)
	n54 := Simd_f32x4_mul(n0, n53)
	_ = Simd_m64_v128_store(m, s2, 0, n54)
	n56 := Simd_m64_v128_load32_zero(m, s0, 20)
	n57 := Simd_i16x8_extmul_low_i8x16_u(n56, [2]uint64{p0, p0h})
	n58 := Simd_i8x16_shuffle(n57, n0, [2]uint64{p1, p1h})
	n59 := Simd_i16x8_extend_low_i8x16_u(n58)
	n60 := Simd_i32x4_extend_low_i16x8_u(n59)
	n61 := Simd_i32x4_mul(n60, [2]uint64{p2, p2h})
	n62 := Simd_i32x4_shr_u(n61, 8)
	n63 := Simd_i32x4_add(n62, [2]uint64{p3, p3h})
	n64 := Simd_f32x4_convert_i32x4_s(n63)
	n65 := Simd_f32x4_mul(n0, n64)
	_ = Simd_m64_v128_store(m, s1+80, 0, n65)
	n67 := Simd_m64_v128_load32_zero(m, s0, 24)
	n68 := Simd_i16x8_extmul_low_i8x16_u(n67, [2]uint64{p0, p0h})
	n69 := Simd_i8x16_shuffle(n68, n0, [2]uint64{p1, p1h})
	n70 := Simd_i16x8_extend_low_i8x16_u(n69)
	n71 := Simd_i32x4_extend_low_i16x8_u(n70)
	n72 := Simd_i32x4_mul(n71, [2]uint64{p2, p2h})
	n73 := Simd_i32x4_shr_u(n72, 8)
	n74 := Simd_i32x4_add(n73, [2]uint64{p3, p3h})
	n75 := Simd_f32x4_convert_i32x4_s(n74)
	n76 := Simd_f32x4_mul(n0, n75)
	_ = Simd_m64_v128_store(m, s1+96, 0, n76)
	n78 := Simd_m64_v128_load32_zero(m, s0, 28)
	n79 := Simd_i16x8_extmul_low_i8x16_u(n78, [2]uint64{p0, p0h})
	n80 := Simd_i8x16_shuffle(n79, n0, [2]uint64{p1, p1h})
	n81 := Simd_i16x8_extend_low_i8x16_u(n80)
	n82 := Simd_i32x4_extend_low_i16x8_u(n81)
	n83 := Simd_i32x4_mul(n82, [2]uint64{p2, p2h})
	n84 := Simd_i32x4_shr_u(n83, 8)
	n85 := Simd_i32x4_add(n84, [2]uint64{p3, p3h})
	n86 := Simd_f32x4_convert_i32x4_s(n85)
	n87 := Simd_f32x4_mul(n0, n86)
	_ = Simd_m64_v128_store(m, s1+112, 0, n87)
	return
}

//go:noinline
func Simd_p_fx57(m *Module, s0 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n2 := Simd_i16x8_mul(n1, [2]uint64{p1, p1h})
	n3 := Simd_i16x8_shr_u(n2, 8)
	n4 := Simd_i16x8_add(n3, [2]uint64{p2, p2h})
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_f32x4_convert_i32x4_s(n5)
	n7 := Simd_f32x4_mul(n0, n6)
	n8 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p3, p3h})
	n9 := Simd_i16x8_mul(n8, [2]uint64{p1, p1h})
	n10 := Simd_i16x8_shr_u(n9, 8)
	n11 := Simd_i16x8_add(n10, [2]uint64{p2, p2h})
	n12 := Simd_i32x4_extend_low_i16x8_s(n11)
	n13 := Simd_f32x4_convert_i32x4_s(n12)
	n14 := Simd_f32x4_mul(n0, n13)
	n15 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n16 := Simd_i16x8_mul(n15, [2]uint64{p1, p1h})
	n17 := Simd_i16x8_shr_u(n16, 8)
	n18 := Simd_i16x8_add(n17, [2]uint64{p2, p2h})
	n19 := Simd_i32x4_extend_low_i16x8_s(n18)
	n20 := Simd_f32x4_convert_i32x4_s(n19)
	n21 := Simd_f32x4_mul(n0, n20)
	n22 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n23 := Simd_i16x8_mul(n22, [2]uint64{p1, p1h})
	n24 := Simd_i16x8_shr_u(n23, 8)
	n25 := Simd_i16x8_add(n24, [2]uint64{p2, p2h})
	n26 := Simd_i32x4_extend_low_i16x8_s(n25)
	n27 := Simd_f32x4_convert_i32x4_s(n26)
	n28 := Simd_f32x4_mul(n0, n27)
	n29 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p6, p6h})
	n30 := Simd_i16x8_mul(n29, [2]uint64{p1, p1h})
	n31 := Simd_i16x8_shr_u(n30, 8)
	n32 := Simd_i16x8_add(n31, [2]uint64{p2, p2h})
	n33 := Simd_i32x4_extend_low_i16x8_s(n32)
	n34 := Simd_f32x4_convert_i32x4_s(n33)
	n35 := Simd_f32x4_mul(n0, n34)
	_ = Simd_m64_v128_store(m, s0, 48, n7)
	_ = Simd_m64_v128_store(m, s0, 32, n14)
	_ = Simd_m64_v128_store(m, s0, 16, n21)
	_ = Simd_m64_v128_store(m, s0, 0, n28)
	_ = Simd_m64_v128_store(m, s0, 304, n35)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx58(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n1 := Simd_i16x8_mul(n0, [2]uint64{p2, p2h})
	n2 := Simd_i16x8_shr_u(n1, 8)
	n3 := Simd_i16x8_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_f32x4_convert_i32x4_s(n4)
	n6 := Simd_f32x4_mul([2]uint64{p0, p0h}, n5)
	n7 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n8 := Simd_i16x8_mul(n7, [2]uint64{p2, p2h})
	n9 := Simd_i16x8_shr_u(n8, 8)
	n10 := Simd_i16x8_add(n9, [2]uint64{p3, p3h})
	n11 := Simd_i32x4_extend_low_i16x8_s(n10)
	n12 := Simd_f32x4_convert_i32x4_s(n11)
	n13 := Simd_f32x4_mul([2]uint64{p0, p0h}, n12)
	n14 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n15 := Simd_i16x8_mul(n14, [2]uint64{p2, p2h})
	n16 := Simd_i16x8_shr_u(n15, 8)
	n17 := Simd_i16x8_add(n16, [2]uint64{p3, p3h})
	n18 := Simd_i32x4_extend_low_i16x8_s(n17)
	n19 := Simd_f32x4_convert_i32x4_s(n18)
	n20 := Simd_f32x4_mul([2]uint64{p0, p0h}, n19)
	n21 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p6, p6h})
	n22 := Simd_i16x8_mul(n21, [2]uint64{p2, p2h})
	n23 := Simd_i16x8_shr_u(n22, 8)
	n24 := Simd_i16x8_add(n23, [2]uint64{p3, p3h})
	n25 := Simd_i32x4_extend_low_i16x8_s(n24)
	n26 := Simd_f32x4_convert_i32x4_s(n25)
	n27 := Simd_f32x4_mul([2]uint64{p0, p0h}, n26)
	_ = Simd_m64_v128_store(m, s0, 288, n6)
	_ = Simd_m64_v128_store(m, s0, 272, n13)
	_ = Simd_m64_v128_store(m, s0, 256, n20)
	_ = Simd_m64_v128_store(m, s0, 240, n27)
	return
}

//go:noinline
func Simd_p_fx59(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n1 := Simd_i16x8_mul(n0, [2]uint64{p2, p2h})
	n2 := Simd_i16x8_shr_u(n1, 8)
	n3 := Simd_i16x8_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_f32x4_convert_i32x4_s(n4)
	n6 := Simd_f32x4_mul([2]uint64{p0, p0h}, n5)
	n7 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n8 := Simd_i16x8_mul(n7, [2]uint64{p2, p2h})
	n9 := Simd_i16x8_shr_u(n8, 8)
	n10 := Simd_i16x8_add(n9, [2]uint64{p3, p3h})
	n11 := Simd_i32x4_extend_low_i16x8_s(n10)
	n12 := Simd_f32x4_convert_i32x4_s(n11)
	n13 := Simd_f32x4_mul([2]uint64{p0, p0h}, n12)
	n14 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n15 := Simd_i16x8_mul(n14, [2]uint64{p2, p2h})
	n16 := Simd_i16x8_shr_u(n15, 8)
	n17 := Simd_i16x8_add(n16, [2]uint64{p3, p3h})
	n18 := Simd_i32x4_extend_low_i16x8_s(n17)
	n19 := Simd_f32x4_convert_i32x4_s(n18)
	n20 := Simd_f32x4_mul([2]uint64{p0, p0h}, n19)
	n21 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p6, p6h})
	n22 := Simd_i16x8_mul(n21, [2]uint64{p2, p2h})
	n23 := Simd_i16x8_shr_u(n22, 8)
	n24 := Simd_i16x8_add(n23, [2]uint64{p3, p3h})
	n25 := Simd_i32x4_extend_low_i16x8_s(n24)
	n26 := Simd_f32x4_convert_i32x4_s(n25)
	n27 := Simd_f32x4_mul([2]uint64{p0, p0h}, n26)
	_ = Simd_m64_v128_store(m, s0, 224, n6)
	_ = Simd_m64_v128_store(m, s0, 208, n13)
	_ = Simd_m64_v128_store(m, s0, 192, n20)
	_ = Simd_m64_v128_store(m, s0, 176, n27)
	return
}

//go:noinline
func Simd_p_fx60(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n1 := Simd_i16x8_mul(n0, [2]uint64{p2, p2h})
	n2 := Simd_i16x8_shr_u(n1, 8)
	n3 := Simd_i16x8_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_f32x4_convert_i32x4_s(n4)
	n6 := Simd_f32x4_mul([2]uint64{p0, p0h}, n5)
	n7 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n8 := Simd_i16x8_mul(n7, [2]uint64{p2, p2h})
	n9 := Simd_i16x8_shr_u(n8, 8)
	n10 := Simd_i16x8_add(n9, [2]uint64{p3, p3h})
	n11 := Simd_i32x4_extend_low_i16x8_s(n10)
	n12 := Simd_f32x4_convert_i32x4_s(n11)
	n13 := Simd_f32x4_mul([2]uint64{p0, p0h}, n12)
	n14 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n15 := Simd_i16x8_mul(n14, [2]uint64{p2, p2h})
	n16 := Simd_i16x8_shr_u(n15, 8)
	n17 := Simd_i16x8_add(n16, [2]uint64{p3, p3h})
	n18 := Simd_i32x4_extend_low_i16x8_s(n17)
	n19 := Simd_f32x4_convert_i32x4_s(n18)
	n20 := Simd_f32x4_mul([2]uint64{p0, p0h}, n19)
	n21 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p6, p6h})
	n22 := Simd_i16x8_mul(n21, [2]uint64{p2, p2h})
	n23 := Simd_i16x8_shr_u(n22, 8)
	n24 := Simd_i16x8_add(n23, [2]uint64{p3, p3h})
	n25 := Simd_i32x4_extend_low_i16x8_s(n24)
	n26 := Simd_f32x4_convert_i32x4_s(n25)
	n27 := Simd_f32x4_mul([2]uint64{p0, p0h}, n26)
	_ = Simd_m64_v128_store(m, s0, 160, n6)
	_ = Simd_m64_v128_store(m, s0, 144, n13)
	_ = Simd_m64_v128_store(m, s0, 128, n20)
	_ = Simd_m64_v128_store(m, s0, 112, n27)
	return
}

//go:noinline
func Simd_p_fx61(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n1 := Simd_i16x8_mul(n0, [2]uint64{p2, p2h})
	n2 := Simd_i16x8_shr_u(n1, 8)
	n3 := Simd_i16x8_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_f32x4_convert_i32x4_s(n4)
	n6 := Simd_f32x4_mul([2]uint64{p0, p0h}, n5)
	n7 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n8 := Simd_i16x8_mul(n7, [2]uint64{p2, p2h})
	n9 := Simd_i16x8_shr_u(n8, 8)
	n10 := Simd_i16x8_add(n9, [2]uint64{p3, p3h})
	n11 := Simd_i32x4_extend_low_i16x8_s(n10)
	n12 := Simd_f32x4_convert_i32x4_s(n11)
	n13 := Simd_f32x4_mul([2]uint64{p0, p0h}, n12)
	n14 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n15 := Simd_i16x8_mul(n14, [2]uint64{p2, p2h})
	n16 := Simd_i16x8_shr_u(n15, 8)
	n17 := Simd_i16x8_add(n16, [2]uint64{p3, p3h})
	n18 := Simd_i32x4_extend_low_i16x8_s(n17)
	n19 := Simd_f32x4_convert_i32x4_s(n18)
	n20 := Simd_f32x4_mul([2]uint64{p0, p0h}, n19)
	_ = Simd_m64_v128_store(m, s0, 96, n6)
	_ = Simd_m64_v128_store(m, s0, 80, n13)
	_ = Simd_m64_v128_store(m, s0, 64, n20)
	return
}

//go:noinline
func Simd_p_fx62(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_i16x8_mul([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_shr_u(n0, 8)
	n2 := Simd_i16x8_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	n6 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n7 := Simd_i16x8_mul(n6, [2]uint64{p2, p2h})
	n8 := Simd_i16x8_shr_u(n7, 8)
	n9 := Simd_i16x8_add(n8, [2]uint64{p3, p3h})
	n10 := Simd_i32x4_extend_low_i16x8_s(n9)
	n11 := Simd_f32x4_convert_i32x4_s(n10)
	n12 := Simd_f32x4_mul([2]uint64{p0, p0h}, n11)
	_ = Simd_m64_v128_store(m, s0, 332, n5)
	_ = Simd_m64_v128_store(m, s0, 348, n12)
	return
}

//go:noinline
func Simd_p_fx63(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 672, n1)
	n3 := Simd_m64_v128_load(m, s0, 16)
	n4 := Simd_f32x4_mul(n3, n3)
	_ = Simd_m64_v128_store(m, s1, 688, n4)
	n6 := Simd_m64_v128_load(m, s0, 32)
	n7 := Simd_f32x4_mul(n6, n6)
	_ = Simd_m64_v128_store(m, s1, 704, n7)
	n9 := Simd_m64_v128_load(m, s0, 48)
	n10 := Simd_f32x4_mul(n9, n9)
	_ = Simd_m64_v128_store(m, s1, 720, n10)
	return n1[0], n1[1], n4[0], n4[1], n7[0], n7[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx64(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 64)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 736, n1)
	n3 := Simd_m64_v128_load(m, s0, 96)
	n4 := Simd_f32x4_mul(n3, n3)
	_ = Simd_m64_v128_store(m, s1, 768, n4)
	n6 := Simd_m64_v128_load(m, s0, 112)
	n7 := Simd_f32x4_mul(n6, n6)
	_ = Simd_m64_v128_store(m, s1, 784, n7)
	n9 := Simd_m64_v128_load(m, s0, 80)
	n10 := Simd_f32x4_mul(n9, n9)
	_ = Simd_m64_v128_store(m, s1, 752, n10)
	return n1[0], n1[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx65(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_f32x4_sqrt([2]uint64{p0, p0h})
	n1 := Simd_f32x4_sqrt([2]uint64{p1, p1h})
	n2 := Simd_f32x4_sqrt([2]uint64{p2, p2h})
	n3 := Simd_f32x4_sqrt([2]uint64{p3, p3h})
	n4 := Simd_f32x4_sqrt([2]uint64{p4, p4h})
	_ = Simd_m64_v128_store(m, s0, 416, n0)
	_ = Simd_m64_v128_store(m, s0, 400, n1)
	_ = Simd_m64_v128_store(m, s0, 384, n2)
	_ = Simd_m64_v128_store(m, s0, 368, n3)
	_ = Simd_m64_v128_store(m, s0, 352, n4)
	n10 := Simd_m64_v128_load(m, s0, 756)
	n11 := Simd_f32x4_sqrt(n10)
	_ = Simd_m64_v128_store(m, s0, 436, n11)
	n13 := Simd_m64_v128_load(m, s0, 772)
	n14 := Simd_f32x4_sqrt(n13)
	_ = Simd_m64_v128_store(m, s0, 452, n14)
	return
}

//go:noinline
func Simd_p_fx66(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+128, 0)
	n1 := Simd_f32x4_ge(n0, [2]uint64{p0, p0h})
	n2 := Simd_v128_not(n1)
	n3 := Simd_f32x4_neg(n0)
	n4 := Simd_v128_bitselect(n0, n3, n1)
	_ = Simd_m64_v128_store(m, s1+128, 0, n4)
	n6 := Simd_m64_v128_load(m, s0+144, 0)
	n7 := Simd_f32x4_ge(n6, [2]uint64{p0, p0h})
	n8 := Simd_f32x4_neg(n6)
	n9 := Simd_v128_bitselect(n6, n8, n7)
	_ = Simd_m64_v128_store(m, s1+144, 0, n9)
	return n0[0], n0[1], n7[0], n7[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx67(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_lane(m, s0+132, 0, 1, [2]uint64{p0, p0h})
	n1 := Simd_m64_scalar_i32_add(s1, 672)
	n2 := Simd_m64_scalar_i32_add(n1, s2)
	n3 := Simd_m64_scalar_i32_add(n2, 128)
	n4 := Simd_m64_v128_load(m, n3, 0)
	n5 := Simd_f32x4_mul(n0, n4)
	n6 := Simd_f32x4_mul(n0, n5)
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx68(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 496)
	_ = Simd_m64_v128_store(m, s0, 528, n0)
	n2 := Simd_m64_v128_load(m, s0, 480)
	_ = Simd_m64_v128_store(m, s0, 512, n2)
	return
}

//go:noinline
func Simd_p_fx69(m *Module, s0 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n1, n0)
	n3 := Simd_f32x4_add(n2, [2]uint64{p0, p0h})
	n4 := Simd_f32x4_mul(n3, [2]uint64{p1, p1h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p2, p2h})
	n6 := Simd_v128_and(n5, [2]uint64{p3, p3h})
	n7 := Simd_i32x4_max_u(n6, [2]uint64{p4, p4h})
	n8 := Simd_i32x4_sub(n7, [2]uint64{p4, p4h})
	n9 := Simd_i32x4_min_u(n8, [2]uint64{p5, p5h})
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx70(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{1084818905618843912, 216736831629295872})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx71(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{216736831696667908, 216736831629295872})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx72(m *Module, s0 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_m64_v128_load_rng(m, s0, 800, 288, 528)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add(n2, [2]uint64{p0, p0h})
	n4 := Simd_f32x4_mul(n3, [2]uint64{p1, p1h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p2, p2h})
	n6 := Simd_v128_and(n5, [2]uint64{p3, p3h})
	n7 := Simd_i32x4_max_u(n6, [2]uint64{p4, p4h})
	n8 := Simd_i32x4_sub(n7, [2]uint64{p4, p4h})
	n9 := Simd_i32x4_min_u(n8, [2]uint64{p5, p5h})
	n10 := Simd_i32x4_shl(n9, 28)
	n11 := Simd_m64_v128_load_nc(m, s0, 288)
	n12 := Simd_v128_or(n10, n11)
	_ = Simd_m64_v128_store(m, s0, 288, n12)
	n14 := Simd_m64_v128_load_rng(m, s0, 816, 304, 528)
	n15 := Simd_f32x4_mul(n0, n14)
	n16 := Simd_f32x4_add(n15, [2]uint64{p0, p0h})
	n17 := Simd_f32x4_mul(n16, [2]uint64{p1, p1h})
	n18 := Simd_f32x4_add(n17, [2]uint64{p2, p2h})
	n19 := Simd_v128_and(n18, [2]uint64{p3, p3h})
	n20 := Simd_i32x4_max_u(n19, [2]uint64{p4, p4h})
	n21 := Simd_i32x4_sub(n20, [2]uint64{p4, p4h})
	n22 := Simd_i32x4_min_u(n21, [2]uint64{p5, p5h})
	n23 := Simd_i32x4_shl(n22, 28)
	n24 := Simd_m64_v128_load_nc(m, s0, 304)
	n25 := Simd_v128_or(n23, n24)
	_ = Simd_m64_v128_store(m, s0, 304, n25)
	return
}

//go:noinline
func Simd_p_fx73(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 544, n1)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx74(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 28)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 572, n1)
	n3 := Simd_m64_v128_load(m, s0, 44)
	n4 := Simd_f32x4_mul(n3, n3)
	_ = Simd_m64_v128_store(m, s1, 588, n4)
	n6 := Simd_m64_v128_load(m, s0, 60)
	n7 := Simd_f32x4_mul(n6, n6)
	_ = Simd_m64_v128_store(m, s1, 604, n7)
	n9 := Simd_m64_v128_load(m, s0, 76)
	n10 := Simd_f32x4_mul(n9, n9)
	_ = Simd_m64_v128_store(m, s1, 620, n10)
	return n1[0], n1[1], n4[0], n4[1], n7[0], n7[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx75(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 92)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 636, n1)
	n3 := Simd_m64_v128_load(m, s0, 108)
	n4 := Simd_f32x4_mul(n3, n3)
	_ = Simd_m64_v128_store(m, s1, 652, n4)
	return n1[0], n1[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx76(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_f32x4_sqrt([2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 224, n0)
	return
}

//go:noinline
func Simd_p_fx77(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_f32x4_sqrt([2]uint64{p0, p0h})
	n1 := Simd_f32x4_sqrt([2]uint64{p1, p1h})
	n2 := Simd_f32x4_sqrt([2]uint64{p2, p2h})
	n3 := Simd_f32x4_sqrt([2]uint64{p3, p3h})
	n4 := Simd_f32x4_sqrt([2]uint64{p4, p4h})
	n5 := Simd_f32x4_sqrt([2]uint64{p5, p5h})
	_ = Simd_m64_v128_store(m, s0, 332, n0)
	_ = Simd_m64_v128_store(m, s0, 316, n1)
	_ = Simd_m64_v128_store(m, s0, 300, n2)
	_ = Simd_m64_v128_store(m, s0, 284, n3)
	_ = Simd_m64_v128_store(m, s0, 268, n4)
	_ = Simd_m64_v128_store(m, s0, 252, n5)
	return
}

//go:noinline
func Simd_p_fx78(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	n1 := Simd_f32x4_ge(n0, [2]uint64{p0, p0h})
	n2 := Simd_f32x4_neg(n0)
	n3 := Simd_v128_bitselect(n0, n2, n1)
	_ = Simd_m64_v128_store(m, s1, 432, n3)
	n5 := Simd_m64_v128_load(m, s0, 0)
	n6 := Simd_f32x4_ge(n5, [2]uint64{p0, p0h})
	n7 := Simd_f32x4_neg(n5)
	n8 := Simd_v128_bitselect(n5, n7, n6)
	_ = Simd_m64_v128_store(m, s1, 416, n8)
	return n1[0], n1[1], n3[0], n3[1], n6[0], n6[1], n8[0], n8[1]
}

//go:noinline
func Simd_p_fx79(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	n1 := Simd_f32x4_ge(n0, [2]uint64{p0, p0h})
	n2 := Simd_f32x4_neg(n0)
	n3 := Simd_v128_bitselect(n0, n2, n1)
	_ = Simd_m64_v128_store(m, s1, 384, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s1, 400, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s1, 448, n3)
	n7 := Simd_m64_v128_load(m, s0, 48)
	n8 := Simd_f32x4_ge(n7, [2]uint64{p0, p0h})
	n9 := Simd_f32x4_neg(n7)
	n10 := Simd_v128_bitselect(n7, n9, n8)
	_ = Simd_m64_v128_store(m, s1, 464, n10)
	return n1[0], n1[1], n3[0], n3[1], n8[0], n8[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx80(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 64)
	n1 := Simd_f32x4_ge(n0, [2]uint64{p0, p0h})
	n2 := Simd_f32x4_neg(n0)
	n3 := Simd_v128_bitselect(n0, n2, n1)
	_ = Simd_m64_v128_store(m, s1, 480, n3)
	n5 := Simd_m64_v128_load(m, s0, 80)
	n6 := Simd_f32x4_ge(n5, [2]uint64{p0, p0h})
	n7 := Simd_f32x4_neg(n5)
	n8 := Simd_v128_bitselect(n5, n7, n6)
	_ = Simd_m64_v128_store(m, s1, 496, n8)
	return n1[0], n1[1], n3[0], n3[1], n6[0], n6[1], n8[0], n8[1]
}

//go:noinline
func Simd_p_fx81(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 96)
	n1 := Simd_f32x4_ge(n0, [2]uint64{p0, p0h})
	n2 := Simd_f32x4_neg(n0)
	n3 := Simd_v128_bitselect(n0, n2, n1)
	_ = Simd_m64_v128_store(m, s1, 512, n3)
	n5 := Simd_m64_v128_load(m, s0, 112)
	n6 := Simd_f32x4_ge(n5, [2]uint64{p0, p0h})
	n7 := Simd_f32x4_neg(n5)
	n8 := Simd_v128_bitselect(n5, n7, n6)
	_ = Simd_m64_v128_store(m, s1, 528, n8)
	return n1[0], n1[1], n3[0], n3[1], n6[0], n6[1], n8[0], n8[1]
}

//go:noinline
func Simd_p_fx82(m *Module, s0 int64, s1 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, 416)
	n1 := Simd_m64_scalar_i32_shl(s1, 4)
	n2 := Simd_m64_scalar_i32_add(n0, n1)
	n3 := Simd_m64_v128_load(m, n2, 0)
	n4 := Simd_f32x4_splat(f0)
	n5 := Simd_f32x4_mul(n4, n3)
	n6 := Simd_f32x4_add(n5, [2]uint64{p0, p0h})
	n7 := Simd_f32x4_mul(n6, [2]uint64{p1, p1h})
	n8 := Simd_f32x4_add(n7, [2]uint64{p2, p2h})
	n9 := Simd_v128_and(n8, [2]uint64{p3, p3h})
	n10 := Simd_i32x4_max_u(n9, [2]uint64{p4, p4h})
	n11 := Simd_i32x4_sub(n10, [2]uint64{p4, p4h})
	n12 := Simd_i32x4_min_u(n11, [2]uint64{p5, p5h})
	return n3[0], n3[1], n12[0], n12[1]
}

//go:noinline
func Simd_p_fx83(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 352)
	_ = Simd_m64_v128_store(m, s0, 384, n0)
	n2 := Simd_m64_v128_load(m, s0, 368)
	_ = Simd_m64_v128_store(m, s0, 400, n2)
	return
}

//go:noinline
func Simd_p_fx84(m *Module, s0 int64, s1 int64, s2 int64, f0 float32, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_mul(n0, [2]uint64{p0, p0h})
	n2 := Simd_m64_scalar_i32_shl(s1, 10)
	n3 := Simd_m64_scalar_i32_add(s0, n2)
	n4 := Simd_m64_scalar_i32_shl(s2, 6)
	n5 := Simd_m64_scalar_i32_add(n3, n4)
	n6 := Simd_m64_v128_load_rng(m, n5, 48, 0, 64)
	n7 := Simd_f32x4_mul(n6, n6)
	n8 := Simd_f32x4_add(n1, n7)
	n9 := Simd_m64_scalar_i32_shl(s1, 10)
	n10 := Simd_m64_scalar_i32_add(s0, n9)
	n11 := Simd_m64_scalar_i32_shl(s2, 6)
	n12 := Simd_m64_scalar_i32_add(n10, n11)
	n13 := Simd_m64_v128_load_nc(m, n12, 32)
	return n1[0], n1[1], n6[0], n6[1], n8[0], n8[1], n13[0], n13[1]
}

//go:noinline
func Simd_p_fx85(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p1, p1h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add([2]uint64{p0, p0h}, n0)
	n2 := Simd_m64_v128_load_nc(m, s0, 16)
	n3 := Simd_f32x4_mul(n2, n2)
	n4 := Simd_f32x4_add([2]uint64{p0, p0h}, n3)
	n5 := Simd_m64_v128_load_nc(m, s0, 0)
	return n1[0], n1[1], n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx86(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p1, p1h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_sqrt(n1)
	n3 := Simd_f32x4_sqrt([2]uint64{p3, p3h})
	n4 := Simd_f32x4_sqrt([2]uint64{p4, p4h})
	n5 := Simd_f32x4_sqrt([2]uint64{p5, p5h})
	n6 := Simd_f32x4_ge([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n7 := Simd_f32x4_neg([2]uint64{p1, p1h})
	n8 := Simd_v128_bitselect([2]uint64{p1, p1h}, n7, n6)
	n9 := Simd_f32x4_ge([2]uint64{p6, p6h}, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s0, 304, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s0, 272, n3)
	_ = Simd_m64_v128_store(m, s0, 256, n4)
	_ = Simd_m64_v128_store(m, s0, 240, n5)
	_ = Simd_m64_v128_store(m, s0, 224, n2)
	_ = Simd_m64_v128_store(m, s0, 320, n8)
	return n1[0], n1[1], n6[0], n6[1], n8[0], n8[1], n9[0], n9[1]
}

//go:noinline
func Simd_p_fx87(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_neg([2]uint64{p0, p0h})
	n1 := Simd_v128_bitselect([2]uint64{p0, p0h}, n0, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_ge([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_neg([2]uint64{p2, p2h})
	n4 := Simd_v128_bitselect([2]uint64{p2, p2h}, n3, n2)
	n5 := Simd_f32x4_ge([2]uint64{p4, p4h}, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 336, n1)
	_ = Simd_m64_v128_store(m, s0, 352, n4)
	return n1[0], n1[1], n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx88(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_neg([2]uint64{p0, p0h})
	n1 := Simd_v128_bitselect([2]uint64{p0, p0h}, n0, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 368, n1)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx89(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_splat(m, s0, 348)
	n1 := Simd_m64_v128_load32_lane(m, s0, 320, 3, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx90(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_splat(m, s0, 380)
	n1 := Simd_m64_v128_load32_lane(m, s1, 0, 3, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx91(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_add(n1, [2]uint64{p4, p4h})
	n3 := Simd_f32x4_mul(n2, [2]uint64{p5, p5h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p6, p6h})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx92(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add(n0, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_mul(n1, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_add(n2, [2]uint64{p4, p4h})
	n4 := Simd_v128_and(n3, [2]uint64{p5, p5h})
	n5 := Simd_i32x4_max_u(n4, [2]uint64{p6, p6h})
	n6 := Simd_i32x4_sub(n5, [2]uint64{p6, p6h})
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx93(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p2, p2h}, [2]uint64{1084818905618843912, 216736831629295872})
	n2 := Simd_v128_or(n0, n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx94(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{216736831696667908, 216736831629295872})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx95(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 288)
	_ = Simd_m64_v128_store(m, s0, 304, n0)
	return
}

//go:noinline
func Simd_p_fx96(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 324)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_add(n1, [2]uint64{p1, p1h})
	n3 := Simd_f32x4_mul(n2, [2]uint64{p2, p2h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p3, p3h})
	n5 := Simd_v128_and(n4, [2]uint64{p4, p4h})
	n6 := Simd_i32x4_max_u(n5, [2]uint64{p5, p5h})
	n7 := Simd_i32x4_sub(n6, [2]uint64{p5, p5h})
	n8 := Simd_i32x4_min_u(n7, [2]uint64{p6, p6h})
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx97(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, [2]uint64{1663540288323457296, 1084818905618843912})
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_mul(n2, [2]uint64{p4, p4h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p5, p5h})
	n5 := Simd_v128_and(n4, [2]uint64{p6, p6h})
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx98(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 356)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_add(n1, [2]uint64{p1, p1h})
	n3 := Simd_f32x4_mul(n2, [2]uint64{p2, p2h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p3, p3h})
	n5 := Simd_v128_and(n4, [2]uint64{p4, p4h})
	n6 := Simd_i32x4_max_u(n5, [2]uint64{p5, p5h})
	n7 := Simd_i32x4_sub(n6, [2]uint64{p5, p5h})
	n8 := Simd_i32x4_min_u(n7, [2]uint64{p6, p6h})
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx99(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_shuffle(n0, n0, [2]uint64{1084818905618843912, 216736831629295872})
	n2 := Simd_v128_or(n0, n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx100(m *Module, s0 int64, s1 int64, s2 int64, f0 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n1, n0)
	n3 := Simd_m64_v128_load(m, s1, 0)
	n4 := Simd_f32x4_mul(n2, n3)
	_ = Simd_m64_v128_store(m, s2, 0, n4)
	return
}

//go:noinline
func Simd_p_fx101(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_add(n0, n1)
	_ = Simd_m64_v128_store(m, s2, 0, n2)
	return
}

//go:noinline
func Simd_p_fx102(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_i32x4_extend_high_i16x8_u(n0)
	n4 := Simd_i32x4_shl(n3, 16)
	n5 := Simd_m64_v128_load(m, s1, 0)
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_add(n2, n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i32x4_extend_high_i16x8_u(n5)
	n12 := Simd_i32x4_shl(n11, 16)
	n13 := Simd_f32x4_add(n4, n12)
	n14 := Simd_i32x4_shr_u(n13, 16)
	n15 := Simd_v128_and(n14, [2]uint64{p1, p1h})
	n16 := Simd_i16x8_narrow_i32x4_u(n9, n14)
	n17 := Simd_v128_or(n16, [2]uint64{p0, p0h})
	n18 := Simd_i32x4_add(n8, n10)
	n19 := Simd_i32x4_add(n18, [2]uint64{p2, p2h})
	n20 := Simd_i32x4_shr_u(n19, 16)
	n21 := Simd_i32x4_add(n13, n15)
	n22 := Simd_i32x4_add(n21, [2]uint64{p2, p2h})
	n23 := Simd_i32x4_shr_u(n22, 16)
	n24 := Simd_i16x8_narrow_i32x4_u(n20, n23)
	n25 := Simd_f32x4_abs(n8)
	n26 := Simd_i32x4_gt_u(n25, [2]uint64{p3, p3h})
	n27 := Simd_f32x4_abs(n13)
	n28 := Simd_i32x4_gt_u(n27, [2]uint64{p3, p3h})
	n29 := Simd_i8x16_shuffle(n26, n28, [2]uint64{940136352262127872, 2097579117671354640})
	n30 := Simd_v128_bitselect(n17, n24, n29)
	_ = Simd_m64_v128_store(m, s2, 0, n30)
	return
}

//go:noinline
func Simd_p_fx103(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_add(n0, n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	return n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx104(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or(n0, [2]uint64{p3, p3h})
	n2 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p4, p4h})
	n3 := Simd_i32x4_add([2]uint64{p1, p1h}, n2)
	n4 := Simd_i32x4_add(n3, [2]uint64{p5, p5h})
	n5 := Simd_i32x4_shr_u(n4, 16)
	n6 := Simd_i8x16_shuffle(n5, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n7 := Simd_f32x4_abs([2]uint64{p1, p1h})
	n8 := Simd_i32x4_gt_u(n7, [2]uint64{p6, p6h})
	n9 := Simd_i8x16_shuffle(n8, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n10 := Simd_v128_bitselect(n1, n6, n9)
	return n10[0], n10[1]
}

//go:noinline
func Simd_p_fx105(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_add(n0, n2)
	_ = Simd_m64_v128_store(m, s2, 0, n3)
	return
}

//go:noinline
func Simd_p_fx106(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_add(n0, n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx107(m *Module, s0 int64, s1 int64, f0 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_add(n1, n0)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx108(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	n5 := Simd_v128_and(n4, [2]uint64{p2, p2h})
	n6 := Simd_i32x4_extend_high_i16x8_u(n0)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_add([2]uint64{p0, p0h}, n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p2, p2h})
	n11 := Simd_i16x8_narrow_i32x4_u(n4, n9)
	n12 := Simd_v128_or(n11, [2]uint64{p1, p1h})
	n13 := Simd_i32x4_add(n3, n5)
	n14 := Simd_i32x4_add(n13, [2]uint64{p3, p3h})
	n15 := Simd_i32x4_shr_u(n14, 16)
	n16 := Simd_i32x4_add(n8, n10)
	n17 := Simd_i32x4_add(n16, [2]uint64{p3, p3h})
	n18 := Simd_i32x4_shr_u(n17, 16)
	n19 := Simd_i16x8_narrow_i32x4_u(n15, n18)
	n20 := Simd_f32x4_abs(n3)
	n21 := Simd_i32x4_gt_u(n20, [2]uint64{p4, p4h})
	n22 := Simd_f32x4_abs(n8)
	n23 := Simd_i32x4_gt_u(n22, [2]uint64{p4, p4h})
	n24 := Simd_i8x16_shuffle(n21, n23, [2]uint64{940136352262127872, 2097579117671354640})
	n25 := Simd_v128_bitselect(n12, n19, n24)
	_ = Simd_m64_v128_store(m, s1, 0, n25)
	return
}

//go:noinline
func Simd_p_fx109(m *Module, s0 int64, f0 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_add(n1, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx110(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_sub(n0, n1)
	_ = Simd_m64_v128_store(m, s2, 0, n2)
	return
}

//go:noinline
func Simd_p_fx111(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_i32x4_extend_high_i16x8_u(n0)
	n4 := Simd_i32x4_shl(n3, 16)
	n5 := Simd_m64_v128_load(m, s1, 0)
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_sub(n2, n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i32x4_extend_high_i16x8_u(n5)
	n12 := Simd_i32x4_shl(n11, 16)
	n13 := Simd_f32x4_sub(n4, n12)
	n14 := Simd_i32x4_shr_u(n13, 16)
	n15 := Simd_v128_and(n14, [2]uint64{p1, p1h})
	n16 := Simd_i16x8_narrow_i32x4_u(n9, n14)
	n17 := Simd_v128_or(n16, [2]uint64{p0, p0h})
	n18 := Simd_i32x4_add(n8, n10)
	n19 := Simd_i32x4_add(n18, [2]uint64{p2, p2h})
	n20 := Simd_i32x4_shr_u(n19, 16)
	n21 := Simd_i32x4_add(n13, n15)
	n22 := Simd_i32x4_add(n21, [2]uint64{p2, p2h})
	n23 := Simd_i32x4_shr_u(n22, 16)
	n24 := Simd_i16x8_narrow_i32x4_u(n20, n23)
	n25 := Simd_f32x4_abs(n8)
	n26 := Simd_i32x4_gt_u(n25, [2]uint64{p3, p3h})
	n27 := Simd_f32x4_abs(n13)
	n28 := Simd_i32x4_gt_u(n27, [2]uint64{p3, p3h})
	n29 := Simd_i8x16_shuffle(n26, n28, [2]uint64{940136352262127872, 2097579117671354640})
	n30 := Simd_v128_bitselect(n17, n24, n29)
	_ = Simd_m64_v128_store(m, s2, 0, n30)
	return
}

//go:noinline
func Simd_p_fx112(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_m64_v128_load(m, s1, 0)
	n3 := Simd_f32x4_sub(n1, n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	return n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx113(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_m64_v128_load(m, s1, 0)
	n3 := Simd_f32x4_sub(n1, n2)
	_ = Simd_m64_v128_store(m, s2, 0, n3)
	return
}

//go:noinline
func Simd_p_fx114(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_mul(n0, n1)
	_ = Simd_m64_v128_store(m, s2, 0, n2)
	return
}

//go:noinline
func Simd_p_fx115(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_i32x4_extend_high_i16x8_u(n0)
	n4 := Simd_i32x4_shl(n3, 16)
	n5 := Simd_m64_v128_load(m, s1, 0)
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_mul(n2, n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i32x4_extend_high_i16x8_u(n5)
	n12 := Simd_i32x4_shl(n11, 16)
	n13 := Simd_f32x4_mul(n4, n12)
	n14 := Simd_i32x4_shr_u(n13, 16)
	n15 := Simd_v128_and(n14, [2]uint64{p1, p1h})
	n16 := Simd_i16x8_narrow_i32x4_u(n9, n14)
	n17 := Simd_v128_or(n16, [2]uint64{p0, p0h})
	n18 := Simd_i32x4_add(n8, n10)
	n19 := Simd_i32x4_add(n18, [2]uint64{p2, p2h})
	n20 := Simd_i32x4_shr_u(n19, 16)
	n21 := Simd_i32x4_add(n13, n15)
	n22 := Simd_i32x4_add(n21, [2]uint64{p2, p2h})
	n23 := Simd_i32x4_shr_u(n22, 16)
	n24 := Simd_i16x8_narrow_i32x4_u(n20, n23)
	n25 := Simd_f32x4_abs(n8)
	n26 := Simd_i32x4_gt_u(n25, [2]uint64{p3, p3h})
	n27 := Simd_f32x4_abs(n13)
	n28 := Simd_i32x4_gt_u(n27, [2]uint64{p3, p3h})
	n29 := Simd_i8x16_shuffle(n26, n28, [2]uint64{940136352262127872, 2097579117671354640})
	n30 := Simd_v128_bitselect(n17, n24, n29)
	_ = Simd_m64_v128_store(m, s2, 0, n30)
	return
}

//go:noinline
func Simd_p_fx116(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_mul(n0, n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	return n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx117(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_mul(n0, n2)
	_ = Simd_m64_v128_store(m, s2, 0, n3)
	return
}

//go:noinline
func Simd_p_fx118(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_div(n0, n1)
	_ = Simd_m64_v128_store(m, s2, 0, n2)
	return
}

//go:noinline
func Simd_p_fx119(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_i32x4_extend_high_i16x8_u(n0)
	n4 := Simd_i32x4_shl(n3, 16)
	n5 := Simd_m64_v128_load(m, s1, 0)
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_div(n2, n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i32x4_extend_high_i16x8_u(n5)
	n12 := Simd_i32x4_shl(n11, 16)
	n13 := Simd_f32x4_div(n4, n12)
	n14 := Simd_i32x4_shr_u(n13, 16)
	n15 := Simd_v128_and(n14, [2]uint64{p1, p1h})
	n16 := Simd_i16x8_narrow_i32x4_u(n9, n14)
	n17 := Simd_v128_or(n16, [2]uint64{p0, p0h})
	n18 := Simd_i32x4_add(n8, n10)
	n19 := Simd_i32x4_add(n18, [2]uint64{p2, p2h})
	n20 := Simd_i32x4_shr_u(n19, 16)
	n21 := Simd_i32x4_add(n13, n15)
	n22 := Simd_i32x4_add(n21, [2]uint64{p2, p2h})
	n23 := Simd_i32x4_shr_u(n22, 16)
	n24 := Simd_i16x8_narrow_i32x4_u(n20, n23)
	n25 := Simd_f32x4_abs(n8)
	n26 := Simd_i32x4_gt_u(n25, [2]uint64{p3, p3h})
	n27 := Simd_f32x4_abs(n13)
	n28 := Simd_i32x4_gt_u(n27, [2]uint64{p3, p3h})
	n29 := Simd_i8x16_shuffle(n26, n28, [2]uint64{940136352262127872, 2097579117671354640})
	n30 := Simd_v128_bitselect(n17, n24, n29)
	_ = Simd_m64_v128_store(m, s2, 0, n30)
	return
}

//go:noinline
func Simd_p_fx120(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_m64_v128_load(m, s1, 0)
	n3 := Simd_f32x4_div(n1, n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	return n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx121(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_m64_v128_load(m, s1, 0)
	n3 := Simd_f32x4_div(n1, n2)
	_ = Simd_m64_v128_store(m, s2, 0, n3)
	return
}

//go:noinline
func Simd_p_fx122(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx123(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_mul(n2, n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_i32x4_extend_high_i16x8_u(n0)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_mul(n7, n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i16x8_narrow_i32x4_u(n4, n9)
	n12 := Simd_v128_or(n11, [2]uint64{p0, p0h})
	n13 := Simd_i32x4_add(n3, n5)
	n14 := Simd_i32x4_add(n13, [2]uint64{p2, p2h})
	n15 := Simd_i32x4_shr_u(n14, 16)
	n16 := Simd_i32x4_add(n8, n10)
	n17 := Simd_i32x4_add(n16, [2]uint64{p2, p2h})
	n18 := Simd_i32x4_shr_u(n17, 16)
	n19 := Simd_i16x8_narrow_i32x4_u(n15, n18)
	n20 := Simd_f32x4_abs(n3)
	n21 := Simd_i32x4_gt_u(n20, [2]uint64{p3, p3h})
	n22 := Simd_f32x4_abs(n8)
	n23 := Simd_i32x4_gt_u(n22, [2]uint64{p3, p3h})
	n24 := Simd_i8x16_shuffle(n21, n23, [2]uint64{940136352262127872, 2097579117671354640})
	n25 := Simd_v128_bitselect(n12, n19, n24)
	_ = Simd_m64_v128_store(m, s1, 0, n25)
	return
}

//go:noinline
func Simd_p_fx124(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_mul(n1, n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx125(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_sqrt(n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx126(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_sqrt(n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_i32x4_extend_high_i16x8_u(n0)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_sqrt(n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i16x8_narrow_i32x4_u(n4, n9)
	n12 := Simd_v128_or(n11, [2]uint64{p0, p0h})
	n13 := Simd_i32x4_add(n3, n5)
	n14 := Simd_i32x4_add(n13, [2]uint64{p2, p2h})
	n15 := Simd_i32x4_shr_u(n14, 16)
	n16 := Simd_i32x4_add(n8, n10)
	n17 := Simd_i32x4_add(n16, [2]uint64{p2, p2h})
	n18 := Simd_i32x4_shr_u(n17, 16)
	n19 := Simd_i16x8_narrow_i32x4_u(n15, n18)
	n20 := Simd_f32x4_abs(n3)
	n21 := Simd_i32x4_gt_u(n20, [2]uint64{p3, p3h})
	n22 := Simd_f32x4_abs(n8)
	n23 := Simd_i32x4_gt_u(n22, [2]uint64{p3, p3h})
	n24 := Simd_i8x16_shuffle(n21, n23, [2]uint64{940136352262127872, 2097579117671354640})
	n25 := Simd_v128_bitselect(n12, n19, n24)
	_ = Simd_m64_v128_store(m, s1, 0, n25)
	return
}

//go:noinline
func Simd_p_fx127(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_sqrt(n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx128(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx129(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i64x2_add(n0, [2]uint64{p0, p0h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx130(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_lt_s([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i64x2_lt_s([2]uint64{p2, p2h}, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{795458214199165184, 1952900979608391952})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx131(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	n3 := Simd_m64_v128_load(m, s0+16, 0)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	_ = Simd_m64_v128_store(m, s0+16, 0, n4)
	n6 := Simd_m64_v128_load(m, s0+32, 0)
	n7 := Simd_f32x4_mul([2]uint64{p0, p0h}, n6)
	_ = Simd_m64_v128_store(m, s0+32, 0, n7)
	n9 := Simd_m64_v128_load(m, s0+48, 0)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	_ = Simd_m64_v128_store(m, s0+48, 0, n10)
	return
}

//go:noinline
func Simd_p_fx132(m *Module, s0 int64, f0 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n1, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx133(m *Module, s0 int64, s1 int64, s2 int64, f0 float32, f1 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f1)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_m64_v128_load(m, s1, 0)
	n4 := Simd_f32x4_add(n2, n3)
	n5 := Simd_f32x4_splat(f0)
	n6 := Simd_f32x4_mul(n5, n4)
	_ = Simd_m64_v128_store(m, s2, 0, n6)
	return
}

//go:noinline
func Simd_p_fx134(m *Module, s0 int64, f0 float32) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 0)
	n1 := Simd_m64_v128_load32_lane(m, s0+1, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s0+2, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s0+3, 0, 3, n2)
	n4 := Simd_f32x4_splat(f0)
	n5 := Simd_f32x4_mul(n4, n3)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx135(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_m64_v128_load(m, s1, 0)
	n3 := Simd_f32x4_add(n1, n2)
	_ = Simd_m64_v128_store(m, s1, 0, n3)
	n5 := Simd_m64_v128_load(m, s0+16, 0)
	n6 := Simd_f32x4_mul([2]uint64{p0, p0h}, n5)
	n7 := Simd_m64_v128_load(m, s1+16, 0)
	n8 := Simd_f32x4_add(n6, n7)
	_ = Simd_m64_v128_store(m, s1+16, 0, n8)
	n10 := Simd_m64_v128_load(m, s0+32, 0)
	n11 := Simd_f32x4_mul([2]uint64{p0, p0h}, n10)
	n12 := Simd_m64_v128_load(m, s1+32, 0)
	n13 := Simd_f32x4_add(n11, n12)
	_ = Simd_m64_v128_store(m, s1+32, 0, n13)
	n15 := Simd_m64_v128_load(m, s0+48, 0)
	n16 := Simd_f32x4_mul([2]uint64{p0, p0h}, n15)
	n17 := Simd_m64_v128_load(m, s1+48, 0)
	n18 := Simd_f32x4_add(n16, n17)
	_ = Simd_m64_v128_store(m, s1+48, 0, n18)
	return
}

//go:noinline
func Simd_p_fx136(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul(n0, [2]uint64{p0, p0h})
	n2 := Simd_m64_v128_load(m, s1, 0)
	n3 := Simd_f32x4_add(n1, n2)
	_ = Simd_m64_v128_store(m, s1, 0, n3)
	return
}

//go:noinline
func Simd_p_fx137(m *Module, s0 int64, s1 int64, f0 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_m64_v128_load(m, s1, 0)
	n4 := Simd_f32x4_add(n2, n3)
	_ = Simd_m64_v128_store(m, s1, 0, n4)
	return
}

//go:noinline
func Simd_p_fx138(m *Module, p0, p0h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_shl(n0, 17)
	return n0[0], n0[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx139(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_s([2]uint64{p0, p0h})
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx140(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_add(n1, [2]uint64{p3, p3h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx141(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 4)
	n1 := Simd_v128_or(n0, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_mul(n1, [2]uint64{p2, p2h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx142(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load32_splat(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx143(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load32_splat(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1+16, 0, n0)
	return
}

//go:noinline
func Simd_p_fx144(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx145(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load(m, n0, 0)
	n2 := Simd_m64_v128_load_rng(m, s2, 0, 0, 32)
	n3 := Simd_f32x4_mul(n1, n2)
	n4 := Simd_f32x4_add([2]uint64{p0, p0h}, n3)
	n5 := Simd_m64_v128_load(m, s3, 0)
	n6 := Simd_m64_v128_load_nc(m, s2+16, 0)
	n7 := Simd_f32x4_mul(n5, n6)
	n8 := Simd_f32x4_add(n4, n7)
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx146(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	_ = Simd_m64_v128_store(m, s0, 0, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s0, 16)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx147(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load(m, n0, 16)
	n2 := Simd_m64_v128_load_rng(m, s2, 0, 0, 32)
	n3 := Simd_f32x4_mul(n1, n2)
	n4 := Simd_f32x4_add([2]uint64{p0, p0h}, n3)
	n5 := Simd_m64_v128_load(m, s3, 16)
	n6 := Simd_m64_v128_load_nc(m, s2+16, 0)
	n7 := Simd_f32x4_mul(n5, n6)
	n8 := Simd_f32x4_add(n4, n7)
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx148(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	_ = Simd_m64_v128_store(m, s0, 16, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s0, 32)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx149(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load(m, n0, 32)
	n2 := Simd_m64_v128_load_rng(m, s2, 0, 0, 32)
	n3 := Simd_f32x4_mul(n1, n2)
	n4 := Simd_f32x4_add([2]uint64{p0, p0h}, n3)
	n5 := Simd_m64_v128_load(m, s3, 32)
	n6 := Simd_m64_v128_load_nc(m, s2+16, 0)
	n7 := Simd_f32x4_mul(n5, n6)
	n8 := Simd_f32x4_add(n4, n7)
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx150(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	_ = Simd_m64_v128_store(m, s0, 32, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s0, 48)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx151(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load(m, n0, 48)
	n2 := Simd_m64_v128_load_rng(m, s2, 0, 0, 32)
	n3 := Simd_f32x4_mul(n1, n2)
	n4 := Simd_f32x4_add([2]uint64{p0, p0h}, n3)
	n5 := Simd_m64_v128_load(m, s3, 48)
	n6 := Simd_m64_v128_load_nc(m, s2+16, 0)
	n7 := Simd_f32x4_mul(n5, n6)
	n8 := Simd_f32x4_add(n4, n7)
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx152(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul([2]uint64{p1, p1h}, n0)
	n2 := Simd_f32x4_add([2]uint64{p0, p0h}, n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	n4 := Simd_m64_v128_load(m, s0+16, 0)
	n5 := Simd_f32x4_mul([2]uint64{p1, p1h}, n4)
	n6 := Simd_f32x4_add([2]uint64{p0, p0h}, n5)
	_ = Simd_m64_v128_store(m, s1+16, 0, n6)
	n8 := Simd_m64_v128_load(m, s0+32, 0)
	n9 := Simd_f32x4_mul([2]uint64{p1, p1h}, n8)
	n10 := Simd_f32x4_add([2]uint64{p0, p0h}, n9)
	_ = Simd_m64_v128_store(m, s1+32, 0, n10)
	n12 := Simd_m64_v128_load(m, s0+48, 0)
	n13 := Simd_f32x4_mul([2]uint64{p1, p1h}, n12)
	n14 := Simd_f32x4_add([2]uint64{p0, p0h}, n13)
	_ = Simd_m64_v128_store(m, s1+48, 0, n14)
	return
}

//go:noinline
func Simd_p_fx153(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul(n0, [2]uint64{p0, p0h})
	n2 := Simd_f32x4_add(n1, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx154(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	return
}

//go:noinline
func Simd_p_fx155(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n1 := Simd_f16x4_cvt(n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	return
}

//go:noinline
func Simd_p_fx156(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx157(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_add(n0, n1)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx158(m *Module, s0 int64, s1 int64, f0 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n1, n0)
	n3 := Simd_m64_v128_load(m, s1, 0)
	n4 := Simd_f32x4_add(n2, n3)
	_ = Simd_m64_v128_store(m, s1, 0, n4)
	return
}

//go:noinline
func Simd_p_fx159(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_pmax(n0, [2]uint64{p0, p0h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx160(m *Module, s0 int64, f0 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n0, n1)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx161(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 64)
	n1 := Simd_m64_scalar_i32_add(s1, s2)
	n2 := Simd_m64_v128_load_rng(m, n1, 0, 0, 64)
	n3 := Simd_f32x4_mul(n0, n2)
	n4 := Simd_f32x4_add(n3, [2]uint64{p0, p0h})
	n5 := Simd_m64_v128_load_nc(m, s0+48, 0)
	n6 := Simd_m64_scalar_i32_add(s1, s2)
	n7 := Simd_m64_scalar_i32_add(n6, 48)
	n8 := Simd_m64_v128_load_nc(m, n7, 0)
	n9 := Simd_f32x4_mul(n5, n8)
	n10 := Simd_f32x4_add(n9, [2]uint64{p1, p1h})
	n11 := Simd_m64_v128_load_nc(m, s0+32, 0)
	n12 := Simd_m64_scalar_i32_add(s1, s2)
	n13 := Simd_m64_scalar_i32_add(n12, 32)
	n14 := Simd_m64_v128_load_nc(m, n13, 0)
	n15 := Simd_f32x4_mul(n11, n14)
	n16 := Simd_f32x4_add(n15, [2]uint64{p2, p2h})
	n17 := Simd_m64_v128_load_nc(m, s0+16, 0)
	n18 := Simd_m64_scalar_i32_add(s1, s2)
	n19 := Simd_m64_scalar_i32_add(n18, 16)
	n20 := Simd_m64_v128_load_nc(m, n19, 0)
	n21 := Simd_f32x4_mul(n17, n20)
	n22 := Simd_f32x4_add(n21, [2]uint64{p3, p3h})
	return n4[0], n4[1], n10[0], n10[1], n16[0], n16[1], n22[0], n22[1]
}

//go:noinline
func Simd_p_fx162(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_f32x4_add(n0, n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx163(m *Module, s0 int64, f0 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_sub(n0, n1)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx164(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_mul(n0, n1)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx165(m *Module, s0 int64, s1 int64, f0 float32, f1 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f1)
	n2 := Simd_f32x4_pmin(n1, n0)
	n3 := Simd_f32x4_splat(f0)
	n4 := Simd_f32x4_pmax(n3, n2)
	_ = Simd_m64_v128_store(m, s1, 0, n4)
	return
}

//go:noinline
func Simd_p_fx166(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 1)
	n2 := Simd_i32x4_max_u(n1, [2]uint64{p2, p2h})
	n3 := Simd_i32x4_shr_u(n2, 1)
	n4 := Simd_v128_and(n3, [2]uint64{p3, p3h})
	n5 := Simd_i32x4_add(n4, [2]uint64{p4, p4h})
	n6 := Simd_f32x4_abs(n0)
	n7 := Simd_f32x4_mul(n6, [2]uint64{p0, p0h})
	n8 := Simd_f32x4_mul(n7, [2]uint64{p1, p1h})
	n9 := Simd_f32x4_add(n8, n5)
	return n0[0], n0[1], n1[0], n1[1], n9[0], n9[1]
}

//go:noinline
func Simd_p_fx167(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p1, p1h}, 13)
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p3, p3h})
	n3 := Simd_i32x4_add(n1, n2)
	n4 := Simd_i32x4_gt_u([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n5 := Simd_v128_bitselect([2]uint64{p0, p0h}, n3, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx168(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 16)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx169(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx170(m *Module, s0 int32, s1 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_add(n0, [2]uint64{p0, p0h})
	n2 := Simd_i32x4_gt_s(n1, [2]uint64{p2, p2h})
	n3 := Simd_i32x4_splat(s1)
	n4 := Simd_i32x4_lt_s(n1, n3)
	n5 := Simd_v128_and(n4, n2)
	n6 := Simd_i32x4_sub([2]uint64{p1, p1h}, n5)
	n7 := Simd_i32x4_add([2]uint64{p0, p0h}, [2]uint64{p3, p3h})
	return n6[0], n6[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx171(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{1084818905618843912, 216736831629295872})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx172(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{216736831696667908, 216736831629295872})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx173(m *Module, s0 int64, s1 int64, f0 float32, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_lt(n0, [2]uint64{p0, p0h})
	n2 := Simd_v128_and(n1, n0)
	n3 := Simd_f32x4_splat(f0)
	n4 := Simd_f32x4_mul(n3, n2)
	n5 := Simd_f32x4_gt(n0, [2]uint64{p0, p0h})
	n6 := Simd_v128_and(n5, n0)
	n7 := Simd_f32x4_add(n4, n6)
	_ = Simd_m64_v128_store(m, s1, 0, n7)
	return
}

//go:noinline
func Simd_p_fx174(m *Module, s0 int64, s1 int64, f0 float32, f1 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_m64_v128_load(m, s1, 0)
	n4 := Simd_f32x4_splat(f1)
	n5 := Simd_f32x4_mul(n4, n3)
	n6 := Simd_f32x4_add(n2, n5)
	_ = Simd_m64_v128_store(m, s0, 0, n6)
	return
}

//go:noinline
func Simd_p_fx175(m *Module, s0 int64, s1 int64, f0 float32) (uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_f32x4_splat(f0)
	n3 := Simd_f32x4_mul(n2, n1)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx176(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_m64_scalar_i32_add(s0, s1)
	n3 := Simd_m64_v128_load16x4_u(m, n2, 0)
	n4 := Simd_f16x4_cvt(n3)
	n5 := Simd_i8x16_shuffle(n1, n4, [2]uint64{p0, p0h})
	return n1[0], n1[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx177(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_scalar_i32_add(n0, s1)
	n2 := Simd_m64_v128_load16x4_u(m, n1, 0)
	n3 := Simd_f16x4_cvt(n2)
	n4 := Simd_m64_scalar_i32_add(s0, s1)
	n5 := Simd_m64_scalar_i32_add(n4, s1)
	n6 := Simd_m64_scalar_i32_add(n5, s1)
	n7 := Simd_m64_v128_load16x4_u(m, n6, 0)
	n8 := Simd_f16x4_cvt(n7)
	n9 := Simd_i8x16_shuffle(n3, n8, [2]uint64{p0, p0h})
	n10 := Simd_i8x16_shuffle([2]uint64{p1, p1h}, n9, [2]uint64{p2, p2h})
	return n3[0], n3[1], n8[0], n8[1], n9[0], n9[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx178(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n1 := Simd_i8x16_shuffle([2]uint64{p4, p4h}, [2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+33536, 0, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0+33280, 0, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx179(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i8x16_shuffle([2]uint64{p3, p3h}, n0, [2]uint64{p4, p4h})
	n2 := Simd_i8x16_shuffle([2]uint64{p3, p3h}, n0, [2]uint64{p5, p5h})
	_ = Simd_m64_v128_store(m, s0+33024, 0, n1)
	_ = Simd_m64_v128_store(m, s0+32768, 0, n2)
	return
}

//go:noinline
func Simd_p_fx180(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 16400)
	n1 := Simd_m64_v128_load_nc(m, s0+16384, 0)
	n2 := Simd_f32x4_add(n0, n1)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx181(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load_nc(m, s0, 160)
	n1 := Simd_m64_v128_load_nc(m, s0, 144)
	n2 := Simd_m64_v128_load_nc(m, s0, 128)
	n3 := Simd_m64_v128_load_nc(m, s0, 112)
	n4 := Simd_m64_v128_load_nc(m, s0, 96)
	n5 := Simd_m64_v128_load_nc(m, s0, 80)
	n6 := Simd_m64_v128_load_nc(m, s0, 64)
	n7 := Simd_m64_v128_load_nc(m, s0, 48)
	n8 := Simd_m64_v128_load_nc(m, s0, 32)
	n9 := Simd_m64_v128_load_nc(m, s0, 16)
	n10 := Simd_m64_v128_load_nc(m, s0, 0)
	n11 := Simd_f32x4_pmax(n10, [2]uint64{p5, p5h})
	n12 := Simd_f32x4_pmax(n9, n11)
	n13 := Simd_f32x4_pmax(n8, n12)
	n14 := Simd_f32x4_pmax(n7, n13)
	n15 := Simd_f32x4_pmax(n6, n14)
	n16 := Simd_f32x4_pmax(n5, n15)
	n17 := Simd_f32x4_pmax(n4, n16)
	n18 := Simd_f32x4_pmax(n3, n17)
	n19 := Simd_f32x4_pmax(n2, n18)
	n20 := Simd_f32x4_pmax(n1, n19)
	n21 := Simd_f32x4_pmax(n0, n20)
	n22 := Simd_f32x4_pmax([2]uint64{p4, p4h}, n21)
	n23 := Simd_f32x4_pmax([2]uint64{p3, p3h}, n22)
	n24 := Simd_f32x4_pmax([2]uint64{p2, p2h}, n23)
	n25 := Simd_f32x4_pmax([2]uint64{p1, p1h}, n24)
	n26 := Simd_f32x4_pmax([2]uint64{p0, p0h}, n25)
	return n26[0], n26[1]
}

//go:noinline
func Simd_p_fx182(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	n3 := Simd_m64_v128_load(m, s1+32784, 0)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	_ = Simd_m64_v128_store(m, s1+32784, 0, n4)
	n6 := Simd_m64_v128_load(m, s1+32800, 0)
	n7 := Simd_f32x4_mul([2]uint64{p0, p0h}, n6)
	_ = Simd_m64_v128_store(m, s1+32800, 0, n7)
	n9 := Simd_m64_v128_load(m, s1+32816, 0)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	_ = Simd_m64_v128_store(m, s1+32816, 0, n10)
	return
}

//go:noinline
func Simd_p_fx183(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	n3 := Simd_m64_v128_load(m, s1+80, 0)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	_ = Simd_m64_v128_store(m, s1+80, 0, n4)
	n6 := Simd_m64_v128_load(m, s1+96, 0)
	n7 := Simd_f32x4_mul([2]uint64{p0, p0h}, n6)
	_ = Simd_m64_v128_store(m, s1+96, 0, n7)
	n9 := Simd_m64_v128_load(m, s1+112, 0)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	_ = Simd_m64_v128_store(m, s1+112, 0, n10)
	return
}

//go:noinline
func Simd_p_fx184(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_scalar_i32_add(s0, s1)
	n2 := Simd_m64_v128_load(m, n1, 0)
	n3 := Simd_f32x4_mul(n0, n2)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	return
}

//go:noinline
func Simd_p_fx185(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_m64_v128_load(m, s1, 0)
	n3 := Simd_f32x4_add(n1, n2)
	_ = Simd_m64_v128_store(m, s1, 0, n3)
	n5 := Simd_m64_v128_load(m, s2+80, 0)
	n6 := Simd_f32x4_mul([2]uint64{p0, p0h}, n5)
	n7 := Simd_m64_v128_load(m, s1+16, 0)
	n8 := Simd_f32x4_add(n6, n7)
	_ = Simd_m64_v128_store(m, s1+16, 0, n8)
	n10 := Simd_m64_v128_load(m, s2+96, 0)
	n11 := Simd_f32x4_mul([2]uint64{p0, p0h}, n10)
	n12 := Simd_m64_v128_load(m, s1+32, 0)
	n13 := Simd_f32x4_add(n11, n12)
	_ = Simd_m64_v128_store(m, s1+32, 0, n13)
	n15 := Simd_m64_v128_load(m, s2+112, 0)
	n16 := Simd_f32x4_mul([2]uint64{p0, p0h}, n15)
	n17 := Simd_m64_v128_load(m, s1+48, 0)
	n18 := Simd_f32x4_add(n16, n17)
	_ = Simd_m64_v128_store(m, s1+48, 0, n18)
	return
}

//go:noinline
func Simd_p_fx186(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64, s7 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_scalar_i32_add(s2, s3)
	n2 := Simd_m64_scalar_i32_add(s1, n1)
	n3 := Simd_m64_v128_load(m, n2, 0)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	n5 := Simd_m64_scalar_i32_add(s5, s6)
	n6 := Simd_m64_scalar_i32_add(s4, n5)
	n7 := Simd_m64_scalar_i32_add(n6, s7)
	n8 := Simd_m64_scalar_i32_add(s1, n7)
	n9 := Simd_m64_v128_load(m, n8, 0)
	n10 := Simd_f32x4_mul([2]uint64{p1, p1h}, n9)
	n11 := Simd_f32x4_add(n4, n10)
	return n0[0], n0[1], n11[0], n11[1]
}

//go:noinline
func Simd_p_fx187(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	_ = Simd_m64_v128_store(m, s0, 0, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s1+16, 0)
	n2 := Simd_m64_v128_load(m, s2+16, 0)
	n3 := Simd_f32x4_mul([2]uint64{p1, p1h}, n2)
	n4 := Simd_m64_v128_load(m, s3+16, 0)
	n5 := Simd_f32x4_mul([2]uint64{p2, p2h}, n4)
	n6 := Simd_f32x4_add(n3, n5)
	_ = Simd_m64_v128_store(m, s0+16, 0, n6)
	n8 := Simd_m64_v128_load(m, s1+32, 0)
	n9 := Simd_m64_v128_load(m, s2+32, 0)
	n10 := Simd_f32x4_mul([2]uint64{p1, p1h}, n9)
	n11 := Simd_m64_v128_load(m, s3+32, 0)
	n12 := Simd_f32x4_mul([2]uint64{p2, p2h}, n11)
	n13 := Simd_f32x4_add(n10, n12)
	_ = Simd_m64_v128_store(m, s0+32, 0, n13)
	return n1[0], n1[1], n6[0], n6[1], n8[0], n8[1], n13[0], n13[1]
}

//go:noinline
func Simd_p_fx188(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n1 := Simd_f32x4_add([2]uint64{p2, p2h}, n0)
	n2 := Simd_m64_v128_load(m, s0+48, 0)
	n3 := Simd_m64_v128_load(m, s1+48, 0)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	n5 := Simd_m64_v128_load(m, s2+48, 0)
	n6 := Simd_f32x4_mul([2]uint64{p1, p1h}, n5)
	n7 := Simd_f32x4_add(n4, n6)
	_ = Simd_m64_v128_store(m, s3+48, 0, n7)
	return n2[0], n2[1], n7[0], n7[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx189(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx190(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_lt(n0, [2]uint64{p1, p1h})
	n2 := Simd_v128_and(n1, [2]uint64{p2, p2h})
	n3 := Simd_f32x4_gt(n0, [2]uint64{p1, p1h})
	n4 := Simd_v128_bitselect([2]uint64{p0, p0h}, n2, n3)
	_ = Simd_m64_v128_store(m, s1, 0, n4)
	return
}

//go:noinline
func Simd_p_fx191(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_lt(n2, [2]uint64{p1, p1h})
	n4 := Simd_v128_and(n3, [2]uint64{p2, p2h})
	n5 := Simd_f32x4_gt(n2, [2]uint64{p1, p1h})
	n6 := Simd_v128_bitselect([2]uint64{p0, p0h}, n4, n5)
	n7 := Simd_i32x4_shr_u(n6, 16)
	n8 := Simd_i32x4_extend_high_i16x8_u(n0)
	n9 := Simd_i32x4_shl(n8, 16)
	n10 := Simd_f32x4_lt(n9, [2]uint64{p1, p1h})
	n11 := Simd_v128_and(n10, [2]uint64{p2, p2h})
	n12 := Simd_f32x4_gt(n9, [2]uint64{p1, p1h})
	n13 := Simd_v128_bitselect([2]uint64{p0, p0h}, n11, n12)
	n14 := Simd_i32x4_shr_u(n13, 16)
	n15 := Simd_i16x8_narrow_i32x4_u(n7, n14)
	n16 := Simd_v128_or(n15, [2]uint64{p3, p3h})
	n17 := Simd_f32x4_abs(n6)
	n18 := Simd_i32x4_gt_u(n17, [2]uint64{p4, p4h})
	n19 := Simd_f32x4_abs(n13)
	n20 := Simd_i32x4_gt_u(n19, [2]uint64{p4, p4h})
	n21 := Simd_i8x16_shuffle(n18, n20, [2]uint64{940136352262127872, 2097579117671354640})
	n22 := Simd_v128_bitselect(n16, n15, n21)
	_ = Simd_m64_v128_store(m, s1, 0, n22)
	return
}

//go:noinline
func Simd_p_fx192(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_lt(n1, [2]uint64{p1, p1h})
	n3 := Simd_v128_and(n2, [2]uint64{p2, p2h})
	n4 := Simd_f32x4_gt(n1, [2]uint64{p1, p1h})
	n5 := Simd_v128_bitselect([2]uint64{p0, p0h}, n3, n4)
	_ = Simd_m64_v128_store(m, s1, 0, n5)
	return
}

//go:noinline
func Simd_p_fx193(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_neg(n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx194(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_neg(n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_i32x4_extend_high_i16x8_u(n0)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_neg(n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i16x8_narrow_i32x4_u(n4, n9)
	n12 := Simd_v128_or(n11, [2]uint64{p0, p0h})
	n13 := Simd_i32x4_add(n3, n5)
	n14 := Simd_i32x4_add(n13, [2]uint64{p2, p2h})
	n15 := Simd_i32x4_shr_u(n14, 16)
	n16 := Simd_i32x4_add(n8, n10)
	n17 := Simd_i32x4_add(n16, [2]uint64{p2, p2h})
	n18 := Simd_i32x4_shr_u(n17, 16)
	n19 := Simd_i16x8_narrow_i32x4_u(n15, n18)
	n20 := Simd_f32x4_abs(n2)
	n21 := Simd_i32x4_gt_u(n20, [2]uint64{p3, p3h})
	n22 := Simd_f32x4_abs(n7)
	n23 := Simd_i32x4_gt_u(n22, [2]uint64{p3, p3h})
	n24 := Simd_i8x16_shuffle(n21, n23, [2]uint64{940136352262127872, 2097579117671354640})
	n25 := Simd_v128_bitselect(n12, n19, n24)
	_ = Simd_m64_v128_store(m, s1, 0, n25)
	return
}

//go:noinline
func Simd_p_fx195(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_neg(n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx196(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_gt(n0, [2]uint64{p0, p0h})
	n2 := Simd_v128_and(n1, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx197(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_gt(n2, [2]uint64{p0, p0h})
	n4 := Simd_v128_and(n3, [2]uint64{p1, p1h})
	n5 := Simd_i32x4_shr_u(n4, 16)
	n6 := Simd_i32x4_extend_high_i16x8_u(n0)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_gt(n7, [2]uint64{p0, p0h})
	n9 := Simd_v128_and(n8, [2]uint64{p1, p1h})
	n10 := Simd_i32x4_shr_u(n9, 16)
	n11 := Simd_i16x8_narrow_i32x4_u(n5, n10)
	n12 := Simd_v128_or(n11, [2]uint64{p2, p2h})
	n13 := Simd_i32x4_gt_u(n4, [2]uint64{p3, p3h})
	n14 := Simd_i32x4_gt_u(n9, [2]uint64{p3, p3h})
	n15 := Simd_i8x16_shuffle(n13, n14, [2]uint64{940136352262127872, 2097579117671354640})
	n16 := Simd_v128_bitselect(n12, n11, n15)
	_ = Simd_m64_v128_store(m, s1, 0, n16)
	return
}

//go:noinline
func Simd_p_fx198(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_gt(n1, [2]uint64{p0, p0h})
	n3 := Simd_v128_and(n2, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s1, 0, n3)
	return
}

//go:noinline
func Simd_p_fx199(m *Module, s0 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_high_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx200(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_shl(n0, 16)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx201(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 16)
	n1 := Simd_v128_and(n0, [2]uint64{p3, p3h})
	n2 := Simd_i32x4_add([2]uint64{p0, p0h}, n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p4, p4h})
	n4 := Simd_i32x4_shr_u(n3, 16)
	n5 := Simd_i32x4_shr_u([2]uint64{p1, p1h}, 16)
	n6 := Simd_v128_and(n5, [2]uint64{p3, p3h})
	n7 := Simd_i16x8_narrow_i32x4_u(n0, n5)
	n8 := Simd_v128_or(n7, [2]uint64{p2, p2h})
	n9 := Simd_i32x4_add([2]uint64{p1, p1h}, n6)
	n10 := Simd_i32x4_add(n9, [2]uint64{p4, p4h})
	n11 := Simd_i32x4_shr_u(n10, 16)
	n12 := Simd_i16x8_narrow_i32x4_u(n4, n11)
	n13 := Simd_f32x4_abs([2]uint64{p0, p0h})
	n14 := Simd_i32x4_gt_u(n13, [2]uint64{p5, p5h})
	n15 := Simd_f32x4_abs([2]uint64{p1, p1h})
	n16 := Simd_i32x4_gt_u(n15, [2]uint64{p5, p5h})
	n17 := Simd_i8x16_shuffle(n14, n16, [2]uint64{940136352262127872, 2097579117671354640})
	n18 := Simd_v128_bitselect(n8, n12, n17)
	_ = Simd_m64_v128_store(m, s0, 0, n18)
	return
}

//go:noinline
func Simd_p_fx202(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_gt(n0, [2]uint64{p0, p0h})
	n2 := Simd_v128_and(n1, n0)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx203(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_gt(n2, [2]uint64{p0, p0h})
	n4 := Simd_v128_and(n3, n2)
	n5 := Simd_i32x4_shr_u(n4, 16)
	n6 := Simd_v128_and(n5, [2]uint64{p2, p2h})
	n7 := Simd_i32x4_extend_high_i16x8_u(n0)
	n8 := Simd_i32x4_shl(n7, 16)
	n9 := Simd_f32x4_gt(n8, [2]uint64{p0, p0h})
	n10 := Simd_v128_and(n9, n8)
	n11 := Simd_i32x4_shr_u(n10, 16)
	n12 := Simd_v128_and(n11, [2]uint64{p2, p2h})
	n13 := Simd_i16x8_narrow_i32x4_u(n5, n11)
	n14 := Simd_v128_or(n13, [2]uint64{p1, p1h})
	n15 := Simd_v128_or(n4, n6)
	n16 := Simd_i32x4_add(n15, [2]uint64{p3, p3h})
	n17 := Simd_i32x4_shr_u(n16, 16)
	n18 := Simd_v128_or(n10, n12)
	n19 := Simd_i32x4_add(n18, [2]uint64{p3, p3h})
	n20 := Simd_i32x4_shr_u(n19, 16)
	n21 := Simd_i16x8_narrow_i32x4_u(n17, n20)
	n22 := Simd_i32x4_gt_u(n4, [2]uint64{p4, p4h})
	n23 := Simd_i32x4_gt_u(n10, [2]uint64{p4, p4h})
	n24 := Simd_i8x16_shuffle(n22, n23, [2]uint64{940136352262127872, 2097579117671354640})
	n25 := Simd_v128_bitselect(n14, n21, n24)
	_ = Simd_m64_v128_store(m, s1, 0, n25)
	return
}

//go:noinline
func Simd_p_fx204(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_gt(n1, [2]uint64{p0, p0h})
	n3 := Simd_v128_and(n2, n1)
	_ = Simd_m64_v128_store(m, s1, 0, n3)
	return
}

//go:noinline
func Simd_p_fx205(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_neg(n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx206(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_div([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	return
}

//go:noinline
func Simd_p_fx207(m *Module, s0 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_high_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_neg(n2)
	return n0[0], n0[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx208(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_neg(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx209(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_div([2]uint64{p0, p0h}, n0)
	n2 := Simd_i32x4_shr_u(n1, 16)
	n3 := Simd_v128_and(n2, [2]uint64{p4, p4h})
	n4 := Simd_i32x4_add(n1, n3)
	n5 := Simd_i32x4_add(n4, [2]uint64{p5, p5h})
	n6 := Simd_i32x4_shr_u(n5, 16)
	n7 := Simd_f32x4_abs(n1)
	n8 := Simd_i32x4_gt_u(n7, [2]uint64{p6, p6h})
	n9 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p2, p2h})
	n10 := Simd_f32x4_div([2]uint64{p0, p0h}, n9)
	n11 := Simd_i32x4_shr_u(n10, 16)
	n12 := Simd_v128_and(n11, [2]uint64{p4, p4h})
	n13 := Simd_i16x8_narrow_i32x4_u(n2, n11)
	n14 := Simd_v128_or(n13, [2]uint64{p3, p3h})
	n15 := Simd_i32x4_add(n10, n12)
	n16 := Simd_i32x4_add(n15, [2]uint64{p5, p5h})
	n17 := Simd_i32x4_shr_u(n16, 16)
	n18 := Simd_i16x8_narrow_i32x4_u(n6, n17)
	n19 := Simd_f32x4_abs(n10)
	n20 := Simd_i32x4_gt_u(n19, [2]uint64{p6, p6h})
	n21 := Simd_i8x16_shuffle(n8, n20, [2]uint64{940136352262127872, 2097579117671354640})
	n22 := Simd_v128_bitselect(n14, n18, n21)
	_ = Simd_m64_v128_store(m, s0, 0, n22)
	return
}

//go:noinline
func Simd_p_fx210(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_neg(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx211(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_sub([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_mul(n1, [2]uint64{p1, p1h})
	n3 := Simd_f32x4_add(n2, [2]uint64{p2, p2h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p3, p3h})
	n5 := Simd_f32x4_mul(n4, [2]uint64{p4, p4h})
	n6 := Simd_f32x4_add(n1, n5)
	n7 := Simd_f32x4_mul(n4, [2]uint64{p5, p5h})
	n8 := Simd_f32x4_add(n6, n7)
	return n0[0], n0[1], n3[0], n3[1], n4[0], n4[1], n8[0], n8[1]
}

//go:noinline
func Simd_p_fx212(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p0, p0h})
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p2, p2h})
	n3 := Simd_f32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p4, p4h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p5, p5h})
	n6 := Simd_f32x4_mul(n0, n5)
	n7 := Simd_f32x4_add(n3, n6)
	n8 := Simd_f32x4_mul(n0, n7)
	n9 := Simd_f32x4_add(n1, n8)
	n10 := Simd_i32x4_shl([2]uint64{p6, p6h}, 23)
	return n9[0], n9[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx213(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_abs([2]uint64{p2, p2h})
	n2 := Simd_f32x4_gt(n1, [2]uint64{p3, p3h})
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx214(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_le([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_i32x4_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_i32x4_sub([2]uint64{p4, p4h}, n1)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx215(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p0, p0h})
	n1 := Simd_f32x4_mul([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_add(n1, [2]uint64{p2, p2h})
	n3 := Simd_f32x4_mul(n2, [2]uint64{p0, p0h})
	n4 := Simd_f32x4_mul([2]uint64{p1, p1h}, [2]uint64{p3, p3h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p3, p3h})
	n6 := Simd_v128_bitselect(n3, n5, [2]uint64{p4, p4h})
	n7 := Simd_f32x4_gt([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n8 := Simd_v128_bitselect(n0, n6, n7)
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx216(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx217(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_f32x4_add([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_div([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	return
}

//go:noinline
func Simd_p_fx218(m *Module, s0 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_neg(n0)
	return n0[0], n0[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx219(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_add(n0, [2]uint64{p0, p0h})
	n2 := Simd_f32x4_div(n1, [2]uint64{p1, p1h})
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx220(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 0, n0)
	return
}

//go:noinline
func Simd_p_fx221(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_high_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_add(n2, [2]uint64{p0, p0h})
	n4 := Simd_f32x4_div(n3, [2]uint64{p1, p1h})
	return n0[0], n0[1], n2[0], n2[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx222(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_add(n1, [2]uint64{p1, p1h})
	n3 := Simd_f32x4_div(n2, [2]uint64{p2, p2h})
	return n1[0], n1[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx223(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_shr_u(n0, 16)
	n2 := Simd_f32x4_mul([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_i32x4_shr_u(n2, 16)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx224(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_narrow_i32x4_u([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx225(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	n2 := Simd_i32x4_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_i32x4_shr_u(n2, 16)
	n4 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p2, p2h})
	n5 := Simd_i32x4_add([2]uint64{p4, p4h}, n4)
	n6 := Simd_i32x4_add(n5, [2]uint64{p3, p3h})
	n7 := Simd_i32x4_shr_u(n6, 16)
	n8 := Simd_i16x8_narrow_i32x4_u(n3, n7)
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx226(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_abs([2]uint64{p0, p0h})
	n1 := Simd_i32x4_gt_u(n0, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_abs([2]uint64{p2, p2h})
	n3 := Simd_i32x4_gt_u(n2, [2]uint64{p1, p1h})
	n4 := Simd_i8x16_shuffle(n1, n3, [2]uint64{940136352262127872, 2097579117671354640})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx227(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_add(n1, [2]uint64{p0, p0h})
	n3 := Simd_f32x4_div(n2, [2]uint64{p1, p1h})
	return n1[0], n1[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx228(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_add(n0, [2]uint64{p0, p0h})
	n2 := Simd_f32x4_div(n1, [2]uint64{p1, p1h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx229(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_high_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_add(n2, [2]uint64{p0, p0h})
	n4 := Simd_f32x4_div(n3, [2]uint64{p1, p1h})
	return n0[0], n0[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx230(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_add(n1, [2]uint64{p1, p1h})
	n3 := Simd_f32x4_div(n2, [2]uint64{p2, p2h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx231(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 16)
	n1 := Simd_v128_and(n0, [2]uint64{p4, p4h})
	n2 := Simd_i16x8_narrow_i32x4_u([2]uint64{p1, p1h}, n0)
	n3 := Simd_v128_or(n2, [2]uint64{p2, p2h})
	n4 := Simd_i32x4_add([2]uint64{p0, p0h}, n1)
	n5 := Simd_i32x4_add(n4, [2]uint64{p5, p5h})
	n6 := Simd_i32x4_shr_u(n5, 16)
	n7 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p4, p4h})
	n8 := Simd_i32x4_add([2]uint64{p3, p3h}, n7)
	n9 := Simd_i32x4_add(n8, [2]uint64{p5, p5h})
	n10 := Simd_i32x4_shr_u(n9, 16)
	n11 := Simd_i16x8_narrow_i32x4_u(n10, n6)
	n12 := Simd_f32x4_abs([2]uint64{p3, p3h})
	n13 := Simd_i32x4_gt_u(n12, [2]uint64{p6, p6h})
	n14 := Simd_f32x4_abs([2]uint64{p0, p0h})
	n15 := Simd_i32x4_gt_u(n14, [2]uint64{p6, p6h})
	n16 := Simd_i8x16_shuffle(n13, n15, [2]uint64{940136352262127872, 2097579117671354640})
	n17 := Simd_v128_bitselect(n3, n11, n16)
	_ = Simd_m64_v128_store(m, s0, 0, n17)
	return
}

//go:noinline
func Simd_p_fx232(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_add(n1, [2]uint64{p0, p0h})
	n3 := Simd_f32x4_div(n2, [2]uint64{p1, p1h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx233(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_floor(n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx234(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_floor(n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_i32x4_extend_high_i16x8_u(n0)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_floor(n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i16x8_narrow_i32x4_u(n4, n9)
	n12 := Simd_v128_or(n11, [2]uint64{p0, p0h})
	n13 := Simd_i32x4_add(n3, n5)
	n14 := Simd_i32x4_add(n13, [2]uint64{p2, p2h})
	n15 := Simd_i32x4_shr_u(n14, 16)
	n16 := Simd_i32x4_add(n8, n10)
	n17 := Simd_i32x4_add(n16, [2]uint64{p2, p2h})
	n18 := Simd_i32x4_shr_u(n17, 16)
	n19 := Simd_i16x8_narrow_i32x4_u(n15, n18)
	n20 := Simd_f32x4_abs(n3)
	n21 := Simd_i32x4_gt_u(n20, [2]uint64{p3, p3h})
	n22 := Simd_f32x4_abs(n8)
	n23 := Simd_i32x4_gt_u(n22, [2]uint64{p3, p3h})
	n24 := Simd_i8x16_shuffle(n21, n23, [2]uint64{940136352262127872, 2097579117671354640})
	n25 := Simd_v128_bitselect(n12, n19, n24)
	_ = Simd_m64_v128_store(m, s1, 0, n25)
	return
}

//go:noinline
func Simd_p_fx235(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_floor(n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx236(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_ceil(n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx237(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_ceil(n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_i32x4_extend_high_i16x8_u(n0)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_ceil(n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i16x8_narrow_i32x4_u(n4, n9)
	n12 := Simd_v128_or(n11, [2]uint64{p0, p0h})
	n13 := Simd_i32x4_add(n3, n5)
	n14 := Simd_i32x4_add(n13, [2]uint64{p2, p2h})
	n15 := Simd_i32x4_shr_u(n14, 16)
	n16 := Simd_i32x4_add(n8, n10)
	n17 := Simd_i32x4_add(n16, [2]uint64{p2, p2h})
	n18 := Simd_i32x4_shr_u(n17, 16)
	n19 := Simd_i16x8_narrow_i32x4_u(n15, n18)
	n20 := Simd_f32x4_abs(n3)
	n21 := Simd_i32x4_gt_u(n20, [2]uint64{p3, p3h})
	n22 := Simd_f32x4_abs(n8)
	n23 := Simd_i32x4_gt_u(n22, [2]uint64{p3, p3h})
	n24 := Simd_i8x16_shuffle(n21, n23, [2]uint64{940136352262127872, 2097579117671354640})
	n25 := Simd_v128_bitselect(n12, n19, n24)
	_ = Simd_m64_v128_store(m, s1, 0, n25)
	return
}

//go:noinline
func Simd_p_fx238(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_ceil(n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx239(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_trunc(n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx240(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_trunc(n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_i32x4_extend_high_i16x8_u(n0)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_trunc(n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i16x8_narrow_i32x4_u(n4, n9)
	n12 := Simd_v128_or(n11, [2]uint64{p0, p0h})
	n13 := Simd_i32x4_add(n3, n5)
	n14 := Simd_i32x4_add(n13, [2]uint64{p2, p2h})
	n15 := Simd_i32x4_shr_u(n14, 16)
	n16 := Simd_i32x4_add(n8, n10)
	n17 := Simd_i32x4_add(n16, [2]uint64{p2, p2h})
	n18 := Simd_i32x4_shr_u(n17, 16)
	n19 := Simd_i16x8_narrow_i32x4_u(n15, n18)
	n20 := Simd_f32x4_abs(n3)
	n21 := Simd_i32x4_gt_u(n20, [2]uint64{p3, p3h})
	n22 := Simd_f32x4_abs(n8)
	n23 := Simd_i32x4_gt_u(n22, [2]uint64{p3, p3h})
	n24 := Simd_i8x16_shuffle(n21, n23, [2]uint64{940136352262127872, 2097579117671354640})
	n25 := Simd_v128_bitselect(n12, n19, n24)
	_ = Simd_m64_v128_store(m, s1, 0, n25)
	return
}

//go:noinline
func Simd_p_fx241(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_trunc(n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx242(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 0, n0)
	return
}

//go:noinline
func Simd_p_fx243(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_shr_u(n0, 16)
	n2 := Simd_v128_and(n1, [2]uint64{p4, p4h})
	n3 := Simd_i32x4_add(n0, n2)
	n4 := Simd_i32x4_add(n3, [2]uint64{p5, p5h})
	n5 := Simd_i32x4_shr_u(n4, 16)
	n6 := Simd_f32x4_abs(n0)
	n7 := Simd_i32x4_gt_u(n6, [2]uint64{p6, p6h})
	n8 := Simd_f32x4_add([2]uint64{p2, p2h}, [2]uint64{p1, p1h})
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p4, p4h})
	n11 := Simd_i16x8_narrow_i32x4_u(n1, n9)
	n12 := Simd_v128_or(n11, [2]uint64{p3, p3h})
	n13 := Simd_i32x4_add(n8, n10)
	n14 := Simd_i32x4_add(n13, [2]uint64{p5, p5h})
	n15 := Simd_i32x4_shr_u(n14, 16)
	n16 := Simd_i16x8_narrow_i32x4_u(n5, n15)
	n17 := Simd_f32x4_abs(n8)
	n18 := Simd_i32x4_gt_u(n17, [2]uint64{p6, p6h})
	n19 := Simd_i8x16_shuffle(n7, n18, [2]uint64{940136352262127872, 2097579117671354640})
	n20 := Simd_v128_bitselect(n12, n16, n19)
	_ = Simd_m64_v128_store(m, s0, 0, n20)
	return
}

//go:noinline
func Simd_p_fx244(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_gt([2]uint64{p0, p0h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n0)
	n2 := Simd_i32x4_shr_u(n1, 16)
	n3 := Simd_f32x4_gt([2]uint64{p3, p3h}, [2]uint64{p2, p2h})
	n4 := Simd_v128_bitselect([2]uint64{p3, p3h}, [2]uint64{p4, p4h}, n3)
	n5 := Simd_i32x4_shr_u(n4, 16)
	return n1[0], n1[1], n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx245(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_abs(n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx246(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 16)
	n3 := Simd_f32x4_abs(n2)
	n4 := Simd_i32x4_shr_u(n3, 16)
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_i32x4_extend_high_i16x8_u(n0)
	n7 := Simd_i32x4_shl(n6, 16)
	n8 := Simd_f32x4_abs(n7)
	n9 := Simd_i32x4_shr_u(n8, 16)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i16x8_narrow_i32x4_u(n4, n9)
	n12 := Simd_v128_or(n11, [2]uint64{p0, p0h})
	n13 := Simd_i32x4_add(n3, n5)
	n14 := Simd_i32x4_add(n13, [2]uint64{p2, p2h})
	n15 := Simd_i32x4_shr_u(n14, 16)
	n16 := Simd_i32x4_add(n8, n10)
	n17 := Simd_i32x4_add(n16, [2]uint64{p2, p2h})
	n18 := Simd_i32x4_shr_u(n17, 16)
	n19 := Simd_i16x8_narrow_i32x4_u(n15, n18)
	n20 := Simd_i32x4_gt_u(n3, [2]uint64{p3, p3h})
	n21 := Simd_i32x4_gt_u(n8, [2]uint64{p3, p3h})
	n22 := Simd_i8x16_shuffle(n20, n21, [2]uint64{940136352262127872, 2097579117671354640})
	n23 := Simd_v128_bitselect(n12, n19, n22)
	_ = Simd_m64_v128_store(m, s1, 0, n23)
	return
}

//go:noinline
func Simd_p_fx247(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_i32x4_shl(n0, 16)
	n2 := Simd_f32x4_abs(n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx248(m *Module, s0 int64, s1 int64, f0 float32, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_pmin(n1, [2]uint64{p0, p0h})
	n3 := Simd_f32x4_neg(n2)
	n4 := Simd_f32x4_splat(f0)
	n5 := Simd_f32x4_mul(n4, n3)
	return n0[0], n0[1], n2[0], n2[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx249(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_f32x4_pmin([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_lt([2]uint64{p1, p1h}, [2]uint64{p0, p0h})
	n2 := Simd_v128_bitselect([2]uint64{p0, p0h}, n0, n1)
	n3 := Simd_f32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_add([2]uint64{p5, p5h}, [2]uint64{p3, p3h})
	n5 := Simd_f32x4_div([2]uint64{p4, p4h}, n4)
	n6 := Simd_f32x4_mul(n3, n5)
	_ = Simd_m64_v128_store(m, s0, 0, n6)
	return
}

//go:noinline
func Simd_p_fx250(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, f0 float32, f1 float32, f2 float32, f3 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f1)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_m64_v128_load(m, s1, 0)
	n4 := Simd_f32x4_splat(f0)
	n5 := Simd_f32x4_mul(n4, n3)
	n6 := Simd_f32x4_splat(f2)
	n7 := Simd_f32x4_mul(n5, n6)
	n8 := Simd_f32x4_add(n2, n5)
	n9 := Simd_f32x4_add(n0, n7)
	n10 := Simd_f32x4_splat(f3)
	n11 := Simd_f32x4_mul(n9, n10)
	_ = Simd_m64_v128_store(m, s2, 0, n8)
	n13 := Simd_m64_v128_load(m, s3, 0)
	n14 := Simd_f32x4_add(n11, n13)
	_ = Simd_m64_v128_store(m, s3, 0, n14)
	return
}

//go:noinline
func Simd_p_fx251(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, f0 float32, f1 float32, f2 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_m64_v128_load(m, s1, 0)
	n4 := Simd_f32x4_splat(f1)
	n5 := Simd_f32x4_mul(n4, n3)
	n6 := Simd_f32x4_add(n2, n5)
	n7 := Simd_f32x4_splat(f2)
	n8 := Simd_f32x4_mul(n6, n7)
	_ = Simd_m64_v128_store(m, s2, 0, n6)
	n10 := Simd_m64_v128_load(m, s3, 0)
	n11 := Simd_f32x4_add(n8, n10)
	_ = Simd_m64_v128_store(m, s3, 0, n11)
	return
}

//go:noinline
func Simd_p_fx252(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0+48, 0, -48, 64)
	n1 := Simd_m64_v128_load_rng(m, s1+48, 0, -48, 64)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	n4 := Simd_m64_v128_load_nc(m, s0+32, 0)
	n5 := Simd_m64_v128_load_nc(m, s1+32, 0)
	n6 := Simd_f32x4_mul(n4, n5)
	n7 := Simd_f32x4_add([2]uint64{p1, p1h}, n6)
	n8 := Simd_m64_v128_load_nc(m, s0+16, 0)
	n9 := Simd_m64_v128_load_nc(m, s1+16, 0)
	n10 := Simd_f32x4_mul(n8, n9)
	n11 := Simd_f32x4_add([2]uint64{p2, p2h}, n10)
	n12 := Simd_m64_v128_load_nc(m, s0, 0)
	n13 := Simd_m64_v128_load_nc(m, s1, 0)
	n14 := Simd_f32x4_mul(n12, n13)
	n15 := Simd_f32x4_add([2]uint64{p3, p3h}, n14)
	return n3[0], n3[1], n7[0], n7[1], n11[0], n11[1], n15[0], n15[1]
}

//go:noinline
func Simd_p_fx253(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_scalar_i32_add(s1, s2)
	n2 := Simd_m64_v128_load(m, n1, 0)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_m64_scalar_i32_add(s1, s3)
	n5 := Simd_m64_v128_load(m, n4, 0)
	n6 := Simd_f32x4_mul([2]uint64{p1, p1h}, n5)
	n7 := Simd_m64_scalar_i32_add(s1, s4)
	n8 := Simd_m64_v128_load(m, n7, 0)
	n9 := Simd_m64_scalar_i32_add(s5, s1)
	n10 := Simd_m64_v128_load(m, n9, 0)
	n11 := Simd_f32x4_mul(n8, n10)
	n12 := Simd_f32x4_add(n6, n11)
	n13 := Simd_f32x4_add(n3, n12)
	n14 := Simd_m64_scalar_i32_add(s1, s6)
	_ = Simd_m64_v128_store(m, n14, 0, n13)
	n16 := Simd_m64_v128_load(m, s0+16, 0)
	n17 := Simd_m64_scalar_i32_add(s1, s2)
	n18 := Simd_m64_scalar_i32_add(n17, 16)
	n19 := Simd_m64_v128_load(m, n18, 0)
	n20 := Simd_f32x4_mul([2]uint64{p0, p0h}, n19)
	n21 := Simd_m64_scalar_i32_add(s1, s3)
	n22 := Simd_m64_scalar_i32_add(n21, 16)
	n23 := Simd_m64_v128_load(m, n22, 0)
	n24 := Simd_f32x4_mul([2]uint64{p1, p1h}, n23)
	n25 := Simd_m64_scalar_i32_add(s1, s4)
	n26 := Simd_m64_scalar_i32_add(n25, 16)
	n27 := Simd_m64_v128_load(m, n26, 0)
	n28 := Simd_m64_scalar_i32_add(s5, s1)
	n29 := Simd_m64_scalar_i32_add(n28, 16)
	n30 := Simd_m64_v128_load(m, n29, 0)
	n31 := Simd_f32x4_mul(n27, n30)
	n32 := Simd_f32x4_add(n24, n31)
	n33 := Simd_f32x4_add(n20, n32)
	n34 := Simd_m64_scalar_i32_add(s1, s6)
	n35 := Simd_m64_scalar_i32_add(n34, 16)
	_ = Simd_m64_v128_store(m, n35, 0, n33)
	return n0[0], n0[1], n13[0], n13[1], n16[0], n16[1], n33[0], n33[1]
}

//go:noinline
func Simd_p_fx254(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+32, 0)
	n1 := Simd_m64_v128_load(m, s1+32, 0)
	n2 := Simd_f32x4_mul([2]uint64{p0, p0h}, n1)
	n3 := Simd_m64_v128_load(m, s2+32, 0)
	n4 := Simd_f32x4_mul([2]uint64{p1, p1h}, n3)
	n5 := Simd_m64_v128_load(m, s3+32, 0)
	n6 := Simd_m64_v128_load(m, s4+32, 0)
	n7 := Simd_f32x4_mul(n5, n6)
	n8 := Simd_f32x4_add(n4, n7)
	n9 := Simd_f32x4_add(n2, n8)
	_ = Simd_m64_v128_store(m, s5+32, 0, n9)
	n11 := Simd_m64_v128_load(m, s0+48, 0)
	n12 := Simd_m64_v128_load(m, s1+48, 0)
	n13 := Simd_f32x4_mul([2]uint64{p0, p0h}, n12)
	n14 := Simd_m64_v128_load(m, s2+48, 0)
	n15 := Simd_f32x4_mul([2]uint64{p1, p1h}, n14)
	n16 := Simd_m64_v128_load(m, s3+48, 0)
	n17 := Simd_m64_v128_load(m, s4+48, 0)
	n18 := Simd_f32x4_mul(n16, n17)
	n19 := Simd_f32x4_add(n15, n18)
	n20 := Simd_f32x4_add(n13, n19)
	_ = Simd_m64_v128_store(m, s5+48, 0, n20)
	return n0[0], n0[1], n9[0], n9[1], n11[0], n11[1], n20[0], n20[1]
}

//go:noinline
func Simd_p_fx255(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64, s7 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 0)
	n1 := Simd_m64_v128_load32_lane(m, s1, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s2, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s3, 0, 3, n2)
	n4 := Simd_f32x4_mul(n3, [2]uint64{p0, p0h})
	n5 := Simd_m64_v128_load32_zero(m, s4, 0)
	n6 := Simd_m64_v128_load32_lane(m, s5, 0, 1, n5)
	n7 := Simd_m64_v128_load32_lane(m, s6, 0, 2, n6)
	n8 := Simd_m64_v128_load32_lane(m, s7, 0, 3, n7)
	n9 := Simd_f32x4_add(n4, n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx256(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_pmax([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n1 := Simd_f32x4_pmax([2]uint64{p1, p1h}, n0)
	n2 := Simd_f32x4_pmax([2]uint64{p0, p0h}, n1)
	n3 := Simd_f32x4_sub([2]uint64{p2, p2h}, n2)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx257(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_add([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_add(n0, [2]uint64{p3, p3h})
	n2 := Simd_f32x4_add(n1, [2]uint64{p4, p4h})
	n3 := Simd_f32x4_div([2]uint64{p0, p0h}, n2)
	n4 := Simd_f32x4_mul([2]uint64{p1, p1h}, n3)
	n5 := Simd_f32x4_add(n4, [2]uint64{p5, p5h})
	n6 := Simd_f32x4_mul([2]uint64{p2, p2h}, n3)
	n7 := Simd_f32x4_add(n6, [2]uint64{p5, p5h})
	n8 := Simd_f32x4_mul([2]uint64{p3, p3h}, n3)
	n9 := Simd_f32x4_add(n8, [2]uint64{p5, p5h})
	n10 := Simd_f32x4_mul([2]uint64{p4, p4h}, n3)
	n11 := Simd_f32x4_add(n10, [2]uint64{p5, p5h})
	n12 := Simd_i8x16_shuffle(n5, n7, [2]uint64{2242261670825954572, 216736831629295872})
	n13 := Simd_i8x16_shuffle(n9, n11, [2]uint64{216736831629295872, 2242261670825954572})
	n14 := Simd_i8x16_shuffle(n12, n13, [2]uint64{p6, p6h})
	n15 := Simd_i8x16_shuffle(n5, n7, [2]uint64{1374179596769034496, 216736831629295872})
	n16 := Simd_i8x16_shuffle(n9, n11, [2]uint64{216736831629295872, 1374179596769034496})
	n17 := Simd_i8x16_shuffle(n15, n16, [2]uint64{p6, p6h})
	n18 := Simd_i8x16_shuffle(n5, n7, [2]uint64{1663540288121341188, 216736831629295872})
	n19 := Simd_i8x16_shuffle(n5, n7, [2]uint64{1952900979473647880, 216736831629295872})
	n20 := Simd_i8x16_shuffle(n9, n11, [2]uint64{216736831629295872, 1663540288121341188})
	n21 := Simd_i8x16_shuffle(n18, n20, [2]uint64{p6, p6h})
	n22 := Simd_i8x16_shuffle(n9, n11, [2]uint64{216736831629295872, 1952900979473647880})
	n23 := Simd_i8x16_shuffle(n19, n22, [2]uint64{p6, p6h})
	n24 := Simd_f32x4_add([2]uint64{p5, p5h}, n17)
	n25 := Simd_f32x4_add(n24, n21)
	n26 := Simd_f32x4_add(n25, n23)
	n27 := Simd_f32x4_add(n14, n26)
	n28 := Simd_f32x4_div([2]uint64{p0, p0h}, n27)
	n29 := Simd_f32x4_mul(n14, n28)
	n30 := Simd_f32x4_mul(n23, n28)
	n31 := Simd_f32x4_mul(n21, n28)
	n32 := Simd_f32x4_mul(n17, n28)
	return n29[0], n29[1], n30[0], n30[1], n31[0], n31[1], n32[0], n32[1]
}

//go:noinline
func Simd_p_fx258(m *Module, f0 float32, f1 float32, f2 float32, f3 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_splat(f1)
	n3 := Simd_f32x4_mul([2]uint64{p1, p1h}, n2)
	n4 := Simd_f32x4_splat(f2)
	n5 := Simd_f32x4_mul([2]uint64{p2, p2h}, n4)
	n6 := Simd_f32x4_splat(f3)
	n7 := Simd_f32x4_mul([2]uint64{p3, p3h}, n6)
	n8 := Simd_f32x4_add([2]uint64{p5, p5h}, n3)
	n9 := Simd_f32x4_add(n8, n5)
	n10 := Simd_f32x4_add(n9, n7)
	n11 := Simd_f32x4_add(n1, n10)
	n12 := Simd_f32x4_div([2]uint64{p4, p4h}, n11)
	n13 := Simd_f32x4_mul(n1, n12)
	n14 := Simd_f32x4_mul(n7, n12)
	n15 := Simd_f32x4_mul(n5, n12)
	n16 := Simd_f32x4_mul(n3, n12)
	return n13[0], n13[1], n14[0], n14[1], n15[0], n15[1], n16[0], n16[1]
}

//go:noinline
func Simd_p_fx259(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_f32x4_sub(n0, n1)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx260(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, f0 float32, f1 float32, f2 float32, f3 float32) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_m64_scalar_i32_add(s1, s2)
	n4 := Simd_m64_v128_load(m, n3, 0)
	n5 := Simd_f32x4_mul(n4, n4)
	n6 := Simd_f32x4_splat(f1)
	n7 := Simd_f32x4_mul(n6, n4)
	n8 := Simd_f32x4_add(n2, n7)
	n9 := Simd_f32x4_splat(f3)
	n10 := Simd_f32x4_mul(n9, n5)
	_ = Simd_m64_v128_store(m, s0, 0, n8)
	n12 := Simd_m64_scalar_i32_add(s3, s2)
	n13 := Simd_m64_v128_load(m, n12, 0)
	n14 := Simd_f32x4_splat(f2)
	n15 := Simd_f32x4_mul(n13, n14)
	n16 := Simd_f32x4_add(n15, n10)
	n17 := Simd_m64_scalar_i32_add(s3, s2)
	_ = Simd_m64_v128_store(m, n17, 0, n16)
	n19 := Simd_m64_scalar_i32_add(s4, s2)
	n20 := Simd_m64_v128_load(m, n19, 0)
	return n8[0], n8[1], n16[0], n16[1], n20[0], n20[1]
}

//go:noinline
func Simd_p_fx261(m *Module, f0 float32, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx262(m *Module, f0 float32, f1 float32, f2 float32, f3 float32, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_splat(f1)
	n2 := Simd_f32x4_mul([2]uint64{p0, p0h}, n1)
	n3 := Simd_f32x4_mul(n0, n2)
	n4 := Simd_f32x4_splat(f2)
	n5 := Simd_f32x4_splat(f3)
	n6 := Simd_f32x4_mul(n5, [2]uint64{p1, p1h})
	n7 := Simd_f32x4_sqrt(n6)
	n8 := Simd_f32x4_add(n4, n7)
	n9 := Simd_f32x4_div(n3, n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx263(m *Module, s0 int64, s1 int64, f0 float32, f1 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_m64_v128_load(m, s1, 0)
	n4 := Simd_f32x4_splat(f1)
	n5 := Simd_f32x4_mul(n4, n3)
	n6 := Simd_f32x4_sub(n2, n5)
	_ = Simd_m64_v128_store(m, s0, 0, n6)
	return
}

//go:noinline
func Simd_p_fx264(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_div([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_sub([2]uint64{p0, p0h}, n1)
	n3 := Simd_f32x4_mul([2]uint64{p2, p2h}, n2)
	n4 := Simd_f32x4_add(n3, [2]uint64{p0, p0h})
	n5 := Simd_m64_v128_load(m, s0, 0)
	n6 := Simd_f32x4_mul(n5, n1)
	n7 := Simd_f32x4_mul(n6, n4)
	_ = Simd_m64_v128_store(m, s1, 0, n7)
	return
}

//go:noinline
func Simd_p_fx265(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_mul(n1, [2]uint64{p1, p1h})
	n3 := Simd_i32x4_shr_u(n2, 8)
	n4 := Simd_i32x4_add(n3, [2]uint64{p2, p2h})
	n5 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n6 := Simd_i32x4_extend_low_i16x8_s(n5)
	n7 := Simd_i32x4_mul(n4, n6)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx266(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_mul(n1, [2]uint64{p1, p1h})
	n3 := Simd_i32x4_shr_u(n2, 8)
	n4 := Simd_i32x4_add(n3, [2]uint64{p2, p2h})
	n5 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n6 := Simd_i32x4_extend_low_i16x8_s(n5)
	n7 := Simd_i32x4_mul(n4, n6)
	n8 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	n10 := Simd_i32x4_mul(n9, [2]uint64{p1, p1h})
	n11 := Simd_i32x4_shr_u(n10, 8)
	n12 := Simd_i32x4_add(n11, [2]uint64{p2, p2h})
	n13 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p5, p5h})
	n14 := Simd_i32x4_extend_low_i16x8_s(n13)
	n15 := Simd_i32x4_mul(n12, n14)
	n16 := Simd_i32x4_add(n7, n15)
	return n16[0], n16[1]
}

//go:noinline
func Simd_p_fx267(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx268(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extmul_low_i8x16_u([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_i16x8_extend_low_i8x16_u(n1)
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_mul(n3, [2]uint64{p2, p2h})
	n5 := Simd_i32x4_shr_u(n4, 8)
	n6 := Simd_i32x4_add(n5, [2]uint64{p4, p4h})
	n7 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p5, p5h})
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_i32x4_mul(n6, n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx269(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	n2 := Simd_i8x16_shuffle(n1, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx270(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx271(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	n2 := Simd_i8x16_shuffle(n1, [2]uint64{p1, p1h}, [2]uint64{p3, p3h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx272(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extmul_low_i8x16_u([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_i16x8_extend_low_i8x16_u(n1)
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_mul(n3, [2]uint64{p4, p4h})
	n5 := Simd_i32x4_shr_u(n4, 8)
	n6 := Simd_i32x4_add(n5, [2]uint64{p5, p5h})
	n7 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p6, p6h})
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_i32x4_mul(n6, n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx273(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extmul_low_i8x16_u([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_i16x8_extend_low_i8x16_u(n1)
	n3 := Simd_i16x8_mul(n2, [2]uint64{p4, p4h})
	n4 := Simd_i16x8_shr_u(n3, 8)
	n5 := Simd_i16x8_add(n4, [2]uint64{p5, p5h})
	n6 := Simd_i16x8_extmul_high_i8x16_u([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n7 := Simd_i8x16_shuffle(n6, [2]uint64{p4, p4h}, [2]uint64{p3, p3h})
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i16x8_mul(n8, [2]uint64{p4, p4h})
	n10 := Simd_i16x8_shr_u(n9, 8)
	n11 := Simd_i16x8_add(n10, [2]uint64{p5, p5h})
	n12 := Simd_m64_v128_load(m, s0, 0)
	n13 := Simd_i16x8_extend_low_i8x16_s(n12)
	n14 := Simd_i16x8_mul(n5, n13)
	n15 := Simd_i32x4_extend_low_i16x8_s(n14)
	n16 := Simd_i16x8_extend_high_i8x16_s(n12)
	n17 := Simd_i16x8_mul(n11, n16)
	n18 := Simd_i32x4_extend_low_i16x8_s(n17)
	n19 := Simd_i32x4_add(n15, n18)
	n20 := Simd_i32x4_extend_high_i16x8_s(n14)
	n21 := Simd_i32x4_extend_high_i16x8_s(n17)
	n22 := Simd_i32x4_add(n19, n20)
	n23 := Simd_i32x4_add(n22, n21)
	n24 := Simd_i8x16_shuffle(n23, [2]uint64{p4, p4h}, [2]uint64{1084818905618843912, 216736831629295872})
	n25 := Simd_i32x4_add(n23, n24)
	return n25[0], n25[1]
}

//go:noinline
func Simd_p_fx274(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{216736831696667908, 216736831629295872})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx275(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i16x8_mul(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_shr_u(n1, 8)
	n3 := Simd_i16x8_add(n2, [2]uint64{p2, p2h})
	n4 := Simd_i16x8_mul(n3, [2]uint64{p3, p3h})
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_i32x4_extend_high_i16x8_s(n4)
	n7 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n8 := Simd_i16x8_mul(n7, [2]uint64{p1, p1h})
	n9 := Simd_i16x8_shr_u(n8, 8)
	n10 := Simd_i16x8_add(n9, [2]uint64{p2, p2h})
	n11 := Simd_i8x16_shuffle(n10, [2]uint64{p1, p1h}, [2]uint64{361417177238079238, 72058693566333184})
	n12 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p5, p5h})
	n13 := Simd_i16x8_mul(n11, n12)
	n14 := Simd_i32x4_extend_low_i16x8_s(n13)
	n15 := Simd_i32x4_add(n14, n5)
	n16 := Simd_i32x4_add(n15, n6)
	n17 := Simd_i8x16_shuffle(n16, n16, [2]uint64{1084818905618843912, 216736831629295872})
	n18 := Simd_i32x4_add(n16, n17)
	return n18[0], n18[1]
}

//go:noinline
func Simd_p_fx276(m *Module, s0 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n1 := Simd_m64_v128_load_nc(m, s0, 16)
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{1952894356634075906, 72058693566333184})
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx277(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i64x2_extend_high_i32x4_u(n0)
	return n0[0], n0[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx278(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_mul(n1, n3)
	n5 := Simd_i32x4_shr_u([2]uint64{p2, p2h}, 9)
	n6 := Simd_i64x2_extend_high_i32x4_u(n5)
	n7 := Simd_i64x2_extend_low_i32x4_u(n5)
	return n4[0], n4[1], n6[0], n6[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx279(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_mul(n1, n3)
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx280(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{1808214010957922560, 72058693566333184})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_v128_and(n1, [2]uint64{p2, p2h})
	n3 := Simd_i64x2_extend_high_i32x4_u(n2)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx281(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{2242255047986382598, 72058693566333184})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_v128_and(n1, [2]uint64{p2, p2h})
	n3 := Simd_i64x2_extend_high_i32x4_u(n2)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx282(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{2097574702310229252, 72058693566333184})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_v128_and(n1, [2]uint64{p2, p2h})
	n3 := Simd_i64x2_extend_high_i32x4_u(n2)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx283(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_neg([2]uint64{p0, p0h})
	n1 := Simd_i8x16_lt_s([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_v128_bitselect(n0, [2]uint64{p0, p0h}, n3)
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx284(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_neg([2]uint64{p0, p0h})
	n1 := Simd_i16x8_lt_s([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_v128_bitselect(n0, [2]uint64{p0, p0h}, n2)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx285(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_neg([2]uint64{p0, p0h})
	n1 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p3, p3h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_v128_bitselect([2]uint64{p0, p0h}, n0, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx286(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_neg([2]uint64{p0, p0h})
	n1 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p3, p3h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_v128_bitselect([2]uint64{p0, p0h}, n0, n4)
	n6 := Simd_i32x4_neg([2]uint64{p4, p4h})
	n7 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p5, p5h})
	n8 := Simd_i8x16_eq(n7, [2]uint64{p3, p3h})
	n9 := Simd_i16x8_extend_low_i8x16_s(n8)
	n10 := Simd_i32x4_extend_low_i16x8_s(n9)
	n11 := Simd_v128_bitselect([2]uint64{p4, p4h}, n6, n10)
	n12 := Simd_i32x4_add(n5, n11)
	return n12[0], n12[1]
}

//go:noinline
func Simd_p_fx287(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 1)
	n3 := Simd_v128_and(n2, [2]uint64{p1, p1h})
	n4 := Simd_v128_or(n3, [2]uint64{p2, p2h})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx288(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx289(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{1084818905618843912, 216736831629295872})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx290(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 0)
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_shl(n2, 6)
	n4 := Simd_v128_and(n3, [2]uint64{p0, p0h})
	n5 := Simd_m64_v128_load(m, s1, 0)
	n6 := Simd_i8x16_shuffle(n5, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n7 := Simd_i16x8_extend_low_i8x16_u(n6)
	n8 := Simd_i32x4_extend_low_i16x8_u(n7)
	n9 := Simd_v128_or(n4, n8)
	n10 := Simd_i64x2_extend_high_i32x4_u(n9)
	return n2[0], n2[1], n5[0], n5[1], n9[0], n9[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx291(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_i32x4_mul(n1, n3)
	n5 := Simd_m64_v128_load(m, s0, 0)
	n6 := Simd_i8x16_shuffle(n5, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	return n4[0], n4[1], n5[0], n5[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx292(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_i32x4_mul(n1, n3)
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx293(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_i32x4_mul(n1, n3)
	n5 := Simd_i32x4_shl([2]uint64{p2, p2h}, 8)
	n6 := Simd_v128_and(n5, [2]uint64{p3, p3h})
	n7 := Simd_i8x16_shuffle([2]uint64{p4, p4h}, [2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	n10 := Simd_v128_or(n6, n9)
	n11 := Simd_i64x2_extend_high_i32x4_u(n10)
	return n4[0], n4[1], n10[0], n10[1], n11[0], n11[1]
}

//go:noinline
func Simd_p_fx294(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_i32x4_mul(n1, n3)
	n5 := Simd_i32x4_shl([2]uint64{p2, p2h}, 2)
	n6 := Simd_v128_and(n5, [2]uint64{p3, p3h})
	n7 := Simd_i8x16_shuffle([2]uint64{p4, p4h}, [2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	n10 := Simd_v128_or(n6, n9)
	n11 := Simd_i64x2_extend_high_i32x4_u(n10)
	n12 := Simd_m64_v128_load32_zero(m, s0+8, 0)
	n13 := Simd_i16x8_extend_low_i8x16_u(n12)
	n14 := Simd_i32x4_extend_low_i16x8_u(n13)
	return n4[0], n4[1], n14[0], n14[1], n10[0], n10[1], n11[0], n11[1]
}

//go:noinline
func Simd_p_fx295(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_i32x4_mul(n1, n3)
	n5 := Simd_i32x4_shl([2]uint64{p2, p2h}, 4)
	n6 := Simd_v128_and(n5, [2]uint64{p3, p3h})
	n7 := Simd_i8x16_shuffle([2]uint64{p4, p4h}, [2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	n10 := Simd_v128_or(n6, n9)
	n11 := Simd_i64x2_extend_high_i32x4_u(n10)
	return n4[0], n4[1], n10[0], n10[1], n11[0], n11[1]
}

//go:noinline
func Simd_p_fx296(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p0, p0h}, 1)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_v128_or(n1, [2]uint64{p2, p2h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx297(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 3)
	n1 := Simd_v128_or(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx298(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 128)
	n1 := Simd_f32x4_abs(n0)
	n2 := Simd_m64_v128_load_nc(m, s1+528, 0)
	n3 := Simd_f32x4_abs(n2)
	n4 := Simd_f32x4_max(n1, n3)
	n5 := Simd_m64_v128_load_nc(m, s1+544, 0)
	n6 := Simd_f32x4_abs(n5)
	n7 := Simd_m64_v128_load_nc(m, s1+560, 0)
	n8 := Simd_f32x4_abs(n7)
	n9 := Simd_f32x4_max(n6, n8)
	n10 := Simd_f32x4_max(n4, n9)
	n11 := Simd_m64_v128_load_nc(m, s1+576, 0)
	n12 := Simd_f32x4_abs(n11)
	n13 := Simd_m64_v128_load_nc(m, s1+592, 0)
	n14 := Simd_f32x4_abs(n13)
	n15 := Simd_f32x4_max(n12, n14)
	n16 := Simd_m64_v128_load_nc(m, s1+608, 0)
	n17 := Simd_f32x4_abs(n16)
	n18 := Simd_m64_v128_load_nc(m, s1+624, 0)
	n19 := Simd_f32x4_abs(n18)
	n20 := Simd_f32x4_max(n17, n19)
	n21 := Simd_f32x4_max(n15, n20)
	n22 := Simd_f32x4_max(n10, n21)
	return n22[0], n22[1]
}

//go:noinline
func Simd_p_fx299(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 416)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{216736831629295872, 216736831629295872})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{506097522981602564, 506097522981602564})
	n3 := Simd_i8x16_shuffle(n0, n1, [2]uint64{795458214333909256, 795458214333909256})
	n4 := Simd_i8x16_shuffle(n0, n1, [2]uint64{1084818905686215948, 1084818905686215948})
	n5 := Simd_m64_v128_load_nc(m, s0, 16)
	n6 := Simd_f32x4_mul(n1, n5)
	n7 := Simd_f32x4_nearest(n6)
	n8 := Simd_i32x4_trunc_sat_f32x4_s(n7)
	n9 := Simd_m64_v128_load_nc(m, s0, 144)
	n10 := Simd_f32x4_mul(n2, n9)
	n11 := Simd_f32x4_nearest(n10)
	n12 := Simd_i32x4_trunc_sat_f32x4_s(n11)
	n13 := Simd_i16x8_narrow_i32x4_s(n8, n12)
	n14 := Simd_m64_v128_load_nc(m, s0, 272)
	n15 := Simd_f32x4_mul(n3, n14)
	n16 := Simd_f32x4_nearest(n15)
	n17 := Simd_i32x4_trunc_sat_f32x4_s(n16)
	n18 := Simd_m64_v128_load_nc(m, s0, 400)
	n19 := Simd_f32x4_mul(n4, n18)
	n20 := Simd_f32x4_nearest(n19)
	n21 := Simd_i32x4_trunc_sat_f32x4_s(n20)
	n22 := Simd_i16x8_narrow_i32x4_s(n17, n21)
	n23 := Simd_i8x16_narrow_i16x8_s(n13, n22)
	_ = Simd_m64_v128_store(m, s1, 8, n23)
	n25 := Simd_m64_v128_load_rng(m, s0, 32, 32, 400)
	n26 := Simd_f32x4_mul(n1, n25)
	n27 := Simd_f32x4_nearest(n26)
	n28 := Simd_i32x4_trunc_sat_f32x4_s(n27)
	n29 := Simd_m64_v128_load_nc(m, s0, 160)
	n30 := Simd_f32x4_mul(n2, n29)
	n31 := Simd_f32x4_nearest(n30)
	n32 := Simd_i32x4_trunc_sat_f32x4_s(n31)
	n33 := Simd_i16x8_narrow_i32x4_s(n28, n32)
	n34 := Simd_m64_v128_load_nc(m, s0, 288)
	n35 := Simd_f32x4_mul(n3, n34)
	n36 := Simd_f32x4_nearest(n35)
	n37 := Simd_i32x4_trunc_sat_f32x4_s(n36)
	n38 := Simd_m64_v128_load_nc(m, s0, 416)
	n39 := Simd_f32x4_mul(n4, n38)
	n40 := Simd_f32x4_nearest(n39)
	n41 := Simd_i32x4_trunc_sat_f32x4_s(n40)
	n42 := Simd_i16x8_narrow_i32x4_s(n37, n41)
	n43 := Simd_i8x16_narrow_i16x8_s(n33, n42)
	_ = Simd_m64_v128_store(m, s1, 24, n43)
	n45 := Simd_m64_v128_load_rng(m, s0, 48, 48, 400)
	n46 := Simd_f32x4_mul(n1, n45)
	n47 := Simd_f32x4_nearest(n46)
	n48 := Simd_i32x4_trunc_sat_f32x4_s(n47)
	n49 := Simd_m64_v128_load_nc(m, s0, 176)
	n50 := Simd_f32x4_mul(n2, n49)
	n51 := Simd_f32x4_nearest(n50)
	n52 := Simd_i32x4_trunc_sat_f32x4_s(n51)
	n53 := Simd_i16x8_narrow_i32x4_s(n48, n52)
	n54 := Simd_m64_v128_load_nc(m, s0, 304)
	n55 := Simd_f32x4_mul(n3, n54)
	n56 := Simd_f32x4_nearest(n55)
	n57 := Simd_i32x4_trunc_sat_f32x4_s(n56)
	n58 := Simd_m64_v128_load_nc(m, s0, 432)
	n59 := Simd_f32x4_mul(n4, n58)
	n60 := Simd_f32x4_nearest(n59)
	n61 := Simd_i32x4_trunc_sat_f32x4_s(n60)
	n62 := Simd_i16x8_narrow_i32x4_s(n57, n61)
	n63 := Simd_i8x16_narrow_i16x8_s(n53, n62)
	_ = Simd_m64_v128_store(m, s1, 40, n63)
	n65 := Simd_m64_v128_load_rng(m, s0, 64, 64, 400)
	n66 := Simd_f32x4_mul(n1, n65)
	n67 := Simd_f32x4_nearest(n66)
	n68 := Simd_i32x4_trunc_sat_f32x4_s(n67)
	n69 := Simd_m64_v128_load_nc(m, s0, 192)
	n70 := Simd_f32x4_mul(n2, n69)
	n71 := Simd_f32x4_nearest(n70)
	n72 := Simd_i32x4_trunc_sat_f32x4_s(n71)
	n73 := Simd_i16x8_narrow_i32x4_s(n68, n72)
	n74 := Simd_m64_v128_load_nc(m, s0, 320)
	n75 := Simd_f32x4_mul(n3, n74)
	n76 := Simd_f32x4_nearest(n75)
	n77 := Simd_i32x4_trunc_sat_f32x4_s(n76)
	n78 := Simd_m64_v128_load_nc(m, s0, 448)
	n79 := Simd_f32x4_mul(n4, n78)
	n80 := Simd_f32x4_nearest(n79)
	n81 := Simd_i32x4_trunc_sat_f32x4_s(n80)
	n82 := Simd_i16x8_narrow_i32x4_s(n77, n81)
	n83 := Simd_i8x16_narrow_i16x8_s(n73, n82)
	_ = Simd_m64_v128_store(m, s1, 56, n83)
	n85 := Simd_m64_v128_load_rng(m, s0, 80, 80, 400)
	n86 := Simd_f32x4_mul(n1, n85)
	n87 := Simd_f32x4_nearest(n86)
	n88 := Simd_i32x4_trunc_sat_f32x4_s(n87)
	n89 := Simd_m64_v128_load_nc(m, s0, 208)
	n90 := Simd_f32x4_mul(n2, n89)
	n91 := Simd_f32x4_nearest(n90)
	n92 := Simd_i32x4_trunc_sat_f32x4_s(n91)
	n93 := Simd_i16x8_narrow_i32x4_s(n88, n92)
	n94 := Simd_m64_v128_load_nc(m, s0, 336)
	n95 := Simd_f32x4_mul(n3, n94)
	n96 := Simd_f32x4_nearest(n95)
	n97 := Simd_i32x4_trunc_sat_f32x4_s(n96)
	n98 := Simd_m64_v128_load_nc(m, s0, 464)
	n99 := Simd_f32x4_mul(n4, n98)
	n100 := Simd_f32x4_nearest(n99)
	n101 := Simd_i32x4_trunc_sat_f32x4_s(n100)
	n102 := Simd_i16x8_narrow_i32x4_s(n97, n101)
	n103 := Simd_i8x16_narrow_i16x8_s(n93, n102)
	_ = Simd_m64_v128_store(m, s1, 72, n103)
	n105 := Simd_m64_v128_load_rng(m, s0, 96, 96, 400)
	n106 := Simd_f32x4_mul(n1, n105)
	n107 := Simd_f32x4_nearest(n106)
	n108 := Simd_i32x4_trunc_sat_f32x4_s(n107)
	n109 := Simd_m64_v128_load_nc(m, s0, 224)
	n110 := Simd_f32x4_mul(n2, n109)
	n111 := Simd_f32x4_nearest(n110)
	n112 := Simd_i32x4_trunc_sat_f32x4_s(n111)
	n113 := Simd_i16x8_narrow_i32x4_s(n108, n112)
	n114 := Simd_m64_v128_load_nc(m, s0, 352)
	n115 := Simd_f32x4_mul(n3, n114)
	n116 := Simd_f32x4_nearest(n115)
	n117 := Simd_i32x4_trunc_sat_f32x4_s(n116)
	n118 := Simd_m64_v128_load_nc(m, s0, 480)
	n119 := Simd_f32x4_mul(n4, n118)
	n120 := Simd_f32x4_nearest(n119)
	n121 := Simd_i32x4_trunc_sat_f32x4_s(n120)
	n122 := Simd_i16x8_narrow_i32x4_s(n117, n121)
	n123 := Simd_i8x16_narrow_i16x8_s(n113, n122)
	_ = Simd_m64_v128_store(m, s1, 88, n123)
	n125 := Simd_m64_v128_load_rng(m, s0, 112, 112, 400)
	n126 := Simd_f32x4_mul(n1, n125)
	n127 := Simd_f32x4_nearest(n126)
	n128 := Simd_i32x4_trunc_sat_f32x4_s(n127)
	n129 := Simd_m64_v128_load_nc(m, s0, 240)
	n130 := Simd_f32x4_mul(n2, n129)
	n131 := Simd_f32x4_nearest(n130)
	n132 := Simd_i32x4_trunc_sat_f32x4_s(n131)
	n133 := Simd_i16x8_narrow_i32x4_s(n128, n132)
	n134 := Simd_m64_v128_load_nc(m, s0, 368)
	n135 := Simd_f32x4_mul(n3, n134)
	n136 := Simd_f32x4_nearest(n135)
	n137 := Simd_i32x4_trunc_sat_f32x4_s(n136)
	n138 := Simd_m64_v128_load_nc(m, s0, 496)
	n139 := Simd_f32x4_mul(n4, n138)
	n140 := Simd_f32x4_nearest(n139)
	n141 := Simd_i32x4_trunc_sat_f32x4_s(n140)
	n142 := Simd_i16x8_narrow_i32x4_s(n137, n141)
	n143 := Simd_i8x16_narrow_i16x8_s(n133, n142)
	_ = Simd_m64_v128_store(m, s1, 104, n143)
	n145 := Simd_m64_v128_load_rng(m, s0, 128, 128, 400)
	n146 := Simd_f32x4_mul(n1, n145)
	n147 := Simd_f32x4_nearest(n146)
	n148 := Simd_i32x4_trunc_sat_f32x4_s(n147)
	n149 := Simd_m64_v128_load_nc(m, s0, 256)
	n150 := Simd_f32x4_mul(n2, n149)
	n151 := Simd_f32x4_nearest(n150)
	n152 := Simd_i32x4_trunc_sat_f32x4_s(n151)
	n153 := Simd_i16x8_narrow_i32x4_s(n148, n152)
	n154 := Simd_m64_v128_load_nc(m, s0, 384)
	n155 := Simd_f32x4_mul(n3, n154)
	n156 := Simd_f32x4_nearest(n155)
	n157 := Simd_i32x4_trunc_sat_f32x4_s(n156)
	n158 := Simd_m64_v128_load_nc(m, s0, 512)
	n159 := Simd_f32x4_mul(n4, n158)
	n160 := Simd_f32x4_nearest(n159)
	n161 := Simd_i32x4_trunc_sat_f32x4_s(n160)
	n162 := Simd_i16x8_narrow_i32x4_s(n157, n161)
	n163 := Simd_i8x16_narrow_i16x8_s(n153, n162)
	_ = Simd_m64_v128_store(m, s1, 120, n163)
	return
}

//go:noinline
func Simd_p_fx300(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 8793760)
	n1 := Simd_m64_v128_load32_lane(m, s1, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s2, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s3, 0, 3, n2)
	n4 := Simd_m64_v128_load32_splat(m, s4, 8793760)
	n5 := Simd_m64_v128_load32_splat(m, s5, 8793760)
	n6 := Simd_m64_v128_load32_splat(m, s6, 8793760)
	return n3[0], n3[1], n4[0], n4[1], n5[0], n5[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx301(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0+136, 0, -64, 80)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_m64_v128_load_nc(m, s0+72, 0)
	n5 := Simd_i8x16_shuffle(n4, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n6 := Simd_i16x8_extend_low_i8x16_s(n5)
	n7 := Simd_i32x4_extend_low_i16x8_s(n6)
	return n0[0], n0[1], n3[0], n3[1], n4[0], n4[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx302(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx303(m *Module, s0 int32, s1 int32, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_mul(n0, [2]uint64{p0, p0h})
	n2 := Simd_i32x4_splat(s1)
	n3 := Simd_i32x4_mul(n2, [2]uint64{p1, p1h})
	n4 := Simd_i32x4_add(n1, n3)
	n5 := Simd_i32x4_shr_s(n4, 4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx304(m *Module, s0 int32, s1 int32, s2 int32, s3 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_mul(n0, [2]uint64{p0, p0h})
	n2 := Simd_i32x4_splat(s1)
	n3 := Simd_i32x4_mul(n2, [2]uint64{p1, p1h})
	n4 := Simd_i32x4_add(n1, n3)
	n5 := Simd_i32x4_shr_s(n4, 4)
	n6 := Simd_i32x4_splat(s2)
	n7 := Simd_i32x4_mul(n6, [2]uint64{p2, p2h})
	n8 := Simd_i32x4_splat(s3)
	n9 := Simd_i32x4_mul(n8, [2]uint64{p3, p3h})
	n10 := Simd_i32x4_add(n7, n9)
	n11 := Simd_i32x4_shr_s(n10, 4)
	n12 := Simd_i32x4_add(n5, n11)
	return n12[0], n12[1]
}

//go:noinline
func Simd_p_fx305(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i8x16_shuffle([2]uint64{p3, p3h}, [2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p6, p6h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx306(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 8793760)
	n1 := Simd_m64_v128_load32_lane(m, s1, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s2, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s3, 0, 3, n2)
	n4 := Simd_m64_v128_load32_splat(m, s4, 8793760)
	return n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx307(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{795458214199165184, 1084818905551471876})
	n1 := Simd_m64_scalar_i32_shl(s1, 4)
	n2 := Simd_m64_scalar_i32_add(s0, n1)
	_ = Simd_m64_v128_store(m, n2, 0, n0)
	return
}

//go:noinline
func Simd_p_fx308(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_splat(int32(s1))
	n1 := Simd_m64_v128_load(m, s0, 0)
	n2 := Simd_i8x16_shuffle(n1, n1, [2]uint64{252119811, 0})
	return n1[0], n1[1], n2[0], n2[1], n0[0], n0[1]
}

//go:noinline
func Simd_p_fx309(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_i8x16_shuffle(n1, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n3 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n4 := Simd_i16x8_extend_low_i8x16_s(n3)
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_i32x4_mul(n2, n5)
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx310(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_i8x16_shuffle(n1, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_i32x4_mul(n2, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx311(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_i8x16_shuffle(n1, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n3 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n4 := Simd_i16x8_extend_low_i8x16_s(n3)
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_i32x4_mul(n2, n5)
	n7 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p5, p5h})
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_i8x16_shuffle(n8, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n10 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p6, p6h})
	n11 := Simd_i32x4_extend_low_i16x8_s(n10)
	n12 := Simd_i32x4_mul(n9, n11)
	n13 := Simd_i32x4_add(n6, n12)
	n14 := Simd_i32x4_shr_s(n13, 4)
	return n14[0], n14[1]
}

//go:noinline
func Simd_p_fx312(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_i8x16_shuffle(n1, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n3 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n4 := Simd_i16x8_extend_low_i8x16_s(n3)
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_i32x4_mul(n2, n5)
	n7 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p5, p5h})
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_i8x16_shuffle(n8, [2]uint64{p3, p3h}, [2]uint64{p2, p2h})
	n10 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p6, p6h})
	n11 := Simd_i32x4_extend_low_i16x8_s(n10)
	n12 := Simd_i32x4_mul(n9, n11)
	n13 := Simd_i32x4_add(n6, n12)
	n14 := Simd_i32x4_shr_s(n13, 4)
	return n14[0], n14[1]
}

//go:noinline
func Simd_p_fx313(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_f16x4_cvt(n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx314(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_scalar_i32_add(s1, s2)
	_ = Simd_m64_v128_store(m, n1, 0, n0)
	n3 := Simd_m64_v128_load(m, s3+-48, 0)
	n4 := Simd_m64_scalar_i32_add(s1, s4)
	n5 := Simd_m64_scalar_i32_add(s2, n4)
	_ = Simd_m64_v128_store(m, n5, 0, n3)
	n7 := Simd_m64_v128_load(m, s3+-32, 0)
	n8 := Simd_m64_scalar_i32_shl(s5, 3)
	n9 := Simd_m64_scalar_i32_add(s1, n8)
	n10 := Simd_m64_scalar_i32_add(s2, n9)
	_ = Simd_m64_v128_store(m, n10, 0, n7)
	n12 := Simd_m64_v128_load(m, s3+-16, 0)
	_ = Simd_m64_v128_store(m, s6, 0, n12)
	return
}

//go:noinline
func Simd_p_fx315(m *Module, s0 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n1 := Simd_m64_v128_load_nc(m, s0, 16)
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{521604871, 0})
	n3 := Simd_i8x16_shuffle(n0, n1, [2]uint64{504761862, 0})
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx316(m *Module, s0 int32, s1 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_i32x4_mul(n0, n3)
	n5 := Simd_i32x4_splat(s1)
	n6 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n7 := Simd_i32x4_extend_low_i16x8_s(n6)
	n8 := Simd_i32x4_mul(n5, n7)
	n9 := Simd_i32x4_add(n4, n8)
	n10 := Simd_i32x4_shr_s(n9, 4)
	return n10[0], n10[1]
}

//go:noinline
func Simd_p_fx317(m *Module, s0 int32, s1 int32, s2 int32, s3 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_i32x4_mul(n0, n3)
	n5 := Simd_i32x4_splat(s1)
	n6 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n7 := Simd_i32x4_extend_low_i16x8_s(n6)
	n8 := Simd_i32x4_mul(n5, n7)
	n9 := Simd_i32x4_add(n4, n8)
	n10 := Simd_i32x4_shr_s(n9, 4)
	n11 := Simd_i32x4_splat(s2)
	n12 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p1, p1h})
	n13 := Simd_i16x8_extend_low_i8x16_s(n12)
	n14 := Simd_i32x4_extend_low_i16x8_s(n13)
	n15 := Simd_i32x4_mul(n11, n14)
	n16 := Simd_i32x4_splat(s3)
	n17 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p4, p4h})
	n18 := Simd_i32x4_extend_low_i16x8_s(n17)
	n19 := Simd_i32x4_mul(n16, n18)
	n20 := Simd_i32x4_add(n15, n19)
	n21 := Simd_i32x4_shr_s(n20, 4)
	n22 := Simd_i32x4_add(n10, n21)
	return n22[0], n22[1]
}

//go:noinline
func Simd_p_fx318(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_m64_v128_load_rng(m, s0+40, 0, 0, 32)
	n3 := Simd_m64_v128_load_nc(m, s0+56, 0)
	n4 := Simd_i8x16_shuffle(n2, n3, [2]uint64{p0, p0h})
	return n1[0], n1[1], n2[0], n2[1], n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx319(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load32_splat(m, s0, 8793760)
	n1 := Simd_m64_v128_load_rng(m, s1+8, 0, 0, 32)
	n2 := Simd_m64_v128_load_nc(m, s1+24, 0)
	n3 := Simd_i8x16_shuffle(n1, n2, [2]uint64{p0, p0h})
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx320(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64, s7 int64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	n2 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	n4 := Simd_m64_v128_load(m, s2, 16)
	n5 := Simd_m64_scalar_i32_add(s4, s5)
	n6 := Simd_m64_scalar_i32_add(s3, n5)
	_ = Simd_m64_v128_store(m, n6, 16, n4)
	n8 := Simd_m64_v128_load(m, s2, 0)
	n9 := Simd_m64_scalar_i32_add(s4, s5)
	n10 := Simd_m64_scalar_i32_add(s3, n9)
	_ = Simd_m64_v128_store(m, n10, 0, n8)
	n12 := Simd_m64_v128_load(m, s6, 16)
	n13 := Simd_m64_scalar_i32_shl(s7, 3)
	n14 := Simd_m64_scalar_i32_add(s5, n13)
	n15 := Simd_m64_scalar_i32_add(s3, n14)
	_ = Simd_m64_v128_store(m, n15, 16, n12)
	n17 := Simd_m64_v128_load(m, s6, 0)
	n18 := Simd_m64_scalar_i32_shl(s7, 3)
	n19 := Simd_m64_scalar_i32_add(s5, n18)
	n20 := Simd_m64_scalar_i32_add(s3, n19)
	_ = Simd_m64_v128_store(m, n20, 0, n17)
	return
}

//go:noinline
func Simd_p_fx321(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	n2 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx322(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n3 := Simd_m64_v128_load_nc(m, s0, 16)
	n4 := Simd_i8x16_shuffle(n2, n3, [2]uint64{521604871, 0})
	return n1[0], n1[1], n2[0], n2[1], n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx323(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) {
	_ = Simd_m64_v128_store(m, s0, 0, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s1, 224)
	n2 := Simd_m64_scalar_i32_add(s2, s3)
	_ = Simd_m64_v128_store(m, n2, 0, n1)
	n4 := Simd_m64_v128_load(m, s1, 240)
	n5 := Simd_m64_scalar_i32_add(s2, s3)
	n6 := Simd_m64_scalar_i32_add(n5, 16)
	_ = Simd_m64_v128_store(m, n6, 0, n4)
	return
}

//go:noinline
func Simd_p_fx324(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 128)
	n1 := Simd_f32x4_sub(n0, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	n3 := Simd_m64_v128_load(m, s0, 144)
	n4 := Simd_f32x4_sub(n3, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s1, 16, n4)
	return
}

//go:noinline
func Simd_p_fx325(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_splat(int32(s0))
	n1 := Simd_i32x4_splat(int32(s2))
	n2 := Simd_m64_v128_load16x4_u(m, s3+24, 0)
	n3 := Simd_f16x4_cvt(n2)
	n4 := Simd_m64_v128_load32_zero(m, s1, 124)
	return n3[0], n3[1], n0[0], n0[1], n4[0], n4[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx326(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p2, p2h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_mul([2]uint64{p1, p1h}, n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	n5 := Simd_f32x4_mul(n4, [2]uint64{p3, p3h})
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx327(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p2, p2h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_mul([2]uint64{p1, p1h}, n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	n5 := Simd_f32x4_mul(n4, [2]uint64{p3, p3h})
	n6 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_i32x4_mul([2]uint64{p4, p4h}, n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	n11 := Simd_f32x4_mul(n10, [2]uint64{p3, p3h})
	n12 := Simd_f32x4_add(n11, [2]uint64{p6, p6h})
	n13 := Simd_f32x4_add(n5, n12)
	return n13[0], n13[1]
}

//go:noinline
func Simd_p_fx328(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0+16, 0)
	n1 := Simd_f16x4_cvt(n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx329(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s3, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_m64_scalar_i32_add(s0, s1)
	n3 := Simd_m64_v128_load(m, n2, 0)
	n4 := Simd_i8x16_shuffle(n3, n3, [2]uint64{252380931, 0})
	n5 := Simd_m64_v128_load32_zero(m, s2, 0)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	return n1[0], n1[1], n3[0], n3[1], n4[0], n4[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx330(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{235537922, 0})
	n1 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{218694913, 0})
	n2 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{201851904, 0})
	n3 := Simd_m64_v128_load32_zero(m, s0+16, 0)
	n4 := Simd_i16x8_extend_low_i8x16_u(n3)
	n5 := Simd_i32x4_extend_low_i16x8_u(n4)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx331(m *Module, s0 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_u(n1)
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_mul(n0, n3)
	n5 := Simd_i32x4_mul(n4, [2]uint64{p2, p2h})
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx332(m *Module, s0 int32, s1 int32, s2 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_u(n1)
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_mul(n0, n3)
	n5 := Simd_i32x4_mul(n4, [2]uint64{p2, p2h})
	n6 := Simd_i32x4_splat(s1)
	n7 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p1, p1h})
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	n10 := Simd_i32x4_mul(n6, n9)
	n11 := Simd_i32x4_mul(n10, [2]uint64{p2, p2h})
	n12 := Simd_i32x4_splat(s2)
	n13 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n14 := Simd_i32x4_extend_low_i16x8_u(n13)
	n15 := Simd_i32x4_mul(n12, n14)
	n16 := Simd_i32x4_mul(n15, [2]uint64{p5, p5h})
	n17 := Simd_i32x4_add(n11, n16)
	n18 := Simd_i32x4_add(n5, n17)
	return n18[0], n18[1]
}

//go:noinline
func Simd_p_fx333(m *Module, s0 int32, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_mul(n0, n2)
	n4 := Simd_i32x4_mul(n3, [2]uint64{p1, p1h})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx334(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load32_splat(m, n0, 0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx335(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_s([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_s([2]uint64{p1, p1h})
	n2 := Simd_i32x4_add(n0, n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx336(m *Module, s0 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_mul([2]uint64{p1, p1h}, n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_f32x4_mul(n3, [2]uint64{p2, p2h})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx337(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 128)
	n1 := Simd_f32x4_sub(n0, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	n3 := Simd_m64_v128_load(m, s0, 144)
	n4 := Simd_f32x4_sub(n3, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s1+16, 0, n4)
	n6 := Simd_m64_v128_load(m, s0, 176)
	n7 := Simd_f32x4_sub(n6, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s2+16, 0, n7)
	n9 := Simd_m64_v128_load(m, s0, 160)
	n10 := Simd_f32x4_sub(n9, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s2, 0, n10)
	n12 := Simd_m64_v128_load(m, s0, 208)
	n13 := Simd_f32x4_sub(n12, [2]uint64{p4, p4h})
	_ = Simd_m64_v128_store(m, s3+16, 0, n13)
	n15 := Simd_m64_v128_load(m, s0, 192)
	return n15[0], n15[1]
}

//go:noinline
func Simd_p_fx338(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_f32x4_sub([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 0, n0)
	n2 := Simd_m64_v128_load(m, s1, 224)
	n3 := Simd_f32x4_sub(n2, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s2, 0, n3)
	n5 := Simd_m64_v128_load(m, s1, 240)
	n6 := Simd_f32x4_sub(n5, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s2+16, 0, n6)
	return
}

//go:noinline
func Simd_p_fx339(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_splat(int32(s0))
	n1 := Simd_m64_v128_load16x4_u(m, s3+24, 0)
	n2 := Simd_f16x4_cvt(n1)
	n3 := Simd_m64_v128_load32_zero(m, s1, 124)
	n4 := Simd_m64_v128_load32_splat(m, s2, 0)
	return n2[0], n2[1], n0[0], n0[1], n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx340(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{504761862, 0})
	n1 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{487918853, 0})
	n2 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{471075844, 0})
	n3 := Simd_m64_v128_load32_zero(m, s0, 0)
	n4 := Simd_i16x8_extend_low_i8x16_u(n3)
	n5 := Simd_i32x4_extend_low_i16x8_u(n4)
	return n5[0], n5[1], n0[0], n0[1], n1[0], n1[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx341(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{437389826, 0})
	n1 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{420546817, 0})
	n2 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{403703808, 0})
	n3 := Simd_m64_v128_load32_zero(m, s0+16, 0)
	n4 := Simd_i16x8_extend_low_i8x16_u(n3)
	n5 := Simd_i32x4_extend_low_i16x8_u(n4)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx342(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_shl(s1, 2)
	n1 := Simd_m64_scalar_i32_add(s0, n0)
	n2 := Simd_m64_v128_load32_splat(m, n1, 0)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx343(m *Module, s0 int64, s1 int64, s2 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s1, s2)
	n1 := Simd_m64_v128_load16x4_u(m, n0, 0)
	n2 := Simd_f16x4_cvt(n1)
	n3 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n4 := Simd_m64_v128_load_nc(m, s0, 16)
	n5 := Simd_i8x16_shuffle(n3, n4, [2]uint64{521604871, 0})
	return n2[0], n2[1], n3[0], n3[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx344(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s3, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_m64_v128_load(m, s0, 0)
	n3 := Simd_i8x16_shuffle(n2, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_low_i8x16_u(n3)
	n5 := Simd_i32x4_extend_low_i16x8_u(n4)
	n6 := Simd_m64_scalar_i32_add(s1, s2)
	n7 := Simd_m64_v128_load(m, n6, 0)
	return n1[0], n1[1], n2[0], n2[1], n5[0], n5[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx345(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx346(m *Module, s0 int32, s1 int32, s2 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, s0)
	n1 := Simd_i32x4_shl(n0, 4)
	n2 := Simd_v128_and(n1, [2]uint64{p1, p1h})
	n3 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n4 := Simd_v128_or(n2, n3)
	n5 := Simd_i32x4_splat(s1)
	n6 := Simd_i32x4_mul(n4, n5)
	n7 := Simd_i32x4_shr_u([2]uint64{p4, p4h}, s0)
	n8 := Simd_i32x4_shl(n7, 4)
	n9 := Simd_v128_and(n8, [2]uint64{p1, p1h})
	n10 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p3, p3h})
	n11 := Simd_v128_or(n9, n10)
	n12 := Simd_i32x4_splat(s2)
	n13 := Simd_i32x4_mul(n11, n12)
	n14 := Simd_i32x4_add(n6, n13)
	return n14[0], n14[1]
}

//go:noinline
func Simd_p_fx347(m *Module, s0 int32, s1 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, s0)
	n1 := Simd_i32x4_shl(n0, 4)
	n2 := Simd_v128_and(n1, [2]uint64{p1, p1h})
	n3 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n4 := Simd_v128_or(n2, n3)
	n5 := Simd_i32x4_splat(s1)
	n6 := Simd_i32x4_mul(n4, n5)
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx348(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx349(m *Module, s0 int32, s1 int32, s2 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, s0)
	n1 := Simd_i32x4_shl(n0, 4)
	n2 := Simd_v128_and(n1, [2]uint64{p1, p1h})
	n3 := Simd_i32x4_shr_u([2]uint64{p2, p2h}, 4)
	n4 := Simd_v128_or(n2, n3)
	n5 := Simd_i32x4_splat(s1)
	n6 := Simd_i32x4_mul(n4, n5)
	n7 := Simd_i32x4_shr_u([2]uint64{p3, p3h}, s0)
	n8 := Simd_i32x4_shl(n7, 4)
	n9 := Simd_v128_and(n8, [2]uint64{p1, p1h})
	n10 := Simd_i32x4_shr_u([2]uint64{p4, p4h}, 4)
	n11 := Simd_v128_or(n9, n10)
	n12 := Simd_i32x4_splat(s2)
	n13 := Simd_i32x4_mul(n11, n12)
	n14 := Simd_i32x4_add(n6, n13)
	return n14[0], n14[1]
}

//go:noinline
func Simd_p_fx350(m *Module, s0 int32, s1 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, s0)
	n1 := Simd_i32x4_shl(n0, 4)
	n2 := Simd_v128_and(n1, [2]uint64{p1, p1h})
	n3 := Simd_i32x4_shr_u([2]uint64{p2, p2h}, 4)
	n4 := Simd_v128_or(n2, n3)
	n5 := Simd_i32x4_splat(s1)
	n6 := Simd_i32x4_mul(n4, n5)
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx351(m *Module, s0 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s0+24, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_m64_v128_load16x4_u(m, s0+16, 0)
	n3 := Simd_f16x4_cvt(n2)
	return n1[0], n1[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx352(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_splat(int32(s0))
	n1 := Simd_i32x4_splat(int32(s2))
	n2 := Simd_m64_v128_load32_zero(m, s1, 0)
	n3 := Simd_i16x8_extend_low_i8x16_u(n2)
	n4 := Simd_i32x4_extend_low_i16x8_u(n3)
	n5 := Simd_i32x4_mul(n0, n4)
	n6 := Simd_f32x4_convert_i32x4_s(n5)
	n7 := Simd_f32x4_mul([2]uint64{p0, p0h}, n6)
	n8 := Simd_f32x4_mul(n7, [2]uint64{p1, p1h})
	n9 := Simd_f32x4_add(n8, [2]uint64{p2, p2h})
	return n0[0], n0[1], n4[0], n4[1], n9[0], n9[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx353(m *Module, s0 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_mul([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_convert_i32x4_s(n0)
	n2 := Simd_f32x4_mul([2]uint64{p0, p0h}, n1)
	n3 := Simd_f32x4_mul(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p4, p4h})
	n5 := Simd_i32x4_splat(s0)
	n6 := Simd_i32x4_mul(n5, [2]uint64{p2, p2h})
	n7 := Simd_f32x4_convert_i32x4_s(n6)
	n8 := Simd_f32x4_mul([2]uint64{p0, p0h}, n7)
	n9 := Simd_f32x4_mul(n8, [2]uint64{p5, p5h})
	n10 := Simd_f32x4_add(n9, [2]uint64{p6, p6h})
	return n4[0], n4[1], n5[0], n5[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx354(m *Module, s0 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_mul(n0, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_f32x4_mul(n3, [2]uint64{p2, p2h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p3, p3h})
	return n0[0], n0[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx355(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_mul([2]uint64{p2, p2h}, n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul([2]uint64{p1, p1h}, n3)
	n5 := Simd_f32x4_mul(n4, [2]uint64{p3, p3h})
	n6 := Simd_f32x4_add(n5, [2]uint64{p4, p4h})
	return n1[0], n1[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx356(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_mul([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_convert_i32x4_s(n0)
	n2 := Simd_f32x4_mul([2]uint64{p0, p0h}, n1)
	n3 := Simd_f32x4_mul(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p4, p4h})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx357(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_f32x4_sub([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 0, n0)
	n2 := Simd_m64_v128_load(m, s1, 240)
	n3 := Simd_f32x4_sub(n2, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s2+16, 0, n3)
	n5 := Simd_m64_v128_load(m, s1, 224)
	n6 := Simd_f32x4_sub(n5, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s2, 0, n6)
	return
}

//go:noinline
func Simd_p_fx358(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n3 := Simd_m64_v128_load_nc(m, s0, 16)
	n4 := Simd_i8x16_shuffle(n2, n3, [2]uint64{p0, p0h})
	n5 := Simd_i16x8_extend_low_i8x16_u(n4)
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	return n1[0], n1[1], n2[0], n2[1], n3[0], n3[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx359(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n4 := Simd_m64_v128_load_nc(m, s0, 16)
	n5 := Simd_i8x16_shuffle(n3, n4, [2]uint64{p0, p0h})
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	return n3[0], n3[1], n4[0], n4[1], n7[0], n7[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx360(m *Module, s0 int32, s1 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, s0)
	n1 := Simd_i32x4_shl(n0, 4)
	n2 := Simd_v128_and(n1, [2]uint64{p1, p1h})
	n3 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n4 := Simd_v128_or(n2, n3)
	n5 := Simd_i32x4_add(n4, [2]uint64{p4, p4h})
	n6 := Simd_i32x4_splat(s1)
	n7 := Simd_i32x4_mul(n5, n6)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx361(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx362(m *Module, s0 int32, s1 int32, s2 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, s0)
	n1 := Simd_i32x4_shl(n0, 4)
	n2 := Simd_v128_and(n1, [2]uint64{p1, p1h})
	n3 := Simd_i32x4_shr_u([2]uint64{p2, p2h}, 4)
	n4 := Simd_v128_or(n2, n3)
	n5 := Simd_i32x4_add(n4, [2]uint64{p3, p3h})
	n6 := Simd_i32x4_splat(s1)
	n7 := Simd_i32x4_mul(n5, n6)
	n8 := Simd_i32x4_shr_u([2]uint64{p4, p4h}, s0)
	n9 := Simd_i32x4_shl(n8, 4)
	n10 := Simd_v128_and(n9, [2]uint64{p1, p1h})
	n11 := Simd_i32x4_shr_u([2]uint64{p5, p5h}, 4)
	n12 := Simd_v128_or(n10, n11)
	n13 := Simd_i32x4_add(n12, [2]uint64{p3, p3h})
	n14 := Simd_i32x4_splat(s2)
	n15 := Simd_i32x4_mul(n13, n14)
	n16 := Simd_i32x4_add(n7, n15)
	return n16[0], n16[1]
}

//go:noinline
func Simd_p_fx363(m *Module, s0 int32, s1 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, s0)
	n1 := Simd_i32x4_shl(n0, 4)
	n2 := Simd_v128_and(n1, [2]uint64{p1, p1h})
	n3 := Simd_i32x4_shr_u([2]uint64{p2, p2h}, 4)
	n4 := Simd_v128_or(n2, n3)
	n5 := Simd_i32x4_add(n4, [2]uint64{p3, p3h})
	n6 := Simd_i32x4_splat(s1)
	n7 := Simd_i32x4_mul(n5, n6)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx364(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s2)
	n1 := Simd_m64_v128_load16x4_u(m, n0, 0)
	n2 := Simd_f16x4_cvt(n1)
	n3 := Simd_m64_scalar_i32_add(s0, s1)
	n4 := Simd_m64_scalar_i32_add(n3, 1168)
	n5 := Simd_m64_v128_load_rng(m, n4, 0, 0, 32)
	n6 := Simd_m64_scalar_i32_add(s0, s1)
	n7 := Simd_m64_scalar_i32_add(n6, 1184)
	n8 := Simd_m64_v128_load_nc(m, n7, 0)
	n9 := Simd_i8x16_shuffle(n5, n8, [2]uint64{p0, p0h})
	n10 := Simd_i16x8_extend_low_i8x16_u(n9)
	n11 := Simd_i32x4_extend_low_i16x8_u(n10)
	return n2[0], n2[1], n5[0], n5[1], n8[0], n8[1], n11[0], n11[1]
}

//go:noinline
func Simd_p_fx365(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) {
	_ = Simd_m64_v128_store(m, s0, 0, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s1, 96)
	n2 := Simd_m64_scalar_i32_add(s2, s3)
	_ = Simd_m64_v128_store(m, n2, 0, n1)
	n4 := Simd_m64_v128_load(m, s1, 112)
	n5 := Simd_m64_scalar_i32_add(s2, s3)
	n6 := Simd_m64_scalar_i32_add(n5, 16)
	_ = Simd_m64_v128_store(m, n6, 0, n4)
	return
}

//go:noinline
func Simd_p_fx366(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i64x2_shl([2]uint64{p0, p0h}, 1)
	n1 := Simd_v128_or(n0, [2]uint64{p1, p1h})
	n2 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n3 := Simd_m64_v128_load_nc(m, s0, 16)
	n4 := Simd_i8x16_shuffle(n2, n3, [2]uint64{521604871, 0})
	return n2[0], n2[1], n3[0], n3[1], n4[0], n4[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx367(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shl([2]uint64{p0, p0h}, 1)
	n1 := Simd_v128_or(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx368(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx369(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_shl(n0, 17)
	n2 := Simd_m64_v128_load(m, s0, 0)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx370(m *Module, s0 int32, s1 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_u(n1)
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_mul(n0, n3)
	n5 := Simd_i32x4_mul(n4, [2]uint64{p2, p2h})
	n6 := Simd_i32x4_splat(s1)
	n7 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p1, p1h})
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	n10 := Simd_i32x4_mul(n6, n9)
	n11 := Simd_i32x4_mul(n10, [2]uint64{p4, p4h})
	n12 := Simd_i32x4_add(n5, n11)
	return n12[0], n12[1]
}

//go:noinline
func Simd_p_fx371(m *Module, s0 int32, s1 int32, s2 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_splat(s1)
	n2 := Simd_i32x4_mul(n1, [2]uint64{p0, p0h})
	n3 := Simd_i32x4_splat(s2)
	n4 := Simd_i32x4_mul(n3, [2]uint64{p1, p1h})
	n5 := Simd_i32x4_add(n2, n4)
	n6 := Simd_f32x4_convert_i32x4_s(n5)
	n7 := Simd_f32x4_mul(n0, n6)
	n8 := Simd_f32x4_mul(n7, [2]uint64{p2, p2h})
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx372(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_sub(n0, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	n3 := Simd_m64_v128_load(m, s0, 16)
	n4 := Simd_f32x4_sub(n3, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s1+16, 0, n4)
	n6 := Simd_m64_v128_load(m, s0, 48)
	n7 := Simd_f32x4_sub(n6, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s2+16, 0, n7)
	n9 := Simd_m64_v128_load(m, s0, 32)
	n10 := Simd_f32x4_sub(n9, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s2, 0, n10)
	n12 := Simd_m64_v128_load(m, s0, 80)
	n13 := Simd_f32x4_sub(n12, [2]uint64{p4, p4h})
	_ = Simd_m64_v128_store(m, s3+16, 0, n13)
	n15 := Simd_m64_v128_load(m, s0, 64)
	return n15[0], n15[1]
}

//go:noinline
func Simd_p_fx373(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_f32x4_sub([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 0, n0)
	n2 := Simd_m64_v128_load(m, s1, 96)
	n3 := Simd_f32x4_sub(n2, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s2, 0, n3)
	n5 := Simd_m64_v128_load(m, s1, 112)
	n6 := Simd_f32x4_sub(n5, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s2+16, 0, n6)
	return
}

//go:noinline
func Simd_p_fx374(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_sub(n0, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	n3 := Simd_m64_v128_load(m, s0, 16)
	n4 := Simd_f32x4_sub(n3, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s1, 16, n4)
	return
}

//go:noinline
func Simd_p_fx375(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_shl(n0, 17)
	n2 := Simd_m64_v128_load32_splat(m, s0, 0)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx376(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_i8x16_shuffle(n1, [2]uint64{p0, p0h}, [2]uint64{506097522914230528, 506097522914230528})
	return n1[0], n1[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx377(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_mul([2]uint64{p0, p0h}, n1)
	n3 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p3, p3h})
	n4 := Simd_i32x4_extend_low_i16x8_u(n3)
	n5 := Simd_i32x4_mul([2]uint64{p2, p2h}, n4)
	n6 := Simd_i32x4_add(n2, n5)
	n7 := Simd_f32x4_convert_i32x4_s(n6)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx378(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_extend_low_i16x8_s([2]uint64{p0, p0h})
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx379(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_add(n1, [2]uint64{p3, p3h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx380(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 4)
	n1 := Simd_v128_or(n0, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_mul(n1, [2]uint64{p2, p2h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx381(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s3, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_m64_v128_load32_splat(m, s0, 8793760)
	n3 := Simd_m64_v128_load32_splat(m, s1, 8793760)
	n4 := Simd_m64_v128_load32_splat(m, s2, 8793760)
	return n1[0], n1[1], n2[0], n2[1], n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx382(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n1 := Simd_m64_v128_load(m, s0, 0)
	n2 := Simd_i8x16_shuffle(n1, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	return n1[0], n1[1], n4[0], n4[1], n0[0], n0[1]
}

//go:noinline
func Simd_p_fx383(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 64)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	return n0[0], n0[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx384(m *Module, s0 int32, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_mul(n0, [2]uint64{p0, p0h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx385(m *Module, s0 int32, s1 int32, s2 int32, s3 int32, s4 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_i32x4_mul(n0, n2)
	n4 := Simd_i32x4_splat(s1)
	n5 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n6 := Simd_i32x4_extend_low_i16x8_s(n5)
	n7 := Simd_i32x4_mul(n4, n6)
	n8 := Simd_i32x4_splat(s2)
	n9 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n10 := Simd_i32x4_extend_low_i16x8_s(n9)
	n11 := Simd_i32x4_mul(n8, n10)
	n12 := Simd_i32x4_splat(s3)
	n13 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n14 := Simd_i32x4_extend_low_i16x8_s(n13)
	n15 := Simd_i32x4_mul(n12, n14)
	n16 := Simd_i32x4_add(n11, n15)
	n17 := Simd_i32x4_add(n7, n16)
	n18 := Simd_i32x4_splat(s4)
	n19 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p4, p4h})
	n20 := Simd_i32x4_extend_low_i16x8_s(n19)
	n21 := Simd_i32x4_mul(n18, n20)
	n22 := Simd_i32x4_add(n17, n21)
	n23 := Simd_i32x4_add(n3, n22)
	return n23[0], n23[1]
}

//go:noinline
func Simd_p_fx386(m *Module, s0 int32, s1 int32, s2 int32, s3 int32, s4 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_mul(n0, [2]uint64{p0, p0h})
	n2 := Simd_i32x4_splat(s1)
	n3 := Simd_i32x4_mul(n2, [2]uint64{p1, p1h})
	n4 := Simd_i32x4_splat(s2)
	n5 := Simd_i32x4_mul(n4, [2]uint64{p2, p2h})
	n6 := Simd_i32x4_splat(s3)
	n7 := Simd_i32x4_mul(n6, [2]uint64{p3, p3h})
	n8 := Simd_i32x4_add(n5, n7)
	n9 := Simd_i32x4_add(n3, n8)
	n10 := Simd_i32x4_splat(s4)
	n11 := Simd_i32x4_mul(n10, [2]uint64{p4, p4h})
	n12 := Simd_i32x4_add(n9, n11)
	n13 := Simd_i32x4_add(n1, n12)
	return n13[0], n13[1]
}

//go:noinline
func Simd_p_fx387(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n3 := Simd_m64_v128_load_nc(m, s0, 16)
	n4 := Simd_i8x16_shuffle(n2, n3, [2]uint64{521604871, 0})
	n5 := Simd_i16x8_extend_low_i8x16_u(n4)
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	return n1[0], n1[1], n2[0], n2[1], n3[0], n3[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx388(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i64x2_extend_high_i32x4_u(n0)
	n2 := Simd_i64x2_extend_low_i32x4_u(n0)
	return n1[0], n1[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx389(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{504761862, 0})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p2, p2h})
	n4 := Simd_i64x2_extend_high_i32x4_u(n3)
	n5 := Simd_i64x2_extend_low_i32x4_u(n3)
	return n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx390(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{487918853, 0})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p2, p2h})
	n4 := Simd_i64x2_extend_high_i32x4_u(n3)
	n5 := Simd_i64x2_extend_low_i32x4_u(n3)
	return n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx391(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{471075844, 0})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p2, p2h})
	n4 := Simd_i64x2_extend_high_i32x4_u(n3)
	n5 := Simd_i64x2_extend_low_i32x4_u(n3)
	return n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx392(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{454232835, 0})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p2, p2h})
	n4 := Simd_i64x2_extend_high_i32x4_u(n3)
	n5 := Simd_i64x2_extend_low_i32x4_u(n3)
	return n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx393(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{437389826, 0})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p2, p2h})
	n4 := Simd_i64x2_extend_high_i32x4_u(n3)
	n5 := Simd_i64x2_extend_low_i32x4_u(n3)
	return n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx394(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{420546817, 0})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p2, p2h})
	n4 := Simd_i64x2_extend_high_i32x4_u(n3)
	n5 := Simd_i64x2_extend_low_i32x4_u(n3)
	return n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx395(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{403703808, 0})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p2, p2h})
	n4 := Simd_i64x2_extend_high_i32x4_u(n3)
	n5 := Simd_i64x2_extend_low_i32x4_u(n3)
	return n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx396(m *Module, p0, p0h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 4)
	n1 := Simd_i64x2_extend_high_i32x4_u(n0)
	n2 := Simd_i64x2_extend_low_i32x4_u(n0)
	return n1[0], n1[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx397(m *Module, s0 int32, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_i32x4_mul(n0, n2)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx398(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64, s7 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	n2 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	n4 := Simd_m64_v128_load(m, s2, 16)
	n5 := Simd_m64_scalar_i32_shl(s4, 2)
	n6 := Simd_m64_scalar_i32_add(n5, s5)
	n7 := Simd_m64_scalar_i32_add(s3, n6)
	_ = Simd_m64_v128_store(m, n7, 16, n4)
	n9 := Simd_m64_v128_load(m, s2, 0)
	n10 := Simd_m64_scalar_i32_shl(s4, 2)
	n11 := Simd_m64_scalar_i32_add(n10, s5)
	n12 := Simd_m64_scalar_i32_add(s3, n11)
	_ = Simd_m64_v128_store(m, n12, 0, n9)
	n14 := Simd_m64_v128_load(m, s6, 16)
	n15 := Simd_m64_scalar_i32_shl(s4, 3)
	n16 := Simd_m64_scalar_i32_add(s5, n15)
	n17 := Simd_m64_scalar_i32_add(s3, n16)
	_ = Simd_m64_v128_store(m, n17, 16, n14)
	n19 := Simd_m64_v128_load(m, s6, 0)
	n20 := Simd_m64_scalar_i32_shl(s4, 3)
	n21 := Simd_m64_scalar_i32_add(s5, n20)
	n22 := Simd_m64_scalar_i32_add(s3, n21)
	_ = Simd_m64_v128_store(m, n22, 0, n19)
	n24 := Simd_m64_v128_load(m, s7, 16)
	return n24[0], n24[1]
}

//go:noinline
func Simd_p_fx399(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	_ = Simd_m64_v128_store(m, s0, 16, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s1, 0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	return
}

//go:noinline
func Simd_p_fx400(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 8793760)
	n1 := Simd_m64_v128_load32_lane(m, s1, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s2, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s3, 0, 3, n2)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx401(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0+72, 0, 0, 80)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{p2, p2h})
	n5 := Simd_i16x8_extend_low_i8x16_s(n4)
	n6 := Simd_i32x4_extend_low_i16x8_s(n5)
	n7 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{p3, p3h})
	n8 := Simd_i16x8_extend_low_i8x16_s(n7)
	n9 := Simd_i32x4_extend_low_i16x8_s(n8)
	n10 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{p4, p4h})
	n11 := Simd_i16x8_extend_low_i8x16_s(n10)
	n12 := Simd_i32x4_extend_low_i16x8_s(n11)
	return n3[0], n3[1], n6[0], n6[1], n9[0], n9[1], n12[0], n12[1]
}

//go:noinline
func Simd_p_fx402(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_nc(m, s0+136, 0)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	return n0[0], n0[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx403(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n1 := Simd_m64_v128_load_nc(m, s0, 16)
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{521604871, 0})
	n3 := Simd_i16x8_extend_low_i8x16_u(n2)
	n4 := Simd_i32x4_extend_low_i16x8_u(n3)
	n5 := Simd_v128_and(n4, [2]uint64{p0, p0h})
	return n0[0], n0[1], n1[0], n1[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx404(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 0)
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_m64_v128_load(m, s1, 0)
	return n0[0], n0[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx405(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p1, p1h}, 23)
	n1 := Simd_i32x4_add(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_lt_u([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_v128_bitselect([2]uint64{p0, p0h}, n1, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx406(m *Module, s0 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0+128, 0, 0, 32)
	n1 := Simd_m64_v128_load_nc(m, s0+144, 0)
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{521604871, 0})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_i8x16_shuffle(n0, n1, [2]uint64{504761862, 0})
	n6 := Simd_i16x8_extend_low_i8x16_s(n5)
	n7 := Simd_i32x4_extend_low_i16x8_s(n6)
	return n0[0], n0[1], n1[0], n1[1], n4[0], n4[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx407(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{487918853, 0})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx408(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{471075844, 0})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx409(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{454232835, 0})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx410(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{437389826, 0})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx411(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{420546817, 0})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx412(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{403703808, 0})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx413(m *Module, s0 int32, s1 int32, s2 int32, s3 int32, s4 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_mul(n0, [2]uint64{p0, p0h})
	n2 := Simd_i32x4_splat(s1)
	n3 := Simd_i32x4_mul(n2, [2]uint64{p1, p1h})
	n4 := Simd_i32x4_splat(s2)
	n5 := Simd_i32x4_mul(n4, [2]uint64{p2, p2h})
	n6 := Simd_i32x4_splat(s3)
	n7 := Simd_i32x4_mul(n6, [2]uint64{p3, p3h})
	n8 := Simd_i32x4_splat(s4)
	n9 := Simd_i32x4_mul(n8, [2]uint64{p4, p4h})
	n10 := Simd_i32x4_add(n7, n9)
	n11 := Simd_i32x4_add(n5, n10)
	n12 := Simd_i32x4_add(n3, n11)
	n13 := Simd_i32x4_add(n1, n12)
	return n13[0], n13[1]
}

//go:noinline
func Simd_p_fx414(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_convert_i32x4_s(n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx415(m *Module, s0 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_shr_u(n0, 16)
	return n0[0], n0[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx416(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_trunc_sat_f32x4_s(n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx417(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_shl(n1, 17)
	n3 := Simd_i32x4_max_u(n2, [2]uint64{p2, p2h})
	n4 := Simd_i32x4_shr_u(n3, 1)
	n5 := Simd_v128_and(n4, [2]uint64{p3, p3h})
	n6 := Simd_i32x4_add(n5, [2]uint64{p4, p4h})
	n7 := Simd_i32x4_shl(n1, 16)
	n8 := Simd_f32x4_abs(n7)
	n9 := Simd_f32x4_mul(n8, [2]uint64{p0, p0h})
	n10 := Simd_f32x4_mul(n9, [2]uint64{p1, p1h})
	n11 := Simd_f32x4_add(n10, n6)
	n12 := Simd_i32x4_extend_high_i16x8_u(n0)
	return n1[0], n1[1], n2[0], n2[1], n11[0], n11[1], n12[0], n12[1]
}

//go:noinline
func Simd_p_fx418(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p0, p0h}, 17)
	n1 := Simd_i32x4_max_u(n0, [2]uint64{p3, p3h})
	n2 := Simd_i32x4_shr_u(n1, 1)
	n3 := Simd_v128_and(n2, [2]uint64{p4, p4h})
	n4 := Simd_i32x4_add(n3, [2]uint64{p5, p5h})
	n5 := Simd_i32x4_shl([2]uint64{p0, p0h}, 16)
	n6 := Simd_f32x4_abs(n5)
	n7 := Simd_f32x4_mul(n6, [2]uint64{p1, p1h})
	n8 := Simd_f32x4_mul(n7, [2]uint64{p2, p2h})
	n9 := Simd_f32x4_add(n8, n4)
	return n0[0], n0[1], n9[0], n9[1]
}

//go:noinline
func Simd_p_fx419(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p1, p1h}, 13)
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p3, p3h})
	n3 := Simd_i32x4_add(n1, n2)
	n4 := Simd_i32x4_gt_u([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n5 := Simd_v128_bitselect([2]uint64{p0, p0h}, n3, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx420(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 328)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx421(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 192, n0)
	return
}

//go:noinline
func Simd_p_fx422(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 192)
	_ = Simd_m64_v128_store(m, s0, 120, n0)
	return
}

//go:noinline
func Simd_p_fx423(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 336)
	_ = Simd_m64_v128_store(m, s0, 200, n0)
	return
}

//go:noinline
func Simd_p_fx424(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 336)
	_ = Simd_m64_v128_store(m, s0, 128, n0)
	return
}

//go:noinline
func Simd_p_fx425(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 96)
	_ = Simd_m64_v128_store(m, s0, 80, n0)
	return
}

//go:noinline
func Simd_p_fx426(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1104)
	_ = Simd_m64_v128_store(m, s0+608, 0, n0)
	return
}

//go:noinline
func Simd_p_fx427(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1032)
	_ = Simd_m64_v128_store(m, s0, 224, n0)
	return
}

//go:noinline
func Simd_p_fx428(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 32, n0)
	return
}

//go:noinline
func Simd_p_fx429(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 24, n0)
	return
}

//go:noinline
func Simd_p_fx430(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 48, n0)
	return
}

//go:noinline
func Simd_p_fx431(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 800)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx432(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 856)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx433(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 944, n0)
	return
}

//go:noinline
func Simd_p_fx434(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 448, n0)
	return
}

//go:noinline
func Simd_p_fx435(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 424)
	_ = Simd_m64_v128_store(m, s0, 360, n0)
	return
}

//go:noinline
func Simd_p_fx436(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 424)
	_ = Simd_m64_v128_store(m, s0, 152, n0)
	return
}

//go:noinline
func Simd_p_fx437(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 424, n0)
	return
}

//go:noinline
func Simd_p_fx438(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 40)
	_ = Simd_m64_v128_store(m, s1, 40, n0)
	n2 := Simd_m64_v128_load(m, s0, 24)
	_ = Simd_m64_v128_store(m, s1, 24, n2)
	n4 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s1, 8, n4)
	return
}

//go:noinline
func Simd_p_fx439(m *Module, s0 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i64x2_extend_low_i32x4_s(n0)
	return n0[0], n0[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx440(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i8x16_eq(n0, [2]uint64{p1, p1h})
	n2 := Simd_v128_or([2]uint64{p0, p0h}, n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx441(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 0)
	n1 := Simd_i8x16_ne(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_u(n1)
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_v128_and(n3, [2]uint64{p2, p2h})
	n5 := Simd_i32x4_add([2]uint64{p0, p0h}, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx442(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{1084818905618843912, 216736831629295872})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	n2 := Simd_i8x16_shuffle(n1, n1, [2]uint64{216736831696667908, 216736831629295872})
	n3 := Simd_i32x4_add(n1, n2)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx443(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 352)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	n2 := Simd_m64_v128_load(m, s0, 336)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx444(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	_ = Simd_m64_v128_store(m, s0, 424, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s1, 24)
	_ = Simd_m64_v128_store(m, s0, 376, n1)
	n3 := Simd_m64_v128_load(m, s1, 8)
	_ = Simd_m64_v128_store(m, s0, 360, n3)
	return
}

//go:noinline
func Simd_p_fx445(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 336)
	_ = Simd_m64_v128_store(m, s0, 64, n0)
	return
}

//go:noinline
func Simd_p_fx446(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 24)
	_ = Simd_m64_v128_store(m, s0, 0, n0)
	return
}

//go:noinline
func Simd_p_fx447(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 96)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx448(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 464, n0)
	return
}

//go:noinline
func Simd_p_fx449(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 64, n0)
	return
}

//go:noinline
func Simd_p_fx450(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 584)
	_ = Simd_m64_v128_store(m, s0, 608, n0)
	return
}

//go:noinline
func Simd_p_fx451(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 656)
	_ = Simd_m64_v128_store(m, s1, 32, n0)
	return
}

//go:noinline
func Simd_p_fx452(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 544)
	_ = Simd_m64_v128_store(m, s0, 168, n0)
	return
}

//go:noinline
func Simd_p_fx453(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_neg(n0)
	n2 := Simd_m64_scalar_i32_add(s0, s1)
	_ = Simd_m64_v128_store(m, n2, 0, n0)
	n4 := Simd_m64_scalar_i32_add(s0, s2)
	_ = Simd_m64_v128_store(m, n4, 0, n0)
	n6 := Simd_m64_scalar_i32_add(s0, s3)
	_ = Simd_m64_v128_store(m, n6, 0, n1)
	return
}

//go:noinline
func Simd_p_fx454(m *Module, s0 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i8x16_shuffle(n0, n0, [2]uint64{1012195045828461056, 0})
	return n0[0], n0[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx455(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx456(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_shr_u([2]uint64{p0, p0h}, 7)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_v128_or(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_shr_u([2]uint64{p0, p0h}, 9)
	n4 := Simd_v128_and(n3, [2]uint64{p1, p1h})
	n5 := Simd_v128_or(n4, [2]uint64{p2, p2h})
	n6 := Simd_i8x16_shuffle(n2, n5, [2]uint64{p3, p3h})
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx457(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_shr_u([2]uint64{p0, p0h}, 11)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_v128_or(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_shr_u([2]uint64{p0, p0h}, 13)
	n4 := Simd_v128_or(n3, [2]uint64{p2, p2h})
	n5 := Simd_i8x16_shuffle(n2, n4, [2]uint64{p3, p3h})
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx458(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{5638, 5895})
	n1 := Simd_i8x16_shuffle([2]uint64{p2, p2h}, [2]uint64{p3, p3h}, [2]uint64{369491968, 386334720})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p4, p4h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx459(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{5124, 5381})
	n1 := Simd_i8x16_shuffle([2]uint64{p2, p2h}, [2]uint64{p3, p3h}, [2]uint64{335806464, 352649216})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p4, p4h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx460(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{4610, 4867})
	n1 := Simd_i8x16_shuffle([2]uint64{p2, p2h}, [2]uint64{p3, p3h}, [2]uint64{302120960, 318963712})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p4, p4h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx461(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{4096, 4353})
	n1 := Simd_i8x16_shuffle([2]uint64{p2, p2h}, [2]uint64{p3, p3h}, [2]uint64{268435456, 285278208})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p4, p4h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx462(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_v128_and(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_add(n2, [2]uint64{p3, p3h})
	return n0[0], n0[1], n1[0], n1[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx463(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_shr_u([2]uint64{p0, p0h}, 8)
	n1 := Simd_i16x8_add(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx464(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shr_u([2]uint64{p0, p0h}, 16)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_i16x8_add(n2, [2]uint64{p1, p1h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx465(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shr_u([2]uint64{p0, p0h}, 24)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_i16x8_add(n2, [2]uint64{p1, p1h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx466(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shr_u([2]uint64{p0, p0h}, 32)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_i16x8_add(n2, [2]uint64{p1, p1h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx467(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shr_u([2]uint64{p0, p0h}, 40)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_i16x8_add(n2, [2]uint64{p1, p1h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx468(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shr_u([2]uint64{p0, p0h}, 48)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_i16x8_add(n2, [2]uint64{p1, p1h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx469(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shr_u([2]uint64{p0, p0h}, 56)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_i16x8_add(n1, [2]uint64{p1, p1h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx470(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_shl([2]uint64{p1, p1h}, 2)
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	n2 := Simd_i16x8_shl([2]uint64{p2, p2h}, 4)
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_i16x8_shl([2]uint64{p3, p3h}, 6)
	n5 := Simd_v128_or(n3, n4)
	n6 := Simd_i16x8_shl([2]uint64{p4, p4h}, 8)
	n7 := Simd_v128_or(n5, n6)
	n8 := Simd_i16x8_shl([2]uint64{p5, p5h}, 10)
	n9 := Simd_v128_or(n7, n8)
	n10 := Simd_i16x8_shl([2]uint64{p6, p6h}, 12)
	n11 := Simd_v128_or(n9, n10)
	return n11[0], n11[1]
}

//go:noinline
func Simd_p_fx471(m *Module, s0 int64, s1 int64, s2 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_splat(int32(s1))
	n1 := Simd_i32x4_splat(int32(s2))
	n2 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n3 := Simd_m64_v128_load_nc(m, s0, 16)
	n4 := Simd_i8x16_shuffle(n2, n3, [2]uint64{521604871, 0})
	n5 := Simd_i16x8_extend_low_i8x16_s(n4)
	n6 := Simd_i32x4_extend_low_i16x8_s(n5)
	n7 := Simd_i32x4_add(n0, n6)
	n8 := Simd_i8x16_shuffle(n2, n3, [2]uint64{504761862, 0})
	n9 := Simd_i16x8_extend_low_i8x16_s(n8)
	n10 := Simd_i32x4_extend_low_i16x8_s(n9)
	n11 := Simd_i32x4_add(n1, n10)
	return n2[0], n2[1], n3[0], n3[1], n7[0], n7[1], n11[0], n11[1]
}

//go:noinline
func Simd_p_fx472(m *Module, s0 int32, s1 int32, s2 int32, s3 int32, s4 int32, s5 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{487918853, 0})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_i32x4_add(n0, n3)
	n5 := Simd_i32x4_mul(n4, n4)
	n6 := Simd_i32x4_splat(s1)
	n7 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{471075844, 0})
	n8 := Simd_i16x8_extend_low_i8x16_s(n7)
	n9 := Simd_i32x4_extend_low_i16x8_s(n8)
	n10 := Simd_i32x4_add(n6, n9)
	n11 := Simd_i32x4_mul(n10, n10)
	n12 := Simd_i32x4_splat(s2)
	n13 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{454232835, 0})
	n14 := Simd_i16x8_extend_low_i8x16_s(n13)
	n15 := Simd_i32x4_extend_low_i16x8_s(n14)
	n16 := Simd_i32x4_add(n12, n15)
	n17 := Simd_i32x4_mul(n16, n16)
	n18 := Simd_i32x4_splat(s3)
	n19 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{437389826, 0})
	n20 := Simd_i16x8_extend_low_i8x16_s(n19)
	n21 := Simd_i32x4_extend_low_i16x8_s(n20)
	n22 := Simd_i32x4_add(n18, n21)
	n23 := Simd_i32x4_mul(n22, n22)
	n24 := Simd_i32x4_splat(s4)
	n25 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{420546817, 0})
	n26 := Simd_i16x8_extend_low_i8x16_s(n25)
	n27 := Simd_i32x4_extend_low_i16x8_s(n26)
	n28 := Simd_i32x4_add(n24, n27)
	n29 := Simd_i32x4_mul(n28, n28)
	n30 := Simd_i32x4_splat(s5)
	n31 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{403703808, 0})
	n32 := Simd_i16x8_extend_low_i8x16_s(n31)
	n33 := Simd_i32x4_extend_low_i16x8_s(n32)
	n34 := Simd_i32x4_add(n30, n33)
	n35 := Simd_i32x4_mul(n34, n34)
	n36 := Simd_i32x4_mul([2]uint64{p2, p2h}, [2]uint64{p2, p2h})
	n37 := Simd_i32x4_mul([2]uint64{p3, p3h}, [2]uint64{p3, p3h})
	n38 := Simd_i32x4_add(n36, n37)
	n39 := Simd_i32x4_add(n38, n5)
	n40 := Simd_i32x4_add(n39, n11)
	n41 := Simd_i32x4_add(n40, n17)
	n42 := Simd_i32x4_add(n41, n23)
	n43 := Simd_i32x4_add(n42, n29)
	n44 := Simd_i32x4_add(n43, n35)
	return n44[0], n44[1]
}

//go:noinline
func Simd_p_fx473(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{1952900979473647880, 2242261670825954572})
	n1 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{1374179596769034496, 1663540288121341188})
	_ = Simd_m64_v128_store(m, s0, 16, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	return
}

//go:noinline
func Simd_p_fx474(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_add(n0, [2]uint64{p3, p3h})
	n2 := Simd_i32x4_trunc_sat_f32x4_s(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p4, p4h})
	n4 := Simd_i16x8_narrow_i32x4_u(n3, [2]uint64{p5, p5h})
	n5 := Simd_i8x16_narrow_i16x8_u(n4, [2]uint64{p5, p5h})
	n6 := Simd_i8x16_min_s(n5, [2]uint64{p6, p6h})
	n7 := Simd_v128_or([2]uint64{p0, p0h}, n6)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx475(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add(n0, [2]uint64{p2, p2h})
	n2 := Simd_i32x4_trunc_sat_f32x4_s(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p3, p3h})
	n4 := Simd_i16x8_narrow_i32x4_u(n3, [2]uint64{p4, p4h})
	n5 := Simd_i8x16_narrow_i16x8_u(n4, [2]uint64{p4, p4h})
	n6 := Simd_i8x16_min_s(n5, [2]uint64{p5, p5h})
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx476(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_sub([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_i32x4_trunc_sat_f32x4_s(n2)
	n4 := Simd_v128_and(n3, [2]uint64{p4, p4h})
	n5 := Simd_i16x8_narrow_i32x4_u(n4, [2]uint64{p5, p5h})
	n6 := Simd_i8x16_narrow_i16x8_u(n5, [2]uint64{p5, p5h})
	n7 := Simd_i8x16_min_s(n6, [2]uint64{p6, p6h})
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx477(m *Module, s0 int64, s1 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_m64_v128_load_rng(m, s0, 0, -64, 80)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add(n2, [2]uint64{p0, p0h})
	n4 := Simd_i32x4_trunc_sat_f32x4_s(n3)
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_i16x8_narrow_i32x4_u(n5, [2]uint64{p2, p2h})
	n7 := Simd_i8x16_narrow_i16x8_u(n6, [2]uint64{p2, p2h})
	n8 := Simd_i8x16_min_s(n7, [2]uint64{p3, p3h})
	n9 := Simd_m64_v128_load_nc(m, s1, 0)
	n10 := Simd_f32x4_mul(n0, n9)
	n11 := Simd_f32x4_add(n10, [2]uint64{p0, p0h})
	n12 := Simd_i32x4_trunc_sat_f32x4_s(n11)
	n13 := Simd_v128_and(n12, [2]uint64{p1, p1h})
	n14 := Simd_i16x8_narrow_i32x4_u(n13, [2]uint64{p2, p2h})
	n15 := Simd_i8x16_narrow_i16x8_u(n14, [2]uint64{p2, p2h})
	n16 := Simd_i8x16_min_s(n15, [2]uint64{p3, p3h})
	return n0[0], n0[1], n8[0], n8[1], n16[0], n16[1]
}

//go:noinline
func Simd_p_fx478(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx479(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_m64_v128_load_rng(m, s0, 0, -96, 112)
	n4 := Simd_f32x4_mul([2]uint64{p2, p2h}, n3)
	n5 := Simd_f32x4_add(n4, [2]uint64{p3, p3h})
	n6 := Simd_i32x4_trunc_sat_f32x4_s(n5)
	n7 := Simd_v128_and(n6, [2]uint64{p4, p4h})
	n8 := Simd_i16x8_narrow_i32x4_u(n7, [2]uint64{p5, p5h})
	n9 := Simd_i8x16_narrow_i16x8_u(n8, [2]uint64{p5, p5h})
	n10 := Simd_i8x16_min_s(n9, [2]uint64{p6, p6h})
	return n2[0], n2[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx480(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_m64_v128_load_nc(m, s0, 0)
	n4 := Simd_f32x4_mul([2]uint64{p2, p2h}, n3)
	n5 := Simd_f32x4_add(n4, [2]uint64{p3, p3h})
	n6 := Simd_i32x4_trunc_sat_f32x4_s(n5)
	n7 := Simd_v128_and(n6, [2]uint64{p4, p4h})
	n8 := Simd_i16x8_narrow_i32x4_u(n7, [2]uint64{p5, p5h})
	n9 := Simd_i8x16_narrow_i16x8_u(n8, [2]uint64{p5, p5h})
	n10 := Simd_i8x16_min_s(n9, [2]uint64{p6, p6h})
	return n2[0], n2[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx481(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_m64_v128_load(m, s0, 0)
	n4 := Simd_f32x4_mul([2]uint64{p2, p2h}, n3)
	n5 := Simd_f32x4_add(n4, [2]uint64{p3, p3h})
	n6 := Simd_i32x4_trunc_sat_f32x4_s(n5)
	n7 := Simd_v128_and(n6, [2]uint64{p4, p4h})
	n8 := Simd_i16x8_narrow_i32x4_u(n7, [2]uint64{p5, p5h})
	n9 := Simd_i8x16_narrow_i16x8_u(n8, [2]uint64{p5, p5h})
	n10 := Simd_i8x16_min_s(n9, [2]uint64{p6, p6h})
	return n2[0], n2[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx482(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_v128_or([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_v128_or(n0, n3)
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx483(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{1084818905618843912, 216736831629295872})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx484(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{216736831696667908, 216736831629295872})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	n2 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n3 := Simd_v128_or([2]uint64{p2, p2h}, n2)
	n4 := Simd_i8x16_shuffle(n1, n3, [2]uint64{1374179596769034496, 216736831629295872})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx485(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	n2 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p2, p2h})
	n3 := Simd_v128_or([2]uint64{p3, p3h}, n2)
	n4 := Simd_i8x16_shuffle(n1, n3, [2]uint64{216736831629295872, 1374179596769034496})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx486(m *Module, s0 int64, s1 int64, f0 float32, f1 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_splat(f1)
	n2 := Simd_m64_v128_load_rng(m, s0, 0, -64, 80)
	n3 := Simd_f32x4_sub(n2, n1)
	n4 := Simd_f32x4_mul(n0, n3)
	n5 := Simd_f32x4_add(n4, [2]uint64{p0, p0h})
	n6 := Simd_i32x4_trunc_sat_f32x4_u(n5)
	n7 := Simd_v128_and(n6, [2]uint64{p1, p1h})
	n8 := Simd_i16x8_narrow_i32x4_u(n7, [2]uint64{p2, p2h})
	n9 := Simd_i8x16_narrow_i16x8_u(n8, [2]uint64{p2, p2h})
	n10 := Simd_m64_v128_load_nc(m, s1, 0)
	n11 := Simd_f32x4_sub(n10, n1)
	n12 := Simd_f32x4_mul(n0, n11)
	n13 := Simd_f32x4_add(n12, [2]uint64{p0, p0h})
	n14 := Simd_i32x4_trunc_sat_f32x4_u(n13)
	n15 := Simd_v128_and(n14, [2]uint64{p1, p1h})
	n16 := Simd_i16x8_narrow_i32x4_u(n15, [2]uint64{p2, p2h})
	n17 := Simd_i8x16_narrow_i16x8_u(n16, [2]uint64{p2, p2h})
	return n0[0], n0[1], n1[0], n1[1], n9[0], n9[1], n17[0], n17[1]
}

//go:noinline
func Simd_p_fx487(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_m64_v128_load_rng(m, s0, 0, -96, 112)
	n4 := Simd_f32x4_sub(n3, [2]uint64{p3, p3h})
	n5 := Simd_f32x4_mul([2]uint64{p2, p2h}, n4)
	n6 := Simd_f32x4_add(n5, [2]uint64{p4, p4h})
	n7 := Simd_i32x4_trunc_sat_f32x4_u(n6)
	n8 := Simd_v128_and(n7, [2]uint64{p5, p5h})
	n9 := Simd_i16x8_narrow_i32x4_u(n8, [2]uint64{p6, p6h})
	n10 := Simd_i8x16_narrow_i16x8_u(n9, [2]uint64{p6, p6h})
	return n2[0], n2[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx488(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_m64_v128_load_nc(m, s0, 0)
	n4 := Simd_f32x4_sub(n3, [2]uint64{p3, p3h})
	n5 := Simd_f32x4_mul([2]uint64{p2, p2h}, n4)
	n6 := Simd_f32x4_add(n5, [2]uint64{p4, p4h})
	n7 := Simd_i32x4_trunc_sat_f32x4_u(n6)
	n8 := Simd_v128_and(n7, [2]uint64{p5, p5h})
	n9 := Simd_i16x8_narrow_i32x4_u(n8, [2]uint64{p6, p6h})
	n10 := Simd_i8x16_narrow_i16x8_u(n9, [2]uint64{p6, p6h})
	return n2[0], n2[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx489(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_m64_v128_load(m, s0, 0)
	n4 := Simd_f32x4_sub(n3, [2]uint64{p3, p3h})
	n5 := Simd_f32x4_mul([2]uint64{p2, p2h}, n4)
	n6 := Simd_f32x4_add(n5, [2]uint64{p4, p4h})
	n7 := Simd_i32x4_trunc_sat_f32x4_u(n6)
	n8 := Simd_v128_and(n7, [2]uint64{p5, p5h})
	n9 := Simd_i16x8_narrow_i32x4_u(n8, [2]uint64{p6, p6h})
	n10 := Simd_i8x16_narrow_i16x8_u(n9, [2]uint64{p6, p6h})
	return n2[0], n2[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx490(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_abs(n0)
	_ = Simd_m64_v128_store(m, s1, 128, n1)
	n3 := Simd_m64_v128_load(m, s0, 16)
	n4 := Simd_f32x4_abs(n3)
	_ = Simd_m64_v128_store(m, s1, 144, n4)
	n6 := Simd_m64_v128_load(m, s0, 32)
	n7 := Simd_f32x4_abs(n6)
	_ = Simd_m64_v128_store(m, s1, 160, n7)
	n9 := Simd_m64_v128_load(m, s0, 48)
	n10 := Simd_f32x4_abs(n9)
	_ = Simd_m64_v128_store(m, s1, 176, n10)
	return n1[0], n1[1], n4[0], n4[1], n7[0], n7[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx491(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 192)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx492(m *Module, s0 int64, s1 int64, f0 float32, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_m64_v128_load_rng(m, s0, 0, 0, 64)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add(n2, [2]uint64{p0, p0h})
	n4 := Simd_v128_and(n3, [2]uint64{p1, p1h})
	n5 := Simd_m64_v128_load_nc(m, s0, 16)
	n6 := Simd_f32x4_mul(n0, n5)
	n7 := Simd_f32x4_add(n6, [2]uint64{p0, p0h})
	n8 := Simd_v128_and(n7, [2]uint64{p1, p1h})
	n9 := Simd_i16x8_narrow_i32x4_u(n4, n8)
	n10 := Simd_m64_v128_load_nc(m, s0, 32)
	n11 := Simd_f32x4_mul(n0, n10)
	n12 := Simd_f32x4_add(n11, [2]uint64{p0, p0h})
	n13 := Simd_v128_and(n12, [2]uint64{p1, p1h})
	n14 := Simd_m64_v128_load_nc(m, s0, 48)
	n15 := Simd_f32x4_mul(n0, n14)
	n16 := Simd_f32x4_add(n15, [2]uint64{p0, p0h})
	n17 := Simd_v128_and(n16, [2]uint64{p1, p1h})
	n18 := Simd_i16x8_narrow_i32x4_u(n13, n17)
	n19 := Simd_i8x16_narrow_i16x8_u(n9, n18)
	_ = Simd_m64_v128_store(m, s1, 0, n19)
	return
}

//go:noinline
func Simd_p_fx493(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	return
}

//go:noinline
func Simd_p_fx494(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p4, p4h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p2, p2h})
	n5 := Simd_v128_and(n4, [2]uint64{p3, p3h})
	n6 := Simd_i16x8_narrow_i32x4_u(n2, n5)
	n7 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p5, p5h})
	n8 := Simd_f32x4_add(n7, [2]uint64{p2, p2h})
	n9 := Simd_v128_and(n8, [2]uint64{p3, p3h})
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p6, p6h})
	n11 := Simd_f32x4_add(n10, [2]uint64{p2, p2h})
	n12 := Simd_v128_and(n11, [2]uint64{p3, p3h})
	n13 := Simd_i16x8_narrow_i32x4_u(n9, n12)
	n14 := Simd_i8x16_narrow_i16x8_u(n6, n13)
	return n14[0], n14[1]
}

//go:noinline
func Simd_p_fx495(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 16, n2)
	n4 := Simd_m64_v128_load_rng(m, s1, 256, 224, 112)
	n5 := Simd_m64_v128_load_nc(m, s1, 224)
	n6 := Simd_m64_v128_load_nc(m, s1, 288)
	n7 := Simd_m64_v128_load_nc(m, s1, 320)
	return n4[0], n4[1], n5[0], n5[1], n6[0], n6[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx496(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 32, n2)
	n4 := Simd_m64_v128_load_rng(m, s1, 368, 336, 112)
	n5 := Simd_m64_v128_load_nc(m, s1, 336)
	n6 := Simd_m64_v128_load_nc(m, s1, 400)
	n7 := Simd_m64_v128_load_nc(m, s1, 432)
	return n4[0], n4[1], n5[0], n5[1], n6[0], n6[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx497(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 48, n2)
	n4 := Simd_m64_v128_load_rng(m, s1, 384, 352, 112)
	n5 := Simd_m64_v128_load_nc(m, s1, 352)
	n6 := Simd_m64_v128_load_nc(m, s1, 416)
	n7 := Simd_m64_v128_load_nc(m, s1, 448)
	return n4[0], n4[1], n5[0], n5[1], n6[0], n6[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx498(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 64, n2)
	return
}

//go:noinline
func Simd_p_fx499(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, f0 float32, f1 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_splat(f1)
	n2 := Simd_m64_scalar_i32_add(s0, s1)
	n3 := Simd_m64_scalar_i32_add(n2, 1024)
	n4 := Simd_m64_v128_load_rng(m, n3, 0, 0, 64)
	n5 := Simd_f32x4_add(n0, n4)
	n6 := Simd_f32x4_div(n5, n1)
	n7 := Simd_f32x4_add(n6, [2]uint64{p0, p0h})
	n8 := Simd_v128_and(n7, [2]uint64{p1, p1h})
	n9 := Simd_i32x4_max_u(n8, [2]uint64{p2, p2h})
	n10 := Simd_i32x4_sub(n9, [2]uint64{p2, p2h})
	n11 := Simd_i32x4_min_u(n10, [2]uint64{p3, p3h})
	n12 := Simd_i8x16_shuffle(n11, [2]uint64{p2, p2h}, [2]uint64{p4, p4h})
	n13 := Simd_m64_scalar_i32_add(s0, s1)
	n14 := Simd_m64_scalar_i32_add(n13, 1040)
	n15 := Simd_m64_v128_load_nc(m, n14, 0)
	n16 := Simd_f32x4_add(n0, n15)
	n17 := Simd_f32x4_div(n16, n1)
	n18 := Simd_f32x4_add(n17, [2]uint64{p0, p0h})
	n19 := Simd_v128_and(n18, [2]uint64{p1, p1h})
	n20 := Simd_i32x4_max_u(n19, [2]uint64{p2, p2h})
	n21 := Simd_i32x4_sub(n20, [2]uint64{p2, p2h})
	n22 := Simd_i32x4_min_u(n21, [2]uint64{p3, p3h})
	n23 := Simd_i8x16_shuffle(n22, [2]uint64{p2, p2h}, [2]uint64{p4, p4h})
	n24 := Simd_i8x16_shuffle(n12, n23, [2]uint64{1374179596769034496, 216736831629295872})
	n25 := Simd_m64_scalar_i32_add(s0, s1)
	n26 := Simd_m64_scalar_i32_add(n25, 1056)
	n27 := Simd_m64_v128_load_nc(m, n26, 0)
	n28 := Simd_f32x4_add(n0, n27)
	n29 := Simd_f32x4_div(n28, n1)
	n30 := Simd_f32x4_add(n29, [2]uint64{p0, p0h})
	n31 := Simd_v128_and(n30, [2]uint64{p1, p1h})
	n32 := Simd_i32x4_max_u(n31, [2]uint64{p2, p2h})
	n33 := Simd_i32x4_sub(n32, [2]uint64{p2, p2h})
	n34 := Simd_i32x4_min_u(n33, [2]uint64{p3, p3h})
	n35 := Simd_i8x16_shuffle(n34, [2]uint64{p2, p2h}, [2]uint64{p4, p4h})
	n36 := Simd_m64_scalar_i32_add(s0, s1)
	n37 := Simd_m64_scalar_i32_add(n36, 1072)
	n38 := Simd_m64_v128_load_nc(m, n37, 0)
	n39 := Simd_f32x4_add(n0, n38)
	n40 := Simd_f32x4_div(n39, n1)
	n41 := Simd_f32x4_add(n40, [2]uint64{p0, p0h})
	n42 := Simd_v128_and(n41, [2]uint64{p1, p1h})
	n43 := Simd_i32x4_max_u(n42, [2]uint64{p2, p2h})
	n44 := Simd_i32x4_sub(n43, [2]uint64{p2, p2h})
	n45 := Simd_i32x4_min_u(n44, [2]uint64{p3, p3h})
	n46 := Simd_i8x16_shuffle(n45, [2]uint64{p2, p2h}, [2]uint64{p4, p4h})
	n47 := Simd_i8x16_shuffle(n35, n46, [2]uint64{216736831629295872, 1374179596769034496})
	n48 := Simd_i8x16_shuffle(n24, n47, [2]uint64{506097522914230528, 2242261671028070680})
	n49 := Simd_m64_scalar_i32_add(s2, 208)
	n50 := Simd_m64_scalar_i32_add(n49, s3)
	_ = Simd_m64_v128_store(m, n50, 0, n48)
	return
}

//go:noinline
func Simd_p_fx500(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i8x16_add(n0, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	return
}

//go:noinline
func Simd_p_fx501(m *Module, s0 int64, s1 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_m64_v128_load_rng(m, s0, 0, 0, 64)
	n2 := Simd_f32x4_div(n1, n0)
	n3 := Simd_f32x4_add(n2, [2]uint64{p0, p0h})
	n4 := Simd_v128_and(n3, [2]uint64{p1, p1h})
	n5 := Simd_i32x4_max_u(n4, [2]uint64{p2, p2h})
	n6 := Simd_i32x4_min_u(n5, [2]uint64{p3, p3h})
	n7 := Simd_v128_and(n6, [2]uint64{p4, p4h})
	n8 := Simd_m64_v128_load_nc(m, s0, 16)
	n9 := Simd_f32x4_div(n8, n0)
	n10 := Simd_f32x4_add(n9, [2]uint64{p0, p0h})
	n11 := Simd_v128_and(n10, [2]uint64{p1, p1h})
	n12 := Simd_i32x4_max_u(n11, [2]uint64{p2, p2h})
	n13 := Simd_i32x4_min_u(n12, [2]uint64{p3, p3h})
	n14 := Simd_v128_and(n13, [2]uint64{p4, p4h})
	n15 := Simd_i16x8_narrow_i32x4_u(n7, n14)
	n16 := Simd_m64_v128_load_nc(m, s0, 32)
	n17 := Simd_f32x4_div(n16, n0)
	n18 := Simd_f32x4_add(n17, [2]uint64{p0, p0h})
	n19 := Simd_v128_and(n18, [2]uint64{p1, p1h})
	n20 := Simd_i32x4_max_u(n19, [2]uint64{p2, p2h})
	n21 := Simd_i32x4_min_u(n20, [2]uint64{p3, p3h})
	n22 := Simd_v128_and(n21, [2]uint64{p4, p4h})
	n23 := Simd_m64_v128_load_nc(m, s0, 48)
	n24 := Simd_f32x4_div(n23, n0)
	n25 := Simd_f32x4_add(n24, [2]uint64{p0, p0h})
	n26 := Simd_v128_and(n25, [2]uint64{p1, p1h})
	n27 := Simd_i32x4_max_u(n26, [2]uint64{p2, p2h})
	n28 := Simd_i32x4_min_u(n27, [2]uint64{p3, p3h})
	n29 := Simd_v128_and(n28, [2]uint64{p4, p4h})
	n30 := Simd_i16x8_narrow_i32x4_u(n22, n29)
	n31 := Simd_i8x16_narrow_i16x8_u(n15, n30)
	n32 := Simd_i8x16_add(n31, [2]uint64{p5, p5h})
	_ = Simd_m64_v128_store(m, s1, 0, n32)
	return
}

//go:noinline
func Simd_p_fx502(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 32, n2)
	n4 := Simd_m64_v128_load_rng(m, s1, 112, 80, 112)
	n5 := Simd_m64_v128_load_nc(m, s1, 80)
	n6 := Simd_m64_v128_load_nc(m, s1, 144)
	n7 := Simd_m64_v128_load_nc(m, s1, 176)
	return n4[0], n4[1], n5[0], n5[1], n6[0], n6[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx503(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 48, n2)
	n4 := Simd_m64_v128_load_rng(m, s1, 224, 192, 112)
	n5 := Simd_m64_v128_load_nc(m, s1, 192)
	n6 := Simd_m64_v128_load_nc(m, s1, 256)
	n7 := Simd_m64_v128_load_nc(m, s1, 288)
	return n4[0], n4[1], n5[0], n5[1], n6[0], n6[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx504(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 64, n2)
	n4 := Simd_m64_v128_load_rng(m, s1, 240, 208, 112)
	n5 := Simd_m64_v128_load_nc(m, s1, 208)
	n6 := Simd_m64_v128_load_nc(m, s1, 272)
	n7 := Simd_m64_v128_load_nc(m, s1, 304)
	return n4[0], n4[1], n5[0], n5[1], n6[0], n6[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx505(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 80, n2)
	return
}

//go:noinline
func Simd_p_fx506(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_i32x4_max_u(n2, [2]uint64{p4, p4h})
	n4 := Simd_i32x4_min_u(n3, [2]uint64{p5, p5h})
	n5 := Simd_v128_and(n4, [2]uint64{p6, p6h})
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx507(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, [2]uint64{506097522914230528, 1663540288323457296})
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and(n2, [2]uint64{p4, p4h})
	n4 := Simd_i32x4_max_u(n3, [2]uint64{p5, p5h})
	n5 := Simd_i32x4_min_u(n4, [2]uint64{p6, p6h})
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx508(m *Module, s0 int64, f0 float32, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_m64_v128_load_rng(m, s0, 0, 0, 64)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add(n2, [2]uint64{p0, p0h})
	n4 := Simd_m64_v128_load_nc(m, s0, 16)
	n5 := Simd_f32x4_mul(n0, n4)
	n6 := Simd_f32x4_add(n5, [2]uint64{p0, p0h})
	n7 := Simd_m64_v128_load_nc(m, s0, 32)
	n8 := Simd_f32x4_mul(n0, n7)
	n9 := Simd_f32x4_add(n8, [2]uint64{p0, p0h})
	n10 := Simd_m64_v128_load_nc(m, s0, 48)
	n11 := Simd_f32x4_mul(n0, n10)
	n12 := Simd_f32x4_add(n11, [2]uint64{p0, p0h})
	return n3[0], n3[1], n6[0], n6[1], n9[0], n9[1], n12[0], n12[1]
}

//go:noinline
func Simd_p_fx509(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_gt_u(n0, [2]uint64{p2, p2h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx510(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p2, p2h})
	n2 := Simd_i16x8_narrow_i32x4_u(n0, n1)
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p2, p2h})
	n4 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p2, p2h})
	n5 := Simd_i16x8_narrow_i32x4_u(n3, n4)
	n6 := Simd_i8x16_narrow_i16x8_u(n2, n5)
	n7 := Simd_v128_bitselect([2]uint64{p0, p0h}, n6, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0, 192, n7)
	return
}

//go:noinline
func Simd_p_fx511(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_gt_u(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p1, p1h})
	n3 := Simd_i32x4_gt_u(n2, [2]uint64{p2, p2h})
	n4 := Simd_i8x16_shuffle(n1, n3, [2]uint64{2024390091656922112, 0})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx512(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_nc(m, s0, 96)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_v128_or([2]uint64{p0, p0h}, n1)
	_ = Simd_m64_v128_store(m, s1, 32, n2)
	n4 := Simd_m64_v128_load_rng(m, s0, 128, 64, 80)
	n5 := Simd_m64_v128_load_nc(m, s0, 64)
	return n0[0], n0[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx513(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	n3 := Simd_m64_v128_load_rng(m, s1, 176, 112, 80)
	n4 := Simd_m64_v128_load_nc(m, s1, 112)
	return n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx514(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 48, n1)
	n3 := Simd_m64_v128_load_rng(m, s1, 144, 80, 80)
	n4 := Simd_m64_v128_load_nc(m, s1, 80)
	return n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx515(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 16, n1)
	return
}

//go:noinline
func Simd_p_fx516(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i8x16_shr_s([2]uint64{p0, p0h}, 2)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_shr_s([2]uint64{p2, p2h}, 4)
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n5 := Simd_v128_or(n3, n4)
	n6 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n7 := Simd_v128_or(n5, n6)
	_ = Simd_m64_v128_store(m, s0, 128, n7)
	return
}

//go:noinline
func Simd_p_fx517(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i8x16_shr_s([2]uint64{p0, p0h}, 2)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_shr_s([2]uint64{p2, p2h}, 4)
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n5 := Simd_v128_or(n3, n4)
	n6 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n7 := Simd_v128_or(n5, n6)
	_ = Simd_m64_v128_store(m, s0, 144, n7)
	return
}

//go:noinline
func Simd_p_fx518(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i8x16_shr_s([2]uint64{p0, p0h}, 2)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_shr_s([2]uint64{p2, p2h}, 4)
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n5 := Simd_v128_or(n3, n4)
	n6 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n7 := Simd_v128_or(n5, n6)
	_ = Simd_m64_v128_store(m, s0, 160, n7)
	return
}

//go:noinline
func Simd_p_fx519(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	n2 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p2, p2h})
	n3 := Simd_v128_or([2]uint64{p3, p3h}, n2)
	_ = Simd_m64_v128_store(m, s0, 96, n1)
	_ = Simd_m64_v128_store(m, s0, 64, n3)
	n6 := Simd_m64_v128_load_rng(m, s1, 272, 208, 80)
	n7 := Simd_m64_v128_load_nc(m, s1, 208)
	return n6[0], n6[1], n7[0], n7[1]
}

//go:noinline
func Simd_p_fx520(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 80, n1)
	n3 := Simd_m64_v128_load_rng(m, s1, 304, 240, 80)
	n4 := Simd_m64_v128_load_nc(m, s1, 240)
	return n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx521(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 112, n1)
	return
}

//go:noinline
func Simd_p_fx522(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i8x16_shr_s([2]uint64{p0, p0h}, 2)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_shr_s([2]uint64{p2, p2h}, 4)
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n5 := Simd_v128_or(n3, n4)
	n6 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n7 := Simd_v128_or(n5, n6)
	_ = Simd_m64_v128_store(m, s0, 176, n7)
	return
}

//go:noinline
func Simd_p_fx523(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 272, n1)
	n3 := Simd_m64_v128_load(m, s0, 16)
	n4 := Simd_f32x4_mul(n3, n3)
	_ = Simd_m64_v128_store(m, s1, 256, n4)
	return n0[0], n0[1], n1[0], n1[1], n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx524(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 240, n1)
	n3 := Simd_m64_v128_load(m, s0, 48)
	n4 := Simd_f32x4_mul(n3, n3)
	return n0[0], n0[1], n1[0], n1[1], n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx525(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{795458214199165184, 216736831629295872})
	n2 := Simd_i64x2_extend_low_i32x4_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx526(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 16)
	n1 := Simd_m64_v128_load32_lane(m, s0+24, 0, 1, n0)
	n2 := Simd_i64x2_extend_low_i32x4_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx527(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, [2]uint64{506097522914230528, 1663540288323457296})
	n1 := Simd_i8x16_sub([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 224, n1)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx528(m *Module, s0 int32, s1 int32, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_splat(s0)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{81907486720, 0})
	n2 := Simd_i16x8_splat(s1)
	n3 := Simd_i8x16_shuffle(n1, n2, [2]uint64{1229482715502477568, 0})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx529(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+48, 0)
	n1 := Simd_f32x4_mul(n0, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx530(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul(n0, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx531(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	n3 := Simd_m64_v128_load(m, s0, 16)
	n4 := Simd_f32x4_mul(n3, n3)
	_ = Simd_m64_v128_store(m, s1, 16, n4)
	n6 := Simd_m64_v128_load(m, s0, 32)
	n7 := Simd_f32x4_mul(n6, n6)
	_ = Simd_m64_v128_store(m, s1, 32, n7)
	n9 := Simd_m64_v128_load(m, s0, 48)
	n10 := Simd_f32x4_mul(n9, n9)
	_ = Simd_m64_v128_store(m, s1, 48, n10)
	n12 := Simd_m64_v128_load(m, s0, 64)
	n13 := Simd_f32x4_mul(n12, n12)
	_ = Simd_m64_v128_store(m, s1, 64, n13)
	n15 := Simd_m64_v128_load(m, s0, 80)
	n16 := Simd_f32x4_mul(n15, n15)
	_ = Simd_m64_v128_store(m, s1, 80, n16)
	n18 := Simd_m64_v128_load(m, s0, 96)
	n19 := Simd_f32x4_mul(n18, n18)
	_ = Simd_m64_v128_store(m, s1, 96, n19)
	return
}

//go:noinline
func Simd_p_fx532(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 112)
	n1 := Simd_f32x4_mul(n0, n0)
	_ = Simd_m64_v128_store(m, s1, 112, n1)
	return
}

//go:noinline
func Simd_p_fx533(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 648)
	_ = Simd_m64_v128_store(m, s0, 40, n0)
	return
}

//go:noinline
func Simd_p_fx534(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 664)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	n2 := Simd_m64_v128_load(m, s0, 648)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx535(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s1, 8, n0)
	n2 := Simd_m64_v128_load(m, s0, 24)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx536(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	_ = Simd_m64_v128_store(m, s0, 24, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s1, 40)
	_ = Simd_m64_v128_store(m, s1, 40, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 40, n1)
	n4 := Simd_m64_v128_load(m, s1, 56)
	_ = Simd_m64_v128_store(m, s0, 56, n4)
	return
}

//go:noinline
func Simd_p_fx537(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 6472)
	_ = Simd_m64_v128_store(m, s1, 8, n0)
	return
}

//go:noinline
func Simd_p_fx538(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 392)
	_ = Simd_m64_v128_store(m, s0, 24, n0)
	return
}

//go:noinline
func Simd_p_fx539(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 376)
	_ = Simd_m64_v128_store(m, s0, 8, n0)
	return
}

//go:noinline
func Simd_p_fx540(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 168)
	_ = Simd_m64_v128_store(m, s1, 168, n0)
	return
}

//go:noinline
func Simd_p_fx541(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s0, 0, [2]uint64{p0, p0h})
	n2 := Simd_m64_scalar_i32_add(s1, s2)
	_ = Simd_m64_v128_store(m, n2, 0, n0)
	n4 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s0, 16, [2]uint64{p0, p0h})
	n6 := Simd_m64_scalar_i32_add(s1, s2)
	_ = Simd_m64_v128_store(m, n6, 16, n4)
	n8 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s0, 32, [2]uint64{p0, p0h})
	n10 := Simd_m64_scalar_i32_add(s1, s2)
	_ = Simd_m64_v128_store(m, n10, 32, n8)
	n12 := Simd_m64_v128_load(m, s0, 48)
	n13 := Simd_m64_scalar_i32_add(s1, s2)
	_ = Simd_m64_v128_store(m, n13, 48, n12)
	return
}

//go:noinline
func Simd_p_fx542(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1272)
	_ = Simd_m64_v128_store(m, s0, 936, n0)
	return
}

//go:noinline
func Simd_p_fx543(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 904)
	_ = Simd_m64_v128_store(m, s0, 800, n0)
	return
}

//go:noinline
func Simd_p_fx544(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1216)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx545(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 832)
	_ = Simd_m64_v128_store(m, s0, 624, n0)
	return
}

//go:noinline
func Simd_p_fx546(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1264)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx547(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1216)
	_ = Simd_m64_v128_store(m, s0, 1240, n0)
	return
}

//go:noinline
func Simd_p_fx548(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1+328, 0, n0)
	return
}

//go:noinline
func Simd_p_fx549(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 512, n0)
	return
}

//go:noinline
func Simd_p_fx550(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 512)
	_ = Simd_m64_v128_store(m, s0, 208, n0)
	return
}

//go:noinline
func Simd_p_fx551(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 512)
	_ = Simd_m64_v128_store(m, s0, 176, n0)
	return
}

//go:noinline
func Simd_p_fx552(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 512)
	_ = Simd_m64_v128_store(m, s0, 272, n0)
	return
}

//go:noinline
func Simd_p_fx553(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 512)
	_ = Simd_m64_v128_store(m, s0, 288, n0)
	return
}

//go:noinline
func Simd_p_fx554(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 512)
	_ = Simd_m64_v128_store(m, s0, 112, n0)
	return
}

//go:noinline
func Simd_p_fx555(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 512)
	_ = Simd_m64_v128_store(m, s0, 16, n0)
	return
}

//go:noinline
func Simd_p_fx556(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1144)
	_ = Simd_m64_v128_store(m, s1, 152, n0)
	return
}

//go:noinline
func Simd_p_fx557(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1008)
	_ = Simd_m64_v128_store(m, s1, 208, n0)
	return
}

//go:noinline
func Simd_p_fx558(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1008)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx559(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 1144)
	_ = Simd_m64_v128_store(m, s1, 48, n0)
	_ = Simd_m64_v128_store(m, s1, 0, [2]uint64{p0, p0h})
	return
}

//go:noinline
func Simd_p_fx560(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load(m, n0, 0)
	_ = Simd_m64_v128_store(m, s2, 0, n1)
	return
}

//go:noinline
func Simd_p_fx561(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	n2 := Simd_m64_v128_load(m, s0+16, 0)
	_ = Simd_m64_v128_store(m, s1+16, 0, n2)
	n4 := Simd_m64_v128_load(m, s0+32, 0)
	_ = Simd_m64_v128_store(m, s1+32, 0, n4)
	n6 := Simd_m64_v128_load(m, s0+48, 0)
	_ = Simd_m64_v128_store(m, s1+48, 0, n6)
	return
}

//go:noinline
func Simd_p_fx562(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, 23052)
	n1 := Simd_m64_scalar_i32_add(n0, s1)
	n2 := Simd_m64_v128_load(m, n1, 0)
	n3 := Simd_i32x4_gt_s(n2, [2]uint64{p0, p0h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx563(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 136, n0)
	return
}

//go:noinline
func Simd_p_fx564(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 112, n0)
	return
}

//go:noinline
func Simd_p_fx565(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 968, n0)
	return
}

//go:noinline
func Simd_p_fx566(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 992, n0)
	return
}

//go:noinline
func Simd_p_fx567(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 1064, n0)
	return
}

//go:noinline
func Simd_p_fx568(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 1016, n0)
	return
}

//go:noinline
func Simd_p_fx569(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 240)
	n1 := Simd_m64_v128_load(m, s1, 1024)
	_ = Simd_m64_v128_store(m, s0, 240, n1)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx570(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1064)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx571(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 1016)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx572(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx573(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_i8x16_shuffle(n1, n1, [2]uint64{p0, p0h})
	n3 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s1, 0, n3)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx574(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s0, 40, n0)
	return
}

//go:noinline
func Simd_p_fx575(m *Module, s0 int32, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_add(n0, [2]uint64{p0, p0h})
	n2 := Simd_i32x4_lt_u(n1, [2]uint64{p1, p1h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx576(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 168, n0)
	return
}

//go:noinline
func Simd_p_fx577(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 104, n0)
	return
}

//go:noinline
func Simd_p_fx578(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 144)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx579(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s1, 8, n0)
	return
}

//go:noinline
func Simd_p_fx580(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 64)
	_ = Simd_m64_v128_store(m, s1, 8, n0)
	return
}

//go:noinline
func Simd_p_fx581(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 176)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx582(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) {
	_ = Simd_m64_v128_store(m, s0, 8, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s1, 32)
	_ = Simd_m64_v128_store(m, s2+160, 0, n1)
	return
}

//go:noinline
func Simd_p_fx583(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	_ = Simd_m64_v128_store(m, s0, 16, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx584(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s1, 32, n0)
	return
}

//go:noinline
func Simd_p_fx585(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 584)
	_ = Simd_m64_v128_store(m, s0, 64, n0)
	return
}

//go:noinline
func Simd_p_fx586(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 280, n0)
	n2 := Simd_m64_v128_load(m, s2, 0)
	_ = Simd_m64_v128_store(m, s1, 264, n2)
	return
}

//go:noinline
func Simd_p_fx587(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 240)
	_ = Simd_m64_v128_store(m, s0, 48, n0)
	return
}

//go:noinline
func Simd_p_fx588(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 224)
	_ = Simd_m64_v128_store(m, s0, 32, n0)
	return
}

//go:noinline
func Simd_p_fx589(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 8, n0)
	return
}

//go:noinline
func Simd_p_fx590(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 600)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx591(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 680, n0)
	return
}

//go:noinline
func Simd_p_fx592(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 72)
	_ = Simd_m64_v128_store(m, s0, 24, n0)
	return
}

//go:noinline
func Simd_p_fx593(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 56)
	_ = Simd_m64_v128_store(m, s0, 56, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 8, n0)
	return
}

//go:noinline
func Simd_p_fx594(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 24)
	_ = Simd_m64_v128_store(m, s0, 48, n0)
	return
}

//go:noinline
func Simd_p_fx595(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s0, 8, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 80, n0)
	return
}

//go:noinline
func Simd_p_fx596(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 48)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	n2 := Simd_m64_v128_load(m, s0+72, 8)
	_ = Simd_m64_v128_store(m, s1, 32, n2)
	return
}

//go:noinline
func Simd_p_fx597(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 96)
	_ = Simd_m64_v128_store(m, s0, 64, n0)
	return
}

//go:noinline
func Simd_p_fx598(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 128, n0)
	return
}

//go:noinline
func Simd_p_fx599(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 224, n0)
	return
}

//go:noinline
func Simd_p_fx600(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 288, n0)
	return
}

//go:noinline
func Simd_p_fx601(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 48)
	_ = Simd_m64_v128_store(m, s0, 120, n0)
	return
}

//go:noinline
func Simd_p_fx602(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 48)
	_ = Simd_m64_v128_store(m, s0, 96, n0)
	return
}

//go:noinline
func Simd_p_fx603(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 48)
	_ = Simd_m64_v128_store(m, s0, 8, n0)
	return
}

//go:noinline
func Simd_p_fx604(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 40)
	_ = Simd_m64_v128_store(m, s0, 88, n0)
	return
}

//go:noinline
func Simd_p_fx605(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 40)
	_ = Simd_m64_v128_store(m, s0, 0, n0)
	return
}

//go:noinline
func Simd_p_fx606(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 120, n0)
	return
}

//go:noinline
func Simd_p_fx607(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	return
}

//go:noinline
func Simd_p_fx608(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 72, n0)
	return
}

//go:noinline
func Simd_p_fx609(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 96, n0)
	return
}

//go:noinline
func Simd_p_fx610(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 64)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx611(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 168)
	_ = Simd_m64_v128_store(m, s0, 88, n0)
	n2 := Simd_m64_v128_load(m, s0, 152)
	_ = Simd_m64_v128_store(m, s0, 72, n2)
	n4 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s0, 56, n4)
	return
}

//go:noinline
func Simd_p_fx612(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 168)
	_ = Simd_m64_v128_store(m, s0, 32, n0)
	n2 := Simd_m64_v128_load(m, s0, 152)
	_ = Simd_m64_v128_store(m, s0, 16, n2)
	n4 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s0, 0, n4)
	return
}

//go:noinline
func Simd_p_fx613(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 168)
	_ = Simd_m64_v128_store(m, s0, 40, n0)
	n2 := Simd_m64_v128_load(m, s0, 152)
	_ = Simd_m64_v128_store(m, s0, 24, n2)
	n4 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s0, 8, n4)
	return
}

//go:noinline
func Simd_p_fx614(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 8, n0)
	n2 := Simd_m64_v128_load(m, s2, 0)
	_ = Simd_m64_v128_store(m, s1, 24, n2)
	n4 := Simd_m64_v128_load(m, s3, 0)
	_ = Simd_m64_v128_store(m, s1, 40, n4)
	return
}

//go:noinline
func Simd_p_fx615(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_f64x2_promote_low_f32x4([2]uint64{p0, p0h})
	n1 := Simd_f64x2_div(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx616(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 88, n0)
	return
}

//go:noinline
func Simd_p_fx617(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 264, n0)
	return
}

//go:noinline
func Simd_p_fx618(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 216, n0)
	return
}

//go:noinline
func Simd_p_fx619(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 192)
	_ = Simd_m64_v128_store(m, s0, 40, n0)
	n2 := Simd_m64_v128_load(m, s0, 176)
	_ = Simd_m64_v128_store(m, s0, 24, n2)
	n4 := Simd_m64_v128_load(m, s0, 160)
	_ = Simd_m64_v128_store(m, s0, 8, n4)
	return
}

//go:noinline
func Simd_p_fx620(m *Module, s0 int64, s1 int64, s2 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, s1)
	_ = Simd_m64_v128_store(m, s2, 92, n0)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx621(m *Module, s0 int64, s1 int64, s2 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, s1)
	_ = Simd_m64_v128_store(m, s2, 108, n0)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx622(m *Module, s0 int64, s1 int64, s2 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, s1)
	_ = Simd_m64_v128_store(m, s2, s1, n0)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx623(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 88)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx624(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 8, n0)
	return
}

//go:noinline
func Simd_p_fx625(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	n1 := Simd_m64_scalar_i32_add(s1, s2)
	_ = Simd_m64_v128_store(m, n1, 8, n0)
	return
}

//go:noinline
func Simd_p_fx626(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_scalar_i32_add(s1, s2)
	_ = Simd_m64_v128_store(m, n1, 8, n0)
	return
}

//go:noinline
func Simd_p_fx627(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	return
}

//go:noinline
func Simd_p_fx628(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 152)
	_ = Simd_m64_v128_store(m, s1, 272, n0)
	return
}

//go:noinline
func Simd_p_fx629(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s1, 2304, n0)
	n2 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 2288, n2)
	return
}

//go:noinline
func Simd_p_fx630(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 48)
	_ = Simd_m64_v128_store(m, s1, 132, n0)
	n2 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s1, 116, n2)
	n4 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 100, n4)
	n6 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 84, n6)
	return
}

//go:noinline
func Simd_p_fx631(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 40, n0)
	n2 := Simd_m64_v128_load(m, s2, 0)
	_ = Simd_m64_v128_store(m, s1, 56, n2)
	n4 := Simd_m64_v128_load(m, s3, 0)
	_ = Simd_m64_v128_store(m, s1, 72, n4)
	return
}

//go:noinline
func Simd_p_fx632(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_m64_v128_load(m, s0, 0)
	n2 := Simd_f32x4_sub(n0, n1)
	n3 := Simd_f32x4_abs(n2)
	n4 := Simd_f32x4_sub([2]uint64{p0, p0h}, n1)
	n5 := Simd_f32x4_abs(n4)
	return n1[0], n1[1], n5[0], n5[1], n0[0], n0[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx633(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p0, p0h})
	n1 := Simd_f32x4_sub(n0, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_abs(n1)
	n3 := Simd_f32x4_pmin([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p4, p4h})
	return n0[0], n0[1], n2[0], n2[1], n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx634(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_sub([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_abs(n0)
	n2 := Simd_f32x4_pmin([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_mul([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_f32x4_sub(n3, [2]uint64{p1, p1h})
	n5 := Simd_f32x4_abs(n4)
	return n1[0], n1[1], n2[0], n2[1], n3[0], n3[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx635(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_pmin([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_pmin(n0, [2]uint64{p5, p5h})
	n2 := Simd_f32x4_mul([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_sub(n2, [2]uint64{p4, p4h})
	n4 := Simd_f32x4_abs(n3)
	return n0[0], n0[1], n2[0], n2[1], n4[0], n4[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx636(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_sub(n0, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_abs(n1)
	n3 := Simd_f32x4_pmin([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p5, p5h})
	return n0[0], n0[1], n2[0], n2[1], n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx637(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_sub([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_abs(n0)
	n2 := Simd_f32x4_pmin([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_pmin(n2, n1)
	n4 := Simd_f32x4_neg([2]uint64{p1, p1h})
	n5 := Simd_f32x4_sub(n4, [2]uint64{p4, p4h})
	n6 := Simd_f32x4_abs(n5)
	return n1[0], n1[1], n2[0], n2[1], n3[0], n3[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx638(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_m64_v128_load(m, s0, 0)
	n2 := Simd_f32x4_sub([2]uint64{p0, p0h}, n1)
	n3 := Simd_f32x4_abs(n2)
	n4 := Simd_f32x4_sub([2]uint64{p2, p2h}, n1)
	n5 := Simd_f32x4_abs(n4)
	return n0[0], n0[1], n1[0], n1[1], n3[0], n3[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx639(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_sub([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_abs(n0)
	n2 := Simd_f32x4_pmin([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_pmin(n2, n1)
	n4 := Simd_f32x4_sub([2]uint64{p4, p4h}, [2]uint64{p1, p1h})
	n5 := Simd_f32x4_abs(n4)
	return n1[0], n1[1], n2[0], n2[1], n5[0], n5[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx640(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_pmin([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_pmin(n0, [2]uint64{p4, p4h})
	n2 := Simd_f32x4_neg([2]uint64{p2, p2h})
	n3 := Simd_f32x4_sub(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_abs(n3)
	n5 := Simd_f32x4_sub([2]uint64{p5, p5h}, [2]uint64{p2, p2h})
	n6 := Simd_f32x4_abs(n5)
	return n0[0], n0[1], n4[0], n4[1], n1[0], n1[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx641(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_pmin([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_pmin(n0, [2]uint64{p4, p4h})
	n2 := Simd_f32x4_sub([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_abs(n2)
	n4 := Simd_f32x4_sub([2]uint64{p5, p5h}, [2]uint64{p3, p3h})
	n5 := Simd_f32x4_abs(n4)
	return n0[0], n0[1], n3[0], n3[1], n1[0], n1[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx642(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_lt([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n2 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx643(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_lt([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx644(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_sub([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_abs(n0)
	n2 := Simd_f32x4_pmin([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_f32x4_lt(n1, n2)
	n4 := Simd_i8x16_shuffle(n3, [2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx645(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_lt([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_v128_and(n1, [2]uint64{p4, p4h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx646(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_m64_v128_load_rng(m, s0+32, 0, -32, 48)
	n2 := Simd_f32x4_sub(n0, n1)
	n3 := Simd_f32x4_abs(n2)
	n4 := Simd_f32x4_sub([2]uint64{p0, p0h}, n1)
	n5 := Simd_f32x4_abs(n4)
	return n1[0], n1[1], n5[0], n5[1], n0[0], n0[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx647(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_m64_v128_load_nc(m, s0, 0)
	n2 := Simd_f32x4_sub([2]uint64{p0, p0h}, n1)
	n3 := Simd_f32x4_abs(n2)
	n4 := Simd_f32x4_sub([2]uint64{p2, p2h}, n1)
	n5 := Simd_f32x4_abs(n4)
	return n0[0], n0[1], n1[0], n1[1], n3[0], n3[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx648(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_v128_and(n1, [2]uint64{p2, p2h})
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p1, p1h}, n4)
	return n1[0], n1[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx649(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p1, p1h}, 4)
	n1 := Simd_i32x4_add(n0, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx650(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_add([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	n5 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n6 := Simd_i8x16_add(n5, [2]uint64{p2, p2h})
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	_ = Simd_m64_v128_store(m, s0, s1, n4)
	_ = Simd_m64_v128_store(m, s2, s1, n10)
	n13 := Simd_m64_v128_load32_zero(m, s3+6, s1)
	return n4[0], n4[1], n10[0], n10[1], n13[0], n13[1]
}

//go:noinline
func Simd_p_fx651(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_add([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	n5 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n6 := Simd_i8x16_add(n5, [2]uint64{p2, p2h})
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	_ = Simd_m64_v128_store(m, s0+80, s1, n4)
	_ = Simd_m64_v128_store(m, s0+16, s1, n10)
	n13 := Simd_m64_v128_load32_zero(m, s2+10, s1)
	return n4[0], n4[1], n10[0], n10[1], n13[0], n13[1]
}

//go:noinline
func Simd_p_fx652(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_add([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	n5 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n6 := Simd_i8x16_add(n5, [2]uint64{p2, p2h})
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	_ = Simd_m64_v128_store(m, s0+96, s1, n4)
	_ = Simd_m64_v128_store(m, s0+32, s1, n10)
	n13 := Simd_m64_v128_load32_zero(m, s2+14, s1)
	return n4[0], n4[1], n10[0], n10[1], n13[0], n13[1]
}

//go:noinline
func Simd_p_fx653(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_add([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	n5 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n6 := Simd_i8x16_add(n5, [2]uint64{p2, p2h})
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	_ = Simd_m64_v128_store(m, s0+112, s1, n4)
	_ = Simd_m64_v128_store(m, s0+48, s1, n10)
	return n4[0], n4[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx654(m *Module, s0 int64, s1 int64, f0 float32, f1 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_splat(f1)
	n2 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_shr_u(n3, int32(s0))
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_f32x4_convert_i32x4_s(n5)
	n7 := Simd_f32x4_mul(n0, n6)
	n8 := Simd_f32x4_add(n7, n1)
	n9 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p2, p2h})
	n10 := Simd_i32x4_extend_low_i16x8_u(n9)
	n11 := Simd_i32x4_shr_u(n10, int32(s0))
	n12 := Simd_v128_and(n11, [2]uint64{p1, p1h})
	n13 := Simd_f32x4_convert_i32x4_s(n12)
	n14 := Simd_f32x4_mul(n0, n13)
	n15 := Simd_f32x4_add(n14, n1)
	n16 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p3, p3h})
	n17 := Simd_i32x4_extend_low_i16x8_u(n16)
	n18 := Simd_i32x4_shr_u(n17, int32(s0))
	n19 := Simd_v128_and(n18, [2]uint64{p1, p1h})
	n20 := Simd_f32x4_convert_i32x4_s(n19)
	n21 := Simd_f32x4_mul(n0, n20)
	n22 := Simd_f32x4_add(n21, n1)
	n23 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n24 := Simd_i32x4_extend_low_i16x8_u(n23)
	n25 := Simd_i32x4_shr_u(n24, int32(s0))
	n26 := Simd_v128_and(n25, [2]uint64{p1, p1h})
	n27 := Simd_f32x4_convert_i32x4_s(n26)
	n28 := Simd_f32x4_mul(n0, n27)
	n29 := Simd_f32x4_add(n28, n1)
	_ = Simd_m64_v128_store(m, s1+48, 0, n8)
	_ = Simd_m64_v128_store(m, s1+32, 0, n15)
	_ = Simd_m64_v128_store(m, s1+16, 0, n22)
	_ = Simd_m64_v128_store(m, s1, 0, n29)
	return
}

//go:noinline
func Simd_p_fx655(m *Module, s0 int64, s1 int64, s2 int64, f0 float32, f1 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_splat(f1)
	n2 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_shr_u(n3, int32(s0))
	n5 := Simd_v128_and(n4, [2]uint64{p1, p1h})
	n6 := Simd_f32x4_convert_i32x4_s(n5)
	n7 := Simd_f32x4_mul(n0, n6)
	n8 := Simd_f32x4_add(n7, n1)
	n9 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p2, p2h})
	n10 := Simd_i32x4_extend_low_i16x8_u(n9)
	n11 := Simd_i32x4_shr_u(n10, int32(s0))
	n12 := Simd_v128_and(n11, [2]uint64{p1, p1h})
	n13 := Simd_f32x4_convert_i32x4_s(n12)
	n14 := Simd_f32x4_mul(n0, n13)
	n15 := Simd_f32x4_add(n14, n1)
	n16 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p3, p3h})
	n17 := Simd_i32x4_extend_low_i16x8_u(n16)
	n18 := Simd_i32x4_shr_u(n17, int32(s0))
	n19 := Simd_v128_and(n18, [2]uint64{p1, p1h})
	n20 := Simd_f32x4_convert_i32x4_s(n19)
	n21 := Simd_f32x4_mul(n0, n20)
	n22 := Simd_f32x4_add(n21, n1)
	n23 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n24 := Simd_i32x4_extend_low_i16x8_u(n23)
	n25 := Simd_i32x4_shr_u(n24, int32(s0))
	n26 := Simd_v128_and(n25, [2]uint64{p1, p1h})
	n27 := Simd_f32x4_convert_i32x4_s(n26)
	n28 := Simd_f32x4_mul(n0, n27)
	n29 := Simd_f32x4_add(n28, n1)
	_ = Simd_m64_v128_store(m, s1+112, 0, n8)
	_ = Simd_m64_v128_store(m, s1+96, 0, n15)
	_ = Simd_m64_v128_store(m, s1+80, 0, n22)
	_ = Simd_m64_v128_store(m, s2, 0, n29)
	return
}

//go:noinline
func Simd_p_fx656(m *Module, s0 int64, s1 int64, s2 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_i8x16_splat(int32(s0))
	n2 := Simd_v128_and([2]uint64{p0, p0h}, n1)
	n3 := Simd_i8x16_eq(n2, [2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_low_i8x16_s(n3)
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_v128_and(n5, [2]uint64{p2, p2h})
	n7 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p3, p3h})
	n8 := Simd_i32x4_extend_low_i16x8_u(n7)
	n9 := Simd_i32x4_shr_u(n8, int32(s1))
	n10 := Simd_v128_and(n9, [2]uint64{p4, p4h})
	n11 := Simd_v128_or(n6, n10)
	n12 := Simd_f32x4_convert_i32x4_s(n11)
	n13 := Simd_f32x4_mul(n0, n12)
	_ = Simd_m64_v128_store(m, s2+48, 0, n13)
	return n0[0], n0[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx657(m *Module, s0 int32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p3, p3h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_v128_and(n3, [2]uint64{p4, p4h})
	n5 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_i32x4_shr_u(n6, s0)
	n8 := Simd_v128_and(n7, [2]uint64{p6, p6h})
	n9 := Simd_v128_or(n4, n8)
	n10 := Simd_f32x4_convert_i32x4_s(n9)
	n11 := Simd_f32x4_mul([2]uint64{p0, p0h}, n10)
	return n11[0], n11[1]
}

//go:noinline
func Simd_p_fx658(m *Module, s0 int64, s1 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_v128_and(n4, [2]uint64{p3, p3h})
	n6 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p4, p4h})
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_i32x4_shr_u(n7, int32(s0))
	n9 := Simd_v128_and(n8, [2]uint64{p5, p5h})
	n10 := Simd_v128_or(n5, n9)
	n11 := Simd_f32x4_convert_i32x4_s(n10)
	n12 := Simd_f32x4_mul(n0, n11)
	_ = Simd_m64_v128_store(m, s1+112, 0, n12)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx659(m *Module, s0 int64, s1 int64, f0 float32) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n2 := Simd_f32x4_abs(n1)
	n3 := Simd_f32x4_add(n0, n2)
	n4 := Simd_m64_v128_load_nc(m, s0, 16)
	n5 := Simd_f32x4_abs(n4)
	n6 := Simd_f32x4_add(n0, n5)
	_ = Simd_m64_v128_store(m, s1, 80, n6)
	n8 := Simd_m64_v128_load(m, s0, 32)
	n9 := Simd_f32x4_abs(n8)
	n10 := Simd_f32x4_add(n0, n9)
	_ = Simd_m64_v128_store(m, s1, 96, n10)
	n12 := Simd_m64_v128_load(m, s0, 48)
	n13 := Simd_f32x4_abs(n12)
	n14 := Simd_f32x4_add(n0, n13)
	_ = Simd_m64_v128_store(m, s1, 112, n14)
	n16 := Simd_m64_v128_load(m, s0, 64)
	n17 := Simd_f32x4_abs(n16)
	n18 := Simd_f32x4_add(n0, n17)
	_ = Simd_m64_v128_store(m, s1, 128, n18)
	n20 := Simd_m64_v128_load(m, s0, 80)
	n21 := Simd_f32x4_abs(n20)
	n22 := Simd_f32x4_add(n0, n21)
	_ = Simd_m64_v128_store(m, s1, 144, n22)
	n24 := Simd_m64_v128_load(m, s0, 96)
	n25 := Simd_f32x4_abs(n24)
	n26 := Simd_f32x4_add(n0, n25)
	_ = Simd_m64_v128_store(m, s1, 160, n26)
	n28 := Simd_m64_v128_load(m, s0, 112)
	n29 := Simd_f32x4_abs(n28)
	n30 := Simd_f32x4_add(n0, n29)
	_ = Simd_m64_v128_store(m, s1, 176, n30)
	_ = Simd_m64_v128_store(m, s1, 64, n3)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx660(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 208)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	n2 := Simd_m64_v128_load(m, s0, 192)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx661(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_div([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_shl(n0, 1)
	n2 := Simd_i32x4_max_u(n1, [2]uint64{p4, p4h})
	n3 := Simd_i32x4_shr_u(n2, 1)
	n4 := Simd_v128_and(n3, [2]uint64{p5, p5h})
	n5 := Simd_i32x4_add(n4, [2]uint64{p6, p6h})
	n6 := Simd_f32x4_abs(n0)
	n7 := Simd_f32x4_mul(n6, [2]uint64{p2, p2h})
	n8 := Simd_f32x4_mul(n7, [2]uint64{p3, p3h})
	n9 := Simd_f32x4_add(n8, n5)
	return n0[0], n0[1], n1[0], n1[1], n9[0], n9[1]
}

//go:noinline
func Simd_p_fx662(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p0, p0h}, 17)
	n1 := Simd_i32x4_lt_u(n0, [2]uint64{p1, p1h})
	return n0[0], n0[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx663(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_div(n0, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and(n2, [2]uint64{p4, p4h})
	n4 := Simd_i32x4_max_u(n3, [2]uint64{p5, p5h})
	n5 := Simd_i32x4_sub(n4, [2]uint64{p5, p5h})
	n6 := Simd_i32x4_min_u(n5, [2]uint64{p6, p6h})
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx664(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_nc(m, s0, 224)
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s1, 16, n1)
	n3 := Simd_m64_v128_load_rng(m, s0, 272, 240, 48)
	n4 := Simd_m64_v128_load_nc(m, s0, 240)
	return n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx665(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 32, n0)
	n2 := Simd_m64_v128_load_rng(m, s1, 320, 288, 48)
	n3 := Simd_m64_v128_load_nc(m, s1, 288)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx666(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 48, n0)
	n2 := Simd_m64_v128_load_rng(m, s1, 336, 304, 48)
	n3 := Simd_m64_v128_load_nc(m, s1, 304)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx667(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 64, n0)
	n2 := Simd_m64_v128_load_rng(m, s1, 384, 352, 48)
	n3 := Simd_m64_v128_load_nc(m, s1, 352)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx668(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 80, n0)
	n2 := Simd_m64_v128_load_rng(m, s1, 400, 368, 48)
	n3 := Simd_m64_v128_load_nc(m, s1, 368)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx669(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 96, n0)
	n2 := Simd_m64_v128_load_rng(m, s1, 448, 416, 48)
	n3 := Simd_m64_v128_load_nc(m, s1, 416)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx670(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 112, n0)
	n2 := Simd_m64_v128_load_rng(m, s1, 464, 432, 48)
	n3 := Simd_m64_v128_load_nc(m, s1, 432)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx671(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 128, n0)
	return
}

//go:noinline
func Simd_p_fx672(m *Module, s0 int64, s1 int64, f0 float32) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n2 := Simd_f32x4_abs(n1)
	n3 := Simd_f32x4_add(n0, n2)
	n4 := Simd_m64_v128_load_nc(m, s0, 16)
	n5 := Simd_f32x4_abs(n4)
	n6 := Simd_f32x4_add(n0, n5)
	_ = Simd_m64_v128_store(m, s1, 48, n6)
	n8 := Simd_m64_v128_load(m, s0, 32)
	n9 := Simd_f32x4_abs(n8)
	n10 := Simd_f32x4_add(n0, n9)
	_ = Simd_m64_v128_store(m, s1, 64, n10)
	n12 := Simd_m64_v128_load(m, s0, 48)
	n13 := Simd_f32x4_abs(n12)
	n14 := Simd_f32x4_add(n0, n13)
	_ = Simd_m64_v128_store(m, s1, 80, n14)
	n16 := Simd_m64_v128_load(m, s0, 64)
	n17 := Simd_f32x4_abs(n16)
	n18 := Simd_f32x4_add(n0, n17)
	_ = Simd_m64_v128_store(m, s1, 96, n18)
	n20 := Simd_m64_v128_load(m, s0, 80)
	n21 := Simd_f32x4_abs(n20)
	n22 := Simd_f32x4_add(n0, n21)
	_ = Simd_m64_v128_store(m, s1, 112, n22)
	n24 := Simd_m64_v128_load(m, s0, 96)
	n25 := Simd_f32x4_abs(n24)
	n26 := Simd_f32x4_add(n0, n25)
	_ = Simd_m64_v128_store(m, s1, 128, n26)
	n28 := Simd_m64_v128_load(m, s0, 112)
	n29 := Simd_f32x4_abs(n28)
	n30 := Simd_f32x4_add(n0, n29)
	_ = Simd_m64_v128_store(m, s1, 144, n30)
	_ = Simd_m64_v128_store(m, s1, 32, n3)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx673(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, f0 float32, f1 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_i8x16_splat(int32(s1))
	n2 := Simd_f32x4_splat(f1)
	n3 := Simd_m64_v128_load32_zero(m, s0, 16)
	n4 := Simd_v128_and(n3, n1)
	n5 := Simd_i8x16_eq(n4, [2]uint64{p0, p0h})
	n6 := Simd_i16x8_extend_low_i8x16_s(n5)
	n7 := Simd_i32x4_extend_low_i16x8_s(n6)
	n8 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n7)
	n9 := Simd_m64_v128_load32_zero(m, s2, 0)
	n10 := Simd_v128_and(n9, [2]uint64{p2, p2h})
	n11 := Simd_i16x8_extend_low_i8x16_u(n10)
	n12 := Simd_i32x4_extend_low_i16x8_u(n11)
	n13 := Simd_v128_or(n8, n12)
	n14 := Simd_f32x4_convert_i32x4_s(n13)
	n15 := Simd_f32x4_mul(n0, n14)
	n16 := Simd_f32x4_add(n15, n2)
	_ = Simd_m64_v128_store(m, s3, 0, n16)
	n18 := Simd_m64_v128_load32_zero(m, s0, 20)
	n19 := Simd_v128_and(n18, n1)
	n20 := Simd_i8x16_eq(n19, [2]uint64{p0, p0h})
	n21 := Simd_i16x8_extend_low_i8x16_s(n20)
	n22 := Simd_i32x4_extend_low_i16x8_s(n21)
	n23 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n22)
	n24 := Simd_m64_v128_load32_zero(m, s2, 4)
	n25 := Simd_v128_and(n24, [2]uint64{p2, p2h})
	n26 := Simd_i16x8_extend_low_i8x16_u(n25)
	n27 := Simd_i32x4_extend_low_i16x8_u(n26)
	n28 := Simd_v128_or(n23, n27)
	n29 := Simd_f32x4_convert_i32x4_s(n28)
	n30 := Simd_f32x4_mul(n0, n29)
	n31 := Simd_f32x4_add(n30, n2)
	_ = Simd_m64_v128_store(m, s3+16, 0, n31)
	n33 := Simd_m64_v128_load32_zero(m, s0, 24)
	n34 := Simd_v128_and(n33, n1)
	n35 := Simd_i8x16_eq(n34, [2]uint64{p0, p0h})
	n36 := Simd_i16x8_extend_low_i8x16_s(n35)
	n37 := Simd_i32x4_extend_low_i16x8_s(n36)
	n38 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n37)
	n39 := Simd_m64_v128_load32_zero(m, s2, 8)
	n40 := Simd_v128_and(n39, [2]uint64{p2, p2h})
	n41 := Simd_i16x8_extend_low_i8x16_u(n40)
	n42 := Simd_i32x4_extend_low_i16x8_u(n41)
	n43 := Simd_v128_or(n38, n42)
	n44 := Simd_f32x4_convert_i32x4_s(n43)
	n45 := Simd_f32x4_mul(n0, n44)
	n46 := Simd_f32x4_add(n45, n2)
	_ = Simd_m64_v128_store(m, s3+32, 0, n46)
	n48 := Simd_m64_v128_load32_zero(m, s0, 28)
	n49 := Simd_v128_and(n48, n1)
	n50 := Simd_i8x16_eq(n49, [2]uint64{p0, p0h})
	n51 := Simd_i16x8_extend_low_i8x16_s(n50)
	n52 := Simd_i32x4_extend_low_i16x8_s(n51)
	n53 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n52)
	n54 := Simd_m64_v128_load32_zero(m, s2, 12)
	n55 := Simd_v128_and(n54, [2]uint64{p2, p2h})
	n56 := Simd_i16x8_extend_low_i8x16_u(n55)
	n57 := Simd_i32x4_extend_low_i16x8_u(n56)
	n58 := Simd_v128_or(n53, n57)
	n59 := Simd_f32x4_convert_i32x4_s(n58)
	n60 := Simd_f32x4_mul(n0, n59)
	n61 := Simd_f32x4_add(n60, n2)
	_ = Simd_m64_v128_store(m, s3+48, 0, n61)
	n63 := Simd_m64_v128_load32_zero(m, s0, 32)
	n64 := Simd_v128_and(n63, n1)
	n65 := Simd_i8x16_eq(n64, [2]uint64{p0, p0h})
	n66 := Simd_i16x8_extend_low_i8x16_s(n65)
	n67 := Simd_i32x4_extend_low_i16x8_s(n66)
	n68 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n67)
	n69 := Simd_m64_v128_load32_zero(m, s2, 16)
	n70 := Simd_v128_and(n69, [2]uint64{p2, p2h})
	n71 := Simd_i16x8_extend_low_i8x16_u(n70)
	n72 := Simd_i32x4_extend_low_i16x8_u(n71)
	n73 := Simd_v128_or(n68, n72)
	n74 := Simd_f32x4_convert_i32x4_s(n73)
	n75 := Simd_f32x4_mul(n0, n74)
	n76 := Simd_f32x4_add(n75, n2)
	_ = Simd_m64_v128_store(m, s4, 0, n76)
	n78 := Simd_m64_v128_load32_zero(m, s0, 36)
	n79 := Simd_v128_and(n78, n1)
	n80 := Simd_i8x16_eq(n79, [2]uint64{p0, p0h})
	n81 := Simd_i16x8_extend_low_i8x16_s(n80)
	n82 := Simd_i32x4_extend_low_i16x8_s(n81)
	n83 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n82)
	n84 := Simd_m64_v128_load32_zero(m, s2, 20)
	n85 := Simd_v128_and(n84, [2]uint64{p2, p2h})
	n86 := Simd_i16x8_extend_low_i8x16_u(n85)
	n87 := Simd_i32x4_extend_low_i16x8_u(n86)
	n88 := Simd_v128_or(n83, n87)
	n89 := Simd_f32x4_convert_i32x4_s(n88)
	n90 := Simd_f32x4_mul(n0, n89)
	n91 := Simd_f32x4_add(n90, n2)
	_ = Simd_m64_v128_store(m, s3+80, 0, n91)
	n93 := Simd_m64_v128_load32_zero(m, s0, 40)
	n94 := Simd_v128_and(n93, n1)
	n95 := Simd_i8x16_eq(n94, [2]uint64{p0, p0h})
	n96 := Simd_i16x8_extend_low_i8x16_s(n95)
	n97 := Simd_i32x4_extend_low_i16x8_s(n96)
	n98 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n97)
	n99 := Simd_m64_v128_load32_zero(m, s2, 24)
	n100 := Simd_v128_and(n99, [2]uint64{p2, p2h})
	n101 := Simd_i16x8_extend_low_i8x16_u(n100)
	n102 := Simd_i32x4_extend_low_i16x8_u(n101)
	n103 := Simd_v128_or(n98, n102)
	n104 := Simd_f32x4_convert_i32x4_s(n103)
	n105 := Simd_f32x4_mul(n0, n104)
	n106 := Simd_f32x4_add(n105, n2)
	_ = Simd_m64_v128_store(m, s3+96, 0, n106)
	n108 := Simd_m64_v128_load32_zero(m, s0, 44)
	n109 := Simd_v128_and(n108, n1)
	n110 := Simd_i8x16_eq(n109, [2]uint64{p0, p0h})
	n111 := Simd_i16x8_extend_low_i8x16_s(n110)
	n112 := Simd_i32x4_extend_low_i16x8_s(n111)
	n113 := Simd_v128_bitselect([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, n112)
	n114 := Simd_m64_v128_load32_zero(m, s2, 28)
	n115 := Simd_v128_and(n114, [2]uint64{p2, p2h})
	n116 := Simd_i16x8_extend_low_i8x16_u(n115)
	n117 := Simd_i32x4_extend_low_i16x8_u(n116)
	n118 := Simd_v128_or(n113, n117)
	n119 := Simd_f32x4_convert_i32x4_s(n118)
	n120 := Simd_f32x4_mul(n0, n119)
	n121 := Simd_f32x4_add(n120, n2)
	_ = Simd_m64_v128_store(m, s3+112, 0, n121)
	return
}

//go:noinline
func Simd_p_fx674(m *Module, s0 int64, s1 int64, s2 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p1, p1h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_v128_bitselect([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, n4)
	n6 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_v128_or(n5, n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	n11 := Simd_f32x4_add(n10, n0)
	_ = Simd_m64_v128_store(m, s0, 0, n11)
	n13 := Simd_m64_v128_load32_zero(m, s1, 20)
	n14 := Simd_m64_v128_load32_zero(m, s2, 4)
	return n0[0], n0[1], n13[0], n13[1], n14[0], n14[1]
}

//go:noinline
func Simd_p_fx675(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_v128_bitselect([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, n3)
	n5 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_v128_or(n4, n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_add(n9, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+16, 0, n10)
	return
}

//go:noinline
func Simd_p_fx676(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_v128_bitselect([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, n3)
	n5 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_v128_or(n4, n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_add(n9, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+32, 0, n10)
	return
}

//go:noinline
func Simd_p_fx677(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_v128_bitselect([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, n3)
	n5 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_v128_or(n4, n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_add(n9, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+48, 0, n10)
	return
}

//go:noinline
func Simd_p_fx678(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_v128_bitselect([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, n3)
	n5 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_v128_or(n4, n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_add(n9, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0, 0, n10)
	return
}

//go:noinline
func Simd_p_fx679(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_v128_bitselect([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, n3)
	n5 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_v128_or(n4, n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_add(n9, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+80, 0, n10)
	return
}

//go:noinline
func Simd_p_fx680(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_v128_bitselect([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, n3)
	n5 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_v128_or(n4, n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_add(n9, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+96, 0, n10)
	return
}

//go:noinline
func Simd_p_fx681(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_v128_bitselect([2]uint64{p1, p1h}, [2]uint64{p2, p2h}, n3)
	n5 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p5, p5h})
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_v128_or(n4, n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_add(n9, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+112, 0, n10)
	return
}

//go:noinline
func Simd_p_fx682(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_or([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_m64_v128_load32_zero(m, s0+128, 0)
	n6 := Simd_v128_and(n5, [2]uint64{p4, p4h})
	n7 := Simd_i8x16_add(n4, n6)
	n8 := Simd_i16x8_extend_low_i8x16_s(n7)
	n9 := Simd_i32x4_extend_low_i16x8_s(n8)
	n10 := Simd_f32x4_convert_i32x4_s(n9)
	n11 := Simd_f32x4_mul(n3, n10)
	_ = Simd_m64_v128_store(m, s1+256, 0, n11)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx683(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_v128_or(n4, [2]uint64{p4, p4h})
	n6 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n7 := Simd_i8x16_add(n5, n6)
	n8 := Simd_i16x8_extend_low_i8x16_s(n7)
	n9 := Simd_i32x4_extend_low_i16x8_s(n8)
	n10 := Simd_f32x4_convert_i32x4_s(n9)
	n11 := Simd_f32x4_mul(n3, n10)
	_ = Simd_m64_v128_store(m, s0, 0, n11)
	return
}

//go:noinline
func Simd_p_fx684(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_m64_v128_load32_zero(m, s0+32, 0)
	n6 := Simd_v128_and(n5, [2]uint64{p4, p4h})
	n7 := Simd_v128_or(n4, n6)
	n8 := Simd_i8x16_add(n7, [2]uint64{p5, p5h})
	n9 := Simd_i16x8_extend_low_i8x16_s(n8)
	n10 := Simd_i32x4_extend_low_i16x8_s(n9)
	n11 := Simd_f32x4_convert_i32x4_s(n10)
	n12 := Simd_f32x4_mul(n3, n11)
	_ = Simd_m64_v128_store(m, s1+128, 0, n12)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx685(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_v128_or(n4, [2]uint64{p4, p4h})
	n6 := Simd_i8x16_add(n5, [2]uint64{p5, p5h})
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul(n3, n9)
	_ = Simd_m64_v128_store(m, s0+384, 0, n10)
	return
}

//go:noinline
func Simd_p_fx686(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_or([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n6 := Simd_i8x16_add(n4, n5)
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul(n3, n9)
	_ = Simd_m64_v128_store(m, s0, 0, n10)
	return
}

//go:noinline
func Simd_p_fx687(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_m64_v128_load32_zero(m, s0+96, 0)
	n6 := Simd_v128_and(n5, [2]uint64{p4, p4h})
	n7 := Simd_v128_or(n4, n6)
	n8 := Simd_i8x16_add(n7, [2]uint64{p5, p5h})
	n9 := Simd_i16x8_extend_low_i8x16_s(n8)
	n10 := Simd_i32x4_extend_low_i16x8_s(n9)
	n11 := Simd_f32x4_convert_i32x4_s(n10)
	n12 := Simd_f32x4_mul(n3, n11)
	_ = Simd_m64_v128_store(m, s1, 0, n12)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx688(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_v128_or(n4, [2]uint64{p4, p4h})
	n6 := Simd_i8x16_add(n5, [2]uint64{p5, p5h})
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i32x4_extend_low_i16x8_s(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul(n3, n9)
	_ = Simd_m64_v128_store(m, s0, 0, n10)
	return
}

//go:noinline
func Simd_p_fx689(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx690(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+128, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx691(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+256, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx692(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+384, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx693(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+512, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx694(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extmul_low_i8x16_u(n0, [2]uint64{p3, p3h})
	n2 := Simd_i16x8_extmul_high_i8x16_u(n0, [2]uint64{p3, p3h})
	n3 := Simd_i8x16_shuffle(n1, n2, [2]uint64{p4, p4h})
	n4 := Simd_i8x16_shuffle([2]uint64{p5, p5h}, [2]uint64{p6, p6h}, [2]uint64{p2, p2h})
	n5 := Simd_i8x16_add(n3, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx695(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extmul_low_i8x16_u([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extmul_high_i8x16_u([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p2, p2h})
	n3 := Simd_i8x16_shuffle([2]uint64{p3, p3h}, [2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_i8x16_add(n2, n3)
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx696(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extmul_low_i8x16_u([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extmul_high_i8x16_u([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p2, p2h})
	n3 := Simd_i8x16_shuffle([2]uint64{p3, p3h}, [2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_i8x16_add(n2, n3)
	n5 := Simd_i16x8_extmul_low_i8x16_u(n4, [2]uint64{p1, p1h})
	n6 := Simd_i8x16_shuffle(n5, [2]uint64{p1, p1h}, [2]uint64{1012195045828461056, 0})
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx697(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+640, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx698(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+704, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx699(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+768, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx700(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+832, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx701(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+896, 0)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx702(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i8x16_add(n0, [2]uint64{p3, p3h})
	n2 := Simd_v128_and(n1, [2]uint64{p4, p4h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx703(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx704(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx705(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+128, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 4)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx706(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+144, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 8)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx707(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+160, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 12)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx708(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+176, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 16)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx709(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+192, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 20)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx710(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+208, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 24)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx711(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+224, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 28)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx712(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+240, 0, n5)
	return
}

//go:noinline
func Simd_p_fx713(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 4)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx714(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+16, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 8)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx715(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+32, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 12)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx716(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+48, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 16)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx717(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 20)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx718(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+80, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 24)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx719(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+96, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1, 28)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx720(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+112, 0, n5)
	return
}

//go:noinline
func Simd_p_fx721(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_f32x4_neg(n3)
	n5 := Simd_i8x16_splat(int32(s0))
	n6 := Simd_v128_and(n5, [2]uint64{p2, p2h})
	n7 := Simd_i8x16_eq(n6, [2]uint64{p3, p3h})
	n8 := Simd_i16x8_extend_low_i8x16_s(n7)
	n9 := Simd_i32x4_extend_low_i16x8_s(n8)
	n10 := Simd_v128_bitselect(n3, n4, n9)
	_ = Simd_m64_v128_store(m, s1+96, 0, n10)
	return
}

//go:noinline
func Simd_p_fx722(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_f32x4_neg(n3)
	n5 := Simd_i8x16_splat(int32(s0))
	n6 := Simd_v128_and(n5, [2]uint64{p2, p2h})
	n7 := Simd_i8x16_eq(n6, [2]uint64{p3, p3h})
	n8 := Simd_i16x8_extend_low_i8x16_s(n7)
	n9 := Simd_i32x4_extend_low_i16x8_s(n8)
	n10 := Simd_v128_bitselect(n3, n4, n9)
	_ = Simd_m64_v128_store(m, s1, 0, n10)
	return
}

//go:noinline
func Simd_p_fx723(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_f32x4_neg(n3)
	n5 := Simd_i8x16_splat(int32(s0))
	n6 := Simd_v128_and(n5, [2]uint64{p2, p2h})
	n7 := Simd_i8x16_eq(n6, [2]uint64{p3, p3h})
	n8 := Simd_i16x8_extend_low_i8x16_s(n7)
	n9 := Simd_i32x4_extend_low_i16x8_s(n8)
	n10 := Simd_v128_bitselect(n3, n4, n9)
	_ = Simd_m64_v128_store(m, s1+32, 0, n10)
	return
}

//go:noinline
func Simd_p_fx724(m *Module, s0 int64, s1 int64, s2 int64, f0 float32, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_i8x16_splat(int32(s1))
	n2 := Simd_v128_and(n1, [2]uint64{p0, p0h})
	n3 := Simd_i8x16_eq(n2, [2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_low_i8x16_s(n3)
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_m64_v128_load32_zero(m, s0, 8550096)
	n7 := Simd_i16x8_extend_low_i8x16_u(n6)
	n8 := Simd_i32x4_extend_low_i16x8_u(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul(n0, n9)
	n11 := Simd_f32x4_neg(n10)
	n12 := Simd_v128_bitselect(n10, n11, n5)
	_ = Simd_m64_v128_store(m, s2, 0, n12)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx725(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i8x16_splat(int32(s1))
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_m64_v128_load32_zero(m, s0, 8550096)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_neg(n9)
	n11 := Simd_v128_bitselect(n9, n10, n4)
	_ = Simd_m64_v128_store(m, s2+32, 0, n11)
	return
}

//go:noinline
func Simd_p_fx726(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i8x16_splat(int32(s1))
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_m64_v128_load32_zero(m, s0, 8550096)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_neg(n9)
	n11 := Simd_v128_bitselect(n9, n10, n4)
	_ = Simd_m64_v128_store(m, s2, 0, n11)
	return
}

//go:noinline
func Simd_p_fx727(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i8x16_splat(int32(s1))
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_m64_v128_load32_zero(m, s0, 8550096)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_neg(n9)
	n11 := Simd_v128_bitselect(n9, n10, n4)
	_ = Simd_m64_v128_store(m, s2+96, 0, n11)
	return
}

//go:noinline
func Simd_p_fx728(m *Module, s0 int64, s1 int64, s2 int64, f0 float32, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_i8x16_splat(int32(s1))
	n2 := Simd_v128_and(n1, [2]uint64{p0, p0h})
	n3 := Simd_i8x16_eq(n2, [2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_low_i8x16_s(n3)
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_m64_v128_load32_zero(m, s0, 8551120)
	n7 := Simd_i16x8_extend_low_i8x16_u(n6)
	n8 := Simd_i32x4_extend_low_i16x8_u(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul(n0, n9)
	n11 := Simd_f32x4_neg(n10)
	n12 := Simd_v128_bitselect(n10, n11, n5)
	_ = Simd_m64_v128_store(m, s2, 0, n12)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx729(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i8x16_splat(int32(s1))
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_m64_v128_load32_zero(m, s0, 8551120)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_neg(n9)
	n11 := Simd_v128_bitselect(n9, n10, n4)
	_ = Simd_m64_v128_store(m, s2+32, 0, n11)
	return
}

//go:noinline
func Simd_p_fx730(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i8x16_splat(int32(s1))
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_m64_v128_load32_zero(m, s0, 8551120)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_neg(n9)
	n11 := Simd_v128_bitselect(n9, n10, n4)
	_ = Simd_m64_v128_store(m, s2, 0, n11)
	return
}

//go:noinline
func Simd_p_fx731(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i8x16_splat(int32(s1))
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_m64_v128_load32_zero(m, s0, 8551120)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_neg(n9)
	n11 := Simd_v128_bitselect(n9, n10, n4)
	_ = Simd_m64_v128_store(m, s2+96, 0, n11)
	return
}

//go:noinline
func Simd_p_fx732(m *Module, s0 int64, s1 int64, s2 int64, f0 float32, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_i8x16_splat(int32(s1))
	n2 := Simd_v128_and(n1, [2]uint64{p0, p0h})
	n3 := Simd_i8x16_eq(n2, [2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_low_i8x16_s(n3)
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_m64_v128_load32_zero(m, s0, 8551120)
	n7 := Simd_i16x8_extend_low_i8x16_u(n6)
	n8 := Simd_i32x4_extend_low_i16x8_u(n7)
	n9 := Simd_f32x4_convert_i32x4_s(n8)
	n10 := Simd_f32x4_mul(n0, n9)
	n11 := Simd_f32x4_neg(n10)
	n12 := Simd_v128_bitselect(n10, n11, n5)
	_ = Simd_m64_v128_store(m, s2+128, 0, n12)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx733(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i8x16_splat(int32(s1))
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_m64_v128_load32_zero(m, s0, 8551120)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_neg(n9)
	n11 := Simd_v128_bitselect(n9, n10, n4)
	_ = Simd_m64_v128_store(m, s2+160, 0, n11)
	return
}

//go:noinline
func Simd_p_fx734(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i8x16_splat(int32(s1))
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_m64_v128_load32_zero(m, s0, 8551120)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_neg(n9)
	n11 := Simd_v128_bitselect(n9, n10, n4)
	_ = Simd_m64_v128_store(m, s2+192, 0, n11)
	return
}

//go:noinline
func Simd_p_fx735(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i8x16_splat(int32(s1))
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p2, p2h})
	n3 := Simd_i16x8_extend_low_i8x16_s(n2)
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_m64_v128_load32_zero(m, s0, 8551120)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_f32x4_neg(n9)
	n11 := Simd_v128_bitselect(n9, n10, n4)
	_ = Simd_m64_v128_store(m, s2+224, 0, n11)
	return
}

//go:noinline
func Simd_p_fx736(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p1, p1h})
	n3 := Simd_v128_or(n2, [2]uint64{p2, p2h})
	n4 := Simd_i8x16_shuffle(n1, n3, [2]uint64{506097522914230528, 1663540288323457296})
	n5 := Simd_i16x8_shr_u([2]uint64{p4, p4h}, 5)
	n6 := Simd_v128_and(n5, [2]uint64{p5, p5h})
	n7 := Simd_v128_or(n6, [2]uint64{p6, p6h})
	n8 := Simd_i16x8_shr_u([2]uint64{p4, p4h}, 8)
	n9 := Simd_v128_and(n8, [2]uint64{p5, p5h})
	n10 := Simd_v128_or(n9, [2]uint64{p6, p6h})
	n11 := Simd_i8x16_narrow_i16x8_u(n7, n10)
	n12 := Simd_i8x16_shuffle(n4, n11, [2]uint64{2095595517207907332, 2240275862884060678})
	n13 := Simd_i8x16_shuffle(n4, n11, [2]uint64{1806234825855600640, 1950915171531753986})
	_ = Simd_m64_v128_store(m, s0, 16, n12)
	_ = Simd_m64_v128_store(m, s0, 0, n13)
	return n11[0], n11[1]
}

//go:noinline
func Simd_p_fx737(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 16)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_i16x8_add(n2, [2]uint64{p1, p1h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx738(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 24)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_i16x8_add(n1, [2]uint64{p1, p1h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx739(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_shl([2]uint64{p1, p1h}, 3)
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	n2 := Simd_i16x8_shl([2]uint64{p2, p2h}, 6)
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_i16x8_shl([2]uint64{p3, p3h}, 9)
	n5 := Simd_v128_or(n3, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx740(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, p0, p0h uint64) {
	n0 := Simd_i32x4_splat(int32(s1))
	n1 := Simd_i32x4_splat(int32(s2))
	n2 := Simd_i32x4_splat(int32(s3))
	n3 := Simd_i32x4_splat(int32(s4))
	n4 := Simd_m64_v128_load(m, s0, 0)
	n5 := Simd_i8x16_shuffle(n4, [2]uint64{p0, p0h}, [2]uint64{252380931, 0})
	n6 := Simd_i16x8_extend_low_i8x16_s(n5)
	n7 := Simd_i32x4_extend_low_i16x8_s(n6)
	n8 := Simd_i32x4_add(n0, n7)
	n9 := Simd_i32x4_mul(n8, n8)
	n10 := Simd_i8x16_shuffle(n4, [2]uint64{p0, p0h}, [2]uint64{235537922, 0})
	n11 := Simd_i16x8_extend_low_i8x16_s(n10)
	n12 := Simd_i32x4_extend_low_i16x8_s(n11)
	n13 := Simd_i32x4_add(n1, n12)
	n14 := Simd_i32x4_mul(n13, n13)
	n15 := Simd_i32x4_add(n9, n14)
	n16 := Simd_i8x16_shuffle(n4, [2]uint64{p0, p0h}, [2]uint64{218694913, 0})
	n17 := Simd_i16x8_extend_low_i8x16_s(n16)
	n18 := Simd_i32x4_extend_low_i16x8_s(n17)
	n19 := Simd_i32x4_add(n2, n18)
	n20 := Simd_i32x4_mul(n19, n19)
	n21 := Simd_i8x16_shuffle(n4, [2]uint64{p0, p0h}, [2]uint64{201851904, 0})
	n22 := Simd_i16x8_extend_low_i8x16_s(n21)
	n23 := Simd_i32x4_extend_low_i16x8_s(n22)
	n24 := Simd_i32x4_add(n3, n23)
	n25 := Simd_i32x4_mul(n24, n24)
	n26 := Simd_i32x4_add(n15, n20)
	n27 := Simd_i32x4_add(n26, n25)
	n28 := Simd_i8x16_shuffle(n27, [2]uint64{p0, p0h}, [2]uint64{1952900979473647880, 2242261670825954572})
	n29 := Simd_i8x16_shuffle(n27, [2]uint64{p0, p0h}, [2]uint64{1374179596769034496, 1663540288121341188})
	_ = Simd_m64_v128_store(m, s5, 16, n28)
	_ = Simd_m64_v128_store(m, s5, 0, n29)
	return
}

//go:noinline
func Simd_p_fx741(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_eq(n0, [2]uint64{p3, p3h})
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_v128_and(n2, [2]uint64{p4, p4h})
	n4 := Simd_i32x4_add([2]uint64{p2, p2h}, n3)
	n5 := Simd_i16x8_gt_u(n0, [2]uint64{p3, p3h})
	n6 := Simd_i32x4_extend_low_i16x8_u(n5)
	n7 := Simd_v128_and(n6, [2]uint64{p4, p4h})
	n8 := Simd_i32x4_add([2]uint64{p5, p5h}, n7)
	return n4[0], n4[1], n8[0], n8[1]
}

//go:noinline
func Simd_p_fx742(m *Module, s0 int64, s1 int64, f0 float32) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_mul(n1, n0)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx743(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_scalar_i32_add(s0, s1)
	n2 := Simd_m64_v128_load(m, n1, 0)
	n3 := Simd_f32x4_add(n0, n2)
	n4 := Simd_f32x4_sub(n0, n2)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	n6 := Simd_m64_scalar_i32_add(s0, s1)
	_ = Simd_m64_v128_store(m, n6, 0, n4)
	return
}

//go:noinline
func Simd_p_fx744(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_scalar_i32_add(s1, s2)
	n2 := Simd_m64_scalar_i32_add(n1, s3)
	n3 := Simd_m64_v128_load(m, n2, 0)
	n4 := Simd_f32x4_add(n0, n3)
	n5 := Simd_f32x4_sub(n0, n3)
	_ = Simd_m64_v128_store(m, s0, 0, n4)
	n7 := Simd_m64_scalar_i32_add(s1, s2)
	n8 := Simd_m64_scalar_i32_add(n7, s3)
	_ = Simd_m64_v128_store(m, n8, 0, n5)
	return
}

//go:noinline
func Simd_p_fx745(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_scalar_i32_add(s0, s1)
	n2 := Simd_m64_v128_load(m, n1, 0)
	n3 := Simd_f32x4_add(n0, n2)
	n4 := Simd_f32x4_sub(n0, n2)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	n6 := Simd_m64_scalar_i32_add(s0, s1)
	_ = Simd_m64_v128_store(m, n6, 0, n4)
	n8 := Simd_m64_v128_load(m, s0+16, 0)
	n9 := Simd_m64_scalar_i32_add(s0, s1)
	n10 := Simd_m64_scalar_i32_add(n9, 16)
	n11 := Simd_m64_v128_load(m, n10, 0)
	n12 := Simd_f32x4_add(n8, n11)
	n13 := Simd_f32x4_sub(n8, n11)
	_ = Simd_m64_v128_store(m, s0+16, 0, n12)
	n15 := Simd_m64_scalar_i32_add(s0, s1)
	n16 := Simd_m64_scalar_i32_add(n15, 16)
	_ = Simd_m64_v128_store(m, n16, 0, n13)
	return
}

//go:noinline
func Simd_p_fx746(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p0, p0h}, 16)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx747(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_div([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_mul(n1, [2]uint64{p2, p2h})
	n3 := Simd_i32x4_shl(n2, 1)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx748(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_abs([2]uint64{p0, p0h})
	n1 := Simd_f32x4_mul(n0, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_mul(n1, [2]uint64{p2, p2h})
	n3 := Simd_i32x4_max_u([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n4 := Simd_i32x4_shr_u(n3, 1)
	n5 := Simd_v128_and(n4, [2]uint64{p5, p5h})
	n6 := Simd_i32x4_add(n5, [2]uint64{p6, p6h})
	n7 := Simd_f32x4_add(n2, n6)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx749(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_mul(n1, [2]uint64{p0, p0h})
	n3 := Simd_f32x4_add(n2, [2]uint64{p3, p3h})
	n4 := Simd_f32x4_mul(n0, n3)
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx750(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_i32x4_shl(n2, 1)
	return n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx751(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p0, p0h}, 23)
	n1 := Simd_i32x4_add(n0, [2]uint64{p1, p1h})
	n2 := Simd_i32x4_add(n0, [2]uint64{p2, p2h})
	n3 := Simd_i32x4_add(n0, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0+9056952, 0, n1)
	_ = Simd_m64_v128_store(m, s0+9056936, 0, n2)
	_ = Simd_m64_v128_store(m, s0+9056920, 0, n3)
	return
}

//go:noinline
func Simd_p_fx752(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 40)
	n1 := Simd_m64_v128_load32_zero(m, s0, 36)
	n2 := Simd_m64_v128_load32_zero(m, s0, 32)
	n3 := Simd_m64_v128_load_rng(m, s1, 0, 0, 32)
	n4 := Simd_v128_and(n3, [2]uint64{p0, p0h})
	n5 := Simd_i16x8_extend_low_i8x16_u(n4)
	n6 := Simd_i16x8_add(n5, [2]uint64{p1, p1h})
	n7 := Simd_i16x8_extend_high_i8x16_u(n4)
	n8 := Simd_i16x8_add(n7, [2]uint64{p1, p1h})
	n9 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n10 := Simd_i16x8_extend_low_i8x16_s(n9)
	n11 := Simd_i16x8_mul(n6, n10)
	n12 := Simd_i32x4_extend_low_i16x8_s(n11)
	n13 := Simd_i16x8_extend_high_i8x16_s(n9)
	n14 := Simd_i16x8_mul(n8, n13)
	n15 := Simd_i32x4_extend_low_i16x8_s(n14)
	n16 := Simd_i32x4_extend_high_i16x8_s(n11)
	n17 := Simd_i32x4_extend_high_i16x8_s(n14)
	n18 := Simd_m64_v128_load_nc(m, s1, 16)
	n19 := Simd_v128_and(n18, [2]uint64{p0, p0h})
	n20 := Simd_i16x8_extend_low_i8x16_u(n19)
	n21 := Simd_i16x8_add(n20, [2]uint64{p1, p1h})
	n22 := Simd_i16x8_extend_high_i8x16_u(n19)
	n23 := Simd_i16x8_add(n22, [2]uint64{p1, p1h})
	n24 := Simd_m64_v128_load_nc(m, s0, 16)
	n25 := Simd_i16x8_extend_low_i8x16_s(n24)
	n26 := Simd_i16x8_mul(n21, n25)
	n27 := Simd_i32x4_extend_low_i16x8_s(n26)
	n28 := Simd_i16x8_extend_high_i8x16_s(n24)
	n29 := Simd_i16x8_mul(n23, n28)
	n30 := Simd_i32x4_extend_low_i16x8_s(n29)
	n31 := Simd_i32x4_add(n12, n27)
	n32 := Simd_i32x4_add(n31, n15)
	n33 := Simd_i32x4_add(n32, n30)
	n34 := Simd_i32x4_add(n33, n16)
	n35 := Simd_i32x4_extend_high_i16x8_s(n26)
	n36 := Simd_i32x4_extend_high_i16x8_s(n29)
	n37 := Simd_i32x4_add(n34, n35)
	n38 := Simd_i32x4_add(n37, n17)
	n39 := Simd_i32x4_add(n38, n36)
	n40 := Simd_i8x16_shuffle(n39, n39, [2]uint64{p2, p2h})
	n41 := Simd_i32x4_add(n39, n40)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n41[0], n41[1]
}

//go:noinline
func Simd_p_fx753(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p2, p2h})
	n4 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_i32x4_mul(n3, n5)
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx754(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p2, p2h})
	n4 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n5 := Simd_i32x4_extend_low_i16x8_s(n4)
	n6 := Simd_i32x4_mul(n3, n5)
	n7 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p1, p1h})
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	n10 := Simd_i32x4_add(n9, [2]uint64{p2, p2h})
	n11 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p5, p5h})
	n12 := Simd_i32x4_extend_low_i16x8_s(n11)
	n13 := Simd_i32x4_mul(n10, n12)
	n14 := Simd_i32x4_add(n13, [2]uint64{p6, p6h})
	n15 := Simd_i32x4_add(n6, n14)
	return n15[0], n15[1]
}

//go:noinline
func Simd_p_fx755(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_add(n1, [2]uint64{p1, p1h})
	n3 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_i32x4_mul(n2, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx756(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_extend_low_i16x8_u(n0)
	n2 := Simd_i32x4_add(n1, [2]uint64{p1, p1h})
	n3 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n4 := Simd_i32x4_extend_low_i16x8_s(n3)
	n5 := Simd_i32x4_mul(n2, n4)
	n6 := Simd_i8x16_shuffle([2]uint64{p3, p3h}, [2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n7 := Simd_i32x4_add([2]uint64{p3, p3h}, n6)
	n8 := Simd_i8x16_shuffle(n7, [2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n9 := Simd_i32x4_add(n5, n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx757(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{72339069014638592, 144680345659310337})
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p3, p3h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx758(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{72339069014638592, 144680345659310337})
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p3, p3h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx759(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_splat(int32(s1))
	n1 := Simd_m64_v128_load32_zero(m, s0, 0)
	n2 := Simd_i8x16_shuffle(n1, n0, [2]uint64{1112397582594, 0})
	n3 := Simd_i8x16_shuffle(n2, [2]uint64{p0, p0h}, [2]uint64{72057594037927936, 144679241852715265})
	n4 := Simd_v128_and(n3, [2]uint64{p1, p1h})
	n5 := Simd_i8x16_eq(n4, [2]uint64{p2, p2h})
	return n2[0], n2[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx760(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{361700864190317572, 434041037028459781})
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p3, p3h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx761(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{217020514202419714, 289360691335463683})
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_eq(n1, [2]uint64{p3, p3h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx762(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shr_u([2]uint64{p0, p0h}, 3)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx763(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shr_u([2]uint64{p0, p0h}, 2)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx764(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 0)
	n1 := Simd_m64_v128_load32_lane(m, s1, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s2, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s3, 0, 3, n2)
	n4 := Simd_m64_v128_load32_splat(m, s4, 0)
	n5 := Simd_f32x4_mul(n3, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx765(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_trunc_sat_f32x4_s([2]uint64{p0, p0h})
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_narrow_i32x4_u(n1, [2]uint64{p2, p2h})
	n3 := Simd_i8x16_narrow_i16x8_u(n2, [2]uint64{p3, p3h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx766(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_shr_u([2]uint64{p0, p0h}, 1)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx767(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+16, 0)
	n1 := Simd_m64_scalar_i32_shl(s2, 2)
	n2 := Simd_m64_scalar_i32_add(s1, n1)
	n3 := Simd_m64_v128_load32_splat(m, n2, 0)
	n4 := Simd_f32x4_mul(n0, n3)
	n5 := Simd_f32x4_add([2]uint64{p0, p0h}, n4)
	n6 := Simd_m64_v128_load(m, s0, 0)
	n7 := Simd_f32x4_mul(n6, n3)
	n8 := Simd_f32x4_add([2]uint64{p2, p2h}, n7)
	n9 := Simd_m64_v128_load32_splat(m, s1, 0)
	n10 := Simd_f32x4_mul(n6, n9)
	n11 := Simd_f32x4_add([2]uint64{p1, p1h}, n10)
	n12 := Simd_f32x4_mul(n0, n9)
	n13 := Simd_f32x4_add([2]uint64{p3, p3h}, n12)
	n14 := Simd_m64_scalar_i32_shl(s3, 2)
	n15 := Simd_m64_scalar_i32_add(s0, n14)
	n16 := Simd_m64_scalar_i32_add(n15, 16)
	n17 := Simd_m64_v128_load(m, n16, 0)
	n18 := Simd_m64_v128_load32_splat(m, s4, 0)
	n19 := Simd_f32x4_mul(n17, n18)
	n20 := Simd_f32x4_add(n5, n19)
	n21 := Simd_m64_scalar_i32_shl(s3, 2)
	n22 := Simd_m64_scalar_i32_add(s0, n21)
	n23 := Simd_m64_v128_load(m, n22, 0)
	n24 := Simd_f32x4_mul(n23, n18)
	n25 := Simd_f32x4_add(n8, n24)
	n26 := Simd_m64_v128_load32_splat(m, s1+4, 0)
	n27 := Simd_f32x4_mul(n23, n26)
	n28 := Simd_f32x4_add(n11, n27)
	n29 := Simd_f32x4_mul(n17, n26)
	n30 := Simd_f32x4_add(n13, n29)
	n31 := Simd_m64_scalar_i32_shl(s3, 2)
	n32 := Simd_m64_scalar_i32_add(s0, n31)
	n33 := Simd_m64_scalar_i32_shl(s3, 2)
	n34 := Simd_m64_scalar_i32_add(n32, n33)
	n35 := Simd_m64_scalar_i32_add(n34, 16)
	n36 := Simd_m64_v128_load(m, n35, 0)
	n37 := Simd_m64_v128_load32_splat(m, s5, 0)
	n38 := Simd_f32x4_mul(n36, n37)
	n39 := Simd_f32x4_add(n20, n38)
	n40 := Simd_m64_scalar_i32_shl(s3, 2)
	n41 := Simd_m64_scalar_i32_add(s0, n40)
	n42 := Simd_m64_scalar_i32_shl(s3, 2)
	n43 := Simd_m64_scalar_i32_add(n41, n42)
	n44 := Simd_m64_v128_load(m, n43, 0)
	n45 := Simd_f32x4_mul(n44, n37)
	n46 := Simd_f32x4_add(n25, n45)
	n47 := Simd_m64_v128_load32_splat(m, s1+8, 0)
	n48 := Simd_f32x4_mul(n44, n47)
	n49 := Simd_f32x4_add(n28, n48)
	n50 := Simd_f32x4_mul(n36, n47)
	n51 := Simd_f32x4_add(n30, n50)
	n52 := Simd_m64_scalar_i32_shl(s3, 2)
	n53 := Simd_m64_scalar_i32_add(s0, n52)
	n54 := Simd_m64_scalar_i32_shl(s3, 2)
	n55 := Simd_m64_scalar_i32_add(n53, n54)
	n56 := Simd_m64_scalar_i32_shl(s3, 2)
	n57 := Simd_m64_scalar_i32_add(n55, n56)
	n58 := Simd_m64_scalar_i32_add(n57, 16)
	n59 := Simd_m64_v128_load(m, n58, 0)
	n60 := Simd_m64_v128_load32_splat(m, s6, 0)
	n61 := Simd_f32x4_mul(n59, n60)
	n62 := Simd_f32x4_add(n39, n61)
	n63 := Simd_m64_scalar_i32_shl(s3, 2)
	n64 := Simd_m64_scalar_i32_add(s0, n63)
	n65 := Simd_m64_scalar_i32_shl(s3, 2)
	n66 := Simd_m64_scalar_i32_add(n64, n65)
	n67 := Simd_m64_scalar_i32_shl(s3, 2)
	n68 := Simd_m64_scalar_i32_add(n66, n67)
	n69 := Simd_m64_v128_load(m, n68, 0)
	n70 := Simd_f32x4_mul(n69, n60)
	n71 := Simd_f32x4_add(n46, n70)
	n72 := Simd_m64_v128_load32_splat(m, s1+12, 0)
	n73 := Simd_f32x4_mul(n69, n72)
	n74 := Simd_f32x4_add(n49, n73)
	n75 := Simd_f32x4_mul(n59, n72)
	n76 := Simd_f32x4_add(n51, n75)
	return n62[0], n62[1], n74[0], n74[1], n71[0], n71[1], n76[0], n76[1]
}

//go:noinline
func Simd_p_fx768(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+16, 0)
	n1 := Simd_m64_scalar_i32_add(s1, s2)
	n2 := Simd_m64_v128_load32_splat(m, n1, 0)
	n3 := Simd_f32x4_mul(n0, n2)
	n4 := Simd_f32x4_add([2]uint64{p0, p0h}, n3)
	n5 := Simd_m64_v128_load(m, s0, 0)
	n6 := Simd_f32x4_mul(n5, n2)
	n7 := Simd_f32x4_add([2]uint64{p2, p2h}, n6)
	n8 := Simd_m64_v128_load32_splat(m, s1, 0)
	n9 := Simd_f32x4_mul(n5, n8)
	n10 := Simd_f32x4_add([2]uint64{p1, p1h}, n9)
	n11 := Simd_f32x4_mul(n0, n8)
	n12 := Simd_f32x4_add([2]uint64{p3, p3h}, n11)
	n13 := Simd_m64_scalar_i32_add(s0, s3)
	n14 := Simd_m64_scalar_i32_add(n13, 16)
	n15 := Simd_m64_v128_load(m, n14, 0)
	n16 := Simd_m64_scalar_i32_add(s1, 4)
	n17 := Simd_m64_scalar_i32_add(n16, s2)
	n18 := Simd_m64_v128_load32_splat(m, n17, 0)
	n19 := Simd_f32x4_mul(n15, n18)
	n20 := Simd_f32x4_add(n4, n19)
	n21 := Simd_m64_scalar_i32_add(s0, s3)
	n22 := Simd_m64_v128_load(m, n21, 0)
	n23 := Simd_f32x4_mul(n22, n18)
	n24 := Simd_f32x4_add(n7, n23)
	n25 := Simd_m64_v128_load32_splat(m, s1+4, 0)
	n26 := Simd_f32x4_mul(n22, n25)
	n27 := Simd_f32x4_add(n10, n26)
	n28 := Simd_f32x4_mul(n15, n25)
	n29 := Simd_f32x4_add(n12, n28)
	n30 := Simd_m64_scalar_i32_add(s0, s3)
	n31 := Simd_m64_scalar_i32_add(n30, s3)
	n32 := Simd_m64_scalar_i32_add(n31, 16)
	n33 := Simd_m64_v128_load(m, n32, 0)
	n34 := Simd_m64_scalar_i32_add(s1, 8)
	n35 := Simd_m64_scalar_i32_add(n34, s2)
	n36 := Simd_m64_v128_load32_splat(m, n35, 0)
	n37 := Simd_f32x4_mul(n33, n36)
	n38 := Simd_f32x4_add(n20, n37)
	n39 := Simd_m64_scalar_i32_add(s0, s3)
	n40 := Simd_m64_scalar_i32_add(n39, s3)
	n41 := Simd_m64_v128_load(m, n40, 0)
	n42 := Simd_f32x4_mul(n41, n36)
	n43 := Simd_f32x4_add(n24, n42)
	n44 := Simd_m64_v128_load32_splat(m, s1+8, 0)
	n45 := Simd_f32x4_mul(n41, n44)
	n46 := Simd_f32x4_add(n27, n45)
	n47 := Simd_f32x4_mul(n33, n44)
	n48 := Simd_f32x4_add(n29, n47)
	n49 := Simd_m64_scalar_i32_add(s0, s3)
	n50 := Simd_m64_scalar_i32_add(n49, s3)
	n51 := Simd_m64_scalar_i32_add(n50, s3)
	n52 := Simd_m64_scalar_i32_add(n51, 16)
	n53 := Simd_m64_v128_load(m, n52, 0)
	n54 := Simd_m64_scalar_i32_add(s1, 12)
	n55 := Simd_m64_scalar_i32_add(n54, s2)
	n56 := Simd_m64_v128_load32_splat(m, n55, 0)
	n57 := Simd_f32x4_mul(n53, n56)
	n58 := Simd_f32x4_add(n38, n57)
	n59 := Simd_m64_scalar_i32_add(s0, s3)
	n60 := Simd_m64_scalar_i32_add(n59, s3)
	n61 := Simd_m64_scalar_i32_add(n60, s3)
	n62 := Simd_m64_v128_load(m, n61, 0)
	n63 := Simd_f32x4_mul(n62, n56)
	n64 := Simd_f32x4_add(n43, n63)
	n65 := Simd_m64_v128_load32_splat(m, s1+12, 0)
	n66 := Simd_f32x4_mul(n62, n65)
	n67 := Simd_f32x4_add(n46, n66)
	n68 := Simd_f32x4_mul(n53, n65)
	n69 := Simd_f32x4_add(n48, n68)
	return n58[0], n58[1], n67[0], n67[1], n64[0], n64[1], n69[0], n69[1]
}

//go:noinline
func Simd_p_fx769(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+16, s1)
	n1 := Simd_m64_scalar_i32_add(s2, s3)
	n2 := Simd_m64_v128_load32_splat(m, n1, s1)
	n3 := Simd_f32x4_mul(n0, n2)
	n4 := Simd_f32x4_add([2]uint64{p0, p0h}, n3)
	n5 := Simd_m64_v128_load(m, s0, s1)
	return n0[0], n0[1], n2[0], n2[1], n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx770(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p1, p1h}, [2]uint64{p3, p3h})
	n1 := Simd_f32x4_add([2]uint64{p2, p2h}, n0)
	n2 := Simd_m64_v128_load32_splat(m, s0, s1)
	n3 := Simd_f32x4_mul([2]uint64{p1, p1h}, n2)
	n4 := Simd_f32x4_add([2]uint64{p0, p0h}, n3)
	n5 := Simd_f32x4_mul([2]uint64{p5, p5h}, n2)
	n6 := Simd_f32x4_add([2]uint64{p4, p4h}, n5)
	return n2[0], n2[1], n4[0], n4[1], n1[0], n1[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx771(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load32_splat(m, s1, 0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	n4 := Simd_m64_scalar_i32_shl(s2, 2)
	n5 := Simd_m64_scalar_i32_add(s1, n4)
	n6 := Simd_m64_v128_load32_splat(m, n5, 0)
	n7 := Simd_f32x4_mul(n0, n6)
	n8 := Simd_f32x4_add([2]uint64{p1, p1h}, n7)
	n9 := Simd_m64_scalar_i32_shl(s3, 2)
	n10 := Simd_m64_scalar_i32_add(s0, n9)
	n11 := Simd_m64_v128_load(m, n10, 0)
	n12 := Simd_m64_v128_load32_splat(m, s1+4, 0)
	n13 := Simd_f32x4_mul(n11, n12)
	n14 := Simd_f32x4_add(n3, n13)
	n15 := Simd_m64_v128_load32_splat(m, s4, 0)
	n16 := Simd_f32x4_mul(n11, n15)
	n17 := Simd_f32x4_add(n8, n16)
	n18 := Simd_m64_scalar_i32_shl(s3, 2)
	n19 := Simd_m64_scalar_i32_add(s0, n18)
	n20 := Simd_m64_scalar_i32_shl(s3, 2)
	n21 := Simd_m64_scalar_i32_add(n19, n20)
	n22 := Simd_m64_v128_load(m, n21, 0)
	n23 := Simd_m64_v128_load32_splat(m, s1+8, 0)
	n24 := Simd_f32x4_mul(n22, n23)
	n25 := Simd_f32x4_add(n14, n24)
	n26 := Simd_m64_v128_load32_splat(m, s5, 0)
	n27 := Simd_f32x4_mul(n22, n26)
	n28 := Simd_f32x4_add(n17, n27)
	n29 := Simd_m64_scalar_i32_shl(s3, 2)
	n30 := Simd_m64_scalar_i32_add(s0, n29)
	n31 := Simd_m64_scalar_i32_shl(s3, 2)
	n32 := Simd_m64_scalar_i32_add(n30, n31)
	n33 := Simd_m64_scalar_i32_shl(s3, 2)
	n34 := Simd_m64_scalar_i32_add(n32, n33)
	n35 := Simd_m64_v128_load(m, n34, 0)
	n36 := Simd_m64_v128_load32_splat(m, s1+12, 0)
	n37 := Simd_f32x4_mul(n35, n36)
	n38 := Simd_f32x4_add(n25, n37)
	n39 := Simd_m64_v128_load32_splat(m, s6, 0)
	n40 := Simd_f32x4_mul(n35, n39)
	n41 := Simd_f32x4_add(n28, n40)
	return n38[0], n38[1], n41[0], n41[1]
}

//go:noinline
func Simd_p_fx772(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load32_splat(m, s1, 0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	n4 := Simd_m64_scalar_i32_add(s1, s2)
	n5 := Simd_m64_v128_load32_splat(m, n4, 0)
	n6 := Simd_f32x4_mul(n0, n5)
	n7 := Simd_f32x4_add([2]uint64{p1, p1h}, n6)
	n8 := Simd_m64_scalar_i32_add(s0, s3)
	n9 := Simd_m64_v128_load(m, n8, 0)
	n10 := Simd_m64_v128_load32_splat(m, s1+4, 0)
	n11 := Simd_f32x4_mul(n9, n10)
	n12 := Simd_f32x4_add(n3, n11)
	n13 := Simd_m64_scalar_i32_add(s1, 4)
	n14 := Simd_m64_scalar_i32_add(n13, s2)
	n15 := Simd_m64_v128_load32_splat(m, n14, 0)
	n16 := Simd_f32x4_mul(n9, n15)
	n17 := Simd_f32x4_add(n7, n16)
	n18 := Simd_m64_scalar_i32_add(s0, s3)
	n19 := Simd_m64_scalar_i32_add(n18, s3)
	n20 := Simd_m64_v128_load(m, n19, 0)
	n21 := Simd_m64_v128_load32_splat(m, s1+8, 0)
	n22 := Simd_f32x4_mul(n20, n21)
	n23 := Simd_f32x4_add(n12, n22)
	n24 := Simd_m64_scalar_i32_add(s1, 8)
	n25 := Simd_m64_scalar_i32_add(n24, s2)
	n26 := Simd_m64_v128_load32_splat(m, n25, 0)
	n27 := Simd_f32x4_mul(n20, n26)
	n28 := Simd_f32x4_add(n17, n27)
	n29 := Simd_m64_scalar_i32_add(s0, s3)
	n30 := Simd_m64_scalar_i32_add(n29, s3)
	n31 := Simd_m64_scalar_i32_add(n30, s3)
	n32 := Simd_m64_v128_load(m, n31, 0)
	n33 := Simd_m64_v128_load32_splat(m, s1+12, 0)
	n34 := Simd_f32x4_mul(n32, n33)
	n35 := Simd_f32x4_add(n23, n34)
	n36 := Simd_m64_scalar_i32_add(s1, 12)
	n37 := Simd_m64_scalar_i32_add(n36, s2)
	n38 := Simd_m64_v128_load32_splat(m, n37, 0)
	n39 := Simd_f32x4_mul(n32, n38)
	n40 := Simd_f32x4_add(n28, n39)
	return n35[0], n35[1], n40[0], n40[1]
}

//go:noinline
func Simd_p_fx773(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, s1)
	n1 := Simd_m64_v128_load32_splat(m, s2, s1)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	n4 := Simd_m64_scalar_i32_add(s2, s3)
	n5 := Simd_m64_v128_load32_splat(m, n4, s1)
	return n0[0], n0[1], n1[0], n1[1], n3[0], n3[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx774(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load32_splat(m, s1, 0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	n4 := Simd_m64_v128_load(m, s0+16, 0)
	n5 := Simd_f32x4_mul(n4, n1)
	n6 := Simd_f32x4_add([2]uint64{p1, p1h}, n5)
	n7 := Simd_m64_scalar_i32_shl(s2, 2)
	n8 := Simd_m64_scalar_i32_add(s0, n7)
	n9 := Simd_m64_v128_load(m, n8, 0)
	n10 := Simd_m64_v128_load32_splat(m, s1+4, 0)
	n11 := Simd_f32x4_mul(n9, n10)
	n12 := Simd_f32x4_add(n3, n11)
	n13 := Simd_m64_scalar_i32_shl(s2, 2)
	n14 := Simd_m64_scalar_i32_add(s0, n13)
	n15 := Simd_m64_scalar_i32_add(n14, 16)
	n16 := Simd_m64_v128_load(m, n15, 0)
	n17 := Simd_f32x4_mul(n16, n10)
	n18 := Simd_f32x4_add(n6, n17)
	n19 := Simd_m64_scalar_i32_shl(s2, 2)
	n20 := Simd_m64_scalar_i32_add(s0, n19)
	n21 := Simd_m64_scalar_i32_shl(s2, 2)
	n22 := Simd_m64_scalar_i32_add(n20, n21)
	n23 := Simd_m64_v128_load(m, n22, 0)
	n24 := Simd_m64_v128_load32_splat(m, s1+8, 0)
	n25 := Simd_f32x4_mul(n23, n24)
	n26 := Simd_f32x4_add(n12, n25)
	n27 := Simd_m64_scalar_i32_shl(s2, 2)
	n28 := Simd_m64_scalar_i32_add(s0, n27)
	n29 := Simd_m64_scalar_i32_shl(s2, 2)
	n30 := Simd_m64_scalar_i32_add(n28, n29)
	n31 := Simd_m64_scalar_i32_add(n30, 16)
	n32 := Simd_m64_v128_load(m, n31, 0)
	n33 := Simd_f32x4_mul(n32, n24)
	n34 := Simd_f32x4_add(n18, n33)
	n35 := Simd_m64_scalar_i32_shl(s2, 2)
	n36 := Simd_m64_scalar_i32_add(s0, n35)
	n37 := Simd_m64_scalar_i32_shl(s2, 2)
	n38 := Simd_m64_scalar_i32_add(n36, n37)
	n39 := Simd_m64_scalar_i32_shl(s2, 2)
	n40 := Simd_m64_scalar_i32_add(n38, n39)
	n41 := Simd_m64_v128_load(m, n40, 0)
	n42 := Simd_m64_v128_load32_splat(m, s1+12, 0)
	n43 := Simd_f32x4_mul(n41, n42)
	n44 := Simd_f32x4_add(n26, n43)
	n45 := Simd_m64_scalar_i32_shl(s2, 2)
	n46 := Simd_m64_scalar_i32_add(s0, n45)
	n47 := Simd_m64_scalar_i32_shl(s2, 2)
	n48 := Simd_m64_scalar_i32_add(n46, n47)
	n49 := Simd_m64_scalar_i32_shl(s2, 2)
	n50 := Simd_m64_scalar_i32_add(n48, n49)
	n51 := Simd_m64_scalar_i32_add(n50, 16)
	n52 := Simd_m64_v128_load(m, n51, 0)
	n53 := Simd_f32x4_mul(n52, n42)
	n54 := Simd_f32x4_add(n34, n53)
	return n44[0], n44[1], n54[0], n54[1]
}

//go:noinline
func Simd_p_fx775(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load32_splat(m, s1, 0)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	n4 := Simd_m64_v128_load(m, s0+16, 0)
	n5 := Simd_f32x4_mul(n4, n1)
	n6 := Simd_f32x4_add([2]uint64{p1, p1h}, n5)
	n7 := Simd_m64_scalar_i32_add(s0, s2)
	n8 := Simd_m64_v128_load(m, n7, 0)
	n9 := Simd_m64_v128_load32_splat(m, s1+4, 0)
	n10 := Simd_f32x4_mul(n8, n9)
	n11 := Simd_f32x4_add(n3, n10)
	n12 := Simd_m64_scalar_i32_add(s0, s2)
	n13 := Simd_m64_scalar_i32_add(n12, 16)
	n14 := Simd_m64_v128_load(m, n13, 0)
	n15 := Simd_f32x4_mul(n14, n9)
	n16 := Simd_f32x4_add(n6, n15)
	n17 := Simd_m64_scalar_i32_add(s0, s2)
	n18 := Simd_m64_scalar_i32_add(n17, s2)
	n19 := Simd_m64_v128_load(m, n18, 0)
	n20 := Simd_m64_v128_load32_splat(m, s1+8, 0)
	n21 := Simd_f32x4_mul(n19, n20)
	n22 := Simd_f32x4_add(n11, n21)
	n23 := Simd_m64_scalar_i32_add(s0, s2)
	n24 := Simd_m64_scalar_i32_add(n23, s2)
	n25 := Simd_m64_scalar_i32_add(n24, 16)
	n26 := Simd_m64_v128_load(m, n25, 0)
	n27 := Simd_f32x4_mul(n26, n20)
	n28 := Simd_f32x4_add(n16, n27)
	n29 := Simd_m64_scalar_i32_add(s0, s2)
	n30 := Simd_m64_scalar_i32_add(n29, s2)
	n31 := Simd_m64_scalar_i32_add(n30, s2)
	n32 := Simd_m64_v128_load(m, n31, 0)
	n33 := Simd_m64_v128_load32_splat(m, s1+12, 0)
	n34 := Simd_f32x4_mul(n32, n33)
	n35 := Simd_f32x4_add(n22, n34)
	n36 := Simd_m64_scalar_i32_add(s0, s2)
	n37 := Simd_m64_scalar_i32_add(n36, s2)
	n38 := Simd_m64_scalar_i32_add(n37, s2)
	n39 := Simd_m64_scalar_i32_add(n38, 16)
	n40 := Simd_m64_v128_load(m, n39, 0)
	n41 := Simd_f32x4_mul(n40, n33)
	n42 := Simd_f32x4_add(n28, n41)
	return n35[0], n35[1], n42[0], n42[1]
}

//go:noinline
func Simd_p_fx776(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, s1)
	n1 := Simd_m64_v128_load32_splat(m, s2, s1)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	n4 := Simd_m64_v128_load(m, s0+16, s1)
	return n0[0], n0[1], n1[0], n1[1], n3[0], n3[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx777(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_add([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_mul([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n3 := Simd_f32x4_add(n1, n2)
	n4 := Simd_f32x4_mul([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n5 := Simd_f32x4_add(n3, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx778(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, s1)
	n1 := Simd_m64_v128_load32_splat(m, s2, s1)
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	return n0[0], n0[1], n1[0], n1[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx779(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load_rng(m, s1, 0, 0, 32)
	n2 := Simd_m64_v128_load_nc(m, s1, 16)
	n3 := Simd_i8x16_shuffle(n1, n2, [2]uint64{795458214199165184, 1952900979608391952})
	n4 := Simd_i8x16_shuffle(n1, n2, [2]uint64{1084818905551471876, 2242261670960698644})
	n5 := Simd_f32x4_mul(n0, n3)
	n6 := Simd_f32x4_mul(n0, n4)
	n7 := Simd_m64_v128_load(m, s2, 0)
	n8 := Simd_f32x4_mul(n7, n4)
	n9 := Simd_f32x4_sub(n5, n8)
	n10 := Simd_f32x4_mul(n3, n7)
	n11 := Simd_f32x4_add(n6, n10)
	_ = Simd_m64_v128_store(m, s3, 0, n9)
	_ = Simd_m64_v128_store(m, s4, 0, n11)
	return
}

//go:noinline
func Simd_p_fx780(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load_rng(m, s1, 0, 0, 32)
	n2 := Simd_m64_v128_load_nc(m, s1, 16)
	n3 := Simd_i8x16_shuffle(n1, n2, [2]uint64{795458214199165184, 1952900979608391952})
	n4 := Simd_i8x16_shuffle(n1, n2, [2]uint64{1084818905551471876, 2242261670960698644})
	n5 := Simd_f32x4_mul(n0, n3)
	n6 := Simd_f32x4_mul(n0, n4)
	n7 := Simd_m64_scalar_i32_add(s0, s2)
	n8 := Simd_m64_v128_load(m, n7, 0)
	n9 := Simd_f32x4_mul(n8, n4)
	n10 := Simd_f32x4_sub(n5, n9)
	n11 := Simd_f32x4_mul(n3, n8)
	n12 := Simd_f32x4_add(n6, n11)
	_ = Simd_m64_v128_store(m, s3, 0, n10)
	n14 := Simd_m64_scalar_i32_add(s3, s2)
	_ = Simd_m64_v128_store(m, n14, 0, n12)
	return
}

//go:noinline
func Simd_p_fx781(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n1 := Simd_m64_v128_load_nc(m, s0, 16)
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p0, p0h})
	n3 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p1, p1h})
	n4 := Simd_m64_v128_load_rng(m, s1, 0, 0, 32)
	n5 := Simd_m64_v128_load_nc(m, s1, 16)
	n6 := Simd_i8x16_shuffle(n4, n5, [2]uint64{p0, p0h})
	n7 := Simd_f32x4_mul(n6, n3)
	n8 := Simd_i8x16_shuffle(n4, n5, [2]uint64{p1, p1h})
	n9 := Simd_f32x4_mul(n2, n6)
	n10 := Simd_f32x4_mul(n3, n8)
	n11 := Simd_f32x4_sub(n9, n10)
	n12 := Simd_f32x4_mul(n2, n8)
	n13 := Simd_f32x4_add(n12, n7)
	n14 := Simd_i8x16_shuffle(n11, n13, [2]uint64{1952900979473647880, 2242261670825954572})
	n15 := Simd_i8x16_shuffle(n11, n13, [2]uint64{1374179596769034496, 1663540288121341188})
	_ = Simd_m64_v128_store(m, s2, 16, n14)
	_ = Simd_m64_v128_store(m, s2, 0, n15)
	return
}

//go:noinline
func Simd_p_fx782(m *Module, s0 int64, s1 int64, f0 float32) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load(m, n0, 0)
	n2 := Simd_f32x4_splat(f0)
	n3 := Simd_f32x4_mul(n1, n2)
	n4 := Simd_m64_v128_load(m, s0, 0)
	n5 := Simd_f32x4_add(n3, n4)
	_ = Simd_m64_v128_store(m, s0, 0, n5)
	return
}

//go:noinline
func Simd_p_fx783(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0+2, 0, 0, 34)
	n1 := Simd_v128_and(n0, [2]uint64{p0, p0h})
	n2 := Simd_i8x16_add(n1, [2]uint64{p1, p1h})
	n3 := Simd_m64_v128_load_rng(m, s1+2, 0, 0, 66)
	return n0[0], n0[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx784(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_m64_v128_load_nc(m, s0+18, 0)
	n2 := Simd_m64_v128_load_nc(m, s1+20, 0)
	n3 := Simd_v128_and(n2, [2]uint64{p2, p2h})
	n4 := Simd_i8x16_add(n3, [2]uint64{p1, p1h})
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx785(m *Module, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n3 := Simd_i32x4_dot_i16x8_s(n1, n2)
	n4 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p1, p1h})
	n5 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p2, p2h})
	n6 := Simd_i32x4_dot_i16x8_s(n4, n5)
	n7 := Simd_i32x4_add(n3, n6)
	n8 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p3, p3h})
	n9 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p4, p4h})
	n10 := Simd_i32x4_dot_i16x8_s(n8, n9)
	n11 := Simd_i32x4_add(n7, n10)
	n12 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n13 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p4, p4h})
	n14 := Simd_i32x4_dot_i16x8_s(n12, n13)
	n15 := Simd_i32x4_add(n11, n14)
	n16 := Simd_f32x4_convert_i32x4_s(n15)
	n17 := Simd_f32x4_mul(n0, n16)
	n18 := Simd_f32x4_add([2]uint64{p0, p0h}, n17)
	return n18[0], n18[1]
}

//go:noinline
func Simd_p_fx786(m *Module, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p0, p0h})
	n2 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n3 := Simd_i32x4_dot_i16x8_s(n1, n2)
	n4 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p0, p0h})
	n5 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p1, p1h})
	n6 := Simd_i32x4_dot_i16x8_s(n4, n5)
	n7 := Simd_i32x4_add(n3, n6)
	n8 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p2, p2h})
	n9 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p3, p3h})
	n10 := Simd_i32x4_dot_i16x8_s(n8, n9)
	n11 := Simd_i32x4_add(n7, n10)
	n12 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n13 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n14 := Simd_i32x4_dot_i16x8_s(n12, n13)
	n15 := Simd_i32x4_add(n11, n14)
	n16 := Simd_f32x4_convert_i32x4_s(n15)
	n17 := Simd_f32x4_mul(n0, n16)
	return n17[0], n17[1]
}

//go:noinline
func Simd_p_fx787(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_ge_s([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 256, n2)
	n4 := Simd_m64_v128_load_rng(m, s1, 0, 0, 80)
	n5 := Simd_m64_v128_load_nc(m, s1, 64)
	return n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx788(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_ge_s([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p4, p4h})
	n4 := Simd_i8x16_eq(n3, [2]uint64{p1, p1h})
	n5 := Simd_v128_and(n4, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s0, 240, n2)
	n7 := Simd_m64_v128_load(m, s1, 32)
	n8 := Simd_v128_and(n7, [2]uint64{p5, p5h})
	n9 := Simd_v128_or(n5, n8)
	_ = Simd_m64_v128_store(m, s0, 16, n9)
	n11 := Simd_m64_v128_load(m, s1, 48)
	return n7[0], n7[1], n11[0], n11[1]
}

//go:noinline
func Simd_p_fx789(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_or(n2, [2]uint64{p4, p4h})
	n4 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p1, p1h})
	n5 := Simd_i8x16_eq(n4, [2]uint64{p2, p2h})
	n6 := Simd_v128_and(n5, [2]uint64{p3, p3h})
	n7 := Simd_v128_or(n6, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0, 128, n3)
	_ = Simd_m64_v128_store(m, s0, 112, n7)
	return
}

//go:noinline
func Simd_p_fx790(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 96, n4)
	return
}

//go:noinline
func Simd_p_fx791(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 80, n4)
	return
}

//go:noinline
func Simd_p_fx792(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 64, n4)
	return
}

//go:noinline
func Simd_p_fx793(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 48, n4)
	return
}

//go:noinline
func Simd_p_fx794(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 32, n4)
	return
}

//go:noinline
func Simd_p_fx795(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 144, n4)
	return
}

//go:noinline
func Simd_p_fx796(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 224, n4)
	return
}

//go:noinline
func Simd_p_fx797(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 208, n4)
	return
}

//go:noinline
func Simd_p_fx798(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 192, n4)
	return
}

//go:noinline
func Simd_p_fx799(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 176, n4)
	return
}

//go:noinline
func Simd_p_fx800(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n4 := Simd_v128_or(n2, n3)
	_ = Simd_m64_v128_store(m, s0, 160, n4)
	return
}

//go:noinline
func Simd_p_fx801(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_mul(n0, [2]uint64{p2, p2h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx802(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_mul(n0, [2]uint64{p2, p2h})
	n2 := Simd_i32x4_extend_high_i16x8_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n1)
	n4 := Simd_i32x4_extend_high_i16x8_s([2]uint64{p4, p4h})
	n5 := Simd_i32x4_add([2]uint64{p3, p3h}, n4)
	n6 := Simd_i32x4_add(n5, n2)
	n7 := Simd_i32x4_extend_low_i16x8_s([2]uint64{p4, p4h})
	n8 := Simd_i32x4_add([2]uint64{p5, p5h}, n7)
	n9 := Simd_i32x4_add(n8, n3)
	return n6[0], n6[1], n9[0], n9[1]
}

//go:noinline
func Simd_p_fx803(m *Module, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_convert_i32x4_s([2]uint64{p1, p1h})
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_convert_i32x4_s([2]uint64{p2, p2h})
	n4 := Simd_f32x4_mul(n0, n3)
	n5 := Simd_f32x4_add(n2, n4)
	n6 := Simd_f32x4_add([2]uint64{p0, p0h}, n5)
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx804(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_add(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 64, n2)
	n4 := Simd_m64_v128_load_rng(m, s1, 144, 16, 144)
	n5 := Simd_m64_v128_load_nc(m, s1, 16)
	return n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx805(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p1, p1h})
	n4 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n5 := Simd_v128_or(n3, n4)
	n6 := Simd_i8x16_add(n5, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 80, n2)
	_ = Simd_m64_v128_store(m, s0, 0, n6)
	return
}

//go:noinline
func Simd_p_fx806(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p1, p1h})
	n4 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n5 := Simd_v128_or(n3, n4)
	n6 := Simd_i8x16_add(n5, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 96, n2)
	_ = Simd_m64_v128_store(m, s0, 32, n6)
	return
}

//go:noinline
func Simd_p_fx807(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_v128_or(n0, n1)
	n3 := Simd_i8x16_add(n2, [2]uint64{p4, p4h})
	_ = Simd_m64_v128_store(m, s0, 16, n3)
	n5 := Simd_m64_v128_load(m, s1, 48)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx808(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p1, p1h})
	n4 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n5 := Simd_v128_or(n3, n4)
	n6 := Simd_i8x16_add(n5, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 112, n2)
	_ = Simd_m64_v128_store(m, s0, 48, n6)
	return
}

//go:noinline
func Simd_p_fx809(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_m64_v128_load_nc(m, s0, 64)
	n2 := Simd_v128_and(n1, [2]uint64{p2, p2h})
	n3 := Simd_v128_or(n0, n2)
	n4 := Simd_i8x16_add(n3, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s1, 128, n4)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx810(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 192, n2)
	n5 := Simd_m64_v128_load(m, s1, 96)
	n6 := Simd_v128_and(n5, [2]uint64{p5, p5h})
	n7 := Simd_v128_or(n3, n6)
	n8 := Simd_i8x16_add(n7, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 160, n8)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx811(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_add(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 224, n2)
	n4 := Simd_m64_v128_load_rng(m, s1, 176, 80, 112)
	n5 := Simd_m64_v128_load_nc(m, s1, 80)
	return n4[0], n4[1], n5[0], n5[1]
}

//go:noinline
func Simd_p_fx812(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_add(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p1, p1h})
	n4 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n5 := Simd_v128_or(n3, n4)
	n6 := Simd_i8x16_add(n5, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 208, n2)
	_ = Simd_m64_v128_store(m, s0, 144, n6)
	return
}

//go:noinline
func Simd_p_fx813(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_m64_v128_load(m, s0, 112)
	n2 := Simd_v128_and(n1, [2]uint64{p2, p2h})
	n3 := Simd_v128_or(n0, n2)
	n4 := Simd_i8x16_add(n3, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s1, 176, n4)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx814(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	n2 := Simd_i8x16_add(n1, [2]uint64{p3, p3h})
	_ = Simd_m64_v128_store(m, s0, 240, n2)
	return
}

//go:noinline
func Simd_p_fx815(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_splat(int32(s3))
	n1 := Simd_m64_scalar_i32_add(s0, s1)
	n2 := Simd_m64_v128_load(m, n1, 0)
	n3 := Simd_m64_scalar_i32_add(s2, s0)
	n4 := Simd_m64_v128_load(m, n3, 0)
	n5 := Simd_i16x8_extmul_low_i8x16_s(n2, n4)
	n6 := Simd_i32x4_extend_high_i16x8_s(n5)
	n7 := Simd_i16x8_extmul_high_i8x16_s(n2, n4)
	n8 := Simd_i32x4_extend_high_i16x8_s(n7)
	n9 := Simd_i32x4_add(n6, n8)
	n10 := Simd_i32x4_mul(n9, n0)
	n11 := Simd_i32x4_add(n10, [2]uint64{p0, p0h})
	n12 := Simd_i32x4_extend_low_i16x8_s(n5)
	n13 := Simd_i32x4_extend_low_i16x8_s(n7)
	n14 := Simd_i32x4_add(n12, n13)
	n15 := Simd_i32x4_mul(n14, n0)
	n16 := Simd_i32x4_add(n15, [2]uint64{p1, p1h})
	return n11[0], n11[1], n16[0], n16[1]
}

//go:noinline
func Simd_p_fx816(m *Module, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_convert_i32x4_s([2]uint64{p0, p0h})
	n2 := Simd_f32x4_mul(n0, n1)
	n3 := Simd_f32x4_add(n2, [2]uint64{p1, p1h})
	n4 := Simd_f32x4_convert_i32x4_s([2]uint64{p2, p2h})
	n5 := Simd_f32x4_mul(n0, n4)
	n6 := Simd_f32x4_add(n5, [2]uint64{p3, p3h})
	return n3[0], n3[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx817(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_scalar_i32_load16_u(m, n0)
	n2 := Simd_m64_scalar_i32_shl(n1, 2)
	n3 := Simd_m64_v128_load32_splat(m, n2, 8793760)
	n4 := Simd_m64_scalar_i32_add(s0, s2)
	n5 := Simd_m64_v128_load16x4_u(m, n4, 0)
	n6 := Simd_f16x4_cvt(n5)
	n7 := Simd_f32x4_mul(n3, n6)
	n8 := Simd_m64_scalar_i32_add(s0, s2)
	n9 := Simd_m64_scalar_i32_add(n8, 120)
	n10 := Simd_m64_v128_load(m, n9, 0)
	n11 := Simd_i16x8_extend_low_i8x16_s(n10)
	n12 := Simd_i16x8_extend_high_i8x16_s(n10)
	n13 := Simd_m64_scalar_i32_add(s0, s1)
	n14 := Simd_m64_scalar_i32_add(n13, 120)
	n15 := Simd_m64_v128_load32_splat(m, n14, 0)
	n16 := Simd_i16x8_extend_low_i8x16_s(n15)
	n17 := Simd_i32x4_dot_i16x8_s(n11, n16)
	n18 := Simd_i16x8_extend_high_i8x16_s(n15)
	n19 := Simd_i32x4_dot_i16x8_s(n12, n18)
	n20 := Simd_m64_scalar_i32_add(s0, s2)
	n21 := Simd_m64_scalar_i32_add(n20, 104)
	n22 := Simd_m64_v128_load(m, n21, 0)
	n23 := Simd_i16x8_extend_low_i8x16_s(n22)
	n24 := Simd_i16x8_extend_high_i8x16_s(n22)
	n25 := Simd_m64_scalar_i32_add(s0, s1)
	n26 := Simd_m64_scalar_i32_add(n25, 104)
	n27 := Simd_m64_v128_load32_splat(m, n26, 0)
	n28 := Simd_i16x8_extend_low_i8x16_s(n27)
	n29 := Simd_i32x4_dot_i16x8_s(n23, n28)
	n30 := Simd_i32x4_add(n17, n29)
	n31 := Simd_i16x8_extend_high_i8x16_s(n27)
	n32 := Simd_i32x4_dot_i16x8_s(n24, n31)
	n33 := Simd_i32x4_add(n19, n32)
	n34 := Simd_m64_scalar_i32_add(s0, s2)
	n35 := Simd_m64_scalar_i32_add(n34, 88)
	n36 := Simd_m64_v128_load(m, n35, 0)
	n37 := Simd_i16x8_extend_low_i8x16_s(n36)
	n38 := Simd_i16x8_extend_high_i8x16_s(n36)
	n39 := Simd_m64_scalar_i32_add(s0, s1)
	n40 := Simd_m64_scalar_i32_add(n39, 88)
	n41 := Simd_m64_v128_load32_splat(m, n40, 0)
	n42 := Simd_i16x8_extend_low_i8x16_s(n41)
	n43 := Simd_i32x4_dot_i16x8_s(n37, n42)
	n44 := Simd_i32x4_add(n30, n43)
	n45 := Simd_i16x8_extend_high_i8x16_s(n41)
	n46 := Simd_i32x4_dot_i16x8_s(n38, n45)
	n47 := Simd_i32x4_add(n33, n46)
	n48 := Simd_m64_scalar_i32_add(s0, s2)
	n49 := Simd_m64_scalar_i32_add(n48, 72)
	n50 := Simd_m64_v128_load(m, n49, 0)
	n51 := Simd_i16x8_extend_low_i8x16_s(n50)
	n52 := Simd_i16x8_extend_high_i8x16_s(n50)
	n53 := Simd_m64_scalar_i32_add(s0, s1)
	n54 := Simd_m64_scalar_i32_add(n53, 72)
	n55 := Simd_m64_v128_load32_splat(m, n54, 0)
	n56 := Simd_i16x8_extend_low_i8x16_s(n55)
	n57 := Simd_i32x4_dot_i16x8_s(n51, n56)
	n58 := Simd_i32x4_add(n44, n57)
	n59 := Simd_i16x8_extend_high_i8x16_s(n55)
	n60 := Simd_i32x4_dot_i16x8_s(n52, n59)
	n61 := Simd_i32x4_add(n47, n60)
	n62 := Simd_m64_scalar_i32_add(s0, s2)
	n63 := Simd_m64_scalar_i32_add(n62, 56)
	n64 := Simd_m64_v128_load(m, n63, 0)
	n65 := Simd_i16x8_extend_low_i8x16_s(n64)
	n66 := Simd_i16x8_extend_high_i8x16_s(n64)
	n67 := Simd_m64_scalar_i32_add(s0, s1)
	n68 := Simd_m64_scalar_i32_add(n67, 56)
	n69 := Simd_m64_v128_load32_splat(m, n68, 0)
	n70 := Simd_i16x8_extend_low_i8x16_s(n69)
	n71 := Simd_i32x4_dot_i16x8_s(n65, n70)
	n72 := Simd_i32x4_add(n58, n71)
	n73 := Simd_i16x8_extend_high_i8x16_s(n69)
	n74 := Simd_i32x4_dot_i16x8_s(n66, n73)
	n75 := Simd_i32x4_add(n61, n74)
	n76 := Simd_m64_scalar_i32_add(s0, s2)
	n77 := Simd_m64_scalar_i32_add(n76, 40)
	n78 := Simd_m64_v128_load(m, n77, 0)
	n79 := Simd_i16x8_extend_low_i8x16_s(n78)
	n80 := Simd_i16x8_extend_high_i8x16_s(n78)
	n81 := Simd_m64_scalar_i32_add(s0, s1)
	n82 := Simd_m64_scalar_i32_add(n81, 40)
	n83 := Simd_m64_v128_load32_splat(m, n82, 0)
	n84 := Simd_i16x8_extend_low_i8x16_s(n83)
	n85 := Simd_i32x4_dot_i16x8_s(n79, n84)
	n86 := Simd_i32x4_add(n72, n85)
	n87 := Simd_i16x8_extend_high_i8x16_s(n83)
	n88 := Simd_i32x4_dot_i16x8_s(n80, n87)
	n89 := Simd_i32x4_add(n75, n88)
	n90 := Simd_m64_scalar_i32_add(s0, s2)
	n91 := Simd_m64_scalar_i32_add(n90, 24)
	n92 := Simd_m64_v128_load(m, n91, 0)
	n93 := Simd_i16x8_extend_low_i8x16_s(n92)
	n94 := Simd_i16x8_extend_high_i8x16_s(n92)
	n95 := Simd_m64_scalar_i32_add(s0, s1)
	n96 := Simd_m64_scalar_i32_add(n95, 24)
	n97 := Simd_m64_v128_load32_splat(m, n96, 0)
	n98 := Simd_i16x8_extend_low_i8x16_s(n97)
	n99 := Simd_i32x4_dot_i16x8_s(n93, n98)
	n100 := Simd_i32x4_add(n86, n99)
	n101 := Simd_i16x8_extend_high_i8x16_s(n97)
	n102 := Simd_i32x4_dot_i16x8_s(n94, n101)
	n103 := Simd_i32x4_add(n89, n102)
	n104 := Simd_m64_scalar_i32_add(s0, s2)
	n105 := Simd_m64_scalar_i32_add(n104, 8)
	n106 := Simd_m64_v128_load(m, n105, 0)
	n107 := Simd_i16x8_extend_low_i8x16_s(n106)
	n108 := Simd_i16x8_extend_high_i8x16_s(n106)
	n109 := Simd_m64_scalar_i32_add(s0, s1)
	n110 := Simd_m64_scalar_i32_add(n109, 8)
	n111 := Simd_m64_v128_load32_splat(m, n110, 0)
	n112 := Simd_i16x8_extend_low_i8x16_s(n111)
	n113 := Simd_i32x4_dot_i16x8_s(n107, n112)
	n114 := Simd_i32x4_add(n100, n113)
	n115 := Simd_i16x8_extend_high_i8x16_s(n111)
	n116 := Simd_i32x4_dot_i16x8_s(n108, n115)
	n117 := Simd_i32x4_add(n103, n116)
	n118 := Simd_i8x16_shuffle(n114, n117, [2]uint64{795458214199165184, 1952900979608391952})
	n119 := Simd_i8x16_shuffle(n114, n117, [2]uint64{1084818905551471876, 2242261670960698644})
	n120 := Simd_i32x4_add(n118, n119)
	n121 := Simd_f32x4_convert_i32x4_s(n120)
	n122 := Simd_f32x4_mul(n7, n121)
	n123 := Simd_f32x4_add([2]uint64{p0, p0h}, n122)
	return n123[0], n123[1]
}

//go:noinline
func Simd_p_fx818(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_scalar_i32_add(n0, 2)
	n2 := Simd_m64_scalar_i32_load16_u(m, n1)
	n3 := Simd_m64_scalar_i32_shl(n2, 2)
	n4 := Simd_m64_v128_load32_splat(m, n3, 8793760)
	n5 := Simd_m64_scalar_i32_add(s0, s2)
	n6 := Simd_m64_v128_load16x4_u(m, n5, 0)
	n7 := Simd_f16x4_cvt(n6)
	n8 := Simd_f32x4_mul(n4, n7)
	n9 := Simd_m64_scalar_i32_add(s0, s2)
	n10 := Simd_m64_scalar_i32_add(n9, 120)
	n11 := Simd_m64_v128_load(m, n10, 0)
	n12 := Simd_i16x8_extend_low_i8x16_s(n11)
	n13 := Simd_i16x8_extend_high_i8x16_s(n11)
	n14 := Simd_m64_scalar_i32_add(s0, s1)
	n15 := Simd_m64_scalar_i32_add(n14, 124)
	n16 := Simd_m64_v128_load32_splat(m, n15, 0)
	n17 := Simd_i16x8_extend_low_i8x16_s(n16)
	n18 := Simd_i32x4_dot_i16x8_s(n12, n17)
	n19 := Simd_i16x8_extend_high_i8x16_s(n16)
	n20 := Simd_i32x4_dot_i16x8_s(n13, n19)
	n21 := Simd_m64_scalar_i32_add(s0, s2)
	n22 := Simd_m64_scalar_i32_add(n21, 104)
	n23 := Simd_m64_v128_load(m, n22, 0)
	n24 := Simd_i16x8_extend_low_i8x16_s(n23)
	n25 := Simd_i16x8_extend_high_i8x16_s(n23)
	n26 := Simd_m64_scalar_i32_add(s0, s1)
	n27 := Simd_m64_scalar_i32_add(n26, 108)
	n28 := Simd_m64_v128_load32_splat(m, n27, 0)
	n29 := Simd_i16x8_extend_low_i8x16_s(n28)
	n30 := Simd_i32x4_dot_i16x8_s(n24, n29)
	n31 := Simd_i32x4_add(n18, n30)
	n32 := Simd_i16x8_extend_high_i8x16_s(n28)
	n33 := Simd_i32x4_dot_i16x8_s(n25, n32)
	n34 := Simd_i32x4_add(n20, n33)
	n35 := Simd_m64_scalar_i32_add(s0, s2)
	n36 := Simd_m64_scalar_i32_add(n35, 88)
	n37 := Simd_m64_v128_load(m, n36, 0)
	n38 := Simd_i16x8_extend_low_i8x16_s(n37)
	n39 := Simd_i16x8_extend_high_i8x16_s(n37)
	n40 := Simd_m64_scalar_i32_add(s0, s1)
	n41 := Simd_m64_scalar_i32_add(n40, 92)
	n42 := Simd_m64_v128_load32_splat(m, n41, 0)
	n43 := Simd_i16x8_extend_low_i8x16_s(n42)
	n44 := Simd_i32x4_dot_i16x8_s(n38, n43)
	n45 := Simd_i32x4_add(n31, n44)
	n46 := Simd_i16x8_extend_high_i8x16_s(n42)
	n47 := Simd_i32x4_dot_i16x8_s(n39, n46)
	n48 := Simd_i32x4_add(n34, n47)
	n49 := Simd_m64_scalar_i32_add(s0, s2)
	n50 := Simd_m64_scalar_i32_add(n49, 72)
	n51 := Simd_m64_v128_load(m, n50, 0)
	n52 := Simd_i16x8_extend_low_i8x16_s(n51)
	n53 := Simd_i16x8_extend_high_i8x16_s(n51)
	n54 := Simd_m64_scalar_i32_add(s0, s1)
	n55 := Simd_m64_scalar_i32_add(n54, 76)
	n56 := Simd_m64_v128_load32_splat(m, n55, 0)
	n57 := Simd_i16x8_extend_low_i8x16_s(n56)
	n58 := Simd_i32x4_dot_i16x8_s(n52, n57)
	n59 := Simd_i32x4_add(n45, n58)
	n60 := Simd_i16x8_extend_high_i8x16_s(n56)
	n61 := Simd_i32x4_dot_i16x8_s(n53, n60)
	n62 := Simd_i32x4_add(n48, n61)
	n63 := Simd_m64_scalar_i32_add(s0, s2)
	n64 := Simd_m64_scalar_i32_add(n63, 56)
	n65 := Simd_m64_v128_load(m, n64, 0)
	n66 := Simd_i16x8_extend_low_i8x16_s(n65)
	n67 := Simd_i16x8_extend_high_i8x16_s(n65)
	n68 := Simd_m64_scalar_i32_add(s0, s1)
	n69 := Simd_m64_scalar_i32_add(n68, 60)
	n70 := Simd_m64_v128_load32_splat(m, n69, 0)
	n71 := Simd_i16x8_extend_low_i8x16_s(n70)
	n72 := Simd_i32x4_dot_i16x8_s(n66, n71)
	n73 := Simd_i32x4_add(n59, n72)
	n74 := Simd_i16x8_extend_high_i8x16_s(n70)
	n75 := Simd_i32x4_dot_i16x8_s(n67, n74)
	n76 := Simd_i32x4_add(n62, n75)
	n77 := Simd_m64_scalar_i32_add(s0, s2)
	n78 := Simd_m64_scalar_i32_add(n77, 40)
	n79 := Simd_m64_v128_load(m, n78, 0)
	n80 := Simd_i16x8_extend_low_i8x16_s(n79)
	n81 := Simd_i16x8_extend_high_i8x16_s(n79)
	n82 := Simd_m64_scalar_i32_add(s0, s1)
	n83 := Simd_m64_scalar_i32_add(n82, 44)
	n84 := Simd_m64_v128_load32_splat(m, n83, 0)
	n85 := Simd_i16x8_extend_low_i8x16_s(n84)
	n86 := Simd_i32x4_dot_i16x8_s(n80, n85)
	n87 := Simd_i32x4_add(n73, n86)
	n88 := Simd_i16x8_extend_high_i8x16_s(n84)
	n89 := Simd_i32x4_dot_i16x8_s(n81, n88)
	n90 := Simd_i32x4_add(n76, n89)
	n91 := Simd_m64_scalar_i32_add(s0, s2)
	n92 := Simd_m64_scalar_i32_add(n91, 24)
	n93 := Simd_m64_v128_load(m, n92, 0)
	n94 := Simd_i16x8_extend_low_i8x16_s(n93)
	n95 := Simd_i16x8_extend_high_i8x16_s(n93)
	n96 := Simd_m64_scalar_i32_add(s0, s1)
	n97 := Simd_m64_scalar_i32_add(n96, 28)
	n98 := Simd_m64_v128_load32_splat(m, n97, 0)
	n99 := Simd_i16x8_extend_low_i8x16_s(n98)
	n100 := Simd_i32x4_dot_i16x8_s(n94, n99)
	n101 := Simd_i32x4_add(n87, n100)
	n102 := Simd_i16x8_extend_high_i8x16_s(n98)
	n103 := Simd_i32x4_dot_i16x8_s(n95, n102)
	n104 := Simd_i32x4_add(n90, n103)
	n105 := Simd_m64_scalar_i32_add(s0, s2)
	n106 := Simd_m64_scalar_i32_add(n105, 8)
	n107 := Simd_m64_v128_load(m, n106, 0)
	n108 := Simd_i16x8_extend_low_i8x16_s(n107)
	n109 := Simd_i16x8_extend_high_i8x16_s(n107)
	n110 := Simd_m64_scalar_i32_add(s0, s1)
	n111 := Simd_m64_scalar_i32_add(n110, 12)
	n112 := Simd_m64_v128_load32_splat(m, n111, 0)
	n113 := Simd_i16x8_extend_low_i8x16_s(n112)
	n114 := Simd_i32x4_dot_i16x8_s(n108, n113)
	n115 := Simd_i32x4_add(n101, n114)
	n116 := Simd_i16x8_extend_high_i8x16_s(n112)
	n117 := Simd_i32x4_dot_i16x8_s(n109, n116)
	n118 := Simd_i32x4_add(n104, n117)
	n119 := Simd_i8x16_shuffle(n115, n118, [2]uint64{795458214199165184, 1952900979608391952})
	n120 := Simd_i8x16_shuffle(n115, n118, [2]uint64{1084818905551471876, 2242261670960698644})
	n121 := Simd_i32x4_add(n119, n120)
	n122 := Simd_f32x4_convert_i32x4_s(n121)
	n123 := Simd_f32x4_mul(n8, n122)
	n124 := Simd_f32x4_add([2]uint64{p0, p0h}, n123)
	return n124[0], n124[1]
}

//go:noinline
func Simd_p_fx819(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_scalar_i32_add(n0, 4)
	n2 := Simd_m64_scalar_i32_load16_u(m, n1)
	n3 := Simd_m64_scalar_i32_shl(n2, 2)
	n4 := Simd_m64_v128_load32_splat(m, n3, 8793760)
	n5 := Simd_m64_scalar_i32_add(s0, s2)
	n6 := Simd_m64_v128_load16x4_u(m, n5, 0)
	n7 := Simd_f16x4_cvt(n6)
	n8 := Simd_f32x4_mul(n4, n7)
	n9 := Simd_m64_scalar_i32_add(s0, s2)
	n10 := Simd_m64_scalar_i32_add(n9, 120)
	n11 := Simd_m64_v128_load(m, n10, 0)
	n12 := Simd_i16x8_extend_low_i8x16_s(n11)
	n13 := Simd_i16x8_extend_high_i8x16_s(n11)
	n14 := Simd_m64_scalar_i32_add(s0, s1)
	n15 := Simd_m64_scalar_i32_add(n14, 128)
	n16 := Simd_m64_v128_load32_splat(m, n15, 0)
	n17 := Simd_i16x8_extend_low_i8x16_s(n16)
	n18 := Simd_i32x4_dot_i16x8_s(n12, n17)
	n19 := Simd_i16x8_extend_high_i8x16_s(n16)
	n20 := Simd_i32x4_dot_i16x8_s(n13, n19)
	n21 := Simd_m64_scalar_i32_add(s0, s2)
	n22 := Simd_m64_scalar_i32_add(n21, 104)
	n23 := Simd_m64_v128_load(m, n22, 0)
	n24 := Simd_i16x8_extend_low_i8x16_s(n23)
	n25 := Simd_i16x8_extend_high_i8x16_s(n23)
	n26 := Simd_m64_scalar_i32_add(s0, s1)
	n27 := Simd_m64_scalar_i32_add(n26, 112)
	n28 := Simd_m64_v128_load32_splat(m, n27, 0)
	n29 := Simd_i16x8_extend_low_i8x16_s(n28)
	n30 := Simd_i32x4_dot_i16x8_s(n24, n29)
	n31 := Simd_i32x4_add(n18, n30)
	n32 := Simd_i16x8_extend_high_i8x16_s(n28)
	n33 := Simd_i32x4_dot_i16x8_s(n25, n32)
	n34 := Simd_i32x4_add(n20, n33)
	n35 := Simd_m64_scalar_i32_add(s0, s2)
	n36 := Simd_m64_scalar_i32_add(n35, 88)
	n37 := Simd_m64_v128_load(m, n36, 0)
	n38 := Simd_i16x8_extend_low_i8x16_s(n37)
	n39 := Simd_i16x8_extend_high_i8x16_s(n37)
	n40 := Simd_m64_scalar_i32_add(s0, s1)
	n41 := Simd_m64_scalar_i32_add(n40, 96)
	n42 := Simd_m64_v128_load32_splat(m, n41, 0)
	n43 := Simd_i16x8_extend_low_i8x16_s(n42)
	n44 := Simd_i32x4_dot_i16x8_s(n38, n43)
	n45 := Simd_i32x4_add(n31, n44)
	n46 := Simd_i16x8_extend_high_i8x16_s(n42)
	n47 := Simd_i32x4_dot_i16x8_s(n39, n46)
	n48 := Simd_i32x4_add(n34, n47)
	n49 := Simd_m64_scalar_i32_add(s0, s2)
	n50 := Simd_m64_scalar_i32_add(n49, 72)
	n51 := Simd_m64_v128_load(m, n50, 0)
	n52 := Simd_i16x8_extend_low_i8x16_s(n51)
	n53 := Simd_i16x8_extend_high_i8x16_s(n51)
	n54 := Simd_m64_scalar_i32_add(s0, s1)
	n55 := Simd_m64_scalar_i32_add(n54, 80)
	n56 := Simd_m64_v128_load32_splat(m, n55, 0)
	n57 := Simd_i16x8_extend_low_i8x16_s(n56)
	n58 := Simd_i32x4_dot_i16x8_s(n52, n57)
	n59 := Simd_i32x4_add(n45, n58)
	n60 := Simd_i16x8_extend_high_i8x16_s(n56)
	n61 := Simd_i32x4_dot_i16x8_s(n53, n60)
	n62 := Simd_i32x4_add(n48, n61)
	n63 := Simd_m64_scalar_i32_add(s0, s2)
	n64 := Simd_m64_scalar_i32_add(n63, 56)
	n65 := Simd_m64_v128_load(m, n64, 0)
	n66 := Simd_i16x8_extend_low_i8x16_s(n65)
	n67 := Simd_i16x8_extend_high_i8x16_s(n65)
	n68 := Simd_m64_v128_load32_splat(m, s3, 0)
	n69 := Simd_i16x8_extend_low_i8x16_s(n68)
	n70 := Simd_i32x4_dot_i16x8_s(n66, n69)
	n71 := Simd_i32x4_add(n59, n70)
	n72 := Simd_i16x8_extend_high_i8x16_s(n68)
	n73 := Simd_i32x4_dot_i16x8_s(n67, n72)
	n74 := Simd_i32x4_add(n62, n73)
	n75 := Simd_m64_scalar_i32_add(s0, s2)
	n76 := Simd_m64_scalar_i32_add(n75, 40)
	n77 := Simd_m64_v128_load(m, n76, 0)
	n78 := Simd_i16x8_extend_low_i8x16_s(n77)
	n79 := Simd_i16x8_extend_high_i8x16_s(n77)
	n80 := Simd_m64_scalar_i32_add(s0, s1)
	n81 := Simd_m64_scalar_i32_add(n80, 48)
	n82 := Simd_m64_v128_load32_splat(m, n81, 0)
	n83 := Simd_i16x8_extend_low_i8x16_s(n82)
	n84 := Simd_i32x4_dot_i16x8_s(n78, n83)
	n85 := Simd_i32x4_add(n71, n84)
	n86 := Simd_i16x8_extend_high_i8x16_s(n82)
	n87 := Simd_i32x4_dot_i16x8_s(n79, n86)
	n88 := Simd_i32x4_add(n74, n87)
	n89 := Simd_m64_scalar_i32_add(s0, s2)
	n90 := Simd_m64_scalar_i32_add(n89, 24)
	n91 := Simd_m64_v128_load(m, n90, 0)
	n92 := Simd_i16x8_extend_low_i8x16_s(n91)
	n93 := Simd_i16x8_extend_high_i8x16_s(n91)
	n94 := Simd_m64_scalar_i32_add(s0, s1)
	n95 := Simd_m64_scalar_i32_add(n94, 32)
	n96 := Simd_m64_v128_load32_splat(m, n95, 0)
	n97 := Simd_i16x8_extend_low_i8x16_s(n96)
	n98 := Simd_i32x4_dot_i16x8_s(n92, n97)
	n99 := Simd_i32x4_add(n85, n98)
	n100 := Simd_i16x8_extend_high_i8x16_s(n96)
	n101 := Simd_i32x4_dot_i16x8_s(n93, n100)
	n102 := Simd_i32x4_add(n88, n101)
	n103 := Simd_m64_scalar_i32_add(s0, s2)
	n104 := Simd_m64_scalar_i32_add(n103, 8)
	n105 := Simd_m64_v128_load(m, n104, 0)
	n106 := Simd_i16x8_extend_low_i8x16_s(n105)
	n107 := Simd_i16x8_extend_high_i8x16_s(n105)
	n108 := Simd_m64_scalar_i32_add(s0, s1)
	n109 := Simd_m64_scalar_i32_add(n108, 16)
	n110 := Simd_m64_v128_load32_splat(m, n109, 0)
	n111 := Simd_i16x8_extend_low_i8x16_s(n110)
	n112 := Simd_i32x4_dot_i16x8_s(n106, n111)
	n113 := Simd_i32x4_add(n99, n112)
	n114 := Simd_i16x8_extend_high_i8x16_s(n110)
	n115 := Simd_i32x4_dot_i16x8_s(n107, n114)
	n116 := Simd_i32x4_add(n102, n115)
	n117 := Simd_i8x16_shuffle(n113, n116, [2]uint64{795458214199165184, 1952900979608391952})
	n118 := Simd_i8x16_shuffle(n113, n116, [2]uint64{1084818905551471876, 2242261670960698644})
	n119 := Simd_i32x4_add(n117, n118)
	n120 := Simd_f32x4_convert_i32x4_s(n119)
	n121 := Simd_f32x4_mul(n8, n120)
	n122 := Simd_f32x4_add([2]uint64{p0, p0h}, n121)
	return n122[0], n122[1]
}

//go:noinline
func Simd_p_fx820(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_scalar_i32_add(n0, 6)
	n2 := Simd_m64_scalar_i32_load16_u(m, n1)
	n3 := Simd_m64_scalar_i32_shl(n2, 2)
	n4 := Simd_m64_v128_load32_splat(m, n3, 8793760)
	n5 := Simd_m64_scalar_i32_add(s0, s2)
	n6 := Simd_m64_v128_load16x4_u(m, n5, 0)
	n7 := Simd_f16x4_cvt(n6)
	n8 := Simd_f32x4_mul(n4, n7)
	n9 := Simd_m64_scalar_i32_add(s0, s2)
	n10 := Simd_m64_scalar_i32_add(n9, 120)
	n11 := Simd_m64_v128_load(m, n10, 0)
	n12 := Simd_i16x8_extend_low_i8x16_s(n11)
	n13 := Simd_i16x8_extend_high_i8x16_s(n11)
	n14 := Simd_m64_scalar_i32_add(s0, s1)
	n15 := Simd_m64_scalar_i32_add(n14, 132)
	n16 := Simd_m64_v128_load32_splat(m, n15, 0)
	n17 := Simd_i16x8_extend_low_i8x16_s(n16)
	n18 := Simd_i32x4_dot_i16x8_s(n12, n17)
	n19 := Simd_i16x8_extend_high_i8x16_s(n16)
	n20 := Simd_i32x4_dot_i16x8_s(n13, n19)
	n21 := Simd_m64_scalar_i32_add(s0, s2)
	n22 := Simd_m64_scalar_i32_add(n21, 104)
	n23 := Simd_m64_v128_load(m, n22, 0)
	n24 := Simd_i16x8_extend_low_i8x16_s(n23)
	n25 := Simd_i16x8_extend_high_i8x16_s(n23)
	n26 := Simd_m64_scalar_i32_add(s0, s1)
	n27 := Simd_m64_scalar_i32_add(n26, 116)
	n28 := Simd_m64_v128_load32_splat(m, n27, 0)
	n29 := Simd_i16x8_extend_low_i8x16_s(n28)
	n30 := Simd_i32x4_dot_i16x8_s(n24, n29)
	n31 := Simd_i32x4_add(n18, n30)
	n32 := Simd_i16x8_extend_high_i8x16_s(n28)
	n33 := Simd_i32x4_dot_i16x8_s(n25, n32)
	n34 := Simd_i32x4_add(n20, n33)
	n35 := Simd_m64_scalar_i32_add(s0, s2)
	n36 := Simd_m64_scalar_i32_add(n35, 88)
	n37 := Simd_m64_v128_load(m, n36, 0)
	n38 := Simd_i16x8_extend_low_i8x16_s(n37)
	n39 := Simd_i16x8_extend_high_i8x16_s(n37)
	n40 := Simd_m64_scalar_i32_add(s0, s1)
	n41 := Simd_m64_scalar_i32_add(n40, 100)
	n42 := Simd_m64_v128_load32_splat(m, n41, 0)
	n43 := Simd_i16x8_extend_low_i8x16_s(n42)
	n44 := Simd_i32x4_dot_i16x8_s(n38, n43)
	n45 := Simd_i32x4_add(n31, n44)
	n46 := Simd_i16x8_extend_high_i8x16_s(n42)
	n47 := Simd_i32x4_dot_i16x8_s(n39, n46)
	n48 := Simd_i32x4_add(n34, n47)
	n49 := Simd_m64_scalar_i32_add(s0, s2)
	n50 := Simd_m64_scalar_i32_add(n49, 72)
	n51 := Simd_m64_v128_load(m, n50, 0)
	n52 := Simd_i16x8_extend_low_i8x16_s(n51)
	n53 := Simd_i16x8_extend_high_i8x16_s(n51)
	n54 := Simd_m64_scalar_i32_add(s0, s1)
	n55 := Simd_m64_scalar_i32_add(n54, 84)
	n56 := Simd_m64_v128_load32_splat(m, n55, 0)
	n57 := Simd_i16x8_extend_low_i8x16_s(n56)
	n58 := Simd_i32x4_dot_i16x8_s(n52, n57)
	n59 := Simd_i32x4_add(n45, n58)
	n60 := Simd_i16x8_extend_high_i8x16_s(n56)
	n61 := Simd_i32x4_dot_i16x8_s(n53, n60)
	n62 := Simd_i32x4_add(n48, n61)
	n63 := Simd_m64_scalar_i32_add(s0, s2)
	n64 := Simd_m64_scalar_i32_add(n63, 56)
	n65 := Simd_m64_v128_load(m, n64, 0)
	n66 := Simd_i16x8_extend_low_i8x16_s(n65)
	n67 := Simd_i16x8_extend_high_i8x16_s(n65)
	n68 := Simd_m64_scalar_i32_add(s0, s1)
	n69 := Simd_m64_scalar_i32_add(n68, 68)
	n70 := Simd_m64_v128_load32_splat(m, n69, 0)
	n71 := Simd_i16x8_extend_low_i8x16_s(n70)
	n72 := Simd_i32x4_dot_i16x8_s(n66, n71)
	n73 := Simd_i32x4_add(n59, n72)
	n74 := Simd_i16x8_extend_high_i8x16_s(n70)
	n75 := Simd_i32x4_dot_i16x8_s(n67, n74)
	n76 := Simd_i32x4_add(n62, n75)
	n77 := Simd_m64_scalar_i32_add(s0, s2)
	n78 := Simd_m64_scalar_i32_add(n77, 40)
	n79 := Simd_m64_v128_load(m, n78, 0)
	n80 := Simd_i16x8_extend_low_i8x16_s(n79)
	n81 := Simd_i16x8_extend_high_i8x16_s(n79)
	n82 := Simd_m64_scalar_i32_add(s0, s1)
	n83 := Simd_m64_scalar_i32_add(n82, 52)
	n84 := Simd_m64_v128_load32_splat(m, n83, 0)
	n85 := Simd_i16x8_extend_low_i8x16_s(n84)
	n86 := Simd_i32x4_dot_i16x8_s(n80, n85)
	n87 := Simd_i32x4_add(n73, n86)
	n88 := Simd_i16x8_extend_high_i8x16_s(n84)
	n89 := Simd_i32x4_dot_i16x8_s(n81, n88)
	n90 := Simd_i32x4_add(n76, n89)
	n91 := Simd_m64_scalar_i32_add(s0, s2)
	n92 := Simd_m64_scalar_i32_add(n91, 24)
	n93 := Simd_m64_v128_load(m, n92, 0)
	n94 := Simd_i16x8_extend_low_i8x16_s(n93)
	n95 := Simd_i16x8_extend_high_i8x16_s(n93)
	n96 := Simd_m64_scalar_i32_add(s0, s1)
	n97 := Simd_m64_scalar_i32_add(n96, 36)
	n98 := Simd_m64_v128_load32_splat(m, n97, 0)
	n99 := Simd_i16x8_extend_low_i8x16_s(n98)
	n100 := Simd_i32x4_dot_i16x8_s(n94, n99)
	n101 := Simd_i32x4_add(n87, n100)
	n102 := Simd_i16x8_extend_high_i8x16_s(n98)
	n103 := Simd_i32x4_dot_i16x8_s(n95, n102)
	n104 := Simd_i32x4_add(n90, n103)
	n105 := Simd_m64_scalar_i32_add(s0, s2)
	n106 := Simd_m64_scalar_i32_add(n105, 8)
	n107 := Simd_m64_v128_load(m, n106, 0)
	n108 := Simd_i16x8_extend_low_i8x16_s(n107)
	n109 := Simd_i16x8_extend_high_i8x16_s(n107)
	n110 := Simd_m64_scalar_i32_add(s0, s1)
	n111 := Simd_m64_scalar_i32_add(n110, 20)
	n112 := Simd_m64_v128_load32_splat(m, n111, 0)
	n113 := Simd_i16x8_extend_low_i8x16_s(n112)
	n114 := Simd_i32x4_dot_i16x8_s(n108, n113)
	n115 := Simd_i32x4_add(n101, n114)
	n116 := Simd_i16x8_extend_high_i8x16_s(n112)
	n117 := Simd_i32x4_dot_i16x8_s(n109, n116)
	n118 := Simd_i32x4_add(n104, n117)
	n119 := Simd_i8x16_shuffle(n115, n118, [2]uint64{795458214199165184, 1952900979608391952})
	n120 := Simd_i8x16_shuffle(n115, n118, [2]uint64{1084818905551471876, 2242261670960698644})
	n121 := Simd_i32x4_add(n119, n120)
	n122 := Simd_f32x4_convert_i32x4_s(n121)
	n123 := Simd_f32x4_mul(n8, n122)
	n124 := Simd_f32x4_add([2]uint64{p0, p0h}, n123)
	return n124[0], n124[1]
}

//go:noinline
func Simd_p_fx821(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 144, n0)
	return
}

//go:noinline
func Simd_p_fx822(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i32x4_max_s([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx823(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{1084818905618843912, 216736831629295872})
	n1 := Simd_i32x4_max_s([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx824(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{216736831696667908, 216736831629295872})
	n1 := Simd_i32x4_max_s([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx825(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i64x2_ne(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n3 := Simd_i64x2_ne(n2, [2]uint64{p2, p2h})
	n4 := Simd_i8x16_shuffle(n1, n3, [2]uint64{795458214199165184, 1952900979608391952})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx826(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load(m, s1, 0)
	n2 := Simd_i32x4_eq(n0, n1)
	n3 := Simd_v128_and([2]uint64{p0, p0h}, n2)
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx827(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p0, p0h}, 31)
	n1 := Simd_i32x4_shr_s(n0, 31)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx828(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	_ = Simd_m64_v128_store(m, s0, 29288, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s0, 29328)
	_ = Simd_m64_v128_store(m, s1, 29328, n1)
	n3 := Simd_m64_v128_load(m, s0, 29312)
	_ = Simd_m64_v128_store(m, s1, 29312, n3)
	return
}

//go:noinline
func Simd_p_fx829(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 29480)
	_ = Simd_m64_v128_store(m, s1, 29480, n0)
	n2 := Simd_m64_v128_load(m, s0, 29464)
	_ = Simd_m64_v128_store(m, s1, 29464, n2)
	n4 := Simd_m64_v128_load(m, s0, 29448)
	_ = Simd_m64_v128_store(m, s1, 29448, n4)
	return
}

//go:noinline
func Simd_p_fx830(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 29536)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx831(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s0, 8, n0)
	return
}

//go:noinline
func Simd_p_fx832(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 29328)
	_ = Simd_m64_v128_store(m, s1, 29328, n0)
	n2 := Simd_m64_v128_load(m, s0, 29312)
	_ = Simd_m64_v128_store(m, s1, 29312, n2)
	return
}

//go:noinline
func Simd_p_fx833(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s1, 32, n0)
	return
}

//go:noinline
func Simd_p_fx834(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 29392, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 29376, n2)
	return
}

//go:noinline
func Simd_p_fx835(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 168, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 152, n2)
	return
}

//go:noinline
func Simd_p_fx836(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 352, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 336, n2)
	return
}

//go:noinline
func Simd_p_fx837(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n1 := Simd_i64x2_ne(n0, [2]uint64{p0, p0h})
	n2 := Simd_m64_v128_load_nc(m, s0, 16)
	n3 := Simd_i64x2_ne(n2, [2]uint64{p0, p0h})
	n4 := Simd_i8x16_shuffle(n1, n3, [2]uint64{795458214199165184, 1952900979608391952})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx838(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_ne([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i64x2_ne([2]uint64{p2, p2h}, [2]uint64{p1, p1h})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{795458214199165184, 1952900979608391952})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx839(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 64)
	_ = Simd_m64_v128_store(m, s0, 24, n0)
	return
}

//go:noinline
func Simd_p_fx840(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 64)
	_ = Simd_m64_v128_store(m, s0, 8, n0)
	return
}

//go:noinline
func Simd_p_fx841(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 272)
	_ = Simd_m64_v128_store(m, s0, 296, n0)
	return
}

//go:noinline
func Simd_p_fx842(m *Module, s0 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n1 := Simd_m64_v128_load_nc(m, s0, 16)
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{1084818905551471876, 2242261670960698644})
	n3 := Simd_i32x4_add(n2, [2]uint64{p0, p0h})
	n4 := Simd_i8x16_shuffle(n0, n1, [2]uint64{795458214199165184, 1952900979608391952})
	n5 := Simd_i32x4_sub(n3, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx843(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{1084818905618843912, 216736831629295872})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx844(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 160)
	_ = Simd_m64_v128_store(m, s0+152, 8, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	return
}

//go:noinline
func Simd_p_fx845(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 184)
	_ = Simd_m64_v128_store(m, s0, 40, n0)
	return
}

//go:noinline
func Simd_p_fx846(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 184)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx847(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	n2 := Simd_m64_v128_load(m, s0, 16)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx848(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	_ = Simd_m64_v128_store(m, s0, 16, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s1, 32)
	_ = Simd_m64_v128_store(m, s1, 32, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 32, n1)
	n4 := Simd_m64_v128_load(m, s1, 48)
	_ = Simd_m64_v128_store(m, s0, 48, n4)
	return
}

//go:noinline
func Simd_p_fx849(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 240)
	_ = Simd_m64_v128_store(m, s0, 24, n0)
	return
}

//go:noinline
func Simd_p_fx850(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_ne([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{795458214199165184, 216736831629295872})
	n2 := Simd_i32x4_sub([2]uint64{p0, p0h}, n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx851(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 152)
	_ = Simd_m64_v128_store(m, s0, 104, n0)
	return
}

//go:noinline
func Simd_p_fx852(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s1, 24, n0)
	return
}

//go:noinline
func Simd_p_fx853(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 48)
	_ = Simd_m64_v128_store(m, s0, 24, n0)
	return
}

//go:noinline
func Simd_p_fx854(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 80)
	_ = Simd_m64_v128_store(m, s0, 80, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 8, n0)
	return
}

//go:noinline
func Simd_p_fx855(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	_ = Simd_m64_v128_store(m, s0, 40, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s1, 16)
	_ = Simd_m64_v128_store(m, s0, 16, n1)
	n3 := Simd_m64_v128_load(m, s1, 0)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	return
}

//go:noinline
func Simd_p_fx856(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, s1)
	_ = Simd_m64_v128_store(m, s2, s1, n0)
	n2 := Simd_m64_v128_load(m, s0, s3)
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx857(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	_ = Simd_m64_v128_store(m, s0, s1, [2]uint64{p0, p0h})
	n1 := Simd_m64_v128_load(m, s2, s3)
	_ = Simd_m64_v128_store(m, s2, s3, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, s3, n1)
	n4 := Simd_m64_v128_load(m, s2, s4)
	_ = Simd_m64_v128_store(m, s0, s4, n4)
	return n1[0], n1[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx858(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 136, n0)
	n2 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 120, n2)
	return
}

//go:noinline
func Simd_p_fx859(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 88)
	_ = Simd_m64_v128_store(m, s0, 112, n0)
	return
}

//go:noinline
func Simd_p_fx860(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 80, n0)
	return
}

//go:noinline
func Simd_p_fx861(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 248, n0)
	return
}

//go:noinline
func Simd_p_fx862(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 272, n0)
	return
}

//go:noinline
func Simd_p_fx863(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 296, n0)
	return
}

//go:noinline
func Simd_p_fx864(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 80, n0)
	n2 := Simd_m64_v128_load(m, s2, 0)
	_ = Simd_m64_v128_store(m, s1, 64, n2)
	n4 := Simd_m64_v128_load(m, s3, 0)
	_ = Simd_m64_v128_store(m, s1, 48, n4)
	n6 := Simd_m64_v128_load(m, s4, 0)
	_ = Simd_m64_v128_store(m, s1, 32, n6)
	return
}

//go:noinline
func Simd_p_fx865(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 152)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	n2 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx866(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 304)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	return
}

//go:noinline
func Simd_p_fx867(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 208)
	_ = Simd_m64_v128_store(m, s1, 40, n0)
	return
}

//go:noinline
func Simd_p_fx868(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 88)
	_ = Simd_m64_v128_store(m, s0, 64, n0)
	n2 := Simd_m64_v128_load(m, s0, 400)
	_ = Simd_m64_v128_store(m, s0, 48, n2)
	return
}

//go:noinline
func Simd_p_fx869(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+96, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	n2 := Simd_m64_v128_load(m, s0+112, 0)
	_ = Simd_m64_v128_store(m, s2, 0, n2)
	n4 := Simd_m64_v128_load(m, s0, 128)
	_ = Simd_m64_v128_store(m, s3, 0, [2]uint64{p0, p0h})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx870(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	n2 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx871(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f64x2_div(n0, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 0, n1)
	return
}

//go:noinline
func Simd_p_fx872(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 72)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx873(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0+32, 0)
	_ = Simd_m64_v128_store(m, s1, 32, n0)
	n2 := Simd_m64_v128_load(m, s0+48, 0)
	_ = Simd_m64_v128_store(m, s1, 48, n2)
	return
}

//go:noinline
func Simd_p_fx874(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 40, n0)
	return
}

//go:noinline
func Simd_p_fx875(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx876(m *Module, s0 int64, s1 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i8x16_shuffle(n0, [2]uint64{p0, p0h}, [2]uint64{579005069656919567, 283686952306183})
	_ = Simd_m64_v128_store(m, s1, 0, n1)
	return
}

//go:noinline
func Simd_p_fx877(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 8)
	n1 := Simd_m64_v128_load(m, s1, 8)
	_ = Simd_m64_v128_store(m, s0, 8, n1)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx878(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_i8x16_add(n0, [2]uint64{p0, p0h})
	n2 := Simd_i8x16_lt_u(n1, [2]uint64{p1, p1h})
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx879(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1+16, 0, n0)
	return
}

//go:noinline
func Simd_p_fx880(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0+32, 0)
	_ = Simd_m64_v128_store(m, s1+32, 0, n0)
	return
}

//go:noinline
func Simd_p_fx881(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx882(m *Module, s0 int32, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(s0)
	n1 := Simd_i32x4_add(n0, [2]uint64{p0, p0h})
	n2 := Simd_i32x4_lt_u(n1, [2]uint64{p1, p1h})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx883(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 0)
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fx884(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 40)
	_ = Simd_m64_v128_store(m, s0, 8, n0)
	return
}

//go:noinline
func Simd_p_fx885(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 64)
	_ = Simd_m64_v128_store(m, s0, 16, n0)
	return
}

//go:noinline
func Simd_p_fx886(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	return
}

//go:noinline
func Simd_p_fx887(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 56)
	_ = Simd_m64_v128_store(m, s0, 8, n0)
	return
}

//go:noinline
func Simd_p_fx888(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 6412)
	_ = Simd_m64_v128_store(m, s1, 32, n0)
	return
}

//go:noinline
func Simd_p_fx889(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 6412)
	_ = Simd_m64_v128_store(m, s1, 48, n0)
	return
}

//go:noinline
func Simd_p_fx890(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i64x2_extend_low_i32x4_s([2]uint64{p1, p1h})
	n1 := Simd_i64x2_add(n0, [2]uint64{p0, p0h})
	n2 := Simd_m64_v128_load(m, s0, s1)
	n3 := Simd_i64x2_add([2]uint64{p0, p0h}, n2)
	_ = Simd_m64_v128_store(m, s2, s3, n3)
	n5 := Simd_m64_v128_load(m, s0, s3)
	n6 := Simd_i64x2_add(n1, n5)
	_ = Simd_m64_v128_store(m, s2, s4, n6)
	return n2[0], n2[1], n3[0], n3[1], n5[0], n5[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx891(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 6412)
	_ = Simd_m64_v128_store(m, s1, 128, n0)
	return
}

//go:noinline
func Simd_p_fx892(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 6412)
	_ = Simd_m64_v128_store(m, s1, 16, n0)
	return
}

//go:noinline
func Simd_p_fx893(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 6412)
	_ = Simd_m64_v128_store(m, s1, 144, n0)
	return
}

//go:noinline
func Simd_p_fx894(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_v128_xor([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 64, n0)
	return
}

//go:noinline
func Simd_p_fx895(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 130)
	n1 := Simd_m64_v128_load32_lane(m, s1, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s0+166, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s2, 0, 3, n2)
	n4 := Simd_v128_xor(n3, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 64, n4)
	return
}

//go:noinline
func Simd_p_fx896(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 134)
	n1 := Simd_m64_v128_load32_lane(m, s0+152, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s0+170, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s0+188, 0, 3, n2)
	n4 := Simd_v128_xor(n3, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 80, n4)
	return
}

//go:noinline
func Simd_p_fx897(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 138)
	n1 := Simd_m64_v128_load32_lane(m, s1, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s0+174, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s2, 0, 3, n2)
	n4 := Simd_v128_xor(n3, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 96, n4)
	return
}

//go:noinline
func Simd_p_fx898(m *Module, s0 int64, p0, p0h uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 142)
	n1 := Simd_m64_v128_load32_lane(m, s0+160, 0, 1, n0)
	n2 := Simd_m64_v128_load32_lane(m, s0+178, 0, 2, n1)
	n3 := Simd_m64_v128_load32_lane(m, s0+196, 0, 3, n2)
	n4 := Simd_v128_xor(n3, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s0, 112, n4)
	return
}

//go:noinline
func Simd_p_fx899(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_v128_xor([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 80, n0)
	return
}

//go:noinline
func Simd_p_fx900(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_v128_xor([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 96, n0)
	return
}

//go:noinline
func Simd_p_fx901(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_v128_xor([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 112, n0)
	return
}

//go:noinline
func Simd_p_fx902(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_v128_xor([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 128, n0)
	return
}

//go:noinline
func Simd_p_fx903(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_v128_xor([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 144, n0)
	return
}

//go:noinline
func Simd_p_fx904(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_v128_xor([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 160, n0)
	return
}

//go:noinline
func Simd_p_fx905(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_v128_xor([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s0, 176, n0)
	return
}

//go:noinline
func Simd_p_fx906(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	_ = Simd_m64_v128_store(m, s0, 116, n1)
	return
}

//go:noinline
func Simd_p_fx907(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s0, 56, n0)
	return
}

//go:noinline
func Simd_p_fx908(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 24)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx909(m *Module, s0 int64, s1 int64, s2 int64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s1, 40, n0)
	n2 := Simd_m64_v128_load(m, s0, 16)
	_ = Simd_m64_v128_store(m, s1, 24, n2)
	n4 := Simd_m64_v128_load(m, s0, s2)
	_ = Simd_m64_v128_store(m, s1, 8, n4)
	return n0[0], n0[1], n2[0], n2[1], n4[0], n4[1]
}

//go:noinline
func Simd_p_fx910(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s1, 80, n0)
	return
}

//go:noinline
func Simd_p_fx911(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_v128_load(m, s0, 192)
	n1 := Simd_i64x2_extend_high_i32x4_s(n0)
	n2 := Simd_i64x2_extend_low_i32x4_s(n0)
	_ = Simd_m64_v128_store(m, s1, 64, n1)
	_ = Simd_m64_v128_store(m, s1, 48, n2)
	n5 := Simd_m64_v128_load(m, s0, 208)
	n6 := Simd_i64x2_extend_high_i32x4_s(n5)
	n7 := Simd_i64x2_extend_low_i32x4_s(n5)
	_ = Simd_m64_v128_store(m, s2, 32, n6)
	_ = Simd_m64_v128_store(m, s2, 16, n7)
	return
}

//go:noinline
func Simd_p_fx912(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s1, 144, n0)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx913(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 144)
	_ = Simd_m64_v128_store(m, s1, 128, n0)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx914(m *Module, s0 int64, s1 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 32)
	_ = Simd_m64_v128_store(m, s1, 64, n0)
	n2 := Simd_m64_v128_load(m, s0, 16)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx915(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_convert_i32x4_s([2]uint64{p2, p2h})
	n2 := Simd_f32x4_sub(n0, n1)
	n3 := Simd_f32x4_convert_i32x4_s([2]uint64{p3, p3h})
	n4 := Simd_f32x4_div(n2, n3)
	n5 := Simd_f32x4_add(n4, [2]uint64{p4, p4h})
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx916(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_i64x2_extend_low_i32x4_s([2]uint64{p1, p1h})
	n1 := Simd_m64_v128_load(m, s0, 32)
	n2 := Simd_i64x2_add(n1, [2]uint64{p0, p0h})
	_ = Simd_m64_v128_store(m, s1, 16, n2)
	n4 := Simd_m64_v128_load(m, s0, 16)
	n5 := Simd_i64x2_add(n4, n0)
	_ = Simd_m64_v128_store(m, s1, 0, n5)
	return
}

//go:noinline
func Simd_p_fx917(m *Module, s0 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s0+16, 0, n0)
	return
}

//go:noinline
func Simd_p_fx918(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0+40, 0)
	_ = Simd_m64_v128_store(m, s1, 0, n0)
	return
}

//go:noinline
func Simd_p_fx919(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	_ = Simd_m64_v128_store(m, s0+16, 0, n0)
	n2 := Simd_m64_v128_load(m, s1, 0)
	_ = Simd_m64_v128_store(m, s1+16, 0, n2)
	n4 := Simd_m64_v128_load(m, s2, 0)
	_ = Simd_m64_v128_store(m, s2+16, 0, n4)
	n6 := Simd_m64_v128_load(m, s3, 0)
	_ = Simd_m64_v128_store(m, s3+16, 0, n6)
	return
}

//go:noinline
func Simd_p_fx920(m *Module, s0 int64, f0 float32, f1 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_eq(n0, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_splat(f1)
	n3 := Simd_f32x4_splat(f0)
	n4 := Simd_v128_bitselect(n3, n2, n1)
	_ = Simd_m64_v128_store(m, s0, 0, n4)
	return
}

//go:noinline
func Simd_p_fx921(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p0, p0h}, 1)
	n1 := Simd_v128_and(n0, [2]uint64{p1, p1h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx922(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i32x4_splat(int32(s0))
	n1 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_i32x4_add(n1, [2]uint64{p2, p2h})
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul(n0, n3)
	n5 := Simd_m64_scalar_i32_add(s1, s2)
	_ = Simd_m64_v128_store(m, n5, 0, n4)
	return
}

//go:noinline
func Simd_p_fx923(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(int32(s0))
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_add(n1, n0)
	n3 := Simd_f32x4_mul([2]uint64{p2, p2h}, [2]uint64{p1, p1h})
	n4 := Simd_f32x4_add(n3, n0)
	n5 := Simd_f32x4_mul([2]uint64{p3, p3h}, [2]uint64{p1, p1h})
	n6 := Simd_f32x4_add(n5, n0)
	n7 := Simd_f32x4_mul([2]uint64{p4, p4h}, [2]uint64{p1, p1h})
	n8 := Simd_f32x4_add(n7, n0)
	_ = Simd_m64_v128_store(m, s1, 80, n2)
	_ = Simd_m64_v128_store(m, s1, 64, n4)
	_ = Simd_m64_v128_store(m, s1, 16, n6)
	_ = Simd_m64_v128_store(m, s1, 0, n8)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx924(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_add(n0, [2]uint64{p2, p2h})
	n2 := Simd_f32x4_mul([2]uint64{p3, p3h}, [2]uint64{p1, p1h})
	n3 := Simd_f32x4_add(n2, [2]uint64{p2, p2h})
	n4 := Simd_f32x4_mul([2]uint64{p4, p4h}, [2]uint64{p1, p1h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p2, p2h})
	n6 := Simd_f32x4_mul([2]uint64{p5, p5h}, [2]uint64{p1, p1h})
	n7 := Simd_f32x4_add(n6, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s0, 112, n1)
	_ = Simd_m64_v128_store(m, s0, 96, n3)
	_ = Simd_m64_v128_store(m, s0, 48, n5)
	_ = Simd_m64_v128_store(m, s0, 32, n7)
	return
}

//go:noinline
func Simd_p_fx925(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p2, p2h})
	n2 := Simd_m64_v128_load32_zero(m, s0+6, 0)
	n3 := Simd_i16x8_extend_low_i8x16_u(n2)
	n4 := Simd_i32x4_extend_low_i16x8_u(n3)
	n5 := Simd_i32x4_shr_u(n4, 4)
	n6 := Simd_v128_or(n0, n5)
	n7 := Simd_i32x4_add(n6, [2]uint64{p3, p3h})
	n8 := Simd_f32x4_convert_i32x4_s(n7)
	n9 := Simd_f32x4_mul([2]uint64{p0, p0h}, n8)
	n10 := Simd_v128_and(n4, [2]uint64{p5, p5h})
	n11 := Simd_v128_or(n1, n10)
	n12 := Simd_i32x4_add(n11, [2]uint64{p3, p3h})
	n13 := Simd_f32x4_convert_i32x4_s(n12)
	n14 := Simd_f32x4_mul([2]uint64{p0, p0h}, n13)
	_ = Simd_m64_v128_store(m, s1, 0, n9)
	_ = Simd_m64_v128_store(m, s2, 0, n14)
	n17 := Simd_m64_v128_load32_zero(m, s0+10, 0)
	n18 := Simd_i16x8_extend_low_i8x16_u(n17)
	n19 := Simd_i32x4_extend_low_i16x8_u(n18)
	return n19[0], n19[1]
}

//go:noinline
func Simd_p_fx926(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i32x4_shr_u([2]uint64{p3, p3h}, 4)
	n2 := Simd_v128_or(n0, n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p4, p4h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	n6 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s0+80, 0, n5)
	n8 := Simd_m64_v128_load32_zero(m, s1+14, 0)
	n9 := Simd_i16x8_extend_low_i8x16_u(n8)
	n10 := Simd_i32x4_extend_low_i16x8_u(n9)
	n11 := Simd_i32x4_shr_u(n10, 4)
	n12 := Simd_v128_or(n6, n11)
	n13 := Simd_i32x4_add(n12, [2]uint64{p4, p4h})
	n14 := Simd_f32x4_convert_i32x4_s(n13)
	n15 := Simd_f32x4_mul([2]uint64{p0, p0h}, n14)
	_ = Simd_m64_v128_store(m, s0+96, 0, n15)
	return n10[0], n10[1]
}

//go:noinline
func Simd_p_fx927(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n2 := Simd_v128_or(n0, n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p5, p5h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+32, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1+18, 0)
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx928(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i32x4_shr_u([2]uint64{p3, p3h}, 4)
	n2 := Simd_v128_or(n0, n1)
	n3 := Simd_i32x4_add(n2, [2]uint64{p4, p4h})
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s0+112, 0, n5)
	return
}

//go:noinline
func Simd_p_fx929(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p1, p1h}, 4)
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_i32x4_add(n3, [2]uint64{p5, p5h})
	n5 := Simd_f32x4_convert_i32x4_s(n4)
	n6 := Simd_f32x4_mul([2]uint64{p0, p0h}, n5)
	_ = Simd_m64_v128_store(m, s0+16, 0, n6)
	return
}

//go:noinline
func Simd_p_fx930(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p1, p1h}, 4)
	n1 := Simd_v128_and(n0, [2]uint64{p2, p2h})
	n2 := Simd_v128_and([2]uint64{p3, p3h}, [2]uint64{p4, p4h})
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_i32x4_add(n3, [2]uint64{p5, p5h})
	n5 := Simd_f32x4_convert_i32x4_s(n4)
	n6 := Simd_f32x4_mul([2]uint64{p0, p0h}, n5)
	_ = Simd_m64_v128_store(m, s0+48, 0, n6)
	return
}

//go:noinline
func Simd_p_fx931(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0+8, 0)
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i32x4_extend_low_i16x8_u(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx932(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_splat(int32(s0))
	n1 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 4)
	n2 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n3 := Simd_v128_or(n1, n2)
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul(n4, [2]uint64{p3, p3h})
	n6 := Simd_f32x4_add(n5, n0)
	_ = Simd_m64_v128_store(m, s1, 64, n6)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx933(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_v128_or(n0, n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul(n3, [2]uint64{p4, p4h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p5, p5h})
	_ = Simd_m64_v128_store(m, s0, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1+12, 0)
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx934(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 4)
	n1 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n0, n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul(n3, [2]uint64{p3, p3h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p4, p4h})
	n6 := Simd_v128_and([2]uint64{p5, p5h}, [2]uint64{p2, p2h})
	_ = Simd_m64_v128_store(m, s0, 80, n5)
	n8 := Simd_m64_v128_load32_zero(m, s1+16, 0)
	n9 := Simd_i16x8_extend_low_i8x16_u(n8)
	n10 := Simd_i32x4_extend_low_i16x8_u(n9)
	n11 := Simd_i32x4_shr_u(n10, 4)
	n12 := Simd_v128_or(n11, n6)
	n13 := Simd_f32x4_convert_i32x4_s(n12)
	n14 := Simd_f32x4_mul(n13, [2]uint64{p3, p3h})
	n15 := Simd_f32x4_add(n14, [2]uint64{p4, p4h})
	_ = Simd_m64_v128_store(m, s0, 96, n15)
	return n10[0], n10[1]
}

//go:noinline
func Simd_p_fx935(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_v128_or(n0, n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul(n3, [2]uint64{p4, p4h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p5, p5h})
	_ = Simd_m64_v128_store(m, s0, 32, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1+20, 0)
	n8 := Simd_i16x8_extend_low_i8x16_u(n7)
	n9 := Simd_i32x4_extend_low_i16x8_u(n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx936(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) {
	n0 := Simd_i32x4_shr_u([2]uint64{p0, p0h}, 4)
	n1 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n2 := Simd_v128_or(n0, n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_mul(n3, [2]uint64{p3, p3h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p4, p4h})
	_ = Simd_m64_v128_store(m, s0, 112, n5)
	return
}

//go:noinline
func Simd_p_fx937(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_shl([2]uint64{p2, p2h}, 4)
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_or(n0, n2)
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul(n4, [2]uint64{p4, p4h})
	n6 := Simd_f32x4_add(n5, [2]uint64{p5, p5h})
	_ = Simd_m64_v128_store(m, s0, 16, n6)
	return
}

//go:noinline
func Simd_p_fx938(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i32x4_shl([2]uint64{p2, p2h}, 4)
	n2 := Simd_v128_and(n1, [2]uint64{p3, p3h})
	n3 := Simd_v128_or(n0, n2)
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul(n4, [2]uint64{p4, p4h})
	n6 := Simd_f32x4_add(n5, [2]uint64{p5, p5h})
	_ = Simd_m64_v128_store(m, s0, 48, n6)
	return
}

//go:noinline
func Simd_p_fx939(m *Module, s0 int64, s1 int64, s2 int64, s3 int64) {
	n0 := Simd_i32x4_splat(int32(s0))
	n1 := Simd_m64_v128_load32_zero(m, s1+2, 0)
	n2 := Simd_i16x8_extend_low_i8x16_s(n1)
	n3 := Simd_i32x4_extend_low_i16x8_s(n2)
	n4 := Simd_f32x4_convert_i32x4_s(n3)
	n5 := Simd_f32x4_mul(n0, n4)
	_ = Simd_m64_v128_store(m, s2, 0, n5)
	n7 := Simd_m64_v128_load32_zero(m, s1+6, 0)
	n8 := Simd_i16x8_extend_low_i8x16_s(n7)
	n9 := Simd_i32x4_extend_low_i16x8_s(n8)
	n10 := Simd_f32x4_convert_i32x4_s(n9)
	n11 := Simd_f32x4_mul(n0, n10)
	_ = Simd_m64_v128_store(m, s2+16, 0, n11)
	n13 := Simd_m64_v128_load32_zero(m, s1+10, 0)
	n14 := Simd_i16x8_extend_low_i8x16_s(n13)
	n15 := Simd_i32x4_extend_low_i16x8_s(n14)
	n16 := Simd_f32x4_convert_i32x4_s(n15)
	n17 := Simd_f32x4_mul(n0, n16)
	_ = Simd_m64_v128_store(m, s2+32, 0, n17)
	n19 := Simd_m64_v128_load32_zero(m, s1+14, 0)
	n20 := Simd_i16x8_extend_low_i8x16_s(n19)
	n21 := Simd_i32x4_extend_low_i16x8_s(n20)
	n22 := Simd_f32x4_convert_i32x4_s(n21)
	n23 := Simd_f32x4_mul(n0, n22)
	_ = Simd_m64_v128_store(m, s2+48, 0, n23)
	n25 := Simd_m64_v128_load32_zero(m, s1+18, 0)
	n26 := Simd_i16x8_extend_low_i8x16_s(n25)
	n27 := Simd_i32x4_extend_low_i16x8_s(n26)
	n28 := Simd_f32x4_convert_i32x4_s(n27)
	n29 := Simd_f32x4_mul(n0, n28)
	_ = Simd_m64_v128_store(m, s3, 0, n29)
	n31 := Simd_m64_v128_load32_zero(m, s1+22, 0)
	n32 := Simd_i16x8_extend_low_i8x16_s(n31)
	n33 := Simd_i32x4_extend_low_i16x8_s(n32)
	n34 := Simd_f32x4_convert_i32x4_s(n33)
	n35 := Simd_f32x4_mul(n0, n34)
	_ = Simd_m64_v128_store(m, s2+80, 0, n35)
	n37 := Simd_m64_v128_load32_zero(m, s1+26, 0)
	n38 := Simd_i16x8_extend_low_i8x16_s(n37)
	n39 := Simd_i32x4_extend_low_i16x8_s(n38)
	n40 := Simd_f32x4_convert_i32x4_s(n39)
	n41 := Simd_f32x4_mul(n0, n40)
	_ = Simd_m64_v128_store(m, s2+96, 0, n41)
	n43 := Simd_m64_v128_load32_zero(m, s1+30, 0)
	n44 := Simd_i16x8_extend_low_i8x16_s(n43)
	n45 := Simd_i32x4_extend_low_i16x8_s(n44)
	n46 := Simd_f32x4_convert_i32x4_s(n45)
	n47 := Simd_f32x4_mul(n0, n46)
	_ = Simd_m64_v128_store(m, s2+112, 0, n47)
	return
}

//go:noinline
func Simd_p_fx940(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_splat(int32(s0))
	n1 := Simd_m64_v128_load32_zero(m, s1+1, 0)
	n2 := Simd_i16x8_extend_low_i8x16_u(n1)
	n3 := Simd_i32x4_extend_low_i16x8_u(n2)
	n4 := Simd_i32x4_shr_u(n3, 4)
	n5 := Simd_i64x2_extend_high_i32x4_u(n4)
	n6 := Simd_i64x2_extend_low_i32x4_u(n4)
	return n0[0], n0[1], n3[0], n3[1], n5[0], n5[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx941(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_i64x2_extend_high_i32x4_u(n4)
	n6 := Simd_i64x2_extend_low_i32x4_u(n4)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	return n5[0], n5[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx942(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	n5 := Simd_m64_v128_load32_zero(m, s1+5, 0)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_i32x4_shr_u(n7, 4)
	n9 := Simd_i64x2_extend_high_i32x4_u(n8)
	n10 := Simd_i64x2_extend_low_i32x4_u(n8)
	return n7[0], n7[1], n9[0], n9[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx943(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_i64x2_extend_high_i32x4_u(n4)
	n6 := Simd_i64x2_extend_low_i32x4_u(n4)
	_ = Simd_m64_v128_store(m, s0+80, 0, n3)
	return n5[0], n5[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx944(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	_ = Simd_m64_v128_store(m, s0+16, 0, n3)
	n5 := Simd_m64_v128_load32_zero(m, s1+9, 0)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_i32x4_shr_u(n7, 4)
	n9 := Simd_i64x2_extend_high_i32x4_u(n8)
	n10 := Simd_i64x2_extend_low_i32x4_u(n8)
	return n7[0], n7[1], n9[0], n9[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx945(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_i64x2_extend_high_i32x4_u(n4)
	n6 := Simd_i64x2_extend_low_i32x4_u(n4)
	_ = Simd_m64_v128_store(m, s0+96, 0, n3)
	return n5[0], n5[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx946(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	_ = Simd_m64_v128_store(m, s0+32, 0, n3)
	n5 := Simd_m64_v128_load32_zero(m, s1+13, 0)
	n6 := Simd_i16x8_extend_low_i8x16_u(n5)
	n7 := Simd_i32x4_extend_low_i16x8_u(n6)
	n8 := Simd_i32x4_shr_u(n7, 4)
	n9 := Simd_i64x2_extend_high_i32x4_u(n8)
	n10 := Simd_i64x2_extend_low_i32x4_u(n8)
	return n7[0], n7[1], n9[0], n9[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx947(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n5 := Simd_i64x2_extend_high_i32x4_u(n4)
	n6 := Simd_i64x2_extend_low_i32x4_u(n4)
	_ = Simd_m64_v128_store(m, s0+112, 0, n3)
	return n5[0], n5[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx948(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	_ = Simd_m64_v128_store(m, s0+48, 0, n3)
	return
}

//go:noinline
func Simd_p_fx949(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 8553172)
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_add([2]uint64{p1, p1h}, n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s1+16, 0, n5)
	return
}

//go:noinline
func Simd_p_fx950(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_add([2]uint64{p1, p1h}, n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	_ = Simd_m64_v128_store(m, s0, 0, n4)
	return
}

//go:noinline
func Simd_p_fx951(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 8553172)
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_add([2]uint64{p1, p1h}, n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s1+48, 0, n5)
	return
}

//go:noinline
func Simd_p_fx952(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_add([2]uint64{p1, p1h}, n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	_ = Simd_m64_v128_store(m, s0+32, 0, n4)
	return
}

//go:noinline
func Simd_p_fx953(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 8553172)
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_add([2]uint64{p1, p1h}, n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s1+80, 0, n5)
	return
}

//go:noinline
func Simd_p_fx954(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load32_zero(m, s0, 8553172)
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i32x4_extend_low_i16x8_s(n1)
	n3 := Simd_f32x4_convert_i32x4_s(n2)
	n4 := Simd_f32x4_add([2]uint64{p1, p1h}, n3)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	_ = Simd_m64_v128_store(m, s1+112, 0, n5)
	return
}

//go:noinline
func Simd_p_fx955(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_add([2]uint64{p1, p1h}, n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	_ = Simd_m64_v128_store(m, s0+96, 0, n4)
	return
}

//go:noinline
func Simd_p_fx956(m *Module, s0 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n1 := Simd_i32x4_extend_low_i16x8_s(n0)
	n2 := Simd_f32x4_convert_i32x4_s(n1)
	n3 := Simd_f32x4_add([2]uint64{p1, p1h}, n2)
	n4 := Simd_f32x4_mul([2]uint64{p0, p0h}, n3)
	n5 := Simd_f32x4_splat(f0)
	_ = Simd_m64_v128_store(m, s0, 0, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx957(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 40)
	_ = Simd_m64_v128_store(m, s1, 32, n0)
	n2 := Simd_m64_v128_load(m, s0, 24)
	_ = Simd_m64_v128_store(m, s1, 16, n2)
	n4 := Simd_m64_v128_load(m, s0, 8)
	_ = Simd_m64_v128_store(m, s1, 0, n4)
	return
}

//go:noinline
func Simd_p_fx958(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_scalar_i32_add(n0, 24)
	n2 := Simd_m64_v128_load16x4_u(m, n1, 0)
	n3 := Simd_f16x4_cvt(n2)
	n4 := Simd_m64_scalar_i32_add(s2, s1)
	n5 := Simd_m64_scalar_i32_add(n4, 24)
	n6 := Simd_m64_v128_load16x4_u(m, n5, 0)
	n7 := Simd_f16x4_cvt(n6)
	n8 := Simd_f32x4_mul(n3, n7)
	n9 := Simd_f32x4_add([2]uint64{p0, p0h}, n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx959(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_scalar_i32_add(n0, 16)
	n2 := Simd_m64_v128_load16x4_u(m, n1, 0)
	n3 := Simd_f16x4_cvt(n2)
	n4 := Simd_m64_scalar_i32_add(s2, s1)
	n5 := Simd_m64_scalar_i32_add(n4, 16)
	n6 := Simd_m64_v128_load16x4_u(m, n5, 0)
	n7 := Simd_f16x4_cvt(n6)
	n8 := Simd_f32x4_mul(n3, n7)
	n9 := Simd_f32x4_add([2]uint64{p0, p0h}, n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx960(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_scalar_i32_add(n0, 8)
	n2 := Simd_m64_v128_load16x4_u(m, n1, 0)
	n3 := Simd_f16x4_cvt(n2)
	n4 := Simd_m64_scalar_i32_add(s2, s1)
	n5 := Simd_m64_scalar_i32_add(n4, 8)
	n6 := Simd_m64_v128_load16x4_u(m, n5, 0)
	n7 := Simd_f16x4_cvt(n6)
	n8 := Simd_f32x4_mul(n3, n7)
	n9 := Simd_f32x4_add([2]uint64{p0, p0h}, n8)
	return n9[0], n9[1]
}

//go:noinline
func Simd_p_fx961(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load16x4_u(m, n0, 0)
	n2 := Simd_f16x4_cvt(n1)
	n3 := Simd_m64_scalar_i32_add(s2, s1)
	n4 := Simd_m64_v128_load16x4_u(m, n3, 0)
	n5 := Simd_f16x4_cvt(n4)
	n6 := Simd_f32x4_mul(n2, n5)
	n7 := Simd_f32x4_add([2]uint64{p0, p0h}, n6)
	return n7[0], n7[1]
}

//go:noinline
func Simd_p_fx962(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_add([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_abs([2]uint64{p2, p2h})
	n2 := Simd_f32x4_gt(n1, [2]uint64{p3, p3h})
	n3 := Simd_m64_v128_load(m, s0, 0)
	return n0[0], n0[1], n3[0], n3[1], n1[0], n1[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx963(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) {
	n0 := Simd_f32x4_add([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n1 := Simd_f32x4_div([2]uint64{p1, p1h}, n0)
	n2 := Simd_f32x4_mul([2]uint64{p0, p0h}, n1)
	_ = Simd_m64_v128_store(m, s0, 0, n2)
	return
}

//go:noinline
func Simd_p_fx964(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) {
	n0 := Simd_f32x4_add([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_f32x4_div([2]uint64{p0, p0h}, n0)
	n2 := Simd_m64_v128_load(m, s0, 0)
	n3 := Simd_f32x4_mul(n2, n1)
	_ = Simd_m64_v128_store(m, s1, 0, n3)
	return
}

//go:noinline
func Simd_p_fx965(m *Module, s0 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_splat(f0)
	n2 := Simd_f32x4_sub(n0, n1)
	n3 := Simd_f32x4_mul(n2, [2]uint64{p0, p0h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p1, p1h})
	n5 := Simd_f32x4_add(n4, [2]uint64{p2, p2h})
	n6 := Simd_f32x4_mul(n5, [2]uint64{p3, p3h})
	n7 := Simd_f32x4_add(n2, n6)
	n8 := Simd_f32x4_mul(n5, [2]uint64{p4, p4h})
	n9 := Simd_f32x4_add(n7, n8)
	n10 := Simd_f32x4_mul(n9, n9)
	return n4[0], n4[1], n5[0], n5[1], n9[0], n9[1], n10[0], n10[1]
}

//go:noinline
func Simd_p_fx966(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p3, p3h})
	n2 := Simd_f32x4_add(n1, [2]uint64{p4, p4h})
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p5, p5h})
	n4 := Simd_f32x4_add(n3, [2]uint64{p6, p6h})
	n5 := Simd_f32x4_mul([2]uint64{p2, p2h}, n4)
	n6 := Simd_f32x4_add(n2, n5)
	n7 := Simd_f32x4_mul([2]uint64{p2, p2h}, n6)
	n8 := Simd_f32x4_add(n0, n7)
	return n8[0], n8[1]
}

//go:noinline
func Simd_p_fx967(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i32x4_shl([2]uint64{p0, p0h}, 23)
	n1 := Simd_i32x4_add(n0, [2]uint64{p1, p1h})
	n2 := Simd_f32x4_abs([2]uint64{p2, p2h})
	n3 := Simd_f32x4_gt(n2, [2]uint64{p3, p3h})
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1], n3[0], n3[1]
}

//go:noinline
func Simd_p_fx968(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n2 := Simd_f16x4_cvt(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, n2)
	n4 := Simd_f32x4_add(n0, n3)
	_ = Simd_m64_v128_store(m, s0, 0, n4)
	n6 := Simd_m64_v128_load(m, s0+16, 0)
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx969(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s2+8, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_f32x4_mul([2]uint64{p1, p1h}, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	n5 := Simd_m64_v128_load(m, s1+32, 0)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx970(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s2+16, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_f32x4_mul([2]uint64{p1, p1h}, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	n5 := Simd_m64_v128_load(m, s1+48, 0)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx971(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_m64_v128_load16x4_u(m, s1+24, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_f32x4_mul([2]uint64{p1, p1h}, n1)
	n3 := Simd_f32x4_add([2]uint64{p0, p0h}, n2)
	_ = Simd_m64_v128_store(m, s0, 0, n3)
	return
}

//go:noinline
func Simd_p_fx972(m *Module, s0 int64, s1 int64, f0 float32) {
	n0 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n1 := Simd_f16x4_cvt(n0)
	n2 := Simd_f32x4_splat(f0)
	n3 := Simd_f32x4_mul(n1, n2)
	n4 := Simd_m64_v128_load(m, s0, 0)
	n5 := Simd_f32x4_add(n3, n4)
	_ = Simd_m64_v128_store(m, s0, 0, n5)
	return
}

//go:noinline
func Simd_p_fx973(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_abs([2]uint64{p0, p0h})
	n1 := Simd_f32x4_abs([2]uint64{p1, p1h})
	n2 := Simd_f32x4_max(n0, n1)
	n3 := Simd_f32x4_abs([2]uint64{p2, p2h})
	n4 := Simd_f32x4_abs([2]uint64{p3, p3h})
	n5 := Simd_f32x4_max(n3, n4)
	n6 := Simd_f32x4_max(n2, n5)
	n7 := Simd_f32x4_abs([2]uint64{p4, p4h})
	n8 := Simd_m64_v128_load_nc(m, s0+80, 0)
	n9 := Simd_f32x4_abs(n8)
	n10 := Simd_f32x4_max(n7, n9)
	n11 := Simd_m64_v128_load_nc(m, s0+96, 0)
	n12 := Simd_f32x4_abs(n11)
	n13 := Simd_m64_v128_load_nc(m, s0+112, 0)
	n14 := Simd_f32x4_abs(n13)
	n15 := Simd_f32x4_max(n12, n14)
	n16 := Simd_f32x4_max(n10, n15)
	n17 := Simd_f32x4_max(n6, n16)
	return n8[0], n8[1], n11[0], n11[1], n13[0], n13[1], n17[0], n17[1]
}

//go:noinline
func Simd_p_fx974(m *Module, s0 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_mul(n0, [2]uint64{p0, p0h})
	n2 := Simd_f32x4_nearest(n1)
	n3 := Simd_i32x4_trunc_sat_f32x4_s(n2)
	n4 := Simd_f32x4_mul(n0, [2]uint64{p1, p1h})
	n5 := Simd_f32x4_nearest(n4)
	n6 := Simd_i32x4_trunc_sat_f32x4_s(n5)
	n7 := Simd_i8x16_shuffle(n3, n6, [2]uint64{p2, p2h})
	n8 := Simd_f32x4_mul(n0, [2]uint64{p3, p3h})
	n9 := Simd_f32x4_nearest(n8)
	n10 := Simd_i32x4_trunc_sat_f32x4_s(n9)
	n11 := Simd_f32x4_mul(n0, [2]uint64{p4, p4h})
	n12 := Simd_f32x4_nearest(n11)
	n13 := Simd_i32x4_trunc_sat_f32x4_s(n12)
	n14 := Simd_i8x16_shuffle(n10, n13, [2]uint64{p5, p5h})
	n15 := Simd_i8x16_shuffle(n7, n14, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+18, 0, n15)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx975(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_nearest(n0)
	n2 := Simd_i32x4_trunc_sat_f32x4_s(n1)
	n3 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p2, p2h})
	n4 := Simd_f32x4_nearest(n3)
	n5 := Simd_i32x4_trunc_sat_f32x4_s(n4)
	n6 := Simd_i8x16_shuffle(n2, n5, [2]uint64{p3, p3h})
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx976(m *Module, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_f32x4_splat(f0)
	n1 := Simd_f32x4_mul(n0, [2]uint64{p0, p0h})
	n2 := Simd_f32x4_nearest(n1)
	n3 := Simd_i32x4_trunc_sat_f32x4_s(n2)
	n4 := Simd_f32x4_mul(n0, [2]uint64{p1, p1h})
	n5 := Simd_f32x4_nearest(n4)
	n6 := Simd_i32x4_trunc_sat_f32x4_s(n5)
	n7 := Simd_f32x4_mul(n0, [2]uint64{p2, p2h})
	n8 := Simd_f32x4_nearest(n7)
	n9 := Simd_i32x4_trunc_sat_f32x4_s(n8)
	return n0[0], n0[1], n3[0], n3[1], n6[0], n6[1], n9[0], n9[1]
}

//go:noinline
func Simd_p_fx977(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_f32x4_mul([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_nearest(n0)
	n2 := Simd_i32x4_trunc_sat_f32x4_s(n1)
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx978(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i8x16_shuffle([2]uint64{p3, p3h}, [2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+20, 0, n2)
	return
}

//go:noinline
func Simd_p_fx979(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_i8x16_shuffle([2]uint64{p3, p3h}, [2]uint64{p4, p4h}, [2]uint64{p5, p5h})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{p6, p6h})
	_ = Simd_m64_v128_store(m, s0+4, 0, n2)
	return
}

//go:noinline
func Simd_p_fx980(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64, p6, p6h uint64) (uint64, uint64) {
	n0 := Simd_i32x4_add([2]uint64{p5, p5h}, [2]uint64{p6, p6h})
	n1 := Simd_i32x4_add([2]uint64{p4, p4h}, n0)
	n2 := Simd_i32x4_add([2]uint64{p3, p3h}, n1)
	n3 := Simd_i32x4_add([2]uint64{p2, p2h}, n2)
	n4 := Simd_i32x4_add([2]uint64{p1, p1h}, n3)
	n5 := Simd_i32x4_add([2]uint64{p0, p0h}, n4)
	return n5[0], n5[1]
}

//go:noinline
func Simd_p_fx981(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_scalar_i32_shl(s1, 10)
	n1 := Simd_m64_scalar_i32_add(s0, n0)
	n2 := Simd_m64_v128_load(m, n1, 0)
	return n2[0], n2[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx982(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 0)
	n1 := Simd_f32x4_pmin([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_pmax([2]uint64{p1, p1h}, n0)
	n3 := Simd_m64_v128_load(m, s1, 0)
	n4 := Simd_f32x4_pmin(n1, n3)
	n5 := Simd_f32x4_pmax(n2, n3)
	n6 := Simd_m64_v128_load(m, s2, 0)
	n7 := Simd_f32x4_pmin(n4, n6)
	n8 := Simd_f32x4_pmax(n5, n6)
	return n7[0], n7[1], n8[0], n8[1]
}

//go:noinline
func Simd_p_fx983(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_pmin([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx984(m *Module, p0, p0h uint64, p1, p1h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_f32x4_pmax([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx985(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 0, 0, 64)
	n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
	n2 := Simd_f32x4_nearest(n1)
	n3 := Simd_i32x4_trunc_sat_f32x4_s(n2)
	n4 := Simd_m64_v128_load_nc(m, s0+16, 0)
	n5 := Simd_f32x4_mul([2]uint64{p0, p0h}, n4)
	n6 := Simd_f32x4_nearest(n5)
	n7 := Simd_i32x4_trunc_sat_f32x4_s(n6)
	n8 := Simd_i16x8_narrow_i32x4_s(n3, n7)
	n9 := Simd_m64_v128_load_nc(m, s0+32, 0)
	n10 := Simd_f32x4_mul([2]uint64{p0, p0h}, n9)
	n11 := Simd_f32x4_nearest(n10)
	n12 := Simd_i32x4_trunc_sat_f32x4_s(n11)
	n13 := Simd_m64_v128_load_nc(m, s0+48, 0)
	n14 := Simd_f32x4_mul([2]uint64{p0, p0h}, n13)
	n15 := Simd_f32x4_nearest(n14)
	n16 := Simd_i32x4_trunc_sat_f32x4_s(n15)
	n17 := Simd_i16x8_narrow_i32x4_s(n12, n16)
	n18 := Simd_i8x16_narrow_i16x8_s(n8, n17)
	n19 := Simd_i16x8_extend_low_i8x16_s(n18)
	n20 := Simd_i16x8_extend_high_i8x16_s(n18)
	n21 := Simd_i16x8_add(n19, n20)
	n22 := Simd_i8x16_shuffle(n21, [2]uint64{p0, p0h}, [2]uint64{1084818905618843912, 72058693566333184})
	n23 := Simd_i16x8_add(n21, n22)
	n24 := Simd_i32x4_extend_low_i16x8_s(n23)
	n25 := Simd_i8x16_shuffle(n24, n24, [2]uint64{1084818905618843912, 506097522914230528})
	n26 := Simd_i32x4_add(n25, n24)
	_ = Simd_m64_v128_store(m, s1, 0, n18)
	return n26[0], n26[1]
}

//go:noinline
func Simd_p_fx986(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{216736831696667908, 795458214401281292})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx987(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0+4, 0)
	n1 := Simd_v128_and(n0, [2]uint64{p0, p0h})
	n2 := Simd_m64_v128_load_rng(m, s1+4, 0, 0, 32)
	return n0[0], n0[1], n1[0], n1[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx988(m *Module, s0 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p2, p2h})
	n2 := Simd_i32x4_dot_i16x8_s(n0, n1)
	n3 := Simd_i16x8_extend_high_i8x16_u([2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p2, p2h})
	n5 := Simd_i32x4_dot_i16x8_s(n3, n4)
	n6 := Simd_i32x4_add(n2, n5)
	n7 := Simd_i16x8_extend_high_i8x16_u([2]uint64{p3, p3h})
	n8 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p3, p3h})
	n9 := Simd_m64_v128_load_nc(m, s0+20, 0)
	n10 := Simd_i16x8_extend_high_i8x16_s(n9)
	n11 := Simd_i32x4_dot_i16x8_s(n7, n10)
	n12 := Simd_i32x4_add(n6, n11)
	n13 := Simd_i16x8_extend_low_i8x16_s(n9)
	n14 := Simd_i32x4_dot_i16x8_s(n8, n13)
	n15 := Simd_i32x4_add(n12, n14)
	n16 := Simd_f32x4_convert_i32x4_s(n15)
	n17 := Simd_f32x4_splat(f0)
	n18 := Simd_f32x4_mul(n17, n16)
	n19 := Simd_f32x4_add([2]uint64{p0, p0h}, n18)
	return n19[0], n19[1]
}

//go:noinline
func Simd_p_fx989(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_sub(n0, [2]uint64{p2, p2h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx990(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i8x16_sub(n0, [2]uint64{p2, p2h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx991(m *Module, s0 int64, s1 int64, s2 int64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load(m, n0, s2)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx992(m *Module, s0 int64, s1 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i8x16_sub([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_high_i8x16_s(n0)
	n2 := Simd_i16x8_extend_low_i8x16_s(n0)
	n3 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n4 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p4, p4h})
	n5 := Simd_i32x4_dot_i16x8_s(n3, n4)
	n6 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p3, p3h})
	n7 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p4, p4h})
	n8 := Simd_i32x4_dot_i16x8_s(n6, n7)
	n9 := Simd_i32x4_add(n5, n8)
	n10 := Simd_m64_v128_load(m, s0+18, s1)
	n11 := Simd_i16x8_extend_high_i8x16_s(n10)
	n12 := Simd_i32x4_dot_i16x8_s(n1, n11)
	n13 := Simd_i32x4_add(n9, n12)
	n14 := Simd_i16x8_extend_low_i8x16_s(n10)
	n15 := Simd_i32x4_dot_i16x8_s(n2, n14)
	n16 := Simd_i32x4_add(n13, n15)
	n17 := Simd_f32x4_convert_i32x4_s(n16)
	n18 := Simd_f32x4_splat(f0)
	n19 := Simd_f32x4_mul(n18, n17)
	n20 := Simd_f32x4_add([2]uint64{p2, p2h}, n19)
	return n0[0], n0[1], n10[0], n10[1], n20[0], n20[1]
}

//go:noinline
func Simd_p_fx993(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_or(n0, [2]uint64{p2, p2h})
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx994(m *Module, s0 int64, f0 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_v128_or([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_high_i8x16_s(n0)
	n2 := Simd_i16x8_extend_low_i8x16_s(n0)
	n3 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p3, p3h})
	n4 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p4, p4h})
	n5 := Simd_i32x4_dot_i16x8_s(n3, n4)
	n6 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p3, p3h})
	n7 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p4, p4h})
	n8 := Simd_i32x4_dot_i16x8_s(n6, n7)
	n9 := Simd_i32x4_add(n5, n8)
	n10 := Simd_m64_v128_load(m, s0+20, 0)
	n11 := Simd_i16x8_extend_high_i8x16_s(n10)
	n12 := Simd_i32x4_dot_i16x8_s(n1, n11)
	n13 := Simd_i32x4_add(n9, n12)
	n14 := Simd_i16x8_extend_low_i8x16_s(n10)
	n15 := Simd_i32x4_dot_i16x8_s(n2, n14)
	n16 := Simd_i32x4_add(n13, n15)
	n17 := Simd_f32x4_convert_i32x4_s(n16)
	n18 := Simd_f32x4_splat(f0)
	n19 := Simd_f32x4_mul(n18, n17)
	n20 := Simd_f32x4_add([2]uint64{p2, p2h}, n19)
	return n20[0], n20[1]
}

//go:noinline
func Simd_p_fx995(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0+2, 0, 0, 134)
	n1 := Simd_i16x8_extend_low_i8x16_s(n0)
	n2 := Simd_i16x8_extend_high_i8x16_s(n0)
	n3 := Simd_m64_v128_load_rng(m, s1+2, 0, 0, 134)
	n4 := Simd_i16x8_extend_low_i8x16_s(n3)
	n5 := Simd_i32x4_dot_i16x8_s(n1, n4)
	n6 := Simd_i16x8_extend_high_i8x16_s(n3)
	n7 := Simd_i32x4_dot_i16x8_s(n2, n6)
	n8 := Simd_i32x4_add(n5, n7)
	n9 := Simd_m64_v128_load_nc(m, s0+18, 0)
	n10 := Simd_i16x8_extend_high_i8x16_s(n9)
	n11 := Simd_i16x8_extend_low_i8x16_s(n9)
	n12 := Simd_m64_v128_load_nc(m, s1+18, 0)
	n13 := Simd_i16x8_extend_high_i8x16_s(n12)
	n14 := Simd_i32x4_dot_i16x8_s(n10, n13)
	n15 := Simd_i32x4_add(n8, n14)
	n16 := Simd_i16x8_extend_low_i8x16_s(n12)
	n17 := Simd_i32x4_dot_i16x8_s(n11, n16)
	n18 := Simd_i32x4_add(n15, n17)
	n19 := Simd_f32x4_convert_i32x4_s(n18)
	n20 := Simd_m64_scalar_i32_load16_u(m, s0)
	n21 := Simd_m64_scalar_i32_shl(n20, 2)
	n22 := Simd_m64_scalar_i32_add(n21, s2)
	n23 := Simd_m64_scalar_f32_load(m, n22)
	n24 := Simd_m64_scalar_i32_load16_u(m, s1)
	n25 := Simd_m64_scalar_i32_shl(n24, 2)
	n26 := Simd_m64_scalar_i32_add(n25, s2)
	n27 := Simd_m64_scalar_f32_load(m, n26)
	n28 := Simd_scalar_f32_mul(n23, n27)
	n29 := Simd_f32x4_splat(n28)
	n30 := Simd_f32x4_mul(n29, n19)
	n31 := Simd_f32x4_add([2]uint64{p0, p0h}, n30)
	n32 := Simd_m64_v128_load_nc(m, s0+36, 0)
	n33 := Simd_i16x8_extend_low_i8x16_s(n32)
	n34 := Simd_i16x8_extend_high_i8x16_s(n32)
	n35 := Simd_m64_v128_load_nc(m, s1+36, 0)
	n36 := Simd_i16x8_extend_low_i8x16_s(n35)
	n37 := Simd_i32x4_dot_i16x8_s(n33, n36)
	n38 := Simd_i16x8_extend_high_i8x16_s(n35)
	n39 := Simd_i32x4_dot_i16x8_s(n34, n38)
	n40 := Simd_i32x4_add(n37, n39)
	n41 := Simd_m64_v128_load_nc(m, s0+52, 0)
	n42 := Simd_i16x8_extend_high_i8x16_s(n41)
	n43 := Simd_i16x8_extend_low_i8x16_s(n41)
	n44 := Simd_m64_v128_load_nc(m, s1+52, 0)
	n45 := Simd_i16x8_extend_high_i8x16_s(n44)
	n46 := Simd_i32x4_dot_i16x8_s(n42, n45)
	n47 := Simd_i32x4_add(n40, n46)
	n48 := Simd_i16x8_extend_low_i8x16_s(n44)
	n49 := Simd_i32x4_dot_i16x8_s(n43, n48)
	n50 := Simd_i32x4_add(n47, n49)
	n51 := Simd_f32x4_convert_i32x4_s(n50)
	n52 := Simd_m64_scalar_i32_load16_u(m, s0+34)
	n53 := Simd_m64_scalar_i32_shl(n52, 2)
	n54 := Simd_m64_scalar_i32_add(n53, s2)
	n55 := Simd_m64_scalar_f32_load(m, n54)
	n56 := Simd_m64_scalar_i32_load16_u(m, s1+34)
	n57 := Simd_m64_scalar_i32_shl(n56, 2)
	n58 := Simd_m64_scalar_i32_add(n57, s2)
	n59 := Simd_m64_scalar_f32_load(m, n58)
	n60 := Simd_scalar_f32_mul(n55, n59)
	n61 := Simd_f32x4_splat(n60)
	n62 := Simd_f32x4_mul(n61, n51)
	n63 := Simd_f32x4_add(n31, n62)
	n64 := Simd_m64_v128_load_nc(m, s0+70, 0)
	n65 := Simd_i16x8_extend_low_i8x16_s(n64)
	n66 := Simd_i16x8_extend_high_i8x16_s(n64)
	n67 := Simd_m64_v128_load_nc(m, s1+70, 0)
	n68 := Simd_i16x8_extend_low_i8x16_s(n67)
	n69 := Simd_i32x4_dot_i16x8_s(n65, n68)
	n70 := Simd_i16x8_extend_high_i8x16_s(n67)
	n71 := Simd_i32x4_dot_i16x8_s(n66, n70)
	n72 := Simd_i32x4_add(n69, n71)
	n73 := Simd_m64_v128_load_nc(m, s0+86, 0)
	n74 := Simd_i16x8_extend_high_i8x16_s(n73)
	n75 := Simd_i16x8_extend_low_i8x16_s(n73)
	n76 := Simd_m64_v128_load_nc(m, s1+86, 0)
	n77 := Simd_i16x8_extend_high_i8x16_s(n76)
	n78 := Simd_i32x4_dot_i16x8_s(n74, n77)
	n79 := Simd_i32x4_add(n72, n78)
	n80 := Simd_i16x8_extend_low_i8x16_s(n76)
	n81 := Simd_i32x4_dot_i16x8_s(n75, n80)
	n82 := Simd_i32x4_add(n79, n81)
	n83 := Simd_f32x4_convert_i32x4_s(n82)
	n84 := Simd_m64_scalar_i32_load16_u(m, s0+68)
	n85 := Simd_m64_scalar_i32_shl(n84, 2)
	n86 := Simd_m64_scalar_i32_add(n85, s2)
	n87 := Simd_m64_scalar_f32_load(m, n86)
	n88 := Simd_m64_scalar_i32_load16_u(m, s1+68)
	n89 := Simd_m64_scalar_i32_shl(n88, 2)
	n90 := Simd_m64_scalar_i32_add(n89, s2)
	n91 := Simd_m64_scalar_f32_load(m, n90)
	n92 := Simd_scalar_f32_mul(n87, n91)
	n93 := Simd_f32x4_splat(n92)
	n94 := Simd_f32x4_mul(n93, n83)
	n95 := Simd_f32x4_add(n63, n94)
	n96 := Simd_m64_v128_load_nc(m, s0+104, 0)
	n97 := Simd_i16x8_extend_low_i8x16_s(n96)
	n98 := Simd_i16x8_extend_high_i8x16_s(n96)
	n99 := Simd_m64_v128_load_nc(m, s1+104, 0)
	n100 := Simd_i16x8_extend_low_i8x16_s(n99)
	n101 := Simd_i32x4_dot_i16x8_s(n97, n100)
	n102 := Simd_i16x8_extend_high_i8x16_s(n99)
	n103 := Simd_i32x4_dot_i16x8_s(n98, n102)
	n104 := Simd_i32x4_add(n101, n103)
	n105 := Simd_m64_v128_load_nc(m, s0+120, 0)
	n106 := Simd_i16x8_extend_high_i8x16_s(n105)
	n107 := Simd_i16x8_extend_low_i8x16_s(n105)
	n108 := Simd_m64_v128_load_nc(m, s1+120, 0)
	n109 := Simd_i16x8_extend_high_i8x16_s(n108)
	n110 := Simd_i32x4_dot_i16x8_s(n106, n109)
	n111 := Simd_i32x4_add(n104, n110)
	n112 := Simd_i16x8_extend_low_i8x16_s(n108)
	n113 := Simd_i32x4_dot_i16x8_s(n107, n112)
	n114 := Simd_i32x4_add(n111, n113)
	n115 := Simd_f32x4_convert_i32x4_s(n114)
	n116 := Simd_m64_scalar_i32_load16_u(m, s0+102)
	n117 := Simd_m64_scalar_i32_shl(n116, 2)
	n118 := Simd_m64_scalar_i32_add(n117, s2)
	n119 := Simd_m64_scalar_f32_load(m, n118)
	n120 := Simd_m64_scalar_i32_load16_u(m, s1+102)
	n121 := Simd_m64_scalar_i32_shl(n120, 2)
	n122 := Simd_m64_scalar_i32_add(n121, s2)
	n123 := Simd_m64_scalar_f32_load(m, n122)
	n124 := Simd_scalar_f32_mul(n119, n123)
	n125 := Simd_f32x4_splat(n124)
	n126 := Simd_f32x4_mul(n125, n115)
	n127 := Simd_f32x4_add(n95, n126)
	return n127[0], n127[1]
}

//go:noinline
func Simd_p_fx996(m *Module, s0 int64, s1 int64, s2 int64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load_rng(m, n0, s2, 0, 32)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx997(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, f0 float32, p0, p0h uint64, p1, p1h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p1, p1h})
	n2 := Simd_m64_scalar_i32_add(s0, s1)
	n3 := Simd_m64_v128_load_rng(m, n2, s2, 0, 32)
	n4 := Simd_i16x8_extend_low_i8x16_s(n3)
	n5 := Simd_i32x4_dot_i16x8_s(n0, n4)
	n6 := Simd_i16x8_extend_high_i8x16_s(n3)
	n7 := Simd_i32x4_dot_i16x8_s(n1, n6)
	n8 := Simd_i32x4_add(n5, n7)
	n9 := Simd_m64_v128_load_nc(m, s3, s2)
	n10 := Simd_i16x8_extend_high_i8x16_s(n9)
	n11 := Simd_i16x8_extend_low_i8x16_s(n9)
	n12 := Simd_m64_v128_load_nc(m, s4, s2)
	n13 := Simd_i16x8_extend_high_i8x16_s(n12)
	n14 := Simd_i32x4_dot_i16x8_s(n10, n13)
	n15 := Simd_i32x4_add(n8, n14)
	n16 := Simd_i16x8_extend_low_i8x16_s(n12)
	n17 := Simd_i32x4_dot_i16x8_s(n11, n16)
	n18 := Simd_i32x4_add(n15, n17)
	n19 := Simd_f32x4_convert_i32x4_s(n18)
	n20 := Simd_f32x4_splat(f0)
	n21 := Simd_f32x4_mul(n20, n19)
	n22 := Simd_f32x4_add([2]uint64{p0, p0h}, n21)
	return n3[0], n3[1], n9[0], n9[1], n12[0], n12[1], n22[0], n22[1]
}

//go:noinline
func Simd_p_fx998(m *Module, s0 int64, s1 int64, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i16x8_extend_high_i8x16_u(n0)
	n3 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_low_i8x16_u(n3)
	n5 := Simd_i16x8_extend_high_i8x16_u(n3)
	n6 := Simd_i32x4_splat(int32(s1))
	n7 := Simd_i32x4_splat(int32(s2))
	n8 := Simd_m64_v128_load_rng(m, s0, 0, 0, 32)
	n9 := Simd_i16x8_extend_low_i8x16_s(n8)
	n10 := Simd_i32x4_dot_i16x8_s(n9, n1)
	n11 := Simd_i16x8_extend_high_i8x16_s(n8)
	n12 := Simd_i32x4_dot_i16x8_s(n11, n2)
	n13 := Simd_i32x4_add(n10, n12)
	n14 := Simd_i32x4_mul(n13, n6)
	n15 := Simd_i32x4_add(n14, [2]uint64{p3, p3h})
	n16 := Simd_m64_v128_load_nc(m, s0+16, 0)
	n17 := Simd_i16x8_extend_low_i8x16_s(n16)
	n18 := Simd_i32x4_dot_i16x8_s(n17, n4)
	n19 := Simd_i16x8_extend_high_i8x16_s(n16)
	n20 := Simd_i32x4_dot_i16x8_s(n19, n5)
	n21 := Simd_i32x4_add(n18, n20)
	n22 := Simd_i32x4_mul(n21, n7)
	n23 := Simd_i32x4_add(n15, n22)
	return n23[0], n23[1]
}

//go:noinline
func Simd_p_fx999(m *Module, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_i8x16_shuffle([2]uint64{p0, p0h}, [2]uint64{p0, p0h}, [2]uint64{1084818905618843912, 506097522914230528})
	n1 := Simd_i32x4_add([2]uint64{p0, p0h}, n0)
	return n1[0], n1[1]
}

//go:noinline
func Simd_p_fx1000(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i32x4_dot_i16x8_s(n0, [2]uint64{p1, p1h})
	n2 := Simd_i16x8_extend_high_i8x16_u([2]uint64{p0, p0h})
	n3 := Simd_i32x4_dot_i16x8_s(n2, [2]uint64{p2, p2h})
	n4 := Simd_i32x4_add(n1, n3)
	n5 := Simd_i8x16_shuffle(n4, n4, [2]uint64{1084818905618843912, 506097522914230528})
	n6 := Simd_i32x4_add(n4, n5)
	return n6[0], n6[1]
}

//go:noinline
func Simd_p_fx1001(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p0, p0h})
	n1 := Simd_i16x8_extend_low_i8x16_s([2]uint64{p1, p1h})
	n2 := Simd_i32x4_dot_i16x8_s(n0, n1)
	n3 := Simd_i16x8_extend_high_i8x16_u([2]uint64{p0, p0h})
	n4 := Simd_i16x8_extend_high_i8x16_s([2]uint64{p1, p1h})
	n5 := Simd_i32x4_dot_i16x8_s(n3, n4)
	n6 := Simd_i32x4_add(n2, n5)
	n7 := Simd_i16x8_extend_low_i8x16_u([2]uint64{p2, p2h})
	n8 := Simd_i16x8_extend_high_i8x16_u([2]uint64{p2, p2h})
	n9 := Simd_m64_v128_load_nc(m, s0+48, 0)
	n10 := Simd_i16x8_extend_low_i8x16_s(n9)
	n11 := Simd_i32x4_dot_i16x8_s(n7, n10)
	n12 := Simd_i32x4_add(n6, n11)
	n13 := Simd_i16x8_extend_high_i8x16_s(n9)
	n14 := Simd_i32x4_dot_i16x8_s(n8, n13)
	n15 := Simd_i32x4_add(n12, n14)
	return n15[0], n15[1]
}

//go:noinline
func Simd_p_fx1002(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i16x8_extend_low_i8x16_u(n0)
	n2 := Simd_i16x8_extend_high_i8x16_u(n0)
	n3 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p1, p1h})
	n4 := Simd_i16x8_extend_low_i8x16_u(n3)
	n5 := Simd_i16x8_extend_high_i8x16_u(n3)
	n6 := Simd_m64_v128_load_nc(m, s0, 0)
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i32x4_dot_i16x8_s(n1, n7)
	n9 := Simd_i16x8_extend_high_i8x16_s(n6)
	n10 := Simd_i32x4_dot_i16x8_s(n2, n9)
	n11 := Simd_i32x4_add(n8, n10)
	n12 := Simd_m64_v128_load_nc(m, s0+16, 0)
	n13 := Simd_i16x8_extend_low_i8x16_s(n12)
	n14 := Simd_i32x4_dot_i16x8_s(n4, n13)
	n15 := Simd_i32x4_add(n11, n14)
	n16 := Simd_i16x8_extend_high_i8x16_s(n12)
	n17 := Simd_i32x4_dot_i16x8_s(n5, n16)
	n18 := Simd_i32x4_add(n15, n17)
	return n18[0], n18[1]
}

//go:noinline
func Simd_p_fx1003(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_v128_or(n0, n1)
	n3 := Simd_i16x8_extend_low_i8x16_u(n2)
	n4 := Simd_i16x8_extend_high_i8x16_u(n2)
	n5 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p3, p3h})
	n6 := Simd_m64_v128_load_rng(m, s0, 0, 0, 64)
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i32x4_dot_i16x8_s(n3, n7)
	n9 := Simd_i16x8_extend_high_i8x16_s(n6)
	n10 := Simd_i32x4_dot_i16x8_s(n4, n9)
	n11 := Simd_i32x4_add(n8, n10)
	n12 := Simd_m64_v128_load_nc(m, s1+16, 0)
	n13 := Simd_v128_and(n12, [2]uint64{p1, p1h})
	n14 := Simd_v128_or(n13, n5)
	n15 := Simd_i16x8_extend_high_i8x16_u(n14)
	n16 := Simd_i16x8_extend_low_i8x16_u(n14)
	n17 := Simd_m64_v128_load_nc(m, s0+16, 0)
	n18 := Simd_i16x8_extend_high_i8x16_s(n17)
	n19 := Simd_i32x4_dot_i16x8_s(n15, n18)
	n20 := Simd_i32x4_add(n11, n19)
	n21 := Simd_i16x8_extend_low_i8x16_s(n17)
	n22 := Simd_i32x4_dot_i16x8_s(n16, n21)
	n23 := Simd_i32x4_add(n20, n22)
	return n12[0], n12[1], n23[0], n23[1]
}

//go:noinline
func Simd_p_fx1004(m *Module, s0 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p1, p1h}, [2]uint64{p2, p2h})
	n1 := Simd_v128_or([2]uint64{p0, p0h}, n0)
	n2 := Simd_i16x8_extend_low_i8x16_u(n1)
	n3 := Simd_i16x8_extend_high_i8x16_u(n1)
	n4 := Simd_v128_and([2]uint64{p4, p4h}, [2]uint64{p2, p2h})
	n5 := Simd_v128_or([2]uint64{p3, p3h}, n4)
	n6 := Simd_i16x8_extend_high_i8x16_u(n5)
	n7 := Simd_i16x8_extend_low_i8x16_u(n5)
	n8 := Simd_m64_v128_load_nc(m, s0+32, 0)
	n9 := Simd_i16x8_extend_low_i8x16_s(n8)
	n10 := Simd_i32x4_dot_i16x8_s(n2, n9)
	n11 := Simd_i16x8_extend_high_i8x16_s(n8)
	n12 := Simd_i32x4_dot_i16x8_s(n3, n11)
	n13 := Simd_i32x4_add(n10, n12)
	n14 := Simd_m64_v128_load_nc(m, s0+48, 0)
	n15 := Simd_i16x8_extend_high_i8x16_s(n14)
	n16 := Simd_i32x4_dot_i16x8_s(n6, n15)
	n17 := Simd_i32x4_add(n13, n16)
	n18 := Simd_i16x8_extend_low_i8x16_s(n14)
	n19 := Simd_i32x4_dot_i16x8_s(n7, n18)
	n20 := Simd_i32x4_add(n17, n19)
	return n20[0], n20[1]
}

//go:noinline
func Simd_p_fx1005(m *Module, s0 int64, s1 int64, p0, p0h uint64) (uint64, uint64) {
	n0 := Simd_m64_scalar_i32_load16_u(m, s0)
	n1 := Simd_m64_scalar_i32_shl(n0, 2)
	n2 := Simd_m64_v128_load32_splat(m, n1, 8793760)
	n3 := Simd_m64_v128_load16x4_u(m, s1, 0)
	n4 := Simd_f16x4_cvt(n3)
	n5 := Simd_f32x4_mul(n2, n4)
	n6 := Simd_m64_v128_load(m, s1+120, 0)
	n7 := Simd_i16x8_extend_low_i8x16_s(n6)
	n8 := Simd_i16x8_extend_high_i8x16_s(n6)
	n9 := Simd_m64_v128_load32_splat(m, s0+30, 0)
	n10 := Simd_i16x8_extend_low_i8x16_s(n9)
	n11 := Simd_i32x4_dot_i16x8_s(n7, n10)
	n12 := Simd_i16x8_extend_high_i8x16_s(n9)
	n13 := Simd_i32x4_dot_i16x8_s(n8, n12)
	n14 := Simd_m64_v128_load(m, s1+104, 0)
	n15 := Simd_i16x8_extend_low_i8x16_s(n14)
	n16 := Simd_i16x8_extend_high_i8x16_s(n14)
	n17 := Simd_m64_v128_load32_splat(m, s0+26, 0)
	n18 := Simd_i16x8_extend_low_i8x16_s(n17)
	n19 := Simd_i32x4_dot_i16x8_s(n15, n18)
	n20 := Simd_i32x4_add(n11, n19)
	n21 := Simd_i16x8_extend_high_i8x16_s(n17)
	n22 := Simd_i32x4_dot_i16x8_s(n16, n21)
	n23 := Simd_i32x4_add(n13, n22)
	n24 := Simd_m64_v128_load(m, s1+88, 0)
	n25 := Simd_i16x8_extend_low_i8x16_s(n24)
	n26 := Simd_i16x8_extend_high_i8x16_s(n24)
	n27 := Simd_m64_v128_load32_splat(m, s0+22, 0)
	n28 := Simd_i16x8_extend_low_i8x16_s(n27)
	n29 := Simd_i32x4_dot_i16x8_s(n25, n28)
	n30 := Simd_i32x4_add(n20, n29)
	n31 := Simd_i16x8_extend_high_i8x16_s(n27)
	n32 := Simd_i32x4_dot_i16x8_s(n26, n31)
	n33 := Simd_i32x4_add(n23, n32)
	n34 := Simd_m64_v128_load(m, s1+72, 0)
	n35 := Simd_i16x8_extend_low_i8x16_s(n34)
	n36 := Simd_i16x8_extend_high_i8x16_s(n34)
	n37 := Simd_m64_v128_load32_splat(m, s0+18, 0)
	n38 := Simd_i16x8_extend_low_i8x16_s(n37)
	n39 := Simd_i32x4_dot_i16x8_s(n35, n38)
	n40 := Simd_i32x4_add(n30, n39)
	n41 := Simd_i16x8_extend_high_i8x16_s(n37)
	n42 := Simd_i32x4_dot_i16x8_s(n36, n41)
	n43 := Simd_i32x4_add(n33, n42)
	n44 := Simd_m64_v128_load(m, s1+56, 0)
	n45 := Simd_i16x8_extend_low_i8x16_s(n44)
	n46 := Simd_i16x8_extend_high_i8x16_s(n44)
	n47 := Simd_m64_v128_load32_splat(m, s0+14, 0)
	n48 := Simd_i16x8_extend_low_i8x16_s(n47)
	n49 := Simd_i32x4_dot_i16x8_s(n45, n48)
	n50 := Simd_i32x4_add(n40, n49)
	n51 := Simd_i16x8_extend_high_i8x16_s(n47)
	n52 := Simd_i32x4_dot_i16x8_s(n46, n51)
	n53 := Simd_i32x4_add(n43, n52)
	n54 := Simd_m64_v128_load(m, s1+40, 0)
	n55 := Simd_i16x8_extend_low_i8x16_s(n54)
	n56 := Simd_i16x8_extend_high_i8x16_s(n54)
	n57 := Simd_m64_v128_load32_splat(m, s0+10, 0)
	n58 := Simd_i16x8_extend_low_i8x16_s(n57)
	n59 := Simd_i32x4_dot_i16x8_s(n55, n58)
	n60 := Simd_i32x4_add(n50, n59)
	n61 := Simd_i16x8_extend_high_i8x16_s(n57)
	n62 := Simd_i32x4_dot_i16x8_s(n56, n61)
	n63 := Simd_i32x4_add(n53, n62)
	n64 := Simd_m64_v128_load(m, s1+24, 0)
	n65 := Simd_i16x8_extend_low_i8x16_s(n64)
	n66 := Simd_i16x8_extend_high_i8x16_s(n64)
	n67 := Simd_m64_v128_load32_splat(m, s0+6, 0)
	n68 := Simd_i16x8_extend_low_i8x16_s(n67)
	n69 := Simd_i32x4_dot_i16x8_s(n65, n68)
	n70 := Simd_i32x4_add(n60, n69)
	n71 := Simd_i16x8_extend_high_i8x16_s(n67)
	n72 := Simd_i32x4_dot_i16x8_s(n66, n71)
	n73 := Simd_i32x4_add(n63, n72)
	n74 := Simd_m64_v128_load(m, s1+8, 0)
	n75 := Simd_i16x8_extend_low_i8x16_s(n74)
	n76 := Simd_i16x8_extend_high_i8x16_s(n74)
	n77 := Simd_m64_v128_load32_splat(m, s0+2, 0)
	n78 := Simd_i16x8_extend_low_i8x16_s(n77)
	n79 := Simd_i32x4_dot_i16x8_s(n75, n78)
	n80 := Simd_i32x4_add(n70, n79)
	n81 := Simd_i16x8_extend_high_i8x16_s(n77)
	n82 := Simd_i32x4_dot_i16x8_s(n76, n81)
	n83 := Simd_i32x4_add(n73, n82)
	n84 := Simd_i8x16_shuffle(n80, n83, [2]uint64{795458214199165184, 1952900979608391952})
	n85 := Simd_i8x16_shuffle(n80, n83, [2]uint64{1084818905551471876, 2242261670960698644})
	n86 := Simd_i32x4_add(n84, n85)
	n87 := Simd_f32x4_convert_i32x4_s(n86)
	n88 := Simd_f32x4_mul(n5, n87)
	n89 := Simd_f32x4_add([2]uint64{p0, p0h}, n88)
	return n89[0], n89[1]
}

//go:noinline
func Simd_p_fx1006(m *Module, s0 int64) (uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 48)
	_ = Simd_m64_v128_store(m, s0, 8, n0)
	return n0[0], n0[1]
}

//go:noinline
func Simd_p_fx1007(m *Module, s0 int64, s1 int64, s2 int64) {
	n0 := Simd_m64_scalar_i32_add(s0, s1)
	n1 := Simd_m64_v128_load(m, n0, 0)
	n2 := Simd_m64_scalar_i32_add(s0, s2)
	_ = Simd_m64_v128_store(m, n2, 0, n1)
	return
}

//go:noinline
func Simd_p_fx1008(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int64, s5 int64, s6 int64, s7 int64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, s1)
	_ = Simd_m64_v128_store(m, s2, s1, n0)
	n2 := Simd_m64_v128_load(m, s3, s1)
	_ = Simd_m64_v128_store(m, s4, s1, n2)
	n4 := Simd_m64_v128_load(m, s5, s1)
	_ = Simd_m64_v128_store(m, s6, s1, n4)
	n6 := Simd_m64_v128_load(m, s7, s1)
	return n0[0], n0[1], n2[0], n2[1], n4[0], n4[1], n6[0], n6[1]
}

//go:noinline
func Simd_p_fx1009(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_v128_and([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i64x2_eq([2]uint64{p0, p0h}, n0)
	n2 := Simd_v128_and([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_i64x2_eq([2]uint64{p2, p2h}, n2)
	n4 := Simd_i8x16_shuffle(n1, n3, [2]uint64{795458214199165184, 1952900979608391952})
	return n4[0], n4[1]
}

//go:noinline
func Simd_p_fx1010(m *Module, s0 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load_rng(m, s0, 16, 0, 32)
	n1 := Simd_m64_v128_load_nc(m, s0, 0)
	return n0[0], n0[1], n1[0], n1[1]
}

//go:noinline
func Simd_p_fx1011(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 144)
	_ = Simd_m64_v128_store(m, s1, 29328, n0)
	n2 := Simd_m64_v128_load(m, s0, 128)
	_ = Simd_m64_v128_store(m, s1, 29312, n2)
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx1012(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_extend_low_i32x4_u([2]uint64{p1, p1h})
	n1 := Simd_i64x2_ne([2]uint64{p0, p0h}, n0)
	n2 := Simd_i64x2_ne([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_i8x16_shuffle(n1, n2, [2]uint64{795458214199165184, 1952900979608391952})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx1013(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_extend_low_i32x4_u([2]uint64{p1, p1h})
	n1 := Simd_i64x2_ne([2]uint64{p0, p0h}, n0)
	n2 := Simd_i64x2_ne([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_i8x16_shuffle(n1, n2, [2]uint64{p4, p4h})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx1014(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_ne([2]uint64{p0, p0h}, [2]uint64{p1, p1h})
	n1 := Simd_i64x2_ne([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n2 := Simd_i8x16_shuffle(n0, n1, [2]uint64{795458214199165184, 1952900979608391952})
	return n2[0], n2[1]
}

//go:noinline
func Simd_p_fx1015(m *Module, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64) {
	n0 := Simd_i64x2_extend_low_i32x4_u([2]uint64{p1, p1h})
	n1 := Simd_i64x2_eq([2]uint64{p0, p0h}, n0)
	n2 := Simd_i64x2_eq([2]uint64{p2, p2h}, [2]uint64{p3, p3h})
	n3 := Simd_i8x16_shuffle(n1, n2, [2]uint64{795458214199165184, 1952900979608391952})
	return n3[0], n3[1]
}

//go:noinline
func Simd_p_fx1016(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 29344, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 29328, n2)
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx1017(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 160, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 144, n2)
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx1018(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 29368, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 29352, n2)
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx1019(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 29384, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 29368, n2)
	return
}

//go:noinline
func Simd_p_fx1020(m *Module, s0 int64, s1 int64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 29360, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 29344, n2)
	return
}

//go:noinline
func Simd_p_fx1021(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 29432, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 29416, n2)
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx1022(m *Module, s0 int64, s1 int64) (uint64, uint64, uint64, uint64) {
	n0 := Simd_m64_v128_load(m, s0, 136)
	_ = Simd_m64_v128_store(m, s1, 168, n0)
	n2 := Simd_m64_v128_load(m, s0, 120)
	_ = Simd_m64_v128_store(m, s1, 152, n2)
	return n0[0], n0[1], n2[0], n2[1]
}

//go:noinline
func Simd_p_fx1023(m *Module, s0 int64, s1 int64, p0, p0h uint64, p1, p1h uint64) {
	n0 := Simd_i32x4_splat(int32(s0))
	n1 := Simd_i32x4_lt_u([2]uint64{p0, p0h}, n0)
	n2 := Simd_v128_and(n1, [2]uint64{p1, p1h})
	_ = Simd_m64_v128_store(m, s1, 0, n2)
	return
}

//go:noinline
func Simd_p_fxl0(m *Module, s0 int32, s1 int32, s2 int64, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64) (uint64, uint64, uint64, uint64, int64) {
	var out0 [2]uint64
	var out1 [2]uint64
	for {
		n0 := Simd_i32x4_splat(s0)
		n1 := Simd_i32x4_add(n0, [2]uint64{p0, p0h})
		n2 := Simd_i32x4_gt_s(n1, [2]uint64{p2, p2h})
		n3 := Simd_i32x4_splat(s1)
		n4 := Simd_i32x4_lt_s(n1, n3)
		n5 := Simd_v128_and(n4, n2)
		n6 := Simd_i32x4_sub([2]uint64{p1, p1h}, n5)
		n7 := Simd_i32x4_add([2]uint64{p0, p0h}, [2]uint64{p3, p3h})
		out0 = n6
		out1 = n7
		p0, p0h = n7[0], n7[1]
		p1, p1h = n6[0], n6[1]
		s2 = s2 - 4
		if s2 == 0 {
			break
		}
	}
	return out0[0], out0[1], out1[0], out1[1], s2
}

//go:noinline
func Simd_p_fxl1(m *Module, s0 int32, f0 float32, f1 float32, f2 float32, f3 float32, p0, p0h uint64, p1, p1h uint64, p2, p2h uint64, p3, p3h uint64, p4, p4h uint64, p5, p5h uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, int32) {
	var out0 [2]uint64
	var out1 [2]uint64
	var out2 [2]uint64
	var out3 [2]uint64
	for {
		n0 := Simd_f32x4_splat(f0)
		n1 := Simd_f32x4_mul([2]uint64{p0, p0h}, n0)
		n2 := Simd_f32x4_splat(f1)
		n3 := Simd_f32x4_mul([2]uint64{p1, p1h}, n2)
		n4 := Simd_f32x4_splat(f2)
		n5 := Simd_f32x4_mul([2]uint64{p2, p2h}, n4)
		n6 := Simd_f32x4_splat(f3)
		n7 := Simd_f32x4_mul([2]uint64{p3, p3h}, n6)
		n8 := Simd_f32x4_add([2]uint64{p5, p5h}, n3)
		n9 := Simd_f32x4_add(n8, n5)
		n10 := Simd_f32x4_add(n9, n7)
		n11 := Simd_f32x4_add(n1, n10)
		n12 := Simd_f32x4_div([2]uint64{p4, p4h}, n11)
		n13 := Simd_f32x4_mul(n1, n12)
		n14 := Simd_f32x4_mul(n7, n12)
		n15 := Simd_f32x4_mul(n5, n12)
		n16 := Simd_f32x4_mul(n3, n12)
		out0 = n13
		out1 = n14
		out2 = n15
		out3 = n16
		p0, p0h = n13[0], n13[1]
		p3, p3h = n14[0], n14[1]
		p2, p2h = n15[0], n15[1]
		p1, p1h = n16[0], n16[1]
		s0 = s0 - 1
		if s0 == 0 {
			break
		}
	}
	return out0[0], out0[1], out1[0], out1[1], out2[0], out2[1], out3[0], out3[1], s0
}

//go:noinline
func Simd_p_fxl2(m *Module, s0 int64, s1 int64, s2 int64, s3 int32, p0, p0h uint64) (uint64, uint64, int64, int32) {
	var out0 [2]uint64
	for {
		n0 := Simd_m64_scalar_i32_add(s0, s1)
		n1 := Simd_m64_scalar_i32_load16_u(m, n0)
		n2 := Simd_m64_scalar_i32_shl(n1, 2)
		n3 := Simd_m64_v128_load32_splat(m, n2, 8793760)
		n4 := Simd_m64_scalar_i32_add(s0, s2)
		n5 := Simd_m64_v128_load16x4_u(m, n4, 0)
		n6 := Simd_f16x4_cvt(n5)
		n7 := Simd_f32x4_mul(n3, n6)
		n8 := Simd_m64_scalar_i32_add(s0, s2)
		n9 := Simd_m64_scalar_i32_add(n8, 120)
		n10 := Simd_m64_v128_load(m, n9, 0)
		n11 := Simd_i16x8_extend_low_i8x16_s(n10)
		n12 := Simd_i16x8_extend_high_i8x16_s(n10)
		n13 := Simd_m64_scalar_i32_add(s0, s1)
		n14 := Simd_m64_scalar_i32_add(n13, 120)
		n15 := Simd_m64_v128_load32_splat(m, n14, 0)
		n16 := Simd_i16x8_extend_low_i8x16_s(n15)
		n17 := Simd_i32x4_dot_i16x8_s(n11, n16)
		n18 := Simd_i16x8_extend_high_i8x16_s(n15)
		n19 := Simd_i32x4_dot_i16x8_s(n12, n18)
		n20 := Simd_m64_scalar_i32_add(s0, s2)
		n21 := Simd_m64_scalar_i32_add(n20, 104)
		n22 := Simd_m64_v128_load(m, n21, 0)
		n23 := Simd_i16x8_extend_low_i8x16_s(n22)
		n24 := Simd_i16x8_extend_high_i8x16_s(n22)
		n25 := Simd_m64_scalar_i32_add(s0, s1)
		n26 := Simd_m64_scalar_i32_add(n25, 104)
		n27 := Simd_m64_v128_load32_splat(m, n26, 0)
		n28 := Simd_i16x8_extend_low_i8x16_s(n27)
		n29 := Simd_i32x4_dot_i16x8_s(n23, n28)
		n30 := Simd_i32x4_add(n17, n29)
		n31 := Simd_i16x8_extend_high_i8x16_s(n27)
		n32 := Simd_i32x4_dot_i16x8_s(n24, n31)
		n33 := Simd_i32x4_add(n19, n32)
		n34 := Simd_m64_scalar_i32_add(s0, s2)
		n35 := Simd_m64_scalar_i32_add(n34, 88)
		n36 := Simd_m64_v128_load(m, n35, 0)
		n37 := Simd_i16x8_extend_low_i8x16_s(n36)
		n38 := Simd_i16x8_extend_high_i8x16_s(n36)
		n39 := Simd_m64_scalar_i32_add(s0, s1)
		n40 := Simd_m64_scalar_i32_add(n39, 88)
		n41 := Simd_m64_v128_load32_splat(m, n40, 0)
		n42 := Simd_i16x8_extend_low_i8x16_s(n41)
		n43 := Simd_i32x4_dot_i16x8_s(n37, n42)
		n44 := Simd_i32x4_add(n30, n43)
		n45 := Simd_i16x8_extend_high_i8x16_s(n41)
		n46 := Simd_i32x4_dot_i16x8_s(n38, n45)
		n47 := Simd_i32x4_add(n33, n46)
		n48 := Simd_m64_scalar_i32_add(s0, s2)
		n49 := Simd_m64_scalar_i32_add(n48, 72)
		n50 := Simd_m64_v128_load(m, n49, 0)
		n51 := Simd_i16x8_extend_low_i8x16_s(n50)
		n52 := Simd_i16x8_extend_high_i8x16_s(n50)
		n53 := Simd_m64_scalar_i32_add(s0, s1)
		n54 := Simd_m64_scalar_i32_add(n53, 72)
		n55 := Simd_m64_v128_load32_splat(m, n54, 0)
		n56 := Simd_i16x8_extend_low_i8x16_s(n55)
		n57 := Simd_i32x4_dot_i16x8_s(n51, n56)
		n58 := Simd_i32x4_add(n44, n57)
		n59 := Simd_i16x8_extend_high_i8x16_s(n55)
		n60 := Simd_i32x4_dot_i16x8_s(n52, n59)
		n61 := Simd_i32x4_add(n47, n60)
		n62 := Simd_m64_scalar_i32_add(s0, s2)
		n63 := Simd_m64_scalar_i32_add(n62, 56)
		n64 := Simd_m64_v128_load(m, n63, 0)
		n65 := Simd_i16x8_extend_low_i8x16_s(n64)
		n66 := Simd_i16x8_extend_high_i8x16_s(n64)
		n67 := Simd_m64_scalar_i32_add(s0, s1)
		n68 := Simd_m64_scalar_i32_add(n67, 56)
		n69 := Simd_m64_v128_load32_splat(m, n68, 0)
		n70 := Simd_i16x8_extend_low_i8x16_s(n69)
		n71 := Simd_i32x4_dot_i16x8_s(n65, n70)
		n72 := Simd_i32x4_add(n58, n71)
		n73 := Simd_i16x8_extend_high_i8x16_s(n69)
		n74 := Simd_i32x4_dot_i16x8_s(n66, n73)
		n75 := Simd_i32x4_add(n61, n74)
		n76 := Simd_m64_scalar_i32_add(s0, s2)
		n77 := Simd_m64_scalar_i32_add(n76, 40)
		n78 := Simd_m64_v128_load(m, n77, 0)
		n79 := Simd_i16x8_extend_low_i8x16_s(n78)
		n80 := Simd_i16x8_extend_high_i8x16_s(n78)
		n81 := Simd_m64_scalar_i32_add(s0, s1)
		n82 := Simd_m64_scalar_i32_add(n81, 40)
		n83 := Simd_m64_v128_load32_splat(m, n82, 0)
		n84 := Simd_i16x8_extend_low_i8x16_s(n83)
		n85 := Simd_i32x4_dot_i16x8_s(n79, n84)
		n86 := Simd_i32x4_add(n72, n85)
		n87 := Simd_i16x8_extend_high_i8x16_s(n83)
		n88 := Simd_i32x4_dot_i16x8_s(n80, n87)
		n89 := Simd_i32x4_add(n75, n88)
		n90 := Simd_m64_scalar_i32_add(s0, s2)
		n91 := Simd_m64_scalar_i32_add(n90, 24)
		n92 := Simd_m64_v128_load(m, n91, 0)
		n93 := Simd_i16x8_extend_low_i8x16_s(n92)
		n94 := Simd_i16x8_extend_high_i8x16_s(n92)
		n95 := Simd_m64_scalar_i32_add(s0, s1)
		n96 := Simd_m64_scalar_i32_add(n95, 24)
		n97 := Simd_m64_v128_load32_splat(m, n96, 0)
		n98 := Simd_i16x8_extend_low_i8x16_s(n97)
		n99 := Simd_i32x4_dot_i16x8_s(n93, n98)
		n100 := Simd_i32x4_add(n86, n99)
		n101 := Simd_i16x8_extend_high_i8x16_s(n97)
		n102 := Simd_i32x4_dot_i16x8_s(n94, n101)
		n103 := Simd_i32x4_add(n89, n102)
		n104 := Simd_m64_scalar_i32_add(s0, s2)
		n105 := Simd_m64_scalar_i32_add(n104, 8)
		n106 := Simd_m64_v128_load(m, n105, 0)
		n107 := Simd_i16x8_extend_low_i8x16_s(n106)
		n108 := Simd_i16x8_extend_high_i8x16_s(n106)
		n109 := Simd_m64_scalar_i32_add(s0, s1)
		n110 := Simd_m64_scalar_i32_add(n109, 8)
		n111 := Simd_m64_v128_load32_splat(m, n110, 0)
		n112 := Simd_i16x8_extend_low_i8x16_s(n111)
		n113 := Simd_i32x4_dot_i16x8_s(n107, n112)
		n114 := Simd_i32x4_add(n100, n113)
		n115 := Simd_i16x8_extend_high_i8x16_s(n111)
		n116 := Simd_i32x4_dot_i16x8_s(n108, n115)
		n117 := Simd_i32x4_add(n103, n116)
		n118 := Simd_i8x16_shuffle(n114, n117, [2]uint64{795458214199165184, 1952900979608391952})
		n119 := Simd_i8x16_shuffle(n114, n117, [2]uint64{1084818905551471876, 2242261670960698644})
		n120 := Simd_i32x4_add(n118, n119)
		n121 := Simd_f32x4_convert_i32x4_s(n120)
		n122 := Simd_f32x4_mul(n7, n121)
		n123 := Simd_f32x4_add([2]uint64{p0, p0h}, n122)
		out0 = n123
		p0, p0h = n123[0], n123[1]
		s0 = s0 + 136
		s3 = s3 - 1
		if s3 <= 1 {
			break
		}
	}
	return out0[0], out0[1], s0, s3
}

//go:noinline
func Simd_p_fxl3(m *Module, s0 int64, s1 int64, s2 int64, s3 int32, p0, p0h uint64) (uint64, uint64, int64, int32) {
	var out0 [2]uint64
	for {
		n0 := Simd_m64_scalar_i32_add(s0, s1)
		n1 := Simd_m64_scalar_i32_add(n0, 2)
		n2 := Simd_m64_scalar_i32_load16_u(m, n1)
		n3 := Simd_m64_scalar_i32_shl(n2, 2)
		n4 := Simd_m64_v128_load32_splat(m, n3, 8793760)
		n5 := Simd_m64_scalar_i32_add(s0, s2)
		n6 := Simd_m64_v128_load16x4_u(m, n5, 0)
		n7 := Simd_f16x4_cvt(n6)
		n8 := Simd_f32x4_mul(n4, n7)
		n9 := Simd_m64_scalar_i32_add(s0, s2)
		n10 := Simd_m64_scalar_i32_add(n9, 120)
		n11 := Simd_m64_v128_load(m, n10, 0)
		n12 := Simd_i16x8_extend_low_i8x16_s(n11)
		n13 := Simd_i16x8_extend_high_i8x16_s(n11)
		n14 := Simd_m64_scalar_i32_add(s0, s1)
		n15 := Simd_m64_scalar_i32_add(n14, 124)
		n16 := Simd_m64_v128_load32_splat(m, n15, 0)
		n17 := Simd_i16x8_extend_low_i8x16_s(n16)
		n18 := Simd_i32x4_dot_i16x8_s(n12, n17)
		n19 := Simd_i16x8_extend_high_i8x16_s(n16)
		n20 := Simd_i32x4_dot_i16x8_s(n13, n19)
		n21 := Simd_m64_scalar_i32_add(s0, s2)
		n22 := Simd_m64_scalar_i32_add(n21, 104)
		n23 := Simd_m64_v128_load(m, n22, 0)
		n24 := Simd_i16x8_extend_low_i8x16_s(n23)
		n25 := Simd_i16x8_extend_high_i8x16_s(n23)
		n26 := Simd_m64_scalar_i32_add(s0, s1)
		n27 := Simd_m64_scalar_i32_add(n26, 108)
		n28 := Simd_m64_v128_load32_splat(m, n27, 0)
		n29 := Simd_i16x8_extend_low_i8x16_s(n28)
		n30 := Simd_i32x4_dot_i16x8_s(n24, n29)
		n31 := Simd_i32x4_add(n18, n30)
		n32 := Simd_i16x8_extend_high_i8x16_s(n28)
		n33 := Simd_i32x4_dot_i16x8_s(n25, n32)
		n34 := Simd_i32x4_add(n20, n33)
		n35 := Simd_m64_scalar_i32_add(s0, s2)
		n36 := Simd_m64_scalar_i32_add(n35, 88)
		n37 := Simd_m64_v128_load(m, n36, 0)
		n38 := Simd_i16x8_extend_low_i8x16_s(n37)
		n39 := Simd_i16x8_extend_high_i8x16_s(n37)
		n40 := Simd_m64_scalar_i32_add(s0, s1)
		n41 := Simd_m64_scalar_i32_add(n40, 92)
		n42 := Simd_m64_v128_load32_splat(m, n41, 0)
		n43 := Simd_i16x8_extend_low_i8x16_s(n42)
		n44 := Simd_i32x4_dot_i16x8_s(n38, n43)
		n45 := Simd_i32x4_add(n31, n44)
		n46 := Simd_i16x8_extend_high_i8x16_s(n42)
		n47 := Simd_i32x4_dot_i16x8_s(n39, n46)
		n48 := Simd_i32x4_add(n34, n47)
		n49 := Simd_m64_scalar_i32_add(s0, s2)
		n50 := Simd_m64_scalar_i32_add(n49, 72)
		n51 := Simd_m64_v128_load(m, n50, 0)
		n52 := Simd_i16x8_extend_low_i8x16_s(n51)
		n53 := Simd_i16x8_extend_high_i8x16_s(n51)
		n54 := Simd_m64_scalar_i32_add(s0, s1)
		n55 := Simd_m64_scalar_i32_add(n54, 76)
		n56 := Simd_m64_v128_load32_splat(m, n55, 0)
		n57 := Simd_i16x8_extend_low_i8x16_s(n56)
		n58 := Simd_i32x4_dot_i16x8_s(n52, n57)
		n59 := Simd_i32x4_add(n45, n58)
		n60 := Simd_i16x8_extend_high_i8x16_s(n56)
		n61 := Simd_i32x4_dot_i16x8_s(n53, n60)
		n62 := Simd_i32x4_add(n48, n61)
		n63 := Simd_m64_scalar_i32_add(s0, s2)
		n64 := Simd_m64_scalar_i32_add(n63, 56)
		n65 := Simd_m64_v128_load(m, n64, 0)
		n66 := Simd_i16x8_extend_low_i8x16_s(n65)
		n67 := Simd_i16x8_extend_high_i8x16_s(n65)
		n68 := Simd_m64_scalar_i32_add(s0, s1)
		n69 := Simd_m64_scalar_i32_add(n68, 60)
		n70 := Simd_m64_v128_load32_splat(m, n69, 0)
		n71 := Simd_i16x8_extend_low_i8x16_s(n70)
		n72 := Simd_i32x4_dot_i16x8_s(n66, n71)
		n73 := Simd_i32x4_add(n59, n72)
		n74 := Simd_i16x8_extend_high_i8x16_s(n70)
		n75 := Simd_i32x4_dot_i16x8_s(n67, n74)
		n76 := Simd_i32x4_add(n62, n75)
		n77 := Simd_m64_scalar_i32_add(s0, s2)
		n78 := Simd_m64_scalar_i32_add(n77, 40)
		n79 := Simd_m64_v128_load(m, n78, 0)
		n80 := Simd_i16x8_extend_low_i8x16_s(n79)
		n81 := Simd_i16x8_extend_high_i8x16_s(n79)
		n82 := Simd_m64_scalar_i32_add(s0, s1)
		n83 := Simd_m64_scalar_i32_add(n82, 44)
		n84 := Simd_m64_v128_load32_splat(m, n83, 0)
		n85 := Simd_i16x8_extend_low_i8x16_s(n84)
		n86 := Simd_i32x4_dot_i16x8_s(n80, n85)
		n87 := Simd_i32x4_add(n73, n86)
		n88 := Simd_i16x8_extend_high_i8x16_s(n84)
		n89 := Simd_i32x4_dot_i16x8_s(n81, n88)
		n90 := Simd_i32x4_add(n76, n89)
		n91 := Simd_m64_scalar_i32_add(s0, s2)
		n92 := Simd_m64_scalar_i32_add(n91, 24)
		n93 := Simd_m64_v128_load(m, n92, 0)
		n94 := Simd_i16x8_extend_low_i8x16_s(n93)
		n95 := Simd_i16x8_extend_high_i8x16_s(n93)
		n96 := Simd_m64_scalar_i32_add(s0, s1)
		n97 := Simd_m64_scalar_i32_add(n96, 28)
		n98 := Simd_m64_v128_load32_splat(m, n97, 0)
		n99 := Simd_i16x8_extend_low_i8x16_s(n98)
		n100 := Simd_i32x4_dot_i16x8_s(n94, n99)
		n101 := Simd_i32x4_add(n87, n100)
		n102 := Simd_i16x8_extend_high_i8x16_s(n98)
		n103 := Simd_i32x4_dot_i16x8_s(n95, n102)
		n104 := Simd_i32x4_add(n90, n103)
		n105 := Simd_m64_scalar_i32_add(s0, s2)
		n106 := Simd_m64_scalar_i32_add(n105, 8)
		n107 := Simd_m64_v128_load(m, n106, 0)
		n108 := Simd_i16x8_extend_low_i8x16_s(n107)
		n109 := Simd_i16x8_extend_high_i8x16_s(n107)
		n110 := Simd_m64_scalar_i32_add(s0, s1)
		n111 := Simd_m64_scalar_i32_add(n110, 12)
		n112 := Simd_m64_v128_load32_splat(m, n111, 0)
		n113 := Simd_i16x8_extend_low_i8x16_s(n112)
		n114 := Simd_i32x4_dot_i16x8_s(n108, n113)
		n115 := Simd_i32x4_add(n101, n114)
		n116 := Simd_i16x8_extend_high_i8x16_s(n112)
		n117 := Simd_i32x4_dot_i16x8_s(n109, n116)
		n118 := Simd_i32x4_add(n104, n117)
		n119 := Simd_i8x16_shuffle(n115, n118, [2]uint64{795458214199165184, 1952900979608391952})
		n120 := Simd_i8x16_shuffle(n115, n118, [2]uint64{1084818905551471876, 2242261670960698644})
		n121 := Simd_i32x4_add(n119, n120)
		n122 := Simd_f32x4_convert_i32x4_s(n121)
		n123 := Simd_f32x4_mul(n8, n122)
		n124 := Simd_f32x4_add([2]uint64{p0, p0h}, n123)
		out0 = n124
		p0, p0h = n124[0], n124[1]
		s0 = s0 + 136
		s3 = s3 - 1
		if s3 <= 1 {
			break
		}
	}
	return out0[0], out0[1], s0, s3
}

//go:noinline
func Simd_p_fxl4(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, s4 int32, p0, p0h uint64) (uint64, uint64, int64, int64, int32) {
	var out0 [2]uint64
	for {
		n0 := Simd_m64_scalar_i32_add(s0, s1)
		n1 := Simd_m64_scalar_i32_add(n0, 4)
		n2 := Simd_m64_scalar_i32_load16_u(m, n1)
		n3 := Simd_m64_scalar_i32_shl(n2, 2)
		n4 := Simd_m64_v128_load32_splat(m, n3, 8793760)
		n5 := Simd_m64_scalar_i32_add(s0, s2)
		n6 := Simd_m64_v128_load16x4_u(m, n5, 0)
		n7 := Simd_f16x4_cvt(n6)
		n8 := Simd_f32x4_mul(n4, n7)
		n9 := Simd_m64_scalar_i32_add(s0, s2)
		n10 := Simd_m64_scalar_i32_add(n9, 120)
		n11 := Simd_m64_v128_load(m, n10, 0)
		n12 := Simd_i16x8_extend_low_i8x16_s(n11)
		n13 := Simd_i16x8_extend_high_i8x16_s(n11)
		n14 := Simd_m64_scalar_i32_add(s0, s1)
		n15 := Simd_m64_scalar_i32_add(n14, 128)
		n16 := Simd_m64_v128_load32_splat(m, n15, 0)
		n17 := Simd_i16x8_extend_low_i8x16_s(n16)
		n18 := Simd_i32x4_dot_i16x8_s(n12, n17)
		n19 := Simd_i16x8_extend_high_i8x16_s(n16)
		n20 := Simd_i32x4_dot_i16x8_s(n13, n19)
		n21 := Simd_m64_scalar_i32_add(s0, s2)
		n22 := Simd_m64_scalar_i32_add(n21, 104)
		n23 := Simd_m64_v128_load(m, n22, 0)
		n24 := Simd_i16x8_extend_low_i8x16_s(n23)
		n25 := Simd_i16x8_extend_high_i8x16_s(n23)
		n26 := Simd_m64_scalar_i32_add(s0, s1)
		n27 := Simd_m64_scalar_i32_add(n26, 112)
		n28 := Simd_m64_v128_load32_splat(m, n27, 0)
		n29 := Simd_i16x8_extend_low_i8x16_s(n28)
		n30 := Simd_i32x4_dot_i16x8_s(n24, n29)
		n31 := Simd_i32x4_add(n18, n30)
		n32 := Simd_i16x8_extend_high_i8x16_s(n28)
		n33 := Simd_i32x4_dot_i16x8_s(n25, n32)
		n34 := Simd_i32x4_add(n20, n33)
		n35 := Simd_m64_scalar_i32_add(s0, s2)
		n36 := Simd_m64_scalar_i32_add(n35, 88)
		n37 := Simd_m64_v128_load(m, n36, 0)
		n38 := Simd_i16x8_extend_low_i8x16_s(n37)
		n39 := Simd_i16x8_extend_high_i8x16_s(n37)
		n40 := Simd_m64_scalar_i32_add(s0, s1)
		n41 := Simd_m64_scalar_i32_add(n40, 96)
		n42 := Simd_m64_v128_load32_splat(m, n41, 0)
		n43 := Simd_i16x8_extend_low_i8x16_s(n42)
		n44 := Simd_i32x4_dot_i16x8_s(n38, n43)
		n45 := Simd_i32x4_add(n31, n44)
		n46 := Simd_i16x8_extend_high_i8x16_s(n42)
		n47 := Simd_i32x4_dot_i16x8_s(n39, n46)
		n48 := Simd_i32x4_add(n34, n47)
		n49 := Simd_m64_scalar_i32_add(s0, s2)
		n50 := Simd_m64_scalar_i32_add(n49, 72)
		n51 := Simd_m64_v128_load(m, n50, 0)
		n52 := Simd_i16x8_extend_low_i8x16_s(n51)
		n53 := Simd_i16x8_extend_high_i8x16_s(n51)
		n54 := Simd_m64_scalar_i32_add(s0, s1)
		n55 := Simd_m64_scalar_i32_add(n54, 80)
		n56 := Simd_m64_v128_load32_splat(m, n55, 0)
		n57 := Simd_i16x8_extend_low_i8x16_s(n56)
		n58 := Simd_i32x4_dot_i16x8_s(n52, n57)
		n59 := Simd_i32x4_add(n45, n58)
		n60 := Simd_i16x8_extend_high_i8x16_s(n56)
		n61 := Simd_i32x4_dot_i16x8_s(n53, n60)
		n62 := Simd_i32x4_add(n48, n61)
		n63 := Simd_m64_scalar_i32_add(s0, s2)
		n64 := Simd_m64_scalar_i32_add(n63, 56)
		n65 := Simd_m64_v128_load(m, n64, 0)
		n66 := Simd_i16x8_extend_low_i8x16_s(n65)
		n67 := Simd_i16x8_extend_high_i8x16_s(n65)
		n68 := Simd_m64_v128_load32_splat(m, s3, 0)
		n69 := Simd_i16x8_extend_low_i8x16_s(n68)
		n70 := Simd_i32x4_dot_i16x8_s(n66, n69)
		n71 := Simd_i32x4_add(n59, n70)
		n72 := Simd_i16x8_extend_high_i8x16_s(n68)
		n73 := Simd_i32x4_dot_i16x8_s(n67, n72)
		n74 := Simd_i32x4_add(n62, n73)
		n75 := Simd_m64_scalar_i32_add(s0, s2)
		n76 := Simd_m64_scalar_i32_add(n75, 40)
		n77 := Simd_m64_v128_load(m, n76, 0)
		n78 := Simd_i16x8_extend_low_i8x16_s(n77)
		n79 := Simd_i16x8_extend_high_i8x16_s(n77)
		n80 := Simd_m64_scalar_i32_add(s0, s1)
		n81 := Simd_m64_scalar_i32_add(n80, 48)
		n82 := Simd_m64_v128_load32_splat(m, n81, 0)
		n83 := Simd_i16x8_extend_low_i8x16_s(n82)
		n84 := Simd_i32x4_dot_i16x8_s(n78, n83)
		n85 := Simd_i32x4_add(n71, n84)
		n86 := Simd_i16x8_extend_high_i8x16_s(n82)
		n87 := Simd_i32x4_dot_i16x8_s(n79, n86)
		n88 := Simd_i32x4_add(n74, n87)
		n89 := Simd_m64_scalar_i32_add(s0, s2)
		n90 := Simd_m64_scalar_i32_add(n89, 24)
		n91 := Simd_m64_v128_load(m, n90, 0)
		n92 := Simd_i16x8_extend_low_i8x16_s(n91)
		n93 := Simd_i16x8_extend_high_i8x16_s(n91)
		n94 := Simd_m64_scalar_i32_add(s0, s1)
		n95 := Simd_m64_scalar_i32_add(n94, 32)
		n96 := Simd_m64_v128_load32_splat(m, n95, 0)
		n97 := Simd_i16x8_extend_low_i8x16_s(n96)
		n98 := Simd_i32x4_dot_i16x8_s(n92, n97)
		n99 := Simd_i32x4_add(n85, n98)
		n100 := Simd_i16x8_extend_high_i8x16_s(n96)
		n101 := Simd_i32x4_dot_i16x8_s(n93, n100)
		n102 := Simd_i32x4_add(n88, n101)
		n103 := Simd_m64_scalar_i32_add(s0, s2)
		n104 := Simd_m64_scalar_i32_add(n103, 8)
		n105 := Simd_m64_v128_load(m, n104, 0)
		n106 := Simd_i16x8_extend_low_i8x16_s(n105)
		n107 := Simd_i16x8_extend_high_i8x16_s(n105)
		n108 := Simd_m64_scalar_i32_add(s0, s1)
		n109 := Simd_m64_scalar_i32_add(n108, 16)
		n110 := Simd_m64_v128_load32_splat(m, n109, 0)
		n111 := Simd_i16x8_extend_low_i8x16_s(n110)
		n112 := Simd_i32x4_dot_i16x8_s(n106, n111)
		n113 := Simd_i32x4_add(n99, n112)
		n114 := Simd_i16x8_extend_high_i8x16_s(n110)
		n115 := Simd_i32x4_dot_i16x8_s(n107, n114)
		n116 := Simd_i32x4_add(n102, n115)
		n117 := Simd_i8x16_shuffle(n113, n116, [2]uint64{795458214199165184, 1952900979608391952})
		n118 := Simd_i8x16_shuffle(n113, n116, [2]uint64{1084818905551471876, 2242261670960698644})
		n119 := Simd_i32x4_add(n117, n118)
		n120 := Simd_f32x4_convert_i32x4_s(n119)
		n121 := Simd_f32x4_mul(n8, n120)
		n122 := Simd_f32x4_add([2]uint64{p0, p0h}, n121)
		out0 = n122
		p0, p0h = n122[0], n122[1]
		s0 = s0 + 136
		s3 = s3 + 136
		s4 = s4 - 1
		if s4 <= 1 {
			break
		}
	}
	return out0[0], out0[1], s0, s3, s4
}

//go:noinline
func Simd_p_fxl5(m *Module, s0 int64, s1 int64, s2 int64, s3 int32, p0, p0h uint64) (uint64, uint64, int64, int32) {
	var out0 [2]uint64
	for {
		n0 := Simd_m64_scalar_i32_add(s0, s1)
		n1 := Simd_m64_scalar_i32_add(n0, 6)
		n2 := Simd_m64_scalar_i32_load16_u(m, n1)
		n3 := Simd_m64_scalar_i32_shl(n2, 2)
		n4 := Simd_m64_v128_load32_splat(m, n3, 8793760)
		n5 := Simd_m64_scalar_i32_add(s0, s2)
		n6 := Simd_m64_v128_load16x4_u(m, n5, 0)
		n7 := Simd_f16x4_cvt(n6)
		n8 := Simd_f32x4_mul(n4, n7)
		n9 := Simd_m64_scalar_i32_add(s0, s2)
		n10 := Simd_m64_scalar_i32_add(n9, 120)
		n11 := Simd_m64_v128_load(m, n10, 0)
		n12 := Simd_i16x8_extend_low_i8x16_s(n11)
		n13 := Simd_i16x8_extend_high_i8x16_s(n11)
		n14 := Simd_m64_scalar_i32_add(s0, s1)
		n15 := Simd_m64_scalar_i32_add(n14, 132)
		n16 := Simd_m64_v128_load32_splat(m, n15, 0)
		n17 := Simd_i16x8_extend_low_i8x16_s(n16)
		n18 := Simd_i32x4_dot_i16x8_s(n12, n17)
		n19 := Simd_i16x8_extend_high_i8x16_s(n16)
		n20 := Simd_i32x4_dot_i16x8_s(n13, n19)
		n21 := Simd_m64_scalar_i32_add(s0, s2)
		n22 := Simd_m64_scalar_i32_add(n21, 104)
		n23 := Simd_m64_v128_load(m, n22, 0)
		n24 := Simd_i16x8_extend_low_i8x16_s(n23)
		n25 := Simd_i16x8_extend_high_i8x16_s(n23)
		n26 := Simd_m64_scalar_i32_add(s0, s1)
		n27 := Simd_m64_scalar_i32_add(n26, 116)
		n28 := Simd_m64_v128_load32_splat(m, n27, 0)
		n29 := Simd_i16x8_extend_low_i8x16_s(n28)
		n30 := Simd_i32x4_dot_i16x8_s(n24, n29)
		n31 := Simd_i32x4_add(n18, n30)
		n32 := Simd_i16x8_extend_high_i8x16_s(n28)
		n33 := Simd_i32x4_dot_i16x8_s(n25, n32)
		n34 := Simd_i32x4_add(n20, n33)
		n35 := Simd_m64_scalar_i32_add(s0, s2)
		n36 := Simd_m64_scalar_i32_add(n35, 88)
		n37 := Simd_m64_v128_load(m, n36, 0)
		n38 := Simd_i16x8_extend_low_i8x16_s(n37)
		n39 := Simd_i16x8_extend_high_i8x16_s(n37)
		n40 := Simd_m64_scalar_i32_add(s0, s1)
		n41 := Simd_m64_scalar_i32_add(n40, 100)
		n42 := Simd_m64_v128_load32_splat(m, n41, 0)
		n43 := Simd_i16x8_extend_low_i8x16_s(n42)
		n44 := Simd_i32x4_dot_i16x8_s(n38, n43)
		n45 := Simd_i32x4_add(n31, n44)
		n46 := Simd_i16x8_extend_high_i8x16_s(n42)
		n47 := Simd_i32x4_dot_i16x8_s(n39, n46)
		n48 := Simd_i32x4_add(n34, n47)
		n49 := Simd_m64_scalar_i32_add(s0, s2)
		n50 := Simd_m64_scalar_i32_add(n49, 72)
		n51 := Simd_m64_v128_load(m, n50, 0)
		n52 := Simd_i16x8_extend_low_i8x16_s(n51)
		n53 := Simd_i16x8_extend_high_i8x16_s(n51)
		n54 := Simd_m64_scalar_i32_add(s0, s1)
		n55 := Simd_m64_scalar_i32_add(n54, 84)
		n56 := Simd_m64_v128_load32_splat(m, n55, 0)
		n57 := Simd_i16x8_extend_low_i8x16_s(n56)
		n58 := Simd_i32x4_dot_i16x8_s(n52, n57)
		n59 := Simd_i32x4_add(n45, n58)
		n60 := Simd_i16x8_extend_high_i8x16_s(n56)
		n61 := Simd_i32x4_dot_i16x8_s(n53, n60)
		n62 := Simd_i32x4_add(n48, n61)
		n63 := Simd_m64_scalar_i32_add(s0, s2)
		n64 := Simd_m64_scalar_i32_add(n63, 56)
		n65 := Simd_m64_v128_load(m, n64, 0)
		n66 := Simd_i16x8_extend_low_i8x16_s(n65)
		n67 := Simd_i16x8_extend_high_i8x16_s(n65)
		n68 := Simd_m64_scalar_i32_add(s0, s1)
		n69 := Simd_m64_scalar_i32_add(n68, 68)
		n70 := Simd_m64_v128_load32_splat(m, n69, 0)
		n71 := Simd_i16x8_extend_low_i8x16_s(n70)
		n72 := Simd_i32x4_dot_i16x8_s(n66, n71)
		n73 := Simd_i32x4_add(n59, n72)
		n74 := Simd_i16x8_extend_high_i8x16_s(n70)
		n75 := Simd_i32x4_dot_i16x8_s(n67, n74)
		n76 := Simd_i32x4_add(n62, n75)
		n77 := Simd_m64_scalar_i32_add(s0, s2)
		n78 := Simd_m64_scalar_i32_add(n77, 40)
		n79 := Simd_m64_v128_load(m, n78, 0)
		n80 := Simd_i16x8_extend_low_i8x16_s(n79)
		n81 := Simd_i16x8_extend_high_i8x16_s(n79)
		n82 := Simd_m64_scalar_i32_add(s0, s1)
		n83 := Simd_m64_scalar_i32_add(n82, 52)
		n84 := Simd_m64_v128_load32_splat(m, n83, 0)
		n85 := Simd_i16x8_extend_low_i8x16_s(n84)
		n86 := Simd_i32x4_dot_i16x8_s(n80, n85)
		n87 := Simd_i32x4_add(n73, n86)
		n88 := Simd_i16x8_extend_high_i8x16_s(n84)
		n89 := Simd_i32x4_dot_i16x8_s(n81, n88)
		n90 := Simd_i32x4_add(n76, n89)
		n91 := Simd_m64_scalar_i32_add(s0, s2)
		n92 := Simd_m64_scalar_i32_add(n91, 24)
		n93 := Simd_m64_v128_load(m, n92, 0)
		n94 := Simd_i16x8_extend_low_i8x16_s(n93)
		n95 := Simd_i16x8_extend_high_i8x16_s(n93)
		n96 := Simd_m64_scalar_i32_add(s0, s1)
		n97 := Simd_m64_scalar_i32_add(n96, 36)
		n98 := Simd_m64_v128_load32_splat(m, n97, 0)
		n99 := Simd_i16x8_extend_low_i8x16_s(n98)
		n100 := Simd_i32x4_dot_i16x8_s(n94, n99)
		n101 := Simd_i32x4_add(n87, n100)
		n102 := Simd_i16x8_extend_high_i8x16_s(n98)
		n103 := Simd_i32x4_dot_i16x8_s(n95, n102)
		n104 := Simd_i32x4_add(n90, n103)
		n105 := Simd_m64_scalar_i32_add(s0, s2)
		n106 := Simd_m64_scalar_i32_add(n105, 8)
		n107 := Simd_m64_v128_load(m, n106, 0)
		n108 := Simd_i16x8_extend_low_i8x16_s(n107)
		n109 := Simd_i16x8_extend_high_i8x16_s(n107)
		n110 := Simd_m64_scalar_i32_add(s0, s1)
		n111 := Simd_m64_scalar_i32_add(n110, 20)
		n112 := Simd_m64_v128_load32_splat(m, n111, 0)
		n113 := Simd_i16x8_extend_low_i8x16_s(n112)
		n114 := Simd_i32x4_dot_i16x8_s(n108, n113)
		n115 := Simd_i32x4_add(n101, n114)
		n116 := Simd_i16x8_extend_high_i8x16_s(n112)
		n117 := Simd_i32x4_dot_i16x8_s(n109, n116)
		n118 := Simd_i32x4_add(n104, n117)
		n119 := Simd_i8x16_shuffle(n115, n118, [2]uint64{795458214199165184, 1952900979608391952})
		n120 := Simd_i8x16_shuffle(n115, n118, [2]uint64{1084818905551471876, 2242261670960698644})
		n121 := Simd_i32x4_add(n119, n120)
		n122 := Simd_f32x4_convert_i32x4_s(n121)
		n123 := Simd_f32x4_mul(n8, n122)
		n124 := Simd_f32x4_add([2]uint64{p0, p0h}, n123)
		out0 = n124
		p0, p0h = n124[0], n124[1]
		s0 = s0 + 136
		s3 = s3 - 1
		if s3 <= 1 {
			break
		}
	}
	return out0[0], out0[1], s0, s3
}

//go:noinline
func Simd_p_fxl6(m *Module, s0 int64, s1 int64, s2 int64, s3 int64, p0, p0h uint64) (uint64, uint64, int64, int64, int64) {
	var out0 [2]uint64
	for uint64(s3) >= uint64(4) {
		n0 := Simd_m64_v128_load_rng(m, s0+2, 0, 0, 134)
		n1 := Simd_i16x8_extend_low_i8x16_s(n0)
		n2 := Simd_i16x8_extend_high_i8x16_s(n0)
		n3 := Simd_m64_v128_load_rng(m, s1+2, 0, 0, 134)
		n4 := Simd_i16x8_extend_low_i8x16_s(n3)
		n5 := Simd_i32x4_dot_i16x8_s(n1, n4)
		n6 := Simd_i16x8_extend_high_i8x16_s(n3)
		n7 := Simd_i32x4_dot_i16x8_s(n2, n6)
		n8 := Simd_i32x4_add(n5, n7)
		n9 := Simd_m64_v128_load_nc(m, s0+18, 0)
		n10 := Simd_i16x8_extend_high_i8x16_s(n9)
		n11 := Simd_i16x8_extend_low_i8x16_s(n9)
		n12 := Simd_m64_v128_load_nc(m, s1+18, 0)
		n13 := Simd_i16x8_extend_high_i8x16_s(n12)
		n14 := Simd_i32x4_dot_i16x8_s(n10, n13)
		n15 := Simd_i32x4_add(n8, n14)
		n16 := Simd_i16x8_extend_low_i8x16_s(n12)
		n17 := Simd_i32x4_dot_i16x8_s(n11, n16)
		n18 := Simd_i32x4_add(n15, n17)
		n19 := Simd_f32x4_convert_i32x4_s(n18)
		n20 := Simd_m64_scalar_i32_load16_u(m, s0)
		n21 := Simd_m64_scalar_i32_shl(n20, 2)
		n22 := Simd_m64_scalar_i32_add(n21, s2)
		n23 := Simd_m64_scalar_f32_load(m, n22)
		n24 := Simd_m64_scalar_i32_load16_u(m, s1)
		n25 := Simd_m64_scalar_i32_shl(n24, 2)
		n26 := Simd_m64_scalar_i32_add(n25, s2)
		n27 := Simd_m64_scalar_f32_load(m, n26)
		n28 := Simd_scalar_f32_mul(n23, n27)
		n29 := Simd_f32x4_splat(n28)
		n30 := Simd_f32x4_mul(n29, n19)
		n31 := Simd_f32x4_add([2]uint64{p0, p0h}, n30)
		n32 := Simd_m64_v128_load_nc(m, s0+36, 0)
		n33 := Simd_i16x8_extend_low_i8x16_s(n32)
		n34 := Simd_i16x8_extend_high_i8x16_s(n32)
		n35 := Simd_m64_v128_load_nc(m, s1+36, 0)
		n36 := Simd_i16x8_extend_low_i8x16_s(n35)
		n37 := Simd_i32x4_dot_i16x8_s(n33, n36)
		n38 := Simd_i16x8_extend_high_i8x16_s(n35)
		n39 := Simd_i32x4_dot_i16x8_s(n34, n38)
		n40 := Simd_i32x4_add(n37, n39)
		n41 := Simd_m64_v128_load_nc(m, s0+52, 0)
		n42 := Simd_i16x8_extend_high_i8x16_s(n41)
		n43 := Simd_i16x8_extend_low_i8x16_s(n41)
		n44 := Simd_m64_v128_load_nc(m, s1+52, 0)
		n45 := Simd_i16x8_extend_high_i8x16_s(n44)
		n46 := Simd_i32x4_dot_i16x8_s(n42, n45)
		n47 := Simd_i32x4_add(n40, n46)
		n48 := Simd_i16x8_extend_low_i8x16_s(n44)
		n49 := Simd_i32x4_dot_i16x8_s(n43, n48)
		n50 := Simd_i32x4_add(n47, n49)
		n51 := Simd_f32x4_convert_i32x4_s(n50)
		n52 := Simd_m64_scalar_i32_load16_u(m, s0+34)
		n53 := Simd_m64_scalar_i32_shl(n52, 2)
		n54 := Simd_m64_scalar_i32_add(n53, s2)
		n55 := Simd_m64_scalar_f32_load(m, n54)
		n56 := Simd_m64_scalar_i32_load16_u(m, s1+34)
		n57 := Simd_m64_scalar_i32_shl(n56, 2)
		n58 := Simd_m64_scalar_i32_add(n57, s2)
		n59 := Simd_m64_scalar_f32_load(m, n58)
		n60 := Simd_scalar_f32_mul(n55, n59)
		n61 := Simd_f32x4_splat(n60)
		n62 := Simd_f32x4_mul(n61, n51)
		n63 := Simd_f32x4_add(n31, n62)
		n64 := Simd_m64_v128_load_nc(m, s0+70, 0)
		n65 := Simd_i16x8_extend_low_i8x16_s(n64)
		n66 := Simd_i16x8_extend_high_i8x16_s(n64)
		n67 := Simd_m64_v128_load_nc(m, s1+70, 0)
		n68 := Simd_i16x8_extend_low_i8x16_s(n67)
		n69 := Simd_i32x4_dot_i16x8_s(n65, n68)
		n70 := Simd_i16x8_extend_high_i8x16_s(n67)
		n71 := Simd_i32x4_dot_i16x8_s(n66, n70)
		n72 := Simd_i32x4_add(n69, n71)
		n73 := Simd_m64_v128_load_nc(m, s0+86, 0)
		n74 := Simd_i16x8_extend_high_i8x16_s(n73)
		n75 := Simd_i16x8_extend_low_i8x16_s(n73)
		n76 := Simd_m64_v128_load_nc(m, s1+86, 0)
		n77 := Simd_i16x8_extend_high_i8x16_s(n76)
		n78 := Simd_i32x4_dot_i16x8_s(n74, n77)
		n79 := Simd_i32x4_add(n72, n78)
		n80 := Simd_i16x8_extend_low_i8x16_s(n76)
		n81 := Simd_i32x4_dot_i16x8_s(n75, n80)
		n82 := Simd_i32x4_add(n79, n81)
		n83 := Simd_f32x4_convert_i32x4_s(n82)
		n84 := Simd_m64_scalar_i32_load16_u(m, s0+68)
		n85 := Simd_m64_scalar_i32_shl(n84, 2)
		n86 := Simd_m64_scalar_i32_add(n85, s2)
		n87 := Simd_m64_scalar_f32_load(m, n86)
		n88 := Simd_m64_scalar_i32_load16_u(m, s1+68)
		n89 := Simd_m64_scalar_i32_shl(n88, 2)
		n90 := Simd_m64_scalar_i32_add(n89, s2)
		n91 := Simd_m64_scalar_f32_load(m, n90)
		n92 := Simd_scalar_f32_mul(n87, n91)
		n93 := Simd_f32x4_splat(n92)
		n94 := Simd_f32x4_mul(n93, n83)
		n95 := Simd_f32x4_add(n63, n94)
		n96 := Simd_m64_v128_load_nc(m, s0+104, 0)
		n97 := Simd_i16x8_extend_low_i8x16_s(n96)
		n98 := Simd_i16x8_extend_high_i8x16_s(n96)
		n99 := Simd_m64_v128_load_nc(m, s1+104, 0)
		n100 := Simd_i16x8_extend_low_i8x16_s(n99)
		n101 := Simd_i32x4_dot_i16x8_s(n97, n100)
		n102 := Simd_i16x8_extend_high_i8x16_s(n99)
		n103 := Simd_i32x4_dot_i16x8_s(n98, n102)
		n104 := Simd_i32x4_add(n101, n103)
		n105 := Simd_m64_v128_load_nc(m, s0+120, 0)
		n106 := Simd_i16x8_extend_high_i8x16_s(n105)
		n107 := Simd_i16x8_extend_low_i8x16_s(n105)
		n108 := Simd_m64_v128_load_nc(m, s1+120, 0)
		n109 := Simd_i16x8_extend_high_i8x16_s(n108)
		n110 := Simd_i32x4_dot_i16x8_s(n106, n109)
		n111 := Simd_i32x4_add(n104, n110)
		n112 := Simd_i16x8_extend_low_i8x16_s(n108)
		n113 := Simd_i32x4_dot_i16x8_s(n107, n112)
		n114 := Simd_i32x4_add(n111, n113)
		n115 := Simd_f32x4_convert_i32x4_s(n114)
		n116 := Simd_m64_scalar_i32_load16_u(m, s0+102)
		n117 := Simd_m64_scalar_i32_shl(n116, 2)
		n118 := Simd_m64_scalar_i32_add(n117, s2)
		n119 := Simd_m64_scalar_f32_load(m, n118)
		n120 := Simd_m64_scalar_i32_load16_u(m, s1+102)
		n121 := Simd_m64_scalar_i32_shl(n120, 2)
		n122 := Simd_m64_scalar_i32_add(n121, s2)
		n123 := Simd_m64_scalar_f32_load(m, n122)
		n124 := Simd_scalar_f32_mul(n119, n123)
		n125 := Simd_f32x4_splat(n124)
		n126 := Simd_f32x4_mul(n125, n115)
		n127 := Simd_f32x4_add(n95, n126)
		out0 = n127
		p0, p0h = n127[0], n127[1]
		s1 = s1 + 136
		s0 = s0 + 136
		s3 = s3 - 4
	}
	return out0[0], out0[1], s1, s0, s3
}

//go:noinline
func Simd_p_fxl7(m *Module, s0 int64, s1 int64, s2 int32, p0, p0h uint64) (uint64, uint64, int64, int64, int32) {
	var out0 [2]uint64
	for {
		n0 := Simd_m64_scalar_i32_load16_u(m, s0)
		n1 := Simd_m64_scalar_i32_shl(n0, 2)
		n2 := Simd_m64_v128_load32_splat(m, n1, 8793760)
		n3 := Simd_m64_v128_load16x4_u(m, s1, 0)
		n4 := Simd_f16x4_cvt(n3)
		n5 := Simd_f32x4_mul(n2, n4)
		n6 := Simd_m64_v128_load(m, s1+120, 0)
		n7 := Simd_i16x8_extend_low_i8x16_s(n6)
		n8 := Simd_i16x8_extend_high_i8x16_s(n6)
		n9 := Simd_m64_v128_load32_splat(m, s0+30, 0)
		n10 := Simd_i16x8_extend_low_i8x16_s(n9)
		n11 := Simd_i32x4_dot_i16x8_s(n7, n10)
		n12 := Simd_i16x8_extend_high_i8x16_s(n9)
		n13 := Simd_i32x4_dot_i16x8_s(n8, n12)
		n14 := Simd_m64_v128_load(m, s1+104, 0)
		n15 := Simd_i16x8_extend_low_i8x16_s(n14)
		n16 := Simd_i16x8_extend_high_i8x16_s(n14)
		n17 := Simd_m64_v128_load32_splat(m, s0+26, 0)
		n18 := Simd_i16x8_extend_low_i8x16_s(n17)
		n19 := Simd_i32x4_dot_i16x8_s(n15, n18)
		n20 := Simd_i32x4_add(n11, n19)
		n21 := Simd_i16x8_extend_high_i8x16_s(n17)
		n22 := Simd_i32x4_dot_i16x8_s(n16, n21)
		n23 := Simd_i32x4_add(n13, n22)
		n24 := Simd_m64_v128_load(m, s1+88, 0)
		n25 := Simd_i16x8_extend_low_i8x16_s(n24)
		n26 := Simd_i16x8_extend_high_i8x16_s(n24)
		n27 := Simd_m64_v128_load32_splat(m, s0+22, 0)
		n28 := Simd_i16x8_extend_low_i8x16_s(n27)
		n29 := Simd_i32x4_dot_i16x8_s(n25, n28)
		n30 := Simd_i32x4_add(n20, n29)
		n31 := Simd_i16x8_extend_high_i8x16_s(n27)
		n32 := Simd_i32x4_dot_i16x8_s(n26, n31)
		n33 := Simd_i32x4_add(n23, n32)
		n34 := Simd_m64_v128_load(m, s1+72, 0)
		n35 := Simd_i16x8_extend_low_i8x16_s(n34)
		n36 := Simd_i16x8_extend_high_i8x16_s(n34)
		n37 := Simd_m64_v128_load32_splat(m, s0+18, 0)
		n38 := Simd_i16x8_extend_low_i8x16_s(n37)
		n39 := Simd_i32x4_dot_i16x8_s(n35, n38)
		n40 := Simd_i32x4_add(n30, n39)
		n41 := Simd_i16x8_extend_high_i8x16_s(n37)
		n42 := Simd_i32x4_dot_i16x8_s(n36, n41)
		n43 := Simd_i32x4_add(n33, n42)
		n44 := Simd_m64_v128_load(m, s1+56, 0)
		n45 := Simd_i16x8_extend_low_i8x16_s(n44)
		n46 := Simd_i16x8_extend_high_i8x16_s(n44)
		n47 := Simd_m64_v128_load32_splat(m, s0+14, 0)
		n48 := Simd_i16x8_extend_low_i8x16_s(n47)
		n49 := Simd_i32x4_dot_i16x8_s(n45, n48)
		n50 := Simd_i32x4_add(n40, n49)
		n51 := Simd_i16x8_extend_high_i8x16_s(n47)
		n52 := Simd_i32x4_dot_i16x8_s(n46, n51)
		n53 := Simd_i32x4_add(n43, n52)
		n54 := Simd_m64_v128_load(m, s1+40, 0)
		n55 := Simd_i16x8_extend_low_i8x16_s(n54)
		n56 := Simd_i16x8_extend_high_i8x16_s(n54)
		n57 := Simd_m64_v128_load32_splat(m, s0+10, 0)
		n58 := Simd_i16x8_extend_low_i8x16_s(n57)
		n59 := Simd_i32x4_dot_i16x8_s(n55, n58)
		n60 := Simd_i32x4_add(n50, n59)
		n61 := Simd_i16x8_extend_high_i8x16_s(n57)
		n62 := Simd_i32x4_dot_i16x8_s(n56, n61)
		n63 := Simd_i32x4_add(n53, n62)
		n64 := Simd_m64_v128_load(m, s1+24, 0)
		n65 := Simd_i16x8_extend_low_i8x16_s(n64)
		n66 := Simd_i16x8_extend_high_i8x16_s(n64)
		n67 := Simd_m64_v128_load32_splat(m, s0+6, 0)
		n68 := Simd_i16x8_extend_low_i8x16_s(n67)
		n69 := Simd_i32x4_dot_i16x8_s(n65, n68)
		n70 := Simd_i32x4_add(n60, n69)
		n71 := Simd_i16x8_extend_high_i8x16_s(n67)
		n72 := Simd_i32x4_dot_i16x8_s(n66, n71)
		n73 := Simd_i32x4_add(n63, n72)
		n74 := Simd_m64_v128_load(m, s1+8, 0)
		n75 := Simd_i16x8_extend_low_i8x16_s(n74)
		n76 := Simd_i16x8_extend_high_i8x16_s(n74)
		n77 := Simd_m64_v128_load32_splat(m, s0+2, 0)
		n78 := Simd_i16x8_extend_low_i8x16_s(n77)
		n79 := Simd_i32x4_dot_i16x8_s(n75, n78)
		n80 := Simd_i32x4_add(n70, n79)
		n81 := Simd_i16x8_extend_high_i8x16_s(n77)
		n82 := Simd_i32x4_dot_i16x8_s(n76, n81)
		n83 := Simd_i32x4_add(n73, n82)
		n84 := Simd_i8x16_shuffle(n80, n83, [2]uint64{795458214199165184, 1952900979608391952})
		n85 := Simd_i8x16_shuffle(n80, n83, [2]uint64{1084818905551471876, 2242261670960698644})
		n86 := Simd_i32x4_add(n84, n85)
		n87 := Simd_f32x4_convert_i32x4_s(n86)
		n88 := Simd_f32x4_mul(n5, n87)
		n89 := Simd_f32x4_add([2]uint64{p0, p0h}, n88)
		out0 = n89
		p0, p0h = n89[0], n89[1]
		s0 = s0 + 34
		s1 = s1 + 136
		s2 = s2 - 1
		if s2 <= 1 {
			break
		}
	}
	return out0[0], out0[1], s0, s1, s2
}

var spinRelaxColdCalls uint32

type ThreadPool struct {
	nextTID atomic.Int32
	wg      sync.WaitGroup

	parkMu sync.Mutex
	parked map[uint64][]chan struct{}
}

// wake releases up to count waiters on ea and reports how many it woke.
func (p *ThreadPool) wake(ea uint64, count int32) int32 {
	p.parkMu.Lock()
	defer p.parkMu.Unlock()
	waiters := p.parked[ea]
	n := int32(len(waiters))
	if count >= 0 && count < n {
		n = count
	}
	for _, ch := range waiters[:n] {
		close(ch)
	}
	if int(n) == len(waiters) {
		delete(p.parked, ea)
	} else {
		p.parked[ea] = waiters[n:]
	}
	return n
}

// SaveGlobals returns the module's mutable globals, in a form that can be handed back
// to RestoreGlobals. It is how a snapshot of an instance captures the state that does not
// live in linear memory.
func SaveGlobals(m *Module) []uint64 {
	g := make([]uint64, 2)
	g[0] = uint64(m.G0)
	g[1] = uint64(m.G1)
	return g
}

// RestoreGlobals puts a snapshot's globals back. A snapshot from a different module (or a
// different build of the same one) has a different global count; rather than
// index out of bounds, take what fits and leave the rest at their declared
// initializers.
func RestoreGlobals(m *Module, g []uint64) {
	if len(g) != 2 {
		return
	}
	m.G0 = int64(g[0])
	m.G1 = int64(g[1])
}

// WasiExitError is the sentinel that the recover layer of SafeInvokeExport
// promotes Proc_exit() panics into, so a wasm-level exit doesn't kill the
// host process and the caller can read the exit code instead.
type WasiExitError struct{ Code int32 }

func (e *WasiExitError) Error() string {
	return "wasi: proc_exit(" + itoa32(e.Code) + ")"
}

// itoa32 is a tiny dependency-free strconv replacement so this file
// doesn't drag in fmt for its sole error path.
func itoa32(v int32) string {
	if v == 0 {
		return "0"
	}
	neg := false
	if v < 0 {
		v = -v
		neg = true
	}
	var buf [12]byte
	i := len(buf)
	for v > 0 {
		i--
		buf[i] = byte('0' + v%10)
		v /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}

// FS is the read/write filesystem backend the WASI host opens files through.
// It abstracts the default os-backed filesystem so an embedder can supply an
// alternative — an in-memory FS, an overlay, a read-only bundle, ... — and
// have every guest path operation (open, stat, mkdir, readdir, write, ...)
// routed to it. It is a write-capable superset of io/fs.FS.
//
// Names are GUEST paths relative to the preopen root: slash-separated, with no
// leading slash (e.g. "encodings/__init__.py", or "" for the root). Methods
// should return the standard fs errors (fs.ErrNotExist, fs.ErrExist,
// fs.ErrPermission) so the host maps them to the right wasi errno.
type FS interface {
	// OpenFile mirrors os.OpenFile: flag is O_RDONLY/O_WRONLY/O_RDWR optionally
	// OR'd with O_CREATE/O_EXCL/O_TRUNC/O_APPEND. The returned File must
	// support the operations the mode implies.
	OpenFile(name string, flag int, perm os.FileMode) (File, error)
	Mkdir(name string, perm os.FileMode) error
	Remove(name string) error
	Rename(oldName, newName string) error
	Stat(name string) (os.FileInfo, error)
	Lstat(name string) (os.FileInfo, error)
	Symlink(oldName, newName string) error
	Readlink(name string) (string, error)
	Link(oldName, newName string) error
}

// File is an open file handle returned by FS.OpenFile. *os.File satisfies it,
// so the default os backend needs no wrapper.
type File interface {
	Read(p []byte) (int, error)
	ReadAt(p []byte, off int64) (int, error)
	Write(p []byte) (int, error)
	WriteAt(p []byte, off int64) (int, error)
	Seek(offset int64, whence int) (int64, error)
	Close() error
	Stat() (os.FileInfo, error)
	ReadDir(n int) ([]os.DirEntry, error)
	Sync() error
	Truncate(size int64) error
	Name() string
}

// osFS is the default FS backend: a thin pass-through to the host filesystem,
// scoped to root (the preopen directory). root "" or "/" means no rewriting.
type osFS struct{ root string }

func (o osFS) join(name string) string {
	if o.root == "" || o.root == "/" {
		return "/" + name
	}
	return filepath.Join(o.root, name)
}
func (o osFS) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	f, err := os.OpenFile(o.join(name), flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
func (o osFS) Mkdir(name string, perm os.FileMode) error { return os.Mkdir(o.join(name), perm) }
func (o osFS) Chmod(name string, mode os.FileMode) error { return os.Chmod(o.join(name), mode) }
func (o osFS) Remove(name string) error                  { return os.Remove(o.join(name)) }
func (o osFS) Rename(a, b string) error                  { return os.Rename(o.join(a), o.join(b)) }
func (o osFS) Stat(name string) (os.FileInfo, error)     { return os.Stat(o.join(name)) }
func (o osFS) Lstat(name string) (os.FileInfo, error)    { return os.Lstat(o.join(name)) }
func (o osFS) Symlink(target, name string) error         { return os.Symlink(target, o.join(name)) }
func (o osFS) Readlink(name string) (string, error)      { return os.Readlink(o.join(name)) }
func (o osFS) Link(a, b string) error                    { return os.Link(o.join(a), o.join(b)) }

// MemFS is an in-memory read/write FS. Each value is an independent tree, so
// two interpreters given separate MemFS values cannot observe each other's
// files (full per-interpreter filesystem isolation, no disk). Build one with
// NewMemFS. Safe for concurrent use.
type MemFS struct {
	mu   sync.Mutex
	root *memNode
}

// NewMemFS returns an empty in-memory filesystem with a root directory.
func NewMemFS() *MemFS {
	return &MemFS{root: &memNode{dir: true, mode: os.ModeDir | 0o755, modTime: time.Unix(0, 0), children: map[string]*memNode{}}}
}

// memNode is a file or directory in a MemFS tree.
type memNode struct {
	name     string
	dir      bool
	mode     os.FileMode
	modTime  time.Time
	data     []byte
	children map[string]*memNode
}

func memSplit(name string) []string {
	name = strings.Trim(name, "/")
	if name == "" {
		return nil
	}
	raw := strings.Split(name, "/")
	out := make([]string, 0, len(raw))
	for _, p := range raw {
		switch p {
		case "", ".":
		case "..":
			if len(out) > 0 {
				out = out[:len(out)-1]
			}
		default:
			out = append(out, p)
		}
	}
	return out
}

// lookup resolves name to a node. Caller holds fsys.mu.
func (fsys *MemFS) lookup(name string) (*memNode, error) {
	n := fsys.root
	for _, part := range memSplit(name) {
		if !n.dir {
			return nil, fs.ErrNotExist
		}
		c, ok := n.children[part]
		if !ok {
			return nil, fs.ErrNotExist
		}
		n = c
	}
	return n, nil
}

// lookupParent resolves the parent dir of name. Caller holds fsys.mu.
func (fsys *MemFS) lookupParent(name string) (*memNode, string, error) {
	parts := memSplit(name)
	if len(parts) == 0 {
		return nil, "", fs.ErrInvalid
	}
	n := fsys.root
	for _, part := range parts[:len(parts)-1] {
		c, ok := n.children[part]
		if !ok || !c.dir {
			return nil, "", fs.ErrNotExist
		}
		n = c
	}
	return n, parts[len(parts)-1], nil
}

func (fsys *MemFS) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	fsys.mu.Lock()
	defer fsys.mu.Unlock()
	node, err := fsys.lookup(name)
	if err != nil {
		if flag&os.O_CREATE == 0 {
			return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrNotExist}
		}
		parent, base, perr := fsys.lookupParent(name)
		if perr != nil {
			return nil, &fs.PathError{Op: "open", Path: name, Err: perr}
		}
		node = &memNode{name: base, mode: perm & 0o777, modTime: time.Now()}
		parent.children[base] = node
	} else {
		if flag&os.O_EXCL != 0 && flag&os.O_CREATE != 0 {
			return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrExist}
		}
		if flag&os.O_TRUNC != 0 && !node.dir {
			node.data = node.data[:0]
			node.modTime = time.Now()
		}
	}
	f := &memFile{fsys: fsys, node: node}
	if flag&os.O_APPEND != 0 {
		f.off = int64(len(node.data))
	}
	return f, nil
}

func (fsys *MemFS) Mkdir(name string, perm os.FileMode) error {
	fsys.mu.Lock()
	defer fsys.mu.Unlock()
	parent, base, err := fsys.lookupParent(name)
	if err != nil {
		return &fs.PathError{Op: "mkdir", Path: name, Err: err}
	}
	if _, ok := parent.children[base]; ok {
		return &fs.PathError{Op: "mkdir", Path: name, Err: fs.ErrExist}
	}
	parent.children[base] = &memNode{name: base, dir: true, mode: os.ModeDir | (perm & 0o777), modTime: time.Now(), children: map[string]*memNode{}}
	return nil
}

func (fsys *MemFS) Remove(name string) error {
	fsys.mu.Lock()
	defer fsys.mu.Unlock()
	parent, base, err := fsys.lookupParent(name)
	if err != nil {
		return &fs.PathError{Op: "remove", Path: name, Err: err}
	}
	n, ok := parent.children[base]
	if !ok {
		return &fs.PathError{Op: "remove", Path: name, Err: fs.ErrNotExist}
	}
	if n.dir && len(n.children) > 0 {
		return &fs.PathError{Op: "remove", Path: name, Err: fs.ErrInvalid}
	}
	delete(parent.children, base)
	return nil
}

func (fsys *MemFS) Rename(oldName, newName string) error {
	fsys.mu.Lock()
	defer fsys.mu.Unlock()
	op, ob, err := fsys.lookupParent(oldName)
	if err != nil {
		return &fs.PathError{Op: "rename", Path: oldName, Err: err}
	}
	node, ok := op.children[ob]
	if !ok {
		return &fs.PathError{Op: "rename", Path: oldName, Err: fs.ErrNotExist}
	}
	np, nb, err := fsys.lookupParent(newName)
	if err != nil {
		return &fs.PathError{Op: "rename", Path: newName, Err: err}
	}
	delete(op.children, ob)
	node.name = nb
	np.children[nb] = node
	return nil
}

func (fsys *MemFS) Stat(name string) (os.FileInfo, error) {
	fsys.mu.Lock()
	defer fsys.mu.Unlock()
	n, err := fsys.lookup(name)
	if err != nil {
		return nil, &fs.PathError{Op: "stat", Path: name, Err: err}
	}
	return n.info(), nil
}

func (fsys *MemFS) Lstat(name string) (os.FileInfo, error) { return fsys.Stat(name) }

// memfs has no symlinks/hardlinks.
func (fsys *MemFS) Symlink(_, _ string) error { return fs.ErrPermission }
func (fsys *MemFS) Link(_, _ string) error    { return fs.ErrPermission }
func (fsys *MemFS) Readlink(name string) (string, error) {
	return "", &fs.PathError{Op: "readlink", Path: name, Err: fs.ErrInvalid}
}

// Chtimes implements the optional chtimesFS capability.
func (fsys *MemFS) Chtimes(name string, _ time.Time, mtime time.Time) error {
	fsys.mu.Lock()
	defer fsys.mu.Unlock()
	n, err := fsys.lookup(name)
	if err != nil {
		return &fs.PathError{Op: "chtimes", Path: name, Err: err}
	}
	n.modTime = mtime
	return nil
}

// MkdirAll creates name and any missing parents. Exposed so embedders can
// populate the FS (e.g. unpack a stdlib bundle) before handing it to a module.
func (fsys *MemFS) MkdirAll(name string, perm os.FileMode) error {
	fsys.mu.Lock()
	defer fsys.mu.Unlock()
	n := fsys.root
	for _, part := range memSplit(name) {
		c, ok := n.children[part]
		if !ok {
			c = &memNode{name: part, dir: true, mode: os.ModeDir | (perm & 0o777), modTime: time.Now(), children: map[string]*memNode{}}
			n.children[part] = c
		} else if !c.dir {
			return &fs.PathError{Op: "mkdir", Path: name, Err: fs.ErrExist}
		}
		n = c
	}
	return nil
}

// WriteFile creates (or overwrites) a file with data, making parent dirs as
// needed. Exposed for pre-populating the FS.
func (fsys *MemFS) WriteFile(name string, data []byte, perm os.FileMode) error {
	if parts := memSplit(name); len(parts) > 1 {
		if err := fsys.MkdirAll(strings.Join(parts[:len(parts)-1], "/"), 0o755); err != nil {
			return err
		}
	}
	fsys.mu.Lock()
	defer fsys.mu.Unlock()
	parent, base, err := fsys.lookupParent(name)
	if err != nil {
		return &fs.PathError{Op: "writefile", Path: name, Err: err}
	}
	cp := make([]byte, len(data))
	copy(cp, data)
	parent.children[base] = &memNode{name: base, mode: perm & 0o777, modTime: time.Now(), data: cp}
	return nil
}

func (n *memNode) info() os.FileInfo {
	if n.dir {
		return memFileInfo{name: n.name, mode: os.ModeDir | (n.mode & 0o777), modTime: n.modTime}
	}
	return memFileInfo{name: n.name, size: int64(len(n.data)), mode: n.mode & 0o777, modTime: n.modTime}
}

type memFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi memFileInfo) Name() string       { return fi.name }
func (fi memFileInfo) Size() int64        { return fi.size }
func (fi memFileInfo) Mode() os.FileMode  { return fi.mode }
func (fi memFileInfo) ModTime() time.Time { return fi.modTime }
func (fi memFileInfo) IsDir() bool        { return fi.mode.IsDir() }
func (fi memFileInfo) Sys() any           { return nil }

type memDirEntry struct{ n *memNode }

func (e memDirEntry) Name() string { return e.n.name }
func (e memDirEntry) IsDir() bool  { return e.n.dir }
func (e memDirEntry) Type() os.FileMode {
	if e.n.dir {
		return os.ModeDir
	}
	return 0
}
func (e memDirEntry) Info() (os.FileInfo, error) { return e.n.info(), nil }

// memFile is an open handle into a MemFS node.
type memFile struct {
	fsys   *MemFS
	node   *memNode
	off    int64
	dirOff int
}

func (f *memFile) Name() string { return f.node.name }
func (f *memFile) Close() error { return nil }
func (f *memFile) Sync() error  { return nil }

func (f *memFile) Stat() (os.FileInfo, error) {
	f.fsys.mu.Lock()
	defer f.fsys.mu.Unlock()
	return f.node.info(), nil
}

func (f *memFile) Read(p []byte) (int, error) {
	f.fsys.mu.Lock()
	defer f.fsys.mu.Unlock()
	if f.node.dir {
		return 0, &fs.PathError{Op: "read", Path: f.node.name, Err: fs.ErrInvalid}
	}
	if f.off >= int64(len(f.node.data)) {
		return 0, io.EOF
	}
	n := copy(p, f.node.data[f.off:])
	f.off += int64(n)
	return n, nil
}

func (f *memFile) ReadAt(p []byte, off int64) (int, error) {
	f.fsys.mu.Lock()
	defer f.fsys.mu.Unlock()
	if off >= int64(len(f.node.data)) {
		return 0, io.EOF
	}
	n := copy(p, f.node.data[off:])
	if n < len(p) {
		return n, io.EOF
	}
	return n, nil
}

// writeAt grows node.data as needed and writes p at off. Caller holds the lock.
func (f *memFile) writeAt(p []byte, off int64) int {
	end := off + int64(len(p))
	if end > int64(len(f.node.data)) {
		grown := make([]byte, end)
		copy(grown, f.node.data)
		f.node.data = grown
	}
	copy(f.node.data[off:], p)
	f.node.modTime = time.Now()
	return len(p)
}

func (f *memFile) Write(p []byte) (int, error) {
	f.fsys.mu.Lock()
	defer f.fsys.mu.Unlock()
	n := f.writeAt(p, f.off)
	f.off += int64(n)
	return n, nil
}

func (f *memFile) WriteAt(p []byte, off int64) (int, error) {
	f.fsys.mu.Lock()
	defer f.fsys.mu.Unlock()
	return f.writeAt(p, off), nil
}

func (f *memFile) Seek(offset int64, whence int) (int64, error) {
	f.fsys.mu.Lock()
	defer f.fsys.mu.Unlock()
	switch whence {
	case io.SeekStart:
		f.off = offset
	case io.SeekCurrent:
		f.off += offset
	case io.SeekEnd:
		f.off = int64(len(f.node.data)) + offset
	}
	return f.off, nil
}

func (f *memFile) Truncate(size int64) error {
	f.fsys.mu.Lock()
	defer f.fsys.mu.Unlock()
	if size <= int64(len(f.node.data)) {
		f.node.data = f.node.data[:size]
	} else {
		grown := make([]byte, size)
		copy(grown, f.node.data)
		f.node.data = grown
	}
	f.node.modTime = time.Now()
	return nil
}

func (f *memFile) ReadDir(n int) ([]os.DirEntry, error) {
	f.fsys.mu.Lock()
	defer f.fsys.mu.Unlock()
	if !f.node.dir {
		return nil, &fs.PathError{Op: "readdir", Path: f.node.name, Err: fs.ErrInvalid}
	}
	names := make([]string, 0, len(f.node.children))
	for name := range f.node.children {
		names = append(names, name)
	}
	sort.Strings(names)
	if f.dirOff >= len(names) {
		if n <= 0 {
			return nil, nil
		}
		return nil, io.EOF
	}
	end := len(names)
	if n > 0 && f.dirOff+n < end {
		end = f.dirOff + n
	}
	out := make([]os.DirEntry, 0, end-f.dirOff)
	for _, name := range names[f.dirOff:end] {
		out = append(out, memDirEntry{f.node.children[name]})
	}
	f.dirOff = end
	return out, nil
}

// wasiOpen is one entry in WasiStubs' fd table. Stdio entries are nil-file
// markers (writes go to the OS handles directly via the WasiStubs fields).
// The conn arm carries a net.Conn for sockets opened via Sock_accept.
type wasiOpen struct {
	f        File
	conn     net.Conn
	listener net.Listener
	isDir    bool
	isSocket bool   // created by Sock_socket, may not yet have a conn
	path     string // guest path relative to the preopen root
	fdflags  int32  // last fdflags set via Path_open or Fd_fdstat_set_flags
	dirCache []os.DirEntry
	// stdio marks an alias of an interpreter stream (1/2/3 = the
	// configured stdin/stdout/stderr; 0 = not an alias). Fd_dup of a bare
	// fd 0/1/2 creates one; closing it never touches the real stream.
	stdio int8
	// refs counts EXTRA table slots sharing this entry (dup/dup2):
	// closeWasiOpen only closes the descriptor when it reaches zero.
	refs int32
}

// WasiStubs is the default Go-native implementation of wasi_snapshot_preview1.
// State is owned per-Module via NewWithWASI / DefaultWASI.
type WasiStubs struct {
	mu sync.Mutex

	// stdin/stdout/stderr back guest fds 0/1/2. They default to the host
	// os.Std* (DefaultWASI) but can be redirected to any io.Reader/io.Writer
	// (an in-process buffer, pipe, ...) via SetStdin/SetStdout/SetStderr, so an
	// embedder can feed input and capture/stream output without touching the
	// host process stdio.
	stdin          io.Reader
	stdout, stderr io.Writer
	fdTable        map[int32]*wasiOpen
	nextFD         int32
	args, env      []string
	monoStart      time.Time
	// preopenDir is the host directory mapped to wasi preopen fd 3.
	// Defaults to "/" (i.e. no rewriting) — the legacy behaviour. Tests
	// can set this via SetPreopenDir to scope filesystem ops to a
	// temporary directory.
	preopenDir string
	// fsHook, when non-nil, is consulted before every filesystem access
	// (Path_open, Path_create_directory, Path_unlink_file). It receives
	// the guest-supplied path (relative to the preopen, e.g. "a.txt" or
	// "sub/a.txt") and whether the access is a write. Returning false
	// denies the operation, which surfaces to the guest as EACCES. This
	// is the host-controlled whitelist hook: the policy itself lives in
	// the embedding application, OUTSIDE the generated runtime.
	fsHook func(path string, write bool) bool
	// netHook, when non-nil, is consulted before every socket operation
	// (Sock_accept, Sock_recv, Sock_send). op is "accept"/"recv"/"send".
	// Returning false denies the operation (EACCES). The same
	// host-controlled-whitelist intent as fsHook, for the network surface.
	netHook func(op string) bool
	// dialHook, when non-nil, is consulted before an OUTBOUND connect
	// (Sock_connect) with the resolved network ("tcp"), the HOST the guest
	// resolved to reach this address (from the preceding Sock_getaddrinfo, or ""
	// if the guest dialed a literal IP), the dotted-quad IP, and the port.
	// Returning false denies the connection (EACCES). Passing the host lets the
	// policy match host+port jointly, which a port-scoped rule needs — the IP
	// alone cannot be tied back to the rule that authorized the name.
	dialHook func(network, host, ip string, port int) bool
	// resolveHook, when non-nil, is consulted before a name lookup
	// (Sock_getaddrinfo) with the requested host. Returning false denies the
	// resolution (the guest sees a gaierror). This is the hostname-level
	// whitelist control point (e.g. block "example.com" by name).
	resolveHook func(host string) bool
	// resolvedHosts maps a resolved dotted-quad IP back to the host name the
	// guest looked it up under (populated by Sock_getaddrinfo, read by
	// Sock_connect), so the dial hook can be given the host. Guarded by mu.
	resolvedHosts map[string]string
	// fsys is the filesystem backend every guest path operation is routed
	// through. Defaults to an osFS scoped to preopenDir (the host filesystem);
	// SetFS swaps in an alternative (e.g. an in-memory FS) so each module can
	// see a private, arbitrary filesystem.
	fsys FS
	// procs tracks host processes spawned via Proc_spawn, keyed by the pid
	// handed back to the guest. nextPID is the handle counter (kept distinct
	// from real OS pids — the guest only ever sees these tokens).
	procs   map[int32]*wasiProc
	nextPID int32
	// execHook, when non-nil, gates every Proc_spawn with the resolved
	// executable path and argv; returning false denies the spawn (EACCES).
	// This is the outbound-process whitelist control point — the analogue of
	// dialHook for sockets. Spawning runs a HOST binary, so a sandbox that
	// enables host processes should always install this.
	execHook func(path string, argv []string) bool
}

// wasiProc is a host process spawned by Proc_spawn. A background goroutine
// Waits on the command and publishes the encoded POSIX status, so Proc_wait
// can support both the blocking (options 0) and non-blocking (WNOHANG) forms
// without holding the WasiStubs lock across the child's lifetime.
type wasiProc struct {
	cmd    *exec.Cmd
	done   chan struct{}
	status int32 // POSIX wait status, valid once done is closed
}

// DefaultWASI returns a WasiStubs configured for typical CLI use: real
// stdio, os.Args, os.Environ(), wall + monotonic clocks. Consumers who
// want a sandboxed setup should construct their own WasiStubs (or any
// Wasi_snapshot_preview1Imports implementation) and pass it to
// NewWithWASI.
func DefaultWASI() *WasiStubs {
	return &WasiStubs{
		stdin:      os.Stdin,
		stdout:     os.Stdout,
		stderr:     os.Stderr,
		fdTable:    map[int32]*wasiOpen{},
		nextFD:     4,
		args:       os.Args,
		env:        os.Environ(),
		monoStart:  time.Now(),
		preopenDir: "/",
		fsys:       osFS{root: "/"},
		procs:      map[int32]*wasiProc{},
		nextPID:    1000,
	}
}

// SetPreopenDir scopes the default (os-backed) filesystem to a host directory.
// Empty string restores the default ("/"), i.e. no rewriting. Tests use this
// to run filesystem syscalls against t.TempDir(). Has no effect once SetFS has
// installed a non-os backend.
func (w *WasiStubs) SetPreopenDir(dir string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if dir == "" {
		dir = "/"
	}
	w.preopenDir = dir
	w.fsys = osFS{root: dir}
}

// SetFS installs a custom filesystem backend. Every guest path operation
// (open, stat, mkdir, readdir, read, write, ...) is then routed to fsys, so a
// caller can give a module a private, arbitrary filesystem — for example an
// in-memory FS so writes never touch disk and are invisible to other modules.
// Pass nil to restore the default os-backed filesystem.
func (w *WasiStubs) SetFS(fsys FS) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if fsys == nil {
		fsys = osFS{root: w.preopenDir}
	}
	w.fsys = fsys
}

// SetFSAccessHook installs a host-controlled filesystem access policy.
// hook is called with the guest path (relative to the preopen) and a
// write flag before each open/create/unlink; returning false denies the
// operation (the guest sees EACCES). Pass nil to clear the policy
// (unrestricted, the default). The hook runs without w.mu held, so it
// may itself call back into the host freely.
func (w *WasiStubs) SetFSAccessHook(hook func(path string, write bool) bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.fsHook = hook
}

// SetNetAccessHook installs a host-controlled network access policy.
// hook is called with the operation name ("accept"/"recv"/"send")
// before each socket operation; returning false denies it (EACCES).
// Pass nil to clear (unrestricted, the default).
//
// NOTE: WASI preview1 has no outbound connect or name resolution, so a
// guest cannot initiate connections regardless of this hook; it governs
// the accept/recv/send surface that preview1 does expose (host-preopened
// listening sockets). Full outbound control requires a host connect
// import, which this runtime does not yet provide.
func (w *WasiStubs) SetNetAccessHook(hook func(op string) bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.netHook = hook
}

// SetDialHook installs a host-controlled OUTBOUND-connection policy. hook is
// called with ("tcp", host, dotted-quad-IP, port) before each Sock_connect,
// where host is the name the guest resolved to reach the IP (from the preceding
// Sock_getaddrinfo) or "" for a literal-IP dial; returning false denies the
// connection (the guest sees a connect EACCES). Pass nil to clear (all outbound
// allowed, the default once outbound is wired).
func (w *WasiStubs) SetDialHook(hook func(network, host, ip string, port int) bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.dialHook = hook
}

// SetResolveHook installs a host-controlled name-resolution policy. hook is
// called with the host being resolved (Sock_getaddrinfo) before the lookup;
// returning false denies it (the guest sees a name-resolution error). Pass nil
// to clear (all lookups allowed). This is where a hostname whitelist such as
// "block example.com" is enforced.
func (w *WasiStubs) SetResolveHook(hook func(host string) bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.resolveHook = hook
}

// SetExecHook installs the process-spawn whitelist consulted by Proc_spawn
// with the executable path and full argv. Returning false denies the spawn
// (the guest's posix_spawn sees EACCES). Spawning runs a HOST binary, so a
// sandbox enabling host processes should always set this.
func (w *WasiStubs) SetExecHook(hook func(path string, argv []string) bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.execHook = hook
}

// readCStr reads a NUL-terminated C string at ptr from linear memory. A nil
// ptr (0) yields "". ok is false on an out-of-bounds or unterminated read.
func (w *WasiStubs) readCStr(m *Module, ptr int32) (s string, ok bool) {
	if ptr == 0 {
		return "", true
	}
	mem := m.Memory
	lo := uint64(uint32(ptr))
	if lo > uint64(len(mem)) {
		return "", false
	}
	rest := mem[lo:]
	for i := 0; i < len(rest); i++ {
		if rest[i] == 0 {
			return string(rest[:i]), true
		}
	}
	return "", false
}

// readCStrArray reads a NULL-terminated array of C-string pointers (a char**)
// at ptr. A nil ptr (0) yields a nil slice. ok is false on a bad read.
func (w *WasiStubs) readCStrArray(m *Module, ptr int32) (out []string, ok bool) {
	if ptr == 0 {
		return nil, true
	}
	for off := ptr; ; off += 4 {
		b := w.memSlice(m, off, 4)
		if b == nil {
			return nil, false
		}
		p := int32(binary.LittleEndian.Uint32(b))
		if p == 0 {
			break
		}
		s, sok := w.readCStr(m, p)
		if !sok {
			return nil, false
		}
		out = append(out, s)
	}
	return out, true
}

// Proc_spawn is a NON-STANDARD host import (module wasi_snapshot_preview1,
// name "proc_spawn") backing the bridge's posix_spawn(). It spawns a HOST
// process: path is the executable, argv/envp are NUL-terminated char** in
// linear memory. The child inherits the interpreter's stdin/stdout/stderr.
// The new pid token is written at pidOutPtr. Returns 0 or a negative errno.
//
// Only stdio inheritance is supported today (no fd remapping / pipes), which
// covers subprocess.run/call with default streams; capture_output via host
// pipes is a follow-up.
func (w *WasiStubs) Proc_spawn(m *Module, pathPtr, argvPtr, envpPtr, stdinFd, stdoutFd, stderrFd, cwdPtr, pidOutPtr int32) int32 {
	path, ok := w.readCStr(m, pathPtr)
	if !ok || path == "" {
		return -_wasiEINVAL
	}
	argv, ok := w.readCStrArray(m, argvPtr)
	if !ok {
		return -_wasiEFAULT
	}
	env, ok := w.readCStrArray(m, envpPtr)
	if !ok {
		return -_wasiEFAULT
	}

	cwd, ok := w.readCStr(m, cwdPtr)
	if !ok {
		return -_wasiEFAULT
	}
	out := w.memSlice(m, pidOutPtr, 4)
	if out == nil {
		return -_wasiEFAULT
	}

	w.mu.Lock()
	hook := w.execHook
	cin := w.childReaderLocked(stdinFd)
	cout := w.childWriterLocked(stdoutFd, w.stdout)
	cerr := w.childWriterLocked(stderrFd, w.stderr)
	w.mu.Unlock()
	if hook != nil && !hook(path, argv) {
		return -_wasiEACCES
	}

	cmd := exec.Command(path)
	if len(argv) > 0 {
		cmd.Args = argv
	} else {
		cmd.Args = []string{path}
	}
	if cwd != "" {
		cmd.Dir = cwd
	}

	cmd.Env = env
	if cmd.Env == nil {
		cmd.Env = []string{}
	}

	cmd.Stdin, cmd.Stdout, cmd.Stderr = cin, cout, cerr
	if err := cmd.Start(); err != nil {
		return -mapExecError(err)
	}

	proc := &wasiProc{cmd: cmd, done: make(chan struct{})}
	go func() {
		werr := cmd.Wait()
		proc.status = encodeWaitStatus(cmd.ProcessState)
		// A non-zero exit or signal surfaces as *exec.ExitError and is the
		// normal path (status already encoded from ProcessState above). Any
		// OTHER error means the wait itself failed; report a 127 exit.
		var exitErr *exec.ExitError
		if werr != nil && !errors.As(werr, &exitErr) {
			proc.status = int32(127) << 8
		}
		close(proc.done)
	}()

	w.mu.Lock()
	if w.procs == nil {
		w.procs = map[int32]*wasiProc{}
	}
	if w.nextPID == 0 {
		w.nextPID = 1000
	}
	pid := w.nextPID
	w.nextPID++
	w.procs[pid] = proc
	w.mu.Unlock()

	binary.LittleEndian.PutUint32(out, uint32(pid))
	return _wasiESUCCESS
}

// childReaderLocked resolves a child stdin source fd. A guest fd whose
// table entry carries a real file (a pipe end, or a guest stdio fd the
// program re-opened onto a file) is used directly; everything else —
// including -1 and an unredirected fd 0 — inherits the interpreter's
// stdin. Caller holds w.mu.
func (w *WasiStubs) childReaderLocked(fd int32) io.Reader {
	if fd >= 0 {
		if op := w.fdTable[fd]; op != nil && op.f != nil {
			return op.f
		}
	}
	return w.stdin
}

// childWriterLocked is the stdout/stderr counterpart of childReaderLocked.
func (w *WasiStubs) childWriterLocked(fd int32, deflt io.Writer) io.Writer {
	if fd >= 0 {
		if op := w.fdTable[fd]; op != nil && op.f != nil {
			return op.f
		}
	}
	return deflt
}

// Pipe is a NON-STANDARD host import (module wasi_snapshot_preview1, name
// "pipe") backing the bridge's pipe()/pipe2(). It creates a host OS pipe and
// registers both ends as guest fds, writing [readFd, writeFd] (two i32) at
// fdsOutPtr. The guest reads the read end via Fd_read; the write end is given
// to a child as its stdout/stderr via Proc_spawn, so subprocess.run can
// capture output. Returns 0 or a negative errno.
func (w *WasiStubs) Pipe(m *Module, fdsOutPtr int32) int32 {
	out := w.memSlice(m, fdsOutPtr, 8)
	if out == nil {
		return -_wasiEFAULT
	}
	r, wr, err := os.Pipe()
	if err != nil {
		return -mapOSError(err)
	}
	w.mu.Lock()
	if w.fdTable == nil {
		w.fdTable = map[int32]*wasiOpen{}
	}
	if w.nextFD < 4 {
		w.nextFD = 4
	}
	rfd := w.nextFD
	w.nextFD++
	wfd := w.nextFD
	w.nextFD++
	w.fdTable[rfd] = &wasiOpen{f: r}
	w.fdTable[wfd] = &wasiOpen{f: wr}
	w.mu.Unlock()
	binary.LittleEndian.PutUint32(out[0:], uint32(rfd))
	binary.LittleEndian.PutUint32(out[4:], uint32(wfd))
	return _wasiESUCCESS
}

// Proc_wait is a NON-STANDARD host import (name "proc_wait") backing the
// bridge's waitpid(). It waits for the process token pid and writes the POSIX
// wait status at statusOutPtr. options is the waitpid() options mask; bit 0
// (WNOHANG) makes it return without blocking when the child is still running
// (the guest sees the documented "0 means no child ready" result, signalled
// by writing pid 0 — encoded by returning EAGAIN). Returns 0, or a negative
// errno (ECHILD for an unknown pid).
func (w *WasiStubs) Proc_wait(m *Module, pid, options, statusOutPtr int32) int32 {
	out := w.memSlice(m, statusOutPtr, 4)
	if out == nil {
		return -_wasiEFAULT
	}
	w.mu.Lock()
	proc := w.procs[pid]
	w.mu.Unlock()
	if proc == nil {
		return -_wasiECHILD
	}
	const wnohang = 1
	if options&wnohang != 0 {
		select {
		case <-proc.done:
		default:

			return -_wasiEAGAIN
		}
	} else {
		<-proc.done
	}
	w.mu.Lock()
	delete(w.procs, pid)
	w.mu.Unlock()
	binary.LittleEndian.PutUint32(out, uint32(proc.status))
	return _wasiESUCCESS
}

// encodeWaitStatus turns a Go ProcessState into a POSIX wait status int (the
// raw value os.waitstatus_to_exitcode decodes): a normal exit N becomes
// (N&0xff)<<8 (WIFEXITED); a signal becomes the low-7-bits signal number
// (WIFSIGNALED).
func encodeWaitStatus(st *os.ProcessState) int32 {
	if st == nil {
		return 0
	}
	if ws, ok := st.Sys().(syscall.WaitStatus); ok {
		if ws.Signaled() {
			return int32(ws.Signal()) & 0x7f
		}
		return int32(ws.ExitStatus()&0xff) << 8
	}
	code := st.ExitCode()
	if code < 0 {
		code = 127
	}
	return int32(code&0xff) << 8
}

// mapExecError maps an exec.Command Start() failure to a wasi errno.
func mapExecError(err error) int32 {
	switch {
	case errors.Is(err, exec.ErrNotFound), errors.Is(err, fs.ErrNotExist):
		return _wasiENOENT
	case errors.Is(err, fs.ErrPermission):
		return _wasiEACCES
	default:
		return _wasiENOENT
	}
}

// Sock_getaddrinfo is a NON-STANDARD host import (module wasi_snapshot_preview1,
// name "sock_getaddrinfo") backing the bridge's getaddrinfo(). It reads the
// host string at (nodePtr,nodeLen), consults the resolve whitelist, resolves it
// to an IPv4 address via Go's resolver (numeric IPs pass through), and writes
// the 4-byte network-order address at outPtr. Returns 0 on success or a
// negative POSIX-ish errno (the bridge maps it to an EAI_* code).
func (w *WasiStubs) Sock_getaddrinfo(m *Module, nodePtr, nodeLen, outPtr int32) int32 {
	host := ""
	if nodeLen > 0 {
		b := w.memSlice(m, nodePtr, nodeLen)
		if b == nil {
			return -_wasiEFAULT
		}
		host = string(b)
	}
	out := w.memSlice(m, outPtr, 4)
	if out == nil {
		return -_wasiEFAULT
	}
	if host == "" {
		binary.LittleEndian.PutUint32(out, 0)
		return _wasiESUCCESS
	}
	w.mu.Lock()
	hook := w.resolveHook
	w.mu.Unlock()
	if hook != nil && !hook(host) {
		return -_wasiEACCES
	}

	if ip := net.ParseIP(host); ip != nil {
		if v4 := ip.To4(); v4 != nil {
			out[0], out[1], out[2], out[3] = v4[0], v4[1], v4[2], v4[3]
			w.recordResolvedHost(v4, host)
			return _wasiESUCCESS
		}
		return -_wasiEAFNOSUPPORT
	}
	ips, err := net.DefaultResolver.LookupIP(context.Background(), "ip4", host)
	if err != nil || len(ips) == 0 {
		return -_wasiENOENT
	}
	v4 := ips[0].To4()
	if v4 == nil {
		return -_wasiEAFNOSUPPORT
	}
	out[0], out[1], out[2], out[3] = v4[0], v4[1], v4[2], v4[3]
	w.recordResolvedHost(v4, host)
	return _wasiESUCCESS
}

// recordResolvedHost remembers that host resolved to v4, so a later Sock_connect
// to that IP can hand the dial hook the host name it was looked up under.
func (w *WasiStubs) recordResolvedHost(v4 net.IP, host string) {
	ip := fmt.Sprintf("%d.%d.%d.%d", v4[0], v4[1], v4[2], v4[3])
	w.mu.Lock()
	if w.resolvedHosts == nil {
		w.resolvedHosts = make(map[string]string)
	}
	w.resolvedHosts[ip] = host
	w.mu.Unlock()
}

// SetEnv overrides the environment the guest sees via environ_get /
// environ_sizes_get. By default DefaultWASI leaks the host process
// os.Environ(); a sandboxed embedding should call SetEnv with an
// explicit (possibly empty) slice of "KEY=VALUE" strings.
func (w *WasiStubs) SetEnv(env []string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.env = append([]string(nil), env...)
}

// SetArgs overrides os.Args as seen by the guest (argv). Mirrors SetEnv.
func (w *WasiStubs) SetArgs(args []string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.args = append([]string(nil), args...)
}

// SetStdin redirects guest fd 0 to r. A nil r leaves the current source.
// Use this to feed input() / sys.stdin from an in-process io.Reader instead
// of the host process stdin.
func (w *WasiStubs) SetStdin(r io.Reader) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if r != nil {
		w.stdin = r
	}
}

// SetStdout redirects guest fd 1 to wr. A nil wr leaves the current sink.
func (w *WasiStubs) SetStdout(wr io.Writer) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if wr != nil {
		w.stdout = wr
	}
}

// SetStderr redirects guest fd 2 to wr. A nil wr leaves the current sink.
func (w *WasiStubs) SetStderr(wr io.Writer) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if wr != nil {
		w.stderr = wr
	}
}

// checkFS consults the FS policy hook (if any). Returns true when the
// access is permitted. Callers must NOT hold w.mu.
func (w *WasiStubs) checkFS(path string, write bool) bool {
	w.mu.Lock()
	hook := w.fsHook
	w.mu.Unlock()
	if hook == nil {
		return true
	}
	return hook(path, write)
}

// checkNet consults the network policy hook (if any). Returns true when
// the operation is permitted. Callers must NOT hold w.mu.
func (w *WasiStubs) checkNet(op string) bool {
	w.mu.Lock()
	hook := w.netHook
	w.mu.Unlock()
	if hook == nil {
		return true
	}
	return hook(op)
}

// memSlice returns m.memory[off : off+n]. Callers must hold any locks
// they need on the wasm side; WasiStubs.mu is independent. Returns an
// empty slice on out-of-range (the wasi function should then return
// EFAULT / EINVAL).
func (w *WasiStubs) memSlice(m *Module, off, n int32) []byte {
	mem := m.Memory
	lo := uint64(uint32(off))
	hi := lo + uint64(uint32(n))
	if hi > uint64(len(mem)) {
		return nil
	}
	return mem[lo:hi]
}

// errno values used below (subset; see wasi-libc errno.h).
const (
	_wasiESUCCESS     int32 = 0
	_wasiE2BIG        int32 = 1
	_wasiEACCES       int32 = 2
	_wasiEAFNOSUPPORT int32 = 5
	_wasiEAGAIN       int32 = 6
	_wasiEBADF        int32 = 8
	_wasiECHILD       int32 = 12
	_wasiECONNREFUSED int32 = 14
	_wasiEISCONN      int32 = 33
	_wasiEBUSY        int32 = 10
	_wasiEEXIST       int32 = 20
	_wasiEFAULT       int32 = 21
	_wasiEINVAL       int32 = 28
	_wasiEIO          int32 = 29
	_wasiEISDIR       int32 = 31
	_wasiENOENT       int32 = 44
	_wasiENOTDIR      int32 = 54
	_wasiENOTSOCK     int32 = 57
	_wasiENOTSUP      int32 = 58
	_wasiENOSYS       int32 = 52
	_wasiEPERM        int32 = 63
	_wasiEPIPE        int32 = 64
)

// mapOSError turns an os/filesystem error into a wasi errno. Used by the
// path-based syscalls so any os.PathError surfaces as the appropriate
// guest-visible code instead of a coarse EIO.
func mapOSError(err error) int32 {
	if err == nil {
		return _wasiESUCCESS
	}
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return _wasiENOENT
	case errors.Is(err, fs.ErrExist):
		return _wasiEEXIST
	case errors.Is(err, fs.ErrPermission):
		return _wasiEACCES
	case errors.Is(err, syscall.ENOTDIR):
		return _wasiENOTDIR
	case errors.Is(err, syscall.EISDIR):
		return _wasiEISDIR
	case errors.Is(err, syscall.EINVAL):
		return _wasiEINVAL
	case errors.Is(err, syscall.EBADF):
		return _wasiEBADF
	case errors.Is(err, syscall.EAGAIN):
		return _wasiEAGAIN
	case errors.Is(err, syscall.EPIPE):
		return _wasiEPIPE
	}
	return _wasiEIO
}

// totalBytes sums len(s)+1 over s in a uint64 and reports whether the
// total fits in an int32 (i.e. is representable as a wasm-side i32
// length). Callers route the result through memSlice and an OOB on a
// pathologically long arg list surfaces as EFAULT to the guest rather
// than a host-side panic via a wrapped-int32 length.
func totalBytesPlusNul(ss []string) (int32, bool) {
	var total uint64
	for _, s := range ss {
		total += uint64(len(s)) + 1
		if total > 0x7fffffff {
			return 0, false
		}
	}
	return int32(total), true
}

func (w *WasiStubs) Args_get(m *Module, argv, argvBuf int32) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()

	argvBytes64 := uint64(len(w.args)) * 4
	if argvBytes64 > 0x7fffffff {
		return _wasiEFAULT
	}
	argvSlice := w.memSlice(m, argv, int32(argvBytes64))
	if argvSlice == nil {
		return _wasiEFAULT
	}
	total, ok := totalBytesPlusNul(w.args)
	if !ok {
		return _wasiEFAULT
	}
	argvBufSlice := w.memSlice(m, argvBuf, total)
	if argvBufSlice == nil {
		return _wasiEFAULT
	}
	bufOff := uint32(0)
	for i, a := range w.args {
		binary.LittleEndian.PutUint32(argvSlice[i*4:], uint32(argvBuf)+bufOff)
		n := copy(argvBufSlice[bufOff:], a)
		if n < len(a) {
			return _wasiEFAULT
		}
		bufOff += uint32(n)
		argvBufSlice[bufOff] = 0
		bufOff++
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Args_sizes_get(m *Module, argcPtr, argvBufLenPtr int32) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	argcSlice := w.memSlice(m, argcPtr, 4)
	bufLenSlice := w.memSlice(m, argvBufLenPtr, 4)
	if argcSlice == nil || bufLenSlice == nil {
		return _wasiEFAULT
	}
	total, ok := totalBytesPlusNul(w.args)
	if !ok {
		return _wasiEFAULT
	}
	binary.LittleEndian.PutUint32(argcSlice, uint32(len(w.args)))
	binary.LittleEndian.PutUint32(bufLenSlice, uint32(total))
	return _wasiESUCCESS
}

func (w *WasiStubs) Environ_get(m *Module, envv, envBuf int32) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	envvBytes64 := uint64(len(w.env)) * 4
	if envvBytes64 > 0x7fffffff {
		return _wasiEFAULT
	}
	envvSlice := w.memSlice(m, envv, int32(envvBytes64))
	if envvSlice == nil {
		return _wasiEFAULT
	}
	total, ok := totalBytesPlusNul(w.env)
	if !ok {
		return _wasiEFAULT
	}
	envBufSlice := w.memSlice(m, envBuf, total)
	if envBufSlice == nil {
		return _wasiEFAULT
	}
	bufOff := uint32(0)
	for i, e := range w.env {
		binary.LittleEndian.PutUint32(envvSlice[i*4:], uint32(envBuf)+bufOff)
		n := copy(envBufSlice[bufOff:], e)
		if n < len(e) {
			return _wasiEFAULT
		}
		bufOff += uint32(n)
		envBufSlice[bufOff] = 0
		bufOff++
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Environ_sizes_get(m *Module, envcPtr, envBufLenPtr int32) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	envcSlice := w.memSlice(m, envcPtr, 4)
	bufLenSlice := w.memSlice(m, envBufLenPtr, 4)
	if envcSlice == nil || bufLenSlice == nil {
		return _wasiEFAULT
	}
	total, ok := totalBytesPlusNul(w.env)
	if !ok {
		return _wasiEFAULT
	}
	binary.LittleEndian.PutUint32(envcSlice, uint32(len(w.env)))
	binary.LittleEndian.PutUint32(bufLenSlice, uint32(total))
	return _wasiESUCCESS
}

func (w *WasiStubs) Clock_res_get(m *Module, clockID int32, resPtr int32) int32 {

	out := w.memSlice(m, resPtr, 8)
	if out == nil {
		return _wasiEFAULT
	}
	binary.LittleEndian.PutUint64(out, 1)
	return _wasiESUCCESS
}

func (w *WasiStubs) Clock_time_get(m *Module, clockID int32, precision int64, timePtr int32) int32 {
	out := w.memSlice(m, timePtr, 8)
	if out == nil {
		return _wasiEFAULT
	}
	nanos, errno := w.clockNanos(clockID)
	if errno != _wasiESUCCESS {
		return errno
	}
	binary.LittleEndian.PutUint64(out, nanos)
	return _wasiESUCCESS
}

// clockNanos is the layout-independent body of clock_time_get, shared
// by the wasm32 and wasm64 bindings.
func (w *WasiStubs) clockNanos(clockID int32) (uint64, int32) {
	switch clockID {
	case 0:
		return uint64(time.Now().UnixNano()), _wasiESUCCESS
	case 1:
		w.mu.Lock()
		nanos := uint64(time.Since(w.monoStart).Nanoseconds())
		w.mu.Unlock()
		return nanos, _wasiESUCCESS
	default:
		return 0, _wasiEINVAL
	}
}

// closeWasiOpen releases every underlying handle held by op and
// joins any Close errors so callers can map them to a wasi errno
// instead of silently dropping the failure.
func closeWasiOpen(op *wasiOpen) error {
	if op.refs > 0 {

		op.refs--
		return nil
	}
	var err error
	if op.f != nil {
		err = errors.Join(err, op.f.Close())
	}
	if op.conn != nil {
		err = errors.Join(err, op.conn.Close())
	}
	if op.listener != nil {
		err = errors.Join(err, op.listener.Close())
	}
	return err
}

func (w *WasiStubs) Fd_close(m *Module, fd int32) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	op := w.fdTable[fd]
	if op == nil {
		return _wasiEBADF
	}
	closeErr := closeWasiOpen(op)
	delete(w.fdTable, fd)
	if closeErr != nil {
		return mapOSError(closeErr)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_fdstat_get(m *Module, fd, ptr int32) int32 {

	out := w.memSlice(m, ptr, 24)
	if out == nil {
		return _wasiEFAULT
	}
	return w.fdstatFill(fd, out)
}

// fdstatFill writes the 24-byte fdstat for fd into out — the shared
// body of the 32- and 64-bit Fd_fdstat_get bindings (the struct holds
// no pointers, so the layout is width-independent).
func (w *WasiStubs) fdstatFill(fd int32, out []byte) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	var ftype byte = 4 // regular file
	var fdflags uint16
	if fd >= 0 && fd <= 2 {
		ftype = 2
	} else if op := w.fdTable[fd]; op != nil {
		if op.isDir {
			ftype = 3
		} else if op.conn != nil {
			ftype = 6
		} else if op.listener != nil {
			ftype = 6
		}
		fdflags = uint16(op.fdflags)
	} else if fd == 3 {
		ftype = 3
	} else if fd >= 4 {
		return _wasiEBADF
	}
	out[0] = ftype
	out[1] = 0
	binary.LittleEndian.PutUint16(out[2:], fdflags)

	binary.LittleEndian.PutUint64(out[8:], ^uint64(0))
	binary.LittleEndian.PutUint64(out[16:], ^uint64(0))
	return _wasiESUCCESS
}

// Fd_fdstat_set_flags maps WASI fdflags to OS file-status flags via the
// per-platform Fcntl wrapper. The flags are also cached on the wasiOpen
// so a subsequent Fd_fdstat_get reflects what the guest set. Stdio fds
// store the requested flags but otherwise no-op; sockets/listeners take
// only the cache update because Go's net layer manages blocking mode
// internally.
func (w *WasiStubs) Fd_fdstat_set_flags(m *Module, fd, flags int32) int32 {
	w.mu.Lock()
	op := w.fdTable[fd]
	if op == nil && fd > 2 {
		w.mu.Unlock()
		return _wasiEBADF
	}
	if op != nil {
		op.fdflags = flags
	}
	w.mu.Unlock()

	_ = op
	_ = flags
	return _wasiESUCCESS
}

// Fd_fdstat_set_rights stores the requested rights on the wasiOpen but
// does not enforce them — the host process is the trust boundary. WASI
// programs that succeed with maximal rights (per Fd_fdstat_get) get the
// same ESUCCESS here.
func (w *WasiStubs) Fd_fdstat_set_rights(m *Module, fd int32, rightsBase, rightsInherit int64) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	if fd >= 0 && fd <= 2 {
		return _wasiESUCCESS
	}
	if w.fdTable[fd] == nil {
		return _wasiEBADF
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_filestat_get(m *Module, fd, ptr int32) int32 {

	out := w.memSlice(m, ptr, 64)
	if out == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {

		for i := range out {
			out[i] = 0
		}

		switch fd {
		case 0, 1, 2:
			out[16] = 2
		case 3:
			out[16] = 3
		}
		return _wasiESUCCESS
	}
	st, err := op.f.Stat()
	if err != nil {
		return mapOSError(err)
	}
	writeFilestat(out, st)
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_filestat_set_size(m *Module, fd int32, size int64) int32 {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return _wasiEBADF
	}
	if err := op.f.Truncate(size); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_filestat_set_times(m *Module, fd int32, atim, mtim int64, fstFlags int32) int32 {

	w.mu.Lock()
	op := w.fdTable[fd]
	fsys := w.fsys
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return _wasiEBADF
	}
	atime, mtime, err := resolveFiletimes(uint64(atim), uint64(mtim), fstFlags, op.f)
	if err != nil {
		return mapOSError(err)
	}

	if cf, ok := fsys.(chtimesFS); ok {
		if err := cf.Chtimes(op.path, atime, mtime); err != nil {
			return mapOSError(err)
		}
	}
	return _wasiESUCCESS
}

// combine64 reconstructs an unsigned 64-bit time value from a pair of
// 32-bit args. WASI signature uses two i32s for the nanosecond timestamp
// in fd_filestat_set_times.
func combine64(hi, lo int32) uint64 {
	return (uint64(uint32(hi)) << 32) | uint64(uint32(lo))
}

// resolveFiletimes decides the (atime, mtime) pair to apply given a
// fstFlags bitmask. Bits 0x2 (ATIME_NOW) and 0x8 (MTIME_NOW) override the
// explicit values with time.Now(). Unset ATIME/MTIME bits keep the
// existing on-disk time, so f.Stat must succeed when those bits are
// unset; the error is returned so the caller can surface it as a wasi
// errno rather than silently writing epoch.
func resolveFiletimes(atimNs, mtimNs uint64, fstFlags int32, f File) (time.Time, time.Time, error) {
	now := time.Now()
	var atime, mtime time.Time

	needCurrent := fstFlags&(0x1|0x2) == 0 || fstFlags&(0x4|0x8) == 0
	if needCurrent {
		st, err := f.Stat()
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		atime = st.ModTime()
		mtime = st.ModTime()
	}
	if fstFlags&0x1 != 0 {
		atime = time.Unix(0, int64(atimNs))
	}
	if fstFlags&0x2 != 0 {
		atime = now
	}
	if fstFlags&0x4 != 0 {
		mtime = time.Unix(0, int64(mtimNs))
	}
	if fstFlags&0x8 != 0 {
		mtime = now
	}
	return atime, mtime, nil
}

func (w *WasiStubs) Fd_prestat_get(m *Module, fd, ptr int32) int32 {

	if fd != 3 {
		return _wasiEBADF
	}

	out := w.memSlice(m, ptr, 8)
	if out == nil {
		return _wasiEFAULT
	}
	out[0] = 0
	binary.LittleEndian.PutUint32(out[4:], 1)
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_prestat_dir_name(m *Module, fd, buf, buflen int32) int32 {
	if fd != 3 {
		return _wasiEBADF
	}
	if buflen < 1 {
		return _wasiESUCCESS
	}
	out := w.memSlice(m, buf, buflen)
	if out == nil {
		return _wasiEFAULT
	}
	out[0] = '/'
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_read(m *Module, fd, iovs, iovsLen, nreadPtr int32) int32 {
	w.mu.Lock()
	src, op := w.fdSrcLocked(fd)
	w.mu.Unlock()
	if src == nil {
		return _wasiEBADF
	}
	bufs, ok := w.iovecSlices(m, iovs, iovsLen)
	nreadSlice := w.memSlice(m, nreadPtr, 4)
	if !ok || nreadSlice == nil {
		return _wasiEFAULT
	}
	_ = op
	binary.LittleEndian.PutUint32(nreadSlice, uint32(readVec(src, bufs)))
	return _wasiESUCCESS
}

// iovecSlices resolves a wasm32 ciovec/iovec array ({u32 ptr, u32 len}
// entries at iovs) into the backing memory windows. Every entry is
// validated before any I/O happens, so a bad iovec faults the whole
// call instead of after a partial transfer.
func (w *WasiStubs) iovecSlices(m *Module, iovs, iovsLen int32) ([][]byte, bool) {

	iovBytes := uint64(uint32(iovsLen)) * 8
	if iovBytes > 0x7fffffff {
		return nil, false
	}
	iovecs := w.memSlice(m, iovs, int32(iovBytes))
	if iovecs == nil {
		return nil, false
	}
	bufs := make([][]byte, 0, iovsLen)
	for i := int32(0); i < iovsLen; i++ {
		bufPtr := binary.LittleEndian.Uint32(iovecs[i*8:])
		bufLen := binary.LittleEndian.Uint32(iovecs[i*8+4:])
		buf := w.memSlice(m, int32(bufPtr), int32(bufLen))
		if buf == nil {
			return nil, false
		}
		bufs = append(bufs, buf)
	}
	return bufs, true
}

// readVec fills bufs from src in order, stopping at the first error
// (EOF included) or short read; returns the bytes read. Shared by the
// wasm32 and wasm64 fd_read bindings — only the iovec layout differs.
func readVec(src io.Reader, bufs [][]byte) uint64 {
	var total uint64
	for _, buf := range bufs {
		n, err := src.Read(buf)
		total += uint64(n)
		if err != nil || n < len(buf) {
			break
		}
	}
	return total
}

// writeVec drains bufs into dst in order, stopping at the first failed
// write; returns the bytes written. Shared like readVec.
func writeVec(dst io.Writer, bufs [][]byte) uint64 {
	var total uint64
	for _, buf := range bufs {
		n, err := dst.Write(buf)
		total += uint64(n)
		if err != nil {
			break
		}
	}
	return total
}

// fdSrcLocked returns the io.Reader for fd and (when applicable) the
// wasiOpen it came from, or nil if fd is invalid. Caller must hold w.mu.
func (w *WasiStubs) fdSrcLocked(fd int32) (io.Reader, *wasiOpen) {

	op := w.fdTable[fd]
	if op == nil {
		if fd == 0 {
			return w.stdin, nil
		}
		return nil, nil
	}
	if op.stdio == 1 {
		return w.stdin, op
	}
	if op.f != nil {
		return op.f, op
	}
	if op.conn != nil {
		return op.conn, op
	}
	return nil, op
}

// fdDstLocked returns the io.Writer for fd or nil if fd is invalid.
// Caller must hold w.mu.
func (w *WasiStubs) fdDstLocked(fd int32) (io.Writer, *wasiOpen) {
	op := w.fdTable[fd]
	if op == nil {
		switch fd {
		case 1:
			return w.stdout, nil
		case 2:
			return w.stderr, nil
		}
		return nil, nil
	}
	switch op.stdio {
	case 2:
		return w.stdout, op
	case 3:
		return w.stderr, op
	}
	if op.f != nil {
		return op.f, op
	}
	if op.conn != nil {
		return op.conn, op
	}
	return nil, op
}

func (w *WasiStubs) Fd_pread(m *Module, fd, iovs, iovsLen int32, offset int64, nreadPtr int32) int32 {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return _wasiEBADF
	}
	iovBytes := uint64(uint32(iovsLen)) * 8
	if iovBytes > 0x7fffffff {
		return _wasiEFAULT
	}
	iovecs := w.memSlice(m, iovs, int32(iovBytes))
	nreadSlice := w.memSlice(m, nreadPtr, 4)
	if iovecs == nil || nreadSlice == nil {
		return _wasiEFAULT
	}
	var total uint32
	curOff := offset
	for i := int32(0); i < iovsLen; i++ {
		bufPtr := binary.LittleEndian.Uint32(iovecs[i*8:])
		bufLen := binary.LittleEndian.Uint32(iovecs[i*8+4:])
		buf := w.memSlice(m, int32(bufPtr), int32(bufLen))
		if buf == nil {
			return _wasiEFAULT
		}
		n, err := op.f.ReadAt(buf, curOff)
		total += uint32(n)
		curOff += int64(n)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			break
		}
		if n < int(bufLen) {
			break
		}
	}
	binary.LittleEndian.PutUint32(nreadSlice, total)
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_pwrite(m *Module, fd, iovs, iovsLen int32, offset int64, nwrittenPtr int32) int32 {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return _wasiEBADF
	}
	iovBytes := uint64(uint32(iovsLen)) * 8
	if iovBytes > 0x7fffffff {
		return _wasiEFAULT
	}
	iovecs := w.memSlice(m, iovs, int32(iovBytes))
	nwSlice := w.memSlice(m, nwrittenPtr, 4)
	if iovecs == nil || nwSlice == nil {
		return _wasiEFAULT
	}
	var total uint32
	curOff := offset
	for i := int32(0); i < iovsLen; i++ {
		bufPtr := binary.LittleEndian.Uint32(iovecs[i*8:])
		bufLen := binary.LittleEndian.Uint32(iovecs[i*8+4:])
		buf := w.memSlice(m, int32(bufPtr), int32(bufLen))
		if buf == nil {
			return _wasiEFAULT
		}
		n, err := op.f.WriteAt(buf, curOff)
		total += uint32(n)
		curOff += int64(n)
		if err != nil {
			break
		}
	}
	binary.LittleEndian.PutUint32(nwSlice, total)
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_seek(m *Module, fd int32, offset int64, whence, newOffPtr int32) int32 {
	out := w.memSlice(m, newOffPtr, 8)
	if out == nil {
		return _wasiEFAULT
	}
	n, errno := w.fdSeek(fd, offset, int(whence))
	if errno != _wasiESUCCESS {
		return errno
	}
	binary.LittleEndian.PutUint64(out, uint64(n))
	return _wasiESUCCESS
}

// fdSeek is the layout-independent body of fd_seek, shared by the
// wasm32 and wasm64 bindings.
func (w *WasiStubs) fdSeek(fd int32, offset int64, whence int) (int64, int32) {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return 0, _wasiEBADF
	}
	n, err := op.f.Seek(offset, whence)
	if err != nil {
		return 0, _wasiEINVAL
	}
	return n, _wasiESUCCESS
}

func (w *WasiStubs) Fd_tell(m *Module, fd, offsetPtr int32) int32 {
	out := w.memSlice(m, offsetPtr, 8)
	if out == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return _wasiEBADF
	}
	n, err := op.f.Seek(0, 1)
	if err != nil {
		return _wasiEIO
	}
	binary.LittleEndian.PutUint64(out, uint64(n))
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_write(m *Module, fd, iovs, iovsLen, nwrittenPtr int32) int32 {
	w.mu.Lock()
	dst, _ := w.fdDstLocked(fd)
	w.mu.Unlock()
	bufs, ok := w.iovecSlices(m, iovs, iovsLen)
	nwrittenSlice := w.memSlice(m, nwrittenPtr, 4)
	if !ok || nwrittenSlice == nil {
		return _wasiEFAULT
	}
	if dst == nil {
		binary.LittleEndian.PutUint32(nwrittenSlice, 0)
		return _wasiEBADF
	}
	binary.LittleEndian.PutUint32(nwrittenSlice, uint32(writeVec(dst, bufs)))
	return _wasiESUCCESS
}
func (w *WasiStubs) Fd_sync(m *Module, fd int32) int32 {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return _wasiEBADF
	}
	if err := op.f.Sync(); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_datasync(m *Module, fd int32) int32 {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return _wasiEBADF
	}

	if err := op.f.Sync(); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_advise(m *Module, fd int32, offset, length int64, advice int32) int32 {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return _wasiEBADF
	}

	_, _, _ = offset, length, advice
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_allocate(m *Module, fd int32, offset, length int64) int32 {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil {
		return _wasiEBADF
	}

	if err := op.f.Truncate(offset + length); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

// Path_chmod is a NON-STANDARD host import (module wasi_snapshot_preview1,
// name "path_chmod") backing a bridge-provided chmod(): WASI preview1 has
// no way to change file modes. The path at (pathPtr,pathLen) is
// preopen-relative, like path_open's. Backends without chmod support
// (MemFS keeps no modes) report ENOSYS. Returns 0 or a negative errno.
func (w *WasiStubs) Path_chmod(m *Module, pathPtr, pathLen, mode int32) int32 {
	pathSlice := w.memSlice(m, pathPtr, pathLen)
	if pathSlice == nil {
		return -_wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	ch, ok := fsys.(interface {
		Chmod(string, os.FileMode) error
	})
	if !ok {
		return -_wasiENOSYS
	}
	if err := ch.Chmod(string(pathSlice), os.FileMode(uint32(mode)&0o7777)); err != nil {
		return -mapOSError(err)
	}
	return _wasiESUCCESS
}

// Path_filestat_mode is a NON-STANDARD host import (module
// wasi_snapshot_preview1, name "path_filestat_mode") backing a
// bridge-provided stat/lstat: WASI's filestat carries no permission
// bits, so the bridge merges the real mode in from here. The path is
// preopen-relative; follow selects stat vs lstat semantics. Writes the
// unix permission bits at modeOutPtr; returns 0 or a negative errno.
func (w *WasiStubs) Path_filestat_mode(m *Module, pathPtr, pathLen, follow, modeOutPtr int32) int32 {
	pathSlice := w.memSlice(m, pathPtr, pathLen)
	out := w.memSlice(m, modeOutPtr, 4)
	if pathSlice == nil || out == nil {
		return -_wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	var fi os.FileInfo
	var err error
	if follow != 0 {
		fi, err = fsys.Stat(string(pathSlice))
	} else {
		fi, err = fsys.Lstat(string(pathSlice))
	}
	if err != nil {
		return -mapOSError(err)
	}
	mode := fi.Mode()
	bits := uint32(mode.Perm())
	if mode&os.ModeSetuid != 0 {
		bits |= 0o4000
	}
	if mode&os.ModeSetgid != 0 {
		bits |= 0o2000
	}
	if mode&os.ModeSticky != 0 {
		bits |= 0o1000
	}
	binary.LittleEndian.PutUint32(out, bits)
	return _wasiESUCCESS
}

// dupSourceLocked resolves the entry a dup of fd should share: the
// existing table entry, or a fresh alias for a bare interpreter stdio fd.
// Caller holds w.mu.
func (w *WasiStubs) dupSourceLocked(fd int32) *wasiOpen {
	if op := w.fdTable[fd]; op != nil {
		return op
	}
	if fd >= 0 && fd <= 2 {
		op := &wasiOpen{stdio: int8(fd + 1)}
		w.fdTable[fd] = op
		return op
	}
	return nil
}

// Fd_dup is a NON-STANDARD host import (module wasi_snapshot_preview1,
// name "fd_dup") backing the bridge's dup(): the new fd shares the same
// open descriptor (offset included), and the underlying file closes only
// when the last sharing fd does. Writes the new fd at outPtr.
func (w *WasiStubs) Fd_dup(m *Module, fd, outPtr int32) int32 {
	out := w.memSlice(m, outPtr, 4)
	if out == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	op := w.dupSourceLocked(fd)
	if op == nil {
		return _wasiEBADF
	}
	nfd := w.nextFD
	w.nextFD++
	w.fdTable[nfd] = op
	op.refs++
	binary.LittleEndian.PutUint32(out, uint32(nfd))
	return _wasiESUCCESS
}

// Fd_dup2 is a NON-STANDARD host import (module wasi_snapshot_preview1,
// name "fd_dup2") backing the bridge's dup2(): to becomes another
// reference to from's descriptor, closing whatever to previously held.
func (w *WasiStubs) Fd_dup2(m *Module, from, to int32) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	src := w.dupSourceLocked(from)
	if src == nil {
		return _wasiEBADF
	}
	if from == to {
		return _wasiESUCCESS
	}
	var closeErr error
	if dst := w.fdTable[to]; dst != nil {
		if dst == src {
			return _wasiESUCCESS
		}
		closeErr = closeWasiOpen(dst)
	}
	w.fdTable[to] = src
	src.refs++
	if closeErr != nil {
		return mapOSError(closeErr)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_renumber(m *Module, from, to int32) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	if from == to {

		if _, ok := w.fdTable[from]; ok {
			return _wasiESUCCESS
		}
		return _wasiEBADF
	}
	src, ok := w.fdTable[from]
	if !ok {
		return _wasiEBADF
	}
	var closeErr error
	if dst, ok2 := w.fdTable[to]; ok2 {
		closeErr = closeWasiOpen(dst)
	}
	w.fdTable[to] = src
	delete(w.fdTable, from)
	if closeErr != nil {
		return mapOSError(closeErr)
	}
	return _wasiESUCCESS
}

// readDirCached lazily caches the directory listing on first
// Fd_readdir, so paged reads (cookie-driven) walk the same snapshot.
func (op *wasiOpen) readDirCached() ([]os.DirEntry, error) {
	if op.dirCache != nil {
		return op.dirCache, nil
	}
	if op.f == nil {
		return nil, syscall.EBADF
	}
	if _, err := op.f.Seek(0, 0); err != nil {
		return nil, err
	}
	entries, err := op.f.ReadDir(-1)
	if err != nil {
		return nil, err
	}

	out := make([]os.DirEntry, 0, len(entries)+2)
	out = append(out, dotEntry(op.path, "."), dotEntry(op.path, ".."))
	out = append(out, entries...)

	sort.SliceStable(out[2:], func(i, j int) bool {
		return out[2+i].Name() < out[2+j].Name()
	})
	op.dirCache = out
	return out, nil
}

// dotEntry produces a synthetic os.DirEntry for "." and "..". Its
// Info() returns the stat of the parent directory (good enough for
// guest-side d_type detection).
func dotEntry(parent, name string) os.DirEntry {
	return &dotDirEntry{name: name, parent: parent}
}

type dotDirEntry struct {
	name, parent string
}

func (d *dotDirEntry) Name() string { return d.name }
func (d *dotDirEntry) IsDir() bool  { return true }
func (d *dotDirEntry) Type() os.FileMode {
	return os.ModeDir
}
func (d *dotDirEntry) Info() (os.FileInfo, error) {
	if d.name == "." {
		return os.Stat(d.parent)
	}
	return os.Stat(filepath.Dir(d.parent))
}

func (w *WasiStubs) Fd_readdir(m *Module, fd, buf, buflen int32, cookie int64, bufusedPtr int32) int32 {
	bufSlice := w.memSlice(m, buf, buflen)
	bufusedSlice := w.memSlice(m, bufusedPtr, 4)
	if bufSlice == nil || bufusedSlice == nil {
		return _wasiEFAULT
	}
	written, errno := w.fdReaddir(fd, bufSlice, cookie)
	if errno != _wasiESUCCESS {
		return errno
	}
	binary.LittleEndian.PutUint32(bufusedSlice, uint32(written))
	return _wasiESUCCESS
}

// fdReaddir is the layout-independent body of fd_readdir: it packs
// dirents into bufSlice starting at the cookie'th entry and returns the
// byte count used. The dirent wire format has no pointer-width fields,
// so wasm32 and wasm64 share it; only the bufused out-pointer differs.
func (w *WasiStubs) fdReaddir(fd int32, bufSlice []byte, cookie int64) (int, int32) {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.f == nil || !op.isDir {
		return 0, _wasiEBADF
	}

	if cookie == 0 {
		op.dirCache = nil
	}
	entries, err := op.readDirCached()
	if err != nil {
		return 0, mapOSError(err)
	}
	startIdx := int(cookie)
	if startIdx < 0 {
		startIdx = 0
	}
	written := 0
	for i := startIdx; i < len(entries); i++ {
		e := entries[i]
		nameBytes := []byte(e.Name())
		// dirent header: d_next u64 + d_ino u64 + d_namlen u32 + d_type u8 + 3 pad = 24 bytes.
		const headerLen = 24
		// os.FileInfo does not expose inode portably; report 0.
		var dtype byte = 4 // regular file
		if e.IsDir() {
			dtype = 3
		} else if e.Type()&os.ModeSymlink != 0 {
			dtype = 7
		} else if e.Type()&os.ModeNamedPipe != 0 {
			dtype = 6
		} else if e.Type()&os.ModeSocket != 0 {
			dtype = 6
		}
		// Assemble the fixed header, then copy header+name into the buffer.
		// When a record does not fully fit we copy as much as fits so that
		// bufused == buflen, which is the wasi-libc signal for "more entries
		// available; call again with the last returned cookie". We must NOT
		// zero-fill the leftover: a zeroed dirent (d_namlen=0, d_next=0) is
		// misread by wasi-libc as end-of-directory and silently truncates the
		// listing (e.g. makes a guest's importer miss standard-library packages).
		var hdr [headerLen]byte
		binary.LittleEndian.PutUint64(hdr[0:], uint64(i+1))

		binary.LittleEndian.PutUint64(hdr[8:], uint64(i)+1)
		binary.LittleEndian.PutUint32(hdr[16:], uint32(len(nameBytes)))
		hdr[20] = dtype
		n := copy(bufSlice[written:], hdr[:])
		written += n
		if n < len(hdr) {
			written = len(bufSlice)
			break
		}
		n = copy(bufSlice[written:], nameBytes)
		written += n
		if n < len(nameBytes) {
			written = len(bufSlice)
			break
		}
	}
	return written, _wasiESUCCESS
}

// Path_open opens a wasm-supplied path and registers it in the fd
// table. The path is resolved against the host filesystem with the same
// rights the host Go process has — wasm2go's default WASI is a thin
// passthrough, not a sandbox. The dirFd == 3 special case keeps the
// "preopen /" convention that wasi-libc requires for its directory
// enumeration, but the path itself is opened verbatim (joined to "/")
// using os.OpenFile. Callers that need a sandbox should provide their
// own Wasi_snapshot_preview1Imports implementation via NewWithWASI.
func (w *WasiStubs) Path_open(m *Module, dirFd, dirflags, pathPtr, pathLen, oflags int32, fsRightsBase, fsRightsInherit int64, fdflags, openedFdPtr int32) int32 {
	if dirFd != 3 {
		return _wasiEBADF
	}
	pathSlice := w.memSlice(m, pathPtr, pathLen)
	outSlice := w.memSlice(m, openedFdPtr, 4)
	if pathSlice == nil || outSlice == nil {
		return _wasiEFAULT
	}
	fd, errno := w.pathOpen(string(pathSlice), dirflags, oflags, fsRightsBase, fdflags)
	if errno != _wasiESUCCESS {
		return errno
	}
	binary.LittleEndian.PutUint32(outSlice, uint32(fd))
	return _wasiESUCCESS
}

// pathOpen is the layout-independent body of path_open: it resolves and
// opens rel, registers the fd, and returns it. Callers own reading the
// path and writing the opened fd at their ABI's pointer width.
func (w *WasiStubs) pathOpen(rel string, dirflags, oflags int32, fsRightsBase int64, fdflags int32) (int32, int32) {
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()

	canRead := fsRightsBase&(1<<1) != 0
	canWrite := fsRightsBase&(1<<6) != 0
	var flag int
	switch {
	case canRead && canWrite:
		flag = os.O_RDWR
	case canWrite && !canRead:
		flag = os.O_WRONLY
	default:
		flag = os.O_RDONLY
	}

	if oflags&0x1 != 0 {
		flag |= os.O_CREATE
	}
	if oflags&0x4 != 0 {
		flag |= os.O_EXCL
	}
	if oflags&0x8 != 0 {
		flag |= os.O_TRUNC
	}

	if fdflags&0x1 != 0 {
		flag |= os.O_APPEND
	}
	if fdflags&(0x2|0x8|0x10) != 0 {
		flag |= os.O_SYNC
	}

	writeAccess := flag&(os.O_WRONLY|os.O_RDWR) != 0 || flag&(os.O_CREATE|os.O_TRUNC) != 0
	if !w.checkFS(rel, writeAccess) {
		return -1, _wasiEACCES
	}

	requireDir := oflags&0x2 != 0
	noFollow := dirflags&0x1 == 0

	if requireDir {

		flag = os.O_RDONLY
	}

	if noFollow {
		if li, lerr := fsys.Lstat(rel); lerr == nil && (li.Mode()&os.ModeSymlink) != 0 {
			return -1, _wasiENOENT
		}
	}
	f, err := fsys.OpenFile(rel, flag, 0o644)
	if err != nil {
		return -1, mapOSError(err)
	}
	st, statErr := f.Stat()
	if statErr != nil {
		return -1, mapOSError(errors.Join(statErr, f.Close()))
	}
	isDir := st.IsDir()
	if requireDir && !isDir {
		if cerr := f.Close(); cerr != nil {
			return -1, mapOSError(cerr)
		}
		return -1, _wasiENOTDIR
	}
	w.mu.Lock()
	fd := w.nextFD
	w.nextFD++
	w.fdTable[fd] = &wasiOpen{f: f, isDir: isDir, path: rel, fdflags: fdflags}
	w.mu.Unlock()
	return fd, _wasiESUCCESS
}

func (w *WasiStubs) Path_create_directory(m *Module, dirFd, pathPtr, pathLen int32) int32 {
	if dirFd != 3 {
		return _wasiEBADF
	}
	pathSlice := w.memSlice(m, pathPtr, pathLen)
	if pathSlice == nil {
		return _wasiEFAULT
	}
	if !w.checkFS(string(pathSlice), true) {
		return _wasiEACCES
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	if err := fsys.Mkdir(string(pathSlice), 0o755); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Path_unlink_file(m *Module, dirFd, pathPtr, pathLen int32) int32 {
	if dirFd != 3 {
		return _wasiEBADF
	}
	pathSlice := w.memSlice(m, pathPtr, pathLen)
	if pathSlice == nil {
		return _wasiEFAULT
	}
	if !w.checkFS(string(pathSlice), true) {
		return _wasiEACCES
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	rel := string(pathSlice)
	st, err := fsys.Lstat(rel)
	if err != nil {
		return mapOSError(err)
	}
	if st.IsDir() {
		return _wasiEISDIR
	}
	if err := fsys.Remove(rel); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Path_remove_directory(m *Module, dirFd, pathPtr, pathLen int32) int32 {
	if dirFd != 3 {
		return _wasiEBADF
	}
	pathSlice := w.memSlice(m, pathPtr, pathLen)
	if pathSlice == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	rel := string(pathSlice)
	st, err := fsys.Lstat(rel)
	if err != nil {
		return mapOSError(err)
	}
	if !st.IsDir() {
		return _wasiENOTDIR
	}
	if err := fsys.Remove(rel); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Path_rename(m *Module, oldFd, oldPathPtr, oldPathLen, newFd, newPathPtr, newPathLen int32) int32 {
	if oldFd != 3 || newFd != 3 {
		return _wasiEBADF
	}
	oldSlice := w.memSlice(m, oldPathPtr, oldPathLen)
	newSlice := w.memSlice(m, newPathPtr, newPathLen)
	if oldSlice == nil || newSlice == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	if err := fsys.Rename(string(oldSlice), string(newSlice)); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Path_filestat_get(m *Module, dirFd, flags, pathPtr, pathLen, outPtr int32) int32 {
	if dirFd != 3 {
		return _wasiEBADF
	}
	pathSlice := w.memSlice(m, pathPtr, pathLen)
	out := w.memSlice(m, outPtr, 64)
	if pathSlice == nil || out == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	rel := string(pathSlice)
	var st os.FileInfo
	var err error
	if flags&0x1 != 0 {
		st, err = fsys.Stat(rel)
	} else {
		st, err = fsys.Lstat(rel)
	}
	if err != nil {
		return mapOSError(err)
	}
	writeFilestat(out, st)
	return _wasiESUCCESS
}

func (w *WasiStubs) Path_filestat_set_times(m *Module, dirFd, flags, pathPtr, pathLen int32, atim, mtim int64, fstFlags int32) int32 {
	if dirFd != 3 {
		return _wasiEBADF
	}
	pathSlice := w.memSlice(m, pathPtr, pathLen)
	if pathSlice == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	rel := string(pathSlice)
	follow := flags&0x1 != 0
	now := time.Now()
	var st os.FileInfo
	var statErr error
	if follow {
		st, statErr = fsys.Stat(rel)
	} else {
		st, statErr = fsys.Lstat(rel)
	}
	if statErr != nil {
		return mapOSError(statErr)
	}
	atime := st.ModTime()
	mtime := st.ModTime()
	if fstFlags&0x1 != 0 {
		atime = time.Unix(0, int64(atim))
	}
	if fstFlags&0x2 != 0 {
		atime = now
	}
	if fstFlags&0x4 != 0 {
		mtime = time.Unix(0, int64(mtim))
	}
	if fstFlags&0x8 != 0 {
		mtime = now
	}

	if cf, ok := fsys.(chtimesFS); ok {
		if err := cf.Chtimes(rel, atime, mtime); err != nil {
			return mapOSError(err)
		}
	}
	return _wasiESUCCESS
}

// chtimesFS is an optional FS capability for backends that track timestamps.
type chtimesFS interface {
	Chtimes(name string, atime, mtime time.Time) error
}

func (o osFS) Chtimes(name string, atime, mtime time.Time) error {
	return os.Chtimes(o.join(name), atime, mtime)
}

func (w *WasiStubs) Path_link(m *Module, oldFd, oldFlags, oldPathPtr, oldPathLen, newFd, newPathPtr, newPathLen int32) int32 {
	if oldFd != 3 || newFd != 3 {
		return _wasiEBADF
	}
	oldSlice := w.memSlice(m, oldPathPtr, oldPathLen)
	newSlice := w.memSlice(m, newPathPtr, newPathLen)
	if oldSlice == nil || newSlice == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	if err := fsys.Link(string(oldSlice), string(newSlice)); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Path_symlink(m *Module, targetPtr, targetLen, dirFd, linkPathPtr, linkPathLen int32) int32 {
	if dirFd != 3 {
		return _wasiEBADF
	}
	targetSlice := w.memSlice(m, targetPtr, targetLen)
	linkSlice := w.memSlice(m, linkPathPtr, linkPathLen)
	if targetSlice == nil || linkSlice == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	if err := fsys.Symlink(string(targetSlice), string(linkSlice)); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Path_readlink(m *Module, dirFd, pathPtr, pathLen, buf, buflen, bufusedPtr int32) int32 {
	if dirFd != 3 {
		return _wasiEBADF
	}
	pathSlice := w.memSlice(m, pathPtr, pathLen)
	bufSlice := w.memSlice(m, buf, buflen)
	bufused := w.memSlice(m, bufusedPtr, 4)
	if pathSlice == nil || bufSlice == nil || bufused == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	target, err := fsys.Readlink(string(pathSlice))
	if err != nil {
		return mapOSError(err)
	}
	n := copy(bufSlice, target)
	binary.LittleEndian.PutUint32(bufused, uint32(n))
	return _wasiESUCCESS
}

func (w *WasiStubs) Random_get(m *Module, buf, bufLen int32) int32 {
	slice := w.memSlice(m, buf, bufLen)
	if slice == nil {
		return _wasiEFAULT
	}
	_, err := rand.Read(slice)
	if err != nil {
		return _wasiEIO
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Sched_yield(m *Module) int32 {
	runtime.Gosched()
	return _wasiESUCCESS
}

// Poll_oneoff decodes the WASI subscription_u records and reproduces the
// requested events.
//
// Each subscription is 48 bytes:
//
//	u64 userdata
//	u8  eventtype  (0=clock, 1=fd_read, 2=fd_write)
//	... per-type payload starting at offset 16
//
// For clock subscriptions, payload at offset 16 is: u32 clock_id, u64
// timeout, u64 precision, u16 sub_clock_flags (bit0=ABSTIME). We sleep
// for `timeout` ns (relative timer) or the diff to `timeout` (absolute
// timer). For fd_read / fd_write subscriptions, payload at offset 16 is
// a u32 fd; we call into the platform Poll helper to wait for
// readiness.
//
// Each emitted event is 32 bytes: u64 userdata, u16 errno, u16
// eventtype, u64 fd_readwrite_nbytes (filled for fd events), u16
// flags, then 6 bytes of padding.
func (w *WasiStubs) Poll_oneoff(m *Module, inPtr, outPtr, nsubs, neventsPtr int32) int32 {
	subsTotal := uint64(uint32(nsubs)) * 48
	if subsTotal > 0x7fffffff {
		return _wasiEFAULT
	}
	subs := w.memSlice(m, inPtr, int32(subsTotal))
	evTotal := uint64(uint32(nsubs)) * 32
	if evTotal > 0x7fffffff {
		return _wasiEFAULT
	}
	events := w.memSlice(m, outPtr, int32(evTotal))
	nev := w.memSlice(m, neventsPtr, 4)
	if subs == nil || events == nil || nev == nil {
		return _wasiEFAULT
	}

	type pollItem struct {
		userdata uint64
		etype    byte
		fd       int32
		isRead   bool
	}
	var minClockNs int64 = -1
	var clockEvents []pollItem
	var fdEvents []pollItem
	for i := int32(0); i < nsubs; i++ {
		base := i * 48
		userdata := binary.LittleEndian.Uint64(subs[base:])
		etype := subs[base+8]
		switch etype {
		case 0:
			timeout := int64(binary.LittleEndian.Uint64(subs[base+24:]))
			flags := binary.LittleEndian.Uint16(subs[base+40:])
			ns := timeout
			if flags&0x1 != 0 {

				ns = timeout - time.Now().UnixNano()
				if ns < 0 {
					ns = 0
				}
			}
			if minClockNs < 0 || ns < minClockNs {
				minClockNs = ns
			}
			clockEvents = append(clockEvents, pollItem{userdata: userdata, etype: 0})
		case 1, 2:
			fd := int32(binary.LittleEndian.Uint32(subs[base+16:]))
			fdEvents = append(fdEvents, pollItem{userdata: userdata, etype: etype, fd: fd, isRead: etype == 1})
		default:

			clockEvents = append(clockEvents, pollItem{userdata: userdata, etype: etype})
		}
	}

	if minClockNs > 0 && len(fdEvents) == 0 {
		time.Sleep(time.Duration(minClockNs))
	}

	written := int32(0)
	for _, ev := range clockEvents {
		if ev.etype == 0 && len(fdEvents) > 0 {
			continue
		}
		writeEvent(events[written:written+32], ev.userdata, ev.etype, 0, 0)
		written += 32
	}
	for _, ev := range fdEvents {
		w.mu.Lock()
		op := w.fdTable[ev.fd]
		w.mu.Unlock()
		var errno int32
		var nbytes uint64
		if op == nil {
			errno = _wasiEBADF
		} else if op.f != nil {

			if ev.isRead {
				if st, err := op.f.Stat(); err == nil {

					if cur, err := op.f.Seek(0, 1); err == nil && st.Size() > cur {
						nbytes = uint64(st.Size() - cur)
					}
				}
			}
		} else if op.conn != nil {

			_ = minClockNs
		}
		writeEvent(events[written:written+32], ev.userdata, ev.etype, uint16(errno), nbytes)
		written += 32
	}

	binary.LittleEndian.PutUint32(nev, uint32(written/32))
	return _wasiESUCCESS
}

func writeEvent(dst []byte, userdata uint64, etype byte, errno uint16, nbytes uint64) {
	for i := range dst {
		dst[i] = 0
	}
	binary.LittleEndian.PutUint64(dst[0:], userdata)
	binary.LittleEndian.PutUint16(dst[8:], errno)
	binary.LittleEndian.PutUint16(dst[10:], uint16(etype))
	binary.LittleEndian.PutUint64(dst[16:], nbytes)
}

func (w *WasiStubs) Proc_exit(m *Module, code int32) {

	panic(&WasiExitError{Code: code})
}

func (w *WasiStubs) Proc_raise(m *Module, sig int32) int32 {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		return mapOSError(err)
	}
	if err := p.Signal(syscall.Signal(sig)); err != nil {
		return mapOSError(err)
	}
	return _wasiESUCCESS
}

// Sock_socket is a NON-STANDARD host import (module wasi_snapshot_preview1,
// name "sock_socket") that backs a libc socket() call wrapped via
// -Wl,--wrap=socket in the guest. WASI preview1 has no way to create an
// outbound socket; this gives the guest a host-managed fd whose connection is
// established later by Sock_connect. domain/type follow the POSIX socket()
// args (AF_INET / SOCK_STREAM); only TCP over IPv4 is supported. Returns the
// new fd, or a negative errno on failure.
func (w *WasiStubs) Sock_socket(m *Module, domain, typ int32) int32 {

	_ = domain
	_ = typ
	w.mu.Lock()
	defer w.mu.Unlock()
	fd := w.nextFD
	w.nextFD++
	w.fdTable[fd] = &wasiOpen{isSocket: true}
	return fd
}

// Sock_connect is a NON-STANDARD host import (module wasi_snapshot_preview1,
// name "sock_connect") backing a libc connect() wrapped via
// -Wl,--wrap=connect. ipBE carries the IPv4 address in network byte order
// exactly as it sat in sockaddr_in.sin_addr.s_addr (so the low byte is the
// first octet); port is host byte order. It consults the dial whitelist,
// dials via Go's net, and attaches the resulting conn to the socket fd so the
// existing Sock_send / Sock_recv / Fd_close paths drive it. Returns 0 or a
// negative errno.
func (w *WasiStubs) Sock_connect(m *Module, fd, ipBE, port int32) int32 {
	u := uint32(ipBE)
	ip := fmt.Sprintf("%d.%d.%d.%d", u&0xff, (u>>8)&0xff, (u>>16)&0xff, (u>>24)&0xff)
	w.mu.Lock()
	op := w.fdTable[fd]
	hook := w.dialHook
	host := w.resolvedHosts[ip]
	w.mu.Unlock()
	if op == nil || !op.isSocket {
		return -_wasiENOTSOCK
	}
	if op.conn != nil {
		return -_wasiEISCONN
	}
	p := int(uint16(port))
	if hook != nil && !hook("tcp", host, ip, p) {
		return -_wasiEACCES
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, strconv.Itoa(p)), 30*time.Second)
	if err != nil {
		return -_wasiECONNREFUSED
	}
	w.mu.Lock()

	if cur := w.fdTable[fd]; cur == op {
		op.conn = conn
		w.mu.Unlock()
		return _wasiESUCCESS
	}
	w.mu.Unlock()
	if cerr := conn.Close(); cerr != nil {
		return mapOSError(cerr)
	}
	return -_wasiEBADF
}

// Sock_accept accepts the next incoming TCP/Unix connection on the
// listener associated with fd, registers it as a new wasiOpen with a
// conn arm, and writes the new fd at fdOutPtr. Returns ENOTSOCK if fd
// isn't a listener.
func (w *WasiStubs) Sock_accept(m *Module, fd, flags, fdOutPtr int32) int32 {
	if !w.checkNet("accept") {
		return _wasiEACCES
	}
	out := w.memSlice(m, fdOutPtr, 4)
	if out == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.listener == nil {
		return _wasiENOTSOCK
	}
	conn, err := op.listener.Accept()
	if err != nil {
		return mapOSError(err)
	}
	w.mu.Lock()
	newFD := w.nextFD
	w.nextFD++
	w.fdTable[newFD] = &wasiOpen{conn: conn}
	w.mu.Unlock()
	binary.LittleEndian.PutUint32(out, uint32(newFD))
	return _wasiESUCCESS
}

func (w *WasiStubs) Sock_recv(m *Module, fd, riData, riDataLen, riFlags, roDataLenPtr, roFlagsPtr int32) int32 {
	if !w.checkNet("recv") {
		return _wasiEACCES
	}
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.conn == nil {
		return _wasiENOTSOCK
	}
	iovBytes := uint64(uint32(riDataLen)) * 8
	if iovBytes > 0x7fffffff {
		return _wasiEFAULT
	}
	iovecs := w.memSlice(m, riData, int32(iovBytes))
	lenOut := w.memSlice(m, roDataLenPtr, 4)

	flagsOut := w.memSlice(m, roFlagsPtr, 2)
	if iovecs == nil || lenOut == nil || flagsOut == nil {
		return _wasiEFAULT
	}
	var total uint32
	for i := int32(0); i < riDataLen; i++ {
		bufPtr := binary.LittleEndian.Uint32(iovecs[i*8:])
		bufLen := binary.LittleEndian.Uint32(iovecs[i*8+4:])
		buf := w.memSlice(m, int32(bufPtr), int32(bufLen))
		if buf == nil {
			return _wasiEFAULT
		}
		n, err := op.conn.Read(buf)
		total += uint32(n)
		if err != nil {
			break
		}
	}
	binary.LittleEndian.PutUint16(flagsOut, 0)
	binary.LittleEndian.PutUint32(lenOut, total)
	return _wasiESUCCESS
}

func (w *WasiStubs) Sock_send(m *Module, fd, siData, siDataLen, siFlags, soDataLenPtr int32) int32 {
	if !w.checkNet("send") {
		return _wasiEACCES
	}
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.conn == nil {
		return _wasiENOTSOCK
	}
	iovBytes := uint64(uint32(siDataLen)) * 8
	if iovBytes > 0x7fffffff {
		return _wasiEFAULT
	}
	iovecs := w.memSlice(m, siData, int32(iovBytes))
	lenOut := w.memSlice(m, soDataLenPtr, 4)
	if iovecs == nil || lenOut == nil {
		return _wasiEFAULT
	}
	var total uint32
	for i := int32(0); i < siDataLen; i++ {
		bufPtr := binary.LittleEndian.Uint32(iovecs[i*8:])
		bufLen := binary.LittleEndian.Uint32(iovecs[i*8+4:])
		buf := w.memSlice(m, int32(bufPtr), int32(bufLen))
		if buf == nil {
			return _wasiEFAULT
		}
		n, err := op.conn.Write(buf)
		total += uint32(n)
		if err != nil {
			break
		}
	}
	binary.LittleEndian.PutUint32(lenOut, total)
	return _wasiESUCCESS
}

func (w *WasiStubs) Sock_shutdown(m *Module, fd, how int32) int32 {
	w.mu.Lock()
	op := w.fdTable[fd]
	w.mu.Unlock()
	if op == nil || op.conn == nil {
		return _wasiENOTSOCK
	}
	type shutdowner interface {
		CloseRead() error
		CloseWrite() error
	}
	sh, ok := op.conn.(shutdowner)
	if !ok {

		if err := op.conn.Close(); err != nil {
			return mapOSError(err)
		}
		return _wasiESUCCESS
	}
	var shErr error
	if how&0x1 != 0 {
		shErr = errors.Join(shErr, sh.CloseRead())
	}
	if how&0x2 != 0 {
		shErr = errors.Join(shErr, sh.CloseWrite())
	}
	if shErr != nil {
		return mapOSError(shErr)
	}
	return _wasiESUCCESS
}

// writeFilestat populates the 64-byte WASI filestat structure from a
// host os.FileInfo. The dev/ino fields come from the per-platform
// wasiPlatformStatSys helper (unix returns Stat_t.Dev/.Ino; Windows
// returns zeros).
func writeFilestat(out []byte, st os.FileInfo) {

	binary.LittleEndian.PutUint64(out[0:], 0)
	binary.LittleEndian.PutUint64(out[8:], 0)
	var ftype byte = 4
	mode := st.Mode()
	switch {
	case mode.IsDir():
		ftype = 3
	case mode&os.ModeSymlink != 0:
		ftype = 7
	case mode&os.ModeNamedPipe != 0:
		ftype = 6
	case mode&os.ModeSocket != 0:
		ftype = 6
	case mode&os.ModeDevice != 0:
		ftype = 1
	case mode&os.ModeCharDevice != 0:
		ftype = 2
	}
	out[16] = ftype
	binary.LittleEndian.PutUint64(out[24:], 1)
	binary.LittleEndian.PutUint64(out[32:], uint64(st.Size()))
	nanos := uint64(st.ModTime().UnixNano())
	binary.LittleEndian.PutUint64(out[40:], nanos)
	binary.LittleEndian.PutUint64(out[48:], nanos)
	binary.LittleEndian.PutUint64(out[56:], nanos)
}

// memSlice64 is memSlice for full-range 64-bit guest pointers.
func (w *WasiStubs) memSlice64(m *Module, off int64, n int64) []byte {
	mem := m.Memory
	lo := uint64(off)
	hi := lo + uint64(n)
	if n < 0 || hi < lo || hi > uint64(len(mem)) {
		return nil
	}
	return mem[lo:hi]
}

func (w *WasiStubs) Clock_time_get64(m *Module, clockID int64, precision int64, timePtr int64) int32 {
	out := w.memSlice64(m, timePtr, 8)
	if out == nil {
		return _wasiEFAULT
	}
	nanos, errno := w.clockNanos(int32(clockID))
	if errno != _wasiESUCCESS {
		return errno
	}
	binary.LittleEndian.PutUint64(out, nanos)
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_close64(m *Module, fd int64) int32 {
	return w.Fd_close(m, int32(fd))
}

func (w *WasiStubs) Sched_yield64(m *Module) int32 {

	return w.Sched_yield(m)
}

func (w *WasiStubs) Fd_fdstat_get64(m *Module, fd int64, ptr int64) int32 {

	out := w.memSlice64(m, ptr, 24)
	if out == nil {
		return _wasiEFAULT
	}
	return w.fdstatFill(int32(fd), out)
}

func (w *WasiStubs) Fd_seek64(m *Module, fd int64, offset int64, whence int64, newOffPtr int64) int32 {
	out := w.memSlice64(m, newOffPtr, 8)
	if out == nil {
		return _wasiEFAULT
	}
	n, errno := w.fdSeek(int32(fd), offset, int(whence))
	if errno != _wasiESUCCESS {
		return errno
	}
	binary.LittleEndian.PutUint64(out, uint64(n))
	return _wasiESUCCESS
}

// iovecSlices64 is iovecSlices for the LP64 iovec layout: {u64 buf,
// u64 len}, 16 bytes per entry.
func (w *WasiStubs) iovecSlices64(m *Module, iovs, iovsLen int64) ([][]byte, bool) {
	if iovsLen < 0 || iovsLen > 1<<20 {
		return nil, false
	}
	iovecs := w.memSlice64(m, iovs, iovsLen*16)
	if iovecs == nil {
		return nil, false
	}
	bufs := make([][]byte, 0, iovsLen)
	for i := int64(0); i < iovsLen; i++ {
		bufPtr := binary.LittleEndian.Uint64(iovecs[i*16:])
		bufLen := binary.LittleEndian.Uint64(iovecs[i*16+8:])
		buf := w.memSlice64(m, int64(bufPtr), int64(bufLen))
		if buf == nil {
			return nil, false
		}
		bufs = append(bufs, buf)
	}
	return bufs, true
}

func (w *WasiStubs) Fd_write64(m *Module, fd int64, iovs int64, iovsLen int64, nwrittenPtr int64) int32 {
	w.mu.Lock()
	dst, _ := w.fdDstLocked(int32(fd))
	w.mu.Unlock()
	bufs, ok := w.iovecSlices64(m, iovs, iovsLen)

	nwrittenSlice := w.memSlice64(m, nwrittenPtr, 8)
	if !ok || nwrittenSlice == nil {
		return _wasiEFAULT
	}
	if dst == nil {
		binary.LittleEndian.PutUint64(nwrittenSlice, 0)
		return _wasiEBADF
	}
	binary.LittleEndian.PutUint64(nwrittenSlice, writeVec(dst, bufs))
	return _wasiESUCCESS
}

func (w *WasiStubs) Proc_exit64(m *Module, code int64) {
	panic(&WasiExitError{Code: int32(code)})
}

// putStrVec64 packs ss as an LP64 char** table (8-byte guest pointers
// at vec) plus NUL-terminated bodies (at buf, guest address bufBase).
// Both slices must already be sized: len(ss)*8 and totalBytesPlusNul.
func putStrVec64(vec, buf []byte, bufBase uint64, ss []string) int32 {
	bufOff := uint64(0)
	for i, s := range ss {
		binary.LittleEndian.PutUint64(vec[i*8:], bufBase+bufOff)
		n := copy(buf[bufOff:], s)
		if n < len(s) {
			return _wasiEFAULT
		}
		bufOff += uint64(n)
		buf[bufOff] = 0
		bufOff++
	}
	return _wasiESUCCESS
}

func (w *WasiStubs) Args_get64(m *Module, argv, argvBuf int64) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	argvSlice := w.memSlice64(m, argv, int64(len(w.args))*8)
	if argvSlice == nil {
		return _wasiEFAULT
	}
	total, ok := totalBytesPlusNul(w.args)
	if !ok {
		return _wasiEFAULT
	}
	argvBufSlice := w.memSlice64(m, argvBuf, int64(total))
	if argvBufSlice == nil {
		return _wasiEFAULT
	}
	return putStrVec64(argvSlice, argvBufSlice, uint64(argvBuf), w.args)
}

func (w *WasiStubs) Args_sizes_get64(m *Module, argcPtr, argvBufLenPtr int64) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	argcSlice := w.memSlice64(m, argcPtr, 8)
	bufLenSlice := w.memSlice64(m, argvBufLenPtr, 8)
	if argcSlice == nil || bufLenSlice == nil {
		return _wasiEFAULT
	}
	total, ok := totalBytesPlusNul(w.args)
	if !ok {
		return _wasiEFAULT
	}
	binary.LittleEndian.PutUint64(argcSlice, uint64(len(w.args)))
	binary.LittleEndian.PutUint64(bufLenSlice, uint64(total))
	return _wasiESUCCESS
}

func (w *WasiStubs) Environ_get64(m *Module, envv, envBuf int64) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	envvSlice := w.memSlice64(m, envv, int64(len(w.env))*8)
	if envvSlice == nil {
		return _wasiEFAULT
	}
	total, ok := totalBytesPlusNul(w.env)
	if !ok {
		return _wasiEFAULT
	}
	envBufSlice := w.memSlice64(m, envBuf, int64(total))
	if envBufSlice == nil {
		return _wasiEFAULT
	}
	return putStrVec64(envvSlice, envBufSlice, uint64(envBuf), w.env)
}

func (w *WasiStubs) Environ_sizes_get64(m *Module, envcPtr, envBufLenPtr int64) int32 {
	w.mu.Lock()
	defer w.mu.Unlock()
	envcSlice := w.memSlice64(m, envcPtr, 8)
	bufLenSlice := w.memSlice64(m, envBufLenPtr, 8)
	if envcSlice == nil || bufLenSlice == nil {
		return _wasiEFAULT
	}
	total, ok := totalBytesPlusNul(w.env)
	if !ok {
		return _wasiEFAULT
	}
	binary.LittleEndian.PutUint64(envcSlice, uint64(len(w.env)))
	binary.LittleEndian.PutUint64(bufLenSlice, uint64(total))
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_fdstat_set_flags64(m *Module, fd, flags int64) int32 {
	return w.Fd_fdstat_set_flags(m, int32(fd), int32(flags))
}

func (w *WasiStubs) Fd_prestat_get64(m *Module, fd, ptr int64) int32 {
	if int32(fd) != 3 {
		return _wasiEBADF
	}

	out := w.memSlice64(m, ptr, 16)
	if out == nil {
		return _wasiEFAULT
	}
	out[0] = 0
	binary.LittleEndian.PutUint64(out[8:], 1)
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_prestat_dir_name64(m *Module, fd, buf, buflen int64) int32 {
	if int32(fd) != 3 {
		return _wasiEBADF
	}
	if buflen < 1 {
		return _wasiESUCCESS
	}
	out := w.memSlice64(m, buf, buflen)
	if out == nil {
		return _wasiEFAULT
	}
	out[0] = '/'
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_read64(m *Module, fd, iovs, iovsLen, nreadPtr int64) int32 {
	w.mu.Lock()
	src, _ := w.fdSrcLocked(int32(fd))
	w.mu.Unlock()
	if src == nil {
		return _wasiEBADF
	}
	bufs, ok := w.iovecSlices64(m, iovs, iovsLen)

	nreadSlice := w.memSlice64(m, nreadPtr, 8)
	if !ok || nreadSlice == nil {
		return _wasiEFAULT
	}
	binary.LittleEndian.PutUint64(nreadSlice, readVec(src, bufs))
	return _wasiESUCCESS
}

func (w *WasiStubs) Fd_readdir64(m *Module, fd, buf, buflen, cookie, bufusedPtr int64) int32 {

	bufSlice := w.memSlice64(m, buf, buflen)
	bufusedSlice := w.memSlice64(m, bufusedPtr, 8)
	if bufSlice == nil || bufusedSlice == nil {
		return _wasiEFAULT
	}
	written, errno := w.fdReaddir(int32(fd), bufSlice, cookie)
	if errno != _wasiESUCCESS {
		return errno
	}
	binary.LittleEndian.PutUint64(bufusedSlice, uint64(written))
	return _wasiESUCCESS
}

func (w *WasiStubs) Path_open64(m *Module, dirFd, dirflags, pathPtr, pathLen, oflags, fsRightsBase, fsRightsInherit, fdflags, openedFdPtr int64) int32 {
	if int32(dirFd) != 3 {
		return _wasiEBADF
	}
	pathSlice := w.memSlice64(m, pathPtr, pathLen)

	outSlice := w.memSlice64(m, openedFdPtr, 4)
	if pathSlice == nil || outSlice == nil {
		return _wasiEFAULT
	}
	fd, errno := w.pathOpen(string(pathSlice), int32(dirflags), int32(oflags), fsRightsBase, int32(fdflags))
	if errno != _wasiESUCCESS {
		return errno
	}
	binary.LittleEndian.PutUint32(outSlice, uint32(fd))
	return _wasiESUCCESS
}

func (w *WasiStubs) Path_filestat_get64(m *Module, dirFd, flags, pathPtr, pathLen, outPtr int64) int32 {
	if int32(dirFd) != 3 {
		return _wasiEBADF
	}
	pathSlice := w.memSlice64(m, pathPtr, pathLen)

	out := w.memSlice64(m, outPtr, 64)
	if pathSlice == nil || out == nil {
		return _wasiEFAULT
	}
	w.mu.Lock()
	fsys := w.fsys
	w.mu.Unlock()
	rel := string(pathSlice)
	var st os.FileInfo
	var err error
	if flags&0x1 != 0 {
		st, err = fsys.Stat(rel)
	} else {
		st, err = fsys.Lstat(rel)
	}
	if err != nil {
		return mapOSError(err)
	}
	writeFilestat(out, st)
	return _wasiESUCCESS
}

func (w *WasiStubs) Random_get64(m *Module, buf, bufLen int64) int32 {
	slice := w.memSlice64(m, buf, bufLen)
	if slice == nil {
		return _wasiEFAULT
	}
	if _, err := rand.Read(slice); err != nil {
		return _wasiEIO
	}
	return _wasiESUCCESS
}
