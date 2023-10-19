class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:
        sr = []
        for c in s:
            if c == '#':
                if sr:
                    sr.pop()
            else:
                sr.append(c)
                
        tr = []
        for c in t:
            if c == '#':
                if tr:
                    tr.pop()
            else:
                tr.append(c)
              
        print(sr, tr)
        return sr == tr
    