class Solution:
    def countPrefixSuffixPairs(self, words: List[str]) -> int:
        def is_prefix_and_suffix(s1: str, s2: str) -> bool:
            n1, n2 = len(s1), len(s2)
            if n1 <= n2 and s1 == s2[:n1] and s1 == s2[n2 - n1:]:
                return True
            return False
        
        result = 0
        for i, word1 in enumerate(words):
            for word2 in words[i + 1:]:
                if is_prefix_and_suffix(word1, word2):
                    result += 1
        return result