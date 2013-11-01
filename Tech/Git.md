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

Content-addressable filesystem = key-value data store

__Blob__

```
$ echo 'test content' | git hash-object -w --stdin
d670460b4b4aece5915caf5c68d12f560a9fe3e4 # SHA-1 hash
```
```
$ find .git/objects -type f
.git/objects/d6/70460b4b4aece5915caf5c68d12f560a9fe3e4
            ^2 ^38
```
```
$ git cat-file -p d670460b4b4aece5915caf5c68d12f560a9fe3e4
test content
```
```
$ git cat-file -t d670460b4b4aece5915caf5c68d12f560a9fe3e4
blob
```

__Tree__

```
$ git cat-file -p master^{tree} # last commit of master branch
100644 blob a906cb2a4a904a152e80877d4088654daad0c859      README
100644 blob 8f94139338f9404f26296befa88755fc2598c289      Rakefile
040000 tree 99f1a6d12cb4b6f19c8655fca46c3ecf317074e0      lib
```

```
$ echo 'new file' > new.txt
$ git update-index --add new.txt # stage
$ git write-tree # write to tree object
0155eb4229851634a0f03eb265b69f5a2d56f341
$ git cat-file -p 0155eb4229851634a0f03eb265b69f5a2d56f341 
100644 blob fa49b077972391ad58037050f2a75f74e3671e92      new.txt
```

Read trees into your staging area by calling `read-tree`.
 
__Commit Objects__



Git has two data structures: a mutable index (also called stage or cache) that caches information about the working directory and the next revision to be committed; and an immutable, append-only object database.

