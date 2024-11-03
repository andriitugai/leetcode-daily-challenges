class Solution:
    def rotateString(self, s: str, goal: str) -> bool:
        if len(s) != len(goal):
            return False
        
        for k in range(len(s)):
            if s[k:]+s[:k] == goal:
                return True
            
        return False