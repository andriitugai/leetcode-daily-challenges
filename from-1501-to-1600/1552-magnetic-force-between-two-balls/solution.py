class Solution:
    def maxDistance(self, position: List[int], m: int) -> int:
        n = len(position)
        position.sort()

        def countballs(dist: int) -> int:
            num_balls = 1
            curr = position[0]
            for pos in position[1:]:
                if pos - curr >= dist:
                    num_balls += 1
                    curr = pos
            return num_balls

        left, right = 1, position[-1] - position[0]
        while left <= right:
            mid = (left + right) // 2
            balls = countballs(mid)
            if balls >= m:
                left = mid + 1
            else:
                right = mid - 1

        return right 