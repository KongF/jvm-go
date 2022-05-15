package conversions

import "jvm-go/ch05/instructions/base"
import "jvm-go/ch05/rtda"

type D2F struct {
	base.NoOperandsInsruction
}
type D2I struct {
	base.NoOperandsInsruction
}
type D2L struct {
	base.NoOperandsInsruction
}

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}
func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}
func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
