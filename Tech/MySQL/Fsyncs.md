### Binlog

* Log of all commands sent to table handler.
* Written before table handlers begins writing data to disk (tablespace - ibdata files)
* Used for replication purposes, manual data recovery. 

### Transaction log

* Internel to innodb
* Writes transaction log -> tablespace -> COMMIT marks complete.

__Redo logs__ are used to apply changes that were made in memory but not flushed to the permanent table records.

__Undo logs__ if transaction is a ROLLBACK, changes will be undone.

On startup:
* Apply _redo log_ first
* Then apply _undo log_

* __Log Sequence Numbers (LSN)__ - unique IDs for each record.
* __Previous Log Sequence Number (Prev LSN)__ - link to last log record
* __Transaction ID__ reference database transaction
* checkpoints - a list of transaction that are open when checkpoint was created.

### Performance

* https://mariadb.com/kb/en/binlog-group-commit-and-innodb_flush_log_at_trx_commit/  
  2 fsyncs instead of 3.

## DoubleWrite

http://www.mysqlperformanceblog.com/2006/08/04/innodb-double-write/

- writes data twice when it writes to table space.
- archive data safety in case of partial page writes.
- log requires pages to be internally consistent. 
- When Innodb flushes pages from buffer pool, it does so by multiple pages.
- several pages are written to double write buffer sequentially, fsync'd
- then pages are written to their real location, fsync'd again.
- inconsistent double write buffer is discarded.
- inconsistent tablespace is recovered from double write buffer.
