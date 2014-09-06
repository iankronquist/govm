Govm
====
A simple 32 bit vm written in Go.
Roughly based off of a [tutorial for C#][0].


Running the assembler:
----------------------
```sh
$ go run assembler/main.go --input some_file.gasm --output test
```

G32 Assmbly language:
---------------------
There are 5 registers:
* `A`: 1 byte
* `B`: 1 byte
* `D`: 2 bytes, a superset of `A` and `B`
* `X`: 2 bytes 
* `Y`: 2 bytes 
    
There are 4 instructions:
* `LDA`
* `LDX`
* `STA`
* `END`

You can define pretty much any sort of label, as long as it ends with a colon.

Here's an example program:
```asm
START:
    LDA 0x41
    LDX 0xA000
    STA X
    END START 
```

[0]: http://www.icemanind.com/VMCS.pdf
