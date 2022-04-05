// File generated using onlygo. DO NOT EDIT!!!
#include "textflag.h"

//func dlopen(path *byte, mode int) (ret uintptr)
TEXT ·dlopen(SB), NOSPLIT, $0-0
	BL runtime·entersyscall(SB)
	MOVD _path+0(FP), R0
	MOVD _mode+8(FP), R1
	CALL _dlopen(SB)
	MOVD R0, ret+16(FP)
	BL runtime·exitsyscall(SB)
	RET

//func dlerror() (ret uintptr)
TEXT ·dlerror(SB), NOSPLIT, $0-0
	BL runtime·entersyscall(SB)
	CALL _dlerror(SB)
	MOVD R0, ret+0(FP)
	BL runtime·exitsyscall(SB)
	RET

//func dlclose(handle uintptr) (ret int)
TEXT ·dlclose(SB), NOSPLIT, $0-0
	BL runtime·entersyscall(SB)
	MOVD _handle+0(FP), R0
	CALL _dlclose(SB)
	MOVD R0, ret+8(FP)
	BL runtime·exitsyscall(SB)
	RET

//func dlsym(handle uintptr, symbol *byte) (ret uintptr)
TEXT ·dlsym(SB), NOSPLIT, $0-0
	BL runtime·entersyscall(SB)
	MOVD _handle+0(FP), R0
	MOVD _symbol+8(FP), R1
	CALL _dlsym(SB)
	MOVD R0, ret+16(FP)
	BL runtime·exitsyscall(SB)
	RET

