package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

type GET_FIELD struct {
	base.Index16Instruction
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	/**三个操作数：
	第一个操作数 uint16索引，来自字节码，通过该索引从当前类的运行时常量池中找到
			一个字段符号索引，解析这个符号引用就可以知道要取类的哪个变量的值。
	第三个操作数是对象引用，从操作数栈中弹出。
	*/
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncomtibleCassChangeError")
	}
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()

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
