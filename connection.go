package cups

/*
#cgo LDFLAGS: -lcups
#include "cups/cups.h"
*/
import "C"
import (
	"unsafe"
)

// Connection is a struct containing information about a CUPS connection
type Connection struct {
	isDefault bool
	address   string
	NumDests  int
	Dests     []*Dest
}

const cupsDestTSize = 32
const cupsOptionTSize = 16

var enumDestsChan chan *Dest

// Refresh updates the list of destinations and their state
func (c *Connection) Refresh() {
	updateDefaultConnection(c)
}

// NewDefaultConnection creates a new default CUPS connection
func NewDefaultConnection() *Connection {
	connection := &Connection{
		isDefault: true,
	}
	updateDefaultConnection(connection)

	return connection
}

// TODO: check for memory leaks
func updateDefaultConnection(c *Connection) {
	var dests *C.cups_dest_t
	destCount := C.cupsGetDests(&dests)
	goDestCount := int(destCount)

	var destsArr []*Dest

	destPtr := uintptr(unsafe.Pointer(dests))
	for i := 0; i < goDestCount; i++ {
		// is this ok?
		dest := (*C.cups_dest_t)(unsafe.Pointer(destPtr))
		d := &Dest{
			Name: C.GoString(dest.name),
		}

		destOptPtr := uintptr(unsafe.Pointer(dest.options))

		options := make(map[string]string)
		for j := 0; j < int(dest.num_options)-1; j++ {
			opt := (*C.cups_option_t)(unsafe.Pointer(destOptPtr))
			options[C.GoString(opt.name)] = C.GoString(opt.value)
			destOptPtr += uintptr(cupsOptionTSize)
		}

		d.options = options

		destsArr = append(destsArr, d)
		destPtr += uintptr(cupsDestTSize)
	}

	// free the pointers
	C.cupsFreeDests(destCount, dests)

	c.NumDests = goDestCount
	c.Dests = destsArr
}
