package lang

import "unsafe"
import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
	native.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)
}

// public final native Class<?> getClass();
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	frame.OperandStack().PushRef(this.Clone())
}
