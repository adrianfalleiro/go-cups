package main

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -lcups
#include <stdio.h>
#include <stdlib.h>
#include "cups/cups.h"
*/
import "C"

import "fmt"

func main() {

	var dests *C.cups_dest_t
	numDests := C.cupsGetDests(&dests)
	defaultDest := C.cupsGetDest(nil, nil, numDests, dests)
	defaultDestName := C.GoString(defaultDest.name)
	fmt.Print(defaultDestName)

}
