lass Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        digits_plus = [0] * len(digits)
        
        extra = 1
        for i in range(len(digits)-1, -1, -1):
            d = digits[i]
            digits_plus[i] = (d + extra) % 10
            extra = (d + extra) // 10
            
        if extra == 1:
            digits_plus = [1] + digits_plus
            
        return digits_plus
        
        