package references

import (
	"jvm-go/ch10/instructions/base"
)
import "jvm-go/ch10/rtda"

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()

	if arrRef == nil {
		panic("java.lang.NullPointerExecption")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
