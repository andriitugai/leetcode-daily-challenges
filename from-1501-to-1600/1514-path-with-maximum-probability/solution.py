class Solution:
    def maxProbability(self, n: int, edges: List[List[int]], succProb: List[float], start: int, end: int) -> float:
        graph = defaultdict(list)
        for (a, b), prob in zip(edges, succProb):
            graph[a].append((b, prob))
            graph[b].append((a, prob))


        probs = [0.0] * n
        probs[start] = 1.0

        q = deque()
        q.append(start)

        while q:
            node = q.popleft()
            for nbr, succ in graph[node]:
                if probs[node] * succ > probs[nbr]:
                    probs[nbr] = probs[node] * succ
                    q.append(nbr)

        return probs[end]