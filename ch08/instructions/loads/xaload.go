package loads

import "jvm-go/ch08/instructions/base"
import "jvm-go/ch08/rtda"
import "jvm-go/ch08/rtda/heap"

type AALOAD struct {
	base.NoOperandsInstruction
}
type BALOAD struct {
	base.NoOperandsInstruction
}
type CALOAD struct {
	base.NoOperandsInstruction
}
type DALOAD struct {
	base.NoOperandsInstruction
}
type FALOAD struct {
	base.NoOperandsInstruction
}
type IALOAD struct {
	base.NoOperandsInstruction
}
type LALOAD struct {
	base.NoOperandsInstruction
}
type SALOAD struct {
	base.NoOperandsInstruction
}

func (self *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
