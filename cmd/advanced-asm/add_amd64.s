// This file contains the Go assembly code for the Add function.

// TEXT directive defines a text (code) section.
// The first argument is the function name, the second is the ABI (Application Binary Interface), and the third is the frame size.
TEXT ·Add(SB), $0-24
	// Load the arguments from the stack into registers.
	MOVQ a+0(FP), AX
	MOVQ b+8(FP), BX
	// Add the values in the registers.
	ADDQ BX, AX
	// Store the result in the return value location.
	MOVQ AX, ret+16(FP)
	// Return from the function.
	RET
