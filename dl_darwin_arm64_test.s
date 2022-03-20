#include "textflag.h"

TEXT ·libc_malloc(SB),NOSPLIT, $0-0
    BL runtime·entersyscall(SB)
    MOVD _size+0(FP), R0
    MOVD ·_malloc(SB), R16
    CALL R16
    MOVD R0, ret+8(FP)
    BL runtime·exitsyscall(SB)
    RET
