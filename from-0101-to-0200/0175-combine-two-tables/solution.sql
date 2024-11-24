# Write your MySQL query statement below
select 
    p.firstName as firstName,
    p.lastName as lastName,
    a.city as city,
    a.state as state
from person p
left join address a on p.personId = a.personId
;