class Solution:
    def isBalanced(self, num: str) -> bool:
        evenSum, oddSum = 0, 0
        for i, d in enumerate(num):
            if i % 2:
                oddSum += int(d)
            else:
                evenSum += int(d)
        return evenSum == oddSum