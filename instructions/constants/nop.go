package constants

import (
	"JVM/instructions/base"
	"JVM/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
