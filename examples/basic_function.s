ORG 0
BUN _Start

Character: HEX 57 ; 'W'

LoadCharacter: HEX 0000
    LDA Character ; Set Return Value
    BUN LoadCharacter I ; Return

PrintCharacter: HEX 0000
    BSA LoadCharacter ; Internal Function Call
    OUT

    BUN PrintCharacter I ; Return

_Start:
    BSA PrintCharacter
    BSA PrintCharacter
    BSA PrintCharacter
    HLT