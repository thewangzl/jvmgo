package control

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
	控制指令
		jsr、ret用于实现finally子句(java6之前)
		return （6条）
		goto	无条件跳转
		tableswitch		switch-case语句，case值可以编码成一个索引表
		lookupswich							  不可以编码成一个索引表
	（11条）
*/
//Access jump table index and jump
type TABLE_SWITCH struct {
	defaultOffset 		int32
	low					int32
	high				int32
	jumpOffsets			[]int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader){
	reader.SkipPadding()					//0-3字节的padding,以保证defaultOffset在字节码中的地址是4的倍数
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame){
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index - self.low])
	}else{
		offset = int(self.jumpOffsets[self.defaultOffset])
	}
	base.Branch(frame, offset)
}