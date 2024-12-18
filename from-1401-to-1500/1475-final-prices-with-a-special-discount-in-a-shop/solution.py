class Solution:
    def finalPrices(self, prices: List[int]) -> List[int]:
        result = prices[::]
        for i, p in enumerate(prices):
            for j in range(i+1, len(prices)):
                if prices[j] <= p:
                    result[i] = p - prices[j]
                    break
        return result