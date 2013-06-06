#include <geos_c.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>

void gogeos_notice_handler(const char *fmt, ...);
void gogeos_error_handler(const char *fmt, ...);
char *gogeos_get_last_error(void);
GEOSContextHandle_t gogeos_initGEOS();
