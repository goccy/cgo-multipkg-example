package internal

/*
#cgo CFLAGS: -I../lib

void FuncA();
void FuncB();
*/
import "C"

func FuncA() {
	C.FuncA()
}

func FuncB() {
	C.FuncB()
}
