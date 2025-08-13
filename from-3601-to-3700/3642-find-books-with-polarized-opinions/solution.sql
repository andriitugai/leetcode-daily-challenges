# Write your MySQL query statement below
with books_5rs as (
    select 
        book_id,
        max(session_rating) - min(session_rating) as rating_spread,
        round(
            sum(
                case
                    when session_rating >=4 or session_rating <= 2 then 1
                    else 0
                end
            ) / count(*)
        , 2) as polarization_score
    from reading_sessions
    group by book_id
    having max(session_rating) >= 4 and min(session_rating) <= 2 and count(*) >= 5
)
select
    b.book_id,
    b.title,
    b.author,
    b.genre,
    b.pages,
    rs.rating_spread,
    rs.polarization_score
from books b
inner join books_5rs rs on b.book_id = rs.book_id
where rs.polarization_score >= 0.6
order by rs.polarization_score desc, b.title desc
