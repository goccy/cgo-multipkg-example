package main

/*
#cgo CFLAGS: -I./

void FuncA();
void FuncB();
*/
import "C"

func main() {
	C.FuncA()
	C.FuncB()
}
