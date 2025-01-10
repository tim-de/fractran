package main

import (
	"fmt"
	"fractran/compiler"
	"fractran/parser"
)

func main() {
    text := "3/2"
    prog, ok := parser.ReadProgram(text)
    if !ok {
        fmt.Println("Failed to parse program")
        return
    }
    //fmt.Println(prog)
    fmt.Println(compiler.CompileProgram(prog, 144))
}
