package main

import (
    "os"
    "log"
    "unsafe"
    ffi "bitbucket.org/binet/go-ffi/pkg/ffi"
)

func main() {
    load_path := php_lib_path()
    lib, err := ffi.NewLibrary(load_path)

    init_runtime, err := lib.Fct("php_embed_init", ffi.Int, []ffi.Type{ffi.Int, ffi.Void})
    if err != nil {
        log.Fatal("Couldn't find php_embed_init")
    }

    init_runtime(0, php_init_args_ptr())
}

func php_lib_path() string {
    php_path := os.Getenv("PHP_LIB_PATH")

    if (php_path == "") {
        log.Fatal("PHP_LIB_PATH unset")
    }

    return php_path
}

func php_init_args_ptr() unsafe.Pointer {
    php_shims, err := ffi.NewLibrary("lib/hacks.so")

    init_args_func = php_shims.Fct("php_init_args", ffi.Void, []ffi.Type{})
    if err != nil {
        log.Fatal("Couldn't find php_init_args")
    }
    return init_args_func()
}
