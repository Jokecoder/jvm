package control

import (
	"JVM/instructions/base"
	"JVM/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction // return void from method
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct {
	base.NoOperandsInstruction // return reference from method
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

type DRETURN struct {
	base.NoOperandsInstruction // return double from method
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(ref)
}

type FRETURN struct {
	base.NoOperandsInstruction // return float from method
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(ref)
}

type IRETURN struct {
	base.NoOperandsInstruction // return int from method
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(ref)
}

type LRETURN struct {
	base.NoOperandsInstruction // return long from method
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(ref)
}
