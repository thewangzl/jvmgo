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
//Branch always
type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame){
	base.Branch(frame, self.Offset)
}