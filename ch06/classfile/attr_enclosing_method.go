package classfile

type EnclosingMethodAttribute struct {
	cp          ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (self *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.methodIndex = reader.readUint16()
}
func (self *EnclosingMethodAttribute) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}
func (self *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if self.methodIndex > 0 {
		return self.cp.getNameAndType(self.methodIndex)
	} else {
		return "", ""
	}
}
