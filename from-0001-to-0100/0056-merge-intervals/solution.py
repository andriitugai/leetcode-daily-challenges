from typing import List

class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda x: x[0])
        result = []
        start, end = intervals[0]
        for s, e in intervals[1:]:
            if s <= end:
                if e > end:
                    end = e
            else:
                result.append([start, end])
                start = s
                end = e

        result.append([start, end])

        return result