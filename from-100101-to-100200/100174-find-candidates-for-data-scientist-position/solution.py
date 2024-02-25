import pandas as pd

def find_candidates(candidates: pd.DataFrame) -> pd.DataFrame:
    candidates = candidates[candidates['skill'].isin(['Python','Tableau','PostgreSQL'])]
    candidates = candidates.groupby(['candidate_id']).count().reset_index()
    return candidates[candidates['skill']==3][['candidate_id']].sort_values('candidate_id')