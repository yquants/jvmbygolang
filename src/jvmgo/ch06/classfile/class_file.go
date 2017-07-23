package classfile

import "fmt"

type ClassFile struct {
	//magic 	uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/**
p24
The information (along with its order) inside of Java Class file has been specified in JSR

the read() method below is designed to read the whole Java Class out following the sequence defined in JSR
*/
func (self *ClassFile) read(reader *ClassReader) {
	//u4(magic)
	self.readAndCheckMagic(reader)
	//u2(minor_version) + u2(major_version)
	self.readAndCheckVersion(reader)
	//u2(constant_pool_count) + ? (constant_pool[constant_pool_count -1])
	self.constantPool = readConstantPool(reader)
	//u2(access_flags)
	self.accessFlags = reader.readUint16()
	//u2(this_class) -> a refercne to constant pool
	self.thisClass = reader.readUint16()
	//u2(super_class) -> a refercne to constant pool
	self.superClass = reader.readUint16()
	//u2(interfaces_count) + u2(interfaces[interfaces_count])
	self.interfaces = reader.readUnit16s()
	//u2(fields_count) + ? (fields[fields_count])
	self.fields = readMembers(reader, self.constantPool)
	//u2(methods_count) + ? (methods[methods_count])
	self.methods = readMembers(reader, self.constantPool)
	//u2(attributes_count) + ? (attributes[attributes_count])
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {

	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
