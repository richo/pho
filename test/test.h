#ifndef __TEST_H
#define __TEST_H

#include <stdlib.h>
#include <stdio.h>

typedef int test;

static int __failure = 0;
#define RUN_TEST(name) do { if (!name()) __failure = 1; } while (0)
#define RESULT return __failure

#define PASS return 1;
#define FAIL(msg) do { fprintf(stderr, "! %s:%d: %s\n", __FILE__, __LINE__, msg); return 0; } while (0)
#define ASSERT(test, msg) if (!(test)) FAIL(msg)

#endif
