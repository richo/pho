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

type PhoContext struct {
    rt *PhoRuntime
    ctx unsafe.Pointer
}

func INIT() PhoRuntime {
    log.Print("initializing php runtime")
    tsrm := unsafe.Pointer(C.init_php());
    return PhoRuntime{tsrm}
}

func (rt *PhoRuntime) NewContext() PhoContext {
    log.Print("intializing php context")
    // TODO
    return PhoContext{rt, nil}
}
