class Solution:
    def maxDepth(self, s: str) -> int:
        ans = 0
        stack = 0
        for c in s:
            if c == '(':
                stack += 1
                ans = max(ans, stack)
            elif c == ')':
                stack -= 1
                
        return ans