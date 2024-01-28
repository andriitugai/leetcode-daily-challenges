class Solution:
    def numSubmatrixSumTarget(self, matrix: List[List[int]], target: int) -> int:
        def getCountWithTarget(arr):
            sum_dict = {0:1}
            ans = 0
            running_sum = 0
            for num in arr:
                running_sum += num
                target_sum = running_sum - target
                if target_sum in sum_dict:
                    ans += sum_dict[target_sum]
                    
                sum_dict[running_sum] = sum_dict.get(running_sum, 0) + 1
                
            return ans
        
            
#         presum = [row[:] for row in matrix]
#         for row in range(len(matrix)):
#             for col in range(1, len(matrix[0])):
#                 presum[row][col] = presum[row][col-1] + matrix[row][col]
        m = len(matrix)
        n = len(matrix[0])
        presum = []
        for i in range(m):
            cur_sum = []
            for j in range(n):
                if j == 0:
                    cur_sum.append(matrix[i][j])
                else:
                    cur_sum.append(matrix[i][j] + cur_sum[-1])
                    
            presum.append(cur_sum)
        
        res = 0
        
        for c2 in range(n):
            for c1 in range(c2+1):
                curr_arr = []
                for row in range(m):
                    row_sum = presum[row][c2] - (presum[row][c1-1] if c1 > 0 else 0)
                    curr_arr.append(row_sum)
                    
                res += getCountWithTarget(curr_arr)
                
        return res