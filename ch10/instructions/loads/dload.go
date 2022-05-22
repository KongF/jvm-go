package loads

import "jvm-go/ch10/instructions/base"
import "jvm-go/ch10/rtda"

type DLOAD struct{ base.Index8Instruction }
type DLOAD_0 struct{ base.NoOperandsInstruction }
type DLOAD_1 struct{ base.NoOperandsInstruction }
type DLOAD_2 struct{ base.NoOperandsInstruction }
type DLOAD_3 struct{ base.NoOperandsInstruction }

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, self.Index)
}
func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}
func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}
func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}
func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
