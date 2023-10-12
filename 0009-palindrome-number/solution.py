lass Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        if x < 10:
            return True
        digs = []
        while x >= 10:
            digs.append(x%10)
            x = x // 10
        digs.append(x)
        
        return digs == digs[::-1]
        