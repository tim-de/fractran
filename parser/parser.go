package parser

import (
	"fmt"
	"fractran/fraction"
	"strings"
)

type Program []fraction.Fraction

func ReadProgram(text string) (prog Program, ok bool) {
    instructions := strings.Split(text, ",")
    for _, instruction := range instructions {
        frac, ok := fraction.ReadFraction(instruction)
        if !ok {
            return Program{}, false
        }
        prog = append(prog, frac)
    }
    return prog, true
}

func (prog Program) String() string {
    if len(prog) == 0 {
        return ""
    }
    str := fmt.Sprint(prog[0])
    for _, frac := range prog[1:] {
        str = fmt.Sprintf("%s; %v", str, frac)
    }
    return str
}
