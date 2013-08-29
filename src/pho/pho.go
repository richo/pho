package main

import (
    "os"
    "log"
    _ "unsafe"
    // "reflect"
    ffi "bitbucket.org/binet/go-ffi/pkg/ffi"
    "C"
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
    // php_get, err := php_shims.Fct("get", ffi.Pointer, []ffi.Type{ffi.Pointer})
    php_get_int, err := php_shims.Fct("get_int_value", ffi.Long, []ffi.Type{ffi.Pointer})
    return_long, err := php_shims.Fct("return_long", ffi.Long, []ffi.Type{})

    init_runtime()
    log.Print("Initialized php runtime")

    test_counter := func() {
        php_eval(`
        function butts() {
            $count = 0;
            while($count < 10) {
                printf("%d\n", $count);
                $count++;
            }
        }

        butts();
`)
    }

    test_counter()

    dump_variable := func(v string) {
        foobar := php_get_int(v);
        butts := foobar.Int()
        log.Printf("Got value of %s: %d", v, butts)
    }


    l := return_long()
    log.Printf("Long Value: %d", l.Int())

    php_eval("$foobar = \"butts\";")
    dump_variable("foobar")
    dump_variable("$foobar")

    log.Print("Evaling echo")
    php_eval(`echo "Butts\n";`)

}
func php_lib_path() string {
    php_path := os.Getenv("PHP_LIB_PATH")

    if (php_path == "") {
        log.Fatal("PHP_LIB_PATH unset")
    }

    return php_path
}
