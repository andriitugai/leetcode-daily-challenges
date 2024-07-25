class Solution:
    def minimumTime(self, n: int, edges: List[List[int]], disappear: List[int]) -> List[int]:
        graph = defaultdict(list)
        heap = [(0, 0)]
        result = [-1] * n

        for u, v, t in edges:
            graph[u].append((t, v))
            graph[v].append((t, u))

        while heap:
            t1, node1 = heappop(heap)
            if result[node1] != -1:
                continue
            result[node1] = t1
            for t2, node2 in graph[node1]:
                t2 += t1
                if t2 >= disappear[node2]:
                    continue
                heappush(heap, (t2, node2))

        return result