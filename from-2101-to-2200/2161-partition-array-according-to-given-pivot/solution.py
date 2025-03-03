class Solution:
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:
        lt, eq, gt = [], [], []
        for num in nums:
            if num < pivot:
                lt.append(num)
            elif num > pivot:
                gt.append(num)
            else:
                eq.append(num)
                
        return lt + eq + gt
        