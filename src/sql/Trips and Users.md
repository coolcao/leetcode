## Trips and Users

### SQL Schema
```sql
Create table If Not Exists Trips (Id int, Client_Id int, Driver_Id int, City_Id int, Status ENUM('completed', 'cancelled_by_driver', 'cancelled_by_client'), Request_at varchar(50))
Create table If Not Exists Users (Users_Id int, Banned varchar(50), Role ENUM('client', 'driver', 'partner'))
Truncate table Trips
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('1', '1', '10', '1', 'completed', '2013-10-01')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('2', '2', '11', '1', 'cancelled_by_driver', '2013-10-01')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('3', '3', '12', '6', 'completed', '2013-10-01')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('4', '4', '13', '6', 'cancelled_by_client', '2013-10-01')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('5', '1', '10', '1', 'completed', '2013-10-02')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('6', '2', '11', '6', 'completed', '2013-10-02')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('7', '3', '12', '6', 'completed', '2013-10-02')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('8', '2', '12', '12', 'completed', '2013-10-03')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('9', '3', '10', '12', 'completed', '2013-10-03')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('10', '4', '13', '12', 'cancelled_by_driver', '2013-10-03')
Truncate table Users
insert into Users (Users_Id, Banned, Role) values ('1', 'No', 'client')
insert into Users (Users_Id, Banned, Role) values ('2', 'Yes', 'client')
insert into Users (Users_Id, Banned, Role) values ('3', 'No', 'client')
insert into Users (Users_Id, Banned, Role) values ('4', 'No', 'client')
insert into Users (Users_Id, Banned, Role) values ('10', 'No', 'driver')
insert into Users (Users_Id, Banned, Role) values ('11', 'No', 'driver')
insert into Users (Users_Id, Banned, Role) values ('12', 'No', 'driver')
insert into Users (Users_Id, Banned, Role) values ('13', 'No', 'driver')
```

The Trips table holds all taxi trips. Each trip has a unique Id, while Client_Id and Driver_Id are both foreign keys to the Users_Id at the Users table. Status is an ENUM type of (‘completed’, ‘cancelled_by_driver’, ‘cancelled_by_client’).

```
+----+-----------+-----------+---------+--------------------+----------+
| Id | Client_Id | Driver_Id | City_Id |        Status      |Request_at|
+----+-----------+-----------+---------+--------------------+----------+
| 1  |     1     |    10     |    1    |     completed      |2013-10-01|
| 2  |     2     |    11     |    1    | cancelled_by_driver|2013-10-01|
| 3  |     3     |    12     |    6    |     completed      |2013-10-01|
| 4  |     4     |    13     |    6    | cancelled_by_client|2013-10-01|
| 5  |     1     |    10     |    1    |     completed      |2013-10-02|
| 6  |     2     |    11     |    6    |     completed      |2013-10-02|
| 7  |     3     |    12     |    6    |     completed      |2013-10-02|
| 8  |     2     |    12     |    12   |     completed      |2013-10-03|
| 9  |     3     |    10     |    12   |     completed      |2013-10-03| 
| 10 |     4     |    13     |    12   | cancelled_by_driver|2013-10-03|
+----+-----------+-----------+---------+--------------------+----------+
```

The Users table holds all users. Each user has an unique Users_Id, and Role is an ENUM type of (‘client’, ‘driver’, ‘partner’).

```
+----------+--------+--------+
| Users_Id | Banned |  Role  |
+----------+--------+--------+
|    1     |   No   | client |
|    2     |   Yes  | client |
|    3     |   No   | client |
|    4     |   No   | client |
|    10    |   No   | driver |
|    11    |   No   | driver |
|    12    |   No   | driver |
|    13    |   No   | driver |
+----------+--------+--------+
```

Write a SQL query to find the cancellation rate of requests made by unbanned users between Oct 1, 2013 and Oct 3, 2013. For the above tables, your SQL query should return the following rows with the cancellation rate being rounded to two decimal places.

```
+------------+-------------------+
|     Day    | Cancellation Rate |
+------------+-------------------+
| 2013-10-01 |       0.33        |
| 2013-10-02 |       0.00        |
| 2013-10-03 |       0.50        |
+------------+-------------------+
```

> Credits:
>
> Special thanks to @cak1erlizhou for contributing this question, writing the problem description and adding part of the test cases.

### 解析
两个表，一个旅程表Trips和一个Users用户表，旅程表保存了客户id, 司机id, 城市id, 状态（状态是枚举，(‘completed’, ‘cancelled_by_driver’, ‘cancelled_by_client’)）以及时间，还有一个表用户表存储的是用户id,用户是否被禁Banned，角色Role。

现在要求写sql查询在 2013-10-01到2013-10-03这一段时间内，所有被未被禁止用户取消的行程的比率。

上面的例子，id为2的用户已被禁止，先从旅程表中过滤掉id=2的所有旅程，然后再计算每天被取消的订单的比率。2013-10-01总共有三单（id=2的那个Trip由于客户id=2，已被禁，所以忽略），被取消的单子有一条，所以比率是 1/3=0.33。同理，2013-10-02没有取消的，为0，2013-10-03总共2单，取消1单，比率为 1/2=0.5.


所以，按照步骤来，先过滤旅程表，将所有被禁用户的订单先过滤掉：

```sql
SELECT t.* from Trips t
where t.Client_Id in (
  select u.Users_Id
  from Users u
  where u.Role = 'client'
    and u.Banned = 'no'
)
  and t.Driver_Id in (
  select u.Users_Id
  from Users u
  where u.Role = 'driver'
    and u.Banned = 'no'
);
```

然后再从这个有效数据的临时表中，统计计算每天的取消率：

```sql
select t.Request_at Day,
       format(count(if(t.Status != 'completed', true, null)) / count(t.id), 2) + 0 as `Cancellation Rate`
from Trips t
where t.Client_Id in (
  select u.Users_Id
  from Users u
  where u.Role = 'client'
    and u.Banned = 'no'
)
  and t.Driver_Id in (
  select u.Users_Id
  from Users u
  where u.Role = 'driver'
    and u.Banned = 'no'
) and t.Request_at >= '2013-10-01' and t.Request_at <= '2013-10-03'
group by t.Request_at;
```

题目要求时间范围是1号到3号，加一个时间限制。统计的时候，根据状态去分别统计被取消订单数量和总的订单数量，然后计算比率即可。加一个format()函数格式化一下小数，题目要求保留2位小数。