package control

import "jvm-go/ch11/instructions/base"
import "jvm-go/ch11/rtda"

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
