package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type DNEG struct{ base.NoOperandsInstruction }
type FNEG struct{ base.NoOperandsInstruction }
type INEG struct{ base.NoOperandsInstruction }
type LNEG struct{ base.NoOperandsInstruction }

func (self *INEG) Execute(frame *rtda.Frame) {
	//todo
	panic("Not Implemented yet")
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	//todo
	panic("Not Implemented yet")
}

func (self *DNEG) Execute(frame *rtda.Frame) {
	//todo
	panic("Not Implemented yet")
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	//todo
	panic("Not Implemented yet")
}
