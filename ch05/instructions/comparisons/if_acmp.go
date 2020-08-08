package comparisons

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

//Branch if reference comparison succeeds

type  IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, self.Offset)
	}
}


type  IF_ACMPNE struct {
	base.BranchInstruction
}


func (self *IF_ACMPNE) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		base.Branch(frame, self.Offset)
	}
}