package a

/*
#cgo CFLAGS: -I../lib

void FuncA();
*/
import "C"
import (
	"example/internal"
)

//export export_SayA
func export_SayA(v C.int) {
	internal.Say(int(v))
}

func FuncA() {
	C.FuncA()
}
