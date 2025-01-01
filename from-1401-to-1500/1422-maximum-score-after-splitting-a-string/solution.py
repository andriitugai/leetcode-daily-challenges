class Solution:
    def maxScore(self, s: str) -> int:
        nOnes = sum(1 if c == '1' else 0 for c in s)
        
        score, maxScore = nOnes, 0
        for c in s[:-1]:
            if c == '0':
                score += 1
            else:
                score -= 1
                
            maxScore = max(score, maxScore)

        return maxScore