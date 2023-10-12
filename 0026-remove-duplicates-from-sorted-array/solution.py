class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        # k, i = 0, 1
        # while True:
        #     while i < len(nums) and nums[i] == nums[k]:
        #         i += 1
        #     if i == len(nums):
        #         break
        #     nums[k+1] = nums[i]
        #     k += 1

        # return k + 1

        slow, fast = 0, 1

        while fast < len(nums):
            if nums[fast] != nums[slow]:
                slow += 1
                if slow != fast:
                    nums[slow] = nums[fast]
            fast += 1

        return slow + 1