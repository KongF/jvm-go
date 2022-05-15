package instructions

import "fmt"
import "jvm-go/ch05/instructions/base"
import . "jvm-go/ch05/instructions/comparisons"
import . "jvm-go/ch05/instructions/constants"
import . "jvm-go/ch05/instructions/control"
import . "jvm-go/ch05/instructions/conversions"
import . "jvm-go/ch05/instructions/extended"
import . "jvm-go/ch05/instructions/loads"
import . "jvm-go/ch05/instructions/math"
import . "jvm-go/ch05/instructions/stack"
import . "jvm-go/ch05/instructions/stores"

var (
	nop         = &NOP{}
	aconst_null = &ACONST_NULL{}
	iconst_m1   = &ICONST_M1{}
	iconst_0    = &ICONST_0{}
	iconst_1    = &ICONST_1{}
	iconst_2    = &ICONST_2{}
	iconst_3    = &ICONST_3{}
	iconst_4    = &ICONST_4{}
	iconst_5    = &ICONST_5{}
	lconst_0    = &LCONST_0{}
	lconst_1    = &LCONST_1{}
	fconst_0    = &FCONST_0{}
	fconst_1    = &FCONST_1{}
	fconst_2    = &FCONST_2{}
	dconst_0    = &DCONST_0{}
	dconst_1    = &DCONST_1{}
	iload_0     = &ILOAD_0{}
	iload_1     = &ILOAD_1{}
	iload_2     = &ILOAD_2{}
	iload_3     = &ILOAD_3{}
	lload_0     = &LLOAD_0{}
	lload_1     = &LLOAD_1{}
	lload_2     = &LLOAD_2{}
	lload_3     = &LLOAD_3{}
	fload_0     = &FLOAD_0{}
	fload_1     = &FLOAD_1{}
	fload_2     = &FLOAD_2{}
	fload_3     = &FLOAD_3{}
	dload_0     = &DLOAD_0{}
	dload_1     = &DLOAD_1{}
	dload_2     = &DLOAD_2{}
	dload_3     = &DLOAD_3{}
	aload_0     = &ALOAD_0{}
	aload_1     = &ALOAD_1{}
	aload_2     = &ALOAD_2{}
	aload_3     = &ALOAD_3{}
	// iaload      = &IALOAD{}
	// laload      = &LALOAD{}
	// faload      = &FALOAD{}
	// daload      = &DALOAD{}
	// aaload      = &AALOAD{}
	// baload      = &BALOAD{}
	// caload      = &CALOAD{}
	// saload      = &SALOAD{}
	istore_0 = &ISTORE_0{}
	istore_1 = &ISTORE_1{}
	istore_2 = &ISTORE_2{}
	istore_3 = &ISTORE_3{}
	lstore_0 = &LSTORE_0{}
	lstore_1 = &LSTORE_1{}
	lstore_2 = &LSTORE_2{}
	lstore_3 = &LSTORE_3{}
	fstore_0 = &FSTORE_0{}
	fstore_1 = &FSTORE_1{}
	fstore_2 = &FSTORE_2{}
	fstore_3 = &FSTORE_3{}
	dstore_0 = &DSTORE_0{}
	dstore_1 = &DSTORE_1{}
	dstore_2 = &DSTORE_2{}
	dstore_3 = &DSTORE_3{}
	astore_0 = &ASTORE_0{}
	astore_1 = &ASTORE_1{}
	astore_2 = &ASTORE_2{}
	astore_3 = &ASTORE_3{}
	// iastore  = &IASTORE{}
	// lastore  = &LASTORE{}
	// fastore  = &FASTORE{}
	// dastore  = &DASTORE{}
	// aastore  = &AASTORE{}
	// bastore  = &BASTORE{}
	// castore  = &CASTORE{}
	// sastore  = &SASTORE{}
	pop     = &POP{}
	pop2    = &POP2{}
	dup     = &DUP{}
	dup_x1  = &DUP_X1{}
	dup_x2  = &DUP_X2{}
	dup2    = &DUP2{}
	dup2_x1 = &DUP2_X1{}
	dup2_x2 = &DUP2_X2{}
	swap    = &SWAP{}
	iadd    = &IADD{}
	ladd    = &LADD{}
	fadd    = &FADD{}
	dadd    = &DADD{}
	isub    = &ISUB{}
	lsub    = &LSUB{}
	fsub    = &FSUB{}
	dsub    = &DSUB{}
	imul    = &IMUL{}
	lmul    = &LMUL{}
	fmul    = &FMUL{}
	dmul    = &DMUL{}
	idiv    = &IDIV{}
	ldiv    = &LDIV{}
	fdiv    = &FDIV{}
	ddiv    = &DDIV{}
	irem    = &IREM{}
	lrem    = &LREM{}
	frem    = &FREM{}
	drem    = &DREM{}
	ineg    = &INEG{}
	lneg    = &LNEG{}
	fneg    = &FNEG{}
	dneg    = &DNEG{}
	ishl    = &ISHL{}
	lshl    = &LSHL{}
	ishr    = &ISHR{}
	lshr    = &LSHR{}
	iushr   = &IUSHR{}
	lushr   = &LUSHR{}
	iand    = &IAND{}
	land    = &LAND{}
	ior     = &IOR{}
	lor     = &LOR{}
	ixor    = &IXOR{}
	lxor    = &LXOR{}
	i2l     = &I2L{}
	i2f     = &I2F{}
	i2d     = &I2D{}
	l2i     = &L2I{}
	l2f     = &L2F{}
	l2d     = &L2D{}
	f2i     = &F2I{}
	f2l     = &F2L{}
	f2d     = &F2D{}
	d2i     = &D2I{}
	d2l     = &D2L{}
	d2f     = &D2F{}
	i2b     = &I2B{}
	i2c     = &I2C{}
	i2s     = &I2S{}
	lcmp    = &LCMP{}
	fcmpl   = &FCMPL{}
	fcmpg   = &FCMPG{}
	dcmpl   = &DCMPL{}
	dcmpg   = &DCMPG{}
	// ireturn = &IRETURN{}
	// lreturn = &LRETURN{}
	// freturn = &FRETURN{}
	// dreturn = &DRETURN{}
	// areturn = &ARETURN{}
	// _return = &RETURN{}
	// arraylength   = &ARRAY_LENGTH{}
	// athrow        = &ATHROW{}
	// monitorenter  = &MONITOR_ENTER{}
	// monitorexit   = &MONITOR_EXIT{}
	// invoke_native = &INVOKE_NATIVE{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return &NOP{}
	case 0x01:
		return &ACONST_NULL{}
	}
}
