package compiler

import (
	"fractran/parser"
    "fractran/fraction"
)

type CompileMode int

const (
    QBE CompileMode = 0
    C CompileMode = iota
    X86_64 CompileMode = iota
)

func CompileProgram(program parser.Program, mode CompileMode) string {
    switch mode {
    case QBE:
        return CompileProgramQBE(program)
    case C:
        return CompileProgramC(program)
    case X86_64:
        return CompileProgramX86_64(program)
    default:
        panic("Invalid compilation mode")
    }
}

// Finds the fixed point multiplier and shift factor
// equivalent by multiplication by the given fraction
// without exceeding 2^max_exponent
func getMultiplier(frac fraction.Fraction, max_exponent int) (multiplier int64, shift int64) {
    for (frac.Numerator() << shift) / frac.Denominator() <= 1 << max_exponent {
        shift += 1
    }
    shift -= 1
    multiplier = (frac.Numerator() << shift) / frac.Denominator() + 1
    return
}
