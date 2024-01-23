class Solution:
    def maxLength(self, arr: List[str]) -> int:
        char_set = set()
        
        def hasDuplicates(cs: set, s: string)-> bool:
            pre_set = set()
            for c in s:
                if c in pre_set or c in cs:
                    return True
                pre_set.add(c)
            return False
        
        def backtrack(i: int):
            if i == len(arr):
                return len(char_set)
            
            ans = 0
            if not hasDuplicates(char_set, arr[i]):
                for c in arr[i]:
                    char_set.add(c)
                ans = backtrack(i+1)
                for c in arr[i]:
                    char_set.remove(c)
                    
            return max(ans, backtrack(i+1))
        
        return backtrack(0)