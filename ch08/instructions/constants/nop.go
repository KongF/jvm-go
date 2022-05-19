package constants

import "jvm-go/ch08/instructions/base"
import "jvm-go/ch08/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
