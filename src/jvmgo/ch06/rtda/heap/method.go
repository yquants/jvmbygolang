package heap

import "jvmgo/ch06/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = uint(codeAttr.MaxStack())
		self.maxLocals = uint(codeAttr.MaxLocals())
		self.code = codeAttr.Code()
	}
}

func (self *Method) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *Method) IsPrivate() bool {
	return 0 != self.accessFlags & ACC_PRIVATE
}

func (self *Method) IsProtected() bool {
	return 0 != self.accessFlags & ACC_PROTECTED
}

func (self *Method) IsStatic() bool {
	return 0 != self.accessFlags & ACC_STATIC
}

func (self *Method) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}
func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags & ACC_SYNCHRONIZED
}

func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags & ACC_BRIDGE
}

func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags & ACC_VARARGS
}

func (self *Method) IsNative() bool {
	return 0 != self.accessFlags & ACC_NATIVE
}

func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

func (self *Method) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) Class() *Class {
	return self.class
}

func (self *Method) Name() string {
	return self.name
}

func (self *Method) Code() []byte {
	return self.code
}
