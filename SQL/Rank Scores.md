## Rank Scores

### SQL Schema
```sql
Create table If Not Exists Scores (Id int, Score DECIMAL(3,2))
Truncate table Scores
insert into Scores (Id, Score) values ('1', '3.5')
insert into Scores (Id, Score) values ('2', '3.65')
insert into Scores (Id, Score) values ('3', '4.0')
insert into Scores (Id, Score) values ('4', '3.85')
insert into Scores (Id, Score) values ('5', '4.0')
insert into Scores (Id, Score) values ('6', '3.65')
```

Write a SQL query to rank scores. If there is a tie between two scores, both should have the same ranking. Note that after a tie, the next ranking number should be the next consecutive integer value. In other words, there should be no "holes" between ranks.

```
+----+-------+
| Id | Score |
+----+-------+
| 1  | 3.50  |
| 2  | 3.65  |
| 3  | 4.00  |
| 4  | 3.85  |
| 5  | 4.00  |
| 6  | 3.65  |
+----+-------+
```
For example, given the above Scores table, your query should generate the following report (order by highest score):

```
+-------+------+
| Score | Rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+
```

有一个成绩表，里面记录了成绩，现要求写一个sql将成绩按照从大到小排序排名，相同的成绩排名相同。例如上面的例子，排名后得的第二个表格。

如果只是简单的倒序输出，那么很简单，`order by score desc` 即可。可是还要记录排名，这怎么办？

sql是可以自定义变量的，那么，如果可以自定义变量，那么我们可以定义一个排名的变量，通过这个变量来记录排名信息。

```sql
select Score,
       @`rank` := if(@pre = (@pre := Score), @`rank`, @`rank` + 1) Rank
from Scores,
     (select @`rank` := 0, @pre := -1) init
order by Score desc;
```

from 的第二个集合 `(select @`rank` := 0, @pre := -1) init` 即是初始化自定义变量的语句。自定义变量使用 @ 符， 使用 := 进行赋值。这里表示定义两个变量， rank并初始化为0， pre并初始化为-1。
rank变量记录当前的排名，pre变量记录迁移条记录的score值，在查询时，通过判断当前值和前一条记录的score值是否相等来计算排名rank（隐含的逻辑是，要按照score进行倒序排序），如果当前score和前一个score相同，那么rank也就相同，如果不同，由于是倒序排序的，所以一定比前一个小，那么rank排名加1即可。