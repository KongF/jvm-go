package reserved

import "jvm-go/ch11/instructions/base"
import "jvm-go/ch11/rtda"
import "jvm-go/ch11/native"
import _ "jvm-go/ch11/native/java/io"
import _ "jvm-go/ch11/native/java/lang"
import _ "jvm-go/ch11/native/java/security"
import _ "jvm-go/ch11/native/java/util/concurrent/atomic"
import _ "jvm-go/ch11/native/sun/io"
import _ "jvm-go/ch11/native/sun/misc"
import _ "jvm-go/ch11/native/sun/reflect"

type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
