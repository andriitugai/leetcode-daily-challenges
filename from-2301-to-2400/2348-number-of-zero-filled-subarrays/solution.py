class Solution:
    def zeroFilledSubarray(self, nums: List[int]) -> int:
        result = 0
        cont = 0
        for n in nums:
            if n == 0:
                cont += 1
                result += cont
            else:
                cont = 0
        return result