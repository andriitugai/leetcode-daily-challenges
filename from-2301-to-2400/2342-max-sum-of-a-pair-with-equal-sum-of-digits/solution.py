class Solution:
    def maximumSum(self, nums: List[int]) -> int:
        class MaxTwo:
            def __init__(self):
                self.el_1 = -1
                self.el_2 = -1

            def add(self, el: int) -> None:
                if el > self.el_1:
                    self.el_2 = self.el_1
                    self.el_1 = el
                elif el > self.el_2:
                    self.el_2 = el

            def getSum(self) -> int:
                if self.el_2 == -1:
                    return -1
                return self.el_1 + self.el_2

        digit_sums = defaultdict(MaxTwo)

        def get_digit_sum(num: int) -> int:
            s = 0
            while num > 0:
                s += num % 10
                num //= 10
            return s

        for num in nums:
            ds = get_digit_sum(num)
            digit_sums[ds].add(num)

        max_sum = -1
        for mt in digit_sums.values():
            mts = mt.getSum()
            if mts > max_sum:
                max_sum = mts

        return max_sum
