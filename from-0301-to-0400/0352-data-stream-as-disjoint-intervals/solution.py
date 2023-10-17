class SummaryRanges:

    def __init__(self):
        self.numset = set()
        
    def addNum(self, value: int) -> None:
        self.numset.add(value)

    def getIntervals(self) -> List[List[int]]:
        nums = sorted(list(self.numset))
        result = []
        i = 0
        while i < len(nums):
            start = nums[i]
            while i < len(nums)-1 and nums[i+1] == nums[i] + 1:
                i += 1

            result.append([start, nums[i]])
            i += 1

        return result


# Your SummaryRanges object will be instantiated and called as such:
# obj = SummaryRanges()
# obj.addNum(value)
# param_2 = obj.getIntervals()