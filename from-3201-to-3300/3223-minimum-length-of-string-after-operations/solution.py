class Solution:
    def minimumLength(self, s: str) -> int:
        counts = Counter(s)
        to_delete = 0
        for _, v in counts.items():
            to_delete += v - 1 if v % 2 == 1 else v - 2

        return len(s) - to_delete