class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:

        indexes = [-1]*400
        res = 0
        start = 0

        for idx in range(len(s)):
            code = ord(s[idx])

            start = max(start, indexes[code]+1)
            res = max(res, idx - start + 1)

            indexes[code] = idx

        return res
            