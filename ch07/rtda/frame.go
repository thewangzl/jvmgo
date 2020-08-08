package rtda

import "fmt"
import "jvmgo/ch07/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars     //局部变量表指针
	operandStack *OperandStack //操作数栈指针
	thread       *Thread       //为了实现跳转
	method       *heap.Method
	nextPC       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) NextPC() int {
	fmt.Printf("NextPC: %s.%s %d\n", self.method.Class().Name(), self.method.Name(), self.nextPC)
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	fmt.Printf("SetNextPC: %s.%s %d -> %d\n", self.method.Class().Name(), self.method.Name(), self.nextPC, nextPC)
	self.nextPC = nextPC
}

func (self *Frame) RevertNextPC() {
	fmt.Printf("RevertNextPC: %s.%s %d -> %d\n", self.method.Class().Name(), self.method.Name(), self.nextPC, self.thread.pc)
	self.nextPC = self.thread.pc
}
