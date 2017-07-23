package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}
type ARETURN struct {
	base.NoOperandsInstruction
}
type IRETURN struct {
	base.NoOperandsInstruction
}
type LRETURN struct {
	base.NoOperandsInstruction
}
type DRETURN struct {
	base.NoOperandsInstruction
}
type FRETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	//hack
	//frame.Thread().PopFrame()
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.PopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokeFrame.OperandStack().PushInt(retVal)
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.PopFrame()
	retVal := currentFrame.OperandStack().PopFloat()
	invokeFrame.OperandStack().PushFloat(retVal)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.PopFrame()
	retVal := currentFrame.OperandStack().PopDouble()
	invokeFrame.OperandStack().PushDouble(retVal)
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.PopFrame()
	retVal := currentFrame.OperandStack().PopLong()
	invokeFrame.OperandStack().PushLong(retVal)
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.PopFrame()
	retVal := currentFrame.OperandStack().PopRef()
	invokeFrame.OperandStack().PushRef(retVal)
}

