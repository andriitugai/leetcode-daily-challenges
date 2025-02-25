class Solution:
    def numOfSubarrays(self, arr: List[int]) -> int:
        MOD = 1000000007
        oddity = [0] * len(arr)
        curr = 0
        for i, num in enumerate(arr):
            curr = (curr + num) % 2
            oddity[i] = curr

        odd, even = 0, 0
        result = 0

        for v in oddity:
            if v == 0:
                even += 1
                result += odd
            else:
                odd += 1
                result += even + 1
            
            result %= MOD

        return result