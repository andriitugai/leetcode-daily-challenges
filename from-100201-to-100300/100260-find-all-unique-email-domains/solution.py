import pandas as pd

def find_unique_email_domains(df: pd.DataFrame) -> pd.DataFrame:    
    df['email_domain'] = df[df['email'].str.endswith('.com')]['email'].str.split('@').str[1]
    result = df.groupby('email_domain').size().reset_index(name='count')
    return result.sort_values(by='email_domain')