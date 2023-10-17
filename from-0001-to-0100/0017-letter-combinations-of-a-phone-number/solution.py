import itertools

class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        
        dtl = {
                "2": "abc", "3": "def", "4": "ghi", "5": "jkl", 
                "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
        
        ans = list(itertools.product(*[dtl[d] for d in digits]))
        
        return [''.join(lst) for lst in ans]