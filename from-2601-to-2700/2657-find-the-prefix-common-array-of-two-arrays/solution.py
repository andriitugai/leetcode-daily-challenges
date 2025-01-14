class Solution:
    def findThePrefixCommonArray(self, A: List[int], B: List[int]) -> List[int]:
        counts = defaultdict(int)
        result = [0] * len(A)
        curr = 0

        for i, (a, b) in enumerate(zip(A, B)):
            counts[a] += 1
            if counts[a] == 2:
                curr += 1

            counts[b] += 1
            if counts[b] == 2:
                curr += 1
            
            result[i] = curr

        return result