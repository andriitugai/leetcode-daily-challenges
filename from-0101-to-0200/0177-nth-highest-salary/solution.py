import pandas as pd

def nth_highest_salary(employee: pd.DataFrame, N: int) -> pd.DataFrame:
    salaries = employee['salary'].drop_duplicates()
    sorted_salaries = salaries.sort_values(ascending=False)
    if N > len(sorted_salaries):
        return pd.DataFrame({'Nth Highest Salary': [None]})

    n_highest = sorted_salaries.iloc[N - 1]
    return pd.DataFrame({'Nth Highest Salary': [n_highest]})