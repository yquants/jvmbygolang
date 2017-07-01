package base

type BytecodeReader struct {
	code []byte
	pc   int //index of current byte reader position
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc ++
	return i
}

func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := self.ReadUint8()
	byte2 := self.ReadUint8()
	return (byte1 << 8) | byte2
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := self.ReadUint8()
	byte2 := self.ReadUint8()
	byte3 := self.ReadUint8()
	byte4 := self.ReadUint8()
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

func (self *BytecodeReader) SkipPadding() {
	for self.pc % 4 != 0 {
		self.ReadUint8()
	}
}

func (self *BytecodeReader) PC() int {
	return self.pc
}
