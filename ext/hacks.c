#include <stdlib.h>
#include <stdio.h>
#include <stdio.h>

#include "hacks.h"

#include <sapi/embed/php_embed.h>

char** php_init_args(void);
struct php_ret_t* zval2go(zval **value);

static const char* init_arg = "embed4";

char** php_init_args(void) {
    char **args;
    args = malloc(sizeof(char*) * 2);
    args[0] = init_arg;
    args[1] = NULL;

    return args;
}

void init_php(void) {
    char **init_args;

    init_args = php_init_args();
    php_embed_init(1, init_args);
}

void eval(char* script) {
    static const char* name = "<EVAL>";
    zend_eval_string(script, NULL, name);
}

void eval_and_print(char* script) {
    void* ret = NULL;

    zend_eval_string(script, ret, "<eval_and_print>");
    printf("eval_and_print> %p\n", ret);
}

void* set_int_value(char* key, long v) {
    zval *value;

    MAKE_STD_ZVAL(value);
    ZVAL_LONG(value, v);

    ZEND_SET_SYMBOL(EG(active_symbol_table), key, value);
}

// Stupid debugging harness to test variable traversal
struct php_ret_t get_int_value(char* key) {
    zval **value;
    struct php_ret_t *ret;

    fprintf(stderr, "C lookup: %s\n", key);

    if(zend_hash_find(EG(active_symbol_table),
                key,
                strlen(key) + 1,
                (void **)&value) == SUCCESS) {
        ret = zval2go(value);
        if (ret->typ == php_int_t) {
            fprintf(stderr, "C int Value: %d\n", ret->data);
            return ret;
        } else if (ret->typ == php_str_t) {
            fprintf(stderr, "C str Value: %s\n", ret->data);
            return ret;
        }
    }

    return NULL;
}

struct php_ret_t* zval2go(zval **value) {
    struct php_ret_t *ret;
    ret = (struct php_ret_t*)malloc(sizeof(struct php_ret_t));
    int len;
    char* str;

    switch(Z_TYPE_P(*value)) {
        case IS_LONG:
            ret->data = Z_LVAL_P(*value);
            ret->typ = php_int_t;
            return ret;
            break;
        case IS_STRING:
            len = Z_STRLEN_P(*value) + 1;
            str = malloc(sizeof(char) * len);
            memset(str, 0, len);
            memcpy(str, Z_STRVAL_P(*value), len - 1);
            ret->data = str;
            ret->typ = php_str_t;
            return ret;
            break;
        default:
            // Not implemented
            return NULL;
            break;
    }
}
