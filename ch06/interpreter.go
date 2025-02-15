package main

import "fmt"
import "jvmgo/ch06/instructions"
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"


func interpret(method *heap.Method){
	//
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame){
	if r := recover(); r != nil {
		// fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		// fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		// panic(r)
	}
}

//循环执行 "计算PC、解码指令、执行指令"
func loop(thread *rtda.Thread, bytecode []byte){
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)
		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)

		inst.Execute(frame)
	}
}