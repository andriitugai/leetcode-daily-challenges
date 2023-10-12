class Solution:
    def longestValidParentheses(self, s: str) -> int:
        if not s or len(s) < 2:
            return 0
        
        dp = [0] * len(s)
        for i in range(1, len(s)):
            if s[i] == ')':
                if s[i-1] == '(':
                    prev_len = 0 if i < 1 else dp[i-2]
                    dp[i] = prev_len + 2
                else:           # s[i-1] == ')'
                    if i > dp[i-1] and s[i - dp[i-1] - 1] == '(':
                        dp[i] = dp[i-1] + 2 + dp[i - dp[i-1] - 2]
        
        return max(dp)
            