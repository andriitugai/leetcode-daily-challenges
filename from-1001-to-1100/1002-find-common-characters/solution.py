class Solution:
    def commonChars(self, words: List[str]) -> List[str]:
        counters = [Counter(list(word)) for word in words]
        
        ans = []
        for k in counters[0].keys():
            if all([k in c for c in counters[1:]]):
                ans.extend([k] * min([c[k] for c in counters]))
                
        return ans
    