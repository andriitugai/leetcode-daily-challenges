class Solution:
    def arrayRankTransform(self, arr: List[int]) -> List[int]:
        if not arr:
            return []

        arrCopy = sorted(arr)

        ranks = dict()
        curNum = arrCopy[0]
        curRnk = 1
        ranks[curNum] = curRnk
        for num in arrCopy[1:]:
            if num != curNum:
                curNum = num
                curRnk += 1
            ranks[curNum] = curRnk

        return [ranks[num] for num in arr]
