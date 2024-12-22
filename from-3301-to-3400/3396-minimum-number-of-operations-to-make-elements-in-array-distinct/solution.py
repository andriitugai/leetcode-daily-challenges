class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        pMap = defaultdict(list)
        for i, num in enumerate(nums):
            pMap[num].append(i)

        posToRemove = -1
        for positions in pMap.values():
            if len(positions) > 1:
                posToRemove = max(posToRemove, positions[-2])

        return 0 if posToRemove == -1 else posToRemove // 3 + 1
