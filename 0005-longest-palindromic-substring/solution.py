class Solution:
    def longestPalindrome(self, s: str) -> str:

        if len(s) == 1:
            return s

        dp = [[0]*len(s) for _ in range(len(s))]
        max_pal = s[0]
        for i in range(len(s)):
            dp[i][i] = 1
        for i in range(len(s)-1):
            if s[i] == s[i+1]:
                dp[i][i+1] = 1
                max_pal = s[i:i+2]

        for offset in range(2, len(s)):
            i = 0
            while i+offset < len(s):
                if s[i] == s[i+offset] and dp[i+1][i+offset-1] == 1:
                    dp[i][i+offset] = 1
                    if len(max_pal) < offset + 1:
                        max_pal = s[i:i+offset+1]
                i += 1

        return max_pal
            