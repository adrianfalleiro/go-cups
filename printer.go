package main

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -lcups
#include "main.h"
#include "cups/cups.h"
*/
import "C"

// Printer is a struct for each printer
type Printer struct {
	Name      string
	IsDefault bool
}

// TestPrint prints a test page
func (p *Printer) TestPrint() int {

	var numOptions C.int
	var options *C.cups_option_t
	var jobID C.int

	// Print a single file
	jobID = C.cupsPrintFile(C.CString(p.Name), C.CString("/usr/share/cups/data/testprint"),
		C.CString("Test Print"), numOptions, options)

	return int(jobID)
}
