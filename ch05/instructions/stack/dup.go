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
//Duplicate the top operand stack value
type DUP struct {
	base.NoOperandsInstruction
}

func (self *DUP) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

//Duplicate the top operand stack value and insert two values down
type DUP_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X1) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

//Duplicate the top operand stack value and insert two or three values down
type DUP_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X2) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

//Duplicate the top one or two operand stack values
type DUP2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top one or two operand stack values and insert two or three values down
type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X1) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

// Duplicate the top one or two operand stack values and insert two, three, or four values down
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X2) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}