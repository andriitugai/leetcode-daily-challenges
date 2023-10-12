class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        candidates.sort()
        
        def dfs(tgt, idx, path):
            nonlocal result, candidates
            if tgt == 0:
                result.append(path)
                return
            
            for i in range(idx, len(candidates)):
                if candidates[i] > tgt:
                    break
                dfs(tgt - candidates[i], i, path + [candidates[i]])
                    
        dfs(target, 0, [])
        return result
            