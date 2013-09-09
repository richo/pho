#include "test.h"
#include "../ext/hacks.h"

test struct_works(void) {
    struct php_ret_t s;

    s.ptr_data = (void*) 1234;
    ASSERT(s.long_data == 1234, "Data doesn't come through the union");
    PASS;
}

int main(int argc, char **argv) {
    RUN_TEST(struct_works);

    RESULT;
}
