package lang

import "math"
import "jvmgo/ch10/native"
import "jvmgo/ch10/rtda"

func init() {
	native.Register("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRawIntBits)
}

//public static native int floatToRawIntBits(float value)
func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}
