### File Systems

__Disk__

* tracks > blocks > sectors
* sectors within a block is set for the partition - 1, 2, 4, 8, 16
* The sector size defaults to 4kb on new HDD

__Extents__

* contiguous blocks on HDD to keep files close together and prevent fragmentation
* part of contiguous blocks after a file can be 'reserved'
* extents help reduce fragmentation and minimize the block lists in the inode entry to allow better file system performance.

```
filefrag -v file1.txt
Filesystem type is: ef53
File size of file1.txt is 178353379 (174174 blocks, blocksize 1024)
Ext logical physical expected length flags
0 0 212993 8192
1 8192 223233 221185 32768
2 40960 256001 6144
3 47104 270337 262145 32768
4 79872 303105 32768
5 112640 335873 32768
6 145408 368641 8192
7 153600 378881 376833 14336
8 167936 397313 393217 4096
9 172032 4338 401409 2048
10 174080 137217 638694 eof
file1.txt: 7 extents found
```

__inode (index node)__

* stores all the information about a file system object (file, device node, socket, pipe, etc.). It does not store the file's data content and file name[1] except for certain cases in modern file systems
* unique number within file system identifies each inode. 
* `ls -i` 
* e.g. `find . -inum 782263 -exec rm -i {} \;`

e.g.
> a single ext4 inode has sufficient space to reference four extents (where each extent represents a set of contiguous blocks). For large files (including those that are fragmented), an inode can reference an index node, each of which can reference a leaf node (referencing multiple extents). This constant depth extent tree provides a rich representation scheme for large, potentially sparse files. The nodes also include self-checking mechanisms to further protect against file system corruption.

http://computer-forensics.sans.org/blog/2010/12/20/digital-forensics-understanding-ext4-part-1-extents

__HFS+__

* http://www.linux.org/threads/hierarchical-file-system-plus-hfs.4493/
* http://www.osxbook.com/software/hfsdebug/
* Maximum number of files per filesystem: Four billion.
* Supports volume sizes to 8EiB.
* Supports file sizes to 8EiB.
* Journaling.
* ACL
* hard linking (transparent symlink)
* transparent compression and encryption
* B-tree for directoy metadata
