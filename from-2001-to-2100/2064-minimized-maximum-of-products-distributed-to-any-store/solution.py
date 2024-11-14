class Solution:
    def minimizedMaximum(self, n: int, quantities: List[int]) -> int:
        
        left, right = 1, max(quantities)
        while left < right:
            mid = (left + right) // 2
            if sum(math.ceil(q / mid) for q in quantities) <= n:
                right = mid
            else:
                left = mid + 1
                
        return left
    
    
    def minimizedMaximum_(self, n, Q):
        beg, end = 0, max(Q)

        while beg + 1 < end:
            mid = (beg + end)//2
            if sum(ceil(i/mid) for i in Q) <= n:
                end = mid
            else:
                beg = mid

        return end