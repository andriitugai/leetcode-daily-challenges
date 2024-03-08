class Solution:
    def maxFrequencyElements(self, nums: List[int]) -> int:
        maxFreq = 0
        cnt = {}
        for num in nums:
            cnt[num] = cnt.get(num, 0) + 1
            if cnt[num] > maxFreq:
                maxFreq = cnt[num]

        result = 0
        for _, freq in cnt.items():
            if freq == maxFreq:
                result += freq

        return result