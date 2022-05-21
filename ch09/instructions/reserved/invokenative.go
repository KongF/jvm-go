package reserved

import "jvm-go/ch09/instructions/base"
import "jvm-go/ch09/rtda"
import "jvm-go/ch09/native"
import _ "jvm-go/ch09/native/java/lang"

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnstatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}
