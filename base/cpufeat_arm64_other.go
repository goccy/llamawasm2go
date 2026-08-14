//go:build arm64 && !darwin && !linux

package base

// CPUDotProd: no detection story on this OS — run the portable
// bodies.
var CPUDotProd = false
