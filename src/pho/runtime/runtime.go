package runtime

import (
    "log"
    // #include "../../../ext/hacks.h"
    "C"
    "sync"
)

type PhoRuntime struct {
    vm *_Ctype_sl_vm_t
}

// TODO Arguments don't make a whole bunch of sense
func (vm PhoRuntime) EvalFile(filename string) int {
    C.eval_file(vm.vm, C.CString(filename))
    return 1
}

var _once_sync sync.Once
func init_slash_runtime() {
    _once_sync.Do(func() {
        log.Print("Initializing slash runtime globally")
        C.init_slash_static()
    })
}

func INIT(name string) PhoRuntime {
    init_slash_runtime()
    log.Printf("initializing slash bowl: %s", name)
    vm := C.init_slash(C.CString(name));
    return PhoRuntime{vm}
}
