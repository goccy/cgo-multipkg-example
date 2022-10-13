package internal

/*
#cgo CFLAGS: -I../lib

void Say(int v);
*/
import "C"

func Say(v int) {
	C.Say(C.int(v))
}
