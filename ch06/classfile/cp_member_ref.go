package classfile


/**
Constant_Fieldref_Info {
	u1 tag
	u2 class_index
	u2 name_and_type_index
}
*/
type ConstantFieldrefInfo struct{
	ConstantMemberrefInfo
}
/**
Constant_Methodref_Info {
	u1 tag
	u2 class_index
	u2 name_and_type_index
}
*/
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

/**
Constant_InterfaceMethodref_Info {
	u1 tag
	u2 class_index
	u2 name_and_type_index
}
*/
type ConstantInterfaceMethodrefInfo struct{
	ConstantMemberrefInfo
}

type ConstantMemberrefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader){
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}


func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

