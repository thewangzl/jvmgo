package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
)

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {
	//todo
}
