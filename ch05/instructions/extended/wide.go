package extended

import (
	"jvm-go/ch05/instructions/base"
	"jvm-go/ch05/instructions/loads"
	"jvm-go/ch05/instructions/math"
	"jvm-go/ch05/instructions/stores"
	"jvm-go/ch05/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: //iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x16: //lload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x17: //fload
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x18: //dload
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x19: //aload
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x36: //istore
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x37: //lstore
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x38: //fstore
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x39: //dstore
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x3a: //astore
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x84: //iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0xa9: //ret
		panic("Unsupported opcode :0xa9!")
	}
}
func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
