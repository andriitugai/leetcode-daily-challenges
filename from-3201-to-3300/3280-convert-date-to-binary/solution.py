class Solution:
    def convertDateToBinary(self, date: str) -> str:
        yy, mm, dd = map(int, date.split('-'))
        return f"{yy:b}-{mm:b}-{dd:b}"