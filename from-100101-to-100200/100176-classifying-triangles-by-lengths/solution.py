import pandas as pd

def get_type(row):
    a, b, c = row['A'], row['B'], row['C']
    if a + b <= c or a + c <= b or b + c <= a:
        return 'Not A Triangle'
    if a == b and b == c:
        return 'Equilateral'
    if a == b or b == c or a == c:
        return 'Isosceles'
    else:
        return 'Scalene'
        
def type_of_triangle(triangles: pd.DataFrame) -> pd.DataFrame:
    triangles['triangle_type'] = triangles.apply(get_type, axis='columns')
    df = triangles[['triangle_type']]
    return df