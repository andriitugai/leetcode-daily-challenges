import pandas as pd

def rising_temperature(weather: pd.DataFrame) -> pd.DataFrame:
    weather.sort_values(by='recordDate', inplace=True)

    temp_diff = weather['temperature'].diff()
    date_diff = weather['recordDate'].diff()

    rise_temp_id = weather[
        (temp_diff > 0) & (date_diff == pd.Timedelta(days=1))
    ]

    return pd.DataFrame(rise_temp_id['id'])