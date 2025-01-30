# Write your MySQL query statement below
select
    user_id,
    email
from users
where regexp_like(email, "^\\w+@[a-zA-Z].*\\.com$")
order by user_id;