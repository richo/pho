#include <stdlib.h>
#include <stdio.h>

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

    zend_val_string(script, ret, "<eval_and_print>");
    printf("eval_and_print> %p\n", ret);
}