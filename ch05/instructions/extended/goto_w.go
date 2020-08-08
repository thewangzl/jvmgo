package extended


import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
	扩展指令
	（6条）
	jsr_w	不讨论
	multianewarray	创建多维数组
*/
//Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}

