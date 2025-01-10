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
    outfile := flag.String("o", "out.ssa", "File name for generated program")
    flag.Parse()
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
    err = os.WriteFile(*outfile, []byte(compiler.CompileProgram(prog, 36)), 0644)
}

func usageFunc() {
    fmt.Fprintf(flag.CommandLine.Output(), "USAGE:\n  fractran [options] <input file>\nOPTIONS:\n")
    flag.PrintDefaults()
}
