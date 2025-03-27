lass Solution:
    def minimumIndex(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        dom, dom_freq = float('-inf'), 0
        for x, f in cnt.items():
            if f > dom_freq:
                dom_freq = f
                dom = x

        n = len(nums)
        left_count = 1 if nums[0] == dom else 0
        right_count = dom_freq - left_count
        for i in range(1, n):
            if left_count * 2 > i and right_count * 2 > n - i:
                return i - 1
            if nums[i] == dom:
                left_count += 1
                right_count -= 1

        return -1
