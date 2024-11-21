import pandas as pd

def second_highest_salary(employee: pd.DataFrame) -> pd.DataFrame:
    salaries = employee['salary'].drop_duplicates()
    sorted_salaries = salaries.sort_values(ascending=False)
    if 2 > len(sorted_salaries):
        return pd.DataFrame({'SecondHighestSalary': [None]})

    n_highest = sorted_salaries.iloc[1]
    return pd.DataFrame({'SecondHighestSalary': [n_highest]})