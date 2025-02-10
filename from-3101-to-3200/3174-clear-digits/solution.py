class Solution:
    def clearDigits(self, s: str) -> str:
        n = len(s)
        marks = [0] * n
        for i, c in enumerate(s):
            if 'a' <= c <= 'z':
                marks[i] = 1
            else:
                for j in range(i - 1, -1, -1):
                    if marks[j] == 1:
                        marks[j] = 0
                        break

        return ''.join(c for c, mark in zip(s, marks) if mark == 1)