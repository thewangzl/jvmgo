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
// Store float into local variable

type FSTORE struct {
	base.Index8Instruction
}

func (self *FSTORE) Execute(frame *rtda.Frame){
	_fstore(frame,uint(self.Index))
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_0) Execute(frame *rtda.Frame){
	_fstore(frame,0)
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_1) Execute(frame *rtda.Frame){
	_fstore(frame,1)
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_2) Execute(frame *rtda.Frame){
	_fstore(frame,2)
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_3) Execute(frame *rtda.Frame){
	_fstore(frame,3)
}

func _fstore(frame *rtda.Frame, index uint){
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}


