package main

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -lcups
#include <stdio.h>
#include <stdlib.h>
#include "cups/cups.h"
void print_dest(void *user_data, unsigned flags, cups_dest_t *dest)
{
  	if (dest->instance)
    	printf("%s/%s\n", dest->name, dest->instance);
	else
		printf("a%s\n", dest->name);
}
int listPrinters()
{
	cups_ptype_t mask = CUPS_PRINTER_LOCAL;
	cups_ptype_t type = CUPS_PRINTER_LOCAL;
	int msec = 0;
	int cancel;

  	return cupsEnumDests(CUPS_DEST_FLAGS_NONE, msec, NULL, type, mask, (cups_dest_cb_t)print_dest, NULL);
}
*/
import "C"

func main() {

	ListPrinters()

}

// ListPrinters function ...
func ListPrinters() {
	C.listPrinters()
}
