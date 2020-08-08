package rtda

import "jvmgo/xxx/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
