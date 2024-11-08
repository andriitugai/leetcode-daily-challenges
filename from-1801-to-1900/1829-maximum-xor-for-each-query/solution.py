class Solution:
    def getMaximumXor(self, nums: List[int], maximumBit: int) -> List[int]:
        xor = 0
        for n in nums:
            xor ^= n

        bitmask = (1 << maximumBit) - 1
        result = []
        for n in reversed(nums):
            result.append(xor ^ bitmask)
            xor ^= n

        return result