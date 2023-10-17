class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        mx = [ [0] * n for _ in range(n)] 
        rs, re = 0, n-1
        cs, ce = 0, n-1
        
        val = 1
        
        while rs <= re and cs <= ce:
            for col in range(cs, ce+1):
                mx[rs][col] = val
                val += 1
            rs += 1
            
            for row in range(rs, re+1):
                mx[row][ce] = val
                val += 1
            ce -= 1
            
            if cs <= ce:
                for col in range(ce, cs-1, -1):
                    mx[re][col] = val
                    val += 1
                re -= 1
                
            if rs <= re:
                for row in range(re, rs-1, -1):
                    mx[row][cs] = val
                    val += 1
                cs += 1
                
        return mx