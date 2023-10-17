class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        coins = [coin for coin in sorted(candidates) if coin <= target]
        result = []
        
        def dfs(tgt, idx, path):
            nonlocal result, coins
            if tgt == 0:
                result.append(path)
                return

            for i in range(idx, len(coins)):
                if i > idx and coins[i] == coins[i-1]:
                    continue
                if coins[i] > tgt:
                    break
                dfs(tgt-coins[i], i+1, path + [coins[i]])
                
        dfs(target, 0, [])
        return result