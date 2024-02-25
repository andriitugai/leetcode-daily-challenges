# Write your MySQL query statement below
with python_skilled as (
    select candidate_id 
    from Candidates 
    where skill = "Python"
), tableau_skilled as (
    select candidate_id 
    from Candidates 
    where skill = "Tableau"
), postgres_skilled as (
    select candidate_id 
    from Candidates 
    where skill = "PostgreSQL"
)
select distinct candidate_id
from Candidates
where candidate_id in (select candidate_id from python_skilled) and
    candidate_id in (select candidate_id from tableau_skilled) and
    candidate_id in (select candidate_id from postgres_skilled)
order by candidate_id;