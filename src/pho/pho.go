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
    _, err := ffi.NewLibrary(load_path)
    if err != nil {
        log.Fatal("Couldn't load libphp5.so")
    }
    php_shims, err := ffi.NewLibrary("lib/hacks.so")
    if err != nil {
        log.Print(err)
        log.Fatal("Couldn't load hacks.so")
    }

    init_runtime, err := php_shims.Fct("init_php", ffi.Int, []ffi.Type{})
    if err != nil {
        log.Fatal("Couldn't find init_php")
    }

    php_eval, err := php_shims.Fct("eval", ffi.Void, []ffi.Type{ffi.Pointer})

    init_runtime()

    test_counter := func() {
        php_eval(`
function count() {
    static $counter = 0;

    while($counter < 100) {
        printf("%d\n", $counter);
        $counter++;
    }
}
count();
`)
    }

    go test_counter()
    go test_counter()
    test_counter()

}
func php_lib_path() string {
    php_path := os.Getenv("PHP_LIB_PATH")

    if (php_path == "") {
        log.Fatal("PHP_LIB_PATH unset")
    }

    return php_path
}
