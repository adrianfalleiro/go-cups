#include <stdio.h>
#include <stdlib.h>
#include "cups/cups.h"

extern void go_callback(void *user_data, unsigned flags, cups_dest_t *dest);

void print_dest(void *user_data, unsigned flags, cups_dest_t *dest)
{
	printf("%s\n", dest->name);
}
int listPrinters()
{
	cups_ptype_t mask = CUPS_PRINTER_LOCAL;
	cups_ptype_t type = CUPS_PRINTER_LOCAL;
	int msec = 0;
	int cancel;

  	return cupsEnumDests(CUPS_DEST_FLAGS_NONE, msec, NULL, type, mask, (cups_dest_cb_t)go_callback, NULL);
}