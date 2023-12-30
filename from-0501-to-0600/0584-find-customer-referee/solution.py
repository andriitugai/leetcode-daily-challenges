import pandas as pd

def find_customer_referee(customer: pd.DataFrame) -> pd.DataFrame:
    customer['referee_id'] = customer['referee_id'].fillna(0)
    df = customer[(customer['referee_id'] != 2)]
    return df[['name']]