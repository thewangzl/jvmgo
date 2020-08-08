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
//Load double from local variable

type DLOAD struct {
	base.Index8Instruction
}

func (self *DLOAD) Execute(frame *rtda.Frame){
	_dload(frame, uint(self.Index))
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_0) Execute(frame *rtda.Frame){
	_dload(frame, 0)
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_1) Execute(frame *rtda.Frame){
	_dload(frame, 1)
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_2) Execute(frame *rtda.Frame){
	_dload(frame, 2)
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_3) Execute(frame *rtda.Frame){
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint){
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}