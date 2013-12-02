package runtime

import (
    "log"
    // #include "../../../ext/hacks.h"
    // #cgo LDFLAGS: -L/home/vagrant/.php/versions/trunk/lib
    "C"
    "unsafe"
)

type PhoRuntime struct {
    // void ***rt
    tsrm unsafe.Pointer
}

type PhoContext struct {
    rt *PhoRuntime
    Context unsafe.Pointer
}

func INIT() PhoRuntime {
    log.Print("initializing php runtime")
    tsrm := unsafe.Pointer(C.init_php());
    return PhoRuntime{tsrm}
}

func INIT2(argc int, argv []string) PhoRuntime {
    var php_argv []*_Ctype_char
    php_argv = make([]*_Ctype_char, argc)

    for i, arg := range argv {
        php_argv[i] = C.CString(arg)
    }
    tsrm := unsafe.Pointer(C.init_php2(C.int(argc), (**_Ctype_char)(&php_argv[0])))
    return PhoRuntime{tsrm}
}
