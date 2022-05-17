package constants

import "jvm-go/ch06/instructions/base"
import "jvm-go/ch06/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
