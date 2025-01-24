class Solution:
    def eventualSafeNodes(self, graph: List[List[int]]) -> List[int]:
        result = []
        safe = dict()

        def dfs(node: int) -> bool:
            if node in safe:
                return safe[node]

            safe[node] = False
            for nei in graph[node]:
                if not dfs(nei):
                    return False
            safe[node] = True
            return True

        for i in range(len(graph)):
            if dfs(i):
                result.append(i)

        return result
