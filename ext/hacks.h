#ifndef _HACKS_H
#define _HACKS_H

enum php_types {
    php_int_t = 1,
    php_str_t
};

union intern_php_type {
    long as_long;
    void* as_ptr;
    char* as_str;
};

struct php_ret_t {
    enum php_types typ;
    union intern_php_type data;
};

struct php_ret_t* get_value(char* key);
void* set_int_value(char*, long);
void*** init_php(void);
void*** init_php2(int, char**);
void eval(char*);
int eval_file(char*);
void* new_interpreter_context(void);
void* set_interpreter_context(void*);

#endif
