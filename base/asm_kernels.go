package base

import _ "embed"

// AsmKernels indexes the assembly overrides this bundle was transpiled with
// (kernels/asm/kernels.json in llama-wasm): every exported kernel a native
// body replaces, with what it computes (role), the tensor type it serves
// (quant) and the architectures that carry a body.
//
//go:embed asm_kernels.json
var AsmKernels []byte
