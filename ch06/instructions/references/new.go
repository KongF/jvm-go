package references

import "jvm-go/ch06/instructions/base"
import "jvm-go/ch06/rtda"
import "jvm-go/ch06/rtda/heap"

type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).((heap.ClassRef))
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
