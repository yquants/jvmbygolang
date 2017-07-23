package rtda

import (
	"jvmgo/ch06/rtda/heap"
	"fmt"
)

type Thread struct {
	pc	int
	stack	*Stack
}

func NewThread()  *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int)  {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame)  {
	fmt.Printf("----------Pushing a frame\n")
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame  {
	fmt.Printf("----------Poping a frame\n")
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame  {
	return self.stack.top()
}

//func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
//	return NewFrame(self, maxLocals, maxStack)
//}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(self, method)
}