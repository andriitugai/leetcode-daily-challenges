class Solution:
    def bagOfTokensScore(self, tokens: List[int], power: int) -> int:
        q = deque(sorted(tokens))
        score, max_score = 0, 0
        
        while q:
            if power >= q[0]:
                power -= q.popleft()
                score += 1
                max_score = max(max_score, score)
            elif score > 0:
                power += q.pop()
                score -= 1
            else:
                break
                
        return max_score
        