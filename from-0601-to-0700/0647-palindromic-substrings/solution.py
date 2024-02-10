class Solution:
    def countSubstrings(self, s: str) -> int:
        slen = len(s)
        if slen < 2:
            return slen
        
        count = 0
        dp = [ [0] * slen for _ in range(slen)]
        for i in range(slen):
            dp[i][i] = 1
            count += 1
            
        for col in range(1, slen):
            for row in range(col):
                if s[row] == s[col]:                # the first and the last chars are equal             
                    if row == col - 1:              # the len of substr == 2      
                        dp[row][col] = 1
                        count += 1
                    elif dp[row+1][col-1] == 1:     # the inner substr is a palindrome
                        dp[row][col] = 1
                        count += 1
                    
        return count
                    
            