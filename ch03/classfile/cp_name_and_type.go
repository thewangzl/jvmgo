package classfile

/**
Constant_NameAndType_info {
	u1	tag
	u2	name_index
	u2	descriptor_index
}
*/
/**
	字段或方法的名称和描述符
	1)类型描述符
		a.基本类型 byte、short、char、int、long、float、double的描述符是首个字母：B、S、C、I、J、F、D
		b.引用类型: L+类的完全限定名+分号
		c.数组：[ + 数组元素类型描述符
	2)字段描述符：字段类型的描述符
	3)方法描述符：(分号分割的参数类型描述符)+返回值类型描述符
*/
type ConstantNameAndTypeInfo struct {
	nameIndex uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}