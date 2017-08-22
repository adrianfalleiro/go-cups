package main

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -lcups
#include "main.h"
#include "cups/cups.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export go_callback
func go_callback(userData unsafe.Pointer, flags C.unsigned, dest *C.cups_dest_t) {
	name := C.GoString(dest.name)
	fmt.Println(name)
}

func main() {

	C.listPrinters()
	C.listPrinters()

	select {}

}
