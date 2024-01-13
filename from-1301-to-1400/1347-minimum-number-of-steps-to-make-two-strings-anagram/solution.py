class Solution:
    def minSteps(self, s: str, t: str) -> int:
        cs, ct = Counter(s), Counter(t)
        
        res = 0
        for k in ct:
            if ct[k] > cs.get(k, 0):
                res += (ct[k] - cs.get(k, 0))
            
        return res
    