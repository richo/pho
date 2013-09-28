#ifndef _HACKS_H
#define _HACKS_H

#include <slash.h>

sl_vm_t* init_slash(const char* name);
void init_slash_static(void);
void eval(sl_vm_t* vm, char* script);
int eval_file(sl_vm_t* vm, char* filename);

#endif
