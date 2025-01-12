package compiler

import (
    "fmt"
    "fractran/fraction"
    "fractran/parser"
)

const c_prologue = `
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>

int step(int64_t*);

int main(int argc, char **argv) {
    if (argc < 2) {
        fprintf(stderr, "Please supply a starting value\n");
        return 1;
    }
    int64_t value = (int64_t) atol(argv[1]);
    if (value == 0) {
        fprintf(stderr, "Invalid starting value '%s'", argv[1]);
        return 1;
    }

    while (step(&value)) {}

    printf("%li\n", value);
    return 0;
}

int step(int64_t *value) {
`

const c_epilogue = `
    return 0;
}
`

const c_instr_fmt = `
    if ((*value * %d) %% %d == 0) {
        *value = (*value * %d) / %d;
        return 1;
    }
`

func CompileProgramC(program parser.Program) string {
    res := c_prologue
    for _, frac := range(program) {
        res = fmt.Sprintf("%s\n%s", res, CompileInstructionC(frac))
    }
    return fmt.Sprintf("%s%s", res, c_epilogue)
}

func CompileInstructionC(frac fraction.Fraction) string {
    return fmt.Sprintf(c_instr_fmt,
        frac.Numerator(), frac.Denominator(),
        frac.Numerator(), frac.Denominator(),
    )
}
