class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        
        cand = []
        result = []
        
        def backtrack(left, right):
            if left == n and right == n:
                result.append("".join(cand))
                return
            
            if left < n:
                cand.append("(")
                backtrack(left+1, right)
                cand.pop()
                
            if right < left:
                cand.append(")")
                backtrack(left, right+1)
                cand.pop()
                
        backtrack(0, 0)
        
        return result
        