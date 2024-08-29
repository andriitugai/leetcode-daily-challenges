class Solution:
    def removeStones(self, stones: List[List[int]]) -> int:
        rows = defaultdict(list)
        cols = defaultdict(list)

        for row, col in stones:
            rows[row].append((row, col))
            cols[col].append((row, col))

        seen = set()
        def dfs(r, c) -> int:
            if (r, c) in seen:
                return 0

            seen.add((r, c))
            for row, col in rows[r] + cols[c]:
                dfs(row, col)

            return 1

        # Count the number of islands
        count = sum([dfs(r, c) for r, c in stones])
        return len(stones) - count