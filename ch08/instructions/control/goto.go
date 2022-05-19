package control

import "jvm-go/ch08/instructions/base"
import "jvm-go/ch08/rtda"

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
