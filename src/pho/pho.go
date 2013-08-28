package main

import (
    "os"
    "fmt"
    ffi "bitbucket.org/binet/go-ffi/pkg/ffi"
)

func main() {
    load_path := php_lib_path()
    lib, err := ffi.NewLibrary(load_path)

    init_runtime := lib.Fct("php_embed_init", ffi.Int, []ffi.Type{ffi.Int, ffi.Char})

    init_runtime(0, "a")
}

func php_lib_path() string {
    php_path := os.Getenv("PHP_LIB_PATH")

    if (php_path == "") {
        fatal("PHP_LIB_PATH unset")
    }

    return php_path
}
