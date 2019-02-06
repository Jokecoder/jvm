package references

import (
	"JVM/instructions/base"
	"JVM/rtda"
)

// hack 调用构造函数初始化对象
type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
