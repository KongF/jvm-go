package control

import "jvm-go/ch07//instructions/base"
import "jvm-go/ch07//rtda"

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offSet int
	if index >= self.low && index <= self.high {
		offSet = int(self.jumpOffsets[index-self.low])
	} else {
		offSet = int(self.defaultOffset)
	}
	base.Branch(frame, offSet)
}
