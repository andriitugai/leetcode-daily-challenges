class Solution:
    def frequencySort(self, nums: List[int]) -> List[int]:
        numf = sorted(collections.Counter(nums).items(), key=lambda x: (x[1], -x[0]))
        res = []
        for num in numf:
            res += ([num[0]] * num[1])
            
        return res
        