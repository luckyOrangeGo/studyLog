# MySQL Learn

## 连接 MySQL 软件

mysql -h localhost -u root -p

-h 代表IP地址 -u 用户名 -p 密码

## 数据库操作

net strat mysql

net stop mysql

### 基本操作

#### 登陆数据库

	-D，--database=name打开指定数据库

	-delimiter=name指定分隔符

	-h，--host=name服务器名称

	-p，--password【=name】密码

	-P，--port=#端口号

	-prompt=name设置提示符

	-u，--user=name用户名

	-V，--version输出版本信息并且退出

- MySQL退出

		mysql>exit；
		mysql>quit；
		mysql>\q；

#### 修改MySQL提示符

连接客户端时通过参数指定

	she11>mysq1-uroot-proot--prompt 提示符

连接上客户端后，通过prompt命令修改

	mysq1>prompt提示符

参数 | 描述
------|------
\D | 完整的日期
\d | 当前数据库
\h | 服务器名称
\u | 当前用户

例如：

	mysq1>PROMPT \u@\h-\d>

### MySQL常用命令

#### MySQL语句的规范

- 关键字与函数名称全部大写

- 数据库名称、表名称、字段名称全部小写

- SQL语句必须以分号结尾


#### 显示当前服务器版本

	SELECT VERSION();

#### 显示当前日期时间

	SELECT NOW();

#### 显示当前用户

	SELECT USER();

#### 创建数据库

    mysql> CREATE DATABASE databaseTest;

创建成功会显示
>Query OK, 1 row affected (0.06 sec)

	CREATE { DATABASE | SCHEMA } [IF NOT EXISTS] db_name  [DEFAULT] CHARACTER SET [=] charset_name
创建数据表

	mysq1>CREATE TABLE tb1(
	->username VARCHAR(20), 
	->age TINYINT UNSIGNED, 
	->salary FLOAT(8,2)UNSIGNED ->);

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

完整命令

	SHOW {DATABASES I SCHEMAS}[LIKE 'pattern'I WHERE expr]

查看警告信息

	SHOW WARNINGS;
	
查看创建信息

	SHOW CREATE DATABASE db_name;

#### 选择数据库

    mysql> USE mysql

显示
>Database changed

---

#### 修改数据库

	ALTER{DATABASE I SCHEMA} [db_name] [DEFAULT] CHARACTER SET [=] charset_name


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