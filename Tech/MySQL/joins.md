https://mariadb.com/kb/en/introduction-to-joins/

```
SELECT * FROM t1 INNER JOIN t2 ON t1.a = t2.b;
------ ------ 
| a    | b    |
------ ------ 
|    2 |    2 |
------ ------ 
1 row in set (0.00 sec)


SELECT * FROM t1 CROSS JOIN t2;
------ ------ 
| a    | b    |
------ ------ 
|    1 |    2 |
|    2 |    2 |
|    3 |    2 |
|    1 |    4 |
|    2 |    4 |
|    3 |    4 |
------ ------ 
6 rows in set (0.00 sec)


SELECT * FROM t1 LEFT JOIN t2 ON t1.a = t2.b;
------ ------ 
| a    | b    |
------ ------ 
|    1 | NULL |
|    2 |    2 |
|    3 | NULL |
------ ------ 
3 rows in set (0.00 sec)


SELECT * FROM t1 RIGHT JOIN t2 ON t1.a = t2.b;
------ ------ 
| b    | a    |
------ ------ 
|    2 |    2 |
|    4 | NULL |
------ ------ 
2 rows in set (0.00 sec)
```
