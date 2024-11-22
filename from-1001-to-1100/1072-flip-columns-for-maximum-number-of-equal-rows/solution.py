class Solution:
    def maxEqualRowsAfterFlips(self, matrix: List[List[int]]) -> int:
        count = defaultdict(int)
        for row in matrix:
            if row[0] == 0:
                key = tuple(row)
            else:
                key = tuple([0 if n else 1 for n in row])

            count[key] += 1
        return max(count.values())