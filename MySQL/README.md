# MySQL Learn

## 连接 MySQL 软件

mysql -h localhost -u root -p

-h 代表IP地址 -u 用户名 -p 密码

## 数据库操作

net strat mysql

net stop mysql

### 基本操作

#### 创建数据库

    mysql> CREATE DATABASE databaseTest;

创建成功会显示
>Query OK, 1 row affected (0.06 sec)

---

#### 查看数据库

    mysql> SHOW DATABASES;

显示

    +--------------------+
    | Database           |
    +--------------------+
    | databasetest       |
    | information_schema |
    | mysql              |
    | performance_schema |
    | sakila             |
    | sys                |
    | world              |
    +--------------------+
    7 rows in set (0.00 sec)

---

#### 选择数据库

    mysql> USE mysql

显示
>Database changed

---

#### 删除数据库

    mysql> DROP DATABASE databasetest;

显示
>Query OK, 0 rows affected (0.04 sec)

### 储存引擎和数据类型

#### 支持引擎显示

    mysql> SHOW ENGINES \g

显示（格式化）

    +--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
    | Engine             | Support | Comment                                                        | Transactions | XA   | Savepoints |
    +--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
    | MEMORY             | YES     | Hash based, stored in memory, useful for temporary tables      | NO           | NO   | NO         |
    | MRG_MYISAM         | YES     | Collection of identical MyISAM tables                          | NO           | NO   | NO         |
    | CSV                | YES     | CSV storage engine                                             | NO           | NO   | NO         |
    | FEDERATED          | NO      | Federated MySQL storage engine                                 | NULL         | NULL | NULL       |
    | PERFORMANCE_SCHEMA | YES     | Performance Schema                                             | NO           | NO   | NO         |
    | MyISAM             | YES     | MyISAM storage engine                                          | NO           | NO   | NO         |
    | InnoDB             | DEFAULT | Supports transactions, row-level locking, and foreign keys     | YES          | YES  | YES        |
    | BLACKHOLE          | YES     | /dev/null storage engine (anything you write to it disappears) | NO           | NO   | NO         |
    | ARCHIVE            | YES     | Archive storage engine                                         | NO           | NO   | NO         |
    +--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
    9 rows in set (0.00 sec)

#### 显示列

    mysql> SHOW ENGINES \G


    *************************** 1. row ***************************
        Engine: MEMORY
        Support: YES
        Comment: Hash based, stored in memory, useful for temporary tables
    Transactions: NO
            XA: NO

    ...
    ...
    ...

    ...
    ...
    ...

    Savepoints: NO
    *************************** 9. row ***************************
        Engine: ARCHIVE
        Support: YES
        Comment: Archive storage engine
    Transactions: NO
            XA: NO
    Savepoints: NO
    9 rows in set (0.00 sec)

### CMD操作

#### 重启MySQL服务

service mysql restart