package compiler

import (
    "fmt"
    "fractran/fraction"
    "fractran/parser"
)

const qbe_prologue = `
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

const qbe_epilogue = `
    ret %init
}

data $fmt = { b "%lu\n", b 0}
data $argerr = { b "Invalid or missing argument\n", b 0 }
`

const qbe_instr_fmt = `
    @inst%d
    %%times_num%d =l mul %%init, %d
    %%tmp%d =l mul %%times_num%d, %d
    %%res%d =l sar %%tmp%d, %d
    %%test%d =l mul %%res%d, %d
    %%cmp%d =w ceql %%times_num%d, %%test%d
    jnz %%cmp%d, @loop, @next%d
    @next%d`

func CompileProgramQBE(program parser.Program) string {
    res := qbe_prologue
    for ix := 0; ix < len(program); ix += 1 {
        res = fmt.Sprintf("%s, @inst%d %%res%d", res, ix, ix)
    }
    res = fmt.Sprintf("%s\n    call $printf(l $fmt, ..., l %%init)", res)
    for pos, frac := range program {
        res = fmt.Sprintf("%s\n%s", res, CompileInstructionQBE(frac, pos))
    }
    return fmt.Sprintf("%s%s", res, qbe_epilogue)
}

func CompileInstructionQBE(frac fraction.Fraction, pos int) string {
    num, den := frac.Numerator(), frac.Denominator()
    // Fixed point multiplication, and a subsequent bit shift are used
    // in place of expensive divide operations
    //div_mul := (68719476736 / den) + 1
    div_mul, shift_factor := getMultiplier(frac, 29)
    //shift_factor := 36
    return fmt.Sprintf(qbe_instr_fmt,
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

