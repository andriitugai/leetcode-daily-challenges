class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        doubled = None
        missed = None
                
        marks = [0] * len(nums)
        for num in nums:
            marks[num - 1] += 1
            if marks[num - 1] == 2:
                doubled = num
            
        for i, mark in enumerate(marks, 1):
            if mark == 0:
                missed = i
                break
                       
        return [doubled, missed]
        