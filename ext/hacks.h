#ifndef _HACKS_H
#define _HACKS_H

typedef enum {
    php_int_t,
    php_str_t
} php_types;

typedef struct {
    php_types type;
    void* data;
} php_ret_t;

#endif
