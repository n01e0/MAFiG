package main

import (
    "fmt"
    "os"
)

func main(){
    f, err := os.Open(os.Args[1])

    if err != nil{
        fmt.Println("File open error")
    }

    defer f.Close()
    buf := make([]byte, 4)
    for{
        n, err := f.Read(buf)

        if n == 0{
            break
        }
        if err != nil {
            break
        }
        fmt.Println(string(buf[:n]))
    }
}
