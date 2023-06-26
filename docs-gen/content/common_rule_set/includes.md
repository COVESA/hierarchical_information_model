---
title: "Includes"
date: 2019-08-04T12:59:44+02:00
weight: 6
---

An include directive in a HIM file will read the file it refers to and the
contents of that file will be inserted into the current buffer in place of the
include directive.  The included file will, in its turn, be scanned for
include directives to be replaced, effectively allowing formation of a tree of includedfiles.

See the figurefor an example of such a tree.

![Include directive](/hierarchical_information_model/images/include_directives.png?width=50pc)<br>
*Figure: Include directives*


The include directive has the following format:

    #include <filename> [prefix]

The ```<filename>``` part specifies the path, relative to the file with the ```#include``` directive, to the vspec file to replace the directive with.

The optional ```[prefix]``` specifies a branch name to be
prepended to all signal entries in the included file. This allows a him file
to be reused multiple times by different files, each file specifying their
own branch to attach the included file to.

An example of an include directive is:

    #include doors.him chassis.doors

The ```doors.him``` section specifies the file to include.

The ```chassis.doors``` section specifies that all signal entries in ```doors.him``` should have their names prefixed with ```chassis.doors```.

If an included him file has node specifications that have
already been defined prior to the included file, the new specifications in the
included file will override the previous specifications.

Complete subtrees can be reused by including them multiple times to different branches.
