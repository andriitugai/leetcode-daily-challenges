class Solution:
    def possibleStringCount(self, word: str) -> int:
        result = 1
        for i, c in enumerate(word):
            if i > 0 and c == word[i-1]:
                result += 1
        return result
        