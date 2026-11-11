# Custom Assembly
> `x16asm`

Because I am not a huge fan of the syntax that Mano's Basic Computer uses for its assembly code, I have opted to switch the `,` character with `:` to follow more modern assembly languages. Additionaly, I have added another directive `ADR` which will place the address of the label it is given in that location. I am not sure if `ADR` is a standard compiler directive for Mano's Basic Computer, however, I added it as it makes life a whole lot easier.

# Define Coroutine / Function
```asm
; Accumulator = Function Argument (unless specific memory addresses defined)
FUNCTION_NAME: HEX 0000 ; Return Address

    ; DO SHIT

    LDA RETURN_VALUE ; Function Return Value Should be placed in the accumulator
    BUN FUNCTION_NAME I ; Return

```

# Pointer Offsets
```asm
LDA HEAP_ADDRESS
ADD OFFSET
STA TEMP_REGISTER

; Set Data at HEAP + OFFSET to 0x0000
CLA
STA TEMP_REGISTER I ; Target Address

TEMP_REGISTER: HEX 0000 ; Virtual Register Allocation
HEAP_ADDRESS: ADR HEAP
HEAP: HEX 0000
```

## Pointer Offset Function
```asm
; Accumulator = Offset
USE_HEAP: HEX 0000

    ; Since + is reflexive, Pointer + Offset = Offset + Pointer
    ADD HEAP_ADDRESS
    STA HEAP_POINTER

    BUN USE_HEAP I ; Return

; USAGE:
LDA HEAP_OFFSET
BSA USE_HEAP

; Access the Value with 'HEAP_POINTER I'
CLA
STA HEAP_POINTER I
```