## Human Traffic of Stadium
### SQL Schema
```sql
Create table If Not Exists stadium (id int, visit_date DATE NULL, people int)
Truncate table stadium
insert into stadium (id, visit_date, people) values ('1', '2017-01-01', '10')
insert into stadium (id, visit_date, people) values ('2', '2017-01-02', '109')
insert into stadium (id, visit_date, people) values ('3', '2017-01-03', '150')
insert into stadium (id, visit_date, people) values ('4', '2017-01-04', '99')
insert into stadium (id, visit_date, people) values ('5', '2017-01-05', '145')
insert into stadium (id, visit_date, people) values ('6', '2017-01-06', '1455')
insert into stadium (id, visit_date, people) values ('7', '2017-01-07', '199')
insert into stadium (id, visit_date, people) values ('8', '2017-01-08', '188')
```

X city built a new stadium, each day many people visit it and the stats are saved as these columns: id, visit_date, people

Please write a query to display the records which have 3 or more consecutive rows and the amount of people more than 100(inclusive).

For example, the table stadium:
```
+------+------------+-----------+
| id   | visit_date | people    |
+------+------------+-----------+
| 1    | 2017-01-01 | 10        |
| 2    | 2017-01-02 | 109       |
| 3    | 2017-01-03 | 150       |
| 4    | 2017-01-04 | 99        |
| 5    | 2017-01-05 | 145       |
| 6    | 2017-01-06 | 1455      |
| 7    | 2017-01-07 | 199       |
| 8    | 2017-01-08 | 188       |
+------+------------+-----------+
```
For the sample data above, the output is:
```
+------+------------+-----------+
| id   | visit_date | people    |
+------+------------+-----------+
| 5    | 2017-01-05 | 145       |
| 6    | 2017-01-06 | 1455      |
| 7    | 2017-01-07 | 199       |
| 8    | 2017-01-08 | 188       |
+------+------------+-----------+
```
Note:
Each day only have one row record, and the dates are increasing with id increasing.

### 解析
有一个表stadium记录了体育场每天参观的人数，字段有id, 展馆日期，人数。每天只有一条记录，而且日期和id一样是按照升序拍的。

现要求写一段sql查询出所有连续超过3天参观人数都超过100人的记录。


这里还是要求连续，所以，我们还是得先做个连续的统计：

```sql
select s.id,
        s.visit_date,
        s.people,
        @count := if(s.people >= 100, @count + 1, 0) as count
from stadium s,
    (select @preDate := '',@count := 0) init
```

先统计一下连续超过100人的天数，结果如下：

```
+----+------------+--------+-------+
| id | visit_date | people | count |
+----+------------+--------+-------+
| 1  | 2017-01-01 | 10     | 0     |
| 2  | 2017-01-02 | 109    | 1     |
| 3  | 2017-01-03 | 150    | 2     |
| 4  | 2017-01-04 | 100    | 3     |
| 5  | 2017-01-05 | 90     | 0     |
| 6  | 2017-01-06 | 101    | 1     |
| 7  | 2017-01-07 | 99     | 0     |
| 8  | 2017-01-08 | 100    | 1     |
+----+------------+--------+-------+
```

从这个结果来看，有哪几天人数超过100很好看出来，但是再怎么写sql过滤出来呢？

`where count>=3`?好像不行，这样的话，会丢失连续的前两天的数据。

那怎么办呢？如果将count这一列提换成有多少天连续超过100天，是不是就能继续走了呢？得到的结果类似如下：

```
+----+------------+--------+-------+
| id | visit_date | people | count |
+----+------------+--------+-------+
| 1  | 2017-01-01 | 10     | 0     |
| 2  | 2017-01-02 | 109    | 3     |
| 3  | 2017-01-03 | 150    | 3     |
| 4  | 2017-01-04 | 100    | 3     |
| 5  | 2017-01-05 | 90     | 0     |
| 6  | 2017-01-06 | 101    | 1     |
| 7  | 2017-01-07 | 99     | 0     |
| 8  | 2017-01-08 | 100    | 1     |
+----+------------+--------+-------+
```

那该怎么做呢？计数是从1慢慢开始往上加的，我在读到某一行的时候，并不知道下面还有几行满足人数大于100这个条件啊。

所以从原数据库里直接搞，可能有点难度，那我们就搞上面这个中间结果，怎么搞？

我们按照时间倒序倒过来，这样读这个中间结果的时候，就是先读的连续记录的最大的值，然后将下面所有的值都改成第一个最大值不就可以了？

```sql
select tmp.id,
        tmp.visit_date,
        tmp.people,
        tmp.count,
        @first := if(@pre = 0, tmp.count, @first)     as first,
        @pre := tmp.count                             as pre,
        @countb := if(tmp.count && @first, @first, 0) as countb

from (
        select s.id,
                s.visit_date,
                s.people,
                @count := if(s.people >= 100, @count + 1, 0) as count
        from stadium s,
            (select @preDate := '',@count := 0) init
    ) tmp,
    (select @pre := 0, @countb := 0, @first := 0) initb
order by tmp.visit_date desc
```

统计的结果如下：

```
+----+------------+--------+-------+-------+-----+--------+
| id | visit_date | people | count | first | pre | countb |
+----+------------+--------+-------+-------+-----+--------+
| 8  | 2017-01-08 | 100    | 1.0   | 1     | 1.0 | 1      |
| 7  | 2017-01-07 | 99     | 0.0   | 1     | 0.0 | 0      |
| 6  | 2017-01-06 | 101    | 1.0   | 1     | 1.0 | 1      |
| 5  | 2017-01-05 | 90     | 0.0   | 1     | 0.0 | 0      |
| 4  | 2017-01-04 | 100    | 3.0   | 3     | 3.0 | 3      |
| 3  | 2017-01-03 | 150    | 2.0   | 3     | 2.0 | 3      |
| 2  | 2017-01-02 | 109    | 1.0   | 3     | 1.0 | 3      |
| 1  | 2017-01-01 | 10     | 0.0   | 3     | 0.0 | 0      |
+----+------------+--------+-------+-------+-----+--------+
```

其中countb这一列就是替换后的，要用到的列。好了，剩下的就是，从上面这个临时表中过滤出所有countb>=3的数据即可。


```sql
select result.id, result.visit_date, result.people
from (
       select tmp.id,
              tmp.visit_date,
              tmp.people,
              tmp.count,
              @first := if(@pre = 0, tmp.count, @first)     as first,
              @pre := tmp.count                             as pre,
              @countb := if(tmp.count && @first, @first, 0) as countb

       from (
              select s.id,
                     s.visit_date,
                     s.people,
                     @count := if(s.people >= 100, @count + 1, 0) as count
              from stadium s,
                   (select @preDate := '',@count := 0) init
            ) tmp,
            (select @pre := 0, @countb := 0, @first := 0) initb
       order by tmp.visit_date desc
     ) result
where result.countb >= 3
order by id;
```

写sql和写业务代码是一样的，要想解决这个问题，按部就班，一步一步，先解决什么，再解决什么，慢慢分析，总能分析出来的。

sql很强大，请不要忽视它。