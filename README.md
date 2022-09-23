# Brainfuck Interpreter in Go

A Go-based interpreter to read and run Brainfuck programs in pure Go in the simplest way.

## Using

```
$ CGO_ENABLED=0 go build -o bf .
$ ./bf
```

### Fibonacci

**./bf fibonacci.bf**

```bf
+++++++++++>+>>>>++++++++++++++++++++++++++++++++++++++
++++++>++++++++++++++++++++++++++++++++<<<<<<[>[>>>>>>+
>+<<<<<<<-]>>>>>>>[<<<<<<<+>>>>>>>-]<[>++++++++++[-<-[>
>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]>[<<[>>>+<<<-]>>[-]]<
<]>>>[>>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]>[<<+>>[-]]<<<
<<<<]>>>>>[++++++++++++++++++++++++++++++++++++++++++++
++++.[-]]++++++++++<[->-<]>++++++++++++++++++++++++++++
++++++++++++++++++++.[-]<<<<<<<<<<<<[>>>+>+<<<<-]>>>>[<
<<<+>>>>-]<-[>>.>.<<<[-]]<<[>>+>+<<<-]>>>[<<<+>>>-]<<[<
+>-]>[<+>-]<<<-]
```

### Hello World

**./bf hello-world.bf**

```
>++++++++[<+++++++++>-]<.
>++++[<+++++++>-]<+.
+++++++..
+++.
>>++++++[<+++++++>-]<++.
------------.
>++++++[<+++++++++>-]<+.
<.
+++.
------.
--------.
>>>++++[<++++++++>-]<+.
```

## Tokens

| Token | Description |
| ----- | --------------------------- |
| >	| Move the pointer to the right |
| <	| Move the pointer to the left |
| +	| Increment the memory cell at the pointer |
| -	| Decrement the memory cell at the pointer |
| .	| Output the character signified by the cell at the pointer |
| ,	| Input a character and store it in the cell at the pointer |
| [	| Jump past the matching ] if the cell at the pointer is 0 |
| ]	| Jump back to the matching [ if the cell at the pointer is nonzero |
| any | All characters other than ><+-.,[] should be considered comments and ignored |

## Brainfuck

Brainfuck is an esoteric programming language created in 1993 by Urban Müller.

Notable for its extreme minimalism, the language consists of only eight simple commands, a data pointer and an instruction pointer. While it is fully Turing complete, it is not intended for practical use, but to challenge and amuse programmers. Brainfuck requires one to break commands into microscopic steps.

The language's name is a reference to the slang term brainfuck, which refers to things so complicated or unusual that they exceed the limits of one's understanding as it was not meant or made for designing actual software but to challenge the boundaries of computer programming.

Read more: https://en.wikipedia.org/wiki/Brainfuck

### Online runner

- https://www.dcode.fr/brainfuck-language

### ASCII Table

https://en.wikipedia.org/wiki/ASCII

You are welcome to add more features and send Pull request.

© Copyright 2022, Max Base
