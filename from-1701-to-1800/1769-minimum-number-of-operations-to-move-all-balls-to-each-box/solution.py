class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        n = len(boxes)
        result = [0] * n

        left_moves = left_balls = 0
        for i in range(n):
            result[i] = left_moves
            if boxes[i] == '1':
                left_balls += 1
            left_moves += left_balls

        right_moves = right_balls = 0
        for i in range(n -1, -1, -1):
            result[i] += right_moves
            if boxes[i] == '1':
                right_balls += 1
            right_moves += right_balls

        return result

        return result