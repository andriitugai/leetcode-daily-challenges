class Solution:
    def tribonacci(self, n: int) -> int:
        tribo = [0, 1, 1] + [0] * (n-2)
        if n < 3:
            return tribo[n]
        
        for i in range(3, n+1):
            tribo[i] = tribo[i-3] + tribo[i-2] + tribo[i-1]
            
        return tribo[-1]
    