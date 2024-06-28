class Solution:
    def maximumImportance(self, n: int, roads: List[List[int]]) -> int:
        edgeCount = [0] * n
        for c1, c2 in roads:
            edgeCount[c1] += 1
            edgeCount[c2] += 1

        result = 0
        label = 1
        for count in sorted(edgeCount):
            result += label * count
            label += 1

        return result