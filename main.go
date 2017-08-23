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

var listPrintersChan chan *Printer

//export listPrinters_go_callback
func listPrinters_go_callback(userData unsafe.Pointer, flags C.unsigned, dest *C.cups_dest_t) C.int {

	printer := &Printer{
		Name:      C.GoString(dest.name),
		IsDefault: !(int(dest.is_default) == 0),
	}

	listPrintersChan <- printer

	return C.int(1)
}

func main() {
	printers := GetPrinters()
	fmt.Println(printers)
}

// GetPrinters returns a list of Printers
func GetPrinters() []*Printer {
	numPrinters := C.numDests()

	listPrintersChan = make(chan *Printer, int(numPrinters))
	C.listPrinters()
	close(listPrintersChan)

	printers := []*Printer{}

	for message := range listPrintersChan {
		printers = append(printers, message)
	}

	return printers
}
