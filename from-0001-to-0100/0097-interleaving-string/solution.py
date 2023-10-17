class Solution:
    def isInterleave(self, s1: str, s2: str, s3: str) -> bool:
#         if not s1 and not s2 and not s3:
#             return True
        
#         if not s3:
#             return False
        
#         if s1 and s1[0] == s3[0]:
#             return self.isInterleave(s1[1:], s2, s3[1:])
        
#         if s2 and s2[0] == s3[0]:
#             return self.isInterleave(s1, s2[1:], s3[1:])
        
#         return False
        m = len(s1)
        n = len(s2)
        
        if m + n != len(s3):
            return False
        
        dp = [[False] * (n+1) for _ in range(m+1)]
        dp[0][0] = True
        
        for i in range(1, m+1):
            dp[i][0] = ( s1[i-1] == s3[i-1] and dp[i-1][0] )
            
        for j in range(1, n+1):
            dp[0][j] = ( s2[j-1] == s3[j-1] and dp[0][j-1] )        
        
        for i in range(1, m+1):
            for j in range(1, n+1):                
                dp[i][j] = (s1[i-1] == s3[i+j-1] and dp[i-1][j]) or (s2[j-1] == s3[i+j-1] and dp[i][j-1])

        return dp[-1][-1]
