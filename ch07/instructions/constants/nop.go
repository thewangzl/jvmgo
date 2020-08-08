package constants

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

//do nothing
type NOP struct {
	base.NoOperandsInstruction
}


func (self *NOP) Execute(frame *rtda.Frame){

}