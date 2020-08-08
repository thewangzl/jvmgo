package stores

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
	存储指令把变量从操作数栈顶弹出，然后存入局部变量表
	（33条）
	按照所操作变量类型分为6类：
		astore系列指令操作引用类型变量
		dstore系列指令操作double类型变量
		fstore系列指令操作float变量
		istore系列指令操作int变量
		lstore系列指令操作long变量
		xastore指令操作数组
*/
// Store long into local variable

type DSTORE struct {
	base.Index8Instruction
}

func (self *DSTORE) Execute(frame *rtda.Frame){
	_dstore(frame,uint(self.Index))
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_0) Execute(frame *rtda.Frame){
	_dstore(frame,0)
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_1) Execute(frame *rtda.Frame){
	_dstore(frame,1)
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_2) Execute(frame *rtda.Frame){
	_dstore(frame,2)
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_3) Execute(frame *rtda.Frame){
	_dstore(frame,3)
}

func _dstore(frame *rtda.Frame, index uint){
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}


