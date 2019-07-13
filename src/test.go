package  main

import (
    "fmt"
    "io/ioutil"
    "os"
)

const (
    INC byte =  73 // I
    DEC byte =  69 // i
    NEX byte = 108 // l
    PRE byte = 106 // j
    RED byte = 124 // |
    RIT byte =  21 // !
    JMP byte =  47 // /
    GOL byte =  92 // \
)

var TAPESIZE = 8192

func main(){
    if len(os.Args) < 2 {
        fmt.Println("Usage: " + os.Args[0] + " FILE")
        return
    }

    src, err := ioutil.ReadFile(os.Args[1])

    if err != nil {
        fmt.Println("File read error")
    }

    tape  := make([]byte, TAPESIZE, TAPESIZE)

    ptr   := 0
    si    := 0
    var iobuf byte = 0

    for {
        switch src[si]{
        case INC:
            ptr ++

        case DEC:
            ptr --

        case NEX:
            tape[ptr] ++

        case PRE:
            tape[ptr] --

        case RED:
            iobuf = tape[ptr]
            fmt.Printf("%c",iobuf)

        case RIT:
            fmt.Scanf("%s",&iobuf)
            tape[ptr] = iobuf

        case JMP:
            if tape[ptr] == 0 {
                jc := 0
                for {
                    si ++
                    if tape[si] == JMP {
                        jc ++
                    } else if tape[si] == GOL {
                        jc --
                        if jc <= 0 {
                            break
                        }
                    }
                }
            }

        case GOL:
            if tape[ptr] != 0 {
                jc := 0
                for {
                    si --
                    if tape[si] == GOL {
                        jc ++
                    } else if tape[si] == JMP {
                        jc --
                        if jc <= 0 {
                            break
                        }
                    }
                }
            }
        }

        si++

        if si >= len(tape){
            break
        }
    }



}
