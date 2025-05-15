class Solution:
    def getLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        idxs = [0]
        for i in range(1, len(groups)):
            if groups[i] != groups[i - 1]:
                idxs.append(i)

        return [words[i] for i in idxs]