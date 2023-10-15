class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        dp = {}
        def get_distance(i, j):
            if (i, j) in dp:
                return dp[i,j]
            if i == len(word1):
                return len(word2) - j
            if j == len(word2):
                return len(word1) - i
            else:
                if word1[i] == word2[j]:
                    dp[i, j] = get_distance(i+1, j+1)
                else:
                    dp[i, j] = min(
                        get_distance(i+1, j),
                        get_distance(i, j+1),
                        get_distance(i+1, j+1)
                    ) + 1
                return dp[i, j]
            
        return get_distance(0, 0)