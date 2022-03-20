package dl

import _ "unsafe"

//go:cgo_import_dynamic libc_dlopen dlopen "/usr/lib/libc.so"

//go:cgo_import_dynamic libc_dlerror dlerror "/usr/lib/libc.so"

//go:cgo_import_dynamic libc_dlclose dlclose "/usr/lib/libc.so"

//go:cgo_import_dynamic libc_dlsym dlsym "/usr/lib/libc.so"
