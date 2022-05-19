package stores

import "jvm-go/ch08/instructions/base"
import "jvm-go/ch08/rtda"
import "jvm-go/ch08/rtda/heap"

type AASTORE struct {
	base.NoOperandsInstruction
}
type BASTORE struct {
	base.NoOperandsInstruction
}
type CASTORE struct {
	base.NoOperandsInstruction
}
type DASTORE struct {
	base.NoOperandsInstruction
}
type FASTORE struct {
	base.NoOperandsInstruction
}
type IASTORE struct {
	base.NoOperandsInstruction
}
type LASTORE struct {
	base.NoOperandsInstruction
}
type SASTORE struct {
	base.NoOperandsInstruction
}

func (self *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}
func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
