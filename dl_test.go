package dl

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

func getLibc() string {
	switch runtime.GOOS {
	case "darwin", "ios":
		return "/usr/lib/libSystem.dylib"
	case "linux":
		return "/usr/lib/libc.so"
	default:
		panic(fmt.Sprintf("unknown GOOS: %s", runtime.GOOS))
	}
}

var _malloc uintptr // pointer to malloc function

func libc_malloc(uintptr) unsafe.Pointer

func TestDL(t *testing.T) {
	var err error
	libc, err := Open(getLibc(), ScopeGlobal)
	if err != nil {
		t.Fatal(err)
		return
	}
	_malloc, err = libc.Lookup("malloc")
	if err != nil {
		t.Fatal(err)
		return
	}
	if _malloc == 0 {
		t.Failed()
	}
	const MallocSize = 5
	if libc_malloc(MallocSize) == nil {
		t.Failed()
	}
	_, err = libc.Lookup("UnknownFunctionName")
	if err == nil {
		t.Failed()
	}
}
