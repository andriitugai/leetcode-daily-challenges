class Solution:
    def findChampion(self, n: int, edges: List[List[int]]) -> int:
        incoming = [0] * n
        for _, dst in edges:
            incoming[dst] += 1

        champion = -1
        for i in range(n):
            if incoming[i] == 0:
                if champion > -1:
                    return -1
                else:
                    champion = i

        return champion