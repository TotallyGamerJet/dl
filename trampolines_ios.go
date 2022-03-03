package dl

import _ "unsafe"

//go:linkname libc_dlopen libc_dlopen
//go:cgo_import_dynamic libc_dlopen dlopen "/usr/lib/libSystem.dylib"

//go:linkname libc_dlerror libc_dlerror
//go:cgo_import_dynamic libc_dlerror dlerror "/usr/lib/libSystem.dylib"

//go:linkname libc_dlclose libc_dlclose
//go:cgo_import_dynamic libc_dlclose dlclose "/usr/lib/libSystem.dylib"

//go:linkname libc_dlsym libc_dlsym
//go:cgo_import_dynamic libc_dlsym dlsym "/usr/lib/libSystem.dylib"
