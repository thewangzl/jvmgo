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
// Store int into local variable

type ISTORE struct {
	base.Index8Instruction
}

func (self *ISTORE) Execute(frame *rtda.Frame){
	_istore(frame,uint(self.Index))
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_0) Execute(frame *rtda.Frame){
	_istore(frame,0)
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_1) Execute(frame *rtda.Frame){
	_istore(frame,1)
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_2) Execute(frame *rtda.Frame){
	_istore(frame,2)
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_3) Execute(frame *rtda.Frame){
	_istore(frame,3)
}

func _istore(frame *rtda.Frame, index uint){
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}


