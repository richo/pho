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
    if err != nil {
        log.Fatal("Couldn't load libphp5.so")
    }
    php_shims, err := ffi.NewLibrary("lib/hacks.so")
    if err != nil {
        log.Print(err)
        log.Fatal("Couldn't load hacks.so")
    }

    init_runtime, err := php_shims.Fct("init_php", ffi.Int, []ffi.Type{ffi.Int, ffi.Void})
    if err != nil {
        log.Fatal("Couldn't find init_php")
    }

    php_eval, err := php_shims.Fct("eval", ffi.Void, []ffi.Type{ffi.Pointer})

    init_runtime()

    php_eval("echo \"butts lol\n\";")

}
func php_lib_path() string {
    php_path := os.Getenv("PHP_LIB_PATH")

    if (php_path == "") {
        log.Fatal("PHP_LIB_PATH unset")
    }

    return php_path
}
