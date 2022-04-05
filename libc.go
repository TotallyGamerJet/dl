package dl

//go:generate onlygo libc.go

//onlygo:open darwin arm64 /usr/lib/libSystem.B.dylib
//onlygo:resolve_with_cgo

//onlygo:open ios arm64 /usr/lib/libSystem.B.dylib
//onlygo:open darwin amd64 /usr/lib/libSystem.B.dylib
//onlygo:open linux amd64 /usr/lib/libc.so

func dlopen(path *byte, mode int) (ret uintptr)

func dlerror() (ret uintptr)

func dlclose(handle uintptr) (ret int)

func dlsym(handle uintptr, symbol *byte) (ret uintptr)
