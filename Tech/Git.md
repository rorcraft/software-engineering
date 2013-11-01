### Git

_git_, which is British English slang roughly equivalent to "unpleasant person". Torvalds said: "I'm an egotistical bastard, and I name all my projects after myself. First 'Linux', now 'git'."[10][11] The man page describes git as "the stupid content tracker"
 
* http://git-scm.com/book/en/Git-Internals
* http://www.infoq.com/presentations/git-index
* http://en.wikipedia.org/wiki/Git_(software)

"plumbing" commands - low-level commands and designed to be chained together UNIX style

"porcelain" commands - user friendly.

```
$ ls .git
HEAD # points to the branch you currently have checked out
branches/ #  isn’t used by newer versions
config
description # only used by GitWeb
hooks/
index # staging area information
info/ # global exclude file for ignored patterns
objects/ #  stores all the content for your database
refs/ # pointers into commit objects in that data (branches)
```
Git has two data structures: a mutable index (also called stage or cache) that caches information about the working directory and the next revision to be committed; and an immutable, append-only object database.

