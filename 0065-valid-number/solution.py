mport re

class Solution:
    def isNumber(self, s: str) -> bool:
        # m = re.match('^[+\-]?[0-9\.eE]+$', s)
        m = re.match('^[+\-]?(([0-9]+[\.]?[0-9]*)|([\.][0-9]+))([eE]{1}[+\-]?[0-9]+)?$', s)
        if m:
            return True
        
        return False
        # return bool(m)