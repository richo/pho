package main

import (
    "os"
    "log"
    // #include "../../ext/hacks.h"
    "C"
    // "reflect"
    php "pho/runtime"
    args "pho/args"
)


func php_eval(s string) {
    log.Printf("PHP> %s", s)
    cstring := C.CString(s)
    C.eval(cstring)

}

func set_int_value(key string, value int) {
    ckey := C.CString(key)
    C.set_int_value(ckey, (C.long)(value))
}

func get_value (key string) *C.struct_php_ret_t {
    ckey := C.CString(key)
    return C.get_value(ckey)
}

func php_eval_file(filename string) {
    C.eval_file(C.CString(filename))
}

func main() {
    php.INIT()
    args := args.Parse(os.Args)

    for _, script := range args.Scripts {
        log.Printf("Evaluating %s", script)
        php_eval_file(script)
    }
}
