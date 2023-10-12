class Solution:
    def romanToInt(self, s: str) -> int:
        romans = {
            'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000
        }
        prev = 0
        result = 0
        i = len(s) - 1
        while i >= 0:
            curr = romans[s[i]]
            if curr >= prev:
                result += curr
                prev = curr
            else:
                result -= curr
                
            i -= 1
            
        return result  