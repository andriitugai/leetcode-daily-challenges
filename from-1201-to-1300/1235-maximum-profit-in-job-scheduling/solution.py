class Solution:
    def jobScheduling_(self, startTime: List[int], endTime: List[int], profit: List[int]) -> int:
        jobs = [(s, e, p) for (s, e), p in zip(zip(startTime, endTime), profit)]
        jobs.sort(key=lambda x: (x[1], x[0]))

        n = len(jobs)
        dp = profit.copy()

        for i in range(1, n):
            s0, e0, p0 = jobs[i]
            dp[i] = max(p0, dp[i-1])
            for j in range(i):
                s1, e1, p1 = jobs[j]
                # if not overlap:
                if s0 >= e1:
                    dp[i] = max(dp[i], dp[j] + p0)

        return max(dp)

    def jobScheduling(self, S, E, profit):
        jobs = sorted(list(zip(S, E, profit)))
        S = [i[0] for i in jobs]
        n = len(jobs)
        dp = [0] * (n + 1)
        for k in range(n-1, -1, -1):
            temp = bisect_left(S, jobs[k][1])
            dp[k] = max(jobs[k][2] + dp[temp], dp[k+1])
        return dp[0]