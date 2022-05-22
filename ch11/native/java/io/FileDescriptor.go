package io

import (
	"jvm-go/ch11/native"
	"jvm-go/ch11/rtda"
)

const fd = "java/io/FileDescriptor"

func init() {
	native.Register(fd, "set", "(I)J", set)
}
func set(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
