package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch04/rtda"
)

type LSTORE struct{ base.Index8Instruction }
type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (self *LSTORE)_lstore(frame *rtda.Frame) {
	_lstore(frame, self.Index)
}

func (self *LSTORE_0)_lstore(frame *rtda.Frame) {
	_lstore(frame, 0)
}

func (self *LSTORE_1)_lstore(frame *rtda.Frame) {
	_lstore(frame, 1)
}

func (self *LSTORE_2)_lstore(frame *rtda.Frame) {
	_lstore(frame, 2)
}

func (self *LSTORE_3)_lstore(frame *rtda.Frame) {
	_lstore(frame, 3)
}