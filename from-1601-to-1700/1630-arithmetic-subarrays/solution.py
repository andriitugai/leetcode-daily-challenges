class Solution:
    def checkArithmeticSubarrays(self, nums: List[int], l: List[int], r: List[int]) -> List[bool]:
        def is_ariphmetic(arr: List[int]) -> bool:
            n = len(arr)
            s = set(arr)
            maxa, mina = max(arr), min(arr)
            if (maxa-mina)%(n-1):
                return False
            step = (maxa-mina)//(n-1)
            return all(mina+step*k in s for k in range(n))

        return [is_ariphmetic(nums[i:j+1]) for i, j in zip(l, r)]