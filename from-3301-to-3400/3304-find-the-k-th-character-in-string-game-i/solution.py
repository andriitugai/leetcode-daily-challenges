class Solution:
    def kthCharacter(self, k: int) -> str:
        s = "a"
        while len(s) < k:
            buff = ""
            for c in s:
                buff += chr(ord(c) + 1)
            s += buff
        return s[k - 1]