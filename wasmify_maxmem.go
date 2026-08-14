package wasm2go

import base "github.com/goccy/llamawasm2go/base"

// SetMaxMemory caps linear-memory growth: memory.grow fails (returns
// -1) rather than taking the module past n bytes. Zero restores the
// module's own ceiling.
func SetMaxMemory(m *base.Module, n uint64) { m.MaxMem = n }
