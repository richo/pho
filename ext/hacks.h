#ifndef _HACKS_H
#define _HACKS_H

typedef enum {
    php_int_t,
    php_str_t
} php_types;

struct php_ret_t {
    php_types typ;
    union {
        long long_data;
        void* ptr_data;
    };
};

#endif
