package instructions

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/instructions/constants"
	"fmt"
	"jvmgo/ch06/instructions/loads"
	"jvmgo/ch06/instructions/stores"
	"jvmgo/ch06/instructions/stack"
	"jvmgo/ch06/instructions/math"
	"jvmgo/ch06/instructions/conversions"
	"jvmgo/ch06/instructions/comparisons"
	"jvmgo/ch06/instructions/control"
	"jvmgo/ch06/instructions/references"
	"jvmgo/ch06/instructions/extended"
)

/*
http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-7.html
 */
func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00: return &constants.NOP{}
	case 0x01: return &constants.ACONST_NULL{}
	case 0x02: return &constants.ICONST_M1{}
	case 0x03: return &constants.ICONST_0{}
	case 0x04: return &constants.ICONST_1{}
	case 0x05: return &constants.ICONST_2{}
	case 0x06: return &constants.ICONST_3{}
	case 0x07: return &constants.ICONST_4{}
	case 0x08: return &constants.ICONST_5{}
	case 0x09: return &constants.LCONST_0{}
	case 0x0A: return &constants.LCONST_1{}
	case 0x0B: return &constants.FCONST_0{}
	case 0x0C: return &constants.FCONST_1{}
	case 0x0D: return &constants.FCONST_2{}
	case 0x0E: return &constants.DCONST_0{}
	case 0x0F: return &constants.DCONST_1{}
	case 0x10: return &constants.BIPUSH{}
	case 0x11: return &constants.SIPUSH{}
	case 0x12: return &constants.LDC{}
	case 0x13: return &constants.LDC_W{}
	case 0x14: return &constants.LDC2_W{}
	case 0x15: return &loads.ALOAD{}
	//case 0x16: return &loads.LLOAD{}
	//case 0x17: return &loads.FLOAD{}
	//case 0x18: return &loads.DLOAD{}
	case 0x19: return &loads.ALOAD{}
	case 0x1A: return &loads.ALOAD_0{}
	case 0x1B: return &loads.ALOAD_1{}
	case 0x1C: return &loads.ALOAD_2{}
	case 0x1D: return &loads.ALOAD_3{}
	//case 0x1E: return &loads.LLOAD_0{}
	//case 0x1F: return &loads.LLOAD_1{}
	//case 0x20: return &loads.LLOAD_2{}
	//case 0x21: return &loads.LLOAD_3{}
	//case 0x22: return &loads.FLOAD_0{}
	//case 0x23: return &loads.FLOAD_1{}
	//case 0x24: return &loads.FLOAD_2{}
	//case 0x25: return &loads.FLOAD_3{}
	//case 0x26: return &loads.DLOAD_0{}
	//case 0x27: return &loads.DLOAD_1{}
	//case 0x28: return &loads.DLOAD_2{}
	//case 0x29: return &loads.DLOAD_3{}
	case 0x2A: return &loads.ALOAD_0{}
	case 0x2B: return &loads.ALOAD_1{}
	case 0x2C: return &loads.ALOAD_2{}
	case 0x2D: return &loads.ALOAD_3{}
	//case 0x2E: return &loads.IALOAD{}
	//case 0x2F: return &loads.LALOAD{}
	//case 0x30: return &loads.FALOAD{}
	//case 0x31: return &loads.DALOAD{}
	//case 0x32: return &loads.AALOAD{}
	//case 0x33: return &loads.BALOAD{}
	//case 0x34: return &loads.CALOAD{}
	//case 0x35: return &loads.SALOAD{}
	case 0x36: return &stores.ASTORE{}
	case 0x37: return &stores.LSTORE{}
	//case 0x38: return &stores.FSTORE{}
	//case 0x39: return &stores.DSTORE{}
	case 0x3A: return &stores.ASTORE{}
	case 0x3B: return &stores.ASTORE_0{}
	case 0x3C: return &stores.ASTORE_1{}
	case 0x3D: return &stores.ASTORE_2{}
	case 0x3E: return &stores.ASTORE_3{}
	case 0x3F: return &stores.LSTORE_0{}
	case 0x40: return &stores.LSTORE_1{}
	case 0x41: return &stores.LSTORE_2{}
	case 0x42: return &stores.LSTORE_3{}
	//case 0x43: return &stores.FSTORE_0{}
	//case 0x44: return &stores.FSTORE_1{}
	//case 0x45: return &stores.FSTORE_2{}
	//case 0x46: return &stores.FSTORE_3{}
	//case 0x47: return &stores.DSTORE_0{}
	//case 0x48: return &stores.DSTORE_1{}
	//case 0x49: return &stores.DSTORE_2{}
	//case 0x4A: return &stores.DSTORE_3{}
	case 0x4B: return &stores.ASTORE_0{}
	case 0x4C: return &stores.ASTORE_1{}
	case 0x4D: return &stores.ASTORE_2{}
	case 0x4E: return &stores.ASTORE_3{}
	//case 0x4F: return &stores.IASTORE{}
	//case 0x50: return &stores.LASTORE{}
	//case 0x51: return &stores.FASTORE{}
	//case 0x52: return &stores.DASTORE{}
	//case 0x53: return &stores.AASTORE{}
	//case 0x54: return &stores.BASTORE{}
	//case 0x55: return &stores.CASTORE{}
	//case 0x56: return &stores.SASTORE{}
	case 0x57: return &stack.POP{}
	case 0x58: return &stack.POP2{}
	case 0x59: return &stack.DUP{}
	case 0x5A: return &stack.DUP_X1{}
	case 0x5B: return &stack.DUP_X2{}
	case 0x5C: return &stack.DUP2{}
	case 0x5D: return &stack.DUP2_X1{}
	case 0x5E: return &stack.DUP2_X2{}
	case 0x5F: return &stack.SWAP{}
	case 0x60: return &math.IADD{}
	case 0x61: return &math.LADD{}
	case 0x62: return &math.FADD{}
	case 0x63: return &math.DADD{}
	case 0x64: return &math.IDIV{}
	case 0x65: return &math.LDIV{}
	case 0x66: return &math.FDIV{}
	case 0x67: return &math.DDIV{}
	case 0x68: return &math.IMUL{}
	case 0x69: return &math.LMUL{}
	case 0x6A: return &math.FMUL{}
	case 0x6B: return &math.DMUL{}
	case 0x6C: return &math.IDIV{}
	case 0x6D: return &math.LDIV{}
	case 0x6E: return &math.FDIV{}
	case 0x6F: return &math.DDIV{}
	case 0x70: return &math.INEG{}
	case 0x71: return &math.LNEG{}
	case 0x72: return &math.FNEG{}
	case 0x73: return &math.DNEG{}
	case 0x74: return &math.INEG{}
	case 0x75: return &math.LNEG{}
	case 0x76: return &math.FNEG{}
	case 0x77: return &math.DNEG{}
	case 0x78: return &math.ISHL{}
	case 0x79: return &math.LSHL{}
	case 0x7A: return &math.ISHR{}
	case 0x7B: return &math.LSHR{}
	case 0x7C: return &math.IUSHR{}
	case 0x7D: return &math.LUSHR{}
	case 0x7E: return &math.IAND{}
	case 0x7F: return &math.LAND{}
	case 0x80: return &math.IOR{}
	case 0x81: return &math.LOR{}
	case 0x82: return &math.IXOR{}
	case 0x83: return &math.LXOR{}
	case 0x84: return &math.IINC{}
	case 0x85: return &conversions.I2L{}
	case 0x86: return &conversions.I2F{}
	case 0x87: return &conversions.I2D{}
	case 0x88: return &conversions.L2I{}
	case 0x89: return &conversions.L2F{}
	case 0x8A: return &conversions.L2D{}
	case 0x8B: return &conversions.F2I{}
	case 0x8C: return &conversions.F2L{}
	case 0x8D: return &conversions.F2D{}
	case 0x8E: return &conversions.D2I{}
	case 0x8F: return &conversions.D2L{}
	case 0x90: return &conversions.D2F{}
	//case 0x91: return &conversions.I2B{}
	//case 0x92: return &conversions.I2C{}
	//case 0x93: return &conversions.I2S{}
	case 0x94: return &comparisons.LCMP{}
	case 0x95: return &comparisons.FCMPL{}
	case 0x96: return &comparisons.FCMPG{}
	case 0x97: return &comparisons.DCMPL{}
	case 0x98: return &comparisons.DCMPG{}
	case 0x99: return &comparisons.IFEQ{}
	case 0x9A: return &comparisons.IFNE{}
	case 0x9B: return &comparisons.IFLT{}
	case 0x9C: return &comparisons.IFGE{}
	case 0x9D: return &comparisons.IFGT{}
	case 0x9E: return &comparisons.IFLE{}
	case 0x9F: return &comparisons.IF_ICMPEQ{}
	case 0xA0: return &comparisons.IF_ICMPNE{}
	case 0xA1: return &comparisons.IF_ICMPLT{}
	case 0xA2: return &comparisons.IF_ICMPGE{}
	case 0xA3: return &comparisons.IF_ICMPGT{}
	case 0xA4: return &comparisons.IF_ICMPLE{}
	case 0xA5: return &comparisons.IF_ACMPEQ{}
	case 0xA6: return &comparisons.IF_ACMPNE{}
	case 0xA7: return &control.GOTO{}
	//case 0xA8: return &constants.JSR{}
	//case 0xA9: return &constants.RET{}
	case 0xAA: return &control.TABLE_SWITCH{}
	case 0xAB: return &control.LOOKUP_SWITCH{}
	case 0xAC: return &control.IRETURN{}
	case 0xAD: return &control.LRETURN{}
	case 0xAE: return &control.FRETURN{}
	case 0xAF: return &control.DRETURN{}
	case 0xB0: return &control.ARETURN{}
	case 0xB1: return &control.RETURN{}
	case 0xB2: return &references.GET_STATIC{}
	case 0xB3: return &references.PUT_STATIC{}
	case 0xB4: return &references.GET_FIELD{}
	case 0xB5: return &references.PUT_FIELD{}
	case 0xB6: return &references.INVOKE_VIRTUAL{}
	case 0xB7: return &references.INVOKE_SPECIAL{}
	//case 0xB8: return &references.INVOKE_STATIC{}
	//case 0xB9: return &references.INVOKE_INTERFACE{}
	//case 0xBA: return &references.INVOKE_DYNAMIC{}
	case 0xBB: return &references.NEW{}
	//case 0xBC: return &constants.NEWARRAY{}
	//case 0xBD: return &constants.ANEWARRAY{}
	//case 0xBE: return &constants.ARRAYLENGTH{}
	//case 0xBF: return &constants.ATHROW{}
	case 0xC0: return &references.CHECK_CAST{}
	case 0xC1: return &references.INSTANCE_OF{}
	//case 0xC2: return &constants.MONITORENTER{}
	//case 0xC3: return &constants.MONITOREXIT{}
	//case 0xC4: return &constants.WIDE{}
	//case 0xC5: return &constants.MULTIANEWARRAY{}
	//case 0xC6: return &constants.IFNULL{}
	//case 0xC7: return &constants.IFNONNULL{}
	case 0xC8: return &extended.GOTO_W{}
	//case 0xC9: return &extended.JSR_W{}
	//case 0xCA: return &extended.BREAKPOINT{}
	/*
case 0xCB: return &constants.ACONST_NULL{}
case 0xCC: return &constants.NOP{}
case 0xCD: return &constants.ACONST_NULL{}
case 0xCE: return &constants.NOP{}
case 0xCF: return &constants.ACONST_NULL{}
case 0xE0: return &constants.NOP{}
case 0xE1: return &constants.ACONST_NULL{}
case 0xE2: return &constants.NOP{}
case 0xE3: return &constants.ACONST_NULL{}
case 0xE4: return &constants.NOP{}
case 0xE5: return &constants.ACONST_NULL{}
case 0xE6: return &constants.NOP{}
case 0xE7: return &constants.ACONST_NULL{}
case 0xE8: return &constants.NOP{}
case 0xE9: return &constants.ACONST_NULL{}
case 0xEA: return &constants.NOP{}
case 0xEB: return &constants.ACONST_NULL{}
case 0xEC: return &constants.NOP{}
case 0xED: return &constants.ACONST_NULL{}
case 0xEE: return &constants.NOP{}
case 0xEF: return &constants.ACONST_NULL{}
case 0xF0: return &constants.NOP{}
case 0xF1: return &constants.ACONST_NULL{}
case 0xF2: return &constants.NOP{}
case 0xF3: return &constants.ACONST_NULL{}
case 0xF4: return &constants.NOP{}
case 0xF5: return &constants.ACONST_NULL{}
case 0xF6: return &constants.NOP{}
case 0xF7: return &constants.ACONST_NULL{}
case 0xF8: return &constants.NOP{}
case 0xF9: return &constants.ACONST_NULL{}
case 0xFA: return &constants.NOP{}
case 0xFB: return &constants.ACONST_NULL{}
case 0xFC: return &constants.NOP{}
case 0xFD: return &constants.ACONST_NULL{}
*/
	//case 0xFE: return &constants.IMPDEP1{}
	//case 0xFF: return &constants.IMPDEP2{}

	default:
		panic(fmt.Errorf("Unsupported opcoe: 0x%x!", opcode))
	}
}
