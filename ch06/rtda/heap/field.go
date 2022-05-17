package heap

import "jvm-go/ch06/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) copyAttributes(field *classfile.MemberInfo) {
	if valAttr := field.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}
func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}
func (self *Field) SlotId() uint {
	return self.slotId
}
