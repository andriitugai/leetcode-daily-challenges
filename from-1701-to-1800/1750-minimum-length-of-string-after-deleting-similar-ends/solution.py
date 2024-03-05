class Solution:
    def minimumLength(self, s: str) -> int:
        if len(s) == 1:
            return 1
        
        left = 0
        right = len(s) - 1
        
        while left < right and s[left] == s[right]:
            while s[left] == s[left+1] and left+1 < right:
                left += 1
            left += 1
            while s[right] == s[right-1] and right - 1 > left:
                right -= 1
            right -= 1
                
        return right - left + 1
        