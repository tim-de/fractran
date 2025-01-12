# FRACTRAN compiler
A compiler for the [FRACTRAN](https://en.wikipedia.org/wiki/FRACTRAN)
programming language devised by John Conway, targeting the
[QBE](https://c9x.me/compile/) compiler backend.

This was largely intended as an exercise to get my head around division
by fixed point multiplication, although it could definitely be implemented
more elegantly than it currently (2025-01-10) is.

The resulting programs take the starting value as a command line argument,
and write the result to stdout

## Usage
The compiler is invoked as follows
```
fractran [options] <input-file>
```
and accepts the following options:
```
-c                  Output C source code
-x                  Output X86_64 assembly
-q                  Output QBE IR (default)
-o <output-file>    Set the file path to write QBE ir to
```

## Issues
While working on this, and comparing it to a fledgling interpreter written
in an interpreted and managed language, it showed some errors. This is due to
some of the intermediate values in the fixed point multiplications overflowing
the 64 bits of the variables used. This affects both the QBE and C modes, as
both use a bitwidth of 64 for all data. This does not, however, affect the
X86_64 mode, as the multiplication instructions in the X86_64 ISA return the
result in two registers, permitting up to 128 bits of result.
