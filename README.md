# FRACTRAN compiler
A compiler for the [FRACTRAN](https://en.wikipedia.org/wiki/FRACTRAN)
programming language devised by John Conway, targeting the
[QBE](https://c9x.me/compile/) compiler backend.

This was largely intended as an exercise to get my head around division
by fixed point multiplication, although it could definitely be implemented
more elegantly than it currently (2025-01-10) is.

Further things to solve include the fact that the program ~~is a hardcoded
string that~~ uses a hardcoded value as its input, but you've got to start
somewhere, right?

##Usage
The compiler is invoked as follows
```
fractran [options] <input-file>
```
and accepts the following options:
```
-o <output-file>    Set the file path to write QBE ir to
```
