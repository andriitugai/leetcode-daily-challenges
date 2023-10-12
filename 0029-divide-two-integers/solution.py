class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == -(1 << 31) and divisor == -1:
            return (1 << 31) - 1
        
        overflown = 2**31-1
        sign = -1
        
        if dividend < 0 and divisor < 0 or dividend > 0 and divisor > 0:
            sign = 1
            
        dividend = abs(dividend)
        divisor = abs(divisor)
        
        # if divisor == 1:
        #     if dividend > overflown:
        #         return overflown
        #     return sign * dividend
        
        ratio = 0
        while dividend >= divisor:
            k = 0
            while dividend - (divisor << 1 << k) >= 0:
                k += 1
            ratio += (1 << k)
            
            dividend -= divisor << k
                        
        return -ratio if sign < 0 else ratio
    