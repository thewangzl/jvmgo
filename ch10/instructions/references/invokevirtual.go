package references

import "fmt"
import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//fmt.Printf("%s.%s - %d\n", resolvedMethod.Class().Name(), resolvedMethod.Name(), resolvedMethod.ArgSlotCount())
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		// hack System.out.println()
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}
	//fmt.Println(ref.Class().Name())

	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) && resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() && ref.Class() != currentClass && !ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

// hack!
func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jStr := stack.PopRef()
		goStr := heap.GoString(jStr)
		fmt.Println(goStr)
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}
