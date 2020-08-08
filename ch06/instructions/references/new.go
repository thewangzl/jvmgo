package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// Create new object
type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame){
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantitionError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}