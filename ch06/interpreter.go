package main

import (
	"fmt"
	"jvm-go/ch06/rtda/heap"
)
import "jvm-go/ch06/instructions"
import "jvm-go/ch06/instructions/base"
import "jvm-go/ch06/rtda"

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		//panic(r)
	}
}
func loop(thread *rtda.Thread, bytecode []byte) {
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
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
