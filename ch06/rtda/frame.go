package rtda

import "jvmgo/ch06/rtda/heap"
type Frame struct {
	lower *Frame
	localVars LocalVars 		//局部变量表指针
	operandStack *OperandStack  //操作数栈指针
	thread *Thread				//为了实现跳转
	method *heap.Method
	nextPC 	int
}

func newFrame(thread *Thread, method *heap.Method) *Frame{
	return &Frame{
		thread : thread,
		method : method,
		localVars : newLocalVars(method.MaxLocals()),
		operandStack : newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) LocalVars() LocalVars{
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack{
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

