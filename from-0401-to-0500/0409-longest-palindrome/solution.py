class Solution:
    def longestPalindrome(self, s: str) -> int:
        cs = Counter(s)
        
        max_odd = 0
        ans = 0
        for c in cs:
            cnum = cs[c]
            if cnum % 2 == 0:
                ans += cnum
            else:
                ans += (cnum - 1)
                if cnum > max_odd:
                    max_odd = 1
                    
        ans += max_odd
        return ans