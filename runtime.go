package dl

import (
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

//go:linkname syscall_syscall syscall.syscall
//go:linkname syscall_syscall6 syscall.syscall6
//go:linkname runtime_gostring runtime.gostring

func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr)              // runtime/sys_darwin.go
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) // runtime/sys_darwin.go
func runtime_gostring(p *byte) string                                           // runtime/string.go

func funcPC(f interface{}) uintptr {
	return reflect.ValueOf(f).Pointer()
}

func cstring(s string) (*byte, error) {
	if strings.IndexByte(s, 0) != -1 {
		return nil, syscall.EINVAL
	}
	a := make([]byte, len(s)+1)
	copy(a, s)
	return &a[0], nil
}

func gostring(p uintptr) string {
	return runtime_gostring(*(**byte)(unsafe.Pointer(&p)))
}
