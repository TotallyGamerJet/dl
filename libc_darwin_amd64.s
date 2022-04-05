// File generated using onlygo. DO NOT EDIT!!!
#include "textflag.h"

//func dlopen(path *byte, mode int) (ret uintptr)
TEXT ·dlopen(SB), NOSPLIT, $0-0
	CALL runtime·entersyscall(SB)
	MOVQ _path+8(SP), DI
	MOVQ _mode+16(SP), SI
	CALL dlopen(SB)
	MOVQ AX, ret+24(SP)
	CALL runtime·exitsyscall(SB)
	RET

//func dlerror() (ret uintptr)
TEXT ·dlerror(SB), NOSPLIT, $0-0
	CALL runtime·entersyscall(SB)
	CALL dlerror(SB)
	MOVQ AX, ret+8(SP)
	CALL runtime·exitsyscall(SB)
	RET

//func dlclose(handle uintptr) (ret int)
TEXT ·dlclose(SB), NOSPLIT, $0-0
	CALL runtime·entersyscall(SB)
	MOVQ _handle+8(SP), DI
	CALL dlclose(SB)
	MOVQ AX, ret+16(SP)
	CALL runtime·exitsyscall(SB)
	RET

//func dlsym(handle uintptr, symbol *byte) (ret uintptr)
TEXT ·dlsym(SB), NOSPLIT, $0-0
	CALL runtime·entersyscall(SB)
	MOVQ _handle+8(SP), DI
	MOVQ _symbol+16(SP), SI
	CALL dlsym(SB)
	MOVQ AX, ret+24(SP)
	CALL runtime·exitsyscall(SB)
	RET

