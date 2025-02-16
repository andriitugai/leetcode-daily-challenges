class Solution:
    def hasSpecialSubstring(self, s: str, k: int) -> bool:
        if k == 1 and (len(s) == 1 or s[0] != s[1]):
            return True
        prev = s[0]
        same = 1
        for curr in s[1:]:
            if curr == prev:
                same += 1
            else:
                if same == k:
                    return True
                same = 1
            prev = curr
        if same == k:
            return True
        return False