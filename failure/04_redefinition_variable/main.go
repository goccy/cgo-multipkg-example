package main

/*
#cgo CFLAGS: -Ilib

void FuncA();
void FuncB();
*/
import "C"

func main() {
	C.FuncA()
	C.FuncB()
}
