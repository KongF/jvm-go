package lang

import "jvm-go/ch09/native"
import "jvm-go/ch09/rtda"

const jlObject = "java/lang/Object"

func init() {
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	//native.Register(jlObject, "hashCode", "()I", hashCode)
	//native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
