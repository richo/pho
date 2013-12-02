package main

import (
    "os"
    "log"
    // #include "../../ext/hacks.h"
    // #cgo LDFLAGS: -L/home/vagrant/.php/versions/trunk/lib
    "C"
    // "reflect"
    "sync"
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

func php_eval_file(filename string) int {
    ret := int(C.eval_file(C.CString(filename)))
    return ret
}

func php_eval_file_in_wg(wg *sync.WaitGroup, filename string) {
    php_eval_file(filename)
    wg.Done()
}

func main() {
    args := args.Parse(os.Args)
    var wg sync.WaitGroup

    argc := len(args.Rest)
    if argc > 0 {
        php.INIT2(argc, args.Rest)
    } else {
        php.INIT()
    }

    for _, script := range args.Goscripts {
        // Setup a new runtime environment, dispatch in a goroutine
        log.Printf("Dispatching %s in a new noodle", script)
        wg.Add(1)
        go php_eval_file_in_wg(&wg, script)
    }

    for _, script := range args.Scripts {
        log.Printf("Evaluating %s", script)
        php_eval_file(script)
    }

    wg.Wait()
}
