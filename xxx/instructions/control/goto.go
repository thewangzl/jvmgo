package control

import "jvmgo/xxx/instructions/base"
import "jvmgo/xxx/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
