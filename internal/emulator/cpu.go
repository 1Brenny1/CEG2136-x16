package emulator

import (
	"fmt"
)

type CPU struct {
	// 16Bit Registers
	DR uint16
	AC uint16
	IR uint16
	TR uint16

	// 12Bit Registers
	AR uint16
	PC uint16

	// 8Bit Registers
	INPR uint8
	OUTR uint8

	E bool // Overflow Error
	I bool // Indirect Addressing
	R bool // Interupt Flag

	D [8]bool // Decoder Output

	FGI bool
	FGO bool

	IEN bool // Interupt Enable

	Memory []uint16

	Running bool
}

func (cpu *CPU) Cycle() {
	//time.Sleep(time.Millisecond * 500)

	// Interrupt Recognition
	if cpu.IEN && (cpu.FGI || cpu.FGO) {
		cpu.R = true
	}

	if cpu.R {
		cpu.Interupt()
	} else {
		cpu.Fetch()
		cpu.Decode()
		cpu.Indirect()
		cpu.Execute()
	}

	if !cpu.FGO {
		fmt.Print(string(cpu.OUTR))
		cpu.FGO = true
	}

	//fmt.Printf("DR: %x\nAC: %x\nIR: %x\nTR: %x\n\nAR: %x\nPC: %x\n=========\n", cpu.DR, cpu.AC, cpu.IR, cpu.TR, cpu.AR, cpu.PC)
}

func (cpu *CPU) Fetch() {
	// R'T0
	cpu.AR = cpu.PC & Mask12
	// R'T1
	cpu.IR = cpu.Memory[cpu.AR]
	cpu.PC = (cpu.PC + 0x001) & Mask12
}

func (cpu *CPU) Decode() { // R'T2
	opcode := (cpu.IR >> 12) & 0x007

	cpu.D = [8]bool{}
	cpu.D[opcode] = true

	cpu.AR = cpu.IR & Mask12
	cpu.I = cpu.IR&0x8000 != 0x0000
}

func (cpu *CPU) Indirect() {
	if !cpu.D[7] && cpu.I { // D7'IT3
		cpu.AR = cpu.Memory[cpu.AR] & Mask12
	}
}

func (cpu *CPU) Interupt() {
	//cpu.R = true

	cpu.AR = 0x000
	cpu.TR = cpu.PC

	cpu.Memory[cpu.AR] = cpu.TR
	cpu.PC = 0x000

	cpu.PC = (cpu.PC + 0x001) & Mask12
	cpu.IEN = false
	cpu.R = false
}

func (cpu *CPU) Execute() {
	if cpu.D[7] {
		if cpu.I {
			cpu.ExecuteIOInstructions()
			return
		}

		cpu.ExecuteRegisterReferenceInstructions()
		return
	}

	cpu.ExecuteMemoryReferenceInstructions()
}

func (cpu *CPU) ExecuteMemoryReferenceInstructions() {
	switch {
	case cpu.D[0]: // AND
		cpu.DR = cpu.Memory[cpu.AR]
		cpu.AC &= cpu.DR

	case cpu.D[1]: // ADD
		cpu.DR = cpu.Memory[cpu.AR]

		result := uint32(cpu.AC) + uint32(cpu.DR)

		cpu.AC = uint16(result)
		cpu.E = result > 0xFFFF

	case cpu.D[2]: // LDA
		cpu.DR = cpu.Memory[cpu.AR]
		cpu.AC = cpu.DR

	case cpu.D[3]: // STA
		cpu.Memory[cpu.AR] = cpu.AC

	case cpu.D[4]: // BUN
		cpu.PC = cpu.AR

	case cpu.D[5]: // BSA
		cpu.Memory[cpu.AR] = cpu.PC
		cpu.AR += 0x001

		cpu.PC = cpu.AR

	case cpu.D[6]: // ISZ
		cpu.DR = cpu.Memory[cpu.AR]
		cpu.DR += 0x0001

		cpu.Memory[cpu.AR] = cpu.DR
		if cpu.DR == 0 {
			cpu.PC = cpu.PC + 0x001
		}
	}
}

func (cpu *CPU) ExecuteRegisterReferenceInstructions() {
	switch cpu.IR {
	case INSTR_CLA:
		cpu.AC = 0x0000

	case INSTR_CLE:
		cpu.E = false

	case INSTR_CMA:
		cpu.AC = ^cpu.AC

	case INSTR_CME:
		cpu.E = !cpu.E

	case INSTR_CIR:
		e := cpu.E
		cpu.E = cpu.AC&0x0001 != 0x0000
		cpu.AC >>= 0x0001
		if e {
			cpu.AC |= 0x8000
		}

	case INSTR_CIL:
		e := cpu.E
		cpu.E = cpu.AC&0x8000 != 0x0000
		cpu.AC <<= 0x0001
		if e {
			cpu.AC |= 0x0001
		}

	case INSTR_INC:
		cpu.AC += 0x0001

	case INSTR_SPA:
		if cpu.AC&0x8000 == 0x0000 {
			cpu.PC = (cpu.PC + 0x001) & Mask12
		}

	case INSTR_SNA:
		if cpu.AC&0x8000 == 0x0001 {
			cpu.PC = (cpu.PC + 0x001) & Mask12
		}

	case INSTR_SZA:
		if cpu.AC == 0x0000 {
			cpu.PC = (cpu.PC + 0x001) & Mask12
		}

	case INSTR_SZE:
		if !cpu.E {
			cpu.PC = (cpu.PC + 0x001) & Mask12
		}

	case INSTR_HLT:
		cpu.Running = false
	}
}

func (cpu *CPU) ExecuteIOInstructions() {
	switch cpu.IR {
	case INSTR_INP:
		cpu.AC = (cpu.AC & Mask8) | uint16(cpu.INPR)
		cpu.FGI = false

	case INSTR_OUT:
		cpu.OUTR = uint8(cpu.AC & 0x00FF)
		cpu.FGO = false

	case INSTR_SKI:
		if cpu.FGI {
			cpu.PC = (cpu.PC + 0x001) & Mask12
		}

	case INSTR_SKO:
		if cpu.FGO {
			cpu.PC = (cpu.PC + 0x001) & Mask12
		}

	case INSTR_ION:
		cpu.IEN = true

	case INSTR_IOF:
		cpu.IEN = false
	}
}
