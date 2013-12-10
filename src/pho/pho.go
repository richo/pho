package main

import (
    "os"
    "log"
    // #include "../../ext/hacks.h"
    "C"
    // "reflect"
    "sync"
    "net"
    php "pho/runtime"
    args "pho/args"
    network "pho/network"
)

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

func noop(interface {}) {
}

func main() {
    var (
        rt php.PhoRuntime
        err error
        sock net.Listener
    )
    args := args.Parse(os.Args)

    argc := len(args.Rest)
    if argc > 0 {
        rt = php.INIT2(argc, args.Rest)
    } else {
        rt = php.INIT()
    }

    if (args.Prefork) {
        if (len(args.Scripts) != 1) {
            log.Panicf("Can't prefork without exactly one script")
        }

        if (args.Address != "" && args.Port != 0) {
            log.Printf("Binding to %s:%d", args.Address, args.Port)
            sock, err = network.MakeSocketFromAddressAndPort(args.Address, args.Port)
        } else if (args.Socket != "") {
            log.Printf("Binding to %s", args.Socket)
            sock, err = network.MakeSocketFromSocketPath(args.Socket)
        } else {
            log.Panicf("Can't prefork without either a bind address and port, or socket path")
        }

        if (err != nil) {
            log.Panicf("Couldn't bind socket");
        }

        noop(sock)
        noop(rt)

        log.Panicf("Prefork not yet implemented")
    }

    for _, script := range args.Scripts {
        log.Printf("Evaluating %s", script)
        php_eval_file(script)
    }
}
