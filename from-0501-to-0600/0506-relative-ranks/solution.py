class Solution:
    def findRelativeRanks(self, score: List[int]) -> List[str]:
        d_scores = {s: i for i, s in enumerate(score)}
        ans = [""] * len(score)
        for p, s in enumerate(sorted(score, reverse=True), start=1):
            place = str(p)
            if p == 1:
                place = "Gold Medal"
            elif p == 2:
                place = "Silver Medal"
            elif p == 3:
                place = "Bronze Medal"
                
            ans[d_scores[s]] = place
            
        return ans