lass Graph:

    def __init__(self, n: int, edges: List[List[int]]):
        self.graph = defaultdict(list)
        self.n=n
        for u,v,c in edges:
            self.graph[u].append((v,c))


    def addEdge(self, edge: List[int]) -> None:
        self.graph[edge[0]].append((edge[1] , edge[2]))


    def shortestPath(self, node1: int, node2: int) -> int:
        dist = [float("inf")]*self.n
        dist[node1]=0
        heap = [[node1,0]]
        res=float("inf")
        while heap:
            node , cost = heappop(heap)
            if node ==node2:
                res = min(res , cost)
            if cost > dist[node]:
                continue
            for neigh , edge_cost in self.graph[node]:
                new_cost=edge_cost + cost
                if new_cost < dist[neigh]:
                    heappush(heap , (neigh , new_cost))
                    dist[neigh] = new_cost
        return res if dist[node2]!=float("inf") else -1


# Your Graph object will be instantiated and called as such:
# obj = Graph(n, edges)
# obj.addEdge(edge)
# param_2 = obj.shortestPath(node1,node2)