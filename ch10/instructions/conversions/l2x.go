package conversions

import "jvm-go/ch10/instructions/base"
import "jvm-go/ch10/rtda"

// Convert long to double
type L2D struct{ base.NoOperandsInstruction }

// Convert long to float
type L2F struct{ base.NoOperandsInstruction }

// Convert long to int
type L2I struct{ base.NoOperandsInstruction }

func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}
func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}
func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
