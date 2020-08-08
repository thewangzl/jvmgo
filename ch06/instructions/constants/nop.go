package constants

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

//do nothing
type NOP struct {
	base.NoOperandsInstruction
}


func (self *NOP) Execute(frame *rtda.Frame){

}