class Solution:
    def maxLength(self, nums: List[int]) -> int:
        n, result = len(nums), 1
        for i in range(n):
            for j in range(i, n):
                sub = nums[i:j + 1]
                gcd = reduce(math.gcd, sub)
                lcm = reduce(math.lcm, sub)
                prod_ = math.prod(sub)
                if prod_ == gcd * lcm:
                    result = max(result, len(sub))

        return result