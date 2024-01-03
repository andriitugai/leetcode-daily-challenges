import pandas as pd

def getDataframeSize(players: pd.DataFrame) -> List[int]:
    return list(players.shape)
    # num_rows, num_columns = players.shape
    # return [num_rows, num_columns]
    