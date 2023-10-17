from typing import List

class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        result = []
        start, end = newInterval

        for interval in intervals:
            s, e = interval
            if s > end:
                result.append([start, end])
                start, end = s, e
            elif e < start:
                result.append(interval)
            else:
                start = min(s,start)
                end = max(e, end)
            
        result.append([start, end])
        return result
        