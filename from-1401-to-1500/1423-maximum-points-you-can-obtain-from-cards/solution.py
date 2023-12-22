class Solution:
    def maxScore(self, cardPoints: List[int], k: int) -> int:
        total = sum(cardPoints)
        n = len(cardPoints)
        if n == k:
            return total

        ans = 0
        cur_sum = sum(cardPoints[:n-k])
        ans = total - cur_sum
        for i in range(1, k+1):
            cur_sum = cur_sum - cardPoints[i-1] + cardPoints[i+n-k-1]
            ans = max(ans, total - cur_sum)

        return ans