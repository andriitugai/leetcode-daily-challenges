class Solution:
    def vowelStrings(self, words: List[str], queries: List[List[int]]) -> List[int]:
        vowels = {'a', 'e', 'i', 'o', 'u'}
        n = len(words)
        prefix = [0] * (n + 1)

        for i in range(n):
            prefix[i + 1] = prefix[i]
            if words[i][0] in vowels and words[i][-1] in vowels:
                prefix[i + 1] += 1

        result = [0] * len(queries)
        for i, (left, right) in enumerate(queries):
            result[i] = prefix[right + 1] - prefix[left]

        return result
