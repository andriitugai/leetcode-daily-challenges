class Solution:
    def findWinners(self, matches: List[List[int]]) -> List[List[int]]:
        loses = dict()
        for winner, loser in matches:
            if winner not in loses:
                loses[winner] = 0
            loses[loser] = loses.get(loser, 0) + 1

        no_lose = sorted([w for w, l in loses.items() if l == 0])
        one_lose = sorted([w for w, l in loses.items() if l == 1])

        return [no_lose, one_lose]