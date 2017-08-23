#include <stdio.h>
#include <stdlib.h>
#include "cups/cups.h"

extern int listPrinters_go_callback(void *user_data, unsigned flags, cups_dest_t *dest);

void print_dest(void *user_data, unsigned flags, cups_dest_t *dest)
{
    printf("%s\n", dest->name);
}
int numDests() {
    cups_dest_t *dests;
    int num_dests = cupsGetDests(&dests);

    free(dests);
    return num_dests;
}
int listPrinters()
{
	cups_ptype_t mask = CUPS_PRINTER_LOCAL;
	cups_ptype_t type = CUPS_PRINTER_LOCAL;
	int msec = 0;
    int cancel;
    int results = 0;

    return cupsEnumDests(CUPS_DEST_FLAGS_NONE, msec, NULL, type, mask, (cups_dest_cb_t)listPrinters_go_callback, NULL);

}