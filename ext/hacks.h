#ifndef _HACKS_H
#define _HACKS_H

typedef enum {
    php_int_t,
    php_str_t
} php_types;

union intern_php_type {
    long as_long;
    void* as_ptr;
    char* as_str;
};

struct php_ret_t {
    php_types typ;
    union intern_php_type data;
};

#endif
