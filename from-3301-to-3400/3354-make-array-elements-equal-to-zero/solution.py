class Solution:
    def countValidSelections(self, nums: List[int]) -> int:
        n = len(nums)
        
        def simulate(arr: List[int], curr: int, dir: int) -> bool:
            nonlocal n
            while 0 <= curr < n:
                if arr[curr] > 0:
                    arr[curr] -= 1
                    dir *= -1
                curr += dir
            return not any(arr)

        result = 0
        for i in range(n):
            if nums[i] == 0:
                if simulate(nums[:], i, 1):
                    result += 1
                if simulate(nums[:], i, -1):
                    result += 1

        return result