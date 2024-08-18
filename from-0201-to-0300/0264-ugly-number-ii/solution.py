class Solution:
    def nthUglyNumber(self, n: int) -> int:
        uglies = [1]
        f2, f3, f5 = 0, 0, 0
        curr =  1
        while curr < n:
            new_ugly = min(uglies[f2]*2, uglies[f3]*3, uglies[f5]*5)
            if new_ugly == uglies[f2]*2:
                f2 += 1
            if new_ugly == uglies[f3]*3:
                f3 += 1
            if new_ugly == uglies[f5]*5:
                f5 += 1

            uglies.append(new_ugly)
            curr += 1

        return uglies[-1]