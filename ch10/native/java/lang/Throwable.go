package lang

import (
	"jvm-go/ch10/native"
	"jvm-go/ch10/rtda"
)

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}
func fillInStackTrace(frame *rtda.Frame) {

}
