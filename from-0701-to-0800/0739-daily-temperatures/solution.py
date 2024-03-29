class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        ans, s = [0]*len(temperatures), deque()
        for cur, cur_tmp in enumerate(temperatures):
            while s and cur_tmp > temperatures[s[-1]]:
                ans[s[-1]] = cur - s[-1]
                s.pop()
            s.append(cur)
        return ans