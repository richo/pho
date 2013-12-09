package args

import (
    "log"
    s "strings"
)

/* Defines a custom parser that deals with the parsing of the pho runtime's flag.
 */

const (
    pNone = 0
    pBin = iota
    pRest = iota
    pAddress = iota
)

type PhoArgs struct {
    Prefork bool
    Bin string
    Address string
    Scripts []string
    Rest []string
}

func initPhoArgs() PhoArgs {
    r := PhoArgs{}
    r.Prefork = false
    return r
}

func Parse(args []string) PhoArgs {
    r := initPhoArgs()

    state := pBin

    for _, i := range args {
        switch state {
        case pBin:
            r.Bin = i
            state = pNone
            break
        case pAddress:
            r.Address = i
            state = pNone
            break
        case pRest:
            r.Rest = append(r.Rest, i)
            break
        case pNone:
            if s.HasPrefix(i, "--") {
                switch i {
                case "--address":
                    state = pAddress
                case "--prefork":
                    r.Prefork = true;
                case "--":
                    state = pRest
                default:
                    log.Panicf("Unrecognised command line switch: %s", i)
                }
                break
            } else {
                r.Scripts = append(r.Scripts, i)
                break
            }
        }
    }
    return r
}
