__EXPLAINED EXTENDED__  reverse compile execution plan into SELECT statemnt. run `SHOW WARNINGS` to see exactly how the query optimizer has transformed the statement.

* EXPLAIN will execute sub queries.
* EXPLAIN is an approximation, can be inaccurate.
* 'filesort' - in-memory sorts or temporary files.
* 'Using temporary' - temporary tables on disk or in memory
* Only for SELECT queries. (5.6 can support UPDATE/INSERT as well)

__Id__ - should most always 1 for SIMPLE queries

__select_type__ - SIMPLE, or PRIMARY - SUBQUERY, DERIVED, UNION

__table__ - name of table to select from

__type__ 
* ALL - table scan
* index - table scan in index order (random IO), if 'Using index' in Extra column, it is using covering index (less expensive).
* range - limited index scan
* ref - index lookup - returns rows that match a single value. 
* eq_ref - index lookup, returns at most a single value.
* const, system - turned a query into a constant
* NULL - requires no table access

__possible_keys__ - list created early in optimization phase.

__key__ - which index it decided to use. 

__key_len__ - number of bytes the database will use in the index. deduce the columns used in the index lookup.

__rows__ - estimates number of rows it'll read to find the desired rows.

__filtered__ - pessimistic estimate of the percentage of rows that will satify some condition on the table.

__EXTRA__
* _Using index_ - using covering index
* _Using where_ - post-filter rows after storage engine retrieves them.
* _Using temporary_ - use a temporary table while sorting 
* _Using filesort_ - use an external sort
* _Range checked_ - means not a good index
