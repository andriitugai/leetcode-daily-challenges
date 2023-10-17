from typing import List

class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        col_left, col_right = 0, len(matrix[0])-1
        row_up, row_dn = 0, len(matrix)-1
        ans = []
        
        while col_left <= col_right and row_up <= row_dn:
            col = col_left
            while col <= col_right:
                ans.append(matrix[row_up][col])
                col += 1
            row_up += 1
            
            row = row_up
            while row <= row_dn:
                ans.append(matrix[row][col_right])
                row += 1
            col_right -= 1
            
            if row_up <= row_dn:
                col = col_right
                while col >= col_left:
                    ans.append(matrix[row_dn][col])
                    col -= 1
                row_dn -= 1
            
            if col_left <= col_right:
                row = row_dn
                while row >= row_up:
                    ans.append(matrix[row][col_left])
                    row -= 1
                col_left += 1
            
        return ans