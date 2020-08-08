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
//Load float from local variable

type FLOAD struct {
	base.Index8Instruction
}

func (self *FLOAD) Execute(frame *rtda.Frame){
	_fload(frame, uint(self.Index))
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_0) Execute(frame *rtda.Frame){
	_fload(frame, 0)
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_1) Execute(frame *rtda.Frame){
	_fload(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_2) Execute(frame *rtda.Frame){
	_fload(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_3) Execute(frame *rtda.Frame){
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint){
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}