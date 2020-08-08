package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
	数学指令包括：算数运算(加法(add)、减法(sub)、乘法(mul)、除法(div)、求余(rem)、取反(neg))、
				位移运算(左移(ISHL、LSHL)、右移(算术右移/有符号右移(ISHR、LSHR)、逻辑右移/无符号右移(IUSHR、LUSHR)))
				布尔运算(按位与(and)、按位或(or)、按位异或(xor))(只能操作int和long变量)
	（37条）

*/
//Add double 
type DADD struct {
	base.NoOperandsInstruction
}

func (self *DADD) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

//Add float 
type FADD struct {
	base.NoOperandsInstruction
}

func (self *FADD) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}


//Add int 
type IADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

//Add long 
type LADD struct {
	base.NoOperandsInstruction
}

func (self *LADD) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}