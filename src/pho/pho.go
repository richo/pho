package main

import (
    "os"
    "log"
    // #include "../../ext/hacks.h"
    // #include <dlfcn.h>
    "C"
    "unsafe"
    // "reflect"
    "flag"
)

type PhoRuntime struct {
    // void ***rt
    rt unsafe.Pointer

}

func init_runtime() PhoRuntime {
    // Idempotent, refcount aside
    C.dlopen(C.CString("libphp.so"), C.RTLD_NOW)
    rt := unsafe.Pointer(C.init_php());
    return PhoRuntime{rt}
}

func php_eval(s string) {
    log.Printf("PHP> %s", s)
    cstring := C.CString(s)
    C.eval(cstring)

}

func set_int_value(key string, value int) {
    ckey := C.CString(key)
    C.set_int_value(ckey, (C.long)(value))
}

func get_int_value (key string) *C.struct_php_ret_t {
    ckey := C.CString(key)
    return C.get_int_value(ckey)
}

func php_eval_file(filename string) {
    C.eval_file(C.CString(filename))
}

func main() {

    // gos := flag.String("go", "", "run a file in a goroutine")

    flag.Parse()

    argv := flag.Args()
    if len(argv) > 1 {
        flag.PrintDefaults()
        os.Exit(1)
    }
    if len(argv) < 1 {
        // TODO repl
        flag.PrintDefaults()
        os.Exit(1)
    }

    init_runtime()
    log.Print("Initialized php runtime")

    php_eval_file(argv[0])
}
