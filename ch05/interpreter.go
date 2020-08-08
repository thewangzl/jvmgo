package main

import "fmt"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/instructions"
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"


func interpret(methodInfo *classfile.MemberInfo){
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	//
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame){
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
/**
	虚拟机解释器逻辑
	do{
		automically calcute pc and fetch opcode at pc;
		if(operands){
			fetch operands;
		}
		execute th action for the opcode;
	} while(there is more to do);

	//java虚拟机伪代码
	for {
		pc := calculatePC()
		opcode := bytecode[pc]
		inst := createInst(opcode)
		inst.fetchOperands(bytecode)
		inst.execute()
	}
*/

//循环执行 "计算PC、解码指令、执行指令"
func loop(thread *rtda.Thread, bytecode []byte){
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//execute
		fmt.Printf("pc:%2d inst:%T %+v\n", pc, inst, inst)
		fmt.Printf("LocalVars:%v  OperandStack:%v\n", frame.LocalVars(),frame.OperandStack())

		inst.Execute(frame)
	}
}