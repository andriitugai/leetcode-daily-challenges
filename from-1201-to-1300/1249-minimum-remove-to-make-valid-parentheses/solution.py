class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        sl = list(s)
        open = 0
        for i in range(len(sl)):
            if sl[i] == '(':
                open += 1
            elif sl[i] == ')':
                if open == 0:
                    sl[i] = '*'
                else:
                    open -= 1

        i = len(sl) - 1
        while i >= 0 and open > 0:
            if sl[i] == '(' and open > 0:
                sl[i] = '*'
                open -= 1
            i -= 1

        return ''.join([c for c in sl if c != '*'])