class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        chmap_st = {}
        chmap_ts = {}
        
        for cs, ct in zip(s, t):
            if cs in chmap_st:
                if chmap_st[cs] != ct:
                    return False
            else:
                chmap_st[cs] = ct
                
            if ct in chmap_ts:
                if chmap_ts[ct] != cs:
                    return False
            else:
                chmap_ts[ct] = cs
                
        return True