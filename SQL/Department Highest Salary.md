## Department Highest Salary

### SQL Schema

```sql
Create table If Not Exists Employee (Id int, Name varchar(255), Salary int, DepartmentId int)
Create table If Not Exists Department (Id int, Name varchar(255))
Truncate table Employee
insert into Employee (Id, Name, Salary, DepartmentId) values ('1', 'Joe', '70000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('2', 'Henry', '80000', '2')
insert into Employee (Id, Name, Salary, DepartmentId) values ('3', 'Sam', '60000', '2')
insert into Employee (Id, Name, Salary, DepartmentId) values ('4', 'Max', '90000', '1')
Truncate table Department
insert into Department (Id, Name) values ('1', 'IT')
insert into Department (Id, Name) values ('2', 'Sales')
```

The Employee table holds all employees. Every employee has an Id, a salary, and there is also a column for the department Id.

```
+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 70000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
+----+-------+--------+--------------+
```
The Department table holds all departments of the company.

```
+----+----------+
| Id | Name     |
+----+----------+
| 1  | IT       |
| 2  | Sales    |
+----+----------+
```

Write a SQL query to find employees who have the highest salary in each of the departments. For the above tables, Max has the highest salary in the IT department and Henry has the highest salary in the Sales department.

```
+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| Sales      | Henry    | 80000  |
+------------+----------+--------+
```

## 解析
有两个表，员工表Employee和部门表Department，分别存储着员工的基本信息和部门名称信息。

写一个sql查询出每个部门的最高薪水的员工名以及其薪水。

这很简单啊，对吧，我们先查第一个表，查出每个部门的最高薪水的人的id，以及其薪水，以及其所在的部门id，然后连接部门表即可。
首先我们做第一步，先查Employee表中每个部门的最高薪水的员工信息。
```sql
select e.Id,
       e.name,
       max(e.Salary) max,
       e.DepartmentId
from Employee e
group by DepartmentId;
```

可是，当你运行的时候，你发现，竟然报错了？？？

```
[42000][1055] Expression #1 of SELECT list is not in GROUP BY clause and contains nonaggregated column 'leetcode.e.Id' which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql_mode=only_full_group_by
```

报错信息说，查询语句里的id字段，不适用于当前的`sql_mode=only_full_group_by`，这是什么鬼？？？

首先，这里我们使用了`group by`进行分组，分组那么就涉及到聚合操作，而id, name在此次查询中并未涉及到聚合操作，因此在`sql_mode=only_full_group_by`模式下会报错。

什么意思呢？当`sql_mode=only_full_group_by`时，我们所有的分组，只能有聚合操作的列才能正确查询。换句话说，select 后面的列只能用group by后面的列以及聚合函数。

为什么会这样？从这个例子中，我们可能看不出什么来，我从一个表中，使用max()函数查找一个列的最大值，并将这一行数据对应的其他列一并取出来似乎逻辑上没什么毛病。首先需要明确的是，max()函数是聚合函数，在这里使用是没问题的。

这里从逻辑上将似乎没什么问题是因为使用的是max()函数，而max()函数返回的是数据表中某一个具体的行数据。如果这里换成求平均值函数avg()函数，你就可能明白了。

例如，我们要求每个部门的平均薪水，按部门输出。这里的平均薪水是将部门所有员工的薪水全加起来然后除以员工数，并不对应着某一具体行数据，那么这个时候，如果你的select后面跟某些具体的列，比如：

```sql
select avg(e.Salary),
       e.DepartmentId,
       e.Id
from Employee e
group by e.DepartmentId;
```
那么，这个id这一列，我该使用哪个呢？所有员工的id的平均数，还是第一个员工的Id 还是最后一个员工的id，没人知道。这也就是说，`sql_mode=only_full_group_by`时，select不能接非聚合函数的列的原因了。

其实这个是可以改的，具体怎么改可以上网查，有设置方法。 **但是建议，不要改，就使用only_full_group_by模式，原因很简单，让所有查询保持逻辑清晰。**

所以，这里我们只能查出这样的结果：
```sql
select e.DepartmentId,
        max(e.Salary) max
from Employee e
group by e.DepartmentId
```
我们先分组查询出每个部门的最高薪水，按照部门id输出：

```
+--------------+-------+
| DepartmentId | max   |
+--------------+-------+
| 1            | 90000 |
| 2            | 80000 |
+--------------+-------+
```

然后再将这个临时的表和Employee和Department表进行联合，查询出对应的员工以及部门名称：

```sql
select d.Name Department,e.Name Employee, e.Salary
from (
       select e.DepartmentId,
              max(e.Salary) max
       from Employee e
       group by e.DepartmentId
     ) tmp
       inner join Employee e on tmp.DepartmentId = e.DepartmentId and tmp.max = e.Salary
       inner join Department d on tmp.DepartmentId = d.Id
order by e.Id desc;
```

这里要注意左连接，右连接以及内连接外连接的区别。不同的连接方式，结果会有很大的不同。 