package main

import "fmt"

func main() {
    var str string
    fmt.Scanf("%s",&str)
    var runes = []rune(str)
    fmt.Printf("%d\n",str[0])
    fmt.Println(string(runes))
    fmt.Println(runes)
}
