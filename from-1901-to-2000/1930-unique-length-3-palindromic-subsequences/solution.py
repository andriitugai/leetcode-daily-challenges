class Solution:
    def countPalindromicSubsequence(self, s: str) -> int:
        firstIndex = [-1] * 26
        lastIndex = [-1] * 26

        for i, c in enumerate(s):
            abcIdx = ord(c) - ord('a')
            if firstIndex[abcIdx] == -1:
                firstIndex[abcIdx] = i
                lastIndex[abcIdx] = i
            else:
                lastIndex[abcIdx] = i

        result = 0
        for i in range(26):
            if firstIndex[i] < lastIndex[i]:
                result += len(set(s[firstIndex[i] + 1:lastIndex[i]]))

        return result