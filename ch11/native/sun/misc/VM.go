package misc

import (
	"jvm-go/ch11/instructions/base"
	"jvm-go/ch11/native"
	"jvm-go/ch11/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}
func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
