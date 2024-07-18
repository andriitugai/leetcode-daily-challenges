import pandas as pd

def sum_daily_odd_even(transactions: pd.DataFrame) -> pd.DataFrame:
    # Create a mask for odd and even amounts
    transactions['odd_sum'] = transactions.apply(lambda x: x['amount'] if x['amount'] % 2 != 0 else 0, axis=1)
    transactions['even_sum'] = transactions.apply(lambda x: x['amount'] if x['amount'] % 2 == 0 else 0, axis=1)
    
    # Group by 'transaction_date' and sum the 'odd_sum' and 'even_sum' columns
    result = transactions.groupby('transaction_date').agg({
        'odd_sum': 'sum',
        'even_sum': 'sum'
    }).reset_index()

    # Sort the DataFrame by 'transaction_date'
    result = result.sort_values('transaction_date')

    return result