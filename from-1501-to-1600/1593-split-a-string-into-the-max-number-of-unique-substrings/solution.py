class Solution:
    def maxUniqueSplit(self, s: str) -> int:
        def dfs(i, currSet):
            if i == len(s):
                return 0
            result = 0
            for j in range(i, len(s)):
                substr = s[i:j + 1]
                if substr in currSet:
                    continue
                currSet.add(substr)
                result = max(result, dfs(j + 1, currSet) + 1)
                currSet.remove(substr)

            return result
        return dfs(0, set())