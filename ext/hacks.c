#include <stdlib.h>
#include <stdio.h>
#include <stdio.h>

#include "hacks.h"

#include <sapi/embed/php_embed.h>

char** php_init_args(void);
struct php_ret_t* zval2go(zval **value);

static void ***tsrm_ls;

static const int php_argc = 1;
static const char* php_argv[] = {"embed4", NULL};

void*** init_php(void) {
    php_embed_init(php_argc, php_argv, &tsrm_ls);
    return tsrm_ls;
}

void eval(char* script) {
    static const char* name = "<EVAL>";
    zend_eval_string(script, NULL, name, tsrm_ls);
}

void eval_file(char* filename) {
  zend_file_handle script;

  script.type = ZEND_HANDLE_FP;
  script.filename = filename;
  script.opened_path = NULL;

  if (!(script.handle.fp = fopen(filename, "rb"))) {
      // XXX Handle error
      return;
  }

  script.free_filename = 0;

  php_execute_script(&script, tsrm_ls);
}


void eval_and_print(char* script) {
    void* ret = NULL;

    zend_eval_string(script, ret, "<eval_and_print>", tsrm_ls);
    printf("eval_and_print> %p\n", ret);
}

void* set_int_value(char* key, long v) {
    zval *value;

    MAKE_STD_ZVAL(value);
    ZVAL_LONG(value, v);

    ZEND_SET_SYMBOL(EG(active_symbol_table), key, value);
}

// Stupid debugging harness to test variable traversal
struct php_ret_t* get_int_value(char* key) {
    zval **value;
    struct php_ret_t *ret;

    fprintf(stderr, "C lookup: %s\n", key);

    if(zend_hash_find(EG(active_symbol_table),
                key,
                strlen(key) + 1,
                (void **)&value) == SUCCESS) {
        ret = zval2go(value);
        fprintf(stderr, "C data location: %p\n", &(ret->data.as_long));
        if (ret->typ == php_int_t) {
            fprintf(stderr, "C int Value: %d\n", ret->data.as_long);
            return ret;
        } else if (ret->typ == php_str_t) {
            fprintf(stderr, "C str Value: %s\n", ret->data.as_ptr);
            return ret;
        }
    }

    return ret;
}

struct php_ret_t* zval2go(zval **value) {
    struct php_ret_t *ret = NULL;
    ret = (struct php_ret_t*)malloc(sizeof(struct php_ret_t));
    int len;
    char* str;

    switch(Z_TYPE_P(*value)) {
        case IS_LONG:
            ret->data.as_long = Z_LVAL_P(*value);
            ret->typ = php_int_t;
            break;
        case IS_STRING:
            len = Z_STRLEN_P(*value) + 1;
            str = malloc(sizeof(char) * len);
            memset(str, 0, len);
            memcpy(str, Z_STRVAL_P(*value), len - 1);
            ret->data.as_str = str;
            ret->typ = php_str_t;
            break;
        default:
            // Not implemented
            break;
    }
    return ret;
}
