# Write your MySQL query statement below
with ranked as (
    select
           status_time as s_time,
           session_status as s_status,
           lead(status_time, 1) over (partition by server_id order by status_time) as e_time,
           lead(session_status, 1) over (partition by server_id order by status_time) as e_status
    from servers
)
select floor(sum(timestampdiff(second, s_time, e_time) / 24 / 3600)) as total_uptime_days
from ranked
where s_status = 'start';
