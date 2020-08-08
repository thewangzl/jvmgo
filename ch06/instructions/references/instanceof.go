package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

//判断对象是否是某个类的实例（或者对象的类是否实现某个接口）
type INSTANCE_OF struct {
	base.Index16Instruction
}

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	/**
	两个操作数，第一个操作数是uint16索引，从方法的字节码中获取，
	通过这个索引可以从当前类的运行时常量池中找到一个类符号引用。
	第二个操作数是对象引用，从操作数栈中弹出。
	*/
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
