package dl

import "unsafe"

var libc_dlopen uintptr

func dlopen(path *byte, mode int) (ret uintptr) {
	r0, _, _ := syscall_syscall(funcPC(libc_dlopen_trampoline), uintptr(unsafe.Pointer(path)), uintptr(mode), 0)
	ret = uintptr(r0)
	return
}
func libc_dlopen_trampoline()

var libc_dlerror uintptr

func dlerror() (ret uintptr) {
	r0, _, _ := syscall_syscall(funcPC(libc_dlerror_trampoline), 0, 0, 0)
	ret = r0
	return
}
func libc_dlerror_trampoline()

var libc_dlclose uintptr

func dlclose(handle uintptr) (ret int) {
	r0, _, _ := syscall_syscall(funcPC(libc_dlclose_trampoline), uintptr(handle), 0, 0)
	ret = int(r0)
	return
}
func libc_dlclose_trampoline()

var libc_dlsym uintptr

func dlsym(handle uintptr, symbol *byte) (ret uintptr) {
	r0, _, _ := syscall_syscall(funcPC(libc_dlsym_trampoline), uintptr(handle), uintptr(unsafe.Pointer(symbol)), 0)
	ret = uintptr(r0)
	return
}
func libc_dlsym_trampoline()
