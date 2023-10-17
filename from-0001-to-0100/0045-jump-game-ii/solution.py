class Solution:
    def jump(self, nums: List[int]) -> int:
        mri = 0  # max reachable index
        steps = 0
        left, right = 0, 0
        while right < len(nums) - 1:
            for pos in range(left, right + 1):
                mri = max(mri, pos + nums[pos])
            left = right + 1
            right = mri
            steps += 1

        return steps