class Solution:
    def openLock(self, deadends: List[str], target: str) -> int:
        if "0000" in deadends:
            return -1

        def neighbors(lock):
            result = []
            nums = list(map(int, lock))
            for i in range(4):
                nei = nums[:]
                nei[i] = (nums[i] + 1) % 10
                cand = "".join(list(map(str, nei)))
                if cand not in visited:
                    result.append(cand)
                nei[i] = (nums[i] + 9) % 10
                cand = "".join(list(map(str, nei)))
                if cand not in visited:
                    result.append(cand)
            return result

        visited = set(deadends)
        q = deque()
        q.append("0000")
        level = 0

        while q:
            for _ in range(len(q)):
                lock = q.popleft()
                if lock == target:
                    return level

                for nei in neighbors(lock):
                    visited.add(nei)
                    q.append(nei)
            level += 1

        return -1