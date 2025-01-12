package compiler

import (
    "fmt"
    "fractran/fraction"
    "fractran/parser"
)

const x86_64_prologue = `
.code64
.global main
.global atol
.global dprintf
.global printf
.global exit

.text
main:
    push %rbp
    cmp $2, %rdi
    jc argerr
    mov 8(%rsi), %rdi
    call atol
    cmp $0, %rax
    jz argerr
loop:
    mov %rax, %r8
`

const x86_64_epilogue = `
    mov %r8, %rsi
    mov $fmt, %rdi
    xor %rax, %rax
    call printf
    xor %rax, %rax
    pop %rbp
    ret

argerr:
    mov $argerrstr, %rsi
    mov $2, %rdi
    xor %rax, %rax
    call dprintf
    mov $1, %rax
    pop %rbp
    ret

.data
argerrstr: .asciz "Invalid starting value\n"
fmt: .asciz "%li\n"
`

const x86_64_instr_fmt = `
    mov %%r8, %%r13
    mov %%r8, %%rax
    mov $%d, %%rcx
    imul %%rcx
    shr $%d, %%rax
    shl $%d, %%rdx
    add %%rdx, %%rax

    mov %%rax, %%r12
    mov $%d, %%rcx
    imul %%rcx
    mov $%d, %%rcx
    imul %%rcx, %%r13
    cmp %%rax, %%r13
    jne 1f
    mov %%r12, %%rax
    jmp loop
1:
`

func CompileProgramX86_64(program parser.Program) string {
    res := x86_64_prologue
    for _, frac := range program {
        res = fmt.Sprintf("%s%s", res, CompileInstructionX86_64(frac))
    }
    return fmt.Sprintf("%s%s", res, x86_64_epilogue)
}

func CompileInstructionX86_64(frac fraction.Fraction) string {
    num, den := frac.Numerator(), frac.Denominator()
    div_mul, shift_factor := getMultiplier(frac, 40)

    return fmt.Sprintf(x86_64_instr_fmt,
        div_mul,
        shift_factor,
        64 - shift_factor,
        den,
        num,
    )
}
