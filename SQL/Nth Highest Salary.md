## Nth Highest Salary

### SQL Schema

```sql
Create table If Not Exists Employee (Id int, Salary int)
Truncate table Employee
insert into Employee (Id, Salary) values ('1', '100')
insert into Employee (Id, Salary) values ('2', '200')
insert into Employee (Id, Salary) values ('3', '300'
```

Write a SQL query to get the nth highest salary from the Employee table.

```
+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
```

For example, given the above Employee table, the nth highest salary where n = 2 is 200. If there is no nth highest salary, then the query should return null.

```
+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+
```

### 解析
这个题目是上一个第二高薪水的一个变形，表结构都是一样的。难度上比第二高薪水稍微难了点。

题目要求写一个sql函数，函数接受一个整型的值n，表示第n高的薪水。

逻辑上还是一样，使用limit和offset实现：
```sql
select distinct e.Salary
    from Employee e
    order by e.Salary desc
    limit n,1;
```

这里n是offset，那么要求第n高的，应该是先倒序后offset为n-1,但这里不能使用n-1，由于sql的特殊语法，我们要使用set进行变量赋值。
```sql
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  declare o int;
  set o = N - 1;
  RETURN (
    select distinct e.Salary
    from Employee e
    order by e.Salary desc
    limit o,1
  );
END;
```
