package compiler

import "fmt"

var blinkBoilerPlate = `MOVW R13, 0x0
MOVW R1, 0x%04X
MOVT R1, 0x0000
STR R1, R13!, 0x4
LDR R1, R13!, 0x4
MOVW R4, 0x0
MOVT R4, 0x3F20
ADD R2, R4, 0x08
LDR R3, R2
ORR R3, R3, 0x00000008
STR R3, R2
ADD R3, R4, 0x1c
MOVW R2, 0x0000
MOVT R2, 0x0020
STR R2, R3
MOVW R7, 0x0000
MOVT R7, 0x004C
STR R13!, R7, 0x4
BL :DELAY
ADD R3, R4, 0x28
MOVW R2, 0x0000
MOVT R2, 0x0020
STR R2, R3
MOVW R7, 0x0000
MOVT R7, 0x000C
STR R13!, R7, 0x4
BL :DELAY
SUB S R1, R1, 0x01
B {PL} 0xFFFFED
MOVW R7, 0x0000
MOVT R7, 0x000d
STR R13!, R7, 0x4
BL :DELAY
B 0xFFFFE1
:DELAY LDR R13!, R7, 0x4
SUB S R7, R7, 0x01
B {PL} 0xFFFFFD
BX R14`

type blinkCompiler struct {
	input    int
	template string
}

func newBlinkCompiler(input int) (iCompiler, error) {
	if input < 1 {
		return nil, fmt.Errorf("must be a number larger than 0")
	}
	return &blinkCompiler{
		input:    input,
		template: blinkBoilerPlate,
	}, nil
}

func (b *blinkCompiler) ToString() string {
	return fmt.Sprintf(blinkBoilerPlate, b.input)
}
