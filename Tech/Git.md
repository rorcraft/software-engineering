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

### 3 main object types: Blob, Tree, Commit, + Tag.

__Content-addressable filesystem = key-value data store__

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

```
$ echo 'first commit' | git commit-tree d8329f # tree object
fdf4fc3344e67ab068f836878b6c4951e3b15f3d
$ echo 'second commit' | git commit-tree 0155eb -p fdf4fc3 # tree object on top of last commit object
cac0cab538b970a37ea1e769cbbde608743bc96d
```

![](http://git-scm.com/figures/18333fig0903-tn.png)

__Data Storage Format__

* header that starts with the type of the object, in this case a blob. Then, it adds a space followed by the size of the content and finally a null byte:
* concatenates the header and the original content and then calculates the SHA-1 checksum of that new content
* compresses the new content with zlib

```
=> "blob 16\000what is up, doc?"
=> "bd9dbf5aae1a3862dd1526723246b20206e5fc37"
=> "x\234K\312\311OR04c(\317H,Q\310,V(-\320QH\311O\266\a\000_\034\a\235"
```

### Branching == References
```
git update-ref refs/heads/master 1a410efbd13591db07496601ebc7a059dd55cfe9
# same as
echo "1a410efbd13591db07496601ebc7a059dd55cfe9" > .git/refs/heads/master 
# same as
git branch master 
```

`git commit`, it creates the commit object, specifying the parent of that commit object to be whatever SHA-1 value the reference in HEAD points to.

```
$ git symbolic-ref HEAD refs/heads/test # sets or print the content of .git/HEAD
$ cat .git/HEAD
ref: refs/heads/test
```

Remote references differ from branches (refs/heads references) mainly in that they can’t be checked out. Git moves them around as bookmarks to the last known state of where those branches were on those servers.

__Tags__

Tag is — a branch that never moves. Tag object points to a commit, it contains a tagger, a date, a message, and a pointer.

You can tag any Git object. In the Git source code, for example, the maintainer has added their GPG public key as a blob object and then tagged it. You can view the public key by running

`$ git cat-file blob junio-gpg-pub`

### Packfiles

The initial format in which Git saves objects on disk is called a loose object format. However, occasionally Git packs up several of these objects into a single binary file called a packfile in order to save space and be more efficient. Git does this if you have too many loose objects around, if you run the git gc command manually, or if you push to a remote server. 

* The packfile is a single file containing the contents of all the objects that were removed from your filesystem. 
* The index is a file that contains offsets into that packfile so you can quickly seek to a specific object.

```
.git/objects/info/packs
.git/objects/pack/pack-7a16e4488ae40c7d2bc56ea2bd43e25212a66c45.idx
.git/objects/pack/pack-7a16e4488ae40c7d2bc56ea2bd43e25212a66c45.pack
```

>  second version of the file is the one that is stored intact, whereas the original version is stored as a delta
Git has two data structures: a mutable index (also called stage or cache) that caches information about the working directory and the next revision to be committed; and an immutable, append-only object database.

```
$ git verify-pack -v \
  .git/objects/pack/pack-7a16e4488ae40c7d2bc56ea2bd43e25212a66c45.idx
```

### Refspec

```
$ git log origin/master
$ git log remotes/origin/master
$ git log refs/remotes/origin/master # same, auto expanded
```

```
[remote "origin"]
       url = git@github.com:schacon/simplegit-progit.git
       fetch = +refs/heads/*:refs/remotes/origin/* # every branch
       fetch = +refs/heads/qa/*:refs/remotes/origin/qa/* # namespace
       fetch = +refs/heads/master:refs/remotes/origin/master # just master
       push = refs/heads/master:refs/heads/qa/master # auto push to qa namespace
```
`+` update the reference even if it isn’t a fast-forward
`*` every branch from remote 
