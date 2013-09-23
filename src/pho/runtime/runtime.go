package runtime

import (
    "log"
    // #include "../../../ext/hacks.h"
    "C"
    "unsafe"
)

type PhoRuntime struct {
    // void ***rt
    tsrm unsafe.Pointer
}

func INIT() PhoRuntime {
    log.Print("initializing php runtime")
    tsrm := unsafe.Pointer(C.init_php());
    return PhoRuntime{tsrm}
}
