package a

/*
#cgo CFLAGS: -I../lib

void FuncA();
*/
import "C"

func FuncA() {
	C.FuncA()
}
