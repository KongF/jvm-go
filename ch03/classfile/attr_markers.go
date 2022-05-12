package classfile

type DeprecateAttribute struct {
	MarkerAttribute
}
type SynthericAttribute struct {
	MarkerAttribute
}
type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	//只用作标记不读取数据
}
