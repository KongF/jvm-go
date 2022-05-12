package classfile

import "math"

/*
CONSTANT_Integer_info/CONSTANT_Float_info {
	u1 tag;
	u4 bytes;
}*/

type ConstantIntegerInfo struct {
	val int32
}
type ConstantFloatInfo struct {
	val float32
}
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}
