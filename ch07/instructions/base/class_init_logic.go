package base

import (
	"jvm-go/ch07/rtda"
	"jvm-go/ch07/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	secheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func secheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}
