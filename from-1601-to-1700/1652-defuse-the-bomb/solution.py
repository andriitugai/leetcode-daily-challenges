class Solution:
    def decrypt(self, code: List[int], k: int) -> List[int]:
        n = len(code)
        result = [0] * n

        for i in range(n):
            if k > 0:
                for j in range(i + 1, i + 1 + k):
                    result[i] += code[j % n]
            elif k < 0:
                for j in range(i - 1, i - 1 + k, -1):
                    result[i] += code[j % n]

        return result