class Solution:
    def lenLongestFibSubseq(self, arr: List[int]) -> int:
        n = len(arr)
        dp = [[0] * n for _ in range(n)]
        max_len = 0
        for curr in range(2, n):
            start, end = 0, curr - 1
            final_sum = arr[curr]
            while start < end:
                curr_sum = arr[start] + arr[end]
                if curr_sum > final_sum:
                    end -= 1
                elif curr_sum < final_sum:
                    start += 1
                else:
                    dp[end][curr] = dp[start][end] + 1
                    max_len = max(max_len, dp[end][curr])
                    end -= 1
                    start += 1

        if max_len == 0:
            return 0
        return max_len + 2