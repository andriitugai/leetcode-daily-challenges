class Solution:
    def maximumLength(self, s: str) -> int:
        subCount = defaultdict(int)
        for i in range(len(s)):
            j = i
            while j < len(s) and s[i] == s[j]:
                subCount[s[i:j+1]] += 1
                j += 1

        maxLen = -1
        for substr, count in subCount.items():
            if count >= 3 and len(substr) > maxLen:
                maxLen = len(substr)
                print(substr)

        return maxLen