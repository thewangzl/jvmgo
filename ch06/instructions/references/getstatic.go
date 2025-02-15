package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

//Get static field from class
type GET_STATIC struct {
	base.Index16Instruction
}

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	/**一个操作数：uint16索引，来自字节码，通过该索引从当前类的运行时常量池中找到
	一个字段符号索引，解析这个符号引用就可以知道要取类的哪个静态变量的值。
	*/
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.
		ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		//todo
	}
}
