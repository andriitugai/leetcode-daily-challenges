from collections import Counter

class Solution:
    from collections import Counter
    
    def isAnagram(self, s: str, t: str) -> bool:
        sc = Counter(s)
        tc = Counter(t)
        
        return sc == tc