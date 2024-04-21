class Solution:
    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        if n == 1:
            return True
            
        g = defaultdict(set)
        for e in edges:
            g[e[0]].add(e[1])
            g[e[1]].add(e[0])

        visited = set()
        q = deque()
        q.append(source)

        while q:
            node = q.popleft()
            visited.add(node)
            for nbr in g[node]:
                if nbr == destination:
                    return True
                if nbr not in visited:
                    q.append(nbr)

        return False
