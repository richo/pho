package main

import (
    "os"
    "fmt"
    // #include "../../ext/hacks.h"
    "C"
    // "reflect"
    "sync"
    "sync/atomic"
    slash "pho/runtime"
    args "pho/args"
)

var vm_id int64 = 0;

func NewVMID() int64 {
    return atomic.AddInt64(&vm_id, 1)
}

func main() {
    args := args.Parse(os.Args)
    var wg sync.WaitGroup

    for _, script := range args.Goscripts {
        // Setup a new runtime environment, dispatch in a goroutine
        go func() {
            wg.Add(1)
            vm_name := fmt.Sprintf("noodle%d", NewVMID())
            vm := slash.INIT(vm_name)
            vm.EvalFile(script)
            wg.Done()
        }()
    }

    for _, script := range args.Scripts {
        vm_name := fmt.Sprintf("noodle%d", NewVMID())
        vm := slash.INIT(vm_name)
        vm.EvalFile(script)
    }

    wg.Wait()
}
