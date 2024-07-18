# Write your MySQL query statement below
with odd_transactions as (
    select 
        transaction_date,
        sum(amount) as odd_sum
    from transactions
    where amount % 2 != 0
    group by transaction_date
), even_transactions as (
    select 
        transaction_date,
        sum(amount) as even_sum
    from transactions
    where amount % 2 = 0
    group by transaction_date
), odd_even_sums as (
    select 
        coalesce(o.transaction_date, e.transaction_date) as transaction_date,
        coalesce(o.odd_sum, 0) as odd_sum,
        coalesce(e.even_sum, 0) as even_sum
    from odd_transactions as o
    left outer join even_transactions as e
        on o.transaction_date = e.transaction_date
    union
    select 
        coalesce(o.transaction_date, e.transaction_date) as transaction_date,
        coalesce(o.odd_sum, 0) as odd_sum,
        coalesce(e.even_sum, 0) as even_sum
    from odd_transactions as o
    right outer join even_transactions as e
        on o.transaction_date = e.transaction_date
)
select transaction_date, odd_sum, even_sum
from odd_even_sums
order by transaction_date
;