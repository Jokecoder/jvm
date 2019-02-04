package math

import (
	"JVM/instructions/base"
	"JVM/rtda"
)

// 局部变量表中int变量加常量

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadUint8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localvars := frame.LocalVars()
	val := localvars.GetInt(self.Index)
	val += self.Const
	localvars.SetInt(self.Index, val)
}
