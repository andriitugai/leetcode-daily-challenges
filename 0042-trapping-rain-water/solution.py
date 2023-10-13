class Solution:
    def trap(self, height: List[int]) -> int:
        if not height:
            return 0
        ans = 0
        size = len(height)
        left_max = [0] * size
        right_max = [0] * size
        
        left_max[0] = height[0]
        for i in range(1, size):
            left_max[i] = max(left_max[i-1], height[i])
        
        right_max[-1] = height[-1]
        for i in range(size-2, -1, -1):
            right_max[i] = max(right_max[i+1], height[i])
            
        # print(left_max)
        # print(right_max)
            
        for i in range(1, size-1):
            ans += (min(left_max[i], right_max[i]) - height[i])
            
        return ans