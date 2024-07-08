class Solution:
    def findTheWinner(self, n: int, k: int) -> int:
        q = deque([i for i in range(1, n+1)])
        cnt = 0
        while q:
            player = q.popleft()
            cnt += 1
            if cnt != k:
                q.append(player)
            else:
                cnt = 0

        return player