package compiler

import (
	"fmt"
	"fractran/fraction"
	"fractran/parser"
)

const prologue = `
export function w $main(w %argc, l %argv) {
    @start
    %argtest =w cugtw %argc, 1
    jnz %argtest, @convert, @argfail
    @argfail
    call $printf(l $argerr)
    jmp @end
    @convert
    %arg1ptr =l add %argv, 8
    %arg1 =l loadl %arg1ptr
    %startval =l call $atol(l %arg1)
    jnz %startval, @run, @argfail
    @run
    %r =l call $run(l %startval)
    call $printf(l $fmt, ..., l %r)
    @end
    ret 0
}

function l $run(l %num) {
    @start
    @loop
    %init =l phi @start %num`

const epilogue = `
    ret %init
}

data $fmt = { b "%lu\n", b 0}
data $argerr = { b "Invalid or missing argument\n", b 0 }
`

const instr_fmt = `
    @inst%d
    %%times_num%d =l mul %%init, %d
    %%tmp%d =l mul %%times_num%d, %d
    %%res%d =l sar %%tmp%d, %d
    %%test%d =l mul %%res%d, %d
    %%cmp%d =w ceql %%times_num%d, %%test%d
    jnz %%cmp%d, @loop, @next%d
    @next%d`

func CompileProgram(prog parser.Program) string {
    res := prologue
    for ix := 0; ix < len(prog); ix += 1 {
        res = fmt.Sprintf("%s, @inst%d %%res%d", res, ix, ix)
    }
    for pos, frac := range prog {
        res = fmt.Sprintf("%s\n%s", res, CompileInstruction(frac, pos))
    }
    return fmt.Sprintf("%s%s", res, epilogue)
}

func CompileInstruction(frac fraction.Fraction, pos int) string {
    num, den := frac.Numerator(), frac.Denominator()
    // Fixed point multiplication, and a subsequent bit shift are used
    // in place of expensive divide operations
    div_mul := (68719476736 / den) + 1
    shift_factor := 36
    return fmt.Sprintf(instr_fmt,
        pos,
        pos, num,
        pos, pos, div_mul,
        pos, pos, shift_factor,
        pos, pos, den,
        pos, pos, pos,
        pos, pos,
        pos,
    )
}
