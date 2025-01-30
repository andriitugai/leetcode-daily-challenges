import pandas as pd
import re

def find_valid_emails(users: pd.DataFrame) -> pd.DataFrame:
    pattern = r"^[\w]+@[a-zA-Z].*\.com"
    valid_emails = users[users['email'].str.match(pattern, na=False)]
    return valid_emails.sort_values(by='user_id').reset_index(drop=True)