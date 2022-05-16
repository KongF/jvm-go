package classfile

type InnerClassesAttribute struct {
	classes []*InnerClassInfo
}
type InnerClassInfo struct {
	innerClassInfoIndex   uint16
	outerClassInfoIndex   uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

func (self *InnerClassesAttribute) readInfo(reader *ClassReader) {
	numberOfClasses := reader.readUint16()
	self.classes = make([]*InnerClassInfo, numberOfClasses)
	for i := range self.classes {
		self.classes[i] = &InnerClassInfo{
			innerClassInfoIndex:   reader.readUint16(),
			outerClassInfoIndex:   reader.readUint16(),
			innerNameIndex:        reader.readUint16(),
			innerClassAccessFlags: reader.readUint16(),
		}
	}
}
