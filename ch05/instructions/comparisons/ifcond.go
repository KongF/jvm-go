package comparisons

import "jvm-go/ch05/instructions/base"
import "jvm-go/ch05/rtda"

/*
ifeq: x==0
ifne: x!=0
iflt: x<0
ifle: x<=0
ifgt: x>0
ifge:x>=0
*/
type IFEQ struct {
	base.BranchInstruction
}
type IFNE struct {
	base.BranchInstruction
}
type IFLT struct {
	base.BranchInstruction
}
type IFLE struct {
	base.BranchInstruction
}
type IFGT struct {
	base.BranchInstruction
}
type IFGE struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}
