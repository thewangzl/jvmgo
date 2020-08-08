package conversions

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
	基本类型强制转换
	(15条)
		i2x系列指令把int变量强制转换为其他类型
		l2x			long
		f2x			float
		d2x			double
*/
// convert double to float
type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// convert double to int
type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

//convert double to long
type D2L struct {
	base.NoOperandsInstruction
}

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}