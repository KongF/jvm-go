package base

import "jvm-go/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}
type NoOperandsInstruction struct {
}

//跳转指令
type BranchInstruction struct {
	Offset int //存放跳转偏移量
}
type Index8Instruction struct {
	Index uint // 局部变量表索引
}
type Index16Instruction struct {
	Index uint
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do 表示没有操作数指令
}
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
