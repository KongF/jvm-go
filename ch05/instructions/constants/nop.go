package constants

import "jvm-go/ch05/instructions/base"
import "jvm-go/ch05/rtda"

type NOP struct {
	base.NoOperandsInsruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
