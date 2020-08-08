package rtda

type Frame struct {
	lower *Frame
	localVars LocalVars 		//局部变量表指针
	operandStack *OperandStack  //操作数栈指针
}

func NewFrame(maxLocals uint , maxStack uint) *Frame{
	return &Frame{
		localVars : newLocalVars(maxLocals),
		operandStack : newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars{
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack{
	return self.operandStack
}