class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        peps = sorted(people)
        ans = 0
        ph = 0
        pt = len(peps)-1
        
        while True:
            if ph == pt:
                return ans + 1
            if ph > pt:
                return ans
            group_weight = peps[pt]
            if peps[pt] + peps[pt-1] <= limit:
                ans += 1
                pt -= 2
            elif peps[pt] + peps[ph] <= limit:
                ans += 1
                pt -= 1
                ph += 1
            else:
                ans += 1
                pt -= 1