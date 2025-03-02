class Solution:
    def largestInteger(self, nums: List[int], k: int) -> int:
        counts = Counter(nums)
        if k == 1:
            uniques = [num for num, count in counts.items() if count == 1]
            if uniques:
                return max(uniques)
            return -1
        elif k == len(nums):
            return max(nums)
        else:
            first, last = nums[0], nums[-1]
            c_first,  c_last = counts[first], counts[last]
            if c_first == 1 and c_last == 1:
                return max(first, last)
            elif c_first == 1:
                return first
            elif c_last == 1:
                return last
        return -1
        