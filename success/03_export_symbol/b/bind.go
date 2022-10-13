package b

/*
#cgo CFLAGS: -I../lib

void FuncB();
*/
import "C"
import (
	"example/internal"
)

//export export_SayB
func export_SayB(v C.int) {
	internal.Say(int(v))
}

func FuncB() {
	C.FuncB()
}
