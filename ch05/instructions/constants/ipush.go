package constants

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
	常量指令把常数推入操作数栈，
	常量来自三个地方：隐含在操作码中、操作数、运行时常量池
*/
/**
	bipush指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶
	sipush指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶
*/
//Push byte
type BIPUSH struct {
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

//Push short
type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}