package math

import (
	"JVM/instructions/base"
	"JVM/rtda"
	"math"
)

// 求余

type IREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

type FREM struct {
	base.NoOperandsInstruction
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}

type LREM struct {
	base.NoOperandsInstruction
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}

type DREM struct {
	base.NoOperandsInstruction
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
