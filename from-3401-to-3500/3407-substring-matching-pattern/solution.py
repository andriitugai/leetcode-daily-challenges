class Solution:
    def hasMatch(self, s: str, p: str) -> bool:
        p1, p2 = p.split('*')
        
        i1 = s.find(p1)
        if i1 < 0:
            return False

        return p2 in s[i1 + len(p1):]
        