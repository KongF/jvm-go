package stack

import "jvm-go/ch11/instructions/base"
import "jvm-go/ch11/rtda"

type POP struct {
	base.NoOperandsInstruction
}
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}