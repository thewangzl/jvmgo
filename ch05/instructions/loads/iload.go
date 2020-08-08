package loads

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
	从局部变量表获取变量，然后推入操作数栈顶
	（33条）
	按照所操作变量类型分为6类：
		aload系列指令操作引用类型变量
		dload系列指令操作double类型变量
		fload系列指令操作float变量
		iload系列指令操作int变量
		lload系列指令操作long变量
		xaload指令操作数组
*/
//Load int from local variable

type ILOAD struct {
	base.Index8Instruction
}

func (self *ILOAD) Execute(frame *rtda.Frame){
	_iload(frame, uint(self.Index))
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_0) Execute(frame *rtda.Frame){
	_iload(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_1) Execute(frame *rtda.Frame){
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_2) Execute(frame *rtda.Frame){
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_3) Execute(frame *rtda.Frame){
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint){
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}