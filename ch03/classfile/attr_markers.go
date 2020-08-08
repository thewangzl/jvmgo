package classfile 

/**
仅做标记,不包含任何数据，所以长度为0
Deprecated_attribute{
	u2	attribute_name_index
	u4	attribute_length		// 0
}
*/
type DeprecatedAttribute struct{
	MarkerAttribute
}

/**
仅做标记,不包含任何数据，所以长度为0
Synthetic_attribute{
	u2	attribute_name_index
	u4	attribute_length		// 0
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}


type MarkerAttribute struct {

}

func (self * MarkerAttribute) readInfo(reader *ClassReader){

}