package main

import (
    "os"
    "log"
    "unsafe"
    // #include "../../ext/hacks.h"
    "C"
    // "reflect"
)

func main() {
    init_runtime := func() {
        C.init_php();
    }

    php_eval := func(s string) {
        log.Printf("PHP> %s", s)
        cstring := C.CString(s)
        C.eval(cstring)

    }

    set_int_value := func(key string, value int) {
        ckey := C.CString(key)
        C.set_int_value(ckey, (C.long)(value))
    }

    get_int_value := func(key string) *C.struct_php_ret_t {
        ckey := C.CString(key)
        return C.get_int_value(ckey)
    }

    init_runtime()
    log.Print("Initialized php runtime")

    log.Print("Setting $max_prints to 10")
    set_int_value("max_prints", 10)

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
        s := get_int_value(v)
        var p = *s
        log.Printf("Got value of %s: %s", v, p)
        log.Printf("Got value of %s: %v", v, p)
        log.Printf("Got value of %s: %#v", v, p)
        var data C.union_intern_php_type = (C.union_intern_php_type)(p.data)

        log.Printf("%s.typ: %d", v, (int)(s.typ))
        log.Printf("%s.typ: %d", v, (C.enum_php_types)((*s).typ))

        switch t {
        case "int":
            var i_val int = *(*int)(unsafe.Pointer(&p.data))
            log.Printf("Got value of %s: %d", v, i_val)
            return
        case "str":
            var s_val string = C.GoString((*C.char)(unsafe.Pointer(&data)))
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
