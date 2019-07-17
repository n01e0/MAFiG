package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 各定義、ここ変えれば言語仕様変えられる
const (
	INC  byte   = 43 // +
	PLUS string = "(*ptr)++;"

	DEC  byte   = 45 // -
	MINS string = "(*ptr)--;"

	NEX  byte   = 62 // >
	MORE string = "ptr++;"

	PRE  byte   = 60 // <
	LESS string = "ptr--;"

	RED  byte   = 44 // ,
	COMA string = "*ptr = inpt();"

	RIT  byte   = 46 // .
	PROD string = "putchar(*ptr);"

	JMP  byte   = 91 // [
	LBRC string = "while (*ptr) {"

	GOL  byte   = 93 // ]
	RBRC string = "}"
)

func main() {
	// 引数で必要なファイルが与えられなかった時にUsage
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " BFsourcefile")
		return
	}
	// ソースファイルを読む
	bfcode, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File read error")
		return
	}
	// 出力(Cソース)ファイル
	ccode, err := os.Create(os.Args[1] + ".c")
	if err != nil {
		fmt.Println("File create error")
		return
	}
    // Cソースファイルの初期設定
	err = cinit(ccode)
	if err != nil {
		fmt.Println("File init error")
		return
	}

	ci := 0 // コードのインデックス
	/* 
    コード解釈
    各文字に対応するコードをCソースファイルに書き込む。
    いちいちエラー処理もしてる
    */
	for ci < len(bfcode) {
		switch bfcode[ci] {
		case INC:
			n, err := ccode.WriteString(PLUS)
			if err != nil|| !(n > 0) {
				fmt.Println("File write error!")
                ccode.Close()
                return
			}
			ci++

		case DEC:
			n, err := ccode.WriteString(MINS)
			if err != nil|| !(n > 0) {
				fmt.Println("File write error!")
                ccode.Close()
                return
			}
			ci++

		case NEX:
			n, err := ccode.WriteString("\n" + MORE)
			if err != nil|| !(n > 0)   {
				fmt.Println("File write error!")
                ccode.Close()
                return
			}
			ci++

		case PRE:
			n, err := ccode.WriteString("\n" + LESS)
			if err != nil|| !(n > 0)  {
				fmt.Println("File write error!")
                ccode.Close()
                return
			}
			ci++

		case RIT:
			n, err := ccode.WriteString("\n" + PROD + "\n")
			if err != nil|| !(n > 0)  {
				fmt.Println("File write error!")
                ccode.Close()
                return
			}
			ci++

		case RED:
			n, err := ccode.WriteString("\n" + COMA + "\n")
			if err != nil|| !(n > 0)  {
				fmt.Println("File write error!")
                ccode.Close()
                return
			}
			ci++

		case JMP:
            n, err := ccode.WriteString("\n" + LBRC + "\n")
			if err != nil|| !(n > 0)  {
				fmt.Println("File write error!")
                ccode.Close()
                return
			}
			ci++

		case GOL:
			n, err := ccode.WriteString("\n" + RBRC + "\n")
			if err != nil|| !(n > 0)  {
				fmt.Println("File write error!")
                ccode.Close()
                return
			}
		    ci++

		default:
			ci++
		}
	}
	ccode.WriteString("putchar(13);putchar(10);\nreturn 0;\n}\n")
	ccode.Close()
}

// Cのソースに共通部分を書き込む(各種関数、宣言等)
func cinit(dest *os.File) error {
	_, err := dest.WriteString("#include <stdio.h>\n#include <stdlib.h>\n\nint inpt(void) {\nint c = getchar();\nif (c == EOF) {\nexit(0);\n}\nreturn c;\n}\n\nint main() {\nchar vmem[32767];\nchar *ptr = vmem;\n")
	/*
	   #include <stdio.h>
	   #include <stdlib.h>

	   int inpt(void) {
	       int c = getchar();
	       if (c == EOF) {
	           exit(0);
	       }
	       return c;
	   }

	   void prnt(int c) {
	       putchar(c);
	   }

	   int main() {
	       char vmem[32767];
	       char *ptr = mem;
	*/
	if err != nil {
		return err
	}
	return nil
}
