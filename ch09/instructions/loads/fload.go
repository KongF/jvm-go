package loads

import "jvm-go/ch09/instructions/base"
import "jvm-go/ch09/rtda"

type FLOAD struct{ base.Index8Instruction }
type FLOAD_0 struct{ base.NoOperandsInstruction }
type FLOAD_1 struct{ base.NoOperandsInstruction }
type FLOAD_2 struct{ base.NoOperandsInstruction }
type FLOAD_3 struct{ base.NoOperandsInstruction }

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, self.Index)
}
func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}
func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}
func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}
func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
