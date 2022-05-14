package math

import "jvm-go/ch05/instructions/base"
import "jvm-go/ch05/rtda"

type IAND struct {
	base.NoOperandsInsruction
}
type LAND struct {
	base.NoOperandsInsruction
}

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}
func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
