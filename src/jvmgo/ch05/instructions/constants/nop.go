package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch04/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame)  {
	// nothing to do
}
