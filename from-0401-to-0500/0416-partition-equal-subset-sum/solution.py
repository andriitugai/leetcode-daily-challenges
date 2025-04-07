class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        total = sum(nums)
        if total % 2 == 1:
            return False

        goal = total // 2
        sums = set([0])

        for num in nums:
            for curr_sum in list(sums):
                new_sum = curr_sum + num
                if new_sum == goal:
                    return True

                sums.add(new_sum)
        
        return False