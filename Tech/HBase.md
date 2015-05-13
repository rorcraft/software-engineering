### Book

http://hbase.apache.org/book.html

### Map / Keyvalue / Cell
HBase maintains maps of Keys to Values (key -> value). Each of these mappings is called a "KeyValue" or a "Cell". You can find a value by its key... That's it.

The unit of storage in HBase consisting of the following fields:

1. row (rowkey)
2. column family
3. column qualifier
4. timestamp
5. type
6. MVCC version
7. value
 
Uniqueness is determined by the combination of row, column family, column qualifier, timestamp, and type.

The natural comparator will perform a bitwise comparison on row, column family, and column qualifier. Less intuitively, it will then treat the greater timestamp as the lesser value with the goal of sorting newer cells first.

### Log splitting
http://blog.cloudera.com/blog/2012/07/hbase-log-splitting/

The process of grouping the WAL edits by region is called log splitting.
 
### Region splitting
http://hortonworks.com/blog/apache-hbase-region-splitting-and-merging/

HBase stores rows of data in tables. 
Tables are split into chunks of rows called “regions”.
Those regions are distributed across the cluster, hosted and made available to client processes by the RegionServer process.
A region is a continuous range within the key space, meaning all rows in the table that sort between the region’s start key and end key are stored in the same region. Regions are non-overlapping, i.e. a single row key belongs to exactly one region at any point in time. A region is only served by a single region server at any point in time, which is how HBase guarantees strong consistency within a single row#. Together with the -ROOT- and .META. regions, a table’s regions effectively form a 3 level B-Tree for the purposes of locating a row within a table.

A Region in turn, consists of many “Stores”, which correspond to column families. A store contains one memstore and zero or more store files. The data for each column family is stored and accessed separately.

### Hfile

http://blog.cloudera.com/blog/2012/06/hbase-io-hfile-input-output/

MapFile is replaced by HFile: a specific map file implementation for HBase. The idea is quite similar to MapFile, but it adds more features than just a plain key/value file. Features such as support for metadata and the index is now kept in the same file.

The data blocks contain the actual key/values as a MapFile.  For each “block close operation” the first key is added to the index, and the index is written on HFile close.

The HFile format also adds two extra “metadata” block types: Meta and FileInfo.  These two key/value blocks are written upon file close.

The Meta block is designed to keep a large amount of data with its key as a String, while FileInfo is a simple Map preferred for small information with keys and values that are both byte-array. Regionserver’s StoreFile uses Meta-Blocks to store a Bloom Filter, and FileInfo for Max SequenceId, Major compaction key and Timerange info. This information is useful to avoid reading the file if there’s no chance that the key is present (Bloom Filter), if the file is too old (Max SequenceId) or if the file is too new (Timerange) to contain what we’re looking for.
