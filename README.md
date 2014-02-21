Header Generator
====

This program generates header guards and a little doxygen for C header files (probably fine for hpp too, haven't devoted any thought to that).

Installing
====

You need Go to build. Clone the repository. If you already have $HOME/bin in your path, just do make install. Otherwise do make, and then copy hgen to some place in your path.

Using in a Project
====

To use this in a project, create an empty .hgenconfig file in the root of your project directory. You may also put your name in the .hgenconfig, in which case it will be inserted into the @author tag in the doxygen header for the file. Note that whatever
you put in the .hgenconfig file is exactly what will appear after the @author tag.

For example, if your project directory is located at /home/shaladdle/fav-project, you can do

```
echo shaladdle >> /home/shaladdle/fav-project/.hgenconfig
```

Once you do that, you can use hgen to make header files anywhere under fav-project. For example, if you did

```
cd fav-project
mkdir module1
touch module1/module1.h
```

hgen would create a header file that looks like this:

```c
/** @file module1.h
 *  @brief Function prototypes for module1.
 *
 *  @author shaladdle
 */

#ifndef __MODULE1_MODULE1_H_
#define __MODULE1_MODULE1_H_

#endif
```

The guards themselves are generated using the path to the root of the project. So if your header file is foo/bar/baz.h, your header guard will look like __FOO_BAR_BAZ_H_.

Features I Want
====
- Add some config options, like default author
- Add some command line flags, so you can specify more explicitly how things
  will end up (module name, author, include commented header guard after #endif)
