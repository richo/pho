pho
===

PHP and Golang. Always a good decision.

### Hacking

Take careful note of where you branch if you hack on this repo. It will undergo and epic rebase at some point.

### Build you some php

```
CC=clang ./configure --enable-pho --enable-debug --without-pear --disable-cli --disable-cgi --enable-embed --prefix ~/.php/versions/trunk
```

Note that on OSX you'll want to merge the pho/master branch of [my fork][1] of
php into your branch, it patches the libtool macros to ensure you get a shared
library and not a Mach Bundle.

Then install it into ./libs in this repo.

### Debugging

```bash
LD_LIBRARY_PATH=~/.php/versions/trunk/lib PHP_LIB_PATH=~/.php/versions/trunk/lib/libphp5.so gdb bin/pho
```

#### What people are saying:

Lets be honest, you going to hell is a forgone conclusion by this point, may as well make the most of it.
~ @rjbone

[1]: https://github.com/richo/php-src
