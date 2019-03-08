## Consecutive Numbers
### SQL Schema
```sql
Create table If Not Exists Logs (Id int, Num int)
Truncate table Logs
insert into Logs (Id, Num) values ('1', '1')
insert into Logs (Id, Num) values ('2', '1')
insert into Logs (Id, Num) values ('3', '1')
insert into Logs (Id, Num) values ('4', '2')
insert into Logs (Id, Num) values ('5', '1')
insert into Logs (Id, Num) values ('6', '2')
insert into Logs (Id, Num) values ('7', '2')
```

Write a SQL query to find all numbers that appear at least three times consecutively.

```
+----+-----+
| Id | Num |
+----+-----+
| 1  |  1  |
| 2  |  1  |
| 3  |  1  |
| 4  |  2  |
| 5  |  1  |
| 6  |  2  |
| 7  |  2  |
+----+-----+
```

For example, given the above Logs table, 1 is the only number that appears consecutively for at least three times.

```
+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+
```

有一个log表，里面存了一些数字，现在要求写一个sql查出所有连续超过至少3次的数字。例如上面一个例子，只有1连续了3次，分别是id=1,2,3。


我们可以通过一个额外的变量来记录相同数字连续重复的次数，先查出重复次数：

```sql
select Num,
        @count := if(@pre = (@pre := Num), @count + 1, 1) as count
from Logs l,
    (select @count := 0, @pre := -100) init
```

这里初始化pre变量时，使用的是 -100， 其实这里不是很恰当，因为题目中并没说这个Num记录的数字是什么范围，如果第一条记录是-100，那有可能会出问题的。

统计出来结果如下：

```
+-----+-------+
| Num | count |
+-----+-------+
| 1   | 1.0   |
| 1   | 2.0   |
| 1   | 3.0   |
| 2   | 1.0   |
| 1   | 1.0   |
| 2   | 1.0   |
| 2   | 2.0   |
+-----+-------+
```

有了这个连续重复出现次数的统计结果，那么就可以直接分组distinct出来了：

```sql
select tmp.Num as ConsecutiveNums
from (
        select Num,
            @count := if(@pre = (@pre := Num), @count + 1, 1) as count
        from Logs l,
            (select @count := 0, @pre := -100) init
    ) tmp
where tmp.count >= 3
group by tmp.Num;
```


