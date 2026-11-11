package assembly_compiler

var keywords = map[string]struct{}{
	"AND": {},
	"ADD": {},
	"LDA": {},
	"STA": {},
	"BUN": {},
	"BSA": {},
	"ISZ": {},

	"CLA": {},
	"CLE": {},
	"CMA": {},
	"CME": {},
	"CIR": {},
	"CIL": {},
	"INC": {},
	"SPA": {},
	"SNA": {},
	"SZA": {},
	"SZE": {},
	"HLT": {},

	"INP": {},
	"OUT": {},
	"SKI": {},
	"SKO": {},
	"ION": {},
	"IOF": {},
}

var directives = map[string]struct{}{
	"DEC": {},
	"HEX": {},
	"ADR": {},

	"ORG": {},
	"END": {},
}

// var memoryInstructions = map[string]struct{}{
// 	"AND": {},
// 	"ADD": {},
// 	"LDA": {},
// 	"STA": {},
// 	"BUN": {},
// 	"BSA": {},
// 	"ISZ": {},
// }

var memoryInstructions = map[string]uint16{
	"AND": 0,
	"ADD": 1,
	"LDA": 2,
	"STA": 3,
	"BUN": 4,
	"BSA": 5,
	"ISZ": 6,
}

var registerInstructions = map[string]uint16{
	"CLA": 0x7800,
	"CLE": 0x7400,
	"CMA": 0x7200,
	"CME": 0x7100,
	"CIR": 0x7080,
	"CIL": 0x7040,
	"INC": 0x7020,
	"SPA": 0x7010,
	"SNA": 0x7008,
	"SZA": 0x7004,
	"SZE": 0x7002,
	"HLT": 0x7001,
}

var ioInstructions = map[string]uint16{
	"INP": 0xF800,
	"OUT": 0xF400,
	"SKI": 0xF200,
	"SKO": 0xF100,
	"ION": 0xF080,
	"IOF": 0xF040,
}
