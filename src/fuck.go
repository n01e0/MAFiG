package main

import(
    "fmt"
    "io/ioutil"
    "os"
)

const (
    INC byte = 62 // +
    DEC byte = 60 // -
    NEX byte = 43 // >
    PRE byte = 45 // <
    RED byte = 44 // ,
    RIT byte = 46 // .
    JMP byte = 91 // [
    GOL byte = 93 // ]
)

var BUFSIZE = 4096 // コード用のバッファサイズ

func main(){
    if len(os.Args) < 2 {
        fmt.Println("Usage: " + os.Args[0] + "file")
        return
    }

    code, err := ioutil.ReadFile(os.Args[1])

    if err != nil {
        fmt.Println("File read error")
    }

    codebuf := make([]byte, BUFSIZE, BUFSIZE) // BUFSIZE分のスライスを確保
    ptr := 0 // コードのポインタ
    ci := 0 // コードのインデックス
    buf := make([]byte, 1) // 入出力用のバッファ
    fmt.Println(code)

    for {
        switch code[ci]{
        case INC:
            ptr++

        case DEC:
            ptr--

        case NEX:
            codebuf[ptr]++

        case PRE:
            codebuf[ptr]--

        case RIT:
            buf[0] = codebuf[ptr]
            os.Stdout.Write(buf)

        case RED:
            os.Stdin.Read(buf)
            codebuf[ptr] = buf[0]

        case JMP:
            if codebuf[ptr] == 0{
                jc := 0     // jump counter
                for {
                    ci ++
                    if code[ci] == JMP {
                        jc++
                    } else if code[ci] == GOL {
                        jc--
                        if jc < 0 {
                            break
                        }
                    }
                }
            }

        case GOL:
            if codebuf[ptr] != 0 {
                jc := 0
                for {
                    ci --
                    if code[ci] == GOL {
                        ci ++
                    } else if code[ci] == JMP {
                        jc --
                        if jc < 0 {
                            break
                        }
                    }
                }
            }
        }

        ci ++

        if ci >= len(code) {
            break
        }
    }
    fmt.Printf("\n")
}

