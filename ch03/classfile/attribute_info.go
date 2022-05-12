package classfile

/*
attribute_info{
	u2	attribute_name_index;
	u4	attribute_length;
	u1	info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttrinbutes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttrinbute(reader, cp)
	}
	return attributes
}
func readAttrinbute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttrinbuteInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}
func newAttrinbuteInfo(attrName string, attLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{}
	case "Syntheic":
		return &SyntheicAttribute{}
	default:
		return &UnparsedAttribute{attrName, attLen, nil}

	}
}
