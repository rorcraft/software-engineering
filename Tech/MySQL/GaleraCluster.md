https://mariadb.com/kb/en/galera/

http://www.percona.com/live/mysql-conference-2013/sessions/how-understand-galera-replication-0

sessions: 
http://www.percona.com/live/mysql-conference-2013/search/node/galera

### Limitations

https://mariadb.com/kb/en/mariadb-galera-cluster-known-limitations/

http://mysql.rjweb.org/doc.php/galera

### Scaled Writing

With 'traditional' replication, especially if using "Statement Based Replication", all writes to the Master are applied to all Slaves, and they are replied serially. Even with a multi-master setup, all writes are applied to all Masters. Hence, there is no way to get "write scaling"; that is, no way to increase the number of writes beyond what a single Master can handle. 

With Galera, there is a moderate degree of write scaling. 
    ⚈  All nodes can perform writes in parallel. 
    ⚈  Writes are replicated via "Row Based Replication", which is decidely faster for single-row operations. 
    ⚈  Writes are applied in parallel on the 'slaves', up to the setting wsrep_slave_threads. This is safe because of the way COMMITs work. 

There is no perfect number for wsrep_slave_threads; it varies with number of CPU cores, client connections, etc. A few dozen is likely to be optimal. Hence, this allows a significant degree of write scaling. If you are I/O-bound, that would be a scaling limitation. 

### AUTO_INCREMENT

By using wsrep_auto_increment_control = ON, the values of auto_increment_increment and auto_increment_offset will be automatically adjusted as nodes come/go. 

Bottom line: There may be gaps in AUTO_INCREMENT values. Consecutive rows, even on one connection, will not have consecutive ids. 

### Certification Based Replication Method

An alternative approach to synchronous database replication using Group Communication and transaction ordering techniques was suggested by a number of researchers (e.g. Database State Machine Approach and Don’t Be Lazy, Be Consistent) and prototype implementations have shown a lot of promise. We combined our experience in synchronous database replication and the latest research in the field to create Galera Replication Toolkit.
Galera replication is a highly transparent and scalable synchronous replication solution for application clustering to achieve high availability and improved performance. Galera-based clusters are:
 
* Highly available
* Highly transparent
* Highly scalable (near linear scalability may be reached depending on the application)
