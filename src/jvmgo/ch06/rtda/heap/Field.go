package heap

import "jvmgo/ch06/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) Descriptor() string {
	return self.descriptor
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) isLongOrDouble() bool {
	//todo
	return self.descriptor == "L" || self.descriptor == "D"
}

func (self *Field) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *Field) IsPrivate() bool {
	return 0 != self.accessFlags & ACC_PRIVATE
}

func (self *Field) IsProtected() bool {
	return 0 != self.accessFlags & ACC_PROTECTED
}

func (self *Field) IsStatic() bool {
	return 0 != self.accessFlags & ACC_STATIC
}

func (self *Field) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags & ACC_VOLATILE
}

func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags & ACC_TRANSIENT
}

func (self *Field) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}

func (self *Field) Class() *Class {
	return self.class
}