class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        n1, n2 = len(s1), len(s2)
        if n1 > n2:
            return False

        cnt1 = Counter(s1)
        cnt2 = Counter(s2[:n1])

        if cnt1 == cnt2:
            return True

        for i in range(n1, n2):
            cnt2[s2[i]] += 1
            cnt2[s2[i - n1]] -= 1
            if cnt2[s2[i - n1]] == 0:
                del cnt2[s2[i - n1]]

            if cnt2 == cnt1:
                return True

        return False