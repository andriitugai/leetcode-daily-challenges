class Solution:
    def largestCombination(self, candidates: List[int]) -> int:
        result = 0
        bitmask = 1
        for i in range(32):
            count = 0
            for n in candidates:
                count += 1 if bitmask & n else 0
            result = max(result, count)
            bitmask <<= 1
        return result