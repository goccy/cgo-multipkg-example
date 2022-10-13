package b

/*
#cgo CFLAGS: -I../lib

void FuncB();
*/
import "C"

func FuncB() {
	C.FuncB()
}
