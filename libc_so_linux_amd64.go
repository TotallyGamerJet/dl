package dl

//go:cgo_import_dynamic _dlopen dlopen "/usr/lib/libc.so"
//go:cgo_import_dynamic _dlerror dlerror "/usr/lib/libc.so"
//go:cgo_import_dynamic _dlclose dlclose "/usr/lib/libc.so"
//go:cgo_import_dynamic _dlsym dlsym "/usr/lib/libc.so"
