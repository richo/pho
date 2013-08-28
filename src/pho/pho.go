package main

import (
    "os"
    "log"
    // "unsafe"
    // "reflect"
    ffi "bitbucket.org/binet/go-ffi/pkg/ffi"
)

func main() {
    load_path := php_lib_path()
    lib, err := ffi.NewLibrary(load_path)
    php_shims, err := ffi.NewLibrary("lib/hacks.so")

    init_runtime, err := lib.Fct("php_embed_init", ffi.Int, []ffi.Type{ffi.Int, ffi.Void})
    if err != nil {
        log.Fatal("Couldn't find php_embed_init")
    }

    args_ptr := php_init_args_ptr()

    log.Printf("args_ptr: %p", args_ptr)
    log.Printf("args_ptr: %p", &php_shims)

    init_runtime(1, args_ptr)
}

func php_lib_path() string {
    php_path := os.Getenv("PHP_LIB_PATH")

    if (php_path == "") {
        log.Fatal("PHP_LIB_PATH unset")
    }

    return php_path
}

func php_init_args_ptr() uintptr {
    php_shims, err := ffi.NewLibrary("lib/hacks.so")

    init_args_func, err := php_shims.Fct("php_init_args", ffi.Void, []ffi.Type{})
    if err != nil {
        log.Fatal("Couldn't find php_init_args")
    }
    ret := init_args_func()
    return ret.UnsafeAddr()
}
