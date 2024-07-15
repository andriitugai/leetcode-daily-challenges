class Solution:
    def getSmallestString(self, s: str) -> str:
        lst = [int(a) for a in list(s)]
        for i in range(len(s) - 1):
            if lst[i] % 2 == lst[i + 1] % 2 and lst[i] > lst[i + 1]:
                lst[i], lst[i + 1] = lst[i + 1], lst[i]
                return "".join([str(a) for a in lst])
        return s