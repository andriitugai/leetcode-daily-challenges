class Solution:
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        # from itertools import zip_longest
        # g1 = (c for w in word1 for c in w)
        # g2 = (c for w in word2 for c in w)
        # for c1, c2 in zip_longest(g1, g2):
        #     if c1 != c2:
        #         return False
        # return True
        return ''.join(word1) == ''.join(word2)