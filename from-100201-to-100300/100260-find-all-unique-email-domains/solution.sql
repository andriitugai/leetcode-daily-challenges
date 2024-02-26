# Write your MySQL query statement below
select 
    substring(email from locate('@', email) + 1) as email_domain,
    count(*) as count
from Emails
where email like '%.com'
group by substring(email from locate('@', email) + 1)
order by 1;