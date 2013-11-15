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
