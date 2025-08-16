class Solution:
    def maximum69Number (self, num: int) -> int:
        nstr = list(str(num))
        for i in range(len(nstr)):
            if nstr[i] == '6':
                nstr[i] = '9'
                break
    
        return int(''.join(nstr))