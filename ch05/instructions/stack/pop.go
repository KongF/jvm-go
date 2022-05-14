package stack

import "jvm-go/ch05/instructions/base"
import "jvm-go/ch05/rtda"

type POP struct {
	base.NoOperandsInsruction
}
type POP2 struct {
	base.NoOperandsInsruction
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
