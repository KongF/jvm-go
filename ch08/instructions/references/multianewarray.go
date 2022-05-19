package references

import "jvm-go/ch08/instructions/base"
import "jvm-go/ch08/rtda"
import "jvm-go/ch08/rtda/heap"

type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}
