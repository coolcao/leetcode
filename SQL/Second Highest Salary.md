## Second Highest Salary

### SQL Schema
```sql
Create table If Not Exists Employee (Id int, Salary int)
Truncate table Employee
insert into Employee (Id, Salary) values ('1', '100')
insert into Employee (Id, Salary) values ('2', '200')
insert into Employee (Id, Salary) values ('3', '300')
```

Write a SQL query to get the second highest salary from the Employee table.

```
+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
```

For example, given the above Employee table, the query should return 200 as the second highest salary. If there is no second highest salary, then the query should return null.

```
+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
```


### 解析

这个题目很简单，就是从一个薪水表中，查出薪水第二高的薪水。第一直觉使用limit分页即可。

```sql
SELECT DISTINCT Salary
    FROM Employee
    ORDER BY Salary DESC
    LIMIT 1 OFFSET 1
```

但是代码一提交，未通过测试，什么鬼，查看结果是这样的：
```
Input: {"headers": {"Employee": ["Id", "Salary"]}, "rows": {"Employee": [[1, 100]]}}
Output: {"headers":["Salary"],"values":[]}
Expected:{"headers":["SecondHighestSalary"],"values":[[null]]}
```

当输入只有一个记录时，应该是查不到第二高薪水的，所以返回一个空数组没错呀，为什么期望返回一个null什么鬼？

好吧，没好好看题目，题目里所，如果没有第二高的薪水，那么应该返回一个null。

所以这里再加一个判断即可。

```sql
SELECT IFNULL(
           (SELECT DISTINCT Salary
            FROM Employee
            ORDER BY Salary DESC
            LIMIT 1 OFFSET 1),
           NULL) AS SecondHighestSalary;
```
