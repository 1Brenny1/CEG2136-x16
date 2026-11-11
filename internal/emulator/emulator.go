package emulator

import "fmt"

const (
	Mask8  uint16 = 0xFF
	Mask12 uint16 = 0xFFF
	Mask16 uint16 = 0xFFFF
)

func Emulate(rom []byte) (*CPU, error) {
	if len(rom)/2 > 4096 {
		return nil, fmt.Errorf("ROM too big: %d/4096", len(rom)/2)
	}

	cpu := &CPU{
		DR: 0x0000,
		AC: 0x0000,
		IR: 0x0000,
		TR: 0x0000,

		AR: 0x000,
		PC: 0x000,

		INPR: 0x00,
		OUTR: 0x00,

		E: false,
		I: false,
		R: false,

		FGI: false,
		FGO: true,

		IEN: false,

		Memory: make([]uint16, 4096),

		Running: true,
	}

	for i, b := range rom {
		if i%2 == 0 {
			cpu.Memory[i/2] = uint16(b) << 8
		} else {
			cpu.Memory[i/2] |= uint16(b)
		}
	}

	for cpu.Running {
		cpu.Cycle()
	}

	return cpu, nil
}
