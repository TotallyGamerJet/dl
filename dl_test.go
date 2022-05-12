package dl

import (
	"fmt"
	"runtime"
	"testing"
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

func TestDL(t *testing.T) {
	var err error
	libc, err := Open(getLibc(), ScopeGlobal)
	if err != nil {
		t.Fatal(err)
		return
	}
	var _malloc uintptr
	_malloc, err = libc.Lookup("malloc")
	if err != nil {
		t.Fatal(err)
		return
	}
	if _malloc == 0 {
		t.Failed()
	}
	const MallocSize = 5
	if ret, _, _ := syscall_syscallX(_malloc, MallocSize, 0, 0); ret == 0 {
		t.Failed()
	}
	_, err = libc.Lookup("UnknownFunctionName")
	if err == nil {
		t.Failed()
	}
}
