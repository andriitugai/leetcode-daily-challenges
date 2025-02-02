class Solution:
    def maxDifference(self, s: str) -> int:
        freq = [0] * 26
        for c in s:
            freq[ord(c) - ord('a')] += 1

        max_odd = 0
        min_even = len(s) + 1
        for f in freq:
            if f > 0:
                if f % 2 and f > max_odd:
                    max_odd = f
                elif f % 2 == 0 and f < min_even:
                    min_even = f

        print(max_odd, min_even)
        return max_odd - min_even
