class Solution:
    def makeFancyString(self, s: str) -> str:
        result = ""
        for i in range(len(s)):
            if i > 1 and s[i] == s[i - 1] and s[i] == s[i - 2]:
                continue
            result += s[i]

        return result