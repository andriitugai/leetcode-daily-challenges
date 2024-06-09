class Solution:
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        remainders = [0] * k
        remainders[0] = 1
        psum = 0
        result = 0
        for num in nums:
            psum += num
            rem = psum % k
            result += remainders[rem]
            remainders[rem] += 1

        return result