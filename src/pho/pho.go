package main

import (
    "os"
    "log"
    "unsafe"
    // "reflect"
    ffi "bitbucket.org/binet/go-ffi/pkg/ffi"
    // #include "../../ext/hacks.h"
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

    _php_eval, err := php_shims.Fct("eval", ffi.Void, []ffi.Type{ffi.Pointer})
    // php_get, err := php_shims.Fct("get", ffi.Pointer, []ffi.Type{ffi.Pointer})
    php_get_int, err := php_shims.Fct("get_int_value", ffi.Pointer, []ffi.Type{ffi.Pointer})
    php_set_int, err := php_shims.Fct("set_int_value", ffi.Pointer, []ffi.Type{ffi.Pointer, ffi.Long})

    php_eval := func(s string) {
        log.Printf("PHP> %s", s)
        _php_eval(s)
    }


    init_runtime()
    log.Print("Initialized php runtime")

    log.Print("Setting $max_prints to 10")
    php_set_int("max_prints", 10)

    test_counter := func() {
        php_eval(`
        function butts($max) {
            $count = 0;
            while($count < $max) {
                printf("%d\n", $count);
                $count++;
            }
        }

        butts($max_prints);
`)
    }

    test_counter()

    dump_variable := func(v string, t string) {
        foobar := php_get_int(v);
        var p *C.struct_php_ret_t = (*C.struct_php_ret_t)(unsafe.Pointer(foobar.UnsafeAddr()))
        switch t {
        case "int":
            var i_val **int = (**int)(p.data)
            log.Printf("Got value of %s: %d", v, **i_val)
            return
        case "str":
            var s_val string = C.GoString((*C.char)(p.data))
            log.Printf("Got value of %s: %s", v, s_val)
            return
        }
    }

    php_eval("$foobar = 15;")
    dump_variable("foobar", "int")
    php_eval("$foobar = \"butts\";")
    dump_variable("foobar", "str")

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
