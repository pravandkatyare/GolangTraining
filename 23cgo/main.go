package main

// #include <stdlib.h>
// #include <stdio.h>
// #include "greeter.h"
import "C"
import "unsafe"

func main() {
	// C.greet()

	// name := C.CString("Foo")
	// defer C.free(unsafe.Pointer(name))

	// C.greetString(name)

	cs := C.struct_Greetee{name: C.CString("FOO"), year: C.int(2024)}
	defer C.free(unsafe.Pointer(&cs))

	C.greetStruct(&cs)
}
