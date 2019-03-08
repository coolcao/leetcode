## Department Top Three Salaries

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

The Employee table holds all employees. Every employee has an Id, and there is also a column for the department Id.

```
+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 70000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
| 5  | Janet | 69000  | 1            |
| 6  | Randy | 85000  | 1            |
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

Write a SQL query to find employees who earn the top three salaries in each of the department. For the above tables, your SQL query should return the following rows.

```
+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| IT         | Randy    | 85000  |
| IT         | Joe      | 70000  |
| Sales      | Henry    | 80000  |
| Sales      | Sam      | 60000  |
+------------+----------+--------+
```

有两个表，员工表和部门表，现在要求写sql查询出每个部门薪水最高的的三个员工。

这个是之前的求每个部门薪水最高员工的题目的进阶版，上一个只需要一个最高的，现在这个要求前三个。

一步一步来，先将员工表按薪水在部门内排序：

```sql
select e.id,
        e.Salary,
        e.DepartmentId,
        @rank :=
            if(@preDepartmentId <> e.DepartmentId, 1, if(@preSalary = e.salary, @rank, @rank + 1)) as rank,
        @preDepartmentId := e.DepartmentId                                                           as preDepartmentId,
        @preSalary := e.Salary                                                                       as PreSalary
from Employee e,
    (select @preSalary := 0, @rank := 0, @preDepartmentId := 0) init
order by e.DepartmentId, e.Salary desc
```

这里我定义三个变量：preSalary记录前一个薪水，preDepartmentId记录前一个部门id，rank记录在部门内的排名。

其实在部门内排序的根据很简单，先判断当前部门和preDepartmentId是否相同，如果不同，那么排名需要重置为1。如果相同，再根据薪水和前一个薪水是否相同，如果薪水相同，那么说明排名和上一个是一样的，不变，如果不同，那么排名rank+1即可。

**当然，这里所有的前提是，先按照部门和薪水排序，薪水按照降序排序。**

结果临时表如下：（这里我多加了几条数据，方便调试）

```
+----+--------+--------------+------+-----------------+-----------+
| id | Salary | DepartmentId | rank | preDepartmentId | PreSalary |
+----+--------+--------------+------+-----------------+-----------+
| 3  | 99000  | 1            | 1    | 1               | 99000     |
| 1  | 75000  | 1            | 2    | 1               | 75000     |
| 4  | 75000  | 1            | 2    | 1               | 75000     |
| 5  | 60000  | 1            | 3    | 1               | 60000     |
| 7  | 60000  | 1            | 3    | 1               | 60000     |
| 2  | 50000  | 1            | 4    | 1               | 50000     |
| 6  | 50000  | 1            | 4    | 1               | 50000     |
| 8  | 88000  | 2            | 1    | 2               | 88000     |
| 10 | 81000  | 2            | 2    | 2               | 81000     |
| 12 | 81000  | 2            | 2    | 2               | 81000     |
| 11 | 79200  | 2            | 3    | 2               | 79200     |
| 9  | 79000  | 2            | 4    | 2               | 79000     |
+----+--------+--------------+------+-----------------+-----------+
```

这里我们可以看到，rank字段已经按照正确的逻辑进行排名了，接下来要求每个部门的薪水前三名的员工信息，那么只需要把这个临时表和员工表，部门表进行连接查询即可，排名前三的话，只需要使用rank过滤即可。

```sql
select d.Name Department, e.Name Employee, tmp.Salary
from (
       select e.id,
              e.Salary,
              e.DepartmentId,
              @rank :=
                  if(@preDepartmentId <> e.DepartmentId, 1, if(@preSalary = e.salary, @rank, @rank + 1)) as rank,
              @preDepartmentId := e.DepartmentId                                                           as preDepartmentId,
              @preSalary := e.Salary                                                                       as PreSalary
       from Employee e,
            (select @preSalary := 0, @rank := 0, @preDepartmentId := 0) init
       order by e.DepartmentId, e.Salary desc
     ) tmp
       inner join Department d on tmp.DepartmentId = d.Id
       inner join Employee e on tmp.Id = e.Id and tmp.Salary = e.Salary
where tmp.rank <= 3
order by tmp.DepartmentId, tmp.rank;

```




