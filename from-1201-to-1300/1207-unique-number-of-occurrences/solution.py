class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        occu = Counter(arr).values()
        return len(occu) == len(set(occu))
        