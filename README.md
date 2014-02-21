Header Generator
====

This program generates header guards and a little doxygen for C header files (probably fine for hpp too, haven't devoted any thought to that).

Installing
====

Clone the repository, and then do a make. This will generate an executable hgen file. Put this somewhere that's in your path (I use $HOME/bin).

Using in a Project
====

To use this in a project, create an empty .hgenconfig file in the root of your project directory. For example, if your project directory is located at /home/shaladdle/fav-project, you can do

```
touch /home/shaladdle/fav-project/.hgenconfig
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
 *  @brief Function prototypes for foo.
 *
 *  @author [your name here]
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
