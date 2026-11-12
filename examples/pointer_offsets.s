ORG 0
BUN _Start

; Since the only way to access an arbritrary value in assembly is through a direct reference in memory
; we need to create our own reference by using a virtual register.
VIRT_REG: HEX 0000

OFFSET: DEC 3

POINTER: ADR POINTER ; Pointer to this instructions address
    HEX 0000
    HEX 0000
    HEX 0057
    HEX 0000
    HEX 0000

_Start:
    LDA POINTER    ; Load Pointer
    ADD OFFSET     ; Add Offset
    STA VIRT_REG   ; Save to Virtual Register

    LDA VIRT_REG I ; Load Value at Pointer + Offset
    OUT
    HLT