class Solution:
    def mySqrt(self, x: int) -> int:
        k = len(bin(x)[2:])

        top = int('1' * (k // 2 + 1), 2)
        low = int('1' + '0' * (k // 2 - 1), 2)

        while top > low:
            root = (top + low) // 2
            if root * root == x:
                return root
            elif root * root < x:
                low = root if low != root else low + 1
            else:
                top = root if top != root else top - 1
                
        if low * low > x: 
            low -= 1
        
        return low