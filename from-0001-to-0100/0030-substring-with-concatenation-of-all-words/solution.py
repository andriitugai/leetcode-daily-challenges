class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        wordsmap = Counter(words)
            
        wordslen = len(words[0])
        wordsnum = len(words)
        
        result = []
        
        for idx in range(len(s) - wordslen*wordsnum + 1):
            seen = []
            found = 0
            for j in range(wordsnum):
                word = s[idx + j*wordslen:idx + (j+1)*wordslen]
                if word in wordsmap:
                    seen.append(word)
                    found += 1
                else:
                    break
                    
            if found == wordsnum and wordsmap == Counter(seen):
                result.append(idx)
            
        return result