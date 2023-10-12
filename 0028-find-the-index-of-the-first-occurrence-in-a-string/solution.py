class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if not needle:
            return 0

        lneed = len(needle)

        for i in range(len(haystack)-len(needle)+1):
            if haystack[i:i+lneed] == needle:
                return i

        return -1