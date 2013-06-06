#include "geos.h"

void gogeos_notice_handler(const char *fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    fprintf(stderr, "NOTICE: ");
    vfprintf(stderr, fmt, ap);
    va_end(ap);
}

#define ERRLEN 256

char gogeos_last_err[ERRLEN];

void gogeos_error_handler(const char *fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    vsnprintf(gogeos_last_err, (size_t) ERRLEN, fmt, ap);
    va_end(ap);
}

char *gogeos_get_last_error(void) {
    return gogeos_last_err;
}

GEOSContextHandle_t gogeos_initGEOS() {
    return initGEOS_r(gogeos_notice_handler, gogeos_error_handler);
}
