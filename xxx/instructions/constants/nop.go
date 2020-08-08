package constants

import "jvmgo/xxx/instructions/base"
import "jvmgo/xxx/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
