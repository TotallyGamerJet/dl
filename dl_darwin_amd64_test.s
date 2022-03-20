#include "textflag.h"

TEXT ·libc_malloc(SB),NOSPLIT, $0-0
	CALL runtime·entersyscall(SB)
	MOVQ _size+8(SP), DI
	MOVD ·_malloc(SB), AX
	CALL AX
	MOVQ AX, ret+16(SP)
	CALL runtime·exitsyscall(SB)
	RET
