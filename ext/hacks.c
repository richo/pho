#include <stdlib.h>
#include <stdio.h>
#include <stdio.h>

#include "hacks.h"

static void
output(sl_vm_t* vm, char* buff, size_t len)
{
    fwrite(buff, len, 1, stdout);
    fflush(stdout);
}

static void
setup_vm_response(sl_vm_t* vm)
{
    sl_response_opts_t res;
    res.descriptive_error_pages = 0;
    res.buffered = 0;
    res.write = output;
    sl_response_set_opts(vm, &res);
}

void init_slash_static(void) {
    sl_static_init();
}

sl_vm_t* init_slash(const char* name) {
    sl_vm_t* vm = sl_init("cli");
    /* sl_gc_set_stack_top(vm->arena, stack_top); */
    setup_vm_response(vm);
    return vm;
}

void eval(sl_vm_t* vm, char* script) {
    // TODO
}

int eval_file(sl_vm_t* vm, char* filename) {
    FILE* fh = fopen(filename, "rb");
    if (!fh)
        perror("Couldn't fopen file");
    if (!vm)
        fprintf(stderr, "No vm D:\n");

    size_t source_cap = 4096;
    size_t source_len = 0;
    uint8_t* source = malloc(source_cap);
    while(!feof(fh) && !ferror(fh)) {
        if(source_len + 4096 > source_cap) {
            source_cap += 4096;
            source = realloc(source, source_cap);
        }
        source_len += fread(source, 1, 4096, fh);
    }
    if(ferror(fh)) {
        perror("Error while reading source file");
    }
    sl_do_string(vm, source, source_len, filename, 0);
    return 0;
}
