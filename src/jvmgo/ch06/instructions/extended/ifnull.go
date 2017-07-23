package extended

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}
type IFNONNLL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame)  {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNLL) Execute(frame *rtda.Frame)  {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}

