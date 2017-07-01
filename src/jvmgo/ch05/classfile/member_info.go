package classfile

/**
The definition of structure of filed/method in .class per JSR
filed_info{
u2(access_flags)
u2(name_index)
u2(descriptor_index)
u2(attributes_count) + ?(attributes[attributes_count])
*/
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// read all members in the .class file
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// read each member in the .class file
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
