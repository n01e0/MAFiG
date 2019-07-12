package main

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

var BUFCODE = 4096

func main(){
    if len(os.Args) < 2 {
        fmt.Println("Usage: " + os.Args[0] + " source file")
        return
    }

    src, err := ioutil.ReadFile(os.Args[1])

    if err != nil {
        fmt.Println("File read error")
    }

    codebuf := make([]byte, BUFCODE, BUFCODE) // bufsize for code
    cpt := 0 // pointer for code
    ci := 0 // index for code
    buf := make([]byte, 1) // buf for io
    fmt.Println(src)

    for {
        switch src[ci]{
        case INC:
            cpt++
            fmt.Printf("+")

        case DEC:
            cpt--
            fmt.Printf("-")

        case NEX:
            codebuf[cpt] ++
            fmt.Printf(">")

        case PRE:
            codebuf[cpt] --
            fmt.Printf("<")

        case RIT:
            buf[0] = codebuf[cpt]
            os.Stdout.Write(buf)

        case RED:
            os.Stdin.Read(buf)
            codebuf[cpt] = buf[0]
            fmt.Printf(",")

        case JMP:
            fmt.Printf("[")
            if codebuf[cpt] == 0 {
                jc := 0
                for {
                    ci ++
                    if src[ci] == JMP {
                        jc ++
                    } else if src[ci] == GOL {
                        jc --
                        if jc < 0 {
                            break
                        }
                    }
                }
            }

        case GOL:
            fmt.Printf("]")
            if codebuf[cpt] != 0 {
                jc := 0
                for {
                    ci--
                    if src[ci] == GOL {
                        ci++
                    } else if src[ci] == JMP {
                        jc--
                        if jc < 0 {
                            break
                        }
                    }
                }
            }
        }

        ci++

        if ci >= len(src){
            break
        }
    }
    fmt.Printf("\n")
}

