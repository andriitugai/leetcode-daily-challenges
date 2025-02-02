class Solution:
    def findValidPair(self, s: str) -> str:
        freq = Counter(s)
        prev_c = s[0]
        prev_n = int(prev_c)
        for i in range(1, len(s)):
            curr_c = s[i]
            curr_n = int(curr_c)
            if curr_c != prev_c and freq[prev_c] == prev_n and freq[curr_c] == curr_n:
                return s[i-1:i+1]
            prev_c, prev_n = curr_c, curr_n
        return ""