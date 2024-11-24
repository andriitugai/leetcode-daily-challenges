import pandas as pd

def duplicate_emails(person: pd.DataFrame) -> pd.DataFrame:
    email_counts = person.groupby('email').size()
    duplicate_emails = email_counts[email_counts > 1].index
    return pd.DataFrame(duplicate_emails, columns=['email'])