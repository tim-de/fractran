package main

import (
	"fmt"
	"fractran/compiler"
	"fractran/parser"
)

func main() {
    text := "455/33, 11/13, 1/11, 3/7, 11/2, 1/3"
    prog, ok := parser.ReadProgram(text)
    if !ok {
        fmt.Println("Failed to parse program")
        return
    }
    //fmt.Println(prog)
    fmt.Println(compiler.CompileProgram(prog, 36))
}
