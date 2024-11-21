# Write your MySQL query statement below
with maxSalary as (
    select departmentId, max(salary) as maxSalary
    from employee
    group by departmentId
)
select
    d.name as Department,
    e.name as Employee,
    s.maxSalary as Salary
from employee e
join department d on d.id = e.departmentId
join maxSalary s on s.departmentId = d.Id and s.maxSalary = e.salary
;