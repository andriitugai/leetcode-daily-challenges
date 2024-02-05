class Solution:
    def firstUniqChar(self, s: str) -> int:
        
        pos = defaultdict(int)
        for c in s:
            pos[c] += 1
            
        for idx, c in enumerate(s):
            if pos[c] == 1:
                return idx
            
        return -1