package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function")
	C.cHello()

	fmt.Println("Going to call another C function")
	myMessage := C.CString("Hello C, from Go")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("Done")
}
