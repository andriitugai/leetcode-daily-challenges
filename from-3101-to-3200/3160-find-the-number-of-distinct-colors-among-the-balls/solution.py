class Solution:
    def queryResults(self, limit: int, queries: List[List[int]]) -> List[int]:
        ball_map = {}
        colors = defaultdict(int)
        result = [0] * len(queries)

        for i, query in enumerate(queries):
            ball, color = query
            if ball in ball_map:
                balance = ball_map[ball]
                if colors[balance] == 1:
                    del colors[balance]
                else:
                    colors[balance] -= 1
            
            ball_map[ball] = color
            colors[color] += 1
            result[i] = len(colors)

        return result
                