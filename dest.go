package cups

/*
#cgo LDFLAGS: -lcups
#include "cups/cups.h"
*/
import "C"

const printerStateIdle = "3"
const printerStatePrinting = "4"
const printerStateStopped = "5"

// Dest is a struct for each printer
type Dest struct {
	Name      string
	IsDefault bool
	options   map[string]string
	// internal keys for status
}

// PrintFile prints a file
func (d *Dest) PrintFile(filename, mimeType string) int {
	// check mime type
	// check file
	return 0
}

// Status returns the status of the dest
func (d *Dest) Status() string {
	var returnMessage string

	// Return status of dest
	value, ok := d.options["printer-state"]

	if ok != true {
		returnMessage = "printer state key does not exist"
	}

	switch value {
	case printerStateIdle:
		returnMessage = "idle"
		break
	case printerStatePrinting:
		returnMessage = "printing"
		break
	case printerStateStopped:
		returnMessage = "stopped"
		break
	default:
		returnMessage = "error"
		break
	}

	return returnMessage

}

// GetOption returns the options
func (d *Dest) GetOption(keyName string) string {
	// Return option
	return ""
}

// GetOptions returns a map of the dest options
func (d *Dest) GetOptions() map[string]string {
	// Return option
	return d.options
}

// TestPrint prints a test page
func (d *Dest) TestPrint() int {

	var numOptions C.int
	var options *C.cups_option_t
	var jobID C.int

	// resolve path/to/test/file
	osxTest := "/usr/share/cups/data/testprint"

	// Print a single file
	jobID = C.cupsPrintFile(C.CString(d.Name), C.CString(osxTest),
		C.CString("Test Print"), numOptions, options)

	return int(jobID)
}
