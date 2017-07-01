package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IXOR struct{ base.NoOperandsInstruction }
type LXOR struct{ base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	result := ^v1 + 1
	stack.PushInt(result)
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	result := ^v1 + 1
	stack.PushLong(result)
}
