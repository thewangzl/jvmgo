package stack

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
	栈指令直接对操作数进行操作
	（9条）
	pop和pop2指令将栈顶变量弹出 (pop只能弹出int、float等占用一个操作数栈位置的变量，
								double和long变量在操作数栈中占两个位置，需要用pop2)
	dup系列指令复制栈顶变量 （6条)
	swap指令交换栈顶两个变量
*/
//Pop the top operand stack value
type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PopSlot()
}

//Pop the top one or two operand stack values
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}