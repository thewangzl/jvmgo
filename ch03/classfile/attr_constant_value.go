package classfile

/**
ConstantValue_attribute {
	u2 attribute_name_index
	u4 attribute_length			//2
	u2 constantvalue_index		//常量池索引
}
*/

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader){
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}

