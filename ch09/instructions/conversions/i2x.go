package conversions

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"

// convert int to byte
type I2B struct {
	base.NoOperandsInstruction
}

func (self *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}

// convert int to char
type I2C struct {
	base.NoOperandsInstruction
}

func (self *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	c := int32(uint64(i))
	stack.PushInt(c)
}

//convert int to short
type I2S struct {
	base.NoOperandsInstruction
}

func (self *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}

//convert int to long
type I2L struct {
	base.NoOperandsInstruction
}

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

//convert int to float
type I2F struct {
	base.NoOperandsInstruction
}

func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

//convert int to double
type I2D struct {
	base.NoOperandsInstruction
}

func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}