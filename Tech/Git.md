### Git

_git_, which is British English slang roughly equivalent to "unpleasant person". Torvalds said: "I'm an egotistical bastard, and I name all my projects after myself. First 'Linux', now 'git'."[10][11] The man page describes git as "the stupid content tracker"
 
* http://git-scm.com/book/en/Git-Internals
* http://www.infoq.com/presentations/git-index
* http://en.wikipedia.org/wiki/Git_(software)
* https://github.com/pluralsight/git-internals-pdf/releases (Peepcode git book)
* http://www.aosabook.org/en/git.html

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

### Protocol

__Dump HTTP__ (github stopped supporting since 2011)

* Does not require a git server, relies on .git directory structure.

```
=> GET info/refs 
ca82a6dff817ec66f44342007202690a93763949     refs/heads/master
=> GET HEAD
ref: refs/heads/master
=> GET objects/ca/82a6dff817ec66f44342007202690a93763949 # because thats the head
(179 bytes of binary data)
=> GET objects/08/5bb3bcb608e1e8451d4b2432f8ecbe6306e7e7 # parent commit
(179 bytes of data)
=> GET objects/cf/da3bf379e4f8dba8717dee55aab78aef7f4daf
(404 - Not Found) # could be in an alternate repository
=> GET objects/info/http-alternates
(empty file)
=> GET objects/info/packs
P pack-816a9b2334da9953e530f27bcac22082a9f5b835.pack
=> GET objects/pack/pack-816a9b2334da9953e530f27bcac22082a9f5b835.idx
(4k of binary data)
=> GET objects/pack/pack-816a9b2334da9953e530f27bcac22082a9f5b835.pack
(13k of binary data)
```

__Smart HTTP__

http://git-scm.com/blog/2010/03/04/smart-http.html

CGI script that is provided with Git called `git-http-backend` on the server.

__SSH__

__Upload__

* `send-pack` runs on client
* `receive-pack` runs on server

```
$ git push origin master # runs send-pack, like below
$ ssh -x git@github.com "git-receive-pack 'schacon/simplegit-progit.git'"
005bca82a6dff817ec66f4437202690a93763949 refs/heads/master report-status delete-refs
003e085bb3bcb608e1e84b2432f8ecbe6306e7e7 refs/heads/topic
0000
```
* Each line starts with a 4-byte hex value specifying how long the rest of the line is. (005b = 91 bytes)
* `send-pack` process determines what commits it has that the server doesn’t. 
* `send-pack` replies with each reference that this push will update

```
0085ca82a6dff817ec66f44342007202690a93763949  15027957951b64cf874c3557a0f3547bd83b3ff6 refs/heads/master report-status
00670000000000000000000000000000000000000000 cdfdb42577e2506715f8cfeacdbabc092bf63e8d refs/heads/experiment
0000
```
* Next, the client uploads a packfile of all the objects the server doesn’t have yet. 
* Finally, the server responds with a success (or failure) indication:

```
000Aunpack ok
```

__Download__

* `fetch-pack` runs on client
* `upload-pack` runs on server
* Git daemon, which listens on a server on port 9418 by default.

In either case, after fetch-pack connects, upload-pack sends back something like this:
```
0088ca82a6dff817ec66f44342007202690a93763949 HEAD\0multi_ack thin-pack \
  side-band side-band-64k ofs-delta shallow no-progress include-tag
003fca82a6dff817ec66f44342007202690a93763949 refs/heads/master
003e085bb3bcb608e1e8451d4b2432f8ecbe6306e7e7 refs/heads/topic
0000 
```
`upload-pack` replies:
```
0054want ca82a6dff817ec66f44342007202690a93763949 ofs-delta
0032have 085bb3bcb608e1e8451d4b2432f8ecbe6306e7e7
0000
0009done
```

### gc

If you run git gc, you’ll no longer have these files in the refs directory. Git will move them for the sake of efficiency into a file named .git/packed-refs that looks like this:
```
$ cat .git/packed-refs
# pack-refs with: peeled
cac0cab538b970a37ea1e769cbbde608743bc96d refs/heads/experiment
ab1afef80fac8e34258ff41fc1b867c702daa24b refs/heads/master
cac0cab538b970a37ea1e769cbbde608743bc96d refs/tags/v1.0
9585191f37f7b0fb9444f35a9bf50de191beadc2 refs/tags/v1.1
^1a410efbd13591db07496601ebc7a059dd55cfe9 # tag directly above is an annotated tag
```

### reflog & recovery

* Git silently records what your HEAD is every time you change it. Each time you commit or change branches, the reflog is updated. The reflog is also updated by the `git update-ref` command.
* `.git/logs/`
* `git log -g` - normal log output for your reflog
* You can recover commits by creating a new branch at that commit: `git branch (branchname) (sha)`

```
$ git fsck --full # shows all objects that aren't pointed to another object
dangling blob d670460b4b4aece5915caf5c68d12f560a9fe3e4
dangling commit ab1afef80fac8e34258ff41fc1b867c702daa24b
dangling tree aea790b9a58f6cf6f2804eeac9f0abbe9631e4c9
dangling blob 7108f7ecb345ee9d0084193f147cdad4d2998293
```

```
$ git count-objects -v
count: 4
size: 16
in-pack: 21
packs: 1
size-pack: 2016 # size of your packfiles in kilobyte
prune-packable: 0
garbage: 0
```

remove file from history
```
$ git filter-branch --index-filter \
   'git rm --cached --ignore-unmatch git.tbz2' -- 6df7640^..   # remove from index instead of file, which require checkout
Rewrite 6df764092f3e7c8f5f94cbe08ee5cf42e92a0289 (1/2)rm 'git.tbz2'
Rewrite da3f30d019005479c99eb4c3406225613985a1db (2/2)
Ref 'refs/heads/master' was rewritten
```
see more filter-branch: http://git-scm.com/book/ch6-4.html

### Merges

http://en.wikipedia.org/wiki/Three-way_merge#Three-way_merge

### Pack heuristics

* http://git.kernel.org/cgit/git/git.git/tree/Documentation/technical/pack-heuristics.txt?id=HEAD
* http://stackoverflow.com/questions/5176225/are-gits-pack-files-deltas-rather-than-snapshots

### Rebase onto

http://pivotallabs.com/git-rebase-onto/

`git rebase --onto (new base) (old base to cut from) [optional](new head to replace)`
`git rebase --onto topicC topicA topicB`
