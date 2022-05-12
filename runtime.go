package dl

import (
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

//go:linkname syscall_syscallX syscall.syscallX
func syscall_syscallX(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) // from runtime/sys_darwin_64.go

func cstring(s string) (*byte, error) {
	if strings.IndexByte(s, 0) != -1 {
		return nil, syscall.EINVAL
	}
	a := make([]byte, len(s)+1)
	copy(a, s)
	return &a[0], nil
}

func gostring(p uintptr) (ret string) {
	var c = *(**byte)(unsafe.Pointer(&p))
	for *c != 0 {
		ret += string(rune(*c))
		c = (*byte)(unsafe.Add(unsafe.Pointer(c), 1))
	}
	return
}

func funcPC(f interface{}) uintptr {
	return reflect.ValueOf(f).Pointer()
}
