#include <stdlib.h>
#include <stdio.h>

#include <sapi/embed/php_embed.h>

char** php_init_args(void);

static const char* init_arg = "embed4";

char** php_init_args(void) {
    char **args;
    args = malloc(sizeof(char) * 2);
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

void* get(char* key)
{
    zval **value;

    if(zend_hash_find(EG(active_symbol_table),
                key,
                strlen(key),
                (void **)&value) == SUCCESS) {
        return zval2go(value);
        // TODO return a Type,ptr array to unpack on the go side.
    }

    return NULL;
}

// Stupid debugging harness to test variable traversal
long get_int_value(char* key) {
{
    zval **value;

    if(zend_hash_find(EG(active_symbol_table),
                key,
                strlen(key),
                (void **)&value) == SUCCESS) {
        return *zval2go(value);
        // TODO return a Type,ptr array to unpack on the go side.
    }

    return NULL;
}

void* zval2go(zval *value) {
  switch(Z_TYPE_P(value)) {
    case IS_LONG:
      return &Z_LVAL_P(value);
      break;
    default:
      // Not implemented
      return NULL;
  }
}
