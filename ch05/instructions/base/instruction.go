package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)

	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {

}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader){
	//
}

//跳转指令
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader){
	self.Offset = int(reader.ReadInt16())
}

/**
存储指令和加载指令需要根据索引存取局部变量表，索引由单字节操作数给出。
*/
//
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader){
	self.Index = uint(reader.ReadInt8())
}

/**
有些指令需要访问运行时常量池，常量池索引由两字节擦作数给出。
*/
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader){
	self.Index = uint(reader.ReadInt16())
}
