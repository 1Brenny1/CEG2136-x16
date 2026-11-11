.PHONY: asm emu

asm:
	go build -o ./build/x16asm.exe ./cmds/x16asm/main.go

tinyc:
	go build -o ./build/tinyc.exe ./cmds/tinyc/main.go

emu:
	go build -o ./build/x16emu.exe ./cmds/x16emu/main.go