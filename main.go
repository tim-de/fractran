package main

import (
	"flag"
	"fmt"
	"fractran/compiler"
	"fractran/parser"
	"os"
)

func main() {
    flag.Usage = usageFunc
    outfile := flag.String("o", "frac.out", "File name for generated program")
    c_mode := flag.Bool("c", false, "Compile to c source")
    x86_64_mode := flag.Bool("x", false, "Compile to x86_64 assembly")
    flag.Parse()
    compile_mode := compiler.QBE
    if *c_mode {
        compile_mode = compiler.C
    }
    if *x86_64_mode {
        compile_mode = compiler.X86_64
    }
    if flag.NArg() < 1 {
        fmt.Fprintln(os.Stderr, "No input file given")
        os.Exit(1)
    }
    infile := flag.Arg(0)
    rawdata, err := os.ReadFile(infile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to read input file '%s'\n", infile)
        os.Exit(1)
    }
    text := string(rawdata)
    prog, ok := parser.ReadProgram(text)
    if !ok {
        fmt.Println("Failed to parse program")
        return
    }
    //fmt.Println(prog)
    //fmt.Println(compiler.CompileProgram(prog, 36))
    err = os.WriteFile(*outfile, []byte(compiler.CompileProgram(prog, compile_mode)), 0644)
}

func usageFunc() {
    fmt.Fprintf(flag.CommandLine.Output(), "USAGE:\n  fractran [options] <input file>\nOPTIONS:\n")
    flag.PrintDefaults()
}
