package comparisons

import "jvm-go/ch05/instructions/base"
import "jvm-go/ch05/rtda"

type DCMPG struct {
	base.NoOperandsInsruction
}
type DCMPL struct {
	base.NoOperandsInsruction
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushDouble(1)
	} else if v1 == v2 {
		stack.PushDouble(0)
	} else if v1 < v2 {
		stack.PushDouble(-1)
	} else if gFlag {
		stack.PushDouble(1)
	} else {
		stack.PushDouble(-1)
	}
}
func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}
func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}
