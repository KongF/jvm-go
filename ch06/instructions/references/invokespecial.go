package references

import "jvm-go/ch06/instructions/base"
import "jvm-go/ch06/rtda"

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
