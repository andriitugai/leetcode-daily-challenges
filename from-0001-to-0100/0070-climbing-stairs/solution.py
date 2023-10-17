class Solution:
    def climbStairs(self, n: int) -> int:

        f1, f2 = 1, 1
        
        for _ in range(n-1):
            f1, f2 = f1 + f2, f1
        
        return f1
