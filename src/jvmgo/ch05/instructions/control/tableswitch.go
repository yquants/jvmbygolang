package control

import "jvmgo/ch05/instructions/base"

type TABLE_SWITCH struct {
	defaultOffset	int32
	low		int32
	high		int32
	jumpOffsets	[]int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader)  {
	
}
