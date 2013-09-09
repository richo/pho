#include "test.h"
#include "../ext/hacks.h"

test struct_works(void) {
    struct php_ret_t s;

    s.data.as_ptr = (void*) 1234;
    ASSERT(s.data.as_long == 1234, "Data doesn't come through the union");
    PASS;
}

int main(int argc, char **argv) {
    RUN_TEST(struct_works);

    RESULT;
}
