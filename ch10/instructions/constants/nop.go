package constants

import "jvm-go/ch10/instructions/base"
import "jvm-go/ch10/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
