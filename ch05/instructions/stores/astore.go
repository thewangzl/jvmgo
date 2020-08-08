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
// Store reference into local variable

type ASTORE struct {
	base.Index8Instruction
}

func (self *ASTORE) Execute(frame *rtda.Frame){
	_astore(frame,uint(self.Index))
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_0) Execute(frame *rtda.Frame){
	_astore(frame,0)
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_1) Execute(frame *rtda.Frame){
	_astore(frame,1)
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_2) Execute(frame *rtda.Frame){
	_astore(frame,2)
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_3) Execute(frame *rtda.Frame){
	_astore(frame,3)
}

func _astore(frame *rtda.Frame, index uint){
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}


