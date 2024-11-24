# Write your MySQL query statement below
-- with cte as (
--     select count(email) as cnt, email 
--     from person
--     group by email
-- )
-- select
--     email
-- from cte
-- where cnt > 1
-- ;

select email
from person 
group by email
having count(email) > 1
;